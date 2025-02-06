// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/gitpod-io/flex-sdk-go"
	"github.com/gitpod-io/flex-sdk-go/internal/testutil"
	"github.com/gitpod-io/flex-sdk-go/option"
)

func TestEnvironmentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.New(context.TODO(), gitpod.EnvironmentNewParams{
		Spec: gitpod.F(gitpod.EnvironmentNewParamsSpec{
			Admission: gitpod.F(gitpod.EnvironmentNewParamsSpecAdmissionAdmissionLevelUnspecified),
			AutomationsFile: gitpod.F(gitpod.EnvironmentNewParamsSpecAutomationsFile{
				AutomationsFilePath: gitpod.F("automationsFilePath"),
				Session:             gitpod.F("session"),
			}),
			Content: gitpod.F(gitpod.EnvironmentNewParamsSpecContent{
				GitEmail:    gitpod.F("gitEmail"),
				GitUsername: gitpod.F("gitUsername"),
				Initializer: gitpod.F(gitpod.EnvironmentNewParamsSpecContentInitializer{
					Specs: gitpod.F([]gitpod.EnvironmentNewParamsSpecContentInitializerSpecUnion{gitpod.EnvironmentNewParamsSpecContentInitializerSpecsContextURL{
						ContextURL: gitpod.F(gitpod.EnvironmentNewParamsSpecContentInitializerSpecsContextURLContextURL{
							URL: gitpod.F("https://example.com"),
						}),
					}}),
				}),
				Session: gitpod.F("session"),
			}),
			DesiredPhase: gitpod.F(gitpod.EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseUnspecified),
			Devcontainer: gitpod.F(gitpod.EnvironmentNewParamsSpecDevcontainer{
				DevcontainerFilePath: gitpod.F("devcontainerFilePath"),
				Session:              gitpod.F("session"),
			}),
			Machine: gitpod.F(gitpod.EnvironmentNewParamsSpecMachine{
				Class:   gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Session: gitpod.F("session"),
			}),
			Ports: gitpod.F([]gitpod.EnvironmentNewParamsSpecPort{{
				Admission: gitpod.F(gitpod.EnvironmentNewParamsSpecPortsAdmissionAdmissionLevelUnspecified),
				Name:      gitpod.F("x"),
				Port:      gitpod.F(int64(1)),
			}}),
			Secrets: gitpod.F([]gitpod.EnvironmentNewParamsSpecSecretUnion{gitpod.EnvironmentNewParamsSpecSecretsObject{
				EnvironmentVariable: gitpod.F("environmentVariable"),
				Name:                gitpod.F("name"),
				Session:             gitpod.F("session"),
				Source:              gitpod.F("source"),
				SourceRef:           gitpod.F("sourceRef"),
			}}),
			SpecVersion: gitpod.F("specVersion"),
			SSHPublicKeys: gitpod.F([]gitpod.EnvironmentNewParamsSpecSSHPublicKey{{
				ID:    gitpod.F("id"),
				Value: gitpod.F("value"),
			}}),
			Timeout: gitpod.F(gitpod.EnvironmentNewParamsSpecTimeout{
				Disconnected: gitpod.F("+9125115.360s"),
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

func TestEnvironmentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Get(context.TODO(), gitpod.EnvironmentGetParams{
		EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnvironmentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Update(context.TODO(), gitpod.EnvironmentUpdateParams{
		EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata:      gitpod.F[any](map[string]interface{}{}),
		Spec: gitpod.F[gitpod.EnvironmentUpdateParamsSpecUnion](gitpod.EnvironmentUpdateParamsSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironment{
			AutomationsFile: gitpod.F[gitpod.EnvironmentUpdateParamsSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileUnion](gitpod.EnvironmentUpdateParamsSpecAutomationsFileIsTheAutomationsFileSpecOfTheEnvironmentAutomationsFileAutomationsFilePathIsThePathToTheAutomationsFileThatIsAppliedInTheEnvironmentRelativeToTheRepoRoot{
				AutomationsFilePath: gitpod.F("automationsFilePath"),
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

func TestEnvironmentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.List(context.TODO(), gitpod.EnvironmentListParams{
		Token:    gitpod.F("token"),
		PageSize: gitpod.F(int64(0)),
		Filter: gitpod.F(gitpod.EnvironmentListParamsFilter{
			CreatorIDs:   gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ProjectIDs:   gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			RunnerIDs:    gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			RunnerKinds:  gitpod.F([]gitpod.EnvironmentListParamsFilterRunnerKind{gitpod.EnvironmentListParamsFilterRunnerKindRunnerKindUnspecified}),
			StatusPhases: gitpod.F([]gitpod.EnvironmentListParamsFilterStatusPhase{gitpod.EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseUnspecified}),
		}),
		OrganizationID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Pagination: gitpod.F(gitpod.EnvironmentListParamsPagination{
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

func TestEnvironmentDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Delete(context.TODO(), gitpod.EnvironmentDeleteParams{
		EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Force:         gitpod.F(true),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnvironmentNewFromProjectWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.NewFromProject(context.TODO(), gitpod.EnvironmentNewFromProjectParams{
		ProjectID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Spec: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpec{
			Admission: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecAdmissionAdmissionLevelUnspecified),
			AutomationsFile: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecAutomationsFile{
				AutomationsFilePath: gitpod.F("automationsFilePath"),
				Session:             gitpod.F("session"),
			}),
			Content: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecContent{
				GitEmail:    gitpod.F("gitEmail"),
				GitUsername: gitpod.F("gitUsername"),
				Initializer: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecContentInitializer{
					Specs: gitpod.F([]gitpod.EnvironmentNewFromProjectParamsSpecContentInitializerSpecUnion{gitpod.EnvironmentNewFromProjectParamsSpecContentInitializerSpecsContextURL{
						ContextURL: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecContentInitializerSpecsContextURLContextURL{
							URL: gitpod.F("https://example.com"),
						}),
					}}),
				}),
				Session: gitpod.F("session"),
			}),
			DesiredPhase: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecDesiredPhaseEnvironmentPhaseUnspecified),
			Devcontainer: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecDevcontainer{
				DevcontainerFilePath: gitpod.F("devcontainerFilePath"),
				Session:              gitpod.F("session"),
			}),
			Machine: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecMachine{
				Class:   gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Session: gitpod.F("session"),
			}),
			Ports: gitpod.F([]gitpod.EnvironmentNewFromProjectParamsSpecPort{{
				Admission: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecPortsAdmissionAdmissionLevelUnspecified),
				Name:      gitpod.F("x"),
				Port:      gitpod.F(int64(1)),
			}}),
			Secrets: gitpod.F([]gitpod.EnvironmentNewFromProjectParamsSpecSecretUnion{gitpod.EnvironmentNewFromProjectParamsSpecSecretsObject{
				EnvironmentVariable: gitpod.F("environmentVariable"),
				Name:                gitpod.F("name"),
				Session:             gitpod.F("session"),
				Source:              gitpod.F("source"),
				SourceRef:           gitpod.F("sourceRef"),
			}}),
			SpecVersion: gitpod.F("specVersion"),
			SSHPublicKeys: gitpod.F([]gitpod.EnvironmentNewFromProjectParamsSpecSSHPublicKey{{
				ID:    gitpod.F("id"),
				Value: gitpod.F("value"),
			}}),
			Timeout: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecTimeout{
				Disconnected: gitpod.F("+9125115.360s"),
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

func TestEnvironmentNewLogsTokenWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.NewLogsToken(context.TODO(), gitpod.EnvironmentNewLogsTokenParams{
		EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnvironmentMarkActiveWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.MarkActive(context.TODO(), gitpod.EnvironmentMarkActiveParams{
		ActivitySignal: gitpod.F(gitpod.EnvironmentMarkActiveParamsActivitySignal{
			Source:    gitpod.F("xxx"),
			Timestamp: gitpod.F(time.Now()),
		}),
		EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnvironmentStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Start(context.TODO(), gitpod.EnvironmentStartParams{
		EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnvironmentStopWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Stop(context.TODO(), gitpod.EnvironmentStopParams{
		EnvironmentID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
