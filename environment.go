// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
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

// EnvironmentService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentService] method instead.
type EnvironmentService struct {
	Options []option.RequestOption
}

// NewEnvironmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnvironmentService(opts ...option.RequestOption) (r *EnvironmentService) {
	r = &EnvironmentService{}
	r.Options = opts
	return
}

// CreateEnvironment creates a new environment and starts it.
func (r *EnvironmentService) New(ctx context.Context, params EnvironmentNewParams, opts ...option.RequestOption) (res *EnvironmentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/CreateEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetEnvironment returns a single environment.
//
// +return NOT_FOUND User does not have access to an environment with the given ID
// +return NOT_FOUND Environment does not exist
func (r *EnvironmentService) Get(ctx context.Context, params EnvironmentGetParams, opts ...option.RequestOption) (res *EnvironmentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/GetEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListEnvironments returns a list of environments that match the query.
func (r *EnvironmentService) List(ctx context.Context, params EnvironmentListParams, opts ...option.RequestOption) (res *EnvironmentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/ListEnvironments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// CreateAbdStartEnvironmentFromProject creates a new environment from a project
// and starts it.
func (r *EnvironmentService) NewFromProject(ctx context.Context, params EnvironmentNewFromProjectParams, opts ...option.RequestOption) (res *EnvironmentNewFromProjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/CreateEnvironmentFromProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// StartEnvironment starts an environment. This function is idempotent, i.e. if the
// environment is already running no error is returned.
func (r *EnvironmentService) Start(ctx context.Context, params EnvironmentStartParams, opts ...option.RequestOption) (res *EnvironmentStartResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/StartEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
	Secrets []EnvironmentNewResponseEnvironmentSpecSecretsUnion `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion EnvironmentNewResponseEnvironmentSpecSpecVersionUnion `json:"specVersion"`
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
	// environment, relative to the repo root.
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
	Specs []EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnion `json:"specs"`
	JSON  environmentNewResponseEnvironmentSpecContentInitializerJSON         `json:"-"`
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

// Union satisfied by
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnknown],
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnknown] or
// [EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnknown].
type EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnion interface {
	implementsEnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentNewResponseEnvironmentSpecContentInitializerSpecsUnion)(nil)).Elem(), "")
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
	// root
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

// Union satisfied by [EnvironmentNewResponseEnvironmentSpecSecretsUnknown],
// [EnvironmentNewResponseEnvironmentSpecSecretsUnknown],
// [EnvironmentNewResponseEnvironmentSpecSecretsUnknown] or
// [EnvironmentNewResponseEnvironmentSpecSecretsUnknown].
type EnvironmentNewResponseEnvironmentSpecSecretsUnion interface {
	implementsEnvironmentNewResponseEnvironmentSpecSecretsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentNewResponseEnvironmentSpecSecretsUnion)(nil)).Elem(), "")
}

// version of the spec. The value of this field has no semantic meaning (e.g. don't
// interpret it as as a timestamp), but it can be used to impose a partial order.
// If a.spec_version < b.spec_version then a was the spec before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentNewResponseEnvironmentSpecSpecVersionUnion interface {
	ImplementsEnvironmentNewResponseEnvironmentSpecSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentNewResponseEnvironmentSpecSpecVersionUnion)(nil)).Elem(),
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
	// but their statuus has different versions. The value of this field has no
	// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
	// to impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion EnvironmentNewResponseEnvironmentStatusStatusVersionUnion `json:"statusVersion"`
	// warning_message contains warnings, e.g. when the environment is present but not
	// in the expected state.
	WarningMessage []string                                    `json:"warningMessage"`
	JSON           environmentNewResponseEnvironmentStatusJSON `json:"-"`
}

// environmentNewResponseEnvironmentStatusJSON contains the JSON metadata for the
// struct [EnvironmentNewResponseEnvironmentStatus]
type environmentNewResponseEnvironmentStatusJSON struct {
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
	Message     string                                                           `json:"message"`
	SpecVersion EnvironmentNewResponseEnvironmentStatusRunnerAckSpecVersionUnion `json:"specVersion"`
	StatusCode  EnvironmentNewResponseEnvironmentStatusRunnerAckStatusCode       `json:"statusCode"`
	JSON        environmentNewResponseEnvironmentStatusRunnerAckJSON             `json:"-"`
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

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentNewResponseEnvironmentStatusRunnerAckSpecVersionUnion interface {
	ImplementsEnvironmentNewResponseEnvironmentStatusRunnerAckSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentNewResponseEnvironmentStatusRunnerAckSpecVersionUnion)(nil)).Elem(),
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

// version of the status update. Environment instances themselves are unversioned,
// but their statuus has different versions. The value of this field has no
// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
// to impose a partial order. If a.status_version < b.status_version then a was the
// status before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentNewResponseEnvironmentStatusStatusVersionUnion interface {
	ImplementsEnvironmentNewResponseEnvironmentStatusStatusVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentNewResponseEnvironmentStatusStatusVersionUnion)(nil)).Elem(),
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
	Secrets []EnvironmentGetResponseEnvironmentSpecSecretsUnion `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion EnvironmentGetResponseEnvironmentSpecSpecVersionUnion `json:"specVersion"`
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
	// environment, relative to the repo root.
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
	Specs []EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion `json:"specs"`
	JSON  environmentGetResponseEnvironmentSpecContentInitializerJSON         `json:"-"`
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

// Union satisfied by
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnknown],
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnknown] or
// [EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnknown].
type EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion interface {
	implementsEnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion)(nil)).Elem(), "")
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
	// root
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

// Union satisfied by [EnvironmentGetResponseEnvironmentSpecSecretsUnknown],
// [EnvironmentGetResponseEnvironmentSpecSecretsUnknown],
// [EnvironmentGetResponseEnvironmentSpecSecretsUnknown] or
// [EnvironmentGetResponseEnvironmentSpecSecretsUnknown].
type EnvironmentGetResponseEnvironmentSpecSecretsUnion interface {
	implementsEnvironmentGetResponseEnvironmentSpecSecretsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentGetResponseEnvironmentSpecSecretsUnion)(nil)).Elem(), "")
}

// version of the spec. The value of this field has no semantic meaning (e.g. don't
// interpret it as as a timestamp), but it can be used to impose a partial order.
// If a.spec_version < b.spec_version then a was the spec before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentGetResponseEnvironmentSpecSpecVersionUnion interface {
	ImplementsEnvironmentGetResponseEnvironmentSpecSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentGetResponseEnvironmentSpecSpecVersionUnion)(nil)).Elem(),
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
	// but their statuus has different versions. The value of this field has no
	// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
	// to impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion EnvironmentGetResponseEnvironmentStatusStatusVersionUnion `json:"statusVersion"`
	// warning_message contains warnings, e.g. when the environment is present but not
	// in the expected state.
	WarningMessage []string                                    `json:"warningMessage"`
	JSON           environmentGetResponseEnvironmentStatusJSON `json:"-"`
}

