// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
)

// OrganizationCustomDomainService contains methods and other services that help
// with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationCustomDomainService] method instead.
type OrganizationCustomDomainService struct {
	Options []option.RequestOption
}

// NewOrganizationCustomDomainService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrganizationCustomDomainService(opts ...option.RequestOption) (r *OrganizationCustomDomainService) {
	r = &OrganizationCustomDomainService{}
	r.Options = opts
	return
}

// Creates a custom domain configuration for an organization.
//
// # Use this method to configure custom domains for organization workspaces
//
// ### Examples
//
// - Configure AWS custom domain:
//
//	Sets up a custom domain with AWS provider.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	domainName: "workspaces.acme-corp.com"
//	provider: CUSTOM_DOMAIN_PROVIDER_AWS
//	awsAccountId: "123456789012"
//	```
func (r *OrganizationCustomDomainService) New(ctx context.Context, body OrganizationCustomDomainNewParams, opts ...option.RequestOption) (res *OrganizationCustomDomainNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/CreateCustomDomain"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a specific custom domain configuration.
//
// # Use this method to view custom domain details
//
// ### Examples
//
// - Get custom domain configuration:
//
//	Retrieves details of a specific custom domain.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *OrganizationCustomDomainService) Get(ctx context.Context, body OrganizationCustomDomainGetParams, opts ...option.RequestOption) (res *OrganizationCustomDomainGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/GetCustomDomain"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates custom domain configuration settings.
//
// Use this method to:
//
// - Update cloud provider settings
// - Change AWS account ID
// - Modify domain configuration
//
// ### Examples
//
// - Update AWS account ID:
//
//	Changes the AWS account ID for the custom domain.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	domainName: "workspaces.acme-corp.com"
//	awsAccountId: "987654321098"
//	```
func (r *OrganizationCustomDomainService) Update(ctx context.Context, body OrganizationCustomDomainUpdateParams, opts ...option.RequestOption) (res *OrganizationCustomDomainUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/UpdateCustomDomain"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes a custom domain configuration from an organization.
//
// Use this method to:
//
// - Disable custom domain functionality
// - Remove outdated configurations
// - Clean up unused domains
//
// ### Examples
//
// - Delete custom domain configuration:
//
//	Removes a specific custom domain configuration.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *OrganizationCustomDomainService) Delete(ctx context.Context, body OrganizationCustomDomainDeleteParams, opts ...option.RequestOption) (res *OrganizationCustomDomainDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/DeleteCustomDomain"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// CustomDomain represents a custom domain configuration for an organization
type CustomDomain struct {
	// id is the unique identifier of the custom domain
	ID string `json:"id,required" format:"uuid"`
	// created_at is when the custom domain was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// domain_name is the custom domain name
	DomainName string `json:"domainName,required"`
	// organization_id is the ID of the organization this custom domain belongs to
	OrganizationID string `json:"organizationId,required" format:"uuid"`
	// updated_at is when the custom domain was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// aws_account_id is the AWS account ID (deprecated: use cloud_account_id)
	//
	// Deprecated: deprecated
	AwsAccountID string `json:"awsAccountId"`
	// cloud_account_id is the unified cloud account identifier (AWS Account ID or GCP
	// Project ID)
	CloudAccountID string `json:"cloudAccountId"`
	// provider is the cloud provider for this custom domain
	Provider CustomDomainProvider `json:"provider"`
	JSON     customDomainJSON     `json:"-"`
}

// customDomainJSON contains the JSON metadata for the struct [CustomDomain]
type customDomainJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	DomainName     apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	AwsAccountID   apijson.Field
	CloudAccountID apijson.Field
	Provider       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customDomainJSON) RawJSON() string {
	return r.raw
}

// CustomDomainProvider represents the cloud provider for custom domain
// configuration
type CustomDomainProvider string

const (
	CustomDomainProviderUnspecified CustomDomainProvider = "CUSTOM_DOMAIN_PROVIDER_UNSPECIFIED"
	CustomDomainProviderAws         CustomDomainProvider = "CUSTOM_DOMAIN_PROVIDER_AWS"
	CustomDomainProviderGcp         CustomDomainProvider = "CUSTOM_DOMAIN_PROVIDER_GCP"
)

func (r CustomDomainProvider) IsKnown() bool {
	switch r {
	case CustomDomainProviderUnspecified, CustomDomainProviderAws, CustomDomainProviderGcp:
		return true
	}
	return false
}

// CreateCustomDomainResponse is the response message for creating a custom domain
type OrganizationCustomDomainNewResponse struct {
	// custom_domain is the created custom domain
	CustomDomain CustomDomain                            `json:"customDomain,required"`
	JSON         organizationCustomDomainNewResponseJSON `json:"-"`
}

// organizationCustomDomainNewResponseJSON contains the JSON metadata for the
// struct [OrganizationCustomDomainNewResponse]
type organizationCustomDomainNewResponseJSON struct {
	CustomDomain apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationCustomDomainNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationCustomDomainNewResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationCustomDomainGetResponse struct {
	// CustomDomain represents a custom domain configuration for an organization
	CustomDomain CustomDomain                            `json:"customDomain,required"`
	JSON         organizationCustomDomainGetResponseJSON `json:"-"`
}

// organizationCustomDomainGetResponseJSON contains the JSON metadata for the
// struct [OrganizationCustomDomainGetResponse]
type organizationCustomDomainGetResponseJSON struct {
	CustomDomain apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationCustomDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationCustomDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

// UpdateCustomDomainResponse is the response message for updating a custom domain
type OrganizationCustomDomainUpdateResponse struct {
	// custom_domain is the updated custom domain
	CustomDomain CustomDomain                               `json:"customDomain,required"`
	JSON         organizationCustomDomainUpdateResponseJSON `json:"-"`
}

// organizationCustomDomainUpdateResponseJSON contains the JSON metadata for the
// struct [OrganizationCustomDomainUpdateResponse]
type organizationCustomDomainUpdateResponseJSON struct {
	CustomDomain apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationCustomDomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationCustomDomainUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationCustomDomainDeleteResponse = interface{}

type OrganizationCustomDomainNewParams struct {
	// domain_name is the custom domain name
	DomainName param.Field[string] `json:"domainName,required"`
	// organization_id is the ID of the organization to create the custom domain for
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
	// aws_account_id is the AWS account ID (deprecated: use cloud_account_id)
	AwsAccountID param.Field[string] `json:"awsAccountId"`
	// cloud_account_id is the unified cloud account identifier (AWS Account ID or GCP
	// Project ID)
	CloudAccountID param.Field[string] `json:"cloudAccountId"`
	// provider is the cloud provider for this custom domain
	Provider param.Field[CustomDomainProvider] `json:"provider"`
}

func (r OrganizationCustomDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationCustomDomainGetParams struct {
	// organization_id is the ID of the organization to retrieve custom domain for
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r OrganizationCustomDomainGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationCustomDomainUpdateParams struct {
	// domain_name is the custom domain name
	DomainName param.Field[string] `json:"domainName,required"`
	// organization_id is the ID of the organization to update custom domain for
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
	// aws_account_id is the AWS account ID (deprecated: use cloud_account_id)
	AwsAccountID param.Field[string] `json:"awsAccountId"`
	// cloud_account_id is the unified cloud account identifier (AWS Account ID or GCP
	// Project ID)
	CloudAccountID param.Field[string] `json:"cloudAccountId"`
	// provider is the cloud provider for this custom domain
	Provider param.Field[CustomDomainProvider] `json:"provider"`
}

func (r OrganizationCustomDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationCustomDomainDeleteParams struct {
	// organization_id is the ID of the organization to delete custom domain for
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r OrganizationCustomDomainDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
