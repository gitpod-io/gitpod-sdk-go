package cache

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"

	"connectrpc.com/connect"
	v1 "github.com/gitpod-io/flex-api-go/v1"
	"github.com/gitpod-io/flex-api-go/v1/v1connect"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/proto"
)

// Custom error types for better error handling
var (
	ErrCacheClosed  = errors.New("cache is closed")
	ErrItemNotFound = errors.New("item not found in cache or data source")
	ErrInvalidInput = errors.New("invalid input parameter")
	ErrSyncFailed   = errors.New("full sync operation failed")
)

// GetListIndex defines the interface for retrieving and listing items.
type GetListIndex[T any] interface {
	// Get retrieves an item from some (typically remote) source.
	Get(context.Context, string) (T, error)

	// List retrieves all items from some (typically remote) source.
	// If maxItems is 0, all items will be returned. If ids are provided,
	// only the items with the given IDs will be returned.
	List(ctx context.Context, maxItems int, ids []string) ([]T, error)

	// Index returns the index of a given item.
	Index(T) string
}

// Informer is an interface that provides methods for retrieving and listing items from a cache.
// The Get method retrieves the value for the given key from the cache. If the value is not
// found in the cache, it will be fetched from the underlying data source and added to the cache.
// If the item isn't found, expect a connect.NotFound error.
// The List method retrieves all items from some (typically remote) source.
// The Close method closes the informer and releases any associated resources.
type Informer[T proto.Message] interface {
	Get(ctx context.Context, key string) (T, error)
	List(ctx context.Context) ([]T, error)
	Close() error
}

type InformerMock[T proto.Message] map[string]T

func (it InformerMock[T]) Get(ctx context.Context, key string) (T, error) {
	if val, ok := it[key]; ok {
		return val, nil
	}
	var dt T
	return dt, ErrItemNotFound
}

func (it InformerMock[T]) List(ctx context.Context) ([]T, error) {
	var items []T
	for _, item := range it {
		items = append(items, item)
	}
	return items, nil
}

func (it InformerMock[T]) Close() error {
	return nil
}

type ResourceCache[T proto.Message] struct {
	cache *expirable.LRU[string, cacheItem[T]]
	gli   GetListIndex[T]
	getsf singleflight.Group

	fullSynced          atomic.Bool
	itemChangedCallback OnCacheItemChangeFunc

	stopChan chan struct{}
	stopOnce sync.Once
	stopped  atomic.Bool
	wg       sync.WaitGroup
}

var _ io.Closer = &ResourceCache[proto.Message]{}

type resourceCacheOptions struct {
	MaxEntries          int
	FullResyncInterval  time.Duration
	ItemChangedCallback OnCacheItemChangeFunc
}

type ResourceCacheOption func(*resourceCacheOptions)

// WithMaxEntries is a ResourceCacheOption that sets the maximum number of entries the cache can hold.
// When the cache reaches the maximum number of entries, the least recently used entries will be evicted
// to make room for new entries.
func WithMaxEntries(maxEntries int) ResourceCacheOption {
	return func(opts *resourceCacheOptions) {
		opts.MaxEntries = maxEntries
	}
}

// WithFullResyncInterval is a ResourceCacheOption that sets the full resync interval for the cache.
// The full resync interval determines how often the cache will fetch all items from the remote source
// to ensure the cache is up-to-date. Setting this to a non-zero value enables the full resync behavior.
func WithFullResyncInterval(fullResyncInterval time.Duration) ResourceCacheOption {
	return func(opts *resourceCacheOptions) {
		opts.FullResyncInterval = fullResyncInterval
	}
}

// WithNoFullSync is a ResourceCacheOption that disables the full resync interval for the cache.
// When this option is used, the cache will not periodically fetch all items from the remote source.
// This can be useful in situations where the remote source is expected to send updates for all
// changed items, and a full resync is not necessary.
func WithNoFullSync() ResourceCacheOption {
	return func(opts *resourceCacheOptions) {
		opts.FullResyncInterval = 0
	}
}

// WithItemChangedCallback is a ResourceCacheOption that sets a callback function to be called
// whenever an item in the cache is changed. The callback function will be passed the current
// context and the key of the cache item that was changed.
func WithItemChangedCallback(onCacheItemChange OnCacheItemChangeFunc) ResourceCacheOption {
	return func(opts *resourceCacheOptions) {
		opts.ItemChangedCallback = onCacheItemChange
	}
}

