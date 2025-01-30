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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByManual],
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostDevcontainerStart],
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostEnvironmentStart].
func (r EnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy) AsUnion() EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByUnion {
	return r.union
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// Union satisfied by
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByManual],
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostDevcontainerStart]
// or
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostEnvironmentStart].
type EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByUnion interface {
	implementsEnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByManual{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostDevcontainerStart{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostEnvironmentStart{}),
		},
	)
}

type EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByManual struct {
	Manual bool                                                                          `json:"manual,required"`
	JSON   environmentAutomationServiceListResponseServicesMetadataTriggeredByManualJSON `json:"-"`
}

// environmentAutomationServiceListResponseServicesMetadataTriggeredByManualJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByManual]
type environmentAutomationServiceListResponseServicesMetadataTriggeredByManualJSON struct {
	Manual      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByManual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesMetadataTriggeredByManualJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByManual) implementsEnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy() {
}

type EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostDevcontainerStart struct {
	PostDevcontainerStart bool                                                                                         `json:"postDevcontainerStart,required"`
	JSON                  environmentAutomationServiceListResponseServicesMetadataTriggeredByPostDevcontainerStartJSON `json:"-"`
}

// environmentAutomationServiceListResponseServicesMetadataTriggeredByPostDevcontainerStartJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostDevcontainerStart]
type environmentAutomationServiceListResponseServicesMetadataTriggeredByPostDevcontainerStartJSON struct {
	PostDevcontainerStart apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostDevcontainerStart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesMetadataTriggeredByPostDevcontainerStartJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostDevcontainerStart) implementsEnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy() {
}

type EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostEnvironmentStart struct {
	PostEnvironmentStart bool                                                                                        `json:"postEnvironmentStart,required"`
	JSON                 environmentAutomationServiceListResponseServicesMetadataTriggeredByPostEnvironmentStartJSON `json:"-"`
}

// environmentAutomationServiceListResponseServicesMetadataTriggeredByPostEnvironmentStartJSON
// contains the JSON metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostEnvironmentStart]
type environmentAutomationServiceListResponseServicesMetadataTriggeredByPostEnvironmentStartJSON struct {
	PostEnvironmentStart apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostEnvironmentStart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesMetadataTriggeredByPostEnvironmentStartJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentAutomationServiceListResponseServicesMetadataTriggeredByPostEnvironmentStart) implementsEnvironmentAutomationServiceListResponseServicesMetadataTriggeredBy() {
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
	Docker EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker `json:"docker,required"`
	JSON   environmentAutomationServiceListResponseServicesSpecRunsOnJSON   `json:"-"`
}

