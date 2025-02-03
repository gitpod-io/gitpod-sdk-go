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

// EnvironmentAutomationServiceService contains methods and other services that
// help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentAutomationServiceService] method instead.
type EnvironmentAutomationServiceService struct {
	Options []option.RequestOption
}

// NewEnvironmentAutomationServiceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEnvironmentAutomationServiceService(opts ...option.RequestOption) (r *EnvironmentAutomationServiceService) {
	r = &EnvironmentAutomationServiceService{}
	r.Options = opts
	return
}

// CreateService
func (r *EnvironmentAutomationServiceService) New(ctx context.Context, params EnvironmentAutomationServiceNewParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/CreateService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetService
func (r *EnvironmentAutomationServiceService) Get(ctx context.Context, params EnvironmentAutomationServiceGetParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/GetService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// UpdateService
func (r *EnvironmentAutomationServiceService) Update(ctx context.Context, params EnvironmentAutomationServiceUpdateParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/UpdateService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListServices
func (r *EnvironmentAutomationServiceService) List(ctx context.Context, params EnvironmentAutomationServiceListParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/ListServices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// DeleteService deletes a service. This call does not block until the service is
// deleted.
//
// If the service is not stopped it will be stopped before deletion.
func (r *EnvironmentAutomationServiceService) Delete(ctx context.Context, params EnvironmentAutomationServiceDeleteParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/DeleteService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// StartService starts a service. This call does not block until the service is
// started.
//
// This call will not error if the service is already running or has been started.
func (r *EnvironmentAutomationServiceService) Start(ctx context.Context, params EnvironmentAutomationServiceStartParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceStartResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StartService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// StopService stops a service. This call does not block until the service is
// stopped.
//
// This call will not error if the service is already stopped or has been stopped.
func (r *EnvironmentAutomationServiceService) Stop(ctx context.Context, params EnvironmentAutomationServiceStopParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceStopResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StopService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type EnvironmentAutomationServiceNewResponse struct {
	Service EnvironmentAutomationServiceNewResponseService `json:"service"`
	JSON    environmentAutomationServiceNewResponseJSON    `json:"-"`
}

// environmentAutomationServiceNewResponseJSON contains the JSON metadata for the
// struct [EnvironmentAutomationServiceNewResponse]
type environmentAutomationServiceNewResponseJSON struct {
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceNewResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceNewResponseService struct {
	ID            string                                                 `json:"id" format:"uuid"`
	EnvironmentID string                                                 `json:"environmentId" format:"uuid"`
	Metadata      EnvironmentAutomationServiceNewResponseServiceMetadata `json:"metadata"`
	Spec          EnvironmentAutomationServiceNewResponseServiceSpec     `json:"spec"`
	Status        EnvironmentAutomationServiceNewResponseServiceStatus   `json:"status"`
	JSON          environmentAutomationServiceNewResponseServiceJSON     `json:"-"`
}

// environmentAutomationServiceNewResponseServiceJSON contains the JSON metadata
// for the struct [EnvironmentAutomationServiceNewResponseService]
type environmentAutomationServiceNewResponseServiceJSON struct {
	ID            apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceNewResponseService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceNewResponseServiceJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceNewResponseServiceMetadata struct {
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
	// creator describes the principal who created the service.
	Creator EnvironmentAutomationServiceNewResponseServiceMetadataCreator `json:"creator"`
	// description is a user-facing description for the service. It can be used to
	// provide context and documentation for the service.
	Description string `json:"description"`
	// name is a user-facing name for the service. Unlike the reference, this field is
	// not unique, and not referenced by the system.
	//
	// This is a short descriptive name for the service.
	Name string `json:"name"`
	// reference is a user-facing identifier for the service which must be unique on
	// the environment.
	//
	// It is used to express dependencies between services, and to identify the service
	// in user interactions (e.g. the CLI).
	Reference string `json:"reference"`
	// triggered_by is a list of trigger that start the service.
	TriggeredBy []EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredBy `json:"triggeredBy"`
	JSON        environmentAutomationServiceNewResponseServiceMetadataJSON          `json:"-"`
}

// environmentAutomationServiceNewResponseServiceMetadataJSON contains the JSON
// metadata for the struct [EnvironmentAutomationServiceNewResponseServiceMetadata]
type environmentAutomationServiceNewResponseServiceMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceNewResponseServiceMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceNewResponseServiceMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created the service.
type EnvironmentAutomationServiceNewResponseServiceMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationServiceNewResponseServiceMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationServiceNewResponseServiceMetadataCreatorJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationServiceNewResponseServiceMetadataCreator]
type environmentAutomationServiceNewResponseServiceMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceNewResponseServiceMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceNewResponseServiceMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipal string

const (
	EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationServiceNewResponseServiceMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredBy struct {
	Manual                bool                                                                  `json:"manual"`
	PostDevcontainerStart bool                                                                  `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                                  `json:"postEnvironmentStart"`
	JSON                  environmentAutomationServiceNewResponseServiceMetadataTriggeredByJSON `json:"-"`
	union                 EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByUnion
}

// environmentAutomationServiceNewResponseServiceMetadataTriggeredByJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredBy]
type environmentAutomationServiceNewResponseServiceMetadataTriggeredByJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r environmentAutomationServiceNewResponseServiceMetadataTriggeredByJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredBy) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredBy{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject],
// [EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject],
// [EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject],
// [EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject].
func (r EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredBy) AsUnion() EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByUnion {
	return r.union
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// Union satisfied by
// [EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject],
// [EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject],
// [EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject] or
// [EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject].
type EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByUnion interface {
	implementsEnvironmentAutomationServiceNewResponseServiceMetadataTriggeredBy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject{}),
		},
	)
}

type EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject struct {
	Manual                bool                                                                        `json:"manual,required"`
	PostDevcontainerStart bool                                                                        `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                                        `json:"postEnvironmentStart"`
	JSON                  environmentAutomationServiceNewResponseServiceMetadataTriggeredByObjectJSON `json:"-"`
}

// environmentAutomationServiceNewResponseServiceMetadataTriggeredByObjectJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject]
type environmentAutomationServiceNewResponseServiceMetadataTriggeredByObjectJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceNewResponseServiceMetadataTriggeredByObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationServiceNewResponseServiceMetadataTriggeredByObject) implementsEnvironmentAutomationServiceNewResponseServiceMetadataTriggeredBy() {
}

type EnvironmentAutomationServiceNewResponseServiceSpec struct {
	// commands contains the commands to start, stop and check the readiness of the
	// service
	Commands EnvironmentAutomationServiceNewResponseServiceSpecCommands `json:"commands"`
	// desired_phase is the phase the service should be in. Used to start or stop the
	// service.
	DesiredPhase EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhase `json:"desiredPhase"`
	// runs_on specifies the environment the service should run on.
	RunsOn EnvironmentAutomationServiceNewResponseServiceSpecRunsOn `json:"runsOn"`
	// session should be changed to trigger a restart of the service. If a service
	// exits it will
	//
	// not be restarted until the session is changed.
	Session string `json:"session"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion string                                                 `json:"specVersion" format:"int64"`
	JSON        environmentAutomationServiceNewResponseServiceSpecJSON `json:"-"`
}

// environmentAutomationServiceNewResponseServiceSpecJSON contains the JSON
// metadata for the struct [EnvironmentAutomationServiceNewResponseServiceSpec]
type environmentAutomationServiceNewResponseServiceSpecJSON struct {
	Commands     apijson.Field
	DesiredPhase apijson.Field
	RunsOn       apijson.Field
	Session      apijson.Field
	SpecVersion  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceNewResponseServiceSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceNewResponseServiceSpecJSON) RawJSON() string {
	return r.raw
}

// commands contains the commands to start, stop and check the readiness of the
// service
type EnvironmentAutomationServiceNewResponseServiceSpecCommands struct {
	// ready is an optional command that is run repeatedly until it exits with a zero
	// exit code.
	//
	// If set, the service will first go into a Starting phase, and then into a Running
	// phase once the ready command exits with a zero exit code.
	Ready string `json:"ready"`
	// start is the command to start and run the service. If start exits, the service
	// will transition to the following phase:
	//
	//   - Stopped: if the exit code is 0
	//   - Failed: if the exit code is not 0 If the stop command is not set, the start
	//     command will receive a SIGTERM signal when the service is requested to stop.
	//     If it does not exit within 2 minutes, it will receive a SIGKILL signal.
	Start string `json:"start"`
	// stop is an optional command that runs when the service is requested to stop.
	//
	// If set, instead of sending a SIGTERM signal to the start command, the stop
	// command will be run. Once the stop command exits, the start command will receive
	// a SIGKILL signal. If the stop command exits with a non-zero exit code, the
	// service will transition to the Failed phase. If the stop command does not exit
	// within 2 minutes, a SIGKILL signal will be sent to both the start and stop
	// commands.
	Stop string                                                         `json:"stop"`
	JSON environmentAutomationServiceNewResponseServiceSpecCommandsJSON `json:"-"`
}

// environmentAutomationServiceNewResponseServiceSpecCommandsJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationServiceNewResponseServiceSpecCommands]
type environmentAutomationServiceNewResponseServiceSpecCommandsJSON struct {
	Ready       apijson.Field
	Start       apijson.Field
	Stop        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceNewResponseServiceSpecCommands) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceNewResponseServiceSpecCommandsJSON) RawJSON() string {
	return r.raw
}

// desired_phase is the phase the service should be in. Used to start or stop the
// service.
type EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhase string

const (
	EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseUnspecified EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhase = "SERVICE_PHASE_UNSPECIFIED"
	EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseStarting    EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhase = "SERVICE_PHASE_STARTING"
	EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseRunning     EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhase = "SERVICE_PHASE_RUNNING"
	EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseStopping    EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhase = "SERVICE_PHASE_STOPPING"
	EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseStopped     EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhase = "SERVICE_PHASE_STOPPED"
	EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseFailed      EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhase = "SERVICE_PHASE_FAILED"
	EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseDeleted     EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhase = "SERVICE_PHASE_DELETED"
)

func (r EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseUnspecified, EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseStarting, EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseRunning, EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseStopping, EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseStopped, EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseFailed, EnvironmentAutomationServiceNewResponseServiceSpecDesiredPhaseServicePhaseDeleted:
		return true
	}
	return false
}

// runs_on specifies the environment the service should run on.
type EnvironmentAutomationServiceNewResponseServiceSpecRunsOn struct {
	// This field can have the runtime type of
	// [EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDockerDocker].
	Docker interface{}                                                  `json:"docker"`
	JSON   environmentAutomationServiceNewResponseServiceSpecRunsOnJSON `json:"-"`
	union  EnvironmentAutomationServiceNewResponseServiceSpecRunsOnUnion
}

// environmentAutomationServiceNewResponseServiceSpecRunsOnJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationServiceNewResponseServiceSpecRunsOn]
type environmentAutomationServiceNewResponseServiceSpecRunsOnJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentAutomationServiceNewResponseServiceSpecRunsOnJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationServiceNewResponseServiceSpecRunsOn) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationServiceNewResponseServiceSpecRunsOn{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationServiceNewResponseServiceSpecRunsOnUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDocker],
// [EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDocker].
func (r EnvironmentAutomationServiceNewResponseServiceSpecRunsOn) AsUnion() EnvironmentAutomationServiceNewResponseServiceSpecRunsOnUnion {
	return r.union
}

// runs_on specifies the environment the service should run on.
//
// Union satisfied by
// [EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDocker] or
// [EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDocker].
type EnvironmentAutomationServiceNewResponseServiceSpecRunsOnUnion interface {
	implementsEnvironmentAutomationServiceNewResponseServiceSpecRunsOn()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationServiceNewResponseServiceSpecRunsOnUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDocker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDocker{}),
		},
	)
}

type EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDocker struct {
	Docker EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDockerDocker `json:"docker,required"`
	JSON   environmentAutomationServiceNewResponseServiceSpecRunsOnDockerJSON   `json:"-"`
}

// environmentAutomationServiceNewResponseServiceSpecRunsOnDockerJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDocker]
type environmentAutomationServiceNewResponseServiceSpecRunsOnDockerJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceNewResponseServiceSpecRunsOnDockerJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDocker) implementsEnvironmentAutomationServiceNewResponseServiceSpecRunsOn() {
}

type EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDockerDocker struct {
	Environment []string                                                                 `json:"environment"`
	Image       string                                                                   `json:"image"`
	JSON        environmentAutomationServiceNewResponseServiceSpecRunsOnDockerDockerJSON `json:"-"`
}

// environmentAutomationServiceNewResponseServiceSpecRunsOnDockerDockerJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDockerDocker]
type environmentAutomationServiceNewResponseServiceSpecRunsOnDockerDockerJSON struct {
	Environment apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceNewResponseServiceSpecRunsOnDockerDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceNewResponseServiceSpecRunsOnDockerDockerJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceNewResponseServiceStatus struct {
	// failure_message summarises why the service failed to operate. If this is
	// non-empty
	//
	// the service has failed to operate and will likely transition to a failed state.
	FailureMessage string `json:"failureMessage"`
	// log_url contains the URL at which the service logs can be accessed.
	LogURL string `json:"logUrl"`
	// phase is the current phase of the service.
	Phase EnvironmentAutomationServiceNewResponseServiceStatusPhase `json:"phase"`
	// session is the current session of the service.
	Session string `json:"session"`
	// version of the status update. Service instances themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string                                                   `json:"statusVersion" format:"int64"`
	JSON          environmentAutomationServiceNewResponseServiceStatusJSON `json:"-"`
}

// environmentAutomationServiceNewResponseServiceStatusJSON contains the JSON
// metadata for the struct [EnvironmentAutomationServiceNewResponseServiceStatus]
type environmentAutomationServiceNewResponseServiceStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	StatusVersion  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceNewResponseServiceStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceNewResponseServiceStatusJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the service.
type EnvironmentAutomationServiceNewResponseServiceStatusPhase string

const (
	EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseUnspecified EnvironmentAutomationServiceNewResponseServiceStatusPhase = "SERVICE_PHASE_UNSPECIFIED"
	EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseStarting    EnvironmentAutomationServiceNewResponseServiceStatusPhase = "SERVICE_PHASE_STARTING"
	EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseRunning     EnvironmentAutomationServiceNewResponseServiceStatusPhase = "SERVICE_PHASE_RUNNING"
	EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseStopping    EnvironmentAutomationServiceNewResponseServiceStatusPhase = "SERVICE_PHASE_STOPPING"
	EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseStopped     EnvironmentAutomationServiceNewResponseServiceStatusPhase = "SERVICE_PHASE_STOPPED"
	EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseFailed      EnvironmentAutomationServiceNewResponseServiceStatusPhase = "SERVICE_PHASE_FAILED"
	EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseDeleted     EnvironmentAutomationServiceNewResponseServiceStatusPhase = "SERVICE_PHASE_DELETED"
)

func (r EnvironmentAutomationServiceNewResponseServiceStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseUnspecified, EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseStarting, EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseRunning, EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseStopping, EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseStopped, EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseFailed, EnvironmentAutomationServiceNewResponseServiceStatusPhaseServicePhaseDeleted:
		return true
	}
	return false
}

type EnvironmentAutomationServiceGetResponse struct {
	Service EnvironmentAutomationServiceGetResponseService `json:"service"`
	JSON    environmentAutomationServiceGetResponseJSON    `json:"-"`
}

// environmentAutomationServiceGetResponseJSON contains the JSON metadata for the
// struct [EnvironmentAutomationServiceGetResponse]
type environmentAutomationServiceGetResponseJSON struct {
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceGetResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceGetResponseService struct {
	ID            string                                                 `json:"id" format:"uuid"`
	EnvironmentID string                                                 `json:"environmentId" format:"uuid"`
	Metadata      EnvironmentAutomationServiceGetResponseServiceMetadata `json:"metadata"`
	Spec          EnvironmentAutomationServiceGetResponseServiceSpec     `json:"spec"`
	Status        EnvironmentAutomationServiceGetResponseServiceStatus   `json:"status"`
	JSON          environmentAutomationServiceGetResponseServiceJSON     `json:"-"`
}

// environmentAutomationServiceGetResponseServiceJSON contains the JSON metadata
// for the struct [EnvironmentAutomationServiceGetResponseService]
type environmentAutomationServiceGetResponseServiceJSON struct {
	ID            apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceGetResponseService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceGetResponseServiceJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceGetResponseServiceMetadata struct {
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
	// creator describes the principal who created the service.
	Creator EnvironmentAutomationServiceGetResponseServiceMetadataCreator `json:"creator"`
	// description is a user-facing description for the service. It can be used to
	// provide context and documentation for the service.
	Description string `json:"description"`
	// name is a user-facing name for the service. Unlike the reference, this field is
	// not unique, and not referenced by the system.
	//
	// This is a short descriptive name for the service.
	Name string `json:"name"`
	// reference is a user-facing identifier for the service which must be unique on
	// the environment.
	//
	// It is used to express dependencies between services, and to identify the service
	// in user interactions (e.g. the CLI).
	Reference string `json:"reference"`
	// triggered_by is a list of trigger that start the service.
	TriggeredBy []EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredBy `json:"triggeredBy"`
	JSON        environmentAutomationServiceGetResponseServiceMetadataJSON          `json:"-"`
}

// environmentAutomationServiceGetResponseServiceMetadataJSON contains the JSON
// metadata for the struct [EnvironmentAutomationServiceGetResponseServiceMetadata]
type environmentAutomationServiceGetResponseServiceMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceGetResponseServiceMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceGetResponseServiceMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created the service.
type EnvironmentAutomationServiceGetResponseServiceMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationServiceGetResponseServiceMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationServiceGetResponseServiceMetadataCreatorJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationServiceGetResponseServiceMetadataCreator]
type environmentAutomationServiceGetResponseServiceMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceGetResponseServiceMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceGetResponseServiceMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipal string

const (
	EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationServiceGetResponseServiceMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredBy struct {
	Manual                bool                                                                  `json:"manual"`
	PostDevcontainerStart bool                                                                  `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                                  `json:"postEnvironmentStart"`
	JSON                  environmentAutomationServiceGetResponseServiceMetadataTriggeredByJSON `json:"-"`
	union                 EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByUnion
}

// environmentAutomationServiceGetResponseServiceMetadataTriggeredByJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredBy]
type environmentAutomationServiceGetResponseServiceMetadataTriggeredByJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r environmentAutomationServiceGetResponseServiceMetadataTriggeredByJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredBy) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredBy{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject],
// [EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject],
// [EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject],
// [EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject].
func (r EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredBy) AsUnion() EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByUnion {
	return r.union
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// Union satisfied by
// [EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject],
// [EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject],
// [EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject] or
// [EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject].
type EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByUnion interface {
	implementsEnvironmentAutomationServiceGetResponseServiceMetadataTriggeredBy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject{}),
		},
	)
}

type EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject struct {
	Manual                bool                                                                        `json:"manual,required"`
	PostDevcontainerStart bool                                                                        `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                                        `json:"postEnvironmentStart"`
	JSON                  environmentAutomationServiceGetResponseServiceMetadataTriggeredByObjectJSON `json:"-"`
}

// environmentAutomationServiceGetResponseServiceMetadataTriggeredByObjectJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject]
type environmentAutomationServiceGetResponseServiceMetadataTriggeredByObjectJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceGetResponseServiceMetadataTriggeredByObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationServiceGetResponseServiceMetadataTriggeredByObject) implementsEnvironmentAutomationServiceGetResponseServiceMetadataTriggeredBy() {
}

type EnvironmentAutomationServiceGetResponseServiceSpec struct {
	// commands contains the commands to start, stop and check the readiness of the
	// service
	Commands EnvironmentAutomationServiceGetResponseServiceSpecCommands `json:"commands"`
	// desired_phase is the phase the service should be in. Used to start or stop the
	// service.
	DesiredPhase EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhase `json:"desiredPhase"`
	// runs_on specifies the environment the service should run on.
	RunsOn EnvironmentAutomationServiceGetResponseServiceSpecRunsOn `json:"runsOn"`
	// session should be changed to trigger a restart of the service. If a service
	// exits it will
	//
	// not be restarted until the session is changed.
	Session string `json:"session"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion string                                                 `json:"specVersion" format:"int64"`
	JSON        environmentAutomationServiceGetResponseServiceSpecJSON `json:"-"`
}

// environmentAutomationServiceGetResponseServiceSpecJSON contains the JSON
// metadata for the struct [EnvironmentAutomationServiceGetResponseServiceSpec]
type environmentAutomationServiceGetResponseServiceSpecJSON struct {
	Commands     apijson.Field
	DesiredPhase apijson.Field
	RunsOn       apijson.Field
	Session      apijson.Field
	SpecVersion  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceGetResponseServiceSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceGetResponseServiceSpecJSON) RawJSON() string {
	return r.raw
}

// commands contains the commands to start, stop and check the readiness of the
// service
type EnvironmentAutomationServiceGetResponseServiceSpecCommands struct {
	// ready is an optional command that is run repeatedly until it exits with a zero
	// exit code.
	//
	// If set, the service will first go into a Starting phase, and then into a Running
	// phase once the ready command exits with a zero exit code.
	Ready string `json:"ready"`
	// start is the command to start and run the service. If start exits, the service
	// will transition to the following phase:
	//
	//   - Stopped: if the exit code is 0
	//   - Failed: if the exit code is not 0 If the stop command is not set, the start
	//     command will receive a SIGTERM signal when the service is requested to stop.
	//     If it does not exit within 2 minutes, it will receive a SIGKILL signal.
	Start string `json:"start"`
	// stop is an optional command that runs when the service is requested to stop.
	//
	// If set, instead of sending a SIGTERM signal to the start command, the stop
	// command will be run. Once the stop command exits, the start command will receive
	// a SIGKILL signal. If the stop command exits with a non-zero exit code, the
	// service will transition to the Failed phase. If the stop command does not exit
	// within 2 minutes, a SIGKILL signal will be sent to both the start and stop
	// commands.
	Stop string                                                         `json:"stop"`
	JSON environmentAutomationServiceGetResponseServiceSpecCommandsJSON `json:"-"`
}

// environmentAutomationServiceGetResponseServiceSpecCommandsJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationServiceGetResponseServiceSpecCommands]
type environmentAutomationServiceGetResponseServiceSpecCommandsJSON struct {
	Ready       apijson.Field
	Start       apijson.Field
	Stop        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceGetResponseServiceSpecCommands) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceGetResponseServiceSpecCommandsJSON) RawJSON() string {
	return r.raw
}

// desired_phase is the phase the service should be in. Used to start or stop the
// service.
type EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhase string

const (
	EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseUnspecified EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhase = "SERVICE_PHASE_UNSPECIFIED"
	EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseStarting    EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhase = "SERVICE_PHASE_STARTING"
	EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseRunning     EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhase = "SERVICE_PHASE_RUNNING"
	EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseStopping    EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhase = "SERVICE_PHASE_STOPPING"
	EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseStopped     EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhase = "SERVICE_PHASE_STOPPED"
	EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseFailed      EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhase = "SERVICE_PHASE_FAILED"
	EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseDeleted     EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhase = "SERVICE_PHASE_DELETED"
)

func (r EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseUnspecified, EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseStarting, EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseRunning, EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseStopping, EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseStopped, EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseFailed, EnvironmentAutomationServiceGetResponseServiceSpecDesiredPhaseServicePhaseDeleted:
		return true
	}
	return false
}

// runs_on specifies the environment the service should run on.
type EnvironmentAutomationServiceGetResponseServiceSpecRunsOn struct {
	// This field can have the runtime type of
	// [EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDockerDocker].
	Docker interface{}                                                  `json:"docker"`
	JSON   environmentAutomationServiceGetResponseServiceSpecRunsOnJSON `json:"-"`
	union  EnvironmentAutomationServiceGetResponseServiceSpecRunsOnUnion
}

// environmentAutomationServiceGetResponseServiceSpecRunsOnJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationServiceGetResponseServiceSpecRunsOn]
type environmentAutomationServiceGetResponseServiceSpecRunsOnJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentAutomationServiceGetResponseServiceSpecRunsOnJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationServiceGetResponseServiceSpecRunsOn) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationServiceGetResponseServiceSpecRunsOn{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationServiceGetResponseServiceSpecRunsOnUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDocker],
// [EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDocker].
func (r EnvironmentAutomationServiceGetResponseServiceSpecRunsOn) AsUnion() EnvironmentAutomationServiceGetResponseServiceSpecRunsOnUnion {
	return r.union
}

// runs_on specifies the environment the service should run on.
//
// Union satisfied by
// [EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDocker] or
// [EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDocker].
type EnvironmentAutomationServiceGetResponseServiceSpecRunsOnUnion interface {
	implementsEnvironmentAutomationServiceGetResponseServiceSpecRunsOn()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationServiceGetResponseServiceSpecRunsOnUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDocker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDocker{}),
		},
	)
}

type EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDocker struct {
	Docker EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDockerDocker `json:"docker,required"`
	JSON   environmentAutomationServiceGetResponseServiceSpecRunsOnDockerJSON   `json:"-"`
}

// environmentAutomationServiceGetResponseServiceSpecRunsOnDockerJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDocker]
type environmentAutomationServiceGetResponseServiceSpecRunsOnDockerJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceGetResponseServiceSpecRunsOnDockerJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDocker) implementsEnvironmentAutomationServiceGetResponseServiceSpecRunsOn() {
}

type EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDockerDocker struct {
	Environment []string                                                                 `json:"environment"`
	Image       string                                                                   `json:"image"`
	JSON        environmentAutomationServiceGetResponseServiceSpecRunsOnDockerDockerJSON `json:"-"`
}

// environmentAutomationServiceGetResponseServiceSpecRunsOnDockerDockerJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDockerDocker]
type environmentAutomationServiceGetResponseServiceSpecRunsOnDockerDockerJSON struct {
	Environment apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceGetResponseServiceSpecRunsOnDockerDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceGetResponseServiceSpecRunsOnDockerDockerJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceGetResponseServiceStatus struct {
	// failure_message summarises why the service failed to operate. If this is
	// non-empty
	//
	// the service has failed to operate and will likely transition to a failed state.
	FailureMessage string `json:"failureMessage"`
	// log_url contains the URL at which the service logs can be accessed.
	LogURL string `json:"logUrl"`
	// phase is the current phase of the service.
	Phase EnvironmentAutomationServiceGetResponseServiceStatusPhase `json:"phase"`
	// session is the current session of the service.
	Session string `json:"session"`
	// version of the status update. Service instances themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string                                                   `json:"statusVersion" format:"int64"`
	JSON          environmentAutomationServiceGetResponseServiceStatusJSON `json:"-"`
}

// environmentAutomationServiceGetResponseServiceStatusJSON contains the JSON
// metadata for the struct [EnvironmentAutomationServiceGetResponseServiceStatus]
type environmentAutomationServiceGetResponseServiceStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	StatusVersion  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceGetResponseServiceStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceGetResponseServiceStatusJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the service.
type EnvironmentAutomationServiceGetResponseServiceStatusPhase string

const (
	EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseUnspecified EnvironmentAutomationServiceGetResponseServiceStatusPhase = "SERVICE_PHASE_UNSPECIFIED"
	EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseStarting    EnvironmentAutomationServiceGetResponseServiceStatusPhase = "SERVICE_PHASE_STARTING"
	EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseRunning     EnvironmentAutomationServiceGetResponseServiceStatusPhase = "SERVICE_PHASE_RUNNING"
	EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseStopping    EnvironmentAutomationServiceGetResponseServiceStatusPhase = "SERVICE_PHASE_STOPPING"
	EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseStopped     EnvironmentAutomationServiceGetResponseServiceStatusPhase = "SERVICE_PHASE_STOPPED"
	EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseFailed      EnvironmentAutomationServiceGetResponseServiceStatusPhase = "SERVICE_PHASE_FAILED"
	EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseDeleted     EnvironmentAutomationServiceGetResponseServiceStatusPhase = "SERVICE_PHASE_DELETED"
)

func (r EnvironmentAutomationServiceGetResponseServiceStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseUnspecified, EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseStarting, EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseRunning, EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseStopping, EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseStopped, EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseFailed, EnvironmentAutomationServiceGetResponseServiceStatusPhaseServicePhaseDeleted:
		return true
	}
	return false
}

