// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
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

// CreateRunnerPolicy creates a new runner policy.
func (r *RunnerPolicyService) New(ctx context.Context, params RunnerPolicyNewParams, opts ...option.RequestOption) (res *RunnerPolicyNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/CreateRunnerPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateRunnerPolicy an existing runner policy.
func (r *RunnerPolicyService) Update(ctx context.Context, params RunnerPolicyUpdateParams, opts ...option.RequestOption) (res *RunnerPolicyUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/UpdateRunnerPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListRunnerPolicies lists runner policies.
func (r *RunnerPolicyService) List(ctx context.Context, params RunnerPolicyListParams, opts ...option.RequestOption) (res *RunnerPolicyListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/ListRunnerPolicies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// DeleteRunnerPolicy deletes a runner policy.
func (r *RunnerPolicyService) Delete(ctx context.Context, params RunnerPolicyDeleteParams, opts ...option.RequestOption) (res *RunnerPolicyDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerService/DeleteRunnerPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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

type RunnerPolicyDeleteResponse = interface{}

type RunnerPolicyNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerPolicyNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// group_id specifies the group_id identifier
	GroupID param.Field[string]                    `json:"groupId" format:"uuid"`
	Role    param.Field[RunnerPolicyNewParamsRole] `json:"role"`
	// runner_id specifies the project identifier
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerPolicyNewParamsConnectProtocolVersion float64

const (
	RunnerPolicyNewParamsConnectProtocolVersion1 RunnerPolicyNewParamsConnectProtocolVersion = 1
)

func (r RunnerPolicyNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerPolicyNewParamsConnectProtocolVersion1:
		return true
	}
	return false
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
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerPolicyUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// group_id specifies the group_id identifier
	GroupID param.Field[string]                       `json:"groupId" format:"uuid"`
	Role    param.Field[RunnerPolicyUpdateParamsRole] `json:"role"`
	// runner_id specifies the project identifier
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerPolicyUpdateParamsConnectProtocolVersion float64

const (
	RunnerPolicyUpdateParamsConnectProtocolVersion1 RunnerPolicyUpdateParamsConnectProtocolVersion = 1
)

func (r RunnerPolicyUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerPolicyUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
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
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[RunnerPolicyListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerPolicyListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[RunnerPolicyListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[RunnerPolicyListParamsConnect] `query:"connect"`
	Message param.Field[string]                        `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [RunnerPolicyListParams]'s query parameters as `url.Values`.
func (r RunnerPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type RunnerPolicyListParamsEncoding string

const (
	RunnerPolicyListParamsEncodingProto RunnerPolicyListParamsEncoding = "proto"
	RunnerPolicyListParamsEncodingJson  RunnerPolicyListParamsEncoding = "json"
)

func (r RunnerPolicyListParamsEncoding) IsKnown() bool {
	switch r {
	case RunnerPolicyListParamsEncodingProto, RunnerPolicyListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerPolicyListParamsConnectProtocolVersion float64

const (
	RunnerPolicyListParamsConnectProtocolVersion1 RunnerPolicyListParamsConnectProtocolVersion = 1
)

func (r RunnerPolicyListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerPolicyListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type RunnerPolicyListParamsCompression string

const (
	RunnerPolicyListParamsCompressionIdentity RunnerPolicyListParamsCompression = "identity"
	RunnerPolicyListParamsCompressionGzip     RunnerPolicyListParamsCompression = "gzip"
	RunnerPolicyListParamsCompressionBr       RunnerPolicyListParamsCompression = "br"
)

func (r RunnerPolicyListParamsCompression) IsKnown() bool {
	switch r {
	case RunnerPolicyListParamsCompressionIdentity, RunnerPolicyListParamsCompressionGzip, RunnerPolicyListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerPolicyListParamsConnect string

const (
	RunnerPolicyListParamsConnectV1 RunnerPolicyListParamsConnect = "v1"
)

func (r RunnerPolicyListParamsConnect) IsKnown() bool {
	switch r {
	case RunnerPolicyListParamsConnectV1:
		return true
	}
	return false
}

type RunnerPolicyDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerPolicyDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// group_id specifies the group_id identifier
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
	// runner_id specifies the project identifier
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerPolicyDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerPolicyDeleteParamsConnectProtocolVersion float64

const (
	RunnerPolicyDeleteParamsConnectProtocolVersion1 RunnerPolicyDeleteParamsConnectProtocolVersion = 1
)

func (r RunnerPolicyDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerPolicyDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}
