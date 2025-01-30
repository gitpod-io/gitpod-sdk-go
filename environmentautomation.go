// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// EnvironmentAutomationService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentAutomationService] method instead.
type EnvironmentAutomationService struct {
	Options        []option.RequestOption
	Tasks          *EnvironmentAutomationTaskService
	TaskExecutions *EnvironmentAutomationTaskExecutionService
	Services       *EnvironmentAutomationServiceService
}

// NewEnvironmentAutomationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnvironmentAutomationService(opts ...option.RequestOption) (r *EnvironmentAutomationService) {
	r = &EnvironmentAutomationService{}
	r.Options = opts
	r.Tasks = NewEnvironmentAutomationTaskService(opts...)
	r.TaskExecutions = NewEnvironmentAutomationTaskExecutionService(opts...)
	r.Services = NewEnvironmentAutomationServiceService(opts...)
	return
}

// UpsertAutomationsFile upserts the automations file for the given environment.
func (r *EnvironmentAutomationService) Upsert(ctx context.Context, body EnvironmentAutomationUpsertParams, opts ...option.RequestOption) (res *EnvironmentAutomationUpsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/UpsertAutomationsFile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EnvironmentAutomationUpsertResponse struct {
	UpdatedServiceIDs []string                                `json:"updatedServiceIds"`
	UpdatedTaskIDs    []string                                `json:"updatedTaskIds"`
	JSON              environmentAutomationUpsertResponseJSON `json:"-"`
}

// environmentAutomationUpsertResponseJSON contains the JSON metadata for the
// struct [EnvironmentAutomationUpsertResponse]
type environmentAutomationUpsertResponseJSON struct {
	UpdatedServiceIDs apijson.Field
	UpdatedTaskIDs    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentAutomationUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentAutomationUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentAutomationUpsertParams struct {
	// WARN: Do not remove any field here, as it will break reading automation yaml
	// files. We error if there are any
	//
	// unknown fields in the yaml (to ensure the yaml is correct), but would break if
	// we removed any fields. This includes marking a field as "reserved" in the proto
	// file, this will also break reading the yaml.
	AutomationsFile param.Field[EnvironmentAutomationUpsertParamsAutomationsFile] `json:"automationsFile"`
	EnvironmentID   param.Field[string]                                           `json:"environmentId" format:"uuid"`
}

func (r EnvironmentAutomationUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WARN: Do not remove any field here, as it will break reading automation yaml
// files. We error if there are any
//
// unknown fields in the yaml (to ensure the yaml is correct), but would break if
// we removed any fields. This includes marking a field as "reserved" in the proto
// file, this will also break reading the yaml.
type EnvironmentAutomationUpsertParamsAutomationsFile struct {
	Services param.Field[map[string]EnvironmentAutomationUpsertParamsAutomationsFileServices] `json:"services"`
	Tasks    param.Field[map[string]EnvironmentAutomationUpsertParamsAutomationsFileTasks]    `json:"tasks"`
}

func (r EnvironmentAutomationUpsertParamsAutomationsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationUpsertParamsAutomationsFileServices struct {
	Commands    param.Field[EnvironmentAutomationUpsertParamsAutomationsFileServicesCommands] `json:"commands"`
	Description param.Field[string]                                                           `json:"description"`
	Name        param.Field[string]                                                           `json:"name"`
	RunsOn      param.Field[EnvironmentAutomationUpsertParamsAutomationsFileServicesRunsOn]   `json:"runsOn"`
	TriggeredBy param.Field[[]string]                                                         `json:"triggeredBy"`
}

func (r EnvironmentAutomationUpsertParamsAutomationsFileServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationUpsertParamsAutomationsFileServicesCommands struct {
	// ready is an optional command that is run repeatedly until it exits with a zero
	// exit code.
	//
	// If set, the service will first go into a Starting phase, and then into a Running
	// phase once the ready command exits with a zero exit code.
	Ready param.Field[string] `json:"ready"`
	// start is the command to start and run the service. If start exits, the service
	// will transition to the following phase:
	//
	//   - Stopped: if the exit code is 0
	//   - Failed: if the exit code is not 0 If the stop command is not set, the start
	//     command will receive a SIGTERM signal when the service is requested to stop.
	//     If it does not exit within 2 minutes, it will receive a SIGKILL signal.
	Start param.Field[string] `json:"start"`
	// stop is an optional command that runs when the service is requested to stop.
	//
	// If set, instead of sending a SIGTERM signal to the start command, the stop
	// command will be run. Once the stop command exits, the start command will receive
	// a SIGKILL signal. If the stop command exits with a non-zero exit code, the
	// service will transition to the Failed phase. If the stop command does not exit
	// within 2 minutes, a SIGKILL signal will be sent to both the start and stop
	// commands.
	Stop param.Field[string] `json:"stop"`
}

func (r EnvironmentAutomationUpsertParamsAutomationsFileServicesCommands) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationUpsertParamsAutomationsFileServicesRunsOn struct {
	Docker param.Field[EnvironmentAutomationUpsertParamsAutomationsFileServicesRunsOnDocker] `json:"docker,required"`
}

func (r EnvironmentAutomationUpsertParamsAutomationsFileServicesRunsOn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationUpsertParamsAutomationsFileServicesRunsOnDocker struct {
	Environment param.Field[[]string] `json:"environment"`
	Image       param.Field[string]   `json:"image"`
}

func (r EnvironmentAutomationUpsertParamsAutomationsFileServicesRunsOnDocker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationUpsertParamsAutomationsFileTasks struct {
	Command     param.Field[string]                                                      `json:"command"`
	DependsOn   param.Field[[]string]                                                    `json:"dependsOn"`
	Description param.Field[string]                                                      `json:"description"`
	Name        param.Field[string]                                                      `json:"name"`
	RunsOn      param.Field[EnvironmentAutomationUpsertParamsAutomationsFileTasksRunsOn] `json:"runsOn"`
	TriggeredBy param.Field[[]string]                                                    `json:"triggeredBy"`
}

func (r EnvironmentAutomationUpsertParamsAutomationsFileTasks) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationUpsertParamsAutomationsFileTasksRunsOn struct {
	Docker param.Field[EnvironmentAutomationUpsertParamsAutomationsFileTasksRunsOnDocker] `json:"docker,required"`
}

func (r EnvironmentAutomationUpsertParamsAutomationsFileTasksRunsOn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentAutomationUpsertParamsAutomationsFileTasksRunsOnDocker struct {
	Environment param.Field[[]string] `json:"environment"`
	Image       param.Field[string]   `json:"image"`
}

func (r EnvironmentAutomationUpsertParamsAutomationsFileTasksRunsOnDocker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
