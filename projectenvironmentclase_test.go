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
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

func TestProjectEnvironmentClaseUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Projects.EnvironmentClases.Update(context.TODO(), gitpod.ProjectEnvironmentClaseUpdateParams{
		ProjectEnvironmentClasses: gitpod.F([]shared.ProjectEnvironmentClassParam{{
			EnvironmentClassID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da041"),
			LocalRunner:        gitpod.F(true),
			Order:              gitpod.F(int64(0)),
		}, {
			EnvironmentClassID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LocalRunner:        gitpod.F(true),
			Order:              gitpod.F(int64(1)),
		}}),
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

func TestProjectEnvironmentClaseListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Projects.EnvironmentClases.List(context.TODO(), gitpod.ProjectEnvironmentClaseListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Pagination: gitpod.F(gitpod.ProjectEnvironmentClaseListParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(20)),
		}),
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
