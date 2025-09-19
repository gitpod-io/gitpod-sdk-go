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

// Validates a runner configuration.
//
// Use this method to:
//
// - Check configuration validity
// - Verify integration settings
// - Validate environment classes
//
// ### Examples
//
// - Validate SCM integration:
//
//	Checks if an SCM integration is valid.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	scmIntegration:
//	  id: "integration-id"
//	  scmId: "github"
//	  host: "github.com"
//	  oauthClientId: "client_id"
//	  oauthPlaintextClientSecret: "client_secret"
//	```
func (r *RunnerConfigurationService) Validate(ctx context.Context, body RunnerConfigurationValidateParams, opts ...option.RequestOption) (res *RunnerConfigurationValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerConfigurationService/ValidateRunnerConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EnvironmentClassValidationResult struct {
	ConfigurationErrors []FieldValidationError               `json:"configurationErrors"`
	DescriptionError    string                               `json:"descriptionError,nullable"`
	DisplayNameError    string                               `json:"displayNameError,nullable"`
	Valid               bool                                 `json:"valid"`
	JSON                environmentClassValidationResultJSON `json:"-"`
}

// environmentClassValidationResultJSON contains the JSON metadata for the struct
// [EnvironmentClassValidationResult]
type environmentClassValidationResultJSON struct {
	ConfigurationErrors apijson.Field
	DescriptionError    apijson.Field
	DisplayNameError    apijson.Field
	Valid               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentClassValidationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassValidationResultJSON) RawJSON() string {
	return r.raw
}

type FieldValidationError struct {
	Error string                   `json:"error"`
	Key   string                   `json:"key"`
	JSON  fieldValidationErrorJSON `json:"-"`
}

// fieldValidationErrorJSON contains the JSON metadata for the struct
// [FieldValidationError]
type fieldValidationErrorJSON struct {
	Error       apijson.Field
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldValidationErrorJSON) RawJSON() string {
	return r.raw
}

type ScmIntegrationValidationResult struct {
	HostError  string                             `json:"hostError,nullable"`
	OAuthError string                             `json:"oauthError,nullable"`
	PatError   string                             `json:"patError,nullable"`
	ScmIDError string                             `json:"scmIdError,nullable"`
	Valid      bool                               `json:"valid"`
	JSON       scmIntegrationValidationResultJSON `json:"-"`
}

// scmIntegrationValidationResultJSON contains the JSON metadata for the struct
// [ScmIntegrationValidationResult]
type scmIntegrationValidationResultJSON struct {
	HostError   apijson.Field
	OAuthError  apijson.Field
	PatError    apijson.Field
	ScmIDError  apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScmIntegrationValidationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scmIntegrationValidationResultJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationValidateResponse struct {
	EnvironmentClass EnvironmentClassValidationResult        `json:"environmentClass"`
	ScmIntegration   ScmIntegrationValidationResult          `json:"scmIntegration"`
	JSON             runnerConfigurationValidateResponseJSON `json:"-"`
}

// runnerConfigurationValidateResponseJSON contains the JSON metadata for the
// struct [RunnerConfigurationValidateResponse]
type runnerConfigurationValidateResponseJSON struct {
	EnvironmentClass apijson.Field
	ScmIntegration   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationValidateParams struct {
	EnvironmentClass param.Field[shared.EnvironmentClassParam]                    `json:"environmentClass"`
	RunnerID         param.Field[string]                                          `json:"runnerId" format:"uuid"`
	ScmIntegration   param.Field[RunnerConfigurationValidateParamsScmIntegration] `json:"scmIntegration"`
}

func (r RunnerConfigurationValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationValidateParamsScmIntegration struct {
	// id is the unique identifier of the SCM integration
	ID   param.Field[string] `json:"id"`
	Host param.Field[string] `json:"host"`
	// issuer_url can be set to override the authentication provider URL, if it doesn't
	// match the SCM host.
	IssuerURL param.Field[string] `json:"issuerUrl"`
	// oauth_client_id is the OAuth app's client ID, if OAuth is configured. If
	// configured, oauth_client_secret must also be set.
	OAuthClientID param.Field[string] `json:"oauthClientId"`
	// oauth_encrypted_client_secret is the OAuth app's client secret encrypted with
	// the runner's public key, if OAuth is configured. This can be used to e.g.
	// validate an already encrypted client secret of an existing SCM integration.
	OAuthEncryptedClientSecret param.Field[string] `json:"oauthEncryptedClientSecret" format:"byte"`
	// oauth_plaintext_client_secret is the OAuth app's client secret in clear text, if
	// OAuth is configured. This can be set to validate any new client secret before it
	// is encrypted and stored. This value will not be stored and get encrypted with
	// the runner's public key before passing it to the runner.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret"`
	Pat                        param.Field[bool]   `json:"pat"`
	// scm_id references the scm_id in the runner's configuration schema that this
	// integration is for
	ScmID param.Field[string] `json:"scmId"`
}

func (r RunnerConfigurationValidateParamsScmIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
