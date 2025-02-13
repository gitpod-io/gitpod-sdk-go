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

func TestEnvironmentAutomationTaskNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Tasks.New(context.TODO(), gitpod.EnvironmentAutomationTaskNewParams{
		DependsOn:     gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: gitpod.F(shared.TaskMetadataParam{
			CreatedAt: gitpod.F(time.Now()),
			Creator: gitpod.F(shared.SubjectParam{
				ID:        gitpod.F("id"),
				Principal: gitpod.F(shared.PrincipalUnspecified),
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
		Spec: gitpod.F(shared.TaskSpecParam{
			Command: gitpod.F("command"),
			RunsOn: gitpod.F(shared.RunsOnParam{
				Docker: gitpod.F(shared.RunsOnDockerParam{
					Environment: gitpod.F([]string{"string"}),
					Image:       gitpod.F("x"),
				}),
			}),
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

func TestEnvironmentAutomationTaskGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Tasks.Get(context.TODO(), gitpod.EnvironmentAutomationTaskGetParams{
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

func TestEnvironmentAutomationTaskUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Tasks.Update(context.TODO(), gitpod.EnvironmentAutomationTaskUpdateParams{
		ID:        gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DependsOn: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		Metadata: gitpod.F(gitpod.EnvironmentAutomationTaskUpdateParamsMetadata{
			Description: gitpod.F("description"),
			Name:        gitpod.F("x"),
			TriggeredBy: gitpod.F(gitpod.EnvironmentAutomationTaskUpdateParamsMetadataTriggeredBy{
				Trigger: gitpod.F([]shared.AutomationTriggerParam{{
					Manual:                gitpod.F(true),
					PostDevcontainerStart: gitpod.F(true),
					PostEnvironmentStart:  gitpod.F(true),
				}}),
			}),
		}),
		Spec: gitpod.F(gitpod.EnvironmentAutomationTaskUpdateParamsSpec{
			Command: gitpod.F("command"),
			RunsOn: gitpod.F(shared.RunsOnParam{
				Docker: gitpod.F(shared.RunsOnDockerParam{
					Environment: gitpod.F([]string{"string"}),
					Image:       gitpod.F("x"),
				}),
			}),
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

func TestEnvironmentAutomationTaskListWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Tasks.List(context.TODO(), gitpod.EnvironmentAutomationTaskListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.EnvironmentAutomationTaskListParamsFilter{
			EnvironmentIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			References:     gitpod.F([]string{"x"}),
			TaskIDs:        gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.EnvironmentAutomationTaskListParamsPagination{
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

func TestEnvironmentAutomationTaskDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Tasks.Delete(context.TODO(), gitpod.EnvironmentAutomationTaskDeleteParams{
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

func TestEnvironmentAutomationTaskStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Tasks.Start(context.TODO(), gitpod.EnvironmentAutomationTaskStartParams{
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
