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
func (r *EnvironmentAutomationTaskService) New(ctx context.Context, params EnvironmentAutomationTaskNewParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/CreateTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetTask
func (r *EnvironmentAutomationTaskService) Get(ctx context.Context, params EnvironmentAutomationTaskGetParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/GetTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
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

type EnvironmentAutomationTaskNewResponse struct {
	Task EnvironmentAutomationTaskNewResponseTask `json:"task"`
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

type EnvironmentAutomationTaskNewResponseTask struct {
	ID string `json:"id" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn     []string                                         `json:"dependsOn" format:"uuid"`
	EnvironmentID string                                           `json:"environmentId" format:"uuid"`
	Metadata      EnvironmentAutomationTaskNewResponseTaskMetadata `json:"metadata"`
	Spec          EnvironmentAutomationTaskNewResponseTaskSpec     `json:"spec"`
	JSON          environmentAutomationTaskNewResponseTaskJSON     `json:"-"`
}

// environmentAutomationTaskNewResponseTaskJSON contains the JSON metadata for the
// struct [EnvironmentAutomationTaskNewResponseTask]
type environmentAutomationTaskNewResponseTaskJSON struct {
	ID            apijson.Field
	DependsOn     apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskNewResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskNewResponseTaskJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskNewResponseTaskMetadata struct {
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
	// creator describes the principal who created the task.
	Creator EnvironmentAutomationTaskNewResponseTaskMetadataCreator `json:"creator"`
	// description is a user-facing description for the task. It can be used to provide
	// context and documentation for the task.
	Description string `json:"description"`
	// name is a user-facing name for the task. Unlike the reference, this field is not
	// unique, and not referenced by the system.
	//
	// This is a short descriptive name for the task.
	Name string `json:"name"`
	// reference is a user-facing identifier for the task which must be unique on the
	// environment.
	//
	// It is used to express dependencies between tasks, and to identify the task in
	// user interactions (e.g. the CLI).
	Reference string `json:"reference"`
	// triggered_by is a list of trigger that start the task.
	TriggeredBy []EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredBy `json:"triggeredBy"`
	JSON        environmentAutomationTaskNewResponseTaskMetadataJSON          `json:"-"`
}

// environmentAutomationTaskNewResponseTaskMetadataJSON contains the JSON metadata
// for the struct [EnvironmentAutomationTaskNewResponseTaskMetadata]
type environmentAutomationTaskNewResponseTaskMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskNewResponseTaskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskNewResponseTaskMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created the task.
type EnvironmentAutomationTaskNewResponseTaskMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationTaskNewResponseTaskMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationTaskNewResponseTaskMetadataCreatorJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskNewResponseTaskMetadataCreator]
type environmentAutomationTaskNewResponseTaskMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskNewResponseTaskMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskNewResponseTaskMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipal string

const (
	EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationTaskNewResponseTaskMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredBy struct {
	Manual                bool                                                            `json:"manual"`
	PostDevcontainerStart bool                                                            `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                            `json:"postEnvironmentStart"`
	JSON                  environmentAutomationTaskNewResponseTaskMetadataTriggeredByJSON `json:"-"`
	union                 EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByUnion
}

// environmentAutomationTaskNewResponseTaskMetadataTriggeredByJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredBy]
type environmentAutomationTaskNewResponseTaskMetadataTriggeredByJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r environmentAutomationTaskNewResponseTaskMetadataTriggeredByJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredBy) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredBy{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject],
// [EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject],
// [EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject],
// [EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject].
func (r EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredBy) AsUnion() EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByUnion {
	return r.union
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// Union satisfied by
// [EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject],
// [EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject],
// [EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject] or
// [EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject].
type EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByUnion interface {
	implementsEnvironmentAutomationTaskNewResponseTaskMetadataTriggeredBy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject{}),
		},
	)
}

type EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject struct {
	Manual                bool                                                                  `json:"manual,required"`
	PostDevcontainerStart bool                                                                  `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                                  `json:"postEnvironmentStart"`
	JSON                  environmentAutomationTaskNewResponseTaskMetadataTriggeredByObjectJSON `json:"-"`
}

// environmentAutomationTaskNewResponseTaskMetadataTriggeredByObjectJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject]
type environmentAutomationTaskNewResponseTaskMetadataTriggeredByObjectJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskNewResponseTaskMetadataTriggeredByObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationTaskNewResponseTaskMetadataTriggeredByObject) implementsEnvironmentAutomationTaskNewResponseTaskMetadataTriggeredBy() {
}

type EnvironmentAutomationTaskNewResponseTaskSpec struct {
	// command contains the command the task should execute
	Command string `json:"command"`
	// runs_on specifies the environment the task should run on.
	RunsOn EnvironmentAutomationTaskNewResponseTaskSpecRunsOn `json:"runsOn"`
	JSON   environmentAutomationTaskNewResponseTaskSpecJSON   `json:"-"`
}

// environmentAutomationTaskNewResponseTaskSpecJSON contains the JSON metadata for
// the struct [EnvironmentAutomationTaskNewResponseTaskSpec]
type environmentAutomationTaskNewResponseTaskSpecJSON struct {
	Command     apijson.Field
	RunsOn      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskNewResponseTaskSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskNewResponseTaskSpecJSON) RawJSON() string {
	return r.raw
}

// runs_on specifies the environment the task should run on.
type EnvironmentAutomationTaskNewResponseTaskSpecRunsOn struct {
	// This field can have the runtime type of
	// [EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDockerDocker].
	Docker interface{}                                            `json:"docker"`
	JSON   environmentAutomationTaskNewResponseTaskSpecRunsOnJSON `json:"-"`
	union  EnvironmentAutomationTaskNewResponseTaskSpecRunsOnUnion
}

// environmentAutomationTaskNewResponseTaskSpecRunsOnJSON contains the JSON
// metadata for the struct [EnvironmentAutomationTaskNewResponseTaskSpecRunsOn]
type environmentAutomationTaskNewResponseTaskSpecRunsOnJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentAutomationTaskNewResponseTaskSpecRunsOnJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationTaskNewResponseTaskSpecRunsOn) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationTaskNewResponseTaskSpecRunsOn{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EnvironmentAutomationTaskNewResponseTaskSpecRunsOnUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDocker],
// [EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDocker].
func (r EnvironmentAutomationTaskNewResponseTaskSpecRunsOn) AsUnion() EnvironmentAutomationTaskNewResponseTaskSpecRunsOnUnion {
	return r.union
}

// runs_on specifies the environment the task should run on.
//
// Union satisfied by [EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDocker] or
// [EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDocker].
type EnvironmentAutomationTaskNewResponseTaskSpecRunsOnUnion interface {
	implementsEnvironmentAutomationTaskNewResponseTaskSpecRunsOn()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskNewResponseTaskSpecRunsOnUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDocker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDocker{}),
		},
	)
}

type EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDocker struct {
	Docker EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDockerDocker `json:"docker,required"`
	JSON   environmentAutomationTaskNewResponseTaskSpecRunsOnDockerJSON   `json:"-"`
}

// environmentAutomationTaskNewResponseTaskSpecRunsOnDockerJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDocker]
type environmentAutomationTaskNewResponseTaskSpecRunsOnDockerJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskNewResponseTaskSpecRunsOnDockerJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDocker) implementsEnvironmentAutomationTaskNewResponseTaskSpecRunsOn() {
}

type EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDockerDocker struct {
	Environment []string                                                           `json:"environment"`
	Image       string                                                             `json:"image"`
	JSON        environmentAutomationTaskNewResponseTaskSpecRunsOnDockerDockerJSON `json:"-"`
}

// environmentAutomationTaskNewResponseTaskSpecRunsOnDockerDockerJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDockerDocker]
type environmentAutomationTaskNewResponseTaskSpecRunsOnDockerDockerJSON struct {
	Environment apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskNewResponseTaskSpecRunsOnDockerDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskNewResponseTaskSpecRunsOnDockerDockerJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskGetResponse struct {
	Task EnvironmentAutomationTaskGetResponseTask `json:"task"`
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

type EnvironmentAutomationTaskGetResponseTask struct {
	ID string `json:"id" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn     []string                                         `json:"dependsOn" format:"uuid"`
	EnvironmentID string                                           `json:"environmentId" format:"uuid"`
	Metadata      EnvironmentAutomationTaskGetResponseTaskMetadata `json:"metadata"`
	Spec          EnvironmentAutomationTaskGetResponseTaskSpec     `json:"spec"`
	JSON          environmentAutomationTaskGetResponseTaskJSON     `json:"-"`
}

