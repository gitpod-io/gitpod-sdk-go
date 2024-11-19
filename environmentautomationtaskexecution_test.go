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

func TestEnvironmentAutomationTaskExecutionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.EnvironmentAutomations.TaskExecutions.Get(context.TODO(), gitpod.EnvironmentAutomationTaskExecutionGetParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskExecutionGetParamsConnectProtocolVersion1),
		Base64:                 gitpod.F("base64"),
		Compression:            gitpod.F("compression"),
		Connect:                gitpod.F("connect"),
		Encoding:               gitpod.F("encoding"),
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

func TestEnvironmentAutomationTaskExecutionListWithOptionalParams(t *testing.T) {
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
	_, err := client.EnvironmentAutomations.TaskExecutions.List(context.TODO(), gitpod.EnvironmentAutomationTaskExecutionListParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskExecutionListParamsConnectProtocolVersion1),
		Base64:                 gitpod.F("base64"),
		Compression:            gitpod.F("compression"),
		Connect:                gitpod.F("connect"),
		Encoding:               gitpod.F("encoding"),
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

func TestEnvironmentAutomationTaskExecutionNewListWithOptionalParams(t *testing.T) {
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
	_, err := client.EnvironmentAutomations.TaskExecutions.NewList(context.TODO(), gitpod.EnvironmentAutomationTaskExecutionNewListParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskExecutionNewListParamsConnectProtocolVersion1),
		Filter: gitpod.F(gitpod.EnvironmentAutomationTaskExecutionNewListParamsFilter{
			EnvironmentIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			Phases:         gitpod.F([]gitpod.EnvironmentAutomationTaskExecutionNewListParamsFilterPhase{gitpod.EnvironmentAutomationTaskExecutionNewListParamsFilterPhaseTaskExecutionPhaseUnspecified}),
			TaskIDs:        gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			TaskReferences: gitpod.F([]string{"string"}),
		}),
		Pagination: gitpod.F(gitpod.EnvironmentAutomationTaskExecutionNewListParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(0)),
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

func TestEnvironmentAutomationTaskExecutionNewGetWithOptionalParams(t *testing.T) {
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
	_, err := client.EnvironmentAutomations.TaskExecutions.NewGet(context.TODO(), gitpod.EnvironmentAutomationTaskExecutionNewGetParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskExecutionNewGetParamsConnectProtocolVersion1),
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

func TestEnvironmentAutomationTaskExecutionStopWithOptionalParams(t *testing.T) {
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
	_, err := client.EnvironmentAutomations.TaskExecutions.Stop(context.TODO(), gitpod.EnvironmentAutomationTaskExecutionStopParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskExecutionStopParamsConnectProtocolVersion1),
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

func TestEnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.EnvironmentAutomations.TaskExecutions.UpdateTaskExecutionStatus(context.TODO(), gitpod.EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParams{
		Body: gitpod.EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsBody{
			ID:             "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			FailureMessage: "failureMessage",
			LogURL:         "logUrl",
			Steps: []gitpod.EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsBodyStep{{
				ID:             gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				FailureMessage: gitpod.F("failureMessage"),
				Phase:          gitpod.F(gitpod.EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsBodyStepsPhaseTaskExecutionPhaseUnspecified),
			}},
		},
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentAutomationTaskExecutionUpdateTaskExecutionStatusParamsConnectProtocolVersion1),
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
