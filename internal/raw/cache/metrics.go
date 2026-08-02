package cache

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// MetricsProvider defines the interface for providing metrics for the ResourceCache.
type MetricsProvider interface {
	// Cache Performance Metrics
	IncrementCacheHits()
	IncrementCacheMisses()
	IncrementCacheEvictions()
	SetCacheSize(size int)
	ObserveCacheGetLatency(latency time.Duration)

	// Synchronization Metrics
	ObserveFullSyncDuration(duration time.Duration)
	IncrementFullSyncFailures()
	SetTimeSinceLastSuccessfulSync(duration time.Duration)

	// External Interactions Metrics
	IncrementExternalGetRequests()
	IncrementExternalListRequests()
	ObserveExternalRequestLatency(latency time.Duration)

	// Event Service Metrics
	IncrementEventsProcessed()
	IncrementEventStreamReconnections()
}

// PrometheusMetricsProvider implements MetricsProvider using Prometheus.
type PrometheusMetricsProvider struct {
	cacheHits                   prometheus.Counter
	cacheMisses                 prometheus.Counter
	cacheEvictions              prometheus.Counter
	cacheSize                   prometheus.Gauge
	cacheGetLatency             prometheus.Histogram
	fullSyncDuration            prometheus.Histogram
	fullSyncFailures            prometheus.Counter
	timeSinceLastSuccessfulSync prometheus.Gauge
	externalGetRequests         prometheus.Counter
	externalListRequests        prometheus.Counter
	externalRequestLatency      prometheus.Histogram
	eventsProcessed             prometheus.Counter
	eventStreamReconnections    prometheus.Counter
}

// NewPrometheusMetricsProvider creates a new PrometheusMetricsProvider and registers its metrics with the provided registry.
func NewPrometheusMetricsProvider(registry prometheus.Registerer) *PrometheusMetricsProvider {
	p := &PrometheusMetricsProvider{
		cacheHits: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "resource_cache_hits_total",
			Help: "The total number of cache hits",
		}),
		cacheMisses: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "resource_cache_misses_total",
			Help: "The total number of cache misses",
		}),
		cacheEvictions: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "resource_cache_evictions_total",
			Help: "The total number of cache evictions",
		}),
		cacheSize: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "resource_cache_size",
			Help: "The current number of items in the cache",
		}),
		cacheGetLatency: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "resource_cache_get_latency_seconds",
			Help:    "The latency of Get operations",
			Buckets: prometheus.DefBuckets,
		}),
		fullSyncDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "resource_cache_full_sync_duration_seconds",
			Help:    "The duration of full sync operations",
			Buckets: prometheus.DefBuckets,
		}),
		fullSyncFailures: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "resource_cache_full_sync_failures_total",
			Help: "The total number of failed full sync operations",
		}),
		timeSinceLastSuccessfulSync: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "resource_cache_time_since_last_successful_sync_seconds",
			Help: "The time elapsed since the last successful full sync",
		}),
		externalGetRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "resource_cache_external_get_requests_total",
			Help: "The total number of Get requests to the underlying data source",
		}),
		externalListRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "resource_cache_external_list_requests_total",
			Help: "The total number of List requests to the underlying data source",
		}),
		externalRequestLatency: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "resource_cache_external_request_latency_seconds",
			Help:    "The latency of requests to the underlying data source",
			Buckets: prometheus.DefBuckets,
		}),
		eventsProcessed: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "resource_cache_events_processed_total",
			Help: "The total number of events processed",
		}),
		eventStreamReconnections: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "resource_cache_event_stream_reconnections_total",
			Help: "The total number of times the event stream had to reconnect",
		}),
	}

	registry.MustRegister(
		p.cacheHits,
		p.cacheMisses,
		p.cacheEvictions,
		p.cacheSize,
		p.cacheGetLatency,
		p.fullSyncDuration,
		p.fullSyncFailures,
		p.timeSinceLastSuccessfulSync,
		p.externalGetRequests,
		p.externalListRequests,
		p.externalRequestLatency,
		p.eventsProcessed,
		p.eventStreamReconnections,
	)

	return p
}

// Implement all MetricsProvider methods for PrometheusMetricsProvider...
func (p *PrometheusMetricsProvider) IncrementCacheHits()      { p.cacheHits.Inc() }
func (p *PrometheusMetricsProvider) IncrementCacheMisses()    { p.cacheMisses.Inc() }
func (p *PrometheusMetricsProvider) IncrementCacheEvictions() { p.cacheEvictions.Inc() }
func (p *PrometheusMetricsProvider) SetCacheSize(size int)    { p.cacheSize.Set(float64(size)) }
func (p *PrometheusMetricsProvider) ObserveCacheGetLatency(latency time.Duration) {
	p.cacheGetLatency.Observe(latency.Seconds())
}
func (p *PrometheusMetricsProvider) ObserveFullSyncDuration(duration time.Duration) {
	p.fullSyncDuration.Observe(duration.Seconds())
}
func (p *PrometheusMetricsProvider) IncrementFullSyncFailures() { p.fullSyncFailures.Inc() }
func (p *PrometheusMetricsProvider) SetTimeSinceLastSuccessfulSync(duration time.Duration) {
	p.timeSinceLastSuccessfulSync.Set(duration.Seconds())
}
func (p *PrometheusMetricsProvider) IncrementExternalGetRequests()  { p.externalGetRequests.Inc() }
func (p *PrometheusMetricsProvider) IncrementExternalListRequests() { p.externalListRequests.Inc() }
func (p *PrometheusMetricsProvider) ObserveExternalRequestLatency(latency time.Duration) {
	p.externalRequestLatency.Observe(latency.Seconds())
}
func (p *PrometheusMetricsProvider) IncrementEventsProcessed() { p.eventsProcessed.Inc() }
func (p *PrometheusMetricsProvider) IncrementEventStreamReconnections() {
	p.eventStreamReconnections.Inc()
}

// NoopMetricsProvider implements MetricsProvider with no-op operations.
type NoopMetricsProvider struct{}

// NewNoopMetricsProvider creates a new NoopMetricsProvider.
func NewNoopMetricsProvider() *NoopMetricsProvider {
	return &NoopMetricsProvider{}
}

// Implement all MetricsProvider methods for NoopMetricsProvider...
func (*NoopMetricsProvider) IncrementCacheHits()                                   {}
func (*NoopMetricsProvider) IncrementCacheMisses()                                 {}
func (*NoopMetricsProvider) IncrementCacheEvictions()                              {}
func (*NoopMetricsProvider) SetCacheSize(size int)                                 {}
func (*NoopMetricsProvider) ObserveCacheGetLatency(latency time.Duration)          {}
func (*NoopMetricsProvider) ObserveFullSyncDuration(duration time.Duration)        {}
func (*NoopMetricsProvider) IncrementFullSyncFailures()                            {}
func (*NoopMetricsProvider) SetTimeSinceLastSuccessfulSync(duration time.Duration) {}
func (*NoopMetricsProvider) IncrementExternalGetRequests()                         {}
func (*NoopMetricsProvider) IncrementExternalListRequests()                        {}
func (*NoopMetricsProvider) ObserveExternalRequestLatency(latency time.Duration)   {}
func (*NoopMetricsProvider) IncrementEventsProcessed()                             {}
func (*NoopMetricsProvider) IncrementEventStreamReconnections()                    {}

var cacheMetricsProvider MetricsProvider = NewNoopMetricsProvider()

// SetMetricsProvider sets the global metrics provider.
func SetMetricsProvider(provider MetricsProvider) {
	cacheMetricsProvider = provider
}
