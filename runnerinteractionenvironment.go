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
	"github.com/stainless-sdks/gitpod-go/shared"
	"github.com/tidwall/gjson"
)

// RunnerInteractionEnvironmentService contains methods and other services that
// help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerInteractionEnvironmentService] method instead.
type RunnerInteractionEnvironmentService struct {
	Options []option.RequestOption
}

// NewRunnerInteractionEnvironmentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRunnerInteractionEnvironmentService(opts ...option.RequestOption) (r *RunnerInteractionEnvironmentService) {
	r = &RunnerInteractionEnvironmentService{}
	r.Options = opts
	return
}

// GetRunnerEnvironment returns the environment given it is owned by the runner.
func (r *RunnerInteractionEnvironmentService) Get(ctx context.Context, params RunnerInteractionEnvironmentGetParams, opts ...option.RequestOption) (res *RunnerInteractionEnvironmentGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/GetRunnerEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListRunnerEnvironments returns the environments this runner is responsible for.
func (r *RunnerInteractionEnvironmentService) List(ctx context.Context, params RunnerInteractionEnvironmentListParams, opts ...option.RequestOption) (res *RunnerInteractionEnvironmentListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/ListRunnerEnvironments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateRunnerEnvironmentStatus updates the status of an environment this runner
// is responsible for.
func (r *RunnerInteractionEnvironmentService) UpdateStatus(ctx context.Context, params RunnerInteractionEnvironmentUpdateStatusParams, opts ...option.RequestOption) (res *RunnerInteractionEnvironmentUpdateStatusResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/UpdateRunnerEnvironmentStatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type RunnerInteractionEnvironmentGetResponse struct {
	Environment RunnerInteractionEnvironmentGetResponseEnvironment `json:"environment"`
	JSON        runnerInteractionEnvironmentGetResponseJSON        `json:"-"`
}

// runnerInteractionEnvironmentGetResponseJSON contains the JSON metadata for the
// struct [RunnerInteractionEnvironmentGetResponse]
type runnerInteractionEnvironmentGetResponseJSON struct {
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionEnvironmentGetResponseEnvironment struct {
	// ID is a unique identifier of this environment. No other environment with the
	// same name must be managed by this environment manager
	ID string `json:"id" format:"uuid"`
	// The environment's access token
	EnvironmentAccessToken string `json:"environmentAccessToken"`
	// EnvironmentMetadata is data associated with an environment that's required for
	// other parts of the system to function
	Metadata RunnerInteractionEnvironmentGetResponseEnvironmentMetadata `json:"metadata"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec RunnerInteractionEnvironmentGetResponseEnvironmentSpec `json:"spec"`
	JSON runnerInteractionEnvironmentGetResponseEnvironmentJSON `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentJSON contains the JSON
// metadata for the struct [RunnerInteractionEnvironmentGetResponseEnvironment]
type runnerInteractionEnvironmentGetResponseEnvironmentJSON struct {
	ID                     apijson.Field
	EnvironmentAccessToken apijson.Field
	Metadata               apijson.Field
	Spec                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentGetResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentMetadata is data associated with an environment that's required for
// other parts of the system to function
type RunnerInteractionEnvironmentGetResponseEnvironmentMetadata struct {
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
	Creator RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreator `json:"creator"`
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
	RunnerID string                                                         `json:"runnerId"`
	JSON     runnerInteractionEnvironmentGetResponseEnvironmentMetadataJSON `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentMetadataJSON contains the JSON
// metadata for the struct
// [RunnerInteractionEnvironmentGetResponseEnvironmentMetadata]
type runnerInteractionEnvironmentGetResponseEnvironmentMetadataJSON struct {
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

func (r *RunnerInteractionEnvironmentGetResponseEnvironmentMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the creator of the environment
type RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipal `json:"principal"`
	JSON      runnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorJSON      `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorJSON contains
// the JSON metadata for the struct
// [RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreator]
type runnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipal string

const (
	RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalUnspecified    RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalAccount        RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalUser           RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_USER"
	RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalRunner         RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalEnvironment    RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalServiceAccount RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalUnspecified, RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalAccount, RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalUser, RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalRunner, RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalEnvironment, RunnerInteractionEnvironmentGetResponseEnvironmentMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type RunnerInteractionEnvironmentGetResponseEnvironmentSpec struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmission `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile RunnerInteractionEnvironmentGetResponseEnvironmentSpecAutomationsFile `json:"automationsFile"`
	// content is the content spec of the environment
	Content RunnerInteractionEnvironmentGetResponseEnvironmentSpecContent `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer RunnerInteractionEnvironmentGetResponseEnvironmentSpecDevcontainer `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine RunnerInteractionEnvironmentGetResponseEnvironmentSpecMachine `json:"machine"`
	// ports is the set of ports which ought to be exposed to the internet
	Ports []RunnerInteractionEnvironmentGetResponseEnvironmentSpecPort `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets []RunnerInteractionEnvironmentGetResponseEnvironmentSpecSecretsUnion `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion RunnerInteractionEnvironmentGetResponseEnvironmentSpecSpecVersionUnion `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys []RunnerInteractionEnvironmentGetResponseEnvironmentSpecSSHPublicKey `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout RunnerInteractionEnvironmentGetResponseEnvironmentSpecTimeout `json:"timeout"`
	JSON    runnerInteractionEnvironmentGetResponseEnvironmentSpecJSON    `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentSpecJSON contains the JSON
// metadata for the struct [RunnerInteractionEnvironmentGetResponseEnvironmentSpec]
type runnerInteractionEnvironmentGetResponseEnvironmentSpecJSON struct {
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

func (r *RunnerInteractionEnvironmentGetResponseEnvironmentSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentSpecJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmission string

const (
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelUnspecified RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelOwnerOnly   RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelEveryone    RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmission) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelUnspecified, RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelOwnerOnly, RunnerInteractionEnvironmentGetResponseEnvironmentSpecAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// automations_file is the automations file spec of the environment
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root.
	AutomationsFilePath string                                                                    `json:"automationsFilePath"`
	Session             string                                                                    `json:"session"`
	JSON                runnerInteractionEnvironmentGetResponseEnvironmentSpecAutomationsFileJSON `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentSpecAutomationsFileJSON
// contains the JSON metadata for the struct
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecAutomationsFile]
type runnerInteractionEnvironmentGetResponseEnvironmentSpecAutomationsFileJSON struct {
	AutomationsFilePath apijson.Field
	Session             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentGetResponseEnvironmentSpecAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentSpecAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// content is the content spec of the environment
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecContent struct {
	// The Git email address
	GitEmail string `json:"gitEmail"`
	// The Git username
	GitUsername string `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer RunnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializer `json:"initializer"`
	Session     string                                                                   `json:"session"`
	JSON        runnerInteractionEnvironmentGetResponseEnvironmentSpecContentJSON        `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentSpecContentJSON contains the
// JSON metadata for the struct
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecContent]
type runnerInteractionEnvironmentGetResponseEnvironmentSpecContentJSON struct {
	GitEmail    apijson.Field
	GitUsername apijson.Field
	Initializer apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentGetResponseEnvironmentSpecContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentSpecContentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentInitializer specifies how an environment is to be initialized
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializer struct {
	Specs []RunnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion `json:"specs"`
	JSON  runnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializerJSON         `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializerJSON
// contains the JSON metadata for the struct
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializer]
type runnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializerJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnknown],
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnknown]
// or
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnknown].
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion interface {
	implementsRunnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*RunnerInteractionEnvironmentGetResponseEnvironmentSpecContentInitializerSpecsUnion)(nil)).Elem(), "")
}

// Phase is the desired phase of the environment
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase string

const (
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUnspecified RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseCreating    RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_CREATING"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStarting    RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STARTING"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseRunning     RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_RUNNING"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUpdating    RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_UPDATING"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopping    RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPING"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopped     RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPED"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleting    RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETING"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleted     RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhase) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUnspecified, RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseCreating, RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStarting, RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseRunning, RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseUpdating, RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopping, RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseStopped, RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleting, RunnerInteractionEnvironmentGetResponseEnvironmentSpecDesiredPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// devcontainer is the devcontainer spec of the environment
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string                                                                 `json:"devcontainerFilePath"`
	Session              string                                                                 `json:"session"`
	JSON                 runnerInteractionEnvironmentGetResponseEnvironmentSpecDevcontainerJSON `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentSpecDevcontainerJSON contains
// the JSON metadata for the struct
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecDevcontainer]
type runnerInteractionEnvironmentGetResponseEnvironmentSpecDevcontainerJSON struct {
	DevcontainerFilePath apijson.Field
	Session              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentGetResponseEnvironmentSpecDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentSpecDevcontainerJSON) RawJSON() string {
	return r.raw
}

// machine is the machine spec of the environment
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecMachine struct {
	// Class denotes the class of the environment we ought to start
	Class   string                                                            `json:"class" format:"uuid"`
	Session string                                                            `json:"session"`
	JSON    runnerInteractionEnvironmentGetResponseEnvironmentSpecMachineJSON `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentSpecMachineJSON contains the
// JSON metadata for the struct
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecMachine]
type runnerInteractionEnvironmentGetResponseEnvironmentSpecMachineJSON struct {
	Class       apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentGetResponseEnvironmentSpecMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentSpecMachineJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionEnvironmentGetResponseEnvironmentSpecPort struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmission `json:"admission"`
	// name of this port
	Name string `json:"name"`
	// port number
	Port int64                                                          `json:"port"`
	JSON runnerInteractionEnvironmentGetResponseEnvironmentSpecPortJSON `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentSpecPortJSON contains the JSON
// metadata for the struct
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecPort]
type runnerInteractionEnvironmentGetResponseEnvironmentSpecPortJSON struct {
	Admission   apijson.Field
	Name        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentGetResponseEnvironmentSpecPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentSpecPortJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmission string

const (
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelUnspecified RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelOwnerOnly   RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelEveryone    RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmission) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelUnspecified, RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelOwnerOnly, RunnerInteractionEnvironmentGetResponseEnvironmentSpecPortsAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// Union satisfied by
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecSecretsUnknown],
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecSecretsUnknown],
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecSecretsUnknown] or
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecSecretsUnknown].
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecSecretsUnion interface {
	implementsRunnerInteractionEnvironmentGetResponseEnvironmentSpecSecretsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*RunnerInteractionEnvironmentGetResponseEnvironmentSpecSecretsUnion)(nil)).Elem(), "")
}

// version of the spec. The value of this field has no semantic meaning (e.g. don't
// interpret it as as a timestamp), but it can be used to impose a partial order.
// If a.spec_version < b.spec_version then a was the spec before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecSpecVersionUnion interface {
	ImplementsRunnerInteractionEnvironmentGetResponseEnvironmentSpecSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerInteractionEnvironmentGetResponseEnvironmentSpecSpecVersionUnion)(nil)).Elem(),
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

type RunnerInteractionEnvironmentGetResponseEnvironmentSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// value is the actual public key in the public key file format
	Value string                                                                 `json:"value"`
	JSON  runnerInteractionEnvironmentGetResponseEnvironmentSpecSSHPublicKeyJSON `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentSpecSSHPublicKeyJSON contains
// the JSON metadata for the struct
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecSSHPublicKey]
type runnerInteractionEnvironmentGetResponseEnvironmentSpecSSHPublicKeyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentGetResponseEnvironmentSpecSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentSpecSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// Timeout configures the environment timeout
type RunnerInteractionEnvironmentGetResponseEnvironmentSpecTimeout struct {
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
	Disconnected string                                                            `json:"disconnected" format:"regex"`
	JSON         runnerInteractionEnvironmentGetResponseEnvironmentSpecTimeoutJSON `json:"-"`
}

// runnerInteractionEnvironmentGetResponseEnvironmentSpecTimeoutJSON contains the
// JSON metadata for the struct
// [RunnerInteractionEnvironmentGetResponseEnvironmentSpecTimeout]
type runnerInteractionEnvironmentGetResponseEnvironmentSpecTimeoutJSON struct {
	Disconnected apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentGetResponseEnvironmentSpecTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentGetResponseEnvironmentSpecTimeoutJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionEnvironmentListResponse struct {
	// The environments running on the runner
	Environments []RunnerInteractionEnvironmentListResponseEnvironment `json:"environments"`
	// pagination contains the pagination options for listing environments
	Pagination RunnerInteractionEnvironmentListResponsePagination `json:"pagination"`
	JSON       runnerInteractionEnvironmentListResponseJSON       `json:"-"`
}

// runnerInteractionEnvironmentListResponseJSON contains the JSON metadata for the
// struct [RunnerInteractionEnvironmentListResponse]
type runnerInteractionEnvironmentListResponseJSON struct {
	Environments apijson.Field
	Pagination   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionEnvironmentListResponseEnvironment struct {
	// ID is a unique identifier of this environment. No other environment with the
	// same name must be managed by this environment manager
	ID string `json:"id" format:"uuid"`
	// The environment's access token
	EnvironmentAccessToken string `json:"environmentAccessToken"`
	// EnvironmentMetadata is data associated with an environment that's required for
	// other parts of the system to function
	Metadata RunnerInteractionEnvironmentListResponseEnvironmentsMetadata `json:"metadata"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec RunnerInteractionEnvironmentListResponseEnvironmentsSpec `json:"spec"`
	JSON runnerInteractionEnvironmentListResponseEnvironmentJSON  `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentJSON contains the JSON
// metadata for the struct [RunnerInteractionEnvironmentListResponseEnvironment]
type runnerInteractionEnvironmentListResponseEnvironmentJSON struct {
	ID                     apijson.Field
	EnvironmentAccessToken apijson.Field
	Metadata               apijson.Field
	Spec                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentMetadata is data associated with an environment that's required for
// other parts of the system to function
type RunnerInteractionEnvironmentListResponseEnvironmentsMetadata struct {
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
	Creator RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreator `json:"creator"`
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
	RunnerID string                                                           `json:"runnerId"`
	JSON     runnerInteractionEnvironmentListResponseEnvironmentsMetadataJSON `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentsMetadataJSON contains the
// JSON metadata for the struct
// [RunnerInteractionEnvironmentListResponseEnvironmentsMetadata]
type runnerInteractionEnvironmentListResponseEnvironmentsMetadataJSON struct {
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

func (r *RunnerInteractionEnvironmentListResponseEnvironmentsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentsMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the creator of the environment
type RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipal `json:"principal"`
	JSON      runnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorJSON      `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorJSON contains
// the JSON metadata for the struct
// [RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreator]
type runnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipal string

const (
	RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalUnspecified    RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalAccount        RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalUser           RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_USER"
	RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalRunner         RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalEnvironment    RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalServiceAccount RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalUnspecified, RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalAccount, RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalUser, RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalRunner, RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalEnvironment, RunnerInteractionEnvironmentListResponseEnvironmentsMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type RunnerInteractionEnvironmentListResponseEnvironmentsSpec struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmission `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile RunnerInteractionEnvironmentListResponseEnvironmentsSpecAutomationsFile `json:"automationsFile"`
	// content is the content spec of the environment
	Content RunnerInteractionEnvironmentListResponseEnvironmentsSpecContent `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer RunnerInteractionEnvironmentListResponseEnvironmentsSpecDevcontainer `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine RunnerInteractionEnvironmentListResponseEnvironmentsSpecMachine `json:"machine"`
	// ports is the set of ports which ought to be exposed to the internet
	Ports []RunnerInteractionEnvironmentListResponseEnvironmentsSpecPort `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets []RunnerInteractionEnvironmentListResponseEnvironmentsSpecSecretsUnion `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion RunnerInteractionEnvironmentListResponseEnvironmentsSpecSpecVersionUnion `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys []RunnerInteractionEnvironmentListResponseEnvironmentsSpecSSHPublicKey `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout RunnerInteractionEnvironmentListResponseEnvironmentsSpecTimeout `json:"timeout"`
	JSON    runnerInteractionEnvironmentListResponseEnvironmentsSpecJSON    `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentsSpecJSON contains the JSON
// metadata for the struct
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpec]
type runnerInteractionEnvironmentListResponseEnvironmentsSpecJSON struct {
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

func (r *RunnerInteractionEnvironmentListResponseEnvironmentsSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentsSpecJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmission string

const (
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelUnspecified RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelOwnerOnly   RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelEveryone    RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmission) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelUnspecified, RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelOwnerOnly, RunnerInteractionEnvironmentListResponseEnvironmentsSpecAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// automations_file is the automations file spec of the environment
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root.
	AutomationsFilePath string                                                                      `json:"automationsFilePath"`
	Session             string                                                                      `json:"session"`
	JSON                runnerInteractionEnvironmentListResponseEnvironmentsSpecAutomationsFileJSON `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentsSpecAutomationsFileJSON
// contains the JSON metadata for the struct
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecAutomationsFile]
type runnerInteractionEnvironmentListResponseEnvironmentsSpecAutomationsFileJSON struct {
	AutomationsFilePath apijson.Field
	Session             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponseEnvironmentsSpecAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentsSpecAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// content is the content spec of the environment
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecContent struct {
	// The Git email address
	GitEmail string `json:"gitEmail"`
	// The Git username
	GitUsername string `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer RunnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializer `json:"initializer"`
	Session     string                                                                     `json:"session"`
	JSON        runnerInteractionEnvironmentListResponseEnvironmentsSpecContentJSON        `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentsSpecContentJSON contains the
// JSON metadata for the struct
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecContent]
type runnerInteractionEnvironmentListResponseEnvironmentsSpecContentJSON struct {
	GitEmail    apijson.Field
	GitUsername apijson.Field
	Initializer apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponseEnvironmentsSpecContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentsSpecContentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentInitializer specifies how an environment is to be initialized
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializer struct {
	Specs []RunnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnion `json:"specs"`
	JSON  runnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializerJSON         `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializerJSON
// contains the JSON metadata for the struct
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializer]
type runnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializerJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnknown],
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnknown]
// or
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnknown].
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnion interface {
	implementsRunnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*RunnerInteractionEnvironmentListResponseEnvironmentsSpecContentInitializerSpecsUnion)(nil)).Elem(), "")
}

// Phase is the desired phase of the environment
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase string

const (
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseUnspecified RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseCreating    RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_CREATING"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStarting    RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_STARTING"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseRunning     RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_RUNNING"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseUpdating    RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_UPDATING"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStopping    RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPING"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStopped     RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_STOPPED"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseDeleting    RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETING"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseDeleted     RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhase) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseUnspecified, RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseCreating, RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStarting, RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseRunning, RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseUpdating, RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStopping, RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseStopped, RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseDeleting, RunnerInteractionEnvironmentListResponseEnvironmentsSpecDesiredPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// devcontainer is the devcontainer spec of the environment
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string                                                                   `json:"devcontainerFilePath"`
	Session              string                                                                   `json:"session"`
	JSON                 runnerInteractionEnvironmentListResponseEnvironmentsSpecDevcontainerJSON `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentsSpecDevcontainerJSON
// contains the JSON metadata for the struct
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecDevcontainer]
type runnerInteractionEnvironmentListResponseEnvironmentsSpecDevcontainerJSON struct {
	DevcontainerFilePath apijson.Field
	Session              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponseEnvironmentsSpecDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentsSpecDevcontainerJSON) RawJSON() string {
	return r.raw
}

// machine is the machine spec of the environment
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecMachine struct {
	// Class denotes the class of the environment we ought to start
	Class   string                                                              `json:"class" format:"uuid"`
	Session string                                                              `json:"session"`
	JSON    runnerInteractionEnvironmentListResponseEnvironmentsSpecMachineJSON `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentsSpecMachineJSON contains the
// JSON metadata for the struct
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecMachine]
type runnerInteractionEnvironmentListResponseEnvironmentsSpecMachineJSON struct {
	Class       apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponseEnvironmentsSpecMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentsSpecMachineJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionEnvironmentListResponseEnvironmentsSpecPort struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmission `json:"admission"`
	// name of this port
	Name string `json:"name"`
	// port number
	Port int64                                                            `json:"port"`
	JSON runnerInteractionEnvironmentListResponseEnvironmentsSpecPortJSON `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentsSpecPortJSON contains the
// JSON metadata for the struct
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecPort]
type runnerInteractionEnvironmentListResponseEnvironmentsSpecPortJSON struct {
	Admission   apijson.Field
	Name        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponseEnvironmentsSpecPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentsSpecPortJSON) RawJSON() string {
	return r.raw
}

// Admission level describes who can access an environment instance and its ports.
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmission string

const (
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelUnspecified RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmission = "ADMISSION_LEVEL_UNSPECIFIED"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelOwnerOnly   RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmission = "ADMISSION_LEVEL_OWNER_ONLY"
	RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelEveryone    RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmission = "ADMISSION_LEVEL_EVERYONE"
)

func (r RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmission) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelUnspecified, RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelOwnerOnly, RunnerInteractionEnvironmentListResponseEnvironmentsSpecPortsAdmissionAdmissionLevelEveryone:
		return true
	}
	return false
}

// Union satisfied by
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecSecretsUnknown],
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecSecretsUnknown],
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecSecretsUnknown] or
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecSecretsUnknown].
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecSecretsUnion interface {
	implementsRunnerInteractionEnvironmentListResponseEnvironmentsSpecSecretsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*RunnerInteractionEnvironmentListResponseEnvironmentsSpecSecretsUnion)(nil)).Elem(), "")
}

// version of the spec. The value of this field has no semantic meaning (e.g. don't
// interpret it as as a timestamp), but it can be used to impose a partial order.
// If a.spec_version < b.spec_version then a was the spec before b.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecSpecVersionUnion interface {
	ImplementsRunnerInteractionEnvironmentListResponseEnvironmentsSpecSpecVersionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerInteractionEnvironmentListResponseEnvironmentsSpecSpecVersionUnion)(nil)).Elem(),
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

type RunnerInteractionEnvironmentListResponseEnvironmentsSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// value is the actual public key in the public key file format
	Value string                                                                   `json:"value"`
	JSON  runnerInteractionEnvironmentListResponseEnvironmentsSpecSSHPublicKeyJSON `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentsSpecSSHPublicKeyJSON
// contains the JSON metadata for the struct
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecSSHPublicKey]
type runnerInteractionEnvironmentListResponseEnvironmentsSpecSSHPublicKeyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponseEnvironmentsSpecSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentsSpecSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// Timeout configures the environment timeout
type RunnerInteractionEnvironmentListResponseEnvironmentsSpecTimeout struct {
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
	Disconnected string                                                              `json:"disconnected" format:"regex"`
	JSON         runnerInteractionEnvironmentListResponseEnvironmentsSpecTimeoutJSON `json:"-"`
}

// runnerInteractionEnvironmentListResponseEnvironmentsSpecTimeoutJSON contains the
// JSON metadata for the struct
// [RunnerInteractionEnvironmentListResponseEnvironmentsSpecTimeout]
type runnerInteractionEnvironmentListResponseEnvironmentsSpecTimeoutJSON struct {
	Disconnected apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponseEnvironmentsSpecTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponseEnvironmentsSpecTimeoutJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing environments
type RunnerInteractionEnvironmentListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                                 `json:"nextToken"`
	JSON      runnerInteractionEnvironmentListResponsePaginationJSON `json:"-"`
}

// runnerInteractionEnvironmentListResponsePaginationJSON contains the JSON
// metadata for the struct [RunnerInteractionEnvironmentListResponsePagination]
type runnerInteractionEnvironmentListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionEnvironmentListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionEnvironmentListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionEnvironmentUpdateStatusResponse = interface{}

type RunnerInteractionEnvironmentGetParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionEnvironmentGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// The environment's ID
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionEnvironmentGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionEnvironmentGetParamsConnectProtocolVersion float64

const (
	RunnerInteractionEnvironmentGetParamsConnectProtocolVersion1 RunnerInteractionEnvironmentGetParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionEnvironmentGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionEnvironmentListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionEnvironmentListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// An optional list of environment IDs to fetch. If this list is empty/not provided
	// all environments that ought to run on the runner are returned.
	EnvironmentIDs param.Field[[]string] `json:"environmentIds" format:"uuid"`
	// pagination contains the pagination options for listing environments
	Pagination param.Field[RunnerInteractionEnvironmentListParamsPagination] `json:"pagination"`
	// The runner's identifier
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionEnvironmentListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionEnvironmentListParamsConnectProtocolVersion float64

const (
	RunnerInteractionEnvironmentListParamsConnectProtocolVersion1 RunnerInteractionEnvironmentListParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionEnvironmentListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// pagination contains the pagination options for listing environments
type RunnerInteractionEnvironmentListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerInteractionEnvironmentListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionEnvironmentUpdateStatusParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionEnvironmentUpdateStatusParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// The environment's ID
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// EnvironmentStatus describes an environment status
	Status param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatus] `json:"status"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionEnvironmentUpdateStatusParamsConnectProtocolVersion float64

const (
	RunnerInteractionEnvironmentUpdateStatusParamsConnectProtocolVersion1 RunnerInteractionEnvironmentUpdateStatusParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// EnvironmentStatus describes an environment status
type RunnerInteractionEnvironmentUpdateStatusParamsStatus struct {
	// automations_file contains the status of the automations file.
	AutomationsFile param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFile] `json:"automationsFile"`
	// content contains the status of the environment content.
	Content param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusContent] `json:"content"`
	// devcontainer contains the status of the devcontainer.
	Devcontainer param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainer] `json:"devcontainer"`
	// environment_url contains the URL at which the environment can be accessed. This
	// field is only set if the environment is running.
	EnvironmentURLs param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLs] `json:"environmentUrls"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage param.Field[[]string] `json:"failureMessage"`
	// machine contains the status of the environment machine
	Machine param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusMachine] `json:"machine"`
	// the phase of an environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase] `json:"phase"`
	// RunnerACK is the acknowledgement from the runner that is has received the
	// environment spec.
	RunnerAck param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAck] `json:"runnerAck"`
	// secrets contains the status of the environment secrets
	Secrets param.Field[[]RunnerInteractionEnvironmentUpdateStatusParamsStatusSecret] `json:"secrets"`
	// ssh_public_keys contains the status of the environment ssh public keys
	SSHPublicKeys param.Field[[]RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKey] `json:"sshPublicKeys"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their statuus has different versions. The value of this field has no
	// semantic meaning (e.g. don't interpret it as as a timestemp), but it can be used
	// to impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusStatusVersionUnion] `json:"statusVersion"`
	// warning_message contains warnings, e.g. when the environment is present but not
	// in the expected state.
	WarningMessage param.Field[[]string] `json:"warningMessage"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// automations_file contains the status of the automations file.
type RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFile struct {
	// automations_file_path is the path to the automations file relative to the repo
	// root.
	AutomationsFilePath param.Field[string] `json:"automationsFilePath"`
	// automations_file_presence indicates how an automations file is present in the
	// environment.
	AutomationsFilePresence param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresence] `json:"automationsFilePresence"`
	// failure_message contains the reason the automations file failed to be applied.
	// This is only set if the phase is FAILED.
	FailureMessage param.Field[string] `json:"failureMessage"`
	// phase is the current phase of the automations file.
	Phase param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhase] `json:"phase"`
	// session is the automations file session that is currently applied in the
	// environment.
	Session param.Field[string] `json:"session"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// automations_file_presence indicates how an automations file is present in the
// environment.
type RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresence string

const (
	RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresencePresenceUnspecified RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresence = "PRESENCE_UNSPECIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresencePresenceAbsent      RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresence = "PRESENCE_ABSENT"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresencePresenceDiscovered  RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresence = "PRESENCE_DISCOVERED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresencePresenceSpecified   RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresence = "PRESENCE_SPECIFIED"
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresence) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresencePresenceUnspecified, RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresencePresenceAbsent, RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresencePresenceDiscovered, RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the automations file.
type RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhase string

const (
	RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseUnspecified  RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhase = "CONTENT_PHASE_UNSPECIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseCreating     RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhase = "CONTENT_PHASE_CREATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseInitializing RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhase = "CONTENT_PHASE_INITIALIZING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseReady        RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhase = "CONTENT_PHASE_READY"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseUpdating     RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhase = "CONTENT_PHASE_UPDATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseFailed       RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhase = "CONTENT_PHASE_FAILED"
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhase) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseUnspecified, RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseCreating, RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseInitializing, RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseReady, RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseUpdating, RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseFailed:
		return true
	}
	return false
}

// content contains the status of the environment content.
type RunnerInteractionEnvironmentUpdateStatusParamsStatusContent struct {
	// content_location_in_machine is the location of the content in the machine
	ContentLocationInMachine param.Field[string] `json:"contentLocationInMachine"`
	// failure_message contains the reason the content initialization failed.
	FailureMessage param.Field[string] `json:"failureMessage"`
	// git is the Git working copy status of the environment. Note: this is a
	// best-effort field and more often than not will not be present. Its absence does
	// not indicate the absence of a working copy.
	Git param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGit] `json:"git"`
	// phase is the current phase of the environment content
	Phase param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhase] `json:"phase"`
	// session is the session that is currently active in the environment.
	Session param.Field[string] `json:"session"`
	// warning_message contains warnings, e.g. when the content is present but not in
	// the expected state.
	WarningMessage param.Field[string] `json:"warningMessage"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// git is the Git working copy status of the environment. Note: this is a
// best-effort field and more often than not will not be present. Its absence does
// not indicate the absence of a working copy.
type RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGit struct {
	// branch is branch we're currently on
	Branch param.Field[string] `json:"branch"`
	// changed_files is an array of changed files in the environment, possibly
	// truncated
	ChangedFiles param.Field[[]RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFile] `json:"changedFiles"`
	// clone_url is the repository url as you would pass it to "git clone". Only HTTPS
	// clone URLs are supported.
	CloneURL param.Field[string] `json:"cloneUrl"`
	// latest_commit is the most recent commit on the current branch
	LatestCommit      param.Field[string] `json:"latestCommit"`
	TotalChangedFiles param.Field[int64]  `json:"totalChangedFiles"`
	// the total number of unpushed changes
	TotalUnpushedCommits param.Field[int64] `json:"totalUnpushedCommits"`
	// unpushed_commits is an array of unpushed changes in the environment, possibly
	// truncated
	UnpushedCommits param.Field[[]string] `json:"unpushedCommits"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFile struct {
	// ChangeType is the type of change that happened to the file
	ChangeType param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeType] `json:"changeType"`
	// path is the path of the file
	Path param.Field[string] `json:"path"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ChangeType is the type of change that happened to the file
type RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeType string

const (
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified        RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNSPECIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeAdded              RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_ADDED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeModified           RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_MODIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeDeleted            RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_DELETED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeRenamed            RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_RENAMED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeCopied             RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_COPIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UPDATED_BUT_UNMERGED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeUntracked          RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNTRACKED"
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeType) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeAdded, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeModified, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeDeleted, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeRenamed, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeCopied, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeUntracked:
		return true
	}
	return false
}

