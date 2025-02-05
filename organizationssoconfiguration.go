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
func (r *OrganizationSSOConfigurationService) List(ctx context.Context, params OrganizationSSOConfigurationListParams, opts ...option.RequestOption) (res *pagination.SSOConfigurationsPage[OrganizationSSOConfigurationListResponse], err error) {
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
func (r *OrganizationSSOConfigurationService) ListAutoPaging(ctx context.Context, params OrganizationSSOConfigurationListParams, opts ...option.RequestOption) *pagination.SSOConfigurationsPageAutoPager[OrganizationSSOConfigurationListResponse] {
	return pagination.NewSSOConfigurationsPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteSSOConfiguration deletes an SSO configuration.
func (r *OrganizationSSOConfigurationService) Delete(ctx context.Context, body OrganizationSSOConfigurationDeleteParams, opts ...option.RequestOption) (res *OrganizationSSOConfigurationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/DeleteSSOConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type OrganizationSSOConfigurationNewResponse struct {
	// sso_configuration is the created SSO configuration
	SSOConfiguration OrganizationSSOConfigurationNewResponseSSOConfiguration `json:"ssoConfiguration"`
	JSON             organizationSSOConfigurationNewResponseJSON             `json:"-"`
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

// sso_configuration is the created SSO configuration
type OrganizationSSOConfigurationNewResponseSSOConfiguration struct {
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
	ProviderType OrganizationSSOConfigurationNewResponseSSOConfigurationProviderType `json:"providerType"`
	// state is the state of the SSO configuration
	State OrganizationSSOConfigurationNewResponseSSOConfigurationState `json:"state"`
	JSON  organizationSSOConfigurationNewResponseSSOConfigurationJSON  `json:"-"`
}

// organizationSSOConfigurationNewResponseSSOConfigurationJSON contains the JSON
// metadata for the struct
// [OrganizationSSOConfigurationNewResponseSSOConfiguration]
type organizationSSOConfigurationNewResponseSSOConfigurationJSON struct {
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

func (r *OrganizationSSOConfigurationNewResponseSSOConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationSSOConfigurationNewResponseSSOConfigurationJSON) RawJSON() string {
	return r.raw
}

// provider_type defines the type of the SSO configuration
type OrganizationSSOConfigurationNewResponseSSOConfigurationProviderType string

const (
	OrganizationSSOConfigurationNewResponseSSOConfigurationProviderTypeProviderTypeUnspecified OrganizationSSOConfigurationNewResponseSSOConfigurationProviderType = "PROVIDER_TYPE_UNSPECIFIED"
	OrganizationSSOConfigurationNewResponseSSOConfigurationProviderTypeProviderTypeBuiltin     OrganizationSSOConfigurationNewResponseSSOConfigurationProviderType = "PROVIDER_TYPE_BUILTIN"
	OrganizationSSOConfigurationNewResponseSSOConfigurationProviderTypeProviderTypeCustom      OrganizationSSOConfigurationNewResponseSSOConfigurationProviderType = "PROVIDER_TYPE_CUSTOM"
)

func (r OrganizationSSOConfigurationNewResponseSSOConfigurationProviderType) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationNewResponseSSOConfigurationProviderTypeProviderTypeUnspecified, OrganizationSSOConfigurationNewResponseSSOConfigurationProviderTypeProviderTypeBuiltin, OrganizationSSOConfigurationNewResponseSSOConfigurationProviderTypeProviderTypeCustom:
		return true
	}
	return false
}

// state is the state of the SSO configuration
type OrganizationSSOConfigurationNewResponseSSOConfigurationState string

const (
	OrganizationSSOConfigurationNewResponseSSOConfigurationStateSSOConfigurationStateUnspecified OrganizationSSOConfigurationNewResponseSSOConfigurationState = "SSO_CONFIGURATION_STATE_UNSPECIFIED"
	OrganizationSSOConfigurationNewResponseSSOConfigurationStateSSOConfigurationStateInactive    OrganizationSSOConfigurationNewResponseSSOConfigurationState = "SSO_CONFIGURATION_STATE_INACTIVE"
	OrganizationSSOConfigurationNewResponseSSOConfigurationStateSSOConfigurationStateActive      OrganizationSSOConfigurationNewResponseSSOConfigurationState = "SSO_CONFIGURATION_STATE_ACTIVE"
)

func (r OrganizationSSOConfigurationNewResponseSSOConfigurationState) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationNewResponseSSOConfigurationStateSSOConfigurationStateUnspecified, OrganizationSSOConfigurationNewResponseSSOConfigurationStateSSOConfigurationStateInactive, OrganizationSSOConfigurationNewResponseSSOConfigurationStateSSOConfigurationStateActive:
		return true
	}
	return false
}

type OrganizationSSOConfigurationGetResponse struct {
	// sso_configuration is the SSO configuration identified by the ID
	SSOConfiguration OrganizationSSOConfigurationGetResponseSSOConfiguration `json:"ssoConfiguration"`
	JSON             organizationSSOConfigurationGetResponseJSON             `json:"-"`
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

// sso_configuration is the SSO configuration identified by the ID
type OrganizationSSOConfigurationGetResponseSSOConfiguration struct {
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
	ProviderType OrganizationSSOConfigurationGetResponseSSOConfigurationProviderType `json:"providerType"`
	// state is the state of the SSO configuration
	State OrganizationSSOConfigurationGetResponseSSOConfigurationState `json:"state"`
	JSON  organizationSSOConfigurationGetResponseSSOConfigurationJSON  `json:"-"`
}

// organizationSSOConfigurationGetResponseSSOConfigurationJSON contains the JSON
// metadata for the struct
// [OrganizationSSOConfigurationGetResponseSSOConfiguration]
type organizationSSOConfigurationGetResponseSSOConfigurationJSON struct {
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

func (r *OrganizationSSOConfigurationGetResponseSSOConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationSSOConfigurationGetResponseSSOConfigurationJSON) RawJSON() string {
	return r.raw
}

// provider_type defines the type of the SSO configuration
type OrganizationSSOConfigurationGetResponseSSOConfigurationProviderType string

const (
	OrganizationSSOConfigurationGetResponseSSOConfigurationProviderTypeProviderTypeUnspecified OrganizationSSOConfigurationGetResponseSSOConfigurationProviderType = "PROVIDER_TYPE_UNSPECIFIED"
	OrganizationSSOConfigurationGetResponseSSOConfigurationProviderTypeProviderTypeBuiltin     OrganizationSSOConfigurationGetResponseSSOConfigurationProviderType = "PROVIDER_TYPE_BUILTIN"
	OrganizationSSOConfigurationGetResponseSSOConfigurationProviderTypeProviderTypeCustom      OrganizationSSOConfigurationGetResponseSSOConfigurationProviderType = "PROVIDER_TYPE_CUSTOM"
)

func (r OrganizationSSOConfigurationGetResponseSSOConfigurationProviderType) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationGetResponseSSOConfigurationProviderTypeProviderTypeUnspecified, OrganizationSSOConfigurationGetResponseSSOConfigurationProviderTypeProviderTypeBuiltin, OrganizationSSOConfigurationGetResponseSSOConfigurationProviderTypeProviderTypeCustom:
		return true
	}
	return false
}

// state is the state of the SSO configuration
type OrganizationSSOConfigurationGetResponseSSOConfigurationState string

const (
	OrganizationSSOConfigurationGetResponseSSOConfigurationStateSSOConfigurationStateUnspecified OrganizationSSOConfigurationGetResponseSSOConfigurationState = "SSO_CONFIGURATION_STATE_UNSPECIFIED"
	OrganizationSSOConfigurationGetResponseSSOConfigurationStateSSOConfigurationStateInactive    OrganizationSSOConfigurationGetResponseSSOConfigurationState = "SSO_CONFIGURATION_STATE_INACTIVE"
	OrganizationSSOConfigurationGetResponseSSOConfigurationStateSSOConfigurationStateActive      OrganizationSSOConfigurationGetResponseSSOConfigurationState = "SSO_CONFIGURATION_STATE_ACTIVE"
)

func (r OrganizationSSOConfigurationGetResponseSSOConfigurationState) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationGetResponseSSOConfigurationStateSSOConfigurationStateUnspecified, OrganizationSSOConfigurationGetResponseSSOConfigurationStateSSOConfigurationStateInactive, OrganizationSSOConfigurationGetResponseSSOConfigurationStateSSOConfigurationStateActive:
		return true
	}
	return false
}

