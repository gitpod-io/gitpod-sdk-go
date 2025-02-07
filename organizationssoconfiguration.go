// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gitpod-io/flex-sdk-go/internal/apijson"
	"github.com/gitpod-io/flex-sdk-go/internal/apiquery"
	"github.com/gitpod-io/flex-sdk-go/internal/param"
	"github.com/gitpod-io/flex-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/flex-sdk-go/option"
	"github.com/gitpod-io/flex-sdk-go/packages/pagination"
)

// OrganizationSSOConfigurationService contains methods and other services that
// help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationSSOConfigurationService] method instead.
type OrganizationSSOConfigurationService struct {
	Options []option.RequestOption
}

// NewOrganizationSSOConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrganizationSSOConfigurationService(opts ...option.RequestOption) (r *OrganizationSSOConfigurationService) {
	r = &OrganizationSSOConfigurationService{}
	r.Options = opts
	return
}

// CreateSSOConfiguration creates a new SSO configuration for the organization.
func (r *OrganizationSSOConfigurationService) New(ctx context.Context, body OrganizationSSOConfigurationNewParams, opts ...option.RequestOption) (res *OrganizationSSOConfigurationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/CreateSSOConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetSSOConfiguration returns an SSO configuration.
func (r *OrganizationSSOConfigurationService) Get(ctx context.Context, body OrganizationSSOConfigurationGetParams, opts ...option.RequestOption) (res *OrganizationSSOConfigurationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/GetSSOConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateSSOConfiguration updates the SSO configuration for the organization.
func (r *OrganizationSSOConfigurationService) Update(ctx context.Context, body OrganizationSSOConfigurationUpdateParams, opts ...option.RequestOption) (res *OrganizationSSOConfigurationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/UpdateSSOConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListSSOConfigurations lists all SSO configurations matching provided filters.
func (r *OrganizationSSOConfigurationService) List(ctx context.Context, params OrganizationSSOConfigurationListParams, opts ...option.RequestOption) (res *pagination.SSOConfigurationsPage[SSOConfiguration], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.OrganizationService/ListSSOConfigurations"
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

// ListSSOConfigurations lists all SSO configurations matching provided filters.
func (r *OrganizationSSOConfigurationService) ListAutoPaging(ctx context.Context, params OrganizationSSOConfigurationListParams, opts ...option.RequestOption) *pagination.SSOConfigurationsPageAutoPager[SSOConfiguration] {
	return pagination.NewSSOConfigurationsPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteSSOConfiguration deletes an SSO configuration.
func (r *OrganizationSSOConfigurationService) Delete(ctx context.Context, body OrganizationSSOConfigurationDeleteParams, opts ...option.RequestOption) (res *OrganizationSSOConfigurationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/DeleteSSOConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProviderType string

const (
	ProviderTypeProviderTypeUnspecified ProviderType = "PROVIDER_TYPE_UNSPECIFIED"
	ProviderTypeProviderTypeBuiltin     ProviderType = "PROVIDER_TYPE_BUILTIN"
	ProviderTypeProviderTypeCustom      ProviderType = "PROVIDER_TYPE_CUSTOM"
)

func (r ProviderType) IsKnown() bool {
	switch r {
	case ProviderTypeProviderTypeUnspecified, ProviderTypeProviderTypeBuiltin, ProviderTypeProviderTypeCustom:
		return true
	}
	return false
}

type SSOConfiguration struct {
	// id is the unique identifier of the SSO configuration
	ID string `json:"id" format:"uuid"`
	// claims are key/value pairs that defines a mapping of claims issued by the IdP.
	Claims map[string]string `json:"claims"`
	// client_id is the client ID of the OIDC application set on the IdP
	ClientID    string `json:"clientId"`
	EmailDomain string `json:"emailDomain"`
	// issuer_url is the URL of the IdP issuer
	IssuerURL      string `json:"issuerUrl"`
	OrganizationID string `json:"organizationId" format:"uuid"`
	// provider_type defines the type of the SSO configuration
	ProviderType ProviderType `json:"providerType"`
	// state is the state of the SSO configuration
	State SSOConfigurationState `json:"state"`
	JSON  ssoConfigurationJSON  `json:"-"`
}

// ssoConfigurationJSON contains the JSON metadata for the struct
// [SSOConfiguration]
type ssoConfigurationJSON struct {
	ID             apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	EmailDomain    apijson.Field
	IssuerURL      apijson.Field
	OrganizationID apijson.Field
	ProviderType   apijson.Field
	State          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SSOConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoConfigurationJSON) RawJSON() string {
	return r.raw
}

type SSOConfigurationState string

const (
	SSOConfigurationStateSSOConfigurationStateUnspecified SSOConfigurationState = "SSO_CONFIGURATION_STATE_UNSPECIFIED"
	SSOConfigurationStateSSOConfigurationStateInactive    SSOConfigurationState = "SSO_CONFIGURATION_STATE_INACTIVE"
	SSOConfigurationStateSSOConfigurationStateActive      SSOConfigurationState = "SSO_CONFIGURATION_STATE_ACTIVE"
)

func (r SSOConfigurationState) IsKnown() bool {
	switch r {
	case SSOConfigurationStateSSOConfigurationStateUnspecified, SSOConfigurationStateSSOConfigurationStateInactive, SSOConfigurationStateSSOConfigurationStateActive:
		return true
	}
	return false
}

type OrganizationSSOConfigurationNewResponse struct {
	// sso_configuration is the created SSO configuration
	SSOConfiguration SSOConfiguration                            `json:"ssoConfiguration"`
	JSON             organizationSSOConfigurationNewResponseJSON `json:"-"`
}

// organizationSSOConfigurationNewResponseJSON contains the JSON metadata for the
// struct [OrganizationSSOConfigurationNewResponse]
type organizationSSOConfigurationNewResponseJSON struct {
	SSOConfiguration apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationSSOConfigurationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationSSOConfigurationNewResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationSSOConfigurationGetResponse struct {
	// sso_configuration is the SSO configuration identified by the ID
	SSOConfiguration SSOConfiguration                            `json:"ssoConfiguration"`
	JSON             organizationSSOConfigurationGetResponseJSON `json:"-"`
}

// organizationSSOConfigurationGetResponseJSON contains the JSON metadata for the
// struct [OrganizationSSOConfigurationGetResponse]
type organizationSSOConfigurationGetResponseJSON struct {
	SSOConfiguration apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationSSOConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationSSOConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationSSOConfigurationUpdateResponse = interface{}

type OrganizationSSOConfigurationDeleteResponse = interface{}

type OrganizationSSOConfigurationNewParams struct {
	// client_id is the client ID of the OIDC application set on the IdP
	ClientID param.Field[string] `json:"clientId"`
	// client_secret is the client secret of the OIDC application set on the IdP
	ClientSecret param.Field[string] `json:"clientSecret"`
	// email_domain is the domain that is allowed to sign in to the organization
	EmailDomain param.Field[string] `json:"emailDomain"`
	// issuer_url is the URL of the IdP issuer
	IssuerURL      param.Field[string] `json:"issuerUrl" format:"uri"`
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
}

func (r OrganizationSSOConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationSSOConfigurationGetParams struct {
	// sso_configuration_id is the ID of the SSO configuration to get
	SSOConfigurationID param.Field[string] `json:"ssoConfigurationId" format:"uuid"`
}

func (r OrganizationSSOConfigurationGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationSSOConfigurationUpdateParams struct {
	// claims are key/value pairs that defines a mapping of claims issued by the IdP.
	Claims param.Field[map[string]string] `json:"claims"`
	// client_id is the client ID of the SSO provider
	ClientID param.Field[string] `json:"clientId"`
	// client_secret is the client secret of the SSO provider
	ClientSecret param.Field[string] `json:"clientSecret"`
	EmailDomain  param.Field[string] `json:"emailDomain"`
	// issuer_url is the URL of the IdP issuer
	IssuerURL param.Field[string] `json:"issuerUrl" format:"uri"`
	// sso_configuration_id is the ID of the SSO configuration to update
	SSOConfigurationID param.Field[string] `json:"ssoConfigurationId" format:"uuid"`
	// state is the state of the SSO configuration
	State param.Field[SSOConfigurationState] `json:"state"`
}

func (r OrganizationSSOConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationSSOConfigurationListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// organization_id is the ID of the organization to list SSO configurations for.
	OrganizationID param.Field[string]                                           `json:"organizationId" format:"uuid"`
	Pagination     param.Field[OrganizationSSOConfigurationListParamsPagination] `json:"pagination"`
}

func (r OrganizationSSOConfigurationListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [OrganizationSSOConfigurationListParams]'s query parameters
// as `url.Values`.
func (r OrganizationSSOConfigurationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationSSOConfigurationListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r OrganizationSSOConfigurationListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationSSOConfigurationDeleteParams struct {
	SSOConfigurationID param.Field[string] `json:"ssoConfigurationId" format:"uuid"`
}

func (r OrganizationSSOConfigurationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
