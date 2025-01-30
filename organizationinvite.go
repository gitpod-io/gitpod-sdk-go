// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
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
	Summary *OrganizationInviteSummaryService
}

// NewOrganizationInviteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationInviteService(opts ...option.RequestOption) (r *OrganizationInviteService) {
	r = &OrganizationInviteService{}
	r.Options = opts
	r.Summary = NewOrganizationInviteSummaryService(opts...)
	return
}

// CreateOrganizationInvite creates an invite for the organization. Any existing
// OrganizationInvites are invalidated and can no longer be used.
func (r *OrganizationInviteService) New(ctx context.Context, body OrganizationInviteNewParams, opts ...option.RequestOption) (res *OrganizationInviteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/CreateOrganizationInvite"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetOrganizationInvite retrieves invite for the organization. If no invite
// exists, a new one is created.
func (r *OrganizationInviteService) Get(ctx context.Context, body OrganizationInviteGetParams, opts ...option.RequestOption) (res *OrganizationInviteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/GetOrganizationInvite"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

type OrganizationInviteNewParams struct {
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
}

func (r OrganizationInviteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationInviteGetParams struct {
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
}

func (r OrganizationInviteGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
