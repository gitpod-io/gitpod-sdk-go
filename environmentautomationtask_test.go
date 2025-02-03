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
	)
	_, err := client.Environments.Automations.Tasks.New(context.TODO(), gitpod.EnvironmentAutomationTaskNewParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsConnectProtocolVersion1),
		DependsOn:              gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		EnvironmentID:          gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsMetadata{
			CreatedAt: gitpod.F(time.Now()),
			Creator: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsMetadataCreator{
				ID:        gitpod.F("id"),
				Principal: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsMetadataCreatorPrincipalPrincipalUnspecified),
			}),
			Description: gitpod.F("description"),
			Name:        gitpod.F("x"),
			Reference:   gitpod.F("reference"),
			TriggeredBy: gitpod.F([]gitpod.EnvironmentAutomationTaskNewParamsMetadataTriggeredByUnion{gitpod.EnvironmentAutomationTaskNewParamsMetadataTriggeredByObject{
				Manual:                gitpod.F(true),
				PostDevcontainerStart: gitpod.F(true),
				PostEnvironmentStart:  gitpod.F(true),
			}}),
		}),
		Spec: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsSpec{
			Command: gitpod.F("command"),
			RunsOn: gitpod.F[gitpod.EnvironmentAutomationTaskNewParamsSpecRunsOnUnion](gitpod.EnvironmentAutomationTaskNewParamsSpecRunsOnDocker{
				Docker: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsSpecRunsOnDockerDocker{
					Environment: gitpod.F([]string{"string"}),
					Image:       gitpod.F("x"),
				}),
			}),
		}),
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
	)
	_, err := client.Environments.Automations.Tasks.Get(context.TODO(), gitpod.EnvironmentAutomationTaskGetParams{
		Encoding:               gitpod.F(gitpod.EnvironmentAutomationTaskGetParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskGetParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.EnvironmentAutomationTaskGetParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.EnvironmentAutomationTaskGetParamsConnectV1),
		Message:                gitpod.F("message"),
		ConnectTimeoutMs:       gitpod.F(0.000000),
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
	)
	_, err := client.Environments.Automations.Tasks.Update(context.TODO(), gitpod.EnvironmentAutomationTaskUpdateParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskUpdateParamsConnectProtocolVersion1),
		ID:                     gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DependsOn:              gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		Metadata:               gitpod.F(gitpod.EnvironmentAutomationTaskUpdateParamsMetadata{}),
		Spec:                   gitpod.F(gitpod.EnvironmentAutomationTaskUpdateParamsSpec{}),
		ConnectTimeoutMs:       gitpod.F(0.000000),
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
	)
	_, err := client.Environments.Automations.Tasks.List(context.TODO(), gitpod.EnvironmentAutomationTaskListParams{
		Encoding:               gitpod.F(gitpod.EnvironmentAutomationTaskListParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskListParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.EnvironmentAutomationTaskListParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.EnvironmentAutomationTaskListParamsConnectV1),
		Message:                gitpod.F("message"),
		ConnectTimeoutMs:       gitpod.F(0.000000),
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
	)
	_, err := client.Environments.Automations.Tasks.Delete(context.TODO(), gitpod.EnvironmentAutomationTaskDeleteParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskDeleteParamsConnectProtocolVersion1),
		ID:                     gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ConnectTimeoutMs:       gitpod.F(0.000000),
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
	)
	_, err := client.Environments.Automations.Tasks.Start(context.TODO(), gitpod.EnvironmentAutomationTaskStartParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskStartParamsConnectProtocolVersion1),
		ID:                     gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ConnectTimeoutMs:       gitpod.F(0.000000),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
