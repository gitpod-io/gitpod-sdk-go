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

// Creates a new runner registration with the server. Registrations are very
// short-lived and must be renewed every 30 seconds.
//
// Use this method to:
//
// - Register organization runners
// - Set up runner configurations
// - Initialize runner credentials
// - Configure auto-updates
//
// ### Examples
//
// - Create cloud runner:
//
//	Creates a new runner in AWS EC2.
//
//	```yaml
//	name: "Production Runner"
//	provider: RUNNER_PROVIDER_AWS_EC2
//	spec:
//	  desiredPhase: RUNNER_PHASE_ACTIVE
//	  configuration:
//	    region: "us-west"
//	    releaseChannel: RUNNER_RELEASE_CHANNEL_STABLE
//	    autoUpdate: true
//	```
//
// - Create local runner:
//
//	Creates a new local runner on Linux.
//
//	```yaml
//	name: "Local Development Runner"
//	provider: RUNNER_PROVIDER_LINUX_HOST
//	spec:
//	  desiredPhase: RUNNER_PHASE_ACTIVE
//	  configuration:
//	    releaseChannel: RUNNER_RELEASE_CHANNEL_LATEST
//	    autoUpdate: true
//	```
func (r *RunnerService) New(ctx context.Context, body RunnerNewParams, opts ...option.RequestOption) (res *RunnerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CreateRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets details about a specific runner.
//
// Use this method to:
//
// - Check runner status
// - View runner configuration
// - Monitor runner health
// - Verify runner capabilities
//
// ### Examples
//
// - Get runner details:
//
//	Retrieves information about a specific runner.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *RunnerService) Get(ctx context.Context, body RunnerGetParams, opts ...option.RequestOption) (res *RunnerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/GetRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates a runner's configuration.
//
// Use this method to:
//
// - Modify runner settings
// - Update release channels
// - Change runner status
// - Configure auto-update settings
//
// ### Examples
//
// - Update configuration:
//
//	Changes runner settings.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	name: "Updated Runner Name"
//	spec:
//	  configuration:
//	    releaseChannel: RUNNER_RELEASE_CHANNEL_LATEST
//	    autoUpdate: true
//	```
func (r *RunnerService) Update(ctx context.Context, body RunnerUpdateParams, opts ...option.RequestOption) (res *RunnerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/UpdateRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all registered runners with optional filtering.
//
// Use this method to:
//
// - View all available runners
// - Filter by runner type
// - Monitor runner status
// - Check runner availability
//
// ### Examples
//
// - List all runners:
//
//	Shows all runners with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter by provider:
//
//	Lists only AWS EC2 runners.
//
//	```yaml
//	filter:
//	  providers: ["RUNNER_PROVIDER_AWS_EC2"]
//	pagination:
//	  pageSize: 20
//	```
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

// Lists all registered runners with optional filtering.
//
// Use this method to:
//
// - View all available runners
// - Filter by runner type
// - Monitor runner status
// - Check runner availability
//
// ### Examples
//
// - List all runners:
//
//	Shows all runners with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter by provider:
//
//	Lists only AWS EC2 runners.
//
//	```yaml
//	filter:
//	  providers: ["RUNNER_PROVIDER_AWS_EC2"]
//	pagination:
//	  pageSize: 20
//	```
func (r *RunnerService) ListAutoPaging(ctx context.Context, params RunnerListParams, opts ...option.RequestOption) *pagination.RunnersPageAutoPager[Runner] {
	return pagination.NewRunnersPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a runner permanently.
//
// Use this method to:
//
// - Remove unused runners
// - Clean up runner registrations
// - Delete obsolete runners
//
// ### Examples
//
// - Delete runner:
//
//	Permanently removes a runner.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *RunnerService) Delete(ctx context.Context, body RunnerDeleteParams, opts ...option.RequestOption) (res *RunnerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/DeleteRunner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Checks if a user is authenticated for a specific host.
//
// Use this method to:
//
// - Verify authentication status
// - Get authentication URLs
// - Check PAT support
//
// ### Examples
//
// - Check authentication:
//
//	Verifies authentication for a host.
//
//	```yaml
//	host: "github.com"
//	```
func (r *RunnerService) CheckAuthenticationForHost(ctx context.Context, body RunnerCheckAuthenticationForHostParams, opts ...option.RequestOption) (res *RunnerCheckAuthenticationForHostResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CheckAuthenticationForHost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a new authentication token for a runner.
//
// Use this method to:
//
// - Generate runner credentials
// - Renew expired tokens
// - Set up runner authentication
//
// Note: This does not expire previously issued tokens.
//
// ### Examples
//
// - Create token:
//
//	Creates a new token for runner authentication.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *RunnerService) NewRunnerToken(ctx context.Context, body RunnerNewRunnerTokenParams, opts ...option.RequestOption) (res *RunnerNewRunnerTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CreateRunnerToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Parses a context URL and returns the parsed result.
//
// Use this method to:
//
// - Validate context URLs
// - Check repository access
// - Verify branch existence
//
// Returns:
//
// - FAILED_PRECONDITION if authentication is required
// - PERMISSION_DENIED if access is not allowed
// - INVALID_ARGUMENT if URL is invalid
// - NOT_FOUND if repository/branch doesn't exist
//
// ### Examples
//
// - Parse URL:
//
//	Parses and validates a context URL.
//
//	```yaml
//	contextUrl: "https://github.com/org/repo/tree/main"
//	```
func (r *RunnerService) ParseContextURL(ctx context.Context, body RunnerParseContextURLParams, opts ...option.RequestOption) (res *RunnerParseContextURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/ParseContextURL"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Runner struct {
	// Time when the Runner was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator is the identity of the creator of the environment
	Creator shared.Subject `json:"creator"`
	// The runner's kind
	Kind RunnerKind `json:"kind"`
	// The runner's name which is shown to users
	Name string `json:"name"`
	// The runner's provider
	Provider RunnerProvider `json:"provider"`
	RunnerID string         `json:"runnerId"`
	// The runner's specification
	Spec RunnerSpec `json:"spec"`
	// The runner's status
	Status RunnerStatus `json:"status"`
	// Time when the Runner was last udpated.
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
	RunnerCapabilitySecretContainerRegistry   RunnerCapability = "RUNNER_CAPABILITY_SECRET_CONTAINER_REGISTRY"
)

func (r RunnerCapability) IsKnown() bool {
	switch r {
	case RunnerCapabilityUnspecified, RunnerCapabilityFetchLocalScmIntegrations, RunnerCapabilitySecretContainerRegistry:
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
	// The runner's reported phase
	Phase RunnerPhase `json:"phase"`
	// region is the region the runner is running in, if applicable.
	Region        string `json:"region"`
	SystemDetails string `json:"systemDetails"`
	// Time when the status was last udpated.
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
	Runner Runner `json:"runner,required"`
	// deprecated, will be removed. Use exchange_token instead.
	//
	// Deprecated: deprecated
	AccessToken string `json:"accessToken"`
	// exchange_token is a one-time use token that should be exchanged by the runner
	// for an access token, using the IdentityService.ExchangeToken rpc. The token
	// expires after 24 hours.
	ExchangeToken string                `json:"exchangeToken"`
	JSON          runnerNewResponseJSON `json:"-"`
}

// runnerNewResponseJSON contains the JSON metadata for the struct
// [RunnerNewResponse]
type runnerNewResponseJSON struct {
	Runner        apijson.Field
	AccessToken   apijson.Field
	ExchangeToken apijson.Field
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
	Runner Runner                `json:"runner,required"`
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
	Authenticated bool `json:"authenticated"`
	// Deprecated: deprecated
	AuthenticationURL string `json:"authenticationUrl"`
	// Deprecated: deprecated
	PatSupported bool `json:"patSupported"`
	// scm_id is the unique identifier of the SCM provider
	ScmID string `json:"scmId"`
	// scm_name is the human-readable name of the SCM provider (e.g., "GitHub",
	// "GitLab")
	ScmName string `json:"scmName"`
	// supports_oauth2 indicates that the host supports OAuth2 authentication
	SupportsOauth2 RunnerCheckAuthenticationForHostResponseSupportsOauth2 `json:"supportsOauth2"`
	// supports_pat indicates that the host supports Personal Access Token
	// authentication
	SupportsPat RunnerCheckAuthenticationForHostResponseSupportsPat `json:"supportsPat"`
	JSON        runnerCheckAuthenticationForHostResponseJSON        `json:"-"`
}

// runnerCheckAuthenticationForHostResponseJSON contains the JSON metadata for the
// struct [RunnerCheckAuthenticationForHostResponse]
type runnerCheckAuthenticationForHostResponseJSON struct {
	Authenticated     apijson.Field
	AuthenticationURL apijson.Field
	PatSupported      apijson.Field
	ScmID             apijson.Field
	ScmName           apijson.Field
	SupportsOauth2    apijson.Field
	SupportsPat       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RunnerCheckAuthenticationForHostResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerCheckAuthenticationForHostResponseJSON) RawJSON() string {
	return r.raw
}

// supports_oauth2 indicates that the host supports OAuth2 authentication
type RunnerCheckAuthenticationForHostResponseSupportsOauth2 struct {
	// auth_url is the URL where users can authenticate
	AuthURL string `json:"authUrl"`
	// docs_url is the URL to the documentation explaining this authentication method
	DocsURL string                                                     `json:"docsUrl"`
	JSON    runnerCheckAuthenticationForHostResponseSupportsOauth2JSON `json:"-"`
}

// runnerCheckAuthenticationForHostResponseSupportsOauth2JSON contains the JSON
// metadata for the struct [RunnerCheckAuthenticationForHostResponseSupportsOauth2]
type runnerCheckAuthenticationForHostResponseSupportsOauth2JSON struct {
	AuthURL     apijson.Field
	DocsURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerCheckAuthenticationForHostResponseSupportsOauth2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerCheckAuthenticationForHostResponseSupportsOauth2JSON) RawJSON() string {
	return r.raw
}

// supports_pat indicates that the host supports Personal Access Token
// authentication
type RunnerCheckAuthenticationForHostResponseSupportsPat struct {
	// create_url is the URL where users can create a new Personal Access Token
	CreateURL string `json:"createUrl"`
	// docs_url is the URL to the documentation explaining PAT usage for this host
	DocsURL string `json:"docsUrl"`
	// example is an example of a Personal Access Token
	Example string `json:"example"`
	// required_scopes is the list of permissions required for the Personal Access
	// Token
	RequiredScopes []string                                                `json:"requiredScopes"`
	JSON           runnerCheckAuthenticationForHostResponseSupportsPatJSON `json:"-"`
}

// runnerCheckAuthenticationForHostResponseSupportsPatJSON contains the JSON
// metadata for the struct [RunnerCheckAuthenticationForHostResponseSupportsPat]
type runnerCheckAuthenticationForHostResponseSupportsPatJSON struct {
	CreateURL      apijson.Field
	DocsURL        apijson.Field
	Example        apijson.Field
	RequiredScopes apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RunnerCheckAuthenticationForHostResponseSupportsPat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerCheckAuthenticationForHostResponseSupportsPatJSON) RawJSON() string {
	return r.raw
}

type RunnerNewRunnerTokenResponse struct {
	// deprecated, will be removed. Use exchange_token instead.
	//
	// Deprecated: deprecated
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
	// The runner's kind This field is optional and here for backwards-compatibility.
	// Use the provider field instead. If provider is set, the runner's kind will be
	// deduced from the provider. Only one of kind and provider must be set.
	Kind param.Field[RunnerKind] `json:"kind"`
	// The runner name for humans
	Name param.Field[string] `json:"name"`
	// The specific implementation type of the runner This field is optional for
	// backwards compatibility but will be required in the future. When specified, kind
	// must not be specified (will be deduced from provider)
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
	// desired_phase can currently only be updated on local-configuration runners, to
	// toggle whether local runners are allowed for running environments in the
	// organization. Set to:
	//
	//   - ACTIVE to enable local runners.
	//   - INACTIVE to disable all local runners. Existing local runners and their
	//     environments will stop, and cannot be started again until the desired_phase is
	//     set to ACTIVE. Use this carefully, as it will affect all users in the
	//     organization who use local runners.
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
