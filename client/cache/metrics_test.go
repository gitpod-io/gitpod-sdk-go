package cache

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	v1 "github.com/gitpod-io/flex-go/v1"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimestampWrapper implements the Requestable interface
type TimestampWrapper struct {
	*timestamppb.Timestamp
}

func (tw TimestampWrapper) GetId() string {
	return tw.AsTime().Format(time.RFC3339)
}

// MetricsRecorder is a simple implementation of MetricsProvider that records metric calls
type MetricsRecorder struct {
	Calls []string
	mu    sync.Mutex
}

func (m *MetricsRecorder) recordCall(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Calls = append(m.Calls, name)
}

func (m *MetricsRecorder) IncrementCacheHits()      { m.recordCall("IncrementCacheHits") }
func (m *MetricsRecorder) IncrementCacheMisses()    { m.recordCall("IncrementCacheMisses") }
func (m *MetricsRecorder) IncrementCacheEvictions() { m.recordCall("IncrementCacheEvictions") }
func (m *MetricsRecorder) SetCacheSize(size int)    { m.recordCall("SetCacheSize") }
func (m *MetricsRecorder) ObserveCacheGetLatency(latency time.Duration) {
	m.recordCall("ObserveCacheGetLatency")
}
func (m *MetricsRecorder) ObserveFullSyncDuration(duration time.Duration) {
	m.recordCall("ObserveFullSyncDuration")
}
func (m *MetricsRecorder) IncrementFullSyncFailures() { m.recordCall("IncrementFullSyncFailures") }
func (m *MetricsRecorder) SetTimeSinceLastSuccessfulSync(duration time.Duration) {
	m.recordCall("SetTimeSinceLastSuccessfulSync")
}
func (m *MetricsRecorder) IncrementExternalGetRequests() {
	m.recordCall("IncrementExternalGetRequests")
}
func (m *MetricsRecorder) IncrementExternalListRequests() {
	m.recordCall("IncrementExternalListRequests")
}
func (m *MetricsRecorder) ObserveExternalRequestLatency(latency time.Duration) {
	m.recordCall("ObserveExternalRequestLatency")
}
func (m *MetricsRecorder) IncrementEventsProcessed() { m.recordCall("IncrementEventsProcessed") }
func (m *MetricsRecorder) IncrementEventStreamReconnections() {
	m.recordCall("IncrementEventStreamReconnections")
}

// MockRequester is a mock implementation of the Requester interface
type MockRequester struct {
	Items map[string]TimestampWrapper
}

func (m *MockRequester) Get(_ context.Context, key string) (TimestampWrapper, error) {
	return m.Items[key], nil
}

func (m *MockRequester) List(_ context.Context, page *v1.PaginationRequest, ids []string) ([]TimestampWrapper, *v1.PaginationResponse, error) {
	var items []TimestampWrapper
	for _, v := range m.Items {
		items = append(items, v)
	}
	return items, &v1.PaginationResponse{}, nil
}

func TestGetListIndexAdapterMetrics(t *testing.T) {
	type Expectation struct {
		Metrics string
	}

	tests := []struct {
		name        string
		action      func(GetListIndex[TimestampWrapper]) error
		expectation Expectation
	}{
		{
			name: "Get",
			action: func(gli GetListIndex[TimestampWrapper]) error {
				_, err := gli.Get(context.Background(), "2023-01-01T00:00:00Z")
				return err
			},
			expectation: Expectation{
				Metrics: "IncrementExternalGetRequests,ObserveExternalRequestLatency",
			},
		},
		{
			name: "List",
			action: func(gli GetListIndex[TimestampWrapper]) error {
				_, err := gli.List(context.Background(), 10, nil)
				return err
			},
			expectation: Expectation{
				Metrics: "IncrementExternalListRequests,ObserveExternalRequestLatency",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metricsRecorder := &MetricsRecorder{}
			SetMetricsProvider(metricsRecorder)

			mockRequester := &MockRequester{
				Items: map[string]TimestampWrapper{
					"2023-01-01T00:00:00Z": {Timestamp: &timestamppb.Timestamp{Seconds: 1672531200}},
				},
			}

			adapter := UseRequester[TimestampWrapper](mockRequester)

			_ = tt.action(adapter) // We're not testing for errors here, just metrics

			actualMetrics := cmp.Diff("", formatMetrics(metricsRecorder.Calls))
			expectedMetrics := cmp.Diff("", tt.expectation.Metrics)

			if actualMetrics != expectedMetrics {
				t.Errorf("Metrics mismatch (-want +got):\n%s", cmp.Diff(tt.expectation.Metrics, formatMetrics(metricsRecorder.Calls)))
			}
		})
	}
}

