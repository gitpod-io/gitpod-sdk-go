// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gitpod-io/flex-sdk-go/internal/apijson"
	"github.com/gitpod-io/flex-sdk-go/internal/apiquery"
	"github.com/gitpod-io/flex-sdk-go/internal/param"
	"github.com/gitpod-io/flex-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/flex-sdk-go/option"
	"github.com/gitpod-io/flex-sdk-go/packages/pagination"
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

// CreateRunnerPolicy creates a new runner policy.
func (r *RunnerPolicyService) New(ctx context.Context, body RunnerPolicyNewParams, opts ...option.RequestOption) (res *RunnerPolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CreateRunnerPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateRunnerPolicy an existing runner policy.
func (r *RunnerPolicyService) Update(ctx context.Context, body RunnerPolicyUpdateParams, opts ...option.RequestOption) (res *RunnerPolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/UpdateRunnerPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListRunnerPolicies lists runner policies.
func (r *RunnerPolicyService) List(ctx context.Context, params RunnerPolicyListParams, opts ...option.RequestOption) (res *pagination.PoliciesPage[RunnerPolicyListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// ListRunnerPolicies lists runner policies.
func (r *RunnerPolicyService) ListAutoPaging(ctx context.Context, params RunnerPolicyListParams, opts ...option.RequestOption) *pagination.PoliciesPageAutoPager[RunnerPolicyListResponse] {
	return pagination.NewPoliciesPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteRunnerPolicy deletes a runner policy.
func (r *RunnerPolicyService) Delete(ctx context.Context, body RunnerPolicyDeleteParams, opts ...option.RequestOption) (res *RunnerPolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/DeleteRunnerPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RunnerPolicyNewResponse struct {
	Policy RunnerPolicyNewResponsePolicy `json:"policy"`
	JSON   runnerPolicyNewResponseJSON   `json:"-"`
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

type RunnerPolicyNewResponsePolicy struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role RunnerPolicyNewResponsePolicyRole `json:"role"`
	JSON runnerPolicyNewResponsePolicyJSON `json:"-"`
}

// runnerPolicyNewResponsePolicyJSON contains the JSON metadata for the struct
// [RunnerPolicyNewResponsePolicy]
type runnerPolicyNewResponsePolicyJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerPolicyNewResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerPolicyNewResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// role is the role assigned to the group
type RunnerPolicyNewResponsePolicyRole string

const (
	RunnerPolicyNewResponsePolicyRoleRunnerRoleUnspecified RunnerPolicyNewResponsePolicyRole = "RUNNER_ROLE_UNSPECIFIED"
	RunnerPolicyNewResponsePolicyRoleRunnerRoleAdmin       RunnerPolicyNewResponsePolicyRole = "RUNNER_ROLE_ADMIN"
	RunnerPolicyNewResponsePolicyRoleRunnerRoleUser        RunnerPolicyNewResponsePolicyRole = "RUNNER_ROLE_USER"
)

func (r RunnerPolicyNewResponsePolicyRole) IsKnown() bool {
	switch r {
	case RunnerPolicyNewResponsePolicyRoleRunnerRoleUnspecified, RunnerPolicyNewResponsePolicyRoleRunnerRoleAdmin, RunnerPolicyNewResponsePolicyRoleRunnerRoleUser:
		return true
	}
	return false
}

type RunnerPolicyUpdateResponse struct {
	Policy RunnerPolicyUpdateResponsePolicy `json:"policy"`
	JSON   runnerPolicyUpdateResponseJSON   `json:"-"`
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

type RunnerPolicyUpdateResponsePolicy struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role RunnerPolicyUpdateResponsePolicyRole `json:"role"`
	JSON runnerPolicyUpdateResponsePolicyJSON `json:"-"`
}

// runnerPolicyUpdateResponsePolicyJSON contains the JSON metadata for the struct
// [RunnerPolicyUpdateResponsePolicy]
type runnerPolicyUpdateResponsePolicyJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerPolicyUpdateResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerPolicyUpdateResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// role is the role assigned to the group
type RunnerPolicyUpdateResponsePolicyRole string

const (
	RunnerPolicyUpdateResponsePolicyRoleRunnerRoleUnspecified RunnerPolicyUpdateResponsePolicyRole = "RUNNER_ROLE_UNSPECIFIED"
	RunnerPolicyUpdateResponsePolicyRoleRunnerRoleAdmin       RunnerPolicyUpdateResponsePolicyRole = "RUNNER_ROLE_ADMIN"
	RunnerPolicyUpdateResponsePolicyRoleRunnerRoleUser        RunnerPolicyUpdateResponsePolicyRole = "RUNNER_ROLE_USER"
)

func (r RunnerPolicyUpdateResponsePolicyRole) IsKnown() bool {
	switch r {
	case RunnerPolicyUpdateResponsePolicyRoleRunnerRoleUnspecified, RunnerPolicyUpdateResponsePolicyRoleRunnerRoleAdmin, RunnerPolicyUpdateResponsePolicyRoleRunnerRoleUser:
		return true
	}
	return false
}

type RunnerPolicyListResponse struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role RunnerPolicyListResponseRole `json:"role"`
	JSON runnerPolicyListResponseJSON `json:"-"`
}

// runnerPolicyListResponseJSON contains the JSON metadata for the struct
// [RunnerPolicyListResponse]
type runnerPolicyListResponseJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

// role is the role assigned to the group
type RunnerPolicyListResponseRole string

const (
	RunnerPolicyListResponseRoleRunnerRoleUnspecified RunnerPolicyListResponseRole = "RUNNER_ROLE_UNSPECIFIED"
	RunnerPolicyListResponseRoleRunnerRoleAdmin       RunnerPolicyListResponseRole = "RUNNER_ROLE_ADMIN"
	RunnerPolicyListResponseRoleRunnerRoleUser        RunnerPolicyListResponseRole = "RUNNER_ROLE_USER"
)

func (r RunnerPolicyListResponseRole) IsKnown() bool {
	switch r {
	case RunnerPolicyListResponseRoleRunnerRoleUnspecified, RunnerPolicyListResponseRoleRunnerRoleAdmin, RunnerPolicyListResponseRoleRunnerRoleUser:
		return true
	}
	return false
}

type RunnerPolicyDeleteResponse = interface{}

type RunnerPolicyNewParams struct {
	// group_id specifies the group_id identifier
	GroupID param.Field[string]                    `json:"groupId" format:"uuid"`
	Role    param.Field[RunnerPolicyNewParamsRole] `json:"role"`
	// runner_id specifies the project identifier
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerPolicyNewParamsRole string

const (
	RunnerPolicyNewParamsRoleRunnerRoleUnspecified RunnerPolicyNewParamsRole = "RUNNER_ROLE_UNSPECIFIED"
	RunnerPolicyNewParamsRoleRunnerRoleAdmin       RunnerPolicyNewParamsRole = "RUNNER_ROLE_ADMIN"
	RunnerPolicyNewParamsRoleRunnerRoleUser        RunnerPolicyNewParamsRole = "RUNNER_ROLE_USER"
)

func (r RunnerPolicyNewParamsRole) IsKnown() bool {
	switch r {
	case RunnerPolicyNewParamsRoleRunnerRoleUnspecified, RunnerPolicyNewParamsRoleRunnerRoleAdmin, RunnerPolicyNewParamsRoleRunnerRoleUser:
		return true
	}
	return false
}

type RunnerPolicyUpdateParams struct {
	// group_id specifies the group_id identifier
	GroupID param.Field[string]                       `json:"groupId" format:"uuid"`
	Role    param.Field[RunnerPolicyUpdateParamsRole] `json:"role"`
	// runner_id specifies the project identifier
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerPolicyUpdateParamsRole string

const (
	RunnerPolicyUpdateParamsRoleRunnerRoleUnspecified RunnerPolicyUpdateParamsRole = "RUNNER_ROLE_UNSPECIFIED"
	RunnerPolicyUpdateParamsRoleRunnerRoleAdmin       RunnerPolicyUpdateParamsRole = "RUNNER_ROLE_ADMIN"
	RunnerPolicyUpdateParamsRoleRunnerRoleUser        RunnerPolicyUpdateParamsRole = "RUNNER_ROLE_USER"
)

func (r RunnerPolicyUpdateParamsRole) IsKnown() bool {
	switch r {
	case RunnerPolicyUpdateParamsRoleRunnerRoleUnspecified, RunnerPolicyUpdateParamsRoleRunnerRoleAdmin, RunnerPolicyUpdateParamsRoleRunnerRoleUser:
		return true
	}
	return false
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
