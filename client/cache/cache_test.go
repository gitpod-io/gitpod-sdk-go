package cache

import (
	context "context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	v1 "github.com/gitpod-io/flex-sdk-go/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"google.golang.org/protobuf/testing/protocmp"
)

func must[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}
	return res
}

func TestGet(t *testing.T) {
	t.Parallel()
	type Expectation struct {
		Error  string
		Result *v1.GetAuthenticatedIdentityResponse
	}
	tests := []struct {
		Name        string
		Expectation Expectation
		GLI         *testGLI
		InitCache   func(cache *expirable.LRU[string, cacheItem[*v1.GetAuthenticatedIdentityResponse]])
		Key         string
	}{
		{
			Name: "not found",
			Expectation: Expectation{
				Error: "item not found in cache or data source: item not found in cache or data source",
			},
			Key: "unknown",
		},
		{
			Name: "cache miss",
			Expectation: Expectation{
				Error: "item not found in cache or data source: item not found in cache or data source",
			},
			GLI: &testGLI{items: map[string]*v1.GetAuthenticatedIdentityResponse{}},
			Key: "key",
		},
		{
			Name: "cache hit",
			Expectation: Expectation{
				Result: &v1.GetAuthenticatedIdentityResponse{OrganizationId: "key"},
			},
			InitCache: func(cache *expirable.LRU[string, cacheItem[*v1.GetAuthenticatedIdentityResponse]]) {
				cache.Add("key", cacheItem[*v1.GetAuthenticatedIdentityResponse]{
					Hash:  must(hashMessage(&v1.GetAuthenticatedIdentityResponse{OrganizationId: "key"})),
					Value: &v1.GetAuthenticatedIdentityResponse{OrganizationId: "key"},
				})
			},
			Key: "key",
		},
		{
			Name: "cache miss, data source hit",
			Expectation: Expectation{
				Result: &v1.GetAuthenticatedIdentityResponse{OrganizationId: "key"},
			},
			GLI: &testGLI{items: map[string]*v1.GetAuthenticatedIdentityResponse{
				"key": {OrganizationId: "key"},
			}},
			Key: "key",
		},
	}

	// nolint:govet
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if test.GLI == nil {
				test.GLI = &testGLI{items: map[string]*v1.GetAuthenticatedIdentityResponse{}}
			}

			var act Expectation

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			cache, err := NewResourceCache(ctx, test.GLI)
			if err != nil {
				t.Fatal(err)
			}
			defer cache.Close()

			if test.InitCache != nil {
				test.InitCache(cache.cache)
			}

			res, err := cache.Get(ctx, test.Key)
			if err != nil {
				act.Error = err.Error()
			}
			act.Result = res

			if diff := cmp.Diff(test.Expectation, act, protocmp.Transform()); diff != "" {
				t.Errorf("Get() mismatch (-want +got):\n%s", diff)
			}

			// Additional check for cache population on data source hit
			if test.Expectation.Error == "" && test.InitCache == nil {
				cachedItem, ok := cache.cache.Get(test.Key)
				if !ok {
					t.Errorf("Item not added to cache")
				} else if diff := cmp.Diff(test.Expectation.Result, cachedItem.Value, protocmp.Transform()); diff != "" {
					t.Errorf("Cached item mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestFullSyncLoop(t *testing.T) {
	t.Parallel()
	type Expectation struct {
		Error         string
		DidSync       bool
		Notifications []string
		Items         []string
	}
	tests := []struct {
		Name        string
		Expectation Expectation
		GLI         *testGLI
		InitCache   func(cache *expirable.LRU[string, cacheItem[*v1.GetAuthenticatedIdentityResponse]])
	}{
		{
			Name: "sync returns no items",
			GLI:  &testGLI{items: map[string]*v1.GetAuthenticatedIdentityResponse{}},
			Expectation: Expectation{
				DidSync: true,
				Items:   []string{},
			},
		},
		{
			Name: "sync does not yield different item",
			GLI: &testGLI{items: map[string]*v1.GetAuthenticatedIdentityResponse{
				"org-id": {OrganizationId: "org-id"},
			}},
			InitCache: func(cache *expirable.LRU[string, cacheItem[*v1.GetAuthenticatedIdentityResponse]]) {
				msg := &v1.GetAuthenticatedIdentityResponse{OrganizationId: "org-id"}
				cache.Add("org-id", cacheItem[*v1.GetAuthenticatedIdentityResponse]{
					Hash:  must(hashMessage(msg)),
					Value: msg,
				})
			},
			Expectation: Expectation{
				DidSync: true,
				Items:   []string{"org-id"},
			},
		},
		{
			Name: "sync yields different item",
			GLI: &testGLI{items: map[string]*v1.GetAuthenticatedIdentityResponse{
				"foobaz": {OrganizationId: "foobaz", Subject: &v1.Subject{Id: "update"}},
			}},
			InitCache: func(cache *expirable.LRU[string, cacheItem[*v1.GetAuthenticatedIdentityResponse]]) {
				msg := &v1.GetAuthenticatedIdentityResponse{OrganizationId: "foobaz", Subject: &v1.Subject{Id: "initial"}}
				cache.Add("foobaz", cacheItem[*v1.GetAuthenticatedIdentityResponse]{
					Hash:  must(hashMessage(msg)),
					Value: msg,
				})
			},
			Expectation: Expectation{
				DidSync:       true,
				Items:         []string{"foobaz"},
				Notifications: []string{"foobaz"},
			},
		},
	}

	// nolint:govet
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			var act Expectation

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			cache, err := NewResourceCache(ctx, test.GLI,
				WithItemChangedCallback(func(_ context.Context, key string) { act.Notifications = append(act.Notifications, key) }),
				WithNoFullSync(),
			)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { cache.Close() })

			if test.InitCache != nil {
				test.InitCache(cache.cache)
			}

			cache.PerformFullSync(500 * time.Millisecond)
			waitCtx, waitCancel := context.WithTimeout(ctx, 500*time.Millisecond)
			act.DidSync = WaitForFullSynchronization(waitCtx, cache)
			waitCancel()

			act.Items = cache.cache.Keys()

			if diff := cmp.Diff(test.Expectation, act); diff != "" {
				t.Errorf("fullSyncLoop() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestResourceCache_FullSync(t *testing.T) {
	t.Parallel()

	gli := &testGLI{
		items: map[string]*v1.GetAuthenticatedIdentityResponse{
			"key1": {OrganizationId: "org1"},
			"key2": {OrganizationId: "org2"},
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cache, err := NewResourceCache[*v1.GetAuthenticatedIdentityResponse](ctx, gli,
		WithFullResyncInterval(100*time.Millisecond),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer cache.Close()

	// Wait for the first sync to complete
	time.Sleep(150 * time.Millisecond)

	if !cache.FullySynchronized() {
		t.Error("Cache should be fully synchronized")
	}

	items, err := cache.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}
}

func TestResourceCache_RemoveItem(t *testing.T) {
	t.Parallel()

	gli := &testGLI{
		items: map[string]*v1.GetAuthenticatedIdentityResponse{
			"key1": {OrganizationId: "org1"},
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	removedKeyCh := make(chan string, 1)
	cache, err := NewResourceCache[*v1.GetAuthenticatedIdentityResponse](ctx, gli,
		WithItemChangedCallback(func(ctx context.Context, key string) {
			select {
			case removedKeyCh <- key:
			default:
			}
		}),
		WithNoFullSync(),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		cache.Close()
	})

	// Add item to cache
	_, err = cache.Get(context.Background(), "key1")
	if err != nil {
		t.Fatal(err)
	}

	// Remove item
	cache.RemoveItem(context.Background(), "key1")

	// Wait for the item changed callback
	select {
	case removedKey := <-removedKeyCh:
		if removedKey != "key1" {
			t.Errorf("Expected invalidated key to be 'key1', got '%s'", removedKey)
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for item changed callback")
	}

	// Check that item is no longer in cache
	items, err := cache.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items after invalidation, got %d", len(items))
	}
}

func TestInvalidateItem(t *testing.T) {
	t.Parallel()
	type Expectation struct {
		Error         string
		Notifications []string
		Items         []string
	}
	tests := []struct {
		Name             string
		Expectation      Expectation
		GLI              *testGLI
		Action           func(cache *ResourceCache[*v1.GetAuthenticatedIdentityResponse]) []*v1.GetAuthenticatedIdentityResponse
		PostInvalidation func(*testGLI)
	}{
		{
			Name: "get post invalidation",
			GLI: &testGLI{
				items: map[string]*v1.GetAuthenticatedIdentityResponse{
					"key1": {OrganizationId: "key1"},
				},
			},
			PostInvalidation: func(gli *testGLI) {
				gli.items = map[string]*v1.GetAuthenticatedIdentityResponse{
					"key1": {OrganizationId: "org1"},
				}
			},
			Action: func(cache *ResourceCache[*v1.GetAuthenticatedIdentityResponse]) []*v1.GetAuthenticatedIdentityResponse {
				r, _ := cache.Get(context.Background(), "key1")
				return []*v1.GetAuthenticatedIdentityResponse{r}
			},
			Expectation: Expectation{
				Notifications: []string{"key1"},
				Items:         []string{"org1"},
			},
		},
		{
			Name: "get removed post invalidation",
			GLI: &testGLI{
				items: map[string]*v1.GetAuthenticatedIdentityResponse{
					"key1": {OrganizationId: "key1"},
				},
			},
			PostInvalidation: func(gli *testGLI) {
				gli.items = map[string]*v1.GetAuthenticatedIdentityResponse{}
			},
			Action: func(cache *ResourceCache[*v1.GetAuthenticatedIdentityResponse]) []*v1.GetAuthenticatedIdentityResponse {
				r, _ := cache.Get(context.Background(), "key1")
				if r == nil {
					return nil
				}
				return []*v1.GetAuthenticatedIdentityResponse{r}
			},
			Expectation: Expectation{
				Notifications: []string{"key1"},
			},
		},
		{
			Name: "list post invalidation",
			GLI: &testGLI{
				items: map[string]*v1.GetAuthenticatedIdentityResponse{
					"key1": {OrganizationId: "key1"},
				},
			},
			Action: func(cache *ResourceCache[*v1.GetAuthenticatedIdentityResponse]) []*v1.GetAuthenticatedIdentityResponse {
				r, _ := cache.List(context.Background())
				return r
			},
			Expectation: Expectation{
				Notifications: []string{"key1"},
				Items:         []string{"key1"},
			},
		},
		{
			Name: "list removed post invalidation",
			GLI: &testGLI{
				items: map[string]*v1.GetAuthenticatedIdentityResponse{
					"key1": {OrganizationId: "key1"},
				},
			},
			PostInvalidation: func(gli *testGLI) {
				gli.items = map[string]*v1.GetAuthenticatedIdentityResponse{}
			},
			Action: func(cache *ResourceCache[*v1.GetAuthenticatedIdentityResponse]) []*v1.GetAuthenticatedIdentityResponse {
				r, _ := cache.List(context.Background())
				return r
			},
			Expectation: Expectation{
				Notifications: []string{"key1"},
			},
		},
		{
			Name: "new item",
			GLI:  &testGLI{},
			PostInvalidation: func(gli *testGLI) {
				gli.items = map[string]*v1.GetAuthenticatedIdentityResponse{
					"key1": {OrganizationId: "key1"},
				}
			},
			Action: func(cache *ResourceCache[*v1.GetAuthenticatedIdentityResponse]) []*v1.GetAuthenticatedIdentityResponse {
				r, _ := cache.Get(context.Background(), "key1")
				return []*v1.GetAuthenticatedIdentityResponse{r}
			},
			Expectation: Expectation{
				Notifications: []string{"key1"},
				Items:         []string{"key1"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var act Expectation

			gli := test.GLI
			if gli == nil {
				gli = &testGLI{}
			}

			ctx := context.Background()
			cache, err := NewResourceCache[*v1.GetAuthenticatedIdentityResponse](ctx, gli,
				WithItemChangedCallback(func(ctx context.Context, key string) {
					act.Notifications = append(act.Notifications, key)
				}),
				WithFullResyncInterval(0),
				WithMaxEntries(100),
			)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { cache.Close() })
			for key, item := range gli.items {
				cache.cache.Add(key, cacheItem[*v1.GetAuthenticatedIdentityResponse]{
					Hash:  must(hashMessage(item)),
					Value: item,
				})
			}

			// invalidate item
			cache.InvalidateItem(context.Background(), "key1")
			if test.PostInvalidation != nil {
				test.PostInvalidation(gli)
			}

			rawItems := test.Action(cache)
			for _, item := range rawItems {
				act.Items = append(act.Items, item.OrganizationId)
			}

			if diff := cmp.Diff(test.Expectation, act); diff != "" {
				t.Errorf("InvalidateItem() mismatch (-want  got):\n%s", diff)
			}
		})
	}
}

func TestResourceCache_Close(t *testing.T) {
	t.Parallel()

	gli := &testGLI{
		items: map[string]*v1.GetAuthenticatedIdentityResponse{
			"key1": {OrganizationId: "org1"},
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cache, err := NewResourceCache[*v1.GetAuthenticatedIdentityResponse](ctx, gli)
	if err != nil {
		t.Fatal(err)
	}

	// Add item to cache
	_, err = cache.Get(context.Background(), "key1")
	if err != nil {
		t.Fatal(err)
	}

	// Close the cache
	err = cache.Close()
	if err != nil {
		t.Fatal(err)
	}

	// Attempt to get item after close
	_, err = cache.Get(context.Background(), "key1")
	if !errors.Is(err, ErrCacheClosed) {
		t.Errorf("Expected ErrCacheClosed, got %v", err)
	}

	// Attempt to list items after close
	_, err = cache.List(context.Background())
	if !errors.Is(err, ErrCacheClosed) {
		t.Errorf("Expected ErrCacheClosed, got %v", err)
	}
}

type testGLI struct {
	items    map[string]*v1.GetAuthenticatedIdentityResponse
	getCount int
	mu       sync.Mutex
	listErr  error
}

func (t *testGLI) Get(ctx context.Context, key string) (*v1.GetAuthenticatedIdentityResponse, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.getCount++
	res, ok := t.items[key]
	if !ok {
		return nil, ErrItemNotFound
	}
	return res, nil
}

func (t *testGLI) List(ctx context.Context, maxItems int, ids []string) ([]*v1.GetAuthenticatedIdentityResponse, error) {
	if t != nil && t.listErr != nil {
		return nil, t.listErr
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	res := make([]*v1.GetAuthenticatedIdentityResponse, 0, len(t.items))

	if len(ids) > 0 {
		for _, id := range ids {
			item, ok := t.items[id]
			if ok {
				res = append(res, item)
			}
			if maxItems != 0 && len(res) == maxItems {
				break
			}
		}
	} else {
		for _, item := range t.items {
			res = append(res, item)
			if maxItems != 0 && len(res) == maxItems {
				break
			}
		}
	}
	return res, nil
}

func (t *testGLI) Index(item *v1.GetAuthenticatedIdentityResponse) string {
	return item.OrganizationId
}

// performanceBudget defines the performance expectations for a benchmark
type performanceBudget struct {
	opsPerSecond float64
	allocsPerOp  int64
	bytesPerOp   int64
}

// checkPerformanceBudget checks if the benchmark results are within the defined budget
func checkPerformanceBudget(b *testing.B, budget performanceBudget) {
	b.Helper()
	result := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for range b.N {
			// Run the benchmark function here
			b.StopTimer()
			b.StartTimer()
		}
	})

	opsPerSecond := float64(result.N) / result.T.Seconds()
	if opsPerSecond < budget.opsPerSecond {
		b.Fatalf("Performance budget exceeded: got %.2f ops/sec, want at least %.2f ops/sec", opsPerSecond, budget.opsPerSecond)
	}
	if result.AllocsPerOp() > budget.allocsPerOp {
		b.Fatalf("Allocation budget exceeded: got %d allocs/op, want at most %d allocs/op", result.AllocsPerOp(), budget.allocsPerOp)
	}
	if result.AllocedBytesPerOp() > budget.bytesPerOp {
		b.Fatalf("Allocation budget exceeded: got %d bytes/op, want at most %d bytes/op", result.AllocedBytesPerOp(), budget.bytesPerOp)
	}
}

// Run benchmarks running:
// go test -benchmem -run='^$' -bench '^(BenchmarkResourceCache_Get|BenchmarkResourceCache_List|BenchmarkResourceCache_FullSync|BenchmarkResourceCache_Concurrent|BenchmarkWithBudgets)$' github.com/gitpod-io/gitpod-next/runner/shared/supervisor/internal/automations
func BenchmarkResourceCache_Get(b *testing.B) {
	gli := &testGLI{
		items: make(map[string]*v1.GetAuthenticatedIdentityResponse),
	}
	for i := range 1000 {
		key := fmt.Sprintf("key-%d", i)
		gli.items[key] = &v1.GetAuthenticatedIdentityResponse{OrganizationId: key}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cache, err := NewResourceCache(ctx, gli)
	if err != nil {
		b.Fatal(err)
	}
	defer cache.Close()

	b.ResetTimer()
	for i := range b.N {
		key := fmt.Sprintf("key-%d", i%1000)
		_, err := cache.Get(context.Background(), key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkResourceCache_List(b *testing.B) {
	gli := &testGLI{
		items: make(map[string]*v1.GetAuthenticatedIdentityResponse),
	}
	for i := range 1000 {
		key := fmt.Sprintf("key-%d", i)
		gli.items[key] = &v1.GetAuthenticatedIdentityResponse{OrganizationId: key}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cache, err := NewResourceCache(ctx, gli)
	if err != nil {
		b.Fatal(err)
	}
	defer cache.Close()

	// Populate the cache
	for key := range gli.items {
		_, err := cache.Get(ctx, key)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for range b.N {
		_, err := cache.List(ctx)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkResourceCache_FullSync(b *testing.B) {
	gli := &testGLI{
		items: make(map[string]*v1.GetAuthenticatedIdentityResponse),
	}
	for i := range 1000 {
		key := fmt.Sprintf("key-%d", i)
		gli.items[key] = &v1.GetAuthenticatedIdentityResponse{OrganizationId: key}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache, err := NewResourceCache(ctx, gli, WithFullResyncInterval(time.Millisecond))
		if err != nil {
			b.Fatal(err)
		}

		// Wait for the first sync to complete
		didSync := WaitForFullSynchronization(ctx, cache)
		if !didSync {
			b.Fatal("Cache did not fully sync")
		}

		cache.Close()
	}
}

func BenchmarkResourceCache_Concurrent(b *testing.B) {
	gli := &testGLI{
		items: make(map[string]*v1.GetAuthenticatedIdentityResponse),
	}
	for i := range 1000 {
		key := fmt.Sprintf("key-%d", i)
		gli.items[key] = &v1.GetAuthenticatedIdentityResponse{OrganizationId: key}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cache, err := NewResourceCache(ctx, gli)
	if err != nil {
		b.Fatal(err)
	}
	defer cache.Close()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("key-%d", i%1000)
			_, err := cache.Get(ctx, key)
			if err != nil {
				b.Error(err)
			}
			i++
		}
	})
}

// Add this function to run all benchmarks with performance budgets
func BenchmarkWithBudgets(b *testing.B) {
	benchmarks := []struct {
		name   string
		bench  func(*testing.B)
		budget performanceBudget
	}{
		{"ResourceCache_Get", BenchmarkResourceCache_Get, performanceBudget{opsPerSecond: 100000, allocsPerOp: 2, bytesPerOp: 64}},
		{"ResourceCache_List", BenchmarkResourceCache_List, performanceBudget{opsPerSecond: 10000, allocsPerOp: 5, bytesPerOp: 1024}},
		{"ResourceCache_FullSync", BenchmarkResourceCache_FullSync, performanceBudget{opsPerSecond: 100, allocsPerOp: 10000, bytesPerOp: 1e6}},
		{"ResourceCache_Concurrent", BenchmarkResourceCache_Concurrent, performanceBudget{opsPerSecond: 500000, allocsPerOp: 2, bytesPerOp: 64}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			result := testing.Benchmark(bm.bench)
			checkPerformanceBudget(b, bm.budget)
			b.Logf("%s: %s", bm.name, result.String())
		})
	}
}

func TestRefetch(t *testing.T) {
	t.Parallel()
	type Expectation struct {
		Error  string
		Result map[string]*v1.GetAuthenticatedIdentityResponse
	}
	tests := []struct {
		Name        string
		Expectation Expectation
		GLI         *testGLI
		InitCache   func(cache *expirable.LRU[string, cacheItem[*v1.GetAuthenticatedIdentityResponse]])
		Keys        []string
	}{
		{
			Name: "empty refetch",
			GLI:  &testGLI{},
			Keys: []string{},
		},
		{
			Name: "successful refetch",
			Expectation: Expectation{
				Result: map[string]*v1.GetAuthenticatedIdentityResponse{
					"key1": {OrganizationId: "key1"},
					"key2": {OrganizationId: "key2"},
				},
			},
			GLI: &testGLI{items: map[string]*v1.GetAuthenticatedIdentityResponse{
				"key1": {OrganizationId: "key1"},
				"key2": {OrganizationId: "key2"},
			}},
			Keys: []string{"key1", "key2"},
		},
		{
			Name: "refetch with GLI error",
			Expectation: Expectation{
				Error: "failed to refetch items: GLI error",
			},
			GLI:  &testGLI{listErr: errors.New("GLI error")},
			Keys: []string{"key1", "key2"},
		},
		{
			Name: "refetch more than 25 keys",
			Expectation: Expectation{
				Result: map[string]*v1.GetAuthenticatedIdentityResponse{
					"key1": {OrganizationId: "key1"},
					"key2": {OrganizationId: "key2"},
				},
			},
			GLI: &testGLI{items: map[string]*v1.GetAuthenticatedIdentityResponse{
				"key1": {OrganizationId: "key1"},
				"key2": {OrganizationId: "key2"},
			}},
			Keys: func() []string {
				keys := make([]string, 26)
				for i := range keys {
					keys[i] = fmt.Sprintf("key%d", i)
				}
				return keys
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var act Expectation

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			cache, err := NewResourceCache(ctx, test.GLI)
			if err != nil {
				t.Fatal(err)
			}
			defer cache.Close()

			if test.InitCache != nil {
				test.InitCache(cache.cache)
			}

			err = cache.Refetch(ctx, test.Keys)
			if err != nil {
				act.Error = err.Error()
			}

			act.Result = make(map[string]*v1.GetAuthenticatedIdentityResponse)
			for _, key := range test.Keys {
				if item, ok := cache.cache.Get(key); ok {
					act.Result[key] = item.Value
				}
			}
			if len(act.Result) == 0 {
				act.Result = nil
			}

			if diff := cmp.Diff(test.Expectation, act, protocmp.Transform()); diff != "" {
				t.Errorf("Refetch() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
