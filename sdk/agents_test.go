package sdk

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	rawclient "github.com/gitpod-io/gitpod-sdk-go/internal/raw"
	"github.com/gitpod-io/gitpod-sdk-go/internal/raw/cache"
	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestAgentSessionWatchResultUsesInformerEvents(t *testing.T) {
	t.Parallel()

	agentExecutionID := "00000000-0000-0000-0000-000000000001"
	running := &v1.AgentExecution{
		Id: agentExecutionID,
		Status: &v1.AgentExecution_Status{
			Phase:           v1.AgentExecution_PHASE_RUNNING,
			CurrentActivity: "checking repository",
		},
	}
	stopped := &v1.AgentExecution{
		Id: agentExecutionID,
		Status: &v1.AgentExecution_Status{
			Phase:           v1.AgentExecution_PHASE_STOPPED,
			CurrentActivity: "done",
		},
	}

	type Expectation struct {
		Result       *v1.AgentExecution
		Updates      []string
		WatchRequest *v1.WatchEventsRequest
		GetRequests  []*v1.GetAgentExecutionRequest
		Err          string
	}

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	sdk := New(mp.Client())
	session := &AgentSession{client: sdk.agents(), id: agentExecutionID}
	events := make(chan *v1.WatchEventsResponse, 1)
	watchRequests := make(chan *v1.WatchEventsRequest, 1)

	var got Expectation
	executions := []*v1.AgentExecution{running, stopped}
	getCalls := 0
	mp.AgentService.EXPECT().
		GetAgentExecution(gomock.Any(), gomock.Any()).
		Times(2).
		DoAndReturn(func(_ context.Context, req *connect.Request[v1.GetAgentExecutionRequest]) (*connect.Response[v1.GetAgentExecutionResponse], error) {
			got.GetRequests = append(got.GetRequests, req.Msg)
			exec := executions[getCalls]
			getCalls++
			return connect.NewResponse(&v1.GetAgentExecutionResponse{AgentExecution: exec}), nil
		})

	watchEvents := func(ctx context.Context, req *connect.Request[v1.WatchEventsRequest]) (cache.EventStream, error) {
		watchRequests <- req.Msg
		return &agentExecutionEventStream{ctx: ctx, events: events}, nil
	}
	result, err := session.watchUntil(t.Context(), "agents.watch_result", mp.Client(), watchEvents, func(_ context.Context, exec *v1.AgentExecution) error {
		got.Updates = append(got.Updates, AgentStatusLine(exec))
		if exec.GetStatus().GetPhase() == v1.AgentExecution_PHASE_RUNNING {
			events <- &v1.WatchEventsResponse{
				ResourceId:   agentExecutionID,
				ResourceType: v1.ResourceType_RESOURCE_TYPE_AGENT_EXECUTION,
				Operation:    v1.ResourceOperation_RESOURCE_OPERATION_UPDATE_STATUS,
			}
		}
		return nil
	}, func(exec *v1.AgentExecution) (bool, error) {
		switch exec.GetStatus().GetPhase() {
		case v1.AgentExecution_PHASE_STOPPED,
			v1.AgentExecution_PHASE_WAITING_FOR_INPUT:
			return true, nil
		default:
			return false, nil
		}
	})
	if err != nil {
		got.Err = err.Error()
	}
	got.Result = result
	got.WatchRequest = <-watchRequests

	expected := Expectation{
		Result: stopped,
		Updates: []string{
			`phase=running activity="checking repository"`,
			`phase=stopped activity="done"`,
		},
		WatchRequest: agentExecutionWatchRequest(agentExecutionID),
		GetRequests: []*v1.GetAgentExecutionRequest{
			{AgentExecutionId: agentExecutionID},
			{AgentExecutionId: agentExecutionID},
		},
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("WatchResult() mismatch (-want +got):\n%s", diff)
	}
}

