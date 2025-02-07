// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gitpod-io/flex-sdk-go/internal/apijson"
	"github.com/gitpod-io/flex-sdk-go/internal/apiquery"
	"github.com/gitpod-io/flex-sdk-go/internal/param"
	"github.com/gitpod-io/flex-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/flex-sdk-go/option"
	"github.com/gitpod-io/flex-sdk-go/packages/pagination"
	"github.com/gitpod-io/flex-sdk-go/shared"
)

// EnvironmentAutomationTaskService contains methods and other services that help
// with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentAutomationTaskService] method instead.
type EnvironmentAutomationTaskService struct {
	Options    []option.RequestOption
	Executions *EnvironmentAutomationTaskExecutionService
}

// NewEnvironmentAutomationTaskService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEnvironmentAutomationTaskService(opts ...option.RequestOption) (r *EnvironmentAutomationTaskService) {
	r = &EnvironmentAutomationTaskService{}
	r.Options = opts
	r.Executions = NewEnvironmentAutomationTaskExecutionService(opts...)
	return
}

// CreateTask
func (r *EnvironmentAutomationTaskService) New(ctx context.Context, body EnvironmentAutomationTaskNewParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/CreateTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetTask
func (r *EnvironmentAutomationTaskService) Get(ctx context.Context, body EnvironmentAutomationTaskGetParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/GetTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateTask
func (r *EnvironmentAutomationTaskService) Update(ctx context.Context, body EnvironmentAutomationTaskUpdateParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/UpdateTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListTasks
func (r *EnvironmentAutomationTaskService) List(ctx context.Context, params EnvironmentAutomationTaskListParams, opts ...option.RequestOption) (res *pagination.TasksPage[Task], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.EnvironmentAutomationService/ListTasks"
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

// ListTasks
func (r *EnvironmentAutomationTaskService) ListAutoPaging(ctx context.Context, params EnvironmentAutomationTaskListParams, opts ...option.RequestOption) *pagination.TasksPageAutoPager[Task] {
	return pagination.NewTasksPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteTask
func (r *EnvironmentAutomationTaskService) Delete(ctx context.Context, body EnvironmentAutomationTaskDeleteParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/DeleteTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// StartTask starts a task, i.e. creates a task execution. This call does not block
// until the task is started; the task will be started asynchronously.
func (r *EnvironmentAutomationTaskService) Start(ctx context.Context, body EnvironmentAutomationTaskStartParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskStartResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StartTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Task struct {
	ID string `json:"id" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn     []string     `json:"dependsOn" format:"uuid"`
	EnvironmentID string       `json:"environmentId" format:"uuid"`
	Metadata      TaskMetadata `json:"metadata"`
	Spec          TaskSpec     `json:"spec"`
	JSON          taskJSON     `json:"-"`
}

// taskJSON contains the JSON metadata for the struct [Task]
type taskJSON struct {
	ID            apijson.Field
	DependsOn     apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Task) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskJSON) RawJSON() string {
	return r.raw
}

type TaskMetadata struct {
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
	Creator shared.Subject `json:"creator"`
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
	TriggeredBy []shared.AutomationTrigger `json:"triggeredBy"`
	JSON        taskMetadataJSON           `json:"-"`
}

// taskMetadataJSON contains the JSON metadata for the struct [TaskMetadata]
type taskMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskMetadataJSON) RawJSON() string {
	return r.raw
}

type TaskMetadataParam struct {
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
	CreatedAt param.Field[time.Time] `json:"createdAt" format:"date-time"`
	// creator describes the principal who created the task.
	Creator param.Field[shared.SubjectParam] `json:"creator"`
	// description is a user-facing description for the task. It can be used to provide
	// context and documentation for the task.
	Description param.Field[string] `json:"description"`
	// name is a user-facing name for the task. Unlike the reference, this field is not
	// unique, and not referenced by the system. This is a short descriptive name for
	// the task.
	Name param.Field[string] `json:"name"`
	// reference is a user-facing identifier for the task which must be unique on the
	// environment. It is used to express dependencies between tasks, and to identify
	// the task in user interactions (e.g. the CLI).
	Reference param.Field[string] `json:"reference"`
	// triggered_by is a list of trigger that start the task.
	TriggeredBy param.Field[[]shared.AutomationTriggerParam] `json:"triggeredBy"`
}

func (r TaskMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TaskSpec struct {
	// command contains the command the task should execute
	Command string `json:"command"`
	// runs_on specifies the environment the task should run on.
	RunsOn shared.RunsOn `json:"runsOn"`
	JSON   taskSpecJSON  `json:"-"`
}

// taskSpecJSON contains the JSON metadata for the struct [TaskSpec]
type taskSpecJSON struct {
	Command     apijson.Field
	RunsOn      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskSpecJSON) RawJSON() string {
	return r.raw
}

type TaskSpecParam struct {
	// command contains the command the task should execute
	Command param.Field[string] `json:"command"`
	// runs_on specifies the environment the task should run on.
	RunsOn param.Field[shared.RunsOnParam] `json:"runsOn"`
}

func (r TaskSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskNewResponse struct {
	Task Task                                     `json:"task"`
	JSON environmentAutomationTaskNewResponseJSON `json:"-"`
}

// environmentAutomationTaskNewResponseJSON contains the JSON metadata for the
// struct [EnvironmentAutomationTaskNewResponse]
type environmentAutomationTaskNewResponseJSON struct {
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskNewResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskGetResponse struct {
	Task Task                                     `json:"task"`
	JSON environmentAutomationTaskGetResponseJSON `json:"-"`
}

// environmentAutomationTaskGetResponseJSON contains the JSON metadata for the
// struct [EnvironmentAutomationTaskGetResponse]
type environmentAutomationTaskGetResponseJSON struct {
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskGetResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskUpdateResponse = interface{}

type EnvironmentAutomationTaskDeleteResponse = interface{}

type EnvironmentAutomationTaskStartResponse struct {
	TaskExecution shared.TaskExecution                       `json:"taskExecution"`
	JSON          environmentAutomationTaskStartResponseJSON `json:"-"`
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

type EnvironmentAutomationTaskNewParams struct {
	DependsOn     param.Field[[]string]          `json:"dependsOn" format:"uuid"`
	EnvironmentID param.Field[string]            `json:"environmentId" format:"uuid"`
	Metadata      param.Field[TaskMetadataParam] `json:"metadata"`
	Spec          param.Field[TaskSpecParam]     `json:"spec"`
}

func (r EnvironmentAutomationTaskNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskGetParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationTaskGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskUpdateParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn param.Field[[]string]                                      `json:"dependsOn" format:"uuid"`
	Metadata  param.Field[EnvironmentAutomationTaskUpdateParamsMetadata] `json:"metadata"`
	Spec      param.Field[EnvironmentAutomationTaskUpdateParamsSpec]     `json:"spec"`
}

func (r EnvironmentAutomationTaskUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskUpdateParamsMetadata struct {
	Description param.Field[string]                                                   `json:"description"`
	Name        param.Field[string]                                                   `json:"name"`
	TriggeredBy param.Field[EnvironmentAutomationTaskUpdateParamsMetadataTriggeredBy] `json:"triggeredBy"`
}

func (r EnvironmentAutomationTaskUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskUpdateParamsMetadataTriggeredBy struct {
	Trigger param.Field[[]shared.AutomationTriggerParam] `json:"trigger"`
}

func (r EnvironmentAutomationTaskUpdateParamsMetadataTriggeredBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskUpdateParamsSpec struct {
	Command param.Field[string]             `json:"command"`
	RunsOn  param.Field[shared.RunsOnParam] `json:"runsOn"`
}

func (r EnvironmentAutomationTaskUpdateParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// filter contains the filter options for listing tasks
	Filter param.Field[EnvironmentAutomationTaskListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing tasks
	Pagination param.Field[EnvironmentAutomationTaskListParamsPagination] `json:"pagination"`
}

func (r EnvironmentAutomationTaskListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [EnvironmentAutomationTaskListParams]'s query parameters as
// `url.Values`.
func (r EnvironmentAutomationTaskListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationTaskDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskStartParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationTaskStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
