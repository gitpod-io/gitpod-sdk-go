// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
)

// An AutomationTrigger represents a trigger for an automation action. The `manual`
// field shows a start button in the UI for manually triggering the automation. The
// `post_machine_start` field indicates that the automation should be triggered
// after the machine has started, before the devcontainer is ready. This is used
// for machine-level services like security agents that need to start early. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started (devcontainer ready). The
// `post_devcontainer_start` field indicates that the automation should be
// triggered after the dev container has started. The `prebuild` field starts the
// automation during a prebuild of an environment. This phase does not have user
// secrets available. The `before_snapshot` field triggers the automation after all
// prebuild tasks complete but before the snapshot is taken. This is used for tasks
// that need to run last during prebuilds, such as IDE warmup. Note: The prebuild
// and before_snapshot triggers can only be used with tasks, not services.
type AutomationTrigger struct {
	BeforeSnapshot        bool                  `json:"beforeSnapshot"`
	Manual                bool                  `json:"manual"`
	PostDevcontainerStart bool                  `json:"postDevcontainerStart"`
	PostEnvironmentStart  bool                  `json:"postEnvironmentStart"`
	PostMachineStart      bool                  `json:"postMachineStart"`
	Prebuild              bool                  `json:"prebuild"`
	JSON                  automationTriggerJSON `json:"-"`
}

// automationTriggerJSON contains the JSON metadata for the struct
// [AutomationTrigger]
type automationTriggerJSON struct {
	BeforeSnapshot        apijson.Field
	Manual                apijson.Field
	PostDevcontainerStart apijson.Field
	PostEnvironmentStart  apijson.Field
	PostMachineStart      apijson.Field
	Prebuild              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AutomationTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationTriggerJSON) RawJSON() string {
	return r.raw
}

// An AutomationTrigger represents a trigger for an automation action. The `manual`
// field shows a start button in the UI for manually triggering the automation. The
// `post_machine_start` field indicates that the automation should be triggered
// after the machine has started, before the devcontainer is ready. This is used
// for machine-level services like security agents that need to start early. The
// `post_environment_start` field indicates that the automation should be triggered
// after the environment has started (devcontainer ready). The
// `post_devcontainer_start` field indicates that the automation should be
// triggered after the dev container has started. The `prebuild` field starts the
// automation during a prebuild of an environment. This phase does not have user
// secrets available. The `before_snapshot` field triggers the automation after all
// prebuild tasks complete but before the snapshot is taken. This is used for tasks
// that need to run last during prebuilds, such as IDE warmup. Note: The prebuild
// and before_snapshot triggers can only be used with tasks, not services.
type AutomationTriggerParam struct {
	BeforeSnapshot        param.Field[bool] `json:"beforeSnapshot"`
	Manual                param.Field[bool] `json:"manual"`
	PostDevcontainerStart param.Field[bool] `json:"postDevcontainerStart"`
	PostEnvironmentStart  param.Field[bool] `json:"postEnvironmentStart"`
	PostMachineStart      param.Field[bool] `json:"postMachineStart"`
	Prebuild              param.Field[bool] `json:"prebuild"`
}

func (r AutomationTriggerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID string `json:"id,required"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID string `json:"runnerId,required"`
	// configuration describes the configuration of the environment class
	Configuration []FieldValue `json:"configuration"`
	// description is a human readable description of the environment class
	Description string `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName string `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled bool                 `json:"enabled"`
	JSON    environmentClassJSON `json:"-"`
}

// environmentClassJSON contains the JSON metadata for the struct
// [EnvironmentClass]
type environmentClassJSON struct {
	ID            apijson.Field
	RunnerID      apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassJSON) RawJSON() string {
	return r.raw
}

type EnvironmentClassParam struct {
	// id is the unique identifier of the environment class
	ID param.Field[string] `json:"id,required"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID param.Field[string] `json:"runnerId,required"`
	// configuration describes the configuration of the environment class
	Configuration param.Field[[]FieldValueParam] `json:"configuration"`
	// description is a human readable description of the environment class
	Description param.Field[string] `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName param.Field[string] `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r EnvironmentClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnvironmentVariableItem represents an environment variable that can be set
// either from a literal value or from a secret reference.
type EnvironmentVariableItem struct {
	// name is the environment variable name.
	Name string `json:"name"`
	// value is a literal string value.
	Value string `json:"value"`
	// value_from specifies a source for the value.
	ValueFrom EnvironmentVariableSource   `json:"valueFrom"`
	JSON      environmentVariableItemJSON `json:"-"`
}

// environmentVariableItemJSON contains the JSON metadata for the struct
// [EnvironmentVariableItem]
type environmentVariableItemJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	ValueFrom   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentVariableItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentVariableItemJSON) RawJSON() string {
	return r.raw
}

// EnvironmentVariableItem represents an environment variable that can be set
// either from a literal value or from a secret reference.
type EnvironmentVariableItemParam struct {
	// name is the environment variable name.
	Name param.Field[string] `json:"name"`
	// value is a literal string value.
	Value param.Field[string] `json:"value"`
	// value_from specifies a source for the value.
	ValueFrom param.Field[EnvironmentVariableSourceParam] `json:"valueFrom"`
}

