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

// Lists completed environment sessions within a specified date range.
//
// Returns a list of environment sessions that were completed within the specified
// date range. Currently running sessions are not included.
//
// Use this method to:
//
// - View environment sessions
// - Filter by project
// - Monitor session activity
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
func (r *UsageService) ListEnvironmentSessions(ctx context.Context, params UsageListEnvironmentSessionsParams, opts ...option.RequestOption) (res *pagination.SessionsPage[EnvironmentSession], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.UsageService/ListEnvironmentSessions"
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

// Lists completed environment sessions within a specified date range.
//
// Returns a list of environment sessions that were completed within the specified
// date range. Currently running sessions are not included.
//
// Use this method to:
//
// - View environment sessions
// - Filter by project
// - Monitor session activity
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
func (r *UsageService) ListEnvironmentSessionsAutoPaging(ctx context.Context, params UsageListEnvironmentSessionsParams, opts ...option.RequestOption) *pagination.SessionsPageAutoPager[EnvironmentSession] {
	return pagination.NewSessionsPageAutoPager(r.ListEnvironmentSessions(ctx, params, opts...))
}

type EnvironmentSession struct {
	// Environment session ID.
	ID string `json:"id"`
	// Time when the session was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Environment class ID associated with the session.
	EnvironmentClassID string `json:"environmentClassId"`
	// Environment ID associated with the session.
	EnvironmentID string `json:"environmentId"`
	// Project ID associated with the session (if available).
	ProjectID string `json:"projectId"`
	// Runner ID associated with the session.
	RunnerID string `json:"runnerId"`
	// Time when the session was stopped.
	StoppedAt time.Time `json:"stoppedAt" format:"date-time"`
	// User ID that created the session.
	UserID string                 `json:"userId"`
	JSON   environmentSessionJSON `json:"-"`
}

// environmentSessionJSON contains the JSON metadata for the struct
// [EnvironmentSession]
type environmentSessionJSON struct {
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

func (r *EnvironmentSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSessionJSON) RawJSON() string {
	return r.raw
}

type UsageListEnvironmentSessionsParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// Filter options.
	Filter param.Field[UsageListEnvironmentSessionsParamsFilter] `json:"filter"`
	// Pagination options.
	Pagination param.Field[UsageListEnvironmentSessionsParamsPagination] `json:"pagination"`
}

func (r UsageListEnvironmentSessionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [UsageListEnvironmentSessionsParams]'s query parameters as
// `url.Values`.
func (r UsageListEnvironmentSessionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter options.
type UsageListEnvironmentSessionsParamsFilter struct {
	// Date range to query sessions within.
	DateRange param.Field[UsageListEnvironmentSessionsParamsFilterDateRange] `json:"dateRange,required"`
	// Optional project ID to filter sessions by.
	ProjectID param.Field[string] `json:"projectId"`
}

func (r UsageListEnvironmentSessionsParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Date range to query sessions within.
type UsageListEnvironmentSessionsParamsFilterDateRange struct {
	// End time of the date range (exclusive).
	EndTime param.Field[time.Time] `json:"endTime,required" format:"date-time"`
	// Start time of the date range (inclusive).
	StartTime param.Field[time.Time] `json:"startTime,required" format:"date-time"`
}

func (r UsageListEnvironmentSessionsParamsFilterDateRange) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Pagination options.
type UsageListEnvironmentSessionsParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r UsageListEnvironmentSessionsParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