type cacheItem[T proto.Message] struct {
	Hash  []byte
	Value T
}

// IsZero returns true if the cacheItem has a nil Hash, indicating that it is an empty/zero value.
func (c cacheItem[T]) IsZero() bool {
	return c.Hash == nil
}

// OnCacheItemChangeFunc is a callback function that is called when an item in the cache is changed.
// The context.Context parameter provides the current context, and the key parameter is the key of the
// cache item that was changed.
type OnCacheItemChangeFunc func(ctx context.Context, key string)

// NewResourceCache creates a new ResourceCache that caches items of type T, indexed by keys of type K.
// The cache will fetch items from the provided GetListIndexWatch implementation, and will periodically
// refresh the cache by fetching all items and adding them to the cache.
func NewResourceCache[T proto.Message](ctx context.Context, gli GetListIndex[T], opts ...ResourceCacheOption) (*ResourceCache[T], error) {
	options := resourceCacheOptions{
		MaxEntries:         0,
		FullResyncInterval: 10 * time.Second,
	}
	for _, opt := range opts {
		opt(&options)
	}

	if options.MaxEntries < 0 {
		return nil, fmt.Errorf("%w: MaxEntries must be positive", ErrInvalidInput)
	}

	res := &ResourceCache[T]{
		gli:                 gli,
		itemChangedCallback: options.ItemChangedCallback,
		stopChan:            make(chan struct{}),
	}
	res.cache = expirable.NewLRU(options.MaxEntries, res.onEviction, options.FullResyncInterval+1*time.Minute)

	res.wg.Add(1)
	go func() {
		defer res.wg.Done()
		if err := res.fullSyncLoop(ctx, options.FullResyncInterval); err != nil {
			slog.ErrorContext(ctx, "Full sync loop exited with error", "err", err)
		}
	}()

	return res, nil
}

// onEviction is called when an item is evicted from the cache.
// It notifies the itemChangedCallback and updates the eviction metric.
func (rc *ResourceCache[T]) onEviction(key string, _ cacheItem[T]) {
	rc.notifyItemChange(context.Background(), key)
	cacheMetricsProvider.IncrementCacheEvictions()
}

// notifyItemChange calls the itemChangedCallback with the given key.
func (rc *ResourceCache[T]) notifyItemChange(ctx context.Context, key string) {
	if rc.itemChangedCallback != nil {
		rc.itemChangedCallback(ctx, key)
	}
}

// FullySynchronized returns true if the cache has been fully synchronized with the underlying data source.
func (rc *ResourceCache[T]) FullySynchronized() bool {
	return rc.fullSynced.Load()
}

// fullSyncLoop periodically refreshes the cache by fetching all items from the underlying GetListIndexWatch
// and adding them to the cache. This ensures the cache stays up-to-date even if items are added or removed
// externally. The loop runs until the ResourceCache is closed.
func (rc *ResourceCache[T]) fullSyncLoop(ctx context.Context, fullResyncInterval time.Duration) error {
	if fullResyncInterval == 0 {
		return nil
	}

	ticker := time.NewTicker(fullResyncInterval)
	defer ticker.Stop()

	for {
		err := rc.PerformFullSync(fullResyncInterval)
		if err != nil {
			slog.ErrorContext(ctx, "Failed to perform full sync", "err", err)
			cacheMetricsProvider.IncrementFullSyncFailures()
			// continue the loop even if sync fails
		}

		select {
		case <-ticker.C:
		case <-rc.stopChan:
			return nil
		}
	}
}

// PerformFullSync fetches all items from the underlying GetListIndexWatch and adds them to the cache.
func (rc *ResourceCache[T]) PerformFullSync(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	start := time.Now()
	defer func() {
		cacheMetricsProvider.ObserveFullSyncDuration(time.Since(start))
	}()

	items, err := rc.gli.List(ctx, 0, nil)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrSyncFailed, err)
	}

	rc.ingestItems(ctx, items)

	rc.fullSynced.Store(true)
	cacheMetricsProvider.SetTimeSinceLastSuccessfulSync(0)
	cacheMetricsProvider.SetCacheSize(rc.cache.Len())
	return nil
}

