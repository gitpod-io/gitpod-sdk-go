// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
	"github.com/tidwall/gjson"
)

// RunnerConfigurationScmIntegrationService contains methods and other services
// that help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerConfigurationScmIntegrationService] method instead.
type RunnerConfigurationScmIntegrationService struct {
	Options []option.RequestOption
}

// NewRunnerConfigurationScmIntegrationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRunnerConfigurationScmIntegrationService(opts ...option.RequestOption) (r *RunnerConfigurationScmIntegrationService) {
	r = &RunnerConfigurationScmIntegrationService{}
	r.Options = opts
	return
}

// CreateSCMIntegration creates a new SCM integration on a runner.
func (r *RunnerConfigurationScmIntegrationService) New(ctx context.Context, params RunnerConfigurationScmIntegrationNewParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/CreateSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetSCMIntegration returns a single SCM integration configured for a runner.
func (r *RunnerConfigurationScmIntegrationService) Get(ctx context.Context, params RunnerConfigurationScmIntegrationGetParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// UpdateSCMIntegration updates an existing SCM integration on a runner.
func (r *RunnerConfigurationScmIntegrationService) Update(ctx context.Context, params RunnerConfigurationScmIntegrationUpdateParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/UpdateSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListSCMIntegrations returns all SCM integrations configured for a runner.
func (r *RunnerConfigurationScmIntegrationService) List(ctx context.Context, params RunnerConfigurationScmIntegrationListParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/ListSCMIntegrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// DeleteSCMIntegration deletes an existing SCM integration on a runner.
func (r *RunnerConfigurationScmIntegrationService) Delete(ctx context.Context, params RunnerConfigurationScmIntegrationDeleteParams, opts ...option.RequestOption) (res *RunnerConfigurationScmIntegrationDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/DeleteSCMIntegration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type RunnerConfigurationScmIntegrationNewResponse struct {
	// id is a uniquely generated identifier for the SCM integration
	ID   string                                           `json:"id" format:"uuid"`
	JSON runnerConfigurationScmIntegrationNewResponseJSON `json:"-"`
}

// runnerConfigurationScmIntegrationNewResponseJSON contains the JSON metadata for
// the struct [RunnerConfigurationScmIntegrationNewResponse]
type runnerConfigurationScmIntegrationNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationNewResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationGetResponse struct {
	Integration RunnerConfigurationScmIntegrationGetResponseIntegration `json:"integration"`
	JSON        runnerConfigurationScmIntegrationGetResponseJSON        `json:"-"`
}

// runnerConfigurationScmIntegrationGetResponseJSON contains the JSON metadata for
// the struct [RunnerConfigurationScmIntegrationGetResponse]
type runnerConfigurationScmIntegrationGetResponseJSON struct {
	Integration apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationGetResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationGetResponseIntegration struct {
	// id is the unique identifier of the SCM integration
	ID   string `json:"id"`
	Host string `json:"host"`
	// This field can have the runtime type of
	// [RunnerConfigurationScmIntegrationGetResponseIntegrationObjectOAuth].
	OAuth    interface{} `json:"oauth"`
	Pat      bool        `json:"pat"`
	RunnerID string      `json:"runnerId"`
	// scm_id references the scm_id in the runner's configuration schema that this
	// integration is for
	ScmID string                                                      `json:"scmId"`
	JSON  runnerConfigurationScmIntegrationGetResponseIntegrationJSON `json:"-"`
	union RunnerConfigurationScmIntegrationGetResponseIntegrationUnion
}

// runnerConfigurationScmIntegrationGetResponseIntegrationJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationScmIntegrationGetResponseIntegration]
type runnerConfigurationScmIntegrationGetResponseIntegrationJSON struct {
	ID          apijson.Field
	Host        apijson.Field
	OAuth       apijson.Field
	Pat         apijson.Field
	RunnerID    apijson.Field
	ScmID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r runnerConfigurationScmIntegrationGetResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationScmIntegrationGetResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationScmIntegrationGetResponseIntegration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RunnerConfigurationScmIntegrationGetResponseIntegrationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationScmIntegrationGetResponseIntegrationObject],
// [RunnerConfigurationScmIntegrationGetResponseIntegrationObject].
func (r RunnerConfigurationScmIntegrationGetResponseIntegration) AsUnion() RunnerConfigurationScmIntegrationGetResponseIntegrationUnion {
	return r.union
}

// Union satisfied by
// [RunnerConfigurationScmIntegrationGetResponseIntegrationObject] or
// [RunnerConfigurationScmIntegrationGetResponseIntegrationObject].
type RunnerConfigurationScmIntegrationGetResponseIntegrationUnion interface {
	implementsRunnerConfigurationScmIntegrationGetResponseIntegration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationScmIntegrationGetResponseIntegrationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationScmIntegrationGetResponseIntegrationObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationScmIntegrationGetResponseIntegrationObject{}),
		},
	)
}

type RunnerConfigurationScmIntegrationGetResponseIntegrationObject struct {
	OAuth RunnerConfigurationScmIntegrationGetResponseIntegrationObjectOAuth `json:"oauth,required"`
	// id is the unique identifier of the SCM integration
	ID       string `json:"id"`
	Host     string `json:"host"`
	Pat      bool   `json:"pat"`
	RunnerID string `json:"runnerId"`
	// scm_id references the scm_id in the runner's configuration schema that this
	// integration is for
	ScmID string                                                            `json:"scmId"`
	JSON  runnerConfigurationScmIntegrationGetResponseIntegrationObjectJSON `json:"-"`
}

// runnerConfigurationScmIntegrationGetResponseIntegrationObjectJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationScmIntegrationGetResponseIntegrationObject]
type runnerConfigurationScmIntegrationGetResponseIntegrationObjectJSON struct {
	OAuth       apijson.Field
	ID          apijson.Field
	Host        apijson.Field
	Pat         apijson.Field
	RunnerID    apijson.Field
	ScmID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationGetResponseIntegrationObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationGetResponseIntegrationObjectJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationScmIntegrationGetResponseIntegrationObject) implementsRunnerConfigurationScmIntegrationGetResponseIntegration() {
}

type RunnerConfigurationScmIntegrationGetResponseIntegrationObjectOAuth struct {
	// client_id is the OAuth app's client ID in clear text.
	ClientID string `json:"clientId"`
	// encrypted_client_secret is the OAuth app's secret encrypted with the runner's
	// public key.
	EncryptedClientSecret string                                                                 `json:"encryptedClientSecret" format:"byte"`
	JSON                  runnerConfigurationScmIntegrationGetResponseIntegrationObjectOAuthJSON `json:"-"`
}

// runnerConfigurationScmIntegrationGetResponseIntegrationObjectOAuthJSON contains
// the JSON metadata for the struct
// [RunnerConfigurationScmIntegrationGetResponseIntegrationObjectOAuth]
type runnerConfigurationScmIntegrationGetResponseIntegrationObjectOAuthJSON struct {
	ClientID              apijson.Field
	EncryptedClientSecret apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationGetResponseIntegrationObjectOAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationGetResponseIntegrationObjectOAuthJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationUpdateResponse = interface{}

type RunnerConfigurationScmIntegrationListResponse struct {
	Integrations []RunnerConfigurationScmIntegrationListResponseIntegration `json:"integrations"`
	// pagination contains the pagination options for listing scm integrations
	Pagination RunnerConfigurationScmIntegrationListResponsePagination `json:"pagination"`
	JSON       runnerConfigurationScmIntegrationListResponseJSON       `json:"-"`
}

// runnerConfigurationScmIntegrationListResponseJSON contains the JSON metadata for
// the struct [RunnerConfigurationScmIntegrationListResponse]
type runnerConfigurationScmIntegrationListResponseJSON struct {
	Integrations apijson.Field
	Pagination   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationListResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationListResponseIntegration struct {
	// id is the unique identifier of the SCM integration
	ID   string `json:"id"`
	Host string `json:"host"`
	// This field can have the runtime type of
	// [RunnerConfigurationScmIntegrationListResponseIntegrationsObjectOAuth].
	OAuth    interface{} `json:"oauth"`
	Pat      bool        `json:"pat"`
	RunnerID string      `json:"runnerId"`
	// scm_id references the scm_id in the runner's configuration schema that this
	// integration is for
	ScmID string                                                       `json:"scmId"`
	JSON  runnerConfigurationScmIntegrationListResponseIntegrationJSON `json:"-"`
	union RunnerConfigurationScmIntegrationListResponseIntegrationsUnion
}

// runnerConfigurationScmIntegrationListResponseIntegrationJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationScmIntegrationListResponseIntegration]
type runnerConfigurationScmIntegrationListResponseIntegrationJSON struct {
	ID          apijson.Field
	Host        apijson.Field
	OAuth       apijson.Field
	Pat         apijson.Field
	RunnerID    apijson.Field
	ScmID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r runnerConfigurationScmIntegrationListResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationScmIntegrationListResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationScmIntegrationListResponseIntegration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RunnerConfigurationScmIntegrationListResponseIntegrationsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationScmIntegrationListResponseIntegrationsObject],
// [RunnerConfigurationScmIntegrationListResponseIntegrationsObject].
func (r RunnerConfigurationScmIntegrationListResponseIntegration) AsUnion() RunnerConfigurationScmIntegrationListResponseIntegrationsUnion {
	return r.union
}

// Union satisfied by
// [RunnerConfigurationScmIntegrationListResponseIntegrationsObject] or
// [RunnerConfigurationScmIntegrationListResponseIntegrationsObject].
type RunnerConfigurationScmIntegrationListResponseIntegrationsUnion interface {
	implementsRunnerConfigurationScmIntegrationListResponseIntegration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationScmIntegrationListResponseIntegrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationScmIntegrationListResponseIntegrationsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationScmIntegrationListResponseIntegrationsObject{}),
		},
	)
}

type RunnerConfigurationScmIntegrationListResponseIntegrationsObject struct {
	OAuth RunnerConfigurationScmIntegrationListResponseIntegrationsObjectOAuth `json:"oauth,required"`
	// id is the unique identifier of the SCM integration
	ID       string `json:"id"`
	Host     string `json:"host"`
	Pat      bool   `json:"pat"`
	RunnerID string `json:"runnerId"`
	// scm_id references the scm_id in the runner's configuration schema that this
	// integration is for
	ScmID string                                                              `json:"scmId"`
	JSON  runnerConfigurationScmIntegrationListResponseIntegrationsObjectJSON `json:"-"`
}

// runnerConfigurationScmIntegrationListResponseIntegrationsObjectJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationScmIntegrationListResponseIntegrationsObject]
type runnerConfigurationScmIntegrationListResponseIntegrationsObjectJSON struct {
	OAuth       apijson.Field
	ID          apijson.Field
	Host        apijson.Field
	Pat         apijson.Field
	RunnerID    apijson.Field
	ScmID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationListResponseIntegrationsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationListResponseIntegrationsObjectJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationScmIntegrationListResponseIntegrationsObject) implementsRunnerConfigurationScmIntegrationListResponseIntegration() {
}

type RunnerConfigurationScmIntegrationListResponseIntegrationsObjectOAuth struct {
	// client_id is the OAuth app's client ID in clear text.
	ClientID string `json:"clientId"`
	// encrypted_client_secret is the OAuth app's secret encrypted with the runner's
	// public key.
	EncryptedClientSecret string                                                                   `json:"encryptedClientSecret" format:"byte"`
	JSON                  runnerConfigurationScmIntegrationListResponseIntegrationsObjectOAuthJSON `json:"-"`
}

// runnerConfigurationScmIntegrationListResponseIntegrationsObjectOAuthJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationScmIntegrationListResponseIntegrationsObjectOAuth]
type runnerConfigurationScmIntegrationListResponseIntegrationsObjectOAuthJSON struct {
	ClientID              apijson.Field
	EncryptedClientSecret apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationListResponseIntegrationsObjectOAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationListResponseIntegrationsObjectOAuthJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing scm integrations
type RunnerConfigurationScmIntegrationListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                                      `json:"nextToken"`
	JSON      runnerConfigurationScmIntegrationListResponsePaginationJSON `json:"-"`
}

// runnerConfigurationScmIntegrationListResponsePaginationJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationScmIntegrationListResponsePagination]
type runnerConfigurationScmIntegrationListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationScmIntegrationListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationScmIntegrationListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationScmIntegrationDeleteResponse = interface{}

type RunnerConfigurationScmIntegrationNewParams struct {
	Body RunnerConfigurationScmIntegrationNewParamsBody `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationScmIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationScmIntegrationNewParamsBody struct {
}

func (r RunnerConfigurationScmIntegrationNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion float64

const (
	RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion1 RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationScmIntegrationGetParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[RunnerConfigurationScmIntegrationGetParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[RunnerConfigurationScmIntegrationGetParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[RunnerConfigurationScmIntegrationGetParamsConnect] `query:"connect"`
	Message param.Field[string]                                            `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [RunnerConfigurationScmIntegrationGetParams]'s query
// parameters as `url.Values`.
func (r RunnerConfigurationScmIntegrationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type RunnerConfigurationScmIntegrationGetParamsEncoding string

const (
	RunnerConfigurationScmIntegrationGetParamsEncodingProto RunnerConfigurationScmIntegrationGetParamsEncoding = "proto"
	RunnerConfigurationScmIntegrationGetParamsEncodingJson  RunnerConfigurationScmIntegrationGetParamsEncoding = "json"
)

func (r RunnerConfigurationScmIntegrationGetParamsEncoding) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationGetParamsEncodingProto, RunnerConfigurationScmIntegrationGetParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion float64

const (
	RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion1 RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type RunnerConfigurationScmIntegrationGetParamsCompression string

const (
	RunnerConfigurationScmIntegrationGetParamsCompressionIdentity RunnerConfigurationScmIntegrationGetParamsCompression = "identity"
	RunnerConfigurationScmIntegrationGetParamsCompressionGzip     RunnerConfigurationScmIntegrationGetParamsCompression = "gzip"
	RunnerConfigurationScmIntegrationGetParamsCompressionBr       RunnerConfigurationScmIntegrationGetParamsCompression = "br"
)

func (r RunnerConfigurationScmIntegrationGetParamsCompression) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationGetParamsCompressionIdentity, RunnerConfigurationScmIntegrationGetParamsCompressionGzip, RunnerConfigurationScmIntegrationGetParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationGetParamsConnect string

const (
	RunnerConfigurationScmIntegrationGetParamsConnectV1 RunnerConfigurationScmIntegrationGetParamsConnect = "v1"
)

func (r RunnerConfigurationScmIntegrationGetParamsConnect) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationGetParamsConnectV1:
		return true
	}
	return false
}

type RunnerConfigurationScmIntegrationUpdateParams struct {
	Body RunnerConfigurationScmIntegrationUpdateParamsBody `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationScmIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationScmIntegrationUpdateParamsBody struct {
}

func (r RunnerConfigurationScmIntegrationUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion float64

const (
	RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion1 RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerConfigurationScmIntegrationListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[RunnerConfigurationScmIntegrationListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[RunnerConfigurationScmIntegrationListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[RunnerConfigurationScmIntegrationListParamsConnect] `query:"connect"`
	Message param.Field[string]                                             `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [RunnerConfigurationScmIntegrationListParams]'s query
// parameters as `url.Values`.
func (r RunnerConfigurationScmIntegrationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type RunnerConfigurationScmIntegrationListParamsEncoding string

const (
	RunnerConfigurationScmIntegrationListParamsEncodingProto RunnerConfigurationScmIntegrationListParamsEncoding = "proto"
	RunnerConfigurationScmIntegrationListParamsEncodingJson  RunnerConfigurationScmIntegrationListParamsEncoding = "json"
)

func (r RunnerConfigurationScmIntegrationListParamsEncoding) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationListParamsEncodingProto, RunnerConfigurationScmIntegrationListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion float64

const (
	RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion1 RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type RunnerConfigurationScmIntegrationListParamsCompression string

const (
	RunnerConfigurationScmIntegrationListParamsCompressionIdentity RunnerConfigurationScmIntegrationListParamsCompression = "identity"
	RunnerConfigurationScmIntegrationListParamsCompressionGzip     RunnerConfigurationScmIntegrationListParamsCompression = "gzip"
	RunnerConfigurationScmIntegrationListParamsCompressionBr       RunnerConfigurationScmIntegrationListParamsCompression = "br"
)

func (r RunnerConfigurationScmIntegrationListParamsCompression) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationListParamsCompressionIdentity, RunnerConfigurationScmIntegrationListParamsCompressionGzip, RunnerConfigurationScmIntegrationListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationListParamsConnect string

const (
	RunnerConfigurationScmIntegrationListParamsConnectV1 RunnerConfigurationScmIntegrationListParamsConnect = "v1"
)

func (r RunnerConfigurationScmIntegrationListParamsConnect) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationListParamsConnectV1:
		return true
	}
	return false
}

type RunnerConfigurationScmIntegrationDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                                              `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationScmIntegrationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion float64

const (
	RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion1 RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}
