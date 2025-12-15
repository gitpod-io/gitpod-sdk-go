// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// GroupMembershipService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupMembershipService] method instead.
type GroupMembershipService struct {
	Options []option.RequestOption
}

// NewGroupMembershipService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGroupMembershipService(opts ...option.RequestOption) (r *GroupMembershipService) {
	r = &GroupMembershipService{}
	r.Options = opts
	return
}

// Creates a membership for a user in a group.
//
// Use this method to:
//
// - Add users to groups
// - Grant group-based permissions to users
//
// ### Examples
//
// - Add a user to a group:
//
//	Creates a membership for a user in a group.
//
//	```yaml
//	groupId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	subject:
//	  id: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	  principal: PRINCIPAL_USER
//	```
//
// ### Authorization
//
// Requires `org:admin` permission on the organization or `group:admin` permission
// on the specific group.
func (r *GroupMembershipService) New(ctx context.Context, body GroupMembershipNewParams, opts ...option.RequestOption) (res *GroupMembershipNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.GroupService/CreateMembership"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all memberships of a group.
//
// Use this method to:
//
// - View all members of a group
// - Audit group membership
//
// ### Examples
//
// - List group members:
//
//	Shows all members of a specific group.
//
//	```yaml
//	groupId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	pagination:
//	  pageSize: 20
//	```
//
// ### Authorization
//
// All organization members can view group membership (transparency model).
func (r *GroupMembershipService) List(ctx context.Context, params GroupMembershipListParams, opts ...option.RequestOption) (res *pagination.MembersPage[GroupMembership], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.GroupService/ListMemberships"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all memberships of a group.
//
// Use this method to:
//
// - View all members of a group
// - Audit group membership
//
// ### Examples
//
// - List group members:
//
//	Shows all members of a specific group.
//
//	```yaml
//	groupId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	pagination:
//	  pageSize: 20
//	```
//
// ### Authorization
//
// All organization members can view group membership (transparency model).
func (r *GroupMembershipService) ListAutoPaging(ctx context.Context, params GroupMembershipListParams, opts ...option.RequestOption) *pagination.MembersPageAutoPager[GroupMembership] {
	return pagination.NewMembersPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a membership for a user in a group.
//
// Use this method to:
//
// - Remove users from groups
// - Revoke group-based permissions
//
// ### Examples
//
// - Remove a user from a group:
//
//	Deletes a membership by its ID.
//
//	```yaml
//	membershipId: "a1b2c3d4-5678-90ab-cdef-1234567890ab"
//	```
//
// ### Authorization
//
// Requires `org:admin` permission on the organization or `group:admin` permission
// on the specific group.
func (r *GroupMembershipService) Delete(ctx context.Context, body GroupMembershipDeleteParams, opts ...option.RequestOption) (res *GroupMembershipDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.GroupService/DeleteMembership"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GroupMembership represents a subject's membership in a group
type GroupMembership struct {
	// Unique identifier for the group membership
	ID string `json:"id" format:"uuid"`
	// Subject's avatar URL
	AvatarURL string `json:"avatarUrl"`
	// Group identifier
	GroupID string `json:"groupId" format:"uuid"`
	// Subject's display name
	Name string `json:"name"`
	// Subject (user, runner, environment, service account, etc.)
	Subject shared.Subject      `json:"subject"`
	JSON    groupMembershipJSON `json:"-"`
}

// groupMembershipJSON contains the JSON metadata for the struct [GroupMembership]
type groupMembershipJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	GroupID     apijson.Field
	Name        apijson.Field
	Subject     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GroupMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupMembershipJSON) RawJSON() string {
	return r.raw
}

type GroupMembershipNewResponse struct {
	// GroupMembership represents a subject's membership in a group
	Member GroupMembership                `json:"member"`
	JSON   groupMembershipNewResponseJSON `json:"-"`
}

// groupMembershipNewResponseJSON contains the JSON metadata for the struct
// [GroupMembershipNewResponse]
type groupMembershipNewResponseJSON struct {
	Member      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GroupMembershipNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupMembershipNewResponseJSON) RawJSON() string {
	return r.raw
}

type GroupMembershipDeleteResponse = interface{}

type GroupMembershipNewParams struct {
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
	// Subject to add to the group
	Subject param.Field[shared.SubjectParam] `json:"subject"`
}

func (r GroupMembershipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GroupMembershipListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	GroupID  param.Field[string] `json:"groupId" format:"uuid"`
	// pagination contains the pagination options for listing memberships
	Pagination param.Field[GroupMembershipListParamsPagination] `json:"pagination"`
}

func (r GroupMembershipListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [GroupMembershipListParams]'s query parameters as
// `url.Values`.
func (r GroupMembershipListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// pagination contains the pagination options for listing memberships
type GroupMembershipListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r GroupMembershipListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GroupMembershipDeleteParams struct {
	// The membership to delete
	MembershipID param.Field[string] `json:"membershipId" format:"uuid"`
}

func (r GroupMembershipDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
