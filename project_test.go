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
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Projects.New(context.TODO(), gitpod.ProjectNewParams{
		EnvironmentClass: gitpod.F[gitpod.ProjectNewParamsEnvironmentClassUnion](gitpod.ProjectNewParamsEnvironmentClassObject{
			EnvironmentClassID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LocalRunner:        gitpod.F(true),
		}),
		Initializer: gitpod.F(gitpod.ProjectNewParamsInitializer{
			Specs: gitpod.F([]gitpod.ProjectNewParamsInitializerSpecUnion{gitpod.ProjectNewParamsInitializerSpecsObject{
				ContextURL: gitpod.F(gitpod.ProjectNewParamsInitializerSpecsObjectContextURL{
					URL: gitpod.F("https://example.com"),
				}),
				Git: gitpod.F(gitpod.ProjectNewParamsInitializerSpecsObjectGit{
					CheckoutLocation:  gitpod.F("checkoutLocation"),
					CloneTarget:       gitpod.F("cloneTarget"),
					RemoteUri:         gitpod.F("remoteUri"),
					TargetMode:        gitpod.F(gitpod.ProjectNewParamsInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified),
					UpstreamRemoteUri: gitpod.F("upstreamRemoteUri"),
				}),
			}}),
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
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Projects.Get(context.TODO(), gitpod.ProjectGetParams{
		Encoding:               gitpod.F(gitpod.ProjectGetParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.ProjectGetParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.ProjectGetParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.ProjectGetParamsConnectV1),
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

func TestProjectUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Update(context.TODO(), gitpod.ProjectUpdateParams{
		Body:                   gitpod.ProjectUpdateParamsBody{},
		ConnectProtocolVersion: gitpod.F(gitpod.ProjectUpdateParamsConnectProtocolVersion1),
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

func TestProjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.List(context.TODO(), gitpod.ProjectListParams{
		Encoding:               gitpod.F(gitpod.ProjectListParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.ProjectListParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.ProjectListParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.ProjectListParamsConnectV1),
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

func TestProjectDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Delete(context.TODO(), gitpod.ProjectDeleteParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ProjectDeleteParamsConnectProtocolVersion1),
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
		option.WithBearerToken("My Bearer Token"),
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
