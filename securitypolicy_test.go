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

func TestSecurityPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.SecurityPolicies.New(context.TODO(), gitpod.SecurityPolicyNewParams{
		Metadata: gitpod.F(gitpod.SecurityPolicyNewParamsMetadata{
			Name: gitpod.F("Veto Exec audit-first"),
		}),
		Spec: gitpod.F(gitpod.SecurityPolicyNewParamsSpec{
			Executables: gitpod.F(gitpod.SecurityPolicyNewParamsSpecExecutables{
				DefaultEffect: gitpod.F(gitpod.SecurityPolicyNewParamsSpecExecutablesDefaultEffectEffectAllow),
				Rules: gitpod.F([]gitpod.SecurityPolicyNewParamsSpecExecutablesRule{{
					Effect: gitpod.F(gitpod.SecurityPolicyNewParamsSpecExecutablesRulesEffectEffectAudit),
					Path:   gitpod.F("npx"),
				}, {
					Effect: gitpod.F(gitpod.SecurityPolicyNewParamsSpecExecutablesRulesEffectEffectBlock),
					Path:   gitpod.F("/usr/bin/curl"),
				}}),
			}),
		}),
		OrganizationID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecurityPolicyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SecurityPolicies.Get(context.TODO(), gitpod.SecurityPolicyGetParams{
		SecurityPolicyID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecurityPolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.SecurityPolicies.Update(context.TODO(), gitpod.SecurityPolicyUpdateParams{
		Metadata: gitpod.F(gitpod.SecurityPolicyUpdateParamsMetadata{
			Name: gitpod.F("x"),
		}),
		SecurityPolicyID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
		Spec: gitpod.F(gitpod.SecurityPolicyUpdateParamsSpec{
			Executables: gitpod.F(gitpod.SecurityPolicyUpdateParamsSpecExecutables{
				DefaultEffect: gitpod.F(gitpod.SecurityPolicyUpdateParamsSpecExecutablesDefaultEffectEffectAllow),
				Rules: gitpod.F([]gitpod.SecurityPolicyUpdateParamsSpecExecutablesRule{{
					Effect: gitpod.F(gitpod.SecurityPolicyUpdateParamsSpecExecutablesRulesEffectEffectBlock),
					Path:   gitpod.F("npx"),
				}, {
					Effect: gitpod.F(gitpod.SecurityPolicyUpdateParamsSpecExecutablesRulesEffectEffectBlock),
					Path:   gitpod.F("/usr/bin/curl"),
				}}),
			}),
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

func TestSecurityPolicyListWithOptionalParams(t *testing.T) {
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
	_, err := client.SecurityPolicies.List(context.TODO(), gitpod.SecurityPolicyListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.SecurityPolicyListParamsFilter{
			OrganizationID:    gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
			Search:            gitpod.F("search"),
			SecurityPolicyIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.SecurityPolicyListParamsPagination{
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

func TestSecurityPolicyDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.SecurityPolicies.Delete(context.TODO(), gitpod.SecurityPolicyDeleteParams{
		SecurityPolicyID: gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
