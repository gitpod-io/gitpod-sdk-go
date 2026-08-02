package sdk

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	rawclient "github.com/gitpod-io/gitpod-sdk-go/internal/raw"
	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestCodexSettings(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Settings *v1.CodexSettings
		Err      string
	}
	tests := []struct {
		Name     string
		Model    v1.CodexOpenAIModel
		Effort   v1.CodexReasoningEffort
		Expected Expectation
	}{
		{Name: "omits_unspecified_settings"},
		{
			Name:  "sets_model",
			Model: v1.CodexOpenAIModel_CODEX_OPEN_AI_MODEL_GPT_5_6_SOL,
			Expected: Expectation{Settings: &v1.CodexSettings{
				Model: v1.CodexOpenAIModel_CODEX_OPEN_AI_MODEL_GPT_5_6_SOL,
			}},
		},
		{
			Name:   "sets_reasoning_effort",
			Effort: v1.CodexReasoningEffort_CODEX_REASONING_EFFORT_HIGH,
			Expected: Expectation{Settings: &v1.CodexSettings{
				ReasoningEffort: v1.CodexReasoningEffort_CODEX_REASONING_EFFORT_HIGH,
			}},
		},
		{
			Name:     "rejects_unknown_model",
			Model:    v1.CodexOpenAIModel(999),
			Expected: Expectation{Err: "model 999 is not supported"},
		},
		{
			Name:     "rejects_unknown_reasoning_effort",
			Effort:   v1.CodexReasoningEffort(999),
			Expected: Expectation{Err: "reasoning effort 999 is not supported"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			var got Expectation
			if err := validateCodexSettings(tc.Model, tc.Effort); err != nil {
				got.Err = err.Error()
			} else {
				got.Settings = codexSettings(tc.Model, tc.Effort)
			}
			if diff := cmp.Diff(tc.Expected, got, protocmp.Transform()); diff != "" {
				t.Errorf("codexSettings() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClientRunCodexScratch(t *testing.T) {
	t.Parallel()

	type Requests struct {
		Create *v1.CreateEnvironmentRequest
		Start  *v1.StartAgentRequest
		Send   *v1.SendToAgentExecutionRequest
	}
	type Expectation struct {
		EnvironmentID string
		ExecutionID   string
		Requests      Requests
		Err           string
	}

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	client := New(mp.Client())
	runningEnvironment := &v1.Environment{
		Id:     "env-1",
		Status: &v1.EnvironmentStatus{Phase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_RUNNING},
	}

	var got Expectation
	gomock.InOrder(
		mp.EnvironmentService.EXPECT().ListEnvironmentClasses(gomock.Any(), gomock.Any()).
			Return(connect.NewResponse(&v1.ListEnvironmentClassesResponse{EnvironmentClasses: []*v1.EnvironmentClass{{Id: "class-1"}}}), nil),
		mp.EnvironmentService.EXPECT().CreateEnvironment(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, req *connect.Request[v1.CreateEnvironmentRequest]) (*connect.Response[v1.CreateEnvironmentResponse], error) {
				got.Requests.Create = req.Msg
				return connect.NewResponse(&v1.CreateEnvironmentResponse{Environment: runningEnvironment}), nil
			}),
		mp.EnvironmentService.EXPECT().GetEnvironment(gomock.Any(), gomock.Any()).
			Return(connect.NewResponse(&v1.GetEnvironmentResponse{Environment: runningEnvironment}), nil),
		mp.AgentService.EXPECT().StartAgent(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, req *connect.Request[v1.StartAgentRequest]) (*connect.Response[v1.StartAgentResponse], error) {
				got.Requests.Start = req.Msg
				return connect.NewResponse(&v1.StartAgentResponse{AgentExecutionId: "exec-1"}), nil
			}),
		mp.AgentService.EXPECT().SendToAgentExecution(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, req *connect.Request[v1.SendToAgentExecutionRequest]) (*connect.Response[v1.SendToAgentExecutionResponse], error) {
				got.Requests.Send = req.Msg
				got.Requests.Send.GetUserInput().Id = ""
				return connect.NewResponse(&v1.SendToAgentExecutionResponse{}), nil
			}),
		mp.AgentService.EXPECT().GetAgentExecution(gomock.Any(), gomock.Any()).
			Return(connect.NewResponse(&v1.GetAgentExecutionResponse{AgentExecution: &v1.AgentExecution{
				Id:     "exec-1",
				Status: &v1.AgentExecution_Status{Phase: v1.AgentExecution_PHASE_RUNNING},
			}}), nil),
	)

	run, err := client.RunCodex(t.Context(), RunCodexOptions{
		Task:            "Build a CLI.",
		EnvironmentName: "scratch-task",
		AgentName:       "build-cli",
		Model:           v1.CodexOpenAIModel_CODEX_OPEN_AI_MODEL_GPT_5_6_SOL,
		ReasoningEffort: v1.CodexReasoningEffort_CODEX_REASONING_EFFORT_HIGH,
	})
	if err != nil {
		got.Err = err.Error()
	}
	if run != nil {
		got.EnvironmentID = run.EnvironmentID()
		got.ExecutionID = run.ID()
	}

	expected := Expectation{
		EnvironmentID: "env-1",
		ExecutionID:   "exec-1",
		Requests: Requests{
			Create: &v1.CreateEnvironmentRequest{
				Spec: &v1.EnvironmentSpec{
					DesiredPhase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_RUNNING,
					Machine:      &v1.EnvironmentSpec_Machine{Class: "class-1"},
				},
				Name: Pointer("scratch-task"),
			},
			Start: &v1.StartAgentRequest{
				AgentId:     codexAppInEnvironmentAgentID,
				CodeContext: &v1.AgentCodeContext{Context: &v1.AgentCodeContext_EnvironmentId{EnvironmentId: "env-1"}},
				Name:        "build-cli",
				CodexSettings: &v1.CodexSettings{
					Model:           v1.CodexOpenAIModel_CODEX_OPEN_AI_MODEL_GPT_5_6_SOL,
					ReasoningEffort: v1.CodexReasoningEffort_CODEX_REASONING_EFFORT_HIGH,
				},
			},
			Send: &v1.SendToAgentExecutionRequest{
				AgentExecutionId: "exec-1",
				Input: &v1.SendToAgentExecutionRequest_UserInput{UserInput: &v1.UserInputBlock{Inputs: []*v1.UserInputBlock_Input{{
					Input: &v1.UserInputBlock_Input_Text{Text: &v1.UserInputBlock_TextInput{Content: "Build a CLI."}},
				}}}},
			},
		},
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("RunCodex() mismatch (-want +got):\n%s", diff)
	}
}

func TestClientRunCodexValidation(t *testing.T) {
	t.Parallel()

	type Expectation struct{ Err string }
	tests := []struct {
		Name     string
		Client   *Client
		Options  RunCodexOptions
		Expected Expectation
	}{
		{Name: "requires_client", Expected: Expectation{Err: "codex.run: sdk requires a management-plane client"}},
		{Name: "requires_task", Client: New(nil), Expected: Expectation{Err: "codex.run: task is required"}},
		{
			Name:     "rejects_invalid_model",
			Client:   New(nil),
			Options:  RunCodexOptions{Task: "Do work", Model: v1.CodexOpenAIModel(999)},
			Expected: Expectation{Err: "codex.run: model 999 is not supported"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			var got Expectation
			_, err := tc.Client.RunCodex(t.Context(), tc.Options)
			if err != nil {
				got.Err = err.Error()
			}
			if diff := cmp.Diff(tc.Expected, got); diff != "" {
				t.Errorf("RunCodex() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