// environmentAutomationServiceListResponseServicesSpecRunsOnJSON contains the JSON
// metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesSpecRunsOn]
type environmentAutomationServiceListResponseServicesSpecRunsOnJSON struct {
	Docker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesSpecRunsOn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesSpecRunsOnJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker struct {
	Environment []string                                                             `json:"environment"`
	Image       string                                                               `json:"image"`
	JSON        environmentAutomationServiceListResponseServicesSpecRunsOnDockerJSON `json:"-"`
}

// environmentAutomationServiceListResponseServicesSpecRunsOnDockerJSON contains
// the JSON metadata for the struct
// [EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker]
type environmentAutomationServiceListResponseServicesSpecRunsOnDockerJSON struct {
	Environment apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationServiceListResponseServicesSpecRunsOnDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationServiceListResponseServicesSpecRunsOnDockerJSON) RawJSON() string {
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

type EnvironmentAutomationServiceUpdateParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationServiceUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                         `json:"id" format:"uuid"`
	Metadata               param.Field[EnvironmentAutomationServiceUpdateParamsMetadataUnion]          `json:"metadata"`
	// Changing the spec of a service is a complex operation. The spec of a service
	//
	// can only be updated if the service is in a stopped state. If the service is
	// running, it must be stopped first.
	Spec param.Field[EnvironmentAutomationServiceUpdateParamsSpecUnion] `json:"spec"`
	// Service status updates are only expected from the executing environment. As a
	// client
	//
	// of this API you are not expected to provide this field. Updating this field
	// requires the `environmentservice:update_status` permission.
	Status param.Field[EnvironmentAutomationServiceUpdateParamsStatusUnion] `json:"status"`
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
	Description param.Field[string]      `json:"description"`
	Name        param.Field[string]      `json:"name"`
	TriggeredBy param.Field[interface{}] `json:"triggeredBy"`
}

func (r EnvironmentAutomationServiceUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsMetadata) implementsEnvironmentAutomationServiceUpdateParamsMetadataUnion() {
}

// Satisfied by [EnvironmentAutomationServiceUpdateParamsMetadataDescription],
// [EnvironmentAutomationServiceUpdateParamsMetadataName],
// [EnvironmentAutomationServiceUpdateParamsMetadataTriggeredBy],
// [EnvironmentAutomationServiceUpdateParamsMetadata].
type EnvironmentAutomationServiceUpdateParamsMetadataUnion interface {
	implementsEnvironmentAutomationServiceUpdateParamsMetadataUnion()
}

type EnvironmentAutomationServiceUpdateParamsMetadataDescription struct {
	Description param.Field[string] `json:"description,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataDescription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataDescription) implementsEnvironmentAutomationServiceUpdateParamsMetadataUnion() {
}

type EnvironmentAutomationServiceUpdateParamsMetadataName struct {
	Name param.Field[string] `json:"name,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataName) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataName) implementsEnvironmentAutomationServiceUpdateParamsMetadataUnion() {
}

type EnvironmentAutomationServiceUpdateParamsMetadataTriggeredBy struct {
	TriggeredBy param.Field[EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredBy] `json:"triggeredBy,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredBy) implementsEnvironmentAutomationServiceUpdateParamsMetadataUnion() {
}

type EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredBy struct {
	Trigger param.Field[[]EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerUnion] `json:"trigger"`
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
type EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTrigger struct {
	Manual                param.Field[bool] `json:"manual"`
	PostDevcontainerStart param.Field[bool] `json:"postDevcontainerStart"`
	PostEnvironmentStart  param.Field[bool] `json:"postEnvironmentStart"`
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTrigger) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTrigger) implementsEnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerUnion() {
}

// An AutomationTrigger represents a trigger for an automation action. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started. The `post_devcontainer_start` field indicates
// that the automation should be triggered after the dev container has started.
//
// Satisfied by
// [EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerManual],
// [EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerPostDevcontainerStart],
// [EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerPostEnvironmentStart],
// [EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTrigger].
type EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerUnion interface {
	implementsEnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerUnion()
}

type EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerManual struct {
	Manual param.Field[bool] `json:"manual,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerManual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerManual) implementsEnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerUnion() {
}

type EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerPostDevcontainerStart struct {
	PostDevcontainerStart param.Field[bool] `json:"postDevcontainerStart,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerPostDevcontainerStart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerPostDevcontainerStart) implementsEnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerUnion() {
}

type EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerPostEnvironmentStart struct {
	PostEnvironmentStart param.Field[bool] `json:"postEnvironmentStart,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerPostEnvironmentStart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerPostEnvironmentStart) implementsEnvironmentAutomationServiceUpdateParamsMetadataTriggeredByTriggeredByTriggerUnion() {
}

// Changing the spec of a service is a complex operation. The spec of a service
//
// can only be updated if the service is in a stopped state. If the service is
// running, it must be stopped first.
type EnvironmentAutomationServiceUpdateParamsSpec struct {
	Commands param.Field[interface{}] `json:"commands"`
	RunsOn   param.Field[interface{}] `json:"runsOn"`
}

func (r EnvironmentAutomationServiceUpdateParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsSpec) implementsEnvironmentAutomationServiceUpdateParamsSpecUnion() {
}

// Changing the spec of a service is a complex operation. The spec of a service
//
// can only be updated if the service is in a stopped state. If the service is
// running, it must be stopped first.
//
// Satisfied by [EnvironmentAutomationServiceUpdateParamsSpecCommands],
// [EnvironmentAutomationServiceUpdateParamsSpecRunsOn],
// [EnvironmentAutomationServiceUpdateParamsSpec].
type EnvironmentAutomationServiceUpdateParamsSpecUnion interface {
	implementsEnvironmentAutomationServiceUpdateParamsSpecUnion()
}

type EnvironmentAutomationServiceUpdateParamsSpecCommands struct {
	Commands param.Field[EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsUnion] `json:"commands,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsSpecCommands) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsSpecCommands) implementsEnvironmentAutomationServiceUpdateParamsSpecUnion() {
}

type EnvironmentAutomationServiceUpdateParamsSpecCommandsCommands struct {
	Ready param.Field[string] `json:"ready"`
	Start param.Field[string] `json:"start"`
	Stop  param.Field[string] `json:"stop"`
}

func (r EnvironmentAutomationServiceUpdateParamsSpecCommandsCommands) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsSpecCommandsCommands) implementsEnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsUnion() {
}

// Satisfied by
// [EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsReady],
// [EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsStart],
// [EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsStop],
// [EnvironmentAutomationServiceUpdateParamsSpecCommandsCommands].
type EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsUnion interface {
	implementsEnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsUnion()
}

type EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsReady struct {
	Ready param.Field[string] `json:"ready,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsReady) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsReady) implementsEnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsUnion() {
}

type EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsStart struct {
	Start param.Field[string] `json:"start,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsStart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsStart) implementsEnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsUnion() {
}

type EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsStop struct {
	Stop param.Field[string] `json:"stop,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsStop) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsStop) implementsEnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsUnion() {
}

type EnvironmentAutomationServiceUpdateParamsSpecRunsOn struct {
	RunsOn param.Field[EnvironmentAutomationServiceUpdateParamsSpecRunsOnRunsOn] `json:"runsOn,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsSpecRunsOn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsSpecRunsOn) implementsEnvironmentAutomationServiceUpdateParamsSpecUnion() {
}

type EnvironmentAutomationServiceUpdateParamsSpecRunsOnRunsOn struct {
	Docker param.Field[EnvironmentAutomationServiceUpdateParamsSpecRunsOnRunsOnDocker] `json:"docker,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsSpecRunsOnRunsOn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationServiceUpdateParamsSpecRunsOnRunsOnDocker struct {
	Environment param.Field[[]string] `json:"environment"`
	Image       param.Field[string]   `json:"image"`
}

func (r EnvironmentAutomationServiceUpdateParamsSpecRunsOnRunsOnDocker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service status updates are only expected from the executing environment. As a
// client
//
// of this API you are not expected to provide this field. Updating this field
// requires the `environmentservice:update_status` permission.
type EnvironmentAutomationServiceUpdateParamsStatus struct {
	FailureMessage param.Field[string]                                              `json:"failureMessage"`
	LogURL         param.Field[string]                                              `json:"logUrl"`
	Phase          param.Field[EnvironmentAutomationServiceUpdateParamsStatusPhase] `json:"phase"`
	Session        param.Field[string]                                              `json:"session"`
}

func (r EnvironmentAutomationServiceUpdateParamsStatus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsStatus) implementsEnvironmentAutomationServiceUpdateParamsStatusUnion() {
}

// Service status updates are only expected from the executing environment. As a
// client
//
// of this API you are not expected to provide this field. Updating this field
// requires the `environmentservice:update_status` permission.
//
// Satisfied by [EnvironmentAutomationServiceUpdateParamsStatusFailureMessage],
// [EnvironmentAutomationServiceUpdateParamsStatusLogURL],
// [EnvironmentAutomationServiceUpdateParamsStatusPhase],
// [EnvironmentAutomationServiceUpdateParamsStatusSession],
// [EnvironmentAutomationServiceUpdateParamsStatus].
type EnvironmentAutomationServiceUpdateParamsStatusUnion interface {
	implementsEnvironmentAutomationServiceUpdateParamsStatusUnion()
}

type EnvironmentAutomationServiceUpdateParamsStatusFailureMessage struct {
	FailureMessage param.Field[string] `json:"failureMessage,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsStatusFailureMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsStatusFailureMessage) implementsEnvironmentAutomationServiceUpdateParamsStatusUnion() {
}

type EnvironmentAutomationServiceUpdateParamsStatusLogURL struct {
	LogURL param.Field[string] `json:"logUrl,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsStatusLogURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsStatusLogURL) implementsEnvironmentAutomationServiceUpdateParamsStatusUnion() {
}

type EnvironmentAutomationServiceUpdateParamsStatusPhase struct {
	Phase param.Field[EnvironmentAutomationServiceUpdateParamsStatusPhasePhase] `json:"phase,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsStatusPhase) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsStatusPhase) implementsEnvironmentAutomationServiceUpdateParamsStatusUnion() {
}

type EnvironmentAutomationServiceUpdateParamsStatusPhasePhase string

const (
	EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseUnspecified EnvironmentAutomationServiceUpdateParamsStatusPhasePhase = "SERVICE_PHASE_UNSPECIFIED"
	EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseStarting    EnvironmentAutomationServiceUpdateParamsStatusPhasePhase = "SERVICE_PHASE_STARTING"
	EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseRunning     EnvironmentAutomationServiceUpdateParamsStatusPhasePhase = "SERVICE_PHASE_RUNNING"
	EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseStopping    EnvironmentAutomationServiceUpdateParamsStatusPhasePhase = "SERVICE_PHASE_STOPPING"
	EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseStopped     EnvironmentAutomationServiceUpdateParamsStatusPhasePhase = "SERVICE_PHASE_STOPPED"
	EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseFailed      EnvironmentAutomationServiceUpdateParamsStatusPhasePhase = "SERVICE_PHASE_FAILED"
	EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseDeleted     EnvironmentAutomationServiceUpdateParamsStatusPhasePhase = "SERVICE_PHASE_DELETED"
)

func (r EnvironmentAutomationServiceUpdateParamsStatusPhasePhase) IsKnown() bool {
	switch r {
	case EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseUnspecified, EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseStarting, EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseRunning, EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseStopping, EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseStopped, EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseFailed, EnvironmentAutomationServiceUpdateParamsStatusPhasePhaseServicePhaseDeleted:
		return true
	}
	return false
}

type EnvironmentAutomationServiceUpdateParamsStatusSession struct {
	Session param.Field[string] `json:"session,required"`
}

func (r EnvironmentAutomationServiceUpdateParamsStatusSession) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentAutomationServiceUpdateParamsStatusSession) implementsEnvironmentAutomationServiceUpdateParamsStatusUnion() {
}

type EnvironmentAutomationServiceListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentAutomationServiceListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// filter contains the filter options for listing services
	Filter param.Field[EnvironmentAutomationServiceListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing services
	Pagination param.Field[EnvironmentAutomationServiceListParamsPagination] `json:"pagination"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentAutomationServiceListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	//
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