// phase is the current phase of the environment content
type RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhase string

const (
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseUnspecified  RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhase = "CONTENT_PHASE_UNSPECIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseCreating     RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhase = "CONTENT_PHASE_CREATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseInitializing RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhase = "CONTENT_PHASE_INITIALIZING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseReady        RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhase = "CONTENT_PHASE_READY"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseUpdating     RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhase = "CONTENT_PHASE_UPDATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseFailed       RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhase = "CONTENT_PHASE_FAILED"
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhase) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseUnspecified, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseCreating, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseInitializing, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseReady, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseUpdating, RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseFailed:
		return true
	}
	return false
}

// devcontainer contains the status of the devcontainer.
type RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainer struct {
	// container_id is the ID of the container.
	ContainerID param.Field[string] `json:"containerId"`
	// container_name is the name of the container that is used to connect to the
	// devcontainer
	ContainerName param.Field[string] `json:"containerName"`
	// devcontainerconfig_in_sync indicates if the devcontainer is up to date w.r.t.
	// the devcontainer config file.
	DevcontainerconfigInSync param.Field[bool] `json:"devcontainerconfigInSync"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath"`
	// devcontainer_file_presence indicates how the devcontainer file is present in the
	// repo.
	DevcontainerFilePresence param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresence] `json:"devcontainerFilePresence"`
	// failure_message contains the reason the devcontainer failed to operate.
	FailureMessage param.Field[string] `json:"failureMessage"`
	// phase is the current phase of the devcontainer
	Phase param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhase] `json:"phase"`
	// remote_user is the user that is used to connect to the devcontainer
	RemoteUser param.Field[string] `json:"remoteUser"`
	// remote_workspace_folder is the folder that is used to connect to the
	// devcontainer
	RemoteWorkspaceFolder param.Field[string] `json:"remoteWorkspaceFolder"`
	// secrets_in_sync indicates if the secrets are up to date w.r.t. the running
	// devcontainer.
	SecretsInSync param.Field[bool] `json:"secretsInSync"`
	// session is the session that is currently active in the devcontainer.
	Session param.Field[string] `json:"session"`
	// warning_message contains warnings, e.g. when the devcontainer is present but not
	// in the expected state.
	WarningMessage param.Field[string] `json:"warningMessage"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// devcontainer_file_presence indicates how the devcontainer file is present in the
// repo.
type RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresence string

const (
	RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresencePresenceUnspecified RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresence = "PRESENCE_UNSPECIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresencePresenceGenerated   RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresence = "PRESENCE_GENERATED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresencePresenceDiscovered  RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresence = "PRESENCE_DISCOVERED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresencePresenceSpecified   RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresence = "PRESENCE_SPECIFIED"
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresence) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresencePresenceUnspecified, RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresencePresenceGenerated, RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresencePresenceDiscovered, RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the devcontainer
type RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhase string

const (
	RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhasePhaseUnspecified RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhase = "PHASE_UNSPECIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhasePhaseCreating    RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhase = "PHASE_CREATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhasePhaseRunning     RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhase = "PHASE_RUNNING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhasePhaseStopped     RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhase = "PHASE_STOPPED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhasePhaseFailed      RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhase = "PHASE_FAILED"
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhase) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhasePhaseUnspecified, RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhasePhaseCreating, RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhasePhaseRunning, RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhasePhaseStopped, RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhasePhaseFailed:
		return true
	}
	return false
}

// environment_url contains the URL at which the environment can be accessed. This
// field is only set if the environment is running.
type RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLs struct {
	// logs is the URL at which the environment logs can be accessed.
	Logs  param.Field[string]                                                                    `json:"logs"`
	Ports param.Field[[]RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLsPort] `json:"ports"`
	// SSH is the URL at which the environment can be accessed via SSH.
	SSH param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLsSSH] `json:"ssh"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLsPort struct {
	// port is the port number of the environment port
	Port param.Field[int64] `json:"port"`
	// url is the URL at which the environment port can be accessed
	URL param.Field[string] `json:"url"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLsPort) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSH is the URL at which the environment can be accessed via SSH.
type RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLsSSH struct {
	URL param.Field[string] `json:"url"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLsSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// machine contains the status of the environment machine
type RunnerInteractionEnvironmentUpdateStatusParamsStatusMachine struct {
	// failure_message contains the reason the machine failed to operate.
	FailureMessage param.Field[string] `json:"failureMessage"`
	// phase is the current phase of the environment machine
	Phase param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhase] `json:"phase"`
	// session is the session that is currently active in the machine.
	Session param.Field[string] `json:"session"`
	// timeout contains the reason the environment has timed out. If this field is
	// empty, the environment has not timed out.
	Timeout param.Field[string] `json:"timeout"`
	// versions contains the versions of components in the machine.
	Versions param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusMachineVersions] `json:"versions"`
	// warning_message contains warnings, e.g. when the machine is present but not in
	// the expected state.
	WarningMessage param.Field[string] `json:"warningMessage"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusMachine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// phase is the current phase of the environment machine
type RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhase string

const (
	RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseUnspecified RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhase = "PHASE_UNSPECIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseCreating    RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhase = "PHASE_CREATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseStarting    RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhase = "PHASE_STARTING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseRunning     RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhase = "PHASE_RUNNING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseStopping    RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhase = "PHASE_STOPPING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseStopped     RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhase = "PHASE_STOPPED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseDeleting    RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhase = "PHASE_DELETING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseDeleted     RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhase = "PHASE_DELETED"
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhase) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseUnspecified, RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseCreating, RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseStarting, RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseRunning, RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseStopping, RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseStopped, RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseDeleting, RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseDeleted:
		return true
	}
	return false
}

// versions contains the versions of components in the machine.
type RunnerInteractionEnvironmentUpdateStatusParamsStatusMachineVersions struct {
	SupervisorCommit  param.Field[string] `json:"supervisorCommit"`
	SupervisorVersion param.Field[string] `json:"supervisorVersion"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusMachineVersions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// the phase of an environment is a simple, high-level summary of where the
