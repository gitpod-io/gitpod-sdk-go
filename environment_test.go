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
					Specs: gitpod.F([]gitpod.EnvironmentNewParamsSpecContentInitializerSpec{{}}),
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
			Secrets:     gitpod.F([]gitpod.EnvironmentNewParamsSpecSecret{{}}),
			SpecVersion: gitpod.F[gitpod.EnvironmentNewParamsSpecSpecVersionUnion](shared.UnionString("string")),
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
	)
	_, err := client.Environments.Get(context.TODO(), gitpod.EnvironmentGetParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentGetParamsConnectProtocolVersion1),
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
	)
	_, err := client.Environments.List(context.TODO(), gitpod.EnvironmentListParams{
		ConnectProtocolVersion: gitpod.F(gitpod.EnvironmentListParamsConnectProtocolVersion1),
		Filter: gitpod.F(gitpod.EnvironmentListParamsFilter{
			CreatorIDs:   gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ProjectIDs:   gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			RunnerIDs:    gitpod.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			StatusPhases: gitpod.F([]gitpod.EnvironmentListParamsFilterStatusPhase{gitpod.EnvironmentListParamsFilterStatusPhaseEnvironmentPhaseUnspecified}),
		}),
		OrganizationID: gitpod.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Pagination: gitpod.F(gitpod.EnvironmentListParamsPagination{
			Token:    gitpod.F("token"),
			PageSize: gitpod.F(int64(0)),
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
					Specs: gitpod.F([]gitpod.EnvironmentNewFromProjectParamsSpecContentInitializerSpec{{}}),
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
			Secrets:     gitpod.F([]gitpod.EnvironmentNewFromProjectParamsSpecSecret{{}}),
			SpecVersion: gitpod.F[gitpod.EnvironmentNewFromProjectParamsSpecSpecVersionUnion](shared.UnionString("string")),
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