// environmentAutomationTaskGetResponseTaskJSON contains the JSON metadata for the
// struct [EnvironmentAutomationTaskGetResponseTask]
type environmentAutomationTaskGetResponseTaskJSON struct {
	ID            apijson.Field
	DependsOn     apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskGetResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskGetResponseTaskJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskGetResponseTaskMetadata struct {
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
	// creator describes the principal who created the task.
	Creator EnvironmentAutomationTaskGetResponseTaskMetadataCreator `json:"creator"`
	// description is a user-facing description for the task. It can be used to provide
	// context and documentation for the task.
	Description string `json:"description"`
	// name is a user-facing name for the task. Unlike the reference, this field is not
	// unique, and not referenced by the system.
	//
	// This is a short descriptive name for the task.
	Name string `json:"name"`
	// reference is a user-facing identifier for the task which must be unique on the
	// environment.
	//
	// It is used to express dependencies between tasks, and to identify the task in
	// user interactions (e.g. the CLI).
	Reference string `json:"reference"`
	// triggered_by is a list of trigger that start the task.
	TriggeredBy []EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredBy `json:"triggeredBy"`
	JSON        environmentAutomationTaskGetResponseTaskMetadataJSON          `json:"-"`
}

// environmentAutomationTaskGetResponseTaskMetadataJSON contains the JSON metadata
// for the struct [EnvironmentAutomationTaskGetResponseTaskMetadata]
type environmentAutomationTaskGetResponseTaskMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskGetResponseTaskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskGetResponseTaskMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created the task.
type EnvironmentAutomationTaskGetResponseTaskMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationTaskGetResponseTaskMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationTaskGetResponseTaskMetadataCreatorJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskGetResponseTaskMetadataCreator]
type environmentAutomationTaskGetResponseTaskMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskGetResponseTaskMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskGetResponseTaskMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipal string

const (
	EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationTaskGetResponseTaskMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredBy struct {
	Manual                bool                                                            `json:"manual"`
	PostDevcontainerStart bool                                                            `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                            `json:"postEnvironmentStart"`
	JSON                  environmentAutomationTaskGetResponseTaskMetadataTriggeredByJSON `json:"-"`
	union                 EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByUnion
}

// environmentAutomationTaskGetResponseTaskMetadataTriggeredByJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredBy]
type environmentAutomationTaskGetResponseTaskMetadataTriggeredByJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r environmentAutomationTaskGetResponseTaskMetadataTriggeredByJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredBy) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredBy{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject],
// [EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject],
// [EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject],
// [EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject].
func (r EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredBy) AsUnion() EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByUnion {
	return r.union
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// Union satisfied by
// [EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject],
// [EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject],
// [EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject] or
// [EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject].
type EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByUnion interface {
	implementsEnvironmentAutomationTaskGetResponseTaskMetadataTriggeredBy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject{}),
		},
	)
}

type EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject struct {
	Manual                bool                                                                  `json:"manual,required"`
	PostDevcontainerStart bool                                                                  `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                                  `json:"postEnvironmentStart"`
	JSON                  environmentAutomationTaskGetResponseTaskMetadataTriggeredByObjectJSON `json:"-"`
}

// environmentAutomationTaskGetResponseTaskMetadataTriggeredByObjectJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject]
type environmentAutomationTaskGetResponseTaskMetadataTriggeredByObjectJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskGetResponseTaskMetadataTriggeredByObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationTaskGetResponseTaskMetadataTriggeredByObject) implementsEnvironmentAutomationTaskGetResponseTaskMetadataTriggeredBy() {
}

type EnvironmentAutomationTaskGetResponseTaskSpec struct {
	// command contains the command the task should execute
	Command string `json:"command"`
	// runs_on specifies the environment the task should run on.
	RunsOn EnvironmentAutomationTaskGetResponseTaskSpecRunsOn `json:"runsOn"`
	JSON   environmentAutomationTaskGetResponseTaskSpecJSON   `json:"-"`
}

// environmentAutomationTaskGetResponseTaskSpecJSON contains the JSON metadata for
// the struct [EnvironmentAutomationTaskGetResponseTaskSpec]
type environmentAutomationTaskGetResponseTaskSpecJSON struct {
	Command     apijson.Field
	RunsOn      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskGetResponseTaskSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskGetResponseTaskSpecJSON) RawJSON() string {
	return r.raw
}

// runs_on specifies the environment the task should run on.
type EnvironmentAutomationTaskGetResponseTaskSpecRunsOn struct {
	// This field can have the runtime type of
	// [EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDockerDocker].
	Docker interface{}                                            `json:"docker"`
	JSON   environmentAutomationTaskGetResponseTaskSpecRunsOnJSON `json:"-"`
	union  EnvironmentAutomationTaskGetResponseTaskSpecRunsOnUnion
}

