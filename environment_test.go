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

func TestEnvironmentNewWithOptionalParams(t *testing.T) {
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
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentNewParamsConnectProtocolVersion1),
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
					Specs: gitpod.F([]gitpod.EnvironmentNewParamsSpecContentInitializerSpecUnion{gitpod.EnvironmentNewParamsSpecContentInitializerSpecsObject{
						ContextURL: gitpod.F(gitpod.EnvironmentNewParamsSpecContentInitializerSpecsObjectContextURL{
							URL: gitpod.F("https://example.com"),
						}),
						Git: gitpod.F(gitpod.EnvironmentNewParamsSpecContentInitializerSpecsObjectGit{
							CheckoutLocation:  gitpod.F("checkoutLocation"),
							CloneTarget:       gitpod.F("cloneTarget"),
							RemoteUri:         gitpod.F("remoteUri"),
							TargetMode:        gitpod.F(gitpod.EnvironmentNewParamsSpecContentInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified),
							UpstreamRemoteUri: gitpod.F("upstreamRemoteUri"),
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
				FilePath:            gitpod.F("filePath"),
				GitCredentialHost:   gitpod.F("gitCredentialHost"),
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

func TestEnvironmentGetWithOptionalParams(t *testing.T) {
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
		Encoding:               gitpod.F(gitpod.EnvironmentGetParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentGetParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.EnvironmentGetParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.EnvironmentGetParamsConnectV1),
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

func TestEnvironmentUpdateWithOptionalParams(t *testing.T) {
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
		Body:                   gitpod.EnvironmentUpdateParamsBody{},
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentUpdateParamsConnectProtocolVersion1),
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

func TestEnvironmentListWithOptionalParams(t *testing.T) {
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
		Encoding:               gitpod.F(gitpod.EnvironmentListParamsEncodingProto),
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentListParamsConnectProtocolVersion1),
		Base64:                 gitpod.F(true),
		Compression:            gitpod.F(gitpod.EnvironmentListParamsCompressionIdentity),
		Connect:                gitpod.F(gitpod.EnvironmentListParamsConnectV1),
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

func TestEnvironmentDeleteWithOptionalParams(t *testing.T) {
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
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentDeleteParamsConnectProtocolVersion1),
		EnvironmentID:          gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Force:                  gitpod.F(true),
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

func TestEnvironmentNewFromProjectWithOptionalParams(t *testing.T) {
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
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentNewFromProjectParamsConnectProtocolVersion1),
		ProjectID:              gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
					Specs: gitpod.F([]gitpod.EnvironmentNewFromProjectParamsSpecContentInitializerSpecUnion{gitpod.EnvironmentNewFromProjectParamsSpecContentInitializerSpecsObject{
						ContextURL: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecContentInitializerSpecsObjectContextURL{
							URL: gitpod.F("https://example.com"),
						}),
						Git: gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecContentInitializerSpecsObjectGit{
							CheckoutLocation:  gitpod.F("checkoutLocation"),
							CloneTarget:       gitpod.F("cloneTarget"),
							RemoteUri:         gitpod.F("remoteUri"),
							TargetMode:        gitpod.F(gitpod.EnvironmentNewFromProjectParamsSpecContentInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified),
							UpstreamRemoteUri: gitpod.F("upstreamRemoteUri"),
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
				FilePath:            gitpod.F("filePath"),
				GitCredentialHost:   gitpod.F("gitCredentialHost"),
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

func TestEnvironmentNewLogsTokenWithOptionalParams(t *testing.T) {
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
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentNewLogsTokenParamsConnectProtocolVersion1),
		EnvironmentID:          gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestEnvironmentMarkActiveWithOptionalParams(t *testing.T) {
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
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentMarkActiveParamsConnectProtocolVersion1),
		ActivitySignal: gitpod.F(gitpod.EnvironmentMarkActiveParamsActivitySignal{
			Source:    gitpod.F("xxx"),
			Timestamp: gitpod.F(time.Now()),
		}),
		EnvironmentID:    gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestEnvironmentStartWithOptionalParams(t *testing.T) {
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
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentStartParamsConnectProtocolVersion1),
		EnvironmentID:          gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestEnvironmentStopWithOptionalParams(t *testing.T) {
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
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentStopParamsConnectProtocolVersion1),
		EnvironmentID:          gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
