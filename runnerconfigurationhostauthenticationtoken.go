// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// RunnerConfigurationHostAuthenticationTokenService contains methods and other
// services that help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerConfigurationHostAuthenticationTokenService] method instead.
type RunnerConfigurationHostAuthenticationTokenService struct {
	Options []option.RequestOption
}

// NewRunnerConfigurationHostAuthenticationTokenService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewRunnerConfigurationHostAuthenticationTokenService(opts ...option.RequestOption) (r *RunnerConfigurationHostAuthenticationTokenService) {
	r = &RunnerConfigurationHostAuthenticationTokenService{}
	r.Options = opts
	return
}

// GetHostAuthenticationToken
func (r *RunnerConfigurationHostAuthenticationTokenService) New(ctx context.Context, params RunnerConfigurationHostAuthenticationTokenNewParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetHostAuthenticationToken
func (r *RunnerConfigurationHostAuthenticationTokenService) Get(ctx context.Context, params RunnerConfigurationHostAuthenticationTokenGetParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// UpdateHostAuthenticationToken
func (r *RunnerConfigurationHostAuthenticationTokenService) Update(ctx context.Context, params RunnerConfigurationHostAuthenticationTokenUpdateParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/UpdateHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListHostAuthenticationTokens
func (r *RunnerConfigurationHostAuthenticationTokenService) List(ctx context.Context, params RunnerConfigurationHostAuthenticationTokenListParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/ListHostAuthenticationTokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// DeleteHostAuthenticationToken
func (r *RunnerConfigurationHostAuthenticationTokenService) Delete(ctx context.Context, params RunnerConfigurationHostAuthenticationTokenDeleteParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/DeleteHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type RunnerConfigurationHostAuthenticationTokenNewResponse struct {
	Token RunnerConfigurationHostAuthenticationTokenNewResponseToken `json:"token"`
	JSON  runnerConfigurationHostAuthenticationTokenNewResponseJSON  `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenNewResponseJSON contains the JSON
// metadata for the struct [RunnerConfigurationHostAuthenticationTokenNewResponse]
type runnerConfigurationHostAuthenticationTokenNewResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenNewResponseToken struct {
	ID       string                                                           `json:"id"`
	Host     string                                                           `json:"host"`
	RunnerID string                                                           `json:"runnerId"`
	Source   RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource `json:"source"`
	UserID   string                                                           `json:"userId"`
	JSON     runnerConfigurationHostAuthenticationTokenNewResponseTokenJSON   `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenNewResponseTokenJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationHostAuthenticationTokenNewResponseToken]
type runnerConfigurationHostAuthenticationTokenNewResponseTokenJSON struct {
	ID          apijson.Field
	Host        apijson.Field
	RunnerID    apijson.Field
	Source      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenNewResponseToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenNewResponseTokenJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource string

const (
	RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourceUnspecified RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_UNSPECIFIED"
	RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourceOAuth       RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_OAUTH"
	RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourcePat         RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_PAT"
)

func (r RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourceUnspecified, RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourceOAuth, RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourcePat:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenGetResponse struct {
	Token RunnerConfigurationHostAuthenticationTokenGetResponseToken `json:"token"`
	JSON  runnerConfigurationHostAuthenticationTokenGetResponseJSON  `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenGetResponseJSON contains the JSON
// metadata for the struct [RunnerConfigurationHostAuthenticationTokenGetResponse]
type runnerConfigurationHostAuthenticationTokenGetResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenGetResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenGetResponseToken struct {
	ID       string                                                           `json:"id"`
	Host     string                                                           `json:"host"`
	RunnerID string                                                           `json:"runnerId"`
	Source   RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource `json:"source"`
	UserID   string                                                           `json:"userId"`
	JSON     runnerConfigurationHostAuthenticationTokenGetResponseTokenJSON   `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenGetResponseTokenJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationHostAuthenticationTokenGetResponseToken]
type runnerConfigurationHostAuthenticationTokenGetResponseTokenJSON struct {
	ID          apijson.Field
	Host        apijson.Field
	RunnerID    apijson.Field
	Source      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenGetResponseToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenGetResponseTokenJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource string

const (
	RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourceUnspecified RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_UNSPECIFIED"
	RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourceOAuth       RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_OAUTH"
	RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourcePat         RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_PAT"
)

func (r RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourceUnspecified, RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourceOAuth, RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourcePat:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenUpdateResponse = interface{}

type RunnerConfigurationHostAuthenticationTokenListResponse struct {
	Pagination RunnerConfigurationHostAuthenticationTokenListResponsePagination `json:"pagination"`
	Tokens     []RunnerConfigurationHostAuthenticationTokenListResponseToken    `json:"tokens"`
	JSON       runnerConfigurationHostAuthenticationTokenListResponseJSON       `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenListResponseJSON contains the JSON
// metadata for the struct [RunnerConfigurationHostAuthenticationTokenListResponse]
type runnerConfigurationHostAuthenticationTokenListResponseJSON struct {
	Pagination  apijson.Field
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenListResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                                               `json:"nextToken"`
	JSON      runnerConfigurationHostAuthenticationTokenListResponsePaginationJSON `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenListResponsePaginationJSON contains
// the JSON metadata for the struct
// [RunnerConfigurationHostAuthenticationTokenListResponsePagination]
type runnerConfigurationHostAuthenticationTokenListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenListResponseToken struct {
	ID       string                                                             `json:"id"`
	Host     string                                                             `json:"host"`
	RunnerID string                                                             `json:"runnerId"`
	Source   RunnerConfigurationHostAuthenticationTokenListResponseTokensSource `json:"source"`
	UserID   string                                                             `json:"userId"`
	JSON     runnerConfigurationHostAuthenticationTokenListResponseTokenJSON    `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenListResponseTokenJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationHostAuthenticationTokenListResponseToken]
type runnerConfigurationHostAuthenticationTokenListResponseTokenJSON struct {
	ID          apijson.Field
	Host        apijson.Field
	RunnerID    apijson.Field
	Source      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenListResponseToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenListResponseTokenJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenListResponseTokensSource string

const (
	RunnerConfigurationHostAuthenticationTokenListResponseTokensSourceHostAuthenticationTokenSourceUnspecified RunnerConfigurationHostAuthenticationTokenListResponseTokensSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_UNSPECIFIED"
	RunnerConfigurationHostAuthenticationTokenListResponseTokensSourceHostAuthenticationTokenSourceOAuth       RunnerConfigurationHostAuthenticationTokenListResponseTokensSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_OAUTH"
	RunnerConfigurationHostAuthenticationTokenListResponseTokensSourceHostAuthenticationTokenSourcePat         RunnerConfigurationHostAuthenticationTokenListResponseTokensSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_PAT"
)

func (r RunnerConfigurationHostAuthenticationTokenListResponseTokensSource) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenListResponseTokensSourceHostAuthenticationTokenSourceUnspecified, RunnerConfigurationHostAuthenticationTokenListResponseTokensSourceHostAuthenticationTokenSourceOAuth, RunnerConfigurationHostAuthenticationTokenListResponseTokensSourceHostAuthenticationTokenSourcePat:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenDeleteResponse = interface{}

type RunnerConfigurationHostAuthenticationTokenNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationHostAuthenticationTokenNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                                    `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationHostAuthenticationTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationHostAuthenticationTokenNewParamsConnectProtocolVersion float64

const (
	RunnerConfigurationHostAuthenticationTokenNewParamsConnectProtocolVersion1 RunnerConfigurationHostAuthenticationTokenNewParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationHostAuthenticationTokenNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenGetParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationHostAuthenticationTokenGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Base64                 param.Field[string]                                                                    `query:"base64"`
	Compression            param.Field[string]                                                                    `query:"compression"`
	Connect                param.Field[string]                                                                    `query:"connect"`
	Encoding               param.Field[string]                                                                    `query:"encoding"`
	Message                param.Field[string]                                                                    `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [RunnerConfigurationHostAuthenticationTokenGetParams]'s
// query parameters as `url.Values`.
func (r RunnerConfigurationHostAuthenticationTokenGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define the version of the Connect protocol
type RunnerConfigurationHostAuthenticationTokenGetParamsConnectProtocolVersion float64

const (
	RunnerConfigurationHostAuthenticationTokenGetParamsConnectProtocolVersion1 RunnerConfigurationHostAuthenticationTokenGetParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationHostAuthenticationTokenGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenUpdateParams struct {
	Body RunnerConfigurationHostAuthenticationTokenUpdateParamsBody `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationHostAuthenticationTokenUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationHostAuthenticationTokenUpdateParamsBody struct {
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationHostAuthenticationTokenUpdateParamsConnectProtocolVersion float64

const (
	RunnerConfigurationHostAuthenticationTokenUpdateParamsConnectProtocolVersion1 RunnerConfigurationHostAuthenticationTokenUpdateParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationHostAuthenticationTokenUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationHostAuthenticationTokenListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Filter                 param.Field[RunnerConfigurationHostAuthenticationTokenListParamsFilter]                 `json:"filter"`
	Pagination             param.Field[RunnerConfigurationHostAuthenticationTokenListParamsPagination]             `json:"pagination"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationHostAuthenticationTokenListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationHostAuthenticationTokenListParamsConnectProtocolVersion float64

const (
	RunnerConfigurationHostAuthenticationTokenListParamsConnectProtocolVersion1 RunnerConfigurationHostAuthenticationTokenListParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationHostAuthenticationTokenListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenListParamsFilter struct {
}

func (r RunnerConfigurationHostAuthenticationTokenListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationHostAuthenticationTokenListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerConfigurationHostAuthenticationTokenListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationHostAuthenticationTokenDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationHostAuthenticationTokenDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                                       `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationHostAuthenticationTokenDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationHostAuthenticationTokenDeleteParamsConnectProtocolVersion float64

const (
	RunnerConfigurationHostAuthenticationTokenDeleteParamsConnectProtocolVersion1 RunnerConfigurationHostAuthenticationTokenDeleteParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationHostAuthenticationTokenDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}
