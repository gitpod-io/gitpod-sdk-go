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

func TestSecretNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Secrets.New(context.TODO(), gitpod.SecretNewParams{
		APIOnly:                        gitpod.F(true),
		ContainerRegistryBasicAuthHost: gitpod.F("containerRegistryBasicAuthHost"),
		EnvironmentVariable:            gitpod.F(true),
		FilePath:                       gitpod.F("filePath"),
		Name:                           gitpod.F("DATABASE_URL"),
		ProjectID:                      gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
		Scope: gitpod.F(gitpod.SecretScopeParam{
			OrganizationID:   gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:        gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ServiceAccountID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			UserID:           gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		Value: gitpod.F("postgresql://user:pass@localhost:5432/db"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecretListWithOptionalParams(t *testing.T) {
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
	_, err := client.Secrets.List(context.TODO(), gitpod.SecretListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.SecretListParamsFilter{
			ProjectIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			Scope: gitpod.F(gitpod.SecretScopeParam{
				OrganizationID:   gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProjectID:        gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
				ServiceAccountID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				UserID:           gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
		}),
		Pagination: gitpod.F(gitpod.SecretListParamsPagination{
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

func TestSecretDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Secrets.Delete(context.TODO(), gitpod.SecretDeleteParams{
		SecretID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecretGetValueWithOptionalParams(t *testing.T) {
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
	_, err := client.Secrets.GetValue(context.TODO(), gitpod.SecretGetValueParams{
		SecretID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecretUpdateValueWithOptionalParams(t *testing.T) {
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
	_, err := client.Secrets.UpdateValue(context.TODO(), gitpod.SecretUpdateValueParams{
		SecretID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
		Value:    gitpod.F("new-secret-value"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
