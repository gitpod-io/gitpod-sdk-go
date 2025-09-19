// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"slices"

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

// Exchanges an exchange token for a new access token.
//
// Use this method to:
//
// - Convert exchange tokens to access tokens
// - Obtain new access credentials
// - Complete token exchange flows
//
// ### Examples
//
// - Exchange token:
//
//	Trades an exchange token for an access token.
//
//	```yaml
//	exchangeToken: "exchange-token-value"
//	```
func (r *IdentityService) ExchangeToken(ctx context.Context, body IdentityExchangeTokenParams, opts ...option.RequestOption) (res *IdentityExchangeTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.IdentityService/ExchangeToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves information about the currently authenticated identity.
//
// Use this method to:
//
// - Get current user information
// - Check authentication status
// - Retrieve organization context
// - Validate authentication principal
//
// ### Examples
//
// - Get current identity:
//
//	Retrieves details about the authenticated user.
//
//	```yaml
//	{}
//	```
func (r *IdentityService) GetAuthenticatedIdentity(ctx context.Context, body IdentityGetAuthenticatedIdentityParams, opts ...option.RequestOption) (res *IdentityGetAuthenticatedIdentityResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.IdentityService/GetAuthenticatedIdentity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets an ID token for authenticating with other services.
//
// Use this method to:
//
// - Obtain authentication tokens for service-to-service calls
// - Access protected resources
// - Generate scoped access tokens
//
// ### Examples
//
// - Get token for single service:
//
//	Retrieves a token for authenticating with one service.
//
//	```yaml
//	audience:
//	  - "https://api.gitpod.io"
//	```
//
// - Get token for multiple services:
//
//	Retrieves a token valid for multiple services.
//
//	```yaml
//	audience:
//	  - "https://api.gitpod.io"
//	  - "https://ws.gitpod.io"
//	```
func (r *IdentityService) GetIDToken(ctx context.Context, body IdentityGetIDTokenParams, opts ...option.RequestOption) (res *IdentityGetIDTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.IdentityService/GetIDToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IDTokenVersion string

const (
	IDTokenVersionUnspecified IDTokenVersion = "ID_TOKEN_VERSION_UNSPECIFIED"
	IDTokenVersionV1          IDTokenVersion = "ID_TOKEN_VERSION_V1"
	IDTokenVersionV2          IDTokenVersion = "ID_TOKEN_VERSION_V2"
)

func (r IDTokenVersion) IsKnown() bool {
	switch r {
	case IDTokenVersionUnspecified, IDTokenVersionV1, IDTokenVersionV2:
		return true
	}
	return false
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
	// version is the version of the ID token.
	Version param.Field[IDTokenVersion] `json:"version"`
}

func (r IdentityGetIDTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
