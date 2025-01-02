// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// RunnerConfigurationService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerConfigurationService] method instead.
type RunnerConfigurationService struct {
	Options                  []option.RequestOption
	HostAuthenticationTokens *RunnerConfigurationHostAuthenticationTokenService
	ConfigurationSchema      *RunnerConfigurationConfigurationSchemaService
	ScmIntegration           *RunnerConfigurationScmIntegrationService
	EnvironmentClasses       *RunnerConfigurationEnvironmentClassService
}

// NewRunnerConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRunnerConfigurationService(opts ...option.RequestOption) (r *RunnerConfigurationService) {
	r = &RunnerConfigurationService{}
	r.Options = opts
	r.HostAuthenticationTokens = NewRunnerConfigurationHostAuthenticationTokenService(opts...)
	r.ConfigurationSchema = NewRunnerConfigurationConfigurationSchemaService(opts...)
	r.ScmIntegration = NewRunnerConfigurationScmIntegrationService(opts...)
	r.EnvironmentClasses = NewRunnerConfigurationEnvironmentClassService(opts...)
	return
}

// ValidateRunnerConfiguration validates a runner configuration (e.g. environment
// class, SCM integration) with the runner.
func (r *RunnerConfigurationService) Validate(ctx context.Context, params RunnerConfigurationValidateParams, opts ...option.RequestOption) (res *RunnerConfigurationValidateResponseUnion, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/ValidateRunnerConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Union satisfied by [RunnerConfigurationValidateResponseUnknown],
// [RunnerConfigurationValidateResponseUnknown] or
// [RunnerConfigurationValidateResponseUnknown].
type RunnerConfigurationValidateResponseUnion interface {
	implementsRunnerConfigurationValidateResponseUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*RunnerConfigurationValidateResponseUnion)(nil)).Elem(), "")
}

type RunnerConfigurationValidateParams struct {
	Body RunnerConfigurationValidateParamsBody `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationValidateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationValidateParamsBody struct {
}

func (r RunnerConfigurationValidateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationValidateParamsConnectProtocolVersion float64

const (
	RunnerConfigurationValidateParamsConnectProtocolVersion1 RunnerConfigurationValidateParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationValidateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationValidateParamsConnectProtocolVersion1:
		return true
	}
	return false
}
