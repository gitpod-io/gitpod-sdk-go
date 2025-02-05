// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/gitpod-io/flex-sdk-go"
	"github.com/gitpod-io/flex-sdk-go/internal/testutil"
	"github.com/gitpod-io/flex-sdk-go/option"
)

func TestRunnerConfigurationValidateWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.Validate(context.TODO(), gitpod.RunnerConfigurationValidateParams{
		Body: gitpod.RunnerConfigurationValidateParamsBodyObject{
			EnvironmentClass: gitpod.F(gitpod.RunnerConfigurationValidateParamsBodyObjectEnvironmentClass{
				ID: gitpod.F("id"),
				Configuration: gitpod.F([]gitpod.RunnerConfigurationValidateParamsBodyObjectEnvironmentClassConfiguration{{
					Key:   gitpod.F("key"),
					Value: gitpod.F("value"),
				}}),
				Description: gitpod.F("xxx"),
				DisplayName: gitpod.F("xxx"),
				Enabled:     gitpod.F(true),
				RunnerID:    gitpod.F("runnerId"),
			}),
			RunnerID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