func TestAgentSessionWatchUntilRunningUsesInformerEvents(t *testing.T) {
	t.Parallel()

	agentExecutionID := "00000000-0000-0000-0000-000000000001"
	pending := &v1.AgentExecution{
		Id:     agentExecutionID,
		Status: &v1.AgentExecution_Status{Phase: v1.AgentExecution_PHASE_PENDING},
	}
	running := &v1.AgentExecution{
		Id:     agentExecutionID,
		Status: &v1.AgentExecution_Status{Phase: v1.AgentExecution_PHASE_RUNNING},
	}

	type Expectation struct {
		Result       *v1.AgentExecution
		WatchRequest *v1.WatchEventsRequest
		GetRequests  []*v1.GetAgentExecutionRequest
		Err          string
	}

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	sdk := New(mp.Client())
	session := &AgentSession{client: sdk.agents(), id: agentExecutionID}
	events := make(chan *v1.WatchEventsResponse, 1)
	watchRequests := make(chan *v1.WatchEventsRequest, 1)
	releaseEvent := make(chan struct{})

	var got Expectation
	executions := []*v1.AgentExecution{pending, pending, running}
	getCalls := 0
	mp.AgentService.EXPECT().
		GetAgentExecution(gomock.Any(), gomock.Any()).
		Times(3).
		DoAndReturn(func(_ context.Context, req *connect.Request[v1.GetAgentExecutionRequest]) (*connect.Response[v1.GetAgentExecutionResponse], error) {
			got.GetRequests = append(got.GetRequests, req.Msg)
			exec := executions[getCalls]
			getCalls++
			return connect.NewResponse(&v1.GetAgentExecutionResponse{AgentExecution: exec}), nil
		})

	watchEvents := func(ctx context.Context, req *connect.Request[v1.WatchEventsRequest]) (cache.EventStream, error) {
		watchRequests <- req.Msg
		return &agentExecutionEventStream{ctx: ctx, events: events}, nil
	}
	conditionCalls := 0
	eventReleased := false
	result, err := session.watchUntil(t.Context(), "agents.wait_running", mp.Client(), watchEvents, nil, func(exec *v1.AgentExecution) (bool, error) {
		if exec.GetStatus().GetPhase() == v1.AgentExecution_PHASE_RUNNING {
			return true, nil
		}
		conditionCalls++
		if conditionCalls == 2 && !eventReleased {
			eventReleased = true
			events <- &v1.WatchEventsResponse{
				ResourceId:   agentExecutionID,
				ResourceType: v1.ResourceType_RESOURCE_TYPE_AGENT_EXECUTION,
				Operation:    v1.ResourceOperation_RESOURCE_OPERATION_UPDATE_STATUS,
			}
			close(releaseEvent)
		}
		return false, nil
	})
	if err != nil {
		got.Err = err.Error()
	}
	got.Result = result
	got.WatchRequest = <-watchRequests
	<-releaseEvent

	expected := Expectation{
		Result:       running,
		WatchRequest: agentExecutionWatchRequest(agentExecutionID),
		GetRequests: []*v1.GetAgentExecutionRequest{
			{AgentExecutionId: agentExecutionID},
			{AgentExecutionId: agentExecutionID},
			{AgentExecutionId: agentExecutionID},
		},
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("watchUntil() mismatch (-want +got):\n%s", diff)
	}
}

func TestAgentSessionWaitRunningHandlesTerminalPhase(t *testing.T) {
	t.Parallel()

	agentExecutionID := "00000000-0000-0000-0000-000000000001"
	waiting := &v1.AgentExecution{
		Id:     agentExecutionID,
		Status: &v1.AgentExecution_Status{Phase: v1.AgentExecution_PHASE_WAITING_FOR_INPUT},
	}

	type Expectation struct {
		Result *v1.AgentExecution
		Err    string
	}

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	sdk := New(mp.Client())
	session := &AgentSession{client: sdk.agents(), id: agentExecutionID}

	var got Expectation
	mp.AgentService.EXPECT().
		GetAgentExecution(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, _ *connect.Request[v1.GetAgentExecutionRequest]) (*connect.Response[v1.GetAgentExecutionResponse], error) {
			return connect.NewResponse(&v1.GetAgentExecutionResponse{AgentExecution: waiting}), nil
		})

	result, err := session.waitRunning(t.Context())
	if err != nil {
		got.Err = err.Error()
	}
	got.Result = result

	expected := Expectation{
		Result: waiting,
		Err:    "agents.wait_running: agent execution 00000000-0000-0000-0000-000000000001 reached waiting for input before running; last status: phase=waiting for input",
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("waitRunning() mismatch (-want +got):\n%s", diff)
	}
}

