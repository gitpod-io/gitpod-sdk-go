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

// AgentService contains methods and other services that help with interacting with
// the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r *AgentService) {
	r = &AgentService{}
	r.Options = opts
	return
}

// Creates a token for conversation access with a specific agent run.
//
// This method generates a temporary token that can be used to securely connect to
// an ongoing agent conversation, for example in a web UI.
//
// ### Examples
//
// - Create a token to join an agent run conversation in a front-end application:
//
//	```yaml
//	agentExecutionId: "6fa1a3c7-fbb7-49d1-ba56-1890dc7c4c35"
//	```
func (r *AgentService) NewExecutionConversationToken(ctx context.Context, body AgentNewExecutionConversationTokenParams, opts ...option.RequestOption) (res *AgentNewExecutionConversationTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AgentService/CreateAgentExecutionConversationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a new prompt.
//
// Use this method to:
//
// - Define new prompts for templates or commands
// - Set up organization-wide prompt libraries
func (r *AgentService) NewPrompt(ctx context.Context, body AgentNewPromptParams, opts ...option.RequestOption) (res *AgentNewPromptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AgentService/CreatePrompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes an agent run.
//
// Use this method to:
//
// - Clean up agent runs that are no longer needed
//
// ### Examples
//
// - Delete an agent run by ID:
//
//	```yaml
//	agentExecutionId: "6fa1a3c7-fbb7-49d1-ba56-1890dc7c4c35"
//	```
func (r *AgentService) DeleteExecution(ctx context.Context, body AgentDeleteExecutionParams, opts ...option.RequestOption) (res *AgentDeleteExecutionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AgentService/DeleteAgentExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes a prompt.
//
// Use this method to:
//
// - Remove unused prompts
func (r *AgentService) DeletePrompt(ctx context.Context, body AgentDeletePromptParams, opts ...option.RequestOption) (res *AgentDeletePromptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AgentService/DeletePrompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all agent runs matching the specified filter.
//
// Use this method to track multiple agent runs and their associated resources.
// Results are ordered by their creation time with the newest first.
//
// ### Examples
//
// - List agent runs by agent ID:
//
//	```yaml
//	filter:
//	  agentIds: ["b8a64cfa-43e2-4b9d-9fb3-07edc63f5971"]
//	pagination:
//	  pageSize: 10
//	```
func (r *AgentService) ListExecutions(ctx context.Context, params AgentListExecutionsParams, opts ...option.RequestOption) (res *pagination.AgentExecutionsPage[AgentExecution], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.AgentService/ListAgentExecutions"
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

// Lists all agent runs matching the specified filter.
//
// Use this method to track multiple agent runs and their associated resources.
// Results are ordered by their creation time with the newest first.
//
// ### Examples
//
// - List agent runs by agent ID:
//
//	```yaml
//	filter:
//	  agentIds: ["b8a64cfa-43e2-4b9d-9fb3-07edc63f5971"]
//	pagination:
//	  pageSize: 10
//	```
func (r *AgentService) ListExecutionsAutoPaging(ctx context.Context, params AgentListExecutionsParams, opts ...option.RequestOption) *pagination.AgentExecutionsPageAutoPager[AgentExecution] {
	return pagination.NewAgentExecutionsPageAutoPager(r.ListExecutions(ctx, params, opts...))
}

// Lists all prompts matching the specified criteria.
//
// Use this method to find and browse prompts across your organization. Results are
// ordered by their creation time with the newest first.
//
// ### Examples
//
// - List all prompts:
//
//	Retrieves all prompts with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 10
//	```
func (r *AgentService) ListPrompts(ctx context.Context, params AgentListPromptsParams, opts ...option.RequestOption) (res *pagination.PromptsPage[Prompt], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.AgentService/ListPrompts"
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

// Lists all prompts matching the specified criteria.
//
// Use this method to find and browse prompts across your organization. Results are
// ordered by their creation time with the newest first.
//
// ### Examples
//
// - List all prompts:
//
//	Retrieves all prompts with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 10
//	```
func (r *AgentService) ListPromptsAutoPaging(ctx context.Context, params AgentListPromptsParams, opts ...option.RequestOption) *pagination.PromptsPageAutoPager[Prompt] {
	return pagination.NewPromptsPageAutoPager(r.ListPrompts(ctx, params, opts...))
}

// Gets details about a specific agent run, including its metadata, specification,
// and status (phase, error messages, and usage statistics).
//
// Use this method to:
//
// - Monitor the run's progress
// - Retrieve the agent's conversation URL
// - Check if an agent run is actively producing output
//
// ### Examples
//
// - Get agent run details by ID:
//
//	```yaml
//	agentExecutionId: "6fa1a3c7-fbb7-49d1-ba56-1890dc7c4c35"
//	```
func (r *AgentService) GetExecution(ctx context.Context, body AgentGetExecutionParams, opts ...option.RequestOption) (res *AgentGetExecutionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AgentService/GetAgentExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets details about a specific prompt including name, description, and prompt
// content.
//
// Use this method to:
//
// - Retrieve prompt details for editing
// - Get prompt content for execution
//
// ### Examples
//
// - Get prompt details:
//
//	```yaml
//	promptId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *AgentService) GetPrompt(ctx context.Context, body AgentGetPromptParams, opts ...option.RequestOption) (res *AgentGetPromptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AgentService/GetPrompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Sends user input to an active agent run.
//
// This method is used to provide interactive or conversation-based input to an
// agent. The agent can respond with output blocks containing text, file changes,
// or tool usage requests.
//
// ### Examples
//
// - Send a text message to an agent:
//
//	```yaml
//	agentExecutionId: "6fa1a3c7-fbb7-49d1-ba56-1890dc7c4c35"
//	userInput:
//	  text:
//	    content: "Generate a report based on the latest logs."
//	```
func (r *AgentService) SendToExecution(ctx context.Context, body AgentSendToExecutionParams, opts ...option.RequestOption) (res *AgentSendToExecutionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AgentService/SendToAgentExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Starts (or triggers) an agent run using a provided agent.
//
// Use this method to:
//
// - Launch an agent based on a known agent
//
// ### Examples
//
// - Start an agent with a project ID:
//
//	```yaml
//	agentId: "b8a64cfa-43e2-4b9d-9fb3-07edc63f5971"
//	codeContext:
//	  projectId: "2d22e4eb-31da-467f-882c-27e21550992f"
//	```
func (r *AgentService) StartExecution(ctx context.Context, body AgentStartExecutionParams, opts ...option.RequestOption) (res *AgentStartExecutionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AgentService/StartAgent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stops an active agent execution.
//
// Use this method to:
//
// - Stop an agent that is currently running
// - Prevent further processing or resource usage
//
// ### Examples
//
// - Stop an agent execution by ID:
//
//	```yaml
//	agentExecutionId: "6fa1a3c7-fbb7-49d1-ba56-1890dc7c4c35"
//	```
func (r *AgentService) StopExecution(ctx context.Context, body AgentStopExecutionParams, opts ...option.RequestOption) (res *AgentStopExecutionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AgentService/StopAgentExecution"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing prompt.
//
// Use this method to:
//
// - Modify prompt content or metadata
// - Change prompt type (template/command)
func (r *AgentService) UpdatePrompt(ctx context.Context, body AgentUpdatePromptParams, opts ...option.RequestOption) (res *AgentUpdatePromptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AgentService/UpdatePrompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AgentCodeContext struct {
	ContextURL    AgentCodeContextContextURL `json:"contextUrl"`
	EnvironmentID string                     `json:"environmentId" format:"uuid"`
	ProjectID     string                     `json:"projectId" format:"uuid"`
	// Pull request context - optional metadata about the PR being worked on This is
	// populated when the agent execution is triggered by a PR workflow or when
	// explicitly provided through the browser extension
	PullRequest AgentCodeContextPullRequest `json:"pullRequest,nullable"`
	JSON        agentCodeContextJSON        `json:"-"`
}

// agentCodeContextJSON contains the JSON metadata for the struct
// [AgentCodeContext]
type agentCodeContextJSON struct {
	ContextURL    apijson.Field
	EnvironmentID apijson.Field
	ProjectID     apijson.Field
	PullRequest   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgentCodeContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentCodeContextJSON) RawJSON() string {
	return r.raw
}

type AgentCodeContextContextURL struct {
	EnvironmentClassID string                         `json:"environmentClassId" format:"uuid"`
	URL                string                         `json:"url" format:"uri"`
	JSON               agentCodeContextContextURLJSON `json:"-"`
}

// agentCodeContextContextURLJSON contains the JSON metadata for the struct
// [AgentCodeContextContextURL]
type agentCodeContextContextURLJSON struct {
	EnvironmentClassID apijson.Field
	URL                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AgentCodeContextContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentCodeContextContextURLJSON) RawJSON() string {
	return r.raw
}

// Pull request context - optional metadata about the PR being worked on This is
// populated when the agent execution is triggered by a PR workflow or when
// explicitly provided through the browser extension
type AgentCodeContextPullRequest struct {
	// Unique identifier from the source system (e.g., "123" for GitHub PR #123)
	ID string `json:"id"`
	// Author name as provided by the SCM system
	Author string `json:"author"`
	// Whether this is a draft pull request
	Draft bool `json:"draft"`
	// Source branch name (the branch being merged from)
	FromBranch string `json:"fromBranch"`
	// Repository information
	Repository AgentCodeContextPullRequestRepository `json:"repository"`
	// Current state of the pull request
	State shared.State `json:"state"`
	// Pull request title
	Title string `json:"title"`
	// Target branch name (the branch being merged into)
	ToBranch string `json:"toBranch"`
	// Pull request URL (e.g., "https://github.com/owner/repo/pull/123")
	URL  string                          `json:"url"`
	JSON agentCodeContextPullRequestJSON `json:"-"`
}

// agentCodeContextPullRequestJSON contains the JSON metadata for the struct
// [AgentCodeContextPullRequest]
type agentCodeContextPullRequestJSON struct {
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

func (r *AgentCodeContextPullRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentCodeContextPullRequestJSON) RawJSON() string {
	return r.raw
}

// Repository information
type AgentCodeContextPullRequestRepository struct {
	CloneURL string                                    `json:"cloneUrl"`
	Host     string                                    `json:"host"`
	Name     string                                    `json:"name"`
	Owner    string                                    `json:"owner"`
	JSON     agentCodeContextPullRequestRepositoryJSON `json:"-"`
}

// agentCodeContextPullRequestRepositoryJSON contains the JSON metadata for the
// struct [AgentCodeContextPullRequestRepository]
type agentCodeContextPullRequestRepositoryJSON struct {
	CloneURL    apijson.Field
	Host        apijson.Field
	Name        apijson.Field
	Owner       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentCodeContextPullRequestRepository) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentCodeContextPullRequestRepositoryJSON) RawJSON() string {
	return r.raw
}

type AgentCodeContextParam struct {
	ContextURL    param.Field[AgentCodeContextContextURLParam] `json:"contextUrl"`
	EnvironmentID param.Field[string]                          `json:"environmentId" format:"uuid"`
	ProjectID     param.Field[string]                          `json:"projectId" format:"uuid"`
	// Pull request context - optional metadata about the PR being worked on This is
	// populated when the agent execution is triggered by a PR workflow or when
	// explicitly provided through the browser extension
	PullRequest param.Field[AgentCodeContextPullRequestParam] `json:"pullRequest"`
}

func (r AgentCodeContextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentCodeContextContextURLParam struct {
	EnvironmentClassID param.Field[string] `json:"environmentClassId" format:"uuid"`
	URL                param.Field[string] `json:"url" format:"uri"`
}

func (r AgentCodeContextContextURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Pull request context - optional metadata about the PR being worked on This is
// populated when the agent execution is triggered by a PR workflow or when
// explicitly provided through the browser extension
type AgentCodeContextPullRequestParam struct {
	// Unique identifier from the source system (e.g., "123" for GitHub PR #123)
	ID param.Field[string] `json:"id"`
	// Author name as provided by the SCM system
	Author param.Field[string] `json:"author"`
	// Whether this is a draft pull request
	Draft param.Field[bool] `json:"draft"`
	// Source branch name (the branch being merged from)
	FromBranch param.Field[string] `json:"fromBranch"`
	// Repository information
	Repository param.Field[AgentCodeContextPullRequestRepositoryParam] `json:"repository"`
	// Current state of the pull request
	State param.Field[shared.State] `json:"state"`
	// Pull request title
	Title param.Field[string] `json:"title"`
	// Target branch name (the branch being merged into)
	ToBranch param.Field[string] `json:"toBranch"`
	// Pull request URL (e.g., "https://github.com/owner/repo/pull/123")
	URL param.Field[string] `json:"url"`
}

func (r AgentCodeContextPullRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Repository information
type AgentCodeContextPullRequestRepositoryParam struct {
	CloneURL param.Field[string] `json:"cloneUrl"`
	Host     param.Field[string] `json:"host"`
	Name     param.Field[string] `json:"name"`
	Owner    param.Field[string] `json:"owner"`
}

func (r AgentCodeContextPullRequestRepositoryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentExecution struct {
	// ID is a unique identifier of this agent run. No other agent run with the same
	// name must be managed by this agent manager
	ID string `json:"id"`
	// Metadata is data associated with this agent that's required for other parts of
	// Gitpod to function
	Metadata AgentExecutionMetadata `json:"metadata"`
	// Spec is the configuration of the agent that's required for the runner to start
	// the agent
	Spec AgentExecutionSpec `json:"spec"`
	// Status is the current status of the agent
	Status AgentExecutionStatus `json:"status"`
	JSON   agentExecutionJSON   `json:"-"`
}

// agentExecutionJSON contains the JSON metadata for the struct [AgentExecution]
type agentExecutionJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionJSON) RawJSON() string {
	return r.raw
}

// Metadata is data associated with this agent that's required for other parts of
// Gitpod to function
type AgentExecutionMetadata struct {
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
	Name        string         `json:"name"`
	// role is the role of the agent execution
	Role AgentExecutionMetadataRole `json:"role"`
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
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// workflow_action_id is set when this agent execution was created as part of a
	// workflow. Used to correlate agent executions with their parent workflow
	// execution action.
	WorkflowActionID string                     `json:"workflowActionId,nullable" format:"uuid"`
	JSON             agentExecutionMetadataJSON `json:"-"`
}

// agentExecutionMetadataJSON contains the JSON metadata for the struct
// [AgentExecutionMetadata]
type agentExecutionMetadataJSON struct {
	CreatedAt        apijson.Field
	Creator          apijson.Field
	Description      apijson.Field
	Name             apijson.Field
	Role             apijson.Field
	UpdatedAt        apijson.Field
	WorkflowActionID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AgentExecutionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionMetadataJSON) RawJSON() string {
	return r.raw
}

// role is the role of the agent execution
type AgentExecutionMetadataRole string

const (
	AgentExecutionMetadataRoleAgentExecutionRoleUnspecified AgentExecutionMetadataRole = "AGENT_EXECUTION_ROLE_UNSPECIFIED"
	AgentExecutionMetadataRoleAgentExecutionRoleDefault     AgentExecutionMetadataRole = "AGENT_EXECUTION_ROLE_DEFAULT"
	AgentExecutionMetadataRoleAgentExecutionRoleWorkflow    AgentExecutionMetadataRole = "AGENT_EXECUTION_ROLE_WORKFLOW"
)

func (r AgentExecutionMetadataRole) IsKnown() bool {
	switch r {
	case AgentExecutionMetadataRoleAgentExecutionRoleUnspecified, AgentExecutionMetadataRoleAgentExecutionRoleDefault, AgentExecutionMetadataRoleAgentExecutionRoleWorkflow:
		return true
	}
	return false
}

// Spec is the configuration of the agent that's required for the runner to start
// the agent
type AgentExecutionSpec struct {
	AgentID     string           `json:"agentId" format:"uuid"`
	CodeContext AgentCodeContext `json:"codeContext"`
	// desired_phase is the desired phase of the agent run
	DesiredPhase AgentExecutionSpecDesiredPhase `json:"desiredPhase"`
	Limits       AgentExecutionSpecLimits       `json:"limits"`
	Session      string                         `json:"session"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion string                 `json:"specVersion"`
	JSON        agentExecutionSpecJSON `json:"-"`
}

// agentExecutionSpecJSON contains the JSON metadata for the struct
// [AgentExecutionSpec]
type agentExecutionSpecJSON struct {
	AgentID      apijson.Field
	CodeContext  apijson.Field
	DesiredPhase apijson.Field
	Limits       apijson.Field
	Session      apijson.Field
	SpecVersion  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AgentExecutionSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionSpecJSON) RawJSON() string {
	return r.raw
}

// desired_phase is the desired phase of the agent run
type AgentExecutionSpecDesiredPhase string

const (
	AgentExecutionSpecDesiredPhasePhaseUnspecified     AgentExecutionSpecDesiredPhase = "PHASE_UNSPECIFIED"
	AgentExecutionSpecDesiredPhasePhasePending         AgentExecutionSpecDesiredPhase = "PHASE_PENDING"
	AgentExecutionSpecDesiredPhasePhaseRunning         AgentExecutionSpecDesiredPhase = "PHASE_RUNNING"
	AgentExecutionSpecDesiredPhasePhaseWaitingForInput AgentExecutionSpecDesiredPhase = "PHASE_WAITING_FOR_INPUT"
	AgentExecutionSpecDesiredPhasePhaseStopped         AgentExecutionSpecDesiredPhase = "PHASE_STOPPED"
)

func (r AgentExecutionSpecDesiredPhase) IsKnown() bool {
	switch r {
	case AgentExecutionSpecDesiredPhasePhaseUnspecified, AgentExecutionSpecDesiredPhasePhasePending, AgentExecutionSpecDesiredPhasePhaseRunning, AgentExecutionSpecDesiredPhasePhaseWaitingForInput, AgentExecutionSpecDesiredPhasePhaseStopped:
		return true
	}
	return false
}

type AgentExecutionSpecLimits struct {
	MaxInputTokens  string                       `json:"maxInputTokens"`
	MaxIterations   string                       `json:"maxIterations"`
	MaxOutputTokens string                       `json:"maxOutputTokens"`
	JSON            agentExecutionSpecLimitsJSON `json:"-"`
}

// agentExecutionSpecLimitsJSON contains the JSON metadata for the struct
// [AgentExecutionSpecLimits]
type agentExecutionSpecLimitsJSON struct {
	MaxInputTokens  apijson.Field
	MaxIterations   apijson.Field
	MaxOutputTokens apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AgentExecutionSpecLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionSpecLimitsJSON) RawJSON() string {
	return r.raw
}

// Status is the current status of the agent
type AgentExecutionStatus struct {
	CachedCreationTokensUsed string `json:"cachedCreationTokensUsed"`
	CachedInputTokensUsed    string `json:"cachedInputTokensUsed"`
	ContextWindowLength      string `json:"contextWindowLength"`
	// conversation_url is the URL to the conversation (all messages exchanged between
	// the agent and the user) of the agent run.
	ConversationURL string `json:"conversationUrl"`
	// current_activity is the current activity description of the agent execution.
	CurrentActivity string `json:"currentActivity"`
	// current_operation is the current operation of the agent execution.
	CurrentOperation AgentExecutionStatusCurrentOperation `json:"currentOperation"`
	// failure_message contains the reason the agent run failed to operate.
	FailureMessage string `json:"failureMessage"`
	// failure_reason contains a structured reason code for the failure.
	FailureReason   AgentExecutionStatusFailureReason `json:"failureReason"`
	InputTokensUsed string                            `json:"inputTokensUsed"`
	Iterations      string                            `json:"iterations"`
	// judgement is the judgement of the agent run produced by the judgement prompt.
	Judgement string `json:"judgement"`
	// mode is the current operational mode of the agent execution. This is set by the
	// agent when entering different modes (e.g., Ralph mode via /ona:ralph command).
	Mode AgentMode `json:"mode"`
	// outputs is a map of key-value pairs that can be set by the agent during
	// execution. Similar to task execution outputs, but with typed values for
	// structured data.
	Outputs          map[string]AgentExecutionStatusOutput `json:"outputs"`
	OutputTokensUsed string                                `json:"outputTokensUsed"`
	Phase            AgentExecutionStatusPhase             `json:"phase"`
	Session          string                                `json:"session"`
	// version of the status. The value of this field has no semantic meaning (e.g.
	// don't interpret it as as a timestamp), but it can be used to impose a partial
	// order. If a.status_version < b.status_version then a was the status before b.
	StatusVersion string `json:"statusVersion"`
	// supported_model is the LLM model being used by the agent execution.
	SupportedModel AgentExecutionStatusSupportedModel `json:"supportedModel"`
	// transcript_url is the URL to the LLM transcript (all messages exchanged between
	// the agent and the LLM) of the agent run.
	TranscriptURL string `json:"transcriptUrl"`
	// used_environments is the list of environments that were used by the agent
	// execution.
	UsedEnvironments []AgentExecutionStatusUsedEnvironment `json:"usedEnvironments"`
	// warning_message contains warnings, e.g. when the LLM is overloaded.
	WarningMessage string                   `json:"warningMessage"`
	JSON           agentExecutionStatusJSON `json:"-"`
}

// agentExecutionStatusJSON contains the JSON metadata for the struct
// [AgentExecutionStatus]
type agentExecutionStatusJSON struct {
	CachedCreationTokensUsed apijson.Field
	CachedInputTokensUsed    apijson.Field
	ContextWindowLength      apijson.Field
	ConversationURL          apijson.Field
	CurrentActivity          apijson.Field
	CurrentOperation         apijson.Field
	FailureMessage           apijson.Field
	FailureReason            apijson.Field
	InputTokensUsed          apijson.Field
	Iterations               apijson.Field
	Judgement                apijson.Field
	Mode                     apijson.Field
	Outputs                  apijson.Field
	OutputTokensUsed         apijson.Field
	Phase                    apijson.Field
	Session                  apijson.Field
	StatusVersion            apijson.Field
	SupportedModel           apijson.Field
	TranscriptURL            apijson.Field
	UsedEnvironments         apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AgentExecutionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionStatusJSON) RawJSON() string {
	return r.raw
}

// current_operation is the current operation of the agent execution.
type AgentExecutionStatusCurrentOperation struct {
	Llm AgentExecutionStatusCurrentOperationLlm `json:"llm"`
	// retries is the number of times the agent run has retried one or more steps
	Retries string                                      `json:"retries"`
	Session string                                      `json:"session"`
	ToolUse AgentExecutionStatusCurrentOperationToolUse `json:"toolUse"`
	JSON    agentExecutionStatusCurrentOperationJSON    `json:"-"`
}

// agentExecutionStatusCurrentOperationJSON contains the JSON metadata for the
// struct [AgentExecutionStatusCurrentOperation]
type agentExecutionStatusCurrentOperationJSON struct {
	Llm         apijson.Field
	Retries     apijson.Field
	Session     apijson.Field
	ToolUse     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentExecutionStatusCurrentOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionStatusCurrentOperationJSON) RawJSON() string {
	return r.raw
}

type AgentExecutionStatusCurrentOperationLlm struct {
	Complete bool                                        `json:"complete"`
	JSON     agentExecutionStatusCurrentOperationLlmJSON `json:"-"`
}

// agentExecutionStatusCurrentOperationLlmJSON contains the JSON metadata for the
// struct [AgentExecutionStatusCurrentOperationLlm]
type agentExecutionStatusCurrentOperationLlmJSON struct {
	Complete    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentExecutionStatusCurrentOperationLlm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionStatusCurrentOperationLlmJSON) RawJSON() string {
	return r.raw
}

type AgentExecutionStatusCurrentOperationToolUse struct {
	Complete bool                                            `json:"complete"`
	ToolName string                                          `json:"toolName"`
	JSON     agentExecutionStatusCurrentOperationToolUseJSON `json:"-"`
}

// agentExecutionStatusCurrentOperationToolUseJSON contains the JSON metadata for
// the struct [AgentExecutionStatusCurrentOperationToolUse]
type agentExecutionStatusCurrentOperationToolUseJSON struct {
	Complete    apijson.Field
	ToolName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentExecutionStatusCurrentOperationToolUse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionStatusCurrentOperationToolUseJSON) RawJSON() string {
	return r.raw
}

// failure_reason contains a structured reason code for the failure.
type AgentExecutionStatusFailureReason string

const (
	AgentExecutionStatusFailureReasonAgentExecutionFailureReasonUnspecified    AgentExecutionStatusFailureReason = "AGENT_EXECUTION_FAILURE_REASON_UNSPECIFIED"
	AgentExecutionStatusFailureReasonAgentExecutionFailureReasonEnvironment    AgentExecutionStatusFailureReason = "AGENT_EXECUTION_FAILURE_REASON_ENVIRONMENT"
	AgentExecutionStatusFailureReasonAgentExecutionFailureReasonService        AgentExecutionStatusFailureReason = "AGENT_EXECUTION_FAILURE_REASON_SERVICE"
	AgentExecutionStatusFailureReasonAgentExecutionFailureReasonLlmIntegration AgentExecutionStatusFailureReason = "AGENT_EXECUTION_FAILURE_REASON_LLM_INTEGRATION"
	AgentExecutionStatusFailureReasonAgentExecutionFailureReasonInternal       AgentExecutionStatusFailureReason = "AGENT_EXECUTION_FAILURE_REASON_INTERNAL"
	AgentExecutionStatusFailureReasonAgentExecutionFailureReasonAgentExecution AgentExecutionStatusFailureReason = "AGENT_EXECUTION_FAILURE_REASON_AGENT_EXECUTION"
)

func (r AgentExecutionStatusFailureReason) IsKnown() bool {
	switch r {
	case AgentExecutionStatusFailureReasonAgentExecutionFailureReasonUnspecified, AgentExecutionStatusFailureReasonAgentExecutionFailureReasonEnvironment, AgentExecutionStatusFailureReasonAgentExecutionFailureReasonService, AgentExecutionStatusFailureReasonAgentExecutionFailureReasonLlmIntegration, AgentExecutionStatusFailureReasonAgentExecutionFailureReasonInternal, AgentExecutionStatusFailureReasonAgentExecutionFailureReasonAgentExecution:
		return true
	}
	return false
}

type AgentExecutionStatusOutput struct {
	BoolValue   bool                           `json:"boolValue"`
	FloatValue  float64                        `json:"floatValue"`
	IntValue    string                         `json:"intValue"`
	StringValue string                         `json:"stringValue"`
	JSON        agentExecutionStatusOutputJSON `json:"-"`
}

// agentExecutionStatusOutputJSON contains the JSON metadata for the struct
// [AgentExecutionStatusOutput]
type agentExecutionStatusOutputJSON struct {
	BoolValue   apijson.Field
	FloatValue  apijson.Field
	IntValue    apijson.Field
	StringValue apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentExecutionStatusOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionStatusOutputJSON) RawJSON() string {
	return r.raw
}

type AgentExecutionStatusPhase string

const (
	AgentExecutionStatusPhasePhaseUnspecified     AgentExecutionStatusPhase = "PHASE_UNSPECIFIED"
	AgentExecutionStatusPhasePhasePending         AgentExecutionStatusPhase = "PHASE_PENDING"
	AgentExecutionStatusPhasePhaseRunning         AgentExecutionStatusPhase = "PHASE_RUNNING"
	AgentExecutionStatusPhasePhaseWaitingForInput AgentExecutionStatusPhase = "PHASE_WAITING_FOR_INPUT"
	AgentExecutionStatusPhasePhaseStopped         AgentExecutionStatusPhase = "PHASE_STOPPED"
)

func (r AgentExecutionStatusPhase) IsKnown() bool {
	switch r {
	case AgentExecutionStatusPhasePhaseUnspecified, AgentExecutionStatusPhasePhasePending, AgentExecutionStatusPhasePhaseRunning, AgentExecutionStatusPhasePhaseWaitingForInput, AgentExecutionStatusPhasePhaseStopped:
		return true
	}
	return false
}

// supported_model is the LLM model being used by the agent execution.
type AgentExecutionStatusSupportedModel string

const (
	AgentExecutionStatusSupportedModelSupportedModelUnspecified       AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_UNSPECIFIED"
	AgentExecutionStatusSupportedModelSupportedModelSonnet3_5         AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_SONNET_3_5"
	AgentExecutionStatusSupportedModelSupportedModelSonnet3_7         AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_SONNET_3_7"
	AgentExecutionStatusSupportedModelSupportedModelSonnet3_7Extended AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_SONNET_3_7_EXTENDED"
	AgentExecutionStatusSupportedModelSupportedModelSonnet4           AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_SONNET_4"
	AgentExecutionStatusSupportedModelSupportedModelSonnet4Extended   AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_SONNET_4_EXTENDED"
	AgentExecutionStatusSupportedModelSupportedModelSonnet4_5         AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_SONNET_4_5"
	AgentExecutionStatusSupportedModelSupportedModelSonnet4_5Extended AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_SONNET_4_5_EXTENDED"
	AgentExecutionStatusSupportedModelSupportedModelOpus4             AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_OPUS_4"
	AgentExecutionStatusSupportedModelSupportedModelOpus4Extended     AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_OPUS_4_EXTENDED"
	AgentExecutionStatusSupportedModelSupportedModelOpus4_5           AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_OPUS_4_5"
	AgentExecutionStatusSupportedModelSupportedModelOpus4_5Extended   AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_OPUS_4_5_EXTENDED"
	AgentExecutionStatusSupportedModelSupportedModelOpenAI4O          AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_OPENAI_4O"
	AgentExecutionStatusSupportedModelSupportedModelOpenAI4OMini      AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_OPENAI_4O_MINI"
	AgentExecutionStatusSupportedModelSupportedModelOpenAIO1          AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_OPENAI_O1"
	AgentExecutionStatusSupportedModelSupportedModelOpenAIO1Mini      AgentExecutionStatusSupportedModel = "SUPPORTED_MODEL_OPENAI_O1_MINI"
)

func (r AgentExecutionStatusSupportedModel) IsKnown() bool {
	switch r {
	case AgentExecutionStatusSupportedModelSupportedModelUnspecified, AgentExecutionStatusSupportedModelSupportedModelSonnet3_5, AgentExecutionStatusSupportedModelSupportedModelSonnet3_7, AgentExecutionStatusSupportedModelSupportedModelSonnet3_7Extended, AgentExecutionStatusSupportedModelSupportedModelSonnet4, AgentExecutionStatusSupportedModelSupportedModelSonnet4Extended, AgentExecutionStatusSupportedModelSupportedModelSonnet4_5, AgentExecutionStatusSupportedModelSupportedModelSonnet4_5Extended, AgentExecutionStatusSupportedModelSupportedModelOpus4, AgentExecutionStatusSupportedModelSupportedModelOpus4Extended, AgentExecutionStatusSupportedModelSupportedModelOpus4_5, AgentExecutionStatusSupportedModelSupportedModelOpus4_5Extended, AgentExecutionStatusSupportedModelSupportedModelOpenAI4O, AgentExecutionStatusSupportedModelSupportedModelOpenAI4OMini, AgentExecutionStatusSupportedModelSupportedModelOpenAIO1, AgentExecutionStatusSupportedModelSupportedModelOpenAIO1Mini:
		return true
	}
	return false
}

type AgentExecutionStatusUsedEnvironment struct {
	CreatedByAgent bool                                    `json:"createdByAgent"`
	EnvironmentID  string                                  `json:"environmentId" format:"uuid"`
	JSON           agentExecutionStatusUsedEnvironmentJSON `json:"-"`
}

// agentExecutionStatusUsedEnvironmentJSON contains the JSON metadata for the
// struct [AgentExecutionStatusUsedEnvironment]
type agentExecutionStatusUsedEnvironmentJSON struct {
	CreatedByAgent apijson.Field
	EnvironmentID  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AgentExecutionStatusUsedEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionStatusUsedEnvironmentJSON) RawJSON() string {
	return r.raw
}

// AgentMode defines the operational mode of an agent
type AgentMode string

const (
	AgentModeUnspecified AgentMode = "AGENT_MODE_UNSPECIFIED"
	AgentModeExecution   AgentMode = "AGENT_MODE_EXECUTION"
	AgentModePlanning    AgentMode = "AGENT_MODE_PLANNING"
	AgentModeRalph       AgentMode = "AGENT_MODE_RALPH"
	AgentModeSpec        AgentMode = "AGENT_MODE_SPEC"
)

func (r AgentMode) IsKnown() bool {
	switch r {
	case AgentModeUnspecified, AgentModeExecution, AgentModePlanning, AgentModeRalph, AgentModeSpec:
		return true
	}
	return false
}

type Prompt struct {
	ID       string         `json:"id"`
	Metadata PromptMetadata `json:"metadata"`
	Spec     PromptSpec     `json:"spec"`
	JSON     promptJSON     `json:"-"`
}

// promptJSON contains the JSON metadata for the struct [Prompt]
type promptJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Prompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptJSON) RawJSON() string {
	return r.raw
}

type PromptMetadata struct {
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
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator is the identity of the prompt creator
	Creator shared.Subject `json:"creator"`
	// description is a description of what the prompt does
	Description string `json:"description"`
	// name is the human readable name of the prompt
	Name string `json:"name"`
	// organization_id is the ID of the organization that contains the prompt
	OrganizationID string `json:"organizationId" format:"uuid"`
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
	UpdatedAt time.Time          `json:"updatedAt" format:"date-time"`
	JSON      promptMetadataJSON `json:"-"`
}

// promptMetadataJSON contains the JSON metadata for the struct [PromptMetadata]
type promptMetadataJSON struct {
	CreatedAt      apijson.Field
	Creator        apijson.Field
	Description    apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PromptMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptMetadataJSON) RawJSON() string {
	return r.raw
}

type PromptSpec struct {
	// command is the unique command string within the organization
	Command string `json:"command"`
	// is_command indicates if this prompt is a command
	IsCommand bool `json:"isCommand"`
	// is_skill indicates if this prompt is a skill (workflow instructions for agents)
	IsSkill bool `json:"isSkill"`
	// is_template indicates if this prompt is a template
	IsTemplate bool `json:"isTemplate"`
	// prompt is the content of the prompt
	Prompt string         `json:"prompt"`
	JSON   promptSpecJSON `json:"-"`
}

// promptSpecJSON contains the JSON metadata for the struct [PromptSpec]
type promptSpecJSON struct {
	Command     apijson.Field
	IsCommand   apijson.Field
	IsSkill     apijson.Field
	IsTemplate  apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptSpecJSON) RawJSON() string {
	return r.raw
}

type UserInputBlockParam struct {
	ID param.Field[string] `json:"id"`
	// Timestamp when this block was created. Used for debugging and support bundles.
	CreatedAt param.Field[time.Time] `json:"createdAt" format:"date-time"`
	// ImageInput allows sending images to the agent. Client must provide the MIME
	// type; backend validates against magic bytes.
	//
	// Deprecated: deprecated
	Image  param.Field[UserInputBlockImageParam]   `json:"image"`
	Inputs param.Field[[]UserInputBlockInputParam] `json:"inputs"`
	// Deprecated: deprecated
	Text param.Field[UserInputBlockTextParam] `json:"text"`
}

func (r UserInputBlockParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ImageInput allows sending images to the agent. Client must provide the MIME
// type; backend validates against magic bytes.
//
// Deprecated: deprecated
type UserInputBlockImageParam struct {
	// Raw image data (max 4MB). Supported formats: PNG, JPEG.
	Data     param.Field[string]                      `json:"data" format:"byte"`
	MimeType param.Field[UserInputBlockImageMimeType] `json:"mimeType"`
}

func (r UserInputBlockImageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserInputBlockImageMimeType string

const (
	UserInputBlockImageMimeTypeImagePng  UserInputBlockImageMimeType = "image/png"
	UserInputBlockImageMimeTypeImageJpeg UserInputBlockImageMimeType = "image/jpeg"
)

func (r UserInputBlockImageMimeType) IsKnown() bool {
	switch r {
	case UserInputBlockImageMimeTypeImagePng, UserInputBlockImageMimeTypeImageJpeg:
		return true
	}
	return false
}

type UserInputBlockInputParam struct {
	// ImageInput allows sending images to the agent. Client must provide the MIME
	// type; backend validates against magic bytes.
	Image param.Field[UserInputBlockInputsImageParam] `json:"image"`
	Text  param.Field[UserInputBlockInputsTextParam]  `json:"text"`
}

func (r UserInputBlockInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ImageInput allows sending images to the agent. Client must provide the MIME
// type; backend validates against magic bytes.
type UserInputBlockInputsImageParam struct {
	// Raw image data (max 4MB). Supported formats: PNG, JPEG.
	Data     param.Field[string]                            `json:"data" format:"byte"`
	MimeType param.Field[UserInputBlockInputsImageMimeType] `json:"mimeType"`
}

func (r UserInputBlockInputsImageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserInputBlockInputsImageMimeType string

const (
	UserInputBlockInputsImageMimeTypeImagePng  UserInputBlockInputsImageMimeType = "image/png"
	UserInputBlockInputsImageMimeTypeImageJpeg UserInputBlockInputsImageMimeType = "image/jpeg"
)

func (r UserInputBlockInputsImageMimeType) IsKnown() bool {
	switch r {
	case UserInputBlockInputsImageMimeTypeImagePng, UserInputBlockInputsImageMimeTypeImageJpeg:
		return true
	}
	return false
}

type UserInputBlockInputsTextParam struct {
	Content param.Field[string] `json:"content"`
}

func (r UserInputBlockInputsTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Deprecated: deprecated
type UserInputBlockTextParam struct {
	Content param.Field[string] `json:"content"`
}

func (r UserInputBlockTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentNewExecutionConversationTokenResponse struct {
	Token string                                         `json:"token"`
	JSON  agentNewExecutionConversationTokenResponseJSON `json:"-"`
}

// agentNewExecutionConversationTokenResponseJSON contains the JSON metadata for
// the struct [AgentNewExecutionConversationTokenResponse]
type agentNewExecutionConversationTokenResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentNewExecutionConversationTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentNewExecutionConversationTokenResponseJSON) RawJSON() string {
	return r.raw
}

type AgentNewPromptResponse struct {
	Prompt Prompt                     `json:"prompt"`
	JSON   agentNewPromptResponseJSON `json:"-"`
}

// agentNewPromptResponseJSON contains the JSON metadata for the struct
// [AgentNewPromptResponse]
type agentNewPromptResponseJSON struct {
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentNewPromptResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentNewPromptResponseJSON) RawJSON() string {
	return r.raw
}

type AgentDeleteExecutionResponse = interface{}

type AgentDeletePromptResponse = interface{}

type AgentGetExecutionResponse struct {
	AgentExecution AgentExecution                `json:"agentExecution"`
	JSON           agentGetExecutionResponseJSON `json:"-"`
}

// agentGetExecutionResponseJSON contains the JSON metadata for the struct
// [AgentGetExecutionResponse]
type agentGetExecutionResponseJSON struct {
	AgentExecution apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AgentGetExecutionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentGetExecutionResponseJSON) RawJSON() string {
	return r.raw
}

type AgentGetPromptResponse struct {
	Prompt Prompt                     `json:"prompt"`
	JSON   agentGetPromptResponseJSON `json:"-"`
}

// agentGetPromptResponseJSON contains the JSON metadata for the struct
// [AgentGetPromptResponse]
type agentGetPromptResponseJSON struct {
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentGetPromptResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentGetPromptResponseJSON) RawJSON() string {
	return r.raw
}

type AgentSendToExecutionResponse = interface{}

type AgentStartExecutionResponse struct {
	AgentExecutionID string                          `json:"agentExecutionId" format:"uuid"`
	JSON             agentStartExecutionResponseJSON `json:"-"`
}

// agentStartExecutionResponseJSON contains the JSON metadata for the struct
// [AgentStartExecutionResponse]
type agentStartExecutionResponseJSON struct {
	AgentExecutionID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AgentStartExecutionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentStartExecutionResponseJSON) RawJSON() string {
	return r.raw
}

type AgentStopExecutionResponse = interface{}

type AgentUpdatePromptResponse struct {
	Prompt Prompt                        `json:"prompt"`
	JSON   agentUpdatePromptResponseJSON `json:"-"`
}

// agentUpdatePromptResponseJSON contains the JSON metadata for the struct
// [AgentUpdatePromptResponse]
type agentUpdatePromptResponseJSON struct {
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentUpdatePromptResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentUpdatePromptResponseJSON) RawJSON() string {
	return r.raw
}

type AgentNewExecutionConversationTokenParams struct {
	AgentExecutionID param.Field[string] `json:"agentExecutionId" format:"uuid"`
}

func (r AgentNewExecutionConversationTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentNewPromptParams struct {
	Command     param.Field[string] `json:"command"`
	Description param.Field[string] `json:"description"`
	IsCommand   param.Field[bool]   `json:"isCommand"`
	IsSkill     param.Field[bool]   `json:"isSkill"`
	IsTemplate  param.Field[bool]   `json:"isTemplate"`
	Name        param.Field[string] `json:"name"`
	Prompt      param.Field[string] `json:"prompt"`
}

func (r AgentNewPromptParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentDeleteExecutionParams struct {
	AgentExecutionID param.Field[string] `json:"agentExecutionId" format:"uuid"`
}

func (r AgentDeleteExecutionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentDeletePromptParams struct {
	PromptID param.Field[string] `json:"promptId" format:"uuid"`
}

func (r AgentDeletePromptParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentListExecutionsParams struct {
	Token      param.Field[string]                              `query:"token"`
	PageSize   param.Field[int64]                               `query:"pageSize"`
	Filter     param.Field[AgentListExecutionsParamsFilter]     `json:"filter"`
	Pagination param.Field[AgentListExecutionsParamsPagination] `json:"pagination"`
}

func (r AgentListExecutionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AgentListExecutionsParams]'s query parameters as
// `url.Values`.
func (r AgentListExecutionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentListExecutionsParamsFilter struct {
	AgentIDs       param.Field[[]string]                                     `json:"agentIds"`
	CreatorIDs     param.Field[[]string]                                     `json:"creatorIds"`
	EnvironmentIDs param.Field[[]string]                                     `json:"environmentIds"`
	ProjectIDs     param.Field[[]string]                                     `json:"projectIds" format:"uuid"`
	Roles          param.Field[[]AgentListExecutionsParamsFilterRole]        `json:"roles"`
	StatusPhases   param.Field[[]AgentListExecutionsParamsFilterStatusPhase] `json:"statusPhases"`
}

func (r AgentListExecutionsParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AgentExecutionRole represents the role of an agent execution
type AgentListExecutionsParamsFilterRole string

const (
	AgentListExecutionsParamsFilterRoleAgentExecutionRoleUnspecified AgentListExecutionsParamsFilterRole = "AGENT_EXECUTION_ROLE_UNSPECIFIED"
	AgentListExecutionsParamsFilterRoleAgentExecutionRoleDefault     AgentListExecutionsParamsFilterRole = "AGENT_EXECUTION_ROLE_DEFAULT"
	AgentListExecutionsParamsFilterRoleAgentExecutionRoleWorkflow    AgentListExecutionsParamsFilterRole = "AGENT_EXECUTION_ROLE_WORKFLOW"
)

func (r AgentListExecutionsParamsFilterRole) IsKnown() bool {
	switch r {
	case AgentListExecutionsParamsFilterRoleAgentExecutionRoleUnspecified, AgentListExecutionsParamsFilterRoleAgentExecutionRoleDefault, AgentListExecutionsParamsFilterRoleAgentExecutionRoleWorkflow:
		return true
	}
	return false
}

type AgentListExecutionsParamsFilterStatusPhase string

const (
	AgentListExecutionsParamsFilterStatusPhasePhaseUnspecified     AgentListExecutionsParamsFilterStatusPhase = "PHASE_UNSPECIFIED"
	AgentListExecutionsParamsFilterStatusPhasePhasePending         AgentListExecutionsParamsFilterStatusPhase = "PHASE_PENDING"
	AgentListExecutionsParamsFilterStatusPhasePhaseRunning         AgentListExecutionsParamsFilterStatusPhase = "PHASE_RUNNING"
	AgentListExecutionsParamsFilterStatusPhasePhaseWaitingForInput AgentListExecutionsParamsFilterStatusPhase = "PHASE_WAITING_FOR_INPUT"
	AgentListExecutionsParamsFilterStatusPhasePhaseStopped         AgentListExecutionsParamsFilterStatusPhase = "PHASE_STOPPED"
)

func (r AgentListExecutionsParamsFilterStatusPhase) IsKnown() bool {
	switch r {
	case AgentListExecutionsParamsFilterStatusPhasePhaseUnspecified, AgentListExecutionsParamsFilterStatusPhasePhasePending, AgentListExecutionsParamsFilterStatusPhasePhaseRunning, AgentListExecutionsParamsFilterStatusPhasePhaseWaitingForInput, AgentListExecutionsParamsFilterStatusPhasePhaseStopped:
		return true
	}
	return false
}

type AgentListExecutionsParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r AgentListExecutionsParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentListPromptsParams struct {
	Token      param.Field[string]                           `query:"token"`
	PageSize   param.Field[int64]                            `query:"pageSize"`
	Filter     param.Field[AgentListPromptsParamsFilter]     `json:"filter"`
	Pagination param.Field[AgentListPromptsParamsPagination] `json:"pagination"`
}

func (r AgentListPromptsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AgentListPromptsParams]'s query parameters as `url.Values`.
func (r AgentListPromptsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentListPromptsParamsFilter struct {
	Command       param.Field[string] `json:"command"`
	CommandPrefix param.Field[string] `json:"commandPrefix"`
	IsCommand     param.Field[bool]   `json:"isCommand"`
	IsSkill       param.Field[bool]   `json:"isSkill"`
	IsTemplate    param.Field[bool]   `json:"isTemplate"`
}

func (r AgentListPromptsParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentListPromptsParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r AgentListPromptsParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentGetExecutionParams struct {
	AgentExecutionID param.Field[string] `json:"agentExecutionId" format:"uuid"`
}

func (r AgentGetExecutionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentGetPromptParams struct {
	PromptID param.Field[string] `json:"promptId" format:"uuid"`
}

func (r AgentGetPromptParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentSendToExecutionParams struct {
	AgentExecutionID param.Field[string]              `json:"agentExecutionId" format:"uuid"`
	UserInput        param.Field[UserInputBlockParam] `json:"userInput"`
}

func (r AgentSendToExecutionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentStartExecutionParams struct {
	AgentID     param.Field[string]                `json:"agentId" format:"uuid"`
	CodeContext param.Field[AgentCodeContextParam] `json:"codeContext"`
	// mode specifies the operational mode for this agent execution If not specified,
	// defaults to AGENT_MODE_EXECUTION
	Mode param.Field[AgentMode] `json:"mode"`
	Name param.Field[string]    `json:"name"`
	// workflow_action_id is an optional reference to the workflow execution action
	// that created this agent execution. Used for tracking and event correlation.
	WorkflowActionID param.Field[string] `json:"workflowActionId" format:"uuid"`
}

func (r AgentStartExecutionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentStopExecutionParams struct {
	AgentExecutionID param.Field[string] `json:"agentExecutionId" format:"uuid"`
}

func (r AgentStopExecutionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentUpdatePromptParams struct {
	// Metadata updates
	Metadata param.Field[AgentUpdatePromptParamsMetadata] `json:"metadata"`
	// The ID of the prompt to update
	PromptID param.Field[string] `json:"promptId" format:"uuid"`
	// Spec updates
	Spec param.Field[AgentUpdatePromptParamsSpec] `json:"spec"`
}

func (r AgentUpdatePromptParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metadata updates
type AgentUpdatePromptParamsMetadata struct {
	// A description of what the prompt does
	Description param.Field[string] `json:"description"`
	// The name of the prompt
	Name param.Field[string] `json:"name"`
}

func (r AgentUpdatePromptParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Spec updates
type AgentUpdatePromptParamsSpec struct {
	// The command string (unique within organization)
	Command param.Field[string] `json:"command"`
	// Whether this prompt is a command
	IsCommand param.Field[bool] `json:"isCommand"`
	// Whether this prompt is a skill
	IsSkill param.Field[bool] `json:"isSkill"`
	// Whether this prompt is a template
	IsTemplate param.Field[bool] `json:"isTemplate"`
	// The prompt content
	Prompt param.Field[string] `json:"prompt"`
}

func (r AgentUpdatePromptParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
