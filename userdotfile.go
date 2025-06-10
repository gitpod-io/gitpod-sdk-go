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

// UserDotfileService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserDotfileService] method instead.
type UserDotfileService struct {
	Options []option.RequestOption
}

// NewUserDotfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserDotfileService(opts ...option.RequestOption) (r *UserDotfileService) {
	r = &UserDotfileService{}
	r.Options = opts
	return
}

// Gets the dotfiles for a user.
//
// Use this method to:
//
// - Retrieve user dotfiles
//
// ### Examples
//
// - Get dotfiles:
//
//	Retrieves the dotfiles for the current user.
//
//	```yaml
//	{}
//	```
func (r *UserDotfileService) Get(ctx context.Context, body UserDotfileGetParams, opts ...option.RequestOption) (res *UserDotfileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.UserService/GetDotfilesConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Sets the dotfiles configuration for a user.
//
// Use this method to:
//
// - Configure user dotfiles
// - Update dotfiles settings
//
// ### Examples
//
// - Set dotfiles configuration:
//
//	Sets the dotfiles configuration for the current user.
//
//	```yaml
//	{ "repository": "https://github.com/gitpod-io/dotfiles" }
//	```
//
// - Remove dotfiles:
//
//	Removes the dotfiles for the current user.
//
//	```yaml
//	{}
//	```
func (r *UserDotfileService) Set(ctx context.Context, body UserDotfileSetParams, opts ...option.RequestOption) (res *UserDotfileSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.UserService/SetDotfilesConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DotfilesConfiguration struct {
	// The URL of a dotfiles repository.
	Repository string                    `json:"repository" format:"uri"`
	JSON       dotfilesConfigurationJSON `json:"-"`
}

// dotfilesConfigurationJSON contains the JSON metadata for the struct
// [DotfilesConfiguration]
type dotfilesConfigurationJSON struct {
	Repository  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DotfilesConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dotfilesConfigurationJSON) RawJSON() string {
	return r.raw
}

type UserDotfileGetResponse struct {
	DotfilesConfiguration DotfilesConfiguration      `json:"dotfilesConfiguration,required"`
	JSON                  userDotfileGetResponseJSON `json:"-"`
}

// userDotfileGetResponseJSON contains the JSON metadata for the struct
// [UserDotfileGetResponse]
type userDotfileGetResponseJSON struct {
	DotfilesConfiguration apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *UserDotfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userDotfileGetResponseJSON) RawJSON() string {
	return r.raw
}

type UserDotfileSetResponse = interface{}

type UserDotfileGetParams struct {
	Empty param.Field[bool] `json:"empty"`
}

func (r UserDotfileGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserDotfileSetParams struct {
	Repository param.Field[string] `json:"repository" format:"uri"`
}

func (r UserDotfileSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
