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
		Spec: gitpod.F(gitpod.EnvironmentSpecParam{
			Admission: gitpod.F(gitpod.AdmissionLevelUnspecified),
			AutomationsFile: gitpod.F(gitpod.EnvironmentSpecAutomationsFileParam{
				AutomationsFilePath: gitpod.F("automationsFilePath"),
				Session:             gitpod.F("session"),
			}),
			Content: gitpod.F(gitpod.EnvironmentSpecContentParam{
				GitEmail:    gitpod.F("gitEmail"),
				GitUsername: gitpod.F("gitUsername"),
				Initializer: gitpod.F(gitpod.EnvironmentInitializerParam{
					Specs: gitpod.F([]gitpod.EnvironmentInitializerSpecParam{{
						ContextURL: gitpod.F(gitpod.EnvironmentInitializerSpecsContextURLParam{
							URL: gitpod.F("https://github.com/gitpod-io/gitpod"),
						}),
						Git: gitpod.F(gitpod.EnvironmentInitializerSpecsGitParam{
							CheckoutLocation:  gitpod.F("checkoutLocation"),
							CloneTarget:       gitpod.F("cloneTarget"),
							RemoteUri:         gitpod.F("remoteUri"),
							TargetMode:        gitpod.F(gitpod.EnvironmentInitializerSpecsGitTargetModeCloneTargetModeUnspecified),
							UpstreamRemoteUri: gitpod.F("upstreamRemoteUri"),
						}),
					}}),
				}),
				Session: gitpod.F("session"),
			}),
			DesiredPhase: gitpod.F(gitpod.EnvironmentPhaseUnspecified),
			Devcontainer: gitpod.F(gitpod.EnvironmentSpecDevcontainerParam{
				DevcontainerFilePath: gitpod.F("devcontainerFilePath"),
				Session:              gitpod.F("session"),
			}),
			Machine: gitpod.F(gitpod.EnvironmentSpecMachineParam{
				Class:   gitpod.F("61000000-0000-0000-0000-000000000000"),
				Session: gitpod.F("session"),
			}),
			Ports: gitpod.F([]gitpod.EnvironmentSpecPortParam{{
				Admission: gitpod.F(gitpod.AdmissionLevelUnspecified),
				Name:      gitpod.F("x"),
				Port:      gitpod.F(int64(1)),
			}}),
			Secrets: gitpod.F([]gitpod.EnvironmentSpecSecretParam{{
				EnvironmentVariable: gitpod.F("environmentVariable"),
				FilePath:            gitpod.F("filePath"),
				GitCredentialHost:   gitpod.F("gitCredentialHost"),
				Name:                gitpod.F("name"),
				Session:             gitpod.F("session"),
				Source:              gitpod.F("source"),
				SourceRef:           gitpod.F("sourceRef"),
			}}),
			SpecVersion: gitpod.F("specVersion"),
			SSHPublicKeys: gitpod.F([]gitpod.EnvironmentSpecSSHPublicKeyParam{{
				ID:    gitpod.F("id"),
				Value: gitpod.F("value"),
			}}),
			Timeout: gitpod.F(gitpod.EnvironmentSpecTimeoutParam{
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
		Spec: gitpod.F(gitpod.EnvironmentUpdateParamsSpec{
			AutomationsFile: gitpod.F(gitpod.EnvironmentUpdateParamsSpecAutomationsFile{
				AutomationsFilePath: gitpod.F("automationsFilePath"),
				Session:             gitpod.F("session"),
			}),
			Content: gitpod.F(gitpod.EnvironmentUpdateParamsSpecContent{
				GitEmail:    gitpod.F("gitEmail"),
				GitUsername: gitpod.F("gitUsername"),
				Initializer: gitpod.F(gitpod.EnvironmentInitializerParam{
					Specs: gitpod.F([]gitpod.EnvironmentInitializerSpecParam{{
						ContextURL: gitpod.F(gitpod.EnvironmentInitializerSpecsContextURLParam{
							URL: gitpod.F("https://example.com"),
						}),
						Git: gitpod.F(gitpod.EnvironmentInitializerSpecsGitParam{
							CheckoutLocation:  gitpod.F("checkoutLocation"),
							CloneTarget:       gitpod.F("cloneTarget"),
							RemoteUri:         gitpod.F("remoteUri"),
							TargetMode:        gitpod.F(gitpod.EnvironmentInitializerSpecsGitTargetModeCloneTargetModeUnspecified),
							UpstreamRemoteUri: gitpod.F("upstreamRemoteUri"),
						}),
					}}),
				}),
				Session: gitpod.F("session"),
			}),
			Devcontainer: gitpod.F(gitpod.EnvironmentUpdateParamsSpecDevcontainer{
				DevcontainerFilePath: gitpod.F("devcontainerFilePath"),
				Session:              gitpod.F("session"),
			}),
			Ports: gitpod.F([]gitpod.EnvironmentUpdateParamsSpecPort{{
				Admission: gitpod.F(gitpod.AdmissionLevelUnspecified),
				Name:      gitpod.F("x"),
				Port:      gitpod.F(int64(1)),
			}}),
			SSHPublicKeys: gitpod.F([]gitpod.EnvironmentUpdateParamsSpecSSHPublicKey{{
				ID:    gitpod.F("id"),
				Value: gitpod.F("value"),
			}}),
			Timeout: gitpod.F(gitpod.EnvironmentUpdateParamsSpecTimeout{
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
			RunnerKinds:  gitpod.F([]gitpod.RunnerKind{gitpod.RunnerKindUnspecified}),
			StatusPhases: gitpod.F([]gitpod.EnvironmentPhase{gitpod.EnvironmentPhaseUnspecified}),
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
		Spec: gitpod.F(gitpod.EnvironmentSpecParam{
			Admission: gitpod.F(gitpod.AdmissionLevelUnspecified),
			AutomationsFile: gitpod.F(gitpod.EnvironmentSpecAutomationsFileParam{
				AutomationsFilePath: gitpod.F("automationsFilePath"),
				Session:             gitpod.F("session"),
			}),
			Content: gitpod.F(gitpod.EnvironmentSpecContentParam{
				GitEmail:    gitpod.F("gitEmail"),
				GitUsername: gitpod.F("gitUsername"),
				Initializer: gitpod.F(gitpod.EnvironmentInitializerParam{
					Specs: gitpod.F([]gitpod.EnvironmentInitializerSpecParam{{
						ContextURL: gitpod.F(gitpod.EnvironmentInitializerSpecsContextURLParam{
							URL: gitpod.F("https://example.com"),
						}),
						Git: gitpod.F(gitpod.EnvironmentInitializerSpecsGitParam{
							CheckoutLocation:  gitpod.F("checkoutLocation"),
							CloneTarget:       gitpod.F("cloneTarget"),
							RemoteUri:         gitpod.F("remoteUri"),
							TargetMode:        gitpod.F(gitpod.EnvironmentInitializerSpecsGitTargetModeCloneTargetModeUnspecified),
							UpstreamRemoteUri: gitpod.F("upstreamRemoteUri"),
						}),
					}}),
				}),
				Session: gitpod.F("session"),
			}),
			DesiredPhase: gitpod.F(gitpod.EnvironmentPhaseUnspecified),
			Devcontainer: gitpod.F(gitpod.EnvironmentSpecDevcontainerParam{
				DevcontainerFilePath: gitpod.F("devcontainerFilePath"),
				Session:              gitpod.F("session"),
			}),
			Machine: gitpod.F(gitpod.EnvironmentSpecMachineParam{
				Class:   gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Session: gitpod.F("session"),
			}),
			Ports: gitpod.F([]gitpod.EnvironmentSpecPortParam{{
				Admission: gitpod.F(gitpod.AdmissionLevelUnspecified),
				Name:      gitpod.F("x"),
				Port:      gitpod.F(int64(1)),
			}}),
			Secrets: gitpod.F([]gitpod.EnvironmentSpecSecretParam{{
				EnvironmentVariable: gitpod.F("environmentVariable"),
				FilePath:            gitpod.F("filePath"),
				GitCredentialHost:   gitpod.F("gitCredentialHost"),
				Name:                gitpod.F("name"),
				Session:             gitpod.F("session"),
				Source:              gitpod.F("source"),
				SourceRef:           gitpod.F("sourceRef"),
			}}),
			SpecVersion: gitpod.F("specVersion"),
			SSHPublicKeys: gitpod.F([]gitpod.EnvironmentSpecSSHPublicKeyParam{{
				ID:    gitpod.F("id"),
				Value: gitpod.F("value"),
			}}),
			Timeout: gitpod.F(gitpod.EnvironmentSpecTimeoutParam{
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
		ActivitySignal: gitpod.F(gitpod.EnvironmentActivitySignalParam{
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