func (rc *ResourceCache[T]) ingestItems(ctx context.Context, items []T) {
	for _, item := range items {
		key := rc.gli.Index(item)

		itemHash, err := hashMessage(item)
		if err != nil {
			slog.WarnContext(ctx, "Failed to hash item during full sync", "key", key, "err", err)
			continue
		}

		elem, ok := rc.cache.Get(key)
		if ok && bytes.Equal(elem.Hash, itemHash) {
			// item is already present and up-to-date
			continue
		}

		rc.cache.Add(key, cacheItem[T]{Hash: itemHash, Value: item})

		// If the element was zero before, it was previously invalidated which would have notified the callback then.
		// Hence we only notify the callback if the element was previously valid.
		if !ok || !elem.IsZero() {
			rc.notifyItemChange(ctx, key)
		}
	}
}

// hashMessage computes the SHA256 hash of the given protobuf message, using deterministic marshaling.
// This ensures that the hash value is consistent across different runs, even if the message fields
// are serialized in a different order.
func hashMessage(m proto.Message) ([]byte, error) {
	h := sha256.New()
	s, err := proto.MarshalOptions{Deterministic: true}.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %w", err)
	}
	return h.Sum(s), nil
}

// Close closes the resource cache, stopping any background tasks and purging the cache.
// After calling Close, the cache can no longer be used.
func (rc *ResourceCache[T]) Close() error {
	rc.stopOnce.Do(func() {
		close(rc.stopChan)
		rc.stopped.Store(true)
	})
	rc.wg.Wait()
	rc.cache.Purge()
	return nil
}

// Get retrieves the value for the given key from the resource cache. If the value is not
// found in the cache, it will be fetched from the underlying data source and added to the cache.
// If the cache has been closed, this method will return an error.
func (rc *ResourceCache[T]) Get(ctx context.Context, key string) (T, error) {
	if rc.stopped.Load() {
		var zero T
		return zero, ErrCacheClosed
	}

	start := time.Now()
	defer func() {
		cacheMetricsProvider.ObserveCacheGetLatency(time.Since(start))
	}()

	notify := true
	if cached, ok := rc.cache.Get(key); ok && !cached.IsZero() {
		cacheMetricsProvider.IncrementCacheHits()
		return cached.Value, nil
	} else if ok {
		notify = false
	}

	cacheMetricsProvider.IncrementCacheMisses()

	val, err, _ := rc.getsf.Do(key, func() (interface{}, error) {
		res, err := rc.gli.Get(ctx, key)
		if err != nil {
			var zero T
			return zero, fmt.Errorf("%w: %w", ErrItemNotFound, err)
		}

		hash, err := hashMessage(res)
		if err != nil {
			return res, fmt.Errorf("failed to hash item: %w", err)
		}

		rc.cache.Add(key, cacheItem[T]{Hash: hash, Value: res})
		if notify {
			rc.notifyItemChange(ctx, key)
		}
		cacheMetricsProvider.SetCacheSize(rc.cache.Len())
		return res, nil
	})
	if err != nil {
		var zero T
		return zero, err
	}

	return val.(T), nil
}

// List retrieves all items currently in the cache. If the cache has been closed, this method will return an error.
//
// List is eventually consistent.
// It might not return items that actually exist, and it might return items that have been deleted.
func (rc *ResourceCache[T]) List(ctx context.Context) ([]T, error) {
	if rc.stopped.Load() {
		return nil, ErrCacheClosed
	}

	var (
		fetchIDs []string
		res      = make([]T, 0, rc.cache.Len())
	)
	for _, id := range rc.cache.Keys() {
		item, ok := rc.cache.Get(id)
		if !ok {
			continue
		}
		if item.IsZero() {
			fetchIDs = append(fetchIDs, id)
		} else {
			res = append(res, item.Value)
		}
	}
	if len(fetchIDs) > 0 {
		err := rc.Refetch(ctx, fetchIDs)
		if err != nil {
			return nil, err
		}

		for _, id := range fetchIDs {
			item, ok := rc.cache.Get(id)
			if !ok {
				continue
			}
			if item.IsZero() {
				continue
			}
			res = append(res, item.Value)
		}
	}

	return res, nil
}