func (r EnvironmentVariableItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnvironmentVariableSource specifies a source for an environment variable value.
type EnvironmentVariableSource struct {
	// secret_ref references a secret by ID.
	SecretRef SecretRef                     `json:"secretRef,required"`
	JSON      environmentVariableSourceJSON `json:"-"`
}

// environmentVariableSourceJSON contains the JSON metadata for the struct
// [EnvironmentVariableSource]
type environmentVariableSourceJSON struct {
	SecretRef   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentVariableSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentVariableSourceJSON) RawJSON() string {
	return r.raw
}

// EnvironmentVariableSource specifies a source for an environment variable value.
type EnvironmentVariableSourceParam struct {
	// secret_ref references a secret by ID.
	SecretRef param.Field[SecretRefParam] `json:"secretRef,required"`
}

func (r EnvironmentVariableSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FieldValue struct {
	Key   string         `json:"key"`
	Value string         `json:"value"`
	JSON  fieldValueJSON `json:"-"`
}

// fieldValueJSON contains the JSON metadata for the struct [FieldValue]
type fieldValueJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldValueJSON) RawJSON() string {
	return r.raw
}

type FieldValueParam struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r FieldValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Gateway represents a system gateway that provides access to services
type Gateway struct {
	// name is the human-readable name of the gateway. name is unique across all
	// gateways.
	Name string `json:"name,required"`
	// url of the gateway
	URL string `json:"url,required"`
	// region is the geographical region where the gateway is located
	Region string      `json:"region"`
	JSON   gatewayJSON `json:"-"`
}

// gatewayJSON contains the JSON metadata for the struct [Gateway]
type gatewayJSON struct {
	Name        apijson.Field
	URL         apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Gateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayJSON) RawJSON() string {
	return r.raw
}

type OrganizationRole string

const (
	OrganizationRoleUnspecified OrganizationRole = "ORGANIZATION_ROLE_UNSPECIFIED"
	OrganizationRoleAdmin       OrganizationRole = "ORGANIZATION_ROLE_ADMIN"
	OrganizationRoleMember      OrganizationRole = "ORGANIZATION_ROLE_MEMBER"
)

func (r OrganizationRole) IsKnown() bool {
	switch r {
	case OrganizationRoleUnspecified, OrganizationRoleAdmin, OrganizationRoleMember:
		return true
	}
	return false
}

type OrganizationTier string

const (
	OrganizationTierUnspecified OrganizationTier = "ORGANIZATION_TIER_UNSPECIFIED"
	OrganizationTierFree        OrganizationTier = "ORGANIZATION_TIER_FREE"
	OrganizationTierEnterprise  OrganizationTier = "ORGANIZATION_TIER_ENTERPRISE"
	OrganizationTierCore        OrganizationTier = "ORGANIZATION_TIER_CORE"
	OrganizationTierFreeOna     OrganizationTier = "ORGANIZATION_TIER_FREE_ONA"
)

func (r OrganizationTier) IsKnown() bool {
	switch r {
	case OrganizationTierUnspecified, OrganizationTierFree, OrganizationTierEnterprise, OrganizationTierCore, OrganizationTierFreeOna:
		return true
	}
	return false
}

type Principal string

const (
	PrincipalUnspecified    Principal = "PRINCIPAL_UNSPECIFIED"
	PrincipalAccount        Principal = "PRINCIPAL_ACCOUNT"
	PrincipalUser           Principal = "PRINCIPAL_USER"
	PrincipalRunner         Principal = "PRINCIPAL_RUNNER"
	PrincipalEnvironment    Principal = "PRINCIPAL_ENVIRONMENT"
	PrincipalServiceAccount Principal = "PRINCIPAL_SERVICE_ACCOUNT"
	PrincipalRunnerManager  Principal = "PRINCIPAL_RUNNER_MANAGER"
	PrincipalAgentExecution Principal = "PRINCIPAL_AGENT_EXECUTION"
)

func (r Principal) IsKnown() bool {
	switch r {
	case PrincipalUnspecified, PrincipalAccount, PrincipalUser, PrincipalRunner, PrincipalEnvironment, PrincipalServiceAccount, PrincipalRunnerManager, PrincipalAgentExecution:
		return true
	}
	return false
}

type ProjectEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool `json:"localRunner"`
	// order is the priority of this entry
	Order int64                       `json:"order"`
	JSON  projectEnvironmentClassJSON `json:"-"`
}

