// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
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

// Creates an invite link for joining an organization. Any existing
// OrganizationInvites are invalidated and can no longer be used.
//
// Use this method to:
//
// - Generate shareable invite links
// - Manage team growth
// - Control organization access
//
// ### Examples
//
// - Create organization invite:
//
//	Generates a new invite link for the organization.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *OrganizationInviteService) New(ctx context.Context, body OrganizationInviteNewParams, opts ...option.RequestOption) (res *OrganizationInviteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/CreateOrganizationInvite"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetOrganizationInvite
func (r *OrganizationInviteService) Get(ctx context.Context, body OrganizationInviteGetParams, opts ...option.RequestOption) (res *OrganizationInviteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/GetOrganizationInvite"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves organization details and membership info based on an invite link.
//
// Use this method to:
//
// - Preview organization details before joining
// - Validate invite link authenticity
// - Check organization size and activity
// - View team information before accepting
//
// ### Examples
//
// - Get invite summary:
//
//	Retrieves organization information from an invite.
//
//	```yaml
//	inviteId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *OrganizationInviteService) GetSummary(ctx context.Context, body OrganizationInviteGetSummaryParams, opts ...option.RequestOption) (res *OrganizationInviteGetSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/GetOrganizationInviteSummary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type OrganizationInvite struct {
	// invite_id is the unique identifier of the invite to join the organization. Use
	// JoinOrganization with this ID to join the organization.
	InviteID string                 `json:"inviteId,required" format:"uuid"`
	JSON     organizationInviteJSON `json:"-"`
}

// organizationInviteJSON contains the JSON metadata for the struct
// [OrganizationInvite]
type organizationInviteJSON struct {
	InviteID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInvite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInviteJSON) RawJSON() string {
	return r.raw
}

type OrganizationInviteNewResponse struct {
	Invite OrganizationInvite                `json:"invite,required"`
	JSON   organizationInviteNewResponseJSON `json:"-"`
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

type OrganizationInviteGetResponse struct {
	Invite OrganizationInvite                `json:"invite,required"`
	JSON   organizationInviteGetResponseJSON `json:"-"`
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

type OrganizationInviteGetSummaryResponse struct {
	OrganizationID          string                                   `json:"organizationId,required" format:"uuid"`
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
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r OrganizationInviteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationInviteGetParams struct {
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r OrganizationInviteGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationInviteGetSummaryParams struct {
	InviteID param.Field[string] `json:"inviteId,required" format:"uuid"`
}

func (r OrganizationInviteGetSummaryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
