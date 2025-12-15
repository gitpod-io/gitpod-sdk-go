// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
)

// ErrorService contains methods and other services that help with interacting with
// the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewErrorService] method instead.
type ErrorService struct {
	Options []option.RequestOption
}

// NewErrorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewErrorService(opts ...option.RequestOption) (r *ErrorService) {
	r = &ErrorService{}
	r.Options = opts
	return
}

// ReportErrors allows clients to report batches of errors that will be sent to
// error reporting systems. The structure is fully compatible with Sentry's event
// payload format.
//
// Use this method to:
//
// - Report client-side errors and exceptions
// - Track application crashes and panics
// - Send error context and metadata for debugging
//
// ### Examples
//
//   - Report a JavaScript error with Sentry-compatible structure: The service
//     accepts events with comprehensive error information including stack traces,
//     identity context, breadcrumbs, and metadata that align with Sentry's event
//     payload format.
func (r *ErrorService) ReportErrors(ctx context.Context, body ErrorReportErrorsParams, opts ...option.RequestOption) (res *ErrorReportErrorsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.ErrorsService/ReportErrors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Breadcrumb information (Sentry-compatible)
type BreadcrumbParam struct {
	// Breadcrumb category
	Category param.Field[string] `json:"category"`
	// Additional breadcrumb data
	Data param.Field[map[string]string] `json:"data"`
	// Breadcrumb level
	Level param.Field[ErrorLevel] `json:"level"`
	// Breadcrumb message
	Message param.Field[string] `json:"message"`
	// When the breadcrumb occurred
	Timestamp param.Field[time.Time] `json:"timestamp" format:"date-time"`
	// Breadcrumb type (e.g., "navigation", "http", "user", "error")
	Type param.Field[string] `json:"type"`
}

func (r BreadcrumbParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ErrorEvent contains comprehensive error information (Sentry-compatible)
type ErrorEventParam struct {
	// Breadcrumbs leading up to the error
	Breadcrumbs param.Field[[]BreadcrumbParam] `json:"breadcrumbs"`
	// Environment (e.g., "production", "staging", "development")
	Environment param.Field[string] `json:"environment"`
	// Unique event identifier (required by Sentry)
	EventID param.Field[string] `json:"eventId"`
	// Exception information (primary error data)
	Exceptions param.Field[[]ExceptionInfoParam] `json:"exceptions"`
	// Additional arbitrary metadata
	Extra param.Field[map[string]string] `json:"extra"`
	// Custom fingerprint for grouping
	Fingerprint param.Field[[]string] `json:"fingerprint"`
	// Identity ID of the user (UUID)
	IdentityID param.Field[string] `json:"identityId"`
	// Error severity level
	Level param.Field[ErrorLevel] `json:"level"`
	// Logger name
	Logger param.Field[string] `json:"logger"`
	// Modules/dependencies information
	Modules param.Field[map[string]string] `json:"modules"`
	// Platform identifier (required by Sentry)
	Platform param.Field[string] `json:"platform"`
	// Release version
	Release param.Field[string] `json:"release"`
	// Request information
	Request param.Field[RequestInfoParam] `json:"request"`
	// SDK information
	SDK param.Field[map[string]string] `json:"sdk"`
	// Server/host name
	ServerName param.Field[string] `json:"serverName"`
	// Tags for filtering and grouping
	Tags param.Field[map[string]string] `json:"tags"`
	// When the event occurred (required by Sentry)
	Timestamp param.Field[time.Time] `json:"timestamp" format:"date-time"`
	// Transaction name (e.g., route name, function name)
	Transaction param.Field[string] `json:"transaction"`
}

func (r ErrorEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Error severity levels (aligned with Sentry levels)
type ErrorLevel string

const (
	ErrorLevelUnspecified ErrorLevel = "ERROR_LEVEL_UNSPECIFIED"
	ErrorLevelDebug       ErrorLevel = "ERROR_LEVEL_DEBUG"
	ErrorLevelInfo        ErrorLevel = "ERROR_LEVEL_INFO"
	ErrorLevelWarning     ErrorLevel = "ERROR_LEVEL_WARNING"
	ErrorLevelError       ErrorLevel = "ERROR_LEVEL_ERROR"
	ErrorLevelFatal       ErrorLevel = "ERROR_LEVEL_FATAL"
)

func (r ErrorLevel) IsKnown() bool {
	switch r {
	case ErrorLevelUnspecified, ErrorLevelDebug, ErrorLevelInfo, ErrorLevelWarning, ErrorLevelError, ErrorLevelFatal:
		return true
	}
	return false
}

// Exception information (Sentry-compatible)
type ExceptionInfoParam struct {
	// Exception mechanism
	Mechanism param.Field[ExceptionMechanismParam] `json:"mechanism"`
	// Module or package where the exception type is defined
	Module param.Field[string] `json:"module"`
	// Stack trace frames
	Stacktrace param.Field[[]StackFrameParam] `json:"stacktrace"`
	// Thread ID if applicable
	ThreadID param.Field[string] `json:"threadId"`
	// Exception type/class name
	Type param.Field[string] `json:"type"`
	// Exception message/value
	Value param.Field[string] `json:"value"`
}

func (r ExceptionInfoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Exception mechanism information (Sentry-compatible)
type ExceptionMechanismParam struct {
	// Additional mechanism-specific data
	Data param.Field[map[string]string] `json:"data"`
	// Human-readable description of the mechanism
	Description param.Field[string] `json:"description"`
	// Whether the exception was handled by user code
	Handled param.Field[bool] `json:"handled"`
	// Whether this is a synthetic exception (created by SDK)
	Synthetic param.Field[bool] `json:"synthetic"`
	// Type of mechanism (e.g., "generic", "promise", "onerror")
	Type param.Field[string] `json:"type"`
}

func (r ExceptionMechanismParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Request information (Sentry-compatible)
type RequestInfoParam struct {
	// Request body (truncated if large)
	Data param.Field[string] `json:"data"`
	// Request headers
	Headers param.Field[map[string]string] `json:"headers"`
	// HTTP method
	Method param.Field[string] `json:"method"`
	// Query parameters
	QueryString param.Field[map[string]string] `json:"queryString"`
	// Request URL
	URL param.Field[string] `json:"url"`
}

func (r RequestInfoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Stack trace frame information (Sentry-compatible)
type StackFrameParam struct {
	// Column number in the line
	Colno       param.Field[int64]  `json:"colno"`
	ContextLine param.Field[string] `json:"contextLine"`
	// File name or path
	Filename param.Field[string] `json:"filename"`
	// Function name
	Function param.Field[string] `json:"function"`
	// Whether this frame is in application code (vs library/framework code)
	InApp param.Field[bool] `json:"inApp"`
	// Line number in the file
	Lineno param.Field[int64] `json:"lineno"`
	// Module or package name
	Module      param.Field[string]   `json:"module"`
	PostContext param.Field[[]string] `json:"postContext"`
	// Source code context around the error line
	PreContext param.Field[[]string] `json:"preContext"`
	// Additional frame-specific variables/locals
	Vars param.Field[map[string]string] `json:"vars"`
}

func (r StackFrameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ErrorReportErrorsResponse = interface{}

type ErrorReportErrorsParams struct {
	// Error events to be reported (batch) - now using Sentry-compatible structure
	Events param.Field[[]ErrorEventParam] `json:"events"`
}

func (r ErrorReportErrorsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
