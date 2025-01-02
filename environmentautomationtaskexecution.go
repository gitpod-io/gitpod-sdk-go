// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
	"github.com/stainless-sdks/gitpod-go/shared"
	"github.com/tidwall/gjson"
)

// EnvironmentAutomationTaskExecutionService contains methods and other services
// that help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentAutomationTaskExecutionService] method instead.
type EnvironmentAutomationTaskExecutionService struct {
	Options []option.RequestOption
}

// NewEnvironmentAutomationTaskExecutionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewEnvironmentAutomationTaskExecutionService(opts ...option.RequestOption) (r *EnvironmentAutomationTaskExecutionService) {
	r = &EnvironmentAutomationTaskExecutionService{}
	r.Options = opts
	return
}

// GetTaskExecution
func (r *EnvironmentAutomationTaskExecutionService) Get(ctx context.Context, params EnvironmentAutomationTaskExecutionGetParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskExecutionGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/GetTaskExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// ListTaskExecutions
func (r *EnvironmentAutomationTaskExecutionService) List(ctx context.Context, params EnvironmentAutomationTaskExecutionListParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskExecutionListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/ListTaskExecutions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// ListTaskExecutions
func (r *EnvironmentAutomationTaskExecutionService) NewList(ctx context.Context, params EnvironmentAutomationTaskExecutionNewListParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskExecutionNewListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/ListTaskExecutions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetTaskExecution
func (r *EnvironmentAutomationTaskExecutionService) NewGet(ctx context.Context, params EnvironmentAutomationTaskExecutionNewGetParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskExecutionNewGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/GetTaskExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// StopTaskExecution
func (r *EnvironmentAutomationTaskExecutionService) Stop(ctx context.Context, params EnvironmentAutomationTaskExecutionStopParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskExecutionStopResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StopTaskExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateTaskExecutionStatus updates the status of a task execution. Only the
// environment executing a task execution is expected to call this function.
func (r *EnvironmentAutomationTaskExecutionService) UpdateTaskExecutionStatus(ctx context.Context, params EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/UpdateTaskExecutionStatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type EnvironmentAutomationTaskExecutionGetResponse struct {
	TaskExecution EnvironmentAutomationTaskExecutionGetResponseTaskExecution `json:"taskExecution"`
	JSON          environmentAutomationTaskExecutionGetResponseJSON          `json:"-"`
}

// environmentAutomationTaskExecutionGetResponseJSON contains the JSON metadata for
// the struct [EnvironmentAutomationTaskExecutionGetResponse]
type environmentAutomationTaskExecutionGetResponseJSON struct {
	TaskExecution apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionGetResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionGetResponseTaskExecution struct {
	ID       string                                                             `json:"id" format:"uuid"`
	Metadata EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadata `json:"metadata"`
	Spec     EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpec     `json:"spec"`
	Status   EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatus   `json:"status"`
	JSON     environmentAutomationTaskExecutionGetResponseTaskExecutionJSON     `json:"-"`
}

// environmentAutomationTaskExecutionGetResponseTaskExecutionJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecution]
type environmentAutomationTaskExecutionGetResponseTaskExecutionJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionGetResponseTaskExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionGetResponseTaskExecutionJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadata struct {
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator describes the principal who created/started the task run.
	Creator EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreator `json:"creator"`
	// environment_id is the ID of the environment in which the task run is executed.
	EnvironmentID string `json:"environmentId" format:"uuid"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// started_by describes the trigger that started the task execution.
	StartedBy string `json:"startedBy"`
	// task_id is the ID of the main task being executed.
	TaskID string                                                                 `json:"taskId" format:"uuid"`
	JSON   environmentAutomationTaskExecutionGetResponseTaskExecutionMetadataJSON `json:"-"`
}

// environmentAutomationTaskExecutionGetResponseTaskExecutionMetadataJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadata]
type environmentAutomationTaskExecutionGetResponseTaskExecutionMetadataJSON struct {
	CompletedAt   apijson.Field
	CreatedAt     apijson.Field
	Creator       apijson.Field
	EnvironmentID apijson.Field
	StartedAt     apijson.Field
	StartedBy     apijson.Field
	TaskID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionGetResponseTaskExecutionMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created/started the task run.
type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreator]
type environmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipal string

const (
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpec struct {
	// desired_phase is the phase the task execution should be in. Used to stop a
	// running task execution early.
	DesiredPhase EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhase `json:"desiredPhase"`
	// plan is a list of groups of steps. The steps in a group are executed
	// concurrently, while the groups are executed sequentially. The order of the
	// groups is the order in which they are executed.
	Plan []EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlan `json:"plan"`
	JSON environmentAutomationTaskExecutionGetResponseTaskExecutionSpecJSON   `json:"-"`
}

// environmentAutomationTaskExecutionGetResponseTaskExecutionSpecJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpec]
type environmentAutomationTaskExecutionGetResponseTaskExecutionSpecJSON struct {
	DesiredPhase apijson.Field
	Plan         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionGetResponseTaskExecutionSpecJSON) RawJSON() string {
	return r.raw
}

// desired_phase is the phase the task execution should be in. Used to stop a
// running task execution early.
type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhase string

const (
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlan struct {
	Steps []EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnion `json:"steps"`
	JSON  environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanJSON         `json:"-"`
}

// environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlan]
type environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnknown],
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnknown]
// or
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnknown].
type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnion interface {
	implementsEnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnion)(nil)).Elem(), "")
}

type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatus struct {
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhase `json:"phase"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their statuus has different versions. The value of this field has no
	// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
	// to impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStatusVersionUnion `json:"statusVersion"`
	// steps provides the status for each individual step of the task execution. If a
	// step is missing it has not yet started.
	Steps []EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStep `json:"steps"`
	JSON  environmentAutomationTaskExecutionGetResponseTaskExecutionStatusJSON   `json:"-"`
}

// environmentAutomationTaskExecutionGetResponseTaskExecutionStatusJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatus]
type environmentAutomationTaskExecutionGetResponseTaskExecutionStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	StatusVersion  apijson.Field
	Steps          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionGetResponseTaskExecutionStatusJSON) RawJSON() string {
	return r.raw
}

// the phase of a environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhase string

const (
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

// version of the status update. Environment instances themselves are unversioned,
// but their statuus has different versions. The value of this field has no
// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
// to impose a partial order. If a.status_version < b.status_version then a was the
// status before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStatusVersionUnion interface {
	ImplementsEnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStatusVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStatusVersionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage string `json:"failureMessage"`
	// the phase of a environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhase `json:"phase"`
	JSON  environmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepJSON   `json:"-"`
}

// environmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStep]
type environmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepJSON struct {
	ID             apijson.Field
	FailureMessage apijson.Field
	Phase          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepJSON) RawJSON() string {
	return r.raw
}

// the phase of a environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhase string

const (
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionListResponse struct {
	Pagination     EnvironmentAutomationTaskExecutionListResponsePagination      `json:"pagination"`
	TaskExecutions []EnvironmentAutomationTaskExecutionListResponseTaskExecution `json:"taskExecutions"`
	JSON           environmentAutomationTaskExecutionListResponseJSON            `json:"-"`
}

// environmentAutomationTaskExecutionListResponseJSON contains the JSON metadata
// for the struct [EnvironmentAutomationTaskExecutionListResponse]
type environmentAutomationTaskExecutionListResponseJSON struct {
	Pagination     apijson.Field
	TaskExecutions apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                                       `json:"nextToken"`
	JSON      environmentAutomationTaskExecutionListResponsePaginationJSON `json:"-"`
}

// environmentAutomationTaskExecutionListResponsePaginationJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponsePagination]
type environmentAutomationTaskExecutionListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionListResponseTaskExecution struct {
	ID       string                                                               `json:"id" format:"uuid"`
	Metadata EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadata `json:"metadata"`
	Spec     EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpec     `json:"spec"`
	Status   EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatus   `json:"status"`
	JSON     environmentAutomationTaskExecutionListResponseTaskExecutionJSON      `json:"-"`
}

// environmentAutomationTaskExecutionListResponseTaskExecutionJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseTaskExecution]
type environmentAutomationTaskExecutionListResponseTaskExecutionJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseTaskExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseTaskExecutionJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadata struct {
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator describes the principal who created/started the task run.
	Creator EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreator `json:"creator"`
	// environment_id is the ID of the environment in which the task run is executed.
	EnvironmentID string `json:"environmentId" format:"uuid"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// started_by describes the trigger that started the task execution.
	StartedBy string `json:"startedBy"`
	// task_id is the ID of the main task being executed.
	TaskID string                                                                   `json:"taskId" format:"uuid"`
	JSON   environmentAutomationTaskExecutionListResponseTaskExecutionsMetadataJSON `json:"-"`
}

// environmentAutomationTaskExecutionListResponseTaskExecutionsMetadataJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadata]
type environmentAutomationTaskExecutionListResponseTaskExecutionsMetadataJSON struct {
	CompletedAt   apijson.Field
	CreatedAt     apijson.Field
	Creator       apijson.Field
	EnvironmentID apijson.Field
	StartedAt     apijson.Field
	StartedBy     apijson.Field
	TaskID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseTaskExecutionsMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created/started the task run.
type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreator]
type environmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipal string

const (
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpec struct {
	// desired_phase is the phase the task execution should be in. Used to stop a
	// running task execution early.
	DesiredPhase EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhase `json:"desiredPhase"`
	// plan is a list of groups of steps. The steps in a group are executed
	// concurrently, while the groups are executed sequentially. The order of the
	// groups is the order in which they are executed.
	Plan []EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlan `json:"plan"`
	JSON environmentAutomationTaskExecutionListResponseTaskExecutionsSpecJSON   `json:"-"`
}

// environmentAutomationTaskExecutionListResponseTaskExecutionsSpecJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpec]
type environmentAutomationTaskExecutionListResponseTaskExecutionsSpecJSON struct {
	DesiredPhase apijson.Field
	Plan         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseTaskExecutionsSpecJSON) RawJSON() string {
	return r.raw
}

// desired_phase is the phase the task execution should be in. Used to stop a
// running task execution early.
type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhase string

const (
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlan struct {
	Steps []EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnion `json:"steps"`
	JSON  environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanJSON         `json:"-"`
}

// environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlan]
type environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnknown],
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnknown]
// or
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnknown].
type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnion interface {
	implementsEnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnion)(nil)).Elem(), "")
}

type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatus struct {
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhase `json:"phase"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their statuus has different versions. The value of this field has no
	// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
	// to impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStatusVersionUnion `json:"statusVersion"`
	// steps provides the status for each individual step of the task execution. If a
	// step is missing it has not yet started.
	Steps []EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStep `json:"steps"`
	JSON  environmentAutomationTaskExecutionListResponseTaskExecutionsStatusJSON   `json:"-"`
}

// environmentAutomationTaskExecutionListResponseTaskExecutionsStatusJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatus]
type environmentAutomationTaskExecutionListResponseTaskExecutionsStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	StatusVersion  apijson.Field
	Steps          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseTaskExecutionsStatusJSON) RawJSON() string {
	return r.raw
}

// the phase of a environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhase string

const (
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

// version of the status update. Environment instances themselves are unversioned,
// but their statuus has different versions. The value of this field has no
// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
// to impose a partial order. If a.status_version < b.status_version then a was the
// status before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStatusVersionUnion interface {
	ImplementsEnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStatusVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStatusVersionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage string `json:"failureMessage"`
	// the phase of a environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhase `json:"phase"`
	JSON  environmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepJSON   `json:"-"`
}

// environmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStep]
type environmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepJSON struct {
	ID             apijson.Field
	FailureMessage apijson.Field
	Phase          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepJSON) RawJSON() string {
	return r.raw
}

// the phase of a environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhase string

const (
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionNewListResponse struct {
	Pagination     EnvironmentAutomationTaskExecutionNewListResponsePagination      `json:"pagination"`
	TaskExecutions []EnvironmentAutomationTaskExecutionNewListResponseTaskExecution `json:"taskExecutions"`
	JSON           environmentAutomationTaskExecutionNewListResponseJSON            `json:"-"`
}

// environmentAutomationTaskExecutionNewListResponseJSON contains the JSON metadata
// for the struct [EnvironmentAutomationTaskExecutionNewListResponse]
type environmentAutomationTaskExecutionNewListResponseJSON struct {
	Pagination     apijson.Field
	TaskExecutions apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionNewListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                                          `json:"nextToken"`
	JSON      environmentAutomationTaskExecutionNewListResponsePaginationJSON `json:"-"`
}

// environmentAutomationTaskExecutionNewListResponsePaginationJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewListResponsePagination]
type environmentAutomationTaskExecutionNewListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionNewListResponseTaskExecution struct {
	ID       string                                                                  `json:"id" format:"uuid"`
	Metadata EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadata `json:"metadata"`
	Spec     EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpec     `json:"spec"`
	Status   EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatus   `json:"status"`
	JSON     environmentAutomationTaskExecutionNewListResponseTaskExecutionJSON      `json:"-"`
}

// environmentAutomationTaskExecutionNewListResponseTaskExecutionJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewListResponseTaskExecution]
type environmentAutomationTaskExecutionNewListResponseTaskExecutionJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewListResponseTaskExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewListResponseTaskExecutionJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadata struct {
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator describes the principal who created/started the task run.
	Creator EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreator `json:"creator"`
	// environment_id is the ID of the environment in which the task run is executed.
	EnvironmentID string `json:"environmentId" format:"uuid"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// started_by describes the trigger that started the task execution.
	StartedBy string `json:"startedBy"`
	// task_id is the ID of the main task being executed.
	TaskID string                                                                      `json:"taskId" format:"uuid"`
	JSON   environmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataJSON `json:"-"`
}

// environmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadata]
type environmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataJSON struct {
	CompletedAt   apijson.Field
	CreatedAt     apijson.Field
	Creator       apijson.Field
	EnvironmentID apijson.Field
	StartedAt     apijson.Field
	StartedBy     apijson.Field
	TaskID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created/started the task run.
type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreator]
type environmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipal string

const (
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpec struct {
	// desired_phase is the phase the task execution should be in. Used to stop a
	// running task execution early.
	DesiredPhase EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhase `json:"desiredPhase"`
	// plan is a list of groups of steps. The steps in a group are executed
	// concurrently, while the groups are executed sequentially. The order of the
	// groups is the order in which they are executed.
	Plan []EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlan `json:"plan"`
	JSON environmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecJSON   `json:"-"`
}

// environmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpec]
type environmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecJSON struct {
	DesiredPhase apijson.Field
	Plan         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecJSON) RawJSON() string {
	return r.raw
}

// desired_phase is the phase the task execution should be in. Used to stop a
// running task execution early.
type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhase string

const (
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecDesiredPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlan struct {
	Steps []EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlanStepsUnion `json:"steps"`
	JSON  environmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlanJSON         `json:"-"`
}

// environmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlanJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlan]
type environmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlanJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlanJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlanStepsUnknown],
// [EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlanStepsUnknown]
// or
// [EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlanStepsUnknown].
type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlanStepsUnion interface {
	implementsEnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlanStepsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsSpecPlanStepsUnion)(nil)).Elem(), "")
}

type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatus struct {
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhase `json:"phase"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their statuus has different versions. The value of this field has no
	// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
	// to impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStatusVersionUnion `json:"statusVersion"`
	// steps provides the status for each individual step of the task execution. If a
	// step is missing it has not yet started.
	Steps []EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStep `json:"steps"`
	JSON  environmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusJSON   `json:"-"`
}

// environmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatus]
type environmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	StatusVersion  apijson.Field
	Steps          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusJSON) RawJSON() string {
	return r.raw
}

// the phase of a environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhase string

const (
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

// version of the status update. Environment instances themselves are unversioned,
// but their statuus has different versions. The value of this field has no
// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
// to impose a partial order. If a.status_version < b.status_version then a was the
// status before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStatusVersionUnion interface {
	ImplementsEnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStatusVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStatusVersionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage string `json:"failureMessage"`
	// the phase of a environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhase `json:"phase"`
	JSON  environmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepJSON   `json:"-"`
}

// environmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStep]
type environmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepJSON struct {
	ID             apijson.Field
	FailureMessage apijson.Field
	Phase          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepJSON) RawJSON() string {
	return r.raw
}

// the phase of a environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhase string

const (
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionNewListResponseTaskExecutionsStatusStepsPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionNewGetResponse struct {
	TaskExecution EnvironmentAutomationTaskExecutionNewGetResponseTaskExecution `json:"taskExecution"`
	JSON          environmentAutomationTaskExecutionNewGetResponseJSON          `json:"-"`
}

// environmentAutomationTaskExecutionNewGetResponseJSON contains the JSON metadata
// for the struct [EnvironmentAutomationTaskExecutionNewGetResponse]
type environmentAutomationTaskExecutionNewGetResponseJSON struct {
	TaskExecution apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewGetResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecution struct {
	ID       string                                                                `json:"id" format:"uuid"`
	Metadata EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadata `json:"metadata"`
	Spec     EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpec     `json:"spec"`
	Status   EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatus   `json:"status"`
	JSON     environmentAutomationTaskExecutionNewGetResponseTaskExecutionJSON     `json:"-"`
}

// environmentAutomationTaskExecutionNewGetResponseTaskExecutionJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewGetResponseTaskExecution]
type environmentAutomationTaskExecutionNewGetResponseTaskExecutionJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewGetResponseTaskExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewGetResponseTaskExecutionJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadata struct {
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator describes the principal who created/started the task run.
	Creator EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreator `json:"creator"`
	// environment_id is the ID of the environment in which the task run is executed.
	EnvironmentID string `json:"environmentId" format:"uuid"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// started_by describes the trigger that started the task execution.
	StartedBy string `json:"startedBy"`
	// task_id is the ID of the main task being executed.
	TaskID string                                                                    `json:"taskId" format:"uuid"`
	JSON   environmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataJSON `json:"-"`
}

// environmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadata]
type environmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataJSON struct {
	CompletedAt   apijson.Field
	CreatedAt     apijson.Field
	Creator       apijson.Field
	EnvironmentID apijson.Field
	StartedAt     apijson.Field
	StartedBy     apijson.Field
	TaskID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created/started the task run.
type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreator]
type environmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipal string

const (
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpec struct {
	// desired_phase is the phase the task execution should be in. Used to stop a
	// running task execution early.
	DesiredPhase EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhase `json:"desiredPhase"`
	// plan is a list of groups of steps. The steps in a group are executed
	// concurrently, while the groups are executed sequentially. The order of the
	// groups is the order in which they are executed.
	Plan []EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlan `json:"plan"`
	JSON environmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecJSON   `json:"-"`
}

// environmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpec]
type environmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecJSON struct {
	DesiredPhase apijson.Field
	Plan         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecJSON) RawJSON() string {
	return r.raw
}

// desired_phase is the phase the task execution should be in. Used to stop a
// running task execution early.
type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhase string

const (
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlan struct {
	Steps []EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlanStepsUnion `json:"steps"`
	JSON  environmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlanJSON         `json:"-"`
}

// environmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlanJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlan]
type environmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlanJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlanJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlanStepsUnknown],
// [EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlanStepsUnknown]
// or
// [EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlanStepsUnknown].
type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlanStepsUnion interface {
	implementsEnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlanStepsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionSpecPlanStepsUnion)(nil)).Elem(), "")
}

type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatus struct {
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhase `json:"phase"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their statuus has different versions. The value of this field has no
	// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
	// to impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStatusVersionUnion `json:"statusVersion"`
	// steps provides the status for each individual step of the task execution. If a
	// step is missing it has not yet started.
	Steps []EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStep `json:"steps"`
	JSON  environmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusJSON   `json:"-"`
}

// environmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatus]
type environmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	StatusVersion  apijson.Field
	Steps          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusJSON) RawJSON() string {
	return r.raw
}

// the phase of a environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhase string

const (
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

// version of the status update. Environment instances themselves are unversioned,
// but their statuus has different versions. The value of this field has no
// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
// to impose a partial order. If a.status_version < b.status_version then a was the
// status before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStatusVersionUnion interface {
	ImplementsEnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStatusVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStatusVersionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage string `json:"failureMessage"`
	// the phase of a environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhase `json:"phase"`
	JSON  environmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepJSON   `json:"-"`
}

// environmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStep]
type environmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepJSON struct {
	ID             apijson.Field
	FailureMessage apijson.Field
	Phase          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepJSON) RawJSON() string {
	return r.raw
}

// the phase of a environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhase string

const (
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionNewGetResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionStopResponse = interface{}

type EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusResponse = interface{}

type EnvironmentAutomationTaskExecutionGetParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskExecutionGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Base64                 param.Field[string]                                                            `query:"base64"`
	Compression            param.Field[string]                                                            `query:"compression"`
	Connect                param.Field[string]                                                            `query:"connect"`
	Encoding               param.Field[string]                                                            `query:"encoding"`
	Message                param.Field[string]                                                            `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [EnvironmentAutomationTaskExecutionGetParams]'s query
// parameters as `url.Values`.
func (r EnvironmentAutomationTaskExecutionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskExecutionGetParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskExecutionGetParamsConnectProtocolVersion1 EnvironmentAutomationTaskExecutionGetParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskExecutionGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskExecutionListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Base64                 param.Field[string]                                                             `query:"base64"`
	Compression            param.Field[string]                                                             `query:"compression"`
	Connect                param.Field[string]                                                             `query:"connect"`
	Encoding               param.Field[string]                                                             `query:"encoding"`
	Message                param.Field[string]                                                             `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [EnvironmentAutomationTaskExecutionListParams]'s query
// parameters as `url.Values`.
func (r EnvironmentAutomationTaskExecutionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskExecutionListParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskExecutionListParamsConnectProtocolVersion1 EnvironmentAutomationTaskExecutionListParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskExecutionListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionNewListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskExecutionNewListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// filter contains the filter options for listing task runs
	Filter param.Field[EnvironmentAutomationTaskExecutionNewListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing task runs
	Pagination param.Field[EnvironmentAutomationTaskExecutionNewListParamsPagination] `json:"pagination"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationTaskExecutionNewListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskExecutionNewListParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskExecutionNewListParamsConnectProtocolVersion1 EnvironmentAutomationTaskExecutionNewListParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskExecutionNewListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionNewListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// filter contains the filter options for listing task runs
type EnvironmentAutomationTaskExecutionNewListParamsFilter struct {
	// environment_ids filters the response to only task runs of these environments
	EnvironmentIDs param.Field[[]string] `json:"environmentIds" format:"uuid"`
	// phases filters the response to only task runs in these phases
	Phases param.Field[[]EnvironmentAutomationTaskExecutionNewListParamsFilterPhase] `json:"phases"`
	// task_ids filters the response to only task runs of these tasks
	TaskIDs param.Field[[]string] `json:"taskIds" format:"uuid"`
	// task_references filters the response to only task runs with this reference
	TaskReferences param.Field[[]string] `json:"taskReferences"`
}

func (r EnvironmentAutomationTaskExecutionNewListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskExecutionNewListParamsFilterPhase string

const (
	EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionNewListParamsFilterPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionNewListParamsFilterPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionNewListParamsFilterPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionNewListParamsFilterPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionNewListParamsFilterPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionNewListParamsFilterPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionNewListParamsFilterPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

// pagination contains the pagination options for listing task runs
type EnvironmentAutomationTaskExecutionNewListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EnvironmentAutomationTaskExecutionNewListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskExecutionNewGetParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskExecutionNewGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                               `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationTaskExecutionNewGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskExecutionNewGetParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskExecutionNewGetParamsConnectProtocolVersion1 EnvironmentAutomationTaskExecutionNewGetParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskExecutionNewGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionNewGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionStopParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskExecutionStopParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                             `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationTaskExecutionStopParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskExecutionStopParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskExecutionStopParamsConnectProtocolVersion1 EnvironmentAutomationTaskExecutionStopParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskExecutionStopParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionStopParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParams struct {
	Body EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsBody `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsBody struct {
}

func (r EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsConnectProtocolVersion1 EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsConnectProtocolVersion1:
		return true
	}
	return false
}