type OrganizationSSOConfigurationUpdateResponse = interface{}

type OrganizationSSOConfigurationListResponse struct {
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
	ProviderType OrganizationSSOConfigurationListResponseProviderType `json:"providerType"`
	// state is the state of the SSO configuration
	State OrganizationSSOConfigurationListResponseState `json:"state"`
	JSON  organizationSSOConfigurationListResponseJSON  `json:"-"`
}

// organizationSSOConfigurationListResponseJSON contains the JSON metadata for the
// struct [OrganizationSSOConfigurationListResponse]
type organizationSSOConfigurationListResponseJSON struct {
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

func (r *OrganizationSSOConfigurationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationSSOConfigurationListResponseJSON) RawJSON() string {
	return r.raw
}

// provider_type defines the type of the SSO configuration
type OrganizationSSOConfigurationListResponseProviderType string

const (
	OrganizationSSOConfigurationListResponseProviderTypeProviderTypeUnspecified OrganizationSSOConfigurationListResponseProviderType = "PROVIDER_TYPE_UNSPECIFIED"
	OrganizationSSOConfigurationListResponseProviderTypeProviderTypeBuiltin     OrganizationSSOConfigurationListResponseProviderType = "PROVIDER_TYPE_BUILTIN"
	OrganizationSSOConfigurationListResponseProviderTypeProviderTypeCustom      OrganizationSSOConfigurationListResponseProviderType = "PROVIDER_TYPE_CUSTOM"
)

func (r OrganizationSSOConfigurationListResponseProviderType) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationListResponseProviderTypeProviderTypeUnspecified, OrganizationSSOConfigurationListResponseProviderTypeProviderTypeBuiltin, OrganizationSSOConfigurationListResponseProviderTypeProviderTypeCustom:
		return true
	}
	return false
}

