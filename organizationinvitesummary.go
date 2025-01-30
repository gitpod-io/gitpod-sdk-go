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

// OrganizationInviteSummaryService contains methods and other services that help
// with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationInviteSummaryService] method instead.
type OrganizationInviteSummaryService struct {
	Options []option.RequestOption
}

// NewOrganizationInviteSummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrganizationInviteSummaryService(opts ...option.RequestOption) (r *OrganizationInviteSummaryService) {
	r = &OrganizationInviteSummaryService{}
	r.Options = opts
	return
}

// GetOrganizationInviteSummary retrieves a summary of the organization based on an
// Invite ID.
//
// Used to discover which organization an invite is for.
func (r *OrganizationInviteSummaryService) Get(ctx context.Context, body OrganizationInviteSummaryGetParams, opts ...option.RequestOption) (res *OrganizationInviteSummaryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/GetOrganizationInviteSummary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type OrganizationInviteSummaryGetResponse struct {
	OrganizationID          string                                   `json:"organizationId" format:"uuid"`
	OrganizationMemberCount int64                                    `json:"organizationMemberCount"`
	OrganizationName        string                                   `json:"organizationName"`
	JSON                    organizationInviteSummaryGetResponseJSON `json:"-"`
}

// organizationInviteSummaryGetResponseJSON contains the JSON metadata for the
// struct [OrganizationInviteSummaryGetResponse]
type organizationInviteSummaryGetResponseJSON struct {
	OrganizationID          apijson.Field
	OrganizationMemberCount apijson.Field
	OrganizationName        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OrganizationInviteSummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInviteSummaryGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationInviteSummaryGetParams struct {
	InviteID param.Field[string] `json:"inviteId" format:"uuid"`
}

func (r OrganizationInviteSummaryGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
