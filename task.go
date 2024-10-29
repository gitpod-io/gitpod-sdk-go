// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// TaskService contains methods and other services that help with interacting with
// the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskService] method instead.
type TaskService struct {
	Options []option.RequestOption
}

// NewTaskService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaskService(opts ...option.RequestOption) (r *TaskService) {
	r = &TaskService{}
	r.Options = opts
	return
}

// CreateTask
func (r *TaskService) New(ctx context.Context, params TaskNewParams, opts ...option.RequestOption) (res *TaskNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/CreateTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetTask
func (r *TaskService) Get(ctx context.Context, params TaskGetParams, opts ...option.RequestOption) (res *TaskGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/GetTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// GetTask
func (r *TaskService) GetNew(ctx context.Context, params TaskGetNewParams, opts ...option.RequestOption) (res *TaskGetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/GetTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type TaskNewResponse struct {
	Task TaskNewResponseTask `json:"task"`
	JSON taskNewResponseJSON `json:"-"`
}

// taskNewResponseJSON contains the JSON metadata for the struct [TaskNewResponse]
type taskNewResponseJSON struct {
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskNewResponseJSON) RawJSON() string {
	return r.raw
}

type TaskNewResponseTask struct {
	ID string `json:"id" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn     []string                    `json:"dependsOn" format:"uuid"`
	EnvironmentID string                      `json:"environmentId" format:"uuid"`
	Metadata      TaskNewResponseTaskMetadata `json:"metadata"`
	Spec          TaskNewResponseTaskSpec     `json:"spec"`
	JSON          taskNewResponseTaskJSON     `json:"-"`
}

// taskNewResponseTaskJSON contains the JSON metadata for the struct
// [TaskNewResponseTask]
type taskNewResponseTaskJSON struct {
	ID            apijson.Field
	DependsOn     apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TaskNewResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskNewResponseTaskJSON) RawJSON() string {
	return r.raw
}

type TaskNewResponseTaskMetadata struct {
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
	Creator TaskNewResponseTaskMetadataCreator `json:"creator"`
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
	TriggeredBy []TaskNewResponseTaskMetadataTriggeredByUnion `json:"triggeredBy"`
	JSON        taskNewResponseTaskMetadataJSON               `json:"-"`
}

// taskNewResponseTaskMetadataJSON contains the JSON metadata for the struct
// [TaskNewResponseTaskMetadata]
type taskNewResponseTaskMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskNewResponseTaskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskNewResponseTaskMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created the task.
type TaskNewResponseTaskMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal TaskNewResponseTaskMetadataCreatorPrincipal `json:"principal"`
	JSON      taskNewResponseTaskMetadataCreatorJSON      `json:"-"`
}

// taskNewResponseTaskMetadataCreatorJSON contains the JSON metadata for the struct
// [TaskNewResponseTaskMetadataCreator]
type taskNewResponseTaskMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskNewResponseTaskMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskNewResponseTaskMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type TaskNewResponseTaskMetadataCreatorPrincipal string

const (
	TaskNewResponseTaskMetadataCreatorPrincipalPrincipalUnspecified    TaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	TaskNewResponseTaskMetadataCreatorPrincipalPrincipalAccount        TaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	TaskNewResponseTaskMetadataCreatorPrincipalPrincipalUser           TaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_USER"
	TaskNewResponseTaskMetadataCreatorPrincipalPrincipalRunner         TaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	TaskNewResponseTaskMetadataCreatorPrincipalPrincipalEnvironment    TaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	TaskNewResponseTaskMetadataCreatorPrincipalPrincipalServiceAccount TaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r TaskNewResponseTaskMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case TaskNewResponseTaskMetadataCreatorPrincipalPrincipalUnspecified, TaskNewResponseTaskMetadataCreatorPrincipalPrincipalAccount, TaskNewResponseTaskMetadataCreatorPrincipalPrincipalUser, TaskNewResponseTaskMetadataCreatorPrincipalPrincipalRunner, TaskNewResponseTaskMetadataCreatorPrincipalPrincipalEnvironment, TaskNewResponseTaskMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// Union satisfied by [TaskNewResponseTaskMetadataTriggeredByUnknown],
// [TaskNewResponseTaskMetadataTriggeredByUnknown],
// [TaskNewResponseTaskMetadataTriggeredByUnknown] or
// [TaskNewResponseTaskMetadataTriggeredByUnknown].
type TaskNewResponseTaskMetadataTriggeredByUnion interface {
	implementsTaskNewResponseTaskMetadataTriggeredByUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*TaskNewResponseTaskMetadataTriggeredByUnion)(nil)).Elem(), "")
}

type TaskNewResponseTaskSpec struct {
	// command contains the command the task should execute
	Command string                      `json:"command"`
	JSON    taskNewResponseTaskSpecJSON `json:"-"`
}

// taskNewResponseTaskSpecJSON contains the JSON metadata for the struct
// [TaskNewResponseTaskSpec]
type taskNewResponseTaskSpecJSON struct {
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskNewResponseTaskSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskNewResponseTaskSpecJSON) RawJSON() string {
	return r.raw
}

type TaskGetResponse struct {
	Task TaskGetResponseTask `json:"task"`
	JSON taskGetResponseJSON `json:"-"`
}

// taskGetResponseJSON contains the JSON metadata for the struct [TaskGetResponse]
type taskGetResponseJSON struct {
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetResponseJSON) RawJSON() string {
	return r.raw
}

type TaskGetResponseTask struct {
	ID string `json:"id" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn     []string                    `json:"dependsOn" format:"uuid"`
	EnvironmentID string                      `json:"environmentId" format:"uuid"`
	Metadata      TaskGetResponseTaskMetadata `json:"metadata"`
	Spec          TaskGetResponseTaskSpec     `json:"spec"`
	JSON          taskGetResponseTaskJSON     `json:"-"`
}

// taskGetResponseTaskJSON contains the JSON metadata for the struct
// [TaskGetResponseTask]
type taskGetResponseTaskJSON struct {
	ID            apijson.Field
	DependsOn     apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TaskGetResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetResponseTaskJSON) RawJSON() string {
	return r.raw
}

type TaskGetResponseTaskMetadata struct {
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
	Creator TaskGetResponseTaskMetadataCreator `json:"creator"`
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
	TriggeredBy []TaskGetResponseTaskMetadataTriggeredByUnion `json:"triggeredBy"`
	JSON        taskGetResponseTaskMetadataJSON               `json:"-"`
}

// taskGetResponseTaskMetadataJSON contains the JSON metadata for the struct
// [TaskGetResponseTaskMetadata]
type taskGetResponseTaskMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskGetResponseTaskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetResponseTaskMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created the task.
type TaskGetResponseTaskMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal TaskGetResponseTaskMetadataCreatorPrincipal `json:"principal"`
	JSON      taskGetResponseTaskMetadataCreatorJSON      `json:"-"`
}

// taskGetResponseTaskMetadataCreatorJSON contains the JSON metadata for the struct
// [TaskGetResponseTaskMetadataCreator]
type taskGetResponseTaskMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskGetResponseTaskMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetResponseTaskMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type TaskGetResponseTaskMetadataCreatorPrincipal string

const (
	TaskGetResponseTaskMetadataCreatorPrincipalPrincipalUnspecified    TaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	TaskGetResponseTaskMetadataCreatorPrincipalPrincipalAccount        TaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	TaskGetResponseTaskMetadataCreatorPrincipalPrincipalUser           TaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_USER"
	TaskGetResponseTaskMetadataCreatorPrincipalPrincipalRunner         TaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	TaskGetResponseTaskMetadataCreatorPrincipalPrincipalEnvironment    TaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	TaskGetResponseTaskMetadataCreatorPrincipalPrincipalServiceAccount TaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r TaskGetResponseTaskMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case TaskGetResponseTaskMetadataCreatorPrincipalPrincipalUnspecified, TaskGetResponseTaskMetadataCreatorPrincipalPrincipalAccount, TaskGetResponseTaskMetadataCreatorPrincipalPrincipalUser, TaskGetResponseTaskMetadataCreatorPrincipalPrincipalRunner, TaskGetResponseTaskMetadataCreatorPrincipalPrincipalEnvironment, TaskGetResponseTaskMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// Union satisfied by [TaskGetResponseTaskMetadataTriggeredByUnknown],
// [TaskGetResponseTaskMetadataTriggeredByUnknown],
// [TaskGetResponseTaskMetadataTriggeredByUnknown] or
// [TaskGetResponseTaskMetadataTriggeredByUnknown].
type TaskGetResponseTaskMetadataTriggeredByUnion interface {
	implementsTaskGetResponseTaskMetadataTriggeredByUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*TaskGetResponseTaskMetadataTriggeredByUnion)(nil)).Elem(), "")
}

