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

// AutomationsFileService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutomationsFileService] method instead.
type AutomationsFileService struct {
	Options []option.RequestOption
}

// NewAutomationsFileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAutomationsFileService(opts ...option.RequestOption) (r *AutomationsFileService) {
	r = &AutomationsFileService{}
	r.Options = opts
	return
}

// UpsertAutomationsFile upserts the automations file for the given environment.
func (r *AutomationsFileService) Upsert(ctx context.Context, params AutomationsFileUpsertParams, opts ...option.RequestOption) (res *AutomationsFileUpsertResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/UpsertAutomationsFile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AutomationsFileUpsertResponse struct {
	UpdatedServiceIDs []string                          `json:"updatedServiceIds"`
	UpdatedTaskIDs    []string                          `json:"updatedTaskIds"`
	JSON              automationsFileUpsertResponseJSON `json:"-"`
}

// automationsFileUpsertResponseJSON contains the JSON metadata for the struct
// [AutomationsFileUpsertResponse]
type automationsFileUpsertResponseJSON struct {
	UpdatedServiceIDs apijson.Field
	UpdatedTaskIDs    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AutomationsFileUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automationsFileUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type AutomationsFileUpsertParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[AutomationsFileUpsertParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// WARN: Do not remove any field here, as it will break reading automation yaml
	// files. We error if there are any unknown fields in the yaml (to ensure the yaml
	// is correct), but would break if we removed any fields. This includes marking a
	// field as "reserved" in the proto file, this will also break reading the yaml.
	AutomationsFile param.Field[AutomationsFileUpsertParamsAutomationsFile] `json:"automationsFile"`
	EnvironmentID   param.Field[string]                                     `json:"environmentId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r AutomationsFileUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type AutomationsFileUpsertParamsConnectProtocolVersion float64

const (
	AutomationsFileUpsertParamsConnectProtocolVersion1 AutomationsFileUpsertParamsConnectProtocolVersion = 1
)

func (r AutomationsFileUpsertParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case AutomationsFileUpsertParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// WARN: Do not remove any field here, as it will break reading automation yaml
// files. We error if there are any unknown fields in the yaml (to ensure the yaml
// is correct), but would break if we removed any fields. This includes marking a
// field as "reserved" in the proto file, this will also break reading the yaml.
type AutomationsFileUpsertParamsAutomationsFile struct {
	Services param.Field[map[string]AutomationsFileUpsertParamsAutomationsFileServices] `json:"services"`
	Tasks    param.Field[map[string]AutomationsFileUpsertParamsAutomationsFileTasks]    `json:"tasks"`
}

func (r AutomationsFileUpsertParamsAutomationsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationsFileUpsertParamsAutomationsFileServices struct {
	Commands    param.Field[AutomationsFileUpsertParamsAutomationsFileServicesCommands]    `json:"commands"`
	Description param.Field[string]                                                        `json:"description"`
	Name        param.Field[string]                                                        `json:"name"`
	RunsOn      param.Field[AutomationsFileUpsertParamsAutomationsFileServicesRunsOnUnion] `json:"runsOn"`
	TriggeredBy param.Field[[]string]                                                      `json:"triggeredBy"`
}

func (r AutomationsFileUpsertParamsAutomationsFileServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationsFileUpsertParamsAutomationsFileServicesCommands struct {
	// ready is an optional command that is run repeatedly until it exits with a zero
	// exit code. If set, the service will first go into a Starting phase, and then
	// into a Running phase once the ready command exits with a zero exit code.
	Ready param.Field[string] `json:"ready"`
	// start is the command to start and run the service. If start exits, the service
	// will transition to the following phase:
	//
	//   - Stopped: if the exit code is 0
	//   - Failed: if the exit code is not 0 If the stop command is not set, the start
	//     command will receive a SIGTERM signal when the service is requested to stop.
	//     If it does not exit within 2 minutes, it will receive a SIGKILL signal.
	Start param.Field[string] `json:"start"`
	// stop is an optional command that runs when the service is requested to stop. If
	// set, instead of sending a SIGTERM signal to the start command, the stop command
	// will be run. Once the stop command exits, the start command will receive a
	// SIGKILL signal. If the stop command exits with a non-zero exit code, the service
	// will transition to the Failed phase. If the stop command does not exit within 2
	// minutes, a SIGKILL signal will be sent to both the start and stop commands.
	Stop param.Field[string] `json:"stop"`
}

func (r AutomationsFileUpsertParamsAutomationsFileServicesCommands) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [AutomationsFileUpsertParamsAutomationsFileServicesRunsOnUnknown],
// [AutomationsFileUpsertParamsAutomationsFileServicesRunsOnUnknown].
type AutomationsFileUpsertParamsAutomationsFileServicesRunsOnUnion interface {
	implementsAutomationsFileUpsertParamsAutomationsFileServicesRunsOnUnion()
}

type AutomationsFileUpsertParamsAutomationsFileTasks struct {
	Command     param.Field[string]                                                     `json:"command"`
	DependsOn   param.Field[[]string]                                                   `json:"dependsOn"`
	Description param.Field[string]                                                     `json:"description"`
	Name        param.Field[string]                                                     `json:"name"`
	RunsOn      param.Field[AutomationsFileUpsertParamsAutomationsFileTasksRunsOnUnion] `json:"runsOn"`
	TriggeredBy param.Field[[]string]                                                   `json:"triggeredBy"`
}

func (r AutomationsFileUpsertParamsAutomationsFileTasks) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [AutomationsFileUpsertParamsAutomationsFileTasksRunsOnUnknown],
// [AutomationsFileUpsertParamsAutomationsFileTasksRunsOnUnknown].
type AutomationsFileUpsertParamsAutomationsFileTasksRunsOnUnion interface {
	implementsAutomationsFileUpsertParamsAutomationsFileTasksRunsOnUnion()
}
