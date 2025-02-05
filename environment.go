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

// EnvironmentService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentService] method instead.
type EnvironmentService struct {
	Options     []option.RequestOption
	Automations *EnvironmentAutomationService
	Classes     *EnvironmentClassService
}

// NewEnvironmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnvironmentService(opts ...option.RequestOption) (r *EnvironmentService) {
	r = &EnvironmentService{}
	r.Options = opts
	r.Automations = NewEnvironmentAutomationService(opts...)
	r.Classes = NewEnvironmentClassService(opts...)
	return
}

// CreateEnvironment creates a new environment and starts it.
func (r *EnvironmentService) New(ctx context.Context, body EnvironmentNewParams, opts ...option.RequestOption) (res *EnvironmentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/CreateEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetEnvironment returns a single environment.
//
// +return NOT_FOUND User does not have access to an environment with the given ID
// +return NOT_FOUND Environment does not exist
func (r *EnvironmentService) Get(ctx context.Context, body EnvironmentGetParams, opts ...option.RequestOption) (res *EnvironmentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/GetEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateEnvironment updates the environment partially.
func (r *EnvironmentService) Update(ctx context.Context, body EnvironmentUpdateParams, opts ...option.RequestOption) (res *EnvironmentUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/UpdateEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListEnvironments returns a list of environments that match the query.
func (r *EnvironmentService) List(ctx context.Context, params EnvironmentListParams, opts ...option.RequestOption) (res *pagination.EnvironmentsPage[EnvironmentListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.EnvironmentService/ListEnvironments"
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

// ListEnvironments returns a list of environments that match the query.
func (r *EnvironmentService) ListAutoPaging(ctx context.Context, params EnvironmentListParams, opts ...option.RequestOption) *pagination.EnvironmentsPageAutoPager[EnvironmentListResponse] {
	return pagination.NewEnvironmentsPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteEnvironment deletes an environment. When the environment is running, it
// will be stopped as well. Deleted environments cannot be started again.
func (r *EnvironmentService) Delete(ctx context.Context, body EnvironmentDeleteParams, opts ...option.RequestOption) (res *EnvironmentDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/DeleteEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// CreateAbdStartEnvironmentFromProject creates a new environment from a project
// and starts it.
func (r *EnvironmentService) NewFromProject(ctx context.Context, body EnvironmentNewFromProjectParams, opts ...option.RequestOption) (res *EnvironmentNewFromProjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/CreateEnvironmentFromProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// CreateEnvironmentLogsToken creates a token that can be used to access the logs
// of an environment.
func (r *EnvironmentService) NewLogsToken(ctx context.Context, body EnvironmentNewLogsTokenParams, opts ...option.RequestOption) (res *EnvironmentNewLogsTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/CreateEnvironmentLogsToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// MarkEnvironmentActive allows tools to signal activity for an environment.
func (r *EnvironmentService) MarkActive(ctx context.Context, body EnvironmentMarkActiveParams, opts ...option.RequestOption) (res *EnvironmentMarkActiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/MarkEnvironmentActive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// StartEnvironment starts an environment. This function is idempotent, i.e. if the
// environment is already running no error is returned.
func (r *EnvironmentService) Start(ctx context.Context, body EnvironmentStartParams, opts ...option.RequestOption) (res *EnvironmentStartResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/StartEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// StopEnvironment stops a running environment.
func (r *EnvironmentService) Stop(ctx context.Context, body EnvironmentStopParams, opts ...option.RequestOption) (res *EnvironmentStopResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/StopEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EnvironmentNewResponse struct {
	// +resource get environment
	Environment EnvironmentNewResponseEnvironment `json:"environment"`
	JSON        environmentNewResponseJSON        `json:"-"`
}

// environmentNewResponseJSON contains the JSON metadata for the struct
// [EnvironmentNewResponse]
type environmentNewResponseJSON struct {
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseJSON) RawJSON() string {
	return r.raw
}

// +resource get environment
type EnvironmentNewResponseEnvironment struct {
	// ID is a unique identifier of this environment. No other environment with the
	// same name must be managed by this environment manager
	ID string `json:"id"`
	// EnvironmentMetadata is data associated with an environment that's required for
	// other parts of the system to function
	Metadata EnvironmentNewResponseEnvironmentMetadata `json:"metadata"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec EnvironmentNewResponseEnvironmentSpec `json:"spec"`
	// EnvironmentStatus describes an environment status
	Status EnvironmentNewResponseEnvironmentStatus `json:"status"`
	JSON   environmentNewResponseEnvironmentJSON   `json:"-"`
}

// environmentNewResponseEnvironmentJSON contains the JSON metadata for the struct
// [EnvironmentNewResponseEnvironment]
type environmentNewResponseEnvironmentJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentMetadata is data associated with an environment that's required for
// other parts of the system to function
type EnvironmentNewResponseEnvironmentMetadata struct {
	// annotations are key/value pairs that gets attached to the environment.
	// +internal - not yet implemented
	Annotations map[string]string `json:"annotations"`
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
	// creator is the identity of the creator of the environment
	Creator EnvironmentNewResponseEnvironmentMetadataCreator `json:"creator"`
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
	LastStartedAt time.Time `json:"lastStartedAt" format:"date-time"`
	// name is the name of the environment as specified by the user
	Name string `json:"name"`
	// organization_id is the ID of the organization that contains the environment
	OrganizationID string `json:"organizationId" format:"uuid"`
	// original_context_url is the normalized URL from which the environment was
	// created
	OriginalContextURL string `json:"originalContextUrl"`
	// If the Environment was started from a project, the project_id will reference the
	// project.
	ProjectID string `json:"projectId"`
	// Runner is the ID of the runner that runs this environment.
	RunnerID string                                        `json:"runnerId"`
	JSON     environmentNewResponseEnvironmentMetadataJSON `json:"-"`
}

// environmentNewResponseEnvironmentMetadataJSON contains the JSON metadata for the
// struct [EnvironmentNewResponseEnvironmentMetadata]
type environmentNewResponseEnvironmentMetadataJSON struct {
	Annotations        apijson.Field
	CreatedAt          apijson.Field
	Creator            apijson.Field
	LastStartedAt      apijson.Field
	Name               apijson.Field
	OrganizationID     apijson.Field
	OriginalContextURL apijson.Field
	ProjectID          apijson.Field
	RunnerID           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the creator of the environment
type EnvironmentNewResponseEnvironmentMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentNewResponseEnvironmentMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentNewResponseEnvironmentMetadataCreatorJSON      `json:"-"`
}

// environmentNewResponseEnvironmentMetadataCreatorJSON contains the JSON metadata
// for the struct [EnvironmentNewResponseEnvironmentMetadataCreator]
type environmentNewResponseEnvironmentMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentNewResponseEnvironmentMetadataCreatorPrincipal string

const (
	EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentNewResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalAccount        EnvironmentNewResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalUser           EnvironmentNewResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalRunner         EnvironmentNewResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentNewResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentNewResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentNewResponseEnvironmentMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalAccount, EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalUser, EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalRunner, EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentNewResponseEnvironmentMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type EnvironmentNewResponseEnvironmentSpec struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission EnvironmentNewResponseEnvironmentSpecAdmission `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile EnvironmentNewResponseEnvironmentSpecAutomationsFile `json:"automationsFile"`
	// content is the content spec of the environment
	Content EnvironmentNewResponseEnvironmentSpecContent `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase EnvironmentNewResponseEnvironmentSpecDesiredPhase `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer EnvironmentNewResponseEnvironmentSpecDevcontainer `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine EnvironmentNewResponseEnvironmentSpecMachine `json:"machine"`
	// ports is the set of ports which ought to be exposed to the internet
	Ports []EnvironmentNewResponseEnvironmentSpecPort `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets []EnvironmentNewResponseEnvironmentSpecSecret `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion string `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys []EnvironmentNewResponseEnvironmentSpecSSHPublicKey `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout EnvironmentNewResponseEnvironmentSpecTimeout `json:"timeout"`
	JSON    environmentNewResponseEnvironmentSpecJSON    `json:"-"`
}

// environmentNewResponseEnvironmentSpecJSON contains the JSON metadata for the
// struct [EnvironmentNewResponseEnvironmentSpec]
type environmentNewResponseEnvironmentSpecJSON struct {
	Admission       apijson.Field
	AutomationsFile apijson.Field
	Content         apijson.Field
	DesiredPhase    apijson.Field
	Devcontainer    apijson.Field
	Machine         apijson.Field
	Ports           apijson.Field
	Secrets         apijson.Field
	SpecVersion     apijson.Field
	SSHPublicKeys   apijson.Field
	Timeout         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentNewResponseEnvironmentSpecAdmission string

const (
	EnvironmentNewResponseEnvironmentSpecAdmissionAdmissionLevelUnspecified EnvironmentNewResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentSpecAdmissionAdmissionLevelOwnerOnly   EnvironmentNewResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentNewResponseEnvironmentSpecAdmissionAdmissionLevelEveryone    EnvironmentNewResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentNewResponseEnvironmentSpecAdmission) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentSpecAdmissionAdmissionLevelUnspecified, EnvironmentNewResponseEnvironmentSpecAdmissionAdmissionLevelOwnerOnly, EnvironmentNewResponseEnvironmentSpecAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// automations_file is the automations file spec of the environment
type EnvironmentNewResponseEnvironmentSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath string                                                   `json:"automationsFilePath"`
	Session             string                                                   `json:"session"`
	JSON                environmentNewResponseEnvironmentSpecAutomationsFileJSON `json:"-"`
}

// environmentNewResponseEnvironmentSpecAutomationsFileJSON contains the JSON
// metadata for the struct [EnvironmentNewResponseEnvironmentSpecAutomationsFile]
type environmentNewResponseEnvironmentSpecAutomationsFileJSON struct {
	AutomationsFilePath apijson.Field
	Session             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// content is the content spec of the environment
type EnvironmentNewResponseEnvironmentSpecContent struct {
	// The Git email address
	GitEmail string `json:"gitEmail"`
	// The Git username
	GitUsername string `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer EnvironmentNewResponseEnvironmentSpecContentInitializer `json:"initializer"`
	Session     string                                                  `json:"session"`
	JSON        environmentNewResponseEnvironmentSpecContentJSON        `json:"-"`
}

// environmentNewResponseEnvironmentSpecContentJSON contains the JSON metadata for
// the struct [EnvironmentNewResponseEnvironmentSpecContent]
type environmentNewResponseEnvironmentSpecContentJSON struct {
	GitEmail    apijson.Field
	GitUsername apijson.Field
	Initializer apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecContentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentInitializer specifies how an environment is to be initialized
type EnvironmentNewResponseEnvironmentSpecContentInitializer struct {
	Specs []EnvironmentNewResponseEnvironmentSpecContentInitializerSpec `json:"specs"`
	JSON  environmentNewResponseEnvironmentSpecContentInitializerJSON   `json:"-"`
}

// environmentNewResponseEnvironmentSpecContentInitializerJSON contains the JSON
// metadata for the struct
// [EnvironmentNewResponseEnvironmentSpecContentInitializer]
type environmentNewResponseEnvironmentSpecContentInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecContentInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecContentInitializerJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewResponseEnvironmentSpecContentInitializerSpec struct {
	// This field can have the runtime type of
	// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGit].
	Git   interface{}                                                     `json:"git"`
	JSON  environmentNewResponseEnvironmentSpecContentInitializerSpecJSON `json:"-"`
	union EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnion
}

// environmentNewResponseEnvironmentSpecContentInitializerSpecJSON contains the
// JSON metadata for the struct
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpec]
type environmentNewResponseEnvironmentSpecContentInitializerSpecJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentNewResponseEnvironmentSpecContentInitializerSpecJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentNewResponseEnvironmentSpecContentInitializerSpec) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentNewResponseEnvironmentSpecContentInitializerSpec{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURL],
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGit].
func (r EnvironmentNewResponseEnvironmentSpecContentInitializerSpec) AsUnion() EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnion {
	return r.union
}

// Union satisfied by
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURL] or
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGit].
type EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnion interface {
	implementsEnvironmentNewResponseEnvironmentSpecContentInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGit{}),
		},
	)
}

type EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURL struct {
	ContextURL EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL `json:"contextUrl,required"`
	JSON       environmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLJSON       `json:"-"`
}

// environmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLJSON
// contains the JSON metadata for the struct
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURL]
type environmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLJSON struct {
	ContextURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURL) implementsEnvironmentNewResponseEnvironmentSpecContentInitializerSpec() {
}

type EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                                               `json:"url" format:"uri"`
	JSON environmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON `json:"-"`
}

// environmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON
// contains the JSON metadata for the struct
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL]
type environmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGit struct {
	Git  EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGit  `json:"git,required"`
	JSON environmentNewResponseEnvironmentSpecContentInitializerSpecsGitJSON `json:"-"`
}

// environmentNewResponseEnvironmentSpecContentInitializerSpecsGitJSON contains the
// JSON metadata for the struct
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGit]
type environmentNewResponseEnvironmentSpecContentInitializerSpecsGitJSON struct {
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecContentInitializerSpecsGitJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGit) implementsEnvironmentNewResponseEnvironmentSpecContentInitializerSpec() {
}

type EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                                                 `json:"upstreamRemoteUri"`
	JSON              environmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitJSON `json:"-"`
}

// environmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitJSON contains
// the JSON metadata for the struct
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGit]
type environmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode string

