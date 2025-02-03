// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
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
//
// short-lived and must be renewed every 30 seconds. Runners can be registered for
// an entire organisation or a single user.
func (r *RunnerService) New(ctx context.Context, params RunnerNewParams, opts ...option.RequestOption) (res *RunnerNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CreateRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetRunner returns a single runner.
func (r *RunnerService) Get(ctx context.Context, params RunnerGetParams, opts ...option.RequestOption) (res *RunnerGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/GetRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// UpdateRunner updates an environment runner.
func (r *RunnerService) Update(ctx context.Context, params RunnerUpdateParams, opts ...option.RequestOption) (res *RunnerUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/UpdateRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListRunners returns all runners registered in the scope.
func (r *RunnerService) List(ctx context.Context, params RunnerListParams, opts ...option.RequestOption) (res *RunnerListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/ListRunners"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// DeleteRunner deletes an environment runner.
func (r *RunnerService) Delete(ctx context.Context, params RunnerDeleteParams, opts ...option.RequestOption) (res *RunnerDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/DeleteRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// CheckAuthenticationForHost asks a runner if the user is authenticated against a
// particular host, e.g. an SCM system.
//
// If not, this function will return a URL that the user should visit to
// authenticate, or indicate that Personal Access Tokens are supported.
func (r *RunnerService) CheckAuthenticationForHost(ctx context.Context, params RunnerCheckAuthenticationForHostParams, opts ...option.RequestOption) (res *RunnerCheckAuthenticationForHostResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CheckAuthenticationForHost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// CreateRunnerToken returns a token that can be used to authenticate as the
//
// runner. Use this call to renew an outdated token - this does not expire any
// previouly issued tokens.
func (r *RunnerService) NewRunnerToken(ctx context.Context, params RunnerNewRunnerTokenParams, opts ...option.RequestOption) (res *RunnerNewRunnerTokenResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CreateRunnerToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
func (r *RunnerService) ParseContextURL(ctx context.Context, params RunnerParseContextURLParams, opts ...option.RequestOption) (res *RunnerParseContextURLResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/ParseContextURL"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
	Name string `json:"name"`
	// RunnerProvider identifies the specific implementation type of a runner. Each
	// provider maps to a specific kind of runner (local or remote), as specified below
	// for each provider.
	Provider RunnerNewResponseRunnerProvider `json:"provider"`
	RunnerID string                          `json:"runnerId"`
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
	Provider    apijson.Field
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

// RunnerProvider identifies the specific implementation type of a runner. Each
// provider maps to a specific kind of runner (local or remote), as specified below
// for each provider.
type RunnerNewResponseRunnerProvider string

const (
	RunnerNewResponseRunnerProviderRunnerProviderUnspecified RunnerNewResponseRunnerProvider = "RUNNER_PROVIDER_UNSPECIFIED"
	RunnerNewResponseRunnerProviderRunnerProviderAwsEc2      RunnerNewResponseRunnerProvider = "RUNNER_PROVIDER_AWS_EC2"
	RunnerNewResponseRunnerProviderRunnerProviderLinuxHost   RunnerNewResponseRunnerProvider = "RUNNER_PROVIDER_LINUX_HOST"
	RunnerNewResponseRunnerProviderRunnerProviderDesktopMac  RunnerNewResponseRunnerProvider = "RUNNER_PROVIDER_DESKTOP_MAC"
)

func (r RunnerNewResponseRunnerProvider) IsKnown() bool {
	switch r {
	case RunnerNewResponseRunnerProviderRunnerProviderUnspecified, RunnerNewResponseRunnerProviderRunnerProviderAwsEc2, RunnerNewResponseRunnerProviderRunnerProviderLinuxHost, RunnerNewResponseRunnerProviderRunnerProviderDesktopMac:
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
	Name string `json:"name"`
	// RunnerProvider identifies the specific implementation type of a runner. Each
	// provider maps to a specific kind of runner (local or remote), as specified below
	// for each provider.
	Provider RunnerGetResponseRunnerProvider `json:"provider"`
	RunnerID string                          `json:"runnerId"`
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
	Provider    apijson.Field
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

// RunnerProvider identifies the specific implementation type of a runner. Each
// provider maps to a specific kind of runner (local or remote), as specified below
// for each provider.
type RunnerGetResponseRunnerProvider string

const (
	RunnerGetResponseRunnerProviderRunnerProviderUnspecified RunnerGetResponseRunnerProvider = "RUNNER_PROVIDER_UNSPECIFIED"
	RunnerGetResponseRunnerProviderRunnerProviderAwsEc2      RunnerGetResponseRunnerProvider = "RUNNER_PROVIDER_AWS_EC2"
	RunnerGetResponseRunnerProviderRunnerProviderLinuxHost   RunnerGetResponseRunnerProvider = "RUNNER_PROVIDER_LINUX_HOST"
	RunnerGetResponseRunnerProviderRunnerProviderDesktopMac  RunnerGetResponseRunnerProvider = "RUNNER_PROVIDER_DESKTOP_MAC"
)

func (r RunnerGetResponseRunnerProvider) IsKnown() bool {
	switch r {
	case RunnerGetResponseRunnerProviderRunnerProviderUnspecified, RunnerGetResponseRunnerProviderRunnerProviderAwsEc2, RunnerGetResponseRunnerProviderRunnerProviderLinuxHost, RunnerGetResponseRunnerProviderRunnerProviderDesktopMac:
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

type RunnerUpdateResponse = interface{}

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
	Name string `json:"name"`
	// RunnerProvider identifies the specific implementation type of a runner. Each
	// provider maps to a specific kind of runner (local or remote), as specified below
	// for each provider.
	Provider RunnerListResponseRunnersProvider `json:"provider"`
	RunnerID string                            `json:"runnerId"`
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
	Provider    apijson.Field
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

// RunnerProvider identifies the specific implementation type of a runner. Each
// provider maps to a specific kind of runner (local or remote), as specified below
// for each provider.
type RunnerListResponseRunnersProvider string

const (
	RunnerListResponseRunnersProviderRunnerProviderUnspecified RunnerListResponseRunnersProvider = "RUNNER_PROVIDER_UNSPECIFIED"
	RunnerListResponseRunnersProviderRunnerProviderAwsEc2      RunnerListResponseRunnersProvider = "RUNNER_PROVIDER_AWS_EC2"
	RunnerListResponseRunnersProviderRunnerProviderLinuxHost   RunnerListResponseRunnersProvider = "RUNNER_PROVIDER_LINUX_HOST"
	RunnerListResponseRunnersProviderRunnerProviderDesktopMac  RunnerListResponseRunnersProvider = "RUNNER_PROVIDER_DESKTOP_MAC"
)

func (r RunnerListResponseRunnersProvider) IsKnown() bool {
	switch r {
	case RunnerListResponseRunnersProviderRunnerProviderUnspecified, RunnerListResponseRunnersProviderRunnerProviderAwsEc2, RunnerListResponseRunnersProviderRunnerProviderLinuxHost, RunnerListResponseRunnersProviderRunnerProviderDesktopMac:
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
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// RunnerKind represents the kind of a runner
	Kind param.Field[RunnerNewParamsKind] `json:"kind"`
	// The runner name for humans
	Name param.Field[string] `json:"name"`
	// RunnerProvider identifies the specific implementation type of a runner. Each
	// provider maps to a specific kind of runner (local or remote), as specified below
	// for each provider.
	Provider param.Field[RunnerNewParamsProvider] `json:"provider"`
	Spec     param.Field[RunnerNewParamsSpec]     `json:"spec"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerNewParamsConnectProtocolVersion float64

const (
	RunnerNewParamsConnectProtocolVersion1 RunnerNewParamsConnectProtocolVersion = 1
)

func (r RunnerNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerNewParamsConnectProtocolVersion1:
		return true
	}
	return false
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

// RunnerProvider identifies the specific implementation type of a runner. Each
// provider maps to a specific kind of runner (local or remote), as specified below
// for each provider.
type RunnerNewParamsProvider string

const (
	RunnerNewParamsProviderRunnerProviderUnspecified RunnerNewParamsProvider = "RUNNER_PROVIDER_UNSPECIFIED"
	RunnerNewParamsProviderRunnerProviderAwsEc2      RunnerNewParamsProvider = "RUNNER_PROVIDER_AWS_EC2"
	RunnerNewParamsProviderRunnerProviderLinuxHost   RunnerNewParamsProvider = "RUNNER_PROVIDER_LINUX_HOST"
	RunnerNewParamsProviderRunnerProviderDesktopMac  RunnerNewParamsProvider = "RUNNER_PROVIDER_DESKTOP_MAC"
)

func (r RunnerNewParamsProvider) IsKnown() bool {
	switch r {
	case RunnerNewParamsProviderRunnerProviderUnspecified, RunnerNewParamsProviderRunnerProviderAwsEc2, RunnerNewParamsProviderRunnerProviderLinuxHost, RunnerNewParamsProviderRunnerProviderDesktopMac:
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
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[RunnerGetParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[RunnerGetParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[RunnerGetParamsConnect] `query:"connect"`
	Message param.Field[string]                 `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [RunnerGetParams]'s query parameters as `url.Values`.
func (r RunnerGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type RunnerGetParamsEncoding string

const (
	RunnerGetParamsEncodingProto RunnerGetParamsEncoding = "proto"
	RunnerGetParamsEncodingJson  RunnerGetParamsEncoding = "json"
)

func (r RunnerGetParamsEncoding) IsKnown() bool {
	switch r {
	case RunnerGetParamsEncodingProto, RunnerGetParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerGetParamsConnectProtocolVersion float64

const (
	RunnerGetParamsConnectProtocolVersion1 RunnerGetParamsConnectProtocolVersion = 1
)

func (r RunnerGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type RunnerGetParamsCompression string

const (
	RunnerGetParamsCompressionIdentity RunnerGetParamsCompression = "identity"
	RunnerGetParamsCompressionGzip     RunnerGetParamsCompression = "gzip"
	RunnerGetParamsCompressionBr       RunnerGetParamsCompression = "br"
)

func (r RunnerGetParamsCompression) IsKnown() bool {
	switch r {
	case RunnerGetParamsCompressionIdentity, RunnerGetParamsCompressionGzip, RunnerGetParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerGetParamsConnect string

const (
	RunnerGetParamsConnectV1 RunnerGetParamsConnect = "v1"
)

func (r RunnerGetParamsConnect) IsKnown() bool {
	switch r {
	case RunnerGetParamsConnectV1:
		return true
	}
	return false
}

type RunnerUpdateParams struct {
	Body RunnerUpdateParamsBody `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerUpdateParamsBody struct {
}

func (r RunnerUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerUpdateParamsConnectProtocolVersion float64

const (
	RunnerUpdateParamsConnectProtocolVersion1 RunnerUpdateParamsConnectProtocolVersion = 1
)

func (r RunnerUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[RunnerListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[RunnerListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[RunnerListParamsConnect] `query:"connect"`
	Message param.Field[string]                  `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [RunnerListParams]'s query parameters as `url.Values`.
func (r RunnerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type RunnerListParamsEncoding string

const (
	RunnerListParamsEncodingProto RunnerListParamsEncoding = "proto"
	RunnerListParamsEncodingJson  RunnerListParamsEncoding = "json"
)

func (r RunnerListParamsEncoding) IsKnown() bool {
	switch r {
	case RunnerListParamsEncodingProto, RunnerListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerListParamsConnectProtocolVersion float64

const (
	RunnerListParamsConnectProtocolVersion1 RunnerListParamsConnectProtocolVersion = 1
)

func (r RunnerListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type RunnerListParamsCompression string

const (
	RunnerListParamsCompressionIdentity RunnerListParamsCompression = "identity"
	RunnerListParamsCompressionGzip     RunnerListParamsCompression = "gzip"
	RunnerListParamsCompressionBr       RunnerListParamsCompression = "br"
)

func (r RunnerListParamsCompression) IsKnown() bool {
	switch r {
	case RunnerListParamsCompressionIdentity, RunnerListParamsCompressionGzip, RunnerListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerListParamsConnect string

const (
	RunnerListParamsConnectV1 RunnerListParamsConnect = "v1"
)

func (r RunnerListParamsConnect) IsKnown() bool {
	switch r {
	case RunnerListParamsConnectV1:
		return true
	}
	return false
}

type RunnerDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// force indicates whether the runner should be deleted forcefully. When force
	// deleting a Runner, all Environments on the runner are also force deleted and
	// regular Runner lifecycle is not respected. Force deleting can result in data
	// loss.
	Force    param.Field[bool]   `json:"force"`
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerDeleteParamsConnectProtocolVersion float64

const (
	RunnerDeleteParamsConnectProtocolVersion1 RunnerDeleteParamsConnectProtocolVersion = 1
)

func (r RunnerDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerCheckAuthenticationForHostParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerCheckAuthenticationForHostParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Host                   param.Field[string]                                                       `json:"host"`
	RunnerID               param.Field[string]                                                       `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerCheckAuthenticationForHostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerCheckAuthenticationForHostParamsConnectProtocolVersion float64

const (
	RunnerCheckAuthenticationForHostParamsConnectProtocolVersion1 RunnerCheckAuthenticationForHostParamsConnectProtocolVersion = 1
)

func (r RunnerCheckAuthenticationForHostParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerCheckAuthenticationForHostParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerNewRunnerTokenParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerNewRunnerTokenParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	RunnerID               param.Field[string]                                           `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerNewRunnerTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerNewRunnerTokenParamsConnectProtocolVersion float64

const (
	RunnerNewRunnerTokenParamsConnectProtocolVersion1 RunnerNewRunnerTokenParamsConnectProtocolVersion = 1
)

func (r RunnerNewRunnerTokenParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerNewRunnerTokenParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerParseContextURLParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerParseContextURLParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ContextURL             param.Field[string]                                            `json:"contextUrl" format:"uri"`
	RunnerID               param.Field[string]                                            `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerParseContextURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerParseContextURLParamsConnectProtocolVersion float64

const (
	RunnerParseContextURLParamsConnectProtocolVersion1 RunnerParseContextURLParamsConnectProtocolVersion = 1
)

func (r RunnerParseContextURLParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerParseContextURLParamsConnectProtocolVersion1:
		return true
	}
	return false
}
