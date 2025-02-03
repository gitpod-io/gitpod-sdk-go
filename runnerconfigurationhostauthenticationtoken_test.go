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
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenNewParamsConnectProtocolVersion1),
		Token:                  gitpod.F("x"),
		ExpiresAt:              gitpod.F(time.Now()),
		Host:                   gitpod.F("x"),
		RefreshToken:           gitpod.F("refreshToken"),
		RunnerID:               gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Source:                 gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenNewParamsSourceHostAuthenticationTokenSourceUnspecified),
		UserID:                 gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
		Encoding:               gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenGetParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenGetParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenGetParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenGetParamsConnectV1),
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

func TestRunnerConfigurationHostAuthenticationTokenUpdateWithOptionalParams(t *testing.T) {
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
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenUpdateParamsConnectProtocolVersion1),
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
		Encoding:               gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenListParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenListParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenListParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenListParamsConnectV1),
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
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationHostAuthenticationTokenDeleteParamsConnectProtocolVersion1),
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
