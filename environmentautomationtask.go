// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
	"github.com/stainless-sdks/gitpod-go/shared"
	"github.com/tidwall/gjson"
)

// EnvironmentAutomationTaskService contains methods and other services that help
// with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentAutomationTaskService] method instead.
type EnvironmentAutomationTaskService struct {
	Options []option.RequestOption
}

// NewEnvironmentAutomationTaskService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEnvironmentAutomationTaskService(opts ...option.RequestOption) (r *EnvironmentAutomationTaskService) {
	r = &EnvironmentAutomationTaskService{}
	r.Options = opts
	return
}

// UpdateTask
func (r *EnvironmentAutomationTaskService) Update(ctx context.Context, params EnvironmentAutomationTaskUpdateParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/UpdateTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListTasks
func (r *EnvironmentAutomationTaskService) List(ctx context.Context, params EnvironmentAutomationTaskListParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/ListTasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// DeleteTask
func (r *EnvironmentAutomationTaskService) Delete(ctx context.Context, params EnvironmentAutomationTaskDeleteParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/DeleteTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// StartTask starts a task, i.e. creates a task execution. This call does not block
// until the task is started; the task will be started asynchronously.
func (r *EnvironmentAutomationTaskService) Start(ctx context.Context, params EnvironmentAutomationTaskStartParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskStartResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StartTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type EnvironmentAutomationTaskUpdateResponse = interface{}

type EnvironmentAutomationTaskListResponse struct {
	Pagination EnvironmentAutomationTaskListResponsePagination `json:"pagination"`
	Tasks      []EnvironmentAutomationTaskListResponseTask     `json:"tasks"`
	JSON       environmentAutomationTaskListResponseJSON       `json:"-"`
}

// environmentAutomationTaskListResponseJSON contains the JSON metadata for the
// struct [EnvironmentAutomationTaskListResponse]
type environmentAutomationTaskListResponseJSON struct {
	Pagination  apijson.Field
	Tasks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                              `json:"nextToken"`
	JSON      environmentAutomationTaskListResponsePaginationJSON `json:"-"`
}

// environmentAutomationTaskListResponsePaginationJSON contains the JSON metadata
// for the struct [EnvironmentAutomationTaskListResponsePagination]
type environmentAutomationTaskListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskListResponseTask struct {
	ID string `json:"id" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn     []string                                           `json:"dependsOn" format:"uuid"`
	EnvironmentID string                                             `json:"environmentId" format:"uuid"`
	Metadata      EnvironmentAutomationTaskListResponseTasksMetadata `json:"metadata"`
	Spec          EnvironmentAutomationTaskListResponseTasksSpec     `json:"spec"`
	JSON          environmentAutomationTaskListResponseTaskJSON      `json:"-"`
}

// environmentAutomationTaskListResponseTaskJSON contains the JSON metadata for the
// struct [EnvironmentAutomationTaskListResponseTask]
type environmentAutomationTaskListResponseTaskJSON struct {
	ID            apijson.Field
	DependsOn     apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskListResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskListResponseTaskJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskListResponseTasksMetadata struct {
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
	// creator describes the principal who created the task.
	Creator EnvironmentAutomationTaskListResponseTasksMetadataCreator `json:"creator"`
	// description is a user-facing description for the task. It can be used to provide
	// context and documentation for the task.
	Description string `json:"description"`
	// name is a user-facing name for the task. Unlike the reference, this field is not
	// unique, and not referenced by the system. This is a short descriptive name for
	// the task.
	Name string `json:"name"`
	// reference is a user-facing identifier for the task which must be unique on the
	// environment. It is used to express dependencies between tasks, and to identify
	// the task in user interactions (e.g. the CLI).
	Reference string `json:"reference"`
	// triggered_by is a list of trigger that start the task.
	TriggeredBy []EnvironmentAutomationTaskListResponseTasksMetadataTriggeredBy `json:"triggeredBy"`
	JSON        environmentAutomationTaskListResponseTasksMetadataJSON          `json:"-"`
}

// environmentAutomationTaskListResponseTasksMetadataJSON contains the JSON
// metadata for the struct [EnvironmentAutomationTaskListResponseTasksMetadata]
type environmentAutomationTaskListResponseTasksMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskListResponseTasksMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskListResponseTasksMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created the task.
type EnvironmentAutomationTaskListResponseTasksMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationTaskListResponseTasksMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationTaskListResponseTasksMetadataCreatorJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskListResponseTasksMetadataCreator]
type environmentAutomationTaskListResponseTasksMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskListResponseTasksMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskListResponseTasksMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipal string

const (
	EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationTaskListResponseTasksMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type EnvironmentAutomationTaskListResponseTasksMetadataTriggeredBy struct {
	JSON environmentAutomationTaskListResponseTasksMetadataTriggeredByJSON `json:"-"`
}

// environmentAutomationTaskListResponseTasksMetadataTriggeredByJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredBy]
type environmentAutomationTaskListResponseTasksMetadataTriggeredByJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskListResponseTasksMetadataTriggeredBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskListResponseTasksMetadataTriggeredByJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskListResponseTasksSpec struct {
	// command contains the command the task should execute
	Command string `json:"command"`
	// runs_on specifies the environment the task should run on.
	RunsOn EnvironmentAutomationTaskListResponseTasksSpecRunsOnUnion `json:"runsOn"`
	JSON   environmentAutomationTaskListResponseTasksSpecJSON        `json:"-"`
}

// environmentAutomationTaskListResponseTasksSpecJSON contains the JSON metadata
// for the struct [EnvironmentAutomationTaskListResponseTasksSpec]
type environmentAutomationTaskListResponseTasksSpecJSON struct {
	Command     apijson.Field
	RunsOn      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskListResponseTasksSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskListResponseTasksSpecJSON) RawJSON() string {
	return r.raw
}

// runs_on specifies the environment the task should run on.
//
// Union satisfied by [EnvironmentAutomationTaskListResponseTasksSpecRunsOnUnknown]
// or [EnvironmentAutomationTaskListResponseTasksSpecRunsOnUnknown].
type EnvironmentAutomationTaskListResponseTasksSpecRunsOnUnion interface {
	implementsEnvironmentAutomationTaskListResponseTasksSpecRunsOnUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentAutomationTaskListResponseTasksSpecRunsOnUnion)(nil)).Elem(), "")
}

type EnvironmentAutomationTaskDeleteResponse = interface{}

type EnvironmentAutomationTaskStartResponse struct {
	TaskExecution EnvironmentAutomationTaskStartResponseTaskExecution `json:"taskExecution"`
	JSON          environmentAutomationTaskStartResponseJSON          `json:"-"`
}

// environmentAutomationTaskStartResponseJSON contains the JSON metadata for the
// struct [EnvironmentAutomationTaskStartResponse]
type environmentAutomationTaskStartResponseJSON struct {
	TaskExecution apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskStartResponseTaskExecution struct {
	ID       string                                                      `json:"id" format:"uuid"`
	Metadata EnvironmentAutomationTaskStartResponseTaskExecutionMetadata `json:"metadata"`
	Spec     EnvironmentAutomationTaskStartResponseTaskExecutionSpec     `json:"spec"`
	Status   EnvironmentAutomationTaskStartResponseTaskExecutionStatus   `json:"status"`
	JSON     environmentAutomationTaskStartResponseTaskExecutionJSON     `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionJSON contains the JSON
// metadata for the struct [EnvironmentAutomationTaskStartResponseTaskExecution]
type environmentAutomationTaskStartResponseTaskExecutionJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskStartResponseTaskExecutionMetadata struct {
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
	Creator EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreator `json:"creator"`
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
	TaskID string                                                          `json:"taskId" format:"uuid"`
	JSON   environmentAutomationTaskStartResponseTaskExecutionMetadataJSON `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionMetadataJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionMetadata]
type environmentAutomationTaskStartResponseTaskExecutionMetadataJSON struct {
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

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created/started the task run.
type EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationTaskStartResponseTaskExecutionMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionMetadataCreatorJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreator]
type environmentAutomationTaskStartResponseTaskExecutionMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipal string

const (
	EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type EnvironmentAutomationTaskStartResponseTaskExecutionSpec struct {
	// desired_phase is the phase the task execution should be in. Used to stop a
	// running task execution early.
	DesiredPhase EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhase `json:"desiredPhase"`
	// plan is a list of groups of steps. The steps in a group are executed
	// concurrently, while the groups are executed sequentially. The order of the
	// groups is the order in which they are executed.
	Plan []EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlan `json:"plan"`
	JSON environmentAutomationTaskStartResponseTaskExecutionSpecJSON   `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionSpecJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpec]
type environmentAutomationTaskStartResponseTaskExecutionSpecJSON struct {
	DesiredPhase apijson.Field
	Plan         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionSpecJSON) RawJSON() string {
	return r.raw
}

// desired_phase is the phase the task execution should be in. Used to stop a
// running task execution early.
type EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhase string

const (
	EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskStartResponseTaskExecutionSpecDesiredPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlan struct {
	Steps []EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStep `json:"steps"`
	JSON  environmentAutomationTaskStartResponseTaskExecutionSpecPlanJSON   `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionSpecPlanJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlan]
type environmentAutomationTaskStartResponseTaskExecutionSpecPlanJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionSpecPlanJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStep struct {
	JSON environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepJSON `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStep]
type environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskStartResponseTaskExecutionStatus struct {
	// failure_message summarises why the task execution failed to operate. If this is
	// non-empty the task execution has failed to operate and will likely transition to
	// a failed state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a task execution represents the aggregated phase of all steps.
	Phase EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhase `json:"phase"`
	// version of the status update. Task executions themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion EnvironmentAutomationTaskStartResponseTaskExecutionStatusStatusVersionUnion `json:"statusVersion" format:"int64"`
	// steps provides the status for each individual step of the task execution. If a
	// step is missing it has not yet started.
	Steps []EnvironmentAutomationTaskStartResponseTaskExecutionStatusStep `json:"steps"`
	JSON  environmentAutomationTaskStartResponseTaskExecutionStatusJSON   `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionStatusJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionStatus]
type environmentAutomationTaskStartResponseTaskExecutionStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	StatusVersion  apijson.Field
	Steps          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionStatusJSON) RawJSON() string {
	return r.raw
}

// the phase of a task execution represents the aggregated phase of all steps.
type EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhase string

const (
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

// version of the status update. Task executions themselves are unversioned, but
// their status has different versions. The value of this field has no semantic
// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
// impose a partial order. If a.status_version < b.status_version then a was the
// status before b.
//
// Union satisfied by [shared.UnionInt] or [shared.UnionString].
type EnvironmentAutomationTaskStartResponseTaskExecutionStatusStatusVersionUnion interface {
	ImplementsEnvironmentAutomationTaskStartResponseTaskExecutionStatusStatusVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskStartResponseTaskExecutionStatusStatusVersionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type EnvironmentAutomationTaskStartResponseTaskExecutionStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the step failed to operate. If this is non-empty
	// the step has failed to operate and will likely transition to a failed state.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the execution step
	Phase EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhase `json:"phase"`
	JSON  environmentAutomationTaskStartResponseTaskExecutionStatusStepJSON   `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionStatusStepJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionStatusStep]
type environmentAutomationTaskStartResponseTaskExecutionStatusStepJSON struct {
	ID             apijson.Field
	FailureMessage apijson.Field
	Phase          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionStatusStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionStatusStepJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the execution step
type EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhase string

const (
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseUnspecified EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhasePending     EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_PENDING"
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseRunning     EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_RUNNING"
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseSucceeded   EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseFailed      EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_FAILED"
	EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseStopped     EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseUnspecified, EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhasePending, EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseRunning, EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseSucceeded, EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseFailed, EnvironmentAutomationTaskStartResponseTaskExecutionStatusStepsPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type EnvironmentAutomationTaskUpdateParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                      `json:"id" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn param.Field[[]string]                                      `json:"dependsOn" format:"uuid"`
	Metadata  param.Field[EnvironmentAutomationTaskUpdateParamsMetadata] `json:"metadata"`
	Spec      param.Field[EnvironmentAutomationTaskUpdateParamsSpec]     `json:"spec"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationTaskUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskUpdateParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskUpdateParamsConnectProtocolVersion1 EnvironmentAutomationTaskUpdateParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentAutomationTaskUpdateParamsMetadata struct {
}

func (r EnvironmentAutomationTaskUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskUpdateParamsSpec struct {
}

func (r EnvironmentAutomationTaskUpdateParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// filter contains the filter options for listing tasks
	Filter param.Field[EnvironmentAutomationTaskListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing tasks
	Pagination param.Field[EnvironmentAutomationTaskListParamsPagination] `json:"pagination"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationTaskListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskListParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskListParamsConnectProtocolVersion1 EnvironmentAutomationTaskListParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// filter contains the filter options for listing tasks
type EnvironmentAutomationTaskListParamsFilter struct {
	// environment_ids filters the response to only tasks of these environments
	EnvironmentIDs param.Field[[]string] `json:"environmentIds" format:"uuid"`
	// references filters the response to only services with these references
	References param.Field[[]string] `json:"references"`
	// task_ids filters the response to only tasks with these IDs
	TaskIDs param.Field[[]string] `json:"taskIds" format:"uuid"`
}

func (r EnvironmentAutomationTaskListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing tasks
type EnvironmentAutomationTaskListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EnvironmentAutomationTaskListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                      `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationTaskDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskDeleteParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskDeleteParamsConnectProtocolVersion1 EnvironmentAutomationTaskDeleteParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentAutomationTaskStartParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskStartParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                     `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationTaskStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskStartParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskStartParamsConnectProtocolVersion1 EnvironmentAutomationTaskStartParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskStartParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskStartParamsConnectProtocolVersion1:
		return true
	}
	return false
}