// projectEnvironmentClassJSON contains the JSON metadata for the struct
// [ProjectEnvironmentClass]
type projectEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	Order              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type ProjectEnvironmentClassParam struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID param.Field[string] `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner param.Field[bool] `json:"localRunner"`
	// order is the priority of this entry
	Order param.Field[int64] `json:"order"`
}

func (r ProjectEnvironmentClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ResourceRole represents roles that can be assigned to groups on resources These
// map directly to the roles defined in backend/db/rule/rbac/role/role.go
type ResourceRole string

const (
	ResourceRoleUnspecified                    ResourceRole = "RESOURCE_ROLE_UNSPECIFIED"
	ResourceRoleOrgAdmin                       ResourceRole = "RESOURCE_ROLE_ORG_ADMIN"
	ResourceRoleOrgMember                      ResourceRole = "RESOURCE_ROLE_ORG_MEMBER"
	ResourceRoleOrgRunnersAdmin                ResourceRole = "RESOURCE_ROLE_ORG_RUNNERS_ADMIN"
	ResourceRoleOrgProjectsAdmin               ResourceRole = "RESOURCE_ROLE_ORG_PROJECTS_ADMIN"
	ResourceRoleOrgAutomationsAdmin            ResourceRole = "RESOURCE_ROLE_ORG_AUTOMATIONS_ADMIN"
	ResourceRoleOrgGroupsAdmin                 ResourceRole = "RESOURCE_ROLE_ORG_GROUPS_ADMIN"
	ResourceRoleGroupAdmin                     ResourceRole = "RESOURCE_ROLE_GROUP_ADMIN"
	ResourceRoleGroupViewer                    ResourceRole = "RESOURCE_ROLE_GROUP_VIEWER"
	ResourceRoleUserIdentity                   ResourceRole = "RESOURCE_ROLE_USER_IDENTITY"
	ResourceRoleUserViewer                     ResourceRole = "RESOURCE_ROLE_USER_VIEWER"
	ResourceRoleUserAdmin                      ResourceRole = "RESOURCE_ROLE_USER_ADMIN"
	ResourceRoleEnvironmentIdentity            ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_IDENTITY"
	ResourceRoleEnvironmentAdmin               ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_ADMIN"
	ResourceRoleEnvironmentUser                ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_USER"
	ResourceRoleEnvironmentViewer              ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_VIEWER"
	ResourceRoleEnvironmentRunner              ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_RUNNER"
	ResourceRoleRunnerIdentity                 ResourceRole = "RESOURCE_ROLE_RUNNER_IDENTITY"
	ResourceRoleRunnerAdmin                    ResourceRole = "RESOURCE_ROLE_RUNNER_ADMIN"
	ResourceRoleRunnerLocalAdmin               ResourceRole = "RESOURCE_ROLE_RUNNER_LOCAL_ADMIN"
	ResourceRoleRunnerManagedAdmin             ResourceRole = "RESOURCE_ROLE_RUNNER_MANAGED_ADMIN"
	ResourceRoleRunnerUser                     ResourceRole = "RESOURCE_ROLE_RUNNER_USER"
	ResourceRoleRunnerConfigurationReader      ResourceRole = "RESOURCE_ROLE_RUNNER_CONFIGURATION_READER"
	ResourceRoleHostAuthenticationTokenAdmin   ResourceRole = "RESOURCE_ROLE_HOST_AUTHENTICATION_TOKEN_ADMIN"
	ResourceRoleHostAuthenticationTokenUpdater ResourceRole = "RESOURCE_ROLE_HOST_AUTHENTICATION_TOKEN_UPDATER"
	ResourceRoleProjectAdmin                   ResourceRole = "RESOURCE_ROLE_PROJECT_ADMIN"
	ResourceRoleProjectUser                    ResourceRole = "RESOURCE_ROLE_PROJECT_USER"
	ResourceRoleProjectEditor                  ResourceRole = "RESOURCE_ROLE_PROJECT_EDITOR"
	ResourceRoleEnvironmentServiceAdmin        ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_SERVICE_ADMIN"
	ResourceRoleEnvironmentServiceViewer       ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_SERVICE_VIEWER"
	ResourceRoleEnvironmentServiceUser         ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_SERVICE_USER"
	ResourceRoleEnvironmentServiceEnv          ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_SERVICE_ENV"
	ResourceRoleEnvironmentTaskAdmin           ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_TASK_ADMIN"
	ResourceRoleEnvironmentTaskViewer          ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_TASK_VIEWER"
	ResourceRoleEnvironmentTaskUser            ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_TASK_USER"
	ResourceRoleEnvironmentTaskEnv             ResourceRole = "RESOURCE_ROLE_ENVIRONMENT_TASK_ENV"
	ResourceRoleServiceAccountIdentity         ResourceRole = "RESOURCE_ROLE_SERVICE_ACCOUNT_IDENTITY"
	ResourceRoleServiceAccountAdmin            ResourceRole = "RESOURCE_ROLE_SERVICE_ACCOUNT_ADMIN"
	ResourceRoleAgentExecutionIdentity         ResourceRole = "RESOURCE_ROLE_AGENT_EXECUTION_IDENTITY"
	ResourceRoleAgentExecutionUser             ResourceRole = "RESOURCE_ROLE_AGENT_EXECUTION_USER"
	ResourceRoleAgentExecutionAdmin            ResourceRole = "RESOURCE_ROLE_AGENT_EXECUTION_ADMIN"
	ResourceRoleAgentExecutionRunner           ResourceRole = "RESOURCE_ROLE_AGENT_EXECUTION_RUNNER"
	ResourceRoleAgentExecutionOutputsReporter  ResourceRole = "RESOURCE_ROLE_AGENT_EXECUTION_OUTPUTS_REPORTER"
	ResourceRoleAgentAdmin                     ResourceRole = "RESOURCE_ROLE_AGENT_ADMIN"
	ResourceRoleAgentViewer                    ResourceRole = "RESOURCE_ROLE_AGENT_VIEWER"
	ResourceRoleAgentExecutor                  ResourceRole = "RESOURCE_ROLE_AGENT_EXECUTOR"
	ResourceRoleWorkflowAdmin                  ResourceRole = "RESOURCE_ROLE_WORKFLOW_ADMIN"
	ResourceRoleWorkflowUser                   ResourceRole = "RESOURCE_ROLE_WORKFLOW_USER"
	ResourceRoleWorkflowViewer                 ResourceRole = "RESOURCE_ROLE_WORKFLOW_VIEWER"
	ResourceRoleWorkflowExecutor               ResourceRole = "RESOURCE_ROLE_WORKFLOW_EXECUTOR"
	ResourceRoleSnapshotAdmin                  ResourceRole = "RESOURCE_ROLE_SNAPSHOT_ADMIN"
	ResourceRoleSnapshotRunner                 ResourceRole = "RESOURCE_ROLE_SNAPSHOT_RUNNER"
	ResourceRoleWebhookAdmin                   ResourceRole = "RESOURCE_ROLE_WEBHOOK_ADMIN"
	ResourceRoleWebhookViewer                  ResourceRole = "RESOURCE_ROLE_WEBHOOK_VIEWER"
)

func (r ResourceRole) IsKnown() bool {
	switch r {
	case ResourceRoleUnspecified, ResourceRoleOrgAdmin, ResourceRoleOrgMember, ResourceRoleOrgRunnersAdmin, ResourceRoleOrgProjectsAdmin, ResourceRoleOrgAutomationsAdmin, ResourceRoleOrgGroupsAdmin, ResourceRoleGroupAdmin, ResourceRoleGroupViewer, ResourceRoleUserIdentity, ResourceRoleUserViewer, ResourceRoleUserAdmin, ResourceRoleEnvironmentIdentity, ResourceRoleEnvironmentAdmin, ResourceRoleEnvironmentUser, ResourceRoleEnvironmentViewer, ResourceRoleEnvironmentRunner, ResourceRoleRunnerIdentity, ResourceRoleRunnerAdmin, ResourceRoleRunnerLocalAdmin, ResourceRoleRunnerManagedAdmin, ResourceRoleRunnerUser, ResourceRoleRunnerConfigurationReader, ResourceRoleHostAuthenticationTokenAdmin, ResourceRoleHostAuthenticationTokenUpdater, ResourceRoleProjectAdmin, ResourceRoleProjectUser, ResourceRoleProjectEditor, ResourceRoleEnvironmentServiceAdmin, ResourceRoleEnvironmentServiceViewer, ResourceRoleEnvironmentServiceUser, ResourceRoleEnvironmentServiceEnv, ResourceRoleEnvironmentTaskAdmin, ResourceRoleEnvironmentTaskViewer, ResourceRoleEnvironmentTaskUser, ResourceRoleEnvironmentTaskEnv, ResourceRoleServiceAccountIdentity, ResourceRoleServiceAccountAdmin, ResourceRoleAgentExecutionIdentity, ResourceRoleAgentExecutionUser, ResourceRoleAgentExecutionAdmin, ResourceRoleAgentExecutionRunner, ResourceRoleAgentExecutionOutputsReporter, ResourceRoleAgentAdmin, ResourceRoleAgentViewer, ResourceRoleAgentExecutor, ResourceRoleWorkflowAdmin, ResourceRoleWorkflowUser, ResourceRoleWorkflowViewer, ResourceRoleWorkflowExecutor, ResourceRoleSnapshotAdmin, ResourceRoleSnapshotRunner, ResourceRoleWebhookAdmin, ResourceRoleWebhookViewer:
		return true
	}
	return false
}

type ResourceType string

const (
	ResourceTypeUnspecified                ResourceType = "RESOURCE_TYPE_UNSPECIFIED"
	ResourceTypeEnvironment                ResourceType = "RESOURCE_TYPE_ENVIRONMENT"
	ResourceTypeRunner                     ResourceType = "RESOURCE_TYPE_RUNNER"
	ResourceTypeProject                    ResourceType = "RESOURCE_TYPE_PROJECT"
	ResourceTypeTask                       ResourceType = "RESOURCE_TYPE_TASK"
	ResourceTypeTaskExecution              ResourceType = "RESOURCE_TYPE_TASK_EXECUTION"
	ResourceTypeService                    ResourceType = "RESOURCE_TYPE_SERVICE"
	ResourceTypeOrganization               ResourceType = "RESOURCE_TYPE_ORGANIZATION"
	ResourceTypeUser                       ResourceType = "RESOURCE_TYPE_USER"
	ResourceTypeEnvironmentClass           ResourceType = "RESOURCE_TYPE_ENVIRONMENT_CLASS"
	ResourceTypeRunnerScmIntegration       ResourceType = "RESOURCE_TYPE_RUNNER_SCM_INTEGRATION"
	ResourceTypeHostAuthenticationToken    ResourceType = "RESOURCE_TYPE_HOST_AUTHENTICATION_TOKEN"
	ResourceTypeGroup                      ResourceType = "RESOURCE_TYPE_GROUP"
	ResourceTypePersonalAccessToken        ResourceType = "RESOURCE_TYPE_PERSONAL_ACCESS_TOKEN"
	ResourceTypeUserPreference             ResourceType = "RESOURCE_TYPE_USER_PREFERENCE"
	ResourceTypeServiceAccount             ResourceType = "RESOURCE_TYPE_SERVICE_ACCOUNT"
	ResourceTypeSecret                     ResourceType = "RESOURCE_TYPE_SECRET"
	ResourceTypeSSOConfig                  ResourceType = "RESOURCE_TYPE_SSO_CONFIG"
	ResourceTypeDomainVerification         ResourceType = "RESOURCE_TYPE_DOMAIN_VERIFICATION"
	ResourceTypeAgentExecution             ResourceType = "RESOURCE_TYPE_AGENT_EXECUTION"
	ResourceTypeRunnerLlmIntegration       ResourceType = "RESOURCE_TYPE_RUNNER_LLM_INTEGRATION"
	ResourceTypeAgent                      ResourceType = "RESOURCE_TYPE_AGENT"
	ResourceTypeEnvironmentSession         ResourceType = "RESOURCE_TYPE_ENVIRONMENT_SESSION"
	ResourceTypeUserSecret                 ResourceType = "RESOURCE_TYPE_USER_SECRET"
	ResourceTypeOrganizationPolicy         ResourceType = "RESOURCE_TYPE_ORGANIZATION_POLICY"
	ResourceTypeOrganizationSecret         ResourceType = "RESOURCE_TYPE_ORGANIZATION_SECRET"
	ResourceTypeProjectEnvironmentClass    ResourceType = "RESOURCE_TYPE_PROJECT_ENVIRONMENT_CLASS"
	ResourceTypeBilling                    ResourceType = "RESOURCE_TYPE_BILLING"
	ResourceTypePrompt                     ResourceType = "RESOURCE_TYPE_PROMPT"
	ResourceTypeCoupon                     ResourceType = "RESOURCE_TYPE_COUPON"
	ResourceTypeCouponRedemption           ResourceType = "RESOURCE_TYPE_COUPON_REDEMPTION"
	ResourceTypeAccount                    ResourceType = "RESOURCE_TYPE_ACCOUNT"
	ResourceTypeIntegration                ResourceType = "RESOURCE_TYPE_INTEGRATION"
	ResourceTypeWorkflow                   ResourceType = "RESOURCE_TYPE_WORKFLOW"
	ResourceTypeWorkflowExecution          ResourceType = "RESOURCE_TYPE_WORKFLOW_EXECUTION"
	ResourceTypeWorkflowExecutionAction    ResourceType = "RESOURCE_TYPE_WORKFLOW_EXECUTION_ACTION"
	ResourceTypeSnapshot                   ResourceType = "RESOURCE_TYPE_SNAPSHOT"
	ResourceTypePrebuild                   ResourceType = "RESOURCE_TYPE_PREBUILD"
	ResourceTypeOrganizationLlmIntegration ResourceType = "RESOURCE_TYPE_ORGANIZATION_LLM_INTEGRATION"
	ResourceTypeCustomDomain               ResourceType = "RESOURCE_TYPE_CUSTOM_DOMAIN"
	ResourceTypeRoleAssignmentChanged      ResourceType = "RESOURCE_TYPE_ROLE_ASSIGNMENT_CHANGED"
	ResourceTypeGroupMembershipChanged     ResourceType = "RESOURCE_TYPE_GROUP_MEMBERSHIP_CHANGED"
	ResourceTypeWebhook                    ResourceType = "RESOURCE_TYPE_WEBHOOK"
	ResourceTypeScimConfiguration          ResourceType = "RESOURCE_TYPE_SCIM_CONFIGURATION"
	ResourceTypeServiceAccountSecret       ResourceType = "RESOURCE_TYPE_SERVICE_ACCOUNT_SECRET"
	ResourceTypeAnnouncementBanner         ResourceType = "RESOURCE_TYPE_ANNOUNCEMENT_BANNER"
	ResourceTypeServiceAccountToken        ResourceType = "RESOURCE_TYPE_SERVICE_ACCOUNT_TOKEN"
	ResourceTypeRoleAssignment             ResourceType = "RESOURCE_TYPE_ROLE_ASSIGNMENT"
)

func (r ResourceType) IsKnown() bool {
	switch r {
	case ResourceTypeUnspecified, ResourceTypeEnvironment, ResourceTypeRunner, ResourceTypeProject, ResourceTypeTask, ResourceTypeTaskExecution, ResourceTypeService, ResourceTypeOrganization, ResourceTypeUser, ResourceTypeEnvironmentClass, ResourceTypeRunnerScmIntegration, ResourceTypeHostAuthenticationToken, ResourceTypeGroup, ResourceTypePersonalAccessToken, ResourceTypeUserPreference, ResourceTypeServiceAccount, ResourceTypeSecret, ResourceTypeSSOConfig, ResourceTypeDomainVerification, ResourceTypeAgentExecution, ResourceTypeRunnerLlmIntegration, ResourceTypeAgent, ResourceTypeEnvironmentSession, ResourceTypeUserSecret, ResourceTypeOrganizationPolicy, ResourceTypeOrganizationSecret, ResourceTypeProjectEnvironmentClass, ResourceTypeBilling, ResourceTypePrompt, ResourceTypeCoupon, ResourceTypeCouponRedemption, ResourceTypeAccount, ResourceTypeIntegration, ResourceTypeWorkflow, ResourceTypeWorkflowExecution, ResourceTypeWorkflowExecutionAction, ResourceTypeSnapshot, ResourceTypePrebuild, ResourceTypeOrganizationLlmIntegration, ResourceTypeCustomDomain, ResourceTypeRoleAssignmentChanged, ResourceTypeGroupMembershipChanged, ResourceTypeWebhook, ResourceTypeScimConfiguration, ResourceTypeServiceAccountSecret, ResourceTypeAnnouncementBanner, ResourceTypeServiceAccountToken, ResourceTypeRoleAssignment:
		return true
	}
	return false
}

type RunsOn struct {
	Docker RunsOnDocker `json:"docker"`
	// Machine runs the service/task directly on the VM/machine level.
	Machine interface{} `json:"machine"`
	JSON    runsOnJSON  `json:"-"`
}

// runsOnJSON contains the JSON metadata for the struct [RunsOn]
type runsOnJSON struct {
	Docker      apijson.Field
	Machine     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunsOn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runsOnJSON) RawJSON() string {
	return r.raw
}

type RunsOnDocker struct {
	Environment []string         `json:"environment"`
	Image       string           `json:"image"`
	JSON        runsOnDockerJSON `json:"-"`
}

// runsOnDockerJSON contains the JSON metadata for the struct [RunsOnDocker]
type runsOnDockerJSON struct {
	Environment apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunsOnDocker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runsOnDockerJSON) RawJSON() string {
	return r.raw
}

type RunsOnParam struct {
	Docker param.Field[RunsOnDockerParam] `json:"docker"`
	// Machine runs the service/task directly on the VM/machine level.
	Machine param.Field[interface{}] `json:"machine"`
}

func (r RunsOnParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunsOnDockerParam struct {
	Environment param.Field[[]string] `json:"environment"`
	Image       param.Field[string]   `json:"image"`
}

func (r RunsOnDockerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SecretRef references a secret by its ID.
type SecretRef struct {
	// id is the UUID of the secret to reference.
	ID   string        `json:"id" format:"uuid"`
	JSON secretRefJSON `json:"-"`
}

// secretRefJSON contains the JSON metadata for the struct [SecretRef]
type secretRefJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretRefJSON) RawJSON() string {
	return r.raw
}

// SecretRef references a secret by its ID.
type SecretRefParam struct {
	// id is the UUID of the secret to reference.
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r SecretRefParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Current state of the pull request
type State string

const (
	StateUnspecified State = "STATE_UNSPECIFIED"
	StateOpen        State = "STATE_OPEN"
	StateClosed      State = "STATE_CLOSED"
	StateMerged      State = "STATE_MERGED"
)

func (r State) IsKnown() bool {
	switch r {
	case StateUnspecified, StateOpen, StateClosed, StateMerged:
		return true
	}
	return false
}

type Subject struct {
	// id is the UUID of the subject
	ID string `json:"id" format:"uuid"`
	// Principal is the principal of the subject
	Principal Principal   `json:"principal"`
	JSON      subjectJSON `json:"-"`
}

// subjectJSON contains the JSON metadata for the struct [Subject]
type subjectJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Subject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subjectJSON) RawJSON() string {
	return r.raw
}

type SubjectParam struct {
	// id is the UUID of the subject
	ID param.Field[string] `json:"id" format:"uuid"`
	// Principal is the principal of the subject
	Principal param.Field[Principal] `json:"principal"`
}

func (r SubjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Task struct {
	ID string `json:"id,required" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn     []string     `json:"dependsOn" format:"uuid"`
	EnvironmentID string       `json:"environmentId" format:"uuid"`
	Metadata      TaskMetadata `json:"metadata"`
	Spec          TaskSpec     `json:"spec"`
	JSON          taskJSON     `json:"-"`
}

