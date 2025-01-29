// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apiform"
	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// RunnerInteractionService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerInteractionService] method instead.
type RunnerInteractionService struct {
	Options      []option.RequestOption
	Environments *RunnerInteractionEnvironmentService
}

// NewRunnerInteractionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRunnerInteractionService(opts ...option.RequestOption) (r *RunnerInteractionService) {
	r = &RunnerInteractionService{}
	r.Options = opts
	r.Environments = NewRunnerInteractionEnvironmentService(opts...)
	return
}

// GetRunnerHostAuthenticationToken returns an authentication token for the given
// host.
func (r *RunnerInteractionService) GetHostAuthenticationTokenValue(ctx context.Context, params RunnerInteractionGetHostAuthenticationTokenValueParams, opts ...option.RequestOption) (res *RunnerInteractionGetHostAuthenticationTokenValueResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/GetHostAuthenticationTokenValue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetLatestVersion returns the latest version of the runner.
func (r *RunnerInteractionService) GetLatestVersion(ctx context.Context, params RunnerInteractionGetLatestVersionParams, opts ...option.RequestOption) (res *RunnerInteractionGetLatestVersionResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/GetLatestVersion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListRunnerEnvironmentClasses returns the environment classes configured for the
// runner.
func (r *RunnerInteractionService) ListRunnerEnvironmentClasses(ctx context.Context, params RunnerInteractionListRunnerEnvironmentClassesParams, opts ...option.RequestOption) (res *RunnerInteractionListRunnerEnvironmentClassesResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/ListRunnerEnvironmentClasses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListRunnerSCMIntegrations returns the SCM integrations configured for the
// runner. For local runners, this returns the SCM integrations configured on the
// organization's local-configuration runner.
func (r *RunnerInteractionService) ListRunnerScmIntegrations(ctx context.Context, params RunnerInteractionListRunnerScmIntegrationsParams, opts ...option.RequestOption) (res *RunnerInteractionListRunnerScmIntegrationsResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/ListRunnerSCMIntegrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// MarkRunnerActive marks a runner as available. This must be called every 30
// seconds to keep the runner active.
func (r *RunnerInteractionService) MarkActive(ctx context.Context, params RunnerInteractionMarkActiveParams, opts ...option.RequestOption) (res *RunnerInteractionMarkActiveResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/MarkRunnerActive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// SendResponse sends a response to a request.
func (r *RunnerInteractionService) SendResponse(ctx context.Context, params RunnerInteractionSendResponseParams, opts ...option.RequestOption) (res *RunnerInteractionSendResponseResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/SendResponse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Signup is called by a runner to sign up with the backend. This is the first call
// a runner makes.
func (r *RunnerInteractionService) Signup(ctx context.Context, params RunnerInteractionSignupParams, opts ...option.RequestOption) (res *RunnerInteractionSignupResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/Signup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateRunnerConfigurationSchema updates the runner's configuration schema.
func (r *RunnerInteractionService) UpdateRunnerConfigurationSchema(ctx context.Context, params RunnerInteractionUpdateRunnerConfigurationSchemaParams, opts ...option.RequestOption) (res *RunnerInteractionUpdateRunnerConfigurationSchemaResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/UpdateRunnerConfigurationSchema"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateRunnerStatus updates the status of the runner.
func (r *RunnerInteractionService) UpdateStatus(ctx context.Context, params RunnerInteractionUpdateStatusParams, opts ...option.RequestOption) (res *RunnerInteractionUpdateStatusResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/UpdateRunnerStatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type RunnerInteractionGetHostAuthenticationTokenValueResponse struct {
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
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
	// The host authentication token's refresh token encrypted as NaCL anonymous sealed
	// box using the runner's public key
	RefreshToken string `json:"refreshToken" format:"byte"`
	// The host authentication token's source
	Source RunnerInteractionGetHostAuthenticationTokenValueResponseSource `json:"source"`
	// The host authentication token's ID
	TokenID string `json:"tokenId"`
	// The authentication token encrypted as NaCL anonymous sealed box using the
	// runner's public key
	Value string                                                       `json:"value" format:"byte"`
	JSON  runnerInteractionGetHostAuthenticationTokenValueResponseJSON `json:"-"`
}

// runnerInteractionGetHostAuthenticationTokenValueResponseJSON contains the JSON
// metadata for the struct
// [RunnerInteractionGetHostAuthenticationTokenValueResponse]
type runnerInteractionGetHostAuthenticationTokenValueResponseJSON struct {
	ExpiresAt    apijson.Field
	RefreshToken apijson.Field
	Source       apijson.Field
	TokenID      apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RunnerInteractionGetHostAuthenticationTokenValueResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionGetHostAuthenticationTokenValueResponseJSON) RawJSON() string {
	return r.raw
}

// The host authentication token's source
type RunnerInteractionGetHostAuthenticationTokenValueResponseSource string

const (
	RunnerInteractionGetHostAuthenticationTokenValueResponseSourceHostAuthenticationTokenSourceUnspecified RunnerInteractionGetHostAuthenticationTokenValueResponseSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_UNSPECIFIED"
	RunnerInteractionGetHostAuthenticationTokenValueResponseSourceHostAuthenticationTokenSourceOAuth       RunnerInteractionGetHostAuthenticationTokenValueResponseSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_OAUTH"
	RunnerInteractionGetHostAuthenticationTokenValueResponseSourceHostAuthenticationTokenSourcePat         RunnerInteractionGetHostAuthenticationTokenValueResponseSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_PAT"
)

func (r RunnerInteractionGetHostAuthenticationTokenValueResponseSource) IsKnown() bool {
	switch r {
	case RunnerInteractionGetHostAuthenticationTokenValueResponseSourceHostAuthenticationTokenSourceUnspecified, RunnerInteractionGetHostAuthenticationTokenValueResponseSourceHostAuthenticationTokenSourceOAuth, RunnerInteractionGetHostAuthenticationTokenValueResponseSourceHostAuthenticationTokenSourcePat:
		return true
	}
	return false
}

type RunnerInteractionGetLatestVersionResponse struct {
	// auto-update indicates if the runner should be updated automatically
	AutoUpdate bool `json:"autoUpdate"`
	// gitpod_cli_download_url is the URL to download the gitpod CLI
	GitpodCliDownloadURL string `json:"gitpodCliDownloadUrl"`
	// The container image of the runner
	RunnerImage string `json:"runnerImage"`
	// supervisor_download_url is the URL to download the supervisor
	SupervisorDownloadURL string `json:"supervisorDownloadUrl"`
	// The latest version of the runner
	Version string                                        `json:"version"`
	JSON    runnerInteractionGetLatestVersionResponseJSON `json:"-"`
}

// runnerInteractionGetLatestVersionResponseJSON contains the JSON metadata for the
// struct [RunnerInteractionGetLatestVersionResponse]
type runnerInteractionGetLatestVersionResponseJSON struct {
	AutoUpdate            apijson.Field
	GitpodCliDownloadURL  apijson.Field
	RunnerImage           apijson.Field
	SupervisorDownloadURL apijson.Field
	Version               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RunnerInteractionGetLatestVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionGetLatestVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionListRunnerEnvironmentClassesResponse struct {
	// The environment classes configured for the runner
	EnvironmentClasses []RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClass `json:"environmentClasses"`
	// pagination contains the pagination options for listing environment classes
	Pagination RunnerInteractionListRunnerEnvironmentClassesResponsePagination `json:"pagination"`
	JSON       runnerInteractionListRunnerEnvironmentClassesResponseJSON       `json:"-"`
}

// runnerInteractionListRunnerEnvironmentClassesResponseJSON contains the JSON
// metadata for the struct [RunnerInteractionListRunnerEnvironmentClassesResponse]
type runnerInteractionListRunnerEnvironmentClassesResponseJSON struct {
	EnvironmentClasses apijson.Field
	Pagination         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerEnvironmentClassesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerEnvironmentClassesResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID string `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration []RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfiguration `json:"configuration"`
	// description is a human readable description of the environment class
	Description string `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName string `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled bool `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID string                                                                    `json:"runnerId"`
	JSON     runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassJSON `json:"-"`
}

// runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassJSON
// contains the JSON metadata for the struct
// [RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClass]
type runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	RunnerID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfiguration struct {
	Key   string                                                                                   `json:"key"`
	Value string                                                                                   `json:"value"`
	JSON  runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfigurationJSON `json:"-"`
}

// runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfigurationJSON
// contains the JSON metadata for the struct
// [RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfiguration]
type runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfigurationJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfigurationJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing environment classes
type RunnerInteractionListRunnerEnvironmentClassesResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                                              `json:"nextToken"`
	JSON      runnerInteractionListRunnerEnvironmentClassesResponsePaginationJSON `json:"-"`
}

// runnerInteractionListRunnerEnvironmentClassesResponsePaginationJSON contains the
// JSON metadata for the struct
// [RunnerInteractionListRunnerEnvironmentClassesResponsePagination]
type runnerInteractionListRunnerEnvironmentClassesResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerEnvironmentClassesResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerEnvironmentClassesResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionListRunnerScmIntegrationsResponse struct {
	// pagination contains the pagination options for listing SCM integrations
	Pagination RunnerInteractionListRunnerScmIntegrationsResponsePagination `json:"pagination"`
	// The SCM integrations configured for the runner
	ScmIntegrations []RunnerInteractionListRunnerScmIntegrationsResponseScmIntegration `json:"scmIntegrations"`
	JSON            runnerInteractionListRunnerScmIntegrationsResponseJSON             `json:"-"`
}

// runnerInteractionListRunnerScmIntegrationsResponseJSON contains the JSON
// metadata for the struct [RunnerInteractionListRunnerScmIntegrationsResponse]
type runnerInteractionListRunnerScmIntegrationsResponseJSON struct {
	Pagination      apijson.Field
	ScmIntegrations apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerScmIntegrationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerScmIntegrationsResponseJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing SCM integrations
type RunnerInteractionListRunnerScmIntegrationsResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                                           `json:"nextToken"`
	JSON      runnerInteractionListRunnerScmIntegrationsResponsePaginationJSON `json:"-"`
}

// runnerInteractionListRunnerScmIntegrationsResponsePaginationJSON contains the
// JSON metadata for the struct
// [RunnerInteractionListRunnerScmIntegrationsResponsePagination]
type runnerInteractionListRunnerScmIntegrationsResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerScmIntegrationsResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerScmIntegrationsResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionListRunnerScmIntegrationsResponseScmIntegration struct {
	JSON runnerInteractionListRunnerScmIntegrationsResponseScmIntegrationJSON `json:"-"`
}

// runnerInteractionListRunnerScmIntegrationsResponseScmIntegrationJSON contains
// the JSON metadata for the struct
// [RunnerInteractionListRunnerScmIntegrationsResponseScmIntegration]
type runnerInteractionListRunnerScmIntegrationsResponseScmIntegrationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerScmIntegrationsResponseScmIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerScmIntegrationsResponseScmIntegrationJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionMarkActiveResponse = interface{}

type RunnerInteractionSendResponseResponse = interface{}

type RunnerInteractionSignupResponse struct {
	// The runner's identity
	RunnerID string                              `json:"runnerId" format:"uuid"`
	JSON     runnerInteractionSignupResponseJSON `json:"-"`
}

// runnerInteractionSignupResponseJSON contains the JSON metadata for the struct
// [RunnerInteractionSignupResponse]
type runnerInteractionSignupResponseJSON struct {
	RunnerID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionSignupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionSignupResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionUpdateRunnerConfigurationSchemaResponse = interface{}

type RunnerInteractionUpdateStatusResponse = interface{}

type RunnerInteractionGetHostAuthenticationTokenValueParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// The host to get the authentication token for
	Host param.Field[string] `json:"host"`
	// The principal's ID to get the authentication token for
	PrincipalID param.Field[string] `json:"principalId" format:"uuid"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionGetHostAuthenticationTokenValueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion float64

const (
	RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion1 RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionGetLatestVersionParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionGetLatestVersionParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// The current version of the runner
	CurrentVersion param.Field[string] `json:"currentVersion"`
	// The version of the infrastructure
	InfrastructureVersion param.Field[string] `json:"infrastructureVersion"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionGetLatestVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionGetLatestVersionParamsConnectProtocolVersion float64

const (
	RunnerInteractionGetLatestVersionParamsConnectProtocolVersion1 RunnerInteractionGetLatestVersionParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionGetLatestVersionParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionGetLatestVersionParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionListRunnerEnvironmentClassesParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Filter                 param.Field[RunnerInteractionListRunnerEnvironmentClassesParamsFilter]                 `json:"filter"`
	// pagination contains the pagination options for listing environment classes
	Pagination param.Field[RunnerInteractionListRunnerEnvironmentClassesParamsPagination] `json:"pagination"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionListRunnerEnvironmentClassesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion float64

const (
	RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion1 RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionListRunnerEnvironmentClassesParamsFilter struct {
	// environment_class_ids filters the response to only environment classes with
	// these IDs
	EnvironmentClassIDs param.Field[[]string] `json:"environmentClassIds" format:"uuid"`
}

func (r RunnerInteractionListRunnerEnvironmentClassesParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing environment classes
type RunnerInteractionListRunnerEnvironmentClassesParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerInteractionListRunnerEnvironmentClassesParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionListRunnerScmIntegrationsParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Filter                 param.Field[RunnerInteractionListRunnerScmIntegrationsParamsFilter]                 `json:"filter"`
	// pagination contains the pagination options for listing SCM integrations
	Pagination param.Field[RunnerInteractionListRunnerScmIntegrationsParamsPagination] `json:"pagination"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionListRunnerScmIntegrationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion float64

const (
	RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion1 RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionListRunnerScmIntegrationsParamsFilter struct {
	// environment_class_ids filters the response to only SCM integrations with these
	// IDs
	ScmIntegrationIDs param.Field[[]string] `json:"scmIntegrationIds" format:"uuid"`
}

func (r RunnerInteractionListRunnerScmIntegrationsParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing SCM integrations
type RunnerInteractionListRunnerScmIntegrationsParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerInteractionListRunnerScmIntegrationsParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionMarkActiveParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionMarkActiveParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionMarkActiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionMarkActiveParamsConnectProtocolVersion float64

const (
	RunnerInteractionMarkActiveParamsConnectProtocolVersion1 RunnerInteractionMarkActiveParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionMarkActiveParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionMarkActiveParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionSendResponseParams struct {
	Body RunnerInteractionSendResponseParamsBody `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionSendResponseParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionSendResponseParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type RunnerInteractionSendResponseParamsBody struct {
}

func (r RunnerInteractionSendResponseParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionSendResponseParamsConnectProtocolVersion float64

const (
	RunnerInteractionSendResponseParamsConnectProtocolVersion1 RunnerInteractionSendResponseParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionSendResponseParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionSendResponseParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionSignupParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionSignupParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// The environment classes this runner has to offer
	EnvironmentClasses param.Field[[]RunnerInteractionSignupParamsEnvironmentClass] `json:"environmentClasses"`
	// The runner's public key. Must be an ECDH public key encoded in PKIX, ASN.1 DER
	// format.
	PublicKey param.Field[string] `json:"publicKey" format:"byte"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionSignupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionSignupParamsConnectProtocolVersion float64

const (
	RunnerInteractionSignupParamsConnectProtocolVersion1 RunnerInteractionSignupParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionSignupParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionSignupParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionSignupParamsEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID param.Field[string] `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration param.Field[[]RunnerInteractionSignupParamsEnvironmentClassesConfiguration] `json:"configuration"`
	// description is a human readable description of the environment class
	Description param.Field[string] `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName param.Field[string] `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled param.Field[bool] `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID param.Field[string] `json:"runnerId"`
}

func (r RunnerInteractionSignupParamsEnvironmentClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionSignupParamsEnvironmentClassesConfiguration struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r RunnerInteractionSignupParamsEnvironmentClassesConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionUpdateRunnerConfigurationSchemaParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// config_schema is the schema for the runner's configuration
	ConfigSchema param.Field[RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchema] `json:"configSchema"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion float64

const (
	RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion1 RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// config_schema is the schema for the runner's configuration
type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchema struct {
	EnvironmentClasses param.Field[[]RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClass] `json:"environmentClasses"`
	RunnerConfig       param.Field[[]RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfig]     `json:"runnerConfig"`
	Scm                param.Field[[]RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScm]              `json:"scm"`
	// The schema version
	Version param.Field[string] `json:"version"`
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClass struct {
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfig struct {
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScm struct {
	DefaultHosts param.Field[[]string]                                                                   `json:"defaultHosts"`
	Name         param.Field[string]                                                                     `json:"name"`
	OAuth        param.Field[RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmOAuth] `json:"oauth"`
	Pat          param.Field[RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmPat]   `json:"pat"`
	ScmID        param.Field[string]                                                                     `json:"scmId"`
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScm) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmOAuth struct {
	// callback_url is the URL the OAuth app will redirect to after the user has
	// authenticated.
	CallbackURL param.Field[string] `json:"callbackUrl"`
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmOAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmPat struct {
	// description is a human-readable description of the PAT.
	Description param.Field[string] `json:"description"`
	// docs_link is a link to the documentation on how to create a PAT for this SCM
	// system.
	DocsLink param.Field[string] `json:"docsLink"`
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmPat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionUpdateStatusParams struct {
	Body RunnerInteractionUpdateStatusParamsBody `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionUpdateStatusParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerInteractionUpdateStatusParamsBody struct {
}

func (r RunnerInteractionUpdateStatusParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionUpdateStatusParamsConnectProtocolVersion float64

const (
	RunnerInteractionUpdateStatusParamsConnectProtocolVersion1 RunnerInteractionUpdateStatusParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionUpdateStatusParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionUpdateStatusParamsConnectProtocolVersion1:
		return true
	}
	return false
}
