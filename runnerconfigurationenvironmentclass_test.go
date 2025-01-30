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

func TestRunnerConfigurationEnvironmentClassNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.EnvironmentClasses.New(context.TODO(), gitpod.RunnerConfigurationEnvironmentClassNewParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationEnvironmentClassNewParamsConnectProtocolVersion1),
		Configuration: gitpod.F([]gitpod.RunnerConfigurationEnvironmentClassNewParamsConfiguration{{
			Key:   gitpod.F("key"),
			Value: gitpod.F("value"),
		}}),
		Description:      gitpod.F("xxx"),
		DisplayName:      gitpod.F("xxx"),
		RunnerID:         gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRunnerConfigurationEnvironmentClassGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.EnvironmentClasses.Get(context.TODO(), gitpod.RunnerConfigurationEnvironmentClassGetParams{
		Encoding:               gitpod.F(gitpod.RunnerConfigurationEnvironmentClassGetParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationEnvironmentClassGetParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.RunnerConfigurationEnvironmentClassGetParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.RunnerConfigurationEnvironmentClassGetParamsConnectV1),
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

func TestRunnerConfigurationEnvironmentClassUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.EnvironmentClasses.Update(context.TODO(), gitpod.RunnerConfigurationEnvironmentClassUpdateParams{
		Body: gitpod.RunnerConfigurationEnvironmentClassUpdateParamsBodyDescription{
			Description: gitpod.F("xxx"),
		},
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationEnvironmentClassUpdateParamsConnectProtocolVersion1),
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

func TestRunnerConfigurationEnvironmentClassListWithOptionalParams(t *testing.T) {
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
	_, err := client.Runners.Configurations.EnvironmentClasses.List(context.TODO(), gitpod.RunnerConfigurationEnvironmentClassListParams{
		Encoding:               gitpod.F(gitpod.RunnerConfigurationEnvironmentClassListParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerConfigurationEnvironmentClassListParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.RunnerConfigurationEnvironmentClassListParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.RunnerConfigurationEnvironmentClassListParamsConnectV1),
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
