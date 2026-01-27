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

func TestGroupRoleAssignmentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Groups.RoleAssignments.New(context.TODO(), gitpod.GroupRoleAssignmentNewParams{
		GroupID:      gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
		ResourceID:   gitpod.F("f53d2330-3795-4c5d-a1f3-453121af9c60"),
		ResourceRole: gitpod.F(shared.ResourceRoleRunnerAdmin),
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

func TestGroupRoleAssignmentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Groups.RoleAssignments.List(context.TODO(), gitpod.GroupRoleAssignmentListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.GroupRoleAssignmentListParamsFilter{
			GroupID:       gitpod.F("groupId"),
			ResourceID:    gitpod.F("resourceId"),
			ResourceRoles: gitpod.F([]shared.ResourceRole{shared.ResourceRoleUnspecified}),
			ResourceTypes: gitpod.F([]shared.ResourceType{shared.ResourceTypeRunner}),
			UserID:        gitpod.F("userId"),
		}),
		Pagination: gitpod.F(gitpod.GroupRoleAssignmentListParamsPagination{
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

func TestGroupRoleAssignmentDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Groups.RoleAssignments.Delete(context.TODO(), gitpod.GroupRoleAssignmentDeleteParams{
		AssignmentID: gitpod.F("a1b2c3d4-5678-90ab-cdef-1234567890ab"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
