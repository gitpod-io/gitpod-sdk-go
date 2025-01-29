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

func TestRunnerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.New(context.TODO(), gitpod.RunnerNewParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerNewParamsConnectProtocolVersion1),
		Kind:                   gitpod.F(gitpod.RunnerNewParamsKindRunnerKindUnspecified),
		Name:                   gitpod.F("xxx"),
		Spec: gitpod.F(gitpod.RunnerNewParamsSpec{
			Configuration: gitpod.F(gitpod.RunnerNewParamsSpecConfiguration{
				AutoUpdate:     gitpod.F(true),
				Region:         gitpod.F("region"),
				ReleaseChannel: gitpod.F(gitpod.RunnerNewParamsSpecConfigurationReleaseChannelRunnerReleaseChannelUnspecified),
			}),
			DesiredPhase: gitpod.F(gitpod.RunnerNewParamsSpecDesiredPhaseRunnerPhaseUnspecified),
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

func TestRunnerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.List(context.TODO(), gitpod.RunnerListParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerListParamsConnectProtocolVersion1),
		Filter: gitpod.F(gitpod.RunnerListParamsFilter{
			CreatorIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			Kinds:      gitpod.F([]gitpod.RunnerListParamsFilterKind{gitpod.RunnerListParamsFilterKindRunnerKindUnspecified}),
		}),
		Pagination: gitpod.F(gitpod.RunnerListParamsPagination{
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

func TestRunnerCheckAuthenticationForHostWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.CheckAuthenticationForHost(context.TODO(), gitpod.RunnerCheckAuthenticationForHostParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerCheckAuthenticationForHostParamsConnectProtocolVersion1),
		Host:                   gitpod.F("host"),
		RunnerID:               gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRunnerNewRunnerTokenWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.NewRunnerToken(context.TODO(), gitpod.RunnerNewRunnerTokenParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerNewRunnerTokenParamsConnectProtocolVersion1),
		RunnerID:               gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRunnerDeleteRunnerWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.DeleteRunner(context.TODO(), gitpod.RunnerDeleteRunnerParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerDeleteRunnerParamsConnectProtocolVersion1),
		Force:                  gitpod.F(true),
		RunnerID:               gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRunnerGetRunnerWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.GetRunner(context.TODO(), gitpod.RunnerGetRunnerParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerGetRunnerParamsConnectProtocolVersion1),
		RunnerID:               gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRunnerParseContextURLWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.ParseContextURL(context.TODO(), gitpod.RunnerParseContextURLParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerParseContextURLParamsConnectProtocolVersion1),
		ContextURL:             gitpod.F("https://example.com"),
		RunnerID:               gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRunnerUpdateRunnerWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.UpdateRunner(context.TODO(), gitpod.RunnerUpdateRunnerParams{
		Body:                   gitpod.RunnerUpdateRunnerParamsBody{},
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerUpdateRunnerParamsConnectProtocolVersion1),
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
