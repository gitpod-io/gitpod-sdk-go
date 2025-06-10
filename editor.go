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

// EditorService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEditorService] method instead.
type EditorService struct {
	Options []option.RequestOption
}

// NewEditorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEditorService(opts ...option.RequestOption) (r *EditorService) {
	r = &EditorService{}
	r.Options = opts
	return
}

// Gets details about a specific editor.
//
// Use this method to:
//
// - View editor information
// - Get editor configuration
//
// ### Examples
//
// - Get editor details:
//
//	Retrieves information about a specific editor.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *EditorService) Get(ctx context.Context, body EditorGetParams, opts ...option.RequestOption) (res *EditorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EditorService/GetEditor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all available code editors, optionally filtered to those allowed in an
// organization.
//
// Use this method to:
//
// - View supported editors
// - Get editor capabilities
// - Browse editor options
// - Check editor availability
//
// ### Examples
//
// - List editors:
//
//	Shows all available editors with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - List editors available to the organization:
//
//	Shows all available editors that are allowed by the policies enforced in the
//	organization with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	filter:
//	  allowedByPolicy: true
//	```
func (r *EditorService) List(ctx context.Context, params EditorListParams, opts ...option.RequestOption) (res *pagination.EditorsPage[Editor], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.EditorService/ListEditors"
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

// Lists all available code editors, optionally filtered to those allowed in an
// organization.
//
// Use this method to:
//
// - View supported editors
// - Get editor capabilities
// - Browse editor options
// - Check editor availability
//
// ### Examples
//
// - List editors:
//
//	Shows all available editors with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - List editors available to the organization:
//
//	Shows all available editors that are allowed by the policies enforced in the
//	organization with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	filter:
//	  allowedByPolicy: true
//	```
func (r *EditorService) ListAutoPaging(ctx context.Context, params EditorListParams, opts ...option.RequestOption) *pagination.EditorsPageAutoPager[Editor] {
	return pagination.NewEditorsPageAutoPager(r.List(ctx, params, opts...))
}

// Resolves the URL for accessing an editor in a specific environment.
//
// Use this method to:
//
// - Get editor access URLs
// - Launch editors for environments
// - Set up editor connections
// - Configure editor access
//
// ### Examples
//
// - Resolve editor URL:
//
//	Gets the URL for accessing an editor in an environment.
//
//	```yaml
//	editorId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *EditorService) ResolveURL(ctx context.Context, body EditorResolveURLParams, opts ...option.RequestOption) (res *EditorResolveURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EditorService/ResolveEditorURL"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Editor struct {
	ID                       string     `json:"id,required" format:"uuid"`
	InstallationInstructions string     `json:"installationInstructions,required"`
	Name                     string     `json:"name,required"`
	URLTemplate              string     `json:"urlTemplate,required"`
	Alias                    string     `json:"alias"`
	IconURL                  string     `json:"iconUrl"`
	ShortDescription         string     `json:"shortDescription"`
	JSON                     editorJSON `json:"-"`
}

// editorJSON contains the JSON metadata for the struct [Editor]
type editorJSON struct {
	ID                       apijson.Field
	InstallationInstructions apijson.Field
	Name                     apijson.Field
	URLTemplate              apijson.Field
	Alias                    apijson.Field
	IconURL                  apijson.Field
	ShortDescription         apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *Editor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editorJSON) RawJSON() string {
	return r.raw
}

type EditorGetResponse struct {
	// editor contains the editor
	Editor Editor                `json:"editor,required"`
	JSON   editorGetResponseJSON `json:"-"`
}

// editorGetResponseJSON contains the JSON metadata for the struct
// [EditorGetResponse]
type editorGetResponseJSON struct {
	Editor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editorGetResponseJSON) RawJSON() string {
	return r.raw
}

type EditorResolveURLResponse struct {
	// url is the resolved editor URL
	URL  string                       `json:"url,required"`
	JSON editorResolveURLResponseJSON `json:"-"`
}

// editorResolveURLResponseJSON contains the JSON metadata for the struct
// [EditorResolveURLResponse]
type editorResolveURLResponseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditorResolveURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editorResolveURLResponseJSON) RawJSON() string {
	return r.raw
}

type EditorGetParams struct {
	// id is the ID of the editor to get
	ID param.Field[string] `json:"id,required"`
}

func (r EditorGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EditorListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// filter contains the filter options for listing editors
	Filter param.Field[EditorListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing environments
	Pagination param.Field[EditorListParamsPagination] `json:"pagination"`
}

func (r EditorListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [EditorListParams]'s query parameters as `url.Values`.
func (r EditorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// filter contains the filter options for listing editors
type EditorListParamsFilter struct {
	// allowed_by_policy filters the response to only editors that are allowed by the
	// policies enforced in the organization
	AllowedByPolicy param.Field[bool] `json:"allowedByPolicy"`
}

func (r EditorListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing environments
type EditorListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EditorListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EditorResolveURLParams struct {
	// editorId is the ID of the editor to resolve the URL for
	EditorID param.Field[string] `json:"editorId,required" format:"uuid"`
	// environmentId is the ID of the environment to resolve the URL for
	EnvironmentID param.Field[string] `json:"environmentId,required" format:"uuid"`
	// organizationId is the ID of the organization to resolve the URL for
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r EditorResolveURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
