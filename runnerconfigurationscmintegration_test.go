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
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Runners.Configurations.ScmIntegrations.New(context.TODO(), gitpod.RunnerConfigurationScmIntegrationNewParams{
		Body: gitpod.RunnerConfigurationScmIntegrationNewParamsBodyOAuthClientIDIsTheOAuthAppSClientIDIfOAuthIsConfiguredIfConfiguredOAuthPlaintextClientSecretMustAlsoBeSet{
			OAuthClientID: gitpod.F("oauthClientId"),
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

func TestRunnerConfigurationScmIntegrationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.ScmIntegrations.Get(context.TODO(), gitpod.RunnerConfigurationScmIntegrationGetParams{
		Encoding:               gitpod.F(gitpod.RunnerConfigurationScmIntegrationGetParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationScmIntegrationGetParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.RunnerConfigurationScmIntegrationGetParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.RunnerConfigurationScmIntegrationGetParamsConnectV1),
		Message:                gitpod.F("message"),
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

func TestRunnerConfigurationScmIntegrationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.ScmIntegrations.Update(context.TODO(), gitpod.RunnerConfigurationScmIntegrationUpdateParams{
		Body: gitpod.RunnerConfigurationScmIntegrationUpdateParamsBodyOAuthClientIDCanBeSetToUpdateTheOAuthAppSClientIDIfAnEmptyStringIsSetTheOAuthConfigurationWillBeRemovedRegardlessOfWhetherAClientSecretIsSetAndAnyExistingHostAuthenticationTokensForTheScmIntegrationSRunnerAndHostThatWereCreatedUsingTheOAuthAppWillBeDeletedThisMightLeadToUsersBeingUnableToAccessTheirRepositoriesUntilTheyReAuthenticate{
			OAuthClientID: gitpod.F("oauthClientId"),
		},
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationScmIntegrationUpdateParamsConnectProtocolVersion1),
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

func TestRunnerConfigurationScmIntegrationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.ScmIntegrations.List(context.TODO(), gitpod.RunnerConfigurationScmIntegrationListParams{
		Encoding:               gitpod.F(gitpod.RunnerConfigurationScmIntegrationListParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationScmIntegrationListParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.RunnerConfigurationScmIntegrationListParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.RunnerConfigurationScmIntegrationListParamsConnectV1),
		Message:                gitpod.F("message"),
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

func TestRunnerConfigurationScmIntegrationDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.ScmIntegrations.Delete(context.TODO(), gitpod.RunnerConfigurationScmIntegrationDeleteParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationScmIntegrationDeleteParamsConnectProtocolVersion1),
		ID:                     gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
