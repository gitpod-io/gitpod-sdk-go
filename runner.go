// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"slices"
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerService/CheckAuthenticationForHost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates an access token for runner logs and debug information.
//
// Generated tokens are valid for one hour and provide runner-specific access
// permissions. The token is scoped to a specific runner and can be used to access
// support bundles.
//
// ### Examples
//
// - Generate runner logs token:
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *RunnerService) NewLogsToken(ctx context.Context, body RunnerNewLogsTokenParams, opts ...option.RequestOption) (res *RunnerNewLogsTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerService/CreateRunnerLogsToken"
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerService/ParseContextURL"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Searches for repositories across all authenticated SCM hosts.
//
// Use this method to:
//
// - List available repositories
// - Search repositories by name or content
// - Discover repositories for environment creation
//
// Returns repositories from all authenticated SCM hosts in natural sort order. If
// no repositories are found, returns an empty list.
//
// ### Examples
//
// - List all repositories:
//
//	Returns up to 25 repositories from all authenticated hosts.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
//
// - Search repositories:
//
//	Searches for repositories matching the query across all hosts.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	searchString: "my-project"
//	limit: 10
//	```
func (r *RunnerService) SearchRepositories(ctx context.Context, body RunnerSearchRepositoriesParams, opts ...option.RequestOption) (res *RunnerSearchRepositoriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerService/SearchRepositories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GatewayInfo struct {
	// Gateway represents a system gateway that provides access to services
	Gateway shared.Gateway `json:"gateway"`
	// latency is the round-trip time of the runner to the gateway in milliseconds.
	Latency string          `json:"latency" format:"regex"`
	JSON    gatewayInfoJSON `json:"-"`
}

// gatewayInfoJSON contains the JSON metadata for the struct [GatewayInfo]
type gatewayInfoJSON struct {
	Gateway     apijson.Field
	Latency     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayInfoJSON) RawJSON() string {
	return r.raw
}

type LogLevel string

const (
	LogLevelUnspecified LogLevel = "LOG_LEVEL_UNSPECIFIED"
	LogLevelDebug       LogLevel = "LOG_LEVEL_DEBUG"
	LogLevelInfo        LogLevel = "LOG_LEVEL_INFO"
	LogLevelWarn        LogLevel = "LOG_LEVEL_WARN"
	LogLevelError       LogLevel = "LOG_LEVEL_ERROR"
)

func (r LogLevel) IsKnown() bool {
	switch r {
	case LogLevelUnspecified, LogLevelDebug, LogLevelInfo, LogLevelWarn, LogLevelError:
		return true
	}
	return false
}

type MetricsConfiguration struct {
	// enabled indicates whether the runner should collect metrics
	Enabled bool `json:"enabled"`
	// password is the password to use for the metrics collector
	Password string `json:"password"`
	// url is the URL of the metrics collector
	URL string `json:"url"`
	// username is the username to use for the metrics collector
	Username string                   `json:"username"`
	JSON     metricsConfigurationJSON `json:"-"`
}

// metricsConfigurationJSON contains the JSON metadata for the struct
// [MetricsConfiguration]
type metricsConfigurationJSON struct {
	Enabled     apijson.Field
	Password    apijson.Field
	URL         apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricsConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricsConfigurationJSON) RawJSON() string {
	return r.raw
}

type MetricsConfigurationParam struct {
	// enabled indicates whether the runner should collect metrics
	Enabled param.Field[bool] `json:"enabled"`
	// password is the password to use for the metrics collector
	Password param.Field[string] `json:"password"`
	// url is the URL of the metrics collector
	URL param.Field[string] `json:"url"`
	// username is the username to use for the metrics collector
	Username param.Field[string] `json:"username"`
}

func (r MetricsConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// The runner manager id specifies the runner manager for the managed runner. This
	// field is only set for managed runners.
	RunnerManagerID string `json:"runnerManagerId" format:"uuid"`
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
	CreatedAt       apijson.Field
	Creator         apijson.Field
	Kind            apijson.Field
	Name            apijson.Field
	Provider        apijson.Field
	RunnerID        apijson.Field
	RunnerManagerID apijson.Field
	Spec            apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
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
	RunnerCapabilityAgentExecution            RunnerCapability = "RUNNER_CAPABILITY_AGENT_EXECUTION"
	RunnerCapabilityAllowEnvTokenPopulation   RunnerCapability = "RUNNER_CAPABILITY_ALLOW_ENV_TOKEN_POPULATION"
	RunnerCapabilityDefaultDevContainerImage  RunnerCapability = "RUNNER_CAPABILITY_DEFAULT_DEV_CONTAINER_IMAGE"
	RunnerCapabilityEnvironmentSnapshot       RunnerCapability = "RUNNER_CAPABILITY_ENVIRONMENT_SNAPSHOT"
)

func (r RunnerCapability) IsKnown() bool {
	switch r {
	case RunnerCapabilityUnspecified, RunnerCapabilityFetchLocalScmIntegrations, RunnerCapabilitySecretContainerRegistry, RunnerCapabilityAgentExecution, RunnerCapabilityAllowEnvTokenPopulation, RunnerCapabilityDefaultDevContainerImage, RunnerCapabilityEnvironmentSnapshot:
		return true
	}
	return false
}

type RunnerConfiguration struct {
	// auto_update indicates whether the runner should automatically update itself.
	AutoUpdate bool `json:"autoUpdate"`
	// devcontainer_image_cache_enabled controls whether the devcontainer build cache
	// is enabled for this runner. Only takes effect on supported runners, currently
	// only AWS EC2 and Gitpod-managed runners.
	DevcontainerImageCacheEnabled bool `json:"devcontainerImageCacheEnabled"`
	// log_level is the log level for the runner
	LogLevel LogLevel `json:"logLevel"`
	// metrics contains configuration for the runner's metrics collection
	Metrics MetricsConfiguration `json:"metrics"`
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
	AutoUpdate                    apijson.Field
	DevcontainerImageCacheEnabled apijson.Field
	LogLevel                      apijson.Field
	Metrics                       apijson.Field
	Region                        apijson.Field
	ReleaseChannel                apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
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
	// devcontainer_image_cache_enabled controls whether the devcontainer build cache
	// is enabled for this runner. Only takes effect on supported runners, currently
	// only AWS EC2 and Gitpod-managed runners.
	DevcontainerImageCacheEnabled param.Field[bool] `json:"devcontainerImageCacheEnabled"`
	// log_level is the log level for the runner
	LogLevel param.Field[LogLevel] `json:"logLevel"`
	// metrics contains configuration for the runner's metrics collection
	Metrics param.Field[MetricsConfigurationParam] `json:"metrics"`
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
	RunnerProviderManaged     RunnerProvider = "RUNNER_PROVIDER_MANAGED"
	RunnerProviderGcp         RunnerProvider = "RUNNER_PROVIDER_GCP"
)

func (r RunnerProvider) IsKnown() bool {
	switch r {
	case RunnerProviderUnspecified, RunnerProviderAwsEc2, RunnerProviderLinuxHost, RunnerProviderDesktopMac, RunnerProviderManaged, RunnerProviderGcp:
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
	DesiredPhase RunnerPhase `json:"desiredPhase"`
	// The runner's variant
	Variant RunnerVariant  `json:"variant"`
	JSON    runnerSpecJSON `json:"-"`
}

// runnerSpecJSON contains the JSON metadata for the struct [RunnerSpec]
type runnerSpecJSON struct {
	Configuration apijson.Field
	DesiredPhase  apijson.Field
	Variant       apijson.Field
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
	// The runner's variant
	Variant param.Field[RunnerVariant] `json:"variant"`
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
	// gateway_info is information about the gateway to which the runner is connected.
	GatewayInfo GatewayInfo `json:"gatewayInfo"`
	// llm_url is the URL of the LLM service to which the runner is connected.
	LlmURL string `json:"llmUrl"`
	LogURL string `json:"logUrl"`
	// The runner's reported message which is shown to users. This message adds more
	// context to the runner's phase.
	Message string `json:"message"`
	// The runner's reported phase
	Phase RunnerPhase `json:"phase"`
	// public_key is the runner's public key used for encryption (32 bytes)
	PublicKey string `json:"publicKey" format:"byte"`
	// region is the region the runner is running in, if applicable.
	Region string `json:"region"`
	// support_bundle_url is the URL at which the runner support bundle can be
	// accessed. This URL provides access to pprof profiles and other debug
	// information. Only available for standalone runners.
	SupportBundleURL string `json:"supportBundleUrl"`
	SystemDetails    string `json:"systemDetails"`
	// Time when the status was last updated.
	UpdatedAt time.Time        `json:"updatedAt" format:"date-time"`
	Version   string           `json:"version"`
	JSON      runnerStatusJSON `json:"-"`
}

// runnerStatusJSON contains the JSON metadata for the struct [RunnerStatus]
type runnerStatusJSON struct {
	AdditionalInfo   apijson.Field
	Capabilities     apijson.Field
	GatewayInfo      apijson.Field
	LlmURL           apijson.Field
	LogURL           apijson.Field
	Message          apijson.Field
	Phase            apijson.Field
	PublicKey        apijson.Field
	Region           apijson.Field
	SupportBundleURL apijson.Field
	SystemDetails    apijson.Field
	UpdatedAt        apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RunnerStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerStatusJSON) RawJSON() string {
	return r.raw
}

type RunnerVariant string

const (
	RunnerVariantUnspecified RunnerVariant = "RUNNER_VARIANT_UNSPECIFIED"
	RunnerVariantStandard    RunnerVariant = "RUNNER_VARIANT_STANDARD"
	RunnerVariantEnterprise  RunnerVariant = "RUNNER_VARIANT_ENTERPRISE"
)

func (r RunnerVariant) IsKnown() bool {
	switch r {
	case RunnerVariantUnspecified, RunnerVariantStandard, RunnerVariantEnterprise:
		return true
	}
	return false
}

type SearchMode string

const (
	SearchModeUnspecified SearchMode = "SEARCH_MODE_UNSPECIFIED"
	SearchModeKeyword     SearchMode = "SEARCH_MODE_KEYWORD"
	SearchModeNative      SearchMode = "SEARCH_MODE_NATIVE"
)

func (r SearchMode) IsKnown() bool {
	switch r {
	case SearchModeUnspecified, SearchModeKeyword, SearchModeNative:
		return true
	}
	return false
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

type RunnerNewLogsTokenResponse struct {
	// access_token is the token that can be used to access the logs and support bundle
	// of the runner
	AccessToken string                         `json:"accessToken,required"`
	JSON        runnerNewLogsTokenResponseJSON `json:"-"`
}

// runnerNewLogsTokenResponseJSON contains the JSON metadata for the struct
// [RunnerNewLogsTokenResponse]
type runnerNewLogsTokenResponseJSON struct {
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerNewLogsTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerNewLogsTokenResponseJSON) RawJSON() string {
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
	Git                RunnerParseContextURLResponseGit   `json:"git"`
	Issue              RunnerParseContextURLResponseIssue `json:"issue"`
	OriginalContextURL string                             `json:"originalContextUrl"`
	// Deprecated: Use top-level PullRequest message instead
	//
	// Deprecated: deprecated
	Pr RunnerParseContextURLResponsePr `json:"pr"`
	// project_ids is a list of projects to which the context URL belongs to.
	ProjectIDs []string `json:"projectIds"`
	// PullRequest represents pull request metadata from source control systems. This
	// message is used across workflow triggers, executions, and agent contexts to
	// maintain consistent PR information throughout the system.
	PullRequest RunnerParseContextURLResponsePullRequest `json:"pullRequest"`
	// scm_id is the unique identifier of the SCM provider (e.g., "github", "gitlab",
	// "bitbucket")
	ScmID string                            `json:"scmId"`
	JSON  runnerParseContextURLResponseJSON `json:"-"`
}

// runnerParseContextURLResponseJSON contains the JSON metadata for the struct
// [RunnerParseContextURLResponse]
type runnerParseContextURLResponseJSON struct {
	Git                apijson.Field
	Issue              apijson.Field
	OriginalContextURL apijson.Field
	Pr                 apijson.Field
	ProjectIDs         apijson.Field
	PullRequest        apijson.Field
	ScmID              apijson.Field
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
	Tag               string                               `json:"tag"`
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
	Tag               apijson.Field
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

type RunnerParseContextURLResponseIssue struct {
	// id is the source system's ID of this issue, e.g. BNFRD-6100
	ID    string                                 `json:"id"`
	Title string                                 `json:"title"`
	JSON  runnerParseContextURLResponseIssueJSON `json:"-"`
}

// runnerParseContextURLResponseIssueJSON contains the JSON metadata for the struct
// [RunnerParseContextURLResponseIssue]
type runnerParseContextURLResponseIssueJSON struct {
	ID          apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerParseContextURLResponseIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerParseContextURLResponseIssueJSON) RawJSON() string {
	return r.raw
}

// Deprecated: Use top-level PullRequest message instead
//
// Deprecated: deprecated
type RunnerParseContextURLResponsePr struct {
	ID         string                              `json:"id"`
	FromBranch string                              `json:"fromBranch"`
	Title      string                              `json:"title"`
	ToBranch   string                              `json:"toBranch"`
	JSON       runnerParseContextURLResponsePrJSON `json:"-"`
}

// runnerParseContextURLResponsePrJSON contains the JSON metadata for the struct
// [RunnerParseContextURLResponsePr]
type runnerParseContextURLResponsePrJSON struct {
	ID          apijson.Field
	FromBranch  apijson.Field
	Title       apijson.Field
	ToBranch    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerParseContextURLResponsePr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerParseContextURLResponsePrJSON) RawJSON() string {
	return r.raw
}

// PullRequest represents pull request metadata from source control systems. This
// message is used across workflow triggers, executions, and agent contexts to
// maintain consistent PR information throughout the system.
type RunnerParseContextURLResponsePullRequest struct {
	// Unique identifier from the source system (e.g., "123" for GitHub PR #123)
	ID string `json:"id"`
	// Author name as provided by the SCM system
	Author string `json:"author"`
	// Source branch name (the branch being merged from)
	FromBranch string `json:"fromBranch"`
	// Repository information
	Repository RunnerParseContextURLResponsePullRequestRepository `json:"repository"`
	// Pull request title
	Title string `json:"title"`
	// Target branch name (the branch being merged into)
	ToBranch string `json:"toBranch"`
	// Pull request URL (e.g., "https://github.com/owner/repo/pull/123")
	URL  string                                       `json:"url"`
	JSON runnerParseContextURLResponsePullRequestJSON `json:"-"`
}

// runnerParseContextURLResponsePullRequestJSON contains the JSON metadata for the
// struct [RunnerParseContextURLResponsePullRequest]
type runnerParseContextURLResponsePullRequestJSON struct {
	ID          apijson.Field
	Author      apijson.Field
	FromBranch  apijson.Field
	Repository  apijson.Field
	Title       apijson.Field
	ToBranch    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerParseContextURLResponsePullRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerParseContextURLResponsePullRequestJSON) RawJSON() string {
	return r.raw
}

// Repository information
type RunnerParseContextURLResponsePullRequestRepository struct {
	CloneURL string                                                 `json:"cloneUrl"`
	Host     string                                                 `json:"host"`
	Name     string                                                 `json:"name"`
	Owner    string                                                 `json:"owner"`
	JSON     runnerParseContextURLResponsePullRequestRepositoryJSON `json:"-"`
}

// runnerParseContextURLResponsePullRequestRepositoryJSON contains the JSON
// metadata for the struct [RunnerParseContextURLResponsePullRequestRepository]
type runnerParseContextURLResponsePullRequestRepositoryJSON struct {
	CloneURL    apijson.Field
	Host        apijson.Field
	Name        apijson.Field
	Owner       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerParseContextURLResponsePullRequestRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerParseContextURLResponsePullRequestRepositoryJSON) RawJSON() string {
	return r.raw
}

type RunnerSearchRepositoriesResponse struct {
	// Deprecated: Use pagination token instead. Total pages can be extracted from
	// token.
	LastPage int64 `json:"lastPage"`
	// Pagination information for the response. Token format:
	// "NEXT_PAGE/TOTAL_PAGES/TOTAL_COUNT" (e.g., "2/40/1000"). Use -1 for unknown
	// values (e.g., "2/-1/-1" when totals unavailable). Empty token means no more
	// pages.
	Pagination RunnerSearchRepositoriesResponsePagination `json:"pagination"`
	// List of repositories matching the search criteria
	Repositories []RunnerSearchRepositoriesResponseRepository `json:"repositories"`
	JSON         runnerSearchRepositoriesResponseJSON         `json:"-"`
}

// runnerSearchRepositoriesResponseJSON contains the JSON metadata for the struct
// [RunnerSearchRepositoriesResponse]
type runnerSearchRepositoriesResponseJSON struct {
	LastPage     apijson.Field
	Pagination   apijson.Field
	Repositories apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RunnerSearchRepositoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerSearchRepositoriesResponseJSON) RawJSON() string {
	return r.raw
}

// Pagination information for the response. Token format:
// "NEXT_PAGE/TOTAL_PAGES/TOTAL_COUNT" (e.g., "2/40/1000"). Use -1 for unknown
// values (e.g., "2/-1/-1" when totals unavailable). Empty token means no more
// pages.
type RunnerSearchRepositoriesResponsePagination struct {
	// Token passed for retrieving the next set of results. Empty if there are no more
	// results
	NextToken string                                         `json:"nextToken"`
	JSON      runnerSearchRepositoriesResponsePaginationJSON `json:"-"`
}

// runnerSearchRepositoriesResponsePaginationJSON contains the JSON metadata for
// the struct [RunnerSearchRepositoriesResponsePagination]
type runnerSearchRepositoriesResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerSearchRepositoriesResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerSearchRepositoriesResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerSearchRepositoriesResponseRepository struct {
	// Repository name (e.g., "my-project")
	Name string `json:"name"`
	// Repository URL (e.g., "https://github.com/owner/my-project")
	URL  string                                         `json:"url"`
	JSON runnerSearchRepositoriesResponseRepositoryJSON `json:"-"`
}

// runnerSearchRepositoriesResponseRepositoryJSON contains the JSON metadata for
// the struct [RunnerSearchRepositoriesResponseRepository]
type runnerSearchRepositoriesResponseRepositoryJSON struct {
	Name        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerSearchRepositoriesResponseRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerSearchRepositoriesResponseRepositoryJSON) RawJSON() string {
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
	Provider param.Field[RunnerProvider] `json:"provider"`
	// The runner manager id specifies the runner manager for the managed runner. This
	// field is mandatory for managed runners, otheriwse should not be set.
	RunnerManagerID param.Field[string]          `json:"runnerManagerId" format:"uuid"`
	Spec            param.Field[RunnerSpecParam] `json:"spec"`
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
	// devcontainer_image_cache_enabled controls whether the shared devcontainer build
	// cache is enabled for this runner.
	DevcontainerImageCacheEnabled param.Field[bool] `json:"devcontainerImageCacheEnabled"`
	// log_level is the log level for the runner
	LogLevel param.Field[LogLevel] `json:"logLevel"`
	// metrics contains configuration for the runner's metrics collection
	Metrics param.Field[RunnerUpdateParamsSpecConfigurationMetrics] `json:"metrics"`
	// The release channel the runner is on
	ReleaseChannel param.Field[RunnerReleaseChannel] `json:"releaseChannel"`
}

func (r RunnerUpdateParamsSpecConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// metrics contains configuration for the runner's metrics collection
type RunnerUpdateParamsSpecConfigurationMetrics struct {
	// enabled indicates whether the runner should collect metrics
	Enabled param.Field[bool] `json:"enabled"`
	// password is the password to use for the metrics collector
	Password param.Field[string] `json:"password"`
	// url is the URL of the metrics collector
	URL param.Field[string] `json:"url"`
	// username is the username to use for the metrics collector
	Username param.Field[string] `json:"username"`
}

func (r RunnerUpdateParamsSpecConfigurationMetrics) MarshalJSON() (data []byte, err error) {
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

type RunnerNewLogsTokenParams struct {
	// runner_id specifies the runner for which the logs token should be created.
	//
	// +required
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerNewLogsTokenParams) MarshalJSON() (data []byte, err error) {
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

type RunnerSearchRepositoriesParams struct {
	// Maximum number of repositories to return. Default: 25, Maximum: 100 Deprecated:
	// Use pagination.page_size instead
	Limit param.Field[int64] `json:"limit"`
	// Pagination parameters for repository search
	Pagination param.Field[RunnerSearchRepositoriesParamsPagination] `json:"pagination"`
	RunnerID   param.Field[string]                                   `json:"runnerId" format:"uuid"`
	// The SCM's host to retrieve repositories from
	ScmHost param.Field[string] `json:"scmHost"`
	// Search mode determines how search_string is interpreted
	SearchMode param.Field[SearchMode] `json:"searchMode"`
	// Search query - interpretation depends on search_mode
	SearchString param.Field[string] `json:"searchString"`
}

func (r RunnerSearchRepositoriesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Pagination parameters for repository search
type RunnerSearchRepositoriesParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerSearchRepositoriesParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
