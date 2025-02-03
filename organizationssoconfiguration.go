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
func (r *OrganizationSSOConfigurationService) New(ctx context.Context, params OrganizationSSOConfigurationNewParams, opts ...option.RequestOption) (res *OrganizationSSOConfigurationNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/CreateSSOConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetSSOConfiguration returns an SSO configuration.
func (r *OrganizationSSOConfigurationService) Get(ctx context.Context, params OrganizationSSOConfigurationGetParams, opts ...option.RequestOption) (res *OrganizationSSOConfigurationGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/GetSSOConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateSSOConfiguration updates the SSO configuration for the organization.
func (r *OrganizationSSOConfigurationService) Update(ctx context.Context, params OrganizationSSOConfigurationUpdateParams, opts ...option.RequestOption) (res *OrganizationSSOConfigurationUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/UpdateSSOConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListSSOConfigurations lists all SSO configurations matching provided filters.
func (r *OrganizationSSOConfigurationService) List(ctx context.Context, params OrganizationSSOConfigurationListParams, opts ...option.RequestOption) (res *OrganizationSSOConfigurationListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/ListSSOConfigurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// DeleteSSOConfiguration deletes an SSO configuration.
func (r *OrganizationSSOConfigurationService) Delete(ctx context.Context, params OrganizationSSOConfigurationDeleteParams, opts ...option.RequestOption) (res *OrganizationSSOConfigurationDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/DeleteSSOConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
	Pagination OrganizationSSOConfigurationListResponsePagination `json:"pagination"`
	// sso_configurations are the SSO configurations for the organization
	SSOConfigurations []OrganizationSSOConfigurationListResponseSSOConfiguration `json:"ssoConfigurations"`
	JSON              organizationSSOConfigurationListResponseJSON               `json:"-"`
}

// organizationSSOConfigurationListResponseJSON contains the JSON metadata for the
// struct [OrganizationSSOConfigurationListResponse]
type organizationSSOConfigurationListResponseJSON struct {
	Pagination        apijson.Field
	SSOConfigurations apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrganizationSSOConfigurationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationSSOConfigurationListResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationSSOConfigurationListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                                 `json:"nextToken"`
	JSON      organizationSSOConfigurationListResponsePaginationJSON `json:"-"`
}

// organizationSSOConfigurationListResponsePaginationJSON contains the JSON
// metadata for the struct [OrganizationSSOConfigurationListResponsePagination]
type organizationSSOConfigurationListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationSSOConfigurationListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationSSOConfigurationListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type OrganizationSSOConfigurationListResponseSSOConfiguration struct {
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
	ProviderType OrganizationSSOConfigurationListResponseSSOConfigurationsProviderType `json:"providerType"`
	// state is the state of the SSO configuration
	State OrganizationSSOConfigurationListResponseSSOConfigurationsState `json:"state"`
	JSON  organizationSSOConfigurationListResponseSSOConfigurationJSON   `json:"-"`
}

// organizationSSOConfigurationListResponseSSOConfigurationJSON contains the JSON
// metadata for the struct
// [OrganizationSSOConfigurationListResponseSSOConfiguration]
type organizationSSOConfigurationListResponseSSOConfigurationJSON struct {
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

func (r *OrganizationSSOConfigurationListResponseSSOConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationSSOConfigurationListResponseSSOConfigurationJSON) RawJSON() string {
	return r.raw
}

// provider_type defines the type of the SSO configuration
type OrganizationSSOConfigurationListResponseSSOConfigurationsProviderType string

const (
	OrganizationSSOConfigurationListResponseSSOConfigurationsProviderTypeProviderTypeUnspecified OrganizationSSOConfigurationListResponseSSOConfigurationsProviderType = "PROVIDER_TYPE_UNSPECIFIED"
	OrganizationSSOConfigurationListResponseSSOConfigurationsProviderTypeProviderTypeBuiltin     OrganizationSSOConfigurationListResponseSSOConfigurationsProviderType = "PROVIDER_TYPE_BUILTIN"
	OrganizationSSOConfigurationListResponseSSOConfigurationsProviderTypeProviderTypeCustom      OrganizationSSOConfigurationListResponseSSOConfigurationsProviderType = "PROVIDER_TYPE_CUSTOM"
)

func (r OrganizationSSOConfigurationListResponseSSOConfigurationsProviderType) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationListResponseSSOConfigurationsProviderTypeProviderTypeUnspecified, OrganizationSSOConfigurationListResponseSSOConfigurationsProviderTypeProviderTypeBuiltin, OrganizationSSOConfigurationListResponseSSOConfigurationsProviderTypeProviderTypeCustom:
		return true
	}
	return false
}

// state is the state of the SSO configuration
type OrganizationSSOConfigurationListResponseSSOConfigurationsState string

const (
	OrganizationSSOConfigurationListResponseSSOConfigurationsStateSSOConfigurationStateUnspecified OrganizationSSOConfigurationListResponseSSOConfigurationsState = "SSO_CONFIGURATION_STATE_UNSPECIFIED"
	OrganizationSSOConfigurationListResponseSSOConfigurationsStateSSOConfigurationStateInactive    OrganizationSSOConfigurationListResponseSSOConfigurationsState = "SSO_CONFIGURATION_STATE_INACTIVE"
	OrganizationSSOConfigurationListResponseSSOConfigurationsStateSSOConfigurationStateActive      OrganizationSSOConfigurationListResponseSSOConfigurationsState = "SSO_CONFIGURATION_STATE_ACTIVE"
)

func (r OrganizationSSOConfigurationListResponseSSOConfigurationsState) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationListResponseSSOConfigurationsStateSSOConfigurationStateUnspecified, OrganizationSSOConfigurationListResponseSSOConfigurationsStateSSOConfigurationStateInactive, OrganizationSSOConfigurationListResponseSSOConfigurationsStateSSOConfigurationStateActive:
		return true
	}
	return false
}

type OrganizationSSOConfigurationDeleteResponse = interface{}

type OrganizationSSOConfigurationNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationSSOConfigurationNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// client_id is the client ID of the OIDC application set on the IdP
	ClientID param.Field[string] `json:"clientId"`
	// client_secret is the client secret of the OIDC application set on the IdP
	ClientSecret param.Field[string] `json:"clientSecret"`
	// email_domain is the domain that is allowed to sign in to the organization
	EmailDomain param.Field[string] `json:"emailDomain"`
	// issuer_url is the URL of the IdP issuer
	IssuerURL      param.Field[string] `json:"issuerUrl" format:"uri"`
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationSSOConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationSSOConfigurationNewParamsConnectProtocolVersion float64

const (
	OrganizationSSOConfigurationNewParamsConnectProtocolVersion1 OrganizationSSOConfigurationNewParamsConnectProtocolVersion = 1
)

func (r OrganizationSSOConfigurationNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationSSOConfigurationGetParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationSSOConfigurationGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// sso_configuration_id is the ID of the SSO configuration to get
	SSOConfigurationID param.Field[string] `json:"ssoConfigurationId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationSSOConfigurationGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationSSOConfigurationGetParamsConnectProtocolVersion float64

const (
	OrganizationSSOConfigurationGetParamsConnectProtocolVersion1 OrganizationSSOConfigurationGetParamsConnectProtocolVersion = 1
)

func (r OrganizationSSOConfigurationGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationSSOConfigurationUpdateParams struct {
	Body OrganizationSSOConfigurationUpdateParamsBodyUnion `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationSSOConfigurationUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
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

// Define the version of the Connect protocol
type OrganizationSSOConfigurationUpdateParamsConnectProtocolVersion float64

const (
	OrganizationSSOConfigurationUpdateParamsConnectProtocolVersion1 OrganizationSSOConfigurationUpdateParamsConnectProtocolVersion = 1
)

func (r OrganizationSSOConfigurationUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationSSOConfigurationListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[OrganizationSSOConfigurationListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationSSOConfigurationListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[OrganizationSSOConfigurationListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[OrganizationSSOConfigurationListParamsConnect] `query:"connect"`
	Message param.Field[string]                                        `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [OrganizationSSOConfigurationListParams]'s query parameters
// as `url.Values`.
func (r OrganizationSSOConfigurationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type OrganizationSSOConfigurationListParamsEncoding string

const (
	OrganizationSSOConfigurationListParamsEncodingProto OrganizationSSOConfigurationListParamsEncoding = "proto"
	OrganizationSSOConfigurationListParamsEncodingJson  OrganizationSSOConfigurationListParamsEncoding = "json"
)

func (r OrganizationSSOConfigurationListParamsEncoding) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationListParamsEncodingProto, OrganizationSSOConfigurationListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type OrganizationSSOConfigurationListParamsConnectProtocolVersion float64

const (
	OrganizationSSOConfigurationListParamsConnectProtocolVersion1 OrganizationSSOConfigurationListParamsConnectProtocolVersion = 1
)

func (r OrganizationSSOConfigurationListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type OrganizationSSOConfigurationListParamsCompression string

const (
	OrganizationSSOConfigurationListParamsCompressionIdentity OrganizationSSOConfigurationListParamsCompression = "identity"
	OrganizationSSOConfigurationListParamsCompressionGzip     OrganizationSSOConfigurationListParamsCompression = "gzip"
	OrganizationSSOConfigurationListParamsCompressionBr       OrganizationSSOConfigurationListParamsCompression = "br"
)

func (r OrganizationSSOConfigurationListParamsCompression) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationListParamsCompressionIdentity, OrganizationSSOConfigurationListParamsCompressionGzip, OrganizationSSOConfigurationListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type OrganizationSSOConfigurationListParamsConnect string

const (
	OrganizationSSOConfigurationListParamsConnectV1 OrganizationSSOConfigurationListParamsConnect = "v1"
)

func (r OrganizationSSOConfigurationListParamsConnect) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationListParamsConnectV1:
		return true
	}
	return false
}

type OrganizationSSOConfigurationDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationSSOConfigurationDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	SSOConfigurationID     param.Field[string]                                                         `json:"ssoConfigurationId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationSSOConfigurationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationSSOConfigurationDeleteParamsConnectProtocolVersion float64

const (
	OrganizationSSOConfigurationDeleteParamsConnectProtocolVersion1 OrganizationSSOConfigurationDeleteParamsConnectProtocolVersion = 1
)

func (r OrganizationSSOConfigurationDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationSSOConfigurationDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}
