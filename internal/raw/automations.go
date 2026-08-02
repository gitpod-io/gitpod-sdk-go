package client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"sync"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"

	"github.com/gitpod-io/gitpod-sdk-go/internal/raw/cache"
	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
)

// ReadAutomationsFile reads an AutomationsFile from the provided io.Reader.
// It decodes the YAML input, marshals it to JSON, and then unmarshals it
// into a v1.AutomationsFile proto.
func ReadAutomationsFile(in io.Reader) (*v1.AutomationsFile, error) {
	raw := make(map[string]any)
	err := yaml.NewDecoder(in).Decode(&raw)
	if err != nil {
		return nil, err
	}
	fc, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}

	var res v1.AutomationsFile
	opts := protojson.UnmarshalOptions{
		// Error if there is any unknown fields, ensures we don't silently ignore e.g. typos.
		DiscardUnknown: false,
	}
	err = opts.Unmarshal(fc, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

var serviceInformerCache = make(map[string]*cache.ResourceCache[*v1.Service])
var serviceInformerCacheMutex sync.RWMutex

func GetServiceInformer(ctx context.Context, client *ManagementPlane, envID string) (*cache.ResourceCache[*v1.Service], error) {
	serviceInformerCacheMutex.RLock()
	defer serviceInformerCacheMutex.RUnlock()

	if serviceInformerCache[envID] != nil {
		return serviceInformerCache[envID], nil
	}

	informer, err := GetServiceInformerWithoutCache(ctx, client, envID)
	if err != nil {
		return nil, err
	}

	serviceInformerCache[envID] = informer
	return informer, nil
}

func GetServiceInformerWithoutCache(ctx context.Context, client *ManagementPlane, envID string) (*cache.ResourceCache[*v1.Service], error) {
	serviceCache, err := cache.NewServiceCache(ctx, client.EnvironmentAutomationService(), &v1.ListServicesRequest_Filter{
		EnvironmentIds: []string{envID},
	}, cache.WithFullResyncInterval(1*time.Minute))
	if err != nil {
		return nil, err
	}

	go func() {
		// Only set up event-based invalidation if EventService is available
		eventService := client.EventService()
		if eventService != nil {
			// Additional check to ensure the service is properly initialized
			defer func() {
				if r := recover(); r != nil {
					slog.WarnContext(ctx, "event based cache invalidation failed due to panic", "panic", r)
				}
			}()

			err := cache.InvalidateFromEventService(ctx, &v1.WatchEventsRequest{
				Scope: &v1.WatchEventsRequest_EnvironmentId{
					EnvironmentId: envID,
				},
			}, cache.AdaptWatchEvents(eventService),
				cache.ForResourceTypes(serviceCache, v1.ResourceType_RESOURCE_TYPE_SERVICE),
			)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}

				slog.ErrorContext(ctx, "event based cache invalidation has failed", "err", err)
			}
		}
	}()

	return serviceCache, nil
}

func SetServiceInformer(envID string, informer *cache.ResourceCache[*v1.Service]) {
	serviceInformerCacheMutex.Lock()
	defer serviceInformerCacheMutex.Unlock()

	serviceInformerCache[envID] = informer
}

var taskExecutionInformerCache = make(map[string]*cache.ResourceCache[*v1.TaskExecution])
var taskExecutionInformerCacheMutex sync.RWMutex

func GetTaskExecutionInformer(ctx context.Context, client *ManagementPlane, envID string) (*cache.ResourceCache[*v1.TaskExecution], error) {
	taskExecutionInformerCacheMutex.RLock()
	defer taskExecutionInformerCacheMutex.RUnlock()

	if taskExecutionInformerCache[envID] != nil {
		return taskExecutionInformerCache[envID], nil
	}

	informer, err := GetTaskExecutionInformerWithoutCache(ctx, client, envID)
	if err != nil {
		return nil, err
	}

	taskExecutionInformerCache[envID] = informer
	return informer, nil
}

func GetTaskExecutionInformerWithoutCache(ctx context.Context, client *ManagementPlane, envID string) (*cache.ResourceCache[*v1.TaskExecution], error) {
	taskExecutionCache, err := cache.NewTaskExecutionCache(ctx, client.EnvironmentAutomationService(), &v1.ListTaskExecutionsRequest_Filter{
		EnvironmentIds: []string{envID},
	}, cache.WithFullResyncInterval(1*time.Minute))
	if err != nil {
		return nil, err
	}

	go func() {
		// Only set up event-based invalidation if EventService is available
		eventService := client.EventService()
		if eventService != nil {
			// Additional check to ensure the service is properly initialized
			defer func() {
				if r := recover(); r != nil {
					slog.WarnContext(ctx, "event based cache invalidation failed due to panic", "panic", r)
				}
			}()

			err := cache.InvalidateFromEventService(ctx, &v1.WatchEventsRequest{
				Scope: &v1.WatchEventsRequest_EnvironmentId{
					EnvironmentId: envID,
				},
			}, cache.AdaptWatchEvents(eventService),
				cache.ForResourceTypes(taskExecutionCache, v1.ResourceType_RESOURCE_TYPE_TASK_EXECUTION),
			)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}

				slog.ErrorContext(ctx, "event based cache invalidation has failed", "err", err)
			}
		}
	}()

	return taskExecutionCache, nil
}

func SetTaskExecutionInformer(envID string, informer *cache.ResourceCache[*v1.TaskExecution]) {
	taskExecutionInformerCacheMutex.Lock()
	defer taskExecutionInformerCacheMutex.Unlock()

	taskExecutionInformerCache[envID] = informer
}

var environmentInformerCache = make(map[string]*cache.ResourceCache[*v1.Environment])
var environmentInformerCacheMutex sync.Mutex

func GetEnvironmentInformer(ctx context.Context, client *ManagementPlane, envID string) (*cache.ResourceCache[*v1.Environment], error) {
	environmentInformerCacheMutex.Lock()
	defer environmentInformerCacheMutex.Unlock()

	if environmentInformerCache[envID] != nil {
		return environmentInformerCache[envID], nil
	}

	fullResyncDuration := 1 * time.Minute
	informer, err := cache.NewEnvironmentCache(ctx, client.EnvironmentService(), envID, cache.WithFullResyncInterval(fullResyncDuration))
	if err != nil {
		return nil, err
	}
	environmentInformerCache[envID] = informer

	go func() {
		// Only set up event-based invalidation if EventService is available
		eventService := client.EventService()
		if eventService != nil {
			// Additional check to ensure the service is properly initialized
			defer func() {
				if r := recover(); r != nil {
					slog.WarnContext(ctx, "event based cache invalidation failed due to panic", "panic", r)
				}
			}()

			err := cache.InvalidateFromEventService(ctx, &v1.WatchEventsRequest{
				Scope: &v1.WatchEventsRequest_EnvironmentId{
					EnvironmentId: envID,
				},
			}, cache.AdaptWatchEvents(eventService),
				cache.ForResourceTypes(informer, v1.ResourceType_RESOURCE_TYPE_ENVIRONMENT),
			)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}

				slog.ErrorContext(ctx, "event based cache invalidation has failed", "err", err)
			}
		}
	}()

	return informer, nil
}
