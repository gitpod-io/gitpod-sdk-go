package client

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"connectrpc.com/connect"
	"golang.org/x/oauth2"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/proto"

	"github.com/gitpod-io/gitpod-sdk-go/internal/raw/tokenproof"
)

const (
	AuthorizationHeader     = "authorization"
	PortAuthorizationHeader = "X-Ona-Authorization"
	UserAgentHeader         = "user-agent"
	BearerPrefix            = "Bearer"
	PrincipalHeader         = "X-Gitpod-Principal"

	// Rate limit headers returned by the server
	HeaderRateLimitRetryAfter = "RateLimit-RetryAfter"
)

func WithCustomUserAgent(userAgent string) connect.Interceptor {
	return withStaticHeaderValue(UserAgentHeader, userAgent)
}

func withOnaUserAgent(userAgentPrefix string) connect.Interceptor {
	return withHeaderValueFn(UserAgentHeader, func(ctx context.Context) string {
		return onaUserAgent(userAgentPrefix, userAgentLayerFromContext(ctx))
	})
}

func WithBearerToken(token string) connect.Interceptor {
	withoutBearer := strings.TrimSpace(strings.TrimPrefix(token, BearerPrefix))

	return TokenSourceInterceptor(oauth2.StaticTokenSource(&oauth2.Token{AccessToken: withoutBearer}))
}

// WithPortAccessToken attaches a token for runner-proxy port authentication.
func WithPortAccessToken(token string) connect.Interceptor {
	return withStaticHeaderValue(PortAuthorizationHeader, token)
}

func TokenSourceInterceptor(source oauth2.TokenSource) connect.Interceptor {
	return withHeaderValueFn(AuthorizationHeader, func(ctx context.Context) string {
		if source == nil {
			slog.DebugContext(ctx, "Token source is nil, no credentials will be attached to requests")
			return ""
		}

		token, err := source.Token()
		if err != nil {
			slog.WarnContext(ctx, "Failed to get token from token source", slog.Any("err", err))
			return ""
		}

		if token.AccessToken == "" {
			slog.WarnContext(ctx, "Token source returned an empty token", slog.Any("token", token))
			return ""
		}

		return fmt.Sprintf("%s %s", BearerPrefix, token.AccessToken)
	})
}

func TokenSourceWithContextInterceptor(source TokenSourceWithContext) connect.Interceptor {
	return withHeaderValueFn(AuthorizationHeader, func(ctx context.Context) string {
		if source == nil {
			slog.DebugContext(ctx, "Token source is nil, no credentials will be attached to requests")
			return ""
		}

		token, err := source.Token(ctx)
		if err != nil {
			slog.WarnContext(ctx, "Failed to get token from token source", slog.Any("err", err))
			return ""
		}

		if token.AccessToken == "" {
			slog.WarnContext(ctx, "Token source returned an empty token", slog.Any("token", token))
			return ""
		}

		return fmt.Sprintf("%s %s", BearerPrefix, token.AccessToken)
	})
}

func TokenSourceWithContextAndProofInterceptor(source TokenSourceWithContext, proofProvider TokenProofProvider) connect.Interceptor {
	return &tokenSourceWithContextAndProofInterceptor{
		source:        source,
		proofProvider: proofProvider,
	}
}

type TokenSourceWithContext interface {
	Token(ctx context.Context) (*oauth2.Token, error)
}

type TokenProofProvider interface {
	Proof(ctx context.Context, procedure string, accessToken string) (string, error)
}

func WithPrincipal(principal string) connect.Interceptor {
	return withStaticHeaderValue(PrincipalHeader, principal)
}

func withStaticHeaderValue(key, value string) connect.Interceptor {
	return &headerInterceptor{
		key:   key,
		value: value,
	}
}

func withHeaderValueFn(key string, valueFn func(ctx context.Context) string) connect.Interceptor {
	return &headerInterceptor{
		key:     key,
		valueFn: valueFn,
	}
}

type headerInterceptor struct {
	key   string
	value string
	// If set, valueFn is used to generate the value for the header instead of value
	valueFn func(ctx context.Context) string
}

type tokenSourceWithContextAndProofInterceptor struct {
	source        TokenSourceWithContext
	proofProvider TokenProofProvider
}

func (i *tokenSourceWithContextAndProofInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
		if ar.Spec().IsClient {
			authorization, proof := i.eval(ctx, ar.Spec().Procedure)
			setHeader(ar.Header(), AuthorizationHeader, authorization)
			setHeader(ar.Header(), tokenproof.HeaderName, proof)
		}

		return next(ctx, ar)
	}
}

func (i *tokenSourceWithContextAndProofInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return connect.StreamingClientFunc(func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		conn := next(ctx, spec)
		authorization, proof := i.eval(ctx, spec.Procedure)
		setHeader(conn.RequestHeader(), AuthorizationHeader, authorization)
		setHeader(conn.RequestHeader(), tokenproof.HeaderName, proof)
		return conn
	})
}

func (i *tokenSourceWithContextAndProofInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

func (i *tokenSourceWithContextAndProofInterceptor) eval(ctx context.Context, procedure string) (authorization string, proof string) {
	if i.source == nil {
		slog.DebugContext(ctx, "Token source is nil, no credentials will be attached to requests")
		return "", ""
	}

	token, err := i.source.Token(ctx)
	if err != nil {
		slog.WarnContext(ctx, "Failed to get token from token source", slog.Any("err", err))
		return "", ""
	}

	if token.AccessToken == "" {
		slog.WarnContext(ctx, "Token source returned an empty token", slog.Any("token", token))
		return "", ""
	}

	if i.proofProvider != nil {
		proof, err = i.proofProvider.Proof(ctx, procedure, token.AccessToken)
		if err != nil {
			slog.WarnContext(ctx, "Failed to attach cnf proof", slog.String("procedure", procedure), slog.Any("err", err))
		}
	}

	return fmt.Sprintf("%s %s", BearerPrefix, token.AccessToken), proof
}

func (i *headerInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
		if ar.Spec().IsClient {
			value := i.eval(ctx)
			setHeader(ar.Header(), i.key, value)
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
		value := i.eval(ctx)
		setHeader(conn.RequestHeader(), i.key, value)
		return conn
	})
}

func (i *headerInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	// We do not attach any credentials to the server side.
	return next
}

func (i *headerInterceptor) eval(ctx context.Context) string {
	if i.valueFn != nil {
		return i.valueFn(ctx)
	}
	return i.value
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
	return connect.StreamingClientFunc(func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		// Check if it's a client RPC, otherwise we are in a server handler.
		if !spec.IsClient {
			return next(ctx, spec)
		}

		rpc := spec.Procedure
		slog.DebugContext(ctx, fmt.Sprintf("Sending streaming RPC %s", rpc))

		return next(ctx, spec)
	})
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

// RateLimitRetryConfig configures the rate limit retry interceptor behavior.
type RateLimitRetryConfig struct {
	// MaxRetries is the maximum number of retry attempts for rate-limited requests.
	// Default is 10 if not set.
	MaxRetries int

	// MaxRetryDelay caps the retry delay to prevent excessively long waits.
	// Default is 60 seconds if not set.
	MaxRetryDelay time.Duration
}

// NewRateLimitRetryInterceptor returns an interceptor that automatically retries
// requests that fail due to rate limiting (ResourceExhausted errors).
// It parses the retry delay from the RetryInfo error detail or RateLimit-RetryAfter header.
// Only affects unary RPCs, not streaming RPCs.
func NewRateLimitRetryInterceptor(cfg RateLimitRetryConfig) connect.UnaryInterceptorFunc {
	if cfg.MaxRetries <= 0 {
		cfg.MaxRetries = 10
	}
	if cfg.MaxRetryDelay <= 0 {
		cfg.MaxRetryDelay = 60 * time.Second
	}

	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
			var lastErr error
			for attempt := 0; attempt <= cfg.MaxRetries; attempt++ {
				resp, err := next(ctx, ar)
				if err == nil {
					return resp, nil
				}

				// Check if this is a rate limit error
				if !isRateLimitError(err) {
					return resp, err
				}

				// Don't retry if we've exhausted attempts
				if attempt >= cfg.MaxRetries {
					lastErr = err
					break
				}

				// Extract retry delay from the error
				retryAfter := extractRetryDelay(err)
				if retryAfter <= 0 {
					// Default to 1 second if no delay specified
					retryAfter = time.Second
				}
				if retryAfter > cfg.MaxRetryDelay {
					retryAfter = cfg.MaxRetryDelay
				}

				rpc := ar.Spec().Procedure
				slog.InfoContext(ctx, "Rate limited, retrying after delay",
					slog.String("rpc", rpc),
					slog.Int("attempt", attempt+1),
					slog.Int("max_retries", cfg.MaxRetries),
					slog.Duration("retry_after", retryAfter),
				)

				// Wait for the retry delay or context cancellation
				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				case <-time.After(retryAfter):
					// Continue to next attempt
				}

				lastErr = err
			}

			return nil, lastErr
		}
	}
}

