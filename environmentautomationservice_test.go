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

func TestEnvironmentAutomationServiceUpdateWithOptionalParams(t *testing.T) {
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
		option.WithConnectProtocolVersion(true),
		option.WithConnectTimeoutHeader(0),
	)
	_, err := client.Environments.Automations.Services.Update(context.TODO(), gitpod.EnvironmentAutomationServiceUpdateParams{
		ID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: gitpod.F[gitpod.EnvironmentAutomationServiceUpdateParamsMetadataUnion](gitpod.EnvironmentAutomationServiceUpdateParamsMetadataDescription{
			Description: gitpod.F("description"),
		}),
		Spec: gitpod.F[gitpod.EnvironmentAutomationServiceUpdateParamsSpecUnion](gitpod.EnvironmentAutomationServiceUpdateParamsSpecCommands{
			Commands: gitpod.F[gitpod.EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsUnion](gitpod.EnvironmentAutomationServiceUpdateParamsSpecCommandsCommandsReady{
				Ready: gitpod.F("ready"),
			}),
		}),
		Status: gitpod.F[gitpod.EnvironmentAutomationServiceUpdateParamsStatusUnion](gitpod.EnvironmentAutomationServiceUpdateParamsStatusFailureMessage{
			FailureMessage: gitpod.F("failureMessage"),
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
		option.WithConnectProtocolVersion(true),
		option.WithConnectTimeoutHeader(0),
	)
	_, err := client.Environments.Automations.Services.List(context.TODO(), gitpod.EnvironmentAutomationServiceListParams{
		Filter: gitpod.F(gitpod.EnvironmentAutomationServiceListParamsFilter{
			EnvironmentIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			References:     gitpod.F([]string{"x"}),
			ServiceIDs:     gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.EnvironmentAutomationServiceListParamsPagination{
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

func TestEnvironmentAutomationServiceDeleteWithOptionalParams(t *testing.T) {
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
		option.WithConnectProtocolVersion(true),
		option.WithConnectTimeoutHeader(0),
	)
	_, err := client.Environments.Automations.Services.Delete(context.TODO(), gitpod.EnvironmentAutomationServiceDeleteParams{
		ID:    gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Force: gitpod.F(true),
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
		option.WithConnectProtocolVersion(true),
		option.WithConnectTimeoutHeader(0),
	)
	_, err := client.Environments.Automations.Services.Start(context.TODO(), gitpod.EnvironmentAutomationServiceStartParams{
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

func TestEnvironmentAutomationServiceStopWithOptionalParams(t *testing.T) {
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
		option.WithConnectProtocolVersion(true),
		option.WithConnectTimeoutHeader(0),
	)
	_, err := client.Environments.Automations.Services.Stop(context.TODO(), gitpod.EnvironmentAutomationServiceStopParams{
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
