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
	"github.com/gitpod-io/flex-sdk-go/shared"
)

// EnvironmentAutomationTaskExecutionService contains methods and other services
// that help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentAutomationTaskExecutionService] method instead.
type EnvironmentAutomationTaskExecutionService struct {
	Options []option.RequestOption
}

// NewEnvironmentAutomationTaskExecutionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewEnvironmentAutomationTaskExecutionService(opts ...option.RequestOption) (r *EnvironmentAutomationTaskExecutionService) {
	r = &EnvironmentAutomationTaskExecutionService{}
	r.Options = opts
	return
}

// GetTaskExecution
func (r *EnvironmentAutomationTaskExecutionService) Get(ctx context.Context, body EnvironmentAutomationTaskExecutionGetParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskExecutionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/GetTaskExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListTaskExecutions
func (r *EnvironmentAutomationTaskExecutionService) List(ctx context.Context, params EnvironmentAutomationTaskExecutionListParams, opts ...option.RequestOption) (res *pagination.TaskExecutionsPage[shared.TaskExecution], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.EnvironmentAutomationService/ListTaskExecutions"
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

// ListTaskExecutions
func (r *EnvironmentAutomationTaskExecutionService) ListAutoPaging(ctx context.Context, params EnvironmentAutomationTaskExecutionListParams, opts ...option.RequestOption) *pagination.TaskExecutionsPageAutoPager[shared.TaskExecution] {
	return pagination.NewTaskExecutionsPageAutoPager(r.List(ctx, params, opts...))
}

// StopTaskExecution
func (r *EnvironmentAutomationTaskExecutionService) Stop(ctx context.Context, body EnvironmentAutomationTaskExecutionStopParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskExecutionStopResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StopTaskExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EnvironmentAutomationTaskExecutionGetResponse struct {
	TaskExecution shared.TaskExecution                              `json:"taskExecution"`
	JSON          environmentAutomationTaskExecutionGetResponseJSON `json:"-"`
}

// environmentAutomationTaskExecutionGetResponseJSON contains the JSON metadata for
// the struct [EnvironmentAutomationTaskExecutionGetResponse]
type environmentAutomationTaskExecutionGetResponseJSON struct {
	TaskExecution apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskExecutionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskExecutionGetResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskExecutionStopResponse = interface{}

type EnvironmentAutomationTaskExecutionGetParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationTaskExecutionGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskExecutionListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// filter contains the filter options for listing task runs
	Filter param.Field[EnvironmentAutomationTaskExecutionListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing task runs
	Pagination param.Field[EnvironmentAutomationTaskExecutionListParamsPagination] `json:"pagination"`
}

func (r EnvironmentAutomationTaskExecutionListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [EnvironmentAutomationTaskExecutionListParams]'s query
// parameters as `url.Values`.
func (r EnvironmentAutomationTaskExecutionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// filter contains the filter options for listing task runs
type EnvironmentAutomationTaskExecutionListParamsFilter struct {
	// environment_ids filters the response to only task runs of these environments
	EnvironmentIDs param.Field[[]string] `json:"environmentIds" format:"uuid"`
	// phases filters the response to only task runs in these phases
	Phases param.Field[[]shared.TaskExecutionPhase] `json:"phases"`
	// task_ids filters the response to only task runs of these tasks
	TaskIDs param.Field[[]string] `json:"taskIds" format:"uuid"`
	// task_references filters the response to only task runs with this reference
	TaskReferences param.Field[[]string] `json:"taskReferences"`
}

func (r EnvironmentAutomationTaskExecutionListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing task runs
type EnvironmentAutomationTaskExecutionListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EnvironmentAutomationTaskExecutionListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskExecutionStopParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationTaskExecutionStopParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
