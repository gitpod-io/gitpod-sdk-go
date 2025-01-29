// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
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

// GetEditor returns the editor with the given ID
func (r *EditorService) Get(ctx context.Context, params EditorGetParams, opts ...option.RequestOption) (res *EditorGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EditorService/GetEditor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListEditors lists all editors available to the caller
func (r *EditorService) List(ctx context.Context, params EditorListParams, opts ...option.RequestOption) (res *EditorListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EditorService/ListEditors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ResolveEditorURL resolves the editor's URL for an environment
func (r *EditorService) ResolveEditorURL(ctx context.Context, params EditorResolveEditorURLParams, opts ...option.RequestOption) (res *EditorResolveEditorURLResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EditorService/ResolveEditorURL"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type EditorGetResponse struct {
	// editor contains the editor
	Editor EditorGetResponseEditor `json:"editor"`
	JSON   editorGetResponseJSON   `json:"-"`
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

// editor contains the editor
type EditorGetResponseEditor struct {
	ID                       string                      `json:"id"`
	IconURL                  string                      `json:"iconUrl"`
	InstallationInstructions string                      `json:"installationInstructions"`
	Name                     string                      `json:"name"`
	ShortDescription         string                      `json:"shortDescription"`
	URLTemplate              string                      `json:"urlTemplate"`
	JSON                     editorGetResponseEditorJSON `json:"-"`
}

// editorGetResponseEditorJSON contains the JSON metadata for the struct
// [EditorGetResponseEditor]
type editorGetResponseEditorJSON struct {
	ID                       apijson.Field
	IconURL                  apijson.Field
	InstallationInstructions apijson.Field
	Name                     apijson.Field
	ShortDescription         apijson.Field
	URLTemplate              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EditorGetResponseEditor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editorGetResponseEditorJSON) RawJSON() string {
	return r.raw
}

type EditorListResponse struct {
	// editors contains the list of editors
	Editors []EditorListResponseEditor `json:"editors"`
	// pagination contains the pagination options for listing environments
	Pagination EditorListResponsePagination `json:"pagination"`
	JSON       editorListResponseJSON       `json:"-"`
}

// editorListResponseJSON contains the JSON metadata for the struct
// [EditorListResponse]
type editorListResponseJSON struct {
	Editors     apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editorListResponseJSON) RawJSON() string {
	return r.raw
}

type EditorListResponseEditor struct {
	ID                       string                       `json:"id"`
	IconURL                  string                       `json:"iconUrl"`
	InstallationInstructions string                       `json:"installationInstructions"`
	Name                     string                       `json:"name"`
	ShortDescription         string                       `json:"shortDescription"`
	URLTemplate              string                       `json:"urlTemplate"`
	JSON                     editorListResponseEditorJSON `json:"-"`
}

// editorListResponseEditorJSON contains the JSON metadata for the struct
// [EditorListResponseEditor]
type editorListResponseEditorJSON struct {
	ID                       apijson.Field
	IconURL                  apijson.Field
	InstallationInstructions apijson.Field
	Name                     apijson.Field
	ShortDescription         apijson.Field
	URLTemplate              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EditorListResponseEditor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editorListResponseEditorJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing environments
type EditorListResponsePagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token string `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize int64                            `json:"pageSize"`
	JSON     editorListResponsePaginationJSON `json:"-"`
}

// editorListResponsePaginationJSON contains the JSON metadata for the struct
// [EditorListResponsePagination]
type editorListResponsePaginationJSON struct {
	Token       apijson.Field
	PageSize    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditorListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editorListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type EditorResolveEditorURLResponse struct {
	// url is the resolved editor URL
	URL  string                             `json:"url"`
	JSON editorResolveEditorURLResponseJSON `json:"-"`
}

// editorResolveEditorURLResponseJSON contains the JSON metadata for the struct
// [EditorResolveEditorURLResponse]
type editorResolveEditorURLResponseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditorResolveEditorURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editorResolveEditorURLResponseJSON) RawJSON() string {
	return r.raw
}

type EditorGetParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EditorGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// id is the ID of the editor to get
	ID param.Field[string] `json:"id"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EditorGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EditorGetParamsConnectProtocolVersion float64

const (
	EditorGetParamsConnectProtocolVersion1 EditorGetParamsConnectProtocolVersion = 1
)

func (r EditorGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EditorGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type EditorListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EditorListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// pagination contains the pagination options for listing environments
	Pagination param.Field[EditorListParamsPagination] `json:"pagination"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EditorListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EditorListParamsConnectProtocolVersion float64

const (
	EditorListParamsConnectProtocolVersion1 EditorListParamsConnectProtocolVersion = 1
)

func (r EditorListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EditorListParamsConnectProtocolVersion1:
		return true
	}
	return false
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

type EditorResolveEditorURLParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EditorResolveEditorURLParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// editorId is the ID of the editor to resolve the URL for
	EditorID param.Field[string] `json:"editorId" format:"uuid"`
	// environmentId is the ID of the environment to resolve the URL for
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
	// organizationId is the ID of the organization to resolve the URL for
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EditorResolveEditorURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type EditorResolveEditorURLParamsConnectProtocolVersion float64

const (
	EditorResolveEditorURLParamsConnectProtocolVersion1 EditorResolveEditorURLParamsConnectProtocolVersion = 1
)

func (r EditorResolveEditorURLParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EditorResolveEditorURLParamsConnectProtocolVersion1:
		return true
	}
	return false
}
