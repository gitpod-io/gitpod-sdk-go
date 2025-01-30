// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
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
func (r *IdentityService) ExchangeToken(ctx context.Context, params IdentityExchangeTokenParams, opts ...option.RequestOption) (res *IdentityExchangeTokenResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.IdentityService/ExchangeToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetAuthenticatedIdentity allows to retrieve the current identity.
func (r *IdentityService) GetAuthenticatedIdentity(ctx context.Context, params IdentityGetAuthenticatedIdentityParams, opts ...option.RequestOption) (res *IdentityGetAuthenticatedIdentityResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.IdentityService/GetAuthenticatedIdentity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetIDToken returns a token that can be used to authenticate the user against the
// other services.
func (r *IdentityService) GetIDToken(ctx context.Context, params IdentityGetIDTokenParams, opts ...option.RequestOption) (res *IdentityGetIDTokenResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.IdentityService/GetIDToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
	Subject IdentityGetAuthenticatedIdentityResponseSubject `json:"subject"`
	JSON    identityGetAuthenticatedIdentityResponseJSON    `json:"-"`
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

// subject is the identity of the current user
type IdentityGetAuthenticatedIdentityResponseSubject struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal IdentityGetAuthenticatedIdentityResponseSubjectPrincipal `json:"principal"`
	JSON      identityGetAuthenticatedIdentityResponseSubjectJSON      `json:"-"`
}

// identityGetAuthenticatedIdentityResponseSubjectJSON contains the JSON metadata
// for the struct [IdentityGetAuthenticatedIdentityResponseSubject]
type identityGetAuthenticatedIdentityResponseSubjectJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityGetAuthenticatedIdentityResponseSubject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityGetAuthenticatedIdentityResponseSubjectJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type IdentityGetAuthenticatedIdentityResponseSubjectPrincipal string

const (
	IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalUnspecified    IdentityGetAuthenticatedIdentityResponseSubjectPrincipal = "PRINCIPAL_UNSPECIFIED"
	IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalAccount        IdentityGetAuthenticatedIdentityResponseSubjectPrincipal = "PRINCIPAL_ACCOUNT"
	IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalUser           IdentityGetAuthenticatedIdentityResponseSubjectPrincipal = "PRINCIPAL_USER"
	IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalRunner         IdentityGetAuthenticatedIdentityResponseSubjectPrincipal = "PRINCIPAL_RUNNER"
	IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalEnvironment    IdentityGetAuthenticatedIdentityResponseSubjectPrincipal = "PRINCIPAL_ENVIRONMENT"
	IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalServiceAccount IdentityGetAuthenticatedIdentityResponseSubjectPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r IdentityGetAuthenticatedIdentityResponseSubjectPrincipal) IsKnown() bool {
	switch r {
	case IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalUnspecified, IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalAccount, IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalUser, IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalRunner, IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalEnvironment, IdentityGetAuthenticatedIdentityResponseSubjectPrincipalPrincipalServiceAccount:
		return true
	}
	return false
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
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[IdentityExchangeTokenParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// exchange_token is the token to exchange
	ExchangeToken param.Field[string] `json:"exchangeToken"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r IdentityExchangeTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type IdentityExchangeTokenParamsConnectProtocolVersion float64

const (
	IdentityExchangeTokenParamsConnectProtocolVersion1 IdentityExchangeTokenParamsConnectProtocolVersion = 1
)

func (r IdentityExchangeTokenParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case IdentityExchangeTokenParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type IdentityGetAuthenticatedIdentityParams struct {
	Body interface{} `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[IdentityGetAuthenticatedIdentityParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r IdentityGetAuthenticatedIdentityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Define the version of the Connect protocol
type IdentityGetAuthenticatedIdentityParamsConnectProtocolVersion float64

const (
	IdentityGetAuthenticatedIdentityParamsConnectProtocolVersion1 IdentityGetAuthenticatedIdentityParamsConnectProtocolVersion = 1
)

func (r IdentityGetAuthenticatedIdentityParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case IdentityGetAuthenticatedIdentityParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type IdentityGetIDTokenParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[IdentityGetIDTokenParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Audience               param.Field[[]string]                                       `json:"audience"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r IdentityGetIDTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type IdentityGetIDTokenParamsConnectProtocolVersion float64

const (
	IdentityGetIDTokenParamsConnectProtocolVersion1 IdentityGetIDTokenParamsConnectProtocolVersion = 1
)

func (r IdentityGetIDTokenParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case IdentityGetIDTokenParamsConnectProtocolVersion1:
		return true
	}
	return false
}
