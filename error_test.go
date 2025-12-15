// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go"
	"github.com/gitpod-io/gitpod-sdk-go/internal/testutil"
	"github.com/gitpod-io/gitpod-sdk-go/option"
)

func TestErrorReportErrorsWithOptionalParams(t *testing.T) {
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
	_, err := client.Errors.ReportErrors(context.TODO(), gitpod.ErrorReportErrorsParams{
		Events: gitpod.F([]gitpod.ErrorEventParam{{
			Breadcrumbs: gitpod.F([]gitpod.BreadcrumbParam{{
				Category: gitpod.F("category"),
				Data: gitpod.F(map[string]string{
					"foo": "string",
				}),
				Level:     gitpod.F(gitpod.ErrorLevelUnspecified),
				Message:   gitpod.F("message"),
				Timestamp: gitpod.F(time.Now()),
				Type:      gitpod.F("type"),
			}}),
			Environment: gitpod.F("environment"),
			EventID:     gitpod.F("210b9798eb53baa4e69d31c1071cf03d"),
			Exceptions: gitpod.F([]gitpod.ExceptionInfoParam{{
				Mechanism: gitpod.F(gitpod.ExceptionMechanismParam{
					Data: gitpod.F(map[string]string{
						"foo": "string",
					}),
					Description: gitpod.F("description"),
					Handled:     gitpod.F(true),
					Synthetic:   gitpod.F(true),
					Type:        gitpod.F("x"),
				}),
				Module: gitpod.F("module"),
				Stacktrace: gitpod.F([]gitpod.StackFrameParam{{
					Colno:       gitpod.F(int64(0)),
					ContextLine: gitpod.F("contextLine"),
					Filename:    gitpod.F("filename"),
					Function:    gitpod.F("function"),
					InApp:       gitpod.F(true),
					Lineno:      gitpod.F(int64(0)),
					Module:      gitpod.F("module"),
					PostContext: gitpod.F([]string{"string"}),
					PreContext:  gitpod.F([]string{"string"}),
					Vars: gitpod.F(map[string]string{
						"foo": "string",
					}),
				}}),
				ThreadID: gitpod.F("threadId"),
				Type:     gitpod.F("x"),
				Value:    gitpod.F("value"),
			}}),
			Extra: gitpod.F(map[string]string{
				"foo": "string",
			}),
			Fingerprint: gitpod.F([]string{"x"}),
			IdentityID:  gitpod.F("ecc2efdd-ddfa-31a9-c6f1-b833d337aa7c"),
			Level:       gitpod.F(gitpod.ErrorLevelUnspecified),
			Logger:      gitpod.F("logger"),
			Modules: gitpod.F(map[string]string{
				"foo": "string",
			}),
			Platform: gitpod.F("x"),
			Release:  gitpod.F("release"),
			Request: gitpod.F(gitpod.RequestInfoParam{
				Data: gitpod.F("data"),
				Headers: gitpod.F(map[string]string{
					"foo": "string",
				}),
				Method: gitpod.F("method"),
				QueryString: gitpod.F(map[string]string{
					"foo": "string",
				}),
				URL: gitpod.F("url"),
			}),
			SDK: gitpod.F(map[string]string{
				"foo": "string",
			}),
			ServerName: gitpod.F("serverName"),
			Tags: gitpod.F(map[string]string{
				"foo": "string",
			}),
			Timestamp:   gitpod.F(time.Now()),
			Transaction: gitpod.F("transaction"),
		}}),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
