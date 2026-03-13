// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// AutomationService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutomationService] method instead.
type AutomationService struct {
	Options []option.RequestOption
}

// NewAutomationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAutomationService(opts ...option.RequestOption) (r *AutomationService) {
	r = &AutomationService{}
	r.Options = opts
	return
}

// Creates a new workflow with specified configuration.
//
// Use this method to:
//
// - Set up automated workflows
// - Configure workflow triggers
// - Define workflow actions and steps
// - Set execution limits and constraints
func (r *AutomationService) New(ctx context.Context, body AutomationNewParams, opts ...option.RequestOption) (res *AutomationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.WorkflowService/CreateWorkflow"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets details about a specific workflow.
//
// Use this method to:
//
// - View workflow configuration
// - Check workflow status
// - Get workflow metadata
//
// ### Examples
//
// - Get workflow details:
//
//	Retrieves information about a specific workflow.
//
//	```yaml
//	workflowId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *AutomationService) Get(ctx context.Context, body AutomationGetParams, opts ...option.RequestOption) (res *AutomationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.WorkflowService/GetWorkflow"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates a workflow's configuration using full replacement semantics.
//
// Update Behavior:
//
// - All provided fields completely replace existing values
// - Optional fields that are not provided remain unchanged
// - Complex fields (triggers, action) are replaced entirely, not merged
// - To remove optional fields, explicitly set them to empty/default values
//
// Use this method to:
//
// - Modify workflow settings
// - Update triggers and actions
// - Change execution limits
// - Update workflow steps
//
// ### Examples
//
// - Update workflow name:
//
//	Changes the workflow's display name.
//
//	```yaml
//	workflowId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	name: "Updated Workflow Name"
//	```
//
// - Replace all triggers:
//
//	Completely replaces the workflow's trigger configuration.
//
//	```yaml
//	workflowId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	triggers:
//	  - manual: {}
//	    context:
//	      projects:
//	        projectIds: ["new-project-id"]
//	```
//
// - Update execution limits:
//
//	Completely replaces the workflow's action configuration.
//
//	```yaml
//	workflowId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	action:
//	  limits:
//	    maxParallel: 10
//	    maxTotal: 100
//	  steps:
//	    - task:
//	        command: "npm test"
//	```
func (r *AutomationService) Update(ctx context.Context, body AutomationUpdateParams, opts ...option.RequestOption) (res *AutomationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.WorkflowService/UpdateWorkflow"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// ListWorkflows
func (r *AutomationService) List(ctx context.Context, params AutomationListParams, opts ...option.RequestOption) (res *pagination.WorkflowsPage[Workflow], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.WorkflowService/ListWorkflows"
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

// ListWorkflows
func (r *AutomationService) ListAutoPaging(ctx context.Context, params AutomationListParams, opts ...option.RequestOption) *pagination.WorkflowsPageAutoPager[Workflow] {
	return pagination.NewWorkflowsPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a workflow permanently.
//
// Use this method to:
//
// - Remove unused workflows
// - Clean up test workflows
// - Delete obsolete configurations
//
// ### Examples
//
// - Delete workflow:
//
//	Permanently removes a workflow.
//
//	```yaml
//	workflowId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *AutomationService) Delete(ctx context.Context, body AutomationDeleteParams, opts ...option.RequestOption) (res *AutomationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.WorkflowService/DeleteWorkflow"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Cancels a running workflow execution.
//
// Use this method to:
//
// - Stop long-running executions
// - Cancel failed executions
// - Manage resource usage
//
// ### Examples
//
// - Cancel execution:
//
//	Stops a running workflow execution.
//
//	```yaml
//	workflowExecutionId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *AutomationService) CancelExecution(ctx context.Context, body AutomationCancelExecutionParams, opts ...option.RequestOption) (res *AutomationCancelExecutionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.WorkflowService/CancelWorkflowExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Cancels a running workflow execution action.
//
// Use this method to:
//
// - Stop long-running actions
// - Cancel failed actions
// - Manage resource usage
//
// ### Examples
//
// - Cancel execution action:
//
//	Stops a running workflow execution action.
//
//	```yaml
//	workflowExecutionActionId: "a1b2c3d4-5e6f-7890-abcd-ef1234567890"
//	```
func (r *AutomationService) CancelExecutionAction(ctx context.Context, body AutomationCancelExecutionActionParams, opts ...option.RequestOption) (res *AutomationCancelExecutionActionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.WorkflowService/CancelWorkflowExecutionAction"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists workflow execution actions with optional filtering.
//
// Use this method to:
//
// - Monitor individual action execution status
// - Debug action failures
// - Track resource usage per action
//
// ### Examples
//
// - List execution actions for workflow execution:
//
//	Shows all execution actions for a specific workflow execution.
//
//	```yaml
//	filter:
//	  workflowExecutionIds: ["d2c94c27-3b76-4a42-b88c-95a85e392c68"]
//	pagination:
//	  pageSize: 20
//	```
func (r *AutomationService) ListExecutionActions(ctx context.Context, params AutomationListExecutionActionsParams, opts ...option.RequestOption) (res *pagination.WorkflowExecutionActionsPage[WorkflowExecutionAction], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.WorkflowService/ListWorkflowExecutionActions"
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

// Lists workflow execution actions with optional filtering.
//
// Use this method to:
//
// - Monitor individual action execution status
// - Debug action failures
// - Track resource usage per action
//
// ### Examples
//
// - List execution actions for workflow execution:
//
//	Shows all execution actions for a specific workflow execution.
//
//	```yaml
//	filter:
//	  workflowExecutionIds: ["d2c94c27-3b76-4a42-b88c-95a85e392c68"]
//	pagination:
//	  pageSize: 20
//	```
func (r *AutomationService) ListExecutionActionsAutoPaging(ctx context.Context, params AutomationListExecutionActionsParams, opts ...option.RequestOption) *pagination.WorkflowExecutionActionsPageAutoPager[WorkflowExecutionAction] {
	return pagination.NewWorkflowExecutionActionsPageAutoPager(r.ListExecutionActions(ctx, params, opts...))
}

// Lists outputs produced by workflow execution actions.
//
// Use this method to:
//
//   - Retrieve test results, coverage metrics, or other structured data from
//     executions
//   - Aggregate outputs across multiple workflow executions
//   - Build dashboards or reports from execution data
//
// ### Examples
//
// - List outputs for a workflow execution:
//
//	Retrieves all outputs produced by actions in the specified execution.
//
//	```yaml
//	filter:
//	  workflowExecutionIds: ["d2c94c27-3b76-4a42-b88c-95a85e392c68"]
//	pagination:
//	  pageSize: 50
//	```
func (r *AutomationService) ListExecutionOutputs(ctx context.Context, params AutomationListExecutionOutputsParams, opts ...option.RequestOption) (res *pagination.OutputsPage[AutomationListExecutionOutputsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.WorkflowService/ListWorkflowExecutionOutputs"
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

// Lists outputs produced by workflow execution actions.
//
// Use this method to:
//
//   - Retrieve test results, coverage metrics, or other structured data from
//     executions
//   - Aggregate outputs across multiple workflow executions
//   - Build dashboards or reports from execution data
//
// ### Examples
//
// - List outputs for a workflow execution:
//
//	Retrieves all outputs produced by actions in the specified execution.
//
//	```yaml
//	filter:
//	  workflowExecutionIds: ["d2c94c27-3b76-4a42-b88c-95a85e392c68"]
//	pagination:
//	  pageSize: 50
//	```
func (r *AutomationService) ListExecutionOutputsAutoPaging(ctx context.Context, params AutomationListExecutionOutputsParams, opts ...option.RequestOption) *pagination.OutputsPageAutoPager[AutomationListExecutionOutputsResponse] {
	return pagination.NewOutputsPageAutoPager(r.ListExecutionOutputs(ctx, params, opts...))
}

// Lists workflow executions with optional filtering.
//
// Use this method to:
//
// - Monitor workflow execution history
// - Track execution status
// - Debug workflow issues
//
// ### Examples
//
// - List executions for workflow:
//
//	Shows all executions for a specific workflow.
//
//	```yaml
//	filter:
//	  workflowIds: ["b0e12f6c-4c67-429d-a4a6-d9838b5da047"]
//	pagination:
//	  pageSize: 20
//	```
func (r *AutomationService) ListExecutions(ctx context.Context, params AutomationListExecutionsParams, opts ...option.RequestOption) (res *pagination.WorkflowExecutionsPage[WorkflowExecution], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.WorkflowService/ListWorkflowExecutions"
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

// Lists workflow executions with optional filtering.
//
// Use this method to:
//
// - Monitor workflow execution history
// - Track execution status
// - Debug workflow issues
//
// ### Examples
//
// - List executions for workflow:
//
//	Shows all executions for a specific workflow.
//
//	```yaml
//	filter:
//	  workflowIds: ["b0e12f6c-4c67-429d-a4a6-d9838b5da047"]
//	pagination:
//	  pageSize: 20
//	```
func (r *AutomationService) ListExecutionsAutoPaging(ctx context.Context, params AutomationListExecutionsParams, opts ...option.RequestOption) *pagination.WorkflowExecutionsPageAutoPager[WorkflowExecution] {
	return pagination.NewWorkflowExecutionsPageAutoPager(r.ListExecutions(ctx, params, opts...))
}

// Gets details about a specific workflow execution.
//
// Use this method to:
//
// - Check execution status
// - View execution results
// - Monitor execution progress
//
// ### Examples
//
// - Get execution details:
//
//	Retrieves information about a specific execution.
//
//	```yaml
//	workflowExecutionId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *AutomationService) GetExecution(ctx context.Context, body AutomationGetExecutionParams, opts ...option.RequestOption) (res *AutomationGetExecutionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.WorkflowService/GetWorkflowExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets details about a specific workflow execution action.
//
// Use this method to:
//
// - Check execution action status
// - View execution action results
// - Monitor execution action progress
//
// ### Examples
//
// - Get execution action details:
//
//	Retrieves information about a specific execution action.
//
//	```yaml
//	workflowExecutionActionId: "a1b2c3d4-5e6f-7890-abcd-ef1234567890"
//	```
func (r *AutomationService) GetExecutionAction(ctx context.Context, body AutomationGetExecutionActionParams, opts ...option.RequestOption) (res *AutomationGetExecutionActionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.WorkflowService/GetWorkflowExecutionAction"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Starts a workflow execution.
//
// Use this method to:
//
// - Start workflow execution on demand
// - Test workflow configurations
// - Run workflows outside of automatic triggers
//
// ### Examples
//
// - Start workflow:
//
//	Starts a workflow execution manually.
//
//	```yaml
//	workflowId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *AutomationService) StartExecution(ctx context.Context, body AutomationStartExecutionParams, opts ...option.RequestOption) (res *AutomationStartExecutionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.WorkflowService/StartWorkflow"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Workflow represents a workflow configuration.
type Workflow struct {
	ID string `json:"id" format:"uuid"`
	// WorkflowMetadata contains workflow metadata.
	Metadata WorkflowMetadata `json:"metadata"`
	Spec     WorkflowSpec     `json:"spec"`
	// Webhook URL for triggering this workflow via HTTP POST Format:
	// {base_url}/workflows/{workflow_id}/webhooks
	WebhookURL string       `json:"webhookUrl"`
	JSON       workflowJSON `json:"-"`
}

// workflowJSON contains the JSON metadata for the struct [Workflow]
type workflowJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	WebhookURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Workflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowJSON) RawJSON() string {
	return r.raw
}

// WorkflowMetadata contains workflow metadata.
type WorkflowMetadata struct {
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	CreatedAt   time.Time      `json:"createdAt" format:"date-time"`
	Creator     shared.Subject `json:"creator"`
	Description string         `json:"description"`
	Executor    shared.Subject `json:"executor"`
	Name        string         `json:"name"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	UpdatedAt time.Time            `json:"updatedAt" format:"date-time"`
	JSON      workflowMetadataJSON `json:"-"`
}

// workflowMetadataJSON contains the JSON metadata for the struct
// [WorkflowMetadata]
type workflowMetadataJSON struct {
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	Executor    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowMetadataJSON) RawJSON() string {
	return r.raw
}

type WorkflowSpec struct {
	// WorkflowAction defines the actions to be executed in a workflow.
	Action WorkflowAction `json:"action"`
	// WorkflowAction defines the actions to be executed in a workflow.
	Report   WorkflowAction    `json:"report"`
	Triggers []WorkflowTrigger `json:"triggers"`
	JSON     workflowSpecJSON  `json:"-"`
}

// workflowSpecJSON contains the JSON metadata for the struct [WorkflowSpec]
type workflowSpecJSON struct {
	Action      apijson.Field
	Report      apijson.Field
	Triggers    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowSpecJSON) RawJSON() string {
	return r.raw
}

// WorkflowAction defines the actions to be executed in a workflow.
type WorkflowAction struct {
	// Limits defines execution limits for workflow actions. Concurrent actions limit
	// cannot exceed total actions limit:
	//
	// ```
	// this.max_parallel <= this.max_total
	// ```
	Limits WorkflowActionLimits `json:"limits" api:"required"`
	// Automation must have between 1 and 50 steps:
	//
	// ```
	// size(this) >= 1 && size(this) <= 50
	// ```
	Steps []WorkflowStep     `json:"steps"`
	JSON  workflowActionJSON `json:"-"`
}

// workflowActionJSON contains the JSON metadata for the struct [WorkflowAction]
type workflowActionJSON struct {
	Limits      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowActionJSON) RawJSON() string {
	return r.raw
}

// Limits defines execution limits for workflow actions. Concurrent actions limit
// cannot exceed total actions limit:
//
// ```
// this.max_parallel <= this.max_total
// ```
type WorkflowActionLimits struct {
	// Maximum parallel actions must be between 1 and 25:
	//
	// ```
	// this >= 1 && this <= 25
	// ```
	MaxParallel int64 `json:"maxParallel"`
	// Maximum total actions must be between 1 and 100:
	//
	// ```
	// this >= 1 && this <= 100
	// ```
	MaxTotal int64 `json:"maxTotal"`
	// PerExecution defines limits per execution action.
	PerExecution WorkflowActionLimitsPerExecution `json:"perExecution"`
	JSON         workflowActionLimitsJSON         `json:"-"`
}

// workflowActionLimitsJSON contains the JSON metadata for the struct
// [WorkflowActionLimits]
type workflowActionLimitsJSON struct {
	MaxParallel  apijson.Field
	MaxTotal     apijson.Field
	PerExecution apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WorkflowActionLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowActionLimitsJSON) RawJSON() string {
	return r.raw
}

// PerExecution defines limits per execution action.
type WorkflowActionLimitsPerExecution struct {
	// Maximum time allowed for a single execution action. Use standard duration format
	// (e.g., "30m" for 30 minutes, "2h" for 2 hours).
	MaxTime string                               `json:"maxTime" format:"regex"`
	JSON    workflowActionLimitsPerExecutionJSON `json:"-"`
}

// workflowActionLimitsPerExecutionJSON contains the JSON metadata for the struct
// [WorkflowActionLimitsPerExecution]
type workflowActionLimitsPerExecutionJSON struct {
	MaxTime     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowActionLimitsPerExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowActionLimitsPerExecutionJSON) RawJSON() string {
	return r.raw
}

// WorkflowAction defines the actions to be executed in a workflow.
type WorkflowActionParam struct {
	// Limits defines execution limits for workflow actions. Concurrent actions limit
	// cannot exceed total actions limit:
	//
	// ```
	// this.max_parallel <= this.max_total
	// ```
	Limits param.Field[WorkflowActionLimitsParam] `json:"limits" api:"required"`
	// Automation must have between 1 and 50 steps:
	//
	// ```
	// size(this) >= 1 && size(this) <= 50
	// ```
	Steps param.Field[[]WorkflowStepParam] `json:"steps"`
}

func (r WorkflowActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Limits defines execution limits for workflow actions. Concurrent actions limit
// cannot exceed total actions limit:
//
// ```
// this.max_parallel <= this.max_total
// ```
type WorkflowActionLimitsParam struct {
	// Maximum parallel actions must be between 1 and 25:
	//
	// ```
	// this >= 1 && this <= 25
	// ```
	MaxParallel param.Field[int64] `json:"maxParallel"`
	// Maximum total actions must be between 1 and 100:
	//
	// ```
	// this >= 1 && this <= 100
	// ```
	MaxTotal param.Field[int64] `json:"maxTotal"`
	// PerExecution defines limits per execution action.
	PerExecution param.Field[WorkflowActionLimitsPerExecutionParam] `json:"perExecution"`
}

func (r WorkflowActionLimitsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PerExecution defines limits per execution action.
type WorkflowActionLimitsPerExecutionParam struct {
	// Maximum time allowed for a single execution action. Use standard duration format
	// (e.g., "30m" for 30 minutes, "2h" for 2 hours).
	MaxTime param.Field[string] `json:"maxTime" format:"regex"`
}

func (r WorkflowActionLimitsPerExecutionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WorkflowExecution represents a workflow execution instance.
type WorkflowExecution struct {
	ID string `json:"id" format:"uuid"`
	// WorkflowExecutionMetadata contains workflow execution metadata.
	Metadata WorkflowExecutionMetadata `json:"metadata"`
	// WorkflowExecutionSpec contains the specification used for this execution.
	Spec WorkflowExecutionSpec `json:"spec"`
	// WorkflowExecutionStatus contains the current status of a workflow execution.
	Status WorkflowExecutionStatus `json:"status"`
	JSON   workflowExecutionJSON   `json:"-"`
}

// workflowExecutionJSON contains the JSON metadata for the struct
// [WorkflowExecution]
type workflowExecutionJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionJSON) RawJSON() string {
	return r.raw
}

// WorkflowExecutionMetadata contains workflow execution metadata.
type WorkflowExecutionMetadata struct {
	Creator  shared.Subject `json:"creator"`
	Executor shared.Subject `json:"executor"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	FinishedAt time.Time `json:"finishedAt" format:"date-time"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	StartedAt  time.Time                     `json:"startedAt" format:"date-time"`
	WorkflowID string                        `json:"workflowId" format:"uuid"`
	JSON       workflowExecutionMetadataJSON `json:"-"`
}

// workflowExecutionMetadataJSON contains the JSON metadata for the struct
// [WorkflowExecutionMetadata]
type workflowExecutionMetadataJSON struct {
	Creator     apijson.Field
	Executor    apijson.Field
	FinishedAt  apijson.Field
	StartedAt   apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionMetadataJSON) RawJSON() string {
	return r.raw
}

// WorkflowExecutionSpec contains the specification used for this execution.
type WorkflowExecutionSpec struct {
	// WorkflowAction defines the actions to be executed in a workflow.
	Action WorkflowAction `json:"action"`
	// WorkflowAction defines the actions to be executed in a workflow.
	Report WorkflowAction `json:"report"`
	// WorkflowExecutionTrigger represents a workflow execution trigger instance.
	Trigger WorkflowExecutionSpecTrigger `json:"trigger"`
	JSON    workflowExecutionSpecJSON    `json:"-"`
}

// workflowExecutionSpecJSON contains the JSON metadata for the struct
// [WorkflowExecutionSpec]
type workflowExecutionSpecJSON struct {
	Action      apijson.Field
	Report      apijson.Field
	Trigger     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionSpecJSON) RawJSON() string {
	return r.raw
}

// WorkflowExecutionTrigger represents a workflow execution trigger instance.
type WorkflowExecutionSpecTrigger struct {
	// Context from the workflow trigger - copied at execution time for immutability.
	// This allows the reconciler to create actions without fetching the workflow
	// definition.
	Context WorkflowTriggerContext `json:"context" api:"required"`
	// Manual trigger - empty message since no additional data needed
	Manual interface{} `json:"manual"`
	// PullRequest represents pull request metadata from source control systems. This
	// message is used across workflow triggers, executions, and agent contexts to
	// maintain consistent PR information throughout the system.
	PullRequest WorkflowExecutionSpecTriggerPullRequest `json:"pullRequest"`
	// Time trigger - just the timestamp when it was triggered
	Time WorkflowExecutionSpecTriggerTime `json:"time"`
	JSON workflowExecutionSpecTriggerJSON `json:"-"`
}

// workflowExecutionSpecTriggerJSON contains the JSON metadata for the struct
// [WorkflowExecutionSpecTrigger]
type workflowExecutionSpecTriggerJSON struct {
	Context     apijson.Field
	Manual      apijson.Field
	PullRequest apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionSpecTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionSpecTriggerJSON) RawJSON() string {
	return r.raw
}

// PullRequest represents pull request metadata from source control systems. This
// message is used across workflow triggers, executions, and agent contexts to
// maintain consistent PR information throughout the system.
type WorkflowExecutionSpecTriggerPullRequest struct {
	// Unique identifier from the source system (e.g., "123" for GitHub PR #123)
	ID string `json:"id"`
	// Author name as provided by the SCM system
	Author string `json:"author"`
	// Whether this is a draft pull request
	Draft bool `json:"draft"`
	// Source branch name (the branch being merged from)
	FromBranch string `json:"fromBranch"`
	// Repository information
	Repository WorkflowExecutionSpecTriggerPullRequestRepository `json:"repository"`
	// Current state of the pull request
	State shared.State `json:"state"`
	// Pull request title
	Title string `json:"title"`
	// Target branch name (the branch being merged into)
	ToBranch string `json:"toBranch"`
	// Pull request URL (e.g., "https://github.com/owner/repo/pull/123")
	URL  string                                      `json:"url"`
	JSON workflowExecutionSpecTriggerPullRequestJSON `json:"-"`
}

// workflowExecutionSpecTriggerPullRequestJSON contains the JSON metadata for the
// struct [WorkflowExecutionSpecTriggerPullRequest]
type workflowExecutionSpecTriggerPullRequestJSON struct {
	ID          apijson.Field
	Author      apijson.Field
	Draft       apijson.Field
	FromBranch  apijson.Field
	Repository  apijson.Field
	State       apijson.Field
	Title       apijson.Field
	ToBranch    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionSpecTriggerPullRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionSpecTriggerPullRequestJSON) RawJSON() string {
	return r.raw
}

// Repository information
type WorkflowExecutionSpecTriggerPullRequestRepository struct {
	CloneURL string                                                `json:"cloneUrl"`
	Host     string                                                `json:"host"`
	Name     string                                                `json:"name"`
	Owner    string                                                `json:"owner"`
	JSON     workflowExecutionSpecTriggerPullRequestRepositoryJSON `json:"-"`
}

// workflowExecutionSpecTriggerPullRequestRepositoryJSON contains the JSON metadata
// for the struct [WorkflowExecutionSpecTriggerPullRequestRepository]
type workflowExecutionSpecTriggerPullRequestRepositoryJSON struct {
	CloneURL    apijson.Field
	Host        apijson.Field
	Name        apijson.Field
	Owner       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionSpecTriggerPullRequestRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionSpecTriggerPullRequestRepositoryJSON) RawJSON() string {
	return r.raw
}

// Time trigger - just the timestamp when it was triggered
type WorkflowExecutionSpecTriggerTime struct {
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	TriggeredAt time.Time                            `json:"triggeredAt" format:"date-time"`
	JSON        workflowExecutionSpecTriggerTimeJSON `json:"-"`
}

// workflowExecutionSpecTriggerTimeJSON contains the JSON metadata for the struct
// [WorkflowExecutionSpecTriggerTime]
type workflowExecutionSpecTriggerTimeJSON struct {
	TriggeredAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionSpecTriggerTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionSpecTriggerTimeJSON) RawJSON() string {
	return r.raw
}

// WorkflowExecutionStatus contains the current status of a workflow execution.
type WorkflowExecutionStatus struct {
	DoneActionCount   int64 `json:"doneActionCount"`
	FailedActionCount int64 `json:"failedActionCount"`
	// Structured failures that caused the workflow execution to fail. Provides
	// detailed error codes, messages, and retry information.
	Failures           []WorkflowExecutionStatusFailure `json:"failures"`
	PendingActionCount int64                            `json:"pendingActionCount"`
	Phase              WorkflowExecutionStatusPhase     `json:"phase"`
	RunningActionCount int64                            `json:"runningActionCount"`
	StoppedActionCount int64                            `json:"stoppedActionCount"`
	// Structured warnings about the workflow execution. Provides detailed warning
	// codes and messages.
	Warnings []WorkflowExecutionStatusWarning `json:"warnings"`
	JSON     workflowExecutionStatusJSON      `json:"-"`
}

// workflowExecutionStatusJSON contains the JSON metadata for the struct
// [WorkflowExecutionStatus]
type workflowExecutionStatusJSON struct {
	DoneActionCount    apijson.Field
	FailedActionCount  apijson.Field
	Failures           apijson.Field
	PendingActionCount apijson.Field
	Phase              apijson.Field
	RunningActionCount apijson.Field
	StoppedActionCount apijson.Field
	Warnings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkflowExecutionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionStatusJSON) RawJSON() string {
	return r.raw
}

// WorkflowError provides structured error information for workflow failures. This
// enables the reconciler to make informed retry decisions and the frontend to
// display actionable error messages.
type WorkflowExecutionStatusFailure struct {
	// Error code identifying the type of error.
	Code WorkflowExecutionStatusFailuresCode `json:"code"`
	// Human-readable error message.
	Message string `json:"message"`
	// Additional metadata about the error. Common keys include:
	//
	// - environment_id: ID of the environment
	// - task_id: ID of the task
	// - service_id: ID of the service
	// - workflow_id: ID of the workflow
	// - workflow_execution_id: ID of the workflow execution
	Meta map[string]string `json:"meta"`
	// Reason explaining why the error occurred. Examples: "not_found", "stopped",
	// "deleted", "creation_failed", "start_failed"
	Reason string `json:"reason"`
	// Retry configuration. If not set, the error is considered non-retriable.
	Retry WorkflowExecutionStatusFailuresRetry `json:"retry" api:"nullable"`
	JSON  workflowExecutionStatusFailureJSON   `json:"-"`
}

// workflowExecutionStatusFailureJSON contains the JSON metadata for the struct
// [WorkflowExecutionStatusFailure]
type workflowExecutionStatusFailureJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Meta        apijson.Field
	Reason      apijson.Field
	Retry       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionStatusFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionStatusFailureJSON) RawJSON() string {
	return r.raw
}

// Error code identifying the type of error.
type WorkflowExecutionStatusFailuresCode string

const (
	WorkflowExecutionStatusFailuresCodeWorkflowErrorCodeUnspecified      WorkflowExecutionStatusFailuresCode = "WORKFLOW_ERROR_CODE_UNSPECIFIED"
	WorkflowExecutionStatusFailuresCodeWorkflowErrorCodeEnvironmentError WorkflowExecutionStatusFailuresCode = "WORKFLOW_ERROR_CODE_ENVIRONMENT_ERROR"
	WorkflowExecutionStatusFailuresCodeWorkflowErrorCodeAgentError       WorkflowExecutionStatusFailuresCode = "WORKFLOW_ERROR_CODE_AGENT_ERROR"
)

func (r WorkflowExecutionStatusFailuresCode) IsKnown() bool {
	switch r {
	case WorkflowExecutionStatusFailuresCodeWorkflowErrorCodeUnspecified, WorkflowExecutionStatusFailuresCodeWorkflowErrorCodeEnvironmentError, WorkflowExecutionStatusFailuresCodeWorkflowErrorCodeAgentError:
		return true
	}
	return false
}

// Retry configuration. If not set, the error is considered non-retriable.
type WorkflowExecutionStatusFailuresRetry struct {
	// Whether the error is retriable.
	Retriable bool `json:"retriable"`
	// Suggested duration to wait before retrying. Only meaningful when retriable is
	// true.
	RetryAfter string                                   `json:"retryAfter" format:"regex"`
	JSON       workflowExecutionStatusFailuresRetryJSON `json:"-"`
}

// workflowExecutionStatusFailuresRetryJSON contains the JSON metadata for the
// struct [WorkflowExecutionStatusFailuresRetry]
type workflowExecutionStatusFailuresRetryJSON struct {
	Retriable   apijson.Field
	RetryAfter  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionStatusFailuresRetry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionStatusFailuresRetryJSON) RawJSON() string {
	return r.raw
}

type WorkflowExecutionStatusPhase string

const (
	WorkflowExecutionStatusPhaseWorkflowExecutionPhaseUnspecified WorkflowExecutionStatusPhase = "WORKFLOW_EXECUTION_PHASE_UNSPECIFIED"
	WorkflowExecutionStatusPhaseWorkflowExecutionPhasePending     WorkflowExecutionStatusPhase = "WORKFLOW_EXECUTION_PHASE_PENDING"
	WorkflowExecutionStatusPhaseWorkflowExecutionPhaseRunning     WorkflowExecutionStatusPhase = "WORKFLOW_EXECUTION_PHASE_RUNNING"
	WorkflowExecutionStatusPhaseWorkflowExecutionPhaseStopping    WorkflowExecutionStatusPhase = "WORKFLOW_EXECUTION_PHASE_STOPPING"
	WorkflowExecutionStatusPhaseWorkflowExecutionPhaseStopped     WorkflowExecutionStatusPhase = "WORKFLOW_EXECUTION_PHASE_STOPPED"
	WorkflowExecutionStatusPhaseWorkflowExecutionPhaseDeleting    WorkflowExecutionStatusPhase = "WORKFLOW_EXECUTION_PHASE_DELETING"
	WorkflowExecutionStatusPhaseWorkflowExecutionPhaseDeleted     WorkflowExecutionStatusPhase = "WORKFLOW_EXECUTION_PHASE_DELETED"
	WorkflowExecutionStatusPhaseWorkflowExecutionPhaseCompleted   WorkflowExecutionStatusPhase = "WORKFLOW_EXECUTION_PHASE_COMPLETED"
)

func (r WorkflowExecutionStatusPhase) IsKnown() bool {
	switch r {
	case WorkflowExecutionStatusPhaseWorkflowExecutionPhaseUnspecified, WorkflowExecutionStatusPhaseWorkflowExecutionPhasePending, WorkflowExecutionStatusPhaseWorkflowExecutionPhaseRunning, WorkflowExecutionStatusPhaseWorkflowExecutionPhaseStopping, WorkflowExecutionStatusPhaseWorkflowExecutionPhaseStopped, WorkflowExecutionStatusPhaseWorkflowExecutionPhaseDeleting, WorkflowExecutionStatusPhaseWorkflowExecutionPhaseDeleted, WorkflowExecutionStatusPhaseWorkflowExecutionPhaseCompleted:
		return true
	}
	return false
}

// WorkflowError provides structured error information for workflow failures. This
// enables the reconciler to make informed retry decisions and the frontend to
// display actionable error messages.
type WorkflowExecutionStatusWarning struct {
	// Error code identifying the type of error.
	Code WorkflowExecutionStatusWarningsCode `json:"code"`
	// Human-readable error message.
	Message string `json:"message"`
	// Additional metadata about the error. Common keys include:
	//
	// - environment_id: ID of the environment
	// - task_id: ID of the task
	// - service_id: ID of the service
	// - workflow_id: ID of the workflow
	// - workflow_execution_id: ID of the workflow execution
	Meta map[string]string `json:"meta"`
	// Reason explaining why the error occurred. Examples: "not_found", "stopped",
	// "deleted", "creation_failed", "start_failed"
	Reason string `json:"reason"`
	// Retry configuration. If not set, the error is considered non-retriable.
	Retry WorkflowExecutionStatusWarningsRetry `json:"retry" api:"nullable"`
	JSON  workflowExecutionStatusWarningJSON   `json:"-"`
}

// workflowExecutionStatusWarningJSON contains the JSON metadata for the struct
// [WorkflowExecutionStatusWarning]
type workflowExecutionStatusWarningJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Meta        apijson.Field
	Reason      apijson.Field
	Retry       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionStatusWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionStatusWarningJSON) RawJSON() string {
	return r.raw
}

// Error code identifying the type of error.
type WorkflowExecutionStatusWarningsCode string

const (
	WorkflowExecutionStatusWarningsCodeWorkflowErrorCodeUnspecified      WorkflowExecutionStatusWarningsCode = "WORKFLOW_ERROR_CODE_UNSPECIFIED"
	WorkflowExecutionStatusWarningsCodeWorkflowErrorCodeEnvironmentError WorkflowExecutionStatusWarningsCode = "WORKFLOW_ERROR_CODE_ENVIRONMENT_ERROR"
	WorkflowExecutionStatusWarningsCodeWorkflowErrorCodeAgentError       WorkflowExecutionStatusWarningsCode = "WORKFLOW_ERROR_CODE_AGENT_ERROR"
)

func (r WorkflowExecutionStatusWarningsCode) IsKnown() bool {
	switch r {
	case WorkflowExecutionStatusWarningsCodeWorkflowErrorCodeUnspecified, WorkflowExecutionStatusWarningsCodeWorkflowErrorCodeEnvironmentError, WorkflowExecutionStatusWarningsCodeWorkflowErrorCodeAgentError:
		return true
	}
	return false
}

// Retry configuration. If not set, the error is considered non-retriable.
type WorkflowExecutionStatusWarningsRetry struct {
	// Whether the error is retriable.
	Retriable bool `json:"retriable"`
	// Suggested duration to wait before retrying. Only meaningful when retriable is
	// true.
	RetryAfter string                                   `json:"retryAfter" format:"regex"`
	JSON       workflowExecutionStatusWarningsRetryJSON `json:"-"`
}

// workflowExecutionStatusWarningsRetryJSON contains the JSON metadata for the
// struct [WorkflowExecutionStatusWarningsRetry]
type workflowExecutionStatusWarningsRetryJSON struct {
	Retriable   apijson.Field
	RetryAfter  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionStatusWarningsRetry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionStatusWarningsRetryJSON) RawJSON() string {
	return r.raw
}

// WorkflowExecutionAction represents a workflow execution action instance.
type WorkflowExecutionAction struct {
	ID string `json:"id" format:"uuid"`
	// WorkflowExecutionActionMetadata contains workflow execution action metadata.
	Metadata WorkflowExecutionActionMetadata `json:"metadata"`
	// WorkflowExecutionActionSpec contains the specification for this execution
	// action.
	Spec WorkflowExecutionActionSpec `json:"spec"`
	// WorkflowExecutionActionStatus contains the current status of a workflow
	// execution action.
	Status WorkflowExecutionActionStatus `json:"status"`
	JSON   workflowExecutionActionJSON   `json:"-"`
}

// workflowExecutionActionJSON contains the JSON metadata for the struct
// [WorkflowExecutionAction]
type workflowExecutionActionJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionJSON) RawJSON() string {
	return r.raw
}

// WorkflowExecutionActionMetadata contains workflow execution action metadata.
type WorkflowExecutionActionMetadata struct {
	// Human-readable name for this action based on its context. Examples:
	// "gitpod-io/gitpod-next" for repository context, "My Project" for project
	// context. Will be empty string for actions created before this field was added.
	ActionName string `json:"actionName"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	FinishedAt time.Time `json:"finishedAt" format:"date-time"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	StartedAt           time.Time                           `json:"startedAt" format:"date-time"`
	WorkflowExecutionID string                              `json:"workflowExecutionId" format:"uuid"`
	WorkflowID          string                              `json:"workflowId" format:"uuid"`
	JSON                workflowExecutionActionMetadataJSON `json:"-"`
}

// workflowExecutionActionMetadataJSON contains the JSON metadata for the struct
// [WorkflowExecutionActionMetadata]
type workflowExecutionActionMetadataJSON struct {
	ActionName          apijson.Field
	FinishedAt          apijson.Field
	StartedAt           apijson.Field
	WorkflowExecutionID apijson.Field
	WorkflowID          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *WorkflowExecutionActionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionMetadataJSON) RawJSON() string {
	return r.raw
}

// WorkflowExecutionActionSpec contains the specification for this execution
// action.
type WorkflowExecutionActionSpec struct {
	// Context for the execution action - specifies where and how the action executes.
	// This is resolved from the workflow trigger context and contains the specific
	// project, repository, or agent context for this execution instance.
	Context AgentCodeContext `json:"context"`
	// PerExecution defines limits per execution action.
	Limits WorkflowExecutionActionSpecLimits `json:"limits"`
	JSON   workflowExecutionActionSpecJSON   `json:"-"`
}

// workflowExecutionActionSpecJSON contains the JSON metadata for the struct
// [WorkflowExecutionActionSpec]
type workflowExecutionActionSpecJSON struct {
	Context     apijson.Field
	Limits      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionActionSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionSpecJSON) RawJSON() string {
	return r.raw
}

// PerExecution defines limits per execution action.
type WorkflowExecutionActionSpecLimits struct {
	// Maximum time allowed for a single execution action. Use standard duration format
	// (e.g., "30m" for 30 minutes, "2h" for 2 hours).
	MaxTime string                                `json:"maxTime" format:"regex"`
	JSON    workflowExecutionActionSpecLimitsJSON `json:"-"`
}

// workflowExecutionActionSpecLimitsJSON contains the JSON metadata for the struct
// [WorkflowExecutionActionSpecLimits]
type workflowExecutionActionSpecLimitsJSON struct {
	MaxTime     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionActionSpecLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionSpecLimitsJSON) RawJSON() string {
	return r.raw
}

// WorkflowExecutionActionStatus contains the current status of a workflow
// execution action.
type WorkflowExecutionActionStatus struct {
	AgentExecutionID string `json:"agentExecutionId"`
	EnvironmentID    string `json:"environmentId" format:"uuid"`
	// Structured failures that caused the workflow execution action to fail. Provides
	// detailed error codes, messages, and retry information.
	Failures []WorkflowExecutionActionStatusFailure `json:"failures"`
	// WorkflowExecutionActionPhase defines the phases of workflow execution action.
	Phase WorkflowExecutionActionStatusPhase `json:"phase"`
	// Step-level progress tracking
	StepStatuses []WorkflowExecutionActionStatusStepStatus `json:"stepStatuses"`
	// Structured warnings about the workflow execution action. Provides detailed
	// warning codes and messages.
	Warnings []WorkflowExecutionActionStatusWarning `json:"warnings"`
	JSON     workflowExecutionActionStatusJSON      `json:"-"`
}

// workflowExecutionActionStatusJSON contains the JSON metadata for the struct
// [WorkflowExecutionActionStatus]
type workflowExecutionActionStatusJSON struct {
	AgentExecutionID apijson.Field
	EnvironmentID    apijson.Field
	Failures         apijson.Field
	Phase            apijson.Field
	StepStatuses     apijson.Field
	Warnings         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkflowExecutionActionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionStatusJSON) RawJSON() string {
	return r.raw
}

// WorkflowError provides structured error information for workflow failures. This
// enables the reconciler to make informed retry decisions and the frontend to
// display actionable error messages.
type WorkflowExecutionActionStatusFailure struct {
	// Error code identifying the type of error.
	Code WorkflowExecutionActionStatusFailuresCode `json:"code"`
	// Human-readable error message.
	Message string `json:"message"`
	// Additional metadata about the error. Common keys include:
	//
	// - environment_id: ID of the environment
	// - task_id: ID of the task
	// - service_id: ID of the service
	// - workflow_id: ID of the workflow
	// - workflow_execution_id: ID of the workflow execution
	Meta map[string]string `json:"meta"`
	// Reason explaining why the error occurred. Examples: "not_found", "stopped",
	// "deleted", "creation_failed", "start_failed"
	Reason string `json:"reason"`
	// Retry configuration. If not set, the error is considered non-retriable.
	Retry WorkflowExecutionActionStatusFailuresRetry `json:"retry" api:"nullable"`
	JSON  workflowExecutionActionStatusFailureJSON   `json:"-"`
}

// workflowExecutionActionStatusFailureJSON contains the JSON metadata for the
// struct [WorkflowExecutionActionStatusFailure]
type workflowExecutionActionStatusFailureJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Meta        apijson.Field
	Reason      apijson.Field
	Retry       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionActionStatusFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionStatusFailureJSON) RawJSON() string {
	return r.raw
}

// Error code identifying the type of error.
type WorkflowExecutionActionStatusFailuresCode string

const (
	WorkflowExecutionActionStatusFailuresCodeWorkflowErrorCodeUnspecified      WorkflowExecutionActionStatusFailuresCode = "WORKFLOW_ERROR_CODE_UNSPECIFIED"
	WorkflowExecutionActionStatusFailuresCodeWorkflowErrorCodeEnvironmentError WorkflowExecutionActionStatusFailuresCode = "WORKFLOW_ERROR_CODE_ENVIRONMENT_ERROR"
	WorkflowExecutionActionStatusFailuresCodeWorkflowErrorCodeAgentError       WorkflowExecutionActionStatusFailuresCode = "WORKFLOW_ERROR_CODE_AGENT_ERROR"
)

func (r WorkflowExecutionActionStatusFailuresCode) IsKnown() bool {
	switch r {
	case WorkflowExecutionActionStatusFailuresCodeWorkflowErrorCodeUnspecified, WorkflowExecutionActionStatusFailuresCodeWorkflowErrorCodeEnvironmentError, WorkflowExecutionActionStatusFailuresCodeWorkflowErrorCodeAgentError:
		return true
	}
	return false
}

// Retry configuration. If not set, the error is considered non-retriable.
type WorkflowExecutionActionStatusFailuresRetry struct {
	// Whether the error is retriable.
	Retriable bool `json:"retriable"`
	// Suggested duration to wait before retrying. Only meaningful when retriable is
	// true.
	RetryAfter string                                         `json:"retryAfter" format:"regex"`
	JSON       workflowExecutionActionStatusFailuresRetryJSON `json:"-"`
}

// workflowExecutionActionStatusFailuresRetryJSON contains the JSON metadata for
// the struct [WorkflowExecutionActionStatusFailuresRetry]
type workflowExecutionActionStatusFailuresRetryJSON struct {
	Retriable   apijson.Field
	RetryAfter  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionActionStatusFailuresRetry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionStatusFailuresRetryJSON) RawJSON() string {
	return r.raw
}

// WorkflowExecutionActionPhase defines the phases of workflow execution action.
type WorkflowExecutionActionStatusPhase string

const (
	WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseUnspecified WorkflowExecutionActionStatusPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_UNSPECIFIED"
	WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhasePending     WorkflowExecutionActionStatusPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_PENDING"
	WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseRunning     WorkflowExecutionActionStatusPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_RUNNING"
	WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseStopping    WorkflowExecutionActionStatusPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_STOPPING"
	WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseStopped     WorkflowExecutionActionStatusPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_STOPPED"
	WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseDeleting    WorkflowExecutionActionStatusPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_DELETING"
	WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseDeleted     WorkflowExecutionActionStatusPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_DELETED"
	WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseDone        WorkflowExecutionActionStatusPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_DONE"
)

func (r WorkflowExecutionActionStatusPhase) IsKnown() bool {
	switch r {
	case WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseUnspecified, WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhasePending, WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseRunning, WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseStopping, WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseStopped, WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseDeleting, WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseDeleted, WorkflowExecutionActionStatusPhaseWorkflowExecutionActionPhaseDone:
		return true
	}
	return false
}

// WorkflowExecutionActionStepStatus represents the status of a single step
// execution.
type WorkflowExecutionActionStatusStepStatus struct {
	// Structured error that caused the step to fail. Provides detailed error code,
	// message, and retry information.
	Error WorkflowExecutionActionStatusStepStatusesError `json:"error"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	FinishedAt time.Time                                      `json:"finishedAt" format:"date-time"`
	Phase      WorkflowExecutionActionStatusStepStatusesPhase `json:"phase"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// The step definition captured at execution time for immutability. This ensures
	// the UI shows the correct step even if the workflow definition changes.
	Step WorkflowStep `json:"step"`
	// Index of the step in the workflow action steps array
	StepIndex int64                                       `json:"stepIndex"`
	JSON      workflowExecutionActionStatusStepStatusJSON `json:"-"`
}

// workflowExecutionActionStatusStepStatusJSON contains the JSON metadata for the
// struct [WorkflowExecutionActionStatusStepStatus]
type workflowExecutionActionStatusStepStatusJSON struct {
	Error       apijson.Field
	FinishedAt  apijson.Field
	Phase       apijson.Field
	StartedAt   apijson.Field
	Step        apijson.Field
	StepIndex   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionActionStatusStepStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionStatusStepStatusJSON) RawJSON() string {
	return r.raw
}

// Structured error that caused the step to fail. Provides detailed error code,
// message, and retry information.
type WorkflowExecutionActionStatusStepStatusesError struct {
	// Error code identifying the type of error.
	Code WorkflowExecutionActionStatusStepStatusesErrorCode `json:"code"`
	// Human-readable error message.
	Message string `json:"message"`
	// Additional metadata about the error. Common keys include:
	//
	// - environment_id: ID of the environment
	// - task_id: ID of the task
	// - service_id: ID of the service
	// - workflow_id: ID of the workflow
	// - workflow_execution_id: ID of the workflow execution
	Meta map[string]string `json:"meta"`
	// Reason explaining why the error occurred. Examples: "not_found", "stopped",
	// "deleted", "creation_failed", "start_failed"
	Reason string `json:"reason"`
	// Retry configuration. If not set, the error is considered non-retriable.
	Retry WorkflowExecutionActionStatusStepStatusesErrorRetry `json:"retry" api:"nullable"`
	JSON  workflowExecutionActionStatusStepStatusesErrorJSON  `json:"-"`
}

// workflowExecutionActionStatusStepStatusesErrorJSON contains the JSON metadata
// for the struct [WorkflowExecutionActionStatusStepStatusesError]
type workflowExecutionActionStatusStepStatusesErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Meta        apijson.Field
	Reason      apijson.Field
	Retry       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionActionStatusStepStatusesError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionStatusStepStatusesErrorJSON) RawJSON() string {
	return r.raw
}

// Error code identifying the type of error.
type WorkflowExecutionActionStatusStepStatusesErrorCode string

const (
	WorkflowExecutionActionStatusStepStatusesErrorCodeWorkflowErrorCodeUnspecified      WorkflowExecutionActionStatusStepStatusesErrorCode = "WORKFLOW_ERROR_CODE_UNSPECIFIED"
	WorkflowExecutionActionStatusStepStatusesErrorCodeWorkflowErrorCodeEnvironmentError WorkflowExecutionActionStatusStepStatusesErrorCode = "WORKFLOW_ERROR_CODE_ENVIRONMENT_ERROR"
	WorkflowExecutionActionStatusStepStatusesErrorCodeWorkflowErrorCodeAgentError       WorkflowExecutionActionStatusStepStatusesErrorCode = "WORKFLOW_ERROR_CODE_AGENT_ERROR"
)

func (r WorkflowExecutionActionStatusStepStatusesErrorCode) IsKnown() bool {
	switch r {
	case WorkflowExecutionActionStatusStepStatusesErrorCodeWorkflowErrorCodeUnspecified, WorkflowExecutionActionStatusStepStatusesErrorCodeWorkflowErrorCodeEnvironmentError, WorkflowExecutionActionStatusStepStatusesErrorCodeWorkflowErrorCodeAgentError:
		return true
	}
	return false
}

// Retry configuration. If not set, the error is considered non-retriable.
type WorkflowExecutionActionStatusStepStatusesErrorRetry struct {
	// Whether the error is retriable.
	Retriable bool `json:"retriable"`
	// Suggested duration to wait before retrying. Only meaningful when retriable is
	// true.
	RetryAfter string                                                  `json:"retryAfter" format:"regex"`
	JSON       workflowExecutionActionStatusStepStatusesErrorRetryJSON `json:"-"`
}

// workflowExecutionActionStatusStepStatusesErrorRetryJSON contains the JSON
// metadata for the struct [WorkflowExecutionActionStatusStepStatusesErrorRetry]
type workflowExecutionActionStatusStepStatusesErrorRetryJSON struct {
	Retriable   apijson.Field
	RetryAfter  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionActionStatusStepStatusesErrorRetry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionStatusStepStatusesErrorRetryJSON) RawJSON() string {
	return r.raw
}

type WorkflowExecutionActionStatusStepStatusesPhase string

const (
	WorkflowExecutionActionStatusStepStatusesPhaseStepPhaseUnspecified WorkflowExecutionActionStatusStepStatusesPhase = "STEP_PHASE_UNSPECIFIED"
	WorkflowExecutionActionStatusStepStatusesPhaseStepPhasePending     WorkflowExecutionActionStatusStepStatusesPhase = "STEP_PHASE_PENDING"
	WorkflowExecutionActionStatusStepStatusesPhaseStepPhaseRunning     WorkflowExecutionActionStatusStepStatusesPhase = "STEP_PHASE_RUNNING"
	WorkflowExecutionActionStatusStepStatusesPhaseStepPhaseDone        WorkflowExecutionActionStatusStepStatusesPhase = "STEP_PHASE_DONE"
	WorkflowExecutionActionStatusStepStatusesPhaseStepPhaseFailed      WorkflowExecutionActionStatusStepStatusesPhase = "STEP_PHASE_FAILED"
	WorkflowExecutionActionStatusStepStatusesPhaseStepPhaseCancelled   WorkflowExecutionActionStatusStepStatusesPhase = "STEP_PHASE_CANCELLED"
)

func (r WorkflowExecutionActionStatusStepStatusesPhase) IsKnown() bool {
	switch r {
	case WorkflowExecutionActionStatusStepStatusesPhaseStepPhaseUnspecified, WorkflowExecutionActionStatusStepStatusesPhaseStepPhasePending, WorkflowExecutionActionStatusStepStatusesPhaseStepPhaseRunning, WorkflowExecutionActionStatusStepStatusesPhaseStepPhaseDone, WorkflowExecutionActionStatusStepStatusesPhaseStepPhaseFailed, WorkflowExecutionActionStatusStepStatusesPhaseStepPhaseCancelled:
		return true
	}
	return false
}

// WorkflowError provides structured error information for workflow failures. This
// enables the reconciler to make informed retry decisions and the frontend to
// display actionable error messages.
type WorkflowExecutionActionStatusWarning struct {
	// Error code identifying the type of error.
	Code WorkflowExecutionActionStatusWarningsCode `json:"code"`
	// Human-readable error message.
	Message string `json:"message"`
	// Additional metadata about the error. Common keys include:
	//
	// - environment_id: ID of the environment
	// - task_id: ID of the task
	// - service_id: ID of the service
	// - workflow_id: ID of the workflow
	// - workflow_execution_id: ID of the workflow execution
	Meta map[string]string `json:"meta"`
	// Reason explaining why the error occurred. Examples: "not_found", "stopped",
	// "deleted", "creation_failed", "start_failed"
	Reason string `json:"reason"`
	// Retry configuration. If not set, the error is considered non-retriable.
	Retry WorkflowExecutionActionStatusWarningsRetry `json:"retry" api:"nullable"`
	JSON  workflowExecutionActionStatusWarningJSON   `json:"-"`
}

// workflowExecutionActionStatusWarningJSON contains the JSON metadata for the
// struct [WorkflowExecutionActionStatusWarning]
type workflowExecutionActionStatusWarningJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Meta        apijson.Field
	Reason      apijson.Field
	Retry       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionActionStatusWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionStatusWarningJSON) RawJSON() string {
	return r.raw
}

// Error code identifying the type of error.
type WorkflowExecutionActionStatusWarningsCode string

const (
	WorkflowExecutionActionStatusWarningsCodeWorkflowErrorCodeUnspecified      WorkflowExecutionActionStatusWarningsCode = "WORKFLOW_ERROR_CODE_UNSPECIFIED"
	WorkflowExecutionActionStatusWarningsCodeWorkflowErrorCodeEnvironmentError WorkflowExecutionActionStatusWarningsCode = "WORKFLOW_ERROR_CODE_ENVIRONMENT_ERROR"
	WorkflowExecutionActionStatusWarningsCodeWorkflowErrorCodeAgentError       WorkflowExecutionActionStatusWarningsCode = "WORKFLOW_ERROR_CODE_AGENT_ERROR"
)

func (r WorkflowExecutionActionStatusWarningsCode) IsKnown() bool {
	switch r {
	case WorkflowExecutionActionStatusWarningsCodeWorkflowErrorCodeUnspecified, WorkflowExecutionActionStatusWarningsCodeWorkflowErrorCodeEnvironmentError, WorkflowExecutionActionStatusWarningsCodeWorkflowErrorCodeAgentError:
		return true
	}
	return false
}

// Retry configuration. If not set, the error is considered non-retriable.
type WorkflowExecutionActionStatusWarningsRetry struct {
	// Whether the error is retriable.
	Retriable bool `json:"retriable"`
	// Suggested duration to wait before retrying. Only meaningful when retriable is
	// true.
	RetryAfter string                                         `json:"retryAfter" format:"regex"`
	JSON       workflowExecutionActionStatusWarningsRetryJSON `json:"-"`
}

// workflowExecutionActionStatusWarningsRetryJSON contains the JSON metadata for
// the struct [WorkflowExecutionActionStatusWarningsRetry]
type workflowExecutionActionStatusWarningsRetryJSON struct {
	Retriable   apijson.Field
	RetryAfter  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowExecutionActionStatusWarningsRetry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowExecutionActionStatusWarningsRetryJSON) RawJSON() string {
	return r.raw
}

// WorkflowStep defines a single step in a workflow action.
type WorkflowStep struct {
	// WorkflowAgentStep represents an agent step that executes with a prompt.
	Agent WorkflowStepAgent `json:"agent"`
	// WorkflowPullRequestStep represents a pull request creation step.
	PullRequest WorkflowStepPullRequest `json:"pullRequest"`
	Report      WorkflowStepReport      `json:"report"`
	// WorkflowTaskStep represents a task step that executes a command.
	Task WorkflowStepTask `json:"task"`
	JSON workflowStepJSON `json:"-"`
}

// workflowStepJSON contains the JSON metadata for the struct [WorkflowStep]
type workflowStepJSON struct {
	Agent       apijson.Field
	PullRequest apijson.Field
	Report      apijson.Field
	Task        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowStepJSON) RawJSON() string {
	return r.raw
}

// WorkflowAgentStep represents an agent step that executes with a prompt.
type WorkflowStepAgent struct {
	// Prompt must be between 1 and 20,000 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 20000
	// ```
	Prompt string                `json:"prompt"`
	JSON   workflowStepAgentJSON `json:"-"`
}

// workflowStepAgentJSON contains the JSON metadata for the struct
// [WorkflowStepAgent]
type workflowStepAgentJSON struct {
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowStepAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowStepAgentJSON) RawJSON() string {
	return r.raw
}

// WorkflowPullRequestStep represents a pull request creation step.
type WorkflowStepPullRequest struct {
	// Branch name must be between 1 and 255 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 255
	// ```
	Branch string `json:"branch"`
	// Description must be at most 20,000 characters:
	//
	// ```
	// size(this) <= 20000
	// ```
	Description string `json:"description"`
	Draft       bool   `json:"draft"`
	// Title must be between 1 and 500 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 500
	// ```
	Title string                      `json:"title"`
	JSON  workflowStepPullRequestJSON `json:"-"`
}

// workflowStepPullRequestJSON contains the JSON metadata for the struct
// [WorkflowStepPullRequest]
type workflowStepPullRequestJSON struct {
	Branch      apijson.Field
	Description apijson.Field
	Draft       apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowStepPullRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowStepPullRequestJSON) RawJSON() string {
	return r.raw
}

type WorkflowStepReport struct {
	// Report must have at least one output:
	//
	// ```
	// size(this) >= 1
	// ```
	Outputs []WorkflowStepReportOutput `json:"outputs"`
	JSON    workflowStepReportJSON     `json:"-"`
}

// workflowStepReportJSON contains the JSON metadata for the struct
// [WorkflowStepReport]
type workflowStepReportJSON struct {
	Outputs     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowStepReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowStepReportJSON) RawJSON() string {
	return r.raw
}

type WorkflowStepReportOutput struct {
	JSON workflowStepReportOutputJSON `json:"-"`
}

// workflowStepReportOutputJSON contains the JSON metadata for the struct
// [WorkflowStepReportOutput]
type workflowStepReportOutputJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowStepReportOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowStepReportOutputJSON) RawJSON() string {
	return r.raw
}

// WorkflowTaskStep represents a task step that executes a command.
type WorkflowStepTask struct {
	// Command must be between 1 and 20,000 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 20000
	// ```
	Command string               `json:"command"`
	JSON    workflowStepTaskJSON `json:"-"`
}

// workflowStepTaskJSON contains the JSON metadata for the struct
// [WorkflowStepTask]
type workflowStepTaskJSON struct {
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowStepTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowStepTaskJSON) RawJSON() string {
	return r.raw
}

// WorkflowStep defines a single step in a workflow action.
type WorkflowStepParam struct {
	// WorkflowAgentStep represents an agent step that executes with a prompt.
	Agent param.Field[WorkflowStepAgentParam] `json:"agent"`
	// WorkflowPullRequestStep represents a pull request creation step.
	PullRequest param.Field[WorkflowStepPullRequestParam] `json:"pullRequest"`
	Report      param.Field[WorkflowStepReportParam]      `json:"report"`
	// WorkflowTaskStep represents a task step that executes a command.
	Task param.Field[WorkflowStepTaskParam] `json:"task"`
}

func (r WorkflowStepParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WorkflowAgentStep represents an agent step that executes with a prompt.
type WorkflowStepAgentParam struct {
	// Prompt must be between 1 and 20,000 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 20000
	// ```
	Prompt param.Field[string] `json:"prompt"`
}

func (r WorkflowStepAgentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WorkflowPullRequestStep represents a pull request creation step.
type WorkflowStepPullRequestParam struct {
	// Branch name must be between 1 and 255 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 255
	// ```
	Branch param.Field[string] `json:"branch"`
	// Description must be at most 20,000 characters:
	//
	// ```
	// size(this) <= 20000
	// ```
	Description param.Field[string] `json:"description"`
	Draft       param.Field[bool]   `json:"draft"`
	// Title must be between 1 and 500 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 500
	// ```
	Title param.Field[string] `json:"title"`
}

func (r WorkflowStepPullRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkflowStepReportParam struct {
	// Report must have at least one output:
	//
	// ```
	// size(this) >= 1
	// ```
	Outputs param.Field[[]WorkflowStepReportOutputParam] `json:"outputs"`
}

func (r WorkflowStepReportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkflowStepReportOutputParam struct {
}

func (r WorkflowStepReportOutputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WorkflowTaskStep represents a task step that executes a command.
type WorkflowStepTaskParam struct {
	// Command must be between 1 and 20,000 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 20000
	// ```
	Command param.Field[string] `json:"command"`
}

func (r WorkflowStepTaskParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WorkflowTrigger defines when a workflow should be executed.
//
// Each trigger type defines a specific condition that will cause the workflow to
// execute:
//
// - Manual: Triggered explicitly by user action via StartWorkflow RPC
// - Time: Triggered automatically based on cron schedule
// - PullRequest: Triggered automatically when specified PR events occur
//
// Trigger Semantics:
//
// - Each trigger instance can create multiple workflow executions
// - Multiple triggers of the same workflow can fire simultaneously
// - Each trigger execution is independent and tracked separately
// - Triggers are evaluated in the context specified by WorkflowTriggerContext
type WorkflowTrigger struct {
	// WorkflowTriggerContext defines the context in which a workflow should run.
	//
	// Context determines where and how the workflow executes:
	//
	// - Projects: Execute in specific project environments
	// - Repositories: Execute in environments created from repository URLs
	// - Agent: Execute in agent-managed environments with custom prompts
	// - FromTrigger: Use context derived from the trigger event (PR-specific)
	//
	// Context Usage by Trigger Type:
	//
	// - Manual: Can use any context type
	// - Time: Typically uses Projects or Repositories context
	// - PullRequest: Can use any context, FromTrigger uses PR repository context
	Context WorkflowTriggerContext `json:"context" api:"required"`
	// Manual trigger - executed when StartWorkflow RPC is called. No additional
	// configuration needed.
	Manual interface{} `json:"manual"`
	// Pull request trigger - executed when specified PR events occur. Only triggers
	// for PRs in repositories matching the trigger context.
	PullRequest WorkflowTriggerPullRequest `json:"pullRequest"`
	// Time-based trigger - executed automatically based on cron schedule. Uses
	// standard cron expression format (minute hour day month weekday).
	Time WorkflowTriggerTime `json:"time"`
	JSON workflowTriggerJSON `json:"-"`
}

// workflowTriggerJSON contains the JSON metadata for the struct [WorkflowTrigger]
type workflowTriggerJSON struct {
	Context     apijson.Field
	Manual      apijson.Field
	PullRequest apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowTriggerJSON) RawJSON() string {
	return r.raw
}

// Pull request trigger - executed when specified PR events occur. Only triggers
// for PRs in repositories matching the trigger context.
type WorkflowTriggerPullRequest struct {
	Events []WorkflowTriggerPullRequestEvent `json:"events"`
	// webhook_id is the optional ID of a webhook that this trigger is bound to. When
	// set, the trigger will be activated when the webhook receives events. This allows
	// multiple workflows to share a single webhook endpoint.
	WebhookID string                         `json:"webhookId" api:"nullable" format:"uuid"`
	JSON      workflowTriggerPullRequestJSON `json:"-"`
}

// workflowTriggerPullRequestJSON contains the JSON metadata for the struct
// [WorkflowTriggerPullRequest]
type workflowTriggerPullRequestJSON struct {
	Events      apijson.Field
	WebhookID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowTriggerPullRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowTriggerPullRequestJSON) RawJSON() string {
	return r.raw
}

// WorkflowPREvent defines pull request events that can trigger workflows.
type WorkflowTriggerPullRequestEvent string

const (
	WorkflowTriggerPullRequestEventPullRequestEventUnspecified    WorkflowTriggerPullRequestEvent = "PULL_REQUEST_EVENT_UNSPECIFIED"
	WorkflowTriggerPullRequestEventPullRequestEventOpened         WorkflowTriggerPullRequestEvent = "PULL_REQUEST_EVENT_OPENED"
	WorkflowTriggerPullRequestEventPullRequestEventUpdated        WorkflowTriggerPullRequestEvent = "PULL_REQUEST_EVENT_UPDATED"
	WorkflowTriggerPullRequestEventPullRequestEventApproved       WorkflowTriggerPullRequestEvent = "PULL_REQUEST_EVENT_APPROVED"
	WorkflowTriggerPullRequestEventPullRequestEventMerged         WorkflowTriggerPullRequestEvent = "PULL_REQUEST_EVENT_MERGED"
	WorkflowTriggerPullRequestEventPullRequestEventClosed         WorkflowTriggerPullRequestEvent = "PULL_REQUEST_EVENT_CLOSED"
	WorkflowTriggerPullRequestEventPullRequestEventReadyForReview WorkflowTriggerPullRequestEvent = "PULL_REQUEST_EVENT_READY_FOR_REVIEW"
)

func (r WorkflowTriggerPullRequestEvent) IsKnown() bool {
	switch r {
	case WorkflowTriggerPullRequestEventPullRequestEventUnspecified, WorkflowTriggerPullRequestEventPullRequestEventOpened, WorkflowTriggerPullRequestEventPullRequestEventUpdated, WorkflowTriggerPullRequestEventPullRequestEventApproved, WorkflowTriggerPullRequestEventPullRequestEventMerged, WorkflowTriggerPullRequestEventPullRequestEventClosed, WorkflowTriggerPullRequestEventPullRequestEventReadyForReview:
		return true
	}
	return false
}

// Time-based trigger - executed automatically based on cron schedule. Uses
// standard cron expression format (minute hour day month weekday).
type WorkflowTriggerTime struct {
	// Cron expression must be between 1 and 100 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 100
	// ```
	CronExpression string                  `json:"cronExpression"`
	JSON           workflowTriggerTimeJSON `json:"-"`
}

// workflowTriggerTimeJSON contains the JSON metadata for the struct
// [WorkflowTriggerTime]
type workflowTriggerTimeJSON struct {
	CronExpression apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WorkflowTriggerTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowTriggerTimeJSON) RawJSON() string {
	return r.raw
}

// WorkflowTrigger defines when a workflow should be executed.
//
// Each trigger type defines a specific condition that will cause the workflow to
// execute:
//
// - Manual: Triggered explicitly by user action via StartWorkflow RPC
// - Time: Triggered automatically based on cron schedule
// - PullRequest: Triggered automatically when specified PR events occur
//
// Trigger Semantics:
//
// - Each trigger instance can create multiple workflow executions
// - Multiple triggers of the same workflow can fire simultaneously
// - Each trigger execution is independent and tracked separately
// - Triggers are evaluated in the context specified by WorkflowTriggerContext
type WorkflowTriggerParam struct {
	// WorkflowTriggerContext defines the context in which a workflow should run.
	//
	// Context determines where and how the workflow executes:
	//
	// - Projects: Execute in specific project environments
	// - Repositories: Execute in environments created from repository URLs
	// - Agent: Execute in agent-managed environments with custom prompts
	// - FromTrigger: Use context derived from the trigger event (PR-specific)
	//
	// Context Usage by Trigger Type:
	//
	// - Manual: Can use any context type
	// - Time: Typically uses Projects or Repositories context
	// - PullRequest: Can use any context, FromTrigger uses PR repository context
	Context param.Field[WorkflowTriggerContextParam] `json:"context" api:"required"`
	// Manual trigger - executed when StartWorkflow RPC is called. No additional
	// configuration needed.
	Manual param.Field[interface{}] `json:"manual"`
	// Pull request trigger - executed when specified PR events occur. Only triggers
	// for PRs in repositories matching the trigger context.
	PullRequest param.Field[WorkflowTriggerPullRequestParam] `json:"pullRequest"`
	// Time-based trigger - executed automatically based on cron schedule. Uses
	// standard cron expression format (minute hour day month weekday).
	Time param.Field[WorkflowTriggerTimeParam] `json:"time"`
}

func (r WorkflowTriggerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Pull request trigger - executed when specified PR events occur. Only triggers
// for PRs in repositories matching the trigger context.
type WorkflowTriggerPullRequestParam struct {
	Events param.Field[[]WorkflowTriggerPullRequestEvent] `json:"events"`
	// webhook_id is the optional ID of a webhook that this trigger is bound to. When
	// set, the trigger will be activated when the webhook receives events. This allows
	// multiple workflows to share a single webhook endpoint.
	WebhookID param.Field[string] `json:"webhookId" format:"uuid"`
}

func (r WorkflowTriggerPullRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Time-based trigger - executed automatically based on cron schedule. Uses
// standard cron expression format (minute hour day month weekday).
type WorkflowTriggerTimeParam struct {
	// Cron expression must be between 1 and 100 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 100
	// ```
	CronExpression param.Field[string] `json:"cronExpression"`
}

func (r WorkflowTriggerTimeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WorkflowTriggerContext defines the context in which a workflow should run.
//
// Context determines where and how the workflow executes:
//
// - Projects: Execute in specific project environments
// - Repositories: Execute in environments created from repository URLs
// - Agent: Execute in agent-managed environments with custom prompts
// - FromTrigger: Use context derived from the trigger event (PR-specific)
//
// Context Usage by Trigger Type:
//
// - Manual: Can use any context type
// - Time: Typically uses Projects or Repositories context
// - PullRequest: Can use any context, FromTrigger uses PR repository context
type WorkflowTriggerContext struct {
	// Execute workflow in agent-managed environments. Agent receives the specified
	// prompt and manages execution context.
	Agent WorkflowTriggerContextAgent `json:"agent"`
	// Use context derived from the trigger event. Currently only supported for
	// PullRequest triggers - uses PR repository context.
	FromTrigger interface{} `json:"fromTrigger"`
	// Execute workflow in specific project environments. Creates environments for each
	// specified project.
	Projects WorkflowTriggerContextProjects `json:"projects"`
	// Execute workflow in environments created from repository URLs. Supports both
	// explicit repository URLs and search patterns.
	Repositories WorkflowTriggerContextRepositories `json:"repositories"`
	JSON         workflowTriggerContextJSON         `json:"-"`
}

// workflowTriggerContextJSON contains the JSON metadata for the struct
// [WorkflowTriggerContext]
type workflowTriggerContextJSON struct {
	Agent        apijson.Field
	FromTrigger  apijson.Field
	Projects     apijson.Field
	Repositories apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WorkflowTriggerContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowTriggerContextJSON) RawJSON() string {
	return r.raw
}

// Execute workflow in agent-managed environments. Agent receives the specified
// prompt and manages execution context.
type WorkflowTriggerContextAgent struct {
	// Agent prompt must be between 1 and 20,000 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 20000
	// ```
	Prompt string                          `json:"prompt"`
	JSON   workflowTriggerContextAgentJSON `json:"-"`
}

// workflowTriggerContextAgentJSON contains the JSON metadata for the struct
// [WorkflowTriggerContextAgent]
type workflowTriggerContextAgentJSON struct {
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowTriggerContextAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowTriggerContextAgentJSON) RawJSON() string {
	return r.raw
}

// Execute workflow in specific project environments. Creates environments for each
// specified project.
type WorkflowTriggerContextProjects struct {
	ProjectIDs []string                           `json:"projectIds" format:"uuid"`
	JSON       workflowTriggerContextProjectsJSON `json:"-"`
}

// workflowTriggerContextProjectsJSON contains the JSON metadata for the struct
// [WorkflowTriggerContextProjects]
type workflowTriggerContextProjectsJSON struct {
	ProjectIDs  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowTriggerContextProjects) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowTriggerContextProjectsJSON) RawJSON() string {
	return r.raw
}

// Execute workflow in environments created from repository URLs. Supports both
// explicit repository URLs and search patterns.
type WorkflowTriggerContextRepositories struct {
	EnvironmentClassID string `json:"environmentClassId" format:"uuid"`
	// RepositorySelector defines how to select repositories for workflow execution.
	// Combines a search string with an SCM host to identify repositories.
	RepoSelector WorkflowTriggerContextRepositoriesRepoSelector `json:"repoSelector"`
	// RepositoryURLs contains a list of explicit repository URLs. Creates one action
	// per repository URL.
	RepositoryURLs WorkflowTriggerContextRepositoriesRepositoryURLs `json:"repositoryUrls"`
	JSON           workflowTriggerContextRepositoriesJSON           `json:"-"`
}

// workflowTriggerContextRepositoriesJSON contains the JSON metadata for the struct
// [WorkflowTriggerContextRepositories]
type workflowTriggerContextRepositoriesJSON struct {
	EnvironmentClassID apijson.Field
	RepoSelector       apijson.Field
	RepositoryURLs     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkflowTriggerContextRepositories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowTriggerContextRepositoriesJSON) RawJSON() string {
	return r.raw
}

// RepositorySelector defines how to select repositories for workflow execution.
// Combines a search string with an SCM host to identify repositories.
type WorkflowTriggerContextRepositoriesRepoSelector struct {
	// Search string to match repositories using SCM-specific search patterns. For
	// GitHub: supports GitHub search syntax (e.g., "org:gitpod-io language:go",
	// "user:octocat stars:>100") For GitLab: supports GitLab search syntax See SCM
	// provider documentation for supported search patterns.
	RepoSearchString string `json:"repoSearchString"`
	// SCM host where the search should be performed (e.g., "github.com", "gitlab.com")
	ScmHost string                                             `json:"scmHost"`
	JSON    workflowTriggerContextRepositoriesRepoSelectorJSON `json:"-"`
}

// workflowTriggerContextRepositoriesRepoSelectorJSON contains the JSON metadata
// for the struct [WorkflowTriggerContextRepositoriesRepoSelector]
type workflowTriggerContextRepositoriesRepoSelectorJSON struct {
	RepoSearchString apijson.Field
	ScmHost          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkflowTriggerContextRepositoriesRepoSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowTriggerContextRepositoriesRepoSelectorJSON) RawJSON() string {
	return r.raw
}

// RepositoryURLs contains a list of explicit repository URLs. Creates one action
// per repository URL.
type WorkflowTriggerContextRepositoriesRepositoryURLs struct {
	RepoURLs []string                                             `json:"repoUrls"`
	JSON     workflowTriggerContextRepositoriesRepositoryURLsJSON `json:"-"`
}

// workflowTriggerContextRepositoriesRepositoryURLsJSON contains the JSON metadata
// for the struct [WorkflowTriggerContextRepositoriesRepositoryURLs]
type workflowTriggerContextRepositoriesRepositoryURLsJSON struct {
	RepoURLs    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowTriggerContextRepositoriesRepositoryURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowTriggerContextRepositoriesRepositoryURLsJSON) RawJSON() string {
	return r.raw
}

// WorkflowTriggerContext defines the context in which a workflow should run.
//
// Context determines where and how the workflow executes:
//
// - Projects: Execute in specific project environments
// - Repositories: Execute in environments created from repository URLs
// - Agent: Execute in agent-managed environments with custom prompts
// - FromTrigger: Use context derived from the trigger event (PR-specific)
//
// Context Usage by Trigger Type:
//
// - Manual: Can use any context type
// - Time: Typically uses Projects or Repositories context
// - PullRequest: Can use any context, FromTrigger uses PR repository context
type WorkflowTriggerContextParam struct {
	// Execute workflow in agent-managed environments. Agent receives the specified
	// prompt and manages execution context.
	Agent param.Field[WorkflowTriggerContextAgentParam] `json:"agent"`
	// Use context derived from the trigger event. Currently only supported for
	// PullRequest triggers - uses PR repository context.
	FromTrigger param.Field[interface{}] `json:"fromTrigger"`
	// Execute workflow in specific project environments. Creates environments for each
	// specified project.
	Projects param.Field[WorkflowTriggerContextProjectsParam] `json:"projects"`
	// Execute workflow in environments created from repository URLs. Supports both
	// explicit repository URLs and search patterns.
	Repositories param.Field[WorkflowTriggerContextRepositoriesParam] `json:"repositories"`
}

func (r WorkflowTriggerContextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Execute workflow in agent-managed environments. Agent receives the specified
// prompt and manages execution context.
type WorkflowTriggerContextAgentParam struct {
	// Agent prompt must be between 1 and 20,000 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 20000
	// ```
	Prompt param.Field[string] `json:"prompt"`
}

func (r WorkflowTriggerContextAgentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Execute workflow in specific project environments. Creates environments for each
// specified project.
type WorkflowTriggerContextProjectsParam struct {
	ProjectIDs param.Field[[]string] `json:"projectIds" format:"uuid"`
}

func (r WorkflowTriggerContextProjectsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Execute workflow in environments created from repository URLs. Supports both
// explicit repository URLs and search patterns.
type WorkflowTriggerContextRepositoriesParam struct {
	EnvironmentClassID param.Field[string] `json:"environmentClassId" format:"uuid"`
	// RepositorySelector defines how to select repositories for workflow execution.
	// Combines a search string with an SCM host to identify repositories.
	RepoSelector param.Field[WorkflowTriggerContextRepositoriesRepoSelectorParam] `json:"repoSelector"`
	// RepositoryURLs contains a list of explicit repository URLs. Creates one action
	// per repository URL.
	RepositoryURLs param.Field[WorkflowTriggerContextRepositoriesRepositoryURLsParam] `json:"repositoryUrls"`
}

func (r WorkflowTriggerContextRepositoriesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// RepositorySelector defines how to select repositories for workflow execution.
// Combines a search string with an SCM host to identify repositories.
type WorkflowTriggerContextRepositoriesRepoSelectorParam struct {
	// Search string to match repositories using SCM-specific search patterns. For
	// GitHub: supports GitHub search syntax (e.g., "org:gitpod-io language:go",
	// "user:octocat stars:>100") For GitLab: supports GitLab search syntax See SCM
	// provider documentation for supported search patterns.
	RepoSearchString param.Field[string] `json:"repoSearchString"`
	// SCM host where the search should be performed (e.g., "github.com", "gitlab.com")
	ScmHost param.Field[string] `json:"scmHost"`
}

func (r WorkflowTriggerContextRepositoriesRepoSelectorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// RepositoryURLs contains a list of explicit repository URLs. Creates one action
// per repository URL.
type WorkflowTriggerContextRepositoriesRepositoryURLsParam struct {
	RepoURLs param.Field[[]string] `json:"repoUrls"`
}

func (r WorkflowTriggerContextRepositoriesRepositoryURLsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationNewResponse struct {
	// Workflow represents a workflow configuration.
	Workflow Workflow                  `json:"workflow"`
	JSON     automationNewResponseJSON `json:"-"`
}

// automationNewResponseJSON contains the JSON metadata for the struct
// [AutomationNewResponse]
type automationNewResponseJSON struct {
	Workflow    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutomationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationNewResponseJSON) RawJSON() string {
	return r.raw
}

type AutomationGetResponse struct {
	// Workflow represents a workflow configuration.
	Workflow Workflow                  `json:"workflow"`
	JSON     automationGetResponseJSON `json:"-"`
}

// automationGetResponseJSON contains the JSON metadata for the struct
// [AutomationGetResponse]
type automationGetResponseJSON struct {
	Workflow    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutomationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationGetResponseJSON) RawJSON() string {
	return r.raw
}

type AutomationUpdateResponse struct {
	// Workflow represents a workflow configuration.
	Workflow Workflow                     `json:"workflow"`
	JSON     automationUpdateResponseJSON `json:"-"`
}

// automationUpdateResponseJSON contains the JSON metadata for the struct
// [AutomationUpdateResponse]
type automationUpdateResponseJSON struct {
	Workflow    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutomationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AutomationDeleteResponse = interface{}

type AutomationCancelExecutionResponse = interface{}

type AutomationCancelExecutionActionResponse = interface{}

type AutomationListExecutionOutputsResponse struct {
	ActionID string                                                 `json:"actionId"`
	Values   map[string]AutomationListExecutionOutputsResponseValue `json:"values"`
	JSON     automationListExecutionOutputsResponseJSON             `json:"-"`
}

// automationListExecutionOutputsResponseJSON contains the JSON metadata for the
// struct [AutomationListExecutionOutputsResponse]
type automationListExecutionOutputsResponseJSON struct {
	ActionID    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutomationListExecutionOutputsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationListExecutionOutputsResponseJSON) RawJSON() string {
	return r.raw
}

type AutomationListExecutionOutputsResponseValue struct {
	BoolValue   bool                                            `json:"boolValue"`
	FloatValue  float64                                         `json:"floatValue"`
	IntValue    string                                          `json:"intValue"`
	StringValue string                                          `json:"stringValue"`
	JSON        automationListExecutionOutputsResponseValueJSON `json:"-"`
}

// automationListExecutionOutputsResponseValueJSON contains the JSON metadata for
// the struct [AutomationListExecutionOutputsResponseValue]
type automationListExecutionOutputsResponseValueJSON struct {
	BoolValue   apijson.Field
	FloatValue  apijson.Field
	IntValue    apijson.Field
	StringValue apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutomationListExecutionOutputsResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationListExecutionOutputsResponseValueJSON) RawJSON() string {
	return r.raw
}

type AutomationGetExecutionResponse struct {
	// WorkflowExecution represents a workflow execution instance.
	WorkflowExecution WorkflowExecution                  `json:"workflowExecution"`
	JSON              automationGetExecutionResponseJSON `json:"-"`
}

// automationGetExecutionResponseJSON contains the JSON metadata for the struct
// [AutomationGetExecutionResponse]
type automationGetExecutionResponseJSON struct {
	WorkflowExecution apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AutomationGetExecutionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationGetExecutionResponseJSON) RawJSON() string {
	return r.raw
}

type AutomationGetExecutionActionResponse struct {
	// WorkflowExecutionAction represents a workflow execution action instance.
	WorkflowExecutionAction WorkflowExecutionAction                  `json:"workflowExecutionAction"`
	JSON                    automationGetExecutionActionResponseJSON `json:"-"`
}

// automationGetExecutionActionResponseJSON contains the JSON metadata for the
// struct [AutomationGetExecutionActionResponse]
type automationGetExecutionActionResponseJSON struct {
	WorkflowExecutionAction apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AutomationGetExecutionActionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationGetExecutionActionResponseJSON) RawJSON() string {
	return r.raw
}

type AutomationStartExecutionResponse struct {
	// WorkflowExecution represents a workflow execution instance.
	WorkflowExecution WorkflowExecution                    `json:"workflowExecution"`
	JSON              automationStartExecutionResponseJSON `json:"-"`
}

// automationStartExecutionResponseJSON contains the JSON metadata for the struct
// [AutomationStartExecutionResponse]
type automationStartExecutionResponseJSON struct {
	WorkflowExecution apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AutomationStartExecutionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationStartExecutionResponseJSON) RawJSON() string {
	return r.raw
}

type AutomationNewParams struct {
	// WorkflowAction defines the actions to be executed in a workflow.
	Action param.Field[WorkflowActionParam] `json:"action" api:"required"`
	// Description must be at most 500 characters:
	//
	// ```
	// size(this) <= 500
	// ```
	Description param.Field[string] `json:"description"`
	// Optional executor for the workflow. If not provided, defaults to the creator.
	// Must be either the caller themselves or a service account.
	Executor param.Field[shared.SubjectParam] `json:"executor"`
	// Name must be between 1 and 80 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 80
	// ```
	Name param.Field[string] `json:"name"`
	// WorkflowAction defines the actions to be executed in a workflow.
	Report param.Field[WorkflowActionParam] `json:"report"`
	// Automation must have between 1 and 10 triggers:
	//
	// ```
	// size(this) >= 1 && size(this) <= 10
	// ```
	Triggers param.Field[[]WorkflowTriggerParam] `json:"triggers"`
}

func (r AutomationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationGetParams struct {
	WorkflowID param.Field[string] `json:"workflowId" format:"uuid"`
}

func (r AutomationGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationUpdateParams struct {
	// WorkflowAction defines the actions to be executed in a workflow.
	Action param.Field[WorkflowActionParam] `json:"action"`
	// Description must be at most 500 characters:
	//
	// ```
	// size(this) <= 500
	// ```
	Description param.Field[string]              `json:"description"`
	Executor    param.Field[shared.SubjectParam] `json:"executor"`
	// Name must be between 1 and 80 characters:
	//
	// ```
	// size(this) >= 1 && size(this) <= 80
	// ```
	Name param.Field[string] `json:"name"`
	// WorkflowAction defines the actions to be executed in a workflow.
	Report param.Field[WorkflowActionParam] `json:"report"`
	// Automation can have at most 10 triggers:
	//
	// ```
	// size(this) <= 10
	// ```
	Triggers   param.Field[[]WorkflowTriggerParam] `json:"triggers"`
	WorkflowID param.Field[string]                 `json:"workflowId" format:"uuid"`
}

func (r AutomationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationListParams struct {
	Token      param.Field[string]                         `query:"token"`
	PageSize   param.Field[int64]                          `query:"pageSize"`
	Filter     param.Field[AutomationListParamsFilter]     `json:"filter"`
	Pagination param.Field[AutomationListParamsPagination] `json:"pagination"`
}

func (r AutomationListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AutomationListParams]'s query parameters as `url.Values`.
func (r AutomationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutomationListParamsFilter struct {
	// creator_ids filters workflows by creator user IDs
	CreatorIDs param.Field[[]string] `json:"creatorIds" format:"uuid"`
	// search performs case-insensitive search across workflow name, description, and
	// ID
	Search param.Field[string] `json:"search"`
	// status_phases filters workflows by the phase of their latest execution. Only
	// workflows whose most recent execution matches one of the specified phases are
	// returned.
	StatusPhases param.Field[[]AutomationListParamsFilterStatusPhase] `json:"statusPhases"`
	WorkflowIDs  param.Field[[]string]                                `json:"workflowIds" format:"uuid"`
}

func (r AutomationListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationListParamsFilterStatusPhase string

const (
	AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseUnspecified AutomationListParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_UNSPECIFIED"
	AutomationListParamsFilterStatusPhaseWorkflowExecutionPhasePending     AutomationListParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_PENDING"
	AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseRunning     AutomationListParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_RUNNING"
	AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseStopping    AutomationListParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_STOPPING"
	AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseStopped     AutomationListParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_STOPPED"
	AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseDeleting    AutomationListParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_DELETING"
	AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseDeleted     AutomationListParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_DELETED"
	AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseCompleted   AutomationListParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_COMPLETED"
)

func (r AutomationListParamsFilterStatusPhase) IsKnown() bool {
	switch r {
	case AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseUnspecified, AutomationListParamsFilterStatusPhaseWorkflowExecutionPhasePending, AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseRunning, AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseStopping, AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseStopped, AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseDeleting, AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseDeleted, AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseCompleted:
		return true
	}
	return false
}

type AutomationListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r AutomationListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationDeleteParams struct {
	// force indicates whether to immediately delete the workflow and all related
	// resources. When true, performs cascading deletion of:
	//
	//   - All workflow executions
	//   - All workflow execution actions
	//   - All environments created by workflow actions
	//   - All agent executions created by workflow actions
	//   - The workflow itself When false (default), marks workflow executions for
	//     deletion and relies on background reconciliation to clean up resources.
	Force      param.Field[bool]   `json:"force"`
	WorkflowID param.Field[string] `json:"workflowId" format:"uuid"`
}

func (r AutomationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationCancelExecutionParams struct {
	WorkflowExecutionID param.Field[string] `json:"workflowExecutionId" format:"uuid"`
}

func (r AutomationCancelExecutionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationCancelExecutionActionParams struct {
	WorkflowExecutionActionID param.Field[string] `json:"workflowExecutionActionId" format:"uuid"`
}

func (r AutomationCancelExecutionActionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationListExecutionActionsParams struct {
	Token      param.Field[string]                                         `query:"token"`
	PageSize   param.Field[int64]                                          `query:"pageSize"`
	Filter     param.Field[AutomationListExecutionActionsParamsFilter]     `json:"filter"`
	Pagination param.Field[AutomationListExecutionActionsParamsPagination] `json:"pagination"`
}

func (r AutomationListExecutionActionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AutomationListExecutionActionsParams]'s query parameters as
// `url.Values`.
func (r AutomationListExecutionActionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutomationListExecutionActionsParamsFilter struct {
	Phases                     param.Field[[]AutomationListExecutionActionsParamsFilterPhase] `json:"phases"`
	WorkflowExecutionActionIDs param.Field[[]string]                                          `json:"workflowExecutionActionIds" format:"uuid"`
	WorkflowExecutionIDs       param.Field[[]string]                                          `json:"workflowExecutionIds" format:"uuid"`
	WorkflowIDs                param.Field[[]string]                                          `json:"workflowIds" format:"uuid"`
}

func (r AutomationListExecutionActionsParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WorkflowExecutionActionPhase defines the phases of workflow execution action.
type AutomationListExecutionActionsParamsFilterPhase string

const (
	AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseUnspecified AutomationListExecutionActionsParamsFilterPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_UNSPECIFIED"
	AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhasePending     AutomationListExecutionActionsParamsFilterPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_PENDING"
	AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseRunning     AutomationListExecutionActionsParamsFilterPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_RUNNING"
	AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseStopping    AutomationListExecutionActionsParamsFilterPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_STOPPING"
	AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseStopped     AutomationListExecutionActionsParamsFilterPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_STOPPED"
	AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseDeleting    AutomationListExecutionActionsParamsFilterPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_DELETING"
	AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseDeleted     AutomationListExecutionActionsParamsFilterPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_DELETED"
	AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseDone        AutomationListExecutionActionsParamsFilterPhase = "WORKFLOW_EXECUTION_ACTION_PHASE_DONE"
)

func (r AutomationListExecutionActionsParamsFilterPhase) IsKnown() bool {
	switch r {
	case AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseUnspecified, AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhasePending, AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseRunning, AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseStopping, AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseStopped, AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseDeleting, AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseDeleted, AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseDone:
		return true
	}
	return false
}

type AutomationListExecutionActionsParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r AutomationListExecutionActionsParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationListExecutionOutputsParams struct {
	Token      param.Field[string]                                         `query:"token"`
	PageSize   param.Field[int64]                                          `query:"pageSize"`
	Filter     param.Field[AutomationListExecutionOutputsParamsFilter]     `json:"filter"`
	Pagination param.Field[AutomationListExecutionOutputsParamsPagination] `json:"pagination"`
}

func (r AutomationListExecutionOutputsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AutomationListExecutionOutputsParams]'s query parameters as
// `url.Values`.
func (r AutomationListExecutionOutputsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutomationListExecutionOutputsParamsFilter struct {
	WorkflowExecutionIDs param.Field[[]string] `json:"workflowExecutionIds" format:"uuid"`
}

func (r AutomationListExecutionOutputsParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationListExecutionOutputsParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r AutomationListExecutionOutputsParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationListExecutionsParams struct {
	Token      param.Field[string]                                   `query:"token"`
	PageSize   param.Field[int64]                                    `query:"pageSize"`
	Filter     param.Field[AutomationListExecutionsParamsFilter]     `json:"filter"`
	Pagination param.Field[AutomationListExecutionsParamsPagination] `json:"pagination"`
	// sort specifies the order of results. When unspecified, results are sorted by
	// operational priority (running first, then failed, then completed, then others).
	// Supported sort fields: startedAt, finishedAt, createdAt.
	Sort param.Field[shared.SortParam] `json:"sort"`
}

func (r AutomationListExecutionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AutomationListExecutionsParams]'s query parameters as
// `url.Values`.
func (r AutomationListExecutionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutomationListExecutionsParamsFilter struct {
	HasFailedActions param.Field[bool] `json:"hasFailedActions"`
	// search performs case-insensitive search across workflow execution ID and trigger
	// type
	Search               param.Field[string]                                            `json:"search"`
	StatusPhases         param.Field[[]AutomationListExecutionsParamsFilterStatusPhase] `json:"statusPhases"`
	WorkflowExecutionIDs param.Field[[]string]                                          `json:"workflowExecutionIds" format:"uuid"`
	WorkflowIDs          param.Field[[]string]                                          `json:"workflowIds" format:"uuid"`
}

func (r AutomationListExecutionsParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationListExecutionsParamsFilterStatusPhase string

const (
	AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseUnspecified AutomationListExecutionsParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_UNSPECIFIED"
	AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhasePending     AutomationListExecutionsParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_PENDING"
	AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseRunning     AutomationListExecutionsParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_RUNNING"
	AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseStopping    AutomationListExecutionsParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_STOPPING"
	AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseStopped     AutomationListExecutionsParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_STOPPED"
	AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseDeleting    AutomationListExecutionsParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_DELETING"
	AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseDeleted     AutomationListExecutionsParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_DELETED"
	AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseCompleted   AutomationListExecutionsParamsFilterStatusPhase = "WORKFLOW_EXECUTION_PHASE_COMPLETED"
)

func (r AutomationListExecutionsParamsFilterStatusPhase) IsKnown() bool {
	switch r {
	case AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseUnspecified, AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhasePending, AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseRunning, AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseStopping, AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseStopped, AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseDeleting, AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseDeleted, AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseCompleted:
		return true
	}
	return false
}

type AutomationListExecutionsParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r AutomationListExecutionsParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationGetExecutionParams struct {
	WorkflowExecutionID param.Field[string] `json:"workflowExecutionId" format:"uuid"`
}

func (r AutomationGetExecutionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationGetExecutionActionParams struct {
	WorkflowExecutionActionID param.Field[string] `json:"workflowExecutionActionId" format:"uuid"`
}

func (r AutomationGetExecutionActionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationStartExecutionParams struct {
	// Optional context override for the execution. When provided, replaces the
	// workflow's default trigger context. User must have appropriate permissions on
	// the overridden resources. Supports Projects, Repositories, and Agent context
	// types. FromTrigger context type is not supported for manual overrides.
	ContextOverride param.Field[WorkflowTriggerContextParam] `json:"contextOverride"`
	// Parameters to substitute into workflow steps using Go template syntax. Use
	// {{ .Parameters.key_name }} in templatable fields (task.command, agent.prompt,
	// pull*request.title/description/branch, trigger context agent.prompt). Keys must
	// match pattern ^[a-zA-Z*][a-zA-Z0-9_]\*$ Maximum 10 parameters allowed. Empty map
	// is treated as no parameters provided.
	Parameters param.Field[map[string]string] `json:"parameters"`
	WorkflowID param.Field[string]            `json:"workflowId" format:"uuid"`
}

func (r AutomationStartExecutionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
