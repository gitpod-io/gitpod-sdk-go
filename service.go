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

// ServiceService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceService] method instead.
type ServiceService struct {
	Options []option.RequestOption
}

// NewServiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewServiceService(opts ...option.RequestOption) (r *ServiceService) {
	r = &ServiceService{}
	r.Options = opts
	return
}

// UpdateService
func (r *ServiceService) Update(ctx context.Context, params ServiceUpdateParams, opts ...option.RequestOption) (res *ServiceUpdateResponse, err error) {
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
func (r *ServiceService) List(ctx context.Context, params ServiceListParams, opts ...option.RequestOption) (res *ServiceListResponse, err error) {
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
// deleted. If the service is not stopped it will be stopped before deletion.
func (r *ServiceService) Delete(ctx context.Context, params ServiceDeleteParams, opts ...option.RequestOption) (res *ServiceDeleteResponse, err error) {
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

// ListServices
func (r *ServiceService) ListNew(ctx context.Context, params ServiceListNewParams, opts ...option.RequestOption) (res *ServiceListNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/ListServices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// StartService starts a service. This call does not block until the service is
// started. This call will not error if the service is already running or has been
// started.
func (r *ServiceService) Start(ctx context.Context, params ServiceStartParams, opts ...option.RequestOption) (res *ServiceStartResponse, err error) {
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
// stopped. This call will not error if the service is already stopped or has been
// stopped.
func (r *ServiceService) Stop(ctx context.Context, params ServiceStopParams, opts ...option.RequestOption) (res *ServiceStopResponse, err error) {
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

type ServiceUpdateResponse = interface{}

type ServiceListResponse struct {
	Pagination ServiceListResponsePagination `json:"pagination"`
	Services   []ServiceListResponseService  `json:"services"`
	JSON       serviceListResponseJSON       `json:"-"`
}

// serviceListResponseJSON contains the JSON metadata for the struct
// [ServiceListResponse]
type serviceListResponseJSON struct {
	Pagination  apijson.Field
	Services    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseJSON) RawJSON() string {
	return r.raw
}

type ServiceListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                            `json:"nextToken"`
	JSON      serviceListResponsePaginationJSON `json:"-"`
}

// serviceListResponsePaginationJSON contains the JSON metadata for the struct
// [ServiceListResponsePagination]
type serviceListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ServiceListResponseService struct {
	ID            string                              `json:"id" format:"uuid"`
	EnvironmentID string                              `json:"environmentId" format:"uuid"`
	Metadata      ServiceListResponseServicesMetadata `json:"metadata"`
	Spec          ServiceListResponseServicesSpec     `json:"spec"`
	Status        ServiceListResponseServicesStatus   `json:"status"`
	JSON          serviceListResponseServiceJSON      `json:"-"`
}

// serviceListResponseServiceJSON contains the JSON metadata for the struct
// [ServiceListResponseService]
type serviceListResponseServiceJSON struct {
	ID            apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ServiceListResponseService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseServiceJSON) RawJSON() string {
	return r.raw
}

type ServiceListResponseServicesMetadata struct {
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
	// creator describes the principal who created the service.
	Creator ServiceListResponseServicesMetadataCreator `json:"creator"`
	// description is a user-facing description for the service. It can be used to
	// provide context and documentation for the service.
	Description string `json:"description"`
	// name is a user-facing name for the service. Unlike the reference, this field is
	// not unique, and not referenced by the system. This is a short descriptive name
	// for the service.
	Name string `json:"name"`
	// reference is a user-facing identifier for the service which must be unique on
	// the environment. It is used to express dependencies between services, and to
	// identify the service in user interactions (e.g. the CLI).
	Reference string `json:"reference"`
	// triggered_by is a list of trigger that start the service.
	TriggeredBy []ServiceListResponseServicesMetadataTriggeredBy `json:"triggeredBy"`
	JSON        serviceListResponseServicesMetadataJSON          `json:"-"`
}

// serviceListResponseServicesMetadataJSON contains the JSON metadata for the
// struct [ServiceListResponseServicesMetadata]
type serviceListResponseServicesMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponseServicesMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseServicesMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created the service.
type ServiceListResponseServicesMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ServiceListResponseServicesMetadataCreatorPrincipal `json:"principal"`
	JSON      serviceListResponseServicesMetadataCreatorJSON      `json:"-"`
}

// serviceListResponseServicesMetadataCreatorJSON contains the JSON metadata for
// the struct [ServiceListResponseServicesMetadataCreator]
type serviceListResponseServicesMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponseServicesMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseServicesMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ServiceListResponseServicesMetadataCreatorPrincipal string

const (
	ServiceListResponseServicesMetadataCreatorPrincipalPrincipalUnspecified    ServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	ServiceListResponseServicesMetadataCreatorPrincipalPrincipalAccount        ServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	ServiceListResponseServicesMetadataCreatorPrincipalPrincipalUser           ServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_USER"
	ServiceListResponseServicesMetadataCreatorPrincipalPrincipalRunner         ServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	ServiceListResponseServicesMetadataCreatorPrincipalPrincipalEnvironment    ServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	ServiceListResponseServicesMetadataCreatorPrincipalPrincipalServiceAccount ServiceListResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ServiceListResponseServicesMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case ServiceListResponseServicesMetadataCreatorPrincipalPrincipalUnspecified, ServiceListResponseServicesMetadataCreatorPrincipalPrincipalAccount, ServiceListResponseServicesMetadataCreatorPrincipalPrincipalUser, ServiceListResponseServicesMetadataCreatorPrincipalPrincipalRunner, ServiceListResponseServicesMetadataCreatorPrincipalPrincipalEnvironment, ServiceListResponseServicesMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the devcontainer has started.
type ServiceListResponseServicesMetadataTriggeredBy struct {
	JSON serviceListResponseServicesMetadataTriggeredByJSON `json:"-"`
}

// serviceListResponseServicesMetadataTriggeredByJSON contains the JSON metadata
// for the struct [ServiceListResponseServicesMetadataTriggeredBy]
type serviceListResponseServicesMetadataTriggeredByJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponseServicesMetadataTriggeredBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseServicesMetadataTriggeredByJSON) RawJSON() string {
	return r.raw
}

type ServiceListResponseServicesSpec struct {
	// commands contains the commands to start, stop and check the readiness of the
	// service
	Commands ServiceListResponseServicesSpecCommands `json:"commands"`
	// Phase is the desired phase of the environment
	DesiredPhase ServiceListResponseServicesSpecDesiredPhase `json:"desiredPhase"`
	// session should be changed to trigger a restart of the service. If a service
	// exits it will not be restarted until the session is changed.
	Session string `json:"session"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion ServiceListResponseServicesSpecSpecVersionUnion `json:"specVersion"`
	JSON        serviceListResponseServicesSpecJSON             `json:"-"`
}

// serviceListResponseServicesSpecJSON contains the JSON metadata for the struct
// [ServiceListResponseServicesSpec]
type serviceListResponseServicesSpecJSON struct {
	Commands     apijson.Field
	DesiredPhase apijson.Field
	Session      apijson.Field
	SpecVersion  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServiceListResponseServicesSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseServicesSpecJSON) RawJSON() string {
	return r.raw
}

// commands contains the commands to start, stop and check the readiness of the
// service
type ServiceListResponseServicesSpecCommands struct {
	// ready is an optional command that is run repeatedly until it exits with a zero
	// exit code. If set, the service will first go into a Starting phase, and then
	// into a Running phase once the ready command exits with a zero exit code.
	Ready string `json:"ready"`
	// start is the command to start and run the service. If start exits, the service
	// will transition to the following phase:
	//
	//   - Stopped: if the exit code is 0
	//   - Failed: if the exit code is not 0 If the stop command is not set, the start
	//     command will receive a SIGTERM signal when the service is requested to stop.
	//     If it does not exit within 2 minutes, it will receive a SIGKILL signal.
	Start string `json:"start"`
	// stop is an optional command that runs when the service is requested to stop. If
	// set, instead of sending a SIGTERM signal to the start command, the stop command
	// will be run. Once the stop command exits, the start command will receive a
	// SIGKILL signal. If the stop command exits with a non-zero exit code, the service
	// will transition to the Failed phase. If the stop command does not exit within 2
	// minutes, a SIGKILL signal will be sent to both the start and stop commands.
	Stop string                                      `json:"stop"`
	JSON serviceListResponseServicesSpecCommandsJSON `json:"-"`
}

// serviceListResponseServicesSpecCommandsJSON contains the JSON metadata for the
// struct [ServiceListResponseServicesSpecCommands]
type serviceListResponseServicesSpecCommandsJSON struct {
	Ready       apijson.Field
	Start       apijson.Field
	Stop        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponseServicesSpecCommands) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseServicesSpecCommandsJSON) RawJSON() string {
	return r.raw
}

// Phase is the desired phase of the environment
type ServiceListResponseServicesSpecDesiredPhase string

const (
	ServiceListResponseServicesSpecDesiredPhaseServicePhaseUnspecified ServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_UNSPECIFIED"
	ServiceListResponseServicesSpecDesiredPhaseServicePhaseStarting    ServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_STARTING"
	ServiceListResponseServicesSpecDesiredPhaseServicePhaseRunning     ServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_RUNNING"
	ServiceListResponseServicesSpecDesiredPhaseServicePhaseStopping    ServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_STOPPING"
	ServiceListResponseServicesSpecDesiredPhaseServicePhaseStopped     ServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_STOPPED"
	ServiceListResponseServicesSpecDesiredPhaseServicePhaseFailed      ServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_FAILED"
	ServiceListResponseServicesSpecDesiredPhaseServicePhaseDeleted     ServiceListResponseServicesSpecDesiredPhase = "SERVICE_PHASE_DELETED"
)

func (r ServiceListResponseServicesSpecDesiredPhase) IsKnown() bool {
	switch r {
	case ServiceListResponseServicesSpecDesiredPhaseServicePhaseUnspecified, ServiceListResponseServicesSpecDesiredPhaseServicePhaseStarting, ServiceListResponseServicesSpecDesiredPhaseServicePhaseRunning, ServiceListResponseServicesSpecDesiredPhaseServicePhaseStopping, ServiceListResponseServicesSpecDesiredPhaseServicePhaseStopped, ServiceListResponseServicesSpecDesiredPhaseServicePhaseFailed, ServiceListResponseServicesSpecDesiredPhaseServicePhaseDeleted:
		return true
	}
	return false
}

// version of the spec. The value of this field has no semantic meaning (e.g. don't
// interpret it as as a timestamp), but it can be used to impose a partial order.
// If a.spec_version < b.spec_version then a was the spec before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type ServiceListResponseServicesSpecSpecVersionUnion interface {
	ImplementsServiceListResponseServicesSpecSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ServiceListResponseServicesSpecSpecVersionUnion)(nil)).Elem(),
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

type ServiceListResponseServicesStatus struct {
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage string `json:"failureMessage"`
	// environment_url contains the URL at which the environment can be accessed. This
	// field is only set if the environment is running.
	LogURL string `json:"logUrl"`
	// the phase of an environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase ServiceListResponseServicesStatusPhase `json:"phase"`
	// session is the current session of the service.
	Session string `json:"session"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their statuus has different versions. The value of this field has no
	// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
	// to impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion ServiceListResponseServicesStatusStatusVersionUnion `json:"statusVersion"`
	JSON          serviceListResponseServicesStatusJSON               `json:"-"`
}

// serviceListResponseServicesStatusJSON contains the JSON metadata for the struct
// [ServiceListResponseServicesStatus]
type serviceListResponseServicesStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	StatusVersion  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ServiceListResponseServicesStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseServicesStatusJSON) RawJSON() string {
	return r.raw
}

// the phase of an environment is a simple, high-level summary of where the
// environment is in its lifecycle
type ServiceListResponseServicesStatusPhase string

const (
	ServiceListResponseServicesStatusPhaseServicePhaseUnspecified ServiceListResponseServicesStatusPhase = "SERVICE_PHASE_UNSPECIFIED"
	ServiceListResponseServicesStatusPhaseServicePhaseStarting    ServiceListResponseServicesStatusPhase = "SERVICE_PHASE_STARTING"
	ServiceListResponseServicesStatusPhaseServicePhaseRunning     ServiceListResponseServicesStatusPhase = "SERVICE_PHASE_RUNNING"
	ServiceListResponseServicesStatusPhaseServicePhaseStopping    ServiceListResponseServicesStatusPhase = "SERVICE_PHASE_STOPPING"
	ServiceListResponseServicesStatusPhaseServicePhaseStopped     ServiceListResponseServicesStatusPhase = "SERVICE_PHASE_STOPPED"
	ServiceListResponseServicesStatusPhaseServicePhaseFailed      ServiceListResponseServicesStatusPhase = "SERVICE_PHASE_FAILED"
	ServiceListResponseServicesStatusPhaseServicePhaseDeleted     ServiceListResponseServicesStatusPhase = "SERVICE_PHASE_DELETED"
)

func (r ServiceListResponseServicesStatusPhase) IsKnown() bool {
	switch r {
	case ServiceListResponseServicesStatusPhaseServicePhaseUnspecified, ServiceListResponseServicesStatusPhaseServicePhaseStarting, ServiceListResponseServicesStatusPhaseServicePhaseRunning, ServiceListResponseServicesStatusPhaseServicePhaseStopping, ServiceListResponseServicesStatusPhaseServicePhaseStopped, ServiceListResponseServicesStatusPhaseServicePhaseFailed, ServiceListResponseServicesStatusPhaseServicePhaseDeleted:
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
type ServiceListResponseServicesStatusStatusVersionUnion interface {
	ImplementsServiceListResponseServicesStatusStatusVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ServiceListResponseServicesStatusStatusVersionUnion)(nil)).Elem(),
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

type ServiceDeleteResponse = interface{}

type ServiceListNewResponse struct {
	Pagination ServiceListNewResponsePagination `json:"pagination"`
	Services   []ServiceListNewResponseService  `json:"services"`
	JSON       serviceListNewResponseJSON       `json:"-"`
}

// serviceListNewResponseJSON contains the JSON metadata for the struct
// [ServiceListNewResponse]
type serviceListNewResponseJSON struct {
	Pagination  apijson.Field
	Services    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListNewResponseJSON) RawJSON() string {
	return r.raw
}

type ServiceListNewResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                               `json:"nextToken"`
	JSON      serviceListNewResponsePaginationJSON `json:"-"`
}

// serviceListNewResponsePaginationJSON contains the JSON metadata for the struct
// [ServiceListNewResponsePagination]
type serviceListNewResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListNewResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListNewResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ServiceListNewResponseService struct {
	ID            string                                 `json:"id" format:"uuid"`
	EnvironmentID string                                 `json:"environmentId" format:"uuid"`
	Metadata      ServiceListNewResponseServicesMetadata `json:"metadata"`
	Spec          ServiceListNewResponseServicesSpec     `json:"spec"`
	Status        ServiceListNewResponseServicesStatus   `json:"status"`
	JSON          serviceListNewResponseServiceJSON      `json:"-"`
}

// serviceListNewResponseServiceJSON contains the JSON metadata for the struct
// [ServiceListNewResponseService]
type serviceListNewResponseServiceJSON struct {
	ID            apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ServiceListNewResponseService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListNewResponseServiceJSON) RawJSON() string {
	return r.raw
}

type ServiceListNewResponseServicesMetadata struct {
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
	// creator describes the principal who created the service.
	Creator ServiceListNewResponseServicesMetadataCreator `json:"creator"`
	// description is a user-facing description for the service. It can be used to
	// provide context and documentation for the service.
	Description string `json:"description"`
	// name is a user-facing name for the service. Unlike the reference, this field is
	// not unique, and not referenced by the system. This is a short descriptive name
	// for the service.
	Name string `json:"name"`
	// reference is a user-facing identifier for the service which must be unique on
	// the environment. It is used to express dependencies between services, and to
	// identify the service in user interactions (e.g. the CLI).
	Reference string `json:"reference"`
	// triggered_by is a list of trigger that start the service.
	TriggeredBy []ServiceListNewResponseServicesMetadataTriggeredBy `json:"triggeredBy"`
	JSON        serviceListNewResponseServicesMetadataJSON          `json:"-"`
}

// serviceListNewResponseServicesMetadataJSON contains the JSON metadata for the
// struct [ServiceListNewResponseServicesMetadata]
type serviceListNewResponseServicesMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListNewResponseServicesMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListNewResponseServicesMetadataJSON) RawJSON() string {
	return r.raw
}

// creator describes the principal who created the service.
type ServiceListNewResponseServicesMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ServiceListNewResponseServicesMetadataCreatorPrincipal `json:"principal"`
	JSON      serviceListNewResponseServicesMetadataCreatorJSON      `json:"-"`
}

// serviceListNewResponseServicesMetadataCreatorJSON contains the JSON metadata for
// the struct [ServiceListNewResponseServicesMetadataCreator]
type serviceListNewResponseServicesMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListNewResponseServicesMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListNewResponseServicesMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ServiceListNewResponseServicesMetadataCreatorPrincipal string

const (
	ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalUnspecified    ServiceListNewResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalAccount        ServiceListNewResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalUser           ServiceListNewResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_USER"
	ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalRunner         ServiceListNewResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalEnvironment    ServiceListNewResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalServiceAccount ServiceListNewResponseServicesMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ServiceListNewResponseServicesMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalUnspecified, ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalAccount, ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalUser, ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalRunner, ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalEnvironment, ServiceListNewResponseServicesMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the devcontainer has started.
type ServiceListNewResponseServicesMetadataTriggeredBy struct {
	JSON serviceListNewResponseServicesMetadataTriggeredByJSON `json:"-"`
}

// serviceListNewResponseServicesMetadataTriggeredByJSON contains the JSON metadata
// for the struct [ServiceListNewResponseServicesMetadataTriggeredBy]
type serviceListNewResponseServicesMetadataTriggeredByJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListNewResponseServicesMetadataTriggeredBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListNewResponseServicesMetadataTriggeredByJSON) RawJSON() string {
	return r.raw
}

type ServiceListNewResponseServicesSpec struct {
	// commands contains the commands to start, stop and check the readiness of the
	// service
	Commands ServiceListNewResponseServicesSpecCommands `json:"commands"`
	// Phase is the desired phase of the environment
	DesiredPhase ServiceListNewResponseServicesSpecDesiredPhase `json:"desiredPhase"`
	// session should be changed to trigger a restart of the service. If a service
	// exits it will not be restarted until the session is changed.
	Session string `json:"session"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion ServiceListNewResponseServicesSpecSpecVersionUnion `json:"specVersion"`
	JSON        serviceListNewResponseServicesSpecJSON             `json:"-"`
}

// serviceListNewResponseServicesSpecJSON contains the JSON metadata for the struct
// [ServiceListNewResponseServicesSpec]
type serviceListNewResponseServicesSpecJSON struct {
	Commands     apijson.Field
	DesiredPhase apijson.Field
	Session      apijson.Field
	SpecVersion  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServiceListNewResponseServicesSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListNewResponseServicesSpecJSON) RawJSON() string {
	return r.raw
}

// commands contains the commands to start, stop and check the readiness of the
// service
type ServiceListNewResponseServicesSpecCommands struct {
	// ready is an optional command that is run repeatedly until it exits with a zero
	// exit code. If set, the service will first go into a Starting phase, and then
	// into a Running phase once the ready command exits with a zero exit code.
	Ready string `json:"ready"`
	// start is the command to start and run the service. If start exits, the service
	// will transition to the following phase:
	//
	//   - Stopped: if the exit code is 0
	//   - Failed: if the exit code is not 0 If the stop command is not set, the start
	//     command will receive a SIGTERM signal when the service is requested to stop.
	//     If it does not exit within 2 minutes, it will receive a SIGKILL signal.
	Start string `json:"start"`
	// stop is an optional command that runs when the service is requested to stop. If
	// set, instead of sending a SIGTERM signal to the start command, the stop command
	// will be run. Once the stop command exits, the start command will receive a
	// SIGKILL signal. If the stop command exits with a non-zero exit code, the service
	// will transition to the Failed phase. If the stop command does not exit within 2
	// minutes, a SIGKILL signal will be sent to both the start and stop commands.
	Stop string                                         `json:"stop"`
	JSON serviceListNewResponseServicesSpecCommandsJSON `json:"-"`
}

// serviceListNewResponseServicesSpecCommandsJSON contains the JSON metadata for
// the struct [ServiceListNewResponseServicesSpecCommands]
type serviceListNewResponseServicesSpecCommandsJSON struct {
	Ready       apijson.Field
	Start       apijson.Field
	Stop        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListNewResponseServicesSpecCommands) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListNewResponseServicesSpecCommandsJSON) RawJSON() string {
	return r.raw
}

// Phase is the desired phase of the environment
type ServiceListNewResponseServicesSpecDesiredPhase string

const (
	ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseUnspecified ServiceListNewResponseServicesSpecDesiredPhase = "SERVICE_PHASE_UNSPECIFIED"
	ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseStarting    ServiceListNewResponseServicesSpecDesiredPhase = "SERVICE_PHASE_STARTING"
	ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseRunning     ServiceListNewResponseServicesSpecDesiredPhase = "SERVICE_PHASE_RUNNING"
	ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseStopping    ServiceListNewResponseServicesSpecDesiredPhase = "SERVICE_PHASE_STOPPING"
	ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseStopped     ServiceListNewResponseServicesSpecDesiredPhase = "SERVICE_PHASE_STOPPED"
	ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseFailed      ServiceListNewResponseServicesSpecDesiredPhase = "SERVICE_PHASE_FAILED"
	ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseDeleted     ServiceListNewResponseServicesSpecDesiredPhase = "SERVICE_PHASE_DELETED"
)

func (r ServiceListNewResponseServicesSpecDesiredPhase) IsKnown() bool {
	switch r {
	case ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseUnspecified, ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseStarting, ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseRunning, ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseStopping, ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseStopped, ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseFailed, ServiceListNewResponseServicesSpecDesiredPhaseServicePhaseDeleted:
		return true
	}
	return false
}

// version of the spec. The value of this field has no semantic meaning (e.g. don't
// interpret it as as a timestamp), but it can be used to impose a partial order.
// If a.spec_version < b.spec_version then a was the spec before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type ServiceListNewResponseServicesSpecSpecVersionUnion interface {
	ImplementsServiceListNewResponseServicesSpecSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ServiceListNewResponseServicesSpecSpecVersionUnion)(nil)).Elem(),
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

type ServiceListNewResponseServicesStatus struct {
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage string `json:"failureMessage"`
	// environment_url contains the URL at which the environment can be accessed. This
	// field is only set if the environment is running.
	LogURL string `json:"logUrl"`
	// the phase of an environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase ServiceListNewResponseServicesStatusPhase `json:"phase"`
	// session is the current session of the service.
	Session string `json:"session"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their statuus has different versions. The value of this field has no
	// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
	// to impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion ServiceListNewResponseServicesStatusStatusVersionUnion `json:"statusVersion"`
	JSON          serviceListNewResponseServicesStatusJSON               `json:"-"`
}

// serviceListNewResponseServicesStatusJSON contains the JSON metadata for the
// struct [ServiceListNewResponseServicesStatus]
type serviceListNewResponseServicesStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	StatusVersion  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ServiceListNewResponseServicesStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListNewResponseServicesStatusJSON) RawJSON() string {
	return r.raw
}