func TestAgentStatusLineIncludesFailureDiagnostics(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Line string
	}

	execution := &v1.AgentExecution{
		Status: &v1.AgentExecution_Status{
			Phase:            v1.AgentExecution_PHASE_STOPPED,
			CurrentActivity:  "starting agent",
			WarningMessage:   "model unavailable",
			FailureReason:    v1.AgentExecutionFailureReason_AGENT_EXECUTION_FAILURE_REASON_SERVICE,
			FailureMessage:   "agent process exited",
			SupportBundleUrl: "https://example.com/support-bundle?token=secret",
		},
	}

	got := Expectation{Line: AgentStatusLine(execution)}
	expected := Expectation{
		Line: `phase=stopped activity="starting agent" warning="model unavailable" failure_reason=AGENT_EXECUTION_FAILURE_REASON_SERVICE failure="agent process exited" support_bundle=https://example.com/support-bundle`,
	}
	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("AgentStatusLine() mismatch (-want +got):\n%s", diff)
	}
}

func TestAgentExecutionLogAttrsOmitsUnspecifiedFailureReason(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Attrs map[string]string
	}

	tests := []struct {
		Name      string
		Execution *v1.AgentExecution
		Expected  Expectation
	}{
		{
			Name: "omits_unspecified_failure_reason",
			Execution: &v1.AgentExecution{
				Id: "exec-1",
				Status: &v1.AgentExecution_Status{
					Phase: v1.AgentExecution_PHASE_RUNNING,
				},
			},
			Expected: Expectation{
				Attrs: map[string]string{
					"activity":           "",
					"agent_execution_id": "exec-1",
					"failure_message":    "",
					"operation":          "agents.watch_result",
					"phase":              "PHASE_RUNNING",
					"status":             "phase=running",
					"support_bundle_url": "",
					"warning":            "",
				},
			},
		},
		{
			Name: "includes_specified_failure_reason",
			Execution: &v1.AgentExecution{
				Id: "exec-1",
				Status: &v1.AgentExecution_Status{
					Phase:         v1.AgentExecution_PHASE_STOPPED,
					FailureReason: v1.AgentExecutionFailureReason_AGENT_EXECUTION_FAILURE_REASON_SERVICE,
				},
			},
			Expected: Expectation{
				Attrs: map[string]string{
					"activity":           "",
					"agent_execution_id": "exec-1",
					"failure_message":    "",
					"failure_reason":     "AGENT_EXECUTION_FAILURE_REASON_SERVICE",
					"operation":          "agents.watch_result",
					"phase":              "PHASE_STOPPED",
					"status":             "phase=stopped failure_reason=AGENT_EXECUTION_FAILURE_REASON_SERVICE",
					"support_bundle_url": "",
					"warning":            "",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			got := Expectation{Attrs: map[string]string{}}
			for _, attr := range agentExecutionLogAttrs("agents.watch_result", tc.Execution) {
				got.Attrs[attr.Key] = attr.Value.Resolve().String()
			}

			if diff := cmp.Diff(tc.Expected, got); diff != "" {
				t.Errorf("agentExecutionLogAttrs() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAgentSessionSendMessage(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Request *v1.SendToAgentExecutionRequest
		Err     string
	}

	tests := []struct {
		Name     string
		Text     string
		Expected Expectation
	}{
		{
			Name: "sends_text_message",
			Text: "Please continue.",
			Expected: Expectation{
				Request: &v1.SendToAgentExecutionRequest{
					AgentExecutionId: "exec-1",
					Input: &v1.SendToAgentExecutionRequest_UserInput{
						UserInput: &v1.UserInputBlock{
							Inputs: []*v1.UserInputBlock_Input{{
								Input: &v1.UserInputBlock_Input_Text{
									Text: &v1.UserInputBlock_TextInput{Content: "Please continue."},
								},
							}},
						},
					},
				},
			},
		},
		{
			Name: "requires_text",
			Expected: Expectation{
				Err: "agents.send_message: text is required",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mp := rawclient.NewMock(ctrl)
			sdk := New(mp.Client())
			session := &AgentSession{client: sdk.agents(), id: "exec-1"}

			var got Expectation
			if tc.Text != "" {
				mp.AgentService.EXPECT().
					SendToAgentExecution(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ context.Context, req *connect.Request[v1.SendToAgentExecutionRequest]) (*connect.Response[v1.SendToAgentExecutionResponse], error) {
						got.Request = req.Msg
						got.Request.GetUserInput().Id = ""
						return connect.NewResponse(&v1.SendToAgentExecutionResponse{}), nil
					})
			}

			err := session.SendMessage(t.Context(), tc.Text)
			if err != nil {
				got.Err = err.Error()
			}

			if diff := cmp.Diff(tc.Expected, got, protocmp.Transform()); diff != "" {
				t.Errorf("SendMessage() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAgentSessionMessageStreamReadsLiveV2Markdown(t *testing.T) {
	t.Parallel()

	type LiveRequest struct {
		Authorization string
		Accept        string
		UserAgent     string
	}
	type CapturedRequests struct {
		Get   *v1.GetAgentExecutionRequest
		Token *v1.CreateAgentExecutionConversationTokenRequest
		Live  LiveRequest
	}
	type Expectation struct {
		Output   string
		Requests CapturedRequests
		Err      string
	}

	liveRequests := make(chan LiveRequest, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		liveRequests <- LiveRequest{
			Authorization: r.Header.Get("Authorization"),
			Accept:        r.Header.Get("Accept"),
			UserAgent:     r.Header.Get("User-Agent"),
		}
		w.Header().Set("Content-Type", "text/event-stream")
		writeAgentLiveBlock(t, w, agentConversationTagUserInput, &v1.UserInputBlock{
			Id: "input-1",
			Inputs: []*v1.UserInputBlock_Input{{
				Input: &v1.UserInputBlock_Input_Text{
					Text: &v1.UserInputBlock_TextInput{Content: "continue please"},
				},
			}},
		})
		writeAgentLiveBlock(t, w, agentConversationTagAgentResponse, &v1.AgentResponseBlock{
			Id:    "block-1",
			Phase: v1.AgentResponseBlock_PHASE_UPDATE,
			Output: &v1.AgentResponseBlock_Text{
				Text: &v1.AgentResponseBlock_TextOutput{Content: "Hello, ", SequenceId: 1},
			},
		})
		writeAgentLiveBlock(t, w, agentConversationTagAgentResponse, &v1.AgentResponseBlock{
			Id:    "block-1",
			Phase: v1.AgentResponseBlock_PHASE_UPDATE,
			Output: &v1.AgentResponseBlock_Text{
				Text: &v1.AgentResponseBlock_TextOutput{Content: "Hello, ", SequenceId: 1},
			},
		})
		writeAgentLiveBlock(t, w, agentConversationTagAgentResponse, &v1.AgentResponseBlock{
			Id:    "block-1",
			Phase: v1.AgentResponseBlock_PHASE_UPDATE,
			Output: &v1.AgentResponseBlock_Text{
				Text: &v1.AgentResponseBlock_TextOutput{Content: "**world**", SequenceId: 2},
			},
		})
		writeAgentLiveBlock(t, w, agentConversationTagAgentResponse, &v1.AgentResponseBlock{
			Id:    "block-1",
			Phase: v1.AgentResponseBlock_PHASE_COMPLETED,
			Output: &v1.AgentResponseBlock_Text{
				Text: &v1.AgentResponseBlock_TextOutput{Content: "Hello, **world**"},
			},
		})
		writeAgentLiveBlock(t, w, agentConversationTagAgentResponse, &v1.AgentResponseBlock{
			Id:    "thought-1",
			Phase: v1.AgentResponseBlock_PHASE_COMPLETED,
			Output: &v1.AgentResponseBlock_Text{
				Text: &v1.AgentResponseBlock_TextOutput{
					Type:    v1.AgentResponseBlock_TextOutput_TYPE_THOUGHTS,
					Content: "private reasoning",
				},
			},
		})
		writeAgentLiveBlock(t, w, agentConversationTagAgentMessage, &v1.AgentMessage{Payload: "subagent note"})
		_, _ = io.WriteString(w, "event: end\n\n")
	}))
	t.Cleanup(server.Close)

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	sdk := New(mp.Client())
	session := &AgentSession{
		client: sdk.agents(),
		id:     "exec-1",
		execution: &v1.AgentExecution{
			Id: "exec-1",
			Status: &v1.AgentExecution_Status{
				ConversationUrls: &v1.AgentExecution_Status_ConversationURLs{Live: server.URL},
			},
		},
	}

	var got Expectation
	mp.AgentService.EXPECT().
		CreateAgentExecutionConversationToken(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *connect.Request[v1.CreateAgentExecutionConversationTokenRequest]) (*connect.Response[v1.CreateAgentExecutionConversationTokenResponse], error) {
			got.Requests.Token = req.Msg
			return connect.NewResponse(&v1.CreateAgentExecutionConversationTokenResponse{Token: "stream-token"}), nil
		})

	stream, err := session.MessageStream(t.Context())
	if err != nil {
		got.Err = err.Error()
	} else {
		t.Cleanup(func() {
			_ = stream.Close()
		})
		output, err := io.ReadAll(stream)
		if err != nil {
			got.Err = err.Error()
		}
		got.Output = string(output)
		got.Requests.Live = <-liveRequests
	}

	expected := Expectation{
		Output: "> continue please\n\nHello, **world**\n\n> subagent note\n\n",
		Requests: CapturedRequests{
			Token: &v1.CreateAgentExecutionConversationTokenRequest{AgentExecutionId: "exec-1"},
			Live: LiveRequest{
				Authorization: "Bearer stream-token",
				Accept:        "text/event-stream",
				UserAgent:     rawclient.SDKUserAgent(),
			},
		},
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("MessageStream() mismatch (-want +got):\n%s", diff)
	}
}

func TestAgentSessionMessageStreamRequiresLiveURL(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		GetRequest *v1.GetAgentExecutionRequest
		Err        string
	}

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	sdk := New(mp.Client())
	session := &AgentSession{client: sdk.agents(), id: "exec-1"}

	var got Expectation
	mp.AgentService.EXPECT().
		GetAgentExecution(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *connect.Request[v1.GetAgentExecutionRequest]) (*connect.Response[v1.GetAgentExecutionResponse], error) {
			got.GetRequest = req.Msg
			return connect.NewResponse(&v1.GetAgentExecutionResponse{
				AgentExecution: &v1.AgentExecution{Id: "exec-1", Status: &v1.AgentExecution_Status{}},
			}), nil
		})

	_, err := session.MessageStream(t.Context())
	if err != nil {
		got.Err = err.Error()
	}

	expected := Expectation{
		GetRequest: &v1.GetAgentExecutionRequest{AgentExecutionId: "exec-1"},
		Err:        "agents.message_stream: agent execution does not expose a v2 live conversation URL",
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("MessageStream() mismatch (-want +got):\n%s", diff)
	}
}

func TestAgentExecutionInformerListUsesExecutionFilter(t *testing.T) {
	t.Parallel()

	agentExecutionID := "00000000-0000-0000-0000-000000000001"
	running := &v1.AgentExecution{
		Id:     agentExecutionID,
		Status: &v1.AgentExecution_Status{Phase: v1.AgentExecution_PHASE_RUNNING},
	}
	stopped := &v1.AgentExecution{
		Id:     agentExecutionID,
		Status: &v1.AgentExecution_Status{Phase: v1.AgentExecution_PHASE_STOPPED},
	}

	type Expectation struct {
		ListRequest *v1.ListAgentExecutionsRequest
		Result      []*v1.AgentExecution
		Err         string
	}

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)

	mp.AgentService.EXPECT().
		GetAgentExecution(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, _ *connect.Request[v1.GetAgentExecutionRequest]) (*connect.Response[v1.GetAgentExecutionResponse], error) {
			return connect.NewResponse(&v1.GetAgentExecutionResponse{AgentExecution: running}), nil
		})

	var got Expectation
	mp.AgentService.EXPECT().
		ListAgentExecutions(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *connect.Request[v1.ListAgentExecutionsRequest]) (*connect.Response[v1.ListAgentExecutionsResponse], error) {
			got.ListRequest = req.Msg
			return connect.NewResponse(&v1.ListAgentExecutionsResponse{
				Pagination:      &v1.PaginationResponse{},
				AgentExecutions: []*v1.AgentExecution{stopped},
			}), nil
		})

	ctx, cancel := context.WithCancel(t.Context())
	t.Cleanup(cancel)
	informer, err := cache.NewAgentExecutionCache(ctx, mp.AgentService, agentExecutionID, cache.WithNoFullSync())
	if err != nil {
		got.Err = err.Error()
	} else {
		t.Cleanup(func() {
			_ = informer.Close()
		})

		if _, err := informer.Get(ctx, agentExecutionID); err != nil {
			got.Err = err.Error()
		} else {
			informer.InvalidateItem(ctx, agentExecutionID)
			got.Result, err = informer.List(ctx)
			if err != nil {
				got.Err = err.Error()
			}
		}
	}

	expected := Expectation{
		ListRequest: &v1.ListAgentExecutionsRequest{
			Pagination: &v1.PaginationRequest{PageSize: 100},
			Filter: &v1.ListAgentExecutionsRequest_Filter{
				AgentExecutionIds: []string{agentExecutionID},
			},
		},
		Result: []*v1.AgentExecution{stopped},
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("Informer.List() mismatch (-want +got):\n%s", diff)
	}
}

