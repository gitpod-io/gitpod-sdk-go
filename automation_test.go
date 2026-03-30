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

func TestAutomationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.New(context.TODO(), gitpod.AutomationNewParams{
		Action: gitpod.F(gitpod.WorkflowActionParam{
			Limits: gitpod.F(gitpod.WorkflowActionLimitsParam{
				MaxParallel: gitpod.F(int64(0)),
				MaxTotal:    gitpod.F(int64(0)),
				PerExecution: gitpod.F(gitpod.WorkflowActionLimitsPerExecutionParam{
					MaxTime: gitpod.F("+9125115.360s"),
				}),
			}),
			Steps: gitpod.F([]gitpod.WorkflowStepParam{{
				Agent: gitpod.F(gitpod.WorkflowStepAgentParam{
					Prompt: gitpod.F("prompt"),
				}),
				PullRequest: gitpod.F(gitpod.WorkflowStepPullRequestParam{
					Branch:      gitpod.F("branch"),
					Description: gitpod.F("description"),
					Draft:       gitpod.F(true),
					Title:       gitpod.F("title"),
				}),
				Report: gitpod.F(gitpod.WorkflowStepReportParam{
					Outputs: gitpod.F([]gitpod.WorkflowStepReportOutputParam{{}}),
				}),
				Task: gitpod.F(gitpod.WorkflowStepTaskParam{
					Command: gitpod.F("command"),
				}),
			}}),
		}),
		Description: gitpod.F("description"),
		Executor: gitpod.F(shared.SubjectParam{
			ID:        gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Principal: gitpod.F(shared.PrincipalUnspecified),
		}),
		Name: gitpod.F("name"),
		Report: gitpod.F(gitpod.WorkflowActionParam{
			Limits: gitpod.F(gitpod.WorkflowActionLimitsParam{
				MaxParallel: gitpod.F(int64(0)),
				MaxTotal:    gitpod.F(int64(0)),
				PerExecution: gitpod.F(gitpod.WorkflowActionLimitsPerExecutionParam{
					MaxTime: gitpod.F("+9125115.360s"),
				}),
			}),
			Steps: gitpod.F([]gitpod.WorkflowStepParam{{
				Agent: gitpod.F(gitpod.WorkflowStepAgentParam{
					Prompt: gitpod.F("prompt"),
				}),
				PullRequest: gitpod.F(gitpod.WorkflowStepPullRequestParam{
					Branch:      gitpod.F("branch"),
					Description: gitpod.F("description"),
					Draft:       gitpod.F(true),
					Title:       gitpod.F("title"),
				}),
				Report: gitpod.F(gitpod.WorkflowStepReportParam{
					Outputs: gitpod.F([]gitpod.WorkflowStepReportOutputParam{{}}),
				}),
				Task: gitpod.F(gitpod.WorkflowStepTaskParam{
					Command: gitpod.F("command"),
				}),
			}}),
		}),
		Triggers: gitpod.F([]gitpod.WorkflowTriggerParam{{
			Context: gitpod.F(gitpod.WorkflowTriggerContextParam{
				Agent: gitpod.F(gitpod.WorkflowTriggerContextAgentParam{
					Prompt: gitpod.F("prompt"),
				}),
				FromTrigger: gitpod.F[any](map[string]interface{}{}),
				Projects: gitpod.F(gitpod.WorkflowTriggerContextProjectsParam{
					ProjectIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				}),
				Repositories: gitpod.F(gitpod.WorkflowTriggerContextRepositoriesParam{
					EnvironmentClassID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					RepoSelector: gitpod.F(gitpod.WorkflowTriggerContextRepositoriesRepoSelectorParam{
						RepoSearchString: gitpod.F("x"),
						ScmHost:          gitpod.F("x"),
					}),
					RepositoryURLs: gitpod.F(gitpod.WorkflowTriggerContextRepositoriesRepositoryURLsParam{
						RepoURLs: gitpod.F([]string{"x"}),
					}),
				}),
			}),
			Manual: gitpod.F[any](map[string]interface{}{}),
			PullRequest: gitpod.F(gitpod.WorkflowTriggerPullRequestParam{
				Events:    gitpod.F([]gitpod.WorkflowTriggerPullRequestEvent{gitpod.WorkflowTriggerPullRequestEventPullRequestEventUnspecified}),
				WebhookID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			Time: gitpod.F(gitpod.WorkflowTriggerTimeParam{
				CronExpression: gitpod.F("cronExpression"),
			}),
		}}),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAutomationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.Get(context.TODO(), gitpod.AutomationGetParams{
		WorkflowID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAutomationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.Update(context.TODO(), gitpod.AutomationUpdateParams{
		Action: gitpod.F(gitpod.WorkflowActionParam{
			Limits: gitpod.F(gitpod.WorkflowActionLimitsParam{
				MaxParallel: gitpod.F(int64(0)),
				MaxTotal:    gitpod.F(int64(0)),
				PerExecution: gitpod.F(gitpod.WorkflowActionLimitsPerExecutionParam{
					MaxTime: gitpod.F("+9125115.360s"),
				}),
			}),
			Steps: gitpod.F([]gitpod.WorkflowStepParam{{
				Agent: gitpod.F(gitpod.WorkflowStepAgentParam{
					Prompt: gitpod.F("prompt"),
				}),
				PullRequest: gitpod.F(gitpod.WorkflowStepPullRequestParam{
					Branch:      gitpod.F("branch"),
					Description: gitpod.F("description"),
					Draft:       gitpod.F(true),
					Title:       gitpod.F("title"),
				}),
				Report: gitpod.F(gitpod.WorkflowStepReportParam{
					Outputs: gitpod.F([]gitpod.WorkflowStepReportOutputParam{{}}),
				}),
				Task: gitpod.F(gitpod.WorkflowStepTaskParam{
					Command: gitpod.F("command"),
				}),
			}}),
		}),
		Description: gitpod.F("description"),
		Executor: gitpod.F(shared.SubjectParam{
			ID:        gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Principal: gitpod.F(shared.PrincipalUnspecified),
		}),
		Name: gitpod.F("name"),
		Report: gitpod.F(gitpod.WorkflowActionParam{
			Limits: gitpod.F(gitpod.WorkflowActionLimitsParam{
				MaxParallel: gitpod.F(int64(0)),
				MaxTotal:    gitpod.F(int64(0)),
				PerExecution: gitpod.F(gitpod.WorkflowActionLimitsPerExecutionParam{
					MaxTime: gitpod.F("+9125115.360s"),
				}),
			}),
			Steps: gitpod.F([]gitpod.WorkflowStepParam{{
				Agent: gitpod.F(gitpod.WorkflowStepAgentParam{
					Prompt: gitpod.F("prompt"),
				}),
				PullRequest: gitpod.F(gitpod.WorkflowStepPullRequestParam{
					Branch:      gitpod.F("branch"),
					Description: gitpod.F("description"),
					Draft:       gitpod.F(true),
					Title:       gitpod.F("title"),
				}),
				Report: gitpod.F(gitpod.WorkflowStepReportParam{
					Outputs: gitpod.F([]gitpod.WorkflowStepReportOutputParam{{}}),
				}),
				Task: gitpod.F(gitpod.WorkflowStepTaskParam{
					Command: gitpod.F("command"),
				}),
			}}),
		}),
		Triggers: gitpod.F([]gitpod.WorkflowTriggerParam{{
			Context: gitpod.F(gitpod.WorkflowTriggerContextParam{
				Agent: gitpod.F(gitpod.WorkflowTriggerContextAgentParam{
					Prompt: gitpod.F("prompt"),
				}),
				FromTrigger: gitpod.F[any](map[string]interface{}{}),
				Projects: gitpod.F(gitpod.WorkflowTriggerContextProjectsParam{
					ProjectIDs: gitpod.F([]string{"new-project-id"}),
				}),
				Repositories: gitpod.F(gitpod.WorkflowTriggerContextRepositoriesParam{
					EnvironmentClassID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					RepoSelector: gitpod.F(gitpod.WorkflowTriggerContextRepositoriesRepoSelectorParam{
						RepoSearchString: gitpod.F("x"),
						ScmHost:          gitpod.F("x"),
					}),
					RepositoryURLs: gitpod.F(gitpod.WorkflowTriggerContextRepositoriesRepositoryURLsParam{
						RepoURLs: gitpod.F([]string{"x"}),
					}),
				}),
			}),
			Manual: gitpod.F[any](map[string]interface{}{}),
			PullRequest: gitpod.F(gitpod.WorkflowTriggerPullRequestParam{
				Events:    gitpod.F([]gitpod.WorkflowTriggerPullRequestEvent{gitpod.WorkflowTriggerPullRequestEventPullRequestEventUnspecified}),
				WebhookID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			Time: gitpod.F(gitpod.WorkflowTriggerTimeParam{
				CronExpression: gitpod.F("cronExpression"),
			}),
		}}),
		WorkflowID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAutomationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.List(context.TODO(), gitpod.AutomationListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.AutomationListParamsFilter{
			CreatorIDs:              gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			HasFailedExecutionSince: gitpod.F(time.Now()),
			Search:                  gitpod.F("search"),
			StatusPhases:            gitpod.F([]gitpod.AutomationListParamsFilterStatusPhase{gitpod.AutomationListParamsFilterStatusPhaseWorkflowExecutionPhaseUnspecified}),
			WorkflowIDs:             gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.AutomationListParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(100)),
		}),
		Sort: gitpod.F(gitpod.AutomationListParamsSort{
			Field: gitpod.F(gitpod.AutomationListParamsSortFieldSortFieldUnspecified),
			Order: gitpod.F(shared.SortOrderUnspecified),
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

func TestAutomationDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.Delete(context.TODO(), gitpod.AutomationDeleteParams{
		Force:      gitpod.F(true),
		WorkflowID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAutomationCancelExecutionWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.CancelExecution(context.TODO(), gitpod.AutomationCancelExecutionParams{
		WorkflowExecutionID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAutomationCancelExecutionActionWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.CancelExecutionAction(context.TODO(), gitpod.AutomationCancelExecutionActionParams{
		WorkflowExecutionActionID: gitpod.F("a1b2c3d4-5e6f-7890-abcd-ef1234567890"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAutomationListExecutionActionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.ListExecutionActions(context.TODO(), gitpod.AutomationListExecutionActionsParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.AutomationListExecutionActionsParamsFilter{
			Phases:                     gitpod.F([]gitpod.AutomationListExecutionActionsParamsFilterPhase{gitpod.AutomationListExecutionActionsParamsFilterPhaseWorkflowExecutionActionPhaseUnspecified}),
			WorkflowExecutionActionIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			WorkflowExecutionIDs:       gitpod.F([]string{"d2c94c27-3b76-4a42-b88c-95a85e392c68"}),
			WorkflowIDs:                gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.AutomationListExecutionActionsParamsPagination{
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

func TestAutomationListExecutionOutputsWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.ListExecutionOutputs(context.TODO(), gitpod.AutomationListExecutionOutputsParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.AutomationListExecutionOutputsParamsFilter{
			WorkflowExecutionIDs: gitpod.F([]string{"d2c94c27-3b76-4a42-b88c-95a85e392c68"}),
		}),
		Pagination: gitpod.F(gitpod.AutomationListExecutionOutputsParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(50)),
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

func TestAutomationListExecutionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.ListExecutions(context.TODO(), gitpod.AutomationListExecutionsParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.AutomationListExecutionsParamsFilter{
			HasFailedActions:     gitpod.F(true),
			Search:               gitpod.F("search"),
			StatusPhases:         gitpod.F([]gitpod.AutomationListExecutionsParamsFilterStatusPhase{gitpod.AutomationListExecutionsParamsFilterStatusPhaseWorkflowExecutionPhaseUnspecified}),
			WorkflowExecutionIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			WorkflowIDs:          gitpod.F([]string{"b0e12f6c-4c67-429d-a4a6-d9838b5da047"}),
		}),
		Pagination: gitpod.F(gitpod.AutomationListExecutionsParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(20)),
		}),
		Sort: gitpod.F(shared.SortParam{
			Field: gitpod.F("field"),
			Order: gitpod.F(shared.SortOrderUnspecified),
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

func TestAutomationGetExecutionWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.GetExecution(context.TODO(), gitpod.AutomationGetExecutionParams{
		WorkflowExecutionID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAutomationGetExecutionActionWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.GetExecutionAction(context.TODO(), gitpod.AutomationGetExecutionActionParams{
		WorkflowExecutionActionID: gitpod.F("a1b2c3d4-5e6f-7890-abcd-ef1234567890"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAutomationStartExecutionWithOptionalParams(t *testing.T) {
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
	_, err := client.Automations.StartExecution(context.TODO(), gitpod.AutomationStartExecutionParams{
		ContextOverride: gitpod.F(gitpod.WorkflowTriggerContextParam{
			Agent: gitpod.F(gitpod.WorkflowTriggerContextAgentParam{
				Prompt: gitpod.F("prompt"),
			}),
			FromTrigger: gitpod.F[any](map[string]interface{}{}),
			Projects: gitpod.F(gitpod.WorkflowTriggerContextProjectsParam{
				ProjectIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			}),
			Repositories: gitpod.F(gitpod.WorkflowTriggerContextRepositoriesParam{
				EnvironmentClassID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RepoSelector: gitpod.F(gitpod.WorkflowTriggerContextRepositoriesRepoSelectorParam{
					RepoSearchString: gitpod.F("x"),
					ScmHost:          gitpod.F("x"),
				}),
				RepositoryURLs: gitpod.F(gitpod.WorkflowTriggerContextRepositoriesRepositoryURLsParam{
					RepoURLs: gitpod.F([]string{"x"}),
				}),
			}),
		}),
		Parameters: gitpod.F(map[string]string{
			"foo": "string",
		}),
		WorkflowID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
