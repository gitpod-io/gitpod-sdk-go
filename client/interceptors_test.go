package client

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"

	v1 "github.com/gitpod-io/flex-sdk-go/v1"
)

func TestNewDeadlineInterceptor(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		req  connect.AnyRequest
		next func(t *testing.T, ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error)
	}{
		{
			name: "sets max 30s deadline if no deadline is set",
			ctx:  context.Background(),
			req:  connect.NewRequest(&v1.ListServicesRequest{}),
			next: func(t *testing.T, ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
				deadline, ok := ctx.Deadline()
				if !ok {
					t.Fatal("expected deadline to be set")
				}
				if time.Until(deadline) > 30*time.Second {
					t.Fatalf("expected deadline to be less than or equal to 30s, got %v", time.Until(deadline))
				}
				return nil, nil
			},
		},
		{
			name: "does not change deadline if less than 30s",
			ctx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				return ctx
			}(),
			req: connect.NewRequest(&v1.ListServicesRequest{}),
			next: func(t *testing.T, ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
				deadline, ok := ctx.Deadline()
				if !ok {
					t.Fatal("expected deadline to be set")
				}
				if time.Until(deadline) > 10*time.Second {
					t.Fatalf("expected deadline to be less than or equal to 10s, got %v", time.Until(deadline))
				}
				return nil, nil
			},
		},
		{
			name: "sets max 30s deadline if more than 30s",
			ctx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
				defer cancel()
				return ctx
			}(),
			req: connect.NewRequest(&v1.ListServicesRequest{}),
			next: func(t *testing.T, ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
				deadline, ok := ctx.Deadline()
				if !ok {
					t.Fatal("expected deadline to be set")
				}
				if time.Until(deadline) > 30*time.Second {
					t.Fatalf("expected deadline to be less than or equal to 30s, got %v", time.Until(deadline))
				}
				return nil, nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interceptor := NewDeadlineInterceptor(30 * time.Second)
			_, err := interceptor(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
				return tt.next(t, ctx, request)
			})(tt.ctx, tt.req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
