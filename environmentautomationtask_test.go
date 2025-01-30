// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/gitpod-go"
	"github.com/stainless-sdks/gitpod-go/internal/testutil"
	"github.com/stainless-sdks/gitpod-go/option"
)

func TestEnvironmentAutomationTaskNewWithOptionalParams(t *testing.T) {
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
		option.WithConnectProtocolVersion(0),
		option.WithConnectTimeoutHeader(0),
	)
	_, err := client.Environments.Automations.Tasks.New(context.TODO(), gitpod.EnvironmentAutomationTaskNewParams{
		DependsOn:     gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsMetadata{
			CreatedAt: gitpod.F(time.Now()),
			Creator: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsMetadataCreator{
				ID:        gitpod.F("id"),
				Principal: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalUnspecified),
			}),
			Description: gitpod.F("description"),
			Name:        gitpod.F("x"),
			Reference:   gitpod.F("reference"),
			TriggeredBy: gitpod.F([]gitpod.EnvironmentAutomationTaskNewParamsMetadataTriggeredByUnion{gitpod.EnvironmentAutomationTaskNewParamsMetadataTriggeredByManual{
				Manual: gitpod.F(true),
			}}),
		}),
		Spec: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsSpec{
			Command: gitpod.F("command"),
			RunsOn: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsSpecRunsOn{
				Docker: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsSpecRunsOnDocker{
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
		option.WithConnectProtocolVersion(0),
		option.WithConnectTimeoutHeader(0),
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
		option.WithConnectProtocolVersion(0),
		option.WithConnectTimeoutHeader(0),
	)
	_, err := client.Environments.Automations.Tasks.Update(context.TODO(), gitpod.EnvironmentAutomationTaskUpdateParams{
		ID:        gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DependsOn: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		Metadata: gitpod.F[gitpod.EnvironmentAutomationTaskUpdateParamsMetadataUnion](gitpod.EnvironmentAutomationTaskUpdateParamsMetadataDescription{
			Description: gitpod.F("description"),
		}),
		Spec: gitpod.F[gitpod.EnvironmentAutomationTaskUpdateParamsSpecUnion](gitpod.EnvironmentAutomationTaskUpdateParamsSpecCommand{
			Command: gitpod.F("command"),
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
		option.WithConnectProtocolVersion(0),
		option.WithConnectTimeoutHeader(0),
	)
	_, err := client.Environments.Automations.Tasks.List(context.TODO(), gitpod.EnvironmentAutomationTaskListParams{
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
		option.WithConnectProtocolVersion(0),
		option.WithConnectTimeoutHeader(0),
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
		option.WithConnectProtocolVersion(0),
		option.WithConnectTimeoutHeader(0),
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
