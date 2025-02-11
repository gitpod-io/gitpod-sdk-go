// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// IdentityService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityService] method instead.
type IdentityService struct {
	Options []option.RequestOption
}

// NewIdentityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIdentityService(opts ...option.RequestOption) (r *IdentityService) {
	r = &IdentityService{}
	r.Options = opts
	return
}

// ExchangeToken trades an exchange token for a new access token.
func (r *IdentityService) ExchangeToken(ctx context.Context, body IdentityExchangeTokenParams, opts ...option.RequestOption) (res *IdentityExchangeTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.IdentityService/ExchangeToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetAuthenticatedIdentity allows to retrieve the current identity.
func (r *IdentityService) GetAuthenticatedIdentity(ctx context.Context, body IdentityGetAuthenticatedIdentityParams, opts ...option.RequestOption) (res *IdentityGetAuthenticatedIdentityResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.IdentityService/GetAuthenticatedIdentity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetIDToken returns a token that can be used to authenticate the user against the
// other services.
func (r *IdentityService) GetIDToken(ctx context.Context, body IdentityGetIDTokenParams, opts ...option.RequestOption) (res *IdentityGetIDTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.IdentityService/GetIDToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IdentityExchangeTokenResponse struct {
	// access_token is the new access token
	AccessToken string                            `json:"accessToken"`
	JSON        identityExchangeTokenResponseJSON `json:"-"`
}

// identityExchangeTokenResponseJSON contains the JSON metadata for the struct
// [IdentityExchangeTokenResponse]
type identityExchangeTokenResponseJSON struct {
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityExchangeTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityExchangeTokenResponseJSON) RawJSON() string {
	return r.raw
}

type IdentityGetAuthenticatedIdentityResponse struct {
	OrganizationID string `json:"organizationId"`
	// subject is the identity of the current user
	Subject shared.Subject                               `json:"subject"`
	JSON    identityGetAuthenticatedIdentityResponseJSON `json:"-"`
}

// identityGetAuthenticatedIdentityResponseJSON contains the JSON metadata for the
// struct [IdentityGetAuthenticatedIdentityResponse]
type identityGetAuthenticatedIdentityResponseJSON struct {
	OrganizationID apijson.Field
	Subject        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityGetAuthenticatedIdentityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityGetAuthenticatedIdentityResponseJSON) RawJSON() string {
	return r.raw
}

type IdentityGetIDTokenResponse struct {
	Token string                         `json:"token"`
	JSON  identityGetIDTokenResponseJSON `json:"-"`
}

// identityGetIDTokenResponseJSON contains the JSON metadata for the struct
// [IdentityGetIDTokenResponse]
type identityGetIDTokenResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityGetIDTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityGetIDTokenResponseJSON) RawJSON() string {
	return r.raw
}

type IdentityExchangeTokenParams struct {
	// exchange_token is the token to exchange
	ExchangeToken param.Field[string] `json:"exchangeToken"`
}

func (r IdentityExchangeTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityGetAuthenticatedIdentityParams struct {
	Empty param.Field[bool] `json:"empty"`
}

func (r IdentityGetAuthenticatedIdentityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityGetIDTokenParams struct {
	Audience param.Field[[]string] `json:"audience"`
}

func (r IdentityGetIDTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