// the phase of an environment is a simple, high-level summary of where the
// environment is in its lifecycle
type ServiceListNewResponseServicesStatusPhase string

const (
	ServiceListNewResponseServicesStatusPhaseServicePhaseUnspecified ServiceListNewResponseServicesStatusPhase = "SERVICE_PHASE_UNSPECIFIED"
	ServiceListNewResponseServicesStatusPhaseServicePhaseStarting    ServiceListNewResponseServicesStatusPhase = "SERVICE_PHASE_STARTING"
	ServiceListNewResponseServicesStatusPhaseServicePhaseRunning     ServiceListNewResponseServicesStatusPhase = "SERVICE_PHASE_RUNNING"
	ServiceListNewResponseServicesStatusPhaseServicePhaseStopping    ServiceListNewResponseServicesStatusPhase = "SERVICE_PHASE_STOPPING"
	ServiceListNewResponseServicesStatusPhaseServicePhaseStopped     ServiceListNewResponseServicesStatusPhase = "SERVICE_PHASE_STOPPED"
	ServiceListNewResponseServicesStatusPhaseServicePhaseFailed      ServiceListNewResponseServicesStatusPhase = "SERVICE_PHASE_FAILED"
	ServiceListNewResponseServicesStatusPhaseServicePhaseDeleted     ServiceListNewResponseServicesStatusPhase = "SERVICE_PHASE_DELETED"
)

