// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/gitpod-io/flex-sdk-go"
	"github.com/gitpod-io/flex-sdk-go/internal/testutil"
	"github.com/gitpod-io/flex-sdk-go/option"
	"github.com/gitpod-io/flex-sdk-go/shared"
)

func TestEnvironmentAutomationServiceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Services.New(context.TODO(), gitpod.EnvironmentAutomationServiceNewParams{
		EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: gitpod.F(gitpod.ServiceMetadataParam{
			CreatedAt: gitpod.F(time.Now()),
			Creator: gitpod.F(shared.SubjectParam{
				ID:        gitpod.F("id"),
				Principal: gitpod.F(shared.PrincipalPrincipalUnspecified),
			}),
			Description: gitpod.F("description"),
			Name:        gitpod.F("x"),
			Reference:   gitpod.F("reference"),
			TriggeredBy: gitpod.F([]shared.AutomationTriggerParam{{
				Manual:                gitpod.F(true),
				PostDevcontainerStart: gitpod.F(true),
				PostEnvironmentStart:  gitpod.F(true),
			}}),
		}),
		Spec: gitpod.F(gitpod.ServiceSpecParam{
			Commands: gitpod.F(gitpod.ServiceSpecCommandsParam{
				Ready: gitpod.F("ready"),
				Start: gitpod.F("x"),
				Stop:  gitpod.F("stop"),
			}),
			DesiredPhase: gitpod.F(gitpod.ServicePhaseServicePhaseUnspecified),
			RunsOn: gitpod.F(shared.RunsOnParam{
				Docker: gitpod.F(shared.RunsOnDockerParam{
					Environment: gitpod.F([]string{"string"}),
					Image:       gitpod.F("x"),
				}),
			}),
			Session:     gitpod.F("session"),
			SpecVersion: gitpod.F("specVersion"),
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

func TestEnvironmentAutomationServiceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Services.Get(context.TODO(), gitpod.EnvironmentAutomationServiceGetParams{
		ID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnvironmentAutomationServiceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Services.Update(context.TODO(), gitpod.EnvironmentAutomationServiceUpdateParams{
		ID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: gitpod.F(gitpod.EnvironmentAutomationServiceUpdateParamsMetadata{
			Description: gitpod.F("description"),
			Name:        gitpod.F("x"),
			TriggeredBy: gitpod.F(gitpod.EnvironmentAutomationServiceUpdateParamsMetadataTriggeredBy{
				Trigger: gitpod.F([]shared.AutomationTriggerParam{{
					Manual:                gitpod.F(true),
					PostDevcontainerStart: gitpod.F(true),
					PostEnvironmentStart:  gitpod.F(true),
				}}),
			}),
		}),
		Spec: gitpod.F(gitpod.EnvironmentAutomationServiceUpdateParamsSpec{
			Commands: gitpod.F(gitpod.EnvironmentAutomationServiceUpdateParamsSpecCommands{
				Ready: gitpod.F("ready"),
				Start: gitpod.F("start"),
				Stop:  gitpod.F("stop"),
			}),
			RunsOn: gitpod.F(shared.RunsOnParam{
				Docker: gitpod.F(shared.RunsOnDockerParam{
					Environment: gitpod.F([]string{"string"}),
					Image:       gitpod.F("x"),
				}),
			}),
		}),
		Status: gitpod.F(gitpod.EnvironmentAutomationServiceUpdateParamsStatus{
			FailureMessage: gitpod.F("failureMessage"),
			LogURL:         gitpod.F("logUrl"),
			Phase:          gitpod.F(gitpod.ServicePhaseServicePhaseUnspecified),
			Session:        gitpod.F("session"),
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

func TestEnvironmentAutomationServiceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Services.List(context.TODO(), gitpod.EnvironmentAutomationServiceListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.EnvironmentAutomationServiceListParamsFilter{
			EnvironmentIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			References:     gitpod.F([]string{"x"}),
			ServiceIDs:     gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.EnvironmentAutomationServiceListParamsPagination{
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

func TestEnvironmentAutomationServiceDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Services.Delete(context.TODO(), gitpod.EnvironmentAutomationServiceDeleteParams{
		ID:    gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Force: gitpod.F(true),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnvironmentAutomationServiceStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Services.Start(context.TODO(), gitpod.EnvironmentAutomationServiceStartParams{
		ID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnvironmentAutomationServiceStopWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Services.Stop(context.TODO(), gitpod.EnvironmentAutomationServiceStopParams{
		ID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