// TransientRetryConfig configures the transient error retry interceptor behavior.
type TransientRetryConfig struct {
	// MaxRetries is the maximum number of retry attempts for transient errors.
	// Default is 3 if not set.
	MaxRetries int

	// InitialBackoff is the delay before the first retry.
	// Default is 500ms if not set.
	InitialBackoff time.Duration

	// MaxBackoff caps the backoff delay.
	// Default is 5 seconds if not set.
	MaxBackoff time.Duration
}

// NewTransientRetryInterceptor returns an interceptor that automatically retries
// unary RPCs that fail with transient errors (CodeUnavailable).
// Uses exponential backoff (2x) between retries.
func NewTransientRetryInterceptor(cfg TransientRetryConfig) connect.UnaryInterceptorFunc {
	if cfg.MaxRetries <= 0 {
		cfg.MaxRetries = 3
	}
	if cfg.InitialBackoff <= 0 {
		cfg.InitialBackoff = 500 * time.Millisecond
	}
	if cfg.MaxBackoff <= 0 {
		cfg.MaxBackoff = 5 * time.Second
	}

	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, ar)
			if err == nil || !isTransientError(err) {
				return resp, err
			}

			backoff := cfg.InitialBackoff
			for attempt := 1; attempt <= cfg.MaxRetries; attempt++ {
				rpc := ar.Spec().Procedure
				slog.WarnContext(ctx, "Transient RPC error, retrying",
					slog.String("rpc", rpc),
					slog.Int("attempt", attempt),
					slog.Int("max_retries", cfg.MaxRetries),
					slog.Duration("backoff", backoff),
					slog.Any("error", err),
				)

				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				case <-time.After(backoff):
				}

				resp, err = next(ctx, ar)
				if err == nil || !isTransientError(err) {
					return resp, err
				}

				backoff *= 2
				if backoff > cfg.MaxBackoff {
					backoff = cfg.MaxBackoff
				}
			}

			return resp, err
		}
	}
}

// isTransientError checks if the error is a transient network error that should be retried.
func isTransientError(err error) bool {
	var connectErr *connect.Error
	if errors.As(err, &connectErr) {
		return connectErr.Code() == connect.CodeUnavailable
	}
	return false
}

// isRateLimitError checks if the error is a rate limit (ResourceExhausted) error.
func isRateLimitError(err error) bool {
	var connectErr *connect.Error
	if errors.As(err, &connectErr) {
		return connectErr.Code() == connect.CodeResourceExhausted
	}
	return false
}

// extractRetryDelay extracts the retry delay from a rate limit error.
// It first checks the RetryInfo error detail, then falls back to the RateLimit-RetryAfter header.
func extractRetryDelay(err error) time.Duration {
	var connectErr *connect.Error
	if !errors.As(err, &connectErr) {
		return 0
	}

	// First, try to get the delay from RetryInfo error detail
	for _, detail := range connectErr.Details() {
		msg, valueErr := detail.Value()
		if valueErr != nil {
			continue
		}
		if retryInfo, ok := msg.(*errdetails.RetryInfo); ok {
			if retryInfo.RetryDelay != nil {
				return retryInfo.RetryDelay.AsDuration()
			}
		}
		// Also try to unmarshal as RetryInfo if the type assertion fails
		retryInfo := new(errdetails.RetryInfo)
		if proto.Unmarshal(detail.Bytes(), retryInfo) == nil && retryInfo.RetryDelay != nil {
			return retryInfo.RetryDelay.AsDuration()
		}
	}

	// Fall back to the header
	if retryAfterStr := connectErr.Meta().Get(HeaderRateLimitRetryAfter); retryAfterStr != "" {
		if seconds, parseErr := strconv.Atoi(retryAfterStr); parseErr == nil && seconds > 0 {
			return time.Duration(seconds) * time.Second
		}
	}

	return 0
}
