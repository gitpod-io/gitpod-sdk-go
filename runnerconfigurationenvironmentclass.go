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

// Creates a new environment class for a runner.
//
// Use this method to:
//
// - Define compute resources
// - Configure environment settings
// - Set up runtime options
//
// ### Examples
//
// - Create environment class:
//
//	Creates a new environment configuration.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	displayName: "Large Instance"
//	description: "8 CPU, 16GB RAM"
//	configuration:
//	  - key: "cpu"
//	    value: "8"
//	  - key: "memory"
//	    value: "16384"
//	```
func (r *RunnerConfigurationEnvironmentClassService) New(ctx context.Context, body RunnerConfigurationEnvironmentClassNewParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerConfigurationService/CreateEnvironmentClass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets details about a specific environment class.
//
// Use this method to:
//
// - View class configuration
// - Check resource settings
// - Verify availability
//
// ### Examples
//
// - Get class details:
//
//	Retrieves information about a specific class.
//
//	```yaml
//	environmentClassId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *RunnerConfigurationEnvironmentClassService) Get(ctx context.Context, body RunnerConfigurationEnvironmentClassGetParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerConfigurationService/GetEnvironmentClass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an environment class.
//
// Use this method to:
//
// - Modify class settings
// - Update resource limits
// - Change availability
//
// ### Examples
//
// - Update class:
//
//	Changes class configuration.
//
//	```yaml
//	environmentClassId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	displayName: "Updated Large Instance"
//	description: "16 CPU, 32GB RAM"
//	enabled: true
//	```
func (r *RunnerConfigurationEnvironmentClassService) Update(ctx context.Context, body RunnerConfigurationEnvironmentClassUpdateParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerConfigurationService/UpdateEnvironmentClass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists environment classes with optional filtering.
//
// Use this method to:
//
// - View available classes
// - Filter by capability
// - Check enabled status
//
// ### Examples
//
// - List all classes:
//
//	Shows all environment classes.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter enabled classes:
//
//	Lists only enabled environment classes.
//
//	```yaml
//	filter:
//	  enabled: true
//	pagination:
//	  pageSize: 20
//	```
//
//	buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE
func (r *RunnerConfigurationEnvironmentClassService) List(ctx context.Context, params RunnerConfigurationEnvironmentClassListParams, opts ...option.RequestOption) (res *pagination.EnvironmentClassesPage[shared.EnvironmentClass], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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

// Lists environment classes with optional filtering.
//
// Use this method to:
//
// - View available classes
// - Filter by capability
// - Check enabled status
//
// ### Examples
//
// - List all classes:
//
//	Shows all environment classes.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter enabled classes:
//
//	Lists only enabled environment classes.
//
//	```yaml
//	filter:
//	  enabled: true
//	pagination:
//	  pageSize: 20
//	```
//
//	buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE
func (r *RunnerConfigurationEnvironmentClassService) ListAutoPaging(ctx context.Context, params RunnerConfigurationEnvironmentClassListParams, opts ...option.RequestOption) *pagination.EnvironmentClassesPageAutoPager[shared.EnvironmentClass] {
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
	EnvironmentClass shared.EnvironmentClass                            `json:"environmentClass"`
	JSON             runnerConfigurationEnvironmentClassGetResponseJSON `json:"-"`
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

type RunnerConfigurationEnvironmentClassUpdateResponse = interface{}

type RunnerConfigurationEnvironmentClassNewParams struct {
	Configuration param.Field[[]shared.FieldValueParam] `json:"configuration"`
	Description   param.Field[string]                   `json:"description"`
	DisplayName   param.Field[string]                   `json:"displayName"`
	RunnerID      param.Field[string]                   `json:"runnerId" format:"uuid"`
}

func (r RunnerConfigurationEnvironmentClassNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationEnvironmentClassGetParams struct {
	EnvironmentClassID param.Field[string] `json:"environmentClassId" format:"uuid"`
}

func (r RunnerConfigurationEnvironmentClassGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationEnvironmentClassUpdateParams struct {
	Description        param.Field[string] `json:"description"`
	DisplayName        param.Field[string] `json:"displayName"`
	Enabled            param.Field[bool]   `json:"enabled"`
	EnvironmentClassID param.Field[string] `json:"environmentClassId" format:"uuid"`
}

func (r RunnerConfigurationEnvironmentClassUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// can_create_environments filters the response to only environment classes that
	// can be used to create new environments by the caller. Unlike enabled, which
	// indicates general availability, this ensures the caller only sees environment
	// classes they are allowed to use.
	CanCreateEnvironments param.Field[bool] `json:"canCreateEnvironments"`
	// enabled filters the response to only enabled or disabled environment classes. If
	// not set, all environment classes are returned.
	Enabled param.Field[bool] `json:"enabled"`
	// runner_ids filters the response to only EnvironmentClasses of these Runner IDs
	RunnerIDs param.Field[[]string] `json:"runnerIds" format:"uuid"`
	// runner_kind filters the response to only environment classes from runners of
	// these kinds.
	RunnerKinds param.Field[[]RunnerKind] `json:"runnerKinds"`
	// runner_providers filters the response to only environment classes from runners
	// of these providers.
	RunnerProviders param.Field[[]RunnerProvider] `json:"runnerProviders"`
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
