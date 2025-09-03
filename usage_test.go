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
)

func TestUsageListEnvironmentRuntimeRecordsWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Usage.ListEnvironmentRuntimeRecords(context.TODO(), gitpod.UsageListEnvironmentRuntimeRecordsParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.UsageListEnvironmentRuntimeRecordsParamsFilter{
			DateRange: gitpod.F(gitpod.UsageListEnvironmentRuntimeRecordsParamsFilterDateRange{
				EndTime:   gitpod.F(time.Now()),
				StartTime: gitpod.F(time.Now()),
			}),
			ProjectID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
		}),
		Pagination: gitpod.F(gitpod.UsageListEnvironmentRuntimeRecordsParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(100)),
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