// InvalidateItem invalidates the item with the given ID in the cache and notifies any listeners
// that the item has changed.
// Invalidated items are not removed from the cache, but they will be re-fetched from the underlying data source
// the next time they are requested through List or Get.
func (rc *ResourceCache[T]) InvalidateItem(ctx context.Context, id string) {
	rc.cache.Add(id, cacheItem[T]{})
	rc.notifyItemChange(ctx, id)
}

// RemoveItem removes the item with the given ID from the cache and notifies any listeners
// that the item has changed. If the item is not found in the cache, this method does nothing.
func (rc *ResourceCache[T]) RemoveItem(ctx context.Context, id string) {
	ok := rc.cache.Remove(id)
	if ok {
		rc.notifyItemChange(ctx, id)
		cacheMetricsProvider.SetCacheSize(rc.cache.Len())
	}
}

// Refetch retrieves the items with the given keys from the underlying data source and updates the cache.
// If more than 25 keys are provided, it will instead perform a full sync of the cache.
func (rc *ResourceCache[T]) Refetch(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return nil
	}

	// We can list max 25 items. If we need to refetch more than that, we'll just fall back to a full reconciliation.
	if len(keys) > 25 {
		var timeout time.Duration
		if deadline, ok := ctx.Deadline(); ok {
			timeout = time.Until(deadline)
		}
		if timeout < 30*time.Second {
			timeout = 30 * time.Second
		}

		return rc.PerformFullSync(timeout)
	}

	items, err := rc.gli.List(ctx, 0, keys)
	if err != nil {
		return fmt.Errorf("failed to refetch items: %w", err)
	}
	rc.ingestItems(ctx, items)

	return nil
}

// Stream is an interface that represents a stream of messages of type Res.
// The Receive method advances the stream to the next message, which can then
// be accessed through the Msg method. The Err method returns the first non-EOF
// error encountered by Receive. The Close method closes the stream.
type EventStream interface {
	// Receive advances the stream to the next message, which will then be
	// available through the Msg method. It returns false when the stream stops,
	// either by reaching the end or by encountering an unexpected error. After
	// Receive returns false, the Err method will return any unexpected error
	// encountered.
	Receive() bool

	// Msg returns the most recent message unmarshaled by a call to Receive.
	Msg() *v1.WatchEventsResponse

	// Err returns the first non-EOF error that was encountered by Receive.
	Err() error

	// Closes the stream, releasing any resources associated with it.
	Close() error
}

// WatchEventsFunc is a function type that represents a method for watching events from the EventServiceClient.
// The function takes a context.Context and a *connect.Request[v1.WatchEventsRequest] as arguments,
// and returns a Stream[v1.WatchEventsResponse] and an error.
type WatchEventsFunc func(context.Context, *connect.Request[v1.WatchEventsRequest]) (EventStream, error)

// AdaptWatchEvents returns a WatchEventsFunc that wraps the provided v1connect.EventServiceClient.
// The returned function calls the WatchEvents method on the provided client and returns the resulting Stream[v1.WatchEventsResponse] and error.
func AdaptWatchEvents(client v1connect.EventServiceClient) WatchEventsFunc {
	return func(ctx context.Context, req *connect.Request[v1.WatchEventsRequest]) (EventStream, error) {
		return client.WatchEvents(ctx, req)
	}
}

// InvalidateItemFunc matches the signature of a ResourceCache[T] InvalidateItem method.
// It's used to chain multiple invalidations to a single event stream client.
type Invalidator interface {
	RemoveItem(ctx context.Context, id string)
	InvalidateItem(ctx context.Context, id string)
}

// InvalidateFromEventService watches for events from the EventServiceClient and removes any matching
// items from the ResourceCache. This allows the cache to be kept up-to-date with changes
// made outside of the cache.
func InvalidateFromEventService(ctx context.Context, req *v1.WatchEventsRequest, watchEvents WatchEventsFunc, invalidators ...Invalidator) error {
	backoff := time.Second

	for ctx.Err() == nil {
		err := handleEventStream(ctx, req, watchEvents, invalidators)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}
			slog.ErrorContext(ctx, "[resource cache] error watching events", "err", err)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(backoff + time.Duration(rand.IntN(1000))*time.Millisecond):
				// exponential backoff with jitter
				backoff *= 2
				if backoff > time.Minute {
					backoff = time.Minute
				}
			}
		} else {
			// reset backoff on successful connection
			backoff = time.Second
		}
	}

	return ctx.Err()
}