// environment is in its lifecycle
type RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase string

const (
	RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseUnspecified RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseCreating    RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase = "ENVIRONMENT_PHASE_CREATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseStarting    RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase = "ENVIRONMENT_PHASE_STARTING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseRunning     RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase = "ENVIRONMENT_PHASE_RUNNING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseUpdating    RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase = "ENVIRONMENT_PHASE_UPDATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseStopping    RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase = "ENVIRONMENT_PHASE_STOPPING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseStopped     RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase = "ENVIRONMENT_PHASE_STOPPED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseDeleting    RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase = "ENVIRONMENT_PHASE_DELETING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseDeleted     RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusPhase) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseUnspecified, RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseCreating, RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseStarting, RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseRunning, RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseUpdating, RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseStopping, RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseStopped, RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseDeleting, RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseDeleted:
		return true
	}
	return false
}

// RunnerACK is the acknowledgement from the runner that is has received the
// environment spec.
type RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAck struct {
	Message     param.Field[string]                                                                        `json:"message"`
	SpecVersion param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckSpecVersionUnion] `json:"specVersion"`
	StatusCode  param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCode]       `json:"statusCode"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckSpecVersionUnion interface {
	ImplementsRunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckSpecVersionUnion()
}

type RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCode string

const (
	RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCodeStatusCodeUnspecified        RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCode = "STATUS_CODE_UNSPECIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCodeStatusCodeOk                 RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCode = "STATUS_CODE_OK"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCodeStatusCodeInvalidResource    RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCode = "STATUS_CODE_INVALID_RESOURCE"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCodeStatusCodeFailedPrecondition RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCode = "STATUS_CODE_FAILED_PRECONDITION"
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCode) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCodeStatusCodeUnspecified, RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCodeStatusCodeOk, RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCodeStatusCodeInvalidResource, RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCodeStatusCodeFailedPrecondition:
		return true
	}
	return false
}

type RunnerInteractionEnvironmentUpdateStatusParamsStatusSecret struct {
	// failure_message contains the reason the secret failed to be materialize.
	FailureMessage param.Field[string]                                                           `json:"failureMessage"`
	Phase          param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhase] `json:"phase"`
	SecretName     param.Field[string]                                                           `json:"secretName"`
	// warning_message contains warnings, e.g. when the secret is present but not in
	// the expected state.
	WarningMessage param.Field[string] `json:"warningMessage"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhase string