// environmentGetResponseEnvironmentStatusJSON contains the JSON metadata for the
// struct [EnvironmentGetResponseEnvironmentStatus]
type environmentGetResponseEnvironmentStatusJSON struct {
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
	Message     string                                                           `json:"message"`
	SpecVersion EnvironmentGetResponseEnvironmentStatusRunnerAckSpecVersionUnion `json:"specVersion"`
	StatusCode  EnvironmentGetResponseEnvironmentStatusRunnerAckStatusCode       `json:"statusCode"`
	JSON        environmentGetResponseEnvironmentStatusRunnerAckJSON             `json:"-"`
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

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentGetResponseEnvironmentStatusRunnerAckSpecVersionUnion interface {
	ImplementsEnvironmentGetResponseEnvironmentStatusRunnerAckSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentGetResponseEnvironmentStatusRunnerAckSpecVersionUnion)(nil)).Elem(),
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

// version of the status update. Environment instances themselves are unversioned,
// but their statuus has different versions. The value of this field has no
// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
// to impose a partial order. If a.status_version < b.status_version then a was the
// status before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentGetResponseEnvironmentStatusStatusVersionUnion interface {
	ImplementsEnvironmentGetResponseEnvironmentStatusStatusVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentGetResponseEnvironmentStatusStatusVersionUnion)(nil)).Elem(),
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

type EnvironmentListResponse struct {
	// environments are the environments that matched the query
	Environments []EnvironmentListResponseEnvironment `json:"environments"`
	// pagination contains the pagination options for listing environments
	Pagination EnvironmentListResponsePagination `json:"pagination"`
	JSON       environmentListResponseJSON       `json:"-"`
}

// environmentListResponseJSON contains the JSON metadata for the struct
// [EnvironmentListResponse]
type environmentListResponseJSON struct {
	Environments apijson.Field
	Pagination   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseJSON) RawJSON() string {
	return r.raw
}

// +resource get environment
type EnvironmentListResponseEnvironment struct {
	// ID is a unique identifier of this environment. No other environment with the
	// same name must be managed by this environment manager
	ID string `json:"id"`
	// EnvironmentMetadata is data associated with an environment that's required for
	// other parts of the system to function
	Metadata EnvironmentListResponseEnvironmentsMetadata `json:"metadata"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec EnvironmentListResponseEnvironmentsSpec `json:"spec"`
	// EnvironmentStatus describes an environment status
	Status EnvironmentListResponseEnvironmentsStatus `json:"status"`
	JSON   environmentListResponseEnvironmentJSON    `json:"-"`
}

// environmentListResponseEnvironmentJSON contains the JSON metadata for the struct
// [EnvironmentListResponseEnvironment]
type environmentListResponseEnvironmentJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentMetadata is data associated with an environment that's required for
// other parts of the system to function
type EnvironmentListResponseEnvironmentsMetadata struct {
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
	Creator EnvironmentListResponseEnvironmentsMetadataCreator `json:"creator"`
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
	RunnerID string                                          `json:"runnerId"`
	JSON     environmentListResponseEnvironmentsMetadataJSON `json:"-"`
}

// environmentListResponseEnvironmentsMetadataJSON contains the JSON metadata for
// the struct [EnvironmentListResponseEnvironmentsMetadata]
type environmentListResponseEnvironmentsMetadataJSON struct {
	Annotations        apijson.Field
	CreatedAt          apijson.Field
	Creator            apijson.Field
	Name               apijson.Field
	OrganizationID     apijson.Field
	OriginalContextURL apijson.Field
	ProjectID          apijson.Field
	RunnerID           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the creator of the environment
type EnvironmentListResponseEnvironmentsMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal EnvironmentListResponseEnvironmentsMetadataCreatorPrincipal `json:"principal"`
	JSON      environmentListResponseEnvironmentsMetadataCreatorJSON      `json:"-"`
}

// environmentListResponseEnvironmentsMetadataCreatorJSON contains the JSON
// metadata for the struct [EnvironmentListResponseEnvironmentsMetadataCreator]
type environmentListResponseEnvironmentsMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type EnvironmentListResponseEnvironmentsMetadataCreatorPrincipal string

const (
	EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalUnspecified    EnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalAccount        EnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalUser           EnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_USER"
	EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalRunner         EnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalEnvironment    EnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalServiceAccount EnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EnvironmentListResponseEnvironmentsMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalUnspecified, EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalAccount, EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalUser, EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalRunner, EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalEnvironment, EnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type EnvironmentListResponseEnvironmentsSpec struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission EnvironmentListResponseEnvironmentsSpecAdmission `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile EnvironmentListResponseEnvironmentsSpecAutomationsFile `json:"automationsFile"`
	// content is the content spec of the environment
	Content EnvironmentListResponseEnvironmentsSpecContent `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase EnvironmentListResponseEnvironmentsSpecDesiredPhase `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer EnvironmentListResponseEnvironmentsSpecDevcontainer `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine EnvironmentListResponseEnvironmentsSpecMachine `json:"machine"`
	// ports is the set of ports which ought to be exposed to the internet
	Ports []EnvironmentListResponseEnvironmentsSpecPort `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets []EnvironmentListResponseEnvironmentsSpecSecretsUnion `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion EnvironmentListResponseEnvironmentsSpecSpecVersionUnion `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys []EnvironmentListResponseEnvironmentsSpecSSHPublicKey `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout EnvironmentListResponseEnvironmentsSpecTimeout `json:"timeout"`
	JSON    environmentListResponseEnvironmentsSpecJSON    `json:"-"`
}

// environmentListResponseEnvironmentsSpecJSON contains the JSON metadata for the
// struct [EnvironmentListResponseEnvironmentsSpec]
type environmentListResponseEnvironmentsSpecJSON struct {
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

func (r *EnvironmentListResponseEnvironmentsSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsSpecJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentListResponseEnvironmentsSpecAdmission string

const (
	EnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelUnspecified EnvironmentListResponseEnvironmentsSpecAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelOwnerOnly   EnvironmentListResponseEnvironmentsSpecAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelEveryone    EnvironmentListResponseEnvironmentsSpecAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentListResponseEnvironmentsSpecAdmission) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelUnspecified, EnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelOwnerOnly, EnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// automations_file is the automations file spec of the environment
type EnvironmentListResponseEnvironmentsSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root.
	AutomationsFilePath string                                                     `json:"automationsFilePath"`
	Session             string                                                     `json:"session"`
	JSON                environmentListResponseEnvironmentsSpecAutomationsFileJSON `json:"-"`
}

// environmentListResponseEnvironmentsSpecAutomationsFileJSON contains the JSON
// metadata for the struct [EnvironmentListResponseEnvironmentsSpecAutomationsFile]
type environmentListResponseEnvironmentsSpecAutomationsFileJSON struct {
	AutomationsFilePath apijson.Field
	Session             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsSpecAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsSpecAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// content is the content spec of the environment
type EnvironmentListResponseEnvironmentsSpecContent struct {
	// The Git email address
	GitEmail string `json:"gitEmail"`
	// The Git username
	GitUsername string `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer EnvironmentListResponseEnvironmentsSpecContentInitializer `json:"initializer"`
	Session     string                                                    `json:"session"`
	JSON        environmentListResponseEnvironmentsSpecContentJSON        `json:"-"`
}

