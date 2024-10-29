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
	"github.com/stainless-sdks/gitpod-go/shared"
)

func TestRunnerInteractionEnvironmentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.Environments.Get(context.TODO(), gitpod.RunnerInteractionEnvironmentGetParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionEnvironmentGetParamsConnectProtocolVersion1),
		EnvironmentID:          gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		RunnerID:               gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRunnerInteractionEnvironmentListWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.Environments.List(context.TODO(), gitpod.RunnerInteractionEnvironmentListParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionEnvironmentListParamsConnectProtocolVersion1),
		EnvironmentIDs:         gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		Pagination: gitpod.F(gitpod.RunnerInteractionEnvironmentListParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(0)),
		}),
		RunnerID:         gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRunnerInteractionEnvironmentUpdateStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.RunnerInteractions.Environments.UpdateStatus(context.TODO(), gitpod.RunnerInteractionEnvironmentUpdateStatusParams{
		ConnectProtocolVersion: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsConnectProtocolVersion1),
		EnvironmentID:          gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		RunnerID:               gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatus{
			AutomationsFile: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFile{
				AutomationsFilePath:     gitpod.F("automationsFilePath"),
				AutomationsFilePresence: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFileAutomationsFilePresencePresenceUnspecified),
				FailureMessage:          gitpod.F("failureMessage"),
				Phase:                   gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusAutomationsFilePhaseContentPhaseUnspecified),
				Session:                 gitpod.F("session"),
			}),
			Content: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusContent{
				ContentLocationInMachine: gitpod.F("contentLocationInMachine"),
				FailureMessage:           gitpod.F("failureMessage"),
				Git: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGit{
					Branch: gitpod.F("branch"),
					ChangedFiles: gitpod.F([]gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFile{{
						ChangeType: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified),
						Path:       gitpod.F("path"),
					}, {
						ChangeType: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified),
						Path:       gitpod.F("path"),
					}, {
						ChangeType: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified),
						Path:       gitpod.F("path"),
					}}),
					CloneURL:             gitpod.F("cloneUrl"),
					LatestCommit:         gitpod.F("latestCommit"),
					TotalChangedFiles:    gitpod.F(int64(0)),
					TotalUnpushedCommits: gitpod.F(int64(0)),
					UnpushedCommits:      gitpod.F([]string{"string", "string", "string"}),
				}),
				Phase:          gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusContentPhaseContentPhaseUnspecified),
				Session:        gitpod.F("session"),
				WarningMessage: gitpod.F("warningMessage"),
			}),
			Devcontainer: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainer{
				ContainerID:              gitpod.F("containerId"),
				ContainerName:            gitpod.F("containerName"),
				DevcontainerconfigInSync: gitpod.F(true),
				DevcontainerFilePath:     gitpod.F("devcontainerFilePath"),
				DevcontainerFilePresence: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerDevcontainerFilePresencePresenceUnspecified),
				FailureMessage:           gitpod.F("failureMessage"),
				Phase:                    gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusDevcontainerPhasePhaseUnspecified),
				RemoteUser:               gitpod.F("remoteUser"),
				RemoteWorkspaceFolder:    gitpod.F("remoteWorkspaceFolder"),
				SecretsInSync:            gitpod.F(true),
				Session:                  gitpod.F("session"),
				WarningMessage:           gitpod.F("warningMessage"),
			}),
			EnvironmentURLs: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLs{
				Logs: gitpod.F("logs"),
				Ports: gitpod.F([]gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLsPort{{
					Port: gitpod.F(int64(1)),
					URL:  gitpod.F("url"),
				}, {
					Port: gitpod.F(int64(1)),
					URL:  gitpod.F("url"),
				}, {
					Port: gitpod.F(int64(1)),
					URL:  gitpod.F("url"),
				}}),
				SSH: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusEnvironmentURLsSSH{
					URL: gitpod.F("url"),
				}),
			}),
			FailureMessage: gitpod.F([]string{"string", "string", "string"}),
			Machine: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusMachine{
				FailureMessage: gitpod.F("failureMessage"),
				Phase:          gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusMachinePhasePhaseUnspecified),
				Session:        gitpod.F("session"),
				Timeout:        gitpod.F("timeout"),
				Versions: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusMachineVersions{
					SupervisorCommit:  gitpod.F("supervisorCommit"),
					SupervisorVersion: gitpod.F("supervisorVersion"),
				}),
				WarningMessage: gitpod.F("warningMessage"),
			}),
			Phase: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusPhaseEnvironmentPhaseUnspecified),
			RunnerAck: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAck{
				Message:     gitpod.F("message"),
				SpecVersion: gitpod.F[gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckSpecVersionUnion](shared.UnionString("string")),
				StatusCode:  gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusRunnerAckStatusCodeStatusCodeUnspecified),
			}),
			Secrets: gitpod.F([]gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusSecret{{
				FailureMessage: gitpod.F("failureMessage"),
				Phase:          gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseUnspecified),
				SecretName:     gitpod.F("secretName"),
				WarningMessage: gitpod.F("warningMessage"),
			}, {
				FailureMessage: gitpod.F("failureMessage"),
				Phase:          gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseUnspecified),
				SecretName:     gitpod.F("secretName"),
				WarningMessage: gitpod.F("warningMessage"),
			}, {
				FailureMessage: gitpod.F("failureMessage"),
				Phase:          gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusSecretsPhaseContentPhaseUnspecified),
				SecretName:     gitpod.F("secretName"),
				WarningMessage: gitpod.F("warningMessage"),
			}}),
			SSHPublicKeys: gitpod.F([]gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKey{{
				ID:    gitpod.F("id"),
				Phase: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseUnspecified),
			}, {
				ID:    gitpod.F("id"),
				Phase: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseUnspecified),
			}, {
				ID:    gitpod.F("id"),
				Phase: gitpod.F(gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusSSHPublicKeysPhaseContentPhaseUnspecified),
			}}),
			StatusVersion:  gitpod.F[gitpod.RunnerInteractionEnvironmentUpdateStatusParamsStatusStatusVersionUnion](shared.UnionString("string")),
			WarningMessage: gitpod.F([]string{"string", "string", "string"}),
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
