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

func TestOrganizationSSOConfigurationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.SSOConfigurations.New(context.TODO(), gitpod.OrganizationSSOConfigurationNewParams{
		ConnectProtocolVersion: gitpod.F(gitpod.OrganizationSSOConfigurationNewParamsConnectProtocolVersion1),
		ClientID:               gitpod.F("x"),
		ClientSecret:           gitpod.F("x"),
		EmailDomain:            gitpod.F("xxxx"),
		IssuerURL:              gitpod.F("https://example.com"),
		OrganizationID:         gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestOrganizationSSOConfigurationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.SSOConfigurations.Get(context.TODO(), gitpod.OrganizationSSOConfigurationGetParams{
		Encoding:               gitpod.F(gitpod.OrganizationSSOConfigurationGetParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.OrganizationSSOConfigurationGetParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.OrganizationSSOConfigurationGetParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.OrganizationSSOConfigurationGetParamsConnectV1),
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

func TestOrganizationSSOConfigurationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.SSOConfigurations.Update(context.TODO(), gitpod.OrganizationSSOConfigurationUpdateParams{
		Body: gitpod.OrganizationSSOConfigurationUpdateParamsBodyClientIDIsTheClientIDOfTheSSOProvider{
			ClientID: gitpod.F("x"),
		},
		ConnectProtocolVersion: gitpod.F(gitpod.OrganizationSSOConfigurationUpdateParamsConnectProtocolVersion1),
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

func TestOrganizationSSOConfigurationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.SSOConfigurations.List(context.TODO(), gitpod.OrganizationSSOConfigurationListParams{
		Encoding:               gitpod.F(gitpod.OrganizationSSOConfigurationListParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.OrganizationSSOConfigurationListParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.OrganizationSSOConfigurationListParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.OrganizationSSOConfigurationListParamsConnectV1),
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

func TestOrganizationSSOConfigurationDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.SSOConfigurations.Delete(context.TODO(), gitpod.OrganizationSSOConfigurationDeleteParams{
		ConnectProtocolVersion: gitpod.F(gitpod.OrganizationSSOConfigurationDeleteParamsConnectProtocolVersion1),
		SSOConfigurationID:     gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
