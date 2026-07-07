// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
)

// UsageService provides usage information about environments, users, and projects.
//
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

// Gets a summary of adoption and usage metrics.
//
// Returns all scalar values, trends, and a sparkline for the Adoption & Usage
// insight category. For full-resolution time series, use the individual time
// series RPCs.
//
// Use this method to:
//
// - Build adoption and usage insight cards
// - Filter adoption metrics by project, user, or team
// - Compare the requested date range against the previous period
//
// ### Example
//
// ```yaml
// dateRange:
//
//	startTime: "2024-01-01T00:00:00Z"
//	endTime: "2024-02-01T00:00:00Z"
//
// projectId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
// ```
func (r *UsageService) GetAdoptionUsageSummary(ctx context.Context, body UsageGetAdoptionUsageSummaryParams, opts ...option.RequestOption) (res *UsageGetAdoptionUsageSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.UsageService/GetAdoptionUsageSummary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets aggregated agent trace summary for the organization or a specific project.
//
// Use this method to:
//
// - Measure agent sessions and line changes
// - Break down agent activity by model
// - Scope agent trace insights to a project, user, or team
//
// ### Example
//
// ```yaml
// dateRange:
//
//	startTime: "2024-01-01T00:00:00Z"
//	endTime: "2024-02-01T00:00:00Z"
//
// projectId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
// ```
func (r *UsageService) GetAgentTraceSummary(ctx context.Context, body UsageGetAgentTraceSummaryParams, opts ...option.RequestOption) (res *UsageGetAgentTraceSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.UsageService/GetAgentTraceSummary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets agent trace data as a time series.
//
// Use this method to:
//
// - Chart agent sessions and line changes over time
// - Select hourly, daily, weekly, or monthly buckets
// - Scope agent trace insights to a project, user, or team
//
// ### Example
//
// ```yaml
// dateRange:
//
//	startTime: "2024-01-01T00:00:00Z"
//	endTime: "2024-02-01T00:00:00Z"
//
// resolution: RESOLUTION_WEEKLY
// projectId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
// ```
func (r *UsageService) GetAgentTraceTimeSeries(ctx context.Context, body UsageGetAgentTraceTimeSeriesParams, opts ...option.RequestOption) (res *UsageGetAgentTraceTimeSeriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.UsageService/GetAgentTraceTimeSeries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets aggregated co-author summary for the organization or a specific project.
//
// Use this method to:
//
// - Measure AI-assisted commits and line changes
// - Scope co-author insights to a project, user, or team
// - Compare the requested date range against the previous period
//
// ### Example
//
// ```yaml
// dateRange:
//
//	startTime: "2024-01-01T00:00:00Z"
//	endTime: "2024-02-01T00:00:00Z"
//
// projectId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
// ```
func (r *UsageService) GetCoAuthorSummary(ctx context.Context, body UsageGetCoAuthorSummaryParams, opts ...option.RequestOption) (res *UsageGetCoAuthorSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.UsageService/GetCoAuthorSummary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets co-author contribution data as a time series.
//
// Use this method to:
//
// - Chart AI-assisted commits and line changes over time
// - Select hourly, daily, weekly, or monthly buckets
// - Scope co-author insights to a project, user, or team
//
// ### Example
//
// ```yaml
// dateRange:
//
//	startTime: "2024-01-01T00:00:00Z"
//	endTime: "2024-02-01T00:00:00Z"
//
// resolution: RESOLUTION_WEEKLY
// projectId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
// ```
func (r *UsageService) GetCoAuthorTimeSeries(ctx context.Context, body UsageGetCoAuthorTimeSeriesParams, opts ...option.RequestOption) (res *UsageGetCoAuthorTimeSeriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.UsageService/GetCoAuthorTimeSeries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets aggregated PR speed summary for the organization or a specific project.
//
// Use this method to:
//
// - Measure pull request lead time and review latency
// - Calculate deployment frequency from merged pull requests
// - Scope PR speed insights to a project, user, or team
//
// ### Example
//
// ```yaml
// dateRange:
//
//	startTime: "2024-01-01T00:00:00Z"
//	endTime: "2024-02-01T00:00:00Z"
//
// projectId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
// ```
func (r *UsageService) GetPrSummary(ctx context.Context, body UsageGetPrSummaryParams, opts ...option.RequestOption) (res *UsageGetPrSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.UsageService/GetPrSummary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets PR speed metrics as a time series.
//
// Use this method to:
//
// - Chart pull request lead time, review latency, and deploy counts
// - Select hourly, daily, weekly, or monthly buckets
// - Scope PR speed insights to a project, user, or team
//
// ### Example
//
// ```yaml
// dateRange:
//
//	startTime: "2024-01-01T00:00:00Z"
//	endTime: "2024-02-01T00:00:00Z"
//
// resolution: RESOLUTION_WEEKLY
// projectId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
// ```
func (r *UsageService) GetPrTimeSeries(ctx context.Context, body UsageGetPrTimeSeriesParams, opts ...option.RequestOption) (res *UsageGetPrTimeSeriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.UsageService/GetPrTimeSeries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
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
	opts = slices.Concat(r.Options, opts)
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

// AgentTraceModelBreakdown contains stats for a single LLM model.
type AgentTraceModelBreakdown struct {
	// Lines added by sessions using this model.
	LinesAdded string `json:"linesAdded"`
	// Lines removed by sessions using this model.
	LinesRemoved string `json:"linesRemoved"`
	// The model these stats are for.
	Model SupportedModel `json:"model"`
	// Number of sessions that used this model.
	Sessions string                       `json:"sessions"`
	JSON     agentTraceModelBreakdownJSON `json:"-"`
}

// agentTraceModelBreakdownJSON contains the JSON metadata for the struct
// [AgentTraceModelBreakdown]
type agentTraceModelBreakdownJSON struct {
	LinesAdded   apijson.Field
	LinesRemoved apijson.Field
	Model        apijson.Field
	Sessions     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AgentTraceModelBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentTraceModelBreakdownJSON) RawJSON() string {
	return r.raw
}

// AgentTraceSummary contains aggregate totals for a date range.
type AgentTraceSummary struct {
	// Per-model breakdown of session stats.
	ByModel []AgentTraceModelBreakdown `json:"byModel"`
	// Total lines added across all sessions.
	TotalLinesAdded string `json:"totalLinesAdded"`
	// Fractional change in total_lines_added compared to the previous period.
	TotalLinesAddedTrend float64 `json:"totalLinesAddedTrend"`
	// Total lines removed across all sessions.
	TotalLinesRemoved string `json:"totalLinesRemoved"`
	// Fractional change in total_lines_removed compared to the previous period.
	TotalLinesRemovedTrend float64 `json:"totalLinesRemovedTrend"`
	// Total number of agent trace sessions in the date range.
	TotalSessions string `json:"totalSessions"`
	// Fractional change in total_sessions compared to the previous period of equal
	// length. Computed as (current - previous) / previous. Zero when there is no
	// previous data.
	TotalSessionsTrend float64               `json:"totalSessionsTrend"`
	JSON               agentTraceSummaryJSON `json:"-"`
}

// agentTraceSummaryJSON contains the JSON metadata for the struct
// [AgentTraceSummary]
type agentTraceSummaryJSON struct {
	ByModel                apijson.Field
	TotalLinesAdded        apijson.Field
	TotalLinesAddedTrend   apijson.Field
	TotalLinesRemoved      apijson.Field
	TotalLinesRemovedTrend apijson.Field
	TotalSessions          apijson.Field
	TotalSessionsTrend     apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AgentTraceSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentTraceSummaryJSON) RawJSON() string {
	return r.raw
}

// AgentTraceTimeBucket contains stats for a single time period.
type AgentTraceTimeBucket struct {
	// Per-model breakdown for this bucket.
	ByModel []AgentTraceModelBreakdown `json:"byModel"`
	// Start of this time bucket.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// Total lines added in this bucket.
	TotalLinesAdded string `json:"totalLinesAdded"`
	// Total lines removed in this bucket.
	TotalLinesRemoved string `json:"totalLinesRemoved"`
	// Number of agent trace sessions in this bucket.
	TotalSessions string                   `json:"totalSessions"`
	JSON          agentTraceTimeBucketJSON `json:"-"`
}

// agentTraceTimeBucketJSON contains the JSON metadata for the struct
// [AgentTraceTimeBucket]
type agentTraceTimeBucketJSON struct {
	ByModel           apijson.Field
	StartTime         apijson.Field
	TotalLinesAdded   apijson.Field
	TotalLinesRemoved apijson.Field
	TotalSessions     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AgentTraceTimeBucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentTraceTimeBucketJSON) RawJSON() string {
	return r.raw
}

// CoAuthorSummary contains aggregate totals for a date range.
type CoAuthorSummary struct {
	// Per-tool breakdown of contribution stats.
	ByTool []ToolBreakdown `json:"byTool"`
	// Number of distinct authors (by author_hash).
	DistinctAuthors string `json:"distinctAuthors"`
	// Fractional change in distinct_authors compared to the previous period.
	DistinctAuthorsTrend float64 `json:"distinctAuthorsTrend"`
	// Total number of commits in the date range.
	TotalCommits string `json:"totalCommits"`
	// Fractional change in total_commits compared to the previous period of equal
	// length. Computed as (current - previous) / previous. Zero when there is no
	// previous data.
	TotalCommitsTrend float64 `json:"totalCommitsTrend"`
	// Total lines added across all commits.
	TotalLinesAdded string `json:"totalLinesAdded"`
	// Fractional change in total_lines_added compared to the previous period.
	TotalLinesAddedTrend float64 `json:"totalLinesAddedTrend"`
	// Total lines removed across all commits.
	TotalLinesRemoved string `json:"totalLinesRemoved"`
	// Fractional change in total_lines_removed compared to the previous period.
	TotalLinesRemovedTrend float64             `json:"totalLinesRemovedTrend"`
	JSON                   coAuthorSummaryJSON `json:"-"`
}

// coAuthorSummaryJSON contains the JSON metadata for the struct [CoAuthorSummary]
type coAuthorSummaryJSON struct {
	ByTool                 apijson.Field
	DistinctAuthors        apijson.Field
	DistinctAuthorsTrend   apijson.Field
	TotalCommits           apijson.Field
	TotalCommitsTrend      apijson.Field
	TotalLinesAdded        apijson.Field
	TotalLinesAddedTrend   apijson.Field
	TotalLinesRemoved      apijson.Field
	TotalLinesRemovedTrend apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CoAuthorSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r coAuthorSummaryJSON) RawJSON() string {
	return r.raw
}

// CoAuthorTimeBucket contains stats for a single time period.
type CoAuthorTimeBucket struct {
	// Ratio of AI-assisted lines added to total lines added (0.0–1.0).
	AIRatio float64 `json:"aiRatio"`
	// Per-tool breakdown for this bucket.
	ByTool []ToolBreakdown `json:"byTool"`
	// Number of distinct authors (by author_hash) in this bucket.
	DistinctAuthors string `json:"distinctAuthors"`
	// Start of this time bucket.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// Total number of commits in this bucket (across all tools).
	TotalCommits string `json:"totalCommits"`
	// Total lines added in this bucket (across all tools).
	TotalLinesAdded string `json:"totalLinesAdded"`
	// Total lines removed in this bucket (across all tools).
	TotalLinesRemoved string                 `json:"totalLinesRemoved"`
	JSON              coAuthorTimeBucketJSON `json:"-"`
}

// coAuthorTimeBucketJSON contains the JSON metadata for the struct
// [CoAuthorTimeBucket]
type coAuthorTimeBucketJSON struct {
	AIRatio           apijson.Field
	ByTool            apijson.Field
	DistinctAuthors   apijson.Field
	StartTime         apijson.Field
	TotalCommits      apijson.Field
	TotalLinesAdded   apijson.Field
	TotalLinesRemoved apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CoAuthorTimeBucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r coAuthorTimeBucketJSON) RawJSON() string {
	return r.raw
}

// CoAuthorTool identifies the AI tool that co-authored a commit. UNSPECIFIED (0)
// is the proto default and must not appear in reported data. Use HUMAN for commits
// with no AI co-author.
type CoAuthorTool string

const (
	CoAuthorToolUnspecified   CoAuthorTool = "CO_AUTHOR_TOOL_UNSPECIFIED"
	CoAuthorToolNoCoauthor    CoAuthorTool = "CO_AUTHOR_TOOL_NO_COAUTHOR"
	CoAuthorToolHumanCoauthor CoAuthorTool = "CO_AUTHOR_TOOL_HUMAN_COAUTHOR"
	CoAuthorToolOna           CoAuthorTool = "CO_AUTHOR_TOOL_ONA"
	CoAuthorToolGitHubCopilot CoAuthorTool = "CO_AUTHOR_TOOL_GITHUB_COPILOT"
	CoAuthorToolCursor        CoAuthorTool = "CO_AUTHOR_TOOL_CURSOR"
	CoAuthorToolOther         CoAuthorTool = "CO_AUTHOR_TOOL_OTHER"
	CoAuthorToolClaude        CoAuthorTool = "CO_AUTHOR_TOOL_CLAUDE"
	CoAuthorToolCodex         CoAuthorTool = "CO_AUTHOR_TOOL_CODEX"
)

func (r CoAuthorTool) IsKnown() bool {
	switch r {
	case CoAuthorToolUnspecified, CoAuthorToolNoCoauthor, CoAuthorToolHumanCoauthor, CoAuthorToolOna, CoAuthorToolGitHubCopilot, CoAuthorToolCursor, CoAuthorToolOther, CoAuthorToolClaude, CoAuthorToolCodex:
		return true
	}
	return false
}

// DateRange specifies a time period for queries.
type DateRangeParam struct {
	// End time of the date range (exclusive).
	EndTime param.Field[time.Time] `json:"endTime" api:"required" format:"date-time"`
	// Start time of the date range (inclusive).
	StartTime param.Field[time.Time] `json:"startTime" api:"required" format:"date-time"`
}

func (r DateRangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// PrSummary contains aggregate PR speed metrics for a date range.
type PrSummary struct {
	// PRs merged to the default branch per week.
	DeploymentFrequency float64 `json:"deploymentFrequency"`
	// Fractional change in deployment_frequency vs previous period. Computed as
	// (current - previous) / previous.
	DeploymentFrequencyTrend float64 `json:"deploymentFrequencyTrend"`
	// Median lead time for changes in seconds (first commit → merge).
	LeadTimeSeconds float64 `json:"leadTimeSeconds"`
	// Fractional change in lead_time_seconds vs previous period. Computed as
	// (current - previous) / previous.
	LeadTimeTrend float64 `json:"leadTimeTrend"`
	// Total PRs merged in the date range.
	PrsMergedCount string `json:"prsMergedCount"`
	// Fractional change in prs_merged_count vs previous period. Computed as (current -
	// previous) / previous.
	PrsMergedTrend float64 `json:"prsMergedTrend"`
	// Median time to first approval in seconds. Zero when no PRs in the range had
	// approvals.
	TimeToFirstApprovalSeconds float64 `json:"timeToFirstApprovalSeconds"`
	// Fractional change in time_to_first_approval_seconds vs previous period. Computed
	// as (current - previous) / previous.
	TimeToFirstApprovalTrend float64       `json:"timeToFirstApprovalTrend"`
	JSON                     prSummaryJSON `json:"-"`
}

// prSummaryJSON contains the JSON metadata for the struct [PrSummary]
type prSummaryJSON struct {
	DeploymentFrequency        apijson.Field
	DeploymentFrequencyTrend   apijson.Field
	LeadTimeSeconds            apijson.Field
	LeadTimeTrend              apijson.Field
	PrsMergedCount             apijson.Field
	PrsMergedTrend             apijson.Field
	TimeToFirstApprovalSeconds apijson.Field
	TimeToFirstApprovalTrend   apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *PrSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prSummaryJSON) RawJSON() string {
	return r.raw
}

// PrTimeBucket contains PR speed metrics for a single time period.
type PrTimeBucket struct {
	// Total number of deploys (merged PRs) in this bucket.
	Deploys string `json:"deploys"`
	// Median lead time in seconds for PRs merged in this bucket.
	LeadTimeSeconds float64 `json:"leadTimeSeconds"`
	// Number of PRs merged in this bucket.
	PrsMergedCount string `json:"prsMergedCount"`
	// Start of this time bucket.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// Median time to first approval in seconds for PRs in this bucket. Zero when no
	// PRs in the bucket had approvals.
	TimeToFirstApprovalSeconds float64          `json:"timeToFirstApprovalSeconds"`
	JSON                       prTimeBucketJSON `json:"-"`
}

// prTimeBucketJSON contains the JSON metadata for the struct [PrTimeBucket]
type prTimeBucketJSON struct {
	Deploys                    apijson.Field
	LeadTimeSeconds            apijson.Field
	PrsMergedCount             apijson.Field
	StartTime                  apijson.Field
	TimeToFirstApprovalSeconds apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *PrTimeBucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prTimeBucketJSON) RawJSON() string {
	return r.raw
}

// Resolution specifies the time granularity for time series data.
type Resolution string

const (
	ResolutionUnspecified Resolution = "RESOLUTION_UNSPECIFIED"
	ResolutionHourly      Resolution = "RESOLUTION_HOURLY"
	ResolutionDaily       Resolution = "RESOLUTION_DAILY"
	ResolutionWeekly      Resolution = "RESOLUTION_WEEKLY"
	ResolutionMonthly     Resolution = "RESOLUTION_MONTHLY"
)

func (r Resolution) IsKnown() bool {
	switch r {
	case ResolutionUnspecified, ResolutionHourly, ResolutionDaily, ResolutionWeekly, ResolutionMonthly:
		return true
	}
	return false
}

// SupportedModel enumerates the LLM models available for agent executions
type SupportedModel string

const (
	SupportedModelUnspecified       SupportedModel = "SUPPORTED_MODEL_UNSPECIFIED"
	SupportedModelSonnet3_5         SupportedModel = "SUPPORTED_MODEL_SONNET_3_5"
	SupportedModelSonnet3_7         SupportedModel = "SUPPORTED_MODEL_SONNET_3_7"
	SupportedModelSonnet3_7Extended SupportedModel = "SUPPORTED_MODEL_SONNET_3_7_EXTENDED"
	SupportedModelSonnet4           SupportedModel = "SUPPORTED_MODEL_SONNET_4"
	SupportedModelSonnet4Extended   SupportedModel = "SUPPORTED_MODEL_SONNET_4_EXTENDED"
	SupportedModelSonnet4_5         SupportedModel = "SUPPORTED_MODEL_SONNET_4_5"
	SupportedModelSonnet4_5Extended SupportedModel = "SUPPORTED_MODEL_SONNET_4_5_EXTENDED"
	SupportedModelSonnet4_6         SupportedModel = "SUPPORTED_MODEL_SONNET_4_6"
	SupportedModelSonnet4_6Extended SupportedModel = "SUPPORTED_MODEL_SONNET_4_6_EXTENDED"
	SupportedModelSonnet5           SupportedModel = "SUPPORTED_MODEL_SONNET_5"
	SupportedModelOpus4             SupportedModel = "SUPPORTED_MODEL_OPUS_4"
	SupportedModelOpus4Extended     SupportedModel = "SUPPORTED_MODEL_OPUS_4_EXTENDED"
	SupportedModelOpus4_5           SupportedModel = "SUPPORTED_MODEL_OPUS_4_5"
	SupportedModelOpus4_5Extended   SupportedModel = "SUPPORTED_MODEL_OPUS_4_5_EXTENDED"
	SupportedModelOpus4_6           SupportedModel = "SUPPORTED_MODEL_OPUS_4_6"
	SupportedModelOpus4_6Extended   SupportedModel = "SUPPORTED_MODEL_OPUS_4_6_EXTENDED"
	SupportedModelOpus4_7           SupportedModel = "SUPPORTED_MODEL_OPUS_4_7"
	SupportedModelOpus4_8           SupportedModel = "SUPPORTED_MODEL_OPUS_4_8"
	SupportedModelHaiku4_5          SupportedModel = "SUPPORTED_MODEL_HAIKU_4_5"
	SupportedModelOpenAI4O          SupportedModel = "SUPPORTED_MODEL_OPENAI_4O"
	SupportedModelOpenAI4OMini      SupportedModel = "SUPPORTED_MODEL_OPENAI_4O_MINI"
	SupportedModelOpenAIO1          SupportedModel = "SUPPORTED_MODEL_OPENAI_O1"
	SupportedModelOpenAIO1Mini      SupportedModel = "SUPPORTED_MODEL_OPENAI_O1_MINI"
	SupportedModelOpenAIAuto        SupportedModel = "SUPPORTED_MODEL_OPENAI_AUTO"
)

func (r SupportedModel) IsKnown() bool {
	switch r {
	case SupportedModelUnspecified, SupportedModelSonnet3_5, SupportedModelSonnet3_7, SupportedModelSonnet3_7Extended, SupportedModelSonnet4, SupportedModelSonnet4Extended, SupportedModelSonnet4_5, SupportedModelSonnet4_5Extended, SupportedModelSonnet4_6, SupportedModelSonnet4_6Extended, SupportedModelSonnet5, SupportedModelOpus4, SupportedModelOpus4Extended, SupportedModelOpus4_5, SupportedModelOpus4_5Extended, SupportedModelOpus4_6, SupportedModelOpus4_6Extended, SupportedModelOpus4_7, SupportedModelOpus4_8, SupportedModelHaiku4_5, SupportedModelOpenAI4O, SupportedModelOpenAI4OMini, SupportedModelOpenAIO1, SupportedModelOpenAIO1Mini, SupportedModelOpenAIAuto:
		return true
	}
	return false
}

type TimeSeriesPoint struct {
	// Timestamp for this data point.
	Time time.Time `json:"time" format:"date-time"`
	// The numerical value for this data point.
	Value int64               `json:"value"`
	JSON  timeSeriesPointJSON `json:"-"`
}

// timeSeriesPointJSON contains the JSON metadata for the struct [TimeSeriesPoint]
type timeSeriesPointJSON struct {
	Time        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimeSeriesPoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timeSeriesPointJSON) RawJSON() string {
	return r.raw
}

// ToolBreakdown contains stats for a single AI tool (or human).
type ToolBreakdown struct {
	// Number of commits attributed to this tool.
	Commits string `json:"commits"`
	// Distinct authors who used this tool.
	DistinctAuthors string `json:"distinctAuthors"`
	// Lines added by this tool.
	LinesAdded string `json:"linesAdded"`
	// Lines removed by this tool.
	LinesRemoved string `json:"linesRemoved"`
	// The tool these stats are for.
	Tool CoAuthorTool      `json:"tool"`
	JSON toolBreakdownJSON `json:"-"`
}

// toolBreakdownJSON contains the JSON metadata for the struct [ToolBreakdown]
type toolBreakdownJSON struct {
	Commits         apijson.Field
	DistinctAuthors apijson.Field
	LinesAdded      apijson.Field
	LinesRemoved    apijson.Field
	Tool            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ToolBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolBreakdownJSON) RawJSON() string {
	return r.raw
}

type UsageGetAdoptionUsageSummaryResponse struct {
	// Count of active users in the date range.
	ActiveUsersCount string `json:"activeUsersCount"`
	// Fractional change in active_users_count vs previous period. Computed as
	// (current - previous) / previous.
	ActiveUsersTrend float64 `json:"activeUsersTrend"`
	// Average environment runtime in seconds per active user.
	EnvRuntimePerUserSeconds float64 `json:"envRuntimePerUserSeconds"`
	// Fractional change in env_runtime_per_user_seconds vs previous period. Computed
	// as (current - previous) / previous.
	EnvRuntimePerUserTrend float64 `json:"envRuntimePerUserTrend"`
	// Count of power users in the date range.
	PowerUsersCount string `json:"powerUsersCount"`
	// Threshold in seconds used to determine power users. Displayed to users so they
	// understand the definition.
	PowerUsersThresholdSeconds string `json:"powerUsersThresholdSeconds"`
	// Fractional change in power_users_count vs previous period. Computed as
	// (current - previous) / previous.
	PowerUsersTrend float64 `json:"powerUsersTrend"`
	// Count of environment sessions (total starts) in the date range.
	SessionsCount string `json:"sessionsCount"`
	// Fractional change in sessions_count vs previous period. Computed as (current -
	// previous) / previous.
	SessionsTrend float64 `json:"sessionsTrend"`
	// Sparkline data for the card's trend line (typically ~4 weekly points).
	Sparkline []TimeSeriesPoint                        `json:"sparkline"`
	JSON      usageGetAdoptionUsageSummaryResponseJSON `json:"-"`
}

// usageGetAdoptionUsageSummaryResponseJSON contains the JSON metadata for the
// struct [UsageGetAdoptionUsageSummaryResponse]
type usageGetAdoptionUsageSummaryResponseJSON struct {
	ActiveUsersCount           apijson.Field
	ActiveUsersTrend           apijson.Field
	EnvRuntimePerUserSeconds   apijson.Field
	EnvRuntimePerUserTrend     apijson.Field
	PowerUsersCount            apijson.Field
	PowerUsersThresholdSeconds apijson.Field
	PowerUsersTrend            apijson.Field
	SessionsCount              apijson.Field
	SessionsTrend              apijson.Field
	Sparkline                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UsageGetAdoptionUsageSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAdoptionUsageSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetAgentTraceSummaryResponse struct {
	// Sparkline data for card rendering.
	Sparkline []TimeSeriesPoint `json:"sparkline"`
	// Summary totals and trends for the requested date range.
	Summary AgentTraceSummary                     `json:"summary"`
	JSON    usageGetAgentTraceSummaryResponseJSON `json:"-"`
}

// usageGetAgentTraceSummaryResponseJSON contains the JSON metadata for the struct
// [UsageGetAgentTraceSummaryResponse]
type usageGetAgentTraceSummaryResponseJSON struct {
	Sparkline   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetAgentTraceSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAgentTraceSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetAgentTraceTimeSeriesResponse struct {
	// Time series of agent trace stats, bucketed by the requested resolution.
	TimeSeries []AgentTraceTimeBucket                   `json:"timeSeries"`
	JSON       usageGetAgentTraceTimeSeriesResponseJSON `json:"-"`
}

// usageGetAgentTraceTimeSeriesResponseJSON contains the JSON metadata for the
// struct [UsageGetAgentTraceTimeSeriesResponse]
type usageGetAgentTraceTimeSeriesResponseJSON struct {
	TimeSeries  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetAgentTraceTimeSeriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetAgentTraceTimeSeriesResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetCoAuthorSummaryResponse struct {
	// Sparkline data for card rendering.
	Sparkline []TimeSeriesPoint `json:"sparkline"`
	// Summary totals and trends for the requested date range.
	Summary CoAuthorSummary                     `json:"summary"`
	JSON    usageGetCoAuthorSummaryResponseJSON `json:"-"`
}

// usageGetCoAuthorSummaryResponseJSON contains the JSON metadata for the struct
// [UsageGetCoAuthorSummaryResponse]
type usageGetCoAuthorSummaryResponseJSON struct {
	Sparkline   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetCoAuthorSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetCoAuthorSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetCoAuthorTimeSeriesResponse struct {
	// Time series of contribution stats, bucketed by the requested resolution.
	TimeSeries []CoAuthorTimeBucket                   `json:"timeSeries"`
	JSON       usageGetCoAuthorTimeSeriesResponseJSON `json:"-"`
}

// usageGetCoAuthorTimeSeriesResponseJSON contains the JSON metadata for the struct
// [UsageGetCoAuthorTimeSeriesResponse]
type usageGetCoAuthorTimeSeriesResponseJSON struct {
	TimeSeries  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetCoAuthorTimeSeriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetCoAuthorTimeSeriesResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetPrSummaryResponse struct {
	// Sparkline data for card rendering.
	Sparkline []TimeSeriesPoint `json:"sparkline"`
	// Summary totals and trends for the requested date range.
	Summary PrSummary                     `json:"summary"`
	JSON    usageGetPrSummaryResponseJSON `json:"-"`
}

// usageGetPrSummaryResponseJSON contains the JSON metadata for the struct
// [UsageGetPrSummaryResponse]
type usageGetPrSummaryResponseJSON struct {
	Sparkline   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetPrSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetPrSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetPrTimeSeriesResponse struct {
	// Time series of PR speed metrics, bucketed by the requested resolution.
	TimeSeries []PrTimeBucket                   `json:"timeSeries"`
	JSON       usageGetPrTimeSeriesResponseJSON `json:"-"`
}

// usageGetPrTimeSeriesResponseJSON contains the JSON metadata for the struct
// [UsageGetPrTimeSeriesResponse]
type usageGetPrTimeSeriesResponseJSON struct {
	TimeSeries  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetPrTimeSeriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetPrTimeSeriesResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetAdoptionUsageSummaryParams struct {
	// Date range to query metrics within.
	DateRange param.Field[DateRangeParam] `json:"dateRange" api:"required"`
	// Optional project ID to filter metrics by.
	ProjectID param.Field[string] `json:"projectId"`
	// Optional team ID to scope results to members of a specific team.
	TeamID param.Field[string] `json:"teamId"`
	// Optional user ID to filter metrics for a specific user (personal insights view).
	UserID param.Field[string] `json:"userId"`
}

func (r UsageGetAdoptionUsageSummaryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageGetAgentTraceSummaryParams struct {
	// Date range to query within.
	DateRange param.Field[DateRangeParam] `json:"dateRange" api:"required"`
	// Optional project ID to scope results.
	ProjectID param.Field[string] `json:"projectId"`
	// Optional team ID to scope results to a specific team. Mutually exclusive with
	// user_id.
	TeamID param.Field[string] `json:"teamId"`
	// Optional user ID to scope results to a specific user. Mutually exclusive with
	// team_id.
	UserID param.Field[string] `json:"userId"`
}

func (r UsageGetAgentTraceSummaryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageGetAgentTraceTimeSeriesParams struct {
	// Date range to query within.
	DateRange param.Field[DateRangeParam] `json:"dateRange" api:"required"`
	// Optional project ID to scope results.
	ProjectID param.Field[string] `json:"projectId"`
	// Time resolution for the series data.
	Resolution param.Field[Resolution] `json:"resolution"`
	// Optional team ID to scope results to a specific team. Mutually exclusive with
	// user_id.
	TeamID param.Field[string] `json:"teamId"`
	// Optional user ID to scope results to a specific user. Mutually exclusive with
	// team_id.
	UserID param.Field[string] `json:"userId"`
}

func (r UsageGetAgentTraceTimeSeriesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageGetCoAuthorSummaryParams struct {
	// Date range to query within.
	DateRange param.Field[DateRangeParam] `json:"dateRange" api:"required"`
	// Optional project ID to scope results.
	ProjectID param.Field[string] `json:"projectId"`
	// Optional team ID to scope results to a specific team. Mutually exclusive with
	// user_id.
	TeamID param.Field[string] `json:"teamId"`
	// Optional user ID to scope results to a specific user. Mutually exclusive with
	// team_id.
	UserID param.Field[string] `json:"userId"`
}

func (r UsageGetCoAuthorSummaryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageGetCoAuthorTimeSeriesParams struct {
	// Date range to query within.
	DateRange param.Field[DateRangeParam] `json:"dateRange" api:"required"`
	// Optional project ID to scope results.
	ProjectID param.Field[string] `json:"projectId"`
	// Time resolution for the series data.
	Resolution param.Field[Resolution] `json:"resolution"`
	// Optional team ID to scope results to a specific team. Mutually exclusive with
	// user_id.
	TeamID param.Field[string] `json:"teamId"`
	// Optional user ID to scope results to a specific user. Mutually exclusive with
	// team_id.
	UserID param.Field[string] `json:"userId"`
}

func (r UsageGetCoAuthorTimeSeriesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageGetPrSummaryParams struct {
	// Date range to query within.
	DateRange param.Field[DateRangeParam] `json:"dateRange" api:"required"`
	// Optional project ID to scope results.
	ProjectID param.Field[string] `json:"projectId"`
	// Optional team ID to scope results to a specific team. Mutually exclusive with
	// user_id.
	TeamID param.Field[string] `json:"teamId"`
	// Optional user ID to scope results to a specific user. Mutually exclusive with
	// team_id.
	UserID param.Field[string] `json:"userId"`
}

func (r UsageGetPrSummaryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageGetPrTimeSeriesParams struct {
	// Date range to query within.
	DateRange param.Field[DateRangeParam] `json:"dateRange" api:"required"`
	// Optional project ID to scope results.
	ProjectID param.Field[string] `json:"projectId"`
	// Time resolution for the series data.
	Resolution param.Field[Resolution] `json:"resolution"`
	// Optional team ID to scope results to a specific team. Mutually exclusive with
	// user_id.
	TeamID param.Field[string] `json:"teamId"`
	// Optional user ID to scope results to a specific user. Mutually exclusive with
	// team_id.
	UserID param.Field[string] `json:"userId"`
}

func (r UsageGetPrTimeSeriesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	DateRange param.Field[DateRangeParam] `json:"dateRange" api:"required"`
	// Optional project ID to filter runtime records by.
	ProjectID param.Field[string] `json:"projectId"`
}

func (r UsageListEnvironmentRuntimeRecordsParamsFilter) MarshalJSON() (data []byte, err error) {
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
