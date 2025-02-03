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

// RunnerConfigurationEnvironmentClassService contains methods and other services
// that help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerConfigurationEnvironmentClassService] method instead.
type RunnerConfigurationEnvironmentClassService struct {
	Options []option.RequestOption
}

// NewRunnerConfigurationEnvironmentClassService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRunnerConfigurationEnvironmentClassService(opts ...option.RequestOption) (r *RunnerConfigurationEnvironmentClassService) {
	r = &RunnerConfigurationEnvironmentClassService{}
	r.Options = opts
	return
}

// CreateEnvironmentClass creates a new environment class on a runner.
func (r *RunnerConfigurationEnvironmentClassService) New(ctx context.Context, params RunnerConfigurationEnvironmentClassNewParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/CreateEnvironmentClass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetEnvironmentClass returns a single environment class configured for a runner.
func (r *RunnerConfigurationEnvironmentClassService) Get(ctx context.Context, params RunnerConfigurationEnvironmentClassGetParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetEnvironmentClass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateEnvironmentClass updates an existing environment class on a runner.
func (r *RunnerConfigurationEnvironmentClassService) Update(ctx context.Context, params RunnerConfigurationEnvironmentClassUpdateParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/UpdateEnvironmentClass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListEnvironmentClasses returns all environment classes configured for a runner.
//
// buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE
func (r *RunnerConfigurationEnvironmentClassService) List(ctx context.Context, params RunnerConfigurationEnvironmentClassListParams, opts ...option.RequestOption) (res *RunnerConfigurationEnvironmentClassListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/ListEnvironmentClasses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type RunnerConfigurationEnvironmentClassNewResponse struct {
	ID   string                                             `json:"id"`
	JSON runnerConfigurationEnvironmentClassNewResponseJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassNewResponseJSON contains the JSON metadata
// for the struct [RunnerConfigurationEnvironmentClassNewResponse]
type runnerConfigurationEnvironmentClassNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassNewResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassGetResponse struct {
	EnvironmentClass RunnerConfigurationEnvironmentClassGetResponseEnvironmentClass `json:"environmentClass"`
	JSON             runnerConfigurationEnvironmentClassGetResponseJSON             `json:"-"`
}

// runnerConfigurationEnvironmentClassGetResponseJSON contains the JSON metadata
// for the struct [RunnerConfigurationEnvironmentClassGetResponse]
type runnerConfigurationEnvironmentClassGetResponseJSON struct {
	EnvironmentClass apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassGetResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassGetResponseEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID string `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration []RunnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfiguration `json:"configuration"`
	// description is a human readable description of the environment class
	Description string `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName string `json:"displayName"`
	// enabled indicates whether the environment class can be used to create
	//
	// new environments.
	Enabled bool `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID string                                                             `json:"runnerId"`
	JSON     runnerConfigurationEnvironmentClassGetResponseEnvironmentClassJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassGetResponseEnvironmentClassJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationEnvironmentClassGetResponseEnvironmentClass]
type runnerConfigurationEnvironmentClassGetResponseEnvironmentClassJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	RunnerID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassGetResponseEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassGetResponseEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfiguration struct {
	Key   string                                                                          `json:"key"`
	Value string                                                                          `json:"value"`
	JSON  runnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfigurationJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfigurationJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfiguration]
type runnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfigurationJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassGetResponseEnvironmentClassConfigurationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassUpdateResponse = interface{}

type RunnerConfigurationEnvironmentClassListResponse struct {
	EnvironmentClasses []RunnerConfigurationEnvironmentClassListResponseEnvironmentClass `json:"environmentClasses"`
	// pagination contains the pagination options for listing environment classes
	Pagination RunnerConfigurationEnvironmentClassListResponsePagination `json:"pagination"`
	JSON       runnerConfigurationEnvironmentClassListResponseJSON       `json:"-"`
}

// runnerConfigurationEnvironmentClassListResponseJSON contains the JSON metadata
// for the struct [RunnerConfigurationEnvironmentClassListResponse]
type runnerConfigurationEnvironmentClassListResponseJSON struct {
	EnvironmentClasses apijson.Field
	Pagination         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassListResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassListResponseEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID string `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration []RunnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfiguration `json:"configuration"`
	// description is a human readable description of the environment class
	Description string `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName string `json:"displayName"`
	// enabled indicates whether the environment class can be used to create
	//
	// new environments.
	Enabled bool `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID string                                                              `json:"runnerId"`
	JSON     runnerConfigurationEnvironmentClassListResponseEnvironmentClassJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassListResponseEnvironmentClassJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationEnvironmentClassListResponseEnvironmentClass]
type runnerConfigurationEnvironmentClassListResponseEnvironmentClassJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	RunnerID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassListResponseEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassListResponseEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfiguration struct {
	Key   string                                                                             `json:"key"`
	Value string                                                                             `json:"value"`
	JSON  runnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfigurationJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfigurationJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfiguration]
type runnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfigurationJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassListResponseEnvironmentClassesConfigurationJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing environment classes
type RunnerConfigurationEnvironmentClassListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                                        `json:"nextToken"`
	JSON      runnerConfigurationEnvironmentClassListResponsePaginationJSON `json:"-"`
}

// runnerConfigurationEnvironmentClassListResponsePaginationJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationEnvironmentClassListResponsePagination]
type runnerConfigurationEnvironmentClassListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationEnvironmentClassListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationEnvironmentClassListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationEnvironmentClassNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationEnvironmentClassNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Configuration          param.Field[[]RunnerConfigurationEnvironmentClassNewParamsConfiguration]        `json:"configuration"`
	Description            param.Field[string]                                                             `json:"description"`
	DisplayName            param.Field[string]                                                             `json:"displayName"`
	RunnerID               param.Field[string]                                                             `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationEnvironmentClassNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationEnvironmentClassNewParamsConnectProtocolVersion float64

const (
	RunnerConfigurationEnvironmentClassNewParamsConnectProtocolVersion1 RunnerConfigurationEnvironmentClassNewParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationEnvironmentClassNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationEnvironmentClassNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationEnvironmentClassNewParamsConfiguration struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r RunnerConfigurationEnvironmentClassNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationEnvironmentClassGetParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationEnvironmentClassGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	EnvironmentClassID     param.Field[string]                                                             `json:"environmentClassId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationEnvironmentClassGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationEnvironmentClassGetParamsConnectProtocolVersion float64

const (
	RunnerConfigurationEnvironmentClassGetParamsConnectProtocolVersion1 RunnerConfigurationEnvironmentClassGetParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationEnvironmentClassGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationEnvironmentClassGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationEnvironmentClassUpdateParams struct {
	Body RunnerConfigurationEnvironmentClassUpdateParamsBodyUnion `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationEnvironmentClassUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationEnvironmentClassUpdateParamsBody struct {
	Description param.Field[string] `json:"description"`
	DisplayName param.Field[string] `json:"displayName"`
	Enabled     param.Field[bool]   `json:"enabled"`
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBody) implementsRunnerConfigurationEnvironmentClassUpdateParamsBodyUnion() {
}

// Satisfied by [RunnerConfigurationEnvironmentClassUpdateParamsBodyDescription],
// [RunnerConfigurationEnvironmentClassUpdateParamsBodyDisplayName],
// [RunnerConfigurationEnvironmentClassUpdateParamsBodyEnabled],
// [RunnerConfigurationEnvironmentClassUpdateParamsBody].
type RunnerConfigurationEnvironmentClassUpdateParamsBodyUnion interface {
	implementsRunnerConfigurationEnvironmentClassUpdateParamsBodyUnion()
}

type RunnerConfigurationEnvironmentClassUpdateParamsBodyDescription struct {
	Description param.Field[string] `json:"description,required"`
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyDescription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyDescription) implementsRunnerConfigurationEnvironmentClassUpdateParamsBodyUnion() {
}

type RunnerConfigurationEnvironmentClassUpdateParamsBodyDisplayName struct {
	DisplayName param.Field[string] `json:"displayName,required"`
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyDisplayName) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyDisplayName) implementsRunnerConfigurationEnvironmentClassUpdateParamsBodyUnion() {
}

type RunnerConfigurationEnvironmentClassUpdateParamsBodyEnabled struct {
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationEnvironmentClassUpdateParamsBodyEnabled) implementsRunnerConfigurationEnvironmentClassUpdateParamsBodyUnion() {
}

// Define the version of the Connect protocol
type RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion float64

const (
	RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion1 RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationEnvironmentClassListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[RunnerConfigurationEnvironmentClassListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[RunnerConfigurationEnvironmentClassListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[RunnerConfigurationEnvironmentClassListParamsConnect] `query:"connect"`
	Message param.Field[string]                                               `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [RunnerConfigurationEnvironmentClassListParams]'s query
// parameters as `url.Values`.
func (r RunnerConfigurationEnvironmentClassListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type RunnerConfigurationEnvironmentClassListParamsEncoding string

const (
	RunnerConfigurationEnvironmentClassListParamsEncodingProto RunnerConfigurationEnvironmentClassListParamsEncoding = "proto"
	RunnerConfigurationEnvironmentClassListParamsEncodingJson  RunnerConfigurationEnvironmentClassListParamsEncoding = "json"
)

func (r RunnerConfigurationEnvironmentClassListParamsEncoding) IsKnown() bool {
	switch r {
	case RunnerConfigurationEnvironmentClassListParamsEncodingProto, RunnerConfigurationEnvironmentClassListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion float64

const (
	RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion1 RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type RunnerConfigurationEnvironmentClassListParamsCompression string

const (
	RunnerConfigurationEnvironmentClassListParamsCompressionIdentity RunnerConfigurationEnvironmentClassListParamsCompression = "identity"
	RunnerConfigurationEnvironmentClassListParamsCompressionGzip     RunnerConfigurationEnvironmentClassListParamsCompression = "gzip"
	RunnerConfigurationEnvironmentClassListParamsCompressionBr       RunnerConfigurationEnvironmentClassListParamsCompression = "br"
)

func (r RunnerConfigurationEnvironmentClassListParamsCompression) IsKnown() bool {
	switch r {
	case RunnerConfigurationEnvironmentClassListParamsCompressionIdentity, RunnerConfigurationEnvironmentClassListParamsCompressionGzip, RunnerConfigurationEnvironmentClassListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerConfigurationEnvironmentClassListParamsConnect string

const (
	RunnerConfigurationEnvironmentClassListParamsConnectV1 RunnerConfigurationEnvironmentClassListParamsConnect = "v1"
)

func (r RunnerConfigurationEnvironmentClassListParamsConnect) IsKnown() bool {
	switch r {
	case RunnerConfigurationEnvironmentClassListParamsConnectV1:
		return true
	}
	return false
}