func (r ServiceListNewResponseServicesStatusPhase) IsKnown() bool {
	switch r {
	case ServiceListNewResponseServicesStatusPhaseServicePhaseUnspecified, ServiceListNewResponseServicesStatusPhaseServicePhaseStarting, ServiceListNewResponseServicesStatusPhaseServicePhaseRunning, ServiceListNewResponseServicesStatusPhaseServicePhaseStopping, ServiceListNewResponseServicesStatusPhaseServicePhaseStopped, ServiceListNewResponseServicesStatusPhaseServicePhaseFailed, ServiceListNewResponseServicesStatusPhaseServicePhaseDeleted:
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
type ServiceListNewResponseServicesStatusStatusVersionUnion interface {
	ImplementsServiceListNewResponseServicesStatusStatusVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ServiceListNewResponseServicesStatusStatusVersionUnion)(nil)).Elem(),
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

type ServiceStartResponse = interface{}

type ServiceStopResponse = interface{}

type ServiceUpdateParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ServiceUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                    `json:"id" format:"uuid"`
	Metadata               param.Field[ServiceUpdateParamsMetadata]               `json:"metadata"`
	// Changing the spec of a service is a complex operation. The spec of a service can
	// only be updated if the service is in a stopped state. If the service is running,
	// it must be stopped first.
	Spec param.Field[ServiceUpdateParamsSpecUnion] `json:"spec"`
	// Service status updates are only expected from the executing environment. As a
	// client of this API you are not expected to provide this field. Updating this
	// field requires the `environmentservice:update_status` permission.
	Status param.Field[ServiceUpdateParamsStatus] `json:"status"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ServiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ServiceUpdateParamsConnectProtocolVersion float64

const (
	ServiceUpdateParamsConnectProtocolVersion1 ServiceUpdateParamsConnectProtocolVersion = 1
)

func (r ServiceUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ServiceUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ServiceUpdateParamsMetadata struct {
}

func (r ServiceUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Changing the spec of a service is a complex operation. The spec of a service can
// only be updated if the service is in a stopped state. If the service is running,
// it must be stopped first.
//
// Satisfied by [ServiceUpdateParamsSpecUnknown], [ServiceUpdateParamsSpecUnknown].
type ServiceUpdateParamsSpecUnion interface {
	implementsServiceUpdateParamsSpecUnion()
}

// Service status updates are only expected from the executing environment. As a
// client of this API you are not expected to provide this field. Updating this
// field requires the `environmentservice:update_status` permission.
type ServiceUpdateParamsStatus struct {
}

func (r ServiceUpdateParamsStatus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ServiceListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Base64                 param.Field[string]                                  `query:"base64"`
	Compression            param.Field[string]                                  `query:"compression"`
	Connect                param.Field[string]                                  `query:"connect"`
	Encoding               param.Field[string]                                  `query:"encoding"`
	Message                param.Field[string]                                  `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [ServiceListParams]'s query parameters as `url.Values`.
func (r ServiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define the version of the Connect protocol
type ServiceListParamsConnectProtocolVersion float64

const (
	ServiceListParamsConnectProtocolVersion1 ServiceListParamsConnectProtocolVersion = 1
)

func (r ServiceListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ServiceListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ServiceDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ServiceDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                    `json:"id" format:"uuid"`
	Force                  param.Field[bool]                                      `json:"force"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ServiceDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ServiceDeleteParamsConnectProtocolVersion float64

const (
	ServiceDeleteParamsConnectProtocolVersion1 ServiceDeleteParamsConnectProtocolVersion = 1
)

func (r ServiceDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ServiceDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ServiceListNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ServiceListNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// filter contains the filter options for listing services
	Filter param.Field[ServiceListNewParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing environments
	Pagination param.Field[ServiceListNewParamsPagination] `json:"pagination"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ServiceListNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ServiceListNewParamsConnectProtocolVersion float64

const (
	ServiceListNewParamsConnectProtocolVersion1 ServiceListNewParamsConnectProtocolVersion = 1
)

func (r ServiceListNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ServiceListNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// filter contains the filter options for listing services
type ServiceListNewParamsFilter struct {
	// environment_ids filters the response to only services of these environments
	EnvironmentIDs param.Field[[]string] `json:"environmentIds" format:"uuid"`
	// references filters the response to only services with these references
	References param.Field[[]string] `json:"references"`
	// service_ids filters the response to only services with these IDs
	ServiceIDs param.Field[[]string] `json:"serviceIds" format:"uuid"`
}

func (r ServiceListNewParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing environments
type ServiceListNewParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r ServiceListNewParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceStartParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ServiceStartParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                   `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ServiceStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ServiceStartParamsConnectProtocolVersion float64

const (
	ServiceStartParamsConnectProtocolVersion1 ServiceStartParamsConnectProtocolVersion = 1
)

func (r ServiceStartParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ServiceStartParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ServiceStopParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ServiceStopParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                  `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ServiceStopParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ServiceStopParamsConnectProtocolVersion float64

const (
	ServiceStopParamsConnectProtocolVersion1 ServiceStopParamsConnectProtocolVersion = 1
)

func (r ServiceStopParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ServiceStopParamsConnectProtocolVersion1:
		return true
	}
	return false
}