func writeAgentLiveBlock(t *testing.T, w io.Writer, tag byte, message proto.Message) {
	t.Helper()

	payload, err := proto.Marshal(message)
	if err != nil {
		t.Fatalf("marshal live block: %v", err)
	}
	frame := append([]byte{tag}, payload...)
	data, err := json.Marshal(agentLiveBlockEvent{Frame: base64.StdEncoding.EncodeToString(frame)})
	if err != nil {
		t.Fatalf("marshal live block event: %v", err)
	}
	_, err = fmt.Fprintf(w, "event: block\ndata: %s\n\n", data)
	if err != nil {
		t.Fatalf("write live block event: %v", err)
	}
}

type agentExecutionEventStream struct {
	ctx    context.Context
	events <-chan *v1.WatchEventsResponse
	msg    *v1.WatchEventsResponse
	err    error
}

func (s *agentExecutionEventStream) Receive() bool {
	select {
	case event, ok := <-s.events:
		if !ok {
			return false
		}
		s.msg = event
		return true
	case <-s.ctx.Done():
		s.err = s.ctx.Err()
		return false
	}
}

func (s *agentExecutionEventStream) Msg() *v1.WatchEventsResponse {
	return s.msg
}

func (s *agentExecutionEventStream) Err() error {
	return s.err
}

func (s *agentExecutionEventStream) Close() error {
	return nil
}
