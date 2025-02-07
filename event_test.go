// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/gitpod-io/flex-sdk-go"
	"github.com/gitpod-io/flex-sdk-go/internal/testutil"
	"github.com/gitpod-io/flex-sdk-go/option"
	"github.com/gitpod-io/flex-sdk-go/shared"
)

func TestEventListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Events.List(context.TODO(), gitpod.EventListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.EventListParamsFilter{
			ActorIDs:        gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ActorPrincipals: gitpod.F([]shared.Principal{shared.PrincipalPrincipalUnspecified}),
			SubjectIDs:      gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			SubjectTypes:    gitpod.F([]gitpod.ResourceType{gitpod.ResourceTypeResourceTypeUnspecified}),
		}),
		Pagination: gitpod.F(gitpod.EventListParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(100)),
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
