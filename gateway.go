// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// GatewayService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayService] method instead.
type GatewayService struct {
	Options []option.RequestOption
}

// NewGatewayService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGatewayService(opts ...option.RequestOption) (r *GatewayService) {
	r = &GatewayService{}
	r.Options = opts
	return
}

// ListGateways
func (r *GatewayService) List(ctx context.Context, params GatewayListParams, opts ...option.RequestOption) (res *pagination.GatewaysPage[shared.Gateway], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.GatewayService/ListGateways"
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

// ListGateways
func (r *GatewayService) ListAutoPaging(ctx context.Context, params GatewayListParams, opts ...option.RequestOption) *pagination.GatewaysPageAutoPager[shared.Gateway] {
	return pagination.NewGatewaysPageAutoPager(r.List(ctx, params, opts...))
}

type GatewayListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// pagination contains the pagination options for listing gateways
	Pagination param.Field[GatewayListParamsPagination] `json:"pagination"`
}

func (r GatewayListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [GatewayListParams]'s query parameters as `url.Values`.
func (r GatewayListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// pagination contains the pagination options for listing gateways
type GatewayListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r GatewayListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
