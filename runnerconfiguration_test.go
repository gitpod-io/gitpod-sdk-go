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
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

func TestRunnerConfigurationValidateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		EnvironmentClass: gitpod.F(shared.EnvironmentClassParam{
			ID:       gitpod.F("id"),
			RunnerID: gitpod.F("runnerId"),
			Configuration: gitpod.F([]shared.FieldValueParam{{
				Key:   gitpod.F("key"),
				Value: gitpod.F("value"),
			}}),
			Description: gitpod.F("xxx"),
			DisplayName: gitpod.F("xxx"),
			Enabled:     gitpod.F(true),
		}),
		RunnerID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
		ScmIntegration: gitpod.F(gitpod.RunnerConfigurationValidateParamsScmIntegration{
			ID:                         gitpod.F("integration-id"),
			Host:                       gitpod.F("github.com"),
			IssuerURL:                  gitpod.F("issuerUrl"),
			OAuthClientID:              gitpod.F("client_id"),
			OAuthEncryptedClientSecret: gitpod.F("U3RhaW5sZXNzIHJvY2tz"),
			OAuthPlaintextClientSecret: gitpod.F("client_secret"),
			Pat:                        gitpod.F(true),
			ScmID:                      gitpod.F("github"),
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
