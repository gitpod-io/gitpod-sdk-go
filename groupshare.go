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

// GroupShareService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupShareService] method instead.
type GroupShareService struct {
	Options []option.RequestOption
}

// NewGroupShareService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGroupShareService(opts ...option.RequestOption) (r *GroupShareService) {
	r = &GroupShareService{}
	r.Options = opts
	return
}

// Shares a resource directly with a principal (user or service account).
//
// Use this method to:
//
//   - Grant a user or service account direct access to a runner, project, or other
//     resource
//   - Share resources without creating and managing groups manually
//
// ### Examples
//
// - Share a runner with a user:
//
//	Grants admin access to a runner for a specific user.
//
//	```yaml
//	resourceType: RESOURCE_TYPE_RUNNER
//	resourceId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	principal: PRINCIPAL_USER
//	principalId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	role: RESOURCE_ROLE_RUNNER_ADMIN
//	```
//
// - Share a runner with a service account:
//
//	Grants user access to a runner for a service account.
//
//	```yaml
//	resourceType: RESOURCE_TYPE_RUNNER
//	resourceId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	principal: PRINCIPAL_SERVICE_ACCOUNT
//	principalId: "a1b2c3d4-5678-90ab-cdef-1234567890ab"
//	role: RESOURCE_ROLE_RUNNER_USER
//	```
//
// ### Authorization
//
// Requires admin role on the specific resource.
func (r *GroupShareService) New(ctx context.Context, body GroupShareNewParams, opts ...option.RequestOption) (res *GroupShareNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.GroupService/ShareResourceWithPrincipal"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes direct access for a principal (user or service account) from a resource.
//
// Use this method to:
//
// - Revoke a principal's direct access to a resource
// - Remove sharing without affecting group-based access
//
// ### Examples
//
// - Remove user access from a runner:
//
//	Revokes a user's direct access to a runner.
//
//	```yaml
//	resourceType: RESOURCE_TYPE_RUNNER
//	resourceId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	principal: PRINCIPAL_USER
//	principalId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	```
//
// ### Authorization
//
// Requires admin role on the specific resource.
func (r *GroupShareService) Delete(ctx context.Context, body GroupShareDeleteParams, opts ...option.RequestOption) (res *GroupShareDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.GroupService/UnshareResourceWithPrincipal"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GroupShareNewResponse = interface{}

type GroupShareDeleteResponse = interface{}

type GroupShareNewParams struct {
	// Type of principal to share with (user or service account)
	Principal param.Field[shared.Principal] `json:"principal"`
	// ID of the principal (user or service account) to share with
	PrincipalID param.Field[string] `json:"principalId" format:"uuid"`
	// ID of the resource to share
	ResourceID param.Field[string] `json:"resourceId" format:"uuid"`
	// Type of resource to share (runner, project, etc.)
	ResourceType param.Field[shared.ResourceType] `json:"resourceType"`
	// Role to grant the principal on the resource
	Role param.Field[shared.ResourceRole] `json:"role"`
}

func (r GroupShareNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GroupShareDeleteParams struct {
	// Type of principal to remove access from (user or service account)
	Principal param.Field[shared.Principal] `json:"principal"`
	// ID of the principal (user or service account) to remove access from
	PrincipalID param.Field[string] `json:"principalId" format:"uuid"`
	// ID of the resource to unshare
	ResourceID param.Field[string] `json:"resourceId" format:"uuid"`
	// Type of resource to unshare
	ResourceType param.Field[shared.ResourceType] `json:"resourceType"`
}

func (r GroupShareDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
