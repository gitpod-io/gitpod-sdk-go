// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
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

// Creates a new SCM integration for a runner.
//
// Use this method to:
//
// - Configure source control access
// - Set up repository integrations
// - Enable code synchronization
//
// ### Examples
//
// - Create GitHub integration:
//
//	Sets up GitHub SCM integration.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	scmId: "github"
//	host: "github.com"
//	oauthClientId: "client_id"
//	oauthPlaintextClientSecret: "client_secret"
//	```
func (r *RunnerConfigurationScmIntegrationService) New(ctx context.Context, body RunnerConfigurationScmIntegrationNewParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/CreateSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets details about a specific SCM integration.
//
// Use this method to:
//
// - View integration settings
// - Check integration status
// - Verify configuration
//
// ### Examples
//
// - Get integration details:
//
//	Retrieves information about a specific integration.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *RunnerConfigurationScmIntegrationService) Get(ctx context.Context, body RunnerConfigurationScmIntegrationGetParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing SCM integration.
//
// Use this method to:
//
// - Modify integration settings
// - Update credentials
// - Change configuration
//
// ### Examples
//
// - Update integration:
//
//	Updates OAuth credentials.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	oauthClientId: "new_client_id"
//	oauthPlaintextClientSecret: "new_client_secret"
//	```
func (r *RunnerConfigurationScmIntegrationService) Update(ctx context.Context, body RunnerConfigurationScmIntegrationUpdateParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/UpdateSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists SCM integrations for a runner.
//
// Use this method to:
//
// - View all integrations
// - Monitor integration status
// - Check available SCMs
//
// ### Examples
//
// - List integrations:
//
//	Shows all SCM integrations.
//
//	```yaml
//	filter:
//	  runnerIds: ["d2c94c27-3b76-4a42-b88c-95a85e392c68"]
//	pagination:
//	  pageSize: 20
//	```
func (r *RunnerConfigurationScmIntegrationService) List(ctx context.Context, params RunnerConfigurationScmIntegrationListParams, opts ...option.RequestOption) (res *pagination.IntegrationsPage[ScmIntegration], err error) {
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

// Lists SCM integrations for a runner.
//
// Use this method to:
//
// - View all integrations
// - Monitor integration status
// - Check available SCMs
//
// ### Examples
//
// - List integrations:
//
//	Shows all SCM integrations.
//
//	```yaml
//	filter:
//	  runnerIds: ["d2c94c27-3b76-4a42-b88c-95a85e392c68"]
//	pagination:
//	  pageSize: 20
//	```
func (r *RunnerConfigurationScmIntegrationService) ListAutoPaging(ctx context.Context, params RunnerConfigurationScmIntegrationListParams, opts ...option.RequestOption) *pagination.IntegrationsPageAutoPager[ScmIntegration] {
	return pagination.NewIntegrationsPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes an SCM integration.
//
// Use this method to:
//
// - Remove unused integrations
// - Clean up configurations
// - Revoke SCM access
//
// ### Examples
//
// - Delete integration:
//
//	Removes an SCM integration.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *RunnerConfigurationScmIntegrationService) Delete(ctx context.Context, body RunnerConfigurationScmIntegrationDeleteParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/DeleteSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ScmIntegration struct {
	// id is the unique identifier of the SCM integration
	ID       string                    `json:"id"`
	Host     string                    `json:"host"`
	OAuth    ScmIntegrationOAuthConfig `json:"oauth,nullable"`
	Pat      bool                      `json:"pat"`
	RunnerID string                    `json:"runnerId"`
	// scm_id references the scm_id in the runner's configuration schema that this
	// integration is for
	ScmID string             `json:"scmId"`
	JSON  scmIntegrationJSON `json:"-"`
}

// scmIntegrationJSON contains the JSON metadata for the struct [ScmIntegration]
type scmIntegrationJSON struct {
	ID          apijson.Field
	Host        apijson.Field
	OAuth       apijson.Field
	Pat         apijson.Field
	RunnerID    apijson.Field
	ScmID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScmIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scmIntegrationJSON) RawJSON() string {
	return r.raw
}

type ScmIntegrationOAuthConfig struct {
	// client_id is the OAuth app's client ID in clear text.
	ClientID string `json:"clientId"`
	// encrypted_client_secret is the OAuth app's secret encrypted with the runner's
	// public key.
	EncryptedClientSecret string `json:"encryptedClientSecret" format:"byte"`
	// issuer_url is used to override the authentication provider URL, if it doesn't
	// match the SCM host.
	//
	// +optional if not set, this account is owned by the installation.
	IssuerURL string                        `json:"issuerUrl"`
	JSON      scmIntegrationOAuthConfigJSON `json:"-"`
}

// scmIntegrationOAuthConfigJSON contains the JSON metadata for the struct
// [ScmIntegrationOAuthConfig]
type scmIntegrationOAuthConfigJSON struct {
	ClientID              apijson.Field
	EncryptedClientSecret apijson.Field
	IssuerURL             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ScmIntegrationOAuthConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scmIntegrationOAuthConfigJSON) RawJSON() string {
	return r.raw
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
	Integration ScmIntegration                                   `json:"integration"`
	JSON        runnerConfigurationScmIntegrationGetResponseJSON `json:"-"`
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

type RunnerConfigurationScmIntegrationUpdateResponse = interface{}

type RunnerConfigurationScmIntegrationDeleteResponse = interface{}

type RunnerConfigurationScmIntegrationNewParams struct {
	Host param.Field[string] `json:"host"`
	// issuer_url can be set to override the authentication provider URL, if it doesn't
	// match the SCM host.
	IssuerURL param.Field[string] `json:"issuerUrl"`
	// oauth_client_id is the OAuth app's client ID, if OAuth is configured. If
	// configured, oauth_plaintext_client_secret must also be set.
	OAuthClientID param.Field[string] `json:"oauthClientId"`
	// oauth_plaintext_client_secret is the OAuth app's client secret in clear text.
	// This will first be encrypted with the runner's public key before being stored.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret"`
	Pat                        param.Field[bool]   `json:"pat"`
	RunnerID                   param.Field[string] `json:"runnerId" format:"uuid"`
	// scm_id references the scm_id in the runner's configuration schema that this
	// integration is for
	ScmID param.Field[string] `json:"scmId"`
}

func (r RunnerConfigurationScmIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationScmIntegrationGetParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r RunnerConfigurationScmIntegrationGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationScmIntegrationUpdateParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// issuer_url can be set to override the authentication provider URL, if it doesn't
	// match the SCM host.
	IssuerURL param.Field[string] `json:"issuerUrl"`
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

func (r RunnerConfigurationScmIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
