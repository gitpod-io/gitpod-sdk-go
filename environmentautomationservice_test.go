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

func TestEnvironmentAutomationServiceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Services.New(context.TODO(), gitpod.EnvironmentAutomationServiceNewParams{
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
		Metadata: gitpod.F(gitpod.ServiceMetadataParam{
			CreatedAt: gitpod.F(time.Now()),
			Creator: gitpod.F(shared.SubjectParam{
				ID:        gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Principal: gitpod.F(shared.PrincipalUnspecified),
			}),
			Description: gitpod.F("Runs the development web server"),
			Name:        gitpod.F("Web Server"),
			Reference:   gitpod.F("web-server"),
			Role:        gitpod.F(gitpod.ServiceRoleUnspecified),
			TriggeredBy: gitpod.F([]shared.AutomationTriggerParam{{
				BeforeSnapshot:        gitpod.F(true),
				Manual:                gitpod.F(true),
				PostDevcontainerStart: gitpod.F(true),
				PostEnvironmentStart:  gitpod.F(true),
				PostMachineStart:      gitpod.F(true),
				Prebuild:              gitpod.F(true),
			}}),
		}),
		Spec: gitpod.F(gitpod.ServiceSpecParam{
			Commands: gitpod.F(gitpod.ServiceSpecCommandsParam{
				Ready: gitpod.F("curl -s http://localhost:3000"),
				Start: gitpod.F("npm run dev"),
				Stop:  gitpod.F("stop"),
			}),
			DesiredPhase: gitpod.F(gitpod.ServicePhaseUnspecified),
			Env: gitpod.F([]shared.EnvironmentVariableItemParam{{
				Name:  gitpod.F("x"),
				Value: gitpod.F("value"),
				ValueFrom: gitpod.F(shared.EnvironmentVariableSourceParam{
					SecretRef: gitpod.F(shared.SecretRefParam{
						ID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					}),
				}),
			}}),
			RunsOn: gitpod.F(shared.RunsOnParam{
				Docker: gitpod.F(shared.RunsOnDockerParam{
					Environment: gitpod.F([]string{"string"}),
					Image:       gitpod.F("x"),
				}),
				Machine: gitpod.F[any](map[string]interface{}{}),
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
	_, err := client.Environments.Automations.Services.Get(context.TODO(), gitpod.EnvironmentAutomationServiceGetParams{
		ID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
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
	_, err := client.Environments.Automations.Services.Update(context.TODO(), gitpod.EnvironmentAutomationServiceUpdateParams{
		ID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
		Metadata: gitpod.F(gitpod.EnvironmentAutomationServiceUpdateParamsMetadata{
			Description: gitpod.F("description"),
			Name:        gitpod.F("x"),
			Role:        gitpod.F(gitpod.ServiceRoleUnspecified),
			TriggeredBy: gitpod.F(gitpod.EnvironmentAutomationServiceUpdateParamsMetadataTriggeredBy{
				Trigger: gitpod.F([]shared.AutomationTriggerParam{{
					BeforeSnapshot:        gitpod.F(true),
					Manual:                gitpod.F(true),
					PostDevcontainerStart: gitpod.F(true),
					PostEnvironmentStart:  gitpod.F(true),
					PostMachineStart:      gitpod.F(true),
					Prebuild:              gitpod.F(true),
				}}),
			}),
		}),
		Spec: gitpod.F(gitpod.EnvironmentAutomationServiceUpdateParamsSpec{
			Commands: gitpod.F(gitpod.EnvironmentAutomationServiceUpdateParamsSpecCommands{
				Ready: gitpod.F("curl -s http://localhost:8080"),
				Start: gitpod.F("npm run start:dev"),
				Stop:  gitpod.F("stop"),
			}),
			Env: gitpod.F([]shared.EnvironmentVariableItemParam{{
				Name:  gitpod.F("x"),
				Value: gitpod.F("value"),
				ValueFrom: gitpod.F(shared.EnvironmentVariableSourceParam{
					SecretRef: gitpod.F(shared.SecretRefParam{
						ID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					}),
				}),
			}}),
			RunsOn: gitpod.F(shared.RunsOnParam{
				Docker: gitpod.F(shared.RunsOnDockerParam{
					Environment: gitpod.F([]string{"string"}),
					Image:       gitpod.F("x"),
				}),
				Machine: gitpod.F[any](map[string]interface{}{}),
			}),
		}),
		Status: gitpod.F(gitpod.EnvironmentAutomationServiceUpdateParamsStatus{
			FailureMessage: gitpod.F("failureMessage"),
			LogURL:         gitpod.F("logUrl"),
			Output: gitpod.F(map[string]string{
				"foo": "string",
			}),
			Phase:   gitpod.F(gitpod.ServicePhaseUnspecified),
			Session: gitpod.F("session"),
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
	_, err := client.Environments.Automations.Services.List(context.TODO(), gitpod.EnvironmentAutomationServiceListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.EnvironmentAutomationServiceListParamsFilter{
			EnvironmentIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			References:     gitpod.F([]string{"web-server", "database"}),
			Roles:          gitpod.F([]gitpod.ServiceRole{gitpod.ServiceRoleUnspecified}),
			ServiceIDs:     gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.EnvironmentAutomationServiceListParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(20)),
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
	_, err := client.Environments.Automations.Services.Delete(context.TODO(), gitpod.EnvironmentAutomationServiceDeleteParams{
		ID:    gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
		Force: gitpod.F(false),
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
	_, err := client.Environments.Automations.Services.Start(context.TODO(), gitpod.EnvironmentAutomationServiceStartParams{
		ID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
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
	_, err := client.Environments.Automations.Services.Stop(context.TODO(), gitpod.EnvironmentAutomationServiceStopParams{
		ID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
