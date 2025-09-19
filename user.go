// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// UserService contains methods and other services that help with interacting with
// the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options  []option.RequestOption
	Dotfiles *UserDotfileService
	Pats     *UserPatService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.Dotfiles = NewUserDotfileService(opts...)
	r.Pats = NewUserPatService(opts...)
	return
}

// Gets information about the currently authenticated user.
//
// Use this method to:
//
// - Get user profile information
// - Check authentication status
// - Retrieve user settings
// - Verify account details
//
// ### Examples
//
// - Get current user:
//
//	Retrieves details about the authenticated user.
//
//	```yaml
//	{}
//	```
func (r *UserService) GetAuthenticatedUser(ctx context.Context, body UserGetAuthenticatedUserParams, opts ...option.RequestOption) (res *UserGetAuthenticatedUserResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.UserService/GetAuthenticatedUser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Sets whether a user account is suspended.
//
// Use this method to:
//
// - Suspend problematic users
// - Reactivate suspended accounts
// - Manage user access
//
// ### Examples
//
// - Suspend user:
//
//	Suspends a user account.
//
//	```yaml
//	userId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	suspended: true
//	```
//
// - Reactivate user:
//
//	Removes suspension from a user account.
//
//	```yaml
//	userId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	suspended: false
//	```
func (r *UserService) SetSuspended(ctx context.Context, body UserSetSuspendedParams, opts ...option.RequestOption) (res *UserSetSuspendedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.UserService/SetSuspended"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type User struct {
	// id is a UUID of the user
	ID string `json:"id,required" format:"uuid"`
	// avatar_url is a link to the user avatar
	AvatarURL string `json:"avatarUrl"`
	// created_at is the creation time
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// name is the full name of the user
	Name string `json:"name"`
	// organization_id is the id of the organization this account is owned by.
	//
	// +optional if not set, this account is owned by the installation.
	OrganizationID string `json:"organizationId" format:"uuid"`
	// status is the status the user is in
	Status shared.UserStatus `json:"status"`
	JSON   userJSON          `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID             apijson.Field
	AvatarURL      apijson.Field
	CreatedAt      apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

type UserGetAuthenticatedUserResponse struct {
	User User                                 `json:"user,required"`
	JSON userGetAuthenticatedUserResponseJSON `json:"-"`
}

// userGetAuthenticatedUserResponseJSON contains the JSON metadata for the struct
// [UserGetAuthenticatedUserResponse]
type userGetAuthenticatedUserResponseJSON struct {
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetAuthenticatedUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetAuthenticatedUserResponseJSON) RawJSON() string {
	return r.raw
}

type UserSetSuspendedResponse = interface{}

type UserGetAuthenticatedUserParams struct {
	Empty param.Field[bool] `json:"empty"`
}

func (r UserGetAuthenticatedUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetSuspendedParams struct {
	Suspended param.Field[bool]   `json:"suspended"`
	UserID    param.Field[string] `json:"userId" format:"uuid"`
}

func (r UserSetSuspendedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