// taskJSON contains the JSON metadata for the struct [Task]
type taskJSON struct {
	ID            apijson.Field
	DependsOn     apijson.Field
	EnvironmentID apijson.Field
	Metadata      apijson.Field
	Spec          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Task) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskJSON) RawJSON() string {
	return r.raw
}

type TaskExecution struct {
	ID       string                `json:"id,required" format:"uuid"`
	Metadata TaskExecutionMetadata `json:"metadata"`
	Spec     TaskExecutionSpec     `json:"spec"`
	Status   TaskExecutionStatus   `json:"status"`
	JSON     taskExecutionJSON     `json:"-"`
}

// taskExecutionJSON contains the JSON metadata for the struct [TaskExecution]
type taskExecutionJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionMetadata struct {
	// completed_at is the time the task execution was done.
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// created_at is the time the task was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator describes the principal who created/started the task run.
	Creator Subject `json:"creator"`
	// environment_id is the ID of the environment in which the task run is executed.
	EnvironmentID string `json:"environmentId" format:"uuid"`
	// started_at is the time the task execution actually started to run.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// started_by describes the trigger that started the task execution.
	StartedBy string `json:"startedBy"`
	// task_id is the ID of the main task being executed.
	TaskID string                    `json:"taskId" format:"uuid"`
	JSON   taskExecutionMetadataJSON `json:"-"`
}

