package client

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"connectrpc.com/connect"
)

const (
	AuthorizationHeader = "authorization"
	UserAgentHeader     = "user-agent"
	BearerPrefix        = "Bearer"
	PrincipalHeader     = "X-Gitpod-Principal"
)

func WithCustomUserAgent(userAgent string) connect.Interceptor {
	return withHeader(UserAgentHeader, userAgent)
}

func WithBearerToken(token string) connect.Interceptor {
	if strings.HasPrefix(token, BearerPrefix) {
		return withHeader(AuthorizationHeader, token)
	}

	return withHeader(AuthorizationHeader, fmt.Sprintf("%s %s", BearerPrefix, token))
}

func WithPrincipal(principal string) connect.Interceptor {
	return withHeader(PrincipalHeader, principal)
}

func withHeader(key, value string) connect.Interceptor {
	return &headerInterceptor{
		key:   key,
		value: value,
	}
}

type headerInterceptor struct {
	key   string
	value string
}

func (i *headerInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
		if ar.Spec().IsClient {
			setHeader(ar.Header(), i.key, i.value)
		}

		return next(ctx, ar)
	}
}

func (i *headerInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return connect.StreamingClientFunc(func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		conn := next(ctx, spec)
		setHeader(conn.RequestHeader(), i.key, i.value)
		return conn
	})
}

func (ic *headerInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	// We do not attach any credentials to the server side.
	return next
}

func setHeader(h http.Header, name, value string) {
	if value != "" && name != "" {
		h.Set(name, value)
	}
}

func NewClientLogInterceptor() connect.Interceptor {
	return &clientLogInterceptor{}
}

type clientLogInterceptor struct {
}

func (*clientLogInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

func (ic *clientLogInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

func (ic *clientLogInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
		// Check if it's a client RPC, otherwise we are in a server handler.
		if !ar.Spec().IsClient {
			return next(ctx, ar)
		}

		rpc := ar.Spec().Procedure
		slog.DebugContext(ctx, fmt.Sprintf("Sending RPC %s: %v", rpc, ar.Any()))

		resp, err := next(ctx, ar)

		if err != nil {
			code := connect.CodeOf(err)

			var exceptional bool
			switch code {
			case connect.CodeDataLoss, connect.CodeUnavailable, connect.CodeUnimplemented, connect.CodeInternal, connect.CodeUnknown:
				exceptional = true
			}

			if exceptional {
				slog.ErrorContext(ctx, fmt.Sprintf("RPC failed %s", rpc), slog.String("rpc.code", code.String()), slog.Any("err", err))
			} else {
				slog.DebugContext(ctx, fmt.Sprintf("RPC failed %s", rpc), slog.String("rpc.code", code.String()), slog.Any("err", err))
			}
		} else {
			slog.DebugContext(ctx, fmt.Sprintf("RPC succeeded %s", rpc), slog.String("rpc.code", "ok"))
		}

		return resp, err
	}
}

// NewDeadlineInterceptor returns an interceptor that adds a maximum deadline to the context if none is set or
// the current deadline is too far in the future.
// Only affects unary RPCs, not streaming RPCs.
func NewDeadlineInterceptor(maxDeadline time.Duration) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
			// Add a maximum deadline to the context if none is set or
			// the current deadline is too far in the future
			deadline, ok := ctx.Deadline()
			if !ok || time.Until(deadline) > maxDeadline {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, maxDeadline)
				defer cancel()
			}
			return next(ctx, ar)
		}
	}
}
