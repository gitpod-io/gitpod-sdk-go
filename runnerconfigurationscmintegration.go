// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
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

// GetSCMIntegration returns a single SCM integration configured for a runner.
func (r *RunnerConfigurationScmIntegrationService) Get(ctx context.Context, params RunnerConfigurationScmIntegrationGetParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateSCMIntegration updates an existing SCM integration on a runner.
func (r *RunnerConfigurationScmIntegrationService) Update(ctx context.Context, params RunnerConfigurationScmIntegrationUpdateParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/UpdateSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListSCMIntegrations returns all SCM integrations configured for a runner.
func (r *RunnerConfigurationScmIntegrationService) List(ctx context.Context, params RunnerConfigurationScmIntegrationListParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/ListSCMIntegrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// DeleteSCMIntegration deletes an existing SCM integration on a runner.
func (r *RunnerConfigurationScmIntegrationService) Delete(ctx context.Context, params RunnerConfigurationScmIntegrationDeleteParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/DeleteSCMIntegration"
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

type RunnerConfigurationScmIntegrationGetResponse struct {
	Integration RunnerConfigurationScmIntegrationGetResponseIntegration `json:"integration"`
	JSON        runnerConfigurationScmIntegrationGetResponseJSON        `json:"-"`
}

// runnerConfigurationScmIntegrationGetResponseJSON contains the JSON metadata for
// the struct [RunnerConfigurationScmIntegrationGetResponse]
type runnerConfigurationScmIntegrationGetResponseJSON struct {
	Integration apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationGetResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationGetResponseIntegration struct {
	OAuth RunnerConfigurationScmIntegrationGetResponseIntegrationOAuth `json:"oauth,required"`
	JSON  runnerConfigurationScmIntegrationGetResponseIntegrationJSON  `json:"-"`
}

// runnerConfigurationScmIntegrationGetResponseIntegrationJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationScmIntegrationGetResponseIntegration]
type runnerConfigurationScmIntegrationGetResponseIntegrationJSON struct {
	OAuth       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationGetResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationGetResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationGetResponseIntegrationOAuth struct {
	// client_id is the OAuth app's client ID in clear text.
	ClientID string `json:"clientId"`
	// encrypted_client_secret is the OAuth app's secret encrypted with the runner's
	// public key.
	EncryptedClientSecret string                                                           `json:"encryptedClientSecret" format:"byte"`
	JSON                  runnerConfigurationScmIntegrationGetResponseIntegrationOAuthJSON `json:"-"`
}

// runnerConfigurationScmIntegrationGetResponseIntegrationOAuthJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationScmIntegrationGetResponseIntegrationOAuth]
type runnerConfigurationScmIntegrationGetResponseIntegrationOAuthJSON struct {
	ClientID              apijson.Field
	EncryptedClientSecret apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationGetResponseIntegrationOAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationGetResponseIntegrationOAuthJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationUpdateResponse = interface{}

type RunnerConfigurationScmIntegrationListResponse struct {
	Integrations []RunnerConfigurationScmIntegrationListResponseIntegration `json:"integrations"`
	// pagination contains the pagination options for listing scm integrations
	Pagination RunnerConfigurationScmIntegrationListResponsePagination `json:"pagination"`
	JSON       runnerConfigurationScmIntegrationListResponseJSON       `json:"-"`
}

// runnerConfigurationScmIntegrationListResponseJSON contains the JSON metadata for
// the struct [RunnerConfigurationScmIntegrationListResponse]
type runnerConfigurationScmIntegrationListResponseJSON struct {
	Integrations apijson.Field
	Pagination   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationListResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationListResponseIntegration struct {
	OAuth RunnerConfigurationScmIntegrationListResponseIntegrationsOAuth `json:"oauth,required"`
	// id is the unique identifier of the SCM integration
	ID       string `json:"id"`
	Host     string `json:"host"`
	Pat      bool   `json:"pat"`
	RunnerID string `json:"runnerId"`
	// scm_id references the scm_id in the runner's configuration schema that this
	// integration is for
	ScmID string                                                       `json:"scmId"`
	JSON  runnerConfigurationScmIntegrationListResponseIntegrationJSON `json:"-"`
}

// runnerConfigurationScmIntegrationListResponseIntegrationJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationScmIntegrationListResponseIntegration]
type runnerConfigurationScmIntegrationListResponseIntegrationJSON struct {
	OAuth       apijson.Field
	ID          apijson.Field
	Host        apijson.Field
	Pat         apijson.Field
	RunnerID    apijson.Field
	ScmID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationListResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationListResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationListResponseIntegrationsOAuth struct {
	// client_id is the OAuth app's client ID in clear text.
	ClientID string `json:"clientId"`
	// encrypted_client_secret is the OAuth app's secret encrypted with the runner's
	// public key.
	EncryptedClientSecret string                                                             `json:"encryptedClientSecret" format:"byte"`
	JSON                  runnerConfigurationScmIntegrationListResponseIntegrationsOAuthJSON `json:"-"`
}

// runnerConfigurationScmIntegrationListResponseIntegrationsOAuthJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationScmIntegrationListResponseIntegrationsOAuth]
type runnerConfigurationScmIntegrationListResponseIntegrationsOAuthJSON struct {
	ClientID              apijson.Field
	EncryptedClientSecret apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationListResponseIntegrationsOAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationListResponseIntegrationsOAuthJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing scm integrations
type RunnerConfigurationScmIntegrationListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                                      `json:"nextToken"`
	JSON      runnerConfigurationScmIntegrationListResponsePaginationJSON `json:"-"`
}

// runnerConfigurationScmIntegrationListResponsePaginationJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationScmIntegrationListResponsePagination]
type runnerConfigurationScmIntegrationListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationDeleteResponse = interface{}

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

// Satisfied by
// [RunnerConfigurationScmIntegrationNewParamsBodyOAuthClientIDIsTheOAuthAppSClientIDIfOAuthIsConfiguredIfConfiguredOAuthPlaintextClientSecretMustAlsoBeSet],
// [RunnerConfigurationScmIntegrationNewParamsBodyOAuthPlaintextClientSecretIsTheOAuthAppSClientSecretInClearTextThisWillFirstBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored],
// [RunnerConfigurationScmIntegrationNewParamsBody].
type RunnerConfigurationScmIntegrationNewParamsBodyUnion interface {
	implementsRunnerConfigurationScmIntegrationNewParamsBodyUnion()
}

type RunnerConfigurationScmIntegrationNewParamsBodyOAuthClientIDIsTheOAuthAppSClientIDIfOAuthIsConfiguredIfConfiguredOAuthPlaintextClientSecretMustAlsoBeSet struct {
	// oauth_client_id is the OAuth app's client ID, if OAuth is configured.
	//
	// If configured, oauth_plaintext_client_secret must also be set.
	OAuthClientID param.Field[string] `json:"oauthClientId,required"`
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthClientIDIsTheOAuthAppSClientIDIfOAuthIsConfiguredIfConfiguredOAuthPlaintextClientSecretMustAlsoBeSet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthClientIDIsTheOAuthAppSClientIDIfOAuthIsConfiguredIfConfiguredOAuthPlaintextClientSecretMustAlsoBeSet) implementsRunnerConfigurationScmIntegrationNewParamsBodyUnion() {
}

type RunnerConfigurationScmIntegrationNewParamsBodyOAuthPlaintextClientSecretIsTheOAuthAppSClientSecretInClearTextThisWillFirstBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored struct {
	// oauth_plaintext_client_secret is the OAuth app's client secret in clear text.
	//
	// This will first be encrypted with the runner's public key before being stored.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret,required"`
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthPlaintextClientSecretIsTheOAuthAppSClientSecretInClearTextThisWillFirstBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthPlaintextClientSecretIsTheOAuthAppSClientSecretInClearTextThisWillFirstBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored) implementsRunnerConfigurationScmIntegrationNewParamsBodyUnion() {
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

type RunnerConfigurationScmIntegrationGetParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                           `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationScmIntegrationGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion float64

const (
	RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion1 RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationScmIntegrationUpdateParams struct {
	Body RunnerConfigurationScmIntegrationUpdateParamsBodyUnion `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationScmIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationScmIntegrationUpdateParamsBody struct {
	// oauth_client_id can be set to update the OAuth app's client ID.
	//
	// If an empty string is set, the OAuth configuration will be removed (regardless
	// of whether a client secret is set), and any existing Host Authentication Tokens
	// for the SCM integration's runner and host that were created using the OAuth app
	// will be deleted. This might lead to users being unable to access their
	// repositories until they re-authenticate.
	OAuthClientID param.Field[string] `json:"oauthClientId"`
	// oauth_plaintext_client_secret can be set to update the OAuth app's client
	// secret.
	//
	// The cleartext secret will be encrypted with the runner's public key before being
	// stored.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret"`
	// pat can be set to enable or disable Personal Access Tokens support.
	//
	// When disabling PATs, any existing Host Authentication Tokens for the SCM
	// integration's runner and host that were created using a PAT will be deleted.
	// This might lead to users being unable to access their repositories until they
	// re-authenticate.
	Pat param.Field[bool] `json:"pat"`
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBody) implementsRunnerConfigurationScmIntegrationUpdateParamsBodyUnion() {
}

// Satisfied by
// [RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthClientIDCanBeSetToUpdateTheOAuthAppSClientIDIfAnEmptyStringIsSetTheOAuthConfigurationWillBeRemovedRegardlessOfWhetherAClientSecretIsSetAndAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingTheOAuthAppWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate],
// [RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthPlaintextClientSecretCanBeSetToUpdateTheOAuthAppSClientSecretTheCleartextSecretWillBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored],
// [RunnerConfigurationScmIntegrationUpdateParamsBodyPatCanBeSetToEnableOrDisablePersonalAccessTokensSupportWhenDisablingPaTsAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingAPatWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate],
// [RunnerConfigurationScmIntegrationUpdateParamsBody].
type RunnerConfigurationScmIntegrationUpdateParamsBodyUnion interface {
	implementsRunnerConfigurationScmIntegrationUpdateParamsBodyUnion()
}

type RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthClientIDCanBeSetToUpdateTheOAuthAppSClientIDIfAnEmptyStringIsSetTheOAuthConfigurationWillBeRemovedRegardlessOfWhetherAClientSecretIsSetAndAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingTheOAuthAppWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate struct {
	// oauth_client_id can be set to update the OAuth app's client ID.
	//
	// If an empty string is set, the OAuth configuration will be removed (regardless
	// of whether a client secret is set), and any existing Host Authentication Tokens
	// for the SCM integration's runner and host that were created using the OAuth app
	// will be deleted. This might lead to users being unable to access their
	// repositories until they re-authenticate.
	OAuthClientID param.Field[string] `json:"oauthClientId,required"`
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthClientIDCanBeSetToUpdateTheOAuthAppSClientIDIfAnEmptyStringIsSetTheOAuthConfigurationWillBeRemovedRegardlessOfWhetherAClientSecretIsSetAndAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingTheOAuthAppWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthClientIDCanBeSetToUpdateTheOAuthAppSClientIDIfAnEmptyStringIsSetTheOAuthConfigurationWillBeRemovedRegardlessOfWhetherAClientSecretIsSetAndAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingTheOAuthAppWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate) implementsRunnerConfigurationScmIntegrationUpdateParamsBodyUnion() {
}

type RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthPlaintextClientSecretCanBeSetToUpdateTheOAuthAppSClientSecretTheCleartextSecretWillBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored struct {
	// oauth_plaintext_client_secret can be set to update the OAuth app's client
	// secret.
	//
	// The cleartext secret will be encrypted with the runner's public key before being
	// stored.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret,required"`
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthPlaintextClientSecretCanBeSetToUpdateTheOAuthAppSClientSecretTheCleartextSecretWillBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthPlaintextClientSecretCanBeSetToUpdateTheOAuthAppSClientSecretTheCleartextSecretWillBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored) implementsRunnerConfigurationScmIntegrationUpdateParamsBodyUnion() {
}

type RunnerConfigurationScmIntegrationUpdateParamsBodyPatCanBeSetToEnableOrDisablePersonalAccessTokensSupportWhenDisablingPaTsAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingAPatWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate struct {
	// pat can be set to enable or disable Personal Access Tokens support.
	//
	// When disabling PATs, any existing Host Authentication Tokens for the SCM
	// integration's runner and host that were created using a PAT will be deleted.
	// This might lead to users being unable to access their repositories until they
	// re-authenticate.
	Pat param.Field[bool] `json:"pat,required"`
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBodyPatCanBeSetToEnableOrDisablePersonalAccessTokensSupportWhenDisablingPaTsAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingAPatWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBodyPatCanBeSetToEnableOrDisablePersonalAccessTokensSupportWhenDisablingPaTsAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingAPatWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate) implementsRunnerConfigurationScmIntegrationUpdateParamsBodyUnion() {
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion float64

const (
	RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion1 RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationScmIntegrationListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[RunnerConfigurationScmIntegrationListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[RunnerConfigurationScmIntegrationListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[RunnerConfigurationScmIntegrationListParamsConnect] `query:"connect"`
	Message param.Field[string]                                             `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [RunnerConfigurationScmIntegrationListParams]'s query
// parameters as `url.Values`.
func (r RunnerConfigurationScmIntegrationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type RunnerConfigurationScmIntegrationListParamsEncoding string

const (
	RunnerConfigurationScmIntegrationListParamsEncodingProto RunnerConfigurationScmIntegrationListParamsEncoding = "proto"
	RunnerConfigurationScmIntegrationListParamsEncodingJson  RunnerConfigurationScmIntegrationListParamsEncoding = "json"
)

func (r RunnerConfigurationScmIntegrationListParamsEncoding) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationListParamsEncodingProto, RunnerConfigurationScmIntegrationListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion float64

const (
	RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion1 RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type RunnerConfigurationScmIntegrationListParamsCompression string

const (
	RunnerConfigurationScmIntegrationListParamsCompressionIdentity RunnerConfigurationScmIntegrationListParamsCompression = "identity"
	RunnerConfigurationScmIntegrationListParamsCompressionGzip     RunnerConfigurationScmIntegrationListParamsCompression = "gzip"
	RunnerConfigurationScmIntegrationListParamsCompressionBr       RunnerConfigurationScmIntegrationListParamsCompression = "br"
)

func (r RunnerConfigurationScmIntegrationListParamsCompression) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationListParamsCompressionIdentity, RunnerConfigurationScmIntegrationListParamsCompressionGzip, RunnerConfigurationScmIntegrationListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationListParamsConnect string

const (
	RunnerConfigurationScmIntegrationListParamsConnectV1 RunnerConfigurationScmIntegrationListParamsConnect = "v1"
)

func (r RunnerConfigurationScmIntegrationListParamsConnect) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationListParamsConnectV1:
		return true
	}
	return false
}

type RunnerConfigurationScmIntegrationDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                              `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationScmIntegrationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion float64

const (
	RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion1 RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}
