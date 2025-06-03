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
				DefaultDevcontainerImage: gitpod.F("defaultDevcontainerImage"),
				DevcontainerFilePath:     gitpod.F("devcontainerFilePath"),
				Dotfiles: gitpod.F(gitpod.EnvironmentSpecDevcontainerDotfilesParam{
					Repository: gitpod.F("https://example.com"),
				}),
				Session: gitpod.F("session"),
			}),
			Machine: gitpod.F(gitpod.EnvironmentSpecMachineParam{
				Class:   gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
				Session: gitpod.F("session"),
			}),
			Ports: gitpod.F([]gitpod.EnvironmentSpecPortParam{{
				Admission: gitpod.F(gitpod.AdmissionLevelUnspecified),
				Name:      gitpod.F("x"),
				Port:      gitpod.F(int64(1)),
			}}),
			Secrets: gitpod.F([]gitpod.EnvironmentSpecSecretParam{{
				ID:                             gitpod.F("id"),
				ContainerRegistryBasicAuthHost: gitpod.F("containerRegistryBasicAuthHost"),
				EnvironmentVariable:            gitpod.F("environmentVariable"),
				FilePath:                       gitpod.F("filePath"),
				GitCredentialHost:              gitpod.F("gitCredentialHost"),
				Name:                           gitpod.F("name"),
				Session:                        gitpod.F("session"),
				Source:                         gitpod.F("source"),
				SourceRef:                      gitpod.F("sourceRef"),
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

func TestEnvironmentGet(t *testing.T) {
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
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
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
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
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
				ID:    gitpod.F("0194b7c1-c954-718d-91a4-9a742aa5fc11"),
				Value: gitpod.F("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI..."),
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
			ArchivalStatus: gitpod.F(gitpod.EnvironmentListParamsFilterArchivalStatusArchivalStatusUnspecified),
			CreatorIDs:     gitpod.F([]string{"f53d2330-3795-4c5d-a1f3-453121af9c60"}),
			ProjectIDs:     gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			RunnerIDs:      gitpod.F([]string{"e6aa9c54-89d3-42c1-ac31-bd8d8f1concentrate"}),
			RunnerKinds:    gitpod.F([]gitpod.RunnerKind{gitpod.RunnerKindUnspecified}),
			StatusPhases:   gitpod.F([]gitpod.EnvironmentPhase{gitpod.EnvironmentPhaseUnspecified}),
		}),
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
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
		Force:         gitpod.F(false),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnvironmentNewEnvironmentToken(t *testing.T) {
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
	_, err := client.Environments.NewEnvironmentToken(context.TODO(), gitpod.EnvironmentNewEnvironmentTokenParams{
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
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
		ProjectID: gitpod.F("b0e12f6c-4c67-429d-a4a6-d9838b5da047"),
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
				DefaultDevcontainerImage: gitpod.F("defaultDevcontainerImage"),
				DevcontainerFilePath:     gitpod.F("devcontainerFilePath"),
				Dotfiles: gitpod.F(gitpod.EnvironmentSpecDevcontainerDotfilesParam{
					Repository: gitpod.F("https://example.com"),
				}),
				Session: gitpod.F("session"),
			}),
			Machine: gitpod.F(gitpod.EnvironmentSpecMachineParam{
				Class:   gitpod.F("d2c94c27-3b76-4a42-b88c-95a85e392c68"),
				Session: gitpod.F("session"),
			}),
			Ports: gitpod.F([]gitpod.EnvironmentSpecPortParam{{
				Admission: gitpod.F(gitpod.AdmissionLevelUnspecified),
				Name:      gitpod.F("x"),
				Port:      gitpod.F(int64(1)),
			}}),
			Secrets: gitpod.F([]gitpod.EnvironmentSpecSecretParam{{
				ID:                             gitpod.F("id"),
				ContainerRegistryBasicAuthHost: gitpod.F("containerRegistryBasicAuthHost"),
				EnvironmentVariable:            gitpod.F("environmentVariable"),
				FilePath:                       gitpod.F("filePath"),
				GitCredentialHost:              gitpod.F("gitCredentialHost"),
				Name:                           gitpod.F("name"),
				Session:                        gitpod.F("session"),
				Source:                         gitpod.F("source"),
				SourceRef:                      gitpod.F("sourceRef"),
			}}),
			SpecVersion: gitpod.F("specVersion"),
			SSHPublicKeys: gitpod.F([]gitpod.EnvironmentSpecSSHPublicKeyParam{{
				ID:    gitpod.F("id"),
				Value: gitpod.F("value"),
			}}),
			Timeout: gitpod.F(gitpod.EnvironmentSpecTimeoutParam{
				Disconnected: gitpod.F("14400s"),
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
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
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
			Source:    gitpod.F("VS Code"),
			Timestamp: gitpod.F(time.Now()),
		}),
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
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
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
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
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnvironmentUnarchiveWithOptionalParams(t *testing.T) {
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
	_, err := client.Environments.Unarchive(context.TODO(), gitpod.EnvironmentUnarchiveParams{
		EnvironmentID: gitpod.F("07e03a28-65a5-4d98-b532-8ea67b188048"),
	})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
