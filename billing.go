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
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// BillingService provides billing and subscription management functionality.
//
// BillingService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingService] method instead.
type BillingService struct {
	Options []option.RequestOption
}

// NewBillingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBillingService(opts ...option.RequestOption) (r *BillingService) {
	r = &BillingService{}
	r.Options = opts
	return
}

// Returns a signed download URL for a CSV export of credit usage.
//
// The URL points to an HTTP endpoint that streams gzip-compressed CSV and is valid
// for five minutes. The download must be made by the same principal that requested
// it, carrying its own bearer token. The export range may cover up to a year.
//
// For organizations without enterprise credit usage enabled (no billing contract
// start date), the export instead contains BYOK cost usage with a different column
// set, and groupBy=RESOURCE is rejected.
//
// Use this method to:
//
// - Export per-user daily credit usage for external reporting
// - Export a per-environment and per-conversation resource breakdown
//
// ### Examples
//
// - Export January's daily summary:
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	dateRange:
//	  startTime: "2024-01-01T00:00:00Z"
//	  endTime: "2024-01-31T00:00:00Z"
//	groupBy: CREDIT_USAGE_EXPORT_GROUP_BY_DAILY_SUMMARY
//	```
//
// ### Authorization
//
// Requires `billing:read_usage` permission on the organization.
func (r *BillingService) GetCreditUsageExport(ctx context.Context, body BillingGetCreditUsageExportParams, opts ...option.RequestOption) (res *BillingGetCreditUsageExportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.BillingService/GetCreditUsageExport"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a daily credit usage report for an enterprise organization.
//
// Each day reports org-wide credits by usage type, plus per-user, per-team,
// per-environment, and per-conversation breakdowns (top consumers with the
// remainder aggregated into an "Others" bucket) and a per-model breakdown of
// intelligence usage.
//
// Use this method to:
//
// - Chart daily credit consumption over a date range
// - Attribute credit usage to users, teams, environments, and conversations
// - Restrict the report to a single user or service account
//
// ### Examples
//
// - Get the report for January:
//
//	Both dates are inclusive and the range must not exceed 31 days.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	dateRange:
//	  startTime: "2024-01-01T00:00:00Z"
//	  endTime: "2024-01-31T00:00:00Z"
//	```
//
// ### Authorization
//
// Requires `billing:read_usage` permission on the organization. A user without it
// can read their own usage by setting filter.subject to their own user identity;
// this self-access path is not available to service accounts.
func (r *BillingService) GetCreditUsageReport(ctx context.Context, body BillingGetCreditUsageReportParams, opts ...option.RequestOption) (res *BillingGetCreditUsageReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.BillingService/GetCreditUsageReport"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns cumulative credit usage for an organization and its teams.
//
// Use this method to:
//
// - Get the total cumulative credit consumption as of a point in time
// - Get per-team cumulative usage with credit allocation (budget) comparison
// - Display team credit summaries on the usage page and team detail page
// - Display user budget utilization when user budgets are enabled
//
// ### Examples
//
// - Get current cumulative usage:
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
//
// - Get cumulative usage as of a specific date:
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	asOf: "2026-03-31T23:59:59Z"
//	```
//
// ### Authorization
//
// Requires `billing:read_usage` permission on the organization.
func (r *BillingService) GetCumulativeCreditUsage(ctx context.Context, body BillingGetCumulativeCreditUsageParams, opts ...option.RequestOption) (res *BillingGetCumulativeCreditUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.BillingService/GetCumulativeCreditUsage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns organization-level enterprise AI usage totals for reporting.
//
// Reports BYOK (bring-your-own-key) token spend: cost in the organization's
// billing currency plus token counts, with a per-model breakdown. Credit-based
// usage from managed models is not included and the credits field is not populated
// by this endpoint.
//
// Use this method to:
//
// - Report total BYOK AI spend (cost and tokens) for a date range
// - Break down organization usage by model
//
// Only available for enterprise organizations.
//
// ### Examples
//
// - Get usage totals for January:
//
//	Returns organization-wide BYOK spend for the month. Both dates are inclusive
//	and the range must not exceed 31 days.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	dateRange:
//	  startTime: "2024-01-01T00:00:00Z"
//	  endTime: "2024-01-31T00:00:00Z"
//	```
//
// ### Authorization
//
// Requires `billing:read_usage` permission on the organization.
func (r *BillingService) GetEnterpriseAIUsageSummary(ctx context.Context, body BillingGetEnterpriseAIUsageSummaryParams, opts ...option.RequestOption) (res *BillingGetEnterpriseAIUsageSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.BillingService/GetEnterpriseAIUsageSummary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns daily enterprise AI usage totals for the organization.
//
// Each day reports BYOK token spend (cost and tokens) with per-user, per-team, and
// per-model breakdowns. Per-user entries cover the top spenders with the remainder
// aggregated into an "Others" bucket; usage not attributed to a user or service
// account appears only in the daily totals. The credits field is not populated by
// this endpoint.
//
// When filter.subject is set the response contains only that subject's usage:
// daily totals and the team breakdown are omitted, and the model breakdown covers
// the subject only.
//
// Use this method to:
//
// - Chart daily BYOK AI spend over a date range
// - Feed daily per-user usage into external dashboards
// - Restrict the response to a single user or service account
//
// Only available for enterprise organizations.
//
// ### Examples
//
// - Get daily usage for January:
//
//	Returns one entry per day with per-user, per-team, and per-model breakdowns.
//	Both dates are inclusive and the range must not exceed 31 days.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	dateRange:
//	  startTime: "2024-01-01T00:00:00Z"
//	  endTime: "2024-01-31T00:00:00Z"
//	```
//
// ### Authorization
//
// Requires `billing:read_usage` permission on the organization.
func (r *BillingService) GetEnterpriseAIUsageTimeSeries(ctx context.Context, body BillingGetEnterpriseAIUsageTimeSeriesParams, opts ...option.RequestOption) (res *BillingGetEnterpriseAIUsageTimeSeriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.BillingService/GetEnterpriseAIUsageTimeSeries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists enterprise AI usage grouped by team.
//
// Reports BYOK token spend per team (cost and tokens) with each team's monthly
// budget when one applies. The credits field is not populated by this endpoint.
//
// Use this method to:
//
// - Compare BYOK AI spend across teams
// - Track team budget utilization
// - Filter usage to specific teams
//
// Only available for enterprise organizations.
//
// ### Examples
//
// - List team usage for January:
//
//	Returns BYOK spend per team with monthly budgets. Both dates are inclusive and
//	the range must not exceed 31 days.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	dateRange:
//	  startTime: "2024-01-01T00:00:00Z"
//	  endTime: "2024-01-31T00:00:00Z"
//	```
//
// ### Authorization
//
// Requires `billing:read_usage` permission on the organization.
func (r *BillingService) ListEnterpriseAITeamUsage(ctx context.Context, params BillingListEnterpriseAITeamUsageParams, opts ...option.RequestOption) (res *pagination.TeamUsagePage[TeamEnterpriseAIUsage], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.BillingService/ListEnterpriseAITeamUsage"
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

// Lists enterprise AI usage grouped by team.
//
// Reports BYOK token spend per team (cost and tokens) with each team's monthly
// budget when one applies. The credits field is not populated by this endpoint.
//
// Use this method to:
//
// - Compare BYOK AI spend across teams
// - Track team budget utilization
// - Filter usage to specific teams
//
// Only available for enterprise organizations.
//
// ### Examples
//
// - List team usage for January:
//
//	Returns BYOK spend per team with monthly budgets. Both dates are inclusive and
//	the range must not exceed 31 days.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	dateRange:
//	  startTime: "2024-01-01T00:00:00Z"
//	  endTime: "2024-01-31T00:00:00Z"
//	```
//
// ### Authorization
//
// Requires `billing:read_usage` permission on the organization.
func (r *BillingService) ListEnterpriseAITeamUsageAutoPaging(ctx context.Context, params BillingListEnterpriseAITeamUsageParams, opts ...option.RequestOption) *pagination.TeamUsagePageAutoPager[TeamEnterpriseAIUsage] {
	return pagination.NewTeamUsagePageAutoPager(r.ListEnterpriseAITeamUsage(ctx, params, opts...))
}

// Lists enterprise AI usage grouped by user with effective monthly budget data.
//
// Reports BYOK token spend (cost and tokens) for each user and service account
// with attributed usage in the date range, including each subject's effective
// monthly budget. Usage not attributed to a user or service account is excluded,
// so the sum across subjects can be less than the organization totals from
// GetEnterpriseAIUsageSummary. The credits field is not populated by this
// endpoint.
//
// Budget fields (month_to_date_usage, utilization_percent, over_budget) are
// computed from usage inside the requested date range measured against the monthly
// limit. Send a range that starts on the first day of the month for true
// month-to-date figures.
//
// Use this method to:
//
// - Export per-user BYOK AI spend to external reporting
// - Identify the highest spenders in the organization
// - Track per-user budget utilization and over-budget users
//
// Only available for enterprise organizations.
//
// ### Examples
//
// - List user usage for January:
//
//	Returns per-user BYOK spend with effective budgets, highest spend first. Both
//	dates are inclusive and the range must not exceed 31 days.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	dateRange:
//	  startTime: "2024-01-01T00:00:00Z"
//	  endTime: "2024-01-31T00:00:00Z"
//	```
//
// ### Authorization
//
// Requires `billing:read_usage` permission on the organization. Callers without it
// can read their own usage by setting filter.subject to themselves.
func (r *BillingService) ListEnterpriseAIUserUsage(ctx context.Context, params BillingListEnterpriseAIUserUsageParams, opts ...option.RequestOption) (res *pagination.UserUsagePage[UserCostBudgetUsage], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.BillingService/ListEnterpriseAIUserUsage"
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

// Lists enterprise AI usage grouped by user with effective monthly budget data.
//
// Reports BYOK token spend (cost and tokens) for each user and service account
// with attributed usage in the date range, including each subject's effective
// monthly budget. Usage not attributed to a user or service account is excluded,
// so the sum across subjects can be less than the organization totals from
// GetEnterpriseAIUsageSummary. The credits field is not populated by this
// endpoint.
//
// Budget fields (month_to_date_usage, utilization_percent, over_budget) are
// computed from usage inside the requested date range measured against the monthly
// limit. Send a range that starts on the first day of the month for true
// month-to-date figures.
//
// Use this method to:
//
// - Export per-user BYOK AI spend to external reporting
// - Identify the highest spenders in the organization
// - Track per-user budget utilization and over-budget users
//
// Only available for enterprise organizations.
//
// ### Examples
//
// - List user usage for January:
//
//	Returns per-user BYOK spend with effective budgets, highest spend first. Both
//	dates are inclusive and the range must not exceed 31 days.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	dateRange:
//	  startTime: "2024-01-01T00:00:00Z"
//	  endTime: "2024-01-31T00:00:00Z"
//	```
//
// ### Authorization
//
// Requires `billing:read_usage` permission on the organization. Callers without it
// can read their own usage by setting filter.subject to themselves.
func (r *BillingService) ListEnterpriseAIUserUsageAutoPaging(ctx context.Context, params BillingListEnterpriseAIUserUsageParams, opts ...option.RequestOption) *pagination.UserUsagePageAutoPager[UserCostBudgetUsage] {
	return pagination.NewUserUsagePageAutoPager(r.ListEnterpriseAIUserUsage(ctx, params, opts...))
}

// Lists per-user month-to-date credit usage with effective monthly budgets.
//
// Results are ordered by total credits descending so the highest spenders appear
// first, with user_id as a stable tiebreaker. Use cursor pagination to walk the
// full set for large organizations.
//
// The default SORT_FIELD_USAGE ordering supports cursor pagination over any number
// of users. Sorting by display name, budget, or budget utilization computes the
// order in memory and is limited to organizations with at most 10,000 users;
// beyond that, use SORT_FIELD_USAGE. Because month-to-date figures are recomputed
// per request, hold a date range stable across a paginated walk to keep page
// tokens valid.
//
// Use this method to:
//
// - Export per-user credit usage to external reporting
// - Identify the highest spenders in the organization
// - Track per-user budget utilization and over-budget users
//
// ### Examples
//
// - List user usage for the current month:
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 50
//	```
//
// ### Authorization
//
// Requires `billing:read_usage` permission on the organization.
func (r *BillingService) ListEnterpriseUserCreditUsage(ctx context.Context, params BillingListEnterpriseUserCreditUsageParams, opts ...option.RequestOption) (res *pagination.UserUsagePage[UserCreditBudgetUsage], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.BillingService/ListEnterpriseUserCreditUsage"
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

// Lists per-user month-to-date credit usage with effective monthly budgets.
//
// Results are ordered by total credits descending so the highest spenders appear
// first, with user_id as a stable tiebreaker. Use cursor pagination to walk the
// full set for large organizations.
//
// The default SORT_FIELD_USAGE ordering supports cursor pagination over any number
// of users. Sorting by display name, budget, or budget utilization computes the
// order in memory and is limited to organizations with at most 10,000 users;
// beyond that, use SORT_FIELD_USAGE. Because month-to-date figures are recomputed
// per request, hold a date range stable across a paginated walk to keep page
// tokens valid.
//
// Use this method to:
//
// - Export per-user credit usage to external reporting
// - Identify the highest spenders in the organization
// - Track per-user budget utilization and over-budget users
//
// ### Examples
//
// - List user usage for the current month:
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 50
//	```
//
// ### Authorization
//
// Requires `billing:read_usage` permission on the organization.
func (r *BillingService) ListEnterpriseUserCreditUsageAutoPaging(ctx context.Context, params BillingListEnterpriseUserCreditUsageParams, opts ...option.RequestOption) *pagination.UserUsagePageAutoPager[UserCreditBudgetUsage] {
	return pagination.NewUserUsagePageAutoPager(r.ListEnterpriseUserCreditUsage(ctx, params, opts...))
}

// AgentExecutionCreditUsage contains a single agent execution's credit usage for a
// day, broken down by type.
type AgentExecutionCreditUsage struct {
	// Empty when representing the "Others" aggregation bucket.
	AgentExecutionID string                        `json:"agentExecutionId"`
	DisplayName      string                        `json:"displayName"`
	Usage            []CreditsByType               `json:"usage"`
	JSON             agentExecutionCreditUsageJSON `json:"-"`
}

// agentExecutionCreditUsageJSON contains the JSON metadata for the struct
// [AgentExecutionCreditUsage]
type agentExecutionCreditUsageJSON struct {
	AgentExecutionID apijson.Field
	DisplayName      apijson.Field
	Usage            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AgentExecutionCreditUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionCreditUsageJSON) RawJSON() string {
	return r.raw
}

type BillingCurrency string

const (
	BillingCurrencyUnspecified BillingCurrency = "BILLING_CURRENCY_UNSPECIFIED"
	BillingCurrencyUsd         BillingCurrency = "BILLING_CURRENCY_USD"
	BillingCurrencyEur         BillingCurrency = "BILLING_CURRENCY_EUR"
	BillingCurrencyGbp         BillingCurrency = "BILLING_CURRENCY_GBP"
)

func (r BillingCurrency) IsKnown() bool {
	switch r {
	case BillingCurrencyUnspecified, BillingCurrencyUsd, BillingCurrencyEur, BillingCurrencyGbp:
		return true
	}
	return false
}

type ByokRateCardTokenType string

const (
	ByokRateCardTokenTypeUnspecified ByokRateCardTokenType = "BYOK_RATE_CARD_TOKEN_TYPE_UNSPECIFIED"
	ByokRateCardTokenTypeInput       ByokRateCardTokenType = "BYOK_RATE_CARD_TOKEN_TYPE_INPUT"
	ByokRateCardTokenTypeOutput      ByokRateCardTokenType = "BYOK_RATE_CARD_TOKEN_TYPE_OUTPUT"
	ByokRateCardTokenTypeCacheRead   ByokRateCardTokenType = "BYOK_RATE_CARD_TOKEN_TYPE_CACHE_READ"
	ByokRateCardTokenTypeCacheWrite  ByokRateCardTokenType = "BYOK_RATE_CARD_TOKEN_TYPE_CACHE_WRITE"
)

func (r ByokRateCardTokenType) IsKnown() bool {
	switch r {
	case ByokRateCardTokenTypeUnspecified, ByokRateCardTokenTypeInput, ByokRateCardTokenTypeOutput, ByokRateCardTokenTypeCacheRead, ByokRateCardTokenTypeCacheWrite:
		return true
	}
	return false
}

// How to group the credit usage export data.
type CreditUsageExportGroupBy string

const (
	CreditUsageExportGroupByUnspecified  CreditUsageExportGroupBy = "CREDIT_USAGE_EXPORT_GROUP_BY_UNSPECIFIED"
	CreditUsageExportGroupByDailySummary CreditUsageExportGroupBy = "CREDIT_USAGE_EXPORT_GROUP_BY_DAILY_SUMMARY"
	CreditUsageExportGroupByResource     CreditUsageExportGroupBy = "CREDIT_USAGE_EXPORT_GROUP_BY_RESOURCE"
)

func (r CreditUsageExportGroupBy) IsKnown() bool {
	switch r {
	case CreditUsageExportGroupByUnspecified, CreditUsageExportGroupByDailySummary, CreditUsageExportGroupByResource:
		return true
	}
	return false
}

// CreditUsageReportFilter narrows the data returned by GetCreditUsageReport.
// Wrapping filters in a message (rather than adding bare fields) lets future
// filters (team, environment, resource kind) be added without further breaking
// changes.
type CreditUsageReportFilterParam struct {
	// Restrict the per-user breakdown to a single subject. The subject must be
	// PRINCIPAL_USER or PRINCIPAL_SERVICE_ACCOUNT and belong to the request's
	// organization. When unset, the report returns the default top-N users + "Others"
	// breakdown.
	//
	// When this field is set:
	//
	//   - daily_usage[*].user_usage contains rows only for the requested subject; no
	//     "Others" aggregation bucket is produced.
	//   - daily_usage[*].org_usage, team_usage, environment_usage, and
	//     conversation_usage are omitted (empty). Callers that need those sections
	//     should issue an unfiltered call.
	//   - period_start and updated_at remain populated.
	Subject param.Field[shared.SubjectParam] `json:"subject"`
}

func (r CreditUsageReportFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CreditsByType contains credits consumed for a single usage type.
type CreditsByType struct {
	Credits float64 `json:"credits"`
	// UsageType identifies the category of usage.
	UsageType UsageType         `json:"usageType"`
	JSON      creditsByTypeJSON `json:"-"`
}

// creditsByTypeJSON contains the JSON metadata for the struct [CreditsByType]
type creditsByTypeJSON struct {
	Credits     apijson.Field
	UsageType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditsByType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditsByTypeJSON) RawJSON() string {
	return r.raw
}

// CumulativeCreditUsage contains cumulative credit consumption totals.
type CumulativeCreditUsage struct {
	// Total credits consumed.
	TotalCredits float64 `json:"totalCredits"`
	// Credits consumed broken down by usage type.
	UsageByType []CreditsByType           `json:"usageByType"`
	JSON        cumulativeCreditUsageJSON `json:"-"`
}

// cumulativeCreditUsageJSON contains the JSON metadata for the struct
// [CumulativeCreditUsage]
type cumulativeCreditUsageJSON struct {
	TotalCredits apijson.Field
	UsageByType  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CumulativeCreditUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cumulativeCreditUsageJSON) RawJSON() string {
	return r.raw
}

// DailyCreditUsage contains credit usage for a single day.
type DailyCreditUsage struct {
	// Per-agent-execution usage for this day (top conversations + "Others"). Empty
	// agent_execution_id represents the "Others" aggregation bucket.
	ConversationUsage []AgentExecutionCreditUsage `json:"conversationUsage"`
	// Start of the day (midnight in the requested timezone).
	Date time.Time `json:"date" format:"date-time"`
	// Per-environment usage for this day (top environments + "Others"). Empty
	// environment_id represents the "Others" aggregation bucket.
	EnvironmentUsage []EnvironmentCreditUsage `json:"environmentUsage"`
	// Org-wide usage broken down by type.
	OrgUsage []CreditsByType `json:"orgUsage"`
	// Per-team usage for this day (top teams + "Others"). Empty team_id represents the
	// "Others" aggregation bucket.
	TeamUsage []TeamCreditUsage `json:"teamUsage"`
	// Org-wide intelligence usage broken down by model.
	UsageByModel []EnterpriseAIUsageByModel `json:"usageByModel"`
	// Per-user usage for this day (top users + "Others").
	UserUsage []UserCreditUsage    `json:"userUsage"`
	JSON      dailyCreditUsageJSON `json:"-"`
}

// dailyCreditUsageJSON contains the JSON metadata for the struct
// [DailyCreditUsage]
type dailyCreditUsageJSON struct {
	ConversationUsage apijson.Field
	Date              apijson.Field
	EnvironmentUsage  apijson.Field
	OrgUsage          apijson.Field
	TeamUsage         apijson.Field
	UsageByModel      apijson.Field
	UserUsage         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DailyCreditUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dailyCreditUsageJSON) RawJSON() string {
	return r.raw
}

type DailyEnterpriseAIUsage struct {
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// budget is unset when no monthly budget applies to the organization.
	Budget    EnterpriseAIUsageBudget `json:"budget"`
	TeamUsage []TeamEnterpriseAIUsage `json:"teamUsage"`
	Usage     EnterpriseAIUsage       `json:"usage"`
	// Usage for this day broken down by model. When the request filters by subject,
	// contains only that subject's model usage.
	UsageByModel []EnterpriseAIUsageByModel `json:"usageByModel"`
	UserUsage    []UserEnterpriseAIUsage    `json:"userUsage"`
	JSON         dailyEnterpriseAIUsageJSON `json:"-"`
}

// dailyEnterpriseAIUsageJSON contains the JSON metadata for the struct
// [DailyEnterpriseAIUsage]
type dailyEnterpriseAIUsageJSON struct {
	Date         apijson.Field
	Budget       apijson.Field
	TeamUsage    apijson.Field
	Usage        apijson.Field
	UsageByModel apijson.Field
	UserUsage    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DailyEnterpriseAIUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dailyEnterpriseAIUsageJSON) RawJSON() string {
	return r.raw
}

type EnterpriseAITokenUsage struct {
	CacheTokens  string                     `json:"cacheTokens"`
	InputTokens  string                     `json:"inputTokens"`
	OutputTokens string                     `json:"outputTokens"`
	TotalTokens  string                     `json:"totalTokens"`
	JSON         enterpriseAITokenUsageJSON `json:"-"`
}

// enterpriseAITokenUsageJSON contains the JSON metadata for the struct
// [EnterpriseAITokenUsage]
type enterpriseAITokenUsageJSON struct {
	CacheTokens  apijson.Field
	InputTokens  apijson.Field
	OutputTokens apijson.Field
	TotalTokens  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnterpriseAITokenUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enterpriseAITokenUsageJSON) RawJSON() string {
	return r.raw
}

type EnterpriseAIUsage struct {
	CostMicrounits string                 `json:"costMicrounits"`
	Credits        float64                `json:"credits"`
	Currency       BillingCurrency        `json:"currency"`
	Tokens         EnterpriseAITokenUsage `json:"tokens"`
	JSON           enterpriseAIUsageJSON  `json:"-"`
}

// enterpriseAIUsageJSON contains the JSON metadata for the struct
// [EnterpriseAIUsage]
type enterpriseAIUsageJSON struct {
	CostMicrounits apijson.Field
	Credits        apijson.Field
	Currency       apijson.Field
	Tokens         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnterpriseAIUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enterpriseAIUsageJSON) RawJSON() string {
	return r.raw
}

type EnterpriseAIUsageBudget struct {
	Currency                   BillingCurrency               `json:"currency"`
	MonthlyCostLimitMicrounits string                        `json:"monthlyCostLimitMicrounits" api:"nullable"`
	MonthlyCreditLimit         string                        `json:"monthlyCreditLimit" api:"nullable"`
	MonthToDateUsage           EnterpriseAIUsage             `json:"monthToDateUsage"`
	Source                     EnterpriseAIUsageBudgetSource `json:"source"`
	UtilizationPercent         float64                       `json:"utilizationPercent"`
	JSON                       enterpriseAIUsageBudgetJSON   `json:"-"`
}

// enterpriseAIUsageBudgetJSON contains the JSON metadata for the struct
// [EnterpriseAIUsageBudget]
type enterpriseAIUsageBudgetJSON struct {
	Currency                   apijson.Field
	MonthlyCostLimitMicrounits apijson.Field
	MonthlyCreditLimit         apijson.Field
	MonthToDateUsage           apijson.Field
	Source                     apijson.Field
	UtilizationPercent         apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *EnterpriseAIUsageBudget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enterpriseAIUsageBudgetJSON) RawJSON() string {
	return r.raw
}

type EnterpriseAIUsageBudgetSource string

const (
	EnterpriseAIUsageBudgetSourceUnspecified  EnterpriseAIUsageBudgetSource = "ENTERPRISE_AI_USAGE_BUDGET_SOURCE_UNSPECIFIED"
	EnterpriseAIUsageBudgetSourceOrganization EnterpriseAIUsageBudgetSource = "ENTERPRISE_AI_USAGE_BUDGET_SOURCE_ORGANIZATION"
	EnterpriseAIUsageBudgetSourceTeam         EnterpriseAIUsageBudgetSource = "ENTERPRISE_AI_USAGE_BUDGET_SOURCE_TEAM"
)

func (r EnterpriseAIUsageBudgetSource) IsKnown() bool {
	switch r {
	case EnterpriseAIUsageBudgetSourceUnspecified, EnterpriseAIUsageBudgetSourceOrganization, EnterpriseAIUsageBudgetSourceTeam:
		return true
	}
	return false
}

type EnterpriseAIUsageByModel struct {
	Model string `json:"model"`
	// Usage excluded from spend because no matching BYOK rate was configured.
	UnpricedUsage            EnterpriseAIUsage              `json:"unpricedUsage"`
	UnpricedUsageByTokenType []EnterpriseAIUsageByTokenType `json:"unpricedUsageByTokenType"`
	Usage                    EnterpriseAIUsage              `json:"usage"`
	UsageByTokenType         []EnterpriseAIUsageByTokenType `json:"usageByTokenType"`
	JSON                     enterpriseAIUsageByModelJSON   `json:"-"`
}

// enterpriseAIUsageByModelJSON contains the JSON metadata for the struct
// [EnterpriseAIUsageByModel]
type enterpriseAIUsageByModelJSON struct {
	Model                    apijson.Field
	UnpricedUsage            apijson.Field
	UnpricedUsageByTokenType apijson.Field
	Usage                    apijson.Field
	UsageByTokenType         apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnterpriseAIUsageByModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enterpriseAIUsageByModelJSON) RawJSON() string {
	return r.raw
}

type EnterpriseAIUsageByTokenType struct {
	TokenType ByokRateCardTokenType            `json:"tokenType"`
	Usage     EnterpriseAIUsage                `json:"usage"`
	JSON      enterpriseAIUsageByTokenTypeJSON `json:"-"`
}

// enterpriseAIUsageByTokenTypeJSON contains the JSON metadata for the struct
// [EnterpriseAIUsageByTokenType]
type enterpriseAIUsageByTokenTypeJSON struct {
	TokenType   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnterpriseAIUsageByTokenType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enterpriseAIUsageByTokenTypeJSON) RawJSON() string {
	return r.raw
}

type EnterpriseAIUsageTimeSeriesFilterParam struct {
	// Restrict the per-user breakdown to a single subject.
	Subject param.Field[shared.SubjectParam] `json:"subject"`
}

func (r EnterpriseAIUsageTimeSeriesFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnterpriseAIUserBudgetPolicySource string

const (
	EnterpriseAIUserBudgetPolicySourceUnspecified  EnterpriseAIUserBudgetPolicySource = "ENTERPRISE_AI_USER_BUDGET_POLICY_SOURCE_UNSPECIFIED"
	EnterpriseAIUserBudgetPolicySourceNone         EnterpriseAIUserBudgetPolicySource = "ENTERPRISE_AI_USER_BUDGET_POLICY_SOURCE_NONE"
	EnterpriseAIUserBudgetPolicySourceOrganization EnterpriseAIUserBudgetPolicySource = "ENTERPRISE_AI_USER_BUDGET_POLICY_SOURCE_ORGANIZATION"
	EnterpriseAIUserBudgetPolicySourceUser         EnterpriseAIUserBudgetPolicySource = "ENTERPRISE_AI_USER_BUDGET_POLICY_SOURCE_USER"
)

func (r EnterpriseAIUserBudgetPolicySource) IsKnown() bool {
	switch r {
	case EnterpriseAIUserBudgetPolicySourceUnspecified, EnterpriseAIUserBudgetPolicySourceNone, EnterpriseAIUserBudgetPolicySourceOrganization, EnterpriseAIUserBudgetPolicySourceUser:
		return true
	}
	return false
}

// EnvironmentCreditUsage contains a single environment's credit usage for a day,
// broken down by type.
type EnvironmentCreditUsage struct {
	DisplayName string `json:"displayName"`
	// Empty when representing the "Others" aggregation bucket.
	EnvironmentID string                     `json:"environmentId"`
	Usage         []CreditsByType            `json:"usage"`
	JSON          environmentCreditUsageJSON `json:"-"`
}

// environmentCreditUsageJSON contains the JSON metadata for the struct
// [EnvironmentCreditUsage]
type environmentCreditUsageJSON struct {
	DisplayName   apijson.Field
	EnvironmentID apijson.Field
	Usage         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentCreditUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentCreditUsageJSON) RawJSON() string {
	return r.raw
}

// TeamCreditUsage contains a single team's credit usage for a day, broken down by
// type.
type TeamCreditUsage struct {
	DisplayName string `json:"displayName"`
	// Empty when representing the "Others" aggregation bucket.
	TeamID string              `json:"teamId"`
	Usage  []CreditsByType     `json:"usage"`
	JSON   teamCreditUsageJSON `json:"-"`
}

// teamCreditUsageJSON contains the JSON metadata for the struct [TeamCreditUsage]
type teamCreditUsageJSON struct {
	DisplayName apijson.Field
	TeamID      apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamCreditUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamCreditUsageJSON) RawJSON() string {
	return r.raw
}

// TeamCumulativeCreditUsage contains a team's cumulative credit usage and
// allocation.
type TeamCumulativeCreditUsage struct {
	// The team's credit allocation (budget) in whole credits, if set. Not set means no
	// allocation has been configured for this team.
	CreditBudget string `json:"creditBudget" api:"nullable"`
	DisplayName  string `json:"displayName"`
	TeamID       string `json:"teamId"`
	// Cumulative credit usage for this team.
	Usage CumulativeCreditUsage         `json:"usage"`
	JSON  teamCumulativeCreditUsageJSON `json:"-"`
}

// teamCumulativeCreditUsageJSON contains the JSON metadata for the struct
// [TeamCumulativeCreditUsage]
type teamCumulativeCreditUsageJSON struct {
	CreditBudget apijson.Field
	DisplayName  apijson.Field
	TeamID       apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TeamCumulativeCreditUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamCumulativeCreditUsageJSON) RawJSON() string {
	return r.raw
}

type TeamEnterpriseAIUsage struct {
	// budget is unset when no monthly budget applies to this team.
	Budget           EnterpriseAIUsageBudget        `json:"budget"`
	DisplayName      string                         `json:"displayName"`
	TeamID           string                         `json:"teamId" format:"uuid"`
	Usage            EnterpriseAIUsage              `json:"usage"`
	UsageByTokenType []EnterpriseAIUsageByTokenType `json:"usageByTokenType"`
	JSON             teamEnterpriseAIUsageJSON      `json:"-"`
}

// teamEnterpriseAIUsageJSON contains the JSON metadata for the struct
// [TeamEnterpriseAIUsage]
type teamEnterpriseAIUsageJSON struct {
	Budget           apijson.Field
	DisplayName      apijson.Field
	TeamID           apijson.Field
	Usage            apijson.Field
	UsageByTokenType apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TeamEnterpriseAIUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamEnterpriseAIUsageJSON) RawJSON() string {
	return r.raw
}

// UsageType identifies the category of usage.
type UsageType string

const (
	UsageTypeUnspecified UsageType = "USAGE_TYPE_UNSPECIFIED"
	UsageTypeEnvironment UsageType = "USAGE_TYPE_ENVIRONMENT"
	UsageTypeAgentic     UsageType = "USAGE_TYPE_AGENTIC"
)

func (r UsageType) IsKnown() bool {
	switch r {
	case UsageTypeUnspecified, UsageTypeEnvironment, UsageTypeAgentic:
		return true
	}
	return false
}

type UserCostBudgetUsage struct {
	BudgetSource               EnterpriseAIUserBudgetPolicySource `json:"budgetSource"`
	Currency                   BillingCurrency                    `json:"currency"`
	DisplayName                string                             `json:"displayName"`
	IsServiceAccount           bool                               `json:"isServiceAccount"`
	MonthlyCostLimitMicrounits string                             `json:"monthlyCostLimitMicrounits" api:"nullable"`
	// Usage within the requested date range. Reflects true month-to-date usage when
	// the range starts on the first day of the month.
	MonthToDateUsage   EnterpriseAIUsage       `json:"monthToDateUsage"`
	NoCap              bool                    `json:"noCap"`
	OverBudget         bool                    `json:"overBudget"`
	UserID             string                  `json:"userId"`
	UtilizationPercent float64                 `json:"utilizationPercent"`
	JSON               userCostBudgetUsageJSON `json:"-"`
}

// userCostBudgetUsageJSON contains the JSON metadata for the struct
// [UserCostBudgetUsage]
type userCostBudgetUsageJSON struct {
	BudgetSource               apijson.Field
	Currency                   apijson.Field
	DisplayName                apijson.Field
	IsServiceAccount           apijson.Field
	MonthlyCostLimitMicrounits apijson.Field
	MonthToDateUsage           apijson.Field
	NoCap                      apijson.Field
	OverBudget                 apijson.Field
	UserID                     apijson.Field
	UtilizationPercent         apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UserCostBudgetUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userCostBudgetUsageJSON) RawJSON() string {
	return r.raw
}

type UserCreditBudgetUsage struct {
	BudgetSource EnterpriseAIUserBudgetPolicySource `json:"budgetSource"`
	CreditBudget string                             `json:"creditBudget" api:"nullable"`
	DisplayName  string                             `json:"displayName"`
	// True when user_id refers to a service account rather than a human user. The
	// dashboard uses this to mark non-human accounts in admin tables.
	IsServiceAccount bool `json:"isServiceAccount"`
	// CumulativeCreditUsage contains cumulative credit consumption totals.
	MonthToDateUsage CumulativeCreditUsage `json:"monthToDateUsage"`
	NoCap            bool                  `json:"noCap"`
	OverBudget       bool                  `json:"overBudget"`
	// Month-to-date intelligence usage broken down by model.
	UsageByModel       []EnterpriseAIUsageByModel `json:"usageByModel"`
	UserID             string                     `json:"userId"`
	UtilizationPercent float64                    `json:"utilizationPercent"`
	JSON               userCreditBudgetUsageJSON  `json:"-"`
}

// userCreditBudgetUsageJSON contains the JSON metadata for the struct
// [UserCreditBudgetUsage]
type userCreditBudgetUsageJSON struct {
	BudgetSource       apijson.Field
	CreditBudget       apijson.Field
	DisplayName        apijson.Field
	IsServiceAccount   apijson.Field
	MonthToDateUsage   apijson.Field
	NoCap              apijson.Field
	OverBudget         apijson.Field
	UsageByModel       apijson.Field
	UserID             apijson.Field
	UtilizationPercent apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserCreditBudgetUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userCreditBudgetUsageJSON) RawJSON() string {
	return r.raw
}

// UserCreditUsage contains a single user's credit usage, broken down by type.
type UserCreditUsage struct {
	DisplayName string          `json:"displayName"`
	Usage       []CreditsByType `json:"usage"`
	// Intelligence usage broken down by model.
	UsageByModel []EnterpriseAIUsageByModel `json:"usageByModel"`
	// Empty when representing the "Others" aggregation bucket.
	UserID string              `json:"userId"`
	JSON   userCreditUsageJSON `json:"-"`
}

// userCreditUsageJSON contains the JSON metadata for the struct [UserCreditUsage]
type userCreditUsageJSON struct {
	DisplayName  apijson.Field
	Usage        apijson.Field
	UsageByModel apijson.Field
	UserID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserCreditUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userCreditUsageJSON) RawJSON() string {
	return r.raw
}

type UserEnterpriseAIUsage struct {
	DisplayName string            `json:"displayName"`
	Usage       EnterpriseAIUsage `json:"usage"`
	// Empty when representing the "Others" aggregation bucket.
	UserID string                    `json:"userId"`
	JSON   userEnterpriseAIUsageJSON `json:"-"`
}

// userEnterpriseAIUsageJSON contains the JSON metadata for the struct
// [UserEnterpriseAIUsage]
type userEnterpriseAIUsageJSON struct {
	DisplayName apijson.Field
	Usage       apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserEnterpriseAIUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEnterpriseAIUsageJSON) RawJSON() string {
	return r.raw
}

type BillingGetCreditUsageExportResponse struct {
	// Signed download URL for the CSV export. Valid for five minutes, and only for the
	// principal that requested it.
	DownloadURL string                                  `json:"downloadUrl"`
	JSON        billingGetCreditUsageExportResponseJSON `json:"-"`
}

// billingGetCreditUsageExportResponseJSON contains the JSON metadata for the
// struct [BillingGetCreditUsageExportResponse]
type billingGetCreditUsageExportResponseJSON struct {
	DownloadURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingGetCreditUsageExportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingGetCreditUsageExportResponseJSON) RawJSON() string {
	return r.raw
}

type BillingGetCreditUsageReportResponse struct {
	// One entry per day in the requested date range.
	DailyUsage []DailyCreditUsage `json:"dailyUsage"`
	// Start of the billing period for this organization. Used by the frontend to
	// filter out months before usage tracking began.
	PeriodStart time.Time `json:"periodStart" api:"nullable" format:"date-time"`
	// When the report data was last computed.
	UpdatedAt time.Time                               `json:"updatedAt" format:"date-time"`
	JSON      billingGetCreditUsageReportResponseJSON `json:"-"`
}

// billingGetCreditUsageReportResponseJSON contains the JSON metadata for the
// struct [BillingGetCreditUsageReportResponse]
type billingGetCreditUsageReportResponseJSON struct {
	DailyUsage  apijson.Field
	PeriodStart apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingGetCreditUsageReportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingGetCreditUsageReportResponseJSON) RawJSON() string {
	return r.raw
}

type BillingGetCumulativeCreditUsageResponse struct {
	// Org-wide cumulative usage, broken down by type and total.
	OrgUsage CumulativeCreditUsage `json:"orgUsage"`
	// Start of the cumulative calculation period. Cumulative totals are computed from
	// this date forward.
	PeriodStart time.Time `json:"periodStart" format:"date-time"`
	// Per-team cumulative usage with credit allocation comparison. Returns all teams
	// (no top-N limit).
	TeamUsage []TeamCumulativeCreditUsage `json:"teamUsage"`
	// Usage by members not assigned to any team.
	UnteamedUsage CumulativeCreditUsage `json:"unteamedUsage"`
	// Per-user month-to-date usage for every user with usage in the period. The budget
	// fields on each entry are populated only when a monthly budget applies to that
	// user. This list is not paginated or capped; for large organizations prefer
	// ListEnterpriseUserCreditUsage.
	UserUsage []UserCreditBudgetUsage                     `json:"userUsage"`
	JSON      billingGetCumulativeCreditUsageResponseJSON `json:"-"`
}

// billingGetCumulativeCreditUsageResponseJSON contains the JSON metadata for the
// struct [BillingGetCumulativeCreditUsageResponse]
type billingGetCumulativeCreditUsageResponseJSON struct {
	OrgUsage      apijson.Field
	PeriodStart   apijson.Field
	TeamUsage     apijson.Field
	UnteamedUsage apijson.Field
	UserUsage     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BillingGetCumulativeCreditUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingGetCumulativeCreditUsageResponseJSON) RawJSON() string {
	return r.raw
}

type BillingGetEnterpriseAIUsageSummaryResponse struct {
	// budget is unset when no monthly budget applies to the organization.
	Budget EnterpriseAIUsageBudget `json:"budget"`
	// calculated_at is the time through which usage has been calculated. Usage after
	// this timestamp may still be processing.
	CalculatedAt time.Time                                      `json:"calculatedAt" format:"date-time"`
	Usage        EnterpriseAIUsage                              `json:"usage"`
	UsageByModel []EnterpriseAIUsageByModel                     `json:"usageByModel"`
	JSON         billingGetEnterpriseAIUsageSummaryResponseJSON `json:"-"`
}

// billingGetEnterpriseAIUsageSummaryResponseJSON contains the JSON metadata for
// the struct [BillingGetEnterpriseAIUsageSummaryResponse]
type billingGetEnterpriseAIUsageSummaryResponseJSON struct {
	Budget       apijson.Field
	CalculatedAt apijson.Field
	Usage        apijson.Field
	UsageByModel apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BillingGetEnterpriseAIUsageSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingGetEnterpriseAIUsageSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type BillingGetEnterpriseAIUsageTimeSeriesResponse struct {
	// calculated_at is the time through which usage has been calculated. Usage after
	// this timestamp may still be processing.
	CalculatedAt time.Time                                         `json:"calculatedAt" format:"date-time"`
	DailyUsage   []DailyEnterpriseAIUsage                          `json:"dailyUsage"`
	JSON         billingGetEnterpriseAIUsageTimeSeriesResponseJSON `json:"-"`
}

// billingGetEnterpriseAIUsageTimeSeriesResponseJSON contains the JSON metadata for
// the struct [BillingGetEnterpriseAIUsageTimeSeriesResponse]
type billingGetEnterpriseAIUsageTimeSeriesResponseJSON struct {
	CalculatedAt apijson.Field
	DailyUsage   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BillingGetEnterpriseAIUsageTimeSeriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingGetEnterpriseAIUsageTimeSeriesResponseJSON) RawJSON() string {
	return r.raw
}

type BillingGetCreditUsageExportParams struct {
	// Date range to export. Both start and end dates are inclusive; time-of-day is
	// ignored. Unlike GetCreditUsageReport, the range may cover up to a year.
	DateRange      param.Field[shared.DateRangeParam] `json:"dateRange" api:"required"`
	OrganizationID param.Field[string]                `json:"organizationId" api:"required" format:"uuid"`
	// How to group the export data. Defaults to DAILY_SUMMARY.
	GroupBy param.Field[CreditUsageExportGroupBy] `json:"groupBy"`
}

func (r BillingGetCreditUsageExportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingGetCreditUsageReportParams struct {
	// Date range for the report. Both start and end dates are inclusive. Time-of-day
	// is ignored; dates are truncated to midnight in the specified timezone. The range
	// must not exceed 31 days.
	DateRange      param.Field[shared.DateRangeParam] `json:"dateRange" api:"required"`
	OrganizationID param.Field[string]                `json:"organizationId" api:"required" format:"uuid"`
	// Optional filter narrowing the returned data. When unset or empty, the response
	// preserves the default behavior (top-N users + "Others"). See
	// CreditUsageReportFilter for per-field response-scoping semantics.
	Filter param.Field[CreditUsageReportFilterParam] `json:"filter"`
	// IANA timezone name (e.g. "America/New_York", "Europe/Berlin") used to bucket
	// daily usage. When empty, defaults to "UTC".
	Timezone param.Field[string] `json:"timezone"`
}

func (r BillingGetCreditUsageReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingGetCumulativeCreditUsageParams struct {
	// organization_id is the ID of the organization to get cumulative usage for.
	OrganizationID param.Field[string] `json:"organizationId" api:"required" format:"uuid"`
	// as_of is the point in time to compute cumulative usage up to. Defaults to now if
	// not set.
	AsOf param.Field[time.Time] `json:"asOf" format:"date-time"`
}

func (r BillingGetCumulativeCreditUsageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingGetEnterpriseAIUsageSummaryParams struct {
	// Date range for the summary. Both start and end dates are inclusive. Time-of-day
	// is ignored; dates are truncated to midnight in the specified timezone.
	DateRange      param.Field[shared.DateRangeParam] `json:"dateRange" api:"required"`
	OrganizationID param.Field[string]                `json:"organizationId" api:"required" format:"uuid"`
	// IANA timezone name used to bucket usage. When empty, defaults to "UTC".
	Timezone param.Field[string] `json:"timezone"`
}

func (r BillingGetEnterpriseAIUsageSummaryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingGetEnterpriseAIUsageTimeSeriesParams struct {
	// Date range for the daily usage series. Both start and end dates are inclusive.
	// Time-of-day is ignored; dates are truncated to midnight in the specified
	// timezone.
	DateRange      param.Field[shared.DateRangeParam]                  `json:"dateRange" api:"required"`
	OrganizationID param.Field[string]                                 `json:"organizationId" api:"required" format:"uuid"`
	Filter         param.Field[EnterpriseAIUsageTimeSeriesFilterParam] `json:"filter"`
	// IANA timezone name used to bucket daily usage. When empty, defaults to "UTC".
	Timezone param.Field[string] `json:"timezone"`
}

func (r BillingGetEnterpriseAIUsageTimeSeriesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingListEnterpriseAITeamUsageParams struct {
	// Date range for the team usage list. Both start and end dates are inclusive.
	// Time-of-day is ignored; dates are truncated to midnight in the specified
	// timezone.
	DateRange      param.Field[shared.DateRangeParam]                            `json:"dateRange" api:"required"`
	OrganizationID param.Field[string]                                           `json:"organizationId" api:"required" format:"uuid"`
	Token          param.Field[string]                                           `query:"token"`
	PageSize       param.Field[int64]                                            `query:"pageSize"`
	Filter         param.Field[BillingListEnterpriseAITeamUsageParamsFilter]     `json:"filter"`
	Pagination     param.Field[BillingListEnterpriseAITeamUsageParamsPagination] `json:"pagination"`
	// IANA timezone name used to bucket usage. When empty, defaults to "UTC".
	Timezone param.Field[string] `json:"timezone"`
}

func (r BillingListEnterpriseAITeamUsageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [BillingListEnterpriseAITeamUsageParams]'s query parameters
// as `url.Values`.
func (r BillingListEnterpriseAITeamUsageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillingListEnterpriseAITeamUsageParamsFilter struct {
	TeamIDs param.Field[[]string] `json:"teamIds" format:"uuid"`
}

func (r BillingListEnterpriseAITeamUsageParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingListEnterpriseAITeamUsageParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r BillingListEnterpriseAITeamUsageParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingListEnterpriseAIUserUsageParams struct {
	// Date range for the user usage list. Both start and end dates are inclusive.
	// Time-of-day is ignored; dates are truncated to midnight in the specified
	// timezone.
	DateRange      param.Field[shared.DateRangeParam] `json:"dateRange" api:"required"`
	OrganizationID param.Field[string]                `json:"organizationId" api:"required" format:"uuid"`
	Token          param.Field[string]                `query:"token"`
	PageSize       param.Field[int64]                 `query:"pageSize"`
	// Optional filter narrowing the returned user usage. When set to a subject, the
	// response contains only usage for that user or service account.
	Filter     param.Field[BillingListEnterpriseAIUserUsageParamsFilter]     `json:"filter"`
	Pagination param.Field[BillingListEnterpriseAIUserUsageParamsPagination] `json:"pagination"`
	// sort controls the ordering of results. Defaults to total spend descending.
	Sort param.Field[BillingListEnterpriseAIUserUsageParamsSort] `json:"sort"`
	// IANA timezone name used to bucket usage. When empty, defaults to "UTC".
	Timezone param.Field[string] `json:"timezone"`
}

func (r BillingListEnterpriseAIUserUsageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [BillingListEnterpriseAIUserUsageParams]'s query parameters
// as `url.Values`.
func (r BillingListEnterpriseAIUserUsageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional filter narrowing the returned user usage. When set to a subject, the
// response contains only usage for that user or service account.
type BillingListEnterpriseAIUserUsageParamsFilter struct {
	// Restrict the user usage list to a single subject. The subject must be
	// PRINCIPAL_USER or PRINCIPAL_SERVICE_ACCOUNT and belong to the request's
	// organization.
	Subject param.Field[shared.SubjectParam] `json:"subject"`
}

func (r BillingListEnterpriseAIUserUsageParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingListEnterpriseAIUserUsageParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r BillingListEnterpriseAIUserUsageParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// sort controls the ordering of results. Defaults to total spend descending.
type BillingListEnterpriseAIUserUsageParamsSort struct {
	Field param.Field[BillingListEnterpriseAIUserUsageParamsSortField] `json:"field"`
	Order param.Field[shared.SortOrder]                                `json:"order"`
}

func (r BillingListEnterpriseAIUserUsageParamsSort) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingListEnterpriseAIUserUsageParamsSortField string

const (
	BillingListEnterpriseAIUserUsageParamsSortFieldSortFieldUnspecified BillingListEnterpriseAIUserUsageParamsSortField = "SORT_FIELD_UNSPECIFIED"
	BillingListEnterpriseAIUserUsageParamsSortFieldSortFieldUsage       BillingListEnterpriseAIUserUsageParamsSortField = "SORT_FIELD_USAGE"
	BillingListEnterpriseAIUserUsageParamsSortFieldSortFieldDisplayName BillingListEnterpriseAIUserUsageParamsSortField = "SORT_FIELD_DISPLAY_NAME"
	BillingListEnterpriseAIUserUsageParamsSortFieldSortFieldBudget      BillingListEnterpriseAIUserUsageParamsSortField = "SORT_FIELD_BUDGET"
	BillingListEnterpriseAIUserUsageParamsSortFieldSortFieldBudgetUsed  BillingListEnterpriseAIUserUsageParamsSortField = "SORT_FIELD_BUDGET_USED"
)

func (r BillingListEnterpriseAIUserUsageParamsSortField) IsKnown() bool {
	switch r {
	case BillingListEnterpriseAIUserUsageParamsSortFieldSortFieldUnspecified, BillingListEnterpriseAIUserUsageParamsSortFieldSortFieldUsage, BillingListEnterpriseAIUserUsageParamsSortFieldSortFieldDisplayName, BillingListEnterpriseAIUserUsageParamsSortFieldSortFieldBudget, BillingListEnterpriseAIUserUsageParamsSortFieldSortFieldBudgetUsed:
		return true
	}
	return false
}

type BillingListEnterpriseUserCreditUsageParams struct {
	// organization_id is the ID of the organization to list user credit usage for.
	OrganizationID param.Field[string] `json:"organizationId" api:"required" format:"uuid"`
	Token          param.Field[string] `query:"token"`
	PageSize       param.Field[int64]  `query:"pageSize"`
	// as_of is the point in time to compute month-to-date usage up to. Defaults to now
	// if not set.
	AsOf       param.Field[time.Time]                                            `json:"asOf" format:"date-time"`
	Pagination param.Field[BillingListEnterpriseUserCreditUsageParamsPagination] `json:"pagination"`
	// sort controls the ordering of results. Defaults to total credits descending.
	Sort param.Field[BillingListEnterpriseUserCreditUsageParamsSort] `json:"sort"`
}

func (r BillingListEnterpriseUserCreditUsageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [BillingListEnterpriseUserCreditUsageParams]'s query
// parameters as `url.Values`.
func (r BillingListEnterpriseUserCreditUsageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillingListEnterpriseUserCreditUsageParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r BillingListEnterpriseUserCreditUsageParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// sort controls the ordering of results. Defaults to total credits descending.
type BillingListEnterpriseUserCreditUsageParamsSort struct {
	Field param.Field[BillingListEnterpriseUserCreditUsageParamsSortField] `json:"field"`
	Order param.Field[shared.SortOrder]                                    `json:"order"`
}

func (r BillingListEnterpriseUserCreditUsageParamsSort) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingListEnterpriseUserCreditUsageParamsSortField string

const (
	BillingListEnterpriseUserCreditUsageParamsSortFieldSortFieldUnspecified BillingListEnterpriseUserCreditUsageParamsSortField = "SORT_FIELD_UNSPECIFIED"
	BillingListEnterpriseUserCreditUsageParamsSortFieldSortFieldUsage       BillingListEnterpriseUserCreditUsageParamsSortField = "SORT_FIELD_USAGE"
	BillingListEnterpriseUserCreditUsageParamsSortFieldSortFieldDisplayName BillingListEnterpriseUserCreditUsageParamsSortField = "SORT_FIELD_DISPLAY_NAME"
	BillingListEnterpriseUserCreditUsageParamsSortFieldSortFieldBudget      BillingListEnterpriseUserCreditUsageParamsSortField = "SORT_FIELD_BUDGET"
	BillingListEnterpriseUserCreditUsageParamsSortFieldSortFieldBudgetUsed  BillingListEnterpriseUserCreditUsageParamsSortField = "SORT_FIELD_BUDGET_USED"
)

func (r BillingListEnterpriseUserCreditUsageParamsSortField) IsKnown() bool {
	switch r {
	case BillingListEnterpriseUserCreditUsageParamsSortFieldSortFieldUnspecified, BillingListEnterpriseUserCreditUsageParamsSortFieldSortFieldUsage, BillingListEnterpriseUserCreditUsageParamsSortFieldSortFieldDisplayName, BillingListEnterpriseUserCreditUsageParamsSortFieldSortFieldBudget, BillingListEnterpriseUserCreditUsageParamsSortFieldSortFieldBudgetUsed:
		return true
	}
	return false
}
