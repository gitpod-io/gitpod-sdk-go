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

// ProjectPolicyService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectPolicyService] method instead.
type ProjectPolicyService struct {
	Options []option.RequestOption
}

// NewProjectPolicyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectPolicyService(opts ...option.RequestOption) (r *ProjectPolicyService) {
	r = &ProjectPolicyService{}
	r.Options = opts
	return
}

// CreateProjectPolicy creates a Project Policy.
func (r *ProjectPolicyService) New(ctx context.Context, params ProjectPolicyNewParams, opts ...option.RequestOption) (res *ProjectPolicyNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/CreateProjectPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateProjectPolicy updates a Project Policy.
func (r *ProjectPolicyService) Update(ctx context.Context, params ProjectPolicyUpdateParams, opts ...option.RequestOption) (res *ProjectPolicyUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/UpdateProjectPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListProjectPolicies lists policies for a project.
func (r *ProjectPolicyService) List(ctx context.Context, params ProjectPolicyListParams, opts ...option.RequestOption) (res *ProjectPolicyListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/ListProjectPolicies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// DeleteProjectPolicy deletes a Project Policy.
func (r *ProjectPolicyService) Delete(ctx context.Context, params ProjectPolicyDeleteParams, opts ...option.RequestOption) (res *ProjectPolicyDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/DeleteProjectPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ProjectPolicyNewResponse struct {
	Policy ProjectPolicyNewResponsePolicy `json:"policy"`
	JSON   projectPolicyNewResponseJSON   `json:"-"`
}

// projectPolicyNewResponseJSON contains the JSON metadata for the struct
// [ProjectPolicyNewResponse]
type projectPolicyNewResponseJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectPolicyNewResponsePolicy struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role ProjectPolicyNewResponsePolicyRole `json:"role"`
	JSON projectPolicyNewResponsePolicyJSON `json:"-"`
}

// projectPolicyNewResponsePolicyJSON contains the JSON metadata for the struct
// [ProjectPolicyNewResponsePolicy]
type projectPolicyNewResponsePolicyJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyNewResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyNewResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// role is the role assigned to the group
type ProjectPolicyNewResponsePolicyRole string

const (
	ProjectPolicyNewResponsePolicyRoleProjectRoleUnspecified ProjectPolicyNewResponsePolicyRole = "PROJECT_ROLE_UNSPECIFIED"
	ProjectPolicyNewResponsePolicyRoleProjectRoleAdmin       ProjectPolicyNewResponsePolicyRole = "PROJECT_ROLE_ADMIN"
	ProjectPolicyNewResponsePolicyRoleProjectRoleUser        ProjectPolicyNewResponsePolicyRole = "PROJECT_ROLE_USER"
)

func (r ProjectPolicyNewResponsePolicyRole) IsKnown() bool {
	switch r {
	case ProjectPolicyNewResponsePolicyRoleProjectRoleUnspecified, ProjectPolicyNewResponsePolicyRoleProjectRoleAdmin, ProjectPolicyNewResponsePolicyRoleProjectRoleUser:
		return true
	}
	return false
}

type ProjectPolicyUpdateResponse struct {
	Policy ProjectPolicyUpdateResponsePolicy `json:"policy"`
	JSON   projectPolicyUpdateResponseJSON   `json:"-"`
}

// projectPolicyUpdateResponseJSON contains the JSON metadata for the struct
// [ProjectPolicyUpdateResponse]
type projectPolicyUpdateResponseJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectPolicyUpdateResponsePolicy struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role ProjectPolicyUpdateResponsePolicyRole `json:"role"`
	JSON projectPolicyUpdateResponsePolicyJSON `json:"-"`
}

// projectPolicyUpdateResponsePolicyJSON contains the JSON metadata for the struct
// [ProjectPolicyUpdateResponsePolicy]
type projectPolicyUpdateResponsePolicyJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyUpdateResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyUpdateResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// role is the role assigned to the group
type ProjectPolicyUpdateResponsePolicyRole string

const (
	ProjectPolicyUpdateResponsePolicyRoleProjectRoleUnspecified ProjectPolicyUpdateResponsePolicyRole = "PROJECT_ROLE_UNSPECIFIED"
	ProjectPolicyUpdateResponsePolicyRoleProjectRoleAdmin       ProjectPolicyUpdateResponsePolicyRole = "PROJECT_ROLE_ADMIN"
	ProjectPolicyUpdateResponsePolicyRoleProjectRoleUser        ProjectPolicyUpdateResponsePolicyRole = "PROJECT_ROLE_USER"
)

func (r ProjectPolicyUpdateResponsePolicyRole) IsKnown() bool {
	switch r {
	case ProjectPolicyUpdateResponsePolicyRoleProjectRoleUnspecified, ProjectPolicyUpdateResponsePolicyRoleProjectRoleAdmin, ProjectPolicyUpdateResponsePolicyRoleProjectRoleUser:
		return true
	}
	return false
}

type ProjectPolicyListResponse struct {
	Pagination ProjectPolicyListResponsePagination `json:"pagination"`
	Policies   []ProjectPolicyListResponsePolicy   `json:"policies"`
	JSON       projectPolicyListResponseJSON       `json:"-"`
}

// projectPolicyListResponseJSON contains the JSON metadata for the struct
// [ProjectPolicyListResponse]
type projectPolicyListResponseJSON struct {
	Pagination  apijson.Field
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectPolicyListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                  `json:"nextToken"`
	JSON      projectPolicyListResponsePaginationJSON `json:"-"`
}

// projectPolicyListResponsePaginationJSON contains the JSON metadata for the
// struct [ProjectPolicyListResponsePagination]
type projectPolicyListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ProjectPolicyListResponsePolicy struct {
	GroupID string `json:"groupId" format:"uuid"`
	// role is the role assigned to the group
	Role ProjectPolicyListResponsePoliciesRole `json:"role"`
	JSON projectPolicyListResponsePolicyJSON   `json:"-"`
}

// projectPolicyListResponsePolicyJSON contains the JSON metadata for the struct
// [ProjectPolicyListResponsePolicy]
type projectPolicyListResponsePolicyJSON struct {
	GroupID     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPolicyListResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPolicyListResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// role is the role assigned to the group
type ProjectPolicyListResponsePoliciesRole string

const (
	ProjectPolicyListResponsePoliciesRoleProjectRoleUnspecified ProjectPolicyListResponsePoliciesRole = "PROJECT_ROLE_UNSPECIFIED"
	ProjectPolicyListResponsePoliciesRoleProjectRoleAdmin       ProjectPolicyListResponsePoliciesRole = "PROJECT_ROLE_ADMIN"
	ProjectPolicyListResponsePoliciesRoleProjectRoleUser        ProjectPolicyListResponsePoliciesRole = "PROJECT_ROLE_USER"
)

func (r ProjectPolicyListResponsePoliciesRole) IsKnown() bool {
	switch r {
	case ProjectPolicyListResponsePoliciesRoleProjectRoleUnspecified, ProjectPolicyListResponsePoliciesRoleProjectRoleAdmin, ProjectPolicyListResponsePoliciesRoleProjectRoleUser:
		return true
	}
	return false
}

type ProjectPolicyDeleteResponse = interface{}

type ProjectPolicyNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ProjectPolicyNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// group_id specifies the group_id identifier
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
	// project_id specifies the project identifier
	ProjectID param.Field[string]                     `json:"projectId" format:"uuid"`
	Role      param.Field[ProjectPolicyNewParamsRole] `json:"role"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ProjectPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ProjectPolicyNewParamsConnectProtocolVersion float64

const (
	ProjectPolicyNewParamsConnectProtocolVersion1 ProjectPolicyNewParamsConnectProtocolVersion = 1
)

func (r ProjectPolicyNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ProjectPolicyNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ProjectPolicyNewParamsRole string

const (
	ProjectPolicyNewParamsRoleProjectRoleUnspecified ProjectPolicyNewParamsRole = "PROJECT_ROLE_UNSPECIFIED"
	ProjectPolicyNewParamsRoleProjectRoleAdmin       ProjectPolicyNewParamsRole = "PROJECT_ROLE_ADMIN"
	ProjectPolicyNewParamsRoleProjectRoleUser        ProjectPolicyNewParamsRole = "PROJECT_ROLE_USER"
)

func (r ProjectPolicyNewParamsRole) IsKnown() bool {
	switch r {
	case ProjectPolicyNewParamsRoleProjectRoleUnspecified, ProjectPolicyNewParamsRoleProjectRoleAdmin, ProjectPolicyNewParamsRoleProjectRoleUser:
		return true
	}
	return false
}

type ProjectPolicyUpdateParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ProjectPolicyUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// group_id specifies the group_id identifier
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
	// project_id specifies the project identifier
	ProjectID param.Field[string]                        `json:"projectId" format:"uuid"`
	Role      param.Field[ProjectPolicyUpdateParamsRole] `json:"role"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ProjectPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ProjectPolicyUpdateParamsConnectProtocolVersion float64

const (
	ProjectPolicyUpdateParamsConnectProtocolVersion1 ProjectPolicyUpdateParamsConnectProtocolVersion = 1
)

func (r ProjectPolicyUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ProjectPolicyUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ProjectPolicyUpdateParamsRole string

const (
	ProjectPolicyUpdateParamsRoleProjectRoleUnspecified ProjectPolicyUpdateParamsRole = "PROJECT_ROLE_UNSPECIFIED"
	ProjectPolicyUpdateParamsRoleProjectRoleAdmin       ProjectPolicyUpdateParamsRole = "PROJECT_ROLE_ADMIN"
	ProjectPolicyUpdateParamsRoleProjectRoleUser        ProjectPolicyUpdateParamsRole = "PROJECT_ROLE_USER"
)

func (r ProjectPolicyUpdateParamsRole) IsKnown() bool {
	switch r {
	case ProjectPolicyUpdateParamsRoleProjectRoleUnspecified, ProjectPolicyUpdateParamsRoleProjectRoleAdmin, ProjectPolicyUpdateParamsRoleProjectRoleUser:
		return true
	}
	return false
}

type ProjectPolicyListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[ProjectPolicyListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ProjectPolicyListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[ProjectPolicyListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[ProjectPolicyListParamsConnect] `query:"connect"`
	Message param.Field[string]                         `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [ProjectPolicyListParams]'s query parameters as
// `url.Values`.
func (r ProjectPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type ProjectPolicyListParamsEncoding string

const (
	ProjectPolicyListParamsEncodingProto ProjectPolicyListParamsEncoding = "proto"
	ProjectPolicyListParamsEncodingJson  ProjectPolicyListParamsEncoding = "json"
)

func (r ProjectPolicyListParamsEncoding) IsKnown() bool {
	switch r {
	case ProjectPolicyListParamsEncodingProto, ProjectPolicyListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type ProjectPolicyListParamsConnectProtocolVersion float64

const (
	ProjectPolicyListParamsConnectProtocolVersion1 ProjectPolicyListParamsConnectProtocolVersion = 1
)

func (r ProjectPolicyListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ProjectPolicyListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type ProjectPolicyListParamsCompression string

const (
	ProjectPolicyListParamsCompressionIdentity ProjectPolicyListParamsCompression = "identity"
	ProjectPolicyListParamsCompressionGzip     ProjectPolicyListParamsCompression = "gzip"
	ProjectPolicyListParamsCompressionBr       ProjectPolicyListParamsCompression = "br"
)

func (r ProjectPolicyListParamsCompression) IsKnown() bool {
	switch r {
	case ProjectPolicyListParamsCompressionIdentity, ProjectPolicyListParamsCompressionGzip, ProjectPolicyListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type ProjectPolicyListParamsConnect string

const (
	ProjectPolicyListParamsConnectV1 ProjectPolicyListParamsConnect = "v1"
)

func (r ProjectPolicyListParamsConnect) IsKnown() bool {
	switch r {
	case ProjectPolicyListParamsConnectV1:
		return true
	}
	return false
}

type ProjectPolicyDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ProjectPolicyDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// group_id specifies the group_id identifier
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ProjectPolicyDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ProjectPolicyDeleteParamsConnectProtocolVersion float64

const (
	ProjectPolicyDeleteParamsConnectProtocolVersion1 ProjectPolicyDeleteParamsConnectProtocolVersion = 1
)

func (r ProjectPolicyDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ProjectPolicyDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}
