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

func TestProjectPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Policies.New(context.TODO(), gitpod.ProjectPolicyNewParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ProjectPolicyNewParamsConnectProtocolVersion1),
		GroupID:                gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:              gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Role:                   gitpod.F(gitpod.ProjectPolicyNewParamsRoleProjectRoleUnspecified),
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

func TestProjectPolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Policies.Update(context.TODO(), gitpod.ProjectPolicyUpdateParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ProjectPolicyUpdateParamsConnectProtocolVersion1),
		GroupID:                gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:              gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Role:                   gitpod.F(gitpod.ProjectPolicyUpdateParamsRoleProjectRoleUnspecified),
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

func TestProjectPolicyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Policies.List(context.TODO(), gitpod.ProjectPolicyListParams{
		Encoding:               gitpod.F(gitpod.ProjectPolicyListParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.ProjectPolicyListParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.ProjectPolicyListParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.ProjectPolicyListParamsConnectV1),
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

func TestProjectPolicyDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Policies.Delete(context.TODO(), gitpod.ProjectPolicyDeleteParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ProjectPolicyDeleteParamsConnectProtocolVersion1),
		GroupID:                gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:              gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
