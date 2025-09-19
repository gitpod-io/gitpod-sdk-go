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

// OrganizationPolicyService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationPolicyService] method instead.
type OrganizationPolicyService struct {
	Options []option.RequestOption
}

// NewOrganizationPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationPolicyService(opts ...option.RequestOption) (r *OrganizationPolicyService) {
	r = &OrganizationPolicyService{}
	r.Options = opts
	return
}

// Gets organization policy settings by organization ID.
//
// Use this method to:
//
// - Retrieve current policy settings for an organization
// - View resource limits and restrictions
// - Check allowed editors and other configurations
//
// ### Examples
//
// - Get organization policies:
//
//	Retrieves policy settings for a specific organization.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *OrganizationPolicyService) Get(ctx context.Context, body OrganizationPolicyGetParams, opts ...option.RequestOption) (res *OrganizationPolicyGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/GetOrganizationPolicies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates organization policy settings.
//
// Use this method to:
//
// - Configure editor restrictions
// - Set environment resource limits
// - Define project creation permissions
// - Customize default configurations
//
// ### Examples
//
// - Update editor policies:
//
//	Restricts available editors and sets a default.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	allowedEditorIds:
//	  - "vscode"
//	  - "jetbrains"
//	defaultEditorId: "vscode"
//	```
//
// - Set environment limits:
//
//	Configures limits for environment usage.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	maximumEnvironmentTimeout: "3600s"
//	maximumRunningEnvironmentsPerUser: "5"
//	maximumEnvironmentsPerUser: "20"
//	```
func (r *OrganizationPolicyService) Update(ctx context.Context, body OrganizationPolicyUpdateParams, opts ...option.RequestOption) (res *OrganizationPolicyUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/UpdateOrganizationPolicies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type OrganizationPolicies struct {
	// allowed_editor_ids is the list of editor IDs that are allowed to be used in the
	// organization
	AllowedEditorIDs []string `json:"allowedEditorIds,required"`
	// allow_local_runners controls whether local runners are allowed to be used in the
	// organization
	AllowLocalRunners bool `json:"allowLocalRunners,required"`
	// default_editor_id is the default editor ID to be used when a user doesn't
	// specify one
	DefaultEditorID string `json:"defaultEditorId,required"`
	// default_environment_image is the default container image when none is defined in
	// repo
	DefaultEnvironmentImage string `json:"defaultEnvironmentImage,required"`
	// maximum_environments_per_user limits total environments (running or stopped) per
	// user
	MaximumEnvironmentsPerUser string `json:"maximumEnvironmentsPerUser,required"`
	// maximum_running_environments_per_user limits simultaneously running environments
	// per user
	MaximumRunningEnvironmentsPerUser string `json:"maximumRunningEnvironmentsPerUser,required"`
	// members_create_projects controls whether members can create projects
	MembersCreateProjects bool `json:"membersCreateProjects,required"`
	// members_require_projects controls whether environments can only be created from
	// projects by non-admin users
	MembersRequireProjects bool `json:"membersRequireProjects,required"`
	// organization_id is the ID of the organization
	OrganizationID string `json:"organizationId,required" format:"uuid"`
	// port_sharing_disabled controls whether port sharing is disabled in the
	// organization
	PortSharingDisabled bool `json:"portSharingDisabled,required"`
	// maximum_environment_timeout controls the maximum timeout allowed for
	// environments in seconds. 0 means no limit (never). Minimum duration is 30
	// minutes.
	MaximumEnvironmentTimeout string                   `json:"maximumEnvironmentTimeout" format:"regex"`
	JSON                      organizationPoliciesJSON `json:"-"`
}

// organizationPoliciesJSON contains the JSON metadata for the struct
// [OrganizationPolicies]
type organizationPoliciesJSON struct {
	AllowedEditorIDs                  apijson.Field
	AllowLocalRunners                 apijson.Field
	DefaultEditorID                   apijson.Field
	DefaultEnvironmentImage           apijson.Field
	MaximumEnvironmentsPerUser        apijson.Field
	MaximumRunningEnvironmentsPerUser apijson.Field
	MembersCreateProjects             apijson.Field
	MembersRequireProjects            apijson.Field
	OrganizationID                    apijson.Field
	PortSharingDisabled               apijson.Field
	MaximumEnvironmentTimeout         apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrganizationPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationPoliciesJSON) RawJSON() string {
	return r.raw
}

type OrganizationPolicyGetResponse struct {
	Policies OrganizationPolicies              `json:"policies,required"`
	JSON     organizationPolicyGetResponseJSON `json:"-"`
}

// organizationPolicyGetResponseJSON contains the JSON metadata for the struct
// [OrganizationPolicyGetResponse]
type organizationPolicyGetResponseJSON struct {
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationPolicyUpdateResponse = interface{}

type OrganizationPolicyGetParams struct {
	// organization_id is the ID of the organization to retrieve policies for
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r OrganizationPolicyGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationPolicyUpdateParams struct {
	// organization_id is the ID of the organization to update policies for
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
	// allowed_editor_ids is the list of editor IDs that are allowed to be used in the
	// organization
	AllowedEditorIDs param.Field[[]string] `json:"allowedEditorIds"`
	// allow_local_runners controls whether local runners are allowed to be used in the
	// organization
	AllowLocalRunners param.Field[bool] `json:"allowLocalRunners"`
	// default_editor_id is the default editor ID to be used when a user doesn't
	// specify one
	DefaultEditorID param.Field[string] `json:"defaultEditorId"`
	// default_environment_image is the default container image when none is defined in
	// repo
	DefaultEnvironmentImage param.Field[string] `json:"defaultEnvironmentImage"`
	// maximum_environments_per_user limits total environments (running or stopped) per
	// user
	MaximumEnvironmentsPerUser param.Field[string] `json:"maximumEnvironmentsPerUser"`
	// maximum_environment_timeout controls the maximum timeout allowed for
	// environments in seconds. 0 means no limit (never). Minimum duration is 30
	// minutes.
	MaximumEnvironmentTimeout param.Field[string] `json:"maximumEnvironmentTimeout" format:"regex"`
	// maximum_running_environments_per_user limits simultaneously running environments
	// per user
	MaximumRunningEnvironmentsPerUser param.Field[string] `json:"maximumRunningEnvironmentsPerUser"`
	// members_create_projects controls whether members can create projects
	MembersCreateProjects param.Field[bool] `json:"membersCreateProjects"`
	// members_require_projects controls whether environments can only be created from
	// projects by non-admin users
	MembersRequireProjects param.Field[bool] `json:"membersRequireProjects"`
	// port_sharing_disabled controls whether port sharing is disabled in the
	// organization
	PortSharingDisabled param.Field[bool] `json:"portSharingDisabled"`
}

func (r OrganizationPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
