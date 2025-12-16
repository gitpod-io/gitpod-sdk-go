// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"slices"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// EnvironmentAutomationService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentAutomationService] method instead.
type EnvironmentAutomationService struct {
	Options  []option.RequestOption
	Services *EnvironmentAutomationServiceService
	Tasks    *EnvironmentAutomationTaskService
}

// NewEnvironmentAutomationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnvironmentAutomationService(opts ...option.RequestOption) (r *EnvironmentAutomationService) {
	r = &EnvironmentAutomationService{}
	r.Options = opts
	r.Services = NewEnvironmentAutomationServiceService(opts...)
	r.Tasks = NewEnvironmentAutomationTaskService(opts...)
	return
}

// Upserts the automations file for the given environment.
//
// Use this method to:
//
// - Configure environment automations
// - Update automation settings
// - Manage automation files
//
// ### Examples
//
// - Update automations file:
//
//	Updates or creates the automations configuration.
//
//	```yaml
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	automationsFile:
//	  services:
//	    web-server:
//	      name: "Web Server"
//	      description: "Development web server"
//	      commands:
//	        start: "npm run dev"
//	        ready: "curl -s http://localhost:3000"
//	      triggeredBy:
//	        - postDevcontainerStart
//	  tasks:
//	    build:
//	      name: "Build Project"
//	      description: "Builds the project artifacts"
//	      command: "npm run build"
//	      triggeredBy:
//	        - postEnvironmentStart
//	```
func (r *EnvironmentAutomationService) Upsert(ctx context.Context, body EnvironmentAutomationUpsertParams, opts ...option.RequestOption) (res *EnvironmentAutomationUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EnvironmentAutomationService/UpsertAutomationsFile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// WARN: Do not remove any field here, as it will break reading automation yaml
// files. We error if there are any unknown fields in the yaml (to ensure the yaml
// is correct), but would break if we removed any fields. This includes marking a
// field as "reserved" in the proto file, this will also break reading the yaml.
type AutomationsFileParam struct {
	Services param.Field[map[string]AutomationsFileServiceParam] `json:"services"`
	Tasks    param.Field[map[string]AutomationsFileTaskParam]    `json:"tasks"`
}

func (r AutomationsFileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationsFileServiceParam struct {
	Commands    param.Field[AutomationsFileServicesCommandsParam] `json:"commands"`
	Description param.Field[string]                               `json:"description"`
	Name        param.Field[string]                               `json:"name"`
	Role        param.Field[AutomationsFileServicesRole]          `json:"role"`
	RunsOn      param.Field[shared.RunsOnParam]                   `json:"runsOn"`
	TriggeredBy param.Field[[]AutomationsFileServicesTriggeredBy] `json:"triggeredBy"`
}

func (r AutomationsFileServiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationsFileServicesCommandsParam struct {
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

func (r AutomationsFileServicesCommandsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationsFileServicesRole string

const (
	AutomationsFileServicesRoleEmpty   AutomationsFileServicesRole = ""
	AutomationsFileServicesRoleDefault AutomationsFileServicesRole = "default"
	AutomationsFileServicesRoleEditor  AutomationsFileServicesRole = "editor"
	AutomationsFileServicesRoleAIAgent AutomationsFileServicesRole = "ai-agent"
)

func (r AutomationsFileServicesRole) IsKnown() bool {
	switch r {
	case AutomationsFileServicesRoleEmpty, AutomationsFileServicesRoleDefault, AutomationsFileServicesRoleEditor, AutomationsFileServicesRoleAIAgent:
		return true
	}
	return false
}

type AutomationsFileServicesTriggeredBy string

const (
	AutomationsFileServicesTriggeredByManual                AutomationsFileServicesTriggeredBy = "manual"
	AutomationsFileServicesTriggeredByPostEnvironmentStart  AutomationsFileServicesTriggeredBy = "postEnvironmentStart"
	AutomationsFileServicesTriggeredByPostDevcontainerStart AutomationsFileServicesTriggeredBy = "postDevcontainerStart"
)

func (r AutomationsFileServicesTriggeredBy) IsKnown() bool {
	switch r {
	case AutomationsFileServicesTriggeredByManual, AutomationsFileServicesTriggeredByPostEnvironmentStart, AutomationsFileServicesTriggeredByPostDevcontainerStart:
		return true
	}
	return false
}

type AutomationsFileTaskParam struct {
	Command     param.Field[string]                            `json:"command"`
	DependsOn   param.Field[[]string]                          `json:"dependsOn"`
	Description param.Field[string]                            `json:"description"`
	Name        param.Field[string]                            `json:"name"`
	RunsOn      param.Field[shared.RunsOnParam]                `json:"runsOn"`
	TriggeredBy param.Field[[]AutomationsFileTasksTriggeredBy] `json:"triggeredBy"`
}

func (r AutomationsFileTaskParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutomationsFileTasksTriggeredBy string

const (
	AutomationsFileTasksTriggeredByManual                AutomationsFileTasksTriggeredBy = "manual"
	AutomationsFileTasksTriggeredByPostEnvironmentStart  AutomationsFileTasksTriggeredBy = "postEnvironmentStart"
	AutomationsFileTasksTriggeredByPostDevcontainerStart AutomationsFileTasksTriggeredBy = "postDevcontainerStart"
	AutomationsFileTasksTriggeredByPrebuild              AutomationsFileTasksTriggeredBy = "prebuild"
)

func (r AutomationsFileTasksTriggeredBy) IsKnown() bool {
	switch r {
	case AutomationsFileTasksTriggeredByManual, AutomationsFileTasksTriggeredByPostEnvironmentStart, AutomationsFileTasksTriggeredByPostDevcontainerStart, AutomationsFileTasksTriggeredByPrebuild:
		return true
	}
	return false
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
	// files. We error if there are any unknown fields in the yaml (to ensure the yaml
	// is correct), but would break if we removed any fields. This includes marking a
	// field as "reserved" in the proto file, this will also break reading the yaml.
	AutomationsFile param.Field[AutomationsFileParam] `json:"automationsFile"`
	EnvironmentID   param.Field[string]               `json:"environmentId" format:"uuid"`
}

func (r EnvironmentAutomationUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
