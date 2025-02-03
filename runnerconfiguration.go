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
	"github.com/tidwall/gjson"
)

// RunnerConfigurationService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerConfigurationService] method instead.
type RunnerConfigurationService struct {
	Options                  []option.RequestOption
	EnvironmentClasses       *RunnerConfigurationEnvironmentClassService
	HostAuthenticationTokens *RunnerConfigurationHostAuthenticationTokenService
	Schema                   *RunnerConfigurationSchemaService
	ScmIntegrations          *RunnerConfigurationScmIntegrationService
}

// NewRunnerConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRunnerConfigurationService(opts ...option.RequestOption) (r *RunnerConfigurationService) {
	r = &RunnerConfigurationService{}
	r.Options = opts
	r.EnvironmentClasses = NewRunnerConfigurationEnvironmentClassService(opts...)
	r.HostAuthenticationTokens = NewRunnerConfigurationHostAuthenticationTokenService(opts...)
	r.Schema = NewRunnerConfigurationSchemaService(opts...)
	r.ScmIntegrations = NewRunnerConfigurationScmIntegrationService(opts...)
	return
}

// ValidateRunnerConfiguration validates a runner configuration (e.g. environment
// class, SCM integration)
//
// with the runner.
func (r *RunnerConfigurationService) Validate(ctx context.Context, params RunnerConfigurationValidateParams, opts ...option.RequestOption) (res *RunnerConfigurationValidateResponse, err error) {
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

type RunnerConfigurationValidateResponse struct {
	// This field can have the runtime type of
	// [RunnerConfigurationValidateResponseObjectEnvironmentClass].
	EnvironmentClass interface{} `json:"environmentClass"`
	// This field can have the runtime type of
	// [RunnerConfigurationValidateResponseObjectScmIntegration].
	ScmIntegration interface{}                             `json:"scmIntegration"`
	JSON           runnerConfigurationValidateResponseJSON `json:"-"`
	union          RunnerConfigurationValidateResponseUnion
}

// runnerConfigurationValidateResponseJSON contains the JSON metadata for the
// struct [RunnerConfigurationValidateResponse]
type runnerConfigurationValidateResponseJSON struct {
	EnvironmentClass apijson.Field
	ScmIntegration   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r runnerConfigurationValidateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationValidateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationValidateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RunnerConfigurationValidateResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationValidateResponseObject],
// [RunnerConfigurationValidateResponseObject],
// [RunnerConfigurationValidateResponseObject].
func (r RunnerConfigurationValidateResponse) AsUnion() RunnerConfigurationValidateResponseUnion {
	return r.union
}

// Union satisfied by [RunnerConfigurationValidateResponseObject],
// [RunnerConfigurationValidateResponseObject] or
// [RunnerConfigurationValidateResponseObject].
type RunnerConfigurationValidateResponseUnion interface {
	implementsRunnerConfigurationValidateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationValidateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationValidateResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationValidateResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationValidateResponseObject{}),
		},
	)
}

type RunnerConfigurationValidateResponseObject struct {
	EnvironmentClass RunnerConfigurationValidateResponseObjectEnvironmentClass `json:"environmentClass,required"`
	ScmIntegration   RunnerConfigurationValidateResponseObjectScmIntegration   `json:"scmIntegration"`
	JSON             runnerConfigurationValidateResponseObjectJSON             `json:"-"`
}

// runnerConfigurationValidateResponseObjectJSON contains the JSON metadata for the
// struct [RunnerConfigurationValidateResponseObject]
type runnerConfigurationValidateResponseObjectJSON struct {
	EnvironmentClass apijson.Field
	ScmIntegration   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationValidateResponseObject) implementsRunnerConfigurationValidateResponse() {}

type RunnerConfigurationValidateResponseObjectEnvironmentClass struct {
	JSON runnerConfigurationValidateResponseObjectEnvironmentClassJSON `json:"-"`
}

// runnerConfigurationValidateResponseObjectEnvironmentClassJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationValidateResponseObjectEnvironmentClass]
type runnerConfigurationValidateResponseObjectEnvironmentClassJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponseObjectEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseObjectEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationValidateResponseObjectScmIntegration struct {
	JSON runnerConfigurationValidateResponseObjectScmIntegrationJSON `json:"-"`
}

// runnerConfigurationValidateResponseObjectScmIntegrationJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationValidateResponseObjectScmIntegration]
type runnerConfigurationValidateResponseObjectScmIntegrationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponseObjectScmIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseObjectScmIntegrationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationValidateParams struct {
	Body RunnerConfigurationValidateParamsBodyUnion `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationValidateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationValidateParamsBody struct {
	EnvironmentClass param.Field[interface{}] `json:"environmentClass"`
	RunnerID         param.Field[string]      `json:"runnerId" format:"uuid"`
	ScmIntegration   param.Field[interface{}] `json:"scmIntegration"`
}

func (r RunnerConfigurationValidateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationValidateParamsBody) implementsRunnerConfigurationValidateParamsBodyUnion() {
}

// Satisfied by [RunnerConfigurationValidateParamsBodyObject],
// [RunnerConfigurationValidateParamsBodyObject],
// [RunnerConfigurationValidateParamsBodyObject],
// [RunnerConfigurationValidateParamsBody].
type RunnerConfigurationValidateParamsBodyUnion interface {
	implementsRunnerConfigurationValidateParamsBodyUnion()
}

type RunnerConfigurationValidateParamsBodyObject struct {
	EnvironmentClass param.Field[RunnerConfigurationValidateParamsBodyObjectEnvironmentClass] `json:"environmentClass,required"`
	RunnerID         param.Field[string]                                                      `json:"runnerId" format:"uuid"`
	ScmIntegration   param.Field[RunnerConfigurationValidateParamsBodyObjectScmIntegration]   `json:"scmIntegration"`
}

func (r RunnerConfigurationValidateParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationValidateParamsBodyObject) implementsRunnerConfigurationValidateParamsBodyUnion() {
}

type RunnerConfigurationValidateParamsBodyObjectEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID param.Field[string] `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration param.Field[[]RunnerConfigurationValidateParamsBodyObjectEnvironmentClassConfiguration] `json:"configuration"`
	// description is a human readable description of the environment class
	Description param.Field[string] `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName param.Field[string] `json:"displayName"`
	// enabled indicates whether the environment class can be used to create
	//
	// new environments.
	Enabled param.Field[bool] `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID param.Field[string] `json:"runnerId"`
}

func (r RunnerConfigurationValidateParamsBodyObjectEnvironmentClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationValidateParamsBodyObjectEnvironmentClassConfiguration struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r RunnerConfigurationValidateParamsBodyObjectEnvironmentClassConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationValidateParamsBodyObjectScmIntegration struct {
}

func (r RunnerConfigurationValidateParamsBodyObjectScmIntegration) MarshalJSON() (data []byte, err error) {
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
