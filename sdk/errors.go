package sdk

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"connectrpc.com/connect"
)

var (
	errSDKClientMissing = errors.New("sdk requires a management-plane client")

	// ErrMissingAPIKey is returned when neither supported API key environment variable is set.
	ErrMissingAPIKey = errors.New("ONA_API_KEY or GITPOD_API_KEY is required")
)

type operationError struct {
	operation string
	err       error
}

func (e operationError) Error() string {
	if e.operation == "" {
		return e.err.Error()
	}
	return fmt.Sprintf("%s: %v", e.operation, e.err)
}

func (e operationError) Unwrap() error {
	return e.err
}

type AuthenticationRequiredError struct{ operationError }
type PermissionDeniedError struct{ operationError }
type NotFoundError struct{ operationError }
type RateLimitedError struct{ operationError }
type UnavailableError struct{ operationError }
type DeadlineExceededError struct{ operationError }
type ValidationError struct{ operationError }
type CapabilityUnavailableError struct{ operationError }
type EnvironmentPolicyError struct{ operationError }
type EnvironmentUnreachableError struct{ operationError }
type SDKError struct{ operationError }

// MapError converts common Connect and context errors into SDK error types while preserving Unwrap.
func MapError(operation string, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, context.Canceled) {
		return err
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return &DeadlineExceededError{operationError{operation: operation, err: err}}
	}

	message := strings.ToLower(err.Error())
	switch connect.CodeOf(err) {
	case connect.CodeUnauthenticated:
		return &AuthenticationRequiredError{operationError{operation: operation, err: err}}
	case connect.CodePermissionDenied:
		return &PermissionDeniedError{operationError{operation: operation, err: err}}
	case connect.CodeNotFound:
		return &NotFoundError{operationError{operation: operation, err: err}}
	case connect.CodeResourceExhausted:
		return &RateLimitedError{operationError{operation: operation, err: err}}
	case connect.CodeUnavailable:
		return &UnavailableError{operationError{operation: operation, err: err}}
	case connect.CodeDeadlineExceeded:
		return &DeadlineExceededError{operationError{operation: operation, err: err}}
	case connect.CodeInvalidArgument:
		return &ValidationError{operationError{operation: operation, err: err}}
	case connect.CodeFailedPrecondition:
		if strings.Contains(message, "capabil") || strings.Contains(message, "unsupported") {
			return &CapabilityUnavailableError{operationError{operation: operation, err: err}}
		}
		return &EnvironmentPolicyError{operationError{operation: operation, err: err}}
	case connect.CodeUnimplemented:
		return &CapabilityUnavailableError{operationError{operation: operation, err: err}}
	default:
		return &SDKError{operationError{operation: operation, err: err}}
	}
}
