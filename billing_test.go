// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go"
	"github.com/gitpod-io/gitpod-sdk-go/internal/testutil"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

func TestBillingGetCreditUsageExportWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gitpod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Billing.GetCreditUsageExport(context.TODO(), gitpod.BillingGetCreditUsageExportParams{
		DateRange: gitpod.F(shared.DateRangeParam{
			EndTime:   gitpod.F(time.Now()),
			StartTime: gitpod.F(time.Now()),
		}),
		OrganizationID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
		GroupBy:        gitpod.F(gitpod.CreditUsageExportGroupByDailySummary),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBillingGetCreditUsageReportWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gitpod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Billing.GetCreditUsageReport(context.TODO(), gitpod.BillingGetCreditUsageReportParams{
		DateRange: gitpod.F(shared.DateRangeParam{
			EndTime:   gitpod.F(time.Now()),
			StartTime: gitpod.F(time.Now()),
		}),
		OrganizationID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
		Filter: gitpod.F(gitpod.CreditUsageReportFilterParam{
			Subject: gitpod.F(shared.SubjectParam{
				ID:        gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Principal: gitpod.F(shared.PrincipalUnspecified),
			}),
		}),
		Timezone: gitpod.F("timezone"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBillingGetCumulativeCreditUsageWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gitpod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Billing.GetCumulativeCreditUsage(context.TODO(), gitpod.BillingGetCumulativeCreditUsageParams{
		OrganizationID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
		AsOf:           gitpod.F(time.Now()),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBillingGetEnterpriseAIUsageSummaryWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gitpod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Billing.GetEnterpriseAIUsageSummary(context.TODO(), gitpod.BillingGetEnterpriseAIUsageSummaryParams{
		DateRange: gitpod.F(shared.DateRangeParam{
			EndTime:   gitpod.F(time.Now()),
			StartTime: gitpod.F(time.Now()),
		}),
		OrganizationID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
		Timezone:       gitpod.F("timezone"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBillingGetEnterpriseAIUsageTimeSeriesWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gitpod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Billing.GetEnterpriseAIUsageTimeSeries(context.TODO(), gitpod.BillingGetEnterpriseAIUsageTimeSeriesParams{
		DateRange: gitpod.F(shared.DateRangeParam{
			EndTime:   gitpod.F(time.Now()),
			StartTime: gitpod.F(time.Now()),
		}),
		OrganizationID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
		Filter: gitpod.F(gitpod.EnterpriseAIUsageTimeSeriesFilterParam{
			Subject: gitpod.F(shared.SubjectParam{
				ID:        gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Principal: gitpod.F(shared.PrincipalUnspecified),
			}),
		}),
		Timezone: gitpod.F("timezone"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBillingListEnterpriseAITeamUsageWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gitpod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Billing.ListEnterpriseAITeamUsage(context.TODO(), gitpod.BillingListEnterpriseAITeamUsageParams{
		DateRange: gitpod.F(shared.DateRangeParam{
			EndTime:   gitpod.F(time.Now()),
			StartTime: gitpod.F(time.Now()),
		}),
		OrganizationID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
		Token:          gitpod.F("token"),
		PageSize:       gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.BillingListEnterpriseAITeamUsageParamsFilter{
			TeamIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.BillingListEnterpriseAITeamUsageParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(100)),
		}),
		Timezone: gitpod.F("timezone"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBillingListEnterpriseAIUserUsageWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gitpod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Billing.ListEnterpriseAIUserUsage(context.TODO(), gitpod.BillingListEnterpriseAIUserUsageParams{
		DateRange: gitpod.F(shared.DateRangeParam{
			EndTime:   gitpod.F(time.Now()),
			StartTime: gitpod.F(time.Now()),
		}),
		OrganizationID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
		Token:          gitpod.F("token"),
		PageSize:       gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.BillingListEnterpriseAIUserUsageParamsFilter{
			Subject: gitpod.F(shared.SubjectParam{
				ID:        gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Principal: gitpod.F(shared.PrincipalUnspecified),
			}),
		}),
		Pagination: gitpod.F(gitpod.BillingListEnterpriseAIUserUsageParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(100)),
		}),
		Sort: gitpod.F(gitpod.BillingListEnterpriseAIUserUsageParamsSort{
			Field: gitpod.F(gitpod.BillingListEnterpriseAIUserUsageParamsSortFieldSortFieldUnspecified),
			Order: gitpod.F(shared.SortOrderUnspecified),
		}),
		Timezone: gitpod.F("timezone"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBillingListEnterpriseUserCreditUsageWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gitpod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Billing.ListEnterpriseUserCreditUsage(context.TODO(), gitpod.BillingListEnterpriseUserCreditUsageParams{
		OrganizationID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
		Token:          gitpod.F("token"),
		PageSize:       gitpod.F(int64(0)),
		AsOf:           gitpod.F(time.Now()),
		Pagination: gitpod.F(gitpod.BillingListEnterpriseUserCreditUsageParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(50)),
		}),
		Sort: gitpod.F(gitpod.BillingListEnterpriseUserCreditUsageParamsSort{
			Field: gitpod.F(gitpod.BillingListEnterpriseUserCreditUsageParamsSortFieldSortFieldUnspecified),
			Order: gitpod.F(shared.SortOrderUnspecified),
		}),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
