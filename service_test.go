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

func TestServiceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Services.Update(context.TODO(), gitpod.ServiceUpdateParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ServiceUpdateParamsConnectProtocolVersion1),
		ID:                     gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata:               gitpod.F(gitpod.ServiceUpdateParamsMetadata{}),
		Spec:                   gitpod.F[gitpod.ServiceUpdateParamsSpecUnion](gitpod.ServiceUpdateParamsSpecUnknown(map[string]interface{}{})),
		Status:                 gitpod.F(gitpod.ServiceUpdateParamsStatus{}),
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

func TestServiceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Services.List(context.TODO(), gitpod.ServiceListParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ServiceListParamsConnectProtocolVersion1),
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

func TestServiceDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Services.Delete(context.TODO(), gitpod.ServiceDeleteParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ServiceDeleteParamsConnectProtocolVersion1),
		ID:                     gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Force:                  gitpod.F(true),
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

func TestServiceListNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Services.ListNew(context.TODO(), gitpod.ServiceListNewParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ServiceListNewParamsConnectProtocolVersion1),
		Filter: gitpod.F(gitpod.ServiceListNewParamsFilter{
			EnvironmentIDs: gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			References:     gitpod.F([]string{"x"}),
			ServiceIDs:     gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
		Pagination: gitpod.F(gitpod.ServiceListNewParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(0)),
		}),
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

func TestServiceStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Services.Start(context.TODO(), gitpod.ServiceStartParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ServiceStartParamsConnectProtocolVersion1),
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

func TestServiceStopWithOptionalParams(t *testing.T) {
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
	_, err := client.Services.Stop(context.TODO(), gitpod.ServiceStopParams{
		ConnectProtocolVersion: gitpod.F(gitpod.ServiceStopParamsConnectProtocolVersion1),
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
