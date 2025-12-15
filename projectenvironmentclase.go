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

// ProjectEnvironmentClaseService contains methods and other services that help
// with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectEnvironmentClaseService] method instead.
type ProjectEnvironmentClaseService struct {
	Options []option.RequestOption
}

// NewProjectEnvironmentClaseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProjectEnvironmentClaseService(opts ...option.RequestOption) (r *ProjectEnvironmentClaseService) {
	r = &ProjectEnvironmentClaseService{}
	r.Options = opts
	return
}

// Updates all environment classes of a project.
//
// Use this method to:
//
// - Modify all environment classea of a project
//
// ### Examples
//
// - Update project environment classes:
//
//	Updates all environment classes for a project.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	projectEnvironmentClasses:
//	  - environmentClassId: "b0e12f6c-4c67-429d-a4a6-d9838b5da041"
//	    order: 0
//	  - localRunner: true
//	    order: 1
//	```
func (r *ProjectEnvironmentClaseService) Update(ctx context.Context, body ProjectEnvironmentClaseUpdateParams, opts ...option.RequestOption) (res *ProjectEnvironmentClaseUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.ProjectService/UpdateProjectEnvironmentClasses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists environment classes of a project.
//
// Use this method to:
//
// - View all environment classes of a project
//
// ### Examples
//
// - List project environment classes:
//
//	Shows all environment classes of a project.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	```
func (r *ProjectEnvironmentClaseService) List(ctx context.Context, params ProjectEnvironmentClaseListParams, opts ...option.RequestOption) (res *pagination.ProjectEnvironmentClassesPage[shared.ProjectEnvironmentClass], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.ProjectService/ListProjectEnvironmentClasses"
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

// Lists environment classes of a project.
//
// Use this method to:
//
// - View all environment classes of a project
//
// ### Examples
//
// - List project environment classes:
//
//	Shows all environment classes of a project.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	```
func (r *ProjectEnvironmentClaseService) ListAutoPaging(ctx context.Context, params ProjectEnvironmentClaseListParams, opts ...option.RequestOption) *pagination.ProjectEnvironmentClassesPageAutoPager[shared.ProjectEnvironmentClass] {
	return pagination.NewProjectEnvironmentClassesPageAutoPager(r.List(ctx, params, opts...))
}

type ProjectEnvironmentClaseUpdateResponse = interface{}

type ProjectEnvironmentClaseUpdateParams struct {
	ProjectEnvironmentClasses param.Field[[]shared.ProjectEnvironmentClassParam] `json:"projectEnvironmentClasses"`
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
}

func (r ProjectEnvironmentClaseUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectEnvironmentClaseListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// pagination contains the pagination options for listing project policies
	Pagination param.Field[ProjectEnvironmentClaseListParamsPagination] `json:"pagination"`
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
}

func (r ProjectEnvironmentClaseListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ProjectEnvironmentClaseListParams]'s query parameters as
// `url.Values`.
func (r ProjectEnvironmentClaseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// pagination contains the pagination options for listing project policies
type ProjectEnvironmentClaseListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r ProjectEnvironmentClaseListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
