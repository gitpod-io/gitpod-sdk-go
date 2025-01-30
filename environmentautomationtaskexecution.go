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
	//
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
	//
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
	//
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
	// concurrently, while the groups are executed sequentially.
	//
	// The order of the groups is the order in which they are executed.
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
	Steps []EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStep `json:"steps"`
	JSON  environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanJSON   `json:"-"`
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

type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// This field can have the runtime type of [[]string].
	DependsOn interface{} `json:"dependsOn"`
	Label     string      `json:"label"`
	ServiceID string      `json:"serviceId" format:"uuid"`
	// This field can have the runtime type of
	// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObjectTask].
	Task  interface{}                                                                `json:"task"`
	JSON  environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepJSON `json:"-"`
	union EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnion
}

// environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStep]
type environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepJSON struct {
	ID          apijson.Field
	DependsOn   apijson.Field
	Label       apijson.Field
	ServiceID   apijson.Field
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStep) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStep{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObject],
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObject].
func (r EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStep) AsUnion() EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnion {
	return r.union
}

// Union satisfied by
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObject]
// or
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObject].
type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnion interface {
	implementsEnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStep()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObject{}),
		},
	)
}

type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObject struct {
	ServiceID string `json:"serviceId,required" format:"uuid"`
	// ID is the ID of the execution step
	ID        string                                                                            `json:"id" format:"uuid"`
	DependsOn []string                                                                          `json:"dependsOn"`
	Label     string                                                                            `json:"label"`
	JSON      environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObjectJSON `json:"-"`
}

// environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObjectJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObject]
type environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObjectJSON struct {
	ServiceID   apijson.Field
	ID          apijson.Field
	DependsOn   apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStepsObject) implementsEnvironmentAutomationTaskExecutionGetResponseTaskExecutionSpecPlanStep() {
}

type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatus struct {
	// failure_message summarises why the task execution failed to operate. If this is
	// non-empty
	//
	// the task execution has failed to operate and will likely transition to a failed
	// state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs
	//
	// or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a task execution represents the aggregated phase of all steps.
	Phase EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhase `json:"phase"`
	// version of the status update. Task executions themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion" format:"int64"`
	// steps provides the status for each individual step of the task execution. If a
	// step is missing it
	//
	// has not yet started.
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

// the phase of a task execution represents the aggregated phase of all steps.
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

type EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the step failed to operate. If this is non-empty
	//
	// the step has failed to operate and will likely transition to a failed state.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the execution step
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

// phase is the current phase of the execution step
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
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
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
	//
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
	//
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
	//
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
	// concurrently, while the groups are executed sequentially.
	//
	// The order of the groups is the order in which they are executed.
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
	Steps []EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStep `json:"steps"`
	JSON  environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanJSON   `json:"-"`
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

type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// This field can have the runtime type of [[]string].
	DependsOn interface{} `json:"dependsOn"`
	Label     string      `json:"label"`
	ServiceID string      `json:"serviceId" format:"uuid"`
	// This field can have the runtime type of
	// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObjectTask].
	Task  interface{}                                                                  `json:"task"`
	JSON  environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepJSON `json:"-"`
	union EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnion
}

// environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStep]
type environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepJSON struct {
	ID          apijson.Field
	DependsOn   apijson.Field
	Label       apijson.Field
	ServiceID   apijson.Field
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStep) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStep{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObject],
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObject].
func (r EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStep) AsUnion() EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnion {
	return r.union
}

// Union satisfied by
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObject]
// or
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObject].
type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnion interface {
	implementsEnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStep()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObject{}),
		},
	)
}

type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObject struct {
	ServiceID string `json:"serviceId,required" format:"uuid"`
	// ID is the ID of the execution step
	ID        string                                                                              `json:"id" format:"uuid"`
	DependsOn []string                                                                            `json:"dependsOn"`
	Label     string                                                                              `json:"label"`
	JSON      environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObjectJSON `json:"-"`
}

// environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObjectJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObject]
type environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObjectJSON struct {
	ServiceID   apijson.Field
	ID          apijson.Field
	DependsOn   apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStepsObject) implementsEnvironmentAutomationTaskExecutionListResponseTaskExecutionsSpecPlanStep() {
}