// state is the state of the SSO configuration
type OrganizationSSOConfigurationListResponseState string

const (
	OrganizationSSOConfigurationListResponseStateSSOConfigurationStateUnspecified OrganizationSSOConfigurationListResponseState = "SSO_CONFIGURATION_STATE_UNSPECIFIED"
	OrganizationSSOConfigurationListResponseStateSSOConfigurationStateInactive    OrganizationSSOConfigurationListResponseState = "SSO_CONFIGURATION_STATE_INACTIVE"
	OrganizationSSOConfigurationListResponseStateSSOConfigurationStateActive      OrganizationSSOConfigurationListResponseState = "SSO_CONFIGURATION_STATE_ACTIVE"
)

func (r OrganizationSSOConfigurationListResponseState) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationListResponseStateSSOConfigurationStateUnspecified, OrganizationSSOConfigurationListResponseStateSSOConfigurationStateInactive, OrganizationSSOConfigurationListResponseStateSSOConfigurationStateActive:
		return true
	}
	return false
}

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
	Body OrganizationSSOConfigurationUpdateParamsBodyUnion `json:"body,required"`
}

func (r OrganizationSSOConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type OrganizationSSOConfigurationUpdateParamsBody struct {
	// client_id is the client ID of the SSO provider
	ClientID param.Field[string] `json:"clientId"`
	// client_secret is the client secret of the SSO provider
	ClientSecret param.Field[string] `json:"clientSecret"`
	EmailDomain  param.Field[string] `json:"emailDomain"`
	// issuer_url is the URL of the IdP issuer
	IssuerURL param.Field[string] `json:"issuerUrl" format:"uri"`
	// state is the state of the SSO configuration
	State param.Field[OrganizationSSOConfigurationUpdateParamsBodyState] `json:"state"`
}

func (r OrganizationSSOConfigurationUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationSSOConfigurationUpdateParamsBody) implementsOrganizationSSOConfigurationUpdateParamsBodyUnion() {
}

// Satisfied by
// [OrganizationSSOConfigurationUpdateParamsBodyClientIDIsTheClientIDOfTheSSOProvider],
// [OrganizationSSOConfigurationUpdateParamsBodyClientSecretIsTheClientSecretOfTheSSOProvider],
// [OrganizationSSOConfigurationUpdateParamsBodyEmailDomain],
// [OrganizationSSOConfigurationUpdateParamsBodyIssuerURLIsTheURLOfTheIDPIssuer],
// [OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfiguration],
// [OrganizationSSOConfigurationUpdateParamsBody].
type OrganizationSSOConfigurationUpdateParamsBodyUnion interface {
	implementsOrganizationSSOConfigurationUpdateParamsBodyUnion()
}

type OrganizationSSOConfigurationUpdateParamsBodyClientIDIsTheClientIDOfTheSSOProvider struct {
	// client_id is the client ID of the SSO provider
	ClientID param.Field[string] `json:"clientId,required"`
}

func (r OrganizationSSOConfigurationUpdateParamsBodyClientIDIsTheClientIDOfTheSSOProvider) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationSSOConfigurationUpdateParamsBodyClientIDIsTheClientIDOfTheSSOProvider) implementsOrganizationSSOConfigurationUpdateParamsBodyUnion() {
}