type TaskGetResponseTaskSpec struct {
	// command contains the command the task should execute
	Command string                      `json:"command"`
	JSON    taskGetResponseTaskSpecJSON `json:"-"`
}

// taskGetResponseTaskSpecJSON contains the JSON metadata for the struct
// [TaskGetResponseTaskSpec]
type taskGetResponseTaskSpecJSON struct {
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskGetResponseTaskSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetResponseTaskSpecJSON) RawJSON() string {
	return r.raw
}

type TaskGetNewResponse struct {
	Task TaskGetNewResponseTask `json:"task"`
	JSON taskGetNewResponseJSON `json:"-"`
}

// taskGetNewResponseJSON contains the JSON metadata for the struct
// [TaskGetNewResponse]
type taskGetNewResponseJSON struct {
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskGetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetNewResponseJSON) RawJSON() string {
	return r.raw
}

type TaskGetNewResponseTask struct {
	ID string `json:"id" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn     []string                       `json:"dependsOn" format:"uuid"`
	EnvironmentID string                         `json:"environmentId" format:"uuid"`
	Metadata      TaskGetNewResponseTaskMetadata `json:"metadata"`
	Spec          TaskGetNewResponseTaskSpec     `json:"spec"`
	JSON          taskGetNewResponseTaskJSON     `json:"-"`
}

// taskGetNewResponseTaskJSON contains the JSON metadata for the struct
// [TaskGetNewResponseTask]
type taskGetNewResponseTaskJSON struct {
	ID            apijson.Field
	DependsOn     apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TaskGetNewResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetNewResponseTaskJSON) RawJSON() string {
	return r.raw
}

type TaskGetNewResponseTaskMetadata struct {
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
	Creator TaskGetNewResponseTaskMetadataCreator `json:"creator"`
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
	TriggeredBy []TaskGetNewResponseTaskMetadataTriggeredByUnion `json:"triggeredBy"`
	JSON        taskGetNewResponseTaskMetadataJSON               `json:"-"`
}

// taskGetNewResponseTaskMetadataJSON contains the JSON metadata for the struct
// [TaskGetNewResponseTaskMetadata]
type taskGetNewResponseTaskMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskGetNewResponseTaskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetNewResponseTaskMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created the task.
type TaskGetNewResponseTaskMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal TaskGetNewResponseTaskMetadataCreatorPrincipal `json:"principal"`
	JSON      taskGetNewResponseTaskMetadataCreatorJSON      `json:"-"`
}

// taskGetNewResponseTaskMetadataCreatorJSON contains the JSON metadata for the
// struct [TaskGetNewResponseTaskMetadataCreator]
type taskGetNewResponseTaskMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskGetNewResponseTaskMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetNewResponseTaskMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type TaskGetNewResponseTaskMetadataCreatorPrincipal string

const (
	TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalUnspecified    TaskGetNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalAccount        TaskGetNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalUser           TaskGetNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_USER"
	TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalRunner         TaskGetNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalEnvironment    TaskGetNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalServiceAccount TaskGetNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r TaskGetNewResponseTaskMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalUnspecified, TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalAccount, TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalUser, TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalRunner, TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalEnvironment, TaskGetNewResponseTaskMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// Union satisfied by [TaskGetNewResponseTaskMetadataTriggeredByUnknown],
// [TaskGetNewResponseTaskMetadataTriggeredByUnknown],
// [TaskGetNewResponseTaskMetadataTriggeredByUnknown] or
// [TaskGetNewResponseTaskMetadataTriggeredByUnknown].
type TaskGetNewResponseTaskMetadataTriggeredByUnion interface {
	implementsTaskGetNewResponseTaskMetadataTriggeredByUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*TaskGetNewResponseTaskMetadataTriggeredByUnion)(nil)).Elem(), "")
}

type TaskGetNewResponseTaskSpec struct {
	// command contains the command the task should execute
	Command string                         `json:"command"`
	JSON    taskGetNewResponseTaskSpecJSON `json:"-"`
}

// taskGetNewResponseTaskSpecJSON contains the JSON metadata for the struct
// [TaskGetNewResponseTaskSpec]
type taskGetNewResponseTaskSpecJSON struct {
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskGetNewResponseTaskSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetNewResponseTaskSpecJSON) RawJSON() string {
	return r.raw
}

type TaskNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[TaskNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	DependsOn              param.Field[[]string]                            `json:"dependsOn" format:"uuid"`
	EnvironmentID          param.Field[string]                              `json:"environmentId" format:"uuid"`
	Metadata               param.Field[TaskNewParamsMetadata]               `json:"metadata"`
	Spec                   param.Field[TaskNewParamsSpec]                   `json:"spec"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r TaskNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type TaskNewParamsConnectProtocolVersion float64

const (
	TaskNewParamsConnectProtocolVersion1 TaskNewParamsConnectProtocolVersion = 1
)

func (r TaskNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case TaskNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type TaskNewParamsMetadata struct {
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
	Creator param.Field[TaskNewParamsMetadataCreator] `json:"creator"`
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
	TriggeredBy param.Field[[]TaskNewParamsMetadataTriggeredByUnion] `json:"triggeredBy"`
}

func (r TaskNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// creator describes the principal who created the task.
type TaskNewParamsMetadataCreator struct {
	// id is the UUID of the subject
	ID param.Field[string] `json:"id"`
	// Principal is the principal of the subject
	Principal param.Field[TaskNewParamsMetadataCreatorPrincipal] `json:"principal"`
}

func (r TaskNewParamsMetadataCreator) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Principal is the principal of the subject
type TaskNewParamsMetadataCreatorPrincipal string

const (
	TaskNewParamsMetadataCreatorPrincipalPrincipalUnspecified    TaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	TaskNewParamsMetadataCreatorPrincipalPrincipalAccount        TaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	TaskNewParamsMetadataCreatorPrincipalPrincipalUser           TaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_USER"
	TaskNewParamsMetadataCreatorPrincipalPrincipalRunner         TaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	TaskNewParamsMetadataCreatorPrincipalPrincipalEnvironment    TaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	TaskNewParamsMetadataCreatorPrincipalPrincipalServiceAccount TaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r TaskNewParamsMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case TaskNewParamsMetadataCreatorPrincipalPrincipalUnspecified, TaskNewParamsMetadataCreatorPrincipalPrincipalAccount, TaskNewParamsMetadataCreatorPrincipalPrincipalUser, TaskNewParamsMetadataCreatorPrincipalPrincipalRunner, TaskNewParamsMetadataCreatorPrincipalPrincipalEnvironment, TaskNewParamsMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// Satisfied by [TaskNewParamsMetadataTriggeredByUnknown],
// [TaskNewParamsMetadataTriggeredByUnknown],
// [TaskNewParamsMetadataTriggeredByUnknown],
// [TaskNewParamsMetadataTriggeredByUnknown].
type TaskNewParamsMetadataTriggeredByUnion interface {
	implementsTaskNewParamsMetadataTriggeredByUnion()
}

type TaskNewParamsSpec struct {
	// command contains the command the task should execute
	Command param.Field[string] `json:"command"`
}

func (r TaskNewParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TaskGetParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[TaskGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Base64                 param.Field[string]                              `query:"base64"`
	Compression            param.Field[string]                              `query:"compression"`
	Connect                param.Field[string]                              `query:"connect"`
	Encoding               param.Field[string]                              `query:"encoding"`
	Message                param.Field[string]                              `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [TaskGetParams]'s query parameters as `url.Values`.
func (r TaskGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define the version of the Connect protocol
type TaskGetParamsConnectProtocolVersion float64

const (
	TaskGetParamsConnectProtocolVersion1 TaskGetParamsConnectProtocolVersion = 1
)

func (r TaskGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case TaskGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type TaskGetNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[TaskGetNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                 `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r TaskGetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type TaskGetNewParamsConnectProtocolVersion float64

const (
	TaskGetNewParamsConnectProtocolVersion1 TaskGetNewParamsConnectProtocolVersion = 1
)

func (r TaskGetNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case TaskGetNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}
