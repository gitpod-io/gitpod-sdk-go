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

func TestOrganizationPolicyGet(t *testing.T) {
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
	_, err := client.Organizations.Policies.Get(context.TODO(), gitpod.OrganizationPolicyGetParams{
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

func TestOrganizationPolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.Policies.Update(context.TODO(), gitpod.OrganizationPolicyUpdateParams{
		OrganizationID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
		AgentPolicy: gitpod.F(gitpod.OrganizationPolicyUpdateParamsAgentPolicy{
			CommandDenyList:        gitpod.F([]string{"string"}),
			McpDisabled:            gitpod.F(true),
			ScmToolsAllowedGroupID: gitpod.F("scmToolsAllowedGroupId"),
			ScmToolsDisabled:       gitpod.F(true),
		}),
		AllowedEditorIDs:                gitpod.F([]string{"string"}),
		AllowLocalRunners:               gitpod.F(true),
		DefaultEditorID:                 gitpod.F("defaultEditorId"),
		DefaultEnvironmentImage:         gitpod.F("defaultEnvironmentImage"),
		DeleteArchivedEnvironmentsAfter: gitpod.F("+9125115.360s"),
		EditorVersionRestrictions: gitpod.F(map[string]gitpod.OrganizationPolicyUpdateParamsEditorVersionRestrictions{
			"foo": {
				AllowedVersions: gitpod.F([]string{"string"}),
			},
		}),
		MaximumEnvironmentLifetime:        gitpod.F("+9125115.360s"),
		MaximumEnvironmentsPerUser:        gitpod.F("20"),
		MaximumEnvironmentTimeout:         gitpod.F("3600s"),
		MaximumRunningEnvironmentsPerUser: gitpod.F("5"),
		MembersCreateProjects:             gitpod.F(true),
		MembersRequireProjects:            gitpod.F(true),
		PortSharingDisabled:               gitpod.F(true),
		RequireCustomDomainAccess:         gitpod.F(true),
		RestrictAccountCreationToScim:     gitpod.F(true),
		SecurityAgentPolicy: gitpod.F(gitpod.OrganizationPolicyUpdateParamsSecurityAgentPolicy{
			Crowdstrike: gitpod.F(gitpod.OrganizationPolicyUpdateParamsSecurityAgentPolicyCrowdstrike{
				AdditionalOptions: gitpod.F(map[string]string{
					"foo": "string",
				}),
				CidSecretID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Enabled:     gitpod.F(true),
				Image:       gitpod.F("image"),
				Tags:        gitpod.F("tags"),
			}),
		}),
		VetoExecPolicy: gitpod.F(gitpod.VetoExecPolicyParam{
			Action:      gitpod.F(gitpod.KernelControlsActionUnspecified),
			Enabled:     gitpod.F(true),
			Executables: gitpod.F([]string{"string"}),
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