type OrganizationSSOConfigurationUpdateParamsBodyClientSecretIsTheClientSecretOfTheSSOProvider struct {
	// client_secret is the client secret of the SSO provider
	ClientSecret param.Field[string] `json:"clientSecret,required"`
}

func (r OrganizationSSOConfigurationUpdateParamsBodyClientSecretIsTheClientSecretOfTheSSOProvider) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationSSOConfigurationUpdateParamsBodyClientSecretIsTheClientSecretOfTheSSOProvider) implementsOrganizationSSOConfigurationUpdateParamsBodyUnion() {
}

type OrganizationSSOConfigurationUpdateParamsBodyEmailDomain struct {
	EmailDomain param.Field[string] `json:"emailDomain,required"`
}

func (r OrganizationSSOConfigurationUpdateParamsBodyEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationSSOConfigurationUpdateParamsBodyEmailDomain) implementsOrganizationSSOConfigurationUpdateParamsBodyUnion() {
}

type OrganizationSSOConfigurationUpdateParamsBodyIssuerURLIsTheURLOfTheIDPIssuer struct {
	// issuer_url is the URL of the IdP issuer
	IssuerURL param.Field[string] `json:"issuerUrl,required" format:"uri"`
}

func (r OrganizationSSOConfigurationUpdateParamsBodyIssuerURLIsTheURLOfTheIDPIssuer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationSSOConfigurationUpdateParamsBodyIssuerURLIsTheURLOfTheIDPIssuer) implementsOrganizationSSOConfigurationUpdateParamsBodyUnion() {
}

type OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfiguration struct {
	// state is the state of the SSO configuration
	State param.Field[OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationState] `json:"state,required"`
}

func (r OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfiguration) implementsOrganizationSSOConfigurationUpdateParamsBodyUnion() {
}

// state is the state of the SSO configuration
type OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationState string

const (
	OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationStateSSOConfigurationStateUnspecified OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationState = "SSO_CONFIGURATION_STATE_UNSPECIFIED"
	OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationStateSSOConfigurationStateInactive    OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationState = "SSO_CONFIGURATION_STATE_INACTIVE"
	OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationStateSSOConfigurationStateActive      OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationState = "SSO_CONFIGURATION_STATE_ACTIVE"
)

func (r OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationState) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationStateSSOConfigurationStateUnspecified, OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationStateSSOConfigurationStateInactive, OrganizationSSOConfigurationUpdateParamsBodyStateIsTheStateOfTheSSOConfigurationStateSSOConfigurationStateActive:
		return true
	}
	return false
}

// state is the state of the SSO configuration
type OrganizationSSOConfigurationUpdateParamsBodyState string

const (
	OrganizationSSOConfigurationUpdateParamsBodyStateSSOConfigurationStateUnspecified OrganizationSSOConfigurationUpdateParamsBodyState = "SSO_CONFIGURATION_STATE_UNSPECIFIED"
	OrganizationSSOConfigurationUpdateParamsBodyStateSSOConfigurationStateInactive    OrganizationSSOConfigurationUpdateParamsBodyState = "SSO_CONFIGURATION_STATE_INACTIVE"
	OrganizationSSOConfigurationUpdateParamsBodyStateSSOConfigurationStateActive      OrganizationSSOConfigurationUpdateParamsBodyState = "SSO_CONFIGURATION_STATE_ACTIVE"
)

func (r OrganizationSSOConfigurationUpdateParamsBodyState) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationUpdateParamsBodyStateSSOConfigurationStateUnspecified, OrganizationSSOConfigurationUpdateParamsBodyStateSSOConfigurationStateInactive, OrganizationSSOConfigurationUpdateParamsBodyStateSSOConfigurationStateActive:
		return true
	}
	return false
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
