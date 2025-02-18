package shared

type ErrorCode string

const (
	ErrorCodeCanceled           ErrorCode = "canceled"
	ErrorCodeUnknown            ErrorCode = "unknown"
	ErrorCodeInvalidArgument    ErrorCode = "invalid_argument"
	ErrorCodeDeadlineExceeded   ErrorCode = "deadline_exceeded"
	ErrorCodeNotFound           ErrorCode = "not_found"
	ErrorCodeAlreadyExists      ErrorCode = "already_exists"
	ErrorCodePermissionDenied   ErrorCode = "permission_denied"
	ErrorCodeResourceExhausted  ErrorCode = "resource_exhausted"
	ErrorCodeFailedPrecondition ErrorCode = "failed_precondition"
	ErrorCodeAborted            ErrorCode = "aborted"
	ErrorCodeOutOfRange         ErrorCode = "out_of_range"
	ErrorCodeUnimplemented      ErrorCode = "unimplemented"
	ErrorCodeInternal           ErrorCode = "internal"
	ErrorCodeUnavailable        ErrorCode = "unavailable"
	ErrorCodeDataLoss           ErrorCode = "data_loss"
	ErrorCodeUnauthenticated    ErrorCode = "unauthenticated"
)
