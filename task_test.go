// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/gitpod-go"
	"github.com/stainless-sdks/gitpod-go/internal/testutil"
	"github.com/stainless-sdks/gitpod-go/option"
)

func TestTaskNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Tasks.New(context.TODO(), gitpod.TaskNewParams{
		ConnectProtocolVersion: gitpod.F(gitpod.TaskNewParamsConnectProtocolVersion1),
		DependsOn:              gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		EnvironmentID:          gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: gitpod.F(gitpod.TaskNewParamsMetadata{
			CreatedAt: gitpod.F(time.Now()),
			Creator: gitpod.F(gitpod.TaskNewParamsMetadataCreator{
				ID:        gitpod.F("id"),
				Principal: gitpod.F(gitpod.TaskNewParamsMetadataCreatorPrincipalPrincipalUnspecified),
			}),
			Description: gitpod.F("description"),
			Name:        gitpod.F("x"),
			Reference:   gitpod.F("reference"),
			TriggeredBy: gitpod.F([]gitpod.TaskNewParamsMetadataTriggeredByUnion{gitpod.TaskNewParamsMetadataTriggeredByUnknown(map[string]interface{}{}), gitpod.TaskNewParamsMetadataTriggeredByUnknown(map[string]interface{}{}), gitpod.TaskNewParamsMetadataTriggeredByUnknown(map[string]interface{}{})}),
		}),
		Spec: gitpod.F(gitpod.TaskNewParamsSpec{
			Command: gitpod.F("command"),
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

func TestTaskGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Tasks.Get(context.TODO(), gitpod.TaskGetParams{
		ConnectProtocolVersion: gitpod.F(gitpod.TaskGetParamsConnectProtocolVersion1),
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

func TestTaskGetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Tasks.GetNew(context.TODO(), gitpod.TaskGetNewParams{
		ConnectProtocolVersion: gitpod.F(gitpod.TaskGetNewParamsConnectProtocolVersion1),
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