const (
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseUnspecified  RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhase = "CONTENT_PHASE_UNSPECIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseCreating     RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhase = "CONTENT_PHASE_CREATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseInitializing RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhase = "CONTENT_PHASE_INITIALIZING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseReady        RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhase = "CONTENT_PHASE_READY"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseUpdating     RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhase = "CONTENT_PHASE_UPDATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseFailed       RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhase = "CONTENT_PHASE_FAILED"
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhase) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseUnspecified, RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseCreating, RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseInitializing, RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseReady, RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseUpdating, RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseFailed:
		return true
	}
	return false
}

type RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID param.Field[string] `json:"id"`
	// phase is the current phase of the public key
	Phase param.Field[RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhase] `json:"phase"`
}

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// phase is the current phase of the public key
type RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhase string

const (
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseUnspecified  RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhase = "CONTENT_PHASE_UNSPECIFIED"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseCreating     RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhase = "CONTENT_PHASE_CREATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseInitializing RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhase = "CONTENT_PHASE_INITIALIZING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseReady        RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhase = "CONTENT_PHASE_READY"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseUpdating     RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhase = "CONTENT_PHASE_UPDATING"
	RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseFailed       RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhase = "CONTENT_PHASE_FAILED"
)

func (r RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhase) IsKnown() bool {
	switch r {
	case RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseUnspecified, RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseCreating, RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseInitializing, RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseReady, RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseUpdating, RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseFailed:
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
// Satisfied by [shared.UnionString], [shared.UnionFloat].
type RunnerInteractionEnvironmentUpdateStatusParamsStatusStatusVersionUnion interface {
	ImplementsRunnerInteractionEnvironmentUpdateStatusParamsStatusStatusVersionUnion()
}
