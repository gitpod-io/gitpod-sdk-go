// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/gitpod-io/flex-sdk-go"
	"github.com/gitpod-io/flex-sdk-go/internal/testutil"
	"github.com/gitpod-io/flex-sdk-go/option"
)

func TestRunnerConfigurationHostAuthenticationTokenNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.HostAuthenticationTokens.New(context.TODO(), gitpod.RunnerConfigurationHostAuthenticationTokenNewParams{
		Token:        gitpod.F("x"),
		ExpiresAt:    gitpod.F(time.Now()),
		Host:         gitpod.F("x"),
		RefreshToken: gitpod.F("refreshToken"),
		RunnerID:     gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Source:       gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenNewParamsSourceHostAuthenticationTokenSourceUnspecified),
		UserID:       gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRunnerConfigurationHostAuthenticationTokenGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.HostAuthenticationTokens.Get(context.TODO(), gitpod.RunnerConfigurationHostAuthenticationTokenGetParams{
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

func TestRunnerConfigurationHostAuthenticationTokenUpdate(t *testing.T) {
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
	_, err := client.Runners.Configurations.HostAuthenticationTokens.Update(context.TODO(), gitpod.RunnerConfigurationHostAuthenticationTokenUpdateParams{
		Body: gitpod.RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyExpiresAt{
			ExpiresAt: gitpod.F(time.Now()),
		},
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRunnerConfigurationHostAuthenticationTokenListWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.HostAuthenticationTokens.List(context.TODO(), gitpod.RunnerConfigurationHostAuthenticationTokenListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F[gitpod.RunnerConfigurationHostAuthenticationTokenListParamsFilterUnion](gitpod.RunnerConfigurationHostAuthenticationTokenListParamsFilterRunnerID{
			RunnerID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		Pagination: gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenListParamsPagination{
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

func TestRunnerConfigurationHostAuthenticationTokenDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.HostAuthenticationTokens.Delete(context.TODO(), gitpod.RunnerConfigurationHostAuthenticationTokenDeleteParams{
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
