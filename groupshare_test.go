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

func TestGroupShareNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Groups.Shares.New(context.TODO(), gitpod.GroupShareNewParams{
		Principal:    gitpod.F(shared.PrincipalServiceAccount),
		PrincipalID:  gitpod.F("a1b2c3d4-5678-90ab-cdef-1234567890ab"),
		ResourceID:   gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
		ResourceType: gitpod.F(shared.ResourceTypeRunner),
		Role:         gitpod.F(shared.ResourceRoleRunnerUser),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGroupShareDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Groups.Shares.Delete(context.TODO(), gitpod.GroupShareDeleteParams{
		Principal:    gitpod.F(shared.PrincipalUser),
		PrincipalID:  gitpod.F("f53d2330-3795-4c5d-a1f3-453121af9c60"),
		ResourceID:   gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
		ResourceType: gitpod.F(shared.ResourceTypeRunner),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