// environmentListResponseEnvironmentsSpecContentJSON contains the JSON metadata
// for the struct [EnvironmentListResponseEnvironmentsSpecContent]
type environmentListResponseEnvironmentsSpecContentJSON struct {
	GitEmail    apijson.Field
	GitUsername apijson.Field
	Initializer apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsSpecContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsSpecContentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentInitializer specifies how an environment is to be initialized
type EnvironmentListResponseEnvironmentsSpecContentInitializer struct {
	Specs []EnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnion `json:"specs"`
	JSON  environmentListResponseEnvironmentsSpecContentInitializerJSON         `json:"-"`
}

// environmentListResponseEnvironmentsSpecContentInitializerJSON contains the JSON
// metadata for the struct
// [EnvironmentListResponseEnvironmentsSpecContentInitializer]
type environmentListResponseEnvironmentsSpecContentInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsSpecContentInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsSpecContentInitializerJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [EnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnknown],
// [EnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnknown] or
// [EnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnknown].
type EnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnion interface {
	implementsEnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnion)(nil)).Elem(), "")
}

// Phase is the desired phase of the environment
type EnvironmentListResponseEnvironmentsSpecDesiredPhase string

const (
	EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseUnspecified EnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseCreating    EnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStarting    EnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseRunning     EnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseUpdating    EnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStopping    EnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStopped     EnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseDeleting    EnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseDeleted     EnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentListResponseEnvironmentsSpecDesiredPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseUnspecified, EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseCreating, EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStarting, EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseRunning, EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseUpdating, EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStopping, EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStopped, EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseDeleting, EnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// devcontainer is the devcontainer spec of the environment
type EnvironmentListResponseEnvironmentsSpecDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string                                                  `json:"devcontainerFilePath"`
	Session              string                                                  `json:"session"`
	JSON                 environmentListResponseEnvironmentsSpecDevcontainerJSON `json:"-"`
}

// environmentListResponseEnvironmentsSpecDevcontainerJSON contains the JSON
// metadata for the struct [EnvironmentListResponseEnvironmentsSpecDevcontainer]
type environmentListResponseEnvironmentsSpecDevcontainerJSON struct {
	DevcontainerFilePath apijson.Field
	Session              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsSpecDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsSpecDevcontainerJSON) RawJSON() string {
	return r.raw
}

// machine is the machine spec of the environment
type EnvironmentListResponseEnvironmentsSpecMachine struct {
	// Class denotes the class of the environment we ought to start
	Class   string                                             `json:"class" format:"uuid"`
	Session string                                             `json:"session"`
	JSON    environmentListResponseEnvironmentsSpecMachineJSON `json:"-"`
}

// environmentListResponseEnvironmentsSpecMachineJSON contains the JSON metadata
// for the struct [EnvironmentListResponseEnvironmentsSpecMachine]
type environmentListResponseEnvironmentsSpecMachineJSON struct {
	Class       apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsSpecMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsSpecMachineJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseEnvironmentsSpecPort struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission EnvironmentListResponseEnvironmentsSpecPortsAdmission `json:"admission"`
	// name of this port
	Name string `json:"name"`
	// port number
	Port int64                                           `json:"port"`
	JSON environmentListResponseEnvironmentsSpecPortJSON `json:"-"`
}

// environmentListResponseEnvironmentsSpecPortJSON contains the JSON metadata for
// the struct [EnvironmentListResponseEnvironmentsSpecPort]
type environmentListResponseEnvironmentsSpecPortJSON struct {
	Admission   apijson.Field
	Name        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsSpecPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsSpecPortJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type EnvironmentListResponseEnvironmentsSpecPortsAdmission string

const (
	EnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelUnspecified EnvironmentListResponseEnvironmentsSpecPortsAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelOwnerOnly   EnvironmentListResponseEnvironmentsSpecPortsAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	EnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelEveryone    EnvironmentListResponseEnvironmentsSpecPortsAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r EnvironmentListResponseEnvironmentsSpecPortsAdmission) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelUnspecified, EnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelOwnerOnly, EnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// Union satisfied by [EnvironmentListResponseEnvironmentsSpecSecretsUnknown],
// [EnvironmentListResponseEnvironmentsSpecSecretsUnknown],
// [EnvironmentListResponseEnvironmentsSpecSecretsUnknown] or
// [EnvironmentListResponseEnvironmentsSpecSecretsUnknown].
type EnvironmentListResponseEnvironmentsSpecSecretsUnion interface {
	implementsEnvironmentListResponseEnvironmentsSpecSecretsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentListResponseEnvironmentsSpecSecretsUnion)(nil)).Elem(), "")
}

// version of the spec. The value of this field has no semantic meaning (e.g. don't
// interpret it as as a timestamp), but it can be used to impose a partial order.
// If a.spec_version < b.spec_version then a was the spec before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentListResponseEnvironmentsSpecSpecVersionUnion interface {
	ImplementsEnvironmentListResponseEnvironmentsSpecSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentListResponseEnvironmentsSpecSpecVersionUnion)(nil)).Elem(),
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

type EnvironmentListResponseEnvironmentsSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// value is the actual public key in the public key file format
	Value string                                                  `json:"value"`
	JSON  environmentListResponseEnvironmentsSpecSSHPublicKeyJSON `json:"-"`
}

// environmentListResponseEnvironmentsSpecSSHPublicKeyJSON contains the JSON
// metadata for the struct [EnvironmentListResponseEnvironmentsSpecSSHPublicKey]
type environmentListResponseEnvironmentsSpecSSHPublicKeyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsSpecSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsSpecSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// Timeout configures the environment timeout
type EnvironmentListResponseEnvironmentsSpecTimeout struct {
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
	Disconnected string                                             `json:"disconnected" format:"regex"`
	JSON         environmentListResponseEnvironmentsSpecTimeoutJSON `json:"-"`
}

