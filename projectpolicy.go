// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gitpod-io/flex-sdk-go/internal/apijson"
	"github.com/gitpod-io/flex-sdk-go/internal/apiquery"
	"github.com/gitpod-io/flex-sdk-go/internal/param"
	"github.com/gitpod-io/flex-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/flex-sdk-go/option"
	"github.com/gitpod-io/flex-sdk-go/packages/pagination"
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

// CreateProjectPolicy creates a Project Policy.
func (r *ProjectPolicyService) New(ctx context.Context, body ProjectPolicyNewParams, opts ...option.RequestOption) (res *ProjectPolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/CreateProjectPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateProjectPolicy updates a Project Policy.
func (r *ProjectPolicyService) Update(ctx context.Context, body ProjectPolicyUpdateParams, opts ...option.RequestOption) (res *ProjectPolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/UpdateProjectPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListProjectPolicies lists policies for a project.
func (r *ProjectPolicyService) List(ctx context.Context, params ProjectPolicyListParams, opts ...option.RequestOption) (res *pagination.PoliciesPage[ProjectPolicyListResponse], err error) {
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

// ListProjectPolicies lists policies for a project.
func (r *ProjectPolicyService) ListAutoPaging(ctx context.Context, params ProjectPolicyListParams, opts ...option.RequestOption) *pagination.PoliciesPageAutoPager[ProjectPolicyListResponse] {
	return pagination.NewPoliciesPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteProjectPolicy deletes a Project Policy.
func (r *ProjectPolicyService) Delete(ctx context.Context, body ProjectPolicyDeleteParams, opts ...option.RequestOption) (res *ProjectPolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/DeleteProjectPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProjectPolicyNewResponse struct {
	Policy ProjectPolicyNewResponsePolicy `json:"policy"`
	JSON   projectPolicyNewResponseJSON   `json:"-"`
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

type ProjectPolicyNewResponsePolicy struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role ProjectPolicyNewResponsePolicyRole `json:"role"`
	JSON projectPolicyNewResponsePolicyJSON `json:"-"`
}

// projectPolicyNewResponsePolicyJSON contains the JSON metadata for the struct
// [ProjectPolicyNewResponsePolicy]
type projectPolicyNewResponsePolicyJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyNewResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyNewResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// role is the role assigned to the group
type ProjectPolicyNewResponsePolicyRole string

const (
	ProjectPolicyNewResponsePolicyRoleProjectRoleUnspecified ProjectPolicyNewResponsePolicyRole = "PROJECT_ROLE_UNSPECIFIED"
	ProjectPolicyNewResponsePolicyRoleProjectRoleAdmin       ProjectPolicyNewResponsePolicyRole = "PROJECT_ROLE_ADMIN"
	ProjectPolicyNewResponsePolicyRoleProjectRoleUser        ProjectPolicyNewResponsePolicyRole = "PROJECT_ROLE_USER"
)

func (r ProjectPolicyNewResponsePolicyRole) IsKnown() bool {
	switch r {
	case ProjectPolicyNewResponsePolicyRoleProjectRoleUnspecified, ProjectPolicyNewResponsePolicyRoleProjectRoleAdmin, ProjectPolicyNewResponsePolicyRoleProjectRoleUser:
		return true
	}
	return false
}

type ProjectPolicyUpdateResponse struct {
	Policy ProjectPolicyUpdateResponsePolicy `json:"policy"`
	JSON   projectPolicyUpdateResponseJSON   `json:"-"`
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

type ProjectPolicyUpdateResponsePolicy struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role ProjectPolicyUpdateResponsePolicyRole `json:"role"`
	JSON projectPolicyUpdateResponsePolicyJSON `json:"-"`
}

// projectPolicyUpdateResponsePolicyJSON contains the JSON metadata for the struct
// [ProjectPolicyUpdateResponsePolicy]
type projectPolicyUpdateResponsePolicyJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyUpdateResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyUpdateResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// role is the role assigned to the group
type ProjectPolicyUpdateResponsePolicyRole string

const (
	ProjectPolicyUpdateResponsePolicyRoleProjectRoleUnspecified ProjectPolicyUpdateResponsePolicyRole = "PROJECT_ROLE_UNSPECIFIED"
	ProjectPolicyUpdateResponsePolicyRoleProjectRoleAdmin       ProjectPolicyUpdateResponsePolicyRole = "PROJECT_ROLE_ADMIN"
	ProjectPolicyUpdateResponsePolicyRoleProjectRoleUser        ProjectPolicyUpdateResponsePolicyRole = "PROJECT_ROLE_USER"
)

func (r ProjectPolicyUpdateResponsePolicyRole) IsKnown() bool {
	switch r {
	case ProjectPolicyUpdateResponsePolicyRoleProjectRoleUnspecified, ProjectPolicyUpdateResponsePolicyRoleProjectRoleAdmin, ProjectPolicyUpdateResponsePolicyRoleProjectRoleUser:
		return true
	}
	return false
}

type ProjectPolicyListResponse struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role ProjectPolicyListResponseRole `json:"role"`
	JSON projectPolicyListResponseJSON `json:"-"`
}

// projectPolicyListResponseJSON contains the JSON metadata for the struct
// [ProjectPolicyListResponse]
type projectPolicyListResponseJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

// role is the role assigned to the group
type ProjectPolicyListResponseRole string

const (
	ProjectPolicyListResponseRoleProjectRoleUnspecified ProjectPolicyListResponseRole = "PROJECT_ROLE_UNSPECIFIED"
	ProjectPolicyListResponseRoleProjectRoleAdmin       ProjectPolicyListResponseRole = "PROJECT_ROLE_ADMIN"
	ProjectPolicyListResponseRoleProjectRoleUser        ProjectPolicyListResponseRole = "PROJECT_ROLE_USER"
)

func (r ProjectPolicyListResponseRole) IsKnown() bool {
	switch r {
	case ProjectPolicyListResponseRoleProjectRoleUnspecified, ProjectPolicyListResponseRoleProjectRoleAdmin, ProjectPolicyListResponseRoleProjectRoleUser:
		return true
	}
	return false
}

type ProjectPolicyDeleteResponse = interface{}

type ProjectPolicyNewParams struct {
	// group_id specifies the group_id identifier
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
	// project_id specifies the project identifier
	ProjectID param.Field[string]                     `json:"projectId" format:"uuid"`
	Role      param.Field[ProjectPolicyNewParamsRole] `json:"role"`
}

func (r ProjectPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectPolicyNewParamsRole string

const (
	ProjectPolicyNewParamsRoleProjectRoleUnspecified ProjectPolicyNewParamsRole = "PROJECT_ROLE_UNSPECIFIED"
	ProjectPolicyNewParamsRoleProjectRoleAdmin       ProjectPolicyNewParamsRole = "PROJECT_ROLE_ADMIN"
	ProjectPolicyNewParamsRoleProjectRoleUser        ProjectPolicyNewParamsRole = "PROJECT_ROLE_USER"
)

func (r ProjectPolicyNewParamsRole) IsKnown() bool {
	switch r {
	case ProjectPolicyNewParamsRoleProjectRoleUnspecified, ProjectPolicyNewParamsRoleProjectRoleAdmin, ProjectPolicyNewParamsRoleProjectRoleUser:
		return true
	}
	return false
}

type ProjectPolicyUpdateParams struct {
	// group_id specifies the group_id identifier
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
	// project_id specifies the project identifier
	ProjectID param.Field[string]                        `json:"projectId" format:"uuid"`
	Role      param.Field[ProjectPolicyUpdateParamsRole] `json:"role"`
}

func (r ProjectPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectPolicyUpdateParamsRole string

const (
	ProjectPolicyUpdateParamsRoleProjectRoleUnspecified ProjectPolicyUpdateParamsRole = "PROJECT_ROLE_UNSPECIFIED"
	ProjectPolicyUpdateParamsRoleProjectRoleAdmin       ProjectPolicyUpdateParamsRole = "PROJECT_ROLE_ADMIN"
	ProjectPolicyUpdateParamsRoleProjectRoleUser        ProjectPolicyUpdateParamsRole = "PROJECT_ROLE_USER"
)

func (r ProjectPolicyUpdateParamsRole) IsKnown() bool {
	switch r {
	case ProjectPolicyUpdateParamsRoleProjectRoleUnspecified, ProjectPolicyUpdateParamsRoleProjectRoleAdmin, ProjectPolicyUpdateParamsRoleProjectRoleUser:
		return true
	}
	return false
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