// environmentAutomationTaskGetResponseTaskSpecRunsOnJSON contains the JSON
// metadata for the struct [EnvironmentAutomationTaskGetResponseTaskSpecRunsOn]
type environmentAutomationTaskGetResponseTaskSpecRunsOnJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentAutomationTaskGetResponseTaskSpecRunsOnJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationTaskGetResponseTaskSpecRunsOn) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationTaskGetResponseTaskSpecRunsOn{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EnvironmentAutomationTaskGetResponseTaskSpecRunsOnUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDocker],
// [EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDocker].
func (r EnvironmentAutomationTaskGetResponseTaskSpecRunsOn) AsUnion() EnvironmentAutomationTaskGetResponseTaskSpecRunsOnUnion {
	return r.union
}

// runs_on specifies the environment the task should run on.
//
// Union satisfied by [EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDocker] or
// [EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDocker].
type EnvironmentAutomationTaskGetResponseTaskSpecRunsOnUnion interface {
	implementsEnvironmentAutomationTaskGetResponseTaskSpecRunsOn()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskGetResponseTaskSpecRunsOnUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDocker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDocker{}),
		},
	)
}

type EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDocker struct {
	Docker EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDockerDocker `json:"docker,required"`
	JSON   environmentAutomationTaskGetResponseTaskSpecRunsOnDockerJSON   `json:"-"`
}

// environmentAutomationTaskGetResponseTaskSpecRunsOnDockerJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDocker]
type environmentAutomationTaskGetResponseTaskSpecRunsOnDockerJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskGetResponseTaskSpecRunsOnDockerJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDocker) implementsEnvironmentAutomationTaskGetResponseTaskSpecRunsOn() {
}

type EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDockerDocker struct {
	Environment []string                                                           `json:"environment"`
	Image       string                                                             `json:"image"`
	JSON        environmentAutomationTaskGetResponseTaskSpecRunsOnDockerDockerJSON `json:"-"`
}

// environmentAutomationTaskGetResponseTaskSpecRunsOnDockerDockerJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDockerDocker]
type environmentAutomationTaskGetResponseTaskSpecRunsOnDockerDockerJSON struct {
	Environment apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskGetResponseTaskSpecRunsOnDockerDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskGetResponseTaskSpecRunsOnDockerDockerJSON) RawJSON() string {
	return r.raw
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
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
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
	// creator describes the principal who created the task.
	Creator EnvironmentAutomationTaskListResponseTasksMetadataCreator `json:"creator"`
	// description is a user-facing description for the task. It can be used to provide
	// context and documentation for the task.
	Description string `json:"description"`
	// name is a user-facing name for the task. Unlike the reference, this field is not
	// unique, and not referenced by the system.
	//
	// This is a short descriptive name for the task.
	Name string `json:"name"`
	// reference is a user-facing identifier for the task which must be unique on the
	// environment.
	//
	// It is used to express dependencies between tasks, and to identify the task in
	// user interactions (e.g. the CLI).
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
	Manual                bool                                                              `json:"manual"`
	PostDevcontainerStart bool                                                              `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                              `json:"postEnvironmentStart"`
	JSON                  environmentAutomationTaskListResponseTasksMetadataTriggeredByJSON `json:"-"`
	union                 EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByUnion
}

// environmentAutomationTaskListResponseTasksMetadataTriggeredByJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredBy]
type environmentAutomationTaskListResponseTasksMetadataTriggeredByJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r environmentAutomationTaskListResponseTasksMetadataTriggeredByJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationTaskListResponseTasksMetadataTriggeredBy) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationTaskListResponseTasksMetadataTriggeredBy{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject],
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject],
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject],
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject].
func (r EnvironmentAutomationTaskListResponseTasksMetadataTriggeredBy) AsUnion() EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByUnion {
	return r.union
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// Union satisfied by
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject],
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject],
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject] or
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject].
type EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByUnion interface {
	implementsEnvironmentAutomationTaskListResponseTasksMetadataTriggeredBy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject{}),
		},
	)
}

type EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject struct {
	Manual                bool                                                                    `json:"manual,required"`
	PostDevcontainerStart bool                                                                    `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                                    `json:"postEnvironmentStart"`
	JSON                  environmentAutomationTaskListResponseTasksMetadataTriggeredByObjectJSON `json:"-"`
}