// environmentListResponseEnvironmentsSpecTimeoutJSON contains the JSON metadata
// for the struct [EnvironmentListResponseEnvironmentsSpecTimeout]
type environmentListResponseEnvironmentsSpecTimeoutJSON struct {
	Disconnected apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsSpecTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsSpecTimeoutJSON) RawJSON() string {
	return r.raw
}

// EnvironmentStatus describes an environment status
type EnvironmentListResponseEnvironmentsStatus struct {
	// automations_file contains the status of the automations file.
	AutomationsFile EnvironmentListResponseEnvironmentsStatusAutomationsFile `json:"automationsFile"`
	// content contains the status of the environment content.
	Content EnvironmentListResponseEnvironmentsStatusContent `json:"content"`
	// devcontainer contains the status of the devcontainer.
	Devcontainer EnvironmentListResponseEnvironmentsStatusDevcontainer `json:"devcontainer"`
	// environment_url contains the URL at which the environment can be accessed. This
	// field is only set if the environment is running.
	EnvironmentURLs EnvironmentListResponseEnvironmentsStatusEnvironmentURLs `json:"environmentUrls"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage []string `json:"failureMessage"`
	// machine contains the status of the environment machine
	Machine EnvironmentListResponseEnvironmentsStatusMachine `json:"machine"`
	// the phase of an environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentListResponseEnvironmentsStatusPhase `json:"phase"`
	// RunnerACK is the acknowledgement from the runner that is has received the
	// environment spec.
	RunnerAck EnvironmentListResponseEnvironmentsStatusRunnerAck `json:"runnerAck"`
	// secrets contains the status of the environment secrets
	Secrets []EnvironmentListResponseEnvironmentsStatusSecret `json:"secrets"`
	// ssh_public_keys contains the status of the environment ssh public keys
	SSHPublicKeys []EnvironmentListResponseEnvironmentsStatusSSHPublicKey `json:"sshPublicKeys"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their statuus has different versions. The value of this field has no
	// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
	// to impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion EnvironmentListResponseEnvironmentsStatusStatusVersionUnion `json:"statusVersion"`
	// warning_message contains warnings, e.g. when the environment is present but not
	// in the expected state.
	WarningMessage []string                                      `json:"warningMessage"`
	JSON           environmentListResponseEnvironmentsStatusJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusJSON contains the JSON metadata for the
// struct [EnvironmentListResponseEnvironmentsStatus]
type environmentListResponseEnvironmentsStatusJSON struct {
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

func (r *EnvironmentListResponseEnvironmentsStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusJSON) RawJSON() string {
	return r.raw
}

// automations_file contains the status of the automations file.
type EnvironmentListResponseEnvironmentsStatusAutomationsFile struct {
	// automations_file_path is the path to the automations file relative to the repo
	// root.
	AutomationsFilePath string `json:"automationsFilePath"`
	// automations_file_presence indicates how an automations file is present in the
	// environment.
	AutomationsFilePresence EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresence `json:"automationsFilePresence"`
	// failure_message contains the reason the automations file failed to be applied.
	// This is only set if the phase is FAILED.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the automations file.
	Phase EnvironmentListResponseEnvironmentsStatusAutomationsFilePhase `json:"phase"`
	// session is the automations file session that is currently applied in the
	// environment.
	Session string                                                       `json:"session"`
	JSON    environmentListResponseEnvironmentsStatusAutomationsFileJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusAutomationsFileJSON contains the JSON
// metadata for the struct
// [EnvironmentListResponseEnvironmentsStatusAutomationsFile]
type environmentListResponseEnvironmentsStatusAutomationsFileJSON struct {
	AutomationsFilePath     apijson.Field
	AutomationsFilePresence apijson.Field
	FailureMessage          apijson.Field
	Phase                   apijson.Field
	Session                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsStatusAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// automations_file_presence indicates how an automations file is present in the
// environment.
type EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresence string

const (
	EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresencePresenceUnspecified EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresencePresenceAbsent      EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresence = "PRESENCE_ABSENT"
	EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresencePresenceDiscovered  EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresencePresenceSpecified   EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresencePresenceUnspecified, EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresencePresenceAbsent, EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresencePresenceDiscovered, EnvironmentListResponseEnvironmentsStatusAutomationsFileAutomationsFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the automations file.
type EnvironmentListResponseEnvironmentsStatusAutomationsFilePhase string

const (
	EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseUnspecified  EnvironmentListResponseEnvironmentsStatusAutomationsFilePhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseCreating     EnvironmentListResponseEnvironmentsStatusAutomationsFilePhase = "CONTENT_PHASE_CREATING"
	EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseInitializing EnvironmentListResponseEnvironmentsStatusAutomationsFilePhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseReady        EnvironmentListResponseEnvironmentsStatusAutomationsFilePhase = "CONTENT_PHASE_READY"
	EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseUpdating     EnvironmentListResponseEnvironmentsStatusAutomationsFilePhase = "CONTENT_PHASE_UPDATING"
	EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseFailed       EnvironmentListResponseEnvironmentsStatusAutomationsFilePhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentListResponseEnvironmentsStatusAutomationsFilePhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseUnspecified, EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseCreating, EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseInitializing, EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseReady, EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseUpdating, EnvironmentListResponseEnvironmentsStatusAutomationsFilePhaseContentPhaseFailed:
		return true
	}
	return false
}

// content contains the status of the environment content.
type EnvironmentListResponseEnvironmentsStatusContent struct {
	// content_location_in_machine is the location of the content in the machine
	ContentLocationInMachine string `json:"contentLocationInMachine"`
	// failure_message contains the reason the content initialization failed.
	FailureMessage string `json:"failureMessage"`
	// git is the Git working copy status of the environment. Note: this is a
	// best-effort field and more often than not will not be present. Its absence does
	// not indicate the absence of a working copy.
	Git EnvironmentListResponseEnvironmentsStatusContentGit `json:"git"`
	// phase is the current phase of the environment content
	Phase EnvironmentListResponseEnvironmentsStatusContentPhase `json:"phase"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the content is present but not in
	// the expected state.
	WarningMessage string                                               `json:"warningMessage"`
	JSON           environmentListResponseEnvironmentsStatusContentJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusContentJSON contains the JSON metadata
// for the struct [EnvironmentListResponseEnvironmentsStatusContent]
type environmentListResponseEnvironmentsStatusContentJSON struct {
	ContentLocationInMachine apijson.Field
	FailureMessage           apijson.Field
	Git                      apijson.Field
	Phase                    apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsStatusContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusContentJSON) RawJSON() string {
	return r.raw
}

// git is the Git working copy status of the environment. Note: this is a
// best-effort field and more often than not will not be present. Its absence does
// not indicate the absence of a working copy.
type EnvironmentListResponseEnvironmentsStatusContentGit struct {
	// branch is branch we're currently on
	Branch string `json:"branch"`
	// changed_files is an array of changed files in the environment, possibly
	// truncated
	ChangedFiles []EnvironmentListResponseEnvironmentsStatusContentGitChangedFile `json:"changedFiles"`
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
	UnpushedCommits []string                                                `json:"unpushedCommits"`
	JSON            environmentListResponseEnvironmentsStatusContentGitJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusContentGitJSON contains the JSON
// metadata for the struct [EnvironmentListResponseEnvironmentsStatusContentGit]
type environmentListResponseEnvironmentsStatusContentGitJSON struct {
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

func (r *EnvironmentListResponseEnvironmentsStatusContentGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusContentGitJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseEnvironmentsStatusContentGitChangedFile struct {
	// ChangeType is the type of change that happened to the file
	ChangeType EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeType `json:"changeType"`
	// path is the path of the file
	Path string                                                             `json:"path"`
	JSON environmentListResponseEnvironmentsStatusContentGitChangedFileJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusContentGitChangedFileJSON contains the
// JSON metadata for the struct
// [EnvironmentListResponseEnvironmentsStatusContentGitChangedFile]
type environmentListResponseEnvironmentsStatusContentGitChangedFileJSON struct {
	ChangeType  apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsStatusContentGitChangedFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusContentGitChangedFileJSON) RawJSON() string {
	return r.raw
}

// ChangeType is the type of change that happened to the file
type EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeType string

const (
	EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified        EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeAdded              EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_ADDED"
	EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeModified           EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_MODIFIED"
	EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeDeleted            EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_DELETED"
	EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeRenamed            EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_RENAMED"
	EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeCopied             EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_COPIED"
	EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UPDATED_BUT_UNMERGED"
	EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeUntracked          EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNTRACKED"
)

func (r EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeType) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified, EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeAdded, EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeModified, EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeDeleted, EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeRenamed, EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeCopied, EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged, EnvironmentListResponseEnvironmentsStatusContentGitChangedFilesChangeTypeChangeTypeUntracked:
		return true
	}
	return false
}

// phase is the current phase of the environment content
type EnvironmentListResponseEnvironmentsStatusContentPhase string

const (
	EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseUnspecified  EnvironmentListResponseEnvironmentsStatusContentPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseCreating     EnvironmentListResponseEnvironmentsStatusContentPhase = "CONTENT_PHASE_CREATING"
	EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseInitializing EnvironmentListResponseEnvironmentsStatusContentPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseReady        EnvironmentListResponseEnvironmentsStatusContentPhase = "CONTENT_PHASE_READY"
	EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseUpdating     EnvironmentListResponseEnvironmentsStatusContentPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseFailed       EnvironmentListResponseEnvironmentsStatusContentPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentListResponseEnvironmentsStatusContentPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseUnspecified, EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseCreating, EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseInitializing, EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseReady, EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseUpdating, EnvironmentListResponseEnvironmentsStatusContentPhaseContentPhaseFailed:
		return true
	}
	return false
}

// devcontainer contains the status of the devcontainer.
type EnvironmentListResponseEnvironmentsStatusDevcontainer struct {
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
	DevcontainerFilePresence EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresence `json:"devcontainerFilePresence"`
	// failure_message contains the reason the devcontainer failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the devcontainer
	Phase EnvironmentListResponseEnvironmentsStatusDevcontainerPhase `json:"phase"`
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
	WarningMessage string                                                    `json:"warningMessage"`
	JSON           environmentListResponseEnvironmentsStatusDevcontainerJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusDevcontainerJSON contains the JSON
// metadata for the struct [EnvironmentListResponseEnvironmentsStatusDevcontainer]
type environmentListResponseEnvironmentsStatusDevcontainerJSON struct {
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

func (r *EnvironmentListResponseEnvironmentsStatusDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusDevcontainerJSON) RawJSON() string {
	return r.raw
}

// devcontainer_file_presence indicates how the devcontainer file is present in the
// repo.
type EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresence string

const (
	EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresencePresenceUnspecified EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresencePresenceGenerated   EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresence = "PRESENCE_GENERATED"
	EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresencePresenceDiscovered  EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresencePresenceSpecified   EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresencePresenceUnspecified, EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresencePresenceGenerated, EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresencePresenceDiscovered, EnvironmentListResponseEnvironmentsStatusDevcontainerDevcontainerFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the devcontainer
type EnvironmentListResponseEnvironmentsStatusDevcontainerPhase string

const (
	EnvironmentListResponseEnvironmentsStatusDevcontainerPhasePhaseUnspecified EnvironmentListResponseEnvironmentsStatusDevcontainerPhase = "PHASE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsStatusDevcontainerPhasePhaseCreating    EnvironmentListResponseEnvironmentsStatusDevcontainerPhase = "PHASE_CREATING"
	EnvironmentListResponseEnvironmentsStatusDevcontainerPhasePhaseRunning     EnvironmentListResponseEnvironmentsStatusDevcontainerPhase = "PHASE_RUNNING"
	EnvironmentListResponseEnvironmentsStatusDevcontainerPhasePhaseStopped     EnvironmentListResponseEnvironmentsStatusDevcontainerPhase = "PHASE_STOPPED"
	EnvironmentListResponseEnvironmentsStatusDevcontainerPhasePhaseFailed      EnvironmentListResponseEnvironmentsStatusDevcontainerPhase = "PHASE_FAILED"
)

func (r EnvironmentListResponseEnvironmentsStatusDevcontainerPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsStatusDevcontainerPhasePhaseUnspecified, EnvironmentListResponseEnvironmentsStatusDevcontainerPhasePhaseCreating, EnvironmentListResponseEnvironmentsStatusDevcontainerPhasePhaseRunning, EnvironmentListResponseEnvironmentsStatusDevcontainerPhasePhaseStopped, EnvironmentListResponseEnvironmentsStatusDevcontainerPhasePhaseFailed:
		return true
	}
	return false
}

// environment_url contains the URL at which the environment can be accessed. This
// field is only set if the environment is running.
type EnvironmentListResponseEnvironmentsStatusEnvironmentURLs struct {
	// logs is the URL at which the environment logs can be accessed.
	Logs  string                                                         `json:"logs"`
	Ports []EnvironmentListResponseEnvironmentsStatusEnvironmentURLsPort `json:"ports"`
	// SSH is the URL at which the environment can be accessed via SSH.
	SSH  EnvironmentListResponseEnvironmentsStatusEnvironmentURLsSSH  `json:"ssh"`
	JSON environmentListResponseEnvironmentsStatusEnvironmentURLsJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusEnvironmentURLsJSON contains the JSON
// metadata for the struct
// [EnvironmentListResponseEnvironmentsStatusEnvironmentURLs]
type environmentListResponseEnvironmentsStatusEnvironmentURLsJSON struct {
	Logs        apijson.Field
	Ports       apijson.Field
	SSH         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsStatusEnvironmentURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusEnvironmentURLsJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseEnvironmentsStatusEnvironmentURLsPort struct {
	// port is the port number of the environment port
	Port int64 `json:"port"`
	// url is the URL at which the environment port can be accessed
	URL  string                                                           `json:"url"`
	JSON environmentListResponseEnvironmentsStatusEnvironmentURLsPortJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusEnvironmentURLsPortJSON contains the
// JSON metadata for the struct
// [EnvironmentListResponseEnvironmentsStatusEnvironmentURLsPort]
type environmentListResponseEnvironmentsStatusEnvironmentURLsPortJSON struct {
	Port        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsStatusEnvironmentURLsPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusEnvironmentURLsPortJSON) RawJSON() string {
	return r.raw
}

// SSH is the URL at which the environment can be accessed via SSH.
type EnvironmentListResponseEnvironmentsStatusEnvironmentURLsSSH struct {
	URL  string                                                          `json:"url"`
	JSON environmentListResponseEnvironmentsStatusEnvironmentURLsSSHJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusEnvironmentURLsSSHJSON contains the
// JSON metadata for the struct
// [EnvironmentListResponseEnvironmentsStatusEnvironmentURLsSSH]
type environmentListResponseEnvironmentsStatusEnvironmentURLsSSHJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsStatusEnvironmentURLsSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusEnvironmentURLsSSHJSON) RawJSON() string {
	return r.raw
}

// machine contains the status of the environment machine
type EnvironmentListResponseEnvironmentsStatusMachine struct {
	// failure_message contains the reason the machine failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the environment machine
	Phase EnvironmentListResponseEnvironmentsStatusMachinePhase `json:"phase"`
	// session is the session that is currently active in the machine.
	Session string `json:"session"`
	// timeout contains the reason the environment has timed out. If this field is
	// empty, the environment has not timed out.
	Timeout string `json:"timeout"`
	// versions contains the versions of components in the machine.
	Versions EnvironmentListResponseEnvironmentsStatusMachineVersions `json:"versions"`
	// warning_message contains warnings, e.g. when the machine is present but not in
	// the expected state.
	WarningMessage string                                               `json:"warningMessage"`
	JSON           environmentListResponseEnvironmentsStatusMachineJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusMachineJSON contains the JSON metadata
// for the struct [EnvironmentListResponseEnvironmentsStatusMachine]
type environmentListResponseEnvironmentsStatusMachineJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	Timeout        apijson.Field
	Versions       apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsStatusMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusMachineJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the environment machine
type EnvironmentListResponseEnvironmentsStatusMachinePhase string

const (
	EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseUnspecified EnvironmentListResponseEnvironmentsStatusMachinePhase = "PHASE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseCreating    EnvironmentListResponseEnvironmentsStatusMachinePhase = "PHASE_CREATING"
	EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseStarting    EnvironmentListResponseEnvironmentsStatusMachinePhase = "PHASE_STARTING"
	EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseRunning     EnvironmentListResponseEnvironmentsStatusMachinePhase = "PHASE_RUNNING"
	EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseStopping    EnvironmentListResponseEnvironmentsStatusMachinePhase = "PHASE_STOPPING"
	EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseStopped     EnvironmentListResponseEnvironmentsStatusMachinePhase = "PHASE_STOPPED"
	EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseDeleting    EnvironmentListResponseEnvironmentsStatusMachinePhase = "PHASE_DELETING"
	EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseDeleted     EnvironmentListResponseEnvironmentsStatusMachinePhase = "PHASE_DELETED"
)

func (r EnvironmentListResponseEnvironmentsStatusMachinePhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseUnspecified, EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseCreating, EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseStarting, EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseRunning, EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseStopping, EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseStopped, EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseDeleting, EnvironmentListResponseEnvironmentsStatusMachinePhasePhaseDeleted:
		return true
	}
	return false
}

// versions contains the versions of components in the machine.
type EnvironmentListResponseEnvironmentsStatusMachineVersions struct {
	SupervisorCommit  string                                                       `json:"supervisorCommit"`
	SupervisorVersion string                                                       `json:"supervisorVersion"`
	JSON              environmentListResponseEnvironmentsStatusMachineVersionsJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusMachineVersionsJSON contains the JSON
// metadata for the struct
// [EnvironmentListResponseEnvironmentsStatusMachineVersions]
type environmentListResponseEnvironmentsStatusMachineVersionsJSON struct {
	SupervisorCommit  apijson.Field
	SupervisorVersion apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsStatusMachineVersions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusMachineVersionsJSON) RawJSON() string {
	return r.raw
}

// the phase of an environment is a simple, high-level summary of where the
// environment is in its lifecycle
type EnvironmentListResponseEnvironmentsStatusPhase string

const (
	EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseUnspecified EnvironmentListResponseEnvironmentsStatusPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseCreating    EnvironmentListResponseEnvironmentsStatusPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseStarting    EnvironmentListResponseEnvironmentsStatusPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseRunning     EnvironmentListResponseEnvironmentsStatusPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseUpdating    EnvironmentListResponseEnvironmentsStatusPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseStopping    EnvironmentListResponseEnvironmentsStatusPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseStopped     EnvironmentListResponseEnvironmentsStatusPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseDeleting    EnvironmentListResponseEnvironmentsStatusPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseDeleted     EnvironmentListResponseEnvironmentsStatusPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentListResponseEnvironmentsStatusPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseUnspecified, EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseCreating, EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseStarting, EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseRunning, EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseUpdating, EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseStopping, EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseStopped, EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseDeleting, EnvironmentListResponseEnvironmentsStatusPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// RunnerACK is the acknowledgement from the runner that is has received the
// environment spec.
type EnvironmentListResponseEnvironmentsStatusRunnerAck struct {
	Message     string                                                             `json:"message"`
	SpecVersion EnvironmentListResponseEnvironmentsStatusRunnerAckSpecVersionUnion `json:"specVersion"`
	StatusCode  EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCode       `json:"statusCode"`
	JSON        environmentListResponseEnvironmentsStatusRunnerAckJSON             `json:"-"`
}

// environmentListResponseEnvironmentsStatusRunnerAckJSON contains the JSON
// metadata for the struct [EnvironmentListResponseEnvironmentsStatusRunnerAck]
type environmentListResponseEnvironmentsStatusRunnerAckJSON struct {
	Message     apijson.Field
	SpecVersion apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsStatusRunnerAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusRunnerAckJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentListResponseEnvironmentsStatusRunnerAckSpecVersionUnion interface {
	ImplementsEnvironmentListResponseEnvironmentsStatusRunnerAckSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentListResponseEnvironmentsStatusRunnerAckSpecVersionUnion)(nil)).Elem(),
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

type EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCode string

const (
	EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCodeStatusCodeUnspecified        EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCode = "STATUS_CODE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCodeStatusCodeOk                 EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCode = "STATUS_CODE_OK"
	EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCodeStatusCodeInvalidResource    EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCode = "STATUS_CODE_INVALID_RESOURCE"
	EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCodeStatusCodeFailedPrecondition EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCode = "STATUS_CODE_FAILED_PRECONDITION"
)

func (r EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCode) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCodeStatusCodeUnspecified, EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCodeStatusCodeOk, EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCodeStatusCodeInvalidResource, EnvironmentListResponseEnvironmentsStatusRunnerAckStatusCodeStatusCodeFailedPrecondition:
		return true
	}
	return false
}

type EnvironmentListResponseEnvironmentsStatusSecret struct {
	// failure_message contains the reason the secret failed to be materialize.
	FailureMessage string                                                `json:"failureMessage"`
	Phase          EnvironmentListResponseEnvironmentsStatusSecretsPhase `json:"phase"`
	SecretName     string                                                `json:"secretName"`
	// warning_message contains warnings, e.g. when the secret is present but not in
	// the expected state.
	WarningMessage string                                              `json:"warningMessage"`
	JSON           environmentListResponseEnvironmentsStatusSecretJSON `json:"-"`
}

// environmentListResponseEnvironmentsStatusSecretJSON contains the JSON metadata
// for the struct [EnvironmentListResponseEnvironmentsStatusSecret]
type environmentListResponseEnvironmentsStatusSecretJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	SecretName     apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsStatusSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusSecretJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseEnvironmentsStatusSecretsPhase string

const (
	EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseUnspecified  EnvironmentListResponseEnvironmentsStatusSecretsPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseCreating     EnvironmentListResponseEnvironmentsStatusSecretsPhase = "CONTENT_PHASE_CREATING"
	EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseInitializing EnvironmentListResponseEnvironmentsStatusSecretsPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseReady        EnvironmentListResponseEnvironmentsStatusSecretsPhase = "CONTENT_PHASE_READY"
	EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseUpdating     EnvironmentListResponseEnvironmentsStatusSecretsPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseFailed       EnvironmentListResponseEnvironmentsStatusSecretsPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentListResponseEnvironmentsStatusSecretsPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseUnspecified, EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseCreating, EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseInitializing, EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseReady, EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseUpdating, EnvironmentListResponseEnvironmentsStatusSecretsPhaseContentPhaseFailed:
		return true
	}
	return false
}

type EnvironmentListResponseEnvironmentsStatusSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// phase is the current phase of the public key
	Phase EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhase `json:"phase"`
	JSON  environmentListResponseEnvironmentsStatusSSHPublicKeyJSON   `json:"-"`
}

// environmentListResponseEnvironmentsStatusSSHPublicKeyJSON contains the JSON
// metadata for the struct [EnvironmentListResponseEnvironmentsStatusSSHPublicKey]
type environmentListResponseEnvironmentsStatusSSHPublicKeyJSON struct {
	ID          apijson.Field
	Phase       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironmentsStatusSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentsStatusSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the public key
type EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhase string

const (
	EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseUnspecified  EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseCreating     EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhase = "CONTENT_PHASE_CREATING"
	EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseInitializing EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseReady        EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhase = "CONTENT_PHASE_READY"
	EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseUpdating     EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseFailed       EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhase) IsKnown() bool {
	switch r {
	case EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseUnspecified, EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseCreating, EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseInitializing, EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseReady, EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseUpdating, EnvironmentListResponseEnvironmentsStatusSSHPublicKeysPhaseContentPhaseFailed:
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
type EnvironmentListResponseEnvironmentsStatusStatusVersionUnion interface {
	ImplementsEnvironmentListResponseEnvironmentsStatusStatusVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentListResponseEnvironmentsStatusStatusVersionUnion)(nil)).Elem(),
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

// pagination contains the pagination options for listing environments
type EnvironmentListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                `json:"nextToken"`
	JSON      environmentListResponsePaginationJSON `json:"-"`
}

// environmentListResponsePaginationJSON contains the JSON metadata for the struct
// [EnvironmentListResponsePagination]
type environmentListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

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
	Secrets []EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnion `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion EnvironmentNewFromProjectResponseEnvironmentSpecSpecVersionUnion `json:"specVersion"`
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
	// environment, relative to the repo root.
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
	Specs []EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnion `json:"specs"`
	JSON  environmentNewFromProjectResponseEnvironmentSpecContentInitializerJSON         `json:"-"`
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

// Union satisfied by
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnknown],
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnknown]
// or
// [EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnknown].
type EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnion interface {
	implementsEnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentNewFromProjectResponseEnvironmentSpecContentInitializerSpecsUnion)(nil)).Elem(), "")
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
	// root
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

// Union satisfied by
// [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnknown],
// [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnknown],
// [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnknown] or
// [EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnknown].
type EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnion interface {
	implementsEnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EnvironmentNewFromProjectResponseEnvironmentSpecSecretsUnion)(nil)).Elem(), "")
}

