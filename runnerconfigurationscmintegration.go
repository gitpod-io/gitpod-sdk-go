// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
	"github.com/stainless-sdks/gitpod-go/packages/pagination"
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

// GetSCMIntegration returns a single SCM integration configured for a runner.
func (r *RunnerConfigurationScmIntegrationService) Get(ctx context.Context, body RunnerConfigurationScmIntegrationGetParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateSCMIntegration updates an existing SCM integration on a runner.
func (r *RunnerConfigurationScmIntegrationService) Update(ctx context.Context, body RunnerConfigurationScmIntegrationUpdateParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/UpdateSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListSCMIntegrations returns all SCM integrations configured for a runner.
func (r *RunnerConfigurationScmIntegrationService) List(ctx context.Context, params RunnerConfigurationScmIntegrationListParams, opts ...option.RequestOption) (res *pagination.IntegrationsPage[RunnerConfigurationScmIntegrationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.RunnerConfigurationService/ListSCMIntegrations"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// ListSCMIntegrations returns all SCM integrations configured for a runner.
func (r *RunnerConfigurationScmIntegrationService) ListAutoPaging(ctx context.Context, params RunnerConfigurationScmIntegrationListParams, opts ...option.RequestOption) *pagination.IntegrationsPageAutoPager[RunnerConfigurationScmIntegrationListResponse] {
	return pagination.NewIntegrationsPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteSCMIntegration deletes an existing SCM integration on a runner.
func (r *RunnerConfigurationScmIntegrationService) Delete(ctx context.Context, body RunnerConfigurationScmIntegrationDeleteParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/DeleteSCMIntegration"
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
	OAuth RunnerConfigurationScmIntegrationListResponseOAuth `json:"oauth,required"`
	// id is the unique identifier of the SCM integration
	ID       string `json:"id"`
	Host     string `json:"host"`
	Pat      bool   `json:"pat"`
	RunnerID string `json:"runnerId"`
	// scm_id references the scm_id in the runner's configuration schema that this
	// integration is for
	ScmID string                                            `json:"scmId"`
	JSON  runnerConfigurationScmIntegrationListResponseJSON `json:"-"`
}

// runnerConfigurationScmIntegrationListResponseJSON contains the JSON metadata for
// the struct [RunnerConfigurationScmIntegrationListResponse]
type runnerConfigurationScmIntegrationListResponseJSON struct {
	OAuth       apijson.Field
	ID          apijson.Field
	Host        apijson.Field
	Pat         apijson.Field
	RunnerID    apijson.Field
	ScmID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationListResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationListResponseOAuth struct {
	// client_id is the OAuth app's client ID in clear text.
	ClientID string `json:"clientId"`
	// encrypted_client_secret is the OAuth app's secret encrypted with the runner's
	// public key.
	EncryptedClientSecret string                                                 `json:"encryptedClientSecret" format:"byte"`
	JSON                  runnerConfigurationScmIntegrationListResponseOAuthJSON `json:"-"`
}

// runnerConfigurationScmIntegrationListResponseOAuthJSON contains the JSON
// metadata for the struct [RunnerConfigurationScmIntegrationListResponseOAuth]
type runnerConfigurationScmIntegrationListResponseOAuthJSON struct {
	ClientID              apijson.Field
	EncryptedClientSecret apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationListResponseOAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationListResponseOAuthJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationDeleteResponse = interface{}

type RunnerConfigurationScmIntegrationNewParams struct {
	Body RunnerConfigurationScmIntegrationNewParamsBodyUnion `json:"body,required"`
}

func (r RunnerConfigurationScmIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationScmIntegrationNewParamsBody struct {
	// oauth_client_id is the OAuth app's client ID, if OAuth is configured. If
	// configured, oauth_plaintext_client_secret must also be set.
	OAuthClientID param.Field[string] `json:"oauthClientId"`
	// oauth_plaintext_client_secret is the OAuth app's client secret in clear text.
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
	// oauth_client_id is the OAuth app's client ID, if OAuth is configured. If
	// configured, oauth_plaintext_client_secret must also be set.
	OAuthClientID param.Field[string] `json:"oauthClientId,required"`
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthClientIDIsTheOAuthAppSClientIDIfOAuthIsConfiguredIfConfiguredOAuthPlaintextClientSecretMustAlsoBeSet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthClientIDIsTheOAuthAppSClientIDIfOAuthIsConfiguredIfConfiguredOAuthPlaintextClientSecretMustAlsoBeSet) implementsRunnerConfigurationScmIntegrationNewParamsBodyUnion() {
}

type RunnerConfigurationScmIntegrationNewParamsBodyOAuthPlaintextClientSecretIsTheOAuthAppSClientSecretInClearTextThisWillFirstBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored struct {
	// oauth_plaintext_client_secret is the OAuth app's client secret in clear text.
	// This will first be encrypted with the runner's public key before being stored.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret,required"`
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthPlaintextClientSecretIsTheOAuthAppSClientSecretInClearTextThisWillFirstBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationNewParamsBodyOAuthPlaintextClientSecretIsTheOAuthAppSClientSecretInClearTextThisWillFirstBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored) implementsRunnerConfigurationScmIntegrationNewParamsBodyUnion() {
}

type RunnerConfigurationScmIntegrationGetParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r RunnerConfigurationScmIntegrationGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationScmIntegrationUpdateParams struct {
	Body RunnerConfigurationScmIntegrationUpdateParamsBodyUnion `json:"body,required"`
}

func (r RunnerConfigurationScmIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationScmIntegrationUpdateParamsBody struct {
	// oauth_client_id can be set to update the OAuth app's client ID. If an empty
	// string is set, the OAuth configuration will be removed (regardless of whether a
	// client secret is set), and any existing Host Authentication Tokens for the SCM
	// integration's runner and host that were created using the OAuth app will be
	// deleted. This might lead to users being unable to access their repositories
	// until they re-authenticate.
	OAuthClientID param.Field[string] `json:"oauthClientId"`
	// oauth_plaintext_client_secret can be set to update the OAuth app's client
	// secret. The cleartext secret will be encrypted with the runner's public key
	// before being stored.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret"`
	// pat can be set to enable or disable Personal Access Tokens support. When
	// disabling PATs, any existing Host Authentication Tokens for the SCM
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
	// oauth_client_id can be set to update the OAuth app's client ID. If an empty
	// string is set, the OAuth configuration will be removed (regardless of whether a
	// client secret is set), and any existing Host Authentication Tokens for the SCM
	// integration's runner and host that were created using the OAuth app will be
	// deleted. This might lead to users being unable to access their repositories
	// until they re-authenticate.
	OAuthClientID param.Field[string] `json:"oauthClientId,required"`
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthClientIDCanBeSetToUpdateTheOAuthAppSClientIDIfAnEmptyStringIsSetTheOAuthConfigurationWillBeRemovedRegardlessOfWhetherAClientSecretIsSetAndAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingTheOAuthAppWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthClientIDCanBeSetToUpdateTheOAuthAppSClientIDIfAnEmptyStringIsSetTheOAuthConfigurationWillBeRemovedRegardlessOfWhetherAClientSecretIsSetAndAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingTheOAuthAppWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate) implementsRunnerConfigurationScmIntegrationUpdateParamsBodyUnion() {
}

type RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthPlaintextClientSecretCanBeSetToUpdateTheOAuthAppSClientSecretTheCleartextSecretWillBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored struct {
	// oauth_plaintext_client_secret can be set to update the OAuth app's client
	// secret. The cleartext secret will be encrypted with the runner's public key
	// before being stored.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret,required"`
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthPlaintextClientSecretCanBeSetToUpdateTheOAuthAppSClientSecretTheCleartextSecretWillBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthPlaintextClientSecretCanBeSetToUpdateTheOAuthAppSClientSecretTheCleartextSecretWillBeEncryptedWithTheRunnerSPublicKeyBeforeBeingStored) implementsRunnerConfigurationScmIntegrationUpdateParamsBodyUnion() {
}

type RunnerConfigurationScmIntegrationUpdateParamsBodyPatCanBeSetToEnableOrDisablePersonalAccessTokensSupportWhenDisablingPaTsAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingAPatWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate struct {
	// pat can be set to enable or disable Personal Access Tokens support. When
	// disabling PATs, any existing Host Authentication Tokens for the SCM
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

type RunnerConfigurationScmIntegrationListParams struct {
	Token    param.Field[string]                                            `query:"token"`
	PageSize param.Field[int64]                                             `query:"pageSize"`
	Filter   param.Field[RunnerConfigurationScmIntegrationListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing scm integrations
	Pagination param.Field[RunnerConfigurationScmIntegrationListParamsPagination] `json:"pagination"`
}

func (r RunnerConfigurationScmIntegrationListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [RunnerConfigurationScmIntegrationListParams]'s query
// parameters as `url.Values`.
func (r RunnerConfigurationScmIntegrationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RunnerConfigurationScmIntegrationListParamsFilter struct {
	// runner_ids filters the response to only SCM integrations of these Runner IDs
	RunnerIDs param.Field[[]string] `json:"runnerIds" format:"uuid"`
}

func (r RunnerConfigurationScmIntegrationListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing scm integrations
type RunnerConfigurationScmIntegrationListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerConfigurationScmIntegrationListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationScmIntegrationDeleteParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r RunnerConfigurationScmIntegrationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