type EnvironmentAutomationServiceUpdateResponse = interface{}

type EnvironmentAutomationServiceListResponse struct {
	Pagination EnvironmentAutomationServiceListResponsePagination `json:"pagination"`
	Services   []EnvironmentAutomationServiceListResponseService  `json:"services"`
	JSON       environmentAutomationServiceListResponseJSON       `json:"-"`
}

// environmentAutomationServiceListResponseJSON contains the JSON metadata for the
// struct [EnvironmentAutomationServiceListResponse]
type environmentAutomationServiceListResponseJSON struct {
	Pagination  apijson.Field
	Services    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                                 `json:"nextToken"`
	JSON      environmentAutomationServiceListResponsePaginationJSON `json:"-"`
}

// environmentAutomationServiceListResponsePaginationJSON contains the JSON
// metadata for the struct [EnvironmentAutomationServiceListResponsePagination]
type environmentAutomationServiceListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceListResponseService struct {
	ID            string                                                   `json:"id" format:"uuid"`
	EnvironmentID string                                                   `json:"environmentId" format:"uuid"`
	Metadata      EnvironmentAutomationServiceListResponseServicesMetadata `json:"metadata"`
	Spec          EnvironmentAutomationServiceListResponseServicesSpec     `json:"spec"`
	Status        EnvironmentAutomationServiceListResponseServicesStatus   `json:"status"`
	JSON          environmentAutomationServiceListResponseServiceJSON      `json:"-"`
}

