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
		RunnerID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ScmIntegration: gitpod.F(gitpod.RunnerConfigurationValidateParamsScmIntegration{
			ID:                         gitpod.F("id"),
			Host:                       gitpod.F("host"),
			OAuthClientID:              gitpod.F("oauthClientId"),
			OAuthEncryptedClientSecret: gitpod.F("U3RhaW5sZXNzIHJvY2tz"),
			OAuthPlaintextClientSecret: gitpod.F("oauthPlaintextClientSecret"),
			Pat:                        gitpod.F(true),
			ScmID:                      gitpod.F("scmId"),
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