type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatus struct {
	// failure_message summarises why the task execution failed to operate. If this is
	// non-empty
	//
	// the task execution has failed to operate and will likely transition to a failed
	// state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs
	//
	// or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a task execution represents the aggregated phase of all steps.
	Phase EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusPhase `json:"phase"`
	// version of the status update. Task executions themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion" format:"int64"`
	// steps provides the status for each individual step of the task execution. If a
	// step is missing it
	//
	// has not yet started.
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

// the phase of a task execution represents the aggregated phase of all steps.
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

type EnvironmentAutomationTaskExecutionListResponseTaskExecutionsStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the step failed to operate. If this is non-empty
	//
	// the step has failed to operate and will likely transition to a failed state.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the execution step
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

// phase is the current phase of the execution step
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

type EnvironmentAutomationTaskExecutionStopResponse = interface{}

type EnvironmentAutomationTaskExecutionGetParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[EnvironmentAutomationTaskExecutionGetParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskExecutionGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[EnvironmentAutomationTaskExecutionGetParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[EnvironmentAutomationTaskExecutionGetParamsConnect] `query:"connect"`
	Message param.Field[string]                                             `query:"message"`
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

// Define which encoding or 'Message-Codec' to use
type EnvironmentAutomationTaskExecutionGetParamsEncoding string

const (
	EnvironmentAutomationTaskExecutionGetParamsEncodingProto EnvironmentAutomationTaskExecutionGetParamsEncoding = "proto"
	EnvironmentAutomationTaskExecutionGetParamsEncodingJson  EnvironmentAutomationTaskExecutionGetParamsEncoding = "json"
)

func (r EnvironmentAutomationTaskExecutionGetParamsEncoding) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionGetParamsEncodingProto, EnvironmentAutomationTaskExecutionGetParamsEncodingJson:
		return true
	}
	return false
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

// Which compression algorithm to use for this request
type EnvironmentAutomationTaskExecutionGetParamsCompression string

const (
	EnvironmentAutomationTaskExecutionGetParamsCompressionIdentity EnvironmentAutomationTaskExecutionGetParamsCompression = "identity"
	EnvironmentAutomationTaskExecutionGetParamsCompressionGzip     EnvironmentAutomationTaskExecutionGetParamsCompression = "gzip"
	EnvironmentAutomationTaskExecutionGetParamsCompressionBr       EnvironmentAutomationTaskExecutionGetParamsCompression = "br"
)

func (r EnvironmentAutomationTaskExecutionGetParamsCompression) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionGetParamsCompressionIdentity, EnvironmentAutomationTaskExecutionGetParamsCompressionGzip, EnvironmentAutomationTaskExecutionGetParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskExecutionGetParamsConnect string

const (
	EnvironmentAutomationTaskExecutionGetParamsConnectV1 EnvironmentAutomationTaskExecutionGetParamsConnect = "v1"
)

func (r EnvironmentAutomationTaskExecutionGetParamsConnect) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionGetParamsConnectV1:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[EnvironmentAutomationTaskExecutionListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskExecutionListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[EnvironmentAutomationTaskExecutionListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[EnvironmentAutomationTaskExecutionListParamsConnect] `query:"connect"`
	Message param.Field[string]                                              `query:"message"`
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

// Define which encoding or 'Message-Codec' to use
type EnvironmentAutomationTaskExecutionListParamsEncoding string

const (
	EnvironmentAutomationTaskExecutionListParamsEncodingProto EnvironmentAutomationTaskExecutionListParamsEncoding = "proto"
	EnvironmentAutomationTaskExecutionListParamsEncodingJson  EnvironmentAutomationTaskExecutionListParamsEncoding = "json"
)

func (r EnvironmentAutomationTaskExecutionListParamsEncoding) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListParamsEncodingProto, EnvironmentAutomationTaskExecutionListParamsEncodingJson:
		return true
	}
	return false
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

// Which compression algorithm to use for this request
type EnvironmentAutomationTaskExecutionListParamsCompression string

const (
	EnvironmentAutomationTaskExecutionListParamsCompressionIdentity EnvironmentAutomationTaskExecutionListParamsCompression = "identity"
	EnvironmentAutomationTaskExecutionListParamsCompressionGzip     EnvironmentAutomationTaskExecutionListParamsCompression = "gzip"
	EnvironmentAutomationTaskExecutionListParamsCompressionBr       EnvironmentAutomationTaskExecutionListParamsCompression = "br"
)

func (r EnvironmentAutomationTaskExecutionListParamsCompression) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListParamsCompressionIdentity, EnvironmentAutomationTaskExecutionListParamsCompressionGzip, EnvironmentAutomationTaskExecutionListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskExecutionListParamsConnect string

const (
	EnvironmentAutomationTaskExecutionListParamsConnectV1 EnvironmentAutomationTaskExecutionListParamsConnect = "v1"
)

func (r EnvironmentAutomationTaskExecutionListParamsConnect) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListParamsConnectV1:
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
