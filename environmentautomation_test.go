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

func TestEnvironmentAutomationUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Automations.Upsert(context.TODO(), gitpod.EnvironmentAutomationUpsertParams{
		AutomationsFile: gitpod.F(gitpod.AutomationsFileParam{
			Services: gitpod.F(map[string]gitpod.AutomationsFileServiceParam{
				"web-server": {
					Commands: gitpod.F(gitpod.AutomationsFileServicesCommandsParam{
						Ready: gitpod.F("curl -s http://localhost:3000"),
						Start: gitpod.F("npm run dev"),
						Stop:  gitpod.F("stop"),
					}),
					Description: gitpod.F("Development web server"),
					Name:        gitpod.F("Web Server"),
					RunsOn: gitpod.F(shared.RunsOnParam{
						Docker: gitpod.F(shared.RunsOnDockerParam{
							Environment: gitpod.F([]string{"string"}),
							Image:       gitpod.F("x"),
						}),
					}),
					TriggeredBy: gitpod.F([]gitpod.AutomationsFileServicesTriggeredBy{gitpod.AutomationsFileServicesTriggeredByPostDevcontainerStart}),
				},
			}),
			Tasks: gitpod.F(map[string]gitpod.AutomationsFileTaskParam{
				"build": {
					Command:     gitpod.F("npm run build"),
					DependsOn:   gitpod.F([]string{"string"}),
					Description: gitpod.F("Builds the project artifacts"),
					Name:        gitpod.F("Build Project"),
					RunsOn: gitpod.F(shared.RunsOnParam{
						Docker: gitpod.F(shared.RunsOnDockerParam{
							Environment: gitpod.F([]string{"string"}),
							Image:       gitpod.F("x"),
						}),
					}),
					TriggeredBy: gitpod.F([]gitpod.AutomationsFileTasksTriggeredBy{gitpod.AutomationsFileTasksTriggeredByPostEnvironmentStart}),
				},
			}),
		}),
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
