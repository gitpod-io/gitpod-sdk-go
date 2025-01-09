// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/gitpod-go"
	"github.com/stainless-sdks/gitpod-go/internal/testutil"
	"github.com/stainless-sdks/gitpod-go/option"
)

func TestRunnerInteractionGetHostAuthenticationTokenValueWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.GetHostAuthenticationTokenValue(context.TODO(), gitpod.RunnerInteractionGetHostAuthenticationTokenValueParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion1),
		Host:                   gitpod.F("host"),
		PrincipalID:            gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		RunnerID:               gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRunnerInteractionGetLatestVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.GetLatestVersion(context.TODO(), gitpod.RunnerInteractionGetLatestVersionParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionGetLatestVersionParamsConnectProtocolVersion1),
		CurrentVersion:         gitpod.F("currentVersion"),
		InfrastructureVersion:  gitpod.F("infrastructureVersion"),
		RunnerID:               gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRunnerInteractionListRunnerEnvironmentClassesWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.ListRunnerEnvironmentClasses(context.TODO(), gitpod.RunnerInteractionListRunnerEnvironmentClassesParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion1),
		Filter: gitpod.F(gitpod.RunnerInteractionListRunnerEnvironmentClassesParamsFilter{
			EnvironmentClassIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.RunnerInteractionListRunnerEnvironmentClassesParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(0)),
		}),
		RunnerID:         gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ConnectTimeoutMs: gitpod.F(0.000000),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRunnerInteractionListRunnerScmIntegrationsWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.ListRunnerScmIntegrations(context.TODO(), gitpod.RunnerInteractionListRunnerScmIntegrationsParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion1),
		Filter: gitpod.F(gitpod.RunnerInteractionListRunnerScmIntegrationsParamsFilter{
			ScmIntegrationIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.RunnerInteractionListRunnerScmIntegrationsParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(0)),
		}),
		RunnerID:         gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ConnectTimeoutMs: gitpod.F(0.000000),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRunnerInteractionMarkActiveWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.MarkActive(context.TODO(), gitpod.RunnerInteractionMarkActiveParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionMarkActiveParamsConnectProtocolVersion1),
		RunnerID:               gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRunnerInteractionSendResponseWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.SendResponse(context.TODO(), gitpod.RunnerInteractionSendResponseParams{
		Body:                   gitpod.RunnerInteractionSendResponseParamsBody{},
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionSendResponseParamsConnectProtocolVersion1),
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

func TestRunnerInteractionSignupWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.Signup(context.TODO(), gitpod.RunnerInteractionSignupParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionSignupParamsConnectProtocolVersion1),
		EnvironmentClasses: gitpod.F([]gitpod.RunnerInteractionSignupParamsEnvironmentClass{{
			ID: gitpod.F("id"),
			Configuration: gitpod.F([]gitpod.RunnerInteractionSignupParamsEnvironmentClassesConfiguration{{
				Key:   gitpod.F("key"),
				Value: gitpod.F("value"),
			}}),
			Description: gitpod.F("xxx"),
			DisplayName: gitpod.F("xxx"),
			Enabled:     gitpod.F(true),
			RunnerID:    gitpod.F("runnerId"),
		}}),
		PublicKey:        gitpod.F("U3RhaW5sZXNzIHJvY2tz"),
		ConnectTimeoutMs: gitpod.F(0.000000),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRunnerInteractionUpdateRunnerConfigurationSchemaWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.UpdateRunnerConfigurationSchema(context.TODO(), gitpod.RunnerInteractionUpdateRunnerConfigurationSchemaParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion1),
		ConfigSchema: gitpod.F(gitpod.RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchema{
			EnvironmentClasses: gitpod.F([]gitpod.RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClass{{}}),
			RunnerConfig:       gitpod.F([]gitpod.RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfig{{}}),
			Scm: gitpod.F([]gitpod.RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScm{{
				DefaultHosts: gitpod.F([]string{"string"}),
				Name:         gitpod.F("name"),
				OAuth: gitpod.F(gitpod.RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmOAuth{
					CallbackURL: gitpod.F("callbackUrl"),
				}),
				Pat: gitpod.F(gitpod.RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmPat{
					Description: gitpod.F("description"),
					DocsLink:    gitpod.F("docsLink"),
				}),
				ScmID: gitpod.F("scmId"),
			}}),
			Version: gitpod.F("version"),
		}),
		RunnerID:         gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ConnectTimeoutMs: gitpod.F(0.000000),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRunnerInteractionUpdateStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.UpdateStatus(context.TODO(), gitpod.RunnerInteractionUpdateStatusParams{
		Body:                   gitpod.RunnerInteractionUpdateStatusParamsBody{},
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionUpdateStatusParamsConnectProtocolVersion1),
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