// taskExecutionMetadataJSON contains the JSON metadata for the struct
// [TaskExecutionMetadata]
type taskExecutionMetadataJSON struct {
	CompletedAt   apijson.Field
	CreatedAt     apijson.Field
	Creator       apijson.Field
	EnvironmentID apijson.Field
	StartedAt     apijson.Field
	StartedBy     apijson.Field
	TaskID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TaskExecutionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionMetadataJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionPhase string

const (
	TaskExecutionPhaseUnspecified TaskExecutionPhase = "TASK_EXECUTION_PHASE_UNSPECIFIED"
	TaskExecutionPhasePending     TaskExecutionPhase = "TASK_EXECUTION_PHASE_PENDING"
	TaskExecutionPhaseRunning     TaskExecutionPhase = "TASK_EXECUTION_PHASE_RUNNING"
	TaskExecutionPhaseSucceeded   TaskExecutionPhase = "TASK_EXECUTION_PHASE_SUCCEEDED"
	TaskExecutionPhaseFailed      TaskExecutionPhase = "TASK_EXECUTION_PHASE_FAILED"
	TaskExecutionPhaseStopped     TaskExecutionPhase = "TASK_EXECUTION_PHASE_STOPPED"
)

func (r TaskExecutionPhase) IsKnown() bool {
	switch r {
	case TaskExecutionPhaseUnspecified, TaskExecutionPhasePending, TaskExecutionPhaseRunning, TaskExecutionPhaseSucceeded, TaskExecutionPhaseFailed, TaskExecutionPhaseStopped:
		return true
	}
	return false
}

type TaskExecutionSpec struct {
	// desired_phase is the phase the task execution should be in. Used to stop a
	// running task execution early.
	DesiredPhase TaskExecutionPhase `json:"desiredPhase"`
	// plan is a list of groups of steps. The steps in a group are executed
	// concurrently, while the groups are executed sequentially. The order of the
	// groups is the order in which they are executed.
	Plan []TaskExecutionSpecPlan `json:"plan"`
	JSON taskExecutionSpecJSON   `json:"-"`
}

// taskExecutionSpecJSON contains the JSON metadata for the struct
// [TaskExecutionSpec]
type taskExecutionSpecJSON struct {
	DesiredPhase apijson.Field
	Plan         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TaskExecutionSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionSpecPlan struct {
	Steps []TaskExecutionSpecPlanStep `json:"steps"`
	JSON  taskExecutionSpecPlanJSON   `json:"-"`
}

// taskExecutionSpecPlanJSON contains the JSON metadata for the struct
// [TaskExecutionSpecPlan]
type taskExecutionSpecPlanJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecutionSpecPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecPlanJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionSpecPlanStep struct {
	// ID is the ID of the execution step
	ID        string                         `json:"id" format:"uuid"`
	DependsOn []string                       `json:"dependsOn"`
	Label     string                         `json:"label"`
	ServiceID string                         `json:"serviceId" format:"uuid"`
	Task      TaskExecutionSpecPlanStepsTask `json:"task"`
	JSON      taskExecutionSpecPlanStepJSON  `json:"-"`
}

// taskExecutionSpecPlanStepJSON contains the JSON metadata for the struct
// [TaskExecutionSpecPlanStep]
type taskExecutionSpecPlanStepJSON struct {
	ID          apijson.Field
	DependsOn   apijson.Field
	Label       apijson.Field
	ServiceID   apijson.Field
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecutionSpecPlanStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecPlanStepJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionSpecPlanStepsTask struct {
	ID   string                             `json:"id" format:"uuid"`
	Spec TaskSpec                           `json:"spec"`
	JSON taskExecutionSpecPlanStepsTaskJSON `json:"-"`
}

// taskExecutionSpecPlanStepsTaskJSON contains the JSON metadata for the struct
// [TaskExecutionSpecPlanStepsTask]
type taskExecutionSpecPlanStepsTaskJSON struct {
	ID          apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecutionSpecPlanStepsTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionSpecPlanStepsTaskJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionStatus struct {
	// failure_message summarises why the task execution failed to operate. If this is
	// non-empty the task execution has failed to operate and will likely transition to
	// a failed state.
	FailureMessage string `json:"failureMessage"`
	// log_url is the URL to the logs of the task's steps. If this is empty, the task
	// either has no logs or has not yet started.
	LogURL string `json:"logUrl"`
	// the phase of a task execution represents the aggregated phase of all steps.
	Phase TaskExecutionPhase `json:"phase"`
	// version of the status update. Task executions themselves are unversioned, but
	// their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion"`
	// steps provides the status for each individual step of the task execution. If a
	// step is missing it has not yet started.
	Steps []TaskExecutionStatusStep `json:"steps"`
	JSON  taskExecutionStatusJSON   `json:"-"`
}

// taskExecutionStatusJSON contains the JSON metadata for the struct
// [TaskExecutionStatus]
type taskExecutionStatusJSON struct {
	FailureMessage apijson.Field
	LogURL         apijson.Field
	Phase          apijson.Field
	StatusVersion  apijson.Field
	Steps          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TaskExecutionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionStatusJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionStatusStep struct {
	// ID is the ID of the execution step
	ID string `json:"id" format:"uuid"`
	// failure_message summarises why the step failed to operate. If this is non-empty
	// the step has failed to operate and will likely transition to a failed state.
	FailureMessage string `json:"failureMessage"`
	// output contains the output of the task execution. setting an output field to
	// empty string will unset it.
	Output map[string]string `json:"output"`
	// phase is the current phase of the execution step
	Phase TaskExecutionPhase          `json:"phase"`
	JSON  taskExecutionStatusStepJSON `json:"-"`
}

// taskExecutionStatusStepJSON contains the JSON metadata for the struct
// [TaskExecutionStatusStep]
type taskExecutionStatusStepJSON struct {
	ID             apijson.Field
	FailureMessage apijson.Field
	Output         apijson.Field
	Phase          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TaskExecutionStatusStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionStatusStepJSON) RawJSON() string {
	return r.raw
}

type TaskMetadata struct {
	// created_at is the time the task was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator describes the principal who created the task.
	Creator Subject `json:"creator"`
	// description is a user-facing description for the task. It can be used to provide
	// context and documentation for the task.
	Description string `json:"description"`
	// name is a user-facing name for the task. Unlike the reference, this field is not
	// unique, and not referenced by the system. This is a short descriptive name for
	// the task.
	Name string `json:"name"`
	// reference is a user-facing identifier for the task which must be unique on the
	// environment. It is used to express dependencies between tasks, and to identify
	// the task in user interactions (e.g. the CLI).
	Reference string `json:"reference"`
	// triggered_by is a list of trigger that start the task.
	TriggeredBy []AutomationTrigger `json:"triggeredBy"`
	JSON        taskMetadataJSON    `json:"-"`
}

// taskMetadataJSON contains the JSON metadata for the struct [TaskMetadata]
type taskMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Reference   apijson.Field
	TriggeredBy apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskMetadataJSON) RawJSON() string {
	return r.raw
}

type TaskMetadataParam struct {
	// created_at is the time the task was created.
	CreatedAt param.Field[time.Time] `json:"createdAt" format:"date-time"`
	// creator describes the principal who created the task.
	Creator param.Field[SubjectParam] `json:"creator"`
	// description is a user-facing description for the task. It can be used to provide
	// context and documentation for the task.
	Description param.Field[string] `json:"description"`
	// name is a user-facing name for the task. Unlike the reference, this field is not
	// unique, and not referenced by the system. This is a short descriptive name for
	// the task.
	Name param.Field[string] `json:"name"`
	// reference is a user-facing identifier for the task which must be unique on the
	// environment. It is used to express dependencies between tasks, and to identify
	// the task in user interactions (e.g. the CLI).
	Reference param.Field[string] `json:"reference"`
	// triggered_by is a list of trigger that start the task.
	TriggeredBy param.Field[[]AutomationTriggerParam] `json:"triggeredBy"`
}

func (r TaskMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TaskSpec struct {
	// command contains the command the task should execute
	Command string `json:"command"`
	// env specifies environment variables for the task.
	Env []EnvironmentVariableItem `json:"env"`
	// runs_on specifies the environment the task should run on.
	RunsOn RunsOn       `json:"runsOn"`
	JSON   taskSpecJSON `json:"-"`
}

// taskSpecJSON contains the JSON metadata for the struct [TaskSpec]
type taskSpecJSON struct {
	Command     apijson.Field
	Env         apijson.Field
	RunsOn      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskSpecJSON) RawJSON() string {
	return r.raw
}

type TaskSpecParam struct {
	// command contains the command the task should execute
	Command param.Field[string] `json:"command"`
	// env specifies environment variables for the task.
	Env param.Field[[]EnvironmentVariableItemParam] `json:"env"`
	// runs_on specifies the environment the task should run on.
	RunsOn param.Field[RunsOnParam] `json:"runsOn"`
}

func (r TaskSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserStatus string

const (
	UserStatusUnspecified UserStatus = "USER_STATUS_UNSPECIFIED"
	UserStatusActive      UserStatus = "USER_STATUS_ACTIVE"
	UserStatusSuspended   UserStatus = "USER_STATUS_SUSPENDED"
	UserStatusLeft        UserStatus = "USER_STATUS_LEFT"
)

func (r UserStatus) IsKnown() bool {
	switch r {
	case UserStatusUnspecified, UserStatusActive, UserStatusSuspended, UserStatusLeft:
		return true
	}
	return false
}
