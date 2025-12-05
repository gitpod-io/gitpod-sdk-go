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
)

// RunnerPolicyService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerPolicyService] method instead.
type RunnerPolicyService struct {
	Options []option.RequestOption
}

// NewRunnerPolicyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRunnerPolicyService(opts ...option.RequestOption) (r *RunnerPolicyService) {
	r = &RunnerPolicyService{}
	r.Options = opts
	return
}

// Creates a new policy for a runner.
//
// Use this method to:
//
// - Set up access controls
// - Define group permissions
// - Configure role-based access
//
// ### Examples
//
// - Create admin policy:
//
//	Grants admin access to a group.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	groupId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	role: RUNNER_ROLE_ADMIN
//	```
func (r *RunnerPolicyService) New(ctx context.Context, body RunnerPolicyNewParams, opts ...option.RequestOption) (res *RunnerPolicyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerService/CreateRunnerPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing runner policy.
//
// Use this method to:
//
// - Modify access levels
// - Change group roles
// - Update permissions
//
// ### Examples
//
// - Update policy role:
//
//	Changes a group's access level.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	groupId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	role: RUNNER_ROLE_USER
//	```
func (r *RunnerPolicyService) Update(ctx context.Context, body RunnerPolicyUpdateParams, opts ...option.RequestOption) (res *RunnerPolicyUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerService/UpdateRunnerPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists policies for a runner.
//
// Use this method to:
//
// - View access controls
// - Check policy configurations
// - Audit permissions
//
// ### Examples
//
// - List policies:
//
//	Shows all policies for a runner.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	pagination:
//	  pageSize: 20
//	```
func (r *RunnerPolicyService) List(ctx context.Context, params RunnerPolicyListParams, opts ...option.RequestOption) (res *pagination.PoliciesPage[RunnerPolicy], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.RunnerService/ListRunnerPolicies"
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

// Lists policies for a runner.
//
// Use this method to:
//
// - View access controls
// - Check policy configurations
// - Audit permissions
//
// ### Examples
//
// - List policies:
//
//	Shows all policies for a runner.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	pagination:
//	  pageSize: 20
//	```
func (r *RunnerPolicyService) ListAutoPaging(ctx context.Context, params RunnerPolicyListParams, opts ...option.RequestOption) *pagination.PoliciesPageAutoPager[RunnerPolicy] {
	return pagination.NewPoliciesPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a runner policy.
//
// Use this method to:
//
// - Remove access controls
// - Revoke permissions
// - Clean up policies
//
// ### Examples
//
// - Delete policy:
//
//	Removes a group's access policy.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	groupId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	```
func (r *RunnerPolicyService) Delete(ctx context.Context, body RunnerPolicyDeleteParams, opts ...option.RequestOption) (res *RunnerPolicyDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerService/DeleteRunnerPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RunnerPolicy struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role RunnerRole       `json:"role"`
	JSON runnerPolicyJSON `json:"-"`
}

// runnerPolicyJSON contains the JSON metadata for the struct [RunnerPolicy]
type runnerPolicyJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerPolicyJSON) RawJSON() string {
	return r.raw
}

type RunnerRole string

const (
	RunnerRoleUnspecified RunnerRole = "RUNNER_ROLE_UNSPECIFIED"
	RunnerRoleAdmin       RunnerRole = "RUNNER_ROLE_ADMIN"
	RunnerRoleUser        RunnerRole = "RUNNER_ROLE_USER"
)

func (r RunnerRole) IsKnown() bool {
	switch r {
	case RunnerRoleUnspecified, RunnerRoleAdmin, RunnerRoleUser:
		return true
	}
	return false
}

type RunnerPolicyNewResponse struct {
	Policy RunnerPolicy                `json:"policy,required"`
	JSON   runnerPolicyNewResponseJSON `json:"-"`
}

// runnerPolicyNewResponseJSON contains the JSON metadata for the struct
// [RunnerPolicyNewResponse]
type runnerPolicyNewResponseJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerPolicyUpdateResponse struct {
	Policy RunnerPolicy                   `json:"policy,required"`
	JSON   runnerPolicyUpdateResponseJSON `json:"-"`
}

// runnerPolicyUpdateResponseJSON contains the JSON metadata for the struct
// [RunnerPolicyUpdateResponse]
type runnerPolicyUpdateResponseJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerPolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerPolicyDeleteResponse = interface{}

type RunnerPolicyNewParams struct {
	// group_id specifies the group_id identifier
	GroupID param.Field[string]     `json:"groupId" format:"uuid"`
	Role    param.Field[RunnerRole] `json:"role"`
	// runner_id specifies the project identifier
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerPolicyUpdateParams struct {
	// group_id specifies the group_id identifier
	GroupID param.Field[string]     `json:"groupId" format:"uuid"`
	Role    param.Field[RunnerRole] `json:"role"`
	// runner_id specifies the project identifier
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerPolicyListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// pagination contains the pagination options for listing project policies
	Pagination param.Field[RunnerPolicyListParamsPagination] `json:"pagination"`
	// runner_id specifies the project identifier
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerPolicyListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [RunnerPolicyListParams]'s query parameters as `url.Values`.
func (r RunnerPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// pagination contains the pagination options for listing project policies
type RunnerPolicyListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerPolicyListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerPolicyDeleteParams struct {
	// group_id specifies the group_id identifier
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
	// runner_id specifies the project identifier
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerPolicyDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
