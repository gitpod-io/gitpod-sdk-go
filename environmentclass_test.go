// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/gitpod-io/gitpod-sdk-go"
	"github.com/gitpod-io/gitpod-sdk-go/internal/testutil"
	"github.com/gitpod-io/gitpod-sdk-go/option"
)

func TestEnvironmentClassListWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Classes.List(context.TODO(), gitpod.EnvironmentClassListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.EnvironmentClassListParamsFilter{
			CanCreateEnvironments: gitpod.F(true),
			Enabled:               gitpod.F(true),
			RunnerIDs:             gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			RunnerKinds:           gitpod.F([]gitpod.RunnerKind{gitpod.RunnerKindUnspecified}),
			RunnerProviders:       gitpod.F([]gitpod.RunnerProvider{gitpod.RunnerProviderUnspecified}),
		}),
		Pagination: gitpod.F(gitpod.EnvironmentClassListParamsPagination{
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
