package client

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/prometheus/client_golang/prometheus"
)

// APICallMetrics holds Prometheus metrics for tracking API calls to the management plane.
type APICallMetrics struct {
	requestsTotal   *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}

// NewAPICallMetrics creates and registers API call metrics with the given registry.
// The returned value should be created once and shared across all clients via Interceptor().
func NewAPICallMetrics(registry prometheus.Registerer) *APICallMetrics {
	m := &APICallMetrics{
		requestsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "gitpod_management_plane_api_requests_completed_total",
			Help: "Total number of completed API requests to the management plane, by procedure and status code.",
		}, []string{"procedure", "code"}),
		requestDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "gitpod_management_plane_api_request_duration_seconds",
			Help:    "Duration of API requests to the management plane, by procedure and status code.",
			Buckets: []float64{0.01, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, 30},
		}, []string{"procedure", "code"}),
	}

	registry.MustRegister(m.requestsTotal, m.requestDuration)

	return m
}

// Interceptor returns a connect.Interceptor that records API call metrics.
func (m *APICallMetrics) Interceptor() connect.Interceptor {
	return &metricsInterceptor{metrics: m}
}

type metricsInterceptor struct {
	metrics *APICallMetrics
}

func (i *metricsInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
		if !ar.Spec().IsClient {
			return next(ctx, ar)
		}

		start := time.Now()
		resp, err := next(ctx, ar)
		duration := time.Since(start)

		procedure := ar.Spec().Procedure
		code := codeFromError(err)

		i.metrics.requestsTotal.WithLabelValues(procedure, code).Inc()
		i.metrics.requestDuration.WithLabelValues(procedure, code).Observe(duration.Seconds())

		return resp, err
	}
}

func (i *metricsInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		if !spec.IsClient {
			return next(ctx, spec)
		}

		// For streaming RPCs we count stream initiation only. Errors that occur
		// during Send/Receive are not captured here because StreamingClientConn
		// doesn't expose a single terminal error the way unary calls do.
		i.metrics.requestsTotal.WithLabelValues(spec.Procedure, "ok").Inc()

		return next(ctx, spec)
	}
}

func (i *metricsInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

func codeFromError(err error) string {
	if err == nil {
		return "ok"
	}
	return connect.CodeOf(err).String()
}
