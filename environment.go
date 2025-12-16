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

// Creates a development environment from a context URL (e.g. Git repository) and
// starts it.
//
// The `class` field must be a valid environment class ID. You can find a list of
// available environment classes with the `ListEnvironmentClasses` method.
//
// ### Examples
//
// - Create from context URL:
//
//	Creates an environment from a Git repository URL with default settings.
//
//	```yaml
//	spec:
//	  machine:
//	    class: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	  content:
//	    initializer:
//	      specs:
//	        - contextUrl:
//	            url: "https://github.com/gitpod-io/gitpod"
//	```
//
// - Create from Git repository:
//
//	Creates an environment from a Git repository with specific branch targeting.
//
//	```yaml
//	spec:
//	  machine:
//	    class: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	  content:
//	    initializer:
//	      specs:
//	        - git:
//	            remoteUri: "https://github.com/gitpod-io/gitpod"
//	            cloneTarget: "main"
//	            targetMode: "CLONE_TARGET_MODE_REMOTE_BRANCH"
//	```
//
// - Create with custom timeout and ports:
//
//	Creates an environment with custom inactivity timeout and exposed port
//	configuration.
//
//	```yaml
//	spec:
//	  machine:
//	    class: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	  content:
//	    initializer:
//	      specs:
//	        - contextUrl:
//	            url: "https://github.com/gitpod-io/gitpod"
//	  timeout:
//	    disconnected: "7200s" # 2 hours in seconds
//	  ports:
//	    - port: 3000
//	      admission: "ADMISSION_LEVEL_EVERYONE"
//	      name: "Web App"
//	```
func (r *EnvironmentService) New(ctx context.Context, body EnvironmentNewParams, opts ...option.RequestOption) (res *EnvironmentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentService/CreateEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets details about a specific environment including its status, configuration,
// and context URL.
//
// Use this method to:
//
// - Check if an environment is ready to use
// - Get connection details for IDE and exposed ports
// - Monitor environment health and resource usage
// - Debug environment setup issues
//
// ### Examples
//
// - Get environment details:
//
//	Retrieves detailed information about a specific environment using its unique
//	identifier.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *EnvironmentService) Get(ctx context.Context, body EnvironmentGetParams, opts ...option.RequestOption) (res *EnvironmentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentService/GetEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an environment's configuration while it is running.
//
// Updates are limited to:
//
// - Git credentials (username, email)
// - SSH public keys
// - Content initialization
// - Port configurations
// - Automation files
// - Environment timeouts
//
// ### Examples
//
// - Update Git credentials:
//
//	Updates the Git configuration for the environment.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	spec:
//	  content:
//	    gitUsername: "example-user"
//	    gitEmail: "user@example.com"
//	```
//
// - Add SSH public key:
//
//	Adds a new SSH public key for authentication.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	spec:
//	  sshPublicKeys:
//	    - id: "0194b7c1-c954-718d-91a4-9a742aa5fc11"
//	      value: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI..."
//	```
//
// - Update content session:
//
//	Updates the content session identifier for the environment.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	spec:
//	  content:
//	    session: "0194b7c1-c954-718d-91a4-9a742aa5fc11"
//	```
//
// Note: Machine class changes require stopping the environment and creating a new
// one.
func (r *EnvironmentService) Update(ctx context.Context, body EnvironmentUpdateParams, opts ...option.RequestOption) (res *EnvironmentUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentService/UpdateEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all environments matching the specified criteria.
//
// Use this method to find and monitor environments across your organization.
// Results are ordered by creation time with newest environments first.
//
// ### Examples
//
// - List running environments for a project:
//
//	Retrieves all running environments for a specific project with pagination.
//
//	```yaml
//	filter:
//	  statusPhases: ["ENVIRONMENT_PHASE_RUNNING"]
//	  projectIds: ["b0e12f6c-4c67-429d-a4a6-d9838b5da047"]
//	pagination:
//	  pageSize: 10
//	```
//
// - List all environments for a specific runner:
//
//	Filters environments by runner ID and creator ID.
//
//	```yaml
//	filter:
//	  runnerIds: ["e6aa9c54-89d3-42c1-ac31-bd8d8f1concentrate"]
//	  creatorIds: ["f53d2330-3795-4c5d-a1f3-453121af9c60"]
//	```
//
// - List stopped and deleted environments:
//
//	Retrieves all environments in stopped or deleted state.
//
//	```yaml
//	filter:
//	  statusPhases: ["ENVIRONMENT_PHASE_STOPPED", "ENVIRONMENT_PHASE_DELETED"]
//	```
func (r *EnvironmentService) List(ctx context.Context, params EnvironmentListParams, opts ...option.RequestOption) (res *pagination.EnvironmentsPage[Environment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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

// Lists all environments matching the specified criteria.
//
// Use this method to find and monitor environments across your organization.
// Results are ordered by creation time with newest environments first.
//
// ### Examples
//
// - List running environments for a project:
//
//	Retrieves all running environments for a specific project with pagination.
//
//	```yaml
//	filter:
//	  statusPhases: ["ENVIRONMENT_PHASE_RUNNING"]
//	  projectIds: ["b0e12f6c-4c67-429d-a4a6-d9838b5da047"]
//	pagination:
//	  pageSize: 10
//	```
//
// - List all environments for a specific runner:
//
//	Filters environments by runner ID and creator ID.
//
//	```yaml
//	filter:
//	  runnerIds: ["e6aa9c54-89d3-42c1-ac31-bd8d8f1concentrate"]
//	  creatorIds: ["f53d2330-3795-4c5d-a1f3-453121af9c60"]
//	```
//
// - List stopped and deleted environments:
//
//	Retrieves all environments in stopped or deleted state.
//
//	```yaml
//	filter:
//	  statusPhases: ["ENVIRONMENT_PHASE_STOPPED", "ENVIRONMENT_PHASE_DELETED"]
//	```
func (r *EnvironmentService) ListAutoPaging(ctx context.Context, params EnvironmentListParams, opts ...option.RequestOption) *pagination.EnvironmentsPageAutoPager[Environment] {
	return pagination.NewEnvironmentsPageAutoPager(r.List(ctx, params, opts...))
}

// Permanently deletes an environment.
//
// Running environments are automatically stopped before deletion. If force is
// true, the environment is deleted immediately without graceful shutdown.
//
// ### Examples
//
// - Delete with graceful shutdown:
//
//	Deletes an environment after gracefully stopping it.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	force: false
//	```
//
// - Force delete:
//
//	Immediately deletes an environment without waiting for graceful shutdown.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	force: true
//	```
func (r *EnvironmentService) Delete(ctx context.Context, body EnvironmentDeleteParams, opts ...option.RequestOption) (res *EnvironmentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentService/DeleteEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates an access token for the environment.
//
// Generated tokens are valid for one hour and provide environment-specific access
// permissions. The token is scoped to a specific environment.
//
// ### Examples
//
// - Generate environment token:
//
//	Creates a temporary access token for accessing an environment.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *EnvironmentService) NewEnvironmentToken(ctx context.Context, body EnvironmentNewEnvironmentTokenParams, opts ...option.RequestOption) (res *EnvironmentNewEnvironmentTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentService/CreateEnvironmentAccessToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates an environment from an existing project configuration and starts it.
//
// This method uses project settings as defaults but allows overriding specific
// configurations. Project settings take precedence over default configurations,
// while custom specifications in the request override project settings.
//
// ### Examples
//
// - Create with project defaults:
//
//	Creates an environment using all default settings from the project
//	configuration.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
//
// - Create with custom compute resources:
//
//	Creates an environment from project with custom machine class and timeout
//	settings.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	spec:
//	  machine:
//	    class: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	  timeout:
//	    disconnected: "14400s" # 4 hours in seconds
//	```
func (r *EnvironmentService) NewFromProject(ctx context.Context, body EnvironmentNewFromProjectParams, opts ...option.RequestOption) (res *EnvironmentNewFromProjectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentService/CreateEnvironmentFromProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates an access token for retrieving environment logs.
//
// Generated tokens are valid for one hour and provide read-only access to the
// environment's logs.
//
// ### Examples
//
// - Generate logs token:
//
//	Creates a temporary access token for retrieving environment logs.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *EnvironmentService) NewLogsToken(ctx context.Context, body EnvironmentNewLogsTokenParams, opts ...option.RequestOption) (res *EnvironmentNewLogsTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentService/CreateEnvironmentLogsToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Records environment activity to prevent automatic shutdown.
//
// Activity signals should be sent every 5 minutes while the environment is
// actively being used. The source must be between 3-80 characters.
//
// ### Examples
//
// - Signal VS Code activity:
//
//	Records VS Code editor activity to prevent environment shutdown.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	activitySignal:
//	  source: "VS Code"
//	  timestamp: "2025-02-12T14:30:00Z"
//	```
func (r *EnvironmentService) MarkActive(ctx context.Context, body EnvironmentMarkActiveParams, opts ...option.RequestOption) (res *EnvironmentMarkActiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentService/MarkEnvironmentActive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Starts a stopped environment.
//
// Use this method to resume work on a previously stopped environment. The
// environment retains its configuration and workspace content from when it was
// stopped.
//
// ### Examples
//
// - Start an environment:
//
//	Resumes a previously stopped environment with its existing configuration.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *EnvironmentService) Start(ctx context.Context, body EnvironmentStartParams, opts ...option.RequestOption) (res *EnvironmentStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentService/StartEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stops a running environment.
//
// Use this method to pause work while preserving the environment's state. The
// environment can be resumed later using StartEnvironment.
//
// ### Examples
//
// - Stop an environment:
//
//	Gracefully stops a running environment while preserving its state.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *EnvironmentService) Stop(ctx context.Context, body EnvironmentStopParams, opts ...option.RequestOption) (res *EnvironmentStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentService/StopEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unarchives an environment.
//
// ### Examples
//
// - Unarchive an environment:
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *EnvironmentService) Unarchive(ctx context.Context, body EnvironmentUnarchiveParams, opts ...option.RequestOption) (res *EnvironmentUnarchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentService/UnarchiveEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Admission level describes who can access an environment instance and its ports.
type AdmissionLevel string

const (
	AdmissionLevelUnspecified AdmissionLevel = "ADMISSION_LEVEL_UNSPECIFIED"
	AdmissionLevelOwnerOnly   AdmissionLevel = "ADMISSION_LEVEL_OWNER_ONLY"
	AdmissionLevelEveryone    AdmissionLevel = "ADMISSION_LEVEL_EVERYONE"
)

func (r AdmissionLevel) IsKnown() bool {
	switch r {
	case AdmissionLevelUnspecified, AdmissionLevelOwnerOnly, AdmissionLevelEveryone:
		return true
	}
	return false
}

// +resource get environment
type Environment struct {
	// ID is a unique identifier of this environment. No other environment with the
	// same name must be managed by this environment manager
	ID string `json:"id,required"`
	// Metadata is data associated with this environment that's required for other
	// parts of Gitpod to function
	Metadata EnvironmentMetadata `json:"metadata"`
	// Spec is the configuration of the environment that's required for the runner to
	// start the environment
	Spec EnvironmentSpec `json:"spec"`
	// Status is the current status of the environment
	Status EnvironmentStatus `json:"status"`
	JSON   environmentJSON   `json:"-"`
}

// environmentJSON contains the JSON metadata for the struct [Environment]
type environmentJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Environment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentActivitySignal used to signal activity for an environment.
type EnvironmentActivitySignal struct {
	// source of the activity signal, such as "VS Code", "SSH", or "Automations". It
	// should be a human-readable string that describes the source of the activity
	// signal.
	Source string `json:"source"`
	// timestamp of when the activity was observed by the source. Only reported every 5
	// minutes. Zero value means no activity was observed.
	Timestamp time.Time                     `json:"timestamp" format:"date-time"`
	JSON      environmentActivitySignalJSON `json:"-"`
}

// environmentActivitySignalJSON contains the JSON metadata for the struct
// [EnvironmentActivitySignal]
type environmentActivitySignalJSON struct {
	Source      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentActivitySignal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentActivitySignalJSON) RawJSON() string {
	return r.raw
}

// EnvironmentActivitySignal used to signal activity for an environment.
type EnvironmentActivitySignalParam struct {
	// source of the activity signal, such as "VS Code", "SSH", or "Automations". It
	// should be a human-readable string that describes the source of the activity
	// signal.
	Source param.Field[string] `json:"source"`
	// timestamp of when the activity was observed by the source. Only reported every 5
	// minutes. Zero value means no activity was observed.
	Timestamp param.Field[time.Time] `json:"timestamp" format:"date-time"`
}

func (r EnvironmentActivitySignalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnvironmentMetadata is data associated with an environment that's required for
// other parts of the system to function
type EnvironmentMetadata struct {
	// annotations are key/value pairs that gets attached to the environment.
	// +internal - not yet implemented
	Annotations map[string]string `json:"annotations"`
	// Time when the Environment was archived. If not set, the environment is not
	// archived.
	ArchivedAt time.Time `json:"archivedAt" format:"date-time"`
	// Time when the Environment was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator is the identity of the creator of the environment
	Creator shared.Subject `json:"creator"`
	// Time when the Environment was last started (i.e. CreateEnvironment or
	// StartEnvironment were called).
	LastStartedAt time.Time `json:"lastStartedAt" format:"date-time"`
	// name is the name of the environment as specified by the user
	Name string `json:"name"`
	// organization_id is the ID of the organization that contains the environment
	OrganizationID string `json:"organizationId" format:"uuid"`
	// original_context_url is the normalized URL from which the environment was
	// created
	OriginalContextURL string `json:"originalContextUrl"`
	// prebuild_id is the ID of the prebuild this environment was created from. Only
	// set if the environment was created from a prebuild.
	PrebuildID string `json:"prebuildId,nullable" format:"uuid"`
	// If the Environment was started from a project, the project_id will reference the
	// project.
	ProjectID string `json:"projectId"`
	// role is the role of the environment
	Role EnvironmentRole `json:"role"`
	// Runner is the ID of the runner that runs this environment.
	RunnerID string                  `json:"runnerId"`
	JSON     environmentMetadataJSON `json:"-"`
}

// environmentMetadataJSON contains the JSON metadata for the struct
// [EnvironmentMetadata]
type environmentMetadataJSON struct {
	Annotations        apijson.Field
	ArchivedAt         apijson.Field
	CreatedAt          apijson.Field
	Creator            apijson.Field
	LastStartedAt      apijson.Field
	Name               apijson.Field
	OrganizationID     apijson.Field
	OriginalContextURL apijson.Field
	PrebuildID         apijson.Field
	ProjectID          apijson.Field
	Role               apijson.Field
	RunnerID           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentMetadataJSON) RawJSON() string {
	return r.raw
}

type EnvironmentPhase string

const (
	EnvironmentPhaseUnspecified EnvironmentPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentPhaseCreating    EnvironmentPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentPhaseStarting    EnvironmentPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentPhaseRunning     EnvironmentPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentPhaseUpdating    EnvironmentPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentPhaseStopping    EnvironmentPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentPhaseStopped     EnvironmentPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentPhaseDeleting    EnvironmentPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentPhaseDeleted     EnvironmentPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentPhase) IsKnown() bool {
	switch r {
	case EnvironmentPhaseUnspecified, EnvironmentPhaseCreating, EnvironmentPhaseStarting, EnvironmentPhaseRunning, EnvironmentPhaseUpdating, EnvironmentPhaseStopping, EnvironmentPhaseStopped, EnvironmentPhaseDeleting, EnvironmentPhaseDeleted:
		return true
	}
	return false
}

// EnvironmentRole represents the role of an environment
type EnvironmentRole string

const (
	EnvironmentRoleUnspecified EnvironmentRole = "ENVIRONMENT_ROLE_UNSPECIFIED"
	EnvironmentRoleDefault     EnvironmentRole = "ENVIRONMENT_ROLE_DEFAULT"
	EnvironmentRolePrebuild    EnvironmentRole = "ENVIRONMENT_ROLE_PREBUILD"
	EnvironmentRoleWorkflow    EnvironmentRole = "ENVIRONMENT_ROLE_WORKFLOW"
)

func (r EnvironmentRole) IsKnown() bool {
	switch r {
	case EnvironmentRoleUnspecified, EnvironmentRoleDefault, EnvironmentRolePrebuild, EnvironmentRoleWorkflow:
		return true
	}
	return false
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type EnvironmentSpec struct {
	// admission controlls who can access the environment and its ports.
	Admission AdmissionLevel `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile EnvironmentSpecAutomationsFile `json:"automationsFile"`
	// content is the content spec of the environment
	Content EnvironmentSpecContent `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase EnvironmentPhase `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer EnvironmentSpecDevcontainer `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine EnvironmentSpecMachine `json:"machine"`
	// ports is the set of ports which ought to be exposed to your network
	Ports []EnvironmentSpecPort `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets []EnvironmentSpecSecret `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion string `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys []EnvironmentSpecSSHPublicKey `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout EnvironmentSpecTimeout `json:"timeout"`
	// workflow_action_id is an optional reference to the workflow execution action
	// that created this environment. Used for tracking and event correlation.
	WorkflowActionID string              `json:"workflowActionId,nullable" format:"uuid"`
	JSON             environmentSpecJSON `json:"-"`
}

// environmentSpecJSON contains the JSON metadata for the struct [EnvironmentSpec]
type environmentSpecJSON struct {
	Admission        apijson.Field
	AutomationsFile  apijson.Field
	Content          apijson.Field
	DesiredPhase     apijson.Field
	Devcontainer     apijson.Field
	Machine          apijson.Field
	Ports            apijson.Field
	Secrets          apijson.Field
	SpecVersion      apijson.Field
	SSHPublicKeys    apijson.Field
	Timeout          apijson.Field
	WorkflowActionID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EnvironmentSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecJSON) RawJSON() string {
	return r.raw
}

// automations_file is the automations file spec of the environment
type EnvironmentSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath string `json:"automationsFilePath"`
	Session             string `json:"session"`
	// trigger_filter specifies which automation triggers should execute. When set,
	// only automations matching these triggers will run. If empty/unset, all triggers
	// are evaluated normally.
	TriggerFilter []shared.AutomationTrigger         `json:"triggerFilter"`
	JSON          environmentSpecAutomationsFileJSON `json:"-"`
}

// environmentSpecAutomationsFileJSON contains the JSON metadata for the struct
// [EnvironmentSpecAutomationsFile]
type environmentSpecAutomationsFileJSON struct {
	AutomationsFilePath apijson.Field
	Session             apijson.Field
	TriggerFilter       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentSpecAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// content is the content spec of the environment
type EnvironmentSpecContent struct {
	// The Git email address
	GitEmail string `json:"gitEmail"`
	// The Git username
	GitUsername string `json:"gitUsername"`
	// initializer configures how the environment is to be initialized
	Initializer EnvironmentInitializer     `json:"initializer"`
	Session     string                     `json:"session"`
	JSON        environmentSpecContentJSON `json:"-"`
}

// environmentSpecContentJSON contains the JSON metadata for the struct
// [EnvironmentSpecContent]
type environmentSpecContentJSON struct {
	GitEmail    apijson.Field
	GitUsername apijson.Field
	Initializer apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentSpecContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecContentJSON) RawJSON() string {
	return r.raw
}

// devcontainer is the devcontainer spec of the environment
type EnvironmentSpecDevcontainer struct {
	// default_devcontainer_image is the default image that is used to start the
	// devcontainer if no devcontainer config file is found
	DefaultDevcontainerImage string `json:"defaultDevcontainerImage"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// Experimental: dotfiles is the dotfiles configuration of the devcontainer
	Dotfiles EnvironmentSpecDevcontainerDotfiles `json:"dotfiles"`
	// lifecycle_stage controls which devcontainer lifecycle commands are executed.
	// Defaults to FULL if not specified.
	LifecycleStage EnvironmentSpecDevcontainerLifecycleStage `json:"lifecycleStage"`
	Session        string                                    `json:"session"`
	JSON           environmentSpecDevcontainerJSON           `json:"-"`
}

// environmentSpecDevcontainerJSON contains the JSON metadata for the struct
// [EnvironmentSpecDevcontainer]
type environmentSpecDevcontainerJSON struct {
	DefaultDevcontainerImage apijson.Field
	DevcontainerFilePath     apijson.Field
	Dotfiles                 apijson.Field
	LifecycleStage           apijson.Field
	Session                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentSpecDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecDevcontainerJSON) RawJSON() string {
	return r.raw
}

// Experimental: dotfiles is the dotfiles configuration of the devcontainer
type EnvironmentSpecDevcontainerDotfiles struct {
	// URL of a dotfiles Git repository (e.g. https://github.com/owner/repository)
	Repository string                                  `json:"repository,required" format:"uri"`
	JSON       environmentSpecDevcontainerDotfilesJSON `json:"-"`
}

// environmentSpecDevcontainerDotfilesJSON contains the JSON metadata for the
// struct [EnvironmentSpecDevcontainerDotfiles]
type environmentSpecDevcontainerDotfilesJSON struct {
	Repository  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentSpecDevcontainerDotfiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecDevcontainerDotfilesJSON) RawJSON() string {
	return r.raw
}

// lifecycle_stage controls which devcontainer lifecycle commands are executed.
// Defaults to FULL if not specified.
type EnvironmentSpecDevcontainerLifecycleStage string

const (
	EnvironmentSpecDevcontainerLifecycleStageLifecycleStageUnspecified EnvironmentSpecDevcontainerLifecycleStage = "LIFECYCLE_STAGE_UNSPECIFIED"
	EnvironmentSpecDevcontainerLifecycleStageLifecycleStageFull        EnvironmentSpecDevcontainerLifecycleStage = "LIFECYCLE_STAGE_FULL"
	EnvironmentSpecDevcontainerLifecycleStageLifecycleStagePrebuild    EnvironmentSpecDevcontainerLifecycleStage = "LIFECYCLE_STAGE_PREBUILD"
)

func (r EnvironmentSpecDevcontainerLifecycleStage) IsKnown() bool {
	switch r {
	case EnvironmentSpecDevcontainerLifecycleStageLifecycleStageUnspecified, EnvironmentSpecDevcontainerLifecycleStageLifecycleStageFull, EnvironmentSpecDevcontainerLifecycleStageLifecycleStagePrebuild:
		return true
	}
	return false
}

// machine is the machine spec of the environment
type EnvironmentSpecMachine struct {
	// Class denotes the class of the environment we ought to start
	Class   string                     `json:"class"`
	Session string                     `json:"session"`
	JSON    environmentSpecMachineJSON `json:"-"`
}

// environmentSpecMachineJSON contains the JSON metadata for the struct
// [EnvironmentSpecMachine]
type environmentSpecMachineJSON struct {
	Class       apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentSpecMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecMachineJSON) RawJSON() string {
	return r.raw
}

type EnvironmentSpecPort struct {
	// policy of this port
	Admission AdmissionLevel `json:"admission"`
	// name of this port
	Name string `json:"name"`
	// port number
	Port int64 `json:"port"`
	// protocol for communication (Gateway proxy → user environment service). this
	// setting only affects the protocol used between Gateway and user environment
	// services.
	Protocol EnvironmentSpecPortsProtocol `json:"protocol"`
	JSON     environmentSpecPortJSON      `json:"-"`
}

// environmentSpecPortJSON contains the JSON metadata for the struct
// [EnvironmentSpecPort]
type environmentSpecPortJSON struct {
	Admission   apijson.Field
	Name        apijson.Field
	Port        apijson.Field
	Protocol    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentSpecPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecPortJSON) RawJSON() string {
	return r.raw
}

// protocol for communication (Gateway proxy → user environment service). this
// setting only affects the protocol used between Gateway and user environment
// services.
type EnvironmentSpecPortsProtocol string

const (
	EnvironmentSpecPortsProtocolProtocolUnspecified EnvironmentSpecPortsProtocol = "PROTOCOL_UNSPECIFIED"
	EnvironmentSpecPortsProtocolProtocolHTTP        EnvironmentSpecPortsProtocol = "PROTOCOL_HTTP"
	EnvironmentSpecPortsProtocolProtocolHTTPS       EnvironmentSpecPortsProtocol = "PROTOCOL_HTTPS"
)

func (r EnvironmentSpecPortsProtocol) IsKnown() bool {
	switch r {
	case EnvironmentSpecPortsProtocolProtocolUnspecified, EnvironmentSpecPortsProtocolProtocolHTTP, EnvironmentSpecPortsProtocolProtocolHTTPS:
		return true
	}
	return false
}

type EnvironmentSpecSecret struct {
	// id is the unique identifier of the secret.
	ID string `json:"id"`
	// api_only indicates the secret is only available via API/CLI. These secrets are
	// resolved but NOT automatically injected into services or devcontainers.
	APIOnly bool `json:"apiOnly"`
	// container_registry_basic_auth_host is the hostname of the container registry
	// that supports basic auth
	ContainerRegistryBasicAuthHost string `json:"containerRegistryBasicAuthHost"`
	EnvironmentVariable            string `json:"environmentVariable"`
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
	SourceRef string                    `json:"sourceRef"`
	JSON      environmentSpecSecretJSON `json:"-"`
}

// environmentSpecSecretJSON contains the JSON metadata for the struct
// [EnvironmentSpecSecret]
type environmentSpecSecretJSON struct {
	ID                             apijson.Field
	APIOnly                        apijson.Field
	ContainerRegistryBasicAuthHost apijson.Field
	EnvironmentVariable            apijson.Field
	FilePath                       apijson.Field
	GitCredentialHost              apijson.Field
	Name                           apijson.Field
	Session                        apijson.Field
	Source                         apijson.Field
	SourceRef                      apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *EnvironmentSpecSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecSecretJSON) RawJSON() string {
	return r.raw
}

type EnvironmentSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// value is the actual public key in the public key file format
	Value string                          `json:"value"`
	JSON  environmentSpecSSHPublicKeyJSON `json:"-"`
}

// environmentSpecSSHPublicKeyJSON contains the JSON metadata for the struct
// [EnvironmentSpecSSHPublicKey]
type environmentSpecSSHPublicKeyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentSpecSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// Timeout configures the environment timeout
type EnvironmentSpecTimeout struct {
	// inacitivity is the maximum time of disconnection before the environment is
	// stopped or paused. Minimum duration is 30 minutes. Set to 0 to disable.
	Disconnected string                     `json:"disconnected" format:"regex"`
	JSON         environmentSpecTimeoutJSON `json:"-"`
}

// environmentSpecTimeoutJSON contains the JSON metadata for the struct
// [EnvironmentSpecTimeout]
type environmentSpecTimeoutJSON struct {
	Disconnected apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentSpecTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecTimeoutJSON) RawJSON() string {
	return r.raw
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type EnvironmentSpecParam struct {
	// admission controlls who can access the environment and its ports.
	Admission param.Field[AdmissionLevel] `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile param.Field[EnvironmentSpecAutomationsFileParam] `json:"automationsFile"`
	// content is the content spec of the environment
	Content param.Field[EnvironmentSpecContentParam] `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase param.Field[EnvironmentPhase] `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer param.Field[EnvironmentSpecDevcontainerParam] `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine param.Field[EnvironmentSpecMachineParam] `json:"machine"`
	// ports is the set of ports which ought to be exposed to your network
	Ports param.Field[[]EnvironmentSpecPortParam] `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets param.Field[[]EnvironmentSpecSecretParam] `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion param.Field[string] `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys param.Field[[]EnvironmentSpecSSHPublicKeyParam] `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout param.Field[EnvironmentSpecTimeoutParam] `json:"timeout"`
	// workflow_action_id is an optional reference to the workflow execution action
	// that created this environment. Used for tracking and event correlation.
	WorkflowActionID param.Field[string] `json:"workflowActionId" format:"uuid"`
}

func (r EnvironmentSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// automations_file is the automations file spec of the environment
type EnvironmentSpecAutomationsFileParam struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath"`
	Session             param.Field[string] `json:"session"`
	// trigger_filter specifies which automation triggers should execute. When set,
	// only automations matching these triggers will run. If empty/unset, all triggers
	// are evaluated normally.
	TriggerFilter param.Field[[]shared.AutomationTriggerParam] `json:"triggerFilter"`
}

func (r EnvironmentSpecAutomationsFileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// content is the content spec of the environment
type EnvironmentSpecContentParam struct {
	// The Git email address
	GitEmail param.Field[string] `json:"gitEmail"`
	// The Git username
	GitUsername param.Field[string] `json:"gitUsername"`
	// initializer configures how the environment is to be initialized
	Initializer param.Field[EnvironmentInitializerParam] `json:"initializer"`
	Session     param.Field[string]                      `json:"session"`
}

func (r EnvironmentSpecContentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// devcontainer is the devcontainer spec of the environment
type EnvironmentSpecDevcontainerParam struct {
	// default_devcontainer_image is the default image that is used to start the
	// devcontainer if no devcontainer config file is found
	DefaultDevcontainerImage param.Field[string] `json:"defaultDevcontainerImage"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath"`
	// Experimental: dotfiles is the dotfiles configuration of the devcontainer
	Dotfiles param.Field[EnvironmentSpecDevcontainerDotfilesParam] `json:"dotfiles"`
	// lifecycle_stage controls which devcontainer lifecycle commands are executed.
	// Defaults to FULL if not specified.
	LifecycleStage param.Field[EnvironmentSpecDevcontainerLifecycleStage] `json:"lifecycleStage"`
	Session        param.Field[string]                                    `json:"session"`
}

func (r EnvironmentSpecDevcontainerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Experimental: dotfiles is the dotfiles configuration of the devcontainer
type EnvironmentSpecDevcontainerDotfilesParam struct {
	// URL of a dotfiles Git repository (e.g. https://github.com/owner/repository)
	Repository param.Field[string] `json:"repository,required" format:"uri"`
}

func (r EnvironmentSpecDevcontainerDotfilesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// machine is the machine spec of the environment
type EnvironmentSpecMachineParam struct {
	// Class denotes the class of the environment we ought to start
	Class   param.Field[string] `json:"class"`
	Session param.Field[string] `json:"session"`
}

func (r EnvironmentSpecMachineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentSpecPortParam struct {
	// policy of this port
	Admission param.Field[AdmissionLevel] `json:"admission"`
	// name of this port
	Name param.Field[string] `json:"name"`
	// port number
	Port param.Field[int64] `json:"port"`
	// protocol for communication (Gateway proxy → user environment service). this
	// setting only affects the protocol used between Gateway and user environment
	// services.
	Protocol param.Field[EnvironmentSpecPortsProtocol] `json:"protocol"`
}

func (r EnvironmentSpecPortParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentSpecSecretParam struct {
	// id is the unique identifier of the secret.
	ID param.Field[string] `json:"id"`
	// api_only indicates the secret is only available via API/CLI. These secrets are
	// resolved but NOT automatically injected into services or devcontainers.
	APIOnly param.Field[bool] `json:"apiOnly"`
	// container_registry_basic_auth_host is the hostname of the container registry
	// that supports basic auth
	ContainerRegistryBasicAuthHost param.Field[string] `json:"containerRegistryBasicAuthHost"`
	EnvironmentVariable            param.Field[string] `json:"environmentVariable"`
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

func (r EnvironmentSpecSecretParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentSpecSSHPublicKeyParam struct {
	// id is the unique identifier of the public key
	ID param.Field[string] `json:"id"`
	// value is the actual public key in the public key file format
	Value param.Field[string] `json:"value"`
}

func (r EnvironmentSpecSSHPublicKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Timeout configures the environment timeout
type EnvironmentSpecTimeoutParam struct {
	// inacitivity is the maximum time of disconnection before the environment is
	// stopped or paused. Minimum duration is 30 minutes. Set to 0 to disable.
	Disconnected param.Field[string] `json:"disconnected" format:"regex"`
}

func (r EnvironmentSpecTimeoutParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnvironmentStatus describes an environment status
type EnvironmentStatus struct {
	// activity_signal is the last activity signal for the environment.
	ActivitySignal EnvironmentActivitySignal `json:"activitySignal"`
	// automations_file contains the status of the automations file.
	AutomationsFile EnvironmentStatusAutomationsFile `json:"automationsFile"`
	// content contains the status of the environment content.
	Content EnvironmentStatusContent `json:"content"`
	// devcontainer contains the status of the devcontainer.
	Devcontainer EnvironmentStatusDevcontainer `json:"devcontainer"`
	// environment_url contains the URL at which the environment can be accessed. This
	// field is only set if the environment is running.
	EnvironmentURLs EnvironmentStatusEnvironmentURLs `json:"environmentUrls"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage []string `json:"failureMessage"`
	// machine contains the status of the environment machine
	Machine EnvironmentStatusMachine `json:"machine"`
	// the phase of an environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentPhase `json:"phase"`
	// runner_ack contains the acknowledgement from the runner that is has received the
	// environment spec.
	RunnerAck EnvironmentStatusRunnerAck `json:"runnerAck"`
	// secrets contains the status of the environment secrets
	Secrets []EnvironmentStatusSecret `json:"secrets"`
	// ssh_public_keys contains the status of the environment ssh public keys
	SSHPublicKeys []EnvironmentStatusSSHPublicKey `json:"sshPublicKeys"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion"`
	// warning_message contains warnings, e.g. when the environment is present but not
	// in the expected state.
	WarningMessage []string              `json:"warningMessage"`
	JSON           environmentStatusJSON `json:"-"`
}

// environmentStatusJSON contains the JSON metadata for the struct
// [EnvironmentStatus]
type environmentStatusJSON struct {
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

func (r *EnvironmentStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusJSON) RawJSON() string {
	return r.raw
}

// automations_file contains the status of the automations file.
type EnvironmentStatusAutomationsFile struct {
	// automations_file_path is the path to the automations file relative to the repo
	// root.
	AutomationsFilePath string `json:"automationsFilePath"`
	// automations_file_presence indicates how an automations file is present in the
	// environment.
	AutomationsFilePresence EnvironmentStatusAutomationsFileAutomationsFilePresence `json:"automationsFilePresence"`
	// failure_message contains the reason the automations file failed to be applied.
	// This is only set if the phase is FAILED.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the automations file.
	Phase EnvironmentStatusAutomationsFilePhase `json:"phase"`
	// session is the automations file session that is currently applied in the
	// environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when no triggers are defined in the
	// automations file.
	WarningMessage string                               `json:"warningMessage"`
	JSON           environmentStatusAutomationsFileJSON `json:"-"`
}

// environmentStatusAutomationsFileJSON contains the JSON metadata for the struct
// [EnvironmentStatusAutomationsFile]
type environmentStatusAutomationsFileJSON struct {
	AutomationsFilePath     apijson.Field
	AutomationsFilePresence apijson.Field
	FailureMessage          apijson.Field
	Phase                   apijson.Field
	Session                 apijson.Field
	WarningMessage          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *EnvironmentStatusAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// automations_file_presence indicates how an automations file is present in the
// environment.
type EnvironmentStatusAutomationsFileAutomationsFilePresence string

const (
	EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceUnspecified EnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceAbsent      EnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_ABSENT"
	EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceDiscovered  EnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceSpecified   EnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentStatusAutomationsFileAutomationsFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceUnspecified, EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceAbsent, EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceDiscovered, EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the automations file.
type EnvironmentStatusAutomationsFilePhase string

const (
	EnvironmentStatusAutomationsFilePhaseContentPhaseUnspecified  EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentStatusAutomationsFilePhaseContentPhaseCreating     EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_CREATING"
	EnvironmentStatusAutomationsFilePhaseContentPhaseInitializing EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentStatusAutomationsFilePhaseContentPhaseReady        EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_READY"
	EnvironmentStatusAutomationsFilePhaseContentPhaseUpdating     EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_UPDATING"
	EnvironmentStatusAutomationsFilePhaseContentPhaseFailed       EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_FAILED"
	EnvironmentStatusAutomationsFilePhaseContentPhaseUnavailable  EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_UNAVAILABLE"
)

func (r EnvironmentStatusAutomationsFilePhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusAutomationsFilePhaseContentPhaseUnspecified, EnvironmentStatusAutomationsFilePhaseContentPhaseCreating, EnvironmentStatusAutomationsFilePhaseContentPhaseInitializing, EnvironmentStatusAutomationsFilePhaseContentPhaseReady, EnvironmentStatusAutomationsFilePhaseContentPhaseUpdating, EnvironmentStatusAutomationsFilePhaseContentPhaseFailed, EnvironmentStatusAutomationsFilePhaseContentPhaseUnavailable:
		return true
	}
	return false
}

// content contains the status of the environment content.
type EnvironmentStatusContent struct {
	// content_location_in_machine is the location of the content in the machine
	ContentLocationInMachine string `json:"contentLocationInMachine"`
	// failure_message contains the reason the content initialization failed.
	FailureMessage string `json:"failureMessage"`
	// git is the Git working copy status of the environment. Note: this is a
	// best-effort field and more often than not will not be present. Its absence does
	// not indicate the absence of a working copy.
	Git EnvironmentStatusContentGit `json:"git"`
	// phase is the current phase of the environment content
	Phase EnvironmentStatusContentPhase `json:"phase"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the content is present but not in
	// the expected state.
	WarningMessage string                       `json:"warningMessage"`
	JSON           environmentStatusContentJSON `json:"-"`
}

// environmentStatusContentJSON contains the JSON metadata for the struct
// [EnvironmentStatusContent]
type environmentStatusContentJSON struct {
	ContentLocationInMachine apijson.Field
	FailureMessage           apijson.Field
	Git                      apijson.Field
	Phase                    apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentStatusContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusContentJSON) RawJSON() string {
	return r.raw
}

// git is the Git working copy status of the environment. Note: this is a
// best-effort field and more often than not will not be present. Its absence does
// not indicate the absence of a working copy.
type EnvironmentStatusContentGit struct {
	// branch is branch we're currently on
	Branch string `json:"branch"`
	// changed_files is an array of changed files in the environment, possibly
	// truncated
	ChangedFiles []EnvironmentStatusContentGitChangedFile `json:"changedFiles"`
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
	UnpushedCommits []string                        `json:"unpushedCommits"`
	JSON            environmentStatusContentGitJSON `json:"-"`
}

// environmentStatusContentGitJSON contains the JSON metadata for the struct
// [EnvironmentStatusContentGit]
type environmentStatusContentGitJSON struct {
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

func (r *EnvironmentStatusContentGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusContentGitJSON) RawJSON() string {
	return r.raw
}

type EnvironmentStatusContentGitChangedFile struct {
	// ChangeType is the type of change that happened to the file
	ChangeType EnvironmentStatusContentGitChangedFilesChangeType `json:"changeType"`
	// path is the path of the file
	Path string                                     `json:"path"`
	JSON environmentStatusContentGitChangedFileJSON `json:"-"`
}

// environmentStatusContentGitChangedFileJSON contains the JSON metadata for the
// struct [EnvironmentStatusContentGitChangedFile]
type environmentStatusContentGitChangedFileJSON struct {
	ChangeType  apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentStatusContentGitChangedFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusContentGitChangedFileJSON) RawJSON() string {
	return r.raw
}

// ChangeType is the type of change that happened to the file
type EnvironmentStatusContentGitChangedFilesChangeType string

const (
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified        EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNSPECIFIED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeAdded              EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_ADDED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeModified           EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_MODIFIED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeDeleted            EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_DELETED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeRenamed            EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_RENAMED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeCopied             EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_COPIED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UPDATED_BUT_UNMERGED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUntracked          EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNTRACKED"
)

func (r EnvironmentStatusContentGitChangedFilesChangeType) IsKnown() bool {
	switch r {
	case EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeAdded, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeModified, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeDeleted, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeRenamed, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeCopied, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUntracked:
		return true
	}
	return false
}

// phase is the current phase of the environment content
type EnvironmentStatusContentPhase string

const (
	EnvironmentStatusContentPhaseContentPhaseUnspecified  EnvironmentStatusContentPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentStatusContentPhaseContentPhaseCreating     EnvironmentStatusContentPhase = "CONTENT_PHASE_CREATING"
	EnvironmentStatusContentPhaseContentPhaseInitializing EnvironmentStatusContentPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentStatusContentPhaseContentPhaseReady        EnvironmentStatusContentPhase = "CONTENT_PHASE_READY"
	EnvironmentStatusContentPhaseContentPhaseUpdating     EnvironmentStatusContentPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentStatusContentPhaseContentPhaseFailed       EnvironmentStatusContentPhase = "CONTENT_PHASE_FAILED"
	EnvironmentStatusContentPhaseContentPhaseUnavailable  EnvironmentStatusContentPhase = "CONTENT_PHASE_UNAVAILABLE"
)

func (r EnvironmentStatusContentPhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusContentPhaseContentPhaseUnspecified, EnvironmentStatusContentPhaseContentPhaseCreating, EnvironmentStatusContentPhaseContentPhaseInitializing, EnvironmentStatusContentPhaseContentPhaseReady, EnvironmentStatusContentPhaseContentPhaseUpdating, EnvironmentStatusContentPhaseContentPhaseFailed, EnvironmentStatusContentPhaseContentPhaseUnavailable:
		return true
	}
	return false
}

// devcontainer contains the status of the devcontainer.
type EnvironmentStatusDevcontainer struct {
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
	DevcontainerFilePresence EnvironmentStatusDevcontainerDevcontainerFilePresence `json:"devcontainerFilePresence"`
	// failure_message contains the reason the devcontainer failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the devcontainer
	Phase EnvironmentStatusDevcontainerPhase `json:"phase"`
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
	WarningMessage string                            `json:"warningMessage"`
	JSON           environmentStatusDevcontainerJSON `json:"-"`
}

// environmentStatusDevcontainerJSON contains the JSON metadata for the struct
// [EnvironmentStatusDevcontainer]
type environmentStatusDevcontainerJSON struct {
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

func (r *EnvironmentStatusDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusDevcontainerJSON) RawJSON() string {
	return r.raw
}

// devcontainer_file_presence indicates how the devcontainer file is present in the
// repo.
type EnvironmentStatusDevcontainerDevcontainerFilePresence string

const (
	EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceUnspecified EnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceGenerated   EnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_GENERATED"
	EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceDiscovered  EnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceSpecified   EnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentStatusDevcontainerDevcontainerFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceUnspecified, EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceGenerated, EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceDiscovered, EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the devcontainer
type EnvironmentStatusDevcontainerPhase string

const (
	EnvironmentStatusDevcontainerPhasePhaseUnspecified EnvironmentStatusDevcontainerPhase = "PHASE_UNSPECIFIED"
	EnvironmentStatusDevcontainerPhasePhaseCreating    EnvironmentStatusDevcontainerPhase = "PHASE_CREATING"
	EnvironmentStatusDevcontainerPhasePhaseRunning     EnvironmentStatusDevcontainerPhase = "PHASE_RUNNING"
	EnvironmentStatusDevcontainerPhasePhaseStopped     EnvironmentStatusDevcontainerPhase = "PHASE_STOPPED"
	EnvironmentStatusDevcontainerPhasePhaseFailed      EnvironmentStatusDevcontainerPhase = "PHASE_FAILED"
)

func (r EnvironmentStatusDevcontainerPhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusDevcontainerPhasePhaseUnspecified, EnvironmentStatusDevcontainerPhasePhaseCreating, EnvironmentStatusDevcontainerPhasePhaseRunning, EnvironmentStatusDevcontainerPhasePhaseStopped, EnvironmentStatusDevcontainerPhasePhaseFailed:
		return true
	}
	return false
}

// environment_url contains the URL at which the environment can be accessed. This
// field is only set if the environment is running.
type EnvironmentStatusEnvironmentURLs struct {
	// logs is the URL at which the environment logs can be accessed.
	Logs string `json:"logs"`
	// ops is the URL at which the environment ops service can be accessed.
	Ops   string                                 `json:"ops"`
	Ports []EnvironmentStatusEnvironmentURLsPort `json:"ports"`
	// SSH is the URL at which the environment can be accessed via SSH.
	SSH EnvironmentStatusEnvironmentURLsSSH `json:"ssh"`
	// support_bundle is the URL at which the environment support bundle can be
	// accessed.
	SupportBundle string                               `json:"supportBundle"`
	JSON          environmentStatusEnvironmentURLsJSON `json:"-"`
}

// environmentStatusEnvironmentURLsJSON contains the JSON metadata for the struct
// [EnvironmentStatusEnvironmentURLs]
type environmentStatusEnvironmentURLsJSON struct {
	Logs          apijson.Field
	Ops           apijson.Field
	Ports         apijson.Field
	SSH           apijson.Field
	SupportBundle apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentStatusEnvironmentURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusEnvironmentURLsJSON) RawJSON() string {
	return r.raw
}

type EnvironmentStatusEnvironmentURLsPort struct {
	// port is the port number of the environment port
	Port int64 `json:"port"`
	// url is the URL at which the environment port can be accessed
	URL  string                                   `json:"url"`
	JSON environmentStatusEnvironmentURLsPortJSON `json:"-"`
}

// environmentStatusEnvironmentURLsPortJSON contains the JSON metadata for the
// struct [EnvironmentStatusEnvironmentURLsPort]
type environmentStatusEnvironmentURLsPortJSON struct {
	Port        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentStatusEnvironmentURLsPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusEnvironmentURLsPortJSON) RawJSON() string {
	return r.raw
}

// SSH is the URL at which the environment can be accessed via SSH.
type EnvironmentStatusEnvironmentURLsSSH struct {
	URL  string                                  `json:"url"`
	JSON environmentStatusEnvironmentURLsSSHJSON `json:"-"`
}

// environmentStatusEnvironmentURLsSSHJSON contains the JSON metadata for the
// struct [EnvironmentStatusEnvironmentURLsSSH]
type environmentStatusEnvironmentURLsSSHJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentStatusEnvironmentURLsSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusEnvironmentURLsSSHJSON) RawJSON() string {
	return r.raw
}

// machine contains the status of the environment machine
type EnvironmentStatusMachine struct {
	// failure_message contains the reason the machine failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the environment machine
	Phase EnvironmentStatusMachinePhase `json:"phase"`
	// session is the session that is currently active in the machine.
	Session string `json:"session"`
	// timeout contains the reason the environment has timed out. If this field is
	// empty, the environment has not timed out.
	Timeout string `json:"timeout"`
	// versions contains the versions of components in the machine.
	Versions EnvironmentStatusMachineVersions `json:"versions"`
	// warning_message contains warnings, e.g. when the machine is present but not in
	// the expected state.
	WarningMessage string                       `json:"warningMessage"`
	JSON           environmentStatusMachineJSON `json:"-"`
}

// environmentStatusMachineJSON contains the JSON metadata for the struct
// [EnvironmentStatusMachine]
type environmentStatusMachineJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	Timeout        apijson.Field
	Versions       apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentStatusMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusMachineJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the environment machine
type EnvironmentStatusMachinePhase string

const (
	EnvironmentStatusMachinePhasePhaseUnspecified EnvironmentStatusMachinePhase = "PHASE_UNSPECIFIED"
	EnvironmentStatusMachinePhasePhaseCreating    EnvironmentStatusMachinePhase = "PHASE_CREATING"
	EnvironmentStatusMachinePhasePhaseStarting    EnvironmentStatusMachinePhase = "PHASE_STARTING"
	EnvironmentStatusMachinePhasePhaseRunning     EnvironmentStatusMachinePhase = "PHASE_RUNNING"
	EnvironmentStatusMachinePhasePhaseStopping    EnvironmentStatusMachinePhase = "PHASE_STOPPING"
	EnvironmentStatusMachinePhasePhaseStopped     EnvironmentStatusMachinePhase = "PHASE_STOPPED"
	EnvironmentStatusMachinePhasePhaseDeleting    EnvironmentStatusMachinePhase = "PHASE_DELETING"
	EnvironmentStatusMachinePhasePhaseDeleted     EnvironmentStatusMachinePhase = "PHASE_DELETED"
)

func (r EnvironmentStatusMachinePhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusMachinePhasePhaseUnspecified, EnvironmentStatusMachinePhasePhaseCreating, EnvironmentStatusMachinePhasePhaseStarting, EnvironmentStatusMachinePhasePhaseRunning, EnvironmentStatusMachinePhasePhaseStopping, EnvironmentStatusMachinePhasePhaseStopped, EnvironmentStatusMachinePhasePhaseDeleting, EnvironmentStatusMachinePhasePhaseDeleted:
		return true
	}
	return false
}

// versions contains the versions of components in the machine.
type EnvironmentStatusMachineVersions struct {
	AmiID             string                               `json:"amiId"`
	SupervisorCommit  string                               `json:"supervisorCommit"`
	SupervisorVersion string                               `json:"supervisorVersion"`
	JSON              environmentStatusMachineVersionsJSON `json:"-"`
}

// environmentStatusMachineVersionsJSON contains the JSON metadata for the struct
// [EnvironmentStatusMachineVersions]
type environmentStatusMachineVersionsJSON struct {
	AmiID             apijson.Field
	SupervisorCommit  apijson.Field
	SupervisorVersion apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentStatusMachineVersions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusMachineVersionsJSON) RawJSON() string {
	return r.raw
}

// runner_ack contains the acknowledgement from the runner that is has received the
// environment spec.
type EnvironmentStatusRunnerAck struct {
	Message     string                               `json:"message"`
	SpecVersion string                               `json:"specVersion"`
	StatusCode  EnvironmentStatusRunnerAckStatusCode `json:"statusCode"`
	JSON        environmentStatusRunnerAckJSON       `json:"-"`
}

// environmentStatusRunnerAckJSON contains the JSON metadata for the struct
// [EnvironmentStatusRunnerAck]
type environmentStatusRunnerAckJSON struct {
	Message     apijson.Field
	SpecVersion apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentStatusRunnerAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusRunnerAckJSON) RawJSON() string {
	return r.raw
}

type EnvironmentStatusRunnerAckStatusCode string

const (
	EnvironmentStatusRunnerAckStatusCodeStatusCodeUnspecified        EnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_UNSPECIFIED"
	EnvironmentStatusRunnerAckStatusCodeStatusCodeOk                 EnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_OK"
	EnvironmentStatusRunnerAckStatusCodeStatusCodeInvalidResource    EnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_INVALID_RESOURCE"
	EnvironmentStatusRunnerAckStatusCodeStatusCodeFailedPrecondition EnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_FAILED_PRECONDITION"
)

func (r EnvironmentStatusRunnerAckStatusCode) IsKnown() bool {
	switch r {
	case EnvironmentStatusRunnerAckStatusCodeStatusCodeUnspecified, EnvironmentStatusRunnerAckStatusCodeStatusCodeOk, EnvironmentStatusRunnerAckStatusCodeStatusCodeInvalidResource, EnvironmentStatusRunnerAckStatusCodeStatusCodeFailedPrecondition:
		return true
	}
	return false
}

type EnvironmentStatusSecret struct {
	// id is the unique identifier of the secret.
	ID string `json:"id"`
	// failure_message contains the reason the secret failed to be materialize.
	FailureMessage string                        `json:"failureMessage"`
	Phase          EnvironmentStatusSecretsPhase `json:"phase"`
	SecretName     string                        `json:"secretName"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the secret is present but not in
	// the expected state.
	WarningMessage string                      `json:"warningMessage"`
	JSON           environmentStatusSecretJSON `json:"-"`
}

// environmentStatusSecretJSON contains the JSON metadata for the struct
// [EnvironmentStatusSecret]
type environmentStatusSecretJSON struct {
	ID             apijson.Field
	FailureMessage apijson.Field
	Phase          apijson.Field
	SecretName     apijson.Field
	Session        apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentStatusSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusSecretJSON) RawJSON() string {
	return r.raw
}

type EnvironmentStatusSecretsPhase string

const (
	EnvironmentStatusSecretsPhaseContentPhaseUnspecified  EnvironmentStatusSecretsPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentStatusSecretsPhaseContentPhaseCreating     EnvironmentStatusSecretsPhase = "CONTENT_PHASE_CREATING"
	EnvironmentStatusSecretsPhaseContentPhaseInitializing EnvironmentStatusSecretsPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentStatusSecretsPhaseContentPhaseReady        EnvironmentStatusSecretsPhase = "CONTENT_PHASE_READY"
	EnvironmentStatusSecretsPhaseContentPhaseUpdating     EnvironmentStatusSecretsPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentStatusSecretsPhaseContentPhaseFailed       EnvironmentStatusSecretsPhase = "CONTENT_PHASE_FAILED"
	EnvironmentStatusSecretsPhaseContentPhaseUnavailable  EnvironmentStatusSecretsPhase = "CONTENT_PHASE_UNAVAILABLE"
)

func (r EnvironmentStatusSecretsPhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusSecretsPhaseContentPhaseUnspecified, EnvironmentStatusSecretsPhaseContentPhaseCreating, EnvironmentStatusSecretsPhaseContentPhaseInitializing, EnvironmentStatusSecretsPhaseContentPhaseReady, EnvironmentStatusSecretsPhaseContentPhaseUpdating, EnvironmentStatusSecretsPhaseContentPhaseFailed, EnvironmentStatusSecretsPhaseContentPhaseUnavailable:
		return true
	}
	return false
}

type EnvironmentStatusSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// phase is the current phase of the public key
	Phase EnvironmentStatusSSHPublicKeysPhase `json:"phase"`
	JSON  environmentStatusSSHPublicKeyJSON   `json:"-"`
}

// environmentStatusSSHPublicKeyJSON contains the JSON metadata for the struct
// [EnvironmentStatusSSHPublicKey]
type environmentStatusSSHPublicKeyJSON struct {
	ID          apijson.Field
	Phase       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentStatusSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the public key
type EnvironmentStatusSSHPublicKeysPhase string

const (
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseUnspecified  EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseCreating     EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_CREATING"
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseInitializing EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseReady        EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_READY"
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseUpdating     EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseFailed       EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_FAILED"
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseUnavailable  EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_UNAVAILABLE"
)

func (r EnvironmentStatusSSHPublicKeysPhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusSSHPublicKeysPhaseContentPhaseUnspecified, EnvironmentStatusSSHPublicKeysPhaseContentPhaseCreating, EnvironmentStatusSSHPublicKeysPhaseContentPhaseInitializing, EnvironmentStatusSSHPublicKeysPhaseContentPhaseReady, EnvironmentStatusSSHPublicKeysPhaseContentPhaseUpdating, EnvironmentStatusSSHPublicKeysPhaseContentPhaseFailed, EnvironmentStatusSSHPublicKeysPhaseContentPhaseUnavailable:
		return true
	}
	return false
}

type EnvironmentNewResponse struct {
	// +resource get environment
	Environment Environment                `json:"environment,required"`
	JSON        environmentNewResponseJSON `json:"-"`
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

type EnvironmentGetResponse struct {
	// +resource get environment
	Environment Environment                `json:"environment,required"`
	JSON        environmentGetResponseJSON `json:"-"`
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

type EnvironmentUpdateResponse = interface{}

type EnvironmentDeleteResponse = interface{}

type EnvironmentNewEnvironmentTokenResponse struct {
	// access_token is the token that can be used for environment authentication
	AccessToken string                                     `json:"accessToken,required"`
	JSON        environmentNewEnvironmentTokenResponseJSON `json:"-"`
}

// environmentNewEnvironmentTokenResponseJSON contains the JSON metadata for the
// struct [EnvironmentNewEnvironmentTokenResponse]
type environmentNewEnvironmentTokenResponseJSON struct {
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewEnvironmentTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewEnvironmentTokenResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewFromProjectResponse struct {
	// +resource get environment
	Environment Environment                           `json:"environment,required"`
	JSON        environmentNewFromProjectResponseJSON `json:"-"`
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

type EnvironmentNewLogsTokenResponse struct {
	// access_token is the token that can be used to access the logs of the environment
	AccessToken string                              `json:"accessToken,required"`
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

type EnvironmentUnarchiveResponse = interface{}

type EnvironmentNewParams struct {
	// spec is the configuration of the environment that's required for the to start
	// the environment
	Spec param.Field[EnvironmentSpecParam] `json:"spec"`
}

func (r EnvironmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentGetParams struct {
	// environment_id specifies the environment to get
	EnvironmentID param.Field[string] `json:"environmentId,required" format:"uuid"`
}

func (r EnvironmentGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParams struct {
	// environment_id specifies which environment should be updated.
	//
	// +required
	EnvironmentID param.Field[string]                          `json:"environmentId" format:"uuid"`
	Metadata      param.Field[EnvironmentUpdateParamsMetadata] `json:"metadata"`
	Spec          param.Field[EnvironmentUpdateParamsSpec]     `json:"spec"`
}

func (r EnvironmentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsMetadata struct {
	// name is the user-defined display name of the environment
	Name param.Field[string] `json:"name"`
}

func (r EnvironmentUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsSpec struct {
	// automations_file is the automations file spec of the environment
	AutomationsFile param.Field[EnvironmentUpdateParamsSpecAutomationsFile] `json:"automationsFile"`
	Content         param.Field[EnvironmentUpdateParamsSpecContent]         `json:"content"`
	Devcontainer    param.Field[EnvironmentUpdateParamsSpecDevcontainer]    `json:"devcontainer"`
	// ports controls port sharing
	Ports param.Field[[]EnvironmentUpdateParamsSpecPort] `json:"ports"`
	// ssh_public_keys are the public keys to update empty array means nothing to
	// update
	SSHPublicKeys param.Field[[]EnvironmentUpdateParamsSpecSSHPublicKey] `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout param.Field[EnvironmentUpdateParamsSpecTimeout] `json:"timeout"`
}

func (r EnvironmentUpdateParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// automations_file is the automations file spec of the environment
type EnvironmentUpdateParamsSpecAutomationsFile struct {
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

func (r EnvironmentUpdateParamsSpecAutomationsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsSpecContent struct {
	// The Git email address
	GitEmail param.Field[string] `json:"gitEmail"`
	// The Git username
	GitUsername param.Field[string] `json:"gitUsername"`
	// initializer configures how the environment is to be initialized
	Initializer param.Field[EnvironmentInitializerParam] `json:"initializer"`
	// session should be changed to trigger a content reinitialization
	Session param.Field[string] `json:"session"`
}

func (r EnvironmentUpdateParamsSpecContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsSpecDevcontainer struct {
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

func (r EnvironmentUpdateParamsSpecDevcontainer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsSpecPort struct {
	// policy of this port
	Admission param.Field[AdmissionLevel] `json:"admission"`
	// name of this port
	Name param.Field[string] `json:"name"`
	// port number
	Port param.Field[int64] `json:"port"`
	// protocol for communication (Gateway proxy → user environment service). this
	// setting only affects the protocol used between Gateway and user environment
	// services.
	Protocol param.Field[EnvironmentUpdateParamsSpecPortsProtocol] `json:"protocol"`
}

func (r EnvironmentUpdateParamsSpecPort) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// protocol for communication (Gateway proxy → user environment service). this
// setting only affects the protocol used between Gateway and user environment
// services.
type EnvironmentUpdateParamsSpecPortsProtocol string

const (
	EnvironmentUpdateParamsSpecPortsProtocolProtocolUnspecified EnvironmentUpdateParamsSpecPortsProtocol = "PROTOCOL_UNSPECIFIED"
	EnvironmentUpdateParamsSpecPortsProtocolProtocolHTTP        EnvironmentUpdateParamsSpecPortsProtocol = "PROTOCOL_HTTP"
	EnvironmentUpdateParamsSpecPortsProtocolProtocolHTTPS       EnvironmentUpdateParamsSpecPortsProtocol = "PROTOCOL_HTTPS"
)

func (r EnvironmentUpdateParamsSpecPortsProtocol) IsKnown() bool {
	switch r {
	case EnvironmentUpdateParamsSpecPortsProtocolProtocolUnspecified, EnvironmentUpdateParamsSpecPortsProtocolProtocolHTTP, EnvironmentUpdateParamsSpecPortsProtocolProtocolHTTPS:
		return true
	}
	return false
}

type EnvironmentUpdateParamsSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID param.Field[string] `json:"id"`
	// value is the actual public key in the public key file format if not provided,
	// the public key will be removed
	Value param.Field[string] `json:"value"`
}

func (r EnvironmentUpdateParamsSpecSSHPublicKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Timeout configures the environment timeout
type EnvironmentUpdateParamsSpecTimeout struct {
	// inacitivity is the maximum time of disconnection before the environment is
	// stopped or paused. Minimum duration is 30 minutes. Set to 0 to disable.
	Disconnected param.Field[string] `json:"disconnected" format:"regex"`
}

func (r EnvironmentUpdateParamsSpecTimeout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentListParams struct {
	Token    param.Field[string]                      `query:"token"`
	PageSize param.Field[int64]                       `query:"pageSize"`
	Filter   param.Field[EnvironmentListParamsFilter] `json:"filter"`
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
	// archival_status filters the response based on environment archive status
	ArchivalStatus param.Field[EnvironmentListParamsFilterArchivalStatus] `json:"archivalStatus"`
	// created_before filters environments created before this timestamp
	CreatedBefore param.Field[time.Time] `json:"createdBefore" format:"date-time"`
	// creator_ids filters the response to only Environments created by specified
	// members
	CreatorIDs param.Field[[]string] `json:"creatorIds" format:"uuid"`
	// project_ids filters the response to only Environments associated with the
	// specified projects
	ProjectIDs param.Field[[]string] `json:"projectIds" format:"uuid"`
	// roles filters the response to only Environments with the specified roles
	Roles param.Field[[]EnvironmentRole] `json:"roles"`
	// runner_ids filters the response to only Environments running on these Runner IDs
	RunnerIDs param.Field[[]string] `json:"runnerIds" format:"uuid"`
	// runner_kinds filters the response to only Environments running on these Runner
	// Kinds
	RunnerKinds param.Field[[]RunnerKind] `json:"runnerKinds"`
	// actual_phases is a list of phases the environment must be in for it to be
	// returned in the API call
	StatusPhases param.Field[[]EnvironmentPhase] `json:"statusPhases"`
}

func (r EnvironmentListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// archival_status filters the response based on environment archive status
type EnvironmentListParamsFilterArchivalStatus string

const (
	EnvironmentListParamsFilterArchivalStatusArchivalStatusUnspecified EnvironmentListParamsFilterArchivalStatus = "ARCHIVAL_STATUS_UNSPECIFIED"
	EnvironmentListParamsFilterArchivalStatusArchivalStatusActive      EnvironmentListParamsFilterArchivalStatus = "ARCHIVAL_STATUS_ACTIVE"
	EnvironmentListParamsFilterArchivalStatusArchivalStatusArchived    EnvironmentListParamsFilterArchivalStatus = "ARCHIVAL_STATUS_ARCHIVED"
	EnvironmentListParamsFilterArchivalStatusArchivalStatusAll         EnvironmentListParamsFilterArchivalStatus = "ARCHIVAL_STATUS_ALL"
)

func (r EnvironmentListParamsFilterArchivalStatus) IsKnown() bool {
	switch r {
	case EnvironmentListParamsFilterArchivalStatusArchivalStatusUnspecified, EnvironmentListParamsFilterArchivalStatusArchivalStatusActive, EnvironmentListParamsFilterArchivalStatusArchivalStatusArchived, EnvironmentListParamsFilterArchivalStatusArchivalStatusAll:
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

type EnvironmentNewEnvironmentTokenParams struct {
	// environment_id specifies the environment for which the access token should be
	// created.
	EnvironmentID param.Field[string] `json:"environmentId,required" format:"uuid"`
}

func (r EnvironmentNewEnvironmentTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewFromProjectParams struct {
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
	// Spec is the configuration of the environment that's required for the runner to
	// start the environment Configuration already defined in the Project will override
	// parts of the spec, if set
	Spec param.Field[EnvironmentSpecParam] `json:"spec"`
}

func (r EnvironmentNewFromProjectParams) MarshalJSON() (data []byte, err error) {
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
	// activity_signal specifies the activity.
	ActivitySignal param.Field[EnvironmentActivitySignalParam] `json:"activitySignal"`
	// The ID of the environment to update activity for.
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentMarkActiveParams) MarshalJSON() (data []byte, err error) {
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

type EnvironmentUnarchiveParams struct {
	// environment_id specifies the environment to unarchive.
	//
	// +required
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentUnarchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