func TestResourceCacheMetrics(t *testing.T) {
	scenarios := []struct {
		name             string
		setup            func(*testGLI)
		actions          func(*ResourceCache[*v1.GetAuthenticatedIdentityResponse], *testGLI)
		expectedMetrics  string
		expectedGetCount int
	}{
		{
			name: "Initial Sync",
			setup: func(gli *testGLI) {
				for i := range 100 {
					key := fmt.Sprintf("key-%d", i)
					gli.items[key] = &v1.GetAuthenticatedIdentityResponse{OrganizationId: key}
				}
			},
			actions: func(rc *ResourceCache[*v1.GetAuthenticatedIdentityResponse], gli *testGLI) {
				// Wait for initial sync to complete
				// No actions
				time.Sleep(10 * time.Millisecond)
			},
			expectedMetrics:  "",
			expectedGetCount: 0,
		},
		{
			name: "Cache Hit and Miss",
			setup: func(gli *testGLI) {
				gli.items["existing"] = &v1.GetAuthenticatedIdentityResponse{OrganizationId: "existing"}
			},
			actions: func(rc *ResourceCache[*v1.GetAuthenticatedIdentityResponse], gli *testGLI) {
				rc.Get(context.Background(), "existing")
				rc.Get(context.Background(), "non-existing")
			},
			expectedMetrics:  "IncrementCacheMisses,SetCacheSize,ObserveCacheGetLatency,IncrementCacheMisses,ObserveCacheGetLatency",
			expectedGetCount: 2, // Both existing and non-existing keys trigger a Get
		},
		{
			name: "Invalidate Item",
			setup: func(gli *testGLI) {
				gli.items["to-invalidate"] = &v1.GetAuthenticatedIdentityResponse{OrganizationId: "to-invalidate"}
			},
			actions: func(rc *ResourceCache[*v1.GetAuthenticatedIdentityResponse], gli *testGLI) {
				rc.Get(context.Background(), "to-invalidate") // Ensure item is in cache
				rc.InvalidateItem(context.Background(), "to-invalidate")
				rc.Get(context.Background(), "to-invalidate") // This should trigger another Get
				time.Sleep(10 * time.Millisecond)             // Wait for cache to update
			},
			expectedMetrics:  "IncrementCacheMisses,SetCacheSize,ObserveCacheGetLatency,IncrementCacheMisses,SetCacheSize,ObserveCacheGetLatency",
			expectedGetCount: 2, // One for initial Get, one after invalidation
		},
		{
			name: "Full Sync with Existing Items",
			setup: func(gli *testGLI) {
				for i := range 100 {
					key := fmt.Sprintf("key-%d", i)
					gli.items[key] = &v1.GetAuthenticatedIdentityResponse{OrganizationId: key}
				}
			},
			actions: func(rc *ResourceCache[*v1.GetAuthenticatedIdentityResponse], gli *testGLI) {
				rc.PerformFullSync(100 * time.Millisecond)
			},
			expectedMetrics:  "SetTimeSinceLastSuccessfulSync,SetCacheSize,ObserveFullSyncDuration",
			expectedGetCount: 0,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			metricsRecorder := &MetricsRecorder{}
			SetMetricsProvider(metricsRecorder)

			gli := &testGLI{
				items: make(map[string]*v1.GetAuthenticatedIdentityResponse),
				mu:    sync.Mutex{},
			}

			scenario.setup(gli)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			cache, err := NewResourceCache[*v1.GetAuthenticatedIdentityResponse](ctx, gli, WithNoFullSync())
			if err != nil {
				t.Fatal(err)
			}
			defer cache.Close()

			scenario.actions(cache, gli)

			actualMetrics := formatMetrics(metricsRecorder.Calls)
			if diff := cmp.Diff(scenario.expectedMetrics, actualMetrics); diff != "" {
				t.Errorf("Metrics mismatch (-want +got):\n%s", diff)
			}

			if gli.getCount != scenario.expectedGetCount {
				t.Errorf("Expected Get count: %d, but got: %d", scenario.expectedGetCount, gli.getCount)
			}
		})
	}
}

// formatMetrics converts a slice of metric calls to a comma-separated string
func formatMetrics(calls []string) string {
	return strings.Join(calls, ",")
}