// environmentAutomationTaskListResponseTasksMetadataTriggeredByObjectJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject]
type environmentAutomationTaskListResponseTasksMetadataTriggeredByObjectJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskListResponseTasksMetadataTriggeredByObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationTaskListResponseTasksMetadataTriggeredByObject) implementsEnvironmentAutomationTaskListResponseTasksMetadataTriggeredBy() {
}

type EnvironmentAutomationTaskListResponseTasksSpec struct {
	// command contains the command the task should execute
	Command string `json:"command"`
	// runs_on specifies the environment the task should run on.
	RunsOn EnvironmentAutomationTaskListResponseTasksSpecRunsOn `json:"runsOn"`
	JSON   environmentAutomationTaskListResponseTasksSpecJSON   `json:"-"`
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
type EnvironmentAutomationTaskListResponseTasksSpecRunsOn struct {
	// This field can have the runtime type of
	// [EnvironmentAutomationTaskListResponseTasksSpecRunsOnDockerDocker].
	Docker interface{}                                              `json:"docker"`
	JSON   environmentAutomationTaskListResponseTasksSpecRunsOnJSON `json:"-"`
	union  EnvironmentAutomationTaskListResponseTasksSpecRunsOnUnion
}

// environmentAutomationTaskListResponseTasksSpecRunsOnJSON contains the JSON
// metadata for the struct [EnvironmentAutomationTaskListResponseTasksSpecRunsOn]
type environmentAutomationTaskListResponseTasksSpecRunsOnJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentAutomationTaskListResponseTasksSpecRunsOnJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationTaskListResponseTasksSpecRunsOn) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationTaskListResponseTasksSpecRunsOn{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EnvironmentAutomationTaskListResponseTasksSpecRunsOnUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationTaskListResponseTasksSpecRunsOnDocker],
// [EnvironmentAutomationTaskListResponseTasksSpecRunsOnDocker].
func (r EnvironmentAutomationTaskListResponseTasksSpecRunsOn) AsUnion() EnvironmentAutomationTaskListResponseTasksSpecRunsOnUnion {
	return r.union
}

// runs_on specifies the environment the task should run on.
//
// Union satisfied by [EnvironmentAutomationTaskListResponseTasksSpecRunsOnDocker]
// or [EnvironmentAutomationTaskListResponseTasksSpecRunsOnDocker].
type EnvironmentAutomationTaskListResponseTasksSpecRunsOnUnion interface {
	implementsEnvironmentAutomationTaskListResponseTasksSpecRunsOn()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskListResponseTasksSpecRunsOnUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskListResponseTasksSpecRunsOnDocker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskListResponseTasksSpecRunsOnDocker{}),
		},
	)
}

type EnvironmentAutomationTaskListResponseTasksSpecRunsOnDocker struct {
	Docker EnvironmentAutomationTaskListResponseTasksSpecRunsOnDockerDocker `json:"docker,required"`
	JSON   environmentAutomationTaskListResponseTasksSpecRunsOnDockerJSON   `json:"-"`
}

// environmentAutomationTaskListResponseTasksSpecRunsOnDockerJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationTaskListResponseTasksSpecRunsOnDocker]
type environmentAutomationTaskListResponseTasksSpecRunsOnDockerJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskListResponseTasksSpecRunsOnDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskListResponseTasksSpecRunsOnDockerJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationTaskListResponseTasksSpecRunsOnDocker) implementsEnvironmentAutomationTaskListResponseTasksSpecRunsOn() {
}

type EnvironmentAutomationTaskListResponseTasksSpecRunsOnDockerDocker struct {
	Environment []string                                                             `json:"environment"`
	Image       string                                                               `json:"image"`
	JSON        environmentAutomationTaskListResponseTasksSpecRunsOnDockerDockerJSON `json:"-"`
}

// environmentAutomationTaskListResponseTasksSpecRunsOnDockerDockerJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationTaskListResponseTasksSpecRunsOnDockerDocker]
type environmentAutomationTaskListResponseTasksSpecRunsOnDockerDockerJSON struct {
	Environment apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskListResponseTasksSpecRunsOnDockerDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskListResponseTasksSpecRunsOnDockerDockerJSON) RawJSON() string {
	return r.raw
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
	Creator EnvironmentAutomationTaskStartResponseTaskExecutionMetadataCreator `json:"creator"`
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
	// concurrently, while the groups are executed sequentially.
	//
	// The order of the groups is the order in which they are executed.
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
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// This field can have the runtime type of [[]string].
	DependsOn interface{} `json:"dependsOn"`
	Label     string      `json:"label"`
	ServiceID string      `json:"serviceId" format:"uuid"`
	// This field can have the runtime type of
	// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTask].
	Task  interface{}                                                         `json:"task"`
	JSON  environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepJSON `json:"-"`
	union EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsUnion
}

// environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStep]
type environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepJSON struct {
	ID          apijson.Field
	DependsOn   apijson.Field
	Label       apijson.Field
	ServiceID   apijson.Field
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStep) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStep{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject],
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject],
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject].
func (r EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStep) AsUnion() EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsUnion {
	return r.union
}

// Union satisfied by
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject],
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject] or
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject].
type EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsUnion interface {
	implementsEnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStep()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject{}),
		},
	)
}

type EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject struct {
	ServiceID string `json:"serviceId,required" format:"uuid"`
	// ID is the ID of the execution step
	ID        string                                                                     `json:"id" format:"uuid"`
	DependsOn []string                                                                   `json:"dependsOn"`
	Label     string                                                                     `json:"label"`
	Task      EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTask `json:"task"`
	JSON      environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectJSON `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject]
type environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectJSON struct {
	ServiceID   apijson.Field
	ID          apijson.Field
	DependsOn   apijson.Field
	Label       apijson.Field
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObject) implementsEnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStep() {
}

type EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTask struct {
	ID   string                                                                         `json:"id" format:"uuid"`
	Spec EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpec `json:"spec"`
	JSON environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskJSON `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTask]
type environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskJSON struct {
	ID          apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpec struct {
	// command contains the command the task should execute
	Command string `json:"command"`
	// runs_on specifies the environment the task should run on.
	RunsOn EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOn `json:"runsOn"`
	JSON   environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecJSON   `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpec]
type environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecJSON struct {
	Command     apijson.Field
	RunsOn      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecJSON) RawJSON() string {
	return r.raw
}

// runs_on specifies the environment the task should run on.
type EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOn struct {
	// This field can have the runtime type of
	// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerDocker].
	Docker interface{}                                                                              `json:"docker"`
	JSON   environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnJSON `json:"-"`
	union  EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnUnion
}

// environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOn]
type environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOn) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOn{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDocker],
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDocker].
func (r EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOn) AsUnion() EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnUnion {
	return r.union
}

// runs_on specifies the environment the task should run on.
//
// Union satisfied by
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDocker]
// or
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDocker].
type EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnUnion interface {
	implementsEnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOn()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDocker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDocker{}),
		},
	)
}

type EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDocker struct {
	Docker EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerDocker `json:"docker,required"`
	JSON   environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerJSON   `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDocker]
type environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDocker) implementsEnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOn() {
}

type EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerDocker struct {
	Environment []string                                                                                             `json:"environment"`
	Image       string                                                                                               `json:"image"`
	JSON        environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerDockerJSON `json:"-"`
}

// environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerDockerJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerDocker]
type environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerDockerJSON struct {
	Environment apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseTaskExecutionSpecPlanStepsObjectTaskSpecRunsOnDockerDockerJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskStartResponseTaskExecutionStatus struct {
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
	Phase EnvironmentAutomationTaskStartResponseTaskExecutionStatusPhase `json:"phase"`
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

type EnvironmentAutomationTaskStartResponseTaskExecutionStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the step failed to operate. If this is non-empty
	//
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

type EnvironmentAutomationTaskNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	DependsOn              param.Field[[]string]                                                 `json:"dependsOn" format:"uuid"`
	EnvironmentID          param.Field[string]                                                   `json:"environmentId" format:"uuid"`
	Metadata               param.Field[EnvironmentAutomationTaskNewParamsMetadata]               `json:"metadata"`
	Spec                   param.Field[EnvironmentAutomationTaskNewParamsSpec]                   `json:"spec"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationTaskNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskNewParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskNewParamsConnectProtocolVersion1 EnvironmentAutomationTaskNewParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentAutomationTaskNewParamsMetadata struct {
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
	CreatedAt param.Field[time.Time] `json:"createdAt" format:"date-time"`
	// creator describes the principal who created the task.
	Creator param.Field[EnvironmentAutomationTaskNewParamsMetadataCreator] `json:"creator"`
	// description is a user-facing description for the task. It can be used to provide
	// context and documentation for the task.
	Description param.Field[string] `json:"description"`
	// name is a user-facing name for the task. Unlike the reference, this field is not
	// unique, and not referenced by the system.
	//
	// This is a short descriptive name for the task.
	Name param.Field[string] `json:"name"`
	// reference is a user-facing identifier for the task which must be unique on the
	// environment.
	//
	// It is used to express dependencies between tasks, and to identify the task in
	// user interactions (e.g. the CLI).
	Reference param.Field[string] `json:"reference"`
	// triggered_by is a list of trigger that start the task.
	TriggeredBy param.Field[[]EnvironmentAutomationTaskNewParamsMetadataTriggeredByUnion] `json:"triggeredBy"`
}

func (r EnvironmentAutomationTaskNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// creator describes the principal who created the task.
type EnvironmentAutomationTaskNewParamsMetadataCreator struct {
	// id is the UUID of the subject
	ID param.Field[string] `json:"id"`
	// Principal is the principal of the subject
	Principal param.Field[EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipal] `json:"principal"`
}

func (r EnvironmentAutomationTaskNewParamsMetadataCreator) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Principal is the principal of the subject
type EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipal string

const (
	EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type EnvironmentAutomationTaskNewParamsMetadataTriggeredBy struct {
	Manual                param.Field[bool] `json:"manual"`
	PostDevcontainerStart param.Field[bool] `json:"postDevcontainerStart"`
	PostEnvironmentStart  param.Field[bool] `json:"postEnvironmentStart"`
}

func (r EnvironmentAutomationTaskNewParamsMetadataTriggeredBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationTaskNewParamsMetadataTriggeredBy) implementsEnvironmentAutomationTaskNewParamsMetadataTriggeredByUnion() {
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// Satisfied by [EnvironmentAutomationTaskNewParamsMetadataTriggeredByObject],
// [EnvironmentAutomationTaskNewParamsMetadataTriggeredByObject],
// [EnvironmentAutomationTaskNewParamsMetadataTriggeredByObject],
// [EnvironmentAutomationTaskNewParamsMetadataTriggeredByObject],
// [EnvironmentAutomationTaskNewParamsMetadataTriggeredBy].
type EnvironmentAutomationTaskNewParamsMetadataTriggeredByUnion interface {
	implementsEnvironmentAutomationTaskNewParamsMetadataTriggeredByUnion()
}

type EnvironmentAutomationTaskNewParamsMetadataTriggeredByObject struct {
	Manual                param.Field[bool] `json:"manual,required"`
	PostDevcontainerStart param.Field[bool] `json:"postDevcontainerStart"`
	PostEnvironmentStart  param.Field[bool] `json:"postEnvironmentStart"`
}

func (r EnvironmentAutomationTaskNewParamsMetadataTriggeredByObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationTaskNewParamsMetadataTriggeredByObject) implementsEnvironmentAutomationTaskNewParamsMetadataTriggeredByUnion() {
}

type EnvironmentAutomationTaskNewParamsSpec struct {
	// command contains the command the task should execute
	Command param.Field[string] `json:"command"`
	// runs_on specifies the environment the task should run on.
	RunsOn param.Field[EnvironmentAutomationTaskNewParamsSpecRunsOnUnion] `json:"runsOn"`
}

func (r EnvironmentAutomationTaskNewParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// runs_on specifies the environment the task should run on.
type EnvironmentAutomationTaskNewParamsSpecRunsOn struct {
	Docker param.Field[interface{}] `json:"docker"`
}

func (r EnvironmentAutomationTaskNewParamsSpecRunsOn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationTaskNewParamsSpecRunsOn) implementsEnvironmentAutomationTaskNewParamsSpecRunsOnUnion() {
}

// runs_on specifies the environment the task should run on.
//
// Satisfied by [EnvironmentAutomationTaskNewParamsSpecRunsOnDocker],
// [EnvironmentAutomationTaskNewParamsSpecRunsOnDocker],
// [EnvironmentAutomationTaskNewParamsSpecRunsOn].
type EnvironmentAutomationTaskNewParamsSpecRunsOnUnion interface {
	implementsEnvironmentAutomationTaskNewParamsSpecRunsOnUnion()
}

type EnvironmentAutomationTaskNewParamsSpecRunsOnDocker struct {
	Docker param.Field[EnvironmentAutomationTaskNewParamsSpecRunsOnDockerDocker] `json:"docker,required"`
}

func (r EnvironmentAutomationTaskNewParamsSpecRunsOnDocker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationTaskNewParamsSpecRunsOnDocker) implementsEnvironmentAutomationTaskNewParamsSpecRunsOnUnion() {
}

type EnvironmentAutomationTaskNewParamsSpecRunsOnDockerDocker struct {
	Environment param.Field[[]string] `json:"environment"`
	Image       param.Field[string]   `json:"image"`
}

func (r EnvironmentAutomationTaskNewParamsSpecRunsOnDockerDocker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskGetParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[EnvironmentAutomationTaskGetParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[EnvironmentAutomationTaskGetParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[EnvironmentAutomationTaskGetParamsConnect] `query:"connect"`
	Message param.Field[string]                                    `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [EnvironmentAutomationTaskGetParams]'s query parameters as
// `url.Values`.
func (r EnvironmentAutomationTaskGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type EnvironmentAutomationTaskGetParamsEncoding string

const (
	EnvironmentAutomationTaskGetParamsEncodingProto EnvironmentAutomationTaskGetParamsEncoding = "proto"
	EnvironmentAutomationTaskGetParamsEncodingJson  EnvironmentAutomationTaskGetParamsEncoding = "json"
)

func (r EnvironmentAutomationTaskGetParamsEncoding) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskGetParamsEncodingProto, EnvironmentAutomationTaskGetParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskGetParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationTaskGetParamsConnectProtocolVersion1 EnvironmentAutomationTaskGetParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationTaskGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type EnvironmentAutomationTaskGetParamsCompression string

const (
	EnvironmentAutomationTaskGetParamsCompressionIdentity EnvironmentAutomationTaskGetParamsCompression = "identity"
	EnvironmentAutomationTaskGetParamsCompressionGzip     EnvironmentAutomationTaskGetParamsCompression = "gzip"
	EnvironmentAutomationTaskGetParamsCompressionBr       EnvironmentAutomationTaskGetParamsCompression = "br"
)

func (r EnvironmentAutomationTaskGetParamsCompression) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskGetParamsCompressionIdentity, EnvironmentAutomationTaskGetParamsCompressionGzip, EnvironmentAutomationTaskGetParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskGetParamsConnect string

const (
	EnvironmentAutomationTaskGetParamsConnectV1 EnvironmentAutomationTaskGetParamsConnect = "v1"
)

func (r EnvironmentAutomationTaskGetParamsConnect) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskGetParamsConnectV1:
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
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[EnvironmentAutomationTaskListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationTaskListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[EnvironmentAutomationTaskListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[EnvironmentAutomationTaskListParamsConnect] `query:"connect"`
	Message param.Field[string]                                     `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [EnvironmentAutomationTaskListParams]'s query parameters as
// `url.Values`.
func (r EnvironmentAutomationTaskListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type EnvironmentAutomationTaskListParamsEncoding string

const (
	EnvironmentAutomationTaskListParamsEncodingProto EnvironmentAutomationTaskListParamsEncoding = "proto"
	EnvironmentAutomationTaskListParamsEncodingJson  EnvironmentAutomationTaskListParamsEncoding = "json"
)

func (r EnvironmentAutomationTaskListParamsEncoding) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskListParamsEncodingProto, EnvironmentAutomationTaskListParamsEncodingJson:
		return true
	}
	return false
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

// Which compression algorithm to use for this request
type EnvironmentAutomationTaskListParamsCompression string

const (
	EnvironmentAutomationTaskListParamsCompressionIdentity EnvironmentAutomationTaskListParamsCompression = "identity"
	EnvironmentAutomationTaskListParamsCompressionGzip     EnvironmentAutomationTaskListParamsCompression = "gzip"
	EnvironmentAutomationTaskListParamsCompressionBr       EnvironmentAutomationTaskListParamsCompression = "br"
)

func (r EnvironmentAutomationTaskListParamsCompression) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskListParamsCompressionIdentity, EnvironmentAutomationTaskListParamsCompressionGzip, EnvironmentAutomationTaskListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EnvironmentAutomationTaskListParamsConnect string

const (
	EnvironmentAutomationTaskListParamsConnectV1 EnvironmentAutomationTaskListParamsConnect = "v1"
)

func (r EnvironmentAutomationTaskListParamsConnect) IsKnown() bool {
	switch r {
	case EnvironmentAutomationTaskListParamsConnectV1:
		return true
	}
	return false
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
