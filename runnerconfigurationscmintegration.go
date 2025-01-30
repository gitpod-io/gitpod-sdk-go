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

// RunnerConfigurationScmIntegrationService contains methods and other services
// that help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerConfigurationScmIntegrationService] method instead.
type RunnerConfigurationScmIntegrationService struct {
	Options []option.RequestOption
}

// NewRunnerConfigurationScmIntegrationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRunnerConfigurationScmIntegrationService(opts ...option.RequestOption) (r *RunnerConfigurationScmIntegrationService) {
	r = &RunnerConfigurationScmIntegrationService{}
	r.Options = opts
	return
}

// CreateSCMIntegration creates a new SCM integration on a runner.
func (r *RunnerConfigurationScmIntegrationService) New(ctx context.Context, params RunnerConfigurationScmIntegrationNewParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/CreateSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type RunnerConfigurationScmIntegrationNewResponse struct {
	// id is a uniquely generated identifier for the SCM integration
	ID   string                                           `json:"id" format:"uuid"`
	JSON runnerConfigurationScmIntegrationNewResponseJSON `json:"-"`
}

// runnerConfigurationScmIntegrationNewResponseJSON contains the JSON metadata for
// the struct [RunnerConfigurationScmIntegrationNewResponse]
type runnerConfigurationScmIntegrationNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationNewResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationNewParams struct {
	Body RunnerConfigurationScmIntegrationNewParamsBodyUnion `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationScmIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationScmIntegrationNewParamsBody struct {
	// oauth_client_id is the OAuth app's client ID, if OAuth is configured.
	//
	// If configured, oauth_plaintext_client_secret must also be set.
	OAuthClientID param.Field[string] `json:"oauthClientId"`
	// oauth_plaintext_client_secret is the OAuth app's client secret in clear text.
	//
	// This will first be encrypted with the runner's public key before being stored.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret"`
}

func (r RunnerConfigurationScmIntegrationNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationNewParamsBody) implementsRunnerConfigurationScmIntegrationNewParamsBodyUnion() {
}

// Satisfied by [RunnerConfigurationScmIntegrationNewParamsBodyOAuthClientID],
// [RunnerConfigurationScmIntegrationNewParamsBodyOAuthPlaintextClientSecret],
// [RunnerConfigurationScmIntegrationNewParamsBody].
type RunnerConfigurationScmIntegrationNewParamsBodyUnion interface {
	implementsRunnerConfigurationScmIntegrationNewParamsBodyUnion()
}

type RunnerConfigurationScmIntegrationNewParamsBodyOAuthClientID struct {
	// oauth_client_id is the OAuth app's client ID, if OAuth is configured.
	//
	// If configured, oauth_plaintext_client_secret must also be set.
	OAuthClientID param.Field[string] `json:"oauthClientId,required"`
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthClientID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthClientID) implementsRunnerConfigurationScmIntegrationNewParamsBodyUnion() {
}

type RunnerConfigurationScmIntegrationNewParamsBodyOAuthPlaintextClientSecret struct {
	// oauth_plaintext_client_secret is the OAuth app's client secret in clear text.
	//
	// This will first be encrypted with the runner's public key before being stored.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret,required"`
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthPlaintextClientSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthPlaintextClientSecret) implementsRunnerConfigurationScmIntegrationNewParamsBodyUnion() {
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion float64

const (
	RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion1 RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}
