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

func TestProjectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.New(context.TODO(), gitpod.ProjectNewParams{
		EnvironmentClass: gitpod.F[gitpod.ProjectNewParamsEnvironmentClassUnion](gitpod.ProjectNewParamsEnvironmentClassUnknown(map[string]interface{}{})),
		Initializer: gitpod.F(gitpod.ProjectNewParamsInitializer{
			Specs: gitpod.F([]gitpod.ProjectNewParamsInitializerSpecUnion{gitpod.ProjectNewParamsInitializerSpecsUnknown(map[string]interface{}{}), gitpod.ProjectNewParamsInitializerSpecsUnknown(map[string]interface{}{}), gitpod.ProjectNewParamsInitializerSpecsUnknown(map[string]interface{}{})}),
		}),
		ConnectProtocolVersion: gitpod.F(gitpod.ProjectNewParamsConnectProtocolVersion1),
		AutomationsFilePath:    gitpod.F("automationsFilePath"),
		DevcontainerFilePath:   gitpod.F("devcontainerFilePath"),
		Name:                   gitpod.F("x"),
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

func TestProjectGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Get(context.TODO(), gitpod.ProjectGetParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ProjectGetParamsConnectProtocolVersion1),
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

func TestProjectNewFromEnvironmentWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.NewFromEnvironment(context.TODO(), gitpod.ProjectNewFromEnvironmentParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ProjectNewFromEnvironmentParamsConnectProtocolVersion1),
		EnvironmentID:          gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Name:                   gitpod.F("x"),
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
