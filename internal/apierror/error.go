package apierror

import "github.com/gitpod-io/gitpod-sdk-go/shared"

type ErrorCode = shared.ErrorCode

const ErrorCodeCanceled = shared.ErrorCodeCanceled
const ErrorCodeUnknown = shared.ErrorCodeUnknown
const ErrorCodeInvalidArgument = shared.ErrorCodeInvalidArgument
const ErrorCodeDeadlineExceeded = shared.ErrorCodeDeadlineExceeded
const ErrorCodeNotFound = shared.ErrorCodeNotFound
const ErrorCodeAlreadyExists = shared.ErrorCodeAlreadyExists
const ErrorCodePermissionDenied = shared.ErrorCodePermissionDenied
const ErrorCodeResourceExhausted = shared.ErrorCodeResourceExhausted
const ErrorCodeFailedPrecondition = shared.ErrorCodeFailedPrecondition
const ErrorCodeAborted = shared.ErrorCodeAborted
const ErrorCodeOutOfRange = shared.ErrorCodeOutOfRange
const ErrorCodeUnimplemented = shared.ErrorCodeUnimplemented
const ErrorCodeInternal = shared.ErrorCodeInternal
const ErrorCodeUnavailable = shared.ErrorCodeUnavailable
const ErrorCodeDataLoss = shared.ErrorCodeDataLoss
const ErrorCodeUnauthenticated = shared.ErrorCodeUnauthenticated
