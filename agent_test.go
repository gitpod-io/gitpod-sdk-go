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
)

func TestAgentNewExecutionConversationTokenWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.NewExecutionConversationToken(context.TODO(), gitpod.AgentNewExecutionConversationTokenParams{
		AgentExecutionID: gitpod.F("6fa1a3c7-fbb7-49d1-ba56-1890dc7c4c35"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentNewPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.NewPrompt(context.TODO(), gitpod.AgentNewPromptParams{
		Command:     gitpod.F("command"),
		Description: gitpod.F("x"),
		IsCommand:   gitpod.F(true),
		IsSkill:     gitpod.F(true),
		IsTemplate:  gitpod.F(true),
		Name:        gitpod.F("x"),
		Prompt:      gitpod.F("x"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentDeleteExecutionWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.DeleteExecution(context.TODO(), gitpod.AgentDeleteExecutionParams{
		AgentExecutionID: gitpod.F("6fa1a3c7-fbb7-49d1-ba56-1890dc7c4c35"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentDeletePromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.DeletePrompt(context.TODO(), gitpod.AgentDeletePromptParams{
		PromptID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentListExecutionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.ListExecutions(context.TODO(), gitpod.AgentListExecutionsParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.AgentListExecutionsParamsFilter{
			AgentIDs:       gitpod.F([]string{"b8a64cfa-43e2-4b9d-9fb3-07edc63f5971"}),
			CreatorIDs:     gitpod.F([]string{"string"}),
			EnvironmentIDs: gitpod.F([]string{"string"}),
			ProjectIDs:     gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			Roles:          gitpod.F([]gitpod.AgentListExecutionsParamsFilterRole{gitpod.AgentListExecutionsParamsFilterRoleAgentExecutionRoleUnspecified}),
			StatusPhases:   gitpod.F([]gitpod.AgentListExecutionsParamsFilterStatusPhase{gitpod.AgentListExecutionsParamsFilterStatusPhasePhaseUnspecified}),
		}),
		Pagination: gitpod.F(gitpod.AgentListExecutionsParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(10)),
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

func TestAgentListPromptsWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.ListPrompts(context.TODO(), gitpod.AgentListPromptsParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.AgentListPromptsParamsFilter{
			Command:       gitpod.F("command"),
			CommandPrefix: gitpod.F("commandPrefix"),
			IsCommand:     gitpod.F(true),
			IsSkill:       gitpod.F(true),
			IsTemplate:    gitpod.F(true),
		}),
		Pagination: gitpod.F(gitpod.AgentListPromptsParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(10)),
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

func TestAgentGetExecutionWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.GetExecution(context.TODO(), gitpod.AgentGetExecutionParams{
		AgentExecutionID: gitpod.F("6fa1a3c7-fbb7-49d1-ba56-1890dc7c4c35"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentGetPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.GetPrompt(context.TODO(), gitpod.AgentGetPromptParams{
		PromptID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentSendToExecutionWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.SendToExecution(context.TODO(), gitpod.AgentSendToExecutionParams{
		AgentExecutionID: gitpod.F("6fa1a3c7-fbb7-49d1-ba56-1890dc7c4c35"),
		UserInput: gitpod.F(gitpod.UserInputBlockParam{
			Text: gitpod.F(gitpod.UserInputBlockTextParam{
				Content: gitpod.F("Generate a report based on the latest logs."),
			}),
			ID:        gitpod.F("id"),
			CreatedAt: gitpod.F(time.Now()),
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

func TestAgentStartExecutionWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.StartExecution(context.TODO(), gitpod.AgentStartExecutionParams{
		AgentID: gitpod.F("b8a64cfa-43e2-4b9d-9fb3-07edc63f5971"),
		CodeContext: gitpod.F(gitpod.AgentCodeContextParam{
			ContextURL: gitpod.F(gitpod.AgentCodeContextContextURLParam{
				EnvironmentClassID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				URL:                gitpod.F("https://example.com"),
			}),
			EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:     gitpod.F("2d22e4eb-31da-467f-882c-27e21550992f"),
			PullRequest: gitpod.F(gitpod.AgentCodeContextPullRequestParam{
				ID:         gitpod.F("id"),
				Author:     gitpod.F("author"),
				FromBranch: gitpod.F("fromBranch"),
				Repository: gitpod.F(gitpod.AgentCodeContextPullRequestRepositoryParam{
					CloneURL: gitpod.F("cloneUrl"),
					Host:     gitpod.F("host"),
					Name:     gitpod.F("name"),
					Owner:    gitpod.F("owner"),
				}),
				Title:    gitpod.F("title"),
				ToBranch: gitpod.F("toBranch"),
				URL:      gitpod.F("url"),
			}),
		}),
		Mode:             gitpod.F(gitpod.AgentModeUnspecified),
		Name:             gitpod.F("name"),
		WorkflowActionID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentStopExecutionWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.StopExecution(context.TODO(), gitpod.AgentStopExecutionParams{
		AgentExecutionID: gitpod.F("6fa1a3c7-fbb7-49d1-ba56-1890dc7c4c35"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentUpdatePromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.UpdatePrompt(context.TODO(), gitpod.AgentUpdatePromptParams{
		Metadata: gitpod.F(gitpod.AgentUpdatePromptParamsMetadata{
			Description: gitpod.F("x"),
			Name:        gitpod.F("name"),
		}),
		PromptID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Spec: gitpod.F(gitpod.AgentUpdatePromptParamsSpec{
			Command:    gitpod.F("command"),
			IsCommand:  gitpod.F(true),
			IsSkill:    gitpod.F(true),
			IsTemplate: gitpod.F(true),
			Prompt:     gitpod.F("prompt"),
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
