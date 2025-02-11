// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// RunnerService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerService] method instead.
type RunnerService struct {
	Options        []option.RequestOption
	Configurations *RunnerConfigurationService
	Policies       *RunnerPolicyService
}

// NewRunnerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRunnerService(opts ...option.RequestOption) (r *RunnerService) {
	r = &RunnerService{}
	r.Options = opts
	r.Configurations = NewRunnerConfigurationService(opts...)
	r.Policies = NewRunnerPolicyService(opts...)
	return
}

// CreateRunner creates a new runner with the server. Registrations are very
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

// UpdateRunner updates an environment runner.
func (r *RunnerService) Update(ctx context.Context, body RunnerUpdateParams, opts ...option.RequestOption) (res *RunnerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/UpdateRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListRunners returns all runners registered in the scope.
func (r *RunnerService) List(ctx context.Context, params RunnerListParams, opts ...option.RequestOption) (res *pagination.RunnersPage[Runner], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.RunnerService/ListRunners"
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

// ListRunners returns all runners registered in the scope.
func (r *RunnerService) ListAutoPaging(ctx context.Context, params RunnerListParams, opts ...option.RequestOption) *pagination.RunnersPageAutoPager[Runner] {
	return pagination.NewRunnersPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteRunner deletes an environment runner.
func (r *RunnerService) Delete(ctx context.Context, body RunnerDeleteParams, opts ...option.RequestOption) (res *RunnerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/DeleteRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// CheckAuthenticationForHost asks a runner if the user is authenticated against a
// particular host, e.g. an SCM system. If not, this function will return a URL
// that the user should visit to authenticate, or indicate that Personal Access
// Tokens are supported.
func (r *RunnerService) CheckAuthenticationForHost(ctx context.Context, body RunnerCheckAuthenticationForHostParams, opts ...option.RequestOption) (res *RunnerCheckAuthenticationForHostResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CheckAuthenticationForHost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// CreateRunnerToken returns a token that can be used to authenticate as the
// runner. Use this call to renew an outdated token - this does not expire any
// previously issued tokens.
func (r *RunnerService) NewRunnerToken(ctx context.Context, body RunnerNewRunnerTokenParams, opts ...option.RequestOption) (res *RunnerNewRunnerTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CreateRunnerToken"
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

type Runner struct {
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
	Creator shared.Subject `json:"creator"`
	// RunnerKind represents the kind of a runner
	Kind RunnerKind `json:"kind"`
	// The runner's name which is shown to users
	Name string `json:"name"`
	// RunnerProvider identifies the specific implementation type of a runner. Each
	// provider maps to a specific kind of runner (local or remote), as specified below
	// for each provider.
	Provider RunnerProvider `json:"provider"`
	RunnerID string         `json:"runnerId"`
	// The runner's specification
	Spec RunnerSpec `json:"spec"`
	// RunnerStatus represents the status of a runner
	Status RunnerStatus `json:"status"`
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
	UpdatedAt time.Time  `json:"updatedAt" format:"date-time"`
	JSON      runnerJSON `json:"-"`
}

// runnerJSON contains the JSON metadata for the struct [Runner]
type runnerJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	Provider    apijson.Field
	RunnerID    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Runner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerJSON) RawJSON() string {
	return r.raw
}

type RunnerCapability string

const (
	RunnerCapabilityUnspecified               RunnerCapability = "RUNNER_CAPABILITY_UNSPECIFIED"
	RunnerCapabilityFetchLocalScmIntegrations RunnerCapability = "RUNNER_CAPABILITY_FETCH_LOCAL_SCM_INTEGRATIONS"
)

func (r RunnerCapability) IsKnown() bool {
	switch r {
	case RunnerCapabilityUnspecified, RunnerCapabilityFetchLocalScmIntegrations:
		return true
	}
	return false
}

type RunnerConfiguration struct {
	// auto_update indicates whether the runner should automatically update itself.
	AutoUpdate bool `json:"autoUpdate"`
	// Region to deploy the runner in, if applicable. This is mainly used for remote
	// runners, and is only a hint. The runner may be deployed in a different region.
	// See the runner's status for the actual region.
	Region string `json:"region"`
	// The release channel the runner is on
	ReleaseChannel RunnerReleaseChannel    `json:"releaseChannel"`
	JSON           runnerConfigurationJSON `json:"-"`
}

// runnerConfigurationJSON contains the JSON metadata for the struct
// [RunnerConfiguration]
type runnerConfigurationJSON struct {
	AutoUpdate     apijson.Field
	Region         apijson.Field
	ReleaseChannel apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RunnerConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationParam struct {
	// auto_update indicates whether the runner should automatically update itself.
	AutoUpdate param.Field[bool] `json:"autoUpdate"`
	// Region to deploy the runner in, if applicable. This is mainly used for remote
	// runners, and is only a hint. The runner may be deployed in a different region.
	// See the runner's status for the actual region.
	Region param.Field[string] `json:"region"`
	// The release channel the runner is on
	ReleaseChannel param.Field[RunnerReleaseChannel] `json:"releaseChannel"`
}

func (r RunnerConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// RunnerKind represents the kind of a runner
type RunnerKind string

const (
	RunnerKindUnspecified        RunnerKind = "RUNNER_KIND_UNSPECIFIED"
	RunnerKindLocal              RunnerKind = "RUNNER_KIND_LOCAL"
	RunnerKindRemote             RunnerKind = "RUNNER_KIND_REMOTE"
	RunnerKindLocalConfiguration RunnerKind = "RUNNER_KIND_LOCAL_CONFIGURATION"
)

func (r RunnerKind) IsKnown() bool {
	switch r {
	case RunnerKindUnspecified, RunnerKindLocal, RunnerKindRemote, RunnerKindLocalConfiguration:
		return true
	}
	return false
}

// RunnerPhase represents the phase a runner is in
type RunnerPhase string

const (
	RunnerPhaseUnspecified RunnerPhase = "RUNNER_PHASE_UNSPECIFIED"
	RunnerPhaseCreated     RunnerPhase = "RUNNER_PHASE_CREATED"
	RunnerPhaseInactive    RunnerPhase = "RUNNER_PHASE_INACTIVE"
	RunnerPhaseActive      RunnerPhase = "RUNNER_PHASE_ACTIVE"
	RunnerPhaseDeleting    RunnerPhase = "RUNNER_PHASE_DELETING"
	RunnerPhaseDeleted     RunnerPhase = "RUNNER_PHASE_DELETED"
	RunnerPhaseDegraded    RunnerPhase = "RUNNER_PHASE_DEGRADED"
)

func (r RunnerPhase) IsKnown() bool {
	switch r {
	case RunnerPhaseUnspecified, RunnerPhaseCreated, RunnerPhaseInactive, RunnerPhaseActive, RunnerPhaseDeleting, RunnerPhaseDeleted, RunnerPhaseDegraded:
		return true
	}
	return false
}

// RunnerProvider identifies the specific implementation type of a runner. Each
// provider maps to a specific kind of runner (local or remote), as specified below
// for each provider.
type RunnerProvider string

const (
	RunnerProviderUnspecified RunnerProvider = "RUNNER_PROVIDER_UNSPECIFIED"
	RunnerProviderAwsEc2      RunnerProvider = "RUNNER_PROVIDER_AWS_EC2"
	RunnerProviderLinuxHost   RunnerProvider = "RUNNER_PROVIDER_LINUX_HOST"
	RunnerProviderDesktopMac  RunnerProvider = "RUNNER_PROVIDER_DESKTOP_MAC"
)

func (r RunnerProvider) IsKnown() bool {
	switch r {
	case RunnerProviderUnspecified, RunnerProviderAwsEc2, RunnerProviderLinuxHost, RunnerProviderDesktopMac:
		return true
	}
	return false
}

type RunnerReleaseChannel string

const (
	RunnerReleaseChannelUnspecified RunnerReleaseChannel = "RUNNER_RELEASE_CHANNEL_UNSPECIFIED"
	RunnerReleaseChannelStable      RunnerReleaseChannel = "RUNNER_RELEASE_CHANNEL_STABLE"
	RunnerReleaseChannelLatest      RunnerReleaseChannel = "RUNNER_RELEASE_CHANNEL_LATEST"
)

func (r RunnerReleaseChannel) IsKnown() bool {
	switch r {
	case RunnerReleaseChannelUnspecified, RunnerReleaseChannelStable, RunnerReleaseChannelLatest:
		return true
	}
	return false
}

type RunnerSpec struct {
	// The runner's configuration
	Configuration RunnerConfiguration `json:"configuration"`
	// RunnerPhase represents the phase a runner is in
	DesiredPhase RunnerPhase    `json:"desiredPhase"`
	JSON         runnerSpecJSON `json:"-"`
}

// runnerSpecJSON contains the JSON metadata for the struct [RunnerSpec]
type runnerSpecJSON struct {
	Configuration apijson.Field
	DesiredPhase  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerSpecJSON) RawJSON() string {
	return r.raw
}

type RunnerSpecParam struct {
	// The runner's configuration
	Configuration param.Field[RunnerConfigurationParam] `json:"configuration"`
	// RunnerPhase represents the phase a runner is in
	DesiredPhase param.Field[RunnerPhase] `json:"desiredPhase"`
}

func (r RunnerSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// RunnerStatus represents the status of a runner
type RunnerStatus struct {
	// additional_info contains additional information about the runner, e.g. a
	// CloudFormation stack URL.
	AdditionalInfo []shared.FieldValue `json:"additionalInfo"`
	// capabilities is a list of capabilities the runner supports.
	Capabilities []RunnerCapability `json:"capabilities"`
	LogURL       string             `json:"logUrl"`
	// The runner's reported message which is shown to users. This message adds more
	// context to the runner's phase.
	Message string `json:"message"`
	// RunnerPhase represents the phase a runner is in
	Phase RunnerPhase `json:"phase"`
	// region is the region the runner is running in, if applicable.
	Region        string `json:"region"`
	SystemDetails string `json:"systemDetails"`
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
	UpdatedAt time.Time        `json:"updatedAt" format:"date-time"`
	Version   string           `json:"version"`
	JSON      runnerStatusJSON `json:"-"`
}

// runnerStatusJSON contains the JSON metadata for the struct [RunnerStatus]
type runnerStatusJSON struct {
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

func (r *RunnerStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerStatusJSON) RawJSON() string {
	return r.raw
}

type RunnerNewResponse struct {
	// deprecated, will be removed. Use exchange_token instead.
	AccessToken string `json:"accessToken"`
	// exchange_token is a one-time use token that should be exchanged by the runner
	// for an access token, using the IdentityService.ExchangeToken rpc. The token
	// expires after 24 hours.
	ExchangeToken string                `json:"exchangeToken"`
	Runner        Runner                `json:"runner"`
	JSON          runnerNewResponseJSON `json:"-"`
}

// runnerNewResponseJSON contains the JSON metadata for the struct
// [RunnerNewResponse]
type runnerNewResponseJSON struct {
	AccessToken   apijson.Field
	ExchangeToken apijson.Field
	Runner        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerNewResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerGetResponse struct {
	Runner Runner                `json:"runner"`
	JSON   runnerGetResponseJSON `json:"-"`
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

type RunnerUpdateResponse = interface{}

type RunnerDeleteResponse = interface{}

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
	// deprecated, will be removed. Use exchange_token instead.
	AccessToken string `json:"accessToken"`
	// exchange_token is a one-time use token that should be exchanged by the runner
	// for an access token, using the IdentityService.ExchangeToken rpc. The token
	// expires after 24 hours.
	ExchangeToken string                           `json:"exchangeToken"`
	JSON          runnerNewRunnerTokenResponseJSON `json:"-"`
}

// runnerNewRunnerTokenResponseJSON contains the JSON metadata for the struct
// [RunnerNewRunnerTokenResponse]
type runnerNewRunnerTokenResponseJSON struct {
	AccessToken   apijson.Field
	ExchangeToken apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerNewRunnerTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerNewRunnerTokenResponseJSON) RawJSON() string {
	return r.raw
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

type RunnerNewParams struct {
	// RunnerKind represents the kind of a runner
	Kind param.Field[RunnerKind] `json:"kind"`
	// The runner name for humans
	Name param.Field[string] `json:"name"`
	// RunnerProvider identifies the specific implementation type of a runner. Each
	// provider maps to a specific kind of runner (local or remote), as specified below
	// for each provider.
	Provider param.Field[RunnerProvider]  `json:"provider"`
	Spec     param.Field[RunnerSpecParam] `json:"spec"`
}

func (r RunnerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerGetParams struct {
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerUpdateParams struct {
	// The runner's name which is shown to users
	Name param.Field[string] `json:"name"`
	// runner_id specifies which runner to be updated.
	//
	// +required
	RunnerID param.Field[string]                 `json:"runnerId" format:"uuid"`
	Spec     param.Field[RunnerUpdateParamsSpec] `json:"spec"`
}

func (r RunnerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerUpdateParamsSpec struct {
	Configuration param.Field[RunnerUpdateParamsSpecConfiguration] `json:"configuration"`
	// RunnerPhase represents the phase a runner is in
	DesiredPhase param.Field[RunnerPhase] `json:"desiredPhase"`
}

func (r RunnerUpdateParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerUpdateParamsSpecConfiguration struct {
	// auto_update indicates whether the runner should automatically update itself.
	AutoUpdate param.Field[bool] `json:"autoUpdate"`
	// The release channel the runner is on
	ReleaseChannel param.Field[RunnerReleaseChannel] `json:"releaseChannel"`
}

func (r RunnerUpdateParamsSpecConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerListParams struct {
	Token    param.Field[string]                 `query:"token"`
	PageSize param.Field[int64]                  `query:"pageSize"`
	Filter   param.Field[RunnerListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing runners
	Pagination param.Field[RunnerListParamsPagination] `json:"pagination"`
}

func (r RunnerListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [RunnerListParams]'s query parameters as `url.Values`.
func (r RunnerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RunnerListParamsFilter struct {
	// creator_ids filters the response to only runner created by specified users
	CreatorIDs param.Field[[]string] `json:"creatorIds" format:"uuid"`
	// kinds filters the response to only runners of the specified kinds
	Kinds param.Field[[]RunnerKind] `json:"kinds"`
	// providers filters the response to only runners of the specified providers
	Providers param.Field[[]RunnerProvider] `json:"providers"`
}

func (r RunnerListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing runners
type RunnerListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerDeleteParams struct {
	// force indicates whether the runner should be deleted forcefully. When force
	// deleting a Runner, all Environments on the runner are also force deleted and
	// regular Runner lifecycle is not respected. Force deleting can result in data
	// loss.
	Force    param.Field[bool]   `json:"force"`
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerDeleteParams) MarshalJSON() (data []byte, err error) {
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

type RunnerParseContextURLParams struct {
	ContextURL param.Field[string] `json:"contextUrl" format:"uri"`
	RunnerID   param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerParseContextURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
