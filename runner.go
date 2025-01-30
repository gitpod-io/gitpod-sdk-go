// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// RunnerService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerService] method instead.
type RunnerService struct {
	Options  []option.RequestOption
	Policies *RunnerPolicyService
}

// NewRunnerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRunnerService(opts ...option.RequestOption) (r *RunnerService) {
	r = &RunnerService{}
	r.Options = opts
	r.Policies = NewRunnerPolicyService(opts...)
	return
}

// CreateRunner creates a new runner with the server. Registrations are very
//
// short-lived and must be renewed every 30 seconds. Runners can be registered for
// an entire organisation or a single user.
func (r *RunnerService) New(ctx context.Context, body RunnerNewParams, opts ...option.RequestOption) (res *RunnerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CreateRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetRunner returns a single runner.
func (r *RunnerService) Get(ctx context.Context, body RunnerGetParams, opts ...option.RequestOption) (res *RunnerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/GetRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListRunners returns all runners registered in the scope.
func (r *RunnerService) List(ctx context.Context, body RunnerListParams, opts ...option.RequestOption) (res *RunnerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/ListRunners"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// CheckAuthenticationForHost asks a runner if the user is authenticated against a
// particular host, e.g. an SCM system.
//
// If not, this function will return a URL that the user should visit to
// authenticate, or indicate that Personal Access Tokens are supported.
func (r *RunnerService) CheckAuthenticationForHost(ctx context.Context, body RunnerCheckAuthenticationForHostParams, opts ...option.RequestOption) (res *RunnerCheckAuthenticationForHostResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CheckAuthenticationForHost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// CreateRunnerToken returns a token that can be used to authenticate as the
//
// runner. Use this call to renew an outdated token - this does not expire any
// previouly issued tokens.
func (r *RunnerService) NewRunnerToken(ctx context.Context, body RunnerNewRunnerTokenParams, opts ...option.RequestOption) (res *RunnerNewRunnerTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CreateRunnerToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// DeleteRunner deletes an environment runner.
func (r *RunnerService) DeleteRunner(ctx context.Context, body RunnerDeleteRunnerParams, opts ...option.RequestOption) (res *RunnerDeleteRunnerResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/DeleteRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetRunner returns a single runner.
func (r *RunnerService) GetRunner(ctx context.Context, body RunnerGetRunnerParams, opts ...option.RequestOption) (res *RunnerGetRunnerResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/GetRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ParseContextURL asks a runner to parse a context URL, and return the parsed
// result.
//
// # This call returns
//
//   - FAILED_PRECONDITION if the user requires authentication on the runner to
//     access the context URL
//   - PERMISSION_DENIED if the user is not allowed to access the context URL using
//     the credentials they have
//   - INVALID_ARGUMENT if the context URL is invalid
//   - NOT_FOUND if the repository or branch indicated by the context URL does not
//     exist
func (r *RunnerService) ParseContextURL(ctx context.Context, body RunnerParseContextURLParams, opts ...option.RequestOption) (res *RunnerParseContextURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/ParseContextURL"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateRunner updates an environment runner.
func (r *RunnerService) UpdateRunner(ctx context.Context, body RunnerUpdateRunnerParams, opts ...option.RequestOption) (res *RunnerUpdateRunnerResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/UpdateRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RunnerNewResponse struct {
	AccessToken string                  `json:"accessToken"`
	Runner      RunnerNewResponseRunner `json:"runner"`
	JSON        runnerNewResponseJSON   `json:"-"`
}

// runnerNewResponseJSON contains the JSON metadata for the struct
// [RunnerNewResponse]
type runnerNewResponseJSON struct {
	AccessToken apijson.Field
	Runner      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerNewResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerNewResponseRunner struct {
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
	// creator is the identity of the creator of the environment
	Creator RunnerNewResponseRunnerCreator `json:"creator"`
	// RunnerKind represents the kind of a runner
	Kind RunnerNewResponseRunnerKind `json:"kind"`
	// The runner's name which is shown to users
	Name     string `json:"name"`
	RunnerID string `json:"runnerId"`
	// The runner's specification
	Spec RunnerNewResponseRunnerSpec `json:"spec"`
	// RunnerStatus represents the status of a runner
	Status RunnerNewResponseRunnerStatus `json:"status"`
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
	UpdatedAt time.Time                   `json:"updatedAt" format:"date-time"`
	JSON      runnerNewResponseRunnerJSON `json:"-"`
}

// runnerNewResponseRunnerJSON contains the JSON metadata for the struct
// [RunnerNewResponseRunner]
type runnerNewResponseRunnerJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	RunnerID    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerNewResponseRunner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerNewResponseRunnerJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the creator of the environment
type RunnerNewResponseRunnerCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal RunnerNewResponseRunnerCreatorPrincipal `json:"principal"`
	JSON      runnerNewResponseRunnerCreatorJSON      `json:"-"`
}

// runnerNewResponseRunnerCreatorJSON contains the JSON metadata for the struct
// [RunnerNewResponseRunnerCreator]
type runnerNewResponseRunnerCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerNewResponseRunnerCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerNewResponseRunnerCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type RunnerNewResponseRunnerCreatorPrincipal string

const (
	RunnerNewResponseRunnerCreatorPrincipalPrincipalUnspecified    RunnerNewResponseRunnerCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	RunnerNewResponseRunnerCreatorPrincipalPrincipalAccount        RunnerNewResponseRunnerCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	RunnerNewResponseRunnerCreatorPrincipalPrincipalUser           RunnerNewResponseRunnerCreatorPrincipal = "PRINCIPAL_USER"
	RunnerNewResponseRunnerCreatorPrincipalPrincipalRunner         RunnerNewResponseRunnerCreatorPrincipal = "PRINCIPAL_RUNNER"
	RunnerNewResponseRunnerCreatorPrincipalPrincipalEnvironment    RunnerNewResponseRunnerCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	RunnerNewResponseRunnerCreatorPrincipalPrincipalServiceAccount RunnerNewResponseRunnerCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r RunnerNewResponseRunnerCreatorPrincipal) IsKnown() bool {
	switch r {
	case RunnerNewResponseRunnerCreatorPrincipalPrincipalUnspecified, RunnerNewResponseRunnerCreatorPrincipalPrincipalAccount, RunnerNewResponseRunnerCreatorPrincipalPrincipalUser, RunnerNewResponseRunnerCreatorPrincipalPrincipalRunner, RunnerNewResponseRunnerCreatorPrincipalPrincipalEnvironment, RunnerNewResponseRunnerCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// RunnerKind represents the kind of a runner
type RunnerNewResponseRunnerKind string

const (
	RunnerNewResponseRunnerKindRunnerKindUnspecified        RunnerNewResponseRunnerKind = "RUNNER_KIND_UNSPECIFIED"
	RunnerNewResponseRunnerKindRunnerKindLocal              RunnerNewResponseRunnerKind = "RUNNER_KIND_LOCAL"
	RunnerNewResponseRunnerKindRunnerKindRemote             RunnerNewResponseRunnerKind = "RUNNER_KIND_REMOTE"
	RunnerNewResponseRunnerKindRunnerKindLocalConfiguration RunnerNewResponseRunnerKind = "RUNNER_KIND_LOCAL_CONFIGURATION"
)

func (r RunnerNewResponseRunnerKind) IsKnown() bool {
	switch r {
	case RunnerNewResponseRunnerKindRunnerKindUnspecified, RunnerNewResponseRunnerKindRunnerKindLocal, RunnerNewResponseRunnerKindRunnerKindRemote, RunnerNewResponseRunnerKindRunnerKindLocalConfiguration:
		return true
	}
	return false
}

// The runner's specification
type RunnerNewResponseRunnerSpec struct {
	// The runner's configuration
	Configuration RunnerNewResponseRunnerSpecConfiguration `json:"configuration"`
	// RunnerPhase represents the phase a runner is in
	DesiredPhase RunnerNewResponseRunnerSpecDesiredPhase `json:"desiredPhase"`
	JSON         runnerNewResponseRunnerSpecJSON         `json:"-"`
}

// runnerNewResponseRunnerSpecJSON contains the JSON metadata for the struct
// [RunnerNewResponseRunnerSpec]
type runnerNewResponseRunnerSpecJSON struct {
	Configuration apijson.Field
	DesiredPhase  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerNewResponseRunnerSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerNewResponseRunnerSpecJSON) RawJSON() string {
	return r.raw
}

// The runner's configuration
type RunnerNewResponseRunnerSpecConfiguration struct {
	// auto_update indicates whether the runner should automatically update itself.
	AutoUpdate bool `json:"autoUpdate"`
	// Region to deploy the runner in, if applicable. This is mainly used for remote
	// runners, and is only a hint. The runner may be deployed in a different region.
	// See the runner's status for the actual region.
	Region string `json:"region"`
	// The release channel the runner is on
	ReleaseChannel RunnerNewResponseRunnerSpecConfigurationReleaseChannel `json:"releaseChannel"`
	JSON           runnerNewResponseRunnerSpecConfigurationJSON           `json:"-"`
}

// runnerNewResponseRunnerSpecConfigurationJSON contains the JSON metadata for the
// struct [RunnerNewResponseRunnerSpecConfiguration]
type runnerNewResponseRunnerSpecConfigurationJSON struct {
	AutoUpdate     apijson.Field
	Region         apijson.Field
	ReleaseChannel apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RunnerNewResponseRunnerSpecConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerNewResponseRunnerSpecConfigurationJSON) RawJSON() string {
	return r.raw
}

// The release channel the runner is on
type RunnerNewResponseRunnerSpecConfigurationReleaseChannel string

const (
	RunnerNewResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelUnspecified RunnerNewResponseRunnerSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_UNSPECIFIED"
	RunnerNewResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelStable      RunnerNewResponseRunnerSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_STABLE"
	RunnerNewResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelLatest      RunnerNewResponseRunnerSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_LATEST"
)

func (r RunnerNewResponseRunnerSpecConfigurationReleaseChannel) IsKnown() bool {
	switch r {
	case RunnerNewResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelUnspecified, RunnerNewResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelStable, RunnerNewResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelLatest:
		return true
	}
	return false
}

// RunnerPhase represents the phase a runner is in
type RunnerNewResponseRunnerSpecDesiredPhase string

const (
	RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseUnspecified RunnerNewResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_UNSPECIFIED"
	RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseCreated     RunnerNewResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_CREATED"
	RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseInactive    RunnerNewResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_INACTIVE"
	RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseActive      RunnerNewResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_ACTIVE"
	RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseDeleting    RunnerNewResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_DELETING"
	RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseDeleted     RunnerNewResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_DELETED"
	RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseDegraded    RunnerNewResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_DEGRADED"
)

func (r RunnerNewResponseRunnerSpecDesiredPhase) IsKnown() bool {
	switch r {
	case RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseUnspecified, RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseCreated, RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseInactive, RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseActive, RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseDeleting, RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseDeleted, RunnerNewResponseRunnerSpecDesiredPhaseRunnerPhaseDegraded:
		return true
	}
	return false
}

// RunnerStatus represents the status of a runner
type RunnerNewResponseRunnerStatus struct {
	// additional_info contains additional information about the runner, e.g. a
	// CloudFormation stack URL.
	AdditionalInfo []RunnerNewResponseRunnerStatusAdditionalInfo `json:"additionalInfo"`
	// capabilities is a list of capabilities the runner supports.
	Capabilities []RunnerNewResponseRunnerStatusCapability `json:"capabilities"`
	LogURL       string                                    `json:"logUrl"`
	// The runner's reported message which is shown to users. This message adds more
	// context to the runner's phase.
	Message string `json:"message"`
	// RunnerPhase represents the phase a runner is in
	Phase RunnerNewResponseRunnerStatusPhase `json:"phase"`
	// region is the region the runner is running in, if applicable.
	Region        string `json:"region"`
	SystemDetails string `json:"systemDetails"`
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
	UpdatedAt time.Time                         `json:"updatedAt" format:"date-time"`
	Version   string                            `json:"version"`
	JSON      runnerNewResponseRunnerStatusJSON `json:"-"`
}

// runnerNewResponseRunnerStatusJSON contains the JSON metadata for the struct
// [RunnerNewResponseRunnerStatus]
type runnerNewResponseRunnerStatusJSON struct {
	AdditionalInfo apijson.Field
	Capabilities   apijson.Field
	LogURL         apijson.Field
	Message        apijson.Field
	Phase          apijson.Field
	Region         apijson.Field
	SystemDetails  apijson.Field
	UpdatedAt      apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RunnerNewResponseRunnerStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerNewResponseRunnerStatusJSON) RawJSON() string {
	return r.raw
}

type RunnerNewResponseRunnerStatusAdditionalInfo struct {
	Key   string                                          `json:"key"`
	Value string                                          `json:"value"`
	JSON  runnerNewResponseRunnerStatusAdditionalInfoJSON `json:"-"`
}

// runnerNewResponseRunnerStatusAdditionalInfoJSON contains the JSON metadata for
// the struct [RunnerNewResponseRunnerStatusAdditionalInfo]
type runnerNewResponseRunnerStatusAdditionalInfoJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerNewResponseRunnerStatusAdditionalInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerNewResponseRunnerStatusAdditionalInfoJSON) RawJSON() string {
	return r.raw
}

type RunnerNewResponseRunnerStatusCapability string

const (
	RunnerNewResponseRunnerStatusCapabilityRunnerCapabilityUnspecified               RunnerNewResponseRunnerStatusCapability = "RUNNER_CAPABILITY_UNSPECIFIED"
	RunnerNewResponseRunnerStatusCapabilityRunnerCapabilityFetchLocalScmIntegrations RunnerNewResponseRunnerStatusCapability = "RUNNER_CAPABILITY_FETCH_LOCAL_SCM_INTEGRATIONS"
)

func (r RunnerNewResponseRunnerStatusCapability) IsKnown() bool {
	switch r {
	case RunnerNewResponseRunnerStatusCapabilityRunnerCapabilityUnspecified, RunnerNewResponseRunnerStatusCapabilityRunnerCapabilityFetchLocalScmIntegrations:
		return true
	}
	return false
}

// RunnerPhase represents the phase a runner is in
type RunnerNewResponseRunnerStatusPhase string

const (
	RunnerNewResponseRunnerStatusPhaseRunnerPhaseUnspecified RunnerNewResponseRunnerStatusPhase = "RUNNER_PHASE_UNSPECIFIED"
	RunnerNewResponseRunnerStatusPhaseRunnerPhaseCreated     RunnerNewResponseRunnerStatusPhase = "RUNNER_PHASE_CREATED"
	RunnerNewResponseRunnerStatusPhaseRunnerPhaseInactive    RunnerNewResponseRunnerStatusPhase = "RUNNER_PHASE_INACTIVE"
	RunnerNewResponseRunnerStatusPhaseRunnerPhaseActive      RunnerNewResponseRunnerStatusPhase = "RUNNER_PHASE_ACTIVE"
	RunnerNewResponseRunnerStatusPhaseRunnerPhaseDeleting    RunnerNewResponseRunnerStatusPhase = "RUNNER_PHASE_DELETING"
	RunnerNewResponseRunnerStatusPhaseRunnerPhaseDeleted     RunnerNewResponseRunnerStatusPhase = "RUNNER_PHASE_DELETED"
	RunnerNewResponseRunnerStatusPhaseRunnerPhaseDegraded    RunnerNewResponseRunnerStatusPhase = "RUNNER_PHASE_DEGRADED"
)

func (r RunnerNewResponseRunnerStatusPhase) IsKnown() bool {
	switch r {
	case RunnerNewResponseRunnerStatusPhaseRunnerPhaseUnspecified, RunnerNewResponseRunnerStatusPhaseRunnerPhaseCreated, RunnerNewResponseRunnerStatusPhaseRunnerPhaseInactive, RunnerNewResponseRunnerStatusPhaseRunnerPhaseActive, RunnerNewResponseRunnerStatusPhaseRunnerPhaseDeleting, RunnerNewResponseRunnerStatusPhaseRunnerPhaseDeleted, RunnerNewResponseRunnerStatusPhaseRunnerPhaseDegraded:
		return true
	}
	return false
}

type RunnerGetResponse struct {
	Runner RunnerGetResponseRunner `json:"runner"`
	JSON   runnerGetResponseJSON   `json:"-"`
}

// runnerGetResponseJSON contains the JSON metadata for the struct
// [RunnerGetResponse]
type runnerGetResponseJSON struct {
	Runner      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerGetResponseRunner struct {
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
	// creator is the identity of the creator of the environment
	Creator RunnerGetResponseRunnerCreator `json:"creator"`
	// RunnerKind represents the kind of a runner
	Kind RunnerGetResponseRunnerKind `json:"kind"`
	// The runner's name which is shown to users
	Name     string `json:"name"`
	RunnerID string `json:"runnerId"`
	// The runner's specification
	Spec RunnerGetResponseRunnerSpec `json:"spec"`
	// RunnerStatus represents the status of a runner
	Status RunnerGetResponseRunnerStatus `json:"status"`
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
	UpdatedAt time.Time                   `json:"updatedAt" format:"date-time"`
	JSON      runnerGetResponseRunnerJSON `json:"-"`
}

// runnerGetResponseRunnerJSON contains the JSON metadata for the struct
// [RunnerGetResponseRunner]
type runnerGetResponseRunnerJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	RunnerID    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerGetResponseRunner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetResponseRunnerJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the creator of the environment
type RunnerGetResponseRunnerCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal RunnerGetResponseRunnerCreatorPrincipal `json:"principal"`
	JSON      runnerGetResponseRunnerCreatorJSON      `json:"-"`
}

// runnerGetResponseRunnerCreatorJSON contains the JSON metadata for the struct
// [RunnerGetResponseRunnerCreator]
type runnerGetResponseRunnerCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerGetResponseRunnerCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetResponseRunnerCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type RunnerGetResponseRunnerCreatorPrincipal string

const (
	RunnerGetResponseRunnerCreatorPrincipalPrincipalUnspecified    RunnerGetResponseRunnerCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	RunnerGetResponseRunnerCreatorPrincipalPrincipalAccount        RunnerGetResponseRunnerCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	RunnerGetResponseRunnerCreatorPrincipalPrincipalUser           RunnerGetResponseRunnerCreatorPrincipal = "PRINCIPAL_USER"
	RunnerGetResponseRunnerCreatorPrincipalPrincipalRunner         RunnerGetResponseRunnerCreatorPrincipal = "PRINCIPAL_RUNNER"
	RunnerGetResponseRunnerCreatorPrincipalPrincipalEnvironment    RunnerGetResponseRunnerCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	RunnerGetResponseRunnerCreatorPrincipalPrincipalServiceAccount RunnerGetResponseRunnerCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r RunnerGetResponseRunnerCreatorPrincipal) IsKnown() bool {
	switch r {
	case RunnerGetResponseRunnerCreatorPrincipalPrincipalUnspecified, RunnerGetResponseRunnerCreatorPrincipalPrincipalAccount, RunnerGetResponseRunnerCreatorPrincipalPrincipalUser, RunnerGetResponseRunnerCreatorPrincipalPrincipalRunner, RunnerGetResponseRunnerCreatorPrincipalPrincipalEnvironment, RunnerGetResponseRunnerCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// RunnerKind represents the kind of a runner
type RunnerGetResponseRunnerKind string

const (
	RunnerGetResponseRunnerKindRunnerKindUnspecified        RunnerGetResponseRunnerKind = "RUNNER_KIND_UNSPECIFIED"
	RunnerGetResponseRunnerKindRunnerKindLocal              RunnerGetResponseRunnerKind = "RUNNER_KIND_LOCAL"
	RunnerGetResponseRunnerKindRunnerKindRemote             RunnerGetResponseRunnerKind = "RUNNER_KIND_REMOTE"
	RunnerGetResponseRunnerKindRunnerKindLocalConfiguration RunnerGetResponseRunnerKind = "RUNNER_KIND_LOCAL_CONFIGURATION"
)

func (r RunnerGetResponseRunnerKind) IsKnown() bool {
	switch r {
	case RunnerGetResponseRunnerKindRunnerKindUnspecified, RunnerGetResponseRunnerKindRunnerKindLocal, RunnerGetResponseRunnerKindRunnerKindRemote, RunnerGetResponseRunnerKindRunnerKindLocalConfiguration:
		return true
	}
	return false
}

// The runner's specification
type RunnerGetResponseRunnerSpec struct {
	// The runner's configuration
	Configuration RunnerGetResponseRunnerSpecConfiguration `json:"configuration"`
	// RunnerPhase represents the phase a runner is in
	DesiredPhase RunnerGetResponseRunnerSpecDesiredPhase `json:"desiredPhase"`
	JSON         runnerGetResponseRunnerSpecJSON         `json:"-"`
}

// runnerGetResponseRunnerSpecJSON contains the JSON metadata for the struct
// [RunnerGetResponseRunnerSpec]
type runnerGetResponseRunnerSpecJSON struct {
	Configuration apijson.Field
	DesiredPhase  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerGetResponseRunnerSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetResponseRunnerSpecJSON) RawJSON() string {
	return r.raw
}

// The runner's configuration
type RunnerGetResponseRunnerSpecConfiguration struct {
	// auto_update indicates whether the runner should automatically update itself.
	AutoUpdate bool `json:"autoUpdate"`
	// Region to deploy the runner in, if applicable. This is mainly used for remote
	// runners, and is only a hint. The runner may be deployed in a different region.
	// See the runner's status for the actual region.
	Region string `json:"region"`
	// The release channel the runner is on
	ReleaseChannel RunnerGetResponseRunnerSpecConfigurationReleaseChannel `json:"releaseChannel"`
	JSON           runnerGetResponseRunnerSpecConfigurationJSON           `json:"-"`
}

// runnerGetResponseRunnerSpecConfigurationJSON contains the JSON metadata for the
// struct [RunnerGetResponseRunnerSpecConfiguration]
type runnerGetResponseRunnerSpecConfigurationJSON struct {
	AutoUpdate     apijson.Field
	Region         apijson.Field
	ReleaseChannel apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RunnerGetResponseRunnerSpecConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetResponseRunnerSpecConfigurationJSON) RawJSON() string {
	return r.raw
}

// The release channel the runner is on
type RunnerGetResponseRunnerSpecConfigurationReleaseChannel string

const (
	RunnerGetResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelUnspecified RunnerGetResponseRunnerSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_UNSPECIFIED"
	RunnerGetResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelStable      RunnerGetResponseRunnerSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_STABLE"
	RunnerGetResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelLatest      RunnerGetResponseRunnerSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_LATEST"
)

func (r RunnerGetResponseRunnerSpecConfigurationReleaseChannel) IsKnown() bool {
	switch r {
	case RunnerGetResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelUnspecified, RunnerGetResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelStable, RunnerGetResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelLatest:
		return true
	}
	return false
}

// RunnerPhase represents the phase a runner is in
type RunnerGetResponseRunnerSpecDesiredPhase string

const (
	RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseUnspecified RunnerGetResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_UNSPECIFIED"
	RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseCreated     RunnerGetResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_CREATED"
	RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseInactive    RunnerGetResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_INACTIVE"
	RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseActive      RunnerGetResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_ACTIVE"
	RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseDeleting    RunnerGetResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_DELETING"
	RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseDeleted     RunnerGetResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_DELETED"
	RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseDegraded    RunnerGetResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_DEGRADED"
)

func (r RunnerGetResponseRunnerSpecDesiredPhase) IsKnown() bool {
	switch r {
	case RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseUnspecified, RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseCreated, RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseInactive, RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseActive, RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseDeleting, RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseDeleted, RunnerGetResponseRunnerSpecDesiredPhaseRunnerPhaseDegraded:
		return true
	}
	return false
}

// RunnerStatus represents the status of a runner
type RunnerGetResponseRunnerStatus struct {
	// additional_info contains additional information about the runner, e.g. a
	// CloudFormation stack URL.
	AdditionalInfo []RunnerGetResponseRunnerStatusAdditionalInfo `json:"additionalInfo"`
	// capabilities is a list of capabilities the runner supports.
	Capabilities []RunnerGetResponseRunnerStatusCapability `json:"capabilities"`
	LogURL       string                                    `json:"logUrl"`
	// The runner's reported message which is shown to users. This message adds more
	// context to the runner's phase.
	Message string `json:"message"`
	// RunnerPhase represents the phase a runner is in
	Phase RunnerGetResponseRunnerStatusPhase `json:"phase"`
	// region is the region the runner is running in, if applicable.
	Region        string `json:"region"`
	SystemDetails string `json:"systemDetails"`
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
	UpdatedAt time.Time                         `json:"updatedAt" format:"date-time"`
	Version   string                            `json:"version"`
	JSON      runnerGetResponseRunnerStatusJSON `json:"-"`
}

// runnerGetResponseRunnerStatusJSON contains the JSON metadata for the struct
// [RunnerGetResponseRunnerStatus]
type runnerGetResponseRunnerStatusJSON struct {
	AdditionalInfo apijson.Field
	Capabilities   apijson.Field
	LogURL         apijson.Field
	Message        apijson.Field
	Phase          apijson.Field
	Region         apijson.Field
	SystemDetails  apijson.Field
	UpdatedAt      apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RunnerGetResponseRunnerStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetResponseRunnerStatusJSON) RawJSON() string {
	return r.raw
}

type RunnerGetResponseRunnerStatusAdditionalInfo struct {
	Key   string                                          `json:"key"`
	Value string                                          `json:"value"`
	JSON  runnerGetResponseRunnerStatusAdditionalInfoJSON `json:"-"`
}

// runnerGetResponseRunnerStatusAdditionalInfoJSON contains the JSON metadata for
// the struct [RunnerGetResponseRunnerStatusAdditionalInfo]
type runnerGetResponseRunnerStatusAdditionalInfoJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerGetResponseRunnerStatusAdditionalInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetResponseRunnerStatusAdditionalInfoJSON) RawJSON() string {
	return r.raw
}

type RunnerGetResponseRunnerStatusCapability string

const (
	RunnerGetResponseRunnerStatusCapabilityRunnerCapabilityUnspecified               RunnerGetResponseRunnerStatusCapability = "RUNNER_CAPABILITY_UNSPECIFIED"
	RunnerGetResponseRunnerStatusCapabilityRunnerCapabilityFetchLocalScmIntegrations RunnerGetResponseRunnerStatusCapability = "RUNNER_CAPABILITY_FETCH_LOCAL_SCM_INTEGRATIONS"
)

func (r RunnerGetResponseRunnerStatusCapability) IsKnown() bool {
	switch r {
	case RunnerGetResponseRunnerStatusCapabilityRunnerCapabilityUnspecified, RunnerGetResponseRunnerStatusCapabilityRunnerCapabilityFetchLocalScmIntegrations:
		return true
	}
	return false
}

// RunnerPhase represents the phase a runner is in
type RunnerGetResponseRunnerStatusPhase string

const (
	RunnerGetResponseRunnerStatusPhaseRunnerPhaseUnspecified RunnerGetResponseRunnerStatusPhase = "RUNNER_PHASE_UNSPECIFIED"
	RunnerGetResponseRunnerStatusPhaseRunnerPhaseCreated     RunnerGetResponseRunnerStatusPhase = "RUNNER_PHASE_CREATED"
	RunnerGetResponseRunnerStatusPhaseRunnerPhaseInactive    RunnerGetResponseRunnerStatusPhase = "RUNNER_PHASE_INACTIVE"
	RunnerGetResponseRunnerStatusPhaseRunnerPhaseActive      RunnerGetResponseRunnerStatusPhase = "RUNNER_PHASE_ACTIVE"
	RunnerGetResponseRunnerStatusPhaseRunnerPhaseDeleting    RunnerGetResponseRunnerStatusPhase = "RUNNER_PHASE_DELETING"
	RunnerGetResponseRunnerStatusPhaseRunnerPhaseDeleted     RunnerGetResponseRunnerStatusPhase = "RUNNER_PHASE_DELETED"
	RunnerGetResponseRunnerStatusPhaseRunnerPhaseDegraded    RunnerGetResponseRunnerStatusPhase = "RUNNER_PHASE_DEGRADED"
)

func (r RunnerGetResponseRunnerStatusPhase) IsKnown() bool {
	switch r {
	case RunnerGetResponseRunnerStatusPhaseRunnerPhaseUnspecified, RunnerGetResponseRunnerStatusPhaseRunnerPhaseCreated, RunnerGetResponseRunnerStatusPhaseRunnerPhaseInactive, RunnerGetResponseRunnerStatusPhaseRunnerPhaseActive, RunnerGetResponseRunnerStatusPhaseRunnerPhaseDeleting, RunnerGetResponseRunnerStatusPhaseRunnerPhaseDeleted, RunnerGetResponseRunnerStatusPhaseRunnerPhaseDegraded:
		return true
	}
	return false
}

type RunnerListResponse struct {
	// pagination contains the pagination options for listing runners
	Pagination RunnerListResponsePagination `json:"pagination"`
	// The runners registered in the scope
	Runners []RunnerListResponseRunner `json:"runners"`
	JSON    runnerListResponseJSON     `json:"-"`
}

// runnerListResponseJSON contains the JSON metadata for the struct
// [RunnerListResponse]
type runnerListResponseJSON struct {
	Pagination  apijson.Field
	Runners     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerListResponseJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing runners
type RunnerListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                           `json:"nextToken"`
	JSON      runnerListResponsePaginationJSON `json:"-"`
}

// runnerListResponsePaginationJSON contains the JSON metadata for the struct
// [RunnerListResponsePagination]
type runnerListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerListResponseRunner struct {
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
	// creator is the identity of the creator of the environment
	Creator RunnerListResponseRunnersCreator `json:"creator"`
	// RunnerKind represents the kind of a runner
	Kind RunnerListResponseRunnersKind `json:"kind"`
	// The runner's name which is shown to users
	Name     string `json:"name"`
	RunnerID string `json:"runnerId"`
	// The runner's specification
	Spec RunnerListResponseRunnersSpec `json:"spec"`
	// RunnerStatus represents the status of a runner
	Status RunnerListResponseRunnersStatus `json:"status"`
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
	UpdatedAt time.Time                    `json:"updatedAt" format:"date-time"`
	JSON      runnerListResponseRunnerJSON `json:"-"`
}

// runnerListResponseRunnerJSON contains the JSON metadata for the struct
// [RunnerListResponseRunner]
type runnerListResponseRunnerJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	RunnerID    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerListResponseRunner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerListResponseRunnerJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the creator of the environment
type RunnerListResponseRunnersCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal RunnerListResponseRunnersCreatorPrincipal `json:"principal"`
	JSON      runnerListResponseRunnersCreatorJSON      `json:"-"`
}

// runnerListResponseRunnersCreatorJSON contains the JSON metadata for the struct
// [RunnerListResponseRunnersCreator]
type runnerListResponseRunnersCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerListResponseRunnersCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerListResponseRunnersCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type RunnerListResponseRunnersCreatorPrincipal string

const (
	RunnerListResponseRunnersCreatorPrincipalPrincipalUnspecified    RunnerListResponseRunnersCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	RunnerListResponseRunnersCreatorPrincipalPrincipalAccount        RunnerListResponseRunnersCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	RunnerListResponseRunnersCreatorPrincipalPrincipalUser           RunnerListResponseRunnersCreatorPrincipal = "PRINCIPAL_USER"
	RunnerListResponseRunnersCreatorPrincipalPrincipalRunner         RunnerListResponseRunnersCreatorPrincipal = "PRINCIPAL_RUNNER"
	RunnerListResponseRunnersCreatorPrincipalPrincipalEnvironment    RunnerListResponseRunnersCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	RunnerListResponseRunnersCreatorPrincipalPrincipalServiceAccount RunnerListResponseRunnersCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r RunnerListResponseRunnersCreatorPrincipal) IsKnown() bool {
	switch r {
	case RunnerListResponseRunnersCreatorPrincipalPrincipalUnspecified, RunnerListResponseRunnersCreatorPrincipalPrincipalAccount, RunnerListResponseRunnersCreatorPrincipalPrincipalUser, RunnerListResponseRunnersCreatorPrincipalPrincipalRunner, RunnerListResponseRunnersCreatorPrincipalPrincipalEnvironment, RunnerListResponseRunnersCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// RunnerKind represents the kind of a runner
type RunnerListResponseRunnersKind string

const (
	RunnerListResponseRunnersKindRunnerKindUnspecified        RunnerListResponseRunnersKind = "RUNNER_KIND_UNSPECIFIED"
	RunnerListResponseRunnersKindRunnerKindLocal              RunnerListResponseRunnersKind = "RUNNER_KIND_LOCAL"
	RunnerListResponseRunnersKindRunnerKindRemote             RunnerListResponseRunnersKind = "RUNNER_KIND_REMOTE"
	RunnerListResponseRunnersKindRunnerKindLocalConfiguration RunnerListResponseRunnersKind = "RUNNER_KIND_LOCAL_CONFIGURATION"
)

func (r RunnerListResponseRunnersKind) IsKnown() bool {
	switch r {
	case RunnerListResponseRunnersKindRunnerKindUnspecified, RunnerListResponseRunnersKindRunnerKindLocal, RunnerListResponseRunnersKindRunnerKindRemote, RunnerListResponseRunnersKindRunnerKindLocalConfiguration:
		return true
	}
	return false
}

// The runner's specification
type RunnerListResponseRunnersSpec struct {
	// The runner's configuration
	Configuration RunnerListResponseRunnersSpecConfiguration `json:"configuration"`
	// RunnerPhase represents the phase a runner is in
	DesiredPhase RunnerListResponseRunnersSpecDesiredPhase `json:"desiredPhase"`
	JSON         runnerListResponseRunnersSpecJSON         `json:"-"`
}

// runnerListResponseRunnersSpecJSON contains the JSON metadata for the struct
// [RunnerListResponseRunnersSpec]
type runnerListResponseRunnersSpecJSON struct {
	Configuration apijson.Field
	DesiredPhase  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerListResponseRunnersSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerListResponseRunnersSpecJSON) RawJSON() string {
	return r.raw
}

// The runner's configuration
type RunnerListResponseRunnersSpecConfiguration struct {
	// auto_update indicates whether the runner should automatically update itself.
	AutoUpdate bool `json:"autoUpdate"`
	// Region to deploy the runner in, if applicable. This is mainly used for remote
	// runners, and is only a hint. The runner may be deployed in a different region.
	// See the runner's status for the actual region.
	Region string `json:"region"`
	// The release channel the runner is on
	ReleaseChannel RunnerListResponseRunnersSpecConfigurationReleaseChannel `json:"releaseChannel"`
	JSON           runnerListResponseRunnersSpecConfigurationJSON           `json:"-"`
}

// runnerListResponseRunnersSpecConfigurationJSON contains the JSON metadata for
// the struct [RunnerListResponseRunnersSpecConfiguration]
type runnerListResponseRunnersSpecConfigurationJSON struct {
	AutoUpdate     apijson.Field
	Region         apijson.Field
	ReleaseChannel apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RunnerListResponseRunnersSpecConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerListResponseRunnersSpecConfigurationJSON) RawJSON() string {
	return r.raw
}

// The release channel the runner is on
type RunnerListResponseRunnersSpecConfigurationReleaseChannel string

const (
	RunnerListResponseRunnersSpecConfigurationReleaseChannelRunnerReleaseChannelUnspecified RunnerListResponseRunnersSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_UNSPECIFIED"
	RunnerListResponseRunnersSpecConfigurationReleaseChannelRunnerReleaseChannelStable      RunnerListResponseRunnersSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_STABLE"
	RunnerListResponseRunnersSpecConfigurationReleaseChannelRunnerReleaseChannelLatest      RunnerListResponseRunnersSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_LATEST"
)

func (r RunnerListResponseRunnersSpecConfigurationReleaseChannel) IsKnown() bool {
	switch r {
	case RunnerListResponseRunnersSpecConfigurationReleaseChannelRunnerReleaseChannelUnspecified, RunnerListResponseRunnersSpecConfigurationReleaseChannelRunnerReleaseChannelStable, RunnerListResponseRunnersSpecConfigurationReleaseChannelRunnerReleaseChannelLatest:
		return true
	}
	return false
}

// RunnerPhase represents the phase a runner is in
type RunnerListResponseRunnersSpecDesiredPhase string

const (
	RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseUnspecified RunnerListResponseRunnersSpecDesiredPhase = "RUNNER_PHASE_UNSPECIFIED"
	RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseCreated     RunnerListResponseRunnersSpecDesiredPhase = "RUNNER_PHASE_CREATED"
	RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseInactive    RunnerListResponseRunnersSpecDesiredPhase = "RUNNER_PHASE_INACTIVE"
	RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseActive      RunnerListResponseRunnersSpecDesiredPhase = "RUNNER_PHASE_ACTIVE"
	RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseDeleting    RunnerListResponseRunnersSpecDesiredPhase = "RUNNER_PHASE_DELETING"
	RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseDeleted     RunnerListResponseRunnersSpecDesiredPhase = "RUNNER_PHASE_DELETED"
	RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseDegraded    RunnerListResponseRunnersSpecDesiredPhase = "RUNNER_PHASE_DEGRADED"
)

func (r RunnerListResponseRunnersSpecDesiredPhase) IsKnown() bool {
	switch r {
	case RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseUnspecified, RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseCreated, RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseInactive, RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseActive, RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseDeleting, RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseDeleted, RunnerListResponseRunnersSpecDesiredPhaseRunnerPhaseDegraded:
		return true
	}
	return false
}

// RunnerStatus represents the status of a runner
type RunnerListResponseRunnersStatus struct {
	// additional_info contains additional information about the runner, e.g. a
	// CloudFormation stack URL.
	AdditionalInfo []RunnerListResponseRunnersStatusAdditionalInfo `json:"additionalInfo"`
	// capabilities is a list of capabilities the runner supports.
	Capabilities []RunnerListResponseRunnersStatusCapability `json:"capabilities"`
	LogURL       string                                      `json:"logUrl"`
	// The runner's reported message which is shown to users. This message adds more
	// context to the runner's phase.
	Message string `json:"message"`
	// RunnerPhase represents the phase a runner is in
	Phase RunnerListResponseRunnersStatusPhase `json:"phase"`
	// region is the region the runner is running in, if applicable.
	Region        string `json:"region"`
	SystemDetails string `json:"systemDetails"`
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
	UpdatedAt time.Time                           `json:"updatedAt" format:"date-time"`
	Version   string                              `json:"version"`
	JSON      runnerListResponseRunnersStatusJSON `json:"-"`
}

// runnerListResponseRunnersStatusJSON contains the JSON metadata for the struct
// [RunnerListResponseRunnersStatus]
type runnerListResponseRunnersStatusJSON struct {
	AdditionalInfo apijson.Field
	Capabilities   apijson.Field
	LogURL         apijson.Field
	Message        apijson.Field
	Phase          apijson.Field
	Region         apijson.Field
	SystemDetails  apijson.Field
	UpdatedAt      apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RunnerListResponseRunnersStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerListResponseRunnersStatusJSON) RawJSON() string {
	return r.raw
}

type RunnerListResponseRunnersStatusAdditionalInfo struct {
	Key   string                                            `json:"key"`
	Value string                                            `json:"value"`
	JSON  runnerListResponseRunnersStatusAdditionalInfoJSON `json:"-"`
}

// runnerListResponseRunnersStatusAdditionalInfoJSON contains the JSON metadata for
// the struct [RunnerListResponseRunnersStatusAdditionalInfo]
type runnerListResponseRunnersStatusAdditionalInfoJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerListResponseRunnersStatusAdditionalInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerListResponseRunnersStatusAdditionalInfoJSON) RawJSON() string {
	return r.raw
}

type RunnerListResponseRunnersStatusCapability string

const (
	RunnerListResponseRunnersStatusCapabilityRunnerCapabilityUnspecified               RunnerListResponseRunnersStatusCapability = "RUNNER_CAPABILITY_UNSPECIFIED"
	RunnerListResponseRunnersStatusCapabilityRunnerCapabilityFetchLocalScmIntegrations RunnerListResponseRunnersStatusCapability = "RUNNER_CAPABILITY_FETCH_LOCAL_SCM_INTEGRATIONS"
)

func (r RunnerListResponseRunnersStatusCapability) IsKnown() bool {
	switch r {
	case RunnerListResponseRunnersStatusCapabilityRunnerCapabilityUnspecified, RunnerListResponseRunnersStatusCapabilityRunnerCapabilityFetchLocalScmIntegrations:
		return true
	}
	return false
}

// RunnerPhase represents the phase a runner is in
type RunnerListResponseRunnersStatusPhase string

const (
	RunnerListResponseRunnersStatusPhaseRunnerPhaseUnspecified RunnerListResponseRunnersStatusPhase = "RUNNER_PHASE_UNSPECIFIED"
	RunnerListResponseRunnersStatusPhaseRunnerPhaseCreated     RunnerListResponseRunnersStatusPhase = "RUNNER_PHASE_CREATED"
	RunnerListResponseRunnersStatusPhaseRunnerPhaseInactive    RunnerListResponseRunnersStatusPhase = "RUNNER_PHASE_INACTIVE"
	RunnerListResponseRunnersStatusPhaseRunnerPhaseActive      RunnerListResponseRunnersStatusPhase = "RUNNER_PHASE_ACTIVE"
	RunnerListResponseRunnersStatusPhaseRunnerPhaseDeleting    RunnerListResponseRunnersStatusPhase = "RUNNER_PHASE_DELETING"
	RunnerListResponseRunnersStatusPhaseRunnerPhaseDeleted     RunnerListResponseRunnersStatusPhase = "RUNNER_PHASE_DELETED"
	RunnerListResponseRunnersStatusPhaseRunnerPhaseDegraded    RunnerListResponseRunnersStatusPhase = "RUNNER_PHASE_DEGRADED"
)

func (r RunnerListResponseRunnersStatusPhase) IsKnown() bool {
	switch r {
	case RunnerListResponseRunnersStatusPhaseRunnerPhaseUnspecified, RunnerListResponseRunnersStatusPhaseRunnerPhaseCreated, RunnerListResponseRunnersStatusPhaseRunnerPhaseInactive, RunnerListResponseRunnersStatusPhaseRunnerPhaseActive, RunnerListResponseRunnersStatusPhaseRunnerPhaseDeleting, RunnerListResponseRunnersStatusPhaseRunnerPhaseDeleted, RunnerListResponseRunnersStatusPhaseRunnerPhaseDegraded:
		return true
	}
	return false
}

type RunnerCheckAuthenticationForHostResponse struct {
	Authenticated     bool                                         `json:"authenticated"`
	AuthenticationURL string                                       `json:"authenticationUrl"`
	PatSupported      bool                                         `json:"patSupported"`
	ScmID             string                                       `json:"scmId"`
	JSON              runnerCheckAuthenticationForHostResponseJSON `json:"-"`
}

// runnerCheckAuthenticationForHostResponseJSON contains the JSON metadata for the
// struct [RunnerCheckAuthenticationForHostResponse]
type runnerCheckAuthenticationForHostResponseJSON struct {
	Authenticated     apijson.Field
	AuthenticationURL apijson.Field
	PatSupported      apijson.Field
	ScmID             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RunnerCheckAuthenticationForHostResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerCheckAuthenticationForHostResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerNewRunnerTokenResponse struct {
	AccessToken string                           `json:"accessToken"`
	JSON        runnerNewRunnerTokenResponseJSON `json:"-"`
}

// runnerNewRunnerTokenResponseJSON contains the JSON metadata for the struct
// [RunnerNewRunnerTokenResponse]
type runnerNewRunnerTokenResponseJSON struct {
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerNewRunnerTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerNewRunnerTokenResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerDeleteRunnerResponse = interface{}

type RunnerGetRunnerResponse struct {
	Runner RunnerGetRunnerResponseRunner `json:"runner"`
	JSON   runnerGetRunnerResponseJSON   `json:"-"`
}

// runnerGetRunnerResponseJSON contains the JSON metadata for the struct
// [RunnerGetRunnerResponse]
type runnerGetRunnerResponseJSON struct {
	Runner      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerGetRunnerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetRunnerResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerGetRunnerResponseRunner struct {
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
	// creator is the identity of the creator of the environment
	Creator RunnerGetRunnerResponseRunnerCreator `json:"creator"`
	// RunnerKind represents the kind of a runner
	Kind RunnerGetRunnerResponseRunnerKind `json:"kind"`
	// The runner's name which is shown to users
	Name     string `json:"name"`
	RunnerID string `json:"runnerId"`
	// The runner's specification
	Spec RunnerGetRunnerResponseRunnerSpec `json:"spec"`
	// RunnerStatus represents the status of a runner
	Status RunnerGetRunnerResponseRunnerStatus `json:"status"`
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
	UpdatedAt time.Time                         `json:"updatedAt" format:"date-time"`
	JSON      runnerGetRunnerResponseRunnerJSON `json:"-"`
}

// runnerGetRunnerResponseRunnerJSON contains the JSON metadata for the struct
// [RunnerGetRunnerResponseRunner]
type runnerGetRunnerResponseRunnerJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	RunnerID    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerGetRunnerResponseRunner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetRunnerResponseRunnerJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the creator of the environment
type RunnerGetRunnerResponseRunnerCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal RunnerGetRunnerResponseRunnerCreatorPrincipal `json:"principal"`
	JSON      runnerGetRunnerResponseRunnerCreatorJSON      `json:"-"`
}

// runnerGetRunnerResponseRunnerCreatorJSON contains the JSON metadata for the
// struct [RunnerGetRunnerResponseRunnerCreator]
type runnerGetRunnerResponseRunnerCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerGetRunnerResponseRunnerCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetRunnerResponseRunnerCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type RunnerGetRunnerResponseRunnerCreatorPrincipal string

const (
	RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalUnspecified    RunnerGetRunnerResponseRunnerCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalAccount        RunnerGetRunnerResponseRunnerCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalUser           RunnerGetRunnerResponseRunnerCreatorPrincipal = "PRINCIPAL_USER"
	RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalRunner         RunnerGetRunnerResponseRunnerCreatorPrincipal = "PRINCIPAL_RUNNER"
	RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalEnvironment    RunnerGetRunnerResponseRunnerCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalServiceAccount RunnerGetRunnerResponseRunnerCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r RunnerGetRunnerResponseRunnerCreatorPrincipal) IsKnown() bool {
	switch r {
	case RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalUnspecified, RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalAccount, RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalUser, RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalRunner, RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalEnvironment, RunnerGetRunnerResponseRunnerCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

// RunnerKind represents the kind of a runner
type RunnerGetRunnerResponseRunnerKind string

const (
	RunnerGetRunnerResponseRunnerKindRunnerKindUnspecified        RunnerGetRunnerResponseRunnerKind = "RUNNER_KIND_UNSPECIFIED"
	RunnerGetRunnerResponseRunnerKindRunnerKindLocal              RunnerGetRunnerResponseRunnerKind = "RUNNER_KIND_LOCAL"
	RunnerGetRunnerResponseRunnerKindRunnerKindRemote             RunnerGetRunnerResponseRunnerKind = "RUNNER_KIND_REMOTE"
	RunnerGetRunnerResponseRunnerKindRunnerKindLocalConfiguration RunnerGetRunnerResponseRunnerKind = "RUNNER_KIND_LOCAL_CONFIGURATION"
)

func (r RunnerGetRunnerResponseRunnerKind) IsKnown() bool {
	switch r {
	case RunnerGetRunnerResponseRunnerKindRunnerKindUnspecified, RunnerGetRunnerResponseRunnerKindRunnerKindLocal, RunnerGetRunnerResponseRunnerKindRunnerKindRemote, RunnerGetRunnerResponseRunnerKindRunnerKindLocalConfiguration:
		return true
	}
	return false
}

// The runner's specification
type RunnerGetRunnerResponseRunnerSpec struct {
	// The runner's configuration
	Configuration RunnerGetRunnerResponseRunnerSpecConfiguration `json:"configuration"`
	// RunnerPhase represents the phase a runner is in
	DesiredPhase RunnerGetRunnerResponseRunnerSpecDesiredPhase `json:"desiredPhase"`
	JSON         runnerGetRunnerResponseRunnerSpecJSON         `json:"-"`
}

// runnerGetRunnerResponseRunnerSpecJSON contains the JSON metadata for the struct
// [RunnerGetRunnerResponseRunnerSpec]
type runnerGetRunnerResponseRunnerSpecJSON struct {
	Configuration apijson.Field
	DesiredPhase  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerGetRunnerResponseRunnerSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetRunnerResponseRunnerSpecJSON) RawJSON() string {
	return r.raw
}

// The runner's configuration
type RunnerGetRunnerResponseRunnerSpecConfiguration struct {
	// auto_update indicates whether the runner should automatically update itself.
	AutoUpdate bool `json:"autoUpdate"`
	// Region to deploy the runner in, if applicable. This is mainly used for remote
	// runners, and is only a hint. The runner may be deployed in a different region.
	// See the runner's status for the actual region.
	Region string `json:"region"`
	// The release channel the runner is on
	ReleaseChannel RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannel `json:"releaseChannel"`
	JSON           runnerGetRunnerResponseRunnerSpecConfigurationJSON           `json:"-"`
}

// runnerGetRunnerResponseRunnerSpecConfigurationJSON contains the JSON metadata
// for the struct [RunnerGetRunnerResponseRunnerSpecConfiguration]
type runnerGetRunnerResponseRunnerSpecConfigurationJSON struct {
	AutoUpdate     apijson.Field
	Region         apijson.Field
	ReleaseChannel apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RunnerGetRunnerResponseRunnerSpecConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetRunnerResponseRunnerSpecConfigurationJSON) RawJSON() string {
	return r.raw
}

// The release channel the runner is on
type RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannel string

const (
	RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelUnspecified RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_UNSPECIFIED"
	RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelStable      RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_STABLE"
	RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelLatest      RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_LATEST"
)

func (r RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannel) IsKnown() bool {
	switch r {
	case RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelUnspecified, RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelStable, RunnerGetRunnerResponseRunnerSpecConfigurationReleaseChannelRunnerReleaseChannelLatest:
		return true
	}
	return false
}

// RunnerPhase represents the phase a runner is in
type RunnerGetRunnerResponseRunnerSpecDesiredPhase string

const (
	RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseUnspecified RunnerGetRunnerResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_UNSPECIFIED"
	RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseCreated     RunnerGetRunnerResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_CREATED"
	RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseInactive    RunnerGetRunnerResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_INACTIVE"
	RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseActive      RunnerGetRunnerResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_ACTIVE"
	RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseDeleting    RunnerGetRunnerResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_DELETING"
	RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseDeleted     RunnerGetRunnerResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_DELETED"
	RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseDegraded    RunnerGetRunnerResponseRunnerSpecDesiredPhase = "RUNNER_PHASE_DEGRADED"
)

func (r RunnerGetRunnerResponseRunnerSpecDesiredPhase) IsKnown() bool {
	switch r {
	case RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseUnspecified, RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseCreated, RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseInactive, RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseActive, RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseDeleting, RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseDeleted, RunnerGetRunnerResponseRunnerSpecDesiredPhaseRunnerPhaseDegraded:
		return true
	}
	return false
}

// RunnerStatus represents the status of a runner
type RunnerGetRunnerResponseRunnerStatus struct {
	// additional_info contains additional information about the runner, e.g. a
	// CloudFormation stack URL.
	AdditionalInfo []RunnerGetRunnerResponseRunnerStatusAdditionalInfo `json:"additionalInfo"`
	// capabilities is a list of capabilities the runner supports.
	Capabilities []RunnerGetRunnerResponseRunnerStatusCapability `json:"capabilities"`
	LogURL       string                                          `json:"logUrl"`
	// The runner's reported message which is shown to users. This message adds more
	// context to the runner's phase.
	Message string `json:"message"`
	// RunnerPhase represents the phase a runner is in
	Phase RunnerGetRunnerResponseRunnerStatusPhase `json:"phase"`
	// region is the region the runner is running in, if applicable.
	Region        string `json:"region"`
	SystemDetails string `json:"systemDetails"`
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
	UpdatedAt time.Time                               `json:"updatedAt" format:"date-time"`
	Version   string                                  `json:"version"`
	JSON      runnerGetRunnerResponseRunnerStatusJSON `json:"-"`
}

// runnerGetRunnerResponseRunnerStatusJSON contains the JSON metadata for the
// struct [RunnerGetRunnerResponseRunnerStatus]
type runnerGetRunnerResponseRunnerStatusJSON struct {
	AdditionalInfo apijson.Field
	Capabilities   apijson.Field
	LogURL         apijson.Field
	Message        apijson.Field
	Phase          apijson.Field
	Region         apijson.Field
	SystemDetails  apijson.Field
	UpdatedAt      apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RunnerGetRunnerResponseRunnerStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetRunnerResponseRunnerStatusJSON) RawJSON() string {
	return r.raw
}

type RunnerGetRunnerResponseRunnerStatusAdditionalInfo struct {
	Key   string                                                `json:"key"`
	Value string                                                `json:"value"`
	JSON  runnerGetRunnerResponseRunnerStatusAdditionalInfoJSON `json:"-"`
}

// runnerGetRunnerResponseRunnerStatusAdditionalInfoJSON contains the JSON metadata
// for the struct [RunnerGetRunnerResponseRunnerStatusAdditionalInfo]
type runnerGetRunnerResponseRunnerStatusAdditionalInfoJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerGetRunnerResponseRunnerStatusAdditionalInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerGetRunnerResponseRunnerStatusAdditionalInfoJSON) RawJSON() string {
	return r.raw
}

type RunnerGetRunnerResponseRunnerStatusCapability string

const (
	RunnerGetRunnerResponseRunnerStatusCapabilityRunnerCapabilityUnspecified               RunnerGetRunnerResponseRunnerStatusCapability = "RUNNER_CAPABILITY_UNSPECIFIED"
	RunnerGetRunnerResponseRunnerStatusCapabilityRunnerCapabilityFetchLocalScmIntegrations RunnerGetRunnerResponseRunnerStatusCapability = "RUNNER_CAPABILITY_FETCH_LOCAL_SCM_INTEGRATIONS"
)

func (r RunnerGetRunnerResponseRunnerStatusCapability) IsKnown() bool {
	switch r {
	case RunnerGetRunnerResponseRunnerStatusCapabilityRunnerCapabilityUnspecified, RunnerGetRunnerResponseRunnerStatusCapabilityRunnerCapabilityFetchLocalScmIntegrations:
		return true
	}
	return false
}

// RunnerPhase represents the phase a runner is in
type RunnerGetRunnerResponseRunnerStatusPhase string

const (
	RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseUnspecified RunnerGetRunnerResponseRunnerStatusPhase = "RUNNER_PHASE_UNSPECIFIED"
	RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseCreated     RunnerGetRunnerResponseRunnerStatusPhase = "RUNNER_PHASE_CREATED"
	RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseInactive    RunnerGetRunnerResponseRunnerStatusPhase = "RUNNER_PHASE_INACTIVE"
	RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseActive      RunnerGetRunnerResponseRunnerStatusPhase = "RUNNER_PHASE_ACTIVE"
	RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseDeleting    RunnerGetRunnerResponseRunnerStatusPhase = "RUNNER_PHASE_DELETING"
	RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseDeleted     RunnerGetRunnerResponseRunnerStatusPhase = "RUNNER_PHASE_DELETED"
	RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseDegraded    RunnerGetRunnerResponseRunnerStatusPhase = "RUNNER_PHASE_DEGRADED"
)

func (r RunnerGetRunnerResponseRunnerStatusPhase) IsKnown() bool {
	switch r {
	case RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseUnspecified, RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseCreated, RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseInactive, RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseActive, RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseDeleting, RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseDeleted, RunnerGetRunnerResponseRunnerStatusPhaseRunnerPhaseDegraded:
		return true
	}
	return false
}

type RunnerParseContextURLResponse struct {
	Git                RunnerParseContextURLResponseGit  `json:"git"`
	OriginalContextURL string                            `json:"originalContextUrl"`
	JSON               runnerParseContextURLResponseJSON `json:"-"`
}

// runnerParseContextURLResponseJSON contains the JSON metadata for the struct
// [RunnerParseContextURLResponse]
type runnerParseContextURLResponseJSON struct {
	Git                apijson.Field
	OriginalContextURL apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RunnerParseContextURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerParseContextURLResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerParseContextURLResponseGit struct {
	Branch            string                               `json:"branch"`
	CloneURL          string                               `json:"cloneUrl"`
	Commit            string                               `json:"commit"`
	Host              string                               `json:"host"`
	Owner             string                               `json:"owner"`
	Repo              string                               `json:"repo"`
	UpstreamRemoteURL string                               `json:"upstreamRemoteUrl"`
	JSON              runnerParseContextURLResponseGitJSON `json:"-"`
}

// runnerParseContextURLResponseGitJSON contains the JSON metadata for the struct
// [RunnerParseContextURLResponseGit]
type runnerParseContextURLResponseGitJSON struct {
	Branch            apijson.Field
	CloneURL          apijson.Field
	Commit            apijson.Field
	Host              apijson.Field
	Owner             apijson.Field
	Repo              apijson.Field
	UpstreamRemoteURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RunnerParseContextURLResponseGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerParseContextURLResponseGitJSON) RawJSON() string {
	return r.raw
}

type RunnerUpdateRunnerResponse = interface{}

type RunnerNewParams struct {
	// RunnerKind represents the kind of a runner
	Kind param.Field[RunnerNewParamsKind] `json:"kind"`
	// The runner name for humans
	Name param.Field[string]              `json:"name"`
	Spec param.Field[RunnerNewParamsSpec] `json:"spec"`
}

func (r RunnerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// RunnerKind represents the kind of a runner
type RunnerNewParamsKind string

const (
	RunnerNewParamsKindRunnerKindUnspecified        RunnerNewParamsKind = "RUNNER_KIND_UNSPECIFIED"
	RunnerNewParamsKindRunnerKindLocal              RunnerNewParamsKind = "RUNNER_KIND_LOCAL"
	RunnerNewParamsKindRunnerKindRemote             RunnerNewParamsKind = "RUNNER_KIND_REMOTE"
	RunnerNewParamsKindRunnerKindLocalConfiguration RunnerNewParamsKind = "RUNNER_KIND_LOCAL_CONFIGURATION"
)

func (r RunnerNewParamsKind) IsKnown() bool {
	switch r {
	case RunnerNewParamsKindRunnerKindUnspecified, RunnerNewParamsKindRunnerKindLocal, RunnerNewParamsKindRunnerKindRemote, RunnerNewParamsKindRunnerKindLocalConfiguration:
		return true
	}
	return false
}

type RunnerNewParamsSpec struct {
	// The runner's configuration
	Configuration param.Field[RunnerNewParamsSpecConfiguration] `json:"configuration"`
	// RunnerPhase represents the phase a runner is in
	DesiredPhase param.Field[RunnerNewParamsSpecDesiredPhase] `json:"desiredPhase"`
}

func (r RunnerNewParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The runner's configuration
type RunnerNewParamsSpecConfiguration struct {
	// auto_update indicates whether the runner should automatically update itself.
	AutoUpdate param.Field[bool] `json:"autoUpdate"`
	// Region to deploy the runner in, if applicable. This is mainly used for remote
	// runners, and is only a hint. The runner may be deployed in a different region.
	// See the runner's status for the actual region.
	Region param.Field[string] `json:"region"`
	// The release channel the runner is on
	ReleaseChannel param.Field[RunnerNewParamsSpecConfigurationReleaseChannel] `json:"releaseChannel"`
}

func (r RunnerNewParamsSpecConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The release channel the runner is on
type RunnerNewParamsSpecConfigurationReleaseChannel string

const (
	RunnerNewParamsSpecConfigurationReleaseChannelRunnerReleaseChannelUnspecified RunnerNewParamsSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_UNSPECIFIED"
	RunnerNewParamsSpecConfigurationReleaseChannelRunnerReleaseChannelStable      RunnerNewParamsSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_STABLE"
	RunnerNewParamsSpecConfigurationReleaseChannelRunnerReleaseChannelLatest      RunnerNewParamsSpecConfigurationReleaseChannel = "RUNNER_RELEASE_CHANNEL_LATEST"
)

func (r RunnerNewParamsSpecConfigurationReleaseChannel) IsKnown() bool {
	switch r {
	case RunnerNewParamsSpecConfigurationReleaseChannelRunnerReleaseChannelUnspecified, RunnerNewParamsSpecConfigurationReleaseChannelRunnerReleaseChannelStable, RunnerNewParamsSpecConfigurationReleaseChannelRunnerReleaseChannelLatest:
		return true
	}
	return false
}

// RunnerPhase represents the phase a runner is in
type RunnerNewParamsSpecDesiredPhase string

const (
	RunnerNewParamsSpecDesiredPhaseRunnerPhaseUnspecified RunnerNewParamsSpecDesiredPhase = "RUNNER_PHASE_UNSPECIFIED"
	RunnerNewParamsSpecDesiredPhaseRunnerPhaseCreated     RunnerNewParamsSpecDesiredPhase = "RUNNER_PHASE_CREATED"
	RunnerNewParamsSpecDesiredPhaseRunnerPhaseInactive    RunnerNewParamsSpecDesiredPhase = "RUNNER_PHASE_INACTIVE"
	RunnerNewParamsSpecDesiredPhaseRunnerPhaseActive      RunnerNewParamsSpecDesiredPhase = "RUNNER_PHASE_ACTIVE"
	RunnerNewParamsSpecDesiredPhaseRunnerPhaseDeleting    RunnerNewParamsSpecDesiredPhase = "RUNNER_PHASE_DELETING"
	RunnerNewParamsSpecDesiredPhaseRunnerPhaseDeleted     RunnerNewParamsSpecDesiredPhase = "RUNNER_PHASE_DELETED"
	RunnerNewParamsSpecDesiredPhaseRunnerPhaseDegraded    RunnerNewParamsSpecDesiredPhase = "RUNNER_PHASE_DEGRADED"
)

func (r RunnerNewParamsSpecDesiredPhase) IsKnown() bool {
	switch r {
	case RunnerNewParamsSpecDesiredPhaseRunnerPhaseUnspecified, RunnerNewParamsSpecDesiredPhaseRunnerPhaseCreated, RunnerNewParamsSpecDesiredPhaseRunnerPhaseInactive, RunnerNewParamsSpecDesiredPhaseRunnerPhaseActive, RunnerNewParamsSpecDesiredPhaseRunnerPhaseDeleting, RunnerNewParamsSpecDesiredPhaseRunnerPhaseDeleted, RunnerNewParamsSpecDesiredPhaseRunnerPhaseDegraded:
		return true
	}
	return false
}

type RunnerGetParams struct {
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerListParams struct {
	Filter param.Field[RunnerListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing runners
	Pagination param.Field[RunnerListParamsPagination] `json:"pagination"`
}

func (r RunnerListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerListParamsFilter struct {
	// creator_ids filters the response to only runner created by specified users
	CreatorIDs param.Field[[]string] `json:"creatorIds" format:"uuid"`
	// kinds filters the response to only runners of the specified kinds
	Kinds param.Field[[]RunnerListParamsFilterKind] `json:"kinds"`
}

func (r RunnerListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// RunnerKind represents the kind of a runner
type RunnerListParamsFilterKind string

const (
	RunnerListParamsFilterKindRunnerKindUnspecified        RunnerListParamsFilterKind = "RUNNER_KIND_UNSPECIFIED"
	RunnerListParamsFilterKindRunnerKindLocal              RunnerListParamsFilterKind = "RUNNER_KIND_LOCAL"
	RunnerListParamsFilterKindRunnerKindRemote             RunnerListParamsFilterKind = "RUNNER_KIND_REMOTE"
	RunnerListParamsFilterKindRunnerKindLocalConfiguration RunnerListParamsFilterKind = "RUNNER_KIND_LOCAL_CONFIGURATION"
)

func (r RunnerListParamsFilterKind) IsKnown() bool {
	switch r {
	case RunnerListParamsFilterKindRunnerKindUnspecified, RunnerListParamsFilterKindRunnerKindLocal, RunnerListParamsFilterKindRunnerKindRemote, RunnerListParamsFilterKindRunnerKindLocalConfiguration:
		return true
	}
	return false
}

// pagination contains the pagination options for listing runners
type RunnerListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	//
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerCheckAuthenticationForHostParams struct {
	Host     param.Field[string] `json:"host"`
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerCheckAuthenticationForHostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerNewRunnerTokenParams struct {
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerNewRunnerTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerDeleteRunnerParams struct {
	// force indicates whether the runner should be deleted forcefully. When force
	// deleting a Runner, all Environments on the runner are also force deleted and
	// regular Runner lifecycle is not respected. Force deleting can result in data
	// loss.
	Force    param.Field[bool]   `json:"force"`
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerDeleteRunnerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerGetRunnerParams struct {
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerGetRunnerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerParseContextURLParams struct {
	ContextURL param.Field[string] `json:"contextUrl" format:"uri"`
	RunnerID   param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerParseContextURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerUpdateRunnerParams struct {
	Body RunnerUpdateRunnerParamsBodyUnion `json:"body,required"`
}

func (r RunnerUpdateRunnerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerUpdateRunnerParamsBody struct {
	// The runner's name which is shown to users
	Name param.Field[string]      `json:"name"`
	Spec param.Field[interface{}] `json:"spec"`
}

func (r RunnerUpdateRunnerParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerUpdateRunnerParamsBody) implementsRunnerUpdateRunnerParamsBodyUnion() {}

// Satisfied by [RunnerUpdateRunnerParamsBodyName],
// [RunnerUpdateRunnerParamsBodySpec], [RunnerUpdateRunnerParamsBody].
type RunnerUpdateRunnerParamsBodyUnion interface {
	implementsRunnerUpdateRunnerParamsBodyUnion()
}

type RunnerUpdateRunnerParamsBodyName struct {
	// The runner's name which is shown to users
	Name param.Field[string] `json:"name,required"`
}

func (r RunnerUpdateRunnerParamsBodyName) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerUpdateRunnerParamsBodyName) implementsRunnerUpdateRunnerParamsBodyUnion() {}

type RunnerUpdateRunnerParamsBodySpec struct {
	Spec param.Field[RunnerUpdateRunnerParamsBodySpecSpecUnion] `json:"spec,required"`
}

func (r RunnerUpdateRunnerParamsBodySpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerUpdateRunnerParamsBodySpec) implementsRunnerUpdateRunnerParamsBodyUnion() {}

type RunnerUpdateRunnerParamsBodySpecSpec struct {
	Configuration param.Field[interface{}] `json:"configuration"`
	// RunnerPhase represents the phase a runner is in
	DesiredPhase param.Field[RunnerUpdateRunnerParamsBodySpecSpecDesiredPhase] `json:"desiredPhase"`
}

func (r RunnerUpdateRunnerParamsBodySpecSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerUpdateRunnerParamsBodySpecSpec) implementsRunnerUpdateRunnerParamsBodySpecSpecUnion() {}

// Satisfied by [RunnerUpdateRunnerParamsBodySpecSpecConfiguration],
// [RunnerUpdateRunnerParamsBodySpecSpecDesiredPhase],
// [RunnerUpdateRunnerParamsBodySpecSpec].
type RunnerUpdateRunnerParamsBodySpecSpecUnion interface {
	implementsRunnerUpdateRunnerParamsBodySpecSpecUnion()
}

type RunnerUpdateRunnerParamsBodySpecSpecConfiguration struct {
	Configuration param.Field[RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationUnion] `json:"configuration,required"`
}

func (r RunnerUpdateRunnerParamsBodySpecSpecConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerUpdateRunnerParamsBodySpecSpecConfiguration) implementsRunnerUpdateRunnerParamsBodySpecSpecUnion() {
}

type RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfiguration struct {
	// auto_update indicates whether the runner should automatically update itself.
	AutoUpdate param.Field[bool] `json:"autoUpdate"`
	// The release channel the runner is on
	ReleaseChannel param.Field[RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannel] `json:"releaseChannel"`
}

func (r RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfiguration) implementsRunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationUnion() {
}

// Satisfied by
// [RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationAutoUpdate],
// [RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannel],
// [RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfiguration].
type RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationUnion interface {
	implementsRunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationUnion()
}

type RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationAutoUpdate struct {
	// auto_update indicates whether the runner should automatically update itself.
	AutoUpdate param.Field[bool] `json:"autoUpdate,required"`
}

func (r RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationAutoUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationAutoUpdate) implementsRunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationUnion() {
}

type RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannel struct {
	// The release channel the runner is on
	ReleaseChannel param.Field[RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannel] `json:"releaseChannel,required"`
}

func (r RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannel) implementsRunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationUnion() {
}

// The release channel the runner is on
type RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannel string

const (
	RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannelRunnerReleaseChannelUnspecified RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannel = "RUNNER_RELEASE_CHANNEL_UNSPECIFIED"
	RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannelRunnerReleaseChannelStable      RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannel = "RUNNER_RELEASE_CHANNEL_STABLE"
	RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannelRunnerReleaseChannelLatest      RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannel = "RUNNER_RELEASE_CHANNEL_LATEST"
)

func (r RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannel) IsKnown() bool {
	switch r {
	case RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannelRunnerReleaseChannelUnspecified, RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannelRunnerReleaseChannelStable, RunnerUpdateRunnerParamsBodySpecSpecConfigurationConfigurationReleaseChannelReleaseChannelRunnerReleaseChannelLatest:
		return true
	}
	return false
}

type RunnerUpdateRunnerParamsBodySpecSpecDesiredPhase struct {
	// RunnerPhase represents the phase a runner is in
	DesiredPhase param.Field[RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhase] `json:"desiredPhase,required"`
}

func (r RunnerUpdateRunnerParamsBodySpecSpecDesiredPhase) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerUpdateRunnerParamsBodySpecSpecDesiredPhase) implementsRunnerUpdateRunnerParamsBodySpecSpecUnion() {
}

// RunnerPhase represents the phase a runner is in
type RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhase string

const (
	RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseUnspecified RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhase = "RUNNER_PHASE_UNSPECIFIED"
	RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseCreated     RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhase = "RUNNER_PHASE_CREATED"
	RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseInactive    RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhase = "RUNNER_PHASE_INACTIVE"
	RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseActive      RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhase = "RUNNER_PHASE_ACTIVE"
	RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseDeleting    RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhase = "RUNNER_PHASE_DELETING"
	RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseDeleted     RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhase = "RUNNER_PHASE_DELETED"
	RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseDegraded    RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhase = "RUNNER_PHASE_DEGRADED"
)

func (r RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhase) IsKnown() bool {
	switch r {
	case RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseUnspecified, RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseCreated, RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseInactive, RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseActive, RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseDeleting, RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseDeleted, RunnerUpdateRunnerParamsBodySpecSpecDesiredPhaseDesiredPhaseRunnerPhaseDegraded:
		return true
	}
	return false
}
