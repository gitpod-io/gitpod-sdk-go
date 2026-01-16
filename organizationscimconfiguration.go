// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
)

// OrganizationScimConfigurationService contains methods and other services that
// help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationScimConfigurationService] method instead.
type OrganizationScimConfigurationService struct {
	Options []option.RequestOption
}

// NewOrganizationScimConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrganizationScimConfigurationService(opts ...option.RequestOption) (r *OrganizationScimConfigurationService) {
	r = &OrganizationScimConfigurationService{}
	r.Options = opts
	return
}

// Creates a new SCIM configuration for automated user provisioning.
//
// Use this method to:
//
// - Set up SCIM 2.0 provisioning from an identity provider
// - Generate a bearer token for SCIM API authentication
// - Link SCIM provisioning to an existing SSO configuration
//
// ### Examples
//
// - Create basic SCIM configuration:
//
//	Creates a SCIM configuration linked to an SSO provider.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	ssoConfigurationId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *OrganizationScimConfigurationService) New(ctx context.Context, body OrganizationScimConfigurationNewParams, opts ...option.RequestOption) (res *OrganizationScimConfigurationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/CreateSCIMConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a specific SCIM configuration.
//
// Use this method to:
//
// - View SCIM configuration details
// - Check if SCIM is enabled
// - Verify SSO linkage
//
// ### Examples
//
// - Get SCIM configuration:
//
//	Retrieves details of a specific SCIM configuration.
//
//	```yaml
//	scimConfigurationId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *OrganizationScimConfigurationService) Get(ctx context.Context, body OrganizationScimConfigurationGetParams, opts ...option.RequestOption) (res *OrganizationScimConfigurationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/GetSCIMConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates a SCIM configuration.
//
// Use this method to:
//
// - Enable or disable SCIM provisioning
// - Link or unlink SSO configuration
// - Update configuration name
//
// ### Examples
//
// - Disable SCIM:
//
//	Disables SCIM provisioning.
//
//	```yaml
//	scimConfigurationId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	enabled: false
//	```
//
// - Link to SSO:
//
//	Links SCIM configuration to an SSO provider.
//
//	```yaml
//	scimConfigurationId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	ssoConfigurationId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	```
func (r *OrganizationScimConfigurationService) Update(ctx context.Context, body OrganizationScimConfigurationUpdateParams, opts ...option.RequestOption) (res *OrganizationScimConfigurationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/UpdateSCIMConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists SCIM configurations for an organization.
//
// Use this method to:
//
// - View all SCIM configurations
// - Monitor provisioning status
// - Audit SCIM settings
//
// ### Examples
//
// - List SCIM configurations:
//
//	Shows all SCIM configurations for an organization.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
func (r *OrganizationScimConfigurationService) List(ctx context.Context, params OrganizationScimConfigurationListParams, opts ...option.RequestOption) (res *pagination.ScimConfigurationsPage[ScimConfiguration], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.OrganizationService/ListSCIMConfigurations"
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

// Lists SCIM configurations for an organization.
//
// Use this method to:
//
// - View all SCIM configurations
// - Monitor provisioning status
// - Audit SCIM settings
//
// ### Examples
//
// - List SCIM configurations:
//
//	Shows all SCIM configurations for an organization.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
func (r *OrganizationScimConfigurationService) ListAutoPaging(ctx context.Context, params OrganizationScimConfigurationListParams, opts ...option.RequestOption) *pagination.ScimConfigurationsPageAutoPager[ScimConfiguration] {
	return pagination.NewScimConfigurationsPageAutoPager(r.List(ctx, params, opts...))
}

// Removes a SCIM configuration from an organization.
//
// Use this method to:
//
// - Disable SCIM provisioning completely
// - Remove unused configurations
// - Clean up after migration
//
// ### Examples
//
// - Delete SCIM configuration:
//
//	Removes a specific SCIM configuration.
//
//	```yaml
//	scimConfigurationId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *OrganizationScimConfigurationService) Delete(ctx context.Context, body OrganizationScimConfigurationDeleteParams, opts ...option.RequestOption) (res *OrganizationScimConfigurationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/DeleteSCIMConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Regenerates the bearer token for a SCIM configuration.
//
// Use this method to:
//
// - Rotate SCIM credentials
// - Recover from token compromise
// - Update IdP configuration
//
// ### Examples
//
// - Regenerate token:
//
//	Creates a new bearer token, invalidating the old one.
//
//	```yaml
//	scimConfigurationId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *OrganizationScimConfigurationService) RegenerateToken(ctx context.Context, body OrganizationScimConfigurationRegenerateTokenParams, opts ...option.RequestOption) (res *OrganizationScimConfigurationRegenerateTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/RegenerateSCIMToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// SCIMConfiguration represents a SCIM 2.0 provisioning configuration
type ScimConfiguration struct {
	// id is the unique identifier of the SCIM configuration
	ID string `json:"id,required" format:"uuid"`
	// created_at is when the SCIM configuration was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// organization_id is the ID of the organization this SCIM configuration belongs to
	OrganizationID string `json:"organizationId,required" format:"uuid"`
	// updated_at is when the SCIM configuration was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// enabled indicates if SCIM provisioning is active
	Enabled bool `json:"enabled"`
	// name is a human-readable name for the SCIM configuration
	Name string `json:"name"`
	// sso_configuration_id is the linked SSO configuration (optional)
	SSOConfigurationID string                `json:"ssoConfigurationId" format:"uuid"`
	JSON               scimConfigurationJSON `json:"-"`
}

// scimConfigurationJSON contains the JSON metadata for the struct
// [ScimConfiguration]
type scimConfigurationJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	OrganizationID     apijson.Field
	UpdatedAt          apijson.Field
	Enabled            apijson.Field
	Name               apijson.Field
	SSOConfigurationID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScimConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigurationJSON) RawJSON() string {
	return r.raw
}

type OrganizationScimConfigurationNewResponse struct {
	// token is the bearer token for SCIM API authentication. This is only returned
	// once during creation - store it securely.
	Token string `json:"token,required"`
	// scim_configuration is the created SCIM configuration
	ScimConfiguration ScimConfiguration                            `json:"scimConfiguration,required"`
	JSON              organizationScimConfigurationNewResponseJSON `json:"-"`
}

// organizationScimConfigurationNewResponseJSON contains the JSON metadata for the
// struct [OrganizationScimConfigurationNewResponse]
type organizationScimConfigurationNewResponseJSON struct {
	Token             apijson.Field
	ScimConfiguration apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrganizationScimConfigurationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationScimConfigurationNewResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationScimConfigurationGetResponse struct {
	// scim_configuration is the SCIM configuration identified by the ID
	ScimConfiguration ScimConfiguration                            `json:"scimConfiguration,required"`
	JSON              organizationScimConfigurationGetResponseJSON `json:"-"`
}

// organizationScimConfigurationGetResponseJSON contains the JSON metadata for the
// struct [OrganizationScimConfigurationGetResponse]
type organizationScimConfigurationGetResponseJSON struct {
	ScimConfiguration apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrganizationScimConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationScimConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationScimConfigurationUpdateResponse struct {
	// scim_configuration is the updated SCIM configuration
	ScimConfiguration ScimConfiguration                               `json:"scimConfiguration,required"`
	JSON              organizationScimConfigurationUpdateResponseJSON `json:"-"`
}

// organizationScimConfigurationUpdateResponseJSON contains the JSON metadata for
// the struct [OrganizationScimConfigurationUpdateResponse]
type organizationScimConfigurationUpdateResponseJSON struct {
	ScimConfiguration apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrganizationScimConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationScimConfigurationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationScimConfigurationDeleteResponse = interface{}

type OrganizationScimConfigurationRegenerateTokenResponse struct {
	// token is the new bearer token for SCIM API authentication. This invalidates the
	// previous token - store it securely.
	Token string                                                   `json:"token,required"`
	JSON  organizationScimConfigurationRegenerateTokenResponseJSON `json:"-"`
}

// organizationScimConfigurationRegenerateTokenResponseJSON contains the JSON
// metadata for the struct [OrganizationScimConfigurationRegenerateTokenResponse]
type organizationScimConfigurationRegenerateTokenResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationScimConfigurationRegenerateTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationScimConfigurationRegenerateTokenResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationScimConfigurationNewParams struct {
	// organization_id is the ID of the organization to create the SCIM configuration
	// for
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
	// sso_configuration_id is the SSO configuration to link (required for user
	// provisioning)
	SSOConfigurationID param.Field[string] `json:"ssoConfigurationId,required" format:"uuid"`
	// name is a human-readable name for the SCIM configuration
	Name param.Field[string] `json:"name"`
}

func (r OrganizationScimConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationScimConfigurationGetParams struct {
	// scim_configuration_id is the ID of the SCIM configuration to get
	ScimConfigurationID param.Field[string] `json:"scimConfigurationId,required" format:"uuid"`
}

func (r OrganizationScimConfigurationGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationScimConfigurationUpdateParams struct {
	// scim_configuration_id is the ID of the SCIM configuration to update
	ScimConfigurationID param.Field[string] `json:"scimConfigurationId,required" format:"uuid"`
	// enabled controls whether SCIM provisioning is active
	Enabled param.Field[bool] `json:"enabled"`
	// name is a human-readable name for the SCIM configuration
	Name param.Field[string] `json:"name"`
	// sso_configuration_id is the SSO configuration to link
	SSOConfigurationID param.Field[string] `json:"ssoConfigurationId" format:"uuid"`
}

func (r OrganizationScimConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationScimConfigurationListParams struct {
	Token      param.Field[string]                                            `query:"token"`
	PageSize   param.Field[int64]                                             `query:"pageSize"`
	Pagination param.Field[OrganizationScimConfigurationListParamsPagination] `json:"pagination"`
}

func (r OrganizationScimConfigurationListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [OrganizationScimConfigurationListParams]'s query parameters
// as `url.Values`.
func (r OrganizationScimConfigurationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationScimConfigurationListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r OrganizationScimConfigurationListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationScimConfigurationDeleteParams struct {
	// scim_configuration_id is the ID of the SCIM configuration to delete
	ScimConfigurationID param.Field[string] `json:"scimConfigurationId,required" format:"uuid"`
}

func (r OrganizationScimConfigurationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationScimConfigurationRegenerateTokenParams struct {
	// scim_configuration_id is the ID of the SCIM configuration to regenerate token
	// for
	ScimConfigurationID param.Field[string] `json:"scimConfigurationId,required" format:"uuid"`
}

func (r OrganizationScimConfigurationRegenerateTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
