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
	"github.com/stainless-sdks/gitpod-go/packages/pagination"
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

// CreateEnvironmentClass creates a new environment class on a runner.
func (r *RunnerConfigurationEnvironmentClassService) New(ctx context.Context, body RunnerConfigurationEnvironmentClassNewParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/CreateEnvironmentClass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetEnvironmentClass returns a single environment class configured for a runner.
func (r *RunnerConfigurationEnvironmentClassService) Get(ctx context.Context, body RunnerConfigurationEnvironmentClassGetParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetEnvironmentClass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateEnvironmentClass updates an existing environment class on a runner.
func (r *RunnerConfigurationEnvironmentClassService) Update(ctx context.Context, body RunnerConfigurationEnvironmentClassUpdateParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/UpdateEnvironmentClass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListEnvironmentClasses returns all environment classes configured for a runner.
// buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE
func (r *RunnerConfigurationEnvironmentClassService) List(ctx context.Context, params RunnerConfigurationEnvironmentClassListParams, opts ...option.RequestOption) (res *pagination.EnvironmentClassesPage[RunnerConfigurationEnvironmentClassListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.RunnerConfigurationService/ListEnvironmentClasses"
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

// ListEnvironmentClasses returns all environment classes configured for a runner.
// buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE
func (r *RunnerConfigurationEnvironmentClassService) ListAutoPaging(ctx context.Context, params RunnerConfigurationEnvironmentClassListParams, opts ...option.RequestOption) *pagination.EnvironmentClassesPageAutoPager[RunnerConfigurationEnvironmentClassListResponse] {
	return pagination.NewEnvironmentClassesPageAutoPager(r.List(ctx, params, opts...))
}

type RunnerConfigurationEnvironmentClassNewResponse struct {
	ID   string                                             `json:"id"`
	JSON runnerConfigurationEnvironmentClassNewResponseJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassNewResponseJSON contains the JSON metadata
// for the struct [RunnerConfigurationEnvironmentClassNewResponse]
type runnerConfigurationEnvironmentClassNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassNewResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassGetResponse struct {
	EnvironmentClass RunnerConfigurationEnvironmentClassGetResponseEnvironmentClass `json:"environmentClass"`
	JSON             runnerConfigurationEnvironmentClassGetResponseJSON             `json:"-"`
}

// runnerConfigurationEnvironmentClassGetResponseJSON contains the JSON metadata
// for the struct [RunnerConfigurationEnvironmentClassGetResponse]
type runnerConfigurationEnvironmentClassGetResponseJSON struct {
	EnvironmentClass apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassGetResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassGetResponseEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID string `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration []RunnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfiguration `json:"configuration"`
	// description is a human readable description of the environment class
	Description string `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName string `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled bool `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID string                                                             `json:"runnerId"`
	JSON     runnerConfigurationEnvironmentClassGetResponseEnvironmentClassJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassGetResponseEnvironmentClassJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationEnvironmentClassGetResponseEnvironmentClass]
type runnerConfigurationEnvironmentClassGetResponseEnvironmentClassJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	RunnerID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassGetResponseEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassGetResponseEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfiguration struct {
	Key   string                                                                          `json:"key"`
	Value string                                                                          `json:"value"`
	JSON  runnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfigurationJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfigurationJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfiguration]
type runnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfigurationJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfigurationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassUpdateResponse = interface{}

type RunnerConfigurationEnvironmentClassListResponse struct {
	// id is the unique identifier of the environment class
	ID string `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration []RunnerConfigurationEnvironmentClassListResponseConfiguration `json:"configuration"`
	// description is a human readable description of the environment class
	Description string `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName string `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled bool `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID string                                              `json:"runnerId"`
	JSON     runnerConfigurationEnvironmentClassListResponseJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassListResponseJSON contains the JSON metadata
// for the struct [RunnerConfigurationEnvironmentClassListResponse]
type runnerConfigurationEnvironmentClassListResponseJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	RunnerID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassListResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassListResponseConfiguration struct {
	Key   string                                                           `json:"key"`
	Value string                                                           `json:"value"`
	JSON  runnerConfigurationEnvironmentClassListResponseConfigurationJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassListResponseConfigurationJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationEnvironmentClassListResponseConfiguration]
type runnerConfigurationEnvironmentClassListResponseConfigurationJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassListResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassListResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassNewParams struct {
	Configuration param.Field[[]RunnerConfigurationEnvironmentClassNewParamsConfiguration] `json:"configuration"`
	Description   param.Field[string]                                                      `json:"description"`
	DisplayName   param.Field[string]                                                      `json:"displayName"`
	RunnerID      param.Field[string]                                                      `json:"runnerId" format:"uuid"`
}

func (r RunnerConfigurationEnvironmentClassNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationEnvironmentClassNewParamsConfiguration struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r RunnerConfigurationEnvironmentClassNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationEnvironmentClassGetParams struct {
	EnvironmentClassID param.Field[string] `json:"environmentClassId" format:"uuid"`
}

func (r RunnerConfigurationEnvironmentClassGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationEnvironmentClassUpdateParams struct {
	Body RunnerConfigurationEnvironmentClassUpdateParamsBodyUnion `json:"body,required"`
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

type RunnerConfigurationEnvironmentClassListParams struct {
	Token    param.Field[string]                                              `query:"token"`
	PageSize param.Field[int64]                                               `query:"pageSize"`
	Filter   param.Field[RunnerConfigurationEnvironmentClassListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing environment classes
	Pagination param.Field[RunnerConfigurationEnvironmentClassListParamsPagination] `json:"pagination"`
}

func (r RunnerConfigurationEnvironmentClassListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [RunnerConfigurationEnvironmentClassListParams]'s query
// parameters as `url.Values`.
func (r RunnerConfigurationEnvironmentClassListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RunnerConfigurationEnvironmentClassListParamsFilter struct {
	// enabled filters the response to only enabled or disabled environment classes. If
	// not set, all environment classes are returned.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RunnerConfigurationEnvironmentClassListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing environment classes
type RunnerConfigurationEnvironmentClassListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerConfigurationEnvironmentClassListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
