// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// OrganizationService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
	Members *OrganizationMemberService
	Invite  *OrganizationInviteService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.Members = NewOrganizationMemberService(opts...)
	r.Invite = NewOrganizationInviteService(opts...)
	return
}

// LeaveOrganization lets the passed user leave an Organization.
func (r *OrganizationService) Leave(ctx context.Context, params OrganizationLeaveParams, opts ...option.RequestOption) (res *OrganizationLeaveResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/LeaveOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// SetRole
func (r *OrganizationService) SetRole(ctx context.Context, params OrganizationSetRoleParams, opts ...option.RequestOption) (res *OrganizationSetRoleResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/SetRole"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type OrganizationLeaveResponse = interface{}

type OrganizationSetRoleResponse = interface{}

type OrganizationLeaveParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationLeaveParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	UserID                 param.Field[string]                                        `json:"userId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationLeaveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationLeaveParamsConnectProtocolVersion float64

const (
	OrganizationLeaveParamsConnectProtocolVersion1 OrganizationLeaveParamsConnectProtocolVersion = 1
)

func (r OrganizationLeaveParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationLeaveParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationSetRoleParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationSetRoleParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	OrganizationID         param.Field[string]                                          `json:"organizationId" format:"uuid"`
	Role                   param.Field[OrganizationSetRoleParamsRole]                   `json:"role"`
	UserID                 param.Field[string]                                          `json:"userId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationSetRoleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationSetRoleParamsConnectProtocolVersion float64

const (
	OrganizationSetRoleParamsConnectProtocolVersion1 OrganizationSetRoleParamsConnectProtocolVersion = 1
)

func (r OrganizationSetRoleParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationSetRoleParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationSetRoleParamsRole string

const (
	OrganizationSetRoleParamsRoleOrganizationRoleUnspecified OrganizationSetRoleParamsRole = "ORGANIZATION_ROLE_UNSPECIFIED"
	OrganizationSetRoleParamsRoleOrganizationRoleAdmin       OrganizationSetRoleParamsRole = "ORGANIZATION_ROLE_ADMIN"
	OrganizationSetRoleParamsRoleOrganizationRoleMember      OrganizationSetRoleParamsRole = "ORGANIZATION_ROLE_MEMBER"
)

func (r OrganizationSetRoleParamsRole) IsKnown() bool {
	switch r {
	case OrganizationSetRoleParamsRoleOrganizationRoleUnspecified, OrganizationSetRoleParamsRoleOrganizationRoleAdmin, OrganizationSetRoleParamsRoleOrganizationRoleMember:
		return true
	}
	return false
}
