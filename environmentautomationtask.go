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
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// EnvironmentAutomationTaskService contains methods and other services that help
// with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentAutomationTaskService] method instead.
type EnvironmentAutomationTaskService struct {
	Options    []option.RequestOption
	Executions *EnvironmentAutomationTaskExecutionService
}

// NewEnvironmentAutomationTaskService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEnvironmentAutomationTaskService(opts ...option.RequestOption) (r *EnvironmentAutomationTaskService) {
	r = &EnvironmentAutomationTaskService{}
	r.Options = opts
	r.Executions = NewEnvironmentAutomationTaskExecutionService(opts...)
	return
}

// Creates a new automation task.
//
// Use this method to:
//
// - Define one-off or scheduled tasks
// - Set up build or test automation
// - Configure task dependencies
// - Specify execution environments
//
// ### Examples
//
// - Create basic task:
//
//	Creates a simple build task.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	metadata:
//	  reference: "build"
//	  name: "Build Project"
//	  description: "Builds the project artifacts"
//	  triggeredBy:
//	    - postEnvironmentStart: true
//	spec:
//	  command: "npm run build"
//	```
//
// - Create task with dependencies:
//
//	Creates a task that depends on other services.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	metadata:
//	  reference: "test"
//	  name: "Run Tests"
//	  description: "Runs the test suite"
//	spec:
//	  command: "npm test"
//	dependsOn: ["d2c94c27-3b76-4a42-b88c-95a85e392c68"]
//	```
func (r *EnvironmentAutomationTaskService) New(ctx context.Context, body EnvironmentAutomationTaskNewParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/CreateTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets details about a specific automation task.
//
// Use this method to:
//
// - Check task configuration
// - View task dependencies
// - Monitor task status
//
// ### Examples
//
// - Get task details:
//
//	Retrieves information about a specific task.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *EnvironmentAutomationTaskService) Get(ctx context.Context, body EnvironmentAutomationTaskGetParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/GetTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an automation task configuration.
//
// Use this method to:
//
// - Modify task commands
// - Update task triggers
// - Change dependencies
// - Adjust execution settings
//
// ### Examples
//
// - Update command:
//
//	Changes the task's command.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	spec:
//	  command: "npm run test:coverage"
//	```
//
// - Update triggers:
//
//	Modifies when the task runs.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	metadata:
//	  triggeredBy:
//	    trigger:
//	      - postEnvironmentStart: true
//	```
func (r *EnvironmentAutomationTaskService) Update(ctx context.Context, body EnvironmentAutomationTaskUpdateParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/UpdateTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists automation tasks with optional filtering.
//
// Use this method to:
//
// - View all tasks in an environment
// - Filter tasks by reference
// - Monitor task status
//
// ### Examples
//
// - List environment tasks:
//
//	Shows all tasks for an environment.
//
//	```yaml
//	filter:
//	  environmentIds: ["07e03a28-65a5-4d98-b532-8ea67b188048"]
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter by reference:
//
//	Lists tasks matching specific references.
//
//	```yaml
//	filter:
//	  references: ["build", "test"]
//	pagination:
//	  pageSize: 20
//	```
func (r *EnvironmentAutomationTaskService) List(ctx context.Context, params EnvironmentAutomationTaskListParams, opts ...option.RequestOption) (res *pagination.TasksPage[shared.Task], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.EnvironmentAutomationService/ListTasks"
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

// Lists automation tasks with optional filtering.
//
// Use this method to:
//
// - View all tasks in an environment
// - Filter tasks by reference
// - Monitor task status
//
// ### Examples
//
// - List environment tasks:
//
//	Shows all tasks for an environment.
//
//	```yaml
//	filter:
//	  environmentIds: ["07e03a28-65a5-4d98-b532-8ea67b188048"]
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter by reference:
//
//	Lists tasks matching specific references.
//
//	```yaml
//	filter:
//	  references: ["build", "test"]
//	pagination:
//	  pageSize: 20
//	```
func (r *EnvironmentAutomationTaskService) ListAutoPaging(ctx context.Context, params EnvironmentAutomationTaskListParams, opts ...option.RequestOption) *pagination.TasksPageAutoPager[shared.Task] {
	return pagination.NewTasksPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes an automation task.
//
// Use this method to:
//
// - Remove unused tasks
// - Clean up task configurations
// - Delete obsolete automations
//
// ### Examples
//
// - Delete task:
//
//	Removes a task and its configuration.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *EnvironmentAutomationTaskService) Delete(ctx context.Context, body EnvironmentAutomationTaskDeleteParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/DeleteTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Starts a task by creating a new task execution. This call does not block until
// the task is started; the task will be started asynchronously.
//
// Use this method to:
//
// - Trigger task execution
// - Run one-off tasks
// - Start scheduled tasks immediately
//
// ### Examples
//
// - Start task:
//
//	Creates a new execution of a task.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *EnvironmentAutomationTaskService) Start(ctx context.Context, body EnvironmentAutomationTaskStartParams, opts ...option.RequestOption) (res *EnvironmentAutomationTaskStartResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StartTask"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EnvironmentAutomationTaskNewResponse struct {
	Task shared.Task                              `json:"task,required"`
	JSON environmentAutomationTaskNewResponseJSON `json:"-"`
}

// environmentAutomationTaskNewResponseJSON contains the JSON metadata for the
// struct [EnvironmentAutomationTaskNewResponse]
type environmentAutomationTaskNewResponseJSON struct {
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskNewResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskGetResponse struct {
	Task shared.Task                              `json:"task,required"`
	JSON environmentAutomationTaskGetResponseJSON `json:"-"`
}

// environmentAutomationTaskGetResponseJSON contains the JSON metadata for the
// struct [EnvironmentAutomationTaskGetResponse]
type environmentAutomationTaskGetResponseJSON struct {
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskGetResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskUpdateResponse = interface{}

type EnvironmentAutomationTaskDeleteResponse = interface{}

type EnvironmentAutomationTaskStartResponse struct {
	TaskExecution shared.TaskExecution                       `json:"taskExecution,required"`
	JSON          environmentAutomationTaskStartResponseJSON `json:"-"`
}

// environmentAutomationTaskStartResponseJSON contains the JSON metadata for the
// struct [EnvironmentAutomationTaskStartResponse]
type environmentAutomationTaskStartResponseJSON struct {
	TaskExecution apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentAutomationTaskStartResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationTaskStartResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationTaskNewParams struct {
	DependsOn     param.Field[[]string]                 `json:"dependsOn" format:"uuid"`
	EnvironmentID param.Field[string]                   `json:"environmentId" format:"uuid"`
	Metadata      param.Field[shared.TaskMetadataParam] `json:"metadata"`
	Spec          param.Field[shared.TaskSpecParam]     `json:"spec"`
}

func (r EnvironmentAutomationTaskNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskGetParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationTaskGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskUpdateParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// dependencies specifies the IDs of the automations this task depends on.
	DependsOn param.Field[[]string]                                      `json:"dependsOn" format:"uuid"`
	Metadata  param.Field[EnvironmentAutomationTaskUpdateParamsMetadata] `json:"metadata"`
	Spec      param.Field[EnvironmentAutomationTaskUpdateParamsSpec]     `json:"spec"`
}

func (r EnvironmentAutomationTaskUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskUpdateParamsMetadata struct {
	Description param.Field[string]                                                   `json:"description"`
	Name        param.Field[string]                                                   `json:"name"`
	TriggeredBy param.Field[EnvironmentAutomationTaskUpdateParamsMetadataTriggeredBy] `json:"triggeredBy"`
}

func (r EnvironmentAutomationTaskUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskUpdateParamsMetadataTriggeredBy struct {
	Trigger param.Field[[]shared.AutomationTriggerParam] `json:"trigger"`
}

func (r EnvironmentAutomationTaskUpdateParamsMetadataTriggeredBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskUpdateParamsSpec struct {
	Command param.Field[string]             `json:"command"`
	RunsOn  param.Field[shared.RunsOnParam] `json:"runsOn"`
}

func (r EnvironmentAutomationTaskUpdateParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// filter contains the filter options for listing tasks
	Filter param.Field[EnvironmentAutomationTaskListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing tasks
	Pagination param.Field[EnvironmentAutomationTaskListParamsPagination] `json:"pagination"`
}

func (r EnvironmentAutomationTaskListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [EnvironmentAutomationTaskListParams]'s query parameters as
// `url.Values`.
func (r EnvironmentAutomationTaskListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// filter contains the filter options for listing tasks
type EnvironmentAutomationTaskListParamsFilter struct {
	// environment_ids filters the response to only tasks of these environments
	EnvironmentIDs param.Field[[]string] `json:"environmentIds" format:"uuid"`
	// references filters the response to only services with these references
	References param.Field[[]string] `json:"references"`
	// task_ids filters the response to only tasks with these IDs
	TaskIDs param.Field[[]string] `json:"taskIds" format:"uuid"`
}

func (r EnvironmentAutomationTaskListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing tasks
type EnvironmentAutomationTaskListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EnvironmentAutomationTaskListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskDeleteParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationTaskDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationTaskStartParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r EnvironmentAutomationTaskStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
