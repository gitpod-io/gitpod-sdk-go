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

func TestAutomationsFileUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.AutomationsFiles.Upsert(context.TODO(), gitpod.AutomationsFileUpsertParams{
		ConnectProtocolVersion: gitpod.F(gitpod.AutomationsFileUpsertParamsConnectProtocolVersion1),
		AutomationsFile: gitpod.F(gitpod.AutomationsFileUpsertParamsAutomationsFile{
			Services: gitpod.F(map[string]gitpod.AutomationsFileUpsertParamsAutomationsFileServices{
				"foo": {
					Commands: gitpod.F(gitpod.AutomationsFileUpsertParamsAutomationsFileServicesCommands{
						Ready: gitpod.F("ready"),
						Start: gitpod.F("x"),
						Stop:  gitpod.F("stop"),
					}),
					Description: gitpod.F("description"),
					Name:        gitpod.F("x"),
					TriggeredBy: gitpod.F([]string{"string", "string", "string"}),
				},
			}),
			Tasks: gitpod.F(map[string]gitpod.AutomationsFileUpsertParamsAutomationsFileTasks{
				"foo": {
					Command:     gitpod.F("x"),
					DependsOn:   gitpod.F([]string{"string", "string", "string"}),
					Description: gitpod.F("description"),
					Name:        gitpod.F("x"),
					TriggeredBy: gitpod.F([]string{"string", "string", "string"}),
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
