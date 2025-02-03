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

func TestEnvironmentAutomationUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Upsert(context.TODO(), gitpod.EnvironmentAutomationUpsertParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationUpsertParamsConnectProtocolVersion1),
		AutomationsFile: gitpod.F(gitpod.EnvironmentAutomationUpsertParamsAutomationsFile{
			Services: gitpod.F(map[string]gitpod.EnvironmentAutomationUpsertParamsAutomationsFileServices{
				"foo": {
					Commands: gitpod.F(gitpod.EnvironmentAutomationUpsertParamsAutomationsFileServicesCommands{
						Ready: gitpod.F("ready"),
						Start: gitpod.F("x"),
						Stop:  gitpod.F("stop"),
					}),
					Description: gitpod.F("description"),
					Name:        gitpod.F("x"),
					RunsOn: gitpod.F[gitpod.EnvironmentAutomationUpsertParamsAutomationsFileServicesRunsOnUnion](gitpod.EnvironmentAutomationUpsertParamsAutomationsFileServicesRunsOnDocker{
						Docker: gitpod.F(gitpod.EnvironmentAutomationUpsertParamsAutomationsFileServicesRunsOnDockerDocker{
							Environment: gitpod.F([]string{"string"}),
							Image:       gitpod.F("x"),
						}),
					}),
					TriggeredBy: gitpod.F([]string{"string"}),
				},
			}),
			Tasks: gitpod.F(map[string]gitpod.EnvironmentAutomationUpsertParamsAutomationsFileTasks{
				"foo": {
					Command:     gitpod.F("x"),
					DependsOn:   gitpod.F([]string{"string"}),
					Description: gitpod.F("description"),
					Name:        gitpod.F("x"),
					RunsOn: gitpod.F[gitpod.EnvironmentAutomationUpsertParamsAutomationsFileTasksRunsOnUnion](gitpod.EnvironmentAutomationUpsertParamsAutomationsFileTasksRunsOnDocker{
						Docker: gitpod.F(gitpod.EnvironmentAutomationUpsertParamsAutomationsFileTasksRunsOnDockerDocker{
							Environment: gitpod.F([]string{"string"}),
							Image:       gitpod.F("x"),
						}),
					}),
					TriggeredBy: gitpod.F([]string{"string"}),
				},
			}),
		}),
		EnvironmentID:    gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