// handleEventStream reads events from the EventServiceClient and invalidates cache items based on the event type.
func handleEventStream(ctx context.Context, req *v1.WatchEventsRequest, watchEvents WatchEventsFunc, invalidators []Invalidator) error {
	stream, err := watchEvents(ctx, connect.NewRequest(req))
	if err != nil {
		return fmt.Errorf("failed to start event stream: %w", err)
	}
	defer stream.Close()

	cacheMetricsProvider.IncrementEventStreamReconnections()

	for stream.Receive() {
		msg := stream.Msg()

		id := msg.GetResourceId()
		if id == "" {
			continue
		}
		for _, invalidator := range invalidators {
			if msg.Operation == v1.ResourceOperation_RESOURCE_OPERATION_DELETE {
				invalidator.RemoveItem(ctx, id)
			} else {
				invalidator.InvalidateItem(ctx, id)
			}
		}
		cacheMetricsProvider.IncrementEventsProcessed()
	}

	return stream.Err()
}

// WaitForFullSynchronization waits for the ResourceCache to be fully synchronized. It will block until the cache is fully synchronized or the context is canceled.
//
// Returns true if the cache is fully synchronized, false otherwise.
func WaitForFullSynchronization[T proto.Message](ctx context.Context, rc *ResourceCache[T]) (synchronized bool) {
	for {
		if ctx.Err() != nil {
			return false
		}
		if rc.FullySynchronized() {
			return true
		}
		time.Sleep(1 * time.Millisecond)
	}
}

// Requestable is an interface that defines the minimum requirements for an object
// that can be cached and retrieved by the ResourceCache. It requires that the
// object implement the proto.Message interface and provide a unique identifier
// via the GetId method.
type Requestable interface {
	proto.Message
	GetId() string
}

// Requester is an interface that defines the minimum requirements for an object
// that can be used to retrieve and list Requestable items. It provides methods
// to get a single item by its ID, and to list multiple items with pagination.
//
// Use this together with UseRequester to create a ResourceCache.
type Requester[T Requestable] interface {
	Get(context.Context, string) (T, error)
	List(ctx context.Context, page *v1.PaginationRequest, ids []string) ([]T, *v1.PaginationResponse, error)
}

// getListIndexAdapter is a helper struct that adapts a Requester implementation to the GetListIndex interface.
type getListIndexAdapter[T Requestable] struct {
	R Requester[T]
}

// Get implements GetListIndex.
func (g *getListIndexAdapter[T]) Get(ctx context.Context, key string) (T, error) {
	cacheMetricsProvider.IncrementExternalGetRequests()
	start := time.Now()
	result, err := g.R.Get(ctx, key)
	cacheMetricsProvider.ObserveExternalRequestLatency(time.Since(start))
	return result, err
}

// Index implements GetListIndex.
func (g *getListIndexAdapter[T]) Index(item T) string {
	return item.GetId()
}

// List implements GetListIndex.
func (g *getListIndexAdapter[T]) List(ctx context.Context, maxItems int, ids []string) ([]T, error) {
	res := make([]T, 0, maxItems)

	var nextToken string
	for {
		cacheMetricsProvider.IncrementExternalListRequests()
		start := time.Now()
		items, page, err := g.R.List(ctx, &v1.PaginationRequest{
			PageSize: 100,
			Token:    nextToken,
		}, ids)
		cacheMetricsProvider.ObserveExternalRequestLatency(time.Since(start))

		if err != nil {
			return nil, err
		}
		if page != nil {
			nextToken = page.NextToken
		}

		res = append(res, items...)
		if maxItems > 0 && len(res) >= maxItems {
			break
		}
		if nextToken == "" {
			break
		}
	}

	return res, nil
}

// UseRequester returns a GetListIndex implementation that wraps the provided Requester.
// The returned GetListIndexAdapter implements the GetListIndex interface by delegating
// to the provided Requester implementation.
func UseRequester[T Requestable](r Requester[T]) GetListIndex[T] {
	return &getListIndexAdapter[T]{
		R: r,
	}
}
