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

// OrganizationInviteService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationInviteService] method instead.
type OrganizationInviteService struct {
	Options []option.RequestOption
}

// NewOrganizationInviteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationInviteService(opts ...option.RequestOption) (r *OrganizationInviteService) {
	r = &OrganizationInviteService{}
	r.Options = opts
	return
}

// CreateOrganizationInvite creates an invite for the organization. Any existing
// OrganizationInvites are invalidated and can no longer be used.
func (r *OrganizationInviteService) New(ctx context.Context, params OrganizationInviteNewParams, opts ...option.RequestOption) (res *OrganizationInviteNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/CreateOrganizationInvite"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetOrganizationInvite retrieves invite for the organization. If no invite
// exists, a new one is created.
func (r *OrganizationInviteService) Get(ctx context.Context, params OrganizationInviteGetParams, opts ...option.RequestOption) (res *OrganizationInviteGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/GetOrganizationInvite"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetOrganizationInviteSummary retrieves a summary of the organization based on an
// Invite ID.
//
// Used to discover which organization an invite is for.
func (r *OrganizationInviteService) GetSummary(ctx context.Context, params OrganizationInviteGetSummaryParams, opts ...option.RequestOption) (res *OrganizationInviteGetSummaryResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/GetOrganizationInviteSummary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type OrganizationInviteNewResponse struct {
	Invite OrganizationInviteNewResponseInvite `json:"invite"`
	JSON   organizationInviteNewResponseJSON   `json:"-"`
}

// organizationInviteNewResponseJSON contains the JSON metadata for the struct
// [OrganizationInviteNewResponse]
type organizationInviteNewResponseJSON struct {
	Invite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInviteNewResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationInviteNewResponseInvite struct {
	// invite_id is the unique identifier of the invite to join the organization.
	//
	// Use JoinOrganization with this ID to join the organization.
	InviteID string                                  `json:"inviteId" format:"uuid"`
	JSON     organizationInviteNewResponseInviteJSON `json:"-"`
}

// organizationInviteNewResponseInviteJSON contains the JSON metadata for the
// struct [OrganizationInviteNewResponseInvite]
type organizationInviteNewResponseInviteJSON struct {
	InviteID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteNewResponseInvite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInviteNewResponseInviteJSON) RawJSON() string {
	return r.raw
}

type OrganizationInviteGetResponse struct {
	Invite OrganizationInviteGetResponseInvite `json:"invite"`
	JSON   organizationInviteGetResponseJSON   `json:"-"`
}

// organizationInviteGetResponseJSON contains the JSON metadata for the struct
// [OrganizationInviteGetResponse]
type organizationInviteGetResponseJSON struct {
	Invite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInviteGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationInviteGetResponseInvite struct {
	// invite_id is the unique identifier of the invite to join the organization.
	//
	// Use JoinOrganization with this ID to join the organization.
	InviteID string                                  `json:"inviteId" format:"uuid"`
	JSON     organizationInviteGetResponseInviteJSON `json:"-"`
}

// organizationInviteGetResponseInviteJSON contains the JSON metadata for the
// struct [OrganizationInviteGetResponseInvite]
type organizationInviteGetResponseInviteJSON struct {
	InviteID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteGetResponseInvite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInviteGetResponseInviteJSON) RawJSON() string {
	return r.raw
}

type OrganizationInviteGetSummaryResponse struct {
	OrganizationID          string                                   `json:"organizationId" format:"uuid"`
	OrganizationMemberCount int64                                    `json:"organizationMemberCount"`
	OrganizationName        string                                   `json:"organizationName"`
	JSON                    organizationInviteGetSummaryResponseJSON `json:"-"`
}

// organizationInviteGetSummaryResponseJSON contains the JSON metadata for the
// struct [OrganizationInviteGetSummaryResponse]
type organizationInviteGetSummaryResponseJSON struct {
	OrganizationID          apijson.Field
	OrganizationMemberCount apijson.Field
	OrganizationName        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OrganizationInviteGetSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInviteGetSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationInviteNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationInviteNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	OrganizationID         param.Field[string]                                            `json:"organizationId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationInviteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationInviteNewParamsConnectProtocolVersion float64

const (
	OrganizationInviteNewParamsConnectProtocolVersion1 OrganizationInviteNewParamsConnectProtocolVersion = 1
)

func (r OrganizationInviteNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationInviteNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationInviteGetParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationInviteGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	OrganizationID         param.Field[string]                                            `json:"organizationId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationInviteGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationInviteGetParamsConnectProtocolVersion float64

const (
	OrganizationInviteGetParamsConnectProtocolVersion1 OrganizationInviteGetParamsConnectProtocolVersion = 1
)

func (r OrganizationInviteGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationInviteGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationInviteGetSummaryParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[OrganizationInviteGetSummaryParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationInviteGetSummaryParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[OrganizationInviteGetSummaryParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[OrganizationInviteGetSummaryParamsConnect] `query:"connect"`
	Message param.Field[string]                                    `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [OrganizationInviteGetSummaryParams]'s query parameters as
// `url.Values`.
func (r OrganizationInviteGetSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type OrganizationInviteGetSummaryParamsEncoding string

const (
	OrganizationInviteGetSummaryParamsEncodingProto OrganizationInviteGetSummaryParamsEncoding = "proto"
	OrganizationInviteGetSummaryParamsEncodingJson  OrganizationInviteGetSummaryParamsEncoding = "json"
)

func (r OrganizationInviteGetSummaryParamsEncoding) IsKnown() bool {
	switch r {
	case OrganizationInviteGetSummaryParamsEncodingProto, OrganizationInviteGetSummaryParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type OrganizationInviteGetSummaryParamsConnectProtocolVersion float64

const (
	OrganizationInviteGetSummaryParamsConnectProtocolVersion1 OrganizationInviteGetSummaryParamsConnectProtocolVersion = 1
)

func (r OrganizationInviteGetSummaryParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationInviteGetSummaryParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type OrganizationInviteGetSummaryParamsCompression string

const (
	OrganizationInviteGetSummaryParamsCompressionIdentity OrganizationInviteGetSummaryParamsCompression = "identity"
	OrganizationInviteGetSummaryParamsCompressionGzip     OrganizationInviteGetSummaryParamsCompression = "gzip"
	OrganizationInviteGetSummaryParamsCompressionBr       OrganizationInviteGetSummaryParamsCompression = "br"
)

func (r OrganizationInviteGetSummaryParamsCompression) IsKnown() bool {
	switch r {
	case OrganizationInviteGetSummaryParamsCompressionIdentity, OrganizationInviteGetSummaryParamsCompressionGzip, OrganizationInviteGetSummaryParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type OrganizationInviteGetSummaryParamsConnect string

const (
	OrganizationInviteGetSummaryParamsConnectV1 OrganizationInviteGetSummaryParamsConnect = "v1"
)

func (r OrganizationInviteGetSummaryParamsConnect) IsKnown() bool {
	switch r {
	case OrganizationInviteGetSummaryParamsConnectV1:
		return true
	}
	return false
}