const (
	EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

// Phase is the desired phase of the environment
type EnvironmentNewResponseEnvironmentSpecDesiredPhase string

const (
	EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUnspecified EnvironmentNewResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseCreating    EnvironmentNewResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStarting    EnvironmentNewResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseRunning     EnvironmentNewResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUpdating    EnvironmentNewResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopping    EnvironmentNewResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopped     EnvironmentNewResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleting    EnvironmentNewResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleted     EnvironmentNewResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentNewResponseEnvironmentSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUnspecified, EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseCreating, EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStarting, EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseRunning, EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUpdating, EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopping, EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopped, EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleting, EnvironmentNewResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// devcontainer is the devcontainer spec of the environment
type EnvironmentNewResponseEnvironmentSpecDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath string                                                `json:"devcontainerFilePath"`
	Session              string                                                `json:"session"`
	JSON                 environmentNewResponseEnvironmentSpecDevcontainerJSON `json:"-"`
}

// environmentNewResponseEnvironmentSpecDevcontainerJSON contains the JSON metadata
// for the struct [EnvironmentNewResponseEnvironmentSpecDevcontainer]
type environmentNewResponseEnvironmentSpecDevcontainerJSON struct {
	DevcontainerFilePath apijson.Field
	Session              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecDevcontainerJSON) RawJSON() string {
	return r.raw
}

// machine is the machine spec of the environment
type EnvironmentNewResponseEnvironmentSpecMachine struct {
	// Class denotes the class of the environment we ought to start
	Class   string                                           `json:"class" format:"uuid"`
	Session string                                           `json:"session"`
	JSON    environmentNewResponseEnvironmentSpecMachineJSON `json:"-"`
}

// environmentNewResponseEnvironmentSpecMachineJSON contains the JSON metadata for
// the struct [EnvironmentNewResponseEnvironmentSpecMachine]
type environmentNewResponseEnvironmentSpecMachineJSON struct {
	Class       apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecMachineJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewResponseEnvironmentSpecPort struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission EnvironmentNewResponseEnvironmentSpecPortsAdmission `json:"admission"`
	// name of this port
	Name string `json:"name"`
	// port number
	Port int64                                         `json:"port"`
	JSON environmentNewResponseEnvironmentSpecPortJSON `json:"-"`
}

// environmentNewResponseEnvironmentSpecPortJSON contains the JSON metadata for the
// struct [EnvironmentNewResponseEnvironmentSpecPort]
type environmentNewResponseEnvironmentSpecPortJSON struct {
	Admission   apijson.Field
	Name        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecPortJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentNewResponseEnvironmentSpecPortsAdmission string

const (
	EnvironmentNewResponseEnvironmentSpecPortsAdmissionAdmissionLevelUnspecified EnvironmentNewResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentSpecPortsAdmissionAdmissionLevelOwnerOnly   EnvironmentNewResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentNewResponseEnvironmentSpecPortsAdmissionAdmissionLevelEveryone    EnvironmentNewResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentNewResponseEnvironmentSpecPortsAdmission) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentSpecPortsAdmissionAdmissionLevelUnspecified, EnvironmentNewResponseEnvironmentSpecPortsAdmissionAdmissionLevelOwnerOnly, EnvironmentNewResponseEnvironmentSpecPortsAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

type EnvironmentNewResponseEnvironmentSpecSecret struct {
	EnvironmentVariable string `json:"environmentVariable"`
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath          string `json:"filePath"`
	GitCredentialHost string `json:"gitCredentialHost"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                          `json:"sourceRef"`
	JSON      environmentNewResponseEnvironmentSpecSecretJSON `json:"-"`
	union     EnvironmentNewResponseEnvironmentSpecSecretsUnion
}

// environmentNewResponseEnvironmentSpecSecretJSON contains the JSON metadata for
// the struct [EnvironmentNewResponseEnvironmentSpecSecret]
type environmentNewResponseEnvironmentSpecSecretJSON struct {
	EnvironmentVariable apijson.Field
	FilePath            apijson.Field
	GitCredentialHost   apijson.Field
	Name                apijson.Field
	Session             apijson.Field
	Source              apijson.Field
	SourceRef           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r environmentNewResponseEnvironmentSpecSecretJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentNewResponseEnvironmentSpecSecret) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentNewResponseEnvironmentSpecSecret{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EnvironmentNewResponseEnvironmentSpecSecretsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentNewResponseEnvironmentSpecSecretsObject],
// [EnvironmentNewResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted],
// [EnvironmentNewResponseEnvironmentSpecSecretsObject].
func (r EnvironmentNewResponseEnvironmentSpecSecret) AsUnion() EnvironmentNewResponseEnvironmentSpecSecretsUnion {
	return r.union
}

// Union satisfied by [EnvironmentNewResponseEnvironmentSpecSecretsObject],
// [EnvironmentNewResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted]
// or [EnvironmentNewResponseEnvironmentSpecSecretsObject].
type EnvironmentNewResponseEnvironmentSpecSecretsUnion interface {
	implementsEnvironmentNewResponseEnvironmentSpecSecret()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentNewResponseEnvironmentSpecSecretsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentNewResponseEnvironmentSpecSecretsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentNewResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentNewResponseEnvironmentSpecSecretsObject{}),
		},
	)
}

type EnvironmentNewResponseEnvironmentSpecSecretsObject struct {
	EnvironmentVariable string `json:"environmentVariable,required"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                                 `json:"sourceRef"`
	JSON      environmentNewResponseEnvironmentSpecSecretsObjectJSON `json:"-"`
}

// environmentNewResponseEnvironmentSpecSecretsObjectJSON contains the JSON
// metadata for the struct [EnvironmentNewResponseEnvironmentSpecSecretsObject]
type environmentNewResponseEnvironmentSpecSecretsObjectJSON struct {
	EnvironmentVariable apijson.Field
	Name                apijson.Field
	Session             apijson.Field
	Source              apijson.Field
	SourceRef           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecSecretsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecSecretsObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentNewResponseEnvironmentSpecSecretsObject) implementsEnvironmentNewResponseEnvironmentSpecSecret() {
}

type EnvironmentNewResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted struct {
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath string `json:"filePath,required"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                                                                                        `json:"sourceRef"`
	JSON      environmentNewResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON `json:"-"`
}

// environmentNewResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON
// contains the JSON metadata for the struct
// [EnvironmentNewResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted]
type environmentNewResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON struct {
	FilePath    apijson.Field
	Name        apijson.Field
	Session     apijson.Field
	Source      apijson.Field
	SourceRef   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentNewResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) implementsEnvironmentNewResponseEnvironmentSpecSecret() {
}

type EnvironmentNewResponseEnvironmentSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// value is the actual public key in the public key file format
	Value string                                                `json:"value"`
	JSON  environmentNewResponseEnvironmentSpecSSHPublicKeyJSON `json:"-"`
}

// environmentNewResponseEnvironmentSpecSSHPublicKeyJSON contains the JSON metadata
// for the struct [EnvironmentNewResponseEnvironmentSpecSSHPublicKey]
type environmentNewResponseEnvironmentSpecSSHPublicKeyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// Timeout configures the environment timeout
type EnvironmentNewResponseEnvironmentSpecTimeout struct {
	// A Duration represents a signed, fixed-length span of time represented as a count
	// of seconds and fractions of seconds at nanosecond resolution. It is independent
	// of any calendar and concepts like "day" or "month". It is related to Timestamp
	// in that the difference between two Timestamp values is a Duration and it can be
	// added or subtracted from a Timestamp. Range is approximately +-10,000 years.
	//
	// # Examples
	//
	// Example 1: Compute Duration from two Timestamps in pseudo code.
	//
	//	Timestamp start = ...;
	//	Timestamp end = ...;
	//	Duration duration = ...;
	//
	//	duration.seconds = end.seconds - start.seconds;
	//	duration.nanos = end.nanos - start.nanos;
	//
	//	if (duration.seconds < 0 && duration.nanos > 0) {
	//	  duration.seconds += 1;
	//	  duration.nanos -= 1000000000;
	//	} else if (duration.seconds > 0 && duration.nanos < 0) {
	//	  duration.seconds -= 1;
	//	  duration.nanos += 1000000000;
	//	}
	//
	// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
	//
	//	Timestamp start = ...;
	//	Duration duration = ...;
	//	Timestamp end = ...;
	//
	//	end.seconds = start.seconds + duration.seconds;
	//	end.nanos = start.nanos + duration.nanos;
	//
	//	if (end.nanos < 0) {
	//	  end.seconds -= 1;
	//	  end.nanos += 1000000000;
	//	} else if (end.nanos >= 1000000000) {
	//	  end.seconds += 1;
	//	  end.nanos -= 1000000000;
	//	}
	//
	// Example 3: Compute Duration from datetime.timedelta in Python.
	//
	//	td = datetime.timedelta(days=3, minutes=10)
	//	duration = Duration()
	//	duration.FromTimedelta(td)
	//
	// # JSON Mapping
	//
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by
	// the number of seconds, with nanoseconds expressed as fractional seconds. For
	// example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as
	// "3.000000001s", and 3 seconds and 1 microsecond should be expressed in JSON
	// format as "3.000001s".
	Disconnected string                                           `json:"disconnected" format:"regex"`
	JSON         environmentNewResponseEnvironmentSpecTimeoutJSON `json:"-"`
}

// environmentNewResponseEnvironmentSpecTimeoutJSON contains the JSON metadata for
// the struct [EnvironmentNewResponseEnvironmentSpecTimeout]
type environmentNewResponseEnvironmentSpecTimeoutJSON struct {
	Disconnected apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentSpecTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentSpecTimeoutJSON) RawJSON() string {
	return r.raw
}

// EnvironmentStatus describes an environment status
type EnvironmentNewResponseEnvironmentStatus struct {
	// EnvironmentActivitySignal used to signal activity for an environment.
	ActivitySignal EnvironmentNewResponseEnvironmentStatusActivitySignal `json:"activitySignal"`
	// automations_file contains the status of the automations file.
	AutomationsFile EnvironmentNewResponseEnvironmentStatusAutomationsFile `json:"automationsFile"`
	// content contains the status of the environment content.
	Content EnvironmentNewResponseEnvironmentStatusContent `json:"content"`
	// devcontainer contains the status of the devcontainer.
	Devcontainer EnvironmentNewResponseEnvironmentStatusDevcontainer `json:"devcontainer"`
	// environment_url contains the URL at which the environment can be accessed. This
	// field is only set if the environment is running.
	EnvironmentURLs EnvironmentNewResponseEnvironmentStatusEnvironmentURLs `json:"environmentUrls"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage []string `json:"failureMessage"`
	// machine contains the status of the environment machine
	Machine EnvironmentNewResponseEnvironmentStatusMachine `json:"machine"`
	// the phase of an environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentNewResponseEnvironmentStatusPhase `json:"phase"`
	// RunnerACK is the acknowledgement from the runner that is has received the
	// environment spec.
	RunnerAck EnvironmentNewResponseEnvironmentStatusRunnerAck `json:"runnerAck"`
	// secrets contains the status of the environment secrets
	Secrets []EnvironmentNewResponseEnvironmentStatusSecret `json:"secrets"`
	// ssh_public_keys contains the status of the environment ssh public keys
	SSHPublicKeys []EnvironmentNewResponseEnvironmentStatusSSHPublicKey `json:"sshPublicKeys"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion"`
	// warning_message contains warnings, e.g. when the environment is present but not
	// in the expected state.
	WarningMessage []string                                    `json:"warningMessage"`
	JSON           environmentNewResponseEnvironmentStatusJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusJSON contains the JSON metadata for the
// struct [EnvironmentNewResponseEnvironmentStatus]
type environmentNewResponseEnvironmentStatusJSON struct {
	ActivitySignal  apijson.Field
	AutomationsFile apijson.Field
	Content         apijson.Field
	Devcontainer    apijson.Field
	EnvironmentURLs apijson.Field
	FailureMessage  apijson.Field
	Machine         apijson.Field
	Phase           apijson.Field
	RunnerAck       apijson.Field
	Secrets         apijson.Field
	SSHPublicKeys   apijson.Field
	StatusVersion   apijson.Field
	WarningMessage  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusJSON) RawJSON() string {
	return r.raw
}

// EnvironmentActivitySignal used to signal activity for an environment.
type EnvironmentNewResponseEnvironmentStatusActivitySignal struct {
	// source of the activity signal, such as "VS Code", "SSH", or "Automations". It
	// should be a human-readable string that describes the source of the activity
	// signal.
	Source string `json:"source"`
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
	Timestamp time.Time                                                 `json:"timestamp" format:"date-time"`
	JSON      environmentNewResponseEnvironmentStatusActivitySignalJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusActivitySignalJSON contains the JSON
// metadata for the struct [EnvironmentNewResponseEnvironmentStatusActivitySignal]
type environmentNewResponseEnvironmentStatusActivitySignalJSON struct {
	Source      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusActivitySignal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusActivitySignalJSON) RawJSON() string {
	return r.raw
}

// automations_file contains the status of the automations file.
type EnvironmentNewResponseEnvironmentStatusAutomationsFile struct {
	// automations_file_path is the path to the automations file relative to the repo
	// root.
	AutomationsFilePath string `json:"automationsFilePath"`
	// automations_file_presence indicates how an automations file is present in the
	// environment.
	AutomationsFilePresence EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresence `json:"automationsFilePresence"`
	// failure_message contains the reason the automations file failed to be applied.
	// This is only set if the phase is FAILED.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the automations file.
	Phase EnvironmentNewResponseEnvironmentStatusAutomationsFilePhase `json:"phase"`
	// session is the automations file session that is currently applied in the
	// environment.
	Session string                                                     `json:"session"`
	JSON    environmentNewResponseEnvironmentStatusAutomationsFileJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusAutomationsFileJSON contains the JSON
// metadata for the struct [EnvironmentNewResponseEnvironmentStatusAutomationsFile]
type environmentNewResponseEnvironmentStatusAutomationsFileJSON struct {
	AutomationsFilePath     apijson.Field
	AutomationsFilePresence apijson.Field
	FailureMessage          apijson.Field
	Phase                   apijson.Field
	Session                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// automations_file_presence indicates how an automations file is present in the
// environment.
type EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresence string

const (
	EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceUnspecified EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceAbsent      EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_ABSENT"
	EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceDiscovered  EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceSpecified   EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceUnspecified, EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceAbsent, EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceDiscovered, EnvironmentNewResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the automations file.
type EnvironmentNewResponseEnvironmentStatusAutomationsFilePhase string

const (
	EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUnspecified  EnvironmentNewResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseCreating     EnvironmentNewResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_CREATING"
	EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseInitializing EnvironmentNewResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseReady        EnvironmentNewResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_READY"
	EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUpdating     EnvironmentNewResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_UPDATING"
	EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseFailed       EnvironmentNewResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentNewResponseEnvironmentStatusAutomationsFilePhase) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUnspecified, EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseCreating, EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseInitializing, EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseReady, EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUpdating, EnvironmentNewResponseEnvironmentStatusAutomationsFilePhaseContentPhaseFailed:
		return true
	}
	return false
}

// content contains the status of the environment content.
type EnvironmentNewResponseEnvironmentStatusContent struct {
	// content_location_in_machine is the location of the content in the machine
	ContentLocationInMachine string `json:"contentLocationInMachine"`
	// failure_message contains the reason the content initialization failed.
	FailureMessage string `json:"failureMessage"`
	// git is the Git working copy status of the environment. Note: this is a
	// best-effort field and more often than not will not be present. Its absence does
	// not indicate the absence of a working copy.
	Git EnvironmentNewResponseEnvironmentStatusContentGit `json:"git"`
	// phase is the current phase of the environment content
	Phase EnvironmentNewResponseEnvironmentStatusContentPhase `json:"phase"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the content is present but not in
	// the expected state.
	WarningMessage string                                             `json:"warningMessage"`
	JSON           environmentNewResponseEnvironmentStatusContentJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusContentJSON contains the JSON metadata
// for the struct [EnvironmentNewResponseEnvironmentStatusContent]
type environmentNewResponseEnvironmentStatusContentJSON struct {
	ContentLocationInMachine apijson.Field
	FailureMessage           apijson.Field
	Git                      apijson.Field
	Phase                    apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusContentJSON) RawJSON() string {
	return r.raw
}

// git is the Git working copy status of the environment. Note: this is a
// best-effort field and more often than not will not be present. Its absence does
// not indicate the absence of a working copy.
type EnvironmentNewResponseEnvironmentStatusContentGit struct {
	// branch is branch we're currently on
	Branch string `json:"branch"`
	// changed_files is an array of changed files in the environment, possibly
	// truncated
	ChangedFiles []EnvironmentNewResponseEnvironmentStatusContentGitChangedFile `json:"changedFiles"`
	// clone_url is the repository url as you would pass it to "git clone". Only HTTPS
	// clone URLs are supported.
	CloneURL string `json:"cloneUrl"`
	// latest_commit is the most recent commit on the current branch
	LatestCommit      string `json:"latestCommit"`
	TotalChangedFiles int64  `json:"totalChangedFiles"`
	// the total number of unpushed changes
	TotalUnpushedCommits int64 `json:"totalUnpushedCommits"`
	// unpushed_commits is an array of unpushed changes in the environment, possibly
	// truncated
	UnpushedCommits []string                                              `json:"unpushedCommits"`
	JSON            environmentNewResponseEnvironmentStatusContentGitJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusContentGitJSON contains the JSON metadata
// for the struct [EnvironmentNewResponseEnvironmentStatusContentGit]
type environmentNewResponseEnvironmentStatusContentGitJSON struct {
	Branch               apijson.Field
	ChangedFiles         apijson.Field
	CloneURL             apijson.Field
	LatestCommit         apijson.Field
	TotalChangedFiles    apijson.Field
	TotalUnpushedCommits apijson.Field
	UnpushedCommits      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusContentGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusContentGitJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewResponseEnvironmentStatusContentGitChangedFile struct {
	// ChangeType is the type of change that happened to the file
	ChangeType EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeType `json:"changeType"`
	// path is the path of the file
	Path string                                                           `json:"path"`
	JSON environmentNewResponseEnvironmentStatusContentGitChangedFileJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusContentGitChangedFileJSON contains the
// JSON metadata for the struct
// [EnvironmentNewResponseEnvironmentStatusContentGitChangedFile]
type environmentNewResponseEnvironmentStatusContentGitChangedFileJSON struct {
	ChangeType  apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusContentGitChangedFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusContentGitChangedFileJSON) RawJSON() string {
	return r.raw
}

// ChangeType is the type of change that happened to the file
type EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeType string

const (
	EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified        EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeAdded              EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_ADDED"
	EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeModified           EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_MODIFIED"
	EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeDeleted            EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_DELETED"
	EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeRenamed            EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_RENAMED"
	EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeCopied             EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_COPIED"
	EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UPDATED_BUT_UNMERGED"
	EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUntracked          EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNTRACKED"
)

func (r EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeType) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified, EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeAdded, EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeModified, EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeDeleted, EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeRenamed, EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeCopied, EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged, EnvironmentNewResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUntracked:
		return true
	}
	return false
}

// phase is the current phase of the environment content
type EnvironmentNewResponseEnvironmentStatusContentPhase string

const (
	EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseUnspecified  EnvironmentNewResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseCreating     EnvironmentNewResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_CREATING"
	EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseInitializing EnvironmentNewResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseReady        EnvironmentNewResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_READY"
	EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseUpdating     EnvironmentNewResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseFailed       EnvironmentNewResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentNewResponseEnvironmentStatusContentPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseUnspecified, EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseCreating, EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseInitializing, EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseReady, EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseUpdating, EnvironmentNewResponseEnvironmentStatusContentPhaseContentPhaseFailed:
		return true
	}
	return false
}

// devcontainer contains the status of the devcontainer.
type EnvironmentNewResponseEnvironmentStatusDevcontainer struct {
	// container_id is the ID of the container.
	ContainerID string `json:"containerId"`
	// container_name is the name of the container that is used to connect to the
	// devcontainer
	ContainerName string `json:"containerName"`
	// devcontainerconfig_in_sync indicates if the devcontainer is up to date w.r.t.
	// the devcontainer config file.
	DevcontainerconfigInSync bool `json:"devcontainerconfigInSync"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// devcontainer_file_presence indicates how the devcontainer file is present in the
	// repo.
	DevcontainerFilePresence EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresence `json:"devcontainerFilePresence"`
	// failure_message contains the reason the devcontainer failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the devcontainer
	Phase EnvironmentNewResponseEnvironmentStatusDevcontainerPhase `json:"phase"`
	// remote_user is the user that is used to connect to the devcontainer
	RemoteUser string `json:"remoteUser"`
	// remote_workspace_folder is the folder that is used to connect to the
	// devcontainer
	RemoteWorkspaceFolder string `json:"remoteWorkspaceFolder"`
	// secrets_in_sync indicates if the secrets are up to date w.r.t. the running
	// devcontainer.
	SecretsInSync bool `json:"secretsInSync"`
	// session is the session that is currently active in the devcontainer.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the devcontainer is present but not
	// in the expected state.
	WarningMessage string                                                  `json:"warningMessage"`
	JSON           environmentNewResponseEnvironmentStatusDevcontainerJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusDevcontainerJSON contains the JSON
// metadata for the struct [EnvironmentNewResponseEnvironmentStatusDevcontainer]
type environmentNewResponseEnvironmentStatusDevcontainerJSON struct {
	ContainerID              apijson.Field
	ContainerName            apijson.Field
	DevcontainerconfigInSync apijson.Field
	DevcontainerFilePath     apijson.Field
	DevcontainerFilePresence apijson.Field
	FailureMessage           apijson.Field
	Phase                    apijson.Field
	RemoteUser               apijson.Field
	RemoteWorkspaceFolder    apijson.Field
	SecretsInSync            apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusDevcontainerJSON) RawJSON() string {
	return r.raw
}

// devcontainer_file_presence indicates how the devcontainer file is present in the
// repo.
type EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresence string

const (
	EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceUnspecified EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceGenerated   EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_GENERATED"
	EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceDiscovered  EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceSpecified   EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceUnspecified, EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceGenerated, EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceDiscovered, EnvironmentNewResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the devcontainer
type EnvironmentNewResponseEnvironmentStatusDevcontainerPhase string

const (
	EnvironmentNewResponseEnvironmentStatusDevcontainerPhasePhaseUnspecified EnvironmentNewResponseEnvironmentStatusDevcontainerPhase = "PHASE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentStatusDevcontainerPhasePhaseCreating    EnvironmentNewResponseEnvironmentStatusDevcontainerPhase = "PHASE_CREATING"
	EnvironmentNewResponseEnvironmentStatusDevcontainerPhasePhaseRunning     EnvironmentNewResponseEnvironmentStatusDevcontainerPhase = "PHASE_RUNNING"
	EnvironmentNewResponseEnvironmentStatusDevcontainerPhasePhaseStopped     EnvironmentNewResponseEnvironmentStatusDevcontainerPhase = "PHASE_STOPPED"
	EnvironmentNewResponseEnvironmentStatusDevcontainerPhasePhaseFailed      EnvironmentNewResponseEnvironmentStatusDevcontainerPhase = "PHASE_FAILED"
)

func (r EnvironmentNewResponseEnvironmentStatusDevcontainerPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentStatusDevcontainerPhasePhaseUnspecified, EnvironmentNewResponseEnvironmentStatusDevcontainerPhasePhaseCreating, EnvironmentNewResponseEnvironmentStatusDevcontainerPhasePhaseRunning, EnvironmentNewResponseEnvironmentStatusDevcontainerPhasePhaseStopped, EnvironmentNewResponseEnvironmentStatusDevcontainerPhasePhaseFailed:
		return true
	}
	return false
}

// environment_url contains the URL at which the environment can be accessed. This
// field is only set if the environment is running.
type EnvironmentNewResponseEnvironmentStatusEnvironmentURLs struct {
	// logs is the URL at which the environment logs can be accessed.
	Logs  string                                                       `json:"logs"`
	Ports []EnvironmentNewResponseEnvironmentStatusEnvironmentURLsPort `json:"ports"`
	// SSH is the URL at which the environment can be accessed via SSH.
	SSH  EnvironmentNewResponseEnvironmentStatusEnvironmentURLsSSH  `json:"ssh"`
	JSON environmentNewResponseEnvironmentStatusEnvironmentURLsJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusEnvironmentURLsJSON contains the JSON
// metadata for the struct [EnvironmentNewResponseEnvironmentStatusEnvironmentURLs]
type environmentNewResponseEnvironmentStatusEnvironmentURLsJSON struct {
	Logs        apijson.Field
	Ports       apijson.Field
	SSH         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusEnvironmentURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusEnvironmentURLsJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewResponseEnvironmentStatusEnvironmentURLsPort struct {
	// port is the port number of the environment port
	Port int64 `json:"port"`
	// url is the URL at which the environment port can be accessed
	URL  string                                                         `json:"url"`
	JSON environmentNewResponseEnvironmentStatusEnvironmentURLsPortJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusEnvironmentURLsPortJSON contains the JSON
// metadata for the struct
// [EnvironmentNewResponseEnvironmentStatusEnvironmentURLsPort]
type environmentNewResponseEnvironmentStatusEnvironmentURLsPortJSON struct {
	Port        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusEnvironmentURLsPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusEnvironmentURLsPortJSON) RawJSON() string {
	return r.raw
}

// SSH is the URL at which the environment can be accessed via SSH.
type EnvironmentNewResponseEnvironmentStatusEnvironmentURLsSSH struct {
	URL  string                                                        `json:"url"`
	JSON environmentNewResponseEnvironmentStatusEnvironmentURLsSSHJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusEnvironmentURLsSSHJSON contains the JSON
// metadata for the struct
// [EnvironmentNewResponseEnvironmentStatusEnvironmentURLsSSH]
type environmentNewResponseEnvironmentStatusEnvironmentURLsSSHJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusEnvironmentURLsSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusEnvironmentURLsSSHJSON) RawJSON() string {
	return r.raw
}

// machine contains the status of the environment machine
type EnvironmentNewResponseEnvironmentStatusMachine struct {
	// failure_message contains the reason the machine failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the environment machine
	Phase EnvironmentNewResponseEnvironmentStatusMachinePhase `json:"phase"`
	// session is the session that is currently active in the machine.
	Session string `json:"session"`
	// timeout contains the reason the environment has timed out. If this field is
	// empty, the environment has not timed out.
	Timeout string `json:"timeout"`
	// versions contains the versions of components in the machine.
	Versions EnvironmentNewResponseEnvironmentStatusMachineVersions `json:"versions"`
	// warning_message contains warnings, e.g. when the machine is present but not in
	// the expected state.
	WarningMessage string                                             `json:"warningMessage"`
	JSON           environmentNewResponseEnvironmentStatusMachineJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusMachineJSON contains the JSON metadata
// for the struct [EnvironmentNewResponseEnvironmentStatusMachine]
type environmentNewResponseEnvironmentStatusMachineJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	Timeout        apijson.Field
	Versions       apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusMachineJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the environment machine
type EnvironmentNewResponseEnvironmentStatusMachinePhase string

const (
	EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseUnspecified EnvironmentNewResponseEnvironmentStatusMachinePhase = "PHASE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseCreating    EnvironmentNewResponseEnvironmentStatusMachinePhase = "PHASE_CREATING"
	EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseStarting    EnvironmentNewResponseEnvironmentStatusMachinePhase = "PHASE_STARTING"
	EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseRunning     EnvironmentNewResponseEnvironmentStatusMachinePhase = "PHASE_RUNNING"
	EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseStopping    EnvironmentNewResponseEnvironmentStatusMachinePhase = "PHASE_STOPPING"
	EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseStopped     EnvironmentNewResponseEnvironmentStatusMachinePhase = "PHASE_STOPPED"
	EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseDeleting    EnvironmentNewResponseEnvironmentStatusMachinePhase = "PHASE_DELETING"
	EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseDeleted     EnvironmentNewResponseEnvironmentStatusMachinePhase = "PHASE_DELETED"
)

func (r EnvironmentNewResponseEnvironmentStatusMachinePhase) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseUnspecified, EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseCreating, EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseStarting, EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseRunning, EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseStopping, EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseStopped, EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseDeleting, EnvironmentNewResponseEnvironmentStatusMachinePhasePhaseDeleted:
		return true
	}
	return false
}

// versions contains the versions of components in the machine.
type EnvironmentNewResponseEnvironmentStatusMachineVersions struct {
	SupervisorCommit  string                                                     `json:"supervisorCommit"`
	SupervisorVersion string                                                     `json:"supervisorVersion"`
	JSON              environmentNewResponseEnvironmentStatusMachineVersionsJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusMachineVersionsJSON contains the JSON
// metadata for the struct [EnvironmentNewResponseEnvironmentStatusMachineVersions]
type environmentNewResponseEnvironmentStatusMachineVersionsJSON struct {
	SupervisorCommit  apijson.Field
	SupervisorVersion apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusMachineVersions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusMachineVersionsJSON) RawJSON() string {
	return r.raw
}

// the phase of an environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentNewResponseEnvironmentStatusPhase string

const (
	EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseUnspecified EnvironmentNewResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseCreating    EnvironmentNewResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseStarting    EnvironmentNewResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseRunning     EnvironmentNewResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseUpdating    EnvironmentNewResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseStopping    EnvironmentNewResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseStopped     EnvironmentNewResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseDeleting    EnvironmentNewResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseDeleted     EnvironmentNewResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentNewResponseEnvironmentStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseUnspecified, EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseCreating, EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseStarting, EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseRunning, EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseUpdating, EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseStopping, EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseStopped, EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseDeleting, EnvironmentNewResponseEnvironmentStatusPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// RunnerACK is the acknowledgement from the runner that is has received the
// environment spec.
type EnvironmentNewResponseEnvironmentStatusRunnerAck struct {
	Message     string                                                     `json:"message"`
	SpecVersion string                                                     `json:"specVersion"`
	StatusCode  EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCode `json:"statusCode"`
	JSON        environmentNewResponseEnvironmentStatusRunnerAckJSON       `json:"-"`
}

// environmentNewResponseEnvironmentStatusRunnerAckJSON contains the JSON metadata
// for the struct [EnvironmentNewResponseEnvironmentStatusRunnerAck]
type environmentNewResponseEnvironmentStatusRunnerAckJSON struct {
	Message     apijson.Field
	SpecVersion apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusRunnerAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusRunnerAckJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCode string

const (
	EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeUnspecified        EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeOk                 EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_OK"
	EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeInvalidResource    EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_INVALID_RESOURCE"
	EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeFailedPrecondition EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_FAILED_PRECONDITION"
)

func (r EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCode) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeUnspecified, EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeOk, EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeInvalidResource, EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeFailedPrecondition:
		return true
	}
	return false
}

type EnvironmentNewResponseEnvironmentStatusSecret struct {
	// failure_message contains the reason the secret failed to be materialize.
	FailureMessage string                                              `json:"failureMessage"`
	Phase          EnvironmentNewResponseEnvironmentStatusSecretsPhase `json:"phase"`
	SecretName     string                                              `json:"secretName"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the secret is present but not in
	// the expected state.
	WarningMessage string                                            `json:"warningMessage"`
	JSON           environmentNewResponseEnvironmentStatusSecretJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusSecretJSON contains the JSON metadata for
// the struct [EnvironmentNewResponseEnvironmentStatusSecret]
type environmentNewResponseEnvironmentStatusSecretJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	SecretName     apijson.Field
	Session        apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusSecretJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewResponseEnvironmentStatusSecretsPhase string

const (
	EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseUnspecified  EnvironmentNewResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseCreating     EnvironmentNewResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_CREATING"
	EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseInitializing EnvironmentNewResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseReady        EnvironmentNewResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_READY"
	EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseUpdating     EnvironmentNewResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseFailed       EnvironmentNewResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentNewResponseEnvironmentStatusSecretsPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseUnspecified, EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseCreating, EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseInitializing, EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseReady, EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseUpdating, EnvironmentNewResponseEnvironmentStatusSecretsPhaseContentPhaseFailed:
		return true
	}
	return false
}

type EnvironmentNewResponseEnvironmentStatusSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// phase is the current phase of the public key
	Phase EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhase `json:"phase"`
	JSON  environmentNewResponseEnvironmentStatusSSHPublicKeyJSON   `json:"-"`
}

// environmentNewResponseEnvironmentStatusSSHPublicKeyJSON contains the JSON
// metadata for the struct [EnvironmentNewResponseEnvironmentStatusSSHPublicKey]
type environmentNewResponseEnvironmentStatusSSHPublicKeyJSON struct {
	ID          apijson.Field
	Phase       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironmentStatusSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentStatusSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the public key
type EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhase string

const (
	EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUnspecified  EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseCreating     EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_CREATING"
	EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseInitializing EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseReady        EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_READY"
	EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUpdating     EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseFailed       EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUnspecified, EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseCreating, EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseInitializing, EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseReady, EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUpdating, EnvironmentNewResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseFailed:
		return true
	}
	return false
}

type EnvironmentGetResponse struct {
	// +resource get environment
	Environment EnvironmentGetResponseEnvironment `json:"environment"`
	JSON        environmentGetResponseJSON        `json:"-"`
}

// environmentGetResponseJSON contains the JSON metadata for the struct
// [EnvironmentGetResponse]
type environmentGetResponseJSON struct {
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseJSON) RawJSON() string {
	return r.raw
}

// +resource get environment
type EnvironmentGetResponseEnvironment struct {
	// ID is a unique identifier of this environment. No other environment with the
	// same name must be managed by this environment manager
	ID string `json:"id"`
	// EnvironmentMetadata is data associated with an environment that's required for
	// other parts of the system to function
	Metadata EnvironmentGetResponseEnvironmentMetadata `json:"metadata"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec EnvironmentGetResponseEnvironmentSpec `json:"spec"`
	// EnvironmentStatus describes an environment status
	Status EnvironmentGetResponseEnvironmentStatus `json:"status"`
	JSON   environmentGetResponseEnvironmentJSON   `json:"-"`
}

// environmentGetResponseEnvironmentJSON contains the JSON metadata for the struct
// [EnvironmentGetResponseEnvironment]
type environmentGetResponseEnvironmentJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentMetadata is data associated with an environment that's required for
// other parts of the system to function
type EnvironmentGetResponseEnvironmentMetadata struct {
	// annotations are key/value pairs that gets attached to the environment.
	// +internal - not yet implemented
	Annotations map[string]string `json:"annotations"`
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
	// creator is the identity of the creator of the environment
	Creator EnvironmentGetResponseEnvironmentMetadataCreator `json:"creator"`
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
	LastStartedAt time.Time `json:"lastStartedAt" format:"date-time"`
	// name is the name of the environment as specified by the user
	Name string `json:"name"`
	// organization_id is the ID of the organization that contains the environment
	OrganizationID string `json:"organizationId" format:"uuid"`
	// original_context_url is the normalized URL from which the environment was
	// created
	OriginalContextURL string `json:"originalContextUrl"`
	// If the Environment was started from a project, the project_id will reference the
	// project.
	ProjectID string `json:"projectId"`
	// Runner is the ID of the runner that runs this environment.
	RunnerID string                                        `json:"runnerId"`
	JSON     environmentGetResponseEnvironmentMetadataJSON `json:"-"`
}

// environmentGetResponseEnvironmentMetadataJSON contains the JSON metadata for the
// struct [EnvironmentGetResponseEnvironmentMetadata]
type environmentGetResponseEnvironmentMetadataJSON struct {
	Annotations        apijson.Field
	CreatedAt          apijson.Field
	Creator            apijson.Field
	LastStartedAt      apijson.Field
	Name               apijson.Field
	OrganizationID     apijson.Field
	OriginalContextURL apijson.Field
	ProjectID          apijson.Field
	RunnerID           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the creator of the environment
type EnvironmentGetResponseEnvironmentMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentGetResponseEnvironmentMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentGetResponseEnvironmentMetadataCreatorJSON      `json:"-"`
}

// environmentGetResponseEnvironmentMetadataCreatorJSON contains the JSON metadata
// for the struct [EnvironmentGetResponseEnvironmentMetadataCreator]
type environmentGetResponseEnvironmentMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentGetResponseEnvironmentMetadataCreatorPrincipal string

const (
	EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalAccount        EnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalUser           EnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalRunner         EnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentGetResponseEnvironmentMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalAccount, EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalUser, EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalRunner, EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type EnvironmentGetResponseEnvironmentSpec struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission EnvironmentGetResponseEnvironmentSpecAdmission `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile EnvironmentGetResponseEnvironmentSpecAutomationsFile `json:"automationsFile"`
	// content is the content spec of the environment
	Content EnvironmentGetResponseEnvironmentSpecContent `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase EnvironmentGetResponseEnvironmentSpecDesiredPhase `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer EnvironmentGetResponseEnvironmentSpecDevcontainer `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine EnvironmentGetResponseEnvironmentSpecMachine `json:"machine"`
	// ports is the set of ports which ought to be exposed to the internet
	Ports []EnvironmentGetResponseEnvironmentSpecPort `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets []EnvironmentGetResponseEnvironmentSpecSecret `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion string `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys []EnvironmentGetResponseEnvironmentSpecSSHPublicKey `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout EnvironmentGetResponseEnvironmentSpecTimeout `json:"timeout"`
	JSON    environmentGetResponseEnvironmentSpecJSON    `json:"-"`
}

// environmentGetResponseEnvironmentSpecJSON contains the JSON metadata for the
// struct [EnvironmentGetResponseEnvironmentSpec]
type environmentGetResponseEnvironmentSpecJSON struct {
	Admission       apijson.Field
	AutomationsFile apijson.Field
	Content         apijson.Field
	DesiredPhase    apijson.Field
	Devcontainer    apijson.Field
	Machine         apijson.Field
	Ports           apijson.Field
	Secrets         apijson.Field
	SpecVersion     apijson.Field
	SSHPublicKeys   apijson.Field
	Timeout         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentGetResponseEnvironmentSpecAdmission string

const (
	EnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelUnspecified EnvironmentGetResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelOwnerOnly   EnvironmentGetResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelEveryone    EnvironmentGetResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentGetResponseEnvironmentSpecAdmission) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelUnspecified, EnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelOwnerOnly, EnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// automations_file is the automations file spec of the environment
type EnvironmentGetResponseEnvironmentSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath string                                                   `json:"automationsFilePath"`
	Session             string                                                   `json:"session"`
	JSON                environmentGetResponseEnvironmentSpecAutomationsFileJSON `json:"-"`
}

// environmentGetResponseEnvironmentSpecAutomationsFileJSON contains the JSON
// metadata for the struct [EnvironmentGetResponseEnvironmentSpecAutomationsFile]
type environmentGetResponseEnvironmentSpecAutomationsFileJSON struct {
	AutomationsFilePath apijson.Field
	Session             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// content is the content spec of the environment
type EnvironmentGetResponseEnvironmentSpecContent struct {
	// The Git email address
	GitEmail string `json:"gitEmail"`
	// The Git username
	GitUsername string `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer EnvironmentGetResponseEnvironmentSpecContentInitializer `json:"initializer"`
	Session     string                                                  `json:"session"`
	JSON        environmentGetResponseEnvironmentSpecContentJSON        `json:"-"`
}

// environmentGetResponseEnvironmentSpecContentJSON contains the JSON metadata for
// the struct [EnvironmentGetResponseEnvironmentSpecContent]
type environmentGetResponseEnvironmentSpecContentJSON struct {
	GitEmail    apijson.Field
	GitUsername apijson.Field
	Initializer apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecContentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentInitializer specifies how an environment is to be initialized
type EnvironmentGetResponseEnvironmentSpecContentInitializer struct {
	Specs []EnvironmentGetResponseEnvironmentSpecContentInitializerSpec `json:"specs"`
	JSON  environmentGetResponseEnvironmentSpecContentInitializerJSON   `json:"-"`
}

// environmentGetResponseEnvironmentSpecContentInitializerJSON contains the JSON
// metadata for the struct
// [EnvironmentGetResponseEnvironmentSpecContentInitializer]
type environmentGetResponseEnvironmentSpecContentInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecContentInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecContentInitializerJSON) RawJSON() string {
	return r.raw
}

type EnvironmentGetResponseEnvironmentSpecContentInitializerSpec struct {
	// This field can have the runtime type of
	// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGit].
	Git   interface{}                                                     `json:"git"`
	JSON  environmentGetResponseEnvironmentSpecContentInitializerSpecJSON `json:"-"`
	union EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion
}

// environmentGetResponseEnvironmentSpecContentInitializerSpecJSON contains the
// JSON metadata for the struct
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpec]
type environmentGetResponseEnvironmentSpecContentInitializerSpecJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentGetResponseEnvironmentSpecContentInitializerSpecJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentGetResponseEnvironmentSpecContentInitializerSpec) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentGetResponseEnvironmentSpecContentInitializerSpec{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURL],
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGit].
func (r EnvironmentGetResponseEnvironmentSpecContentInitializerSpec) AsUnion() EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion {
	return r.union
}

// Union satisfied by
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURL] or
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGit].
type EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion interface {
	implementsEnvironmentGetResponseEnvironmentSpecContentInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGit{}),
		},
	)
}

type EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURL struct {
	ContextURL EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL `json:"contextUrl,required"`
	JSON       environmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLJSON       `json:"-"`
}

// environmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLJSON
// contains the JSON metadata for the struct
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURL]
type environmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLJSON struct {
	ContextURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURL) implementsEnvironmentGetResponseEnvironmentSpecContentInitializerSpec() {
}

type EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                                               `json:"url" format:"uri"`
	JSON environmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON `json:"-"`
}

// environmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON
// contains the JSON metadata for the struct
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL]
type environmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON) RawJSON() string {
	return r.raw
}

type EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGit struct {
	Git  EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGit  `json:"git,required"`
	JSON environmentGetResponseEnvironmentSpecContentInitializerSpecsGitJSON `json:"-"`
}

// environmentGetResponseEnvironmentSpecContentInitializerSpecsGitJSON contains the
// JSON metadata for the struct
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGit]
type environmentGetResponseEnvironmentSpecContentInitializerSpecsGitJSON struct {
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecContentInitializerSpecsGitJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGit) implementsEnvironmentGetResponseEnvironmentSpecContentInitializerSpec() {
}

type EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                                                 `json:"upstreamRemoteUri"`
	JSON              environmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitJSON `json:"-"`
}

// environmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitJSON contains
// the JSON metadata for the struct
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGit]
type environmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode string

const (
	EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

// Phase is the desired phase of the environment
type EnvironmentGetResponseEnvironmentSpecDesiredPhase string

const (
	EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUnspecified EnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseCreating    EnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStarting    EnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseRunning     EnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUpdating    EnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopping    EnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopped     EnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleting    EnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleted     EnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentGetResponseEnvironmentSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUnspecified, EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseCreating, EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStarting, EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseRunning, EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUpdating, EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopping, EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopped, EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleting, EnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// devcontainer is the devcontainer spec of the environment
type EnvironmentGetResponseEnvironmentSpecDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath string                                                `json:"devcontainerFilePath"`
	Session              string                                                `json:"session"`
	JSON                 environmentGetResponseEnvironmentSpecDevcontainerJSON `json:"-"`
}

// environmentGetResponseEnvironmentSpecDevcontainerJSON contains the JSON metadata
// for the struct [EnvironmentGetResponseEnvironmentSpecDevcontainer]
type environmentGetResponseEnvironmentSpecDevcontainerJSON struct {
	DevcontainerFilePath apijson.Field
	Session              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecDevcontainerJSON) RawJSON() string {
	return r.raw
}

// machine is the machine spec of the environment
type EnvironmentGetResponseEnvironmentSpecMachine struct {
	// Class denotes the class of the environment we ought to start
	Class   string                                           `json:"class" format:"uuid"`
	Session string                                           `json:"session"`
	JSON    environmentGetResponseEnvironmentSpecMachineJSON `json:"-"`
}

// environmentGetResponseEnvironmentSpecMachineJSON contains the JSON metadata for
// the struct [EnvironmentGetResponseEnvironmentSpecMachine]
type environmentGetResponseEnvironmentSpecMachineJSON struct {
	Class       apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecMachineJSON) RawJSON() string {
	return r.raw
}

type EnvironmentGetResponseEnvironmentSpecPort struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission EnvironmentGetResponseEnvironmentSpecPortsAdmission `json:"admission"`
	// name of this port
	Name string `json:"name"`
	// port number
	Port int64                                         `json:"port"`
	JSON environmentGetResponseEnvironmentSpecPortJSON `json:"-"`
}

// environmentGetResponseEnvironmentSpecPortJSON contains the JSON metadata for the
// struct [EnvironmentGetResponseEnvironmentSpecPort]
type environmentGetResponseEnvironmentSpecPortJSON struct {
	Admission   apijson.Field
	Name        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecPortJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentGetResponseEnvironmentSpecPortsAdmission string

const (
	EnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelUnspecified EnvironmentGetResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelOwnerOnly   EnvironmentGetResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelEveryone    EnvironmentGetResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentGetResponseEnvironmentSpecPortsAdmission) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelUnspecified, EnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelOwnerOnly, EnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

type EnvironmentGetResponseEnvironmentSpecSecret struct {
	EnvironmentVariable string `json:"environmentVariable"`
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath          string `json:"filePath"`
	GitCredentialHost string `json:"gitCredentialHost"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                          `json:"sourceRef"`
	JSON      environmentGetResponseEnvironmentSpecSecretJSON `json:"-"`
	union     EnvironmentGetResponseEnvironmentSpecSecretsUnion
}

// environmentGetResponseEnvironmentSpecSecretJSON contains the JSON metadata for
// the struct [EnvironmentGetResponseEnvironmentSpecSecret]
type environmentGetResponseEnvironmentSpecSecretJSON struct {
	EnvironmentVariable apijson.Field
	FilePath            apijson.Field
	GitCredentialHost   apijson.Field
	Name                apijson.Field
	Session             apijson.Field
	Source              apijson.Field
	SourceRef           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r environmentGetResponseEnvironmentSpecSecretJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentGetResponseEnvironmentSpecSecret) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentGetResponseEnvironmentSpecSecret{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EnvironmentGetResponseEnvironmentSpecSecretsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentGetResponseEnvironmentSpecSecretsObject],
// [EnvironmentGetResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted],
// [EnvironmentGetResponseEnvironmentSpecSecretsObject].
func (r EnvironmentGetResponseEnvironmentSpecSecret) AsUnion() EnvironmentGetResponseEnvironmentSpecSecretsUnion {
	return r.union
}

// Union satisfied by [EnvironmentGetResponseEnvironmentSpecSecretsObject],
// [EnvironmentGetResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted]
// or [EnvironmentGetResponseEnvironmentSpecSecretsObject].
type EnvironmentGetResponseEnvironmentSpecSecretsUnion interface {
	implementsEnvironmentGetResponseEnvironmentSpecSecret()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentGetResponseEnvironmentSpecSecretsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentGetResponseEnvironmentSpecSecretsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentGetResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentGetResponseEnvironmentSpecSecretsObject{}),
		},
	)
}

type EnvironmentGetResponseEnvironmentSpecSecretsObject struct {
	EnvironmentVariable string `json:"environmentVariable,required"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                                 `json:"sourceRef"`
	JSON      environmentGetResponseEnvironmentSpecSecretsObjectJSON `json:"-"`
}

// environmentGetResponseEnvironmentSpecSecretsObjectJSON contains the JSON
// metadata for the struct [EnvironmentGetResponseEnvironmentSpecSecretsObject]
type environmentGetResponseEnvironmentSpecSecretsObjectJSON struct {
	EnvironmentVariable apijson.Field
	Name                apijson.Field
	Session             apijson.Field
	Source              apijson.Field
	SourceRef           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecSecretsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecSecretsObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentGetResponseEnvironmentSpecSecretsObject) implementsEnvironmentGetResponseEnvironmentSpecSecret() {
}

type EnvironmentGetResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted struct {
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath string `json:"filePath,required"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                                                                                        `json:"sourceRef"`
	JSON      environmentGetResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON `json:"-"`
}

// environmentGetResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON
// contains the JSON metadata for the struct
// [EnvironmentGetResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted]
type environmentGetResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON struct {
	FilePath    apijson.Field
	Name        apijson.Field
	Session     apijson.Field
	Source      apijson.Field
	SourceRef   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentGetResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) implementsEnvironmentGetResponseEnvironmentSpecSecret() {
}

type EnvironmentGetResponseEnvironmentSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// value is the actual public key in the public key file format
	Value string                                                `json:"value"`
	JSON  environmentGetResponseEnvironmentSpecSSHPublicKeyJSON `json:"-"`
}

// environmentGetResponseEnvironmentSpecSSHPublicKeyJSON contains the JSON metadata
// for the struct [EnvironmentGetResponseEnvironmentSpecSSHPublicKey]
type environmentGetResponseEnvironmentSpecSSHPublicKeyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// Timeout configures the environment timeout
type EnvironmentGetResponseEnvironmentSpecTimeout struct {
	// A Duration represents a signed, fixed-length span of time represented as a count
	// of seconds and fractions of seconds at nanosecond resolution. It is independent
	// of any calendar and concepts like "day" or "month". It is related to Timestamp
	// in that the difference between two Timestamp values is a Duration and it can be
	// added or subtracted from a Timestamp. Range is approximately +-10,000 years.
	//
	// # Examples
	//
	// Example 1: Compute Duration from two Timestamps in pseudo code.
	//
	//	Timestamp start = ...;
	//	Timestamp end = ...;
	//	Duration duration = ...;
	//
	//	duration.seconds = end.seconds - start.seconds;
	//	duration.nanos = end.nanos - start.nanos;
	//
	//	if (duration.seconds < 0 && duration.nanos > 0) {
	//	  duration.seconds += 1;
	//	  duration.nanos -= 1000000000;
	//	} else if (duration.seconds > 0 && duration.nanos < 0) {
	//	  duration.seconds -= 1;
	//	  duration.nanos += 1000000000;
	//	}
	//
	// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
	//
	//	Timestamp start = ...;
	//	Duration duration = ...;
	//	Timestamp end = ...;
	//
	//	end.seconds = start.seconds + duration.seconds;
	//	end.nanos = start.nanos + duration.nanos;
	//
	//	if (end.nanos < 0) {
	//	  end.seconds -= 1;
	//	  end.nanos += 1000000000;
	//	} else if (end.nanos >= 1000000000) {
	//	  end.seconds += 1;
	//	  end.nanos -= 1000000000;
	//	}
	//
	// Example 3: Compute Duration from datetime.timedelta in Python.
	//
	//	td = datetime.timedelta(days=3, minutes=10)
	//	duration = Duration()
	//	duration.FromTimedelta(td)
	//
	// # JSON Mapping
	//
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by
	// the number of seconds, with nanoseconds expressed as fractional seconds. For
	// example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as
	// "3.000000001s", and 3 seconds and 1 microsecond should be expressed in JSON
	// format as "3.000001s".
	Disconnected string                                           `json:"disconnected" format:"regex"`
	JSON         environmentGetResponseEnvironmentSpecTimeoutJSON `json:"-"`
}

// environmentGetResponseEnvironmentSpecTimeoutJSON contains the JSON metadata for
// the struct [EnvironmentGetResponseEnvironmentSpecTimeout]
type environmentGetResponseEnvironmentSpecTimeoutJSON struct {
	Disconnected apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentSpecTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentSpecTimeoutJSON) RawJSON() string {
	return r.raw
}

// EnvironmentStatus describes an environment status
type EnvironmentGetResponseEnvironmentStatus struct {
	// EnvironmentActivitySignal used to signal activity for an environment.
	ActivitySignal EnvironmentGetResponseEnvironmentStatusActivitySignal `json:"activitySignal"`
	// automations_file contains the status of the automations file.
	AutomationsFile EnvironmentGetResponseEnvironmentStatusAutomationsFile `json:"automationsFile"`
	// content contains the status of the environment content.
	Content EnvironmentGetResponseEnvironmentStatusContent `json:"content"`
	// devcontainer contains the status of the devcontainer.
	Devcontainer EnvironmentGetResponseEnvironmentStatusDevcontainer `json:"devcontainer"`
	// environment_url contains the URL at which the environment can be accessed. This
	// field is only set if the environment is running.
	EnvironmentURLs EnvironmentGetResponseEnvironmentStatusEnvironmentURLs `json:"environmentUrls"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage []string `json:"failureMessage"`
	// machine contains the status of the environment machine
	Machine EnvironmentGetResponseEnvironmentStatusMachine `json:"machine"`
	// the phase of an environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentGetResponseEnvironmentStatusPhase `json:"phase"`
	// RunnerACK is the acknowledgement from the runner that is has received the
	// environment spec.
	RunnerAck EnvironmentGetResponseEnvironmentStatusRunnerAck `json:"runnerAck"`
	// secrets contains the status of the environment secrets
	Secrets []EnvironmentGetResponseEnvironmentStatusSecret `json:"secrets"`
	// ssh_public_keys contains the status of the environment ssh public keys
	SSHPublicKeys []EnvironmentGetResponseEnvironmentStatusSSHPublicKey `json:"sshPublicKeys"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion"`
	// warning_message contains warnings, e.g. when the environment is present but not
	// in the expected state.
	WarningMessage []string                                    `json:"warningMessage"`
	JSON           environmentGetResponseEnvironmentStatusJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusJSON contains the JSON metadata for the
// struct [EnvironmentGetResponseEnvironmentStatus]
type environmentGetResponseEnvironmentStatusJSON struct {
	ActivitySignal  apijson.Field
	AutomationsFile apijson.Field
	Content         apijson.Field
	Devcontainer    apijson.Field
	EnvironmentURLs apijson.Field
	FailureMessage  apijson.Field
	Machine         apijson.Field
	Phase           apijson.Field
	RunnerAck       apijson.Field
	Secrets         apijson.Field
	SSHPublicKeys   apijson.Field
	StatusVersion   apijson.Field
	WarningMessage  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusJSON) RawJSON() string {
	return r.raw
}

// EnvironmentActivitySignal used to signal activity for an environment.
type EnvironmentGetResponseEnvironmentStatusActivitySignal struct {
	// source of the activity signal, such as "VS Code", "SSH", or "Automations". It
	// should be a human-readable string that describes the source of the activity
	// signal.
	Source string `json:"source"`
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
	Timestamp time.Time                                                 `json:"timestamp" format:"date-time"`
	JSON      environmentGetResponseEnvironmentStatusActivitySignalJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusActivitySignalJSON contains the JSON
// metadata for the struct [EnvironmentGetResponseEnvironmentStatusActivitySignal]
type environmentGetResponseEnvironmentStatusActivitySignalJSON struct {
	Source      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusActivitySignal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusActivitySignalJSON) RawJSON() string {
	return r.raw
}

// automations_file contains the status of the automations file.
type EnvironmentGetResponseEnvironmentStatusAutomationsFile struct {
	// automations_file_path is the path to the automations file relative to the repo
	// root.
	AutomationsFilePath string `json:"automationsFilePath"`
	// automations_file_presence indicates how an automations file is present in the
	// environment.
	AutomationsFilePresence EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresence `json:"automationsFilePresence"`
	// failure_message contains the reason the automations file failed to be applied.
	// This is only set if the phase is FAILED.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the automations file.
	Phase EnvironmentGetResponseEnvironmentStatusAutomationsFilePhase `json:"phase"`
	// session is the automations file session that is currently applied in the
	// environment.
	Session string                                                     `json:"session"`
	JSON    environmentGetResponseEnvironmentStatusAutomationsFileJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusAutomationsFileJSON contains the JSON
// metadata for the struct [EnvironmentGetResponseEnvironmentStatusAutomationsFile]
type environmentGetResponseEnvironmentStatusAutomationsFileJSON struct {
	AutomationsFilePath     apijson.Field
	AutomationsFilePresence apijson.Field
	FailureMessage          apijson.Field
	Phase                   apijson.Field
	Session                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// automations_file_presence indicates how an automations file is present in the
// environment.
type EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresence string

const (
	EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceUnspecified EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceAbsent      EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_ABSENT"
	EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceDiscovered  EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceSpecified   EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceUnspecified, EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceAbsent, EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceDiscovered, EnvironmentGetResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the automations file.
type EnvironmentGetResponseEnvironmentStatusAutomationsFilePhase string

const (
	EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUnspecified  EnvironmentGetResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseCreating     EnvironmentGetResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_CREATING"
	EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseInitializing EnvironmentGetResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseReady        EnvironmentGetResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_READY"
	EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUpdating     EnvironmentGetResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_UPDATING"
	EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseFailed       EnvironmentGetResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentGetResponseEnvironmentStatusAutomationsFilePhase) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUnspecified, EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseCreating, EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseInitializing, EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseReady, EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUpdating, EnvironmentGetResponseEnvironmentStatusAutomationsFilePhaseContentPhaseFailed:
		return true
	}
	return false
}

// content contains the status of the environment content.
type EnvironmentGetResponseEnvironmentStatusContent struct {
	// content_location_in_machine is the location of the content in the machine
	ContentLocationInMachine string `json:"contentLocationInMachine"`
	// failure_message contains the reason the content initialization failed.
	FailureMessage string `json:"failureMessage"`
	// git is the Git working copy status of the environment. Note: this is a
	// best-effort field and more often than not will not be present. Its absence does
	// not indicate the absence of a working copy.
	Git EnvironmentGetResponseEnvironmentStatusContentGit `json:"git"`
	// phase is the current phase of the environment content
	Phase EnvironmentGetResponseEnvironmentStatusContentPhase `json:"phase"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the content is present but not in
	// the expected state.
	WarningMessage string                                             `json:"warningMessage"`
	JSON           environmentGetResponseEnvironmentStatusContentJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusContentJSON contains the JSON metadata
// for the struct [EnvironmentGetResponseEnvironmentStatusContent]
type environmentGetResponseEnvironmentStatusContentJSON struct {
	ContentLocationInMachine apijson.Field
	FailureMessage           apijson.Field
	Git                      apijson.Field
	Phase                    apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusContentJSON) RawJSON() string {
	return r.raw
}

// git is the Git working copy status of the environment. Note: this is a
// best-effort field and more often than not will not be present. Its absence does
// not indicate the absence of a working copy.
type EnvironmentGetResponseEnvironmentStatusContentGit struct {
	// branch is branch we're currently on
	Branch string `json:"branch"`
	// changed_files is an array of changed files in the environment, possibly
	// truncated
	ChangedFiles []EnvironmentGetResponseEnvironmentStatusContentGitChangedFile `json:"changedFiles"`
	// clone_url is the repository url as you would pass it to "git clone". Only HTTPS
	// clone URLs are supported.
	CloneURL string `json:"cloneUrl"`
	// latest_commit is the most recent commit on the current branch
	LatestCommit      string `json:"latestCommit"`
	TotalChangedFiles int64  `json:"totalChangedFiles"`
	// the total number of unpushed changes
	TotalUnpushedCommits int64 `json:"totalUnpushedCommits"`
	// unpushed_commits is an array of unpushed changes in the environment, possibly
	// truncated
	UnpushedCommits []string                                              `json:"unpushedCommits"`
	JSON            environmentGetResponseEnvironmentStatusContentGitJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusContentGitJSON contains the JSON metadata
// for the struct [EnvironmentGetResponseEnvironmentStatusContentGit]
type environmentGetResponseEnvironmentStatusContentGitJSON struct {
	Branch               apijson.Field
	ChangedFiles         apijson.Field
	CloneURL             apijson.Field
	LatestCommit         apijson.Field
	TotalChangedFiles    apijson.Field
	TotalUnpushedCommits apijson.Field
	UnpushedCommits      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusContentGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusContentGitJSON) RawJSON() string {
	return r.raw
}

type EnvironmentGetResponseEnvironmentStatusContentGitChangedFile struct {
	// ChangeType is the type of change that happened to the file
	ChangeType EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeType `json:"changeType"`
	// path is the path of the file
	Path string                                                           `json:"path"`
	JSON environmentGetResponseEnvironmentStatusContentGitChangedFileJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusContentGitChangedFileJSON contains the
// JSON metadata for the struct
// [EnvironmentGetResponseEnvironmentStatusContentGitChangedFile]
type environmentGetResponseEnvironmentStatusContentGitChangedFileJSON struct {
	ChangeType  apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusContentGitChangedFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusContentGitChangedFileJSON) RawJSON() string {
	return r.raw
}

// ChangeType is the type of change that happened to the file
type EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeType string

const (
	EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified        EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeAdded              EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_ADDED"
	EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeModified           EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_MODIFIED"
	EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeDeleted            EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_DELETED"
	EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeRenamed            EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_RENAMED"
	EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeCopied             EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_COPIED"
	EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UPDATED_BUT_UNMERGED"
	EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUntracked          EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNTRACKED"
)

func (r EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeType) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified, EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeAdded, EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeModified, EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeDeleted, EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeRenamed, EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeCopied, EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged, EnvironmentGetResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUntracked:
		return true
	}
	return false
}

// phase is the current phase of the environment content
type EnvironmentGetResponseEnvironmentStatusContentPhase string

const (
	EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseUnspecified  EnvironmentGetResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseCreating     EnvironmentGetResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_CREATING"
	EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseInitializing EnvironmentGetResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseReady        EnvironmentGetResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_READY"
	EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseUpdating     EnvironmentGetResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseFailed       EnvironmentGetResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentGetResponseEnvironmentStatusContentPhase) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseUnspecified, EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseCreating, EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseInitializing, EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseReady, EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseUpdating, EnvironmentGetResponseEnvironmentStatusContentPhaseContentPhaseFailed:
		return true
	}
	return false
}

// devcontainer contains the status of the devcontainer.
type EnvironmentGetResponseEnvironmentStatusDevcontainer struct {
	// container_id is the ID of the container.
	ContainerID string `json:"containerId"`
	// container_name is the name of the container that is used to connect to the
	// devcontainer
	ContainerName string `json:"containerName"`
	// devcontainerconfig_in_sync indicates if the devcontainer is up to date w.r.t.
	// the devcontainer config file.
	DevcontainerconfigInSync bool `json:"devcontainerconfigInSync"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// devcontainer_file_presence indicates how the devcontainer file is present in the
	// repo.
	DevcontainerFilePresence EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresence `json:"devcontainerFilePresence"`
	// failure_message contains the reason the devcontainer failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the devcontainer
	Phase EnvironmentGetResponseEnvironmentStatusDevcontainerPhase `json:"phase"`
	// remote_user is the user that is used to connect to the devcontainer
	RemoteUser string `json:"remoteUser"`
	// remote_workspace_folder is the folder that is used to connect to the
	// devcontainer
	RemoteWorkspaceFolder string `json:"remoteWorkspaceFolder"`
	// secrets_in_sync indicates if the secrets are up to date w.r.t. the running
	// devcontainer.
	SecretsInSync bool `json:"secretsInSync"`
	// session is the session that is currently active in the devcontainer.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the devcontainer is present but not
	// in the expected state.
	WarningMessage string                                                  `json:"warningMessage"`
	JSON           environmentGetResponseEnvironmentStatusDevcontainerJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusDevcontainerJSON contains the JSON
// metadata for the struct [EnvironmentGetResponseEnvironmentStatusDevcontainer]
type environmentGetResponseEnvironmentStatusDevcontainerJSON struct {
	ContainerID              apijson.Field
	ContainerName            apijson.Field
	DevcontainerconfigInSync apijson.Field
	DevcontainerFilePath     apijson.Field
	DevcontainerFilePresence apijson.Field
	FailureMessage           apijson.Field
	Phase                    apijson.Field
	RemoteUser               apijson.Field
	RemoteWorkspaceFolder    apijson.Field
	SecretsInSync            apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusDevcontainerJSON) RawJSON() string {
	return r.raw
}

// devcontainer_file_presence indicates how the devcontainer file is present in the
// repo.
type EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresence string

const (
	EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceUnspecified EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceGenerated   EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_GENERATED"
	EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceDiscovered  EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceSpecified   EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceUnspecified, EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceGenerated, EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceDiscovered, EnvironmentGetResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the devcontainer
type EnvironmentGetResponseEnvironmentStatusDevcontainerPhase string

const (
	EnvironmentGetResponseEnvironmentStatusDevcontainerPhasePhaseUnspecified EnvironmentGetResponseEnvironmentStatusDevcontainerPhase = "PHASE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentStatusDevcontainerPhasePhaseCreating    EnvironmentGetResponseEnvironmentStatusDevcontainerPhase = "PHASE_CREATING"
	EnvironmentGetResponseEnvironmentStatusDevcontainerPhasePhaseRunning     EnvironmentGetResponseEnvironmentStatusDevcontainerPhase = "PHASE_RUNNING"
	EnvironmentGetResponseEnvironmentStatusDevcontainerPhasePhaseStopped     EnvironmentGetResponseEnvironmentStatusDevcontainerPhase = "PHASE_STOPPED"
	EnvironmentGetResponseEnvironmentStatusDevcontainerPhasePhaseFailed      EnvironmentGetResponseEnvironmentStatusDevcontainerPhase = "PHASE_FAILED"
)

func (r EnvironmentGetResponseEnvironmentStatusDevcontainerPhase) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentStatusDevcontainerPhasePhaseUnspecified, EnvironmentGetResponseEnvironmentStatusDevcontainerPhasePhaseCreating, EnvironmentGetResponseEnvironmentStatusDevcontainerPhasePhaseRunning, EnvironmentGetResponseEnvironmentStatusDevcontainerPhasePhaseStopped, EnvironmentGetResponseEnvironmentStatusDevcontainerPhasePhaseFailed:
		return true
	}
	return false
}

// environment_url contains the URL at which the environment can be accessed. This
// field is only set if the environment is running.
type EnvironmentGetResponseEnvironmentStatusEnvironmentURLs struct {
	// logs is the URL at which the environment logs can be accessed.
	Logs  string                                                       `json:"logs"`
	Ports []EnvironmentGetResponseEnvironmentStatusEnvironmentURLsPort `json:"ports"`
	// SSH is the URL at which the environment can be accessed via SSH.
	SSH  EnvironmentGetResponseEnvironmentStatusEnvironmentURLsSSH  `json:"ssh"`
	JSON environmentGetResponseEnvironmentStatusEnvironmentURLsJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusEnvironmentURLsJSON contains the JSON
// metadata for the struct [EnvironmentGetResponseEnvironmentStatusEnvironmentURLs]
type environmentGetResponseEnvironmentStatusEnvironmentURLsJSON struct {
	Logs        apijson.Field
	Ports       apijson.Field
	SSH         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusEnvironmentURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusEnvironmentURLsJSON) RawJSON() string {
	return r.raw
}

type EnvironmentGetResponseEnvironmentStatusEnvironmentURLsPort struct {
	// port is the port number of the environment port
	Port int64 `json:"port"`
	// url is the URL at which the environment port can be accessed
	URL  string                                                         `json:"url"`
	JSON environmentGetResponseEnvironmentStatusEnvironmentURLsPortJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusEnvironmentURLsPortJSON contains the JSON
// metadata for the struct
// [EnvironmentGetResponseEnvironmentStatusEnvironmentURLsPort]
type environmentGetResponseEnvironmentStatusEnvironmentURLsPortJSON struct {
	Port        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusEnvironmentURLsPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusEnvironmentURLsPortJSON) RawJSON() string {
	return r.raw
}

// SSH is the URL at which the environment can be accessed via SSH.
type EnvironmentGetResponseEnvironmentStatusEnvironmentURLsSSH struct {
	URL  string                                                        `json:"url"`
	JSON environmentGetResponseEnvironmentStatusEnvironmentURLsSSHJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusEnvironmentURLsSSHJSON contains the JSON
// metadata for the struct
// [EnvironmentGetResponseEnvironmentStatusEnvironmentURLsSSH]
type environmentGetResponseEnvironmentStatusEnvironmentURLsSSHJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusEnvironmentURLsSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusEnvironmentURLsSSHJSON) RawJSON() string {
	return r.raw
}

// machine contains the status of the environment machine
type EnvironmentGetResponseEnvironmentStatusMachine struct {
	// failure_message contains the reason the machine failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the environment machine
	Phase EnvironmentGetResponseEnvironmentStatusMachinePhase `json:"phase"`
	// session is the session that is currently active in the machine.
	Session string `json:"session"`
	// timeout contains the reason the environment has timed out. If this field is
	// empty, the environment has not timed out.
	Timeout string `json:"timeout"`
	// versions contains the versions of components in the machine.
	Versions EnvironmentGetResponseEnvironmentStatusMachineVersions `json:"versions"`
	// warning_message contains warnings, e.g. when the machine is present but not in
	// the expected state.
	WarningMessage string                                             `json:"warningMessage"`
	JSON           environmentGetResponseEnvironmentStatusMachineJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusMachineJSON contains the JSON metadata
// for the struct [EnvironmentGetResponseEnvironmentStatusMachine]
type environmentGetResponseEnvironmentStatusMachineJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	Timeout        apijson.Field
	Versions       apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusMachineJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the environment machine
type EnvironmentGetResponseEnvironmentStatusMachinePhase string

const (
	EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseUnspecified EnvironmentGetResponseEnvironmentStatusMachinePhase = "PHASE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseCreating    EnvironmentGetResponseEnvironmentStatusMachinePhase = "PHASE_CREATING"
	EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseStarting    EnvironmentGetResponseEnvironmentStatusMachinePhase = "PHASE_STARTING"
	EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseRunning     EnvironmentGetResponseEnvironmentStatusMachinePhase = "PHASE_RUNNING"
	EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseStopping    EnvironmentGetResponseEnvironmentStatusMachinePhase = "PHASE_STOPPING"
	EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseStopped     EnvironmentGetResponseEnvironmentStatusMachinePhase = "PHASE_STOPPED"
	EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseDeleting    EnvironmentGetResponseEnvironmentStatusMachinePhase = "PHASE_DELETING"
	EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseDeleted     EnvironmentGetResponseEnvironmentStatusMachinePhase = "PHASE_DELETED"
)

func (r EnvironmentGetResponseEnvironmentStatusMachinePhase) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseUnspecified, EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseCreating, EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseStarting, EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseRunning, EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseStopping, EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseStopped, EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseDeleting, EnvironmentGetResponseEnvironmentStatusMachinePhasePhaseDeleted:
		return true
	}
	return false
}

// versions contains the versions of components in the machine.
type EnvironmentGetResponseEnvironmentStatusMachineVersions struct {
	SupervisorCommit  string                                                     `json:"supervisorCommit"`
	SupervisorVersion string                                                     `json:"supervisorVersion"`
	JSON              environmentGetResponseEnvironmentStatusMachineVersionsJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusMachineVersionsJSON contains the JSON
// metadata for the struct [EnvironmentGetResponseEnvironmentStatusMachineVersions]
type environmentGetResponseEnvironmentStatusMachineVersionsJSON struct {
	SupervisorCommit  apijson.Field
	SupervisorVersion apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusMachineVersions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusMachineVersionsJSON) RawJSON() string {
	return r.raw
}

// the phase of an environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentGetResponseEnvironmentStatusPhase string

const (
	EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseUnspecified EnvironmentGetResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseCreating    EnvironmentGetResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseStarting    EnvironmentGetResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseRunning     EnvironmentGetResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseUpdating    EnvironmentGetResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseStopping    EnvironmentGetResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseStopped     EnvironmentGetResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseDeleting    EnvironmentGetResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseDeleted     EnvironmentGetResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentGetResponseEnvironmentStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseUnspecified, EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseCreating, EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseStarting, EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseRunning, EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseUpdating, EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseStopping, EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseStopped, EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseDeleting, EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// RunnerACK is the acknowledgement from the runner that is has received the
// environment spec.
type EnvironmentGetResponseEnvironmentStatusRunnerAck struct {
	Message     string                                                     `json:"message"`
	SpecVersion string                                                     `json:"specVersion"`
	StatusCode  EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCode `json:"statusCode"`
	JSON        environmentGetResponseEnvironmentStatusRunnerAckJSON       `json:"-"`
}

// environmentGetResponseEnvironmentStatusRunnerAckJSON contains the JSON metadata
// for the struct [EnvironmentGetResponseEnvironmentStatusRunnerAck]
type environmentGetResponseEnvironmentStatusRunnerAckJSON struct {
	Message     apijson.Field
	SpecVersion apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusRunnerAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusRunnerAckJSON) RawJSON() string {
	return r.raw
}

type EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCode string

const (
	EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeUnspecified        EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeOk                 EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_OK"
	EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeInvalidResource    EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_INVALID_RESOURCE"
	EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeFailedPrecondition EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_FAILED_PRECONDITION"
)

func (r EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCode) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeUnspecified, EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeOk, EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeInvalidResource, EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeFailedPrecondition:
		return true
	}
	return false
}

type EnvironmentGetResponseEnvironmentStatusSecret struct {
	// failure_message contains the reason the secret failed to be materialize.
	FailureMessage string                                              `json:"failureMessage"`
	Phase          EnvironmentGetResponseEnvironmentStatusSecretsPhase `json:"phase"`
	SecretName     string                                              `json:"secretName"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the secret is present but not in
	// the expected state.
	WarningMessage string                                            `json:"warningMessage"`
	JSON           environmentGetResponseEnvironmentStatusSecretJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusSecretJSON contains the JSON metadata for
// the struct [EnvironmentGetResponseEnvironmentStatusSecret]
type environmentGetResponseEnvironmentStatusSecretJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	SecretName     apijson.Field
	Session        apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusSecretJSON) RawJSON() string {
	return r.raw
}

type EnvironmentGetResponseEnvironmentStatusSecretsPhase string

const (
	EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseUnspecified  EnvironmentGetResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseCreating     EnvironmentGetResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_CREATING"
	EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseInitializing EnvironmentGetResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseReady        EnvironmentGetResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_READY"
	EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseUpdating     EnvironmentGetResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseFailed       EnvironmentGetResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentGetResponseEnvironmentStatusSecretsPhase) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseUnspecified, EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseCreating, EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseInitializing, EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseReady, EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseUpdating, EnvironmentGetResponseEnvironmentStatusSecretsPhaseContentPhaseFailed:
		return true
	}
	return false
}

type EnvironmentGetResponseEnvironmentStatusSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// phase is the current phase of the public key
	Phase EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhase `json:"phase"`
	JSON  environmentGetResponseEnvironmentStatusSSHPublicKeyJSON   `json:"-"`
}

// environmentGetResponseEnvironmentStatusSSHPublicKeyJSON contains the JSON
// metadata for the struct [EnvironmentGetResponseEnvironmentStatusSSHPublicKey]
type environmentGetResponseEnvironmentStatusSSHPublicKeyJSON struct {
	ID          apijson.Field
	Phase       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponseEnvironmentStatusSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseEnvironmentStatusSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the public key
type EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhase string

const (
	EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUnspecified  EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseCreating     EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_CREATING"
	EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseInitializing EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseReady        EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_READY"
	EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUpdating     EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseFailed       EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhase) IsKnown() bool {
	switch r {
	case EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUnspecified, EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseCreating, EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseInitializing, EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseReady, EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUpdating, EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseFailed:
		return true
	}
	return false
}

type EnvironmentUpdateResponse = interface{}

// +resource get environment
type EnvironmentListResponse struct {
	// ID is a unique identifier of this environment. No other environment with the
	// same name must be managed by this environment manager
	ID string `json:"id"`
	// EnvironmentMetadata is data associated with an environment that's required for
	// other parts of the system to function
	Metadata EnvironmentListResponseMetadata `json:"metadata"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec EnvironmentListResponseSpec `json:"spec"`
	// EnvironmentStatus describes an environment status
	Status EnvironmentListResponseStatus `json:"status"`
	JSON   environmentListResponseJSON   `json:"-"`
}

// environmentListResponseJSON contains the JSON metadata for the struct
// [EnvironmentListResponse]
type environmentListResponseJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseJSON) RawJSON() string {
	return r.raw
}

// EnvironmentMetadata is data associated with an environment that's required for
// other parts of the system to function
type EnvironmentListResponseMetadata struct {
	// annotations are key/value pairs that gets attached to the environment.
	// +internal - not yet implemented
	Annotations map[string]string `json:"annotations"`
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
	// creator is the identity of the creator of the environment
	Creator EnvironmentListResponseMetadataCreator `json:"creator"`
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
	LastStartedAt time.Time `json:"lastStartedAt" format:"date-time"`
	// name is the name of the environment as specified by the user
	Name string `json:"name"`
	// organization_id is the ID of the organization that contains the environment
	OrganizationID string `json:"organizationId" format:"uuid"`
	// original_context_url is the normalized URL from which the environment was
	// created
	OriginalContextURL string `json:"originalContextUrl"`
	// If the Environment was started from a project, the project_id will reference the
	// project.
	ProjectID string `json:"projectId"`
	// Runner is the ID of the runner that runs this environment.
	RunnerID string                              `json:"runnerId"`
	JSON     environmentListResponseMetadataJSON `json:"-"`
}

// environmentListResponseMetadataJSON contains the JSON metadata for the struct
// [EnvironmentListResponseMetadata]
type environmentListResponseMetadataJSON struct {
	Annotations        apijson.Field
	CreatedAt          apijson.Field
	Creator            apijson.Field
	LastStartedAt      apijson.Field
	Name               apijson.Field
	OrganizationID     apijson.Field
	OriginalContextURL apijson.Field
	ProjectID          apijson.Field
	RunnerID           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentListResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the creator of the environment
type EnvironmentListResponseMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentListResponseMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentListResponseMetadataCreatorJSON      `json:"-"`
}

// environmentListResponseMetadataCreatorJSON contains the JSON metadata for the
// struct [EnvironmentListResponseMetadataCreator]
type environmentListResponseMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentListResponseMetadataCreatorPrincipal string

const (
	EnvironmentListResponseMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentListResponseMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentListResponseMetadataCreatorPrincipalPrincipalAccount        EnvironmentListResponseMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentListResponseMetadataCreatorPrincipalPrincipalUser           EnvironmentListResponseMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentListResponseMetadataCreatorPrincipalPrincipalRunner         EnvironmentListResponseMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentListResponseMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentListResponseMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentListResponseMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentListResponseMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentListResponseMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentListResponseMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentListResponseMetadataCreatorPrincipalPrincipalAccount, EnvironmentListResponseMetadataCreatorPrincipalPrincipalUser, EnvironmentListResponseMetadataCreatorPrincipalPrincipalRunner, EnvironmentListResponseMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentListResponseMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type EnvironmentListResponseSpec struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission EnvironmentListResponseSpecAdmission `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile EnvironmentListResponseSpecAutomationsFile `json:"automationsFile"`
	// content is the content spec of the environment
	Content EnvironmentListResponseSpecContent `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase EnvironmentListResponseSpecDesiredPhase `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer EnvironmentListResponseSpecDevcontainer `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine EnvironmentListResponseSpecMachine `json:"machine"`
	// ports is the set of ports which ought to be exposed to the internet
	Ports []EnvironmentListResponseSpecPort `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets []EnvironmentListResponseSpecSecret `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion string `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys []EnvironmentListResponseSpecSSHPublicKey `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout EnvironmentListResponseSpecTimeout `json:"timeout"`
	JSON    environmentListResponseSpecJSON    `json:"-"`
}

// environmentListResponseSpecJSON contains the JSON metadata for the struct
// [EnvironmentListResponseSpec]
type environmentListResponseSpecJSON struct {
	Admission       apijson.Field
	AutomationsFile apijson.Field
	Content         apijson.Field
	DesiredPhase    apijson.Field
	Devcontainer    apijson.Field
	Machine         apijson.Field
	Ports           apijson.Field
	Secrets         apijson.Field
	SpecVersion     apijson.Field
	SSHPublicKeys   apijson.Field
	Timeout         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EnvironmentListResponseSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentListResponseSpecAdmission string

const (
	EnvironmentListResponseSpecAdmissionAdmissionLevelUnspecified EnvironmentListResponseSpecAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentListResponseSpecAdmissionAdmissionLevelOwnerOnly   EnvironmentListResponseSpecAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentListResponseSpecAdmissionAdmissionLevelEveryone    EnvironmentListResponseSpecAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentListResponseSpecAdmission) IsKnown() bool {
	switch r {
	case EnvironmentListResponseSpecAdmissionAdmissionLevelUnspecified, EnvironmentListResponseSpecAdmissionAdmissionLevelOwnerOnly, EnvironmentListResponseSpecAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// automations_file is the automations file spec of the environment
type EnvironmentListResponseSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath string                                         `json:"automationsFilePath"`
	Session             string                                         `json:"session"`
	JSON                environmentListResponseSpecAutomationsFileJSON `json:"-"`
}

// environmentListResponseSpecAutomationsFileJSON contains the JSON metadata for
// the struct [EnvironmentListResponseSpecAutomationsFile]
type environmentListResponseSpecAutomationsFileJSON struct {
	AutomationsFilePath apijson.Field
	Session             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// content is the content spec of the environment
type EnvironmentListResponseSpecContent struct {
	// The Git email address
	GitEmail string `json:"gitEmail"`
	// The Git username
	GitUsername string `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer EnvironmentListResponseSpecContentInitializer `json:"initializer"`
	Session     string                                        `json:"session"`
	JSON        environmentListResponseSpecContentJSON        `json:"-"`
}

// environmentListResponseSpecContentJSON contains the JSON metadata for the struct
// [EnvironmentListResponseSpecContent]
type environmentListResponseSpecContentJSON struct {
	GitEmail    apijson.Field
	GitUsername apijson.Field
	Initializer apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecContentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentInitializer specifies how an environment is to be initialized
type EnvironmentListResponseSpecContentInitializer struct {
	Specs []EnvironmentListResponseSpecContentInitializerSpec `json:"specs"`
	JSON  environmentListResponseSpecContentInitializerJSON   `json:"-"`
}

// environmentListResponseSpecContentInitializerJSON contains the JSON metadata for
// the struct [EnvironmentListResponseSpecContentInitializer]
type environmentListResponseSpecContentInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecContentInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecContentInitializerJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseSpecContentInitializerSpec struct {
	// This field can have the runtime type of
	// [EnvironmentListResponseSpecContentInitializerSpecsContextURLContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [EnvironmentListResponseSpecContentInitializerSpecsGitGit].
	Git   interface{}                                           `json:"git"`
	JSON  environmentListResponseSpecContentInitializerSpecJSON `json:"-"`
	union EnvironmentListResponseSpecContentInitializerSpecsUnion
}

// environmentListResponseSpecContentInitializerSpecJSON contains the JSON metadata
// for the struct [EnvironmentListResponseSpecContentInitializerSpec]
type environmentListResponseSpecContentInitializerSpecJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentListResponseSpecContentInitializerSpecJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentListResponseSpecContentInitializerSpec) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentListResponseSpecContentInitializerSpec{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EnvironmentListResponseSpecContentInitializerSpecsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentListResponseSpecContentInitializerSpecsContextURL],
// [EnvironmentListResponseSpecContentInitializerSpecsGit].
func (r EnvironmentListResponseSpecContentInitializerSpec) AsUnion() EnvironmentListResponseSpecContentInitializerSpecsUnion {
	return r.union
}

// Union satisfied by
// [EnvironmentListResponseSpecContentInitializerSpecsContextURL] or
// [EnvironmentListResponseSpecContentInitializerSpecsGit].
type EnvironmentListResponseSpecContentInitializerSpecsUnion interface {
	implementsEnvironmentListResponseSpecContentInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentListResponseSpecContentInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentListResponseSpecContentInitializerSpecsContextURL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentListResponseSpecContentInitializerSpecsGit{}),
		},
	)
}

type EnvironmentListResponseSpecContentInitializerSpecsContextURL struct {
	ContextURL EnvironmentListResponseSpecContentInitializerSpecsContextURLContextURL `json:"contextUrl,required"`
	JSON       environmentListResponseSpecContentInitializerSpecsContextURLJSON       `json:"-"`
}

// environmentListResponseSpecContentInitializerSpecsContextURLJSON contains the
// JSON metadata for the struct
// [EnvironmentListResponseSpecContentInitializerSpecsContextURL]
type environmentListResponseSpecContentInitializerSpecsContextURLJSON struct {
	ContextURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecContentInitializerSpecsContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecContentInitializerSpecsContextURLJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentListResponseSpecContentInitializerSpecsContextURL) implementsEnvironmentListResponseSpecContentInitializerSpec() {
}

type EnvironmentListResponseSpecContentInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                                     `json:"url" format:"uri"`
	JSON environmentListResponseSpecContentInitializerSpecsContextURLContextURLJSON `json:"-"`
}

// environmentListResponseSpecContentInitializerSpecsContextURLContextURLJSON
// contains the JSON metadata for the struct
// [EnvironmentListResponseSpecContentInitializerSpecsContextURLContextURL]
type environmentListResponseSpecContentInitializerSpecsContextURLContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecContentInitializerSpecsContextURLContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecContentInitializerSpecsContextURLContextURLJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseSpecContentInitializerSpecsGit struct {
	Git  EnvironmentListResponseSpecContentInitializerSpecsGitGit  `json:"git,required"`
	JSON environmentListResponseSpecContentInitializerSpecsGitJSON `json:"-"`
}

// environmentListResponseSpecContentInitializerSpecsGitJSON contains the JSON
// metadata for the struct [EnvironmentListResponseSpecContentInitializerSpecsGit]
type environmentListResponseSpecContentInitializerSpecsGitJSON struct {
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecContentInitializerSpecsGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecContentInitializerSpecsGitJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentListResponseSpecContentInitializerSpecsGit) implementsEnvironmentListResponseSpecContentInitializerSpec() {
}

type EnvironmentListResponseSpecContentInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                                       `json:"upstreamRemoteUri"`
	JSON              environmentListResponseSpecContentInitializerSpecsGitGitJSON `json:"-"`
}

// environmentListResponseSpecContentInitializerSpecsGitGitJSON contains the JSON
// metadata for the struct
// [EnvironmentListResponseSpecContentInitializerSpecsGitGit]
type environmentListResponseSpecContentInitializerSpecsGitGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecContentInitializerSpecsGitGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecContentInitializerSpecsGitGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetMode string

const (
	EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, EnvironmentListResponseSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

// Phase is the desired phase of the environment
type EnvironmentListResponseSpecDesiredPhase string

const (
	EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseUnspecified EnvironmentListResponseSpecDesiredPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseCreating    EnvironmentListResponseSpecDesiredPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseStarting    EnvironmentListResponseSpecDesiredPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseRunning     EnvironmentListResponseSpecDesiredPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseUpdating    EnvironmentListResponseSpecDesiredPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseStopping    EnvironmentListResponseSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseStopped     EnvironmentListResponseSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseDeleting    EnvironmentListResponseSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseDeleted     EnvironmentListResponseSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentListResponseSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseUnspecified, EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseCreating, EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseStarting, EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseRunning, EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseUpdating, EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseStopping, EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseStopped, EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseDeleting, EnvironmentListResponseSpecDesiredPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// devcontainer is the devcontainer spec of the environment
type EnvironmentListResponseSpecDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath string                                      `json:"devcontainerFilePath"`
	Session              string                                      `json:"session"`
	JSON                 environmentListResponseSpecDevcontainerJSON `json:"-"`
}

// environmentListResponseSpecDevcontainerJSON contains the JSON metadata for the
// struct [EnvironmentListResponseSpecDevcontainer]
type environmentListResponseSpecDevcontainerJSON struct {
	DevcontainerFilePath apijson.Field
	Session              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecDevcontainerJSON) RawJSON() string {
	return r.raw
}

// machine is the machine spec of the environment
type EnvironmentListResponseSpecMachine struct {
	// Class denotes the class of the environment we ought to start
	Class   string                                 `json:"class" format:"uuid"`
	Session string                                 `json:"session"`
	JSON    environmentListResponseSpecMachineJSON `json:"-"`
}

// environmentListResponseSpecMachineJSON contains the JSON metadata for the struct
// [EnvironmentListResponseSpecMachine]
type environmentListResponseSpecMachineJSON struct {
	Class       apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecMachineJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseSpecPort struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission EnvironmentListResponseSpecPortsAdmission `json:"admission"`
	// name of this port
	Name string `json:"name"`
	// port number
	Port int64                               `json:"port"`
	JSON environmentListResponseSpecPortJSON `json:"-"`
}

// environmentListResponseSpecPortJSON contains the JSON metadata for the struct
// [EnvironmentListResponseSpecPort]
type environmentListResponseSpecPortJSON struct {
	Admission   apijson.Field
	Name        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecPortJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentListResponseSpecPortsAdmission string

const (
	EnvironmentListResponseSpecPortsAdmissionAdmissionLevelUnspecified EnvironmentListResponseSpecPortsAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentListResponseSpecPortsAdmissionAdmissionLevelOwnerOnly   EnvironmentListResponseSpecPortsAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentListResponseSpecPortsAdmissionAdmissionLevelEveryone    EnvironmentListResponseSpecPortsAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentListResponseSpecPortsAdmission) IsKnown() bool {
	switch r {
	case EnvironmentListResponseSpecPortsAdmissionAdmissionLevelUnspecified, EnvironmentListResponseSpecPortsAdmissionAdmissionLevelOwnerOnly, EnvironmentListResponseSpecPortsAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

type EnvironmentListResponseSpecSecret struct {
	EnvironmentVariable string `json:"environmentVariable"`
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath          string `json:"filePath"`
	GitCredentialHost string `json:"gitCredentialHost"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                `json:"sourceRef"`
	JSON      environmentListResponseSpecSecretJSON `json:"-"`
	union     EnvironmentListResponseSpecSecretsUnion
}

// environmentListResponseSpecSecretJSON contains the JSON metadata for the struct
// [EnvironmentListResponseSpecSecret]
type environmentListResponseSpecSecretJSON struct {
	EnvironmentVariable apijson.Field
	FilePath            apijson.Field
	GitCredentialHost   apijson.Field
	Name                apijson.Field
	Session             apijson.Field
	Source              apijson.Field
	SourceRef           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r environmentListResponseSpecSecretJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentListResponseSpecSecret) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentListResponseSpecSecret{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EnvironmentListResponseSpecSecretsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentListResponseSpecSecretsObject],
// [EnvironmentListResponseSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted],
// [EnvironmentListResponseSpecSecretsObject].
func (r EnvironmentListResponseSpecSecret) AsUnion() EnvironmentListResponseSpecSecretsUnion {
	return r.union
}

// Union satisfied by [EnvironmentListResponseSpecSecretsObject],
// [EnvironmentListResponseSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted]
// or [EnvironmentListResponseSpecSecretsObject].
type EnvironmentListResponseSpecSecretsUnion interface {
	implementsEnvironmentListResponseSpecSecret()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentListResponseSpecSecretsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentListResponseSpecSecretsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentListResponseSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentListResponseSpecSecretsObject{}),
		},
	)
}

type EnvironmentListResponseSpecSecretsObject struct {
	EnvironmentVariable string `json:"environmentVariable,required"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                       `json:"sourceRef"`
	JSON      environmentListResponseSpecSecretsObjectJSON `json:"-"`
}

// environmentListResponseSpecSecretsObjectJSON contains the JSON metadata for the
// struct [EnvironmentListResponseSpecSecretsObject]
type environmentListResponseSpecSecretsObjectJSON struct {
	EnvironmentVariable apijson.Field
	Name                apijson.Field
	Session             apijson.Field
	Source              apijson.Field
	SourceRef           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecSecretsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecSecretsObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentListResponseSpecSecretsObject) implementsEnvironmentListResponseSpecSecret() {}

type EnvironmentListResponseSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted struct {
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath string `json:"filePath,required"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                                                                              `json:"sourceRef"`
	JSON      environmentListResponseSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON `json:"-"`
}

// environmentListResponseSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON
// contains the JSON metadata for the struct
// [EnvironmentListResponseSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted]
type environmentListResponseSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON struct {
	FilePath    apijson.Field
	Name        apijson.Field
	Session     apijson.Field
	Source      apijson.Field
	SourceRef   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentListResponseSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) implementsEnvironmentListResponseSpecSecret() {
}

type EnvironmentListResponseSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// value is the actual public key in the public key file format
	Value string                                      `json:"value"`
	JSON  environmentListResponseSpecSSHPublicKeyJSON `json:"-"`
}

// environmentListResponseSpecSSHPublicKeyJSON contains the JSON metadata for the
// struct [EnvironmentListResponseSpecSSHPublicKey]
type environmentListResponseSpecSSHPublicKeyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// Timeout configures the environment timeout
type EnvironmentListResponseSpecTimeout struct {
	// A Duration represents a signed, fixed-length span of time represented as a count
	// of seconds and fractions of seconds at nanosecond resolution. It is independent
	// of any calendar and concepts like "day" or "month". It is related to Timestamp
	// in that the difference between two Timestamp values is a Duration and it can be
	// added or subtracted from a Timestamp. Range is approximately +-10,000 years.
	//
	// # Examples
	//
	// Example 1: Compute Duration from two Timestamps in pseudo code.
	//
	//	Timestamp start = ...;
	//	Timestamp end = ...;
	//	Duration duration = ...;
	//
	//	duration.seconds = end.seconds - start.seconds;
	//	duration.nanos = end.nanos - start.nanos;
	//
	//	if (duration.seconds < 0 && duration.nanos > 0) {
	//	  duration.seconds += 1;
	//	  duration.nanos -= 1000000000;
	//	} else if (duration.seconds > 0 && duration.nanos < 0) {
	//	  duration.seconds -= 1;
	//	  duration.nanos += 1000000000;
	//	}
	//
	// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
	//
	//	Timestamp start = ...;
	//	Duration duration = ...;
	//	Timestamp end = ...;
	//
	//	end.seconds = start.seconds + duration.seconds;
	//	end.nanos = start.nanos + duration.nanos;
	//
	//	if (end.nanos < 0) {
	//	  end.seconds -= 1;
	//	  end.nanos += 1000000000;
	//	} else if (end.nanos >= 1000000000) {
	//	  end.seconds += 1;
	//	  end.nanos -= 1000000000;
	//	}
	//
	// Example 3: Compute Duration from datetime.timedelta in Python.
	//
	//	td = datetime.timedelta(days=3, minutes=10)
	//	duration = Duration()
	//	duration.FromTimedelta(td)
	//
	// # JSON Mapping
	//
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by
	// the number of seconds, with nanoseconds expressed as fractional seconds. For
	// example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as
	// "3.000000001s", and 3 seconds and 1 microsecond should be expressed in JSON
	// format as "3.000001s".
	Disconnected string                                 `json:"disconnected" format:"regex"`
	JSON         environmentListResponseSpecTimeoutJSON `json:"-"`
}

// environmentListResponseSpecTimeoutJSON contains the JSON metadata for the struct
// [EnvironmentListResponseSpecTimeout]
type environmentListResponseSpecTimeoutJSON struct {
	Disconnected apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentListResponseSpecTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseSpecTimeoutJSON) RawJSON() string {
	return r.raw
}

// EnvironmentStatus describes an environment status
type EnvironmentListResponseStatus struct {
	// EnvironmentActivitySignal used to signal activity for an environment.
	ActivitySignal EnvironmentListResponseStatusActivitySignal `json:"activitySignal"`
	// automations_file contains the status of the automations file.
	AutomationsFile EnvironmentListResponseStatusAutomationsFile `json:"automationsFile"`
	// content contains the status of the environment content.
	Content EnvironmentListResponseStatusContent `json:"content"`
	// devcontainer contains the status of the devcontainer.
	Devcontainer EnvironmentListResponseStatusDevcontainer `json:"devcontainer"`
	// environment_url contains the URL at which the environment can be accessed. This
	// field is only set if the environment is running.
	EnvironmentURLs EnvironmentListResponseStatusEnvironmentURLs `json:"environmentUrls"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage []string `json:"failureMessage"`
	// machine contains the status of the environment machine
	Machine EnvironmentListResponseStatusMachine `json:"machine"`
	// the phase of an environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentListResponseStatusPhase `json:"phase"`
	// RunnerACK is the acknowledgement from the runner that is has received the
	// environment spec.
	RunnerAck EnvironmentListResponseStatusRunnerAck `json:"runnerAck"`
	// secrets contains the status of the environment secrets
	Secrets []EnvironmentListResponseStatusSecret `json:"secrets"`
	// ssh_public_keys contains the status of the environment ssh public keys
	SSHPublicKeys []EnvironmentListResponseStatusSSHPublicKey `json:"sshPublicKeys"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion"`
	// warning_message contains warnings, e.g. when the environment is present but not
	// in the expected state.
	WarningMessage []string                          `json:"warningMessage"`
	JSON           environmentListResponseStatusJSON `json:"-"`
}

// environmentListResponseStatusJSON contains the JSON metadata for the struct
// [EnvironmentListResponseStatus]
type environmentListResponseStatusJSON struct {
	ActivitySignal  apijson.Field
	AutomationsFile apijson.Field
	Content         apijson.Field
	Devcontainer    apijson.Field
	EnvironmentURLs apijson.Field
	FailureMessage  apijson.Field
	Machine         apijson.Field
	Phase           apijson.Field
	RunnerAck       apijson.Field
	Secrets         apijson.Field
	SSHPublicKeys   apijson.Field
	StatusVersion   apijson.Field
	WarningMessage  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EnvironmentListResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusJSON) RawJSON() string {
	return r.raw
}

// EnvironmentActivitySignal used to signal activity for an environment.
type EnvironmentListResponseStatusActivitySignal struct {
	// source of the activity signal, such as "VS Code", "SSH", or "Automations". It
	// should be a human-readable string that describes the source of the activity
	// signal.
	Source string `json:"source"`
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
	Timestamp time.Time                                       `json:"timestamp" format:"date-time"`
	JSON      environmentListResponseStatusActivitySignalJSON `json:"-"`
}

// environmentListResponseStatusActivitySignalJSON contains the JSON metadata for
// the struct [EnvironmentListResponseStatusActivitySignal]
type environmentListResponseStatusActivitySignalJSON struct {
	Source      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusActivitySignal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusActivitySignalJSON) RawJSON() string {
	return r.raw
}

// automations_file contains the status of the automations file.
type EnvironmentListResponseStatusAutomationsFile struct {
	// automations_file_path is the path to the automations file relative to the repo
	// root.
	AutomationsFilePath string `json:"automationsFilePath"`
	// automations_file_presence indicates how an automations file is present in the
	// environment.
	AutomationsFilePresence EnvironmentListResponseStatusAutomationsFileAutomationsFilePresence `json:"automationsFilePresence"`
	// failure_message contains the reason the automations file failed to be applied.
	// This is only set if the phase is FAILED.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the automations file.
	Phase EnvironmentListResponseStatusAutomationsFilePhase `json:"phase"`
	// session is the automations file session that is currently applied in the
	// environment.
	Session string                                           `json:"session"`
	JSON    environmentListResponseStatusAutomationsFileJSON `json:"-"`
}

// environmentListResponseStatusAutomationsFileJSON contains the JSON metadata for
// the struct [EnvironmentListResponseStatusAutomationsFile]
type environmentListResponseStatusAutomationsFileJSON struct {
	AutomationsFilePath     apijson.Field
	AutomationsFilePresence apijson.Field
	FailureMessage          apijson.Field
	Phase                   apijson.Field
	Session                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// automations_file_presence indicates how an automations file is present in the
// environment.
type EnvironmentListResponseStatusAutomationsFileAutomationsFilePresence string

const (
	EnvironmentListResponseStatusAutomationsFileAutomationsFilePresencePresenceUnspecified EnvironmentListResponseStatusAutomationsFileAutomationsFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentListResponseStatusAutomationsFileAutomationsFilePresencePresenceAbsent      EnvironmentListResponseStatusAutomationsFileAutomationsFilePresence = "PRESENCE_ABSENT"
	EnvironmentListResponseStatusAutomationsFileAutomationsFilePresencePresenceDiscovered  EnvironmentListResponseStatusAutomationsFileAutomationsFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentListResponseStatusAutomationsFileAutomationsFilePresencePresenceSpecified   EnvironmentListResponseStatusAutomationsFileAutomationsFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentListResponseStatusAutomationsFileAutomationsFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentListResponseStatusAutomationsFileAutomationsFilePresencePresenceUnspecified, EnvironmentListResponseStatusAutomationsFileAutomationsFilePresencePresenceAbsent, EnvironmentListResponseStatusAutomationsFileAutomationsFilePresencePresenceDiscovered, EnvironmentListResponseStatusAutomationsFileAutomationsFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the automations file.
type EnvironmentListResponseStatusAutomationsFilePhase string

const (
	EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseUnspecified  EnvironmentListResponseStatusAutomationsFilePhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseCreating     EnvironmentListResponseStatusAutomationsFilePhase = "CONTENT_PHASE_CREATING"
	EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseInitializing EnvironmentListResponseStatusAutomationsFilePhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseReady        EnvironmentListResponseStatusAutomationsFilePhase = "CONTENT_PHASE_READY"
	EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseUpdating     EnvironmentListResponseStatusAutomationsFilePhase = "CONTENT_PHASE_UPDATING"
	EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseFailed       EnvironmentListResponseStatusAutomationsFilePhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentListResponseStatusAutomationsFilePhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseUnspecified, EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseCreating, EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseInitializing, EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseReady, EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseUpdating, EnvironmentListResponseStatusAutomationsFilePhaseContentPhaseFailed:
		return true
	}
	return false
}

// content contains the status of the environment content.
type EnvironmentListResponseStatusContent struct {
	// content_location_in_machine is the location of the content in the machine
	ContentLocationInMachine string `json:"contentLocationInMachine"`
	// failure_message contains the reason the content initialization failed.
	FailureMessage string `json:"failureMessage"`
	// git is the Git working copy status of the environment. Note: this is a
	// best-effort field and more often than not will not be present. Its absence does
	// not indicate the absence of a working copy.
	Git EnvironmentListResponseStatusContentGit `json:"git"`
	// phase is the current phase of the environment content
	Phase EnvironmentListResponseStatusContentPhase `json:"phase"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the content is present but not in
	// the expected state.
	WarningMessage string                                   `json:"warningMessage"`
	JSON           environmentListResponseStatusContentJSON `json:"-"`
}

// environmentListResponseStatusContentJSON contains the JSON metadata for the
// struct [EnvironmentListResponseStatusContent]
type environmentListResponseStatusContentJSON struct {
	ContentLocationInMachine apijson.Field
	FailureMessage           apijson.Field
	Git                      apijson.Field
	Phase                    apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusContentJSON) RawJSON() string {
	return r.raw
}

// git is the Git working copy status of the environment. Note: this is a
// best-effort field and more often than not will not be present. Its absence does
// not indicate the absence of a working copy.
type EnvironmentListResponseStatusContentGit struct {
	// branch is branch we're currently on
	Branch string `json:"branch"`
	// changed_files is an array of changed files in the environment, possibly
	// truncated
	ChangedFiles []EnvironmentListResponseStatusContentGitChangedFile `json:"changedFiles"`
	// clone_url is the repository url as you would pass it to "git clone". Only HTTPS
	// clone URLs are supported.
	CloneURL string `json:"cloneUrl"`
	// latest_commit is the most recent commit on the current branch
	LatestCommit      string `json:"latestCommit"`
	TotalChangedFiles int64  `json:"totalChangedFiles"`
	// the total number of unpushed changes
	TotalUnpushedCommits int64 `json:"totalUnpushedCommits"`
	// unpushed_commits is an array of unpushed changes in the environment, possibly
	// truncated
	UnpushedCommits []string                                    `json:"unpushedCommits"`
	JSON            environmentListResponseStatusContentGitJSON `json:"-"`
}

// environmentListResponseStatusContentGitJSON contains the JSON metadata for the
// struct [EnvironmentListResponseStatusContentGit]
type environmentListResponseStatusContentGitJSON struct {
	Branch               apijson.Field
	ChangedFiles         apijson.Field
	CloneURL             apijson.Field
	LatestCommit         apijson.Field
	TotalChangedFiles    apijson.Field
	TotalUnpushedCommits apijson.Field
	UnpushedCommits      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusContentGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusContentGitJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseStatusContentGitChangedFile struct {
	// ChangeType is the type of change that happened to the file
	ChangeType EnvironmentListResponseStatusContentGitChangedFilesChangeType `json:"changeType"`
	// path is the path of the file
	Path string                                                 `json:"path"`
	JSON environmentListResponseStatusContentGitChangedFileJSON `json:"-"`
}

// environmentListResponseStatusContentGitChangedFileJSON contains the JSON
// metadata for the struct [EnvironmentListResponseStatusContentGitChangedFile]
type environmentListResponseStatusContentGitChangedFileJSON struct {
	ChangeType  apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusContentGitChangedFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusContentGitChangedFileJSON) RawJSON() string {
	return r.raw
}

// ChangeType is the type of change that happened to the file
type EnvironmentListResponseStatusContentGitChangedFilesChangeType string

const (
	EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified        EnvironmentListResponseStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNSPECIFIED"
	EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeAdded              EnvironmentListResponseStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_ADDED"
	EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeModified           EnvironmentListResponseStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_MODIFIED"
	EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeDeleted            EnvironmentListResponseStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_DELETED"
	EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeRenamed            EnvironmentListResponseStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_RENAMED"
	EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeCopied             EnvironmentListResponseStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_COPIED"
	EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged EnvironmentListResponseStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UPDATED_BUT_UNMERGED"
	EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeUntracked          EnvironmentListResponseStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNTRACKED"
)

func (r EnvironmentListResponseStatusContentGitChangedFilesChangeType) IsKnown() bool {
	switch r {
	case EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified, EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeAdded, EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeModified, EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeDeleted, EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeRenamed, EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeCopied, EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged, EnvironmentListResponseStatusContentGitChangedFilesChangeTypeChangeTypeUntracked:
		return true
	}
	return false
}

// phase is the current phase of the environment content
type EnvironmentListResponseStatusContentPhase string

const (
	EnvironmentListResponseStatusContentPhaseContentPhaseUnspecified  EnvironmentListResponseStatusContentPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseStatusContentPhaseContentPhaseCreating     EnvironmentListResponseStatusContentPhase = "CONTENT_PHASE_CREATING"
	EnvironmentListResponseStatusContentPhaseContentPhaseInitializing EnvironmentListResponseStatusContentPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentListResponseStatusContentPhaseContentPhaseReady        EnvironmentListResponseStatusContentPhase = "CONTENT_PHASE_READY"
	EnvironmentListResponseStatusContentPhaseContentPhaseUpdating     EnvironmentListResponseStatusContentPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentListResponseStatusContentPhaseContentPhaseFailed       EnvironmentListResponseStatusContentPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentListResponseStatusContentPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseStatusContentPhaseContentPhaseUnspecified, EnvironmentListResponseStatusContentPhaseContentPhaseCreating, EnvironmentListResponseStatusContentPhaseContentPhaseInitializing, EnvironmentListResponseStatusContentPhaseContentPhaseReady, EnvironmentListResponseStatusContentPhaseContentPhaseUpdating, EnvironmentListResponseStatusContentPhaseContentPhaseFailed:
		return true
	}
	return false
}

// devcontainer contains the status of the devcontainer.
type EnvironmentListResponseStatusDevcontainer struct {
	// container_id is the ID of the container.
	ContainerID string `json:"containerId"`
	// container_name is the name of the container that is used to connect to the
	// devcontainer
	ContainerName string `json:"containerName"`
	// devcontainerconfig_in_sync indicates if the devcontainer is up to date w.r.t.
	// the devcontainer config file.
	DevcontainerconfigInSync bool `json:"devcontainerconfigInSync"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// devcontainer_file_presence indicates how the devcontainer file is present in the
	// repo.
	DevcontainerFilePresence EnvironmentListResponseStatusDevcontainerDevcontainerFilePresence `json:"devcontainerFilePresence"`
	// failure_message contains the reason the devcontainer failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the devcontainer
	Phase EnvironmentListResponseStatusDevcontainerPhase `json:"phase"`
	// remote_user is the user that is used to connect to the devcontainer
	RemoteUser string `json:"remoteUser"`
	// remote_workspace_folder is the folder that is used to connect to the
	// devcontainer
	RemoteWorkspaceFolder string `json:"remoteWorkspaceFolder"`
	// secrets_in_sync indicates if the secrets are up to date w.r.t. the running
	// devcontainer.
	SecretsInSync bool `json:"secretsInSync"`
	// session is the session that is currently active in the devcontainer.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the devcontainer is present but not
	// in the expected state.
	WarningMessage string                                        `json:"warningMessage"`
	JSON           environmentListResponseStatusDevcontainerJSON `json:"-"`
}

// environmentListResponseStatusDevcontainerJSON contains the JSON metadata for the
// struct [EnvironmentListResponseStatusDevcontainer]
type environmentListResponseStatusDevcontainerJSON struct {
	ContainerID              apijson.Field
	ContainerName            apijson.Field
	DevcontainerconfigInSync apijson.Field
	DevcontainerFilePath     apijson.Field
	DevcontainerFilePresence apijson.Field
	FailureMessage           apijson.Field
	Phase                    apijson.Field
	RemoteUser               apijson.Field
	RemoteWorkspaceFolder    apijson.Field
	SecretsInSync            apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusDevcontainerJSON) RawJSON() string {
	return r.raw
}

// devcontainer_file_presence indicates how the devcontainer file is present in the
// repo.
type EnvironmentListResponseStatusDevcontainerDevcontainerFilePresence string

const (
	EnvironmentListResponseStatusDevcontainerDevcontainerFilePresencePresenceUnspecified EnvironmentListResponseStatusDevcontainerDevcontainerFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentListResponseStatusDevcontainerDevcontainerFilePresencePresenceGenerated   EnvironmentListResponseStatusDevcontainerDevcontainerFilePresence = "PRESENCE_GENERATED"
	EnvironmentListResponseStatusDevcontainerDevcontainerFilePresencePresenceDiscovered  EnvironmentListResponseStatusDevcontainerDevcontainerFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentListResponseStatusDevcontainerDevcontainerFilePresencePresenceSpecified   EnvironmentListResponseStatusDevcontainerDevcontainerFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentListResponseStatusDevcontainerDevcontainerFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentListResponseStatusDevcontainerDevcontainerFilePresencePresenceUnspecified, EnvironmentListResponseStatusDevcontainerDevcontainerFilePresencePresenceGenerated, EnvironmentListResponseStatusDevcontainerDevcontainerFilePresencePresenceDiscovered, EnvironmentListResponseStatusDevcontainerDevcontainerFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the devcontainer
type EnvironmentListResponseStatusDevcontainerPhase string

const (
	EnvironmentListResponseStatusDevcontainerPhasePhaseUnspecified EnvironmentListResponseStatusDevcontainerPhase = "PHASE_UNSPECIFIED"
	EnvironmentListResponseStatusDevcontainerPhasePhaseCreating    EnvironmentListResponseStatusDevcontainerPhase = "PHASE_CREATING"
	EnvironmentListResponseStatusDevcontainerPhasePhaseRunning     EnvironmentListResponseStatusDevcontainerPhase = "PHASE_RUNNING"
	EnvironmentListResponseStatusDevcontainerPhasePhaseStopped     EnvironmentListResponseStatusDevcontainerPhase = "PHASE_STOPPED"
	EnvironmentListResponseStatusDevcontainerPhasePhaseFailed      EnvironmentListResponseStatusDevcontainerPhase = "PHASE_FAILED"
)

func (r EnvironmentListResponseStatusDevcontainerPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseStatusDevcontainerPhasePhaseUnspecified, EnvironmentListResponseStatusDevcontainerPhasePhaseCreating, EnvironmentListResponseStatusDevcontainerPhasePhaseRunning, EnvironmentListResponseStatusDevcontainerPhasePhaseStopped, EnvironmentListResponseStatusDevcontainerPhasePhaseFailed:
		return true
	}
	return false
}

// environment_url contains the URL at which the environment can be accessed. This
// field is only set if the environment is running.
type EnvironmentListResponseStatusEnvironmentURLs struct {
	// logs is the URL at which the environment logs can be accessed.
	Logs  string                                             `json:"logs"`
	Ports []EnvironmentListResponseStatusEnvironmentURLsPort `json:"ports"`
	// SSH is the URL at which the environment can be accessed via SSH.
	SSH  EnvironmentListResponseStatusEnvironmentURLsSSH  `json:"ssh"`
	JSON environmentListResponseStatusEnvironmentURLsJSON `json:"-"`
}

// environmentListResponseStatusEnvironmentURLsJSON contains the JSON metadata for
// the struct [EnvironmentListResponseStatusEnvironmentURLs]
type environmentListResponseStatusEnvironmentURLsJSON struct {
	Logs        apijson.Field
	Ports       apijson.Field
	SSH         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusEnvironmentURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusEnvironmentURLsJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseStatusEnvironmentURLsPort struct {
	// port is the port number of the environment port
	Port int64 `json:"port"`
	// url is the URL at which the environment port can be accessed
	URL  string                                               `json:"url"`
	JSON environmentListResponseStatusEnvironmentURLsPortJSON `json:"-"`
}

// environmentListResponseStatusEnvironmentURLsPortJSON contains the JSON metadata
// for the struct [EnvironmentListResponseStatusEnvironmentURLsPort]
type environmentListResponseStatusEnvironmentURLsPortJSON struct {
	Port        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusEnvironmentURLsPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusEnvironmentURLsPortJSON) RawJSON() string {
	return r.raw
}

// SSH is the URL at which the environment can be accessed via SSH.
type EnvironmentListResponseStatusEnvironmentURLsSSH struct {
	URL  string                                              `json:"url"`
	JSON environmentListResponseStatusEnvironmentURLsSSHJSON `json:"-"`
}

// environmentListResponseStatusEnvironmentURLsSSHJSON contains the JSON metadata
// for the struct [EnvironmentListResponseStatusEnvironmentURLsSSH]
type environmentListResponseStatusEnvironmentURLsSSHJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusEnvironmentURLsSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusEnvironmentURLsSSHJSON) RawJSON() string {
	return r.raw
}

// machine contains the status of the environment machine
type EnvironmentListResponseStatusMachine struct {
	// failure_message contains the reason the machine failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the environment machine
	Phase EnvironmentListResponseStatusMachinePhase `json:"phase"`
	// session is the session that is currently active in the machine.
	Session string `json:"session"`
	// timeout contains the reason the environment has timed out. If this field is
	// empty, the environment has not timed out.
	Timeout string `json:"timeout"`
	// versions contains the versions of components in the machine.
	Versions EnvironmentListResponseStatusMachineVersions `json:"versions"`
	// warning_message contains warnings, e.g. when the machine is present but not in
	// the expected state.
	WarningMessage string                                   `json:"warningMessage"`
	JSON           environmentListResponseStatusMachineJSON `json:"-"`
}

// environmentListResponseStatusMachineJSON contains the JSON metadata for the
// struct [EnvironmentListResponseStatusMachine]
type environmentListResponseStatusMachineJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	Timeout        apijson.Field
	Versions       apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusMachineJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the environment machine
type EnvironmentListResponseStatusMachinePhase string

const (
	EnvironmentListResponseStatusMachinePhasePhaseUnspecified EnvironmentListResponseStatusMachinePhase = "PHASE_UNSPECIFIED"
	EnvironmentListResponseStatusMachinePhasePhaseCreating    EnvironmentListResponseStatusMachinePhase = "PHASE_CREATING"
	EnvironmentListResponseStatusMachinePhasePhaseStarting    EnvironmentListResponseStatusMachinePhase = "PHASE_STARTING"
	EnvironmentListResponseStatusMachinePhasePhaseRunning     EnvironmentListResponseStatusMachinePhase = "PHASE_RUNNING"
	EnvironmentListResponseStatusMachinePhasePhaseStopping    EnvironmentListResponseStatusMachinePhase = "PHASE_STOPPING"
	EnvironmentListResponseStatusMachinePhasePhaseStopped     EnvironmentListResponseStatusMachinePhase = "PHASE_STOPPED"
	EnvironmentListResponseStatusMachinePhasePhaseDeleting    EnvironmentListResponseStatusMachinePhase = "PHASE_DELETING"
	EnvironmentListResponseStatusMachinePhasePhaseDeleted     EnvironmentListResponseStatusMachinePhase = "PHASE_DELETED"
)

func (r EnvironmentListResponseStatusMachinePhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseStatusMachinePhasePhaseUnspecified, EnvironmentListResponseStatusMachinePhasePhaseCreating, EnvironmentListResponseStatusMachinePhasePhaseStarting, EnvironmentListResponseStatusMachinePhasePhaseRunning, EnvironmentListResponseStatusMachinePhasePhaseStopping, EnvironmentListResponseStatusMachinePhasePhaseStopped, EnvironmentListResponseStatusMachinePhasePhaseDeleting, EnvironmentListResponseStatusMachinePhasePhaseDeleted:
		return true
	}
	return false
}

// versions contains the versions of components in the machine.
type EnvironmentListResponseStatusMachineVersions struct {
	SupervisorCommit  string                                           `json:"supervisorCommit"`
	SupervisorVersion string                                           `json:"supervisorVersion"`
	JSON              environmentListResponseStatusMachineVersionsJSON `json:"-"`
}

// environmentListResponseStatusMachineVersionsJSON contains the JSON metadata for
// the struct [EnvironmentListResponseStatusMachineVersions]
type environmentListResponseStatusMachineVersionsJSON struct {
	SupervisorCommit  apijson.Field
	SupervisorVersion apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusMachineVersions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusMachineVersionsJSON) RawJSON() string {
	return r.raw
}

// the phase of an environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentListResponseStatusPhase string

const (
	EnvironmentListResponseStatusPhaseEnvironmentPhaseUnspecified EnvironmentListResponseStatusPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseStatusPhaseEnvironmentPhaseCreating    EnvironmentListResponseStatusPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentListResponseStatusPhaseEnvironmentPhaseStarting    EnvironmentListResponseStatusPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentListResponseStatusPhaseEnvironmentPhaseRunning     EnvironmentListResponseStatusPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentListResponseStatusPhaseEnvironmentPhaseUpdating    EnvironmentListResponseStatusPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentListResponseStatusPhaseEnvironmentPhaseStopping    EnvironmentListResponseStatusPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentListResponseStatusPhaseEnvironmentPhaseStopped     EnvironmentListResponseStatusPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentListResponseStatusPhaseEnvironmentPhaseDeleting    EnvironmentListResponseStatusPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentListResponseStatusPhaseEnvironmentPhaseDeleted     EnvironmentListResponseStatusPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentListResponseStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseStatusPhaseEnvironmentPhaseUnspecified, EnvironmentListResponseStatusPhaseEnvironmentPhaseCreating, EnvironmentListResponseStatusPhaseEnvironmentPhaseStarting, EnvironmentListResponseStatusPhaseEnvironmentPhaseRunning, EnvironmentListResponseStatusPhaseEnvironmentPhaseUpdating, EnvironmentListResponseStatusPhaseEnvironmentPhaseStopping, EnvironmentListResponseStatusPhaseEnvironmentPhaseStopped, EnvironmentListResponseStatusPhaseEnvironmentPhaseDeleting, EnvironmentListResponseStatusPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// RunnerACK is the acknowledgement from the runner that is has received the
// environment spec.
type EnvironmentListResponseStatusRunnerAck struct {
	Message     string                                           `json:"message"`
	SpecVersion string                                           `json:"specVersion"`
	StatusCode  EnvironmentListResponseStatusRunnerAckStatusCode `json:"statusCode"`
	JSON        environmentListResponseStatusRunnerAckJSON       `json:"-"`
}

// environmentListResponseStatusRunnerAckJSON contains the JSON metadata for the
// struct [EnvironmentListResponseStatusRunnerAck]
type environmentListResponseStatusRunnerAckJSON struct {
	Message     apijson.Field
	SpecVersion apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusRunnerAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusRunnerAckJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseStatusRunnerAckStatusCode string

const (
	EnvironmentListResponseStatusRunnerAckStatusCodeStatusCodeUnspecified        EnvironmentListResponseStatusRunnerAckStatusCode = "STATUS_CODE_UNSPECIFIED"
	EnvironmentListResponseStatusRunnerAckStatusCodeStatusCodeOk                 EnvironmentListResponseStatusRunnerAckStatusCode = "STATUS_CODE_OK"
	EnvironmentListResponseStatusRunnerAckStatusCodeStatusCodeInvalidResource    EnvironmentListResponseStatusRunnerAckStatusCode = "STATUS_CODE_INVALID_RESOURCE"
	EnvironmentListResponseStatusRunnerAckStatusCodeStatusCodeFailedPrecondition EnvironmentListResponseStatusRunnerAckStatusCode = "STATUS_CODE_FAILED_PRECONDITION"
)

func (r EnvironmentListResponseStatusRunnerAckStatusCode) IsKnown() bool {
	switch r {
	case EnvironmentListResponseStatusRunnerAckStatusCodeStatusCodeUnspecified, EnvironmentListResponseStatusRunnerAckStatusCodeStatusCodeOk, EnvironmentListResponseStatusRunnerAckStatusCodeStatusCodeInvalidResource, EnvironmentListResponseStatusRunnerAckStatusCodeStatusCodeFailedPrecondition:
		return true
	}
	return false
}

type EnvironmentListResponseStatusSecret struct {
	// failure_message contains the reason the secret failed to be materialize.
	FailureMessage string                                    `json:"failureMessage"`
	Phase          EnvironmentListResponseStatusSecretsPhase `json:"phase"`
	SecretName     string                                    `json:"secretName"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the secret is present but not in
	// the expected state.
	WarningMessage string                                  `json:"warningMessage"`
	JSON           environmentListResponseStatusSecretJSON `json:"-"`
}

// environmentListResponseStatusSecretJSON contains the JSON metadata for the
// struct [EnvironmentListResponseStatusSecret]
type environmentListResponseStatusSecretJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	SecretName     apijson.Field
	Session        apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusSecretJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseStatusSecretsPhase string

const (
	EnvironmentListResponseStatusSecretsPhaseContentPhaseUnspecified  EnvironmentListResponseStatusSecretsPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseStatusSecretsPhaseContentPhaseCreating     EnvironmentListResponseStatusSecretsPhase = "CONTENT_PHASE_CREATING"
	EnvironmentListResponseStatusSecretsPhaseContentPhaseInitializing EnvironmentListResponseStatusSecretsPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentListResponseStatusSecretsPhaseContentPhaseReady        EnvironmentListResponseStatusSecretsPhase = "CONTENT_PHASE_READY"
	EnvironmentListResponseStatusSecretsPhaseContentPhaseUpdating     EnvironmentListResponseStatusSecretsPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentListResponseStatusSecretsPhaseContentPhaseFailed       EnvironmentListResponseStatusSecretsPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentListResponseStatusSecretsPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseStatusSecretsPhaseContentPhaseUnspecified, EnvironmentListResponseStatusSecretsPhaseContentPhaseCreating, EnvironmentListResponseStatusSecretsPhaseContentPhaseInitializing, EnvironmentListResponseStatusSecretsPhaseContentPhaseReady, EnvironmentListResponseStatusSecretsPhaseContentPhaseUpdating, EnvironmentListResponseStatusSecretsPhaseContentPhaseFailed:
		return true
	}
	return false
}

type EnvironmentListResponseStatusSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// phase is the current phase of the public key
	Phase EnvironmentListResponseStatusSSHPublicKeysPhase `json:"phase"`
	JSON  environmentListResponseStatusSSHPublicKeyJSON   `json:"-"`
}

// environmentListResponseStatusSSHPublicKeyJSON contains the JSON metadata for the
// struct [EnvironmentListResponseStatusSSHPublicKey]
type environmentListResponseStatusSSHPublicKeyJSON struct {
	ID          apijson.Field
	Phase       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseStatusSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseStatusSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the public key
type EnvironmentListResponseStatusSSHPublicKeysPhase string

const (
	EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseUnspecified  EnvironmentListResponseStatusSSHPublicKeysPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseCreating     EnvironmentListResponseStatusSSHPublicKeysPhase = "CONTENT_PHASE_CREATING"
	EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseInitializing EnvironmentListResponseStatusSSHPublicKeysPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseReady        EnvironmentListResponseStatusSSHPublicKeysPhase = "CONTENT_PHASE_READY"
	EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseUpdating     EnvironmentListResponseStatusSSHPublicKeysPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseFailed       EnvironmentListResponseStatusSSHPublicKeysPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentListResponseStatusSSHPublicKeysPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseUnspecified, EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseCreating, EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseInitializing, EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseReady, EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseUpdating, EnvironmentListResponseStatusSSHPublicKeysPhaseContentPhaseFailed:
		return true
	}
	return false
}

type EnvironmentDeleteResponse = interface{}

type EnvironmentNewFromProjectResponse struct {
	// +resource get environment
	Environment EnvironmentNewFromProjectResponseEnvironment `json:"environment"`
	JSON        environmentNewFromProjectResponseJSON        `json:"-"`
}

// environmentNewFromProjectResponseJSON contains the JSON metadata for the struct
// [EnvironmentNewFromProjectResponse]
type environmentNewFromProjectResponseJSON struct {
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseJSON) RawJSON() string {
	return r.raw
}

// +resource get environment
type EnvironmentNewFromProjectResponseEnvironment struct {
	// ID is a unique identifier of this environment. No other environment with the
	// same name must be managed by this environment manager
	ID string `json:"id"`
	// EnvironmentMetadata is data associated with an environment that's required for
	// other parts of the system to function
	Metadata EnvironmentNewFromProjectResponseEnvironmentMetadata `json:"metadata"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec EnvironmentNewFromProjectResponseEnvironmentSpec `json:"spec"`
	// EnvironmentStatus describes an environment status
	Status EnvironmentNewFromProjectResponseEnvironmentStatus `json:"status"`
	JSON   environmentNewFromProjectResponseEnvironmentJSON   `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentJSON contains the JSON metadata for
// the struct [EnvironmentNewFromProjectResponseEnvironment]
type environmentNewFromProjectResponseEnvironmentJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentMetadata is data associated with an environment that's required for
// other parts of the system to function
type EnvironmentNewFromProjectResponseEnvironmentMetadata struct {
	// annotations are key/value pairs that gets attached to the environment.
	// +internal - not yet implemented
	Annotations map[string]string `json:"annotations"`
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
	// creator is the identity of the creator of the environment
	Creator EnvironmentNewFromProjectResponseEnvironmentMetadataCreator `json:"creator"`
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
	LastStartedAt time.Time `json:"lastStartedAt" format:"date-time"`
	// name is the name of the environment as specified by the user
	Name string `json:"name"`
	// organization_id is the ID of the organization that contains the environment
	OrganizationID string `json:"organizationId" format:"uuid"`
	// original_context_url is the normalized URL from which the environment was
	// created
	OriginalContextURL string `json:"originalContextUrl"`
	// If the Environment was started from a project, the project_id will reference the
	// project.
	ProjectID string `json:"projectId"`
	// Runner is the ID of the runner that runs this environment.
	RunnerID string                                                   `json:"runnerId"`
	JSON     environmentNewFromProjectResponseEnvironmentMetadataJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentMetadataJSON contains the JSON
// metadata for the struct [EnvironmentNewFromProjectResponseEnvironmentMetadata]
type environmentNewFromProjectResponseEnvironmentMetadataJSON struct {
	Annotations        apijson.Field
	CreatedAt          apijson.Field
	Creator            apijson.Field
	LastStartedAt      apijson.Field
	Name               apijson.Field
	OrganizationID     apijson.Field
	OriginalContextURL apijson.Field
	ProjectID          apijson.Field
	RunnerID           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the creator of the environment
type EnvironmentNewFromProjectResponseEnvironmentMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentNewFromProjectResponseEnvironmentMetadataCreatorJSON      `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentMetadataCreatorJSON contains the
// JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentMetadataCreator]
type environmentNewFromProjectResponseEnvironmentMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipal string

const (
	EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalAccount        EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalUser           EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalRunner         EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalAccount, EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalUser, EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalRunner, EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentNewFromProjectResponseEnvironmentMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type EnvironmentNewFromProjectResponseEnvironmentSpec struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission EnvironmentNewFromProjectResponseEnvironmentSpecAdmission `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile EnvironmentNewFromProjectResponseEnvironmentSpecAutomationsFile `json:"automationsFile"`
	// content is the content spec of the environment
	Content EnvironmentNewFromProjectResponseEnvironmentSpecContent `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer EnvironmentNewFromProjectResponseEnvironmentSpecDevcontainer `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine EnvironmentNewFromProjectResponseEnvironmentSpecMachine `json:"machine"`
	// ports is the set of ports which ought to be exposed to the internet
	Ports []EnvironmentNewFromProjectResponseEnvironmentSpecPort `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets []EnvironmentNewFromProjectResponseEnvironmentSpecSecret `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion string `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys []EnvironmentNewFromProjectResponseEnvironmentSpecSSHPublicKey `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout EnvironmentNewFromProjectResponseEnvironmentSpecTimeout `json:"timeout"`
	JSON    environmentNewFromProjectResponseEnvironmentSpecJSON    `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecJSON contains the JSON metadata
// for the struct [EnvironmentNewFromProjectResponseEnvironmentSpec]
type environmentNewFromProjectResponseEnvironmentSpecJSON struct {
	Admission       apijson.Field
	AutomationsFile apijson.Field
	Content         apijson.Field
	DesiredPhase    apijson.Field
	Devcontainer    apijson.Field
	Machine         apijson.Field
	Ports           apijson.Field
	Secrets         apijson.Field
	SpecVersion     apijson.Field
	SSHPublicKeys   apijson.Field
	Timeout         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentNewFromProjectResponseEnvironmentSpecAdmission string

const (
	EnvironmentNewFromProjectResponseEnvironmentSpecAdmissionAdmissionLevelUnspecified EnvironmentNewFromProjectResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentSpecAdmissionAdmissionLevelOwnerOnly   EnvironmentNewFromProjectResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentNewFromProjectResponseEnvironmentSpecAdmissionAdmissionLevelEveryone    EnvironmentNewFromProjectResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentNewFromProjectResponseEnvironmentSpecAdmission) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentSpecAdmissionAdmissionLevelUnspecified, EnvironmentNewFromProjectResponseEnvironmentSpecAdmissionAdmissionLevelOwnerOnly, EnvironmentNewFromProjectResponseEnvironmentSpecAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// automations_file is the automations file spec of the environment
type EnvironmentNewFromProjectResponseEnvironmentSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath string                                                              `json:"automationsFilePath"`
	Session             string                                                              `json:"session"`
	JSON                environmentNewFromProjectResponseEnvironmentSpecAutomationsFileJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecAutomationsFileJSON contains the
// JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecAutomationsFile]
type environmentNewFromProjectResponseEnvironmentSpecAutomationsFileJSON struct {
	AutomationsFilePath apijson.Field
	Session             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// content is the content spec of the environment
type EnvironmentNewFromProjectResponseEnvironmentSpecContent struct {
	// The Git email address
	GitEmail string `json:"gitEmail"`
	// The Git username
	GitUsername string `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializer `json:"initializer"`
	Session     string                                                             `json:"session"`
	JSON        environmentNewFromProjectResponseEnvironmentSpecContentJSON        `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecContentJSON contains the JSON
// metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecContent]
type environmentNewFromProjectResponseEnvironmentSpecContentJSON struct {
	GitEmail    apijson.Field
	GitUsername apijson.Field
	Initializer apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecContentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentInitializer specifies how an environment is to be initialized
type EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializer struct {
	Specs []EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpec `json:"specs"`
	JSON  environmentNewFromProjectResponseEnvironmentSpecContentInitializerJSON   `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecContentInitializerJSON contains
// the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializer]
type environmentNewFromProjectResponseEnvironmentSpecContentInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecContentInitializerJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpec struct {
	// This field can have the runtime type of
	// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGit].
	Git   interface{}                                                                `json:"git"`
	JSON  environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecJSON `json:"-"`
	union EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnion
}

// environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecJSON
// contains the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpec]
type environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpec) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpec{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURL],
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGit].
func (r EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpec) AsUnion() EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnion {
	return r.union
}

// Union satisfied by
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURL]
// or [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGit].
type EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnion interface {
	implementsEnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGit{}),
		},
	)
}

type EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURL struct {
	ContextURL EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL `json:"contextUrl,required"`
	JSON       environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLJSON       `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLJSON
// contains the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURL]
type environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLJSON struct {
	ContextURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURL) implementsEnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpec() {
}

type EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                                                          `json:"url" format:"uri"`
	JSON environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON
// contains the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL]
type environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsContextURLContextURLJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGit struct {
	Git  EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGit  `json:"git,required"`
	JSON environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitJSON
// contains the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGit]
type environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitJSON struct {
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGit) implementsEnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpec() {
}

type EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                                                            `json:"upstreamRemoteUri"`
	JSON              environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitJSON
// contains the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGit]
type environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode string

const (
	EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

// Phase is the desired phase of the environment
type EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase string

const (
	EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUnspecified EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseCreating    EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStarting    EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseRunning     EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUpdating    EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopping    EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopped     EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleting    EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleted     EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUnspecified, EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseCreating, EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStarting, EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseRunning, EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUpdating, EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopping, EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopped, EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleting, EnvironmentNewFromProjectResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// devcontainer is the devcontainer spec of the environment
type EnvironmentNewFromProjectResponseEnvironmentSpecDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath string                                                           `json:"devcontainerFilePath"`
	Session              string                                                           `json:"session"`
	JSON                 environmentNewFromProjectResponseEnvironmentSpecDevcontainerJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecDevcontainerJSON contains the
// JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecDevcontainer]
type environmentNewFromProjectResponseEnvironmentSpecDevcontainerJSON struct {
	DevcontainerFilePath apijson.Field
	Session              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecDevcontainerJSON) RawJSON() string {
	return r.raw
}

// machine is the machine spec of the environment
type EnvironmentNewFromProjectResponseEnvironmentSpecMachine struct {
	// Class denotes the class of the environment we ought to start
	Class   string                                                      `json:"class" format:"uuid"`
	Session string                                                      `json:"session"`
	JSON    environmentNewFromProjectResponseEnvironmentSpecMachineJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecMachineJSON contains the JSON
// metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecMachine]
type environmentNewFromProjectResponseEnvironmentSpecMachineJSON struct {
	Class       apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecMachineJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewFromProjectResponseEnvironmentSpecPort struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmission `json:"admission"`
	// name of this port
	Name string `json:"name"`
	// port number
	Port int64                                                    `json:"port"`
	JSON environmentNewFromProjectResponseEnvironmentSpecPortJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecPortJSON contains the JSON
// metadata for the struct [EnvironmentNewFromProjectResponseEnvironmentSpecPort]
type environmentNewFromProjectResponseEnvironmentSpecPortJSON struct {
	Admission   apijson.Field
	Name        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecPortJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmission string

const (
	EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmissionAdmissionLevelUnspecified EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmissionAdmissionLevelOwnerOnly   EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmissionAdmissionLevelEveryone    EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmission) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmissionAdmissionLevelUnspecified, EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmissionAdmissionLevelOwnerOnly, EnvironmentNewFromProjectResponseEnvironmentSpecPortsAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

type EnvironmentNewFromProjectResponseEnvironmentSpecSecret struct {
	EnvironmentVariable string `json:"environmentVariable"`
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath          string `json:"filePath"`
	GitCredentialHost string `json:"gitCredentialHost"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                                     `json:"sourceRef"`
	JSON      environmentNewFromProjectResponseEnvironmentSpecSecretJSON `json:"-"`
	union     EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnion
}

// environmentNewFromProjectResponseEnvironmentSpecSecretJSON contains the JSON
// metadata for the struct [EnvironmentNewFromProjectResponseEnvironmentSpecSecret]
type environmentNewFromProjectResponseEnvironmentSpecSecretJSON struct {
	EnvironmentVariable apijson.Field
	FilePath            apijson.Field
	GitCredentialHost   apijson.Field
	Name                apijson.Field
	Session             apijson.Field
	Source              apijson.Field
	SourceRef           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r environmentNewFromProjectResponseEnvironmentSpecSecretJSON) RawJSON() string {
	return r.raw
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecSecret) UnmarshalJSON(data []byte) (err error) {
	*r = EnvironmentNewFromProjectResponseEnvironmentSpecSecret{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsObject],
// [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted],
// [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsObject].
func (r EnvironmentNewFromProjectResponseEnvironmentSpecSecret) AsUnion() EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnion {
	return r.union
}

// Union satisfied by
// [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsObject],
// [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted]
// or [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsObject].
type EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnion interface {
	implementsEnvironmentNewFromProjectResponseEnvironmentSpecSecret()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentNewFromProjectResponseEnvironmentSpecSecretsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentNewFromProjectResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EnvironmentNewFromProjectResponseEnvironmentSpecSecretsObject{}),
		},
	)
}

type EnvironmentNewFromProjectResponseEnvironmentSpecSecretsObject struct {
	EnvironmentVariable string `json:"environmentVariable,required"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                                            `json:"sourceRef"`
	JSON      environmentNewFromProjectResponseEnvironmentSpecSecretsObjectJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecSecretsObjectJSON contains the
// JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsObject]
type environmentNewFromProjectResponseEnvironmentSpecSecretsObjectJSON struct {
	EnvironmentVariable apijson.Field
	Name                apijson.Field
	Session             apijson.Field
	Source              apijson.Field
	SourceRef           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecSecretsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecSecretsObjectJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentNewFromProjectResponseEnvironmentSpecSecretsObject) implementsEnvironmentNewFromProjectResponseEnvironmentSpecSecret() {
}

type EnvironmentNewFromProjectResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted struct {
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath string `json:"filePath,required"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                                                                                                                   `json:"sourceRef"`
	JSON      environmentNewFromProjectResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON
// contains the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted]
type environmentNewFromProjectResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON struct {
	FilePath    apijson.Field
	Name        apijson.Field
	Session     apijson.Field
	Source      apijson.Field
	SourceRef   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMountedJSON) RawJSON() string {
	return r.raw
}

func (r EnvironmentNewFromProjectResponseEnvironmentSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) implementsEnvironmentNewFromProjectResponseEnvironmentSpecSecret() {
}

type EnvironmentNewFromProjectResponseEnvironmentSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// value is the actual public key in the public key file format
	Value string                                                           `json:"value"`
	JSON  environmentNewFromProjectResponseEnvironmentSpecSSHPublicKeyJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecSSHPublicKeyJSON contains the
// JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecSSHPublicKey]
type environmentNewFromProjectResponseEnvironmentSpecSSHPublicKeyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// Timeout configures the environment timeout
type EnvironmentNewFromProjectResponseEnvironmentSpecTimeout struct {
	// A Duration represents a signed, fixed-length span of time represented as a count
	// of seconds and fractions of seconds at nanosecond resolution. It is independent
	// of any calendar and concepts like "day" or "month". It is related to Timestamp
	// in that the difference between two Timestamp values is a Duration and it can be
	// added or subtracted from a Timestamp. Range is approximately +-10,000 years.
	//
	// # Examples
	//
	// Example 1: Compute Duration from two Timestamps in pseudo code.
	//
	//	Timestamp start = ...;
	//	Timestamp end = ...;
	//	Duration duration = ...;
	//
	//	duration.seconds = end.seconds - start.seconds;
	//	duration.nanos = end.nanos - start.nanos;
	//
	//	if (duration.seconds < 0 && duration.nanos > 0) {
	//	  duration.seconds += 1;
	//	  duration.nanos -= 1000000000;
	//	} else if (duration.seconds > 0 && duration.nanos < 0) {
	//	  duration.seconds -= 1;
	//	  duration.nanos += 1000000000;
	//	}
	//
	// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
	//
	//	Timestamp start = ...;
	//	Duration duration = ...;
	//	Timestamp end = ...;
	//
	//	end.seconds = start.seconds + duration.seconds;
	//	end.nanos = start.nanos + duration.nanos;
	//
	//	if (end.nanos < 0) {
	//	  end.seconds -= 1;
	//	  end.nanos += 1000000000;
	//	} else if (end.nanos >= 1000000000) {
	//	  end.seconds += 1;
	//	  end.nanos -= 1000000000;
	//	}
	//
	// Example 3: Compute Duration from datetime.timedelta in Python.
	//
	//	td = datetime.timedelta(days=3, minutes=10)
	//	duration = Duration()
	//	duration.FromTimedelta(td)
	//
	// # JSON Mapping
	//
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by
	// the number of seconds, with nanoseconds expressed as fractional seconds. For
	// example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as
	// "3.000000001s", and 3 seconds and 1 microsecond should be expressed in JSON
	// format as "3.000001s".
	Disconnected string                                                      `json:"disconnected" format:"regex"`
	JSON         environmentNewFromProjectResponseEnvironmentSpecTimeoutJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentSpecTimeoutJSON contains the JSON
// metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentSpecTimeout]
type environmentNewFromProjectResponseEnvironmentSpecTimeoutJSON struct {
	Disconnected apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentSpecTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentSpecTimeoutJSON) RawJSON() string {
	return r.raw
}

// EnvironmentStatus describes an environment status
type EnvironmentNewFromProjectResponseEnvironmentStatus struct {
	// EnvironmentActivitySignal used to signal activity for an environment.
	ActivitySignal EnvironmentNewFromProjectResponseEnvironmentStatusActivitySignal `json:"activitySignal"`
	// automations_file contains the status of the automations file.
	AutomationsFile EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFile `json:"automationsFile"`
	// content contains the status of the environment content.
	Content EnvironmentNewFromProjectResponseEnvironmentStatusContent `json:"content"`
	// devcontainer contains the status of the devcontainer.
	Devcontainer EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainer `json:"devcontainer"`
	// environment_url contains the URL at which the environment can be accessed. This
	// field is only set if the environment is running.
	EnvironmentURLs EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLs `json:"environmentUrls"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage []string `json:"failureMessage"`
	// machine contains the status of the environment machine
	Machine EnvironmentNewFromProjectResponseEnvironmentStatusMachine `json:"machine"`
	// the phase of an environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentNewFromProjectResponseEnvironmentStatusPhase `json:"phase"`
	// RunnerACK is the acknowledgement from the runner that is has received the
	// environment spec.
	RunnerAck EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAck `json:"runnerAck"`
	// secrets contains the status of the environment secrets
	Secrets []EnvironmentNewFromProjectResponseEnvironmentStatusSecret `json:"secrets"`
	// ssh_public_keys contains the status of the environment ssh public keys
	SSHPublicKeys []EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKey `json:"sshPublicKeys"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion"`
	// warning_message contains warnings, e.g. when the environment is present but not
	// in the expected state.
	WarningMessage []string                                               `json:"warningMessage"`
	JSON           environmentNewFromProjectResponseEnvironmentStatusJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusJSON contains the JSON
// metadata for the struct [EnvironmentNewFromProjectResponseEnvironmentStatus]
type environmentNewFromProjectResponseEnvironmentStatusJSON struct {
	ActivitySignal  apijson.Field
	AutomationsFile apijson.Field
	Content         apijson.Field
	Devcontainer    apijson.Field
	EnvironmentURLs apijson.Field
	FailureMessage  apijson.Field
	Machine         apijson.Field
	Phase           apijson.Field
	RunnerAck       apijson.Field
	Secrets         apijson.Field
	SSHPublicKeys   apijson.Field
	StatusVersion   apijson.Field
	WarningMessage  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusJSON) RawJSON() string {
	return r.raw
}

// EnvironmentActivitySignal used to signal activity for an environment.
type EnvironmentNewFromProjectResponseEnvironmentStatusActivitySignal struct {
	// source of the activity signal, such as "VS Code", "SSH", or "Automations". It
	// should be a human-readable string that describes the source of the activity
	// signal.
	Source string `json:"source"`
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
	Timestamp time.Time                                                            `json:"timestamp" format:"date-time"`
	JSON      environmentNewFromProjectResponseEnvironmentStatusActivitySignalJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusActivitySignalJSON contains
// the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusActivitySignal]
type environmentNewFromProjectResponseEnvironmentStatusActivitySignalJSON struct {
	Source      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusActivitySignal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusActivitySignalJSON) RawJSON() string {
	return r.raw
}

// automations_file contains the status of the automations file.
type EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFile struct {
	// automations_file_path is the path to the automations file relative to the repo
	// root.
	AutomationsFilePath string `json:"automationsFilePath"`
	// automations_file_presence indicates how an automations file is present in the
	// environment.
	AutomationsFilePresence EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresence `json:"automationsFilePresence"`
	// failure_message contains the reason the automations file failed to be applied.
	// This is only set if the phase is FAILED.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the automations file.
	Phase EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhase `json:"phase"`
	// session is the automations file session that is currently applied in the
	// environment.
	Session string                                                                `json:"session"`
	JSON    environmentNewFromProjectResponseEnvironmentStatusAutomationsFileJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusAutomationsFileJSON contains
// the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFile]
type environmentNewFromProjectResponseEnvironmentStatusAutomationsFileJSON struct {
	AutomationsFilePath     apijson.Field
	AutomationsFilePresence apijson.Field
	FailureMessage          apijson.Field
	Phase                   apijson.Field
	Session                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// automations_file_presence indicates how an automations file is present in the
// environment.
type EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresence string

const (
	EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceUnspecified EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceAbsent      EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_ABSENT"
	EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceDiscovered  EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceSpecified   EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceUnspecified, EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceAbsent, EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceDiscovered, EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFileAutomationsFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the automations file.
type EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhase string

const (
	EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUnspecified  EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseCreating     EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_CREATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseInitializing EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseReady        EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_READY"
	EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUpdating     EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_UPDATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseFailed       EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhase) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUnspecified, EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseCreating, EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseInitializing, EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseReady, EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseUpdating, EnvironmentNewFromProjectResponseEnvironmentStatusAutomationsFilePhaseContentPhaseFailed:
		return true
	}
	return false
}

// content contains the status of the environment content.
type EnvironmentNewFromProjectResponseEnvironmentStatusContent struct {
	// content_location_in_machine is the location of the content in the machine
	ContentLocationInMachine string `json:"contentLocationInMachine"`
	// failure_message contains the reason the content initialization failed.
	FailureMessage string `json:"failureMessage"`
	// git is the Git working copy status of the environment. Note: this is a
	// best-effort field and more often than not will not be present. Its absence does
	// not indicate the absence of a working copy.
	Git EnvironmentNewFromProjectResponseEnvironmentStatusContentGit `json:"git"`
	// phase is the current phase of the environment content
	Phase EnvironmentNewFromProjectResponseEnvironmentStatusContentPhase `json:"phase"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the content is present but not in
	// the expected state.
	WarningMessage string                                                        `json:"warningMessage"`
	JSON           environmentNewFromProjectResponseEnvironmentStatusContentJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusContentJSON contains the JSON
// metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusContent]
type environmentNewFromProjectResponseEnvironmentStatusContentJSON struct {
	ContentLocationInMachine apijson.Field
	FailureMessage           apijson.Field
	Git                      apijson.Field
	Phase                    apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusContentJSON) RawJSON() string {
	return r.raw
}

// git is the Git working copy status of the environment. Note: this is a
// best-effort field and more often than not will not be present. Its absence does
// not indicate the absence of a working copy.
type EnvironmentNewFromProjectResponseEnvironmentStatusContentGit struct {
	// branch is branch we're currently on
	Branch string `json:"branch"`
	// changed_files is an array of changed files in the environment, possibly
	// truncated
	ChangedFiles []EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFile `json:"changedFiles"`
	// clone_url is the repository url as you would pass it to "git clone". Only HTTPS
	// clone URLs are supported.
	CloneURL string `json:"cloneUrl"`
	// latest_commit is the most recent commit on the current branch
	LatestCommit      string `json:"latestCommit"`
	TotalChangedFiles int64  `json:"totalChangedFiles"`
	// the total number of unpushed changes
	TotalUnpushedCommits int64 `json:"totalUnpushedCommits"`
	// unpushed_commits is an array of unpushed changes in the environment, possibly
	// truncated
	UnpushedCommits []string                                                         `json:"unpushedCommits"`
	JSON            environmentNewFromProjectResponseEnvironmentStatusContentGitJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusContentGitJSON contains the
// JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusContentGit]
type environmentNewFromProjectResponseEnvironmentStatusContentGitJSON struct {
	Branch               apijson.Field
	ChangedFiles         apijson.Field
	CloneURL             apijson.Field
	LatestCommit         apijson.Field
	TotalChangedFiles    apijson.Field
	TotalUnpushedCommits apijson.Field
	UnpushedCommits      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusContentGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusContentGitJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFile struct {
	// ChangeType is the type of change that happened to the file
	ChangeType EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeType `json:"changeType"`
	// path is the path of the file
	Path string                                                                      `json:"path"`
	JSON environmentNewFromProjectResponseEnvironmentStatusContentGitChangedFileJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusContentGitChangedFileJSON
// contains the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFile]
type environmentNewFromProjectResponseEnvironmentStatusContentGitChangedFileJSON struct {
	ChangeType  apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusContentGitChangedFileJSON) RawJSON() string {
	return r.raw
}

// ChangeType is the type of change that happened to the file
type EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeType string

const (
	EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified        EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeAdded              EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_ADDED"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeModified           EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_MODIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeDeleted            EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_DELETED"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeRenamed            EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_RENAMED"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeCopied             EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_COPIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UPDATED_BUT_UNMERGED"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUntracked          EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNTRACKED"
)

func (r EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeType) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified, EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeAdded, EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeModified, EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeDeleted, EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeRenamed, EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeCopied, EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged, EnvironmentNewFromProjectResponseEnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUntracked:
		return true
	}
	return false
}

// phase is the current phase of the environment content
type EnvironmentNewFromProjectResponseEnvironmentStatusContentPhase string

const (
	EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseUnspecified  EnvironmentNewFromProjectResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseCreating     EnvironmentNewFromProjectResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_CREATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseInitializing EnvironmentNewFromProjectResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseReady        EnvironmentNewFromProjectResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_READY"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseUpdating     EnvironmentNewFromProjectResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseFailed       EnvironmentNewFromProjectResponseEnvironmentStatusContentPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentNewFromProjectResponseEnvironmentStatusContentPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseUnspecified, EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseCreating, EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseInitializing, EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseReady, EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseUpdating, EnvironmentNewFromProjectResponseEnvironmentStatusContentPhaseContentPhaseFailed:
		return true
	}
	return false
}

// devcontainer contains the status of the devcontainer.
type EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainer struct {
	// container_id is the ID of the container.
	ContainerID string `json:"containerId"`
	// container_name is the name of the container that is used to connect to the
	// devcontainer
	ContainerName string `json:"containerName"`
	// devcontainerconfig_in_sync indicates if the devcontainer is up to date w.r.t.
	// the devcontainer config file.
	DevcontainerconfigInSync bool `json:"devcontainerconfigInSync"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// devcontainer_file_presence indicates how the devcontainer file is present in the
	// repo.
	DevcontainerFilePresence EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresence `json:"devcontainerFilePresence"`
	// failure_message contains the reason the devcontainer failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the devcontainer
	Phase EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhase `json:"phase"`
	// remote_user is the user that is used to connect to the devcontainer
	RemoteUser string `json:"remoteUser"`
	// remote_workspace_folder is the folder that is used to connect to the
	// devcontainer
	RemoteWorkspaceFolder string `json:"remoteWorkspaceFolder"`
	// secrets_in_sync indicates if the secrets are up to date w.r.t. the running
	// devcontainer.
	SecretsInSync bool `json:"secretsInSync"`
	// session is the session that is currently active in the devcontainer.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the devcontainer is present but not
	// in the expected state.
	WarningMessage string                                                             `json:"warningMessage"`
	JSON           environmentNewFromProjectResponseEnvironmentStatusDevcontainerJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusDevcontainerJSON contains the
// JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainer]
type environmentNewFromProjectResponseEnvironmentStatusDevcontainerJSON struct {
	ContainerID              apijson.Field
	ContainerName            apijson.Field
	DevcontainerconfigInSync apijson.Field
	DevcontainerFilePath     apijson.Field
	DevcontainerFilePresence apijson.Field
	FailureMessage           apijson.Field
	Phase                    apijson.Field
	RemoteUser               apijson.Field
	RemoteWorkspaceFolder    apijson.Field
	SecretsInSync            apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusDevcontainerJSON) RawJSON() string {
	return r.raw
}

// devcontainer_file_presence indicates how the devcontainer file is present in the
// repo.
type EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresence string

const (
	EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceUnspecified EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceGenerated   EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_GENERATED"
	EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceDiscovered  EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceSpecified   EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceUnspecified, EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceGenerated, EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceDiscovered, EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerDevcontainerFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the devcontainer
type EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhase string

const (
	EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhasePhaseUnspecified EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhase = "PHASE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhasePhaseCreating    EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhase = "PHASE_CREATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhasePhaseRunning     EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhase = "PHASE_RUNNING"
	EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhasePhaseStopped     EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhase = "PHASE_STOPPED"
	EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhasePhaseFailed      EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhase = "PHASE_FAILED"
)

func (r EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhasePhaseUnspecified, EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhasePhaseCreating, EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhasePhaseRunning, EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhasePhaseStopped, EnvironmentNewFromProjectResponseEnvironmentStatusDevcontainerPhasePhaseFailed:
		return true
	}
	return false
}

// environment_url contains the URL at which the environment can be accessed. This
// field is only set if the environment is running.
type EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLs struct {
	// logs is the URL at which the environment logs can be accessed.
	Logs  string                                                                  `json:"logs"`
	Ports []EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsPort `json:"ports"`
	// SSH is the URL at which the environment can be accessed via SSH.
	SSH  EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsSSH  `json:"ssh"`
	JSON environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsJSON contains
// the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLs]
type environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsJSON struct {
	Logs        apijson.Field
	Ports       apijson.Field
	SSH         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsPort struct {
	// port is the port number of the environment port
	Port int64 `json:"port"`
	// url is the URL at which the environment port can be accessed
	URL  string                                                                    `json:"url"`
	JSON environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsPortJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsPortJSON
// contains the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsPort]
type environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsPortJSON struct {
	Port        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsPortJSON) RawJSON() string {
	return r.raw
}

// SSH is the URL at which the environment can be accessed via SSH.
type EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsSSH struct {
	URL  string                                                                   `json:"url"`
	JSON environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsSSHJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsSSHJSON
// contains the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsSSH]
type environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsSSHJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusEnvironmentURLsSSHJSON) RawJSON() string {
	return r.raw
}

// machine contains the status of the environment machine
type EnvironmentNewFromProjectResponseEnvironmentStatusMachine struct {
	// failure_message contains the reason the machine failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the environment machine
	Phase EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhase `json:"phase"`
	// session is the session that is currently active in the machine.
	Session string `json:"session"`
	// timeout contains the reason the environment has timed out. If this field is
	// empty, the environment has not timed out.
	Timeout string `json:"timeout"`
	// versions contains the versions of components in the machine.
	Versions EnvironmentNewFromProjectResponseEnvironmentStatusMachineVersions `json:"versions"`
	// warning_message contains warnings, e.g. when the machine is present but not in
	// the expected state.
	WarningMessage string                                                        `json:"warningMessage"`
	JSON           environmentNewFromProjectResponseEnvironmentStatusMachineJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusMachineJSON contains the JSON
// metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusMachine]
type environmentNewFromProjectResponseEnvironmentStatusMachineJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	Timeout        apijson.Field
	Versions       apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusMachineJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the environment machine
type EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhase string

const (
	EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseUnspecified EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhase = "PHASE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseCreating    EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhase = "PHASE_CREATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseStarting    EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhase = "PHASE_STARTING"
	EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseRunning     EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhase = "PHASE_RUNNING"
	EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseStopping    EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhase = "PHASE_STOPPING"
	EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseStopped     EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhase = "PHASE_STOPPED"
	EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseDeleting    EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhase = "PHASE_DELETING"
	EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseDeleted     EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhase = "PHASE_DELETED"
)

func (r EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhase) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseUnspecified, EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseCreating, EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseStarting, EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseRunning, EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseStopping, EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseStopped, EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseDeleting, EnvironmentNewFromProjectResponseEnvironmentStatusMachinePhasePhaseDeleted:
		return true
	}
	return false
}

// versions contains the versions of components in the machine.
type EnvironmentNewFromProjectResponseEnvironmentStatusMachineVersions struct {
	SupervisorCommit  string                                                                `json:"supervisorCommit"`
	SupervisorVersion string                                                                `json:"supervisorVersion"`
	JSON              environmentNewFromProjectResponseEnvironmentStatusMachineVersionsJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusMachineVersionsJSON contains
// the JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusMachineVersions]
type environmentNewFromProjectResponseEnvironmentStatusMachineVersionsJSON struct {
	SupervisorCommit  apijson.Field
	SupervisorVersion apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusMachineVersions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusMachineVersionsJSON) RawJSON() string {
	return r.raw
}

// the phase of an environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentNewFromProjectResponseEnvironmentStatusPhase string

const (
	EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseUnspecified EnvironmentNewFromProjectResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseCreating    EnvironmentNewFromProjectResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseStarting    EnvironmentNewFromProjectResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseRunning     EnvironmentNewFromProjectResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseUpdating    EnvironmentNewFromProjectResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseStopping    EnvironmentNewFromProjectResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseStopped     EnvironmentNewFromProjectResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseDeleting    EnvironmentNewFromProjectResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseDeleted     EnvironmentNewFromProjectResponseEnvironmentStatusPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentNewFromProjectResponseEnvironmentStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseUnspecified, EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseCreating, EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseStarting, EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseRunning, EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseUpdating, EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseStopping, EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseStopped, EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseDeleting, EnvironmentNewFromProjectResponseEnvironmentStatusPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// RunnerACK is the acknowledgement from the runner that is has received the
// environment spec.
type EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAck struct {
	Message     string                                                                `json:"message"`
	SpecVersion string                                                                `json:"specVersion"`
	StatusCode  EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCode `json:"statusCode"`
	JSON        environmentNewFromProjectResponseEnvironmentStatusRunnerAckJSON       `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusRunnerAckJSON contains the
// JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAck]
type environmentNewFromProjectResponseEnvironmentStatusRunnerAckJSON struct {
	Message     apijson.Field
	SpecVersion apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusRunnerAckJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCode string

const (
	EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeUnspecified        EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeOk                 EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_OK"
	EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeInvalidResource    EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_INVALID_RESOURCE"
	EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeFailedPrecondition EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_FAILED_PRECONDITION"
)

func (r EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCode) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeUnspecified, EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeOk, EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeInvalidResource, EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCodeStatusCodeFailedPrecondition:
		return true
	}
	return false
}

type EnvironmentNewFromProjectResponseEnvironmentStatusSecret struct {
	// failure_message contains the reason the secret failed to be materialize.
	FailureMessage string                                                         `json:"failureMessage"`
	Phase          EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhase `json:"phase"`
	SecretName     string                                                         `json:"secretName"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the secret is present but not in
	// the expected state.
	WarningMessage string                                                       `json:"warningMessage"`
	JSON           environmentNewFromProjectResponseEnvironmentStatusSecretJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusSecretJSON contains the JSON
// metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusSecret]
type environmentNewFromProjectResponseEnvironmentStatusSecretJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	SecretName     apijson.Field
	Session        apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusSecretJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhase string

const (
	EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseUnspecified  EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseCreating     EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_CREATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseInitializing EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseReady        EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_READY"
	EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseUpdating     EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseFailed       EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseUnspecified, EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseCreating, EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseInitializing, EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseReady, EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseUpdating, EnvironmentNewFromProjectResponseEnvironmentStatusSecretsPhaseContentPhaseFailed:
		return true
	}
	return false
}

type EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// phase is the current phase of the public key
	Phase EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhase `json:"phase"`
	JSON  environmentNewFromProjectResponseEnvironmentStatusSSHPublicKeyJSON   `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusSSHPublicKeyJSON contains the
// JSON metadata for the struct
// [EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKey]
type environmentNewFromProjectResponseEnvironmentStatusSSHPublicKeyJSON struct {
	ID          apijson.Field
	Phase       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseEnvironmentStatusSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the public key
type EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhase string

const (
	EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUnspecified  EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseCreating     EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_CREATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseInitializing EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseReady        EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_READY"
	EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUpdating     EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseFailed       EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUnspecified, EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseCreating, EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseInitializing, EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseReady, EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseUpdating, EnvironmentNewFromProjectResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseFailed:
		return true
	}
	return false
}

type EnvironmentNewLogsTokenResponse struct {
	// access_token is the token that can be used to access the logs of the environment
	AccessToken string                              `json:"accessToken"`
	JSON        environmentNewLogsTokenResponseJSON `json:"-"`
}

// environmentNewLogsTokenResponseJSON contains the JSON metadata for the struct
// [EnvironmentNewLogsTokenResponse]
type environmentNewLogsTokenResponseJSON struct {
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewLogsTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewLogsTokenResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentMarkActiveResponse = interface{}

type EnvironmentStartResponse = interface{}

type EnvironmentStopResponse = interface{}

type EnvironmentNewParams struct {
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec param.Field[EnvironmentNewParamsSpec] `json:"spec"`
}

func (r EnvironmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type EnvironmentNewParamsSpec struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission param.Field[EnvironmentNewParamsSpecAdmission] `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile param.Field[EnvironmentNewParamsSpecAutomationsFile] `json:"automationsFile"`
	// content is the content spec of the environment
	Content param.Field[EnvironmentNewParamsSpecContent] `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase param.Field[EnvironmentNewParamsSpecDesiredPhase] `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer param.Field[EnvironmentNewParamsSpecDevcontainer] `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine param.Field[EnvironmentNewParamsSpecMachine] `json:"machine"`
	// ports is the set of ports which ought to be exposed to the internet
	Ports param.Field[[]EnvironmentNewParamsSpecPort] `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets param.Field[[]EnvironmentNewParamsSpecSecretUnion] `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion param.Field[string] `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys param.Field[[]EnvironmentNewParamsSpecSSHPublicKey] `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout param.Field[EnvironmentNewParamsSpecTimeout] `json:"timeout"`
}

func (r EnvironmentNewParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentNewParamsSpecAdmission string

const (
	EnvironmentNewParamsSpecAdmissionAdmissionLevelUnspecified EnvironmentNewParamsSpecAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentNewParamsSpecAdmissionAdmissionLevelOwnerOnly   EnvironmentNewParamsSpecAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentNewParamsSpecAdmissionAdmissionLevelEveryone    EnvironmentNewParamsSpecAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentNewParamsSpecAdmission) IsKnown() bool {
	switch r {
	case EnvironmentNewParamsSpecAdmissionAdmissionLevelUnspecified, EnvironmentNewParamsSpecAdmissionAdmissionLevelOwnerOnly, EnvironmentNewParamsSpecAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// automations_file is the automations file spec of the environment
type EnvironmentNewParamsSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath"`
	Session             param.Field[string] `json:"session"`
}

func (r EnvironmentNewParamsSpecAutomationsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// content is the content spec of the environment
type EnvironmentNewParamsSpecContent struct {
	// The Git email address
	GitEmail param.Field[string] `json:"gitEmail"`
	// The Git username
	GitUsername param.Field[string] `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer param.Field[EnvironmentNewParamsSpecContentInitializer] `json:"initializer"`
	Session     param.Field[string]                                     `json:"session"`
}

func (r EnvironmentNewParamsSpecContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnvironmentInitializer specifies how an environment is to be initialized
type EnvironmentNewParamsSpecContentInitializer struct {
	Specs param.Field[[]EnvironmentNewParamsSpecContentInitializerSpecUnion] `json:"specs"`
}

func (r EnvironmentNewParamsSpecContentInitializer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewParamsSpecContentInitializerSpec struct {
	ContextURL param.Field[interface{}] `json:"contextUrl"`
	Git        param.Field[interface{}] `json:"git"`
}

func (r EnvironmentNewParamsSpecContentInitializerSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewParamsSpecContentInitializerSpec) implementsEnvironmentNewParamsSpecContentInitializerSpecUnion() {
}

// Satisfied by [EnvironmentNewParamsSpecContentInitializerSpecsContextURL],
// [EnvironmentNewParamsSpecContentInitializerSpecsGit],
// [EnvironmentNewParamsSpecContentInitializerSpec].
type EnvironmentNewParamsSpecContentInitializerSpecUnion interface {
	implementsEnvironmentNewParamsSpecContentInitializerSpecUnion()
}

type EnvironmentNewParamsSpecContentInitializerSpecsContextURL struct {
	ContextURL param.Field[EnvironmentNewParamsSpecContentInitializerSpecsContextURLContextURL] `json:"contextUrl,required"`
}

func (r EnvironmentNewParamsSpecContentInitializerSpecsContextURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewParamsSpecContentInitializerSpecsContextURL) implementsEnvironmentNewParamsSpecContentInitializerSpecUnion() {
}

type EnvironmentNewParamsSpecContentInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r EnvironmentNewParamsSpecContentInitializerSpecsContextURLContextURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewParamsSpecContentInitializerSpecsGit struct {
	Git param.Field[EnvironmentNewParamsSpecContentInitializerSpecsGitGit] `json:"git,required"`
}

func (r EnvironmentNewParamsSpecContentInitializerSpecsGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewParamsSpecContentInitializerSpecsGit) implementsEnvironmentNewParamsSpecContentInitializerSpecUnion() {
}

type EnvironmentNewParamsSpecContentInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation param.Field[string] `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget param.Field[string] `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri param.Field[string] `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode param.Field[EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetMode] `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri param.Field[string] `json:"upstreamRemoteUri"`
}

func (r EnvironmentNewParamsSpecContentInitializerSpecsGitGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetMode string

const (
	EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, EnvironmentNewParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

// Phase is the desired phase of the environment
type EnvironmentNewParamsSpecDesiredPhase string

const (
	EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseUnspecified EnvironmentNewParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseCreating    EnvironmentNewParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseStarting    EnvironmentNewParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseRunning     EnvironmentNewParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseUpdating    EnvironmentNewParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseStopping    EnvironmentNewParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseStopped     EnvironmentNewParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseDeleting    EnvironmentNewParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseDeleted     EnvironmentNewParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentNewParamsSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseUnspecified, EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseCreating, EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseStarting, EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseRunning, EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseUpdating, EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseStopping, EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseStopped, EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseDeleting, EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// devcontainer is the devcontainer spec of the environment
type EnvironmentNewParamsSpecDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath"`
	Session              param.Field[string] `json:"session"`
}

func (r EnvironmentNewParamsSpecDevcontainer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// machine is the machine spec of the environment
type EnvironmentNewParamsSpecMachine struct {
	// Class denotes the class of the environment we ought to start
	Class   param.Field[string] `json:"class" format:"uuid"`
	Session param.Field[string] `json:"session"`
}

func (r EnvironmentNewParamsSpecMachine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewParamsSpecPort struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission param.Field[EnvironmentNewParamsSpecPortsAdmission] `json:"admission"`
	// name of this port
	Name param.Field[string] `json:"name"`
	// port number
	Port param.Field[int64] `json:"port"`
}

func (r EnvironmentNewParamsSpecPort) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentNewParamsSpecPortsAdmission string

const (
	EnvironmentNewParamsSpecPortsAdmissionAdmissionLevelUnspecified EnvironmentNewParamsSpecPortsAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentNewParamsSpecPortsAdmissionAdmissionLevelOwnerOnly   EnvironmentNewParamsSpecPortsAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentNewParamsSpecPortsAdmissionAdmissionLevelEveryone    EnvironmentNewParamsSpecPortsAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentNewParamsSpecPortsAdmission) IsKnown() bool {
	switch r {
	case EnvironmentNewParamsSpecPortsAdmissionAdmissionLevelUnspecified, EnvironmentNewParamsSpecPortsAdmissionAdmissionLevelOwnerOnly, EnvironmentNewParamsSpecPortsAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

type EnvironmentNewParamsSpecSecret struct {
	EnvironmentVariable param.Field[string] `json:"environmentVariable"`
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath          param.Field[string] `json:"filePath"`
	GitCredentialHost param.Field[string] `json:"gitCredentialHost"`
	// name is the human readable description of the secret
	Name param.Field[string] `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session param.Field[string] `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source param.Field[string] `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef param.Field[string] `json:"sourceRef"`
}

func (r EnvironmentNewParamsSpecSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewParamsSpecSecret) implementsEnvironmentNewParamsSpecSecretUnion() {}

// Satisfied by [EnvironmentNewParamsSpecSecretsObject],
// [EnvironmentNewParamsSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted],
// [EnvironmentNewParamsSpecSecretsObject], [EnvironmentNewParamsSpecSecret].
type EnvironmentNewParamsSpecSecretUnion interface {
	implementsEnvironmentNewParamsSpecSecretUnion()
}

type EnvironmentNewParamsSpecSecretsObject struct {
	EnvironmentVariable param.Field[string] `json:"environmentVariable,required"`
	// name is the human readable description of the secret
	Name param.Field[string] `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session param.Field[string] `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source param.Field[string] `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef param.Field[string] `json:"sourceRef"`
}

func (r EnvironmentNewParamsSpecSecretsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewParamsSpecSecretsObject) implementsEnvironmentNewParamsSpecSecretUnion() {}

type EnvironmentNewParamsSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted struct {
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath param.Field[string] `json:"filePath,required"`
	// name is the human readable description of the secret
	Name param.Field[string] `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session param.Field[string] `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source param.Field[string] `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef param.Field[string] `json:"sourceRef"`
}

func (r EnvironmentNewParamsSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewParamsSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) implementsEnvironmentNewParamsSpecSecretUnion() {
}

type EnvironmentNewParamsSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID param.Field[string] `json:"id"`
	// value is the actual public key in the public key file format
	Value param.Field[string] `json:"value"`
}

func (r EnvironmentNewParamsSpecSSHPublicKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Timeout configures the environment timeout
type EnvironmentNewParamsSpecTimeout struct {
	// A Duration represents a signed, fixed-length span of time represented as a count
	// of seconds and fractions of seconds at nanosecond resolution. It is independent
	// of any calendar and concepts like "day" or "month". It is related to Timestamp
	// in that the difference between two Timestamp values is a Duration and it can be
	// added or subtracted from a Timestamp. Range is approximately +-10,000 years.
	//
	// # Examples
	//
	// Example 1: Compute Duration from two Timestamps in pseudo code.
	//
	//	Timestamp start = ...;
	//	Timestamp end = ...;
	//	Duration duration = ...;
	//
	//	duration.seconds = end.seconds - start.seconds;
	//	duration.nanos = end.nanos - start.nanos;
	//
	//	if (duration.seconds < 0 && duration.nanos > 0) {
	//	  duration.seconds += 1;
	//	  duration.nanos -= 1000000000;
	//	} else if (duration.seconds > 0 && duration.nanos < 0) {
	//	  duration.seconds -= 1;
	//	  duration.nanos += 1000000000;
	//	}
	//
	// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
	//
	//	Timestamp start = ...;
	//	Duration duration = ...;
	//	Timestamp end = ...;
	//
	//	end.seconds = start.seconds + duration.seconds;
	//	end.nanos = start.nanos + duration.nanos;
	//
	//	if (end.nanos < 0) {
	//	  end.seconds -= 1;
	//	  end.nanos += 1000000000;
	//	} else if (end.nanos >= 1000000000) {
	//	  end.seconds += 1;
	//	  end.nanos -= 1000000000;
	//	}
	//
	// Example 3: Compute Duration from datetime.timedelta in Python.
	//
	//	td = datetime.timedelta(days=3, minutes=10)
	//	duration = Duration()
	//	duration.FromTimedelta(td)
	//
	// # JSON Mapping
	//
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by
	// the number of seconds, with nanoseconds expressed as fractional seconds. For
	// example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as
	// "3.000000001s", and 3 seconds and 1 microsecond should be expressed in JSON
	// format as "3.000001s".
	Disconnected param.Field[string] `json:"disconnected" format:"regex"`
}

func (r EnvironmentNewParamsSpecTimeout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentGetParams struct {
	// environment_id specifies the environment to get
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParams struct {
	Body EnvironmentUpdateParamsBodyUnion `json:"body,required"`
}

func (r EnvironmentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type EnvironmentUpdateParamsBody struct {
	Metadata param.Field[interface{}] `json:"metadata"`
	Spec     param.Field[interface{}] `json:"spec"`
}

func (r EnvironmentUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBody) implementsEnvironmentUpdateParamsBodyUnion() {}

// Satisfied by [EnvironmentUpdateParamsBodyMetadata],
// [EnvironmentUpdateParamsBodySpec], [EnvironmentUpdateParamsBody].
type EnvironmentUpdateParamsBodyUnion interface {
	implementsEnvironmentUpdateParamsBodyUnion()
}

type EnvironmentUpdateParamsBodyMetadata struct {
	Metadata param.Field[interface{}] `json:"metadata,required"`
}

func (r EnvironmentUpdateParamsBodyMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodyMetadata) implementsEnvironmentUpdateParamsBodyUnion() {}

type EnvironmentUpdateParamsBodySpec struct {
	Spec param.Field[EnvironmentUpdateParamsBodySpecSpecUnion] `json:"spec,required"`
}

func (r EnvironmentUpdateParamsBodySpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpec) implementsEnvironmentUpdateParamsBodyUnion() {}

type EnvironmentUpdateParamsBodySpecSpec struct {
	AutomationsFile param.Field[interface{}] `json:"automationsFile"`
	Content         param.Field[interface{}] `json:"content"`
	Devcontainer    param.Field[interface{}] `json:"devcontainer"`
	Timeout         param.Field[interface{}] `json:"timeout"`
}

func (r EnvironmentUpdateParamsBodySpecSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpec) implementsEnvironmentUpdateParamsBodySpecSpecUnion() {}

// Satisfied by
// [EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironment],
// [EnvironmentUpdateParamsBodySpecSpecContent],
// [EnvironmentUpdateParamsBodySpecSpecDevcontainer],
// [EnvironmentUpdateParamsBodySpecSpecTimeoutConfiguresTheEnvironmentTimeout],
// [EnvironmentUpdateParamsBodySpecSpec].
type EnvironmentUpdateParamsBodySpecSpecUnion interface {
	implementsEnvironmentUpdateParamsBodySpecSpecUnion()
}

type EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironment struct {
	// automations_file is the automations file spec of the environment
	AutomationsFile param.Field[EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileUnion] `json:"automationsFile,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironment) implementsEnvironmentUpdateParamsBodySpecSpecUnion() {
}

// automations_file is the automations file spec of the environment
type EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath"`
	Session             param.Field[string] `json:"session"`
}

func (r EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFile) implementsEnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileUnion() {
}

// automations_file is the automations file spec of the environment
//
// Satisfied by
// [EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileAutomationsFilePathIsThePathToTheAutomationsFileThatIsAppliedInTheEnvironmentRelativeToTheRepoRoot],
// [EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileSession],
// [EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFile].
type EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileUnion interface {
	implementsEnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileUnion()
}

type EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileAutomationsFilePathIsThePathToTheAutomationsFileThatIsAppliedInTheEnvironmentRelativeToTheRepoRoot struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileAutomationsFilePathIsThePathToTheAutomationsFileThatIsAppliedInTheEnvironmentRelativeToTheRepoRoot) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileAutomationsFilePathIsThePathToTheAutomationsFileThatIsAppliedInTheEnvironmentRelativeToTheRepoRoot) implementsEnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileUnion() {
}

type EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileSession struct {
	Session param.Field[string] `json:"session,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileSession) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileSession) implementsEnvironmentUpdateParamsBodySpecSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileUnion() {
}

type EnvironmentUpdateParamsBodySpecSpecContent struct {
	Content param.Field[EnvironmentUpdateParamsBodySpecSpecContentContentUnion] `json:"content,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecContent) implementsEnvironmentUpdateParamsBodySpecSpecUnion() {
}

type EnvironmentUpdateParamsBodySpecSpecContentContent struct {
	// The Git email address
	GitEmail param.Field[string] `json:"gitEmail"`
	// The Git username
	GitUsername param.Field[string]      `json:"gitUsername"`
	Initializer param.Field[interface{}] `json:"initializer"`
	// session should be changed to trigger a content reinitialization
	Session param.Field[string] `json:"session"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContent) implementsEnvironmentUpdateParamsBodySpecSpecContentContentUnion() {
}

// Satisfied by
// [EnvironmentUpdateParamsBodySpecSpecContentContentTheGitEmailAddress],
// [EnvironmentUpdateParamsBodySpecSpecContentContentTheGitUsername],
// [EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitialized],
// [EnvironmentUpdateParamsBodySpecSpecContentContentSessionShouldBeChangedToTriggerAContentReinitialization],
// [EnvironmentUpdateParamsBodySpecSpecContentContent].
type EnvironmentUpdateParamsBodySpecSpecContentContentUnion interface {
	implementsEnvironmentUpdateParamsBodySpecSpecContentContentUnion()
}

type EnvironmentUpdateParamsBodySpecSpecContentContentTheGitEmailAddress struct {
	// The Git email address
	GitEmail param.Field[string] `json:"gitEmail,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentTheGitEmailAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentTheGitEmailAddress) implementsEnvironmentUpdateParamsBodySpecSpecContentContentUnion() {
}

type EnvironmentUpdateParamsBodySpecSpecContentContentTheGitUsername struct {
	// The Git username
	GitUsername param.Field[string] `json:"gitUsername,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentTheGitUsername) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentTheGitUsername) implementsEnvironmentUpdateParamsBodySpecSpecContentContentUnion() {
}

type EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitialized struct {
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer param.Field[EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializer] `json:"initializer,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitialized) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitialized) implementsEnvironmentUpdateParamsBodySpecSpecContentContentUnion() {
}

// EnvironmentInitializer specifies how an environment is to be initialized
type EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializer struct {
	Specs param.Field[[]EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecUnion] `json:"specs"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpec struct {
	ContextURL param.Field[interface{}] `json:"contextUrl"`
	Git        param.Field[interface{}] `json:"git"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpec) implementsEnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecUnion() {
}

// Satisfied by
// [EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsContextURL],
// [EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGit],
// [EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpec].
type EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecUnion interface {
	implementsEnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecUnion()
}

type EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsContextURL struct {
	ContextURL param.Field[EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsContextURLContextURL] `json:"contextUrl,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsContextURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsContextURL) implementsEnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecUnion() {
}

type EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsContextURLContextURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGit struct {
	Git param.Field[EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGit] `json:"git,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGit) implementsEnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecUnion() {
}

type EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation param.Field[string] `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget param.Field[string] `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri param.Field[string] `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode param.Field[EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetMode] `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri param.Field[string] `json:"upstreamRemoteUri"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetMode string

const (
	EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, EnvironmentUpdateParamsBodySpecSpecContentContentInitializerConfiguresHowTheEnvironmentIsToBeInitializedInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type EnvironmentUpdateParamsBodySpecSpecContentContentSessionShouldBeChangedToTriggerAContentReinitialization struct {
	// session should be changed to trigger a content reinitialization
	Session param.Field[string] `json:"session,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentSessionShouldBeChangedToTriggerAContentReinitialization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecContentContentSessionShouldBeChangedToTriggerAContentReinitialization) implementsEnvironmentUpdateParamsBodySpecSpecContentContentUnion() {
}

type EnvironmentUpdateParamsBodySpecSpecDevcontainer struct {
	Devcontainer param.Field[EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerUnion] `json:"devcontainer,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecDevcontainer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecDevcontainer) implementsEnvironmentUpdateParamsBodySpecSpecUnion() {
}

type EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath"`
	// session should be changed to trigger a devcontainer rebuild
	Session param.Field[string] `json:"session"`
}

func (r EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainer) implementsEnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerUnion() {
}

// Satisfied by
// [EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerDevcontainerFilePathIsThePathToTheDevcontainerFileRelativeToTheRepoRoot],
// [EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerSessionShouldBeChangedToTriggerADevcontainerRebuild],
// [EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainer].
type EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerUnion interface {
	implementsEnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerUnion()
}

type EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerDevcontainerFilePathIsThePathToTheDevcontainerFileRelativeToTheRepoRoot struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerDevcontainerFilePathIsThePathToTheDevcontainerFileRelativeToTheRepoRoot) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerDevcontainerFilePathIsThePathToTheDevcontainerFileRelativeToTheRepoRoot) implementsEnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerUnion() {
}

type EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerSessionShouldBeChangedToTriggerADevcontainerRebuild struct {
	// session should be changed to trigger a devcontainer rebuild
	Session param.Field[string] `json:"session,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerSessionShouldBeChangedToTriggerADevcontainerRebuild) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerSessionShouldBeChangedToTriggerADevcontainerRebuild) implementsEnvironmentUpdateParamsBodySpecSpecDevcontainerDevcontainerUnion() {
}

type EnvironmentUpdateParamsBodySpecSpecTimeoutConfiguresTheEnvironmentTimeout struct {
	// Timeout configures the environment timeout
	Timeout param.Field[EnvironmentUpdateParamsBodySpecSpecTimeoutConfiguresTheEnvironmentTimeoutTimeout] `json:"timeout,required"`
}

func (r EnvironmentUpdateParamsBodySpecSpecTimeoutConfiguresTheEnvironmentTimeout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentUpdateParamsBodySpecSpecTimeoutConfiguresTheEnvironmentTimeout) implementsEnvironmentUpdateParamsBodySpecSpecUnion() {
}

// Timeout configures the environment timeout
type EnvironmentUpdateParamsBodySpecSpecTimeoutConfiguresTheEnvironmentTimeoutTimeout struct {
	// A Duration represents a signed, fixed-length span of time represented as a count
	// of seconds and fractions of seconds at nanosecond resolution. It is independent
	// of any calendar and concepts like "day" or "month". It is related to Timestamp
	// in that the difference between two Timestamp values is a Duration and it can be
	// added or subtracted from a Timestamp. Range is approximately +-10,000 years.
	//
	// # Examples
	//
	// Example 1: Compute Duration from two Timestamps in pseudo code.
	//
	//	Timestamp start = ...;
	//	Timestamp end = ...;
	//	Duration duration = ...;
	//
	//	duration.seconds = end.seconds - start.seconds;
	//	duration.nanos = end.nanos - start.nanos;
	//
	//	if (duration.seconds < 0 && duration.nanos > 0) {
	//	  duration.seconds += 1;
	//	  duration.nanos -= 1000000000;
	//	} else if (duration.seconds > 0 && duration.nanos < 0) {
	//	  duration.seconds -= 1;
	//	  duration.nanos += 1000000000;
	//	}
	//
	// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
	//
	//	Timestamp start = ...;
	//	Duration duration = ...;
	//	Timestamp end = ...;
	//
	//	end.seconds = start.seconds + duration.seconds;
	//	end.nanos = start.nanos + duration.nanos;
	//
	//	if (end.nanos < 0) {
	//	  end.seconds -= 1;
	//	  end.nanos += 1000000000;
	//	} else if (end.nanos >= 1000000000) {
	//	  end.seconds += 1;
	//	  end.nanos -= 1000000000;
	//	}
	//
	// Example 3: Compute Duration from datetime.timedelta in Python.
	//
	//	td = datetime.timedelta(days=3, minutes=10)
	//	duration = Duration()
	//	duration.FromTimedelta(td)
	//
	// # JSON Mapping
	//
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by
	// the number of seconds, with nanoseconds expressed as fractional seconds. For
	// example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as
	// "3.000000001s", and 3 seconds and 1 microsecond should be expressed in JSON
	// format as "3.000001s".
	Disconnected param.Field[string] `json:"disconnected,required" format:"regex"`
}

func (r EnvironmentUpdateParamsBodySpecSpecTimeoutConfiguresTheEnvironmentTimeoutTimeout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentListParams struct {
	Token    param.Field[string]                      `query:"token"`
	PageSize param.Field[int64]                       `query:"pageSize"`
	Filter   param.Field[EnvironmentListParamsFilter] `json:"filter"`
	// organization_id is the ID of the organization that contains the environments
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
	// pagination contains the pagination options for listing environments
	Pagination param.Field[EnvironmentListParamsPagination] `json:"pagination"`
}

func (r EnvironmentListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [EnvironmentListParams]'s query parameters as `url.Values`.
func (r EnvironmentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EnvironmentListParamsFilter struct {
	// creator_ids filters the response to only Environments created by specified
	// members
	CreatorIDs param.Field[[]string] `json:"creatorIds" format:"uuid"`
	// project_ids filters the response to only Environments associated with the
	// specified projects
	ProjectIDs param.Field[[]string] `json:"projectIds" format:"uuid"`
	// runner_ids filters the response to only Environments running on these Runner IDs
	RunnerIDs param.Field[[]string] `json:"runnerIds" format:"uuid"`
	// runner_kinds filters the response to only Environments running on these Runner
	// Kinds
	RunnerKinds param.Field[[]EnvironmentListParamsFilterRunnerKind] `json:"runnerKinds"`
	// actual_phases is a list of phases the environment must be in for it to be
	// returned in the API call
	StatusPhases param.Field[[]EnvironmentListParamsFilterStatusPhase] `json:"statusPhases"`
}

func (r EnvironmentListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// RunnerKind represents the kind of a runner
type EnvironmentListParamsFilterRunnerKind string

const (
	EnvironmentListParamsFilterRunnerKindRunnerKindUnspecified        EnvironmentListParamsFilterRunnerKind = "RUNNER_KIND_UNSPECIFIED"
	EnvironmentListParamsFilterRunnerKindRunnerKindLocal              EnvironmentListParamsFilterRunnerKind = "RUNNER_KIND_LOCAL"
	EnvironmentListParamsFilterRunnerKindRunnerKindRemote             EnvironmentListParamsFilterRunnerKind = "RUNNER_KIND_REMOTE"
	EnvironmentListParamsFilterRunnerKindRunnerKindLocalConfiguration EnvironmentListParamsFilterRunnerKind = "RUNNER_KIND_LOCAL_CONFIGURATION"
)

func (r EnvironmentListParamsFilterRunnerKind) IsKnown() bool {
	switch r {
	case EnvironmentListParamsFilterRunnerKindRunnerKindUnspecified, EnvironmentListParamsFilterRunnerKindRunnerKindLocal, EnvironmentListParamsFilterRunnerKindRunnerKindRemote, EnvironmentListParamsFilterRunnerKindRunnerKindLocalConfiguration:
		return true
	}
	return false
}

type EnvironmentListParamsFilterStatusPhase string

const (
	EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseUnspecified EnvironmentListParamsFilterStatusPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseCreating    EnvironmentListParamsFilterStatusPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseStarting    EnvironmentListParamsFilterStatusPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseRunning     EnvironmentListParamsFilterStatusPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseUpdating    EnvironmentListParamsFilterStatusPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseStopping    EnvironmentListParamsFilterStatusPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseStopped     EnvironmentListParamsFilterStatusPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseDeleting    EnvironmentListParamsFilterStatusPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseDeleted     EnvironmentListParamsFilterStatusPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentListParamsFilterStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseUnspecified, EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseCreating, EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseStarting, EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseRunning, EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseUpdating, EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseStopping, EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseStopped, EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseDeleting, EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// pagination contains the pagination options for listing environments
type EnvironmentListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EnvironmentListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentDeleteParams struct {
	// environment_id specifies the environment that is going to delete.
	//
	// +required
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
	// force indicates whether the environment should be deleted forcefully When force
	// deleting an Environment, the Environment is removed immediately and environment
	// lifecycle is not respected. Force deleting can result in data loss on the
	// environment.
	Force param.Field[bool] `json:"force"`
}

func (r EnvironmentDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewFromProjectParams struct {
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec param.Field[EnvironmentNewFromProjectParamsSpec] `json:"spec"`
}

func (r EnvironmentNewFromProjectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type EnvironmentNewFromProjectParamsSpec struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission param.Field[EnvironmentNewFromProjectParamsSpecAdmission] `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile param.Field[EnvironmentNewFromProjectParamsSpecAutomationsFile] `json:"automationsFile"`
	// content is the content spec of the environment
	Content param.Field[EnvironmentNewFromProjectParamsSpecContent] `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase param.Field[EnvironmentNewFromProjectParamsSpecDesiredPhase] `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer param.Field[EnvironmentNewFromProjectParamsSpecDevcontainer] `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine param.Field[EnvironmentNewFromProjectParamsSpecMachine] `json:"machine"`
	// ports is the set of ports which ought to be exposed to the internet
	Ports param.Field[[]EnvironmentNewFromProjectParamsSpecPort] `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets param.Field[[]EnvironmentNewFromProjectParamsSpecSecretUnion] `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion param.Field[string] `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys param.Field[[]EnvironmentNewFromProjectParamsSpecSSHPublicKey] `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout param.Field[EnvironmentNewFromProjectParamsSpecTimeout] `json:"timeout"`
}

func (r EnvironmentNewFromProjectParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentNewFromProjectParamsSpecAdmission string

const (
	EnvironmentNewFromProjectParamsSpecAdmissionAdmissionLevelUnspecified EnvironmentNewFromProjectParamsSpecAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentNewFromProjectParamsSpecAdmissionAdmissionLevelOwnerOnly   EnvironmentNewFromProjectParamsSpecAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentNewFromProjectParamsSpecAdmissionAdmissionLevelEveryone    EnvironmentNewFromProjectParamsSpecAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentNewFromProjectParamsSpecAdmission) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectParamsSpecAdmissionAdmissionLevelUnspecified, EnvironmentNewFromProjectParamsSpecAdmissionAdmissionLevelOwnerOnly, EnvironmentNewFromProjectParamsSpecAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// automations_file is the automations file spec of the environment
type EnvironmentNewFromProjectParamsSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath"`
	Session             param.Field[string] `json:"session"`
}

func (r EnvironmentNewFromProjectParamsSpecAutomationsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// content is the content spec of the environment
type EnvironmentNewFromProjectParamsSpecContent struct {
	// The Git email address
	GitEmail param.Field[string] `json:"gitEmail"`
	// The Git username
	GitUsername param.Field[string] `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer param.Field[EnvironmentNewFromProjectParamsSpecContentInitializer] `json:"initializer"`
	Session     param.Field[string]                                                `json:"session"`
}

func (r EnvironmentNewFromProjectParamsSpecContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnvironmentInitializer specifies how an environment is to be initialized
type EnvironmentNewFromProjectParamsSpecContentInitializer struct {
	Specs param.Field[[]EnvironmentNewFromProjectParamsSpecContentInitializerSpecUnion] `json:"specs"`
}

func (r EnvironmentNewFromProjectParamsSpecContentInitializer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewFromProjectParamsSpecContentInitializerSpec struct {
	ContextURL param.Field[interface{}] `json:"contextUrl"`
	Git        param.Field[interface{}] `json:"git"`
}

func (r EnvironmentNewFromProjectParamsSpecContentInitializerSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewFromProjectParamsSpecContentInitializerSpec) implementsEnvironmentNewFromProjectParamsSpecContentInitializerSpecUnion() {
}

// Satisfied by
// [EnvironmentNewFromProjectParamsSpecContentInitializerSpecsContextURL],
// [EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGit],
// [EnvironmentNewFromProjectParamsSpecContentInitializerSpec].
type EnvironmentNewFromProjectParamsSpecContentInitializerSpecUnion interface {
	implementsEnvironmentNewFromProjectParamsSpecContentInitializerSpecUnion()
}

type EnvironmentNewFromProjectParamsSpecContentInitializerSpecsContextURL struct {
	ContextURL param.Field[EnvironmentNewFromProjectParamsSpecContentInitializerSpecsContextURLContextURL] `json:"contextUrl,required"`
}

func (r EnvironmentNewFromProjectParamsSpecContentInitializerSpecsContextURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewFromProjectParamsSpecContentInitializerSpecsContextURL) implementsEnvironmentNewFromProjectParamsSpecContentInitializerSpecUnion() {
}

type EnvironmentNewFromProjectParamsSpecContentInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r EnvironmentNewFromProjectParamsSpecContentInitializerSpecsContextURLContextURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGit struct {
	Git param.Field[EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGit] `json:"git,required"`
}

func (r EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGit) implementsEnvironmentNewFromProjectParamsSpecContentInitializerSpecUnion() {
}

type EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation param.Field[string] `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget param.Field[string] `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri param.Field[string] `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode param.Field[EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetMode] `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri param.Field[string] `json:"upstreamRemoteUri"`
}

func (r EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetMode string

const (
	EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, EnvironmentNewFromProjectParamsSpecContentInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

// Phase is the desired phase of the environment
type EnvironmentNewFromProjectParamsSpecDesiredPhase string

const (
	EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseUnspecified EnvironmentNewFromProjectParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseCreating    EnvironmentNewFromProjectParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseStarting    EnvironmentNewFromProjectParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseRunning     EnvironmentNewFromProjectParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseUpdating    EnvironmentNewFromProjectParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseStopping    EnvironmentNewFromProjectParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseStopped     EnvironmentNewFromProjectParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseDeleting    EnvironmentNewFromProjectParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseDeleted     EnvironmentNewFromProjectParamsSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentNewFromProjectParamsSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseUnspecified, EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseCreating, EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseStarting, EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseRunning, EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseUpdating, EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseStopping, EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseStopped, EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseDeleting, EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// devcontainer is the devcontainer spec of the environment
type EnvironmentNewFromProjectParamsSpecDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath"`
	Session              param.Field[string] `json:"session"`
}

func (r EnvironmentNewFromProjectParamsSpecDevcontainer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// machine is the machine spec of the environment
type EnvironmentNewFromProjectParamsSpecMachine struct {
	// Class denotes the class of the environment we ought to start
	Class   param.Field[string] `json:"class" format:"uuid"`
	Session param.Field[string] `json:"session"`
}

func (r EnvironmentNewFromProjectParamsSpecMachine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewFromProjectParamsSpecPort struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission param.Field[EnvironmentNewFromProjectParamsSpecPortsAdmission] `json:"admission"`
	// name of this port
	Name param.Field[string] `json:"name"`
	// port number
	Port param.Field[int64] `json:"port"`
}

func (r EnvironmentNewFromProjectParamsSpecPort) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentNewFromProjectParamsSpecPortsAdmission string

const (
	EnvironmentNewFromProjectParamsSpecPortsAdmissionAdmissionLevelUnspecified EnvironmentNewFromProjectParamsSpecPortsAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentNewFromProjectParamsSpecPortsAdmissionAdmissionLevelOwnerOnly   EnvironmentNewFromProjectParamsSpecPortsAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentNewFromProjectParamsSpecPortsAdmissionAdmissionLevelEveryone    EnvironmentNewFromProjectParamsSpecPortsAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentNewFromProjectParamsSpecPortsAdmission) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectParamsSpecPortsAdmissionAdmissionLevelUnspecified, EnvironmentNewFromProjectParamsSpecPortsAdmissionAdmissionLevelOwnerOnly, EnvironmentNewFromProjectParamsSpecPortsAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

type EnvironmentNewFromProjectParamsSpecSecret struct {
	EnvironmentVariable param.Field[string] `json:"environmentVariable"`
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath          param.Field[string] `json:"filePath"`
	GitCredentialHost param.Field[string] `json:"gitCredentialHost"`
	// name is the human readable description of the secret
	Name param.Field[string] `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session param.Field[string] `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source param.Field[string] `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef param.Field[string] `json:"sourceRef"`
}

func (r EnvironmentNewFromProjectParamsSpecSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewFromProjectParamsSpecSecret) implementsEnvironmentNewFromProjectParamsSpecSecretUnion() {
}

// Satisfied by [EnvironmentNewFromProjectParamsSpecSecretsObject],
// [EnvironmentNewFromProjectParamsSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted],
// [EnvironmentNewFromProjectParamsSpecSecretsObject],
// [EnvironmentNewFromProjectParamsSpecSecret].
type EnvironmentNewFromProjectParamsSpecSecretUnion interface {
	implementsEnvironmentNewFromProjectParamsSpecSecretUnion()
}

type EnvironmentNewFromProjectParamsSpecSecretsObject struct {
	EnvironmentVariable param.Field[string] `json:"environmentVariable,required"`
	// name is the human readable description of the secret
	Name param.Field[string] `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session param.Field[string] `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source param.Field[string] `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef param.Field[string] `json:"sourceRef"`
}

func (r EnvironmentNewFromProjectParamsSpecSecretsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewFromProjectParamsSpecSecretsObject) implementsEnvironmentNewFromProjectParamsSpecSecretUnion() {
}

type EnvironmentNewFromProjectParamsSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted struct {
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath param.Field[string] `json:"filePath,required"`
	// name is the human readable description of the secret
	Name param.Field[string] `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session param.Field[string] `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source param.Field[string] `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef param.Field[string] `json:"sourceRef"`
}

func (r EnvironmentNewFromProjectParamsSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnvironmentNewFromProjectParamsSpecSecretsFilePathIsThePathInsideTheDevcontainerWhereTheSecretIsMounted) implementsEnvironmentNewFromProjectParamsSpecSecretUnion() {
}

type EnvironmentNewFromProjectParamsSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID param.Field[string] `json:"id"`
	// value is the actual public key in the public key file format
	Value param.Field[string] `json:"value"`
}

func (r EnvironmentNewFromProjectParamsSpecSSHPublicKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Timeout configures the environment timeout
type EnvironmentNewFromProjectParamsSpecTimeout struct {
	// A Duration represents a signed, fixed-length span of time represented as a count
	// of seconds and fractions of seconds at nanosecond resolution. It is independent
	// of any calendar and concepts like "day" or "month". It is related to Timestamp
	// in that the difference between two Timestamp values is a Duration and it can be
	// added or subtracted from a Timestamp. Range is approximately +-10,000 years.
	//
	// # Examples
	//
	// Example 1: Compute Duration from two Timestamps in pseudo code.
	//
	//	Timestamp start = ...;
	//	Timestamp end = ...;
	//	Duration duration = ...;
	//
	//	duration.seconds = end.seconds - start.seconds;
	//	duration.nanos = end.nanos - start.nanos;
	//
	//	if (duration.seconds < 0 && duration.nanos > 0) {
	//	  duration.seconds += 1;
	//	  duration.nanos -= 1000000000;
	//	} else if (duration.seconds > 0 && duration.nanos < 0) {
	//	  duration.seconds -= 1;
	//	  duration.nanos += 1000000000;
	//	}
	//
	// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
	//
	//	Timestamp start = ...;
	//	Duration duration = ...;
	//	Timestamp end = ...;
	//
	//	end.seconds = start.seconds + duration.seconds;
	//	end.nanos = start.nanos + duration.nanos;
	//
	//	if (end.nanos < 0) {
	//	  end.seconds -= 1;
	//	  end.nanos += 1000000000;
	//	} else if (end.nanos >= 1000000000) {
	//	  end.seconds += 1;
	//	  end.nanos -= 1000000000;
	//	}
	//
	// Example 3: Compute Duration from datetime.timedelta in Python.
	//
	//	td = datetime.timedelta(days=3, minutes=10)
	//	duration = Duration()
	//	duration.FromTimedelta(td)
	//
	// # JSON Mapping
	//
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by
	// the number of seconds, with nanoseconds expressed as fractional seconds. For
	// example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as
	// "3.000000001s", and 3 seconds and 1 microsecond should be expressed in JSON
	// format as "3.000001s".
	Disconnected param.Field[string] `json:"disconnected" format:"regex"`
}

func (r EnvironmentNewFromProjectParamsSpecTimeout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewLogsTokenParams struct {
	// environment_id specifies the environment for which the logs token should be
	// created.
	//
	// +required
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentNewLogsTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentMarkActiveParams struct {
	// EnvironmentActivitySignal used to signal activity for an environment.
	ActivitySignal param.Field[EnvironmentMarkActiveParamsActivitySignal] `json:"activitySignal"`
	// The ID of the environment to update activity for.
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentMarkActiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnvironmentActivitySignal used to signal activity for an environment.
type EnvironmentMarkActiveParamsActivitySignal struct {
	// source of the activity signal, such as "VS Code", "SSH", or "Automations". It
	// should be a human-readable string that describes the source of the activity
	// signal.
	Source param.Field[string] `json:"source"`
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
	Timestamp param.Field[time.Time] `json:"timestamp" format:"date-time"`
}

func (r EnvironmentMarkActiveParamsActivitySignal) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentStartParams struct {
	// environment_id specifies which environment should be started.
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentStopParams struct {
	// environment_id specifies which environment should be stopped.
	//
	// +required
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentStopParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
