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
func (r *AccountService) Get(ctx context.Context, body AccountGetParams, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/GetAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// DeleteAccount deletes an Account. To Delete an Account, the Account must not be
// an active member of any Organization.
func (r *AccountService) Delete(ctx context.Context, body AccountDeleteParams, opts ...option.RequestOption) (res *AccountDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/DeleteAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetSSOLoginURL returns the URL to redirect the user to for SSO login.
func (r *AccountService) GetSSOLoginURL(ctx context.Context, body AccountGetSSOLoginURLParams, opts ...option.RequestOption) (res *AccountGetSSOLoginURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/GetSSOLoginURL"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListLoginProviders returns the list of login providers matching the provided
// filters.
func (r *AccountService) ListLoginProviders(ctx context.Context, params AccountListLoginProvidersParams, opts ...option.RequestOption) (res *pagination.LoginProvidersPage[AccountListLoginProvidersResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.AccountService/ListLoginProviders"
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

// ListLoginProviders returns the list of login providers matching the provided
// filters.
func (r *AccountService) ListLoginProvidersAutoPaging(ctx context.Context, params AccountListLoginProvidersParams, opts ...option.RequestOption) *pagination.LoginProvidersPageAutoPager[AccountListLoginProvidersResponse] {
	return pagination.NewLoginProvidersPageAutoPager(r.ListLoginProviders(ctx, params, opts...))
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
	// login_url is the URL to redirect the browser agent to for login
	LoginURL string `json:"loginUrl"`
	// provider is the provider used by this login method, e.g. "github", "google",
	// "custom"
	Provider string                                `json:"provider"`
	JSON     accountListLoginProvidersResponseJSON `json:"-"`
}

// accountListLoginProvidersResponseJSON contains the JSON metadata for the struct
// [AccountListLoginProvidersResponse]
type accountListLoginProvidersResponseJSON struct {
	LoginURL    apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListLoginProvidersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListLoginProvidersResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDeleteParams struct {
	AccountID param.Field[string] `json:"accountId" format:"uuid"`
}

func (r AccountDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGetSSOLoginURLParams struct {
	// return_to is the URL the user will be redirected to after login
	ReturnTo param.Field[string] `json:"returnTo,required" format:"uri"`
	// email is the email the user wants to login with
	Email param.Field[string] `json:"email" format:"email"`
}

func (r AccountGetSSOLoginURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountListLoginProvidersParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// filter contains the filter options for listing login methods
	Filter param.Field[AccountListLoginProvidersParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing login methods
	Pagination param.Field[AccountListLoginProvidersParamsPagination] `json:"pagination"`
}

func (r AccountListLoginProvidersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountListLoginProvidersParams]'s query parameters as
// `url.Values`.
func (r AccountListLoginProvidersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// filter contains the filter options for listing login methods
type AccountListLoginProvidersParamsFilter struct {
	// invite_id is the ID of the invite URL the user wants to login with
	InviteID param.Field[string] `json:"inviteId" format:"uuid"`
}

func (r AccountListLoginProvidersParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing login methods
type AccountListLoginProvidersParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r AccountListLoginProvidersParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
