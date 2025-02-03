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

// AccountService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// GetAccount retrieves a single Account.
func (r *AccountService) Get(ctx context.Context, params AccountGetParams, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/GetAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// DeleteAccount deletes an Account. To Delete an Account, the Account must not be
// an active member of any Organization.
func (r *AccountService) Delete(ctx context.Context, params AccountDeleteParams, opts ...option.RequestOption) (res *AccountDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/DeleteAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetSSOLoginURL returns the URL to redirect the user to for SSO login.
func (r *AccountService) GetSSOLoginURL(ctx context.Context, params AccountGetSSOLoginURLParams, opts ...option.RequestOption) (res *AccountGetSSOLoginURLResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/GetSSOLoginURL"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListLoginProviders returns the list of login providers matching the provided
// filters.
func (r *AccountService) ListLoginProviders(ctx context.Context, params AccountListLoginProvidersParams, opts ...option.RequestOption) (res *AccountListLoginProvidersResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/ListLoginProviders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type AccountGetResponse struct {
	Account AccountGetResponseAccount `json:"account"`
	JSON    accountGetResponseJSON    `json:"-"`
}

// accountGetResponseJSON contains the JSON metadata for the struct
// [AccountGetResponse]
type accountGetResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetResponseAccount struct {
	// organization_id is the ID of the organization the account is owned by if it's
	// created through custom SSO
	OrganizationID string                        `json:"organizationId,required"`
	JSON           accountGetResponseAccountJSON `json:"-"`
}

// accountGetResponseAccountJSON contains the JSON metadata for the struct
// [AccountGetResponseAccount]
type accountGetResponseAccountJSON struct {
	OrganizationID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountGetResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseAccountJSON) RawJSON() string {
	return r.raw
}

type AccountDeleteResponse = interface{}

type AccountGetSSOLoginURLResponse struct {
	// login_url is the URL to redirect the user to for SSO login
	LoginURL string                            `json:"loginUrl"`
	JSON     accountGetSSOLoginURLResponseJSON `json:"-"`
}

// accountGetSSOLoginURLResponseJSON contains the JSON metadata for the struct
// [AccountGetSSOLoginURLResponse]
type accountGetSSOLoginURLResponseJSON struct {
	LoginURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetSSOLoginURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetSSOLoginURLResponseJSON) RawJSON() string {
	return r.raw
}

type AccountListLoginProvidersResponse struct {
	LoginProviders []AccountListLoginProvidersResponseLoginProvider `json:"loginProviders"`
	Pagination     AccountListLoginProvidersResponsePagination      `json:"pagination"`
	JSON           accountListLoginProvidersResponseJSON            `json:"-"`
}

// accountListLoginProvidersResponseJSON contains the JSON metadata for the struct
// [AccountListLoginProvidersResponse]
type accountListLoginProvidersResponseJSON struct {
	LoginProviders apijson.Field
	Pagination     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountListLoginProvidersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListLoginProvidersResponseJSON) RawJSON() string {
	return r.raw
}

type AccountListLoginProvidersResponseLoginProvider struct {
	// login_url is the URL to redirect the browser agent to for login
	LoginURL string `json:"loginUrl"`
	// provider is the provider used by this login method, e.g. "github", "google",
	// "custom"
	Provider string                                             `json:"provider"`
	JSON     accountListLoginProvidersResponseLoginProviderJSON `json:"-"`
}

// accountListLoginProvidersResponseLoginProviderJSON contains the JSON metadata
// for the struct [AccountListLoginProvidersResponseLoginProvider]
type accountListLoginProvidersResponseLoginProviderJSON struct {
	LoginURL    apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListLoginProvidersResponseLoginProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListLoginProvidersResponseLoginProviderJSON) RawJSON() string {
	return r.raw
}

type AccountListLoginProvidersResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                          `json:"nextToken"`
	JSON      accountListLoginProvidersResponsePaginationJSON `json:"-"`
}

// accountListLoginProvidersResponsePaginationJSON contains the JSON metadata for
// the struct [AccountListLoginProvidersResponsePagination]
type accountListLoginProvidersResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListLoginProvidersResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListLoginProvidersResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type AccountGetParams struct {
	Body interface{} `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[AccountGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r AccountGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Define the version of the Connect protocol
type AccountGetParamsConnectProtocolVersion float64

const (
	AccountGetParamsConnectProtocolVersion1 AccountGetParamsConnectProtocolVersion = 1
)

func (r AccountGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case AccountGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type AccountDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[AccountDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	AccountID              param.Field[string]                                    `json:"accountId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r AccountDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type AccountDeleteParamsConnectProtocolVersion float64

const (
	AccountDeleteParamsConnectProtocolVersion1 AccountDeleteParamsConnectProtocolVersion = 1
)

func (r AccountDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case AccountDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type AccountGetSSOLoginURLParams struct {
	// return_to is the URL the user will be redirected to after login
	ReturnTo param.Field[string] `json:"returnTo,required" format:"uri"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[AccountGetSSOLoginURLParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// email is the email the user wants to login with
	Email param.Field[string] `json:"email" format:"email"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r AccountGetSSOLoginURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type AccountGetSSOLoginURLParamsConnectProtocolVersion float64

const (
	AccountGetSSOLoginURLParamsConnectProtocolVersion1 AccountGetSSOLoginURLParamsConnectProtocolVersion = 1
)

func (r AccountGetSSOLoginURLParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case AccountGetSSOLoginURLParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type AccountListLoginProvidersParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[AccountListLoginProvidersParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[AccountListLoginProvidersParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[AccountListLoginProvidersParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[AccountListLoginProvidersParamsConnect] `query:"connect"`
	Message param.Field[string]                                 `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [AccountListLoginProvidersParams]'s query parameters as
// `url.Values`.
func (r AccountListLoginProvidersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type AccountListLoginProvidersParamsEncoding string

const (
	AccountListLoginProvidersParamsEncodingProto AccountListLoginProvidersParamsEncoding = "proto"
	AccountListLoginProvidersParamsEncodingJson  AccountListLoginProvidersParamsEncoding = "json"
)

func (r AccountListLoginProvidersParamsEncoding) IsKnown() bool {
	switch r {
	case AccountListLoginProvidersParamsEncodingProto, AccountListLoginProvidersParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type AccountListLoginProvidersParamsConnectProtocolVersion float64

const (
	AccountListLoginProvidersParamsConnectProtocolVersion1 AccountListLoginProvidersParamsConnectProtocolVersion = 1
)

func (r AccountListLoginProvidersParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case AccountListLoginProvidersParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type AccountListLoginProvidersParamsCompression string

const (
	AccountListLoginProvidersParamsCompressionIdentity AccountListLoginProvidersParamsCompression = "identity"
	AccountListLoginProvidersParamsCompressionGzip     AccountListLoginProvidersParamsCompression = "gzip"
	AccountListLoginProvidersParamsCompressionBr       AccountListLoginProvidersParamsCompression = "br"
)

func (r AccountListLoginProvidersParamsCompression) IsKnown() bool {
	switch r {
	case AccountListLoginProvidersParamsCompressionIdentity, AccountListLoginProvidersParamsCompressionGzip, AccountListLoginProvidersParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type AccountListLoginProvidersParamsConnect string

const (
	AccountListLoginProvidersParamsConnectV1 AccountListLoginProvidersParamsConnect = "v1"
)

func (r AccountListLoginProvidersParamsConnect) IsKnown() bool {
	switch r {
	case AccountListLoginProvidersParamsConnectV1:
		return true
	}
	return false
}
