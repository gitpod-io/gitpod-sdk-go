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

// RunnerConfigurationEnvironmentClassService contains methods and other services
// that help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerConfigurationEnvironmentClassService] method instead.
type RunnerConfigurationEnvironmentClassService struct {
	Options []option.RequestOption
}

// NewRunnerConfigurationEnvironmentClassService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRunnerConfigurationEnvironmentClassService(opts ...option.RequestOption) (r *RunnerConfigurationEnvironmentClassService) {
	r = &RunnerConfigurationEnvironmentClassService{}
	r.Options = opts
	return
}

// UpdateEnvironmentClass updates an existing environment class on a runner.
func (r *RunnerConfigurationEnvironmentClassService) Update(ctx context.Context, params RunnerConfigurationEnvironmentClassUpdateParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/UpdateEnvironmentClass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListEnvironmentClasses returns all environment classes configured for a runner.
//
// buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE
func (r *RunnerConfigurationEnvironmentClassService) List(ctx context.Context, params RunnerConfigurationEnvironmentClassListParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/ListEnvironmentClasses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type RunnerConfigurationEnvironmentClassUpdateResponse = interface{}

type RunnerConfigurationEnvironmentClassListResponse struct {
	EnvironmentClasses []RunnerConfigurationEnvironmentClassListResponseEnvironmentClass `json:"environmentClasses"`
	// pagination contains the pagination options for listing environment classes
	Pagination RunnerConfigurationEnvironmentClassListResponsePagination `json:"pagination"`
	JSON       runnerConfigurationEnvironmentClassListResponseJSON       `json:"-"`
}

// runnerConfigurationEnvironmentClassListResponseJSON contains the JSON metadata
// for the struct [RunnerConfigurationEnvironmentClassListResponse]
type runnerConfigurationEnvironmentClassListResponseJSON struct {
	EnvironmentClasses apijson.Field
	Pagination         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassListResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassListResponseEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID string `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration []RunnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfiguration `json:"configuration"`
	// description is a human readable description of the environment class
	Description string `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName string `json:"displayName"`
	// enabled indicates whether the environment class can be used to create
	//
	// new environments.
	Enabled bool `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID string                                                              `json:"runnerId"`
	JSON     runnerConfigurationEnvironmentClassListResponseEnvironmentClassJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassListResponseEnvironmentClassJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationEnvironmentClassListResponseEnvironmentClass]
type runnerConfigurationEnvironmentClassListResponseEnvironmentClassJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	RunnerID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassListResponseEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassListResponseEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfiguration struct {
	Key   string                                                                             `json:"key"`
	Value string                                                                             `json:"value"`
	JSON  runnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfigurationJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfigurationJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfiguration]
type runnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfigurationJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfigurationJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing environment classes
type RunnerConfigurationEnvironmentClassListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                                        `json:"nextToken"`
	JSON      runnerConfigurationEnvironmentClassListResponsePaginationJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassListResponsePaginationJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationEnvironmentClassListResponsePagination]
type runnerConfigurationEnvironmentClassListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassUpdateParams struct {
	Body RunnerConfigurationEnvironmentClassUpdateParamsBodyUnion `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationEnvironmentClassUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationEnvironmentClassUpdateParamsBody struct {
	Description param.Field[string] `json:"description"`
	DisplayName param.Field[string] `json:"displayName"`
	Enabled     param.Field[bool]   `json:"enabled"`
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBody) implementsRunnerConfigurationEnvironmentClassUpdateParamsBodyUnion() {
}

// Satisfied by [RunnerConfigurationEnvironmentClassUpdateParamsBodyDescription],
// [RunnerConfigurationEnvironmentClassUpdateParamsBodyDisplayName],
// [RunnerConfigurationEnvironmentClassUpdateParamsBodyEnabled],
// [RunnerConfigurationEnvironmentClassUpdateParamsBody].
type RunnerConfigurationEnvironmentClassUpdateParamsBodyUnion interface {
	implementsRunnerConfigurationEnvironmentClassUpdateParamsBodyUnion()
}

type RunnerConfigurationEnvironmentClassUpdateParamsBodyDescription struct {
	Description param.Field[string] `json:"description,required"`
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyDescription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyDescription) implementsRunnerConfigurationEnvironmentClassUpdateParamsBodyUnion() {
}

type RunnerConfigurationEnvironmentClassUpdateParamsBodyDisplayName struct {
	DisplayName param.Field[string] `json:"displayName,required"`
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyDisplayName) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyDisplayName) implementsRunnerConfigurationEnvironmentClassUpdateParamsBodyUnion() {
}

type RunnerConfigurationEnvironmentClassUpdateParamsBodyEnabled struct {
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyEnabled) implementsRunnerConfigurationEnvironmentClassUpdateParamsBodyUnion() {
}

// Define the version of the Connect protocol
type RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion float64

const (
	RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion1 RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationEnvironmentClassListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Filter                 param.Field[RunnerConfigurationEnvironmentClassListParamsFilter]                 `json:"filter"`
	// pagination contains the pagination options for listing environment classes
	Pagination param.Field[RunnerConfigurationEnvironmentClassListParamsPagination] `json:"pagination"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationEnvironmentClassListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion float64

const (
	RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion1 RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationEnvironmentClassListParamsFilter struct {
	// enabled filters the response to only enabled or disabled environment classes.
	//
	// If not set, all environment classes are returned.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RunnerConfigurationEnvironmentClassListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing environment classes
type RunnerConfigurationEnvironmentClassListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	//
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerConfigurationEnvironmentClassListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
