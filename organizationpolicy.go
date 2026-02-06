// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"slices"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
)

// OrganizationPolicyService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationPolicyService] method instead.
type OrganizationPolicyService struct {
	Options []option.RequestOption
}

// NewOrganizationPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationPolicyService(opts ...option.RequestOption) (r *OrganizationPolicyService) {
	r = &OrganizationPolicyService{}
	r.Options = opts
	return
}

// Gets organization policy settings by organization ID.
//
// Use this method to:
//
// - Retrieve current policy settings for an organization
// - View resource limits and restrictions
// - Check allowed editors and other configurations
//
// ### Examples
//
// - Get organization policies:
//
//	Retrieves policy settings for a specific organization.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *OrganizationPolicyService) Get(ctx context.Context, body OrganizationPolicyGetParams, opts ...option.RequestOption) (res *OrganizationPolicyGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/GetOrganizationPolicies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates organization policy settings.
//
// Use this method to:
//
// - Configure editor restrictions
// - Set environment resource limits
// - Define project creation permissions
// - Customize default configurations
//
// ### Examples
//
// - Update editor policies:
//
//	Restricts available editors and sets a default.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	allowedEditorIds:
//	  - "vscode"
//	  - "jetbrains"
//	defaultEditorId: "vscode"
//	```
//
// - Set environment limits:
//
//	Configures limits for environment usage.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	maximumEnvironmentTimeout: "3600s"
//	maximumRunningEnvironmentsPerUser: "5"
//	maximumEnvironmentsPerUser: "20"
//	```
func (r *OrganizationPolicyService) Update(ctx context.Context, body OrganizationPolicyUpdateParams, opts ...option.RequestOption) (res *OrganizationPolicyUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/UpdateOrganizationPolicies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// AgentPolicy contains agent-specific policy settings for an organization
type AgentPolicy struct {
	// command_deny_list contains a list of commands that agents are not allowed to
	// execute
	CommandDenyList []string `json:"commandDenyList,required"`
	// mcp_disabled controls whether MCP (Model Context Protocol) is disabled for
	// agents
	McpDisabled bool `json:"mcpDisabled,required"`
	// scm_tools_disabled controls whether SCM (Source Control Management) tools are
	// disabled for agents
	ScmToolsDisabled bool `json:"scmToolsDisabled,required"`
	// scm_tools_allowed_group_id restricts SCM tools access to members of this group.
	// Empty means no restriction (all users can use SCM tools if not disabled).
	ScmToolsAllowedGroupID string          `json:"scmToolsAllowedGroupId"`
	JSON                   agentPolicyJSON `json:"-"`
}

// agentPolicyJSON contains the JSON metadata for the struct [AgentPolicy]
type agentPolicyJSON struct {
	CommandDenyList        apijson.Field
	McpDisabled            apijson.Field
	ScmToolsDisabled       apijson.Field
	ScmToolsAllowedGroupID apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AgentPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentPolicyJSON) RawJSON() string {
	return r.raw
}

// CrowdStrikeConfig configures CrowdStrike Falcon sensor deployment
type CrowdStrikeConfig struct {
	// additional*options contains additional FALCONCTL_OPT*\* options as key-value
	// pairs. Keys should NOT include the FALCONCTL*OPT* prefix.
	AdditionalOptions map[string]string `json:"additionalOptions"`
	// cid_secret_id references an organization secret containing the Customer ID
	// (CID).
	CidSecretID string `json:"cidSecretId" format:"uuid"`
	// enabled controls whether CrowdStrike Falcon is deployed to environments
	Enabled bool `json:"enabled"`
	// image is the CrowdStrike Falcon sensor container image reference
	Image string `json:"image"`
	// tags are optional tags to apply to the Falcon sensor (comma-separated)
	Tags string                `json:"tags"`
	JSON crowdStrikeConfigJSON `json:"-"`
}

// crowdStrikeConfigJSON contains the JSON metadata for the struct
// [CrowdStrikeConfig]
type crowdStrikeConfigJSON struct {
	AdditionalOptions apijson.Field
	CidSecretID       apijson.Field
	Enabled           apijson.Field
	Image             apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CrowdStrikeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crowdStrikeConfigJSON) RawJSON() string {
	return r.raw
}

// ExecutableDenyList contains executables that are blocked from execution in
// environments.
type ExecutableDenyList struct {
	// enabled controls whether executable blocking is active
	Enabled bool `json:"enabled"`
	// executables is the list of executable paths or names to block
	Executables []string               `json:"executables"`
	JSON        executableDenyListJSON `json:"-"`
}

// executableDenyListJSON contains the JSON metadata for the struct
// [ExecutableDenyList]
type executableDenyListJSON struct {
	Enabled     apijson.Field
	Executables apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutableDenyList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executableDenyListJSON) RawJSON() string {
	return r.raw
}

// ExecutableDenyList contains executables that are blocked from execution in
// environments.
type ExecutableDenyListParam struct {
	// enabled controls whether executable blocking is active
	Enabled param.Field[bool] `json:"enabled"`
	// executables is the list of executable paths or names to block
	Executables param.Field[[]string] `json:"executables"`
}

func (r ExecutableDenyListParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationPolicies struct {
	// agent_policy contains agent-specific policy settings
	AgentPolicy AgentPolicy `json:"agentPolicy,required"`
	// allowed_editor_ids is the list of editor IDs that are allowed to be used in the
	// organization
	AllowedEditorIDs []string `json:"allowedEditorIds,required"`
	// allow_local_runners controls whether local runners are allowed to be used in the
	// organization
	AllowLocalRunners bool `json:"allowLocalRunners,required"`
	// default_editor_id is the default editor ID to be used when a user doesn't
	// specify one
	DefaultEditorID string `json:"defaultEditorId,required"`
	// default_environment_image is the default container image when none is defined in
	// repo
	DefaultEnvironmentImage string `json:"defaultEnvironmentImage,required"`
	// maximum_environments_per_user limits total environments (running or stopped) per
	// user
	MaximumEnvironmentsPerUser string `json:"maximumEnvironmentsPerUser,required"`
	// maximum_running_environments_per_user limits simultaneously running environments
	// per user
	MaximumRunningEnvironmentsPerUser string `json:"maximumRunningEnvironmentsPerUser,required"`
	// members_create_projects controls whether members can create projects
	MembersCreateProjects bool `json:"membersCreateProjects,required"`
	// members_require_projects controls whether environments can only be created from
	// projects by non-admin users
	MembersRequireProjects bool `json:"membersRequireProjects,required"`
	// organization_id is the ID of the organization
	OrganizationID string `json:"organizationId,required" format:"uuid"`
	// port_sharing_disabled controls whether port sharing is disabled in the
	// organization
	PortSharingDisabled bool `json:"portSharingDisabled,required"`
	// require_custom_domain_access controls whether users must access via custom
	// domain when one is configured. When true, access via app.gitpod.io is blocked.
	RequireCustomDomainAccess bool `json:"requireCustomDomainAccess,required"`
	// restrict_account_creation_to_scim controls whether account creation is
	// restricted to SCIM-provisioned users only. When true and SCIM is configured for
	// the organization, only users provisioned via SCIM can create accounts.
	RestrictAccountCreationToScim bool `json:"restrictAccountCreationToScim,required"`
	// delete_archived_environments_after controls how long archived environments are
	// kept before automatic deletion. 0 means no automatic deletion. Maximum duration
	// is 4 weeks (2419200 seconds).
	DeleteArchivedEnvironmentsAfter string `json:"deleteArchivedEnvironmentsAfter" format:"regex"`
	// editor_version_restrictions restricts which editor versions can be used. Maps
	// editor ID to version policy, editor_version_restrictions not set means no
	// restrictions. If empty or not set for an editor, we will use the latest version
	// of the editor
	EditorVersionRestrictions map[string]OrganizationPoliciesEditorVersionRestriction `json:"editorVersionRestrictions"`
	// executable_deny_list contains executables that are blocked from execution in
	// environments.
	ExecutableDenyList ExecutableDenyList `json:"executableDenyList"`
	// maximum_environment_lifetime controls for how long environments are allowed to
	// be reused. 0 means no maximum lifetime. Maximum duration is 180 days (15552000
	// seconds).
	MaximumEnvironmentLifetime string `json:"maximumEnvironmentLifetime" format:"regex"`
	// maximum_environment_timeout controls the maximum timeout allowed for
	// environments in seconds. 0 means no limit (never). Minimum duration is 30
	// minutes (1800 seconds).
	MaximumEnvironmentTimeout string `json:"maximumEnvironmentTimeout" format:"regex"`
	// security_agent_policy contains security agent configuration for the
	// organization. When configured, security agents are automatically deployed to all
	// environments.
	SecurityAgentPolicy SecurityAgentPolicy      `json:"securityAgentPolicy"`
	JSON                organizationPoliciesJSON `json:"-"`
}

// organizationPoliciesJSON contains the JSON metadata for the struct
// [OrganizationPolicies]
type organizationPoliciesJSON struct {
	AgentPolicy                       apijson.Field
	AllowedEditorIDs                  apijson.Field
	AllowLocalRunners                 apijson.Field
	DefaultEditorID                   apijson.Field
	DefaultEnvironmentImage           apijson.Field
	MaximumEnvironmentsPerUser        apijson.Field
	MaximumRunningEnvironmentsPerUser apijson.Field
	MembersCreateProjects             apijson.Field
	MembersRequireProjects            apijson.Field
	OrganizationID                    apijson.Field
	PortSharingDisabled               apijson.Field
	RequireCustomDomainAccess         apijson.Field
	RestrictAccountCreationToScim     apijson.Field
	DeleteArchivedEnvironmentsAfter   apijson.Field
	EditorVersionRestrictions         apijson.Field
	ExecutableDenyList                apijson.Field
	MaximumEnvironmentLifetime        apijson.Field
	MaximumEnvironmentTimeout         apijson.Field
	SecurityAgentPolicy               apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrganizationPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationPoliciesJSON) RawJSON() string {
	return r.raw
}

// EditorVersionPolicy defines the version policy for a specific editor
type OrganizationPoliciesEditorVersionRestriction struct {
	// allowed_versions lists the versions that are allowed If empty, we will use the
	// latest version of the editor
	//
	// Examples for JetBrains: `["2025.2", "2025.1", "2024.3"]`
	AllowedVersions []string                                         `json:"allowedVersions"`
	JSON            organizationPoliciesEditorVersionRestrictionJSON `json:"-"`
}

// organizationPoliciesEditorVersionRestrictionJSON contains the JSON metadata for
// the struct [OrganizationPoliciesEditorVersionRestriction]
type organizationPoliciesEditorVersionRestrictionJSON struct {
	AllowedVersions apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *OrganizationPoliciesEditorVersionRestriction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationPoliciesEditorVersionRestrictionJSON) RawJSON() string {
	return r.raw
}

// SecurityAgentPolicy contains security agent configuration for an organization.
// When enabled, security agents are automatically deployed to all environments.
type SecurityAgentPolicy struct {
	// crowdstrike contains CrowdStrike Falcon configuration
	Crowdstrike CrowdStrikeConfig       `json:"crowdstrike"`
	JSON        securityAgentPolicyJSON `json:"-"`
}

// securityAgentPolicyJSON contains the JSON metadata for the struct
// [SecurityAgentPolicy]
type securityAgentPolicyJSON struct {
	Crowdstrike apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecurityAgentPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityAgentPolicyJSON) RawJSON() string {
	return r.raw
}

type OrganizationPolicyGetResponse struct {
	Policies OrganizationPolicies              `json:"policies,required"`
	JSON     organizationPolicyGetResponseJSON `json:"-"`
}

// organizationPolicyGetResponseJSON contains the JSON metadata for the struct
// [OrganizationPolicyGetResponse]
type organizationPolicyGetResponseJSON struct {
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationPolicyUpdateResponse = interface{}

type OrganizationPolicyGetParams struct {
	// organization_id is the ID of the organization to retrieve policies for
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r OrganizationPolicyGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationPolicyUpdateParams struct {
	// organization_id is the ID of the organization to update policies for
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
	// agent_policy contains agent-specific policy settings
	AgentPolicy param.Field[OrganizationPolicyUpdateParamsAgentPolicy] `json:"agentPolicy"`
	// allowed_editor_ids is the list of editor IDs that are allowed to be used in the
	// organization
	AllowedEditorIDs param.Field[[]string] `json:"allowedEditorIds"`
	// allow_local_runners controls whether local runners are allowed to be used in the
	// organization
	AllowLocalRunners param.Field[bool] `json:"allowLocalRunners"`
	// default_editor_id is the default editor ID to be used when a user doesn't
	// specify one
	DefaultEditorID param.Field[string] `json:"defaultEditorId"`
	// default_environment_image is the default container image when none is defined in
	// repo
	DefaultEnvironmentImage param.Field[string] `json:"defaultEnvironmentImage"`
	// delete_archived_environments_after controls how long archived environments are
	// kept before automatic deletion. 0 means no automatic deletion. Maximum duration
	// is 4 weeks (2419200 seconds).
	DeleteArchivedEnvironmentsAfter param.Field[string] `json:"deleteArchivedEnvironmentsAfter" format:"regex"`
	// editor_version_restrictions restricts which editor versions can be used. Maps
	// editor ID to version policy with allowed major versions.
	EditorVersionRestrictions param.Field[map[string]OrganizationPolicyUpdateParamsEditorVersionRestrictions] `json:"editorVersionRestrictions"`
	// executable_deny_list contains executables that are blocked from execution in
	// environments.
	ExecutableDenyList param.Field[ExecutableDenyListParam] `json:"executableDenyList"`
	// maximum_environment_lifetime controls for how long environments are allowed to
	// be reused. 0 means no maximum lifetime. Maximum duration is 180 days (15552000
	// seconds).
	MaximumEnvironmentLifetime param.Field[string] `json:"maximumEnvironmentLifetime" format:"regex"`
	// maximum_environments_per_user limits total environments (running or stopped) per
	// user
	MaximumEnvironmentsPerUser param.Field[string] `json:"maximumEnvironmentsPerUser"`
	// maximum_environment_timeout controls the maximum timeout allowed for
	// environments in seconds. 0 means no limit (never). Minimum duration is 30
	// minutes (1800 seconds).
	MaximumEnvironmentTimeout param.Field[string] `json:"maximumEnvironmentTimeout" format:"regex"`
	// maximum_running_environments_per_user limits simultaneously running environments
	// per user
	MaximumRunningEnvironmentsPerUser param.Field[string] `json:"maximumRunningEnvironmentsPerUser"`
	// members_create_projects controls whether members can create projects
	MembersCreateProjects param.Field[bool] `json:"membersCreateProjects"`
	// members_require_projects controls whether environments can only be created from
	// projects by non-admin users
	MembersRequireProjects param.Field[bool] `json:"membersRequireProjects"`
	// port_sharing_disabled controls whether port sharing is disabled in the
	// organization
	PortSharingDisabled param.Field[bool] `json:"portSharingDisabled"`
	// require_custom_domain_access controls whether users must access via custom
	// domain when one is configured. When true, access via app.gitpod.io is blocked.
	RequireCustomDomainAccess param.Field[bool] `json:"requireCustomDomainAccess"`
	// restrict_account_creation_to_scim controls whether account creation is
	// restricted to SCIM-provisioned users only. When true and SCIM is configured for
	// the organization, only users provisioned via SCIM can create accounts.
	RestrictAccountCreationToScim param.Field[bool] `json:"restrictAccountCreationToScim"`
	// security_agent_policy contains security agent configuration updates
	SecurityAgentPolicy param.Field[OrganizationPolicyUpdateParamsSecurityAgentPolicy] `json:"securityAgentPolicy"`
}

func (r OrganizationPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// agent_policy contains agent-specific policy settings
type OrganizationPolicyUpdateParamsAgentPolicy struct {
	// command_deny_list contains a list of commands that agents are not allowed to
	// execute
	CommandDenyList param.Field[[]string] `json:"commandDenyList"`
	// mcp_disabled controls whether MCP (Model Context Protocol) is disabled for
	// agents
	McpDisabled param.Field[bool] `json:"mcpDisabled"`
	// scm_tools_allowed_group_id restricts SCM tools access to members of this group.
	// Empty means no restriction (all users can use SCM tools if not disabled).
	ScmToolsAllowedGroupID param.Field[string] `json:"scmToolsAllowedGroupId"`
	// scm_tools_disabled controls whether SCM (Source Control Management) tools are
	// disabled for agents
	ScmToolsDisabled param.Field[bool] `json:"scmToolsDisabled"`
}

func (r OrganizationPolicyUpdateParamsAgentPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EditorVersionPolicy defines the version policy for a specific editor
type OrganizationPolicyUpdateParamsEditorVersionRestrictions struct {
	// allowed_versions lists the versions that are allowed If empty, we will use the
	// latest version of the editor
	//
	// Examples for JetBrains: `["2025.2", "2025.1", "2024.3"]`
	AllowedVersions param.Field[[]string] `json:"allowedVersions"`
}

func (r OrganizationPolicyUpdateParamsEditorVersionRestrictions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// security_agent_policy contains security agent configuration updates
type OrganizationPolicyUpdateParamsSecurityAgentPolicy struct {
	// crowdstrike contains CrowdStrike Falcon configuration updates
	Crowdstrike param.Field[OrganizationPolicyUpdateParamsSecurityAgentPolicyCrowdstrike] `json:"crowdstrike"`
}

func (r OrganizationPolicyUpdateParamsSecurityAgentPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// crowdstrike contains CrowdStrike Falcon configuration updates
type OrganizationPolicyUpdateParamsSecurityAgentPolicyCrowdstrike struct {
	// additional*options contains additional FALCONCTL_OPT*\* options as key-value
	// pairs
	AdditionalOptions param.Field[map[string]string] `json:"additionalOptions"`
	// cid_secret_id references an organization secret containing the Customer ID (CID)
	CidSecretID param.Field[string] `json:"cidSecretId" format:"uuid"`
	// enabled controls whether CrowdStrike Falcon is deployed to environments
	Enabled param.Field[bool] `json:"enabled"`
	// image is the CrowdStrike Falcon sensor container image reference
	Image param.Field[string] `json:"image"`
	// tags are optional tags to apply to the Falcon sensor
	Tags param.Field[string] `json:"tags"`
}

func (r OrganizationPolicyUpdateParamsSecurityAgentPolicyCrowdstrike) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
