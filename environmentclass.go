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
// details a user is able to use based on the query buf:lint:ignore
// RPC_REQUEST_RESPONSE_UNIQUE
func (r *EnvironmentClassService) List(ctx context.Context, params EnvironmentClassListParams, opts ...option.RequestOption) (res *pagination.EnvironmentClassesPage[shared.EnvironmentClass], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.EnvironmentService/ListEnvironmentClasses"
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

// ListEnvironmentClasses returns the list of environment classes with runner
// details a user is able to use based on the query buf:lint:ignore
// RPC_REQUEST_RESPONSE_UNIQUE
func (r *EnvironmentClassService) ListAutoPaging(ctx context.Context, params EnvironmentClassListParams, opts ...option.RequestOption) *pagination.EnvironmentClassesPageAutoPager[shared.EnvironmentClass] {
	return pagination.NewEnvironmentClassesPageAutoPager(r.List(ctx, params, opts...))
}

type EnvironmentClassListParams struct {
	Token    param.Field[string]                           `query:"token"`
	PageSize param.Field[int64]                            `query:"pageSize"`
	Filter   param.Field[EnvironmentClassListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing environment classes
	Pagination param.Field[EnvironmentClassListParamsPagination] `json:"pagination"`
}

func (r EnvironmentClassListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [EnvironmentClassListParams]'s query parameters as
// `url.Values`.
func (r EnvironmentClassListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EnvironmentClassListParamsFilter struct {
	// enabled filters the response to only enabled or disabled environment classes. If
	// not set, all environment classes are returned.
	Enabled param.Field[bool] `json:"enabled"`
	// runner_ids filters the response to only EnvironmentClasses of these Runner IDs
	RunnerIDs param.Field[[]string] `json:"runnerIds" format:"uuid"`
}

func (r EnvironmentClassListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing environment classes
type EnvironmentClassListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EnvironmentClassListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
