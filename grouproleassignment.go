// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// GroupRoleAssignmentService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupRoleAssignmentService] method instead.
type GroupRoleAssignmentService struct {
	Options []option.RequestOption
}

// NewGroupRoleAssignmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGroupRoleAssignmentService(opts ...option.RequestOption) (r *GroupRoleAssignmentService) {
	r = &GroupRoleAssignmentService{}
	r.Options = opts
	return
}

// Creates a role assignment for a group on a resource.
//
// Use this method to:
//
// - Assign specific roles to groups on runners, projects, or environments
// - Grant group-based access to resources
//
// ### Examples
//
// - Assign admin role on a runner:
//
//	Grants the group admin access to a runner.
//
//	```yaml
//	groupId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	resourceType: RESOURCE_TYPE_RUNNER
//	resourceId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	resourceRole: RESOURCE_ROLE_RUNNER_ADMIN
//	```
//
// - Assign user role on a project:
//
//	Grants the group user access to a project.
//
//	```yaml
//	groupId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	resourceType: RESOURCE_TYPE_PROJECT
//	resourceId: "a1b2c3d4-5678-90ab-cdef-1234567890ab"
//	resourceRole: RESOURCE_ROLE_PROJECT_USER
//	```
//
// ### Authorization
//
// Requires admin role on the specific resource.
func (r *GroupRoleAssignmentService) New(ctx context.Context, body GroupRoleAssignmentNewParams, opts ...option.RequestOption) (res *GroupRoleAssignmentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.GroupService/CreateRoleAssignment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists role assignments for a group or resource.
//
// Use this method to:
//
// - View all role assignments for a group
// - Audit resource access
// - Check which groups have access to resources
//
// ### Examples
//
// - List role assignments for a group:
//
//	Shows all role assignments for a specific group.
//
//	```yaml
//	filter:
//	  groupId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	pagination:
//	  pageSize: 20
//	```
//
// - List role assignments by resource type:
//
//	Shows all role assignments for runners.
//
//	```yaml
//	filter:
//	  resourceTypes:
//	    - RESOURCE_TYPE_RUNNER
//	pagination:
//	  pageSize: 20
//	```
//
// ### Authorization
//
// All organization members can view role assignments (transparency model).
func (r *GroupRoleAssignmentService) List(ctx context.Context, params GroupRoleAssignmentListParams, opts ...option.RequestOption) (res *pagination.AssignmentsPage[RoleAssignment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.GroupService/ListRoleAssignments"
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

// Lists role assignments for a group or resource.
//
// Use this method to:
//
// - View all role assignments for a group
// - Audit resource access
// - Check which groups have access to resources
//
// ### Examples
//
// - List role assignments for a group:
//
//	Shows all role assignments for a specific group.
//
//	```yaml
//	filter:
//	  groupId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	pagination:
//	  pageSize: 20
//	```
//
// - List role assignments by resource type:
//
//	Shows all role assignments for runners.
//
//	```yaml
//	filter:
//	  resourceTypes:
//	    - RESOURCE_TYPE_RUNNER
//	pagination:
//	  pageSize: 20
//	```
//
// ### Authorization
//
// All organization members can view role assignments (transparency model).
func (r *GroupRoleAssignmentService) ListAutoPaging(ctx context.Context, params GroupRoleAssignmentListParams, opts ...option.RequestOption) *pagination.AssignmentsPageAutoPager[RoleAssignment] {
	return pagination.NewAssignmentsPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a role assignment.
//
// Use this method to:
//
// - Remove group access to resources
// - Revoke role-based permissions
//
// ### Examples
//
// - Delete a role assignment:
//
//	Removes a role assignment by its ID.
//
//	```yaml
//	assignmentId: "a1b2c3d4-5678-90ab-cdef-1234567890ab"
//	```
//
// ### Authorization
//
// Requires admin role on the specific resource.
func (r *GroupRoleAssignmentService) Delete(ctx context.Context, body GroupRoleAssignmentDeleteParams, opts ...option.RequestOption) (res *GroupRoleAssignmentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.GroupService/DeleteRoleAssignment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ResourceRole represents roles that can be assigned to groups on resources These
// map directly to the roles defined in backend/db/rule/rbac/role/role.go
type ResourceRole string

const (
	ResourceRoleUnspecified                    ResourceRole = "RESOURCE_ROLE_UNSPECIFIED"
	ResourceRoleOrgAdmin                       ResourceRole = "RESOURCE_ROLE_ORG_ADMIN"
	ResourceRoleOrgMember                      ResourceRole = "RESOURCE_ROLE_ORG_MEMBER"
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
)

func (r ResourceRole) IsKnown() bool {
	switch r {
	case ResourceRoleUnspecified, ResourceRoleOrgAdmin, ResourceRoleOrgMember, ResourceRoleGroupAdmin, ResourceRoleGroupViewer, ResourceRoleUserIdentity, ResourceRoleUserViewer, ResourceRoleUserAdmin, ResourceRoleEnvironmentIdentity, ResourceRoleEnvironmentAdmin, ResourceRoleEnvironmentUser, ResourceRoleEnvironmentViewer, ResourceRoleEnvironmentRunner, ResourceRoleRunnerIdentity, ResourceRoleRunnerAdmin, ResourceRoleRunnerLocalAdmin, ResourceRoleRunnerManagedAdmin, ResourceRoleRunnerUser, ResourceRoleRunnerConfigurationReader, ResourceRoleHostAuthenticationTokenAdmin, ResourceRoleHostAuthenticationTokenUpdater, ResourceRoleProjectAdmin, ResourceRoleProjectUser, ResourceRoleProjectEditor, ResourceRoleEnvironmentServiceAdmin, ResourceRoleEnvironmentServiceViewer, ResourceRoleEnvironmentServiceUser, ResourceRoleEnvironmentServiceEnv, ResourceRoleEnvironmentTaskAdmin, ResourceRoleEnvironmentTaskViewer, ResourceRoleEnvironmentTaskUser, ResourceRoleEnvironmentTaskEnv, ResourceRoleServiceAccountIdentity, ResourceRoleServiceAccountAdmin, ResourceRoleAgentExecutionIdentity, ResourceRoleAgentExecutionUser, ResourceRoleAgentExecutionAdmin, ResourceRoleAgentExecutionRunner, ResourceRoleAgentExecutionOutputsReporter, ResourceRoleAgentAdmin, ResourceRoleAgentViewer, ResourceRoleAgentExecutor, ResourceRoleWorkflowAdmin, ResourceRoleWorkflowUser, ResourceRoleWorkflowViewer, ResourceRoleWorkflowExecutor, ResourceRoleSnapshotAdmin, ResourceRoleSnapshotRunner:
		return true
	}
	return false
}

// RoleAssignment represents a role assigned to a group on a specific resource
type RoleAssignment struct {
	// Unique identifier for the role assignment
	ID string `json:"id" format:"uuid"`
	// Group identifier
	GroupID string `json:"groupId" format:"uuid"`
	// Organization identifier
	OrganizationID string `json:"organizationId" format:"uuid"`
	// Resource identifier
	ResourceID string `json:"resourceId" format:"uuid"`
	// Role assigned to the group on this resource
	ResourceRole ResourceRole `json:"resourceRole"`
	// Type of resource (runner, project, environment, etc.)
	ResourceType shared.ResourceType `json:"resourceType"`
	JSON         roleAssignmentJSON  `json:"-"`
}

// roleAssignmentJSON contains the JSON metadata for the struct [RoleAssignment]
type roleAssignmentJSON struct {
	ID             apijson.Field
	GroupID        apijson.Field
	OrganizationID apijson.Field
	ResourceID     apijson.Field
	ResourceRole   apijson.Field
	ResourceType   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RoleAssignment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleAssignmentJSON) RawJSON() string {
	return r.raw
}

type GroupRoleAssignmentNewResponse struct {
	// RoleAssignment represents a role assigned to a group on a specific resource
	Assignment RoleAssignment                     `json:"assignment"`
	JSON       groupRoleAssignmentNewResponseJSON `json:"-"`
}

// groupRoleAssignmentNewResponseJSON contains the JSON metadata for the struct
// [GroupRoleAssignmentNewResponse]
type groupRoleAssignmentNewResponseJSON struct {
	Assignment  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GroupRoleAssignmentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupRoleAssignmentNewResponseJSON) RawJSON() string {
	return r.raw
}

type GroupRoleAssignmentDeleteResponse = interface{}

type GroupRoleAssignmentNewParams struct {
	GroupID    param.Field[string] `json:"groupId" format:"uuid"`
	ResourceID param.Field[string] `json:"resourceId" format:"uuid"`
	// ResourceRole represents roles that can be assigned to groups on resources These
	// map directly to the roles defined in backend/db/rule/rbac/role/role.go
	ResourceRole param.Field[ResourceRole]        `json:"resourceRole"`
	ResourceType param.Field[shared.ResourceType] `json:"resourceType"`
}

func (r GroupRoleAssignmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GroupRoleAssignmentListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// Filter parameters
	Filter param.Field[GroupRoleAssignmentListParamsFilter] `json:"filter"`
	// Pagination parameters
	Pagination param.Field[GroupRoleAssignmentListParamsPagination] `json:"pagination"`
}

func (r GroupRoleAssignmentListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [GroupRoleAssignmentListParams]'s query parameters as
// `url.Values`.
func (r GroupRoleAssignmentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter parameters
type GroupRoleAssignmentListParamsFilter struct {
	// group_id filters the response to only role assignments for this specific group
	// Empty string is allowed and means no filtering by group
	GroupID param.Field[string] `json:"groupId"`
	// resource_roles filters the response to only role assignments with these specific
	// roles
	ResourceRoles param.Field[[]ResourceRole] `json:"resourceRoles"`
	// resource_types filters the response to only role assignments for these resource
	// types
	ResourceTypes param.Field[[]shared.ResourceType] `json:"resourceTypes"`
	// user_id filters the response to only role assignments for groups that this user
	// is a member of Empty string is allowed and means no filtering by user
	UserID param.Field[string] `json:"userId"`
}

func (r GroupRoleAssignmentListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Pagination parameters
type GroupRoleAssignmentListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r GroupRoleAssignmentListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GroupRoleAssignmentDeleteParams struct {
	AssignmentID param.Field[string] `json:"assignmentId" format:"uuid"`
}

func (r GroupRoleAssignmentDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
