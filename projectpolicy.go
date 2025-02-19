// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
)

// ProjectPolicyService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectPolicyService] method instead.
type ProjectPolicyService struct {
	Options []option.RequestOption
}

// NewProjectPolicyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectPolicyService(opts ...option.RequestOption) (r *ProjectPolicyService) {
	r = &ProjectPolicyService{}
	r.Options = opts
	return
}

// Creates a new policy for a project.
//
// Use this method to:
//
// - Set up access controls
// - Define group permissions
// - Configure role-based access
//
// ### Examples
//
// - Create admin policy:
//
//	Grants admin access to a group.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	groupId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	role: PROJECT_ROLE_ADMIN
//	```
func (r *ProjectPolicyService) New(ctx context.Context, body ProjectPolicyNewParams, opts ...option.RequestOption) (res *ProjectPolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/CreateProjectPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing project policy.
//
// Use this method to:
//
// - Modify access levels
// - Change group roles
// - Update permissions
//
// ### Examples
//
// - Update policy role:
//
//	Changes a group's access level.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	groupId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	role: PROJECT_ROLE_EDITOR
//	```
func (r *ProjectPolicyService) Update(ctx context.Context, body ProjectPolicyUpdateParams, opts ...option.RequestOption) (res *ProjectPolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/UpdateProjectPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists policies for a project.
//
// Use this method to:
//
// - View access controls
// - Check policy configurations
// - Audit permissions
//
// ### Examples
//
// - List policies:
//
//	Shows all policies for a project.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	```
func (r *ProjectPolicyService) List(ctx context.Context, params ProjectPolicyListParams, opts ...option.RequestOption) (res *pagination.PoliciesPage[ProjectPolicy], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.ProjectService/ListProjectPolicies"
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

// Lists policies for a project.
//
// Use this method to:
//
// - View access controls
// - Check policy configurations
// - Audit permissions
//
// ### Examples
//
// - List policies:
//
//	Shows all policies for a project.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	```
func (r *ProjectPolicyService) ListAutoPaging(ctx context.Context, params ProjectPolicyListParams, opts ...option.RequestOption) *pagination.PoliciesPageAutoPager[ProjectPolicy] {
	return pagination.NewPoliciesPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a project policy.
//
// Use this method to:
//
// - Remove access controls
// - Revoke permissions
// - Clean up policies
//
// ### Examples
//
// - Delete policy:
//
//	Removes a group's access policy.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	groupId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	```
func (r *ProjectPolicyService) Delete(ctx context.Context, body ProjectPolicyDeleteParams, opts ...option.RequestOption) (res *ProjectPolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/DeleteProjectPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProjectPolicy struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role ProjectRole       `json:"role"`
	JSON projectPolicyJSON `json:"-"`
}

// projectPolicyJSON contains the JSON metadata for the struct [ProjectPolicy]
type projectPolicyJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyJSON) RawJSON() string {
	return r.raw
}

type ProjectRole string

const (
	ProjectRoleUnspecified ProjectRole = "PROJECT_ROLE_UNSPECIFIED"
	ProjectRoleAdmin       ProjectRole = "PROJECT_ROLE_ADMIN"
	ProjectRoleUser        ProjectRole = "PROJECT_ROLE_USER"
	ProjectRoleEditor      ProjectRole = "PROJECT_ROLE_EDITOR"
)

func (r ProjectRole) IsKnown() bool {
	switch r {
	case ProjectRoleUnspecified, ProjectRoleAdmin, ProjectRoleUser, ProjectRoleEditor:
		return true
	}
	return false
}

type ProjectPolicyNewResponse struct {
	Policy ProjectPolicy                `json:"policy"`
	JSON   projectPolicyNewResponseJSON `json:"-"`
}

// projectPolicyNewResponseJSON contains the JSON metadata for the struct
// [ProjectPolicyNewResponse]
type projectPolicyNewResponseJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectPolicyUpdateResponse struct {
	Policy ProjectPolicy                   `json:"policy"`
	JSON   projectPolicyUpdateResponseJSON `json:"-"`
}

// projectPolicyUpdateResponseJSON contains the JSON metadata for the struct
// [ProjectPolicyUpdateResponse]
type projectPolicyUpdateResponseJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectPolicyDeleteResponse = interface{}

type ProjectPolicyNewParams struct {
	// group_id specifies the group_id identifier
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
	// project_id specifies the project identifier
	ProjectID param.Field[string]      `json:"projectId" format:"uuid"`
	Role      param.Field[ProjectRole] `json:"role"`
}

func (r ProjectPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectPolicyUpdateParams struct {
	// group_id specifies the group_id identifier
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
	// project_id specifies the project identifier
	ProjectID param.Field[string]      `json:"projectId" format:"uuid"`
	Role      param.Field[ProjectRole] `json:"role"`
}

func (r ProjectPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectPolicyListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// pagination contains the pagination options for listing project policies
	Pagination param.Field[ProjectPolicyListParamsPagination] `json:"pagination"`
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
}

func (r ProjectPolicyListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ProjectPolicyListParams]'s query parameters as
// `url.Values`.
func (r ProjectPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// pagination contains the pagination options for listing project policies
type ProjectPolicyListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r ProjectPolicyListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectPolicyDeleteParams struct {
	// group_id specifies the group_id identifier
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
}

func (r ProjectPolicyDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
