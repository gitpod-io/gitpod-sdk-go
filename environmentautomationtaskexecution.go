// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/gitpod-io/flex-sdk-go/internal/apijson"
	"github.com/gitpod-io/flex-sdk-go/internal/apiquery"
	"github.com/gitpod-io/flex-sdk-go/internal/param"
	"github.com/gitpod-io/flex-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/flex-sdk-go/option"
	"github.com/gitpod-io/flex-sdk-go/packages/pagination"
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
func (r *EnvironmentAutomationTaskExecutionService) Get(ctx context.Context, body EnvironmentAutomationTaskExecutionGetParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskExecutionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/GetTaskExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListTaskExecutions
func (r *EnvironmentAutomationTaskExecutionService) List(ctx context.Context, params EnvironmentAutomationTaskExecutionListParams, opts ...option.RequestOption) (res *pagination.TaskExecutionsPage[EnvironmentAutomationTaskExecutionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.EnvironmentAutomationService/ListTaskExecutions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// ListTaskExecutions
func (r *EnvironmentAutomationTaskExecutionService) ListAutoPaging(ctx context.Context, params EnvironmentAutomationTaskExecutionListParams, opts ...option.RequestOption) *pagination.TaskExecutionsPageAutoPager[EnvironmentAutomationTaskExecutionListResponse] {
	return pagination.NewTaskExecutionsPageAutoPager(r.List(ctx, params, opts...))
}

// StopTaskExecution
func (r *EnvironmentAutomationTaskExecutionService) Stop(ctx context.Context, body EnvironmentAutomationTaskExecutionStopParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskExecutionStopResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StopTaskExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
	// non-empty the task execution has failed to operate and will likely transition to
	// a failed state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a task execution represents the aggregated phase of all steps.
	Phase EnvironmentAutomationTaskExecutionGetResponseTaskExecutionStatusPhase `json:"phase"`
	// version of the status update. Task executions themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion"`
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
	ID       string                                                 `json:"id" format:"uuid"`
	Metadata EnvironmentAutomationTaskExecutionListResponseMetadata `json:"metadata"`
	Spec     EnvironmentAutomationTaskExecutionListResponseSpec     `json:"spec"`
	Status   EnvironmentAutomationTaskExecutionListResponseStatus   `json:"status"`
	JSON     environmentAutomationTaskExecutionListResponseJSON     `json:"-"`
}

// environmentAutomationTaskExecutionListResponseJSON contains the JSON metadata
// for the struct [EnvironmentAutomationTaskExecutionListResponse]
type environmentAutomationTaskExecutionListResponseJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionListResponseMetadata struct {
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
	Creator EnvironmentAutomationTaskExecutionListResponseMetadataCreator `json:"creator"`
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
	TaskID string                                                     `json:"taskId" format:"uuid"`
	JSON   environmentAutomationTaskExecutionListResponseMetadataJSON `json:"-"`
}

// environmentAutomationTaskExecutionListResponseMetadataJSON contains the JSON
// metadata for the struct [EnvironmentAutomationTaskExecutionListResponseMetadata]
type environmentAutomationTaskExecutionListResponseMetadataJSON struct {
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

func (r *EnvironmentAutomationTaskExecutionListResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created/started the task run.
type EnvironmentAutomationTaskExecutionListResponseMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationTaskExecutionListResponseMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationTaskExecutionListResponseMetadataCreatorJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseMetadataCreator]
type environmentAutomationTaskExecutionListResponseMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipal string

const (
	EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationTaskExecutionListResponseMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionListResponseSpec struct {
	// desired_phase is the phase the task execution should be in. Used to stop a
	// running task execution early.
	DesiredPhase EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhase `json:"desiredPhase"`
	// plan is a list of groups of steps. The steps in a group are executed
	// concurrently, while the groups are executed sequentially. The order of the
	// groups is the order in which they are executed.
	Plan []EnvironmentAutomationTaskExecutionListResponseSpecPlan `json:"plan"`
	JSON environmentAutomationTaskExecutionListResponseSpecJSON   `json:"-"`
}

// environmentAutomationTaskExecutionListResponseSpecJSON contains the JSON
// metadata for the struct [EnvironmentAutomationTaskExecutionListResponseSpec]
type environmentAutomationTaskExecutionListResponseSpecJSON struct {
	DesiredPhase apijson.Field
	Plan         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseSpecJSON) RawJSON() string {
	return r.raw
}

// desired_phase is the phase the task execution should be in. Used to stop a
// running task execution early.
type EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhase string

const (
	EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionListResponseSpecDesiredPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionListResponseSpecPlan struct {
	Steps []EnvironmentAutomationTaskExecutionListResponseSpecPlanStep `json:"steps"`
	JSON  environmentAutomationTaskExecutionListResponseSpecPlanJSON   `json:"-"`
}

// environmentAutomationTaskExecutionListResponseSpecPlanJSON contains the JSON
// metadata for the struct [EnvironmentAutomationTaskExecutionListResponseSpecPlan]
type environmentAutomationTaskExecutionListResponseSpecPlanJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseSpecPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseSpecPlanJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionListResponseSpecPlanStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// This field can have the runtime type of [[]string].
	DependsOn interface{} `json:"dependsOn"`
	Label     string      `json:"label"`
	ServiceID string      `json:"serviceId" format:"uuid"`
	// This field can have the runtime type of
	// [EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsObjectTask].
	Task  interface{}                                                    `json:"task"`
	JSON  environmentAutomationTaskExecutionListResponseSpecPlanStepJSON `json:"-"`
	union EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsUnion
}

// environmentAutomationTaskExecutionListResponseSpecPlanStepJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseSpecPlanStep]
type environmentAutomationTaskExecutionListResponseSpecPlanStepJSON struct {
	ID          apijson.Field
	DependsOn   apijson.Field
	Label       apijson.Field
	ServiceID   apijson.Field
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentAutomationTaskExecutionListResponseSpecPlanStepJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationTaskExecutionListResponseSpecPlanStep) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationTaskExecutionListResponseSpecPlanStep{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsObject],
// [EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsObject].
func (r EnvironmentAutomationTaskExecutionListResponseSpecPlanStep) AsUnion() EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsUnion {
	return r.union
}

// Union satisfied by
// [EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsObject] or
// [EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsObject].
type EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsUnion interface {
	implementsEnvironmentAutomationTaskExecutionListResponseSpecPlanStep()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsObject{}),
		},
	)
}

type EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsObject struct {
	ServiceID string `json:"serviceId,required" format:"uuid"`
	// ID is the ID of the execution step
	ID        string                                                                `json:"id" format:"uuid"`
	DependsOn []string                                                              `json:"dependsOn"`
	Label     string                                                                `json:"label"`
	JSON      environmentAutomationTaskExecutionListResponseSpecPlanStepsObjectJSON `json:"-"`
}

// environmentAutomationTaskExecutionListResponseSpecPlanStepsObjectJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsObject]
type environmentAutomationTaskExecutionListResponseSpecPlanStepsObjectJSON struct {
	ServiceID   apijson.Field
	ID          apijson.Field
	DependsOn   apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseSpecPlanStepsObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationTaskExecutionListResponseSpecPlanStepsObject) implementsEnvironmentAutomationTaskExecutionListResponseSpecPlanStep() {
}

type EnvironmentAutomationTaskExecutionListResponseStatus struct {
	// failure_message summarises why the task execution failed to operate. If this is
	// non-empty the task execution has failed to operate and will likely transition to
	// a failed state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a task execution represents the aggregated phase of all steps.
	Phase EnvironmentAutomationTaskExecutionListResponseStatusPhase `json:"phase"`
	// version of the status update. Task executions themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion"`
	// steps provides the status for each individual step of the task execution. If a
	// step is missing it has not yet started.
	Steps []EnvironmentAutomationTaskExecutionListResponseStatusStep `json:"steps"`
	JSON  environmentAutomationTaskExecutionListResponseStatusJSON   `json:"-"`
}

// environmentAutomationTaskExecutionListResponseStatusJSON contains the JSON
// metadata for the struct [EnvironmentAutomationTaskExecutionListResponseStatus]
type environmentAutomationTaskExecutionListResponseStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	StatusVersion  apijson.Field
	Steps          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseStatusJSON) RawJSON() string {
	return r.raw
}

// the phase of a task execution represents the aggregated phase of all steps.
type EnvironmentAutomationTaskExecutionListResponseStatusPhase string

const (
	EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionListResponseStatusPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionListResponseStatusPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionListResponseStatusPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionListResponseStatusPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionListResponseStatusPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionListResponseStatusPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionListResponseStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionListResponseStatusPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionListResponseStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the step failed to operate. If this is non-empty
	// the step has failed to operate and will likely transition to a failed state.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the execution step
	Phase EnvironmentAutomationTaskExecutionListResponseStatusStepsPhase `json:"phase"`
	JSON  environmentAutomationTaskExecutionListResponseStatusStepJSON   `json:"-"`
}

// environmentAutomationTaskExecutionListResponseStatusStepJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskExecutionListResponseStatusStep]
type environmentAutomationTaskExecutionListResponseStatusStepJSON struct {
	ID             apijson.Field
	FailureMessage apijson.Field
	Phase          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionListResponseStatusStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionListResponseStatusStepJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the execution step
type EnvironmentAutomationTaskExecutionListResponseStatusStepsPhase string

const (
	EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionListResponseStatusStepsPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionListResponseStatusStepsPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionListResponseStatusStepsPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionListResponseStatusStepsPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionListResponseStatusStepsPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionListResponseStatusStepsPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionListResponseStatusStepsPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionListResponseStatusStepsPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskExecutionStopResponse = interface{}

type EnvironmentAutomationTaskExecutionGetParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationTaskExecutionGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskExecutionListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// filter contains the filter options for listing task runs
	Filter param.Field[EnvironmentAutomationTaskExecutionListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing task runs
	Pagination param.Field[EnvironmentAutomationTaskExecutionListParamsPagination] `json:"pagination"`
}

func (r EnvironmentAutomationTaskExecutionListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [EnvironmentAutomationTaskExecutionListParams]'s query
// parameters as `url.Values`.
func (r EnvironmentAutomationTaskExecutionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// filter contains the filter options for listing task runs
type EnvironmentAutomationTaskExecutionListParamsFilter struct {
	// environment_ids filters the response to only task runs of these environments
	EnvironmentIDs param.Field[[]string] `json:"environmentIds" format:"uuid"`
	// phases filters the response to only task runs in these phases
	Phases param.Field[[]EnvironmentAutomationTaskExecutionListParamsFilterPhase] `json:"phases"`
	// task_ids filters the response to only task runs of these tasks
	TaskIDs param.Field[[]string] `json:"taskIds" format:"uuid"`
	// task_references filters the response to only task runs with this reference
	TaskReferences param.Field[[]string] `json:"taskReferences"`
}

func (r EnvironmentAutomationTaskExecutionListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskExecutionListParamsFilterPhase string

const (
	EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskExecutionListParamsFilterPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskExecutionListParamsFilterPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskExecutionListParamsFilterPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskExecutionListParamsFilterPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskExecutionListParamsFilterPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskExecutionListParamsFilterPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskExecutionListParamsFilterPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskExecutionListParamsFilterPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

// pagination contains the pagination options for listing task runs
type EnvironmentAutomationTaskExecutionListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EnvironmentAutomationTaskExecutionListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskExecutionStopParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationTaskExecutionStopParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
