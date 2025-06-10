// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
)

// UsageService contains methods and other services that help with interacting with
// the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageService] method instead.
type UsageService struct {
	Options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r *UsageService) {
	r = &UsageService{}
	r.Options = opts
	return
}

// Lists completed environment runtime records within a specified date range.
//
// Returns a list of environment runtime records that were completed within the
// specified date range. Records of currently running environments are not
// included.
//
// Use this method to:
//
// - View environment runtime records
// - Filter by project
// - Create custom usage reports
//
// ### Example
//
// ```yaml
// filter:
//
//	projectId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	dateRange:
//	  startTime: "2024-01-01T00:00:00Z"
//	  endTime: "2024-01-02T00:00:00Z"
//
// pagination:
//
//	pageSize: 100
//
// ```
func (r *UsageService) ListEnvironmentRuntimeRecords(ctx context.Context, params UsageListEnvironmentRuntimeRecordsParams, opts ...option.RequestOption) (res *pagination.RecordsPage[EnvironmentUsageRecord], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.UsageService/ListEnvironmentUsageRecords"
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

// Lists completed environment runtime records within a specified date range.
//
// Returns a list of environment runtime records that were completed within the
// specified date range. Records of currently running environments are not
// included.
//
// Use this method to:
//
// - View environment runtime records
// - Filter by project
// - Create custom usage reports
//
// ### Example
//
// ```yaml
// filter:
//
//	projectId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	dateRange:
//	  startTime: "2024-01-01T00:00:00Z"
//	  endTime: "2024-01-02T00:00:00Z"
//
// pagination:
//
//	pageSize: 100
//
// ```
func (r *UsageService) ListEnvironmentRuntimeRecordsAutoPaging(ctx context.Context, params UsageListEnvironmentRuntimeRecordsParams, opts ...option.RequestOption) *pagination.RecordsPageAutoPager[EnvironmentUsageRecord] {
	return pagination.NewRecordsPageAutoPager(r.ListEnvironmentRuntimeRecords(ctx, params, opts...))
}

// EnvironmentUsageRecord represents a record of an environment from start to stop.
type EnvironmentUsageRecord struct {
	// Environment usage record ID.
	ID string `json:"id"`
	// Time when the environment was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Environment class ID associated with the record.
	EnvironmentClassID string `json:"environmentClassId"`
	// Environment ID associated with the record.
	EnvironmentID string `json:"environmentId"`
	// Project ID associated with the environment (if available).
	ProjectID string `json:"projectId"`
	// Runner ID associated with the environment.
	RunnerID string `json:"runnerId"`
	// Time when the environment was stopped.
	StoppedAt time.Time `json:"stoppedAt" format:"date-time"`
	// User ID is the ID of the user who created the environment associated with the
	// record.
	UserID string                     `json:"userId"`
	JSON   environmentUsageRecordJSON `json:"-"`
}

// environmentUsageRecordJSON contains the JSON metadata for the struct
// [EnvironmentUsageRecord]
type environmentUsageRecordJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	EnvironmentClassID apijson.Field
	EnvironmentID      apijson.Field
	ProjectID          apijson.Field
	RunnerID           apijson.Field
	StoppedAt          apijson.Field
	UserID             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentUsageRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentUsageRecordJSON) RawJSON() string {
	return r.raw
}

type UsageListEnvironmentRuntimeRecordsParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// Filter options.
	Filter param.Field[UsageListEnvironmentRuntimeRecordsParamsFilter] `json:"filter"`
	// Pagination options.
	Pagination param.Field[UsageListEnvironmentRuntimeRecordsParamsPagination] `json:"pagination"`
}

func (r UsageListEnvironmentRuntimeRecordsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [UsageListEnvironmentRuntimeRecordsParams]'s query
// parameters as `url.Values`.
func (r UsageListEnvironmentRuntimeRecordsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter options.
type UsageListEnvironmentRuntimeRecordsParamsFilter struct {
	// Date range to query runtime records within.
	DateRange param.Field[UsageListEnvironmentRuntimeRecordsParamsFilterDateRange] `json:"dateRange,required"`
	// Optional project ID to filter runtime records by.
	ProjectID param.Field[string] `json:"projectId"`
}

func (r UsageListEnvironmentRuntimeRecordsParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Date range to query runtime records within.
type UsageListEnvironmentRuntimeRecordsParamsFilterDateRange struct {
	// End time of the date range (exclusive).
	EndTime param.Field[time.Time] `json:"endTime,required" format:"date-time"`
	// Start time of the date range (inclusive).
	StartTime param.Field[time.Time] `json:"startTime,required" format:"date-time"`
}

func (r UsageListEnvironmentRuntimeRecordsParamsFilterDateRange) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Pagination options.
type UsageListEnvironmentRuntimeRecordsParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r UsageListEnvironmentRuntimeRecordsParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
