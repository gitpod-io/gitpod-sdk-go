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

// EnvironmentClassService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentClassService] method instead.
type EnvironmentClassService struct {
	Options []option.RequestOption
}

// NewEnvironmentClassService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnvironmentClassService(opts ...option.RequestOption) (r *EnvironmentClassService) {
	r = &EnvironmentClassService{}
	r.Options = opts
	return
}

// ListEnvironmentClasses returns the list of environment classes with runner
// details a user is able to use based on the
//
// query buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE
func (r *EnvironmentClassService) List(ctx context.Context, params EnvironmentClassListParams, opts ...option.RequestOption) (res *EnvironmentClassListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/ListEnvironmentClasses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type EnvironmentClassListResponse struct {
	EnvironmentClasses []EnvironmentClassListResponseEnvironmentClass `json:"environmentClasses"`
	// pagination contains the pagination options for listing environment classes
	Pagination EnvironmentClassListResponsePagination `json:"pagination"`
	JSON       environmentClassListResponseJSON       `json:"-"`
}

// environmentClassListResponseJSON contains the JSON metadata for the struct
// [EnvironmentClassListResponse]
type environmentClassListResponseJSON struct {
	EnvironmentClasses apijson.Field
	Pagination         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentClassListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentClassListResponseEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID string `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration []EnvironmentClassListResponseEnvironmentClassesConfiguration `json:"configuration"`
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
	RunnerID string                                           `json:"runnerId"`
	JSON     environmentClassListResponseEnvironmentClassJSON `json:"-"`
}

// environmentClassListResponseEnvironmentClassJSON contains the JSON metadata for
// the struct [EnvironmentClassListResponseEnvironmentClass]
type environmentClassListResponseEnvironmentClassJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	RunnerID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentClassListResponseEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassListResponseEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type EnvironmentClassListResponseEnvironmentClassesConfiguration struct {
	Key   string                                                          `json:"key"`
	Value string                                                          `json:"value"`
	JSON  environmentClassListResponseEnvironmentClassesConfigurationJSON `json:"-"`
}

// environmentClassListResponseEnvironmentClassesConfigurationJSON contains the
// JSON metadata for the struct
// [EnvironmentClassListResponseEnvironmentClassesConfiguration]
type environmentClassListResponseEnvironmentClassesConfigurationJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentClassListResponseEnvironmentClassesConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassListResponseEnvironmentClassesConfigurationJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing environment classes
type EnvironmentClassListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                     `json:"nextToken"`
	JSON      environmentClassListResponsePaginationJSON `json:"-"`
}

// environmentClassListResponsePaginationJSON contains the JSON metadata for the
// struct [EnvironmentClassListResponsePagination]
type environmentClassListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentClassListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type EnvironmentClassListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[EnvironmentClassListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EnvironmentClassListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[EnvironmentClassListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[EnvironmentClassListParamsConnect] `query:"connect"`
	Message param.Field[string]                            `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [EnvironmentClassListParams]'s query parameters as
// `url.Values`.
func (r EnvironmentClassListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type EnvironmentClassListParamsEncoding string

const (
	EnvironmentClassListParamsEncodingProto EnvironmentClassListParamsEncoding = "proto"
	EnvironmentClassListParamsEncodingJson  EnvironmentClassListParamsEncoding = "json"
)

func (r EnvironmentClassListParamsEncoding) IsKnown() bool {
	switch r {
	case EnvironmentClassListParamsEncodingProto, EnvironmentClassListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EnvironmentClassListParamsConnectProtocolVersion float64

const (
	EnvironmentClassListParamsConnectProtocolVersion1 EnvironmentClassListParamsConnectProtocolVersion = 1
)

func (r EnvironmentClassListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EnvironmentClassListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type EnvironmentClassListParamsCompression string

const (
	EnvironmentClassListParamsCompressionIdentity EnvironmentClassListParamsCompression = "identity"
	EnvironmentClassListParamsCompressionGzip     EnvironmentClassListParamsCompression = "gzip"
	EnvironmentClassListParamsCompressionBr       EnvironmentClassListParamsCompression = "br"
)

func (r EnvironmentClassListParamsCompression) IsKnown() bool {
	switch r {
	case EnvironmentClassListParamsCompressionIdentity, EnvironmentClassListParamsCompressionGzip, EnvironmentClassListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EnvironmentClassListParamsConnect string

const (
	EnvironmentClassListParamsConnectV1 EnvironmentClassListParamsConnect = "v1"
)

func (r EnvironmentClassListParamsConnect) IsKnown() bool {
	switch r {
	case EnvironmentClassListParamsConnectV1:
		return true
	}
	return false
}
