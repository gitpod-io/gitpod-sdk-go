package client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"

	"github.com/gitpod-io/flex-go/client/cache"
	v1 "github.com/gitpod-io/flex-go/v1"
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

var serviceInformerCache = make(map[string]cache.Informer[*v1.Service])

func GetServiceInformer(ctx context.Context, client *ManagementPlane, envID string) (cache.Informer[*v1.Service], error) {
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

func GetServiceInformerWithoutCache(ctx context.Context, client *ManagementPlane, envID string) (cache.Informer[*v1.Service], error) {
	serviceCache, err := cache.NewServiceCache(ctx, client.EnvironmentAutomationService(), &v1.ListServicesRequest_Filter{
		EnvironmentIds: []string{envID},
	}, cache.WithFullResyncInterval(1*time.Minute))
	if err != nil {
		return nil, err
	}

	go func() {
		err := cache.InvalidateFromEventService(ctx, &v1.WatchEventsRequest{
			Scope: &v1.WatchEventsRequest_EnvironmentId{
				EnvironmentId: envID,
			},
		}, cache.AdaptWatchEvents(client.EventService()), serviceCache)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}

			slog.ErrorContext(ctx, "event based cache invalidation has failed", "err", err)
		}
	}()

	return serviceCache, nil
}

func SetServiceInformer(envID string, informer cache.Informer[*v1.Service]) {
	serviceInformerCache[envID] = informer
}

var taskExecutionInformerCache = make(map[string]cache.Informer[*v1.TaskExecution])

func GetTaskExecutionInformer(ctx context.Context, client *ManagementPlane, envID string) (cache.Informer[*v1.TaskExecution], error) {
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

func GetTaskExecutionInformerWithoutCache(ctx context.Context, client *ManagementPlane, envID string) (cache.Informer[*v1.TaskExecution], error) {
	taskExecutionCache, err := cache.NewTaskExecutionCache(ctx, client.EnvironmentAutomationService(), &v1.ListTaskExecutionsRequest_Filter{
		EnvironmentIds: []string{envID},
	}, cache.WithFullResyncInterval(1*time.Minute))
	if err != nil {
		return nil, err
	}

	go func() {
		err := cache.InvalidateFromEventService(ctx, &v1.WatchEventsRequest{
			Scope: &v1.WatchEventsRequest_EnvironmentId{
				EnvironmentId: envID,
			},
		}, cache.AdaptWatchEvents(client.EventService()), taskExecutionCache)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}

			slog.ErrorContext(ctx, "event based cache invalidation has failed", "err", err)
		}
	}()

	return taskExecutionCache, nil
}

func SetTaskExecutionInformer(envID string, informer cache.Informer[*v1.TaskExecution]) {
	taskExecutionInformerCache[envID] = informer
}
