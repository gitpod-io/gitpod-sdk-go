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

// ListRunnerPolicies lists runner policies.
func (r *RunnerPolicyService) List(ctx context.Context, body RunnerPolicyListParams, opts ...option.RequestOption) (res *RunnerPolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/ListRunnerPolicies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RunnerPolicyListResponse struct {
	Pagination RunnerPolicyListResponsePagination `json:"pagination"`
	Policies   []RunnerPolicyListResponsePolicy   `json:"policies"`
	JSON       runnerPolicyListResponseJSON       `json:"-"`
}

// runnerPolicyListResponseJSON contains the JSON metadata for the struct
// [RunnerPolicyListResponse]
type runnerPolicyListResponseJSON struct {
	Pagination  apijson.Field
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerPolicyListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                 `json:"nextToken"`
	JSON      runnerPolicyListResponsePaginationJSON `json:"-"`
}

// runnerPolicyListResponsePaginationJSON contains the JSON metadata for the struct
// [RunnerPolicyListResponsePagination]
type runnerPolicyListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerPolicyListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerPolicyListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerPolicyListResponsePolicy struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role RunnerPolicyListResponsePoliciesRole `json:"role"`
	JSON runnerPolicyListResponsePolicyJSON   `json:"-"`
}

// runnerPolicyListResponsePolicyJSON contains the JSON metadata for the struct
// [RunnerPolicyListResponsePolicy]
type runnerPolicyListResponsePolicyJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerPolicyListResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerPolicyListResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// role is the role assigned to the group
type RunnerPolicyListResponsePoliciesRole string

const (
	RunnerPolicyListResponsePoliciesRoleRunnerRoleUnspecified RunnerPolicyListResponsePoliciesRole = "RUNNER_ROLE_UNSPECIFIED"
	RunnerPolicyListResponsePoliciesRoleRunnerRoleAdmin       RunnerPolicyListResponsePoliciesRole = "RUNNER_ROLE_ADMIN"
	RunnerPolicyListResponsePoliciesRoleRunnerRoleUser        RunnerPolicyListResponsePoliciesRole = "RUNNER_ROLE_USER"
)

func (r RunnerPolicyListResponsePoliciesRole) IsKnown() bool {
	switch r {
	case RunnerPolicyListResponsePoliciesRoleRunnerRoleUnspecified, RunnerPolicyListResponsePoliciesRoleRunnerRoleAdmin, RunnerPolicyListResponsePoliciesRoleRunnerRoleUser:
		return true
	}
	return false
}

type RunnerPolicyListParams struct {
	// pagination contains the pagination options for listing project policies
	Pagination param.Field[RunnerPolicyListParamsPagination] `json:"pagination"`
	// runner_id specifies the project identifier
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerPolicyListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing project policies
type RunnerPolicyListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	//
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerPolicyListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
