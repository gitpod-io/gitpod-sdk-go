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
func (r *RunnerConfigurationScmIntegrationService) New(ctx context.Context, body RunnerConfigurationScmIntegrationNewParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/CreateSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
