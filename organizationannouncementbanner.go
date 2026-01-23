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
)

// OrganizationAnnouncementBannerService contains methods and other services that
// help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationAnnouncementBannerService] method instead.
type OrganizationAnnouncementBannerService struct {
	Options []option.RequestOption
}

// NewOrganizationAnnouncementBannerService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrganizationAnnouncementBannerService(opts ...option.RequestOption) (r *OrganizationAnnouncementBannerService) {
	r = &OrganizationAnnouncementBannerService{}
	r.Options = opts
	return
}

// Updates the announcement banner configuration for an organization.
//
// Use this method to configure the announcement banner displayed to all users.
// Only organization admins can update the banner. Requires Enterprise tier.
//
// ### Examples
//
// - Enable announcement banner:
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	message: "Scheduled maintenance on Saturday 10pm-2am UTC"
//	enabled: true
//	```
//
// - Disable announcement banner:
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	enabled: false
//	```
func (r *OrganizationAnnouncementBannerService) Update(ctx context.Context, body OrganizationAnnouncementBannerUpdateParams, opts ...option.RequestOption) (res *OrganizationAnnouncementBannerUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/UpdateAnnouncementBanner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the announcement banner configuration for an organization.
//
// Use this method to fetch the current announcement banner settings. All
// organization members can read the banner configuration.
//
// ### Examples
//
// - Get announcement banner:
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *OrganizationAnnouncementBannerService) Get(ctx context.Context, body OrganizationAnnouncementBannerGetParams, opts ...option.RequestOption) (res *OrganizationAnnouncementBannerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/GetAnnouncementBanner"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AnnouncementBanner struct {
	// organization_id is the ID of the organization
	OrganizationID string `json:"organizationId,required" format:"uuid"`
	// enabled controls whether the banner is displayed
	Enabled bool `json:"enabled"`
	// message is the banner message displayed to users. Supports basic Markdown.
	Message string                 `json:"message"`
	JSON    announcementBannerJSON `json:"-"`
}

// announcementBannerJSON contains the JSON metadata for the struct
// [AnnouncementBanner]
type announcementBannerJSON struct {
	OrganizationID apijson.Field
	Enabled        apijson.Field
	Message        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AnnouncementBanner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r announcementBannerJSON) RawJSON() string {
	return r.raw
}

type OrganizationAnnouncementBannerUpdateResponse struct {
	// banner is the updated announcement banner configuration
	Banner AnnouncementBanner                               `json:"banner,required"`
	JSON   organizationAnnouncementBannerUpdateResponseJSON `json:"-"`
}

// organizationAnnouncementBannerUpdateResponseJSON contains the JSON metadata for
// the struct [OrganizationAnnouncementBannerUpdateResponse]
type organizationAnnouncementBannerUpdateResponseJSON struct {
	Banner      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationAnnouncementBannerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAnnouncementBannerUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationAnnouncementBannerGetResponse struct {
	// banner is the announcement banner configuration
	Banner AnnouncementBanner                            `json:"banner,required"`
	JSON   organizationAnnouncementBannerGetResponseJSON `json:"-"`
}

// organizationAnnouncementBannerGetResponseJSON contains the JSON metadata for the
// struct [OrganizationAnnouncementBannerGetResponse]
type organizationAnnouncementBannerGetResponseJSON struct {
	Banner      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationAnnouncementBannerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAnnouncementBannerGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationAnnouncementBannerUpdateParams struct {
	// organization_id is the ID of the organization
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
	// enabled controls whether the banner is displayed
	Enabled param.Field[bool] `json:"enabled"`
	// message is the banner message. Supports basic Markdown. Maximum 1000 characters.
	Message param.Field[string] `json:"message"`
}

func (r OrganizationAnnouncementBannerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationAnnouncementBannerGetParams struct {
	// organization_id is the ID of the organization
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r OrganizationAnnouncementBannerGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