// version of the spec. The value of this field has no semantic meaning (e.g. don't
// interpret it as as a timestamp), but it can be used to impose a partial order.
// If a.spec_version < b.spec_version then a was the spec before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentNewFromProjectResponseEnvironmentSpecSpecVersionUnion interface {
	ImplementsEnvironmentNewFromProjectResponseEnvironmentSpecSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentNewFromProjectResponseEnvironmentSpecSpecVersionUnion)(nil)).Elem(),
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
	// but their statuus has different versions. The value of this field has no
	// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
	// to impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion EnvironmentNewFromProjectResponseEnvironmentStatusStatusVersionUnion `json:"statusVersion"`
	// warning_message contains warnings, e.g. when the environment is present but not
	// in the expected state.
	WarningMessage []string                                               `json:"warningMessage"`
	JSON           environmentNewFromProjectResponseEnvironmentStatusJSON `json:"-"`
}

// environmentNewFromProjectResponseEnvironmentStatusJSON contains the JSON
// metadata for the struct [EnvironmentNewFromProjectResponseEnvironmentStatus]
type environmentNewFromProjectResponseEnvironmentStatusJSON struct {
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
	Message     string                                                                      `json:"message"`
	SpecVersion EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckSpecVersionUnion `json:"specVersion"`
	StatusCode  EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckStatusCode       `json:"statusCode"`
	JSON        environmentNewFromProjectResponseEnvironmentStatusRunnerAckJSON             `json:"-"`
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

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckSpecVersionUnion interface {
	ImplementsEnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentNewFromProjectResponseEnvironmentStatusRunnerAckSpecVersionUnion)(nil)).Elem(),
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