// environmentAutomationServiceListResponseServiceJSON contains the JSON metadata
// for the struct [EnvironmentAutomationServiceListResponseService]
type environmentAutomationServiceListResponseServiceJSON struct {
	ID            apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServiceJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceListResponseServicesMetadata struct {
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
	// creator describes the principal who created the service.
	Creator EnvironmentAutomationServiceListResponseServicesMetadataCreator `json:"creator"`
	// description is a user-facing description for the service. It can be used to
	// provide context and documentation for the service.
	Description string `json:"description"`
	// name is a user-facing name for the service. Unlike the reference, this field is
	// not unique, and not referenced by the system.
	//
	// This is a short descriptive name for the service.
	Name string `json:"name"`
	// reference is a user-facing identifier for the service which must be unique on
	// the environment.
	//
	// It is used to express dependencies between services, and to identify the service
	// in user interactions (e.g. the CLI).
	Reference string `json:"reference"`
	// triggered_by is a list of trigger that start the service.
	TriggeredBy []EnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy `json:"triggeredBy"`
	JSON        environmentAutomationServiceListResponseServicesMetadataJSON          `json:"-"`
}

// environmentAutomationServiceListResponseServicesMetadataJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesMetadata]
type environmentAutomationServiceListResponseServicesMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created the service.
type EnvironmentAutomationServiceListResponseServicesMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentAutomationServiceListResponseServicesMetadataCreatorJSON      `json:"-"`
}

