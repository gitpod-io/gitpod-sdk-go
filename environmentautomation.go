// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"github.com/stainless-sdks/gitpod-go/option"
)

// EnvironmentAutomationService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentAutomationService] method instead.
type EnvironmentAutomationService struct {
	Options  []option.RequestOption
	Tasks    *EnvironmentAutomationTaskService
	Services *EnvironmentAutomationServiceService
}

// NewEnvironmentAutomationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnvironmentAutomationService(opts ...option.RequestOption) (r *EnvironmentAutomationService) {
	r = &EnvironmentAutomationService{}
	r.Options = opts
	r.Tasks = NewEnvironmentAutomationTaskService(opts...)
	r.Services = NewEnvironmentAutomationServiceService(opts...)
	return
}
