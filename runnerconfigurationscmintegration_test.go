// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gitpod-go"
	"github.com/stainless-sdks/gitpod-go/internal/testutil"
	"github.com/stainless-sdks/gitpod-go/option"
)

func TestRunnerConfigurationScmIntegrationNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gitpod.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.RunnerConfigurations.ScmIntegration.New(context.TODO(), gitpod.RunnerConfigurationScmIntegrationNewParams{
		Body: gitpod.RunnerConfigurationScmIntegrationNewParamsBody{
			Host:                       "host",
			OAuthClientID:              "oauthClientId",
			OAuthPlaintextClientSecret: "oauthPlaintextClientSecret",
			Pat:                        true,
			RunnerID:                   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ScmID:                      "scmId",
		},
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationScmIntegrationNewParamsConnectProtocolVersion1),
		ConnectTimeoutMs:       gitpod.F(0.000000),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
