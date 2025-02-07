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
func (r *EnvironmentAutomationServiceService) New(ctx context.Context, body EnvironmentAutomationServiceNewParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/CreateService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetService
func (r *EnvironmentAutomationServiceService) Get(ctx context.Context, body EnvironmentAutomationServiceGetParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/GetService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateService
func (r *EnvironmentAutomationServiceService) Update(ctx context.Context, body EnvironmentAutomationServiceUpdateParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/UpdateService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListServices
func (r *EnvironmentAutomationServiceService) List(ctx context.Context, params EnvironmentAutomationServiceListParams, opts ...option.RequestOption) (res *pagination.ServicesPage[Service], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.EnvironmentAutomationService/ListServices"
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

// ListServices
func (r *EnvironmentAutomationServiceService) ListAutoPaging(ctx context.Context, params EnvironmentAutomationServiceListParams, opts ...option.RequestOption) *pagination.ServicesPageAutoPager[Service] {
	return pagination.NewServicesPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteService deletes a service. This call does not block until the service is
// deleted. If the service is not stopped it will be stopped before deletion.
func (r *EnvironmentAutomationServiceService) Delete(ctx context.Context, body EnvironmentAutomationServiceDeleteParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/DeleteService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// StartService starts a service. This call does not block until the service is
// started. This call will not error if the service is already running or has been
// started.
func (r *EnvironmentAutomationServiceService) Start(ctx context.Context, body EnvironmentAutomationServiceStartParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceStartResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StartService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// StopService stops a service. This call does not block until the service is
// stopped. This call will not error if the service is already stopped or has been
// stopped.
func (r *EnvironmentAutomationServiceService) Stop(ctx context.Context, body EnvironmentAutomationServiceStopParams, opts ...option.RequestOption) (res *EnvironmentAutomationServiceStopResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StopService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Service struct {
	ID            string          `json:"id" format:"uuid"`
	EnvironmentID string          `json:"environmentId" format:"uuid"`
	Metadata      ServiceMetadata `json:"metadata"`
	Spec          ServiceSpec     `json:"spec"`
	Status        ServiceStatus   `json:"status"`
	JSON          serviceJSON     `json:"-"`
}

// serviceJSON contains the JSON metadata for the struct [Service]
type serviceJSON struct {
	ID            apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Service) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceJSON) RawJSON() string {
	return r.raw
}

type ServiceMetadata struct {
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
	Creator shared.Subject `json:"creator"`
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
	TriggeredBy []shared.AutomationTrigger `json:"triggeredBy"`
	JSON        serviceMetadataJSON        `json:"-"`
}

// serviceMetadataJSON contains the JSON metadata for the struct [ServiceMetadata]
type serviceMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceMetadataJSON) RawJSON() string {
	return r.raw
}

type ServiceMetadataParam struct {
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
	// creator describes the principal who created the service.
	Creator param.Field[shared.SubjectParam] `json:"creator"`
	// description is a user-facing description for the service. It can be used to
	// provide context and documentation for the service.
	Description param.Field[string] `json:"description"`
	// name is a user-facing name for the service. Unlike the reference, this field is
	// not unique, and not referenced by the system. This is a short descriptive name
	// for the service.
	Name param.Field[string] `json:"name"`
	// reference is a user-facing identifier for the service which must be unique on
	// the environment. It is used to express dependencies between services, and to
	// identify the service in user interactions (e.g. the CLI).
	Reference param.Field[string] `json:"reference"`
	// triggered_by is a list of trigger that start the service.
	TriggeredBy param.Field[[]shared.AutomationTriggerParam] `json:"triggeredBy"`
}

func (r ServiceMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServicePhase string

const (
	ServicePhaseUnspecified ServicePhase = "SERVICE_PHASE_UNSPECIFIED"
	ServicePhaseStarting    ServicePhase = "SERVICE_PHASE_STARTING"
	ServicePhaseRunning     ServicePhase = "SERVICE_PHASE_RUNNING"
	ServicePhaseStopping    ServicePhase = "SERVICE_PHASE_STOPPING"
	ServicePhaseStopped     ServicePhase = "SERVICE_PHASE_STOPPED"
	ServicePhaseFailed      ServicePhase = "SERVICE_PHASE_FAILED"
	ServicePhaseDeleted     ServicePhase = "SERVICE_PHASE_DELETED"
)

func (r ServicePhase) IsKnown() bool {
	switch r {
	case ServicePhaseUnspecified, ServicePhaseStarting, ServicePhaseRunning, ServicePhaseStopping, ServicePhaseStopped, ServicePhaseFailed, ServicePhaseDeleted:
		return true
	}
	return false
}

type ServiceSpec struct {
	// commands contains the commands to start, stop and check the readiness of the
	// service
	Commands ServiceSpecCommands `json:"commands"`
	// desired_phase is the phase the service should be in. Used to start or stop the
	// service.
	DesiredPhase ServicePhase `json:"desiredPhase"`
	// runs_on specifies the environment the service should run on.
	RunsOn shared.RunsOn `json:"runsOn"`
	// session should be changed to trigger a restart of the service. If a service
	// exits it will not be restarted until the session is changed.
	Session string `json:"session"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion string          `json:"specVersion"`
	JSON        serviceSpecJSON `json:"-"`
}

// serviceSpecJSON contains the JSON metadata for the struct [ServiceSpec]
type serviceSpecJSON struct {
	Commands     apijson.Field
	DesiredPhase apijson.Field
	RunsOn       apijson.Field
	Session      apijson.Field
	SpecVersion  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServiceSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceSpecJSON) RawJSON() string {
	return r.raw
}

// commands contains the commands to start, stop and check the readiness of the
// service
type ServiceSpecCommands struct {
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
	Stop string                  `json:"stop"`
	JSON serviceSpecCommandsJSON `json:"-"`
}

// serviceSpecCommandsJSON contains the JSON metadata for the struct
// [ServiceSpecCommands]
type serviceSpecCommandsJSON struct {
	Ready       apijson.Field
	Start       apijson.Field
	Stop        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceSpecCommands) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceSpecCommandsJSON) RawJSON() string {
	return r.raw
}

type ServiceSpecParam struct {
	// commands contains the commands to start, stop and check the readiness of the
	// service
	Commands param.Field[ServiceSpecCommandsParam] `json:"commands"`
	// desired_phase is the phase the service should be in. Used to start or stop the
	// service.
	DesiredPhase param.Field[ServicePhase] `json:"desiredPhase"`
	// runs_on specifies the environment the service should run on.
	RunsOn param.Field[shared.RunsOnParam] `json:"runsOn"`
	// session should be changed to trigger a restart of the service. If a service
	// exits it will not be restarted until the session is changed.
	Session param.Field[string] `json:"session"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion param.Field[string] `json:"specVersion"`
}

func (r ServiceSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// commands contains the commands to start, stop and check the readiness of the
// service
type ServiceSpecCommandsParam struct {
	// ready is an optional command that is run repeatedly until it exits with a zero
	// exit code. If set, the service will first go into a Starting phase, and then
	// into a Running phase once the ready command exits with a zero exit code.
	Ready param.Field[string] `json:"ready"`
	// start is the command to start and run the service. If start exits, the service
	// will transition to the following phase:
	//
	//   - Stopped: if the exit code is 0
	//   - Failed: if the exit code is not 0 If the stop command is not set, the start
	//     command will receive a SIGTERM signal when the service is requested to stop.
	//     If it does not exit within 2 minutes, it will receive a SIGKILL signal.
	Start param.Field[string] `json:"start"`
	// stop is an optional command that runs when the service is requested to stop. If
	// set, instead of sending a SIGTERM signal to the start command, the stop command
	// will be run. Once the stop command exits, the start command will receive a
	// SIGKILL signal. If the stop command exits with a non-zero exit code, the service
	// will transition to the Failed phase. If the stop command does not exit within 2
	// minutes, a SIGKILL signal will be sent to both the start and stop commands.
	Stop param.Field[string] `json:"stop"`
}

func (r ServiceSpecCommandsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceStatus struct {
	// failure_message summarises why the service failed to operate. If this is
	// non-empty the service has failed to operate and will likely transition to a
	// failed state.
	FailureMessage string `json:"failureMessage"`
	// log_url contains the URL at which the service logs can be accessed.
	LogURL string `json:"logUrl"`
	// phase is the current phase of the service.
	Phase ServicePhase `json:"phase"`
	// session is the current session of the service.
	Session string `json:"session"`
	// version of the status update. Service instances themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string            `json:"statusVersion"`
	JSON          serviceStatusJSON `json:"-"`
}

// serviceStatusJSON contains the JSON metadata for the struct [ServiceStatus]
type serviceStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	StatusVersion  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ServiceStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceStatusJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceNewResponse struct {
	Service Service                                     `json:"service"`
	JSON    environmentAutomationServiceNewResponseJSON `json:"-"`
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

type EnvironmentAutomationServiceGetResponse struct {
	Service Service                                     `json:"service"`
	JSON    environmentAutomationServiceGetResponseJSON `json:"-"`
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

type EnvironmentAutomationServiceUpdateResponse = interface{}

type EnvironmentAutomationServiceDeleteResponse = interface{}

type EnvironmentAutomationServiceStartResponse = interface{}

type EnvironmentAutomationServiceStopResponse = interface{}

type EnvironmentAutomationServiceNewParams struct {
	EnvironmentID param.Field[string]               `json:"environmentId" format:"uuid"`
	Metadata      param.Field[ServiceMetadataParam] `json:"metadata"`
	Spec          param.Field[ServiceSpecParam]     `json:"spec"`
}

func (r EnvironmentAutomationServiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceGetParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationServiceGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceUpdateParams struct {
	ID       param.Field[string]                                           `json:"id" format:"uuid"`
	Metadata param.Field[EnvironmentAutomationServiceUpdateParamsMetadata] `json:"metadata"`
	// Changing the spec of a service is a complex operation. The spec of a service can
	// only be updated if the service is in a stopped state. If the service is running,
	// it must be stopped first.
	Spec param.Field[EnvironmentAutomationServiceUpdateParamsSpec] `json:"spec"`
	// Service status updates are only expected from the executing environment. As a
	// client of this API you are not expected to provide this field. Updating this
	// field requires the `environmentservice:update_status` permission.
	Status param.Field[EnvironmentAutomationServiceUpdateParamsStatus] `json:"status"`
}

func (r EnvironmentAutomationServiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceUpdateParamsMetadata struct {
	Description param.Field[string]                                                      `json:"description"`
	Name        param.Field[string]                                                      `json:"name"`
	TriggeredBy param.Field[EnvironmentAutomationServiceUpdateParamsMetadataTriggeredBy] `json:"triggeredBy"`
}

func (r EnvironmentAutomationServiceUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceUpdateParamsMetadataTriggeredBy struct {
	Trigger param.Field[[]shared.AutomationTriggerParam] `json:"trigger"`
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Changing the spec of a service is a complex operation. The spec of a service can
// only be updated if the service is in a stopped state. If the service is running,
// it must be stopped first.
type EnvironmentAutomationServiceUpdateParamsSpec struct {
	Commands param.Field[EnvironmentAutomationServiceUpdateParamsSpecCommands] `json:"commands"`
	RunsOn   param.Field[shared.RunsOnParam]                                   `json:"runsOn"`
}

func (r EnvironmentAutomationServiceUpdateParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceUpdateParamsSpecCommands struct {
	Ready param.Field[string] `json:"ready"`
	Start param.Field[string] `json:"start"`
	Stop  param.Field[string] `json:"stop"`
}

func (r EnvironmentAutomationServiceUpdateParamsSpecCommands) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service status updates are only expected from the executing environment. As a
// client of this API you are not expected to provide this field. Updating this
// field requires the `environmentservice:update_status` permission.
type EnvironmentAutomationServiceUpdateParamsStatus struct {
	FailureMessage param.Field[string]       `json:"failureMessage"`
	LogURL         param.Field[string]       `json:"logUrl"`
	Phase          param.Field[ServicePhase] `json:"phase"`
	Session        param.Field[string]       `json:"session"`
}

func (r EnvironmentAutomationServiceUpdateParamsStatus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// filter contains the filter options for listing services
	Filter param.Field[EnvironmentAutomationServiceListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing services
	Pagination param.Field[EnvironmentAutomationServiceListParamsPagination] `json:"pagination"`
}

func (r EnvironmentAutomationServiceListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [EnvironmentAutomationServiceListParams]'s query parameters
// as `url.Values`.
func (r EnvironmentAutomationServiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// filter contains the filter options for listing services
type EnvironmentAutomationServiceListParamsFilter struct {
	// environment_ids filters the response to only services of these environments
	EnvironmentIDs param.Field[[]string] `json:"environmentIds" format:"uuid"`
	// references filters the response to only services with these references
	References param.Field[[]string] `json:"references"`
	// service_ids filters the response to only services with these IDs
	ServiceIDs param.Field[[]string] `json:"serviceIds" format:"uuid"`
}

func (r EnvironmentAutomationServiceListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing services
type EnvironmentAutomationServiceListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EnvironmentAutomationServiceListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceDeleteParams struct {
	ID    param.Field[string] `json:"id" format:"uuid"`
	Force param.Field[bool]   `json:"force"`
}

func (r EnvironmentAutomationServiceDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceStartParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationServiceStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceStopParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationServiceStopParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
