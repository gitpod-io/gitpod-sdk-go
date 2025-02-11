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
				"foo": {
					Commands: gitpod.F(gitpod.AutomationsFileServicesCommandsParam{
						Ready: gitpod.F("ready"),
						Start: gitpod.F("x"),
						Stop:  gitpod.F("stop"),
					}),
					Description: gitpod.F("description"),
					Name:        gitpod.F("x"),
					RunsOn: gitpod.F(shared.RunsOnParam{
						Docker: gitpod.F(shared.RunsOnDockerParam{
							Environment: gitpod.F([]string{"string"}),
							Image:       gitpod.F("x"),
						}),
					}),
					TriggeredBy: gitpod.F([]gitpod.AutomationsFileServicesTriggeredBy{gitpod.AutomationsFileServicesTriggeredByManual}),
				},
			}),
			Tasks: gitpod.F(map[string]gitpod.AutomationsFileTaskParam{
				"foo": {
					Command:     gitpod.F("x"),
					DependsOn:   gitpod.F([]string{"string"}),
					Description: gitpod.F("description"),
					Name:        gitpod.F("x"),
					RunsOn: gitpod.F(shared.RunsOnParam{
						Docker: gitpod.F(shared.RunsOnDockerParam{
							Environment: gitpod.F([]string{"string"}),
							Image:       gitpod.F("x"),
						}),
					}),
					TriggeredBy: gitpod.F([]gitpod.AutomationsFileTasksTriggeredBy{gitpod.AutomationsFileTasksTriggeredByManual}),
				},
			}),
		}),
		EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
