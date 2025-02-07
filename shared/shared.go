// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/gitpod-io/flex-sdk-go"
	"github.com/gitpod-io/flex-sdk-go/internal/apijson"
	"github.com/gitpod-io/flex-sdk-go/internal/param"
)

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type AutomationTrigger struct {
	Manual                bool                  `json:"manual"`
	PostDevcontainerStart bool                  `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                  `json:"postEnvironmentStart"`
	JSON                  automationTriggerJSON `json:"-"`
}

// automationTriggerJSON contains the JSON metadata for the struct
// [AutomationTrigger]
type automationTriggerJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AutomationTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationTriggerJSON) RawJSON() string {
	return r.raw
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type AutomationTriggerParam struct {
	Manual                param.Field[bool] `json:"manual"`
	PostDevcontainerStart param.Field[bool] `json:"postDevcontainerStart"`
	PostEnvironmentStart  param.Field[bool] `json:"postEnvironmentStart"`
}

func (r AutomationTriggerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID string `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration []FieldValue `json:"configuration"`
	// description is a human readable description of the environment class
	Description string `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName string `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled bool `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID string               `json:"runnerId"`
	JSON     environmentClassJSON `json:"-"`
}

// environmentClassJSON contains the JSON metadata for the struct
// [EnvironmentClass]
type environmentClassJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	RunnerID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassJSON) RawJSON() string {
	return r.raw
}

type EnvironmentClassParam struct {
	// id is the unique identifier of the environment class
	ID param.Field[string] `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration param.Field[[]FieldValueParam] `json:"configuration"`
	// description is a human readable description of the environment class
	Description param.Field[string] `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName param.Field[string] `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled param.Field[bool] `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID param.Field[string] `json:"runnerId"`
}

func (r EnvironmentClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FieldValue struct {
	Key   string         `json:"key"`
	Value string         `json:"value"`
	JSON  fieldValueJSON `json:"-"`
}

// fieldValueJSON contains the JSON metadata for the struct [FieldValue]
type fieldValueJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldValueJSON) RawJSON() string {
	return r.raw
}

type FieldValueParam struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r FieldValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationRole string

const (
	OrganizationRoleOrganizationRoleUnspecified OrganizationRole = "ORGANIZATION_ROLE_UNSPECIFIED"
	OrganizationRoleOrganizationRoleAdmin       OrganizationRole = "ORGANIZATION_ROLE_ADMIN"
	OrganizationRoleOrganizationRoleMember      OrganizationRole = "ORGANIZATION_ROLE_MEMBER"
)

func (r OrganizationRole) IsKnown() bool {
	switch r {
	case OrganizationRoleOrganizationRoleUnspecified, OrganizationRoleOrganizationRoleAdmin, OrganizationRoleOrganizationRoleMember:
		return true
	}
	return false
}

type Principal string

const (
	PrincipalPrincipalUnspecified    Principal = "PRINCIPAL_UNSPECIFIED"
	PrincipalPrincipalAccount        Principal = "PRINCIPAL_ACCOUNT"
	PrincipalPrincipalUser           Principal = "PRINCIPAL_USER"
	PrincipalPrincipalRunner         Principal = "PRINCIPAL_RUNNER"
	PrincipalPrincipalEnvironment    Principal = "PRINCIPAL_ENVIRONMENT"
	PrincipalPrincipalServiceAccount Principal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r Principal) IsKnown() bool {
	switch r {
	case PrincipalPrincipalUnspecified, PrincipalPrincipalAccount, PrincipalPrincipalUser, PrincipalPrincipalRunner, PrincipalPrincipalEnvironment, PrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type RunsOn struct {
	Docker RunsOnDocker `json:"docker,required"`
	JSON   runsOnJSON   `json:"-"`
}

// runsOnJSON contains the JSON metadata for the struct [RunsOn]
type runsOnJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunsOn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runsOnJSON) RawJSON() string {
	return r.raw
}

type RunsOnDocker struct {
	Environment []string         `json:"environment"`
	Image       string           `json:"image"`
	JSON        runsOnDockerJSON `json:"-"`
}

// runsOnDockerJSON contains the JSON metadata for the struct [RunsOnDocker]
type runsOnDockerJSON struct {
	Environment apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunsOnDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runsOnDockerJSON) RawJSON() string {
	return r.raw
}

type RunsOnParam struct {
	Docker param.Field[RunsOnDockerParam] `json:"docker,required"`
}

func (r RunsOnParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunsOnDockerParam struct {
	Environment param.Field[[]string] `json:"environment"`
	Image       param.Field[string]   `json:"image"`
}

func (r RunsOnDockerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Subject struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal Principal   `json:"principal"`
	JSON      subjectJSON `json:"-"`
}

// subjectJSON contains the JSON metadata for the struct [Subject]
type subjectJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Subject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subjectJSON) RawJSON() string {
	return r.raw
}

type SubjectParam struct {
	// id is the UUID of the subject
	ID param.Field[string] `json:"id"`
	// Principal is the principal of the subject
	Principal param.Field[Principal] `json:"principal"`
}

func (r SubjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TaskExecution struct {
	ID       string                `json:"id" format:"uuid"`
	Metadata TaskExecutionMetadata `json:"metadata"`
	Spec     TaskExecutionSpec     `json:"spec"`
	Status   TaskExecutionStatus   `json:"status"`
	JSON     taskExecutionJSON     `json:"-"`
}

// taskExecutionJSON contains the JSON metadata for the struct [TaskExecution]
type taskExecutionJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionMetadata struct {
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
	Creator Subject `json:"creator"`
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
	TaskID string                    `json:"taskId" format:"uuid"`
	JSON   taskExecutionMetadataJSON `json:"-"`
}

// taskExecutionMetadataJSON contains the JSON metadata for the struct
// [TaskExecutionMetadata]
type taskExecutionMetadataJSON struct {
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

func (r *TaskExecutionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionMetadataJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionPhase string

const (
	TaskExecutionPhaseTaskExecutionPhaseUnspecified TaskExecutionPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	TaskExecutionPhaseTaskExecutionPhasePending     TaskExecutionPhase = "TASK_EXECUTION_PHASE_PENDING"
	TaskExecutionPhaseTaskExecutionPhaseRunning     TaskExecutionPhase = "TASK_EXECUTION_PHASE_RUNNING"
	TaskExecutionPhaseTaskExecutionPhaseSucceeded   TaskExecutionPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	TaskExecutionPhaseTaskExecutionPhaseFailed      TaskExecutionPhase = "TASK_EXECUTION_PHASE_FAILED"
	TaskExecutionPhaseTaskExecutionPhaseStopped     TaskExecutionPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r TaskExecutionPhase) IsKnown() bool {
	switch r {
	case TaskExecutionPhaseTaskExecutionPhaseUnspecified, TaskExecutionPhaseTaskExecutionPhasePending, TaskExecutionPhaseTaskExecutionPhaseRunning, TaskExecutionPhaseTaskExecutionPhaseSucceeded, TaskExecutionPhaseTaskExecutionPhaseFailed, TaskExecutionPhaseTaskExecutionPhaseStopped:
		return true
	}
	return false
}

type TaskExecutionSpec struct {
	// desired_phase is the phase the task execution should be in. Used to stop a
	// running task execution early.
	DesiredPhase TaskExecutionPhase `json:"desiredPhase"`
	// plan is a list of groups of steps. The steps in a group are executed
	// concurrently, while the groups are executed sequentially. The order of the
	// groups is the order in which they are executed.
	Plan []TaskExecutionSpecPlan `json:"plan"`
	JSON taskExecutionSpecJSON   `json:"-"`
}

// taskExecutionSpecJSON contains the JSON metadata for the struct
// [TaskExecutionSpec]
type taskExecutionSpecJSON struct {
	DesiredPhase apijson.Field
	Plan         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TaskExecutionSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionSpecPlan struct {
	Steps []TaskExecutionSpecPlanStep `json:"steps"`
	JSON  taskExecutionSpecPlanJSON   `json:"-"`
}

// taskExecutionSpecPlanJSON contains the JSON metadata for the struct
// [TaskExecutionSpecPlan]
type taskExecutionSpecPlanJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecutionSpecPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecPlanJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionSpecPlanStep struct {
	// ID is the ID of the execution step
	ID        string                         `json:"id" format:"uuid"`
	DependsOn []string                       `json:"dependsOn"`
	Label     string                         `json:"label"`
	ServiceID string                         `json:"serviceId" format:"uuid"`
	Task      TaskExecutionSpecPlanStepsTask `json:"task"`
	JSON      taskExecutionSpecPlanStepJSON  `json:"-"`
}

// taskExecutionSpecPlanStepJSON contains the JSON metadata for the struct
// [TaskExecutionSpecPlanStep]
type taskExecutionSpecPlanStepJSON struct {
	ID          apijson.Field
	DependsOn   apijson.Field
	Label       apijson.Field
	ServiceID   apijson.Field
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecutionSpecPlanStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecPlanStepJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionSpecPlanStepsTask struct {
	ID   string                             `json:"id" format:"uuid"`
	Spec gitpod.TaskSpec                    `json:"spec"`
	JSON taskExecutionSpecPlanStepsTaskJSON `json:"-"`
}

// taskExecutionSpecPlanStepsTaskJSON contains the JSON metadata for the struct
// [TaskExecutionSpecPlanStepsTask]
type taskExecutionSpecPlanStepsTaskJSON struct {
	ID          apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecutionSpecPlanStepsTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecPlanStepsTaskJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionStatus struct {
	// failure_message summarises why the task execution failed to operate. If this is
	// non-empty the task execution has failed to operate and will likely transition to
	// a failed state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a task execution represents the aggregated phase of all steps.
	Phase TaskExecutionPhase `json:"phase"`
	// version of the status update. Task executions themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion"`
	// steps provides the status for each individual step of the task execution. If a
	// step is missing it has not yet started.
	Steps []TaskExecutionStatusStep `json:"steps"`
	JSON  taskExecutionStatusJSON   `json:"-"`
}