// environmentAutomationServiceListResponseServicesMetadataCreatorJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesMetadataCreator]
type environmentAutomationServiceListResponseServicesMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipal string

const (
	EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationServiceListResponseServicesMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type EnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy struct {
	Manual                bool                                                                    `json:"manual"`
	PostDevcontainerStart bool                                                                    `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                                    `json:"postEnvironmentStart"`
	JSON                  environmentAutomationServiceListResponseServicesMetadataTriggeredByJSON `json:"-"`
	union                 EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByUnion
}

// environmentAutomationServiceListResponseServicesMetadataTriggeredByJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy]
type environmentAutomationServiceListResponseServicesMetadataTriggeredByJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r environmentAutomationServiceListResponseServicesMetadataTriggeredByJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject],
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject],
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject],
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject].
func (r EnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy) AsUnion() EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByUnion {
	return r.union
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// Union satisfied by
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject],
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject],
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject] or
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject].
type EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByUnion interface {
	implementsEnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject{}),
		},
	)
}

type EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject struct {
	Manual                bool                                                                          `json:"manual,required"`
	PostDevcontainerStart bool                                                                          `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                                                                          `json:"postEnvironmentStart"`
	JSON                  environmentAutomationServiceListResponseServicesMetadataTriggeredByObjectJSON `json:"-"`
}

// environmentAutomationServiceListResponseServicesMetadataTriggeredByObjectJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject]
type environmentAutomationServiceListResponseServicesMetadataTriggeredByObjectJSON struct {
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesMetadataTriggeredByObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByObject) implementsEnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy() {
}

type EnvironmentAutomationServiceListResponseServicesSpec struct {
	// commands contains the commands to start, stop and check the readiness of the
	// service
	Commands EnvironmentAutomationServiceListResponseServicesSpecCommands `json:"commands"`
	// desired_phase is the phase the service should be in. Used to start or stop the
	// service.
	DesiredPhase EnvironmentAutomationServiceListResponseServicesSpecDesiredPhase `json:"desiredPhase"`
	// runs_on specifies the environment the service should run on.
	RunsOn EnvironmentAutomationServiceListResponseServicesSpecRunsOn `json:"runsOn"`
	// session should be changed to trigger a restart of the service. If a service
	// exits it will
	//
	// not be restarted until the session is changed.
	Session string `json:"session"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion string                                                   `json:"specVersion" format:"int64"`
	JSON        environmentAutomationServiceListResponseServicesSpecJSON `json:"-"`
}

// environmentAutomationServiceListResponseServicesSpecJSON contains the JSON
// metadata for the struct [EnvironmentAutomationServiceListResponseServicesSpec]
type environmentAutomationServiceListResponseServicesSpecJSON struct {
	Commands     apijson.Field
	DesiredPhase apijson.Field
	RunsOn       apijson.Field
	Session      apijson.Field
	SpecVersion  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesSpecJSON) RawJSON() string {
	return r.raw
}

// commands contains the commands to start, stop and check the readiness of the
// service
type EnvironmentAutomationServiceListResponseServicesSpecCommands struct {
	// ready is an optional command that is run repeatedly until it exits with a zero
	// exit code.
	//
	// If set, the service will first go into a Starting phase, and then into a Running
	// phase once the ready command exits with a zero exit code.
	Ready string `json:"ready"`
	// start is the command to start and run the service. If start exits, the service
	// will transition to the following phase:
	//
	//   - Stopped: if the exit code is 0
	//   - Failed: if the exit code is not 0 If the stop command is not set, the start
	//     command will receive a SIGTERM signal when the service is requested to stop.
	//     If it does not exit within 2 minutes, it will receive a SIGKILL signal.
	Start string `json:"start"`
	// stop is an optional command that runs when the service is requested to stop.
	//
	// If set, instead of sending a SIGTERM signal to the start command, the stop
	// command will be run. Once the stop command exits, the start command will receive
	// a SIGKILL signal. If the stop command exits with a non-zero exit code, the
	// service will transition to the Failed phase. If the stop command does not exit
	// within 2 minutes, a SIGKILL signal will be sent to both the start and stop
	// commands.
	Stop string                                                           `json:"stop"`
	JSON environmentAutomationServiceListResponseServicesSpecCommandsJSON `json:"-"`
}

// environmentAutomationServiceListResponseServicesSpecCommandsJSON contains the
// JSON metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesSpecCommands]
type environmentAutomationServiceListResponseServicesSpecCommandsJSON struct {
	Ready       apijson.Field
	Start       apijson.Field
	Stop        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesSpecCommands) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesSpecCommandsJSON) RawJSON() string {
	return r.raw
}

// desired_phase is the phase the service should be in. Used to start or stop the
// service.
type EnvironmentAutomationServiceListResponseServicesSpecDesiredPhase string

const (
	EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseUnspecified EnvironmentAutomationServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_UNSPECIFIED"
	EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseStarting    EnvironmentAutomationServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_STARTING"
	EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseRunning     EnvironmentAutomationServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_RUNNING"
	EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseStopping    EnvironmentAutomationServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_STOPPING"
	EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseStopped     EnvironmentAutomationServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_STOPPED"
	EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseFailed      EnvironmentAutomationServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_FAILED"
	EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseDeleted     EnvironmentAutomationServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_DELETED"
)

func (r EnvironmentAutomationServiceListResponseServicesSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseUnspecified, EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseStarting, EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseRunning, EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseStopping, EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseStopped, EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseFailed, EnvironmentAutomationServiceListResponseServicesSpecDesiredPhaseServicePhaseDeleted:
		return true
	}
	return false
}

// runs_on specifies the environment the service should run on.
type EnvironmentAutomationServiceListResponseServicesSpecRunsOn struct {
	// This field can have the runtime type of
	// [EnvironmentAutomationServiceListResponseServicesSpecRunsOnDockerDocker].
	Docker interface{}                                                    `json:"docker"`
	JSON   environmentAutomationServiceListResponseServicesSpecRunsOnJSON `json:"-"`
	union  EnvironmentAutomationServiceListResponseServicesSpecRunsOnUnion
}

// environmentAutomationServiceListResponseServicesSpecRunsOnJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesSpecRunsOn]
type environmentAutomationServiceListResponseServicesSpecRunsOnJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentAutomationServiceListResponseServicesSpecRunsOnJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentAutomationServiceListResponseServicesSpecRunsOn) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentAutomationServiceListResponseServicesSpecRunsOn{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentAutomationServiceListResponseServicesSpecRunsOnUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker],
// [EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker].
func (r EnvironmentAutomationServiceListResponseServicesSpecRunsOn) AsUnion() EnvironmentAutomationServiceListResponseServicesSpecRunsOnUnion {
	return r.union
}

// runs_on specifies the environment the service should run on.
//
// Union satisfied by
// [EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker] or
// [EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker].
type EnvironmentAutomationServiceListResponseServicesSpecRunsOnUnion interface {
	implementsEnvironmentAutomationServiceListResponseServicesSpecRunsOn()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationServiceListResponseServicesSpecRunsOnUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker{}),
		},
	)
}

type EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker struct {
	Docker EnvironmentAutomationServiceListResponseServicesSpecRunsOnDockerDocker `json:"docker,required"`
	JSON   environmentAutomationServiceListResponseServicesSpecRunsOnDockerJSON   `json:"-"`
}

// environmentAutomationServiceListResponseServicesSpecRunsOnDockerJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker]
type environmentAutomationServiceListResponseServicesSpecRunsOnDockerJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesSpecRunsOnDockerJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker) implementsEnvironmentAutomationServiceListResponseServicesSpecRunsOn() {
}

type EnvironmentAutomationServiceListResponseServicesSpecRunsOnDockerDocker struct {
	Environment []string                                                                   `json:"environment"`
	Image       string                                                                     `json:"image"`
	JSON        environmentAutomationServiceListResponseServicesSpecRunsOnDockerDockerJSON `json:"-"`
}

// environmentAutomationServiceListResponseServicesSpecRunsOnDockerDockerJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesSpecRunsOnDockerDocker]
type environmentAutomationServiceListResponseServicesSpecRunsOnDockerDockerJSON struct {
	Environment apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesSpecRunsOnDockerDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesSpecRunsOnDockerDockerJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceListResponseServicesStatus struct {
	// failure_message summarises why the service failed to operate. If this is
	// non-empty
	//
	// the service has failed to operate and will likely transition to a failed state.
	FailureMessage string `json:"failureMessage"`
	// log_url contains the URL at which the service logs can be accessed.
	LogURL string `json:"logUrl"`
	// phase is the current phase of the service.
	Phase EnvironmentAutomationServiceListResponseServicesStatusPhase `json:"phase"`
	// session is the current session of the service.
	Session string `json:"session"`
	// version of the status update. Service instances themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string                                                     `json:"statusVersion" format:"int64"`
	JSON          environmentAutomationServiceListResponseServicesStatusJSON `json:"-"`
}

// environmentAutomationServiceListResponseServicesStatusJSON contains the JSON
// metadata for the struct [EnvironmentAutomationServiceListResponseServicesStatus]
type environmentAutomationServiceListResponseServicesStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	StatusVersion  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesStatusJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the service.
type EnvironmentAutomationServiceListResponseServicesStatusPhase string

const (
	EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseUnspecified EnvironmentAutomationServiceListResponseServicesStatusPhase = "SERVICE_PHASE_UNSPECIFIED"
	EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseStarting    EnvironmentAutomationServiceListResponseServicesStatusPhase = "SERVICE_PHASE_STARTING"
	EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseRunning     EnvironmentAutomationServiceListResponseServicesStatusPhase = "SERVICE_PHASE_RUNNING"
	EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseStopping    EnvironmentAutomationServiceListResponseServicesStatusPhase = "SERVICE_PHASE_STOPPING"
	EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseStopped     EnvironmentAutomationServiceListResponseServicesStatusPhase = "SERVICE_PHASE_STOPPED"
	EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseFailed      EnvironmentAutomationServiceListResponseServicesStatusPhase = "SERVICE_PHASE_FAILED"
	EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseDeleted     EnvironmentAutomationServiceListResponseServicesStatusPhase = "SERVICE_PHASE_DELETED"
)

func (r EnvironmentAutomationServiceListResponseServicesStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseUnspecified, EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseStarting, EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseRunning, EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseStopping, EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseStopped, EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseFailed, EnvironmentAutomationServiceListResponseServicesStatusPhaseServicePhaseDeleted:
		return true
	}
	return false
}

type EnvironmentAutomationServiceDeleteResponse = interface{}

type EnvironmentAutomationServiceStartResponse = interface{}

type EnvironmentAutomationServiceStopResponse = interface{}

type EnvironmentAutomationServiceNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationServiceNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	EnvironmentID          param.Field[string]                                                      `json:"environmentId" format:"uuid"`
	Metadata               param.Field[EnvironmentAutomationServiceNewParamsMetadata]               `json:"metadata"`
	Spec                   param.Field[EnvironmentAutomationServiceNewParamsSpec]                   `json:"spec"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationServiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationServiceNewParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationServiceNewParamsConnectProtocolVersion1 EnvironmentAutomationServiceNewParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationServiceNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentAutomationServiceNewParamsMetadata struct {
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
	// creator describes the principal who created the service.
	Creator param.Field[EnvironmentAutomationServiceNewParamsMetadataCreator] `json:"creator"`
	// description is a user-facing description for the service. It can be used to
	// provide context and documentation for the service.
	Description param.Field[string] `json:"description"`
	// name is a user-facing name for the service. Unlike the reference, this field is
	// not unique, and not referenced by the system.
	//
	// This is a short descriptive name for the service.
	Name param.Field[string] `json:"name"`
	// reference is a user-facing identifier for the service which must be unique on
	// the environment.
	//
	// It is used to express dependencies between services, and to identify the service
	// in user interactions (e.g. the CLI).
	Reference param.Field[string] `json:"reference"`
	// triggered_by is a list of trigger that start the service.
	TriggeredBy param.Field[[]EnvironmentAutomationServiceNewParamsMetadataTriggeredByUnion] `json:"triggeredBy"`
}

func (r EnvironmentAutomationServiceNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// creator describes the principal who created the service.
type EnvironmentAutomationServiceNewParamsMetadataCreator struct {
	// id is the UUID of the subject
	ID param.Field[string] `json:"id"`
	// Principal is the principal of the subject
	Principal param.Field[EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipal] `json:"principal"`
}

func (r EnvironmentAutomationServiceNewParamsMetadataCreator) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Principal is the principal of the subject
type EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipal string

const (
	EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalAccount        EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalUser           EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalRunner         EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalAccount, EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalUser, EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalRunner, EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentAutomationServiceNewParamsMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type EnvironmentAutomationServiceNewParamsMetadataTriggeredBy struct {
	Manual                param.Field[bool] `json:"manual"`
	PostDevcontainerStart param.Field[bool] `json:"postDevcontainerStart"`
	PostEnvironmentStart  param.Field[bool] `json:"postEnvironmentStart"`
}

func (r EnvironmentAutomationServiceNewParamsMetadataTriggeredBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceNewParamsMetadataTriggeredBy) implementsEnvironmentAutomationServiceNewParamsMetadataTriggeredByUnion() {
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// Satisfied by [EnvironmentAutomationServiceNewParamsMetadataTriggeredByObject],
// [EnvironmentAutomationServiceNewParamsMetadataTriggeredByObject],
// [EnvironmentAutomationServiceNewParamsMetadataTriggeredByObject],
// [EnvironmentAutomationServiceNewParamsMetadataTriggeredByObject],
// [EnvironmentAutomationServiceNewParamsMetadataTriggeredBy].
type EnvironmentAutomationServiceNewParamsMetadataTriggeredByUnion interface {
	implementsEnvironmentAutomationServiceNewParamsMetadataTriggeredByUnion()
}

type EnvironmentAutomationServiceNewParamsMetadataTriggeredByObject struct {
	Manual                param.Field[bool] `json:"manual,required"`
	PostDevcontainerStart param.Field[bool] `json:"postDevcontainerStart"`
	PostEnvironmentStart  param.Field[bool] `json:"postEnvironmentStart"`
}

func (r EnvironmentAutomationServiceNewParamsMetadataTriggeredByObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceNewParamsMetadataTriggeredByObject) implementsEnvironmentAutomationServiceNewParamsMetadataTriggeredByUnion() {
}

type EnvironmentAutomationServiceNewParamsSpec struct {
	// commands contains the commands to start, stop and check the readiness of the
	// service
	Commands param.Field[EnvironmentAutomationServiceNewParamsSpecCommands] `json:"commands"`
	// desired_phase is the phase the service should be in. Used to start or stop the
	// service.
	DesiredPhase param.Field[EnvironmentAutomationServiceNewParamsSpecDesiredPhase] `json:"desiredPhase"`
	// runs_on specifies the environment the service should run on.
	RunsOn param.Field[EnvironmentAutomationServiceNewParamsSpecRunsOnUnion] `json:"runsOn"`
	// session should be changed to trigger a restart of the service. If a service
	// exits it will
	//
	// not be restarted until the session is changed.
	Session param.Field[string] `json:"session"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion param.Field[string] `json:"specVersion" format:"int64"`
}

func (r EnvironmentAutomationServiceNewParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// commands contains the commands to start, stop and check the readiness of the
// service
type EnvironmentAutomationServiceNewParamsSpecCommands struct {
	// ready is an optional command that is run repeatedly until it exits with a zero
	// exit code.
	//
	// If set, the service will first go into a Starting phase, and then into a Running
	// phase once the ready command exits with a zero exit code.
	Ready param.Field[string] `json:"ready"`
	// start is the command to start and run the service. If start exits, the service
	// will transition to the following phase:
	//
	//   - Stopped: if the exit code is 0
	//   - Failed: if the exit code is not 0 If the stop command is not set, the start
	//     command will receive a SIGTERM signal when the service is requested to stop.
	//     If it does not exit within 2 minutes, it will receive a SIGKILL signal.
	Start param.Field[string] `json:"start"`
	// stop is an optional command that runs when the service is requested to stop.
	//
	// If set, instead of sending a SIGTERM signal to the start command, the stop
	// command will be run. Once the stop command exits, the start command will receive
	// a SIGKILL signal. If the stop command exits with a non-zero exit code, the
	// service will transition to the Failed phase. If the stop command does not exit
	// within 2 minutes, a SIGKILL signal will be sent to both the start and stop
	// commands.
	Stop param.Field[string] `json:"stop"`
}

func (r EnvironmentAutomationServiceNewParamsSpecCommands) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// desired_phase is the phase the service should be in. Used to start or stop the
// service.
type EnvironmentAutomationServiceNewParamsSpecDesiredPhase string

const (
	EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseUnspecified EnvironmentAutomationServiceNewParamsSpecDesiredPhase = "SERVICE_PHASE_UNSPECIFIED"
	EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseStarting    EnvironmentAutomationServiceNewParamsSpecDesiredPhase = "SERVICE_PHASE_STARTING"
	EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseRunning     EnvironmentAutomationServiceNewParamsSpecDesiredPhase = "SERVICE_PHASE_RUNNING"
	EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseStopping    EnvironmentAutomationServiceNewParamsSpecDesiredPhase = "SERVICE_PHASE_STOPPING"
	EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseStopped     EnvironmentAutomationServiceNewParamsSpecDesiredPhase = "SERVICE_PHASE_STOPPED"
	EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseFailed      EnvironmentAutomationServiceNewParamsSpecDesiredPhase = "SERVICE_PHASE_FAILED"
	EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseDeleted     EnvironmentAutomationServiceNewParamsSpecDesiredPhase = "SERVICE_PHASE_DELETED"
)

func (r EnvironmentAutomationServiceNewParamsSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseUnspecified, EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseStarting, EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseRunning, EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseStopping, EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseStopped, EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseFailed, EnvironmentAutomationServiceNewParamsSpecDesiredPhaseServicePhaseDeleted:
		return true
	}
	return false
}

// runs_on specifies the environment the service should run on.
type EnvironmentAutomationServiceNewParamsSpecRunsOn struct {
	Docker param.Field[interface{}] `json:"docker"`
}

func (r EnvironmentAutomationServiceNewParamsSpecRunsOn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceNewParamsSpecRunsOn) implementsEnvironmentAutomationServiceNewParamsSpecRunsOnUnion() {
}

// runs_on specifies the environment the service should run on.
//
// Satisfied by [EnvironmentAutomationServiceNewParamsSpecRunsOnDocker],
// [EnvironmentAutomationServiceNewParamsSpecRunsOnDocker],
// [EnvironmentAutomationServiceNewParamsSpecRunsOn].
type EnvironmentAutomationServiceNewParamsSpecRunsOnUnion interface {
	implementsEnvironmentAutomationServiceNewParamsSpecRunsOnUnion()
}

type EnvironmentAutomationServiceNewParamsSpecRunsOnDocker struct {
	Docker param.Field[EnvironmentAutomationServiceNewParamsSpecRunsOnDockerDocker] `json:"docker,required"`
}

func (r EnvironmentAutomationServiceNewParamsSpecRunsOnDocker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceNewParamsSpecRunsOnDocker) implementsEnvironmentAutomationServiceNewParamsSpecRunsOnUnion() {
}

type EnvironmentAutomationServiceNewParamsSpecRunsOnDockerDocker struct {
	Environment param.Field[[]string] `json:"environment"`
	Image       param.Field[string]   `json:"image"`
}

func (r EnvironmentAutomationServiceNewParamsSpecRunsOnDockerDocker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceGetParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[EnvironmentAutomationServiceGetParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationServiceGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[EnvironmentAutomationServiceGetParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[EnvironmentAutomationServiceGetParamsConnect] `query:"connect"`
	Message param.Field[string]                                       `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [EnvironmentAutomationServiceGetParams]'s query parameters
// as `url.Values`.
func (r EnvironmentAutomationServiceGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type EnvironmentAutomationServiceGetParamsEncoding string

const (
	EnvironmentAutomationServiceGetParamsEncodingProto EnvironmentAutomationServiceGetParamsEncoding = "proto"
	EnvironmentAutomationServiceGetParamsEncodingJson  EnvironmentAutomationServiceGetParamsEncoding = "json"
)

func (r EnvironmentAutomationServiceGetParamsEncoding) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceGetParamsEncodingProto, EnvironmentAutomationServiceGetParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EnvironmentAutomationServiceGetParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationServiceGetParamsConnectProtocolVersion1 EnvironmentAutomationServiceGetParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationServiceGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type EnvironmentAutomationServiceGetParamsCompression string

const (
	EnvironmentAutomationServiceGetParamsCompressionIdentity EnvironmentAutomationServiceGetParamsCompression = "identity"
	EnvironmentAutomationServiceGetParamsCompressionGzip     EnvironmentAutomationServiceGetParamsCompression = "gzip"
	EnvironmentAutomationServiceGetParamsCompressionBr       EnvironmentAutomationServiceGetParamsCompression = "br"
)

func (r EnvironmentAutomationServiceGetParamsCompression) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceGetParamsCompressionIdentity, EnvironmentAutomationServiceGetParamsCompressionGzip, EnvironmentAutomationServiceGetParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EnvironmentAutomationServiceGetParamsConnect string

const (
	EnvironmentAutomationServiceGetParamsConnectV1 EnvironmentAutomationServiceGetParamsConnect = "v1"
)

func (r EnvironmentAutomationServiceGetParamsConnect) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceGetParamsConnectV1:
		return true
	}
	return false
}

type EnvironmentAutomationServiceUpdateParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationServiceUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                         `json:"id" format:"uuid"`
	Metadata               param.Field[EnvironmentAutomationServiceUpdateParamsMetadata]               `json:"metadata"`
	// Changing the spec of a service is a complex operation. The spec of a service
	//
	// can only be updated if the service is in a stopped state. If the service is
	// running, it must be stopped first.
	Spec param.Field[EnvironmentAutomationServiceUpdateParamsSpec] `json:"spec"`
	// Service status updates are only expected from the executing environment. As a
	// client
	//
	// of this API you are not expected to provide this field. Updating this field
	// requires the `environmentservice:update_status` permission.
	Status param.Field[EnvironmentAutomationServiceUpdateParamsStatus] `json:"status"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationServiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationServiceUpdateParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationServiceUpdateParamsConnectProtocolVersion1 EnvironmentAutomationServiceUpdateParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationServiceUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentAutomationServiceUpdateParamsMetadata struct {
}

func (r EnvironmentAutomationServiceUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Changing the spec of a service is a complex operation. The spec of a service
//
// can only be updated if the service is in a stopped state. If the service is
// running, it must be stopped first.
type EnvironmentAutomationServiceUpdateParamsSpec struct {
}

func (r EnvironmentAutomationServiceUpdateParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service status updates are only expected from the executing environment. As a
// client
//
// of this API you are not expected to provide this field. Updating this field
// requires the `environmentservice:update_status` permission.
type EnvironmentAutomationServiceUpdateParamsStatus struct {
}

func (r EnvironmentAutomationServiceUpdateParamsStatus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[EnvironmentAutomationServiceListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationServiceListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[EnvironmentAutomationServiceListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[EnvironmentAutomationServiceListParamsConnect] `query:"connect"`
	Message param.Field[string]                                        `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [EnvironmentAutomationServiceListParams]'s query parameters
// as `url.Values`.
func (r EnvironmentAutomationServiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type EnvironmentAutomationServiceListParamsEncoding string

const (
	EnvironmentAutomationServiceListParamsEncodingProto EnvironmentAutomationServiceListParamsEncoding = "proto"
	EnvironmentAutomationServiceListParamsEncodingJson  EnvironmentAutomationServiceListParamsEncoding = "json"
)

func (r EnvironmentAutomationServiceListParamsEncoding) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceListParamsEncodingProto, EnvironmentAutomationServiceListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EnvironmentAutomationServiceListParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationServiceListParamsConnectProtocolVersion1 EnvironmentAutomationServiceListParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationServiceListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type EnvironmentAutomationServiceListParamsCompression string

const (
	EnvironmentAutomationServiceListParamsCompressionIdentity EnvironmentAutomationServiceListParamsCompression = "identity"
	EnvironmentAutomationServiceListParamsCompressionGzip     EnvironmentAutomationServiceListParamsCompression = "gzip"
	EnvironmentAutomationServiceListParamsCompressionBr       EnvironmentAutomationServiceListParamsCompression = "br"
)

func (r EnvironmentAutomationServiceListParamsCompression) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceListParamsCompressionIdentity, EnvironmentAutomationServiceListParamsCompressionGzip, EnvironmentAutomationServiceListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EnvironmentAutomationServiceListParamsConnect string

const (
	EnvironmentAutomationServiceListParamsConnectV1 EnvironmentAutomationServiceListParamsConnect = "v1"
)

func (r EnvironmentAutomationServiceListParamsConnect) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceListParamsConnectV1:
		return true
	}
	return false
}

type EnvironmentAutomationServiceDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationServiceDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                         `json:"id" format:"uuid"`
	Force                  param.Field[bool]                                                           `json:"force"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationServiceDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationServiceDeleteParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationServiceDeleteParamsConnectProtocolVersion1 EnvironmentAutomationServiceDeleteParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationServiceDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentAutomationServiceStartParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationServiceStartParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                        `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationServiceStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationServiceStartParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationServiceStartParamsConnectProtocolVersion1 EnvironmentAutomationServiceStartParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationServiceStartParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceStartParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentAutomationServiceStopParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationServiceStopParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                       `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationServiceStopParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentAutomationServiceStopParamsConnectProtocolVersion float64

const (
	EnvironmentAutomationServiceStopParamsConnectProtocolVersion1 EnvironmentAutomationServiceStopParamsConnectProtocolVersion = 1
)

func (r EnvironmentAutomationServiceStopParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceStopParamsConnectProtocolVersion1:
		return true
	}
	return false
}