// version of the status update. Environment instances themselves are unversioned,
// but their statuus has different versions. The value of this field has no
// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
// to impose a partial order. If a.status_version < b.status_version then a was the
// status before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type EnvironmentNewFromProjectResponseEnvironmentStatusStatusVersionUnion interface {
	ImplementsEnvironmentNewFromProjectResponseEnvironmentStatusStatusVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EnvironmentNewFromProjectResponseEnvironmentStatusStatusVersionUnion)(nil)).Elem(),
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

type EnvironmentStartResponse = interface{}

type EnvironmentNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec param.Field[EnvironmentNewParamsSpec] `json:"spec"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentNewParamsConnectProtocolVersion float64

const (
	EnvironmentNewParamsConnectProtocolVersion1 EnvironmentNewParamsConnectProtocolVersion = 1
)

func (r EnvironmentNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentNewParamsConnectProtocolVersion1:
		return true
	}
	return false
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
	SpecVersion param.Field[EnvironmentNewParamsSpecSpecVersionUnion] `json:"specVersion"`
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
	// environment, relative to the repo root.
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

// Satisfied by [EnvironmentNewParamsSpecContentInitializerSpecsUnknown],
// [EnvironmentNewParamsSpecContentInitializerSpecsUnknown],
// [EnvironmentNewParamsSpecContentInitializerSpecsUnknown].
type EnvironmentNewParamsSpecContentInitializerSpecUnion interface {
	implementsEnvironmentNewParamsSpecContentInitializerSpecUnion()
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
	// root
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

// Satisfied by [EnvironmentNewParamsSpecSecretsUnknown],
// [EnvironmentNewParamsSpecSecretsUnknown],
// [EnvironmentNewParamsSpecSecretsUnknown],
// [EnvironmentNewParamsSpecSecretsUnknown].
type EnvironmentNewParamsSpecSecretUnion interface {
	implementsEnvironmentNewParamsSpecSecretUnion()
}

// version of the spec. The value of this field has no semantic meaning (e.g. don't
// interpret it as as a timestamp), but it can be used to impose a partial order.
// If a.spec_version < b.spec_version then a was the spec before b.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat].
type EnvironmentNewParamsSpecSpecVersionUnion interface {
	ImplementsEnvironmentNewParamsSpecSpecVersionUnion()
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
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// environment_id specifies the environment to get
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentGetParamsConnectProtocolVersion float64

const (
	EnvironmentGetParamsConnectProtocolVersion1 EnvironmentGetParamsConnectProtocolVersion = 1
)

func (r EnvironmentGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EnvironmentListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Filter                 param.Field[EnvironmentListParamsFilter]                 `json:"filter"`
	// organization_id is the ID of the organization that contains the environments
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
	// pagination contains the pagination options for listing environments
	Pagination param.Field[EnvironmentListParamsPagination] `json:"pagination"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentListParamsConnectProtocolVersion float64

const (
	EnvironmentListParamsConnectProtocolVersion1 EnvironmentListParamsConnectProtocolVersion = 1
)

func (r EnvironmentListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentListParamsConnectProtocolVersion1:
		return true
	}
	return false
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
	// actual_phases is a list of phases the environment must be in for it to be
	// returned in the API call
	StatusPhases param.Field[[]EnvironmentListParamsFilterStatusPhase] `json:"statusPhases"`
}

func (r EnvironmentListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type EnvironmentNewFromProjectParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentNewFromProjectParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ProjectID              param.Field[string]                                                `json:"projectId" format:"uuid"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec param.Field[EnvironmentNewFromProjectParamsSpec] `json:"spec"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentNewFromProjectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentNewFromProjectParamsConnectProtocolVersion float64

const (
	EnvironmentNewFromProjectParamsConnectProtocolVersion1 EnvironmentNewFromProjectParamsConnectProtocolVersion = 1
)

func (r EnvironmentNewFromProjectParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentNewFromProjectParamsConnectProtocolVersion1:
		return true
	}
	return false
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
	SpecVersion param.Field[EnvironmentNewFromProjectParamsSpecSpecVersionUnion] `json:"specVersion"`
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
	// environment, relative to the repo root.
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

// Satisfied by
// [EnvironmentNewFromProjectParamsSpecContentInitializerSpecsUnknown],
// [EnvironmentNewFromProjectParamsSpecContentInitializerSpecsUnknown],
// [EnvironmentNewFromProjectParamsSpecContentInitializerSpecsUnknown].
type EnvironmentNewFromProjectParamsSpecContentInitializerSpecUnion interface {
	implementsEnvironmentNewFromProjectParamsSpecContentInitializerSpecUnion()
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
	// root
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

// Satisfied by [EnvironmentNewFromProjectParamsSpecSecretsUnknown],
// [EnvironmentNewFromProjectParamsSpecSecretsUnknown],
// [EnvironmentNewFromProjectParamsSpecSecretsUnknown],
// [EnvironmentNewFromProjectParamsSpecSecretsUnknown].
type EnvironmentNewFromProjectParamsSpecSecretUnion interface {
	implementsEnvironmentNewFromProjectParamsSpecSecretUnion()
}

// version of the spec. The value of this field has no semantic meaning (e.g. don't
// interpret it as as a timestamp), but it can be used to impose a partial order.
// If a.spec_version < b.spec_version then a was the spec before b.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat].
type EnvironmentNewFromProjectParamsSpecSpecVersionUnion interface {
	ImplementsEnvironmentNewFromProjectParamsSpecSpecVersionUnion()
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

type EnvironmentStartParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentStartParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// environment_id specifies which environment should be started.
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EnvironmentStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EnvironmentStartParamsConnectProtocolVersion float64

const (
	EnvironmentStartParamsConnectProtocolVersion1 EnvironmentStartParamsConnectProtocolVersion = 1
)

func (r EnvironmentStartParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentStartParamsConnectProtocolVersion1:
		return true
	}
	return false
}