// taskExecutionStatusJSON contains the JSON metadata for the struct
// [TaskExecutionStatus]
type taskExecutionStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	StatusVersion  apijson.Field
	Steps          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TaskExecutionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionStatusJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the step failed to operate. If this is non-empty
	// the step has failed to operate and will likely transition to a failed state.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the execution step
	Phase TaskExecutionPhase          `json:"phase"`
	JSON  taskExecutionStatusStepJSON `json:"-"`
}

// taskExecutionStatusStepJSON contains the JSON metadata for the struct
// [TaskExecutionStatusStep]
type taskExecutionStatusStepJSON struct {
	ID             apijson.Field
	FailureMessage apijson.Field
	Phase          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TaskExecutionStatusStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionStatusStepJSON) RawJSON() string {
	return r.raw
}

type UserStatus string

const (
	UserStatusUserStatusUnspecified UserStatus = "USER_STATUS_UNSPECIFIED"
	UserStatusUserStatusActive      UserStatus = "USER_STATUS_ACTIVE"
	UserStatusUserStatusSuspended   UserStatus = "USER_STATUS_SUSPENDED"
	UserStatusUserStatusLeft        UserStatus = "USER_STATUS_LEFT"
)

func (r UserStatus) IsKnown() bool {
	switch r {
	case UserStatusUserStatusUnspecified, UserStatusUserStatusActive, UserStatusUserStatusSuspended, UserStatusUserStatusLeft:
		return true
	}
	return false
}
