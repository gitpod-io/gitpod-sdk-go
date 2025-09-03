// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/gitpod-io/gitpod-sdk-go"
	"github.com/gitpod-io/gitpod-sdk-go/internal/testutil"
	"github.com/gitpod-io/gitpod-sdk-go/option"
)

func TestProjectNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
		EnvironmentClass: gitpod.F(gitpod.ProjectEnvironmentClassParam{
			EnvironmentClassID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
			LocalRunner:        gitpod.F(true),
		}),
		Initializer: gitpod.F(gitpod.EnvironmentInitializerParam{
			Specs: gitpod.F([]gitpod.EnvironmentInitializerSpecParam{{
				ContextURL: gitpod.F(gitpod.EnvironmentInitializerSpecsContextURLParam{
					URL: gitpod.F("https://example.com"),
				}),
				Git: gitpod.F(gitpod.EnvironmentInitializerSpecsGitParam{
					CheckoutLocation:  gitpod.F("checkoutLocation"),
					CloneTarget:       gitpod.F("cloneTarget"),
					RemoteUri:         gitpod.F("https://github.com/org/repo"),
					TargetMode:        gitpod.F(gitpod.EnvironmentInitializerSpecsGitTargetModeCloneTargetModeUnspecified),
					UpstreamRemoteUri: gitpod.F("upstreamRemoteUri"),
				}),
			}}),
		}),
		AutomationsFilePath:  gitpod.F("automationsFilePath"),
		DevcontainerFilePath: gitpod.F("devcontainerFilePath"),
		Name:                 gitpod.F("Web Application"),
		TechnicalDescription: gitpod.F("technicalDescription"),
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
	t.Skip("Prism tests are disabled")
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
		ProjectID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
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
	t.Skip("Prism tests are disabled")
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
		AutomationsFilePath:  gitpod.F("automationsFilePath"),
		DevcontainerFilePath: gitpod.F("devcontainerFilePath"),
		EnvironmentClass: gitpod.F(gitpod.ProjectEnvironmentClassParam{
			EnvironmentClassID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
			LocalRunner:        gitpod.F(true),
		}),
		Initializer: gitpod.F(gitpod.EnvironmentInitializerParam{
			Specs: gitpod.F([]gitpod.EnvironmentInitializerSpecParam{{
				ContextURL: gitpod.F(gitpod.EnvironmentInitializerSpecsContextURLParam{
					URL: gitpod.F("https://example.com"),
				}),
				Git: gitpod.F(gitpod.EnvironmentInitializerSpecsGitParam{
					CheckoutLocation:  gitpod.F("checkoutLocation"),
					CloneTarget:       gitpod.F("cloneTarget"),
					RemoteUri:         gitpod.F("remoteUri"),
					TargetMode:        gitpod.F(gitpod.EnvironmentInitializerSpecsGitTargetModeCloneTargetModeUnspecified),
					UpstreamRemoteUri: gitpod.F("upstreamRemoteUri"),
				}),
			}}),
		}),
		Name:                 gitpod.F("x"),
		ProjectID:            gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
		TechnicalDescription: gitpod.F("technicalDescription"),
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
	t.Skip("Prism tests are disabled")
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
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.ProjectListParamsFilter{
			ProjectIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.ProjectListParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(20)),
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

func TestProjectDeleteWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
		ProjectID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
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
	t.Skip("Prism tests are disabled")
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
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
		Name:          gitpod.F("Frontend Project"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
