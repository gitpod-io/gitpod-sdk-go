package cache_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/gitpod-io/flex-sdk-go/client/cache"
	v1 "github.com/gitpod-io/flex-sdk-go/v1"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

//go:generate go run go.uber.org/mock/mockgen@v0.5.0 -package cache_test -destination cache_mock_test.go github.com/gitpod-io/flex-sdk-go/client/cache Invalidator,EventStream

func TestInvalidateFromEventService(t *testing.T) {
	t.Parallel()
	type Expectation struct {
		Error         string
		Invalidations []string
		Removals      []string
	}
	tests := []struct {
		Name        string
		Expectation Expectation
		Events      []*v1.WatchEventsResponse
	}{
		{
			Name: "no events",
		},
		{
			Name: "invalidate",
			Expectation: Expectation{
				Invalidations: []string{"foo", "bar"},
			},
			Events: []*v1.WatchEventsResponse{
				{
					Operation:    v1.ResourceOperation_RESOURCE_OPERATION_CREATE,
					ResourceType: v1.ResourceType_RESOURCE_TYPE_ENVIRONMENT,
					ResourceId:   "foo",
				},
				{
					Operation:    v1.ResourceOperation_RESOURCE_OPERATION_UPDATE,
					ResourceType: v1.ResourceType_RESOURCE_TYPE_ENVIRONMENT,
					ResourceId:   "bar",
				},
			},
		},
		{
			Name: "remove",
			Expectation: Expectation{
				Removals: []string{"foo"},
			},
			Events: []*v1.WatchEventsResponse{
				{
					Operation:    v1.ResourceOperation_RESOURCE_OPERATION_DELETE,
					ResourceType: v1.ResourceType_RESOURCE_TYPE_ENVIRONMENT,
					ResourceId:   "foo",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var act Expectation

			ctrl := gomock.NewController(t)
			t.Cleanup(ctrl.Finish)

			invalidator := NewMockInvalidator(ctrl)
			invalidator.EXPECT().InvalidateItem(gomock.Any(), gomock.Any()).Return().Do(func(ctx context.Context, key string) { act.Invalidations = append(act.Invalidations, key) }).AnyTimes()
			invalidator.EXPECT().RemoveItem(gomock.Any(), gomock.Any()).Return().Do(func(ctx context.Context, key string) { act.Removals = append(act.Removals, key) }).AnyTimes()

			ctx, cancel := context.WithCancel(context.Background())

			var eventIdx int
			stream := NewMockEventStream(ctrl)
			stream.EXPECT().Receive().DoAndReturn(func() bool {
				if eventIdx >= len(test.Events) {
					// we're done - give the last event some time
					go func() {
						time.Sleep(100 * time.Millisecond)
						cancel()
					}()
					return false
				}
				eventIdx++
				return true
			}).AnyTimes()
			stream.EXPECT().Close().AnyTimes()
			stream.EXPECT().Err().Return(nil).AnyTimes()
			stream.EXPECT().Msg().DoAndReturn(func() (*v1.WatchEventsResponse, error) {
				if eventIdx > len(test.Events) {
					return nil, nil
				}
				return test.Events[eventIdx-1], nil
			}).AnyTimes()

			err := cache.InvalidateFromEventService(ctx, &v1.WatchEventsRequest{
				Scope: &v1.WatchEventsRequest_EnvironmentId{
					EnvironmentId: "foo",
				},
			}, func(ctx context.Context, r *connect.Request[v1.WatchEventsRequest]) (cache.EventStream, error) {
				return stream, nil
			}, invalidator)
			if err != nil && !errors.Is(err, context.Canceled) {
				act.Error = err.Error()
			}

			if diff := cmp.Diff(test.Expectation, act); diff != "" {
				t.Errorf("InvalidateFromEventService() mismatch (-want  got):\n%s", diff)
			}
		})
	}
}
