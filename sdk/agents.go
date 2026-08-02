package sdk

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	rawclient "github.com/gitpod-io/gitpod-sdk-go/internal/raw"
	"github.com/gitpod-io/gitpod-sdk-go/internal/raw/cache"
	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
)

const codexAppInEnvironmentAgentID = "00000000-0000-0000-0000-000000007800"

const (
	agentConversationTagAgentResponse byte = 1
	agentConversationTagUserInput     byte = 2
	agentConversationTagAgentMessage  byte = 3
)

type agentClient struct {
	sdk *Client
}

type startCodexOptions struct {
	EnvironmentID   string
	Prompt          string
	PromptInputID   string
	Name            string
	Model           v1.CodexOpenAIModel
	ReasoningEffort v1.CodexReasoningEffort
}

// EnvironmentCodexOptions starts Codex inside an environment.
type EnvironmentCodexOptions struct {
	// Prompt is the required first user message sent to Codex before waiting for it to run.
	Prompt string
	Name   string
	// Model selects the Codex model. Unspecified uses the backend default.
	Model v1.CodexOpenAIModel
	// ReasoningEffort selects the Codex reasoning effort. Unspecified uses the backend default.
	ReasoningEffort v1.CodexReasoningEffort
}

// AgentSession is a handle to a Codex agent execution.
type AgentSession struct {
	client    *agentClient
	id        string
	execution *v1.AgentExecution
}

// AgentExecutionUpdateFunc is called with the latest agent execution snapshot.
type AgentExecutionUpdateFunc func(context.Context, *v1.AgentExecution) error

type agentLiveBlockEvent struct {
	Frame string `json:"frame"`
}

type agentMessageStream struct {
	*io.PipeReader
	cancel context.CancelFunc
	body   io.Closer
}

type agentMarkdownStreamRenderer struct {
	textBlocks      map[string]agentRenderedTextBlock
	openTextBlockID string
}

type agentRenderedTextBlock struct {
	sawDelta       bool
	lastSequenceID uint64
}

// StartCodex starts Codex inside this environment.
func (e *Environment) StartCodex(ctx context.Context, opts EnvironmentCodexOptions) (*AgentSession, error) {
	operation := "environments.start_codex"
	if e == nil || e.sdk == nil || e.environment == nil {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("environment is not initialized")}}
	}
	if strings.TrimSpace(opts.Prompt) == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("prompt is required")}}
	}
	if err := validateCodexSettings(opts.Model, opts.ReasoningEffort); err != nil {
		return nil, &ValidationError{operationError{operation: operation, err: err}}
	}
	ctx = e.sdk.requestContext(ctx)
	e.sdk.logger().DebugContext(ctx, "starting Codex inside environment",
		"operation", operation,
		"environment_id", e.environment.idString(),
		"name", opts.Name,
		"prompt_bytes", len(opts.Prompt),
	)
	session, err := e.sdk.agents().startCodex(ctx, startCodexOptions{
		EnvironmentID:   e.environment.idString(),
		Prompt:          opts.Prompt,
		Name:            opts.Name,
		Model:           opts.Model,
		ReasoningEffort: opts.ReasoningEffort,
	})
	if err != nil {
		return nil, err
	}
	if _, err := session.waitRunning(ctx); err != nil {
		if session.execution != nil {
			e.sdk.logger().DebugContext(ctx, "Codex inside environment did not reach running",
				"operation", operation,
				"environment_id", e.environment.idString(),
				"agent_execution_id", session.ID(),
				"status", AgentStatusLine(session.execution),
				"err", err,
			)
		}
		return session, err
	}
	e.sdk.logger().DebugContext(ctx, "Codex inside environment is running",
		"operation", operation,
		"environment_id", e.environment.idString(),
		"agent_execution_id", session.ID(),
	)
	return session, nil
}

func (c *agentClient) startCodex(ctx context.Context, opts startCodexOptions) (*AgentSession, error) {
	operation := "codex.start"
	raw, err := c.raw(operation)
	if err != nil {
		return nil, err
	}
	codeContext, err := buildCodexCodeContext(opts)
	if err != nil {
		return nil, &ValidationError{operationError{operation: operation, err: err}}
	}
	if strings.TrimSpace(opts.Prompt) == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("prompt is required")}}
	}

	c.sdk.logger().DebugContext(ctx, "starting Codex agent",
		"operation", operation,
		"environment_id", opts.EnvironmentID,
		"name", opts.Name,
		"has_prompt", opts.Prompt != "",
	)
	resp, err := raw.AgentService().StartAgent(ctx, connect.NewRequest(&v1.StartAgentRequest{
		AgentId:       codexAppInEnvironmentAgentID,
		CodeContext:   codeContext,
		Name:          opts.Name,
		CodexSettings: codexSettings(opts.Model, opts.ReasoningEffort),
	}))
	if err != nil {
		return nil, MapError(operation, err)
	}

	session := &AgentSession{
		client: c,
		id:     resp.Msg.GetAgentExecutionId(),
	}
	c.sdk.logger().DebugContext(ctx, "Codex agent started",
		"operation", operation,
		"agent_execution_id", session.id,
		"environment_id", opts.EnvironmentID,
	)
	if err := session.sendText(ctx, opts.Prompt, opts.PromptInputID); err != nil {
		return session, err
	}
	return session, nil
}

func codexSettings(model v1.CodexOpenAIModel, effort v1.CodexReasoningEffort) *v1.CodexSettings {
	if model == v1.CodexOpenAIModel_CODEX_OPEN_AI_MODEL_UNSPECIFIED && effort == v1.CodexReasoningEffort_CODEX_REASONING_EFFORT_UNSPECIFIED {
		return nil
	}
	return &v1.CodexSettings{Model: model, ReasoningEffort: effort}
}

func validateCodexSettings(model v1.CodexOpenAIModel, effort v1.CodexReasoningEffort) error {
	if _, ok := v1.CodexOpenAIModel_name[int32(model)]; !ok {
		return fmt.Errorf("model %d is not supported", model)
	}
	if _, ok := v1.CodexReasoningEffort_name[int32(effort)]; !ok {
		return fmt.Errorf("reasoning effort %d is not supported", effort)
	}
	return nil
}

// ID returns the agent execution ID.
func (s *AgentSession) ID() string {
	if s == nil {
		return ""
	}
	return s.id
}

// SendMessage sends a text message to the agent execution.
func (s *AgentSession) SendMessage(ctx context.Context, text string) error {
	return s.sendText(ctx, text, "")
}

// MessageStream returns a live markdown stream for new conversation messages.
//
// The stream starts at the current live position and does not fetch history.
func (s *AgentSession) MessageStream(ctx context.Context) (io.ReadCloser, error) {
	operation := "agents.message_stream"
	raw, err := s.raw(operation)
	if err != nil {
		return nil, err
	}
	ctx = s.client.sdk.requestContext(ctx)
	if raw.AgentService() == nil {
		return nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("agent service is not configured")}}
	}

	s.client.sdk.logger().DebugContext(ctx, "opening agent message stream",
		"operation", operation,
		"agent_execution_id", s.id,
	)
	var liveURL string
	if s.execution != nil {
		liveURL = s.execution.GetStatus().GetConversationUrls().GetLive()
	}
	if liveURL == "" {
		executionResp, err := raw.AgentService().GetAgentExecution(ctx, connect.NewRequest(&v1.GetAgentExecutionRequest{
			AgentExecutionId: s.id,
		}))
		if err != nil {
			return nil, MapError(operation, err)
		}
		s.execution = executionResp.Msg.GetAgentExecution()
		liveURL = s.execution.GetStatus().GetConversationUrls().GetLive()
	}
	if liveURL == "" {
		return nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("agent execution does not expose a v2 live conversation URL")}}
	}

	tokenResp, err := raw.AgentService().CreateAgentExecutionConversationToken(ctx, connect.NewRequest(&v1.CreateAgentExecutionConversationTokenRequest{
		AgentExecutionId: s.id,
	}))
	if err != nil {
		return nil, MapError(operation, err)
	}
	token := tokenResp.Msg.GetToken()
	if token == "" {
		return nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("agent conversation token response did not include a token")}}
	}

	streamCtx, cancel := context.WithCancel(ctx)
	req, err := http.NewRequestWithContext(streamCtx, http.MethodGet, liveURL, nil)
	if err != nil {
		cancel()
		return nil, &ValidationError{operationError{operation: operation, err: fmt.Errorf("create live stream request for %s: %w", safeURLForLog(liveURL), err)}}
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("User-Agent", rawclient.SDKUserAgent())

	resp, err := s.client.sdk.config().httpClient.Do(req) //nolint:bodyclose // The returned stream owns the body on success.
	if err != nil {
		cancel()
		return nil, &UnavailableError{operationError{operation: operation, err: fmt.Errorf("connect to live agent message stream at %s: %w", safeURLForLog(liveURL), err)}}
	}
	if resp.StatusCode == http.StatusNotFound {
		_ = resp.Body.Close()
		cancel()
		return nil, &CapabilityUnavailableError{operationError{operation: operation, err: fmt.Errorf("v2 live agent message stream is not available at %s", safeURLForLog(liveURL))}}
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		_ = resp.Body.Close()
		cancel()
		return nil, &UnavailableError{operationError{operation: operation, err: fmt.Errorf("live agent message stream at %s returned HTTP %d", safeURLForLog(liveURL), resp.StatusCode)}}
	}

	reader, writer := io.Pipe()
	stream := &agentMessageStream{PipeReader: reader, cancel: cancel, body: resp.Body}
	go func() {
		defer func() { _ = resp.Body.Close() }()
		renderer := agentMarkdownStreamRenderer{textBlocks: map[string]agentRenderedTextBlock{}}
		err := renderer.consume(streamCtx, resp.Body, writer)
		if err != nil && !errors.Is(err, context.Canceled) {
			_ = writer.CloseWithError(err)
			return
		}
		_ = writer.Close()
	}()

	return stream, nil
}

func (s *agentMessageStream) Close() error {
	s.cancel()
	if s.body != nil {
		_ = s.body.Close()
	}
	return s.PipeReader.Close()
}

func (r *agentMarkdownStreamRenderer) consume(ctx context.Context, body io.Reader, writer *io.PipeWriter) error {
	input := bufio.NewReader(body)
	var event string
	var dataLines []string

	for {
		line, err := input.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
			if line == "" {
				done, err := r.consumeEvent(ctx, writer, event, dataLines)
				if err != nil || done {
					return err
				}
				event = ""
				dataLines = nil
			} else if strings.HasPrefix(line, ":") {
				continue
			} else {
				field, value, ok := strings.Cut(line, ":")
				if !ok {
					continue
				}
				value = strings.TrimPrefix(value, " ")
				switch field {
				case "event":
					event = value
				case "data":
					dataLines = append(dataLines, value)
				}
			}
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				if event != "" || len(dataLines) > 0 {
					if _, err := r.consumeEvent(ctx, writer, event, dataLines); err != nil {
						return err
					}
				}
				if tail := r.closeOpenTextBlock(); tail != "" {
					if _, err := io.WriteString(writer, tail); err != nil {
						return err
					}
				}
				return nil
			}
			if ctx.Err() != nil {
				return ctx.Err()
			}
			return err
		}
	}
}

func (r *agentMarkdownStreamRenderer) consumeEvent(ctx context.Context, writer *io.PipeWriter, event string, dataLines []string) (bool, error) {
	switch event {
	case "end":
		if tail := r.closeOpenTextBlock(); tail != "" {
			if _, err := io.WriteString(writer, tail); err != nil {
				return true, err
			}
		}
		return true, nil
	case "block":
		if len(dataLines) == 0 {
			return false, nil
		}
		var payload agentLiveBlockEvent
		if err := json.Unmarshal([]byte(strings.Join(dataLines, "")), &payload); err != nil {
			return false, fmt.Errorf("parse live agent message stream block payload: %w", err)
		}
		if payload.Frame == "" {
			return false, nil
		}
		frame, err := base64.StdEncoding.DecodeString(payload.Frame)
		if err != nil {
			return false, fmt.Errorf("decode live agent message stream block payload: %w", err)
		}
		rendered, err := r.renderFrame(frame)
		if err != nil {
			return false, err
		}
		if rendered == "" {
			return false, nil
		}
		if ctx.Err() != nil {
			return false, ctx.Err()
		}
		_, err = io.WriteString(writer, rendered)
		return false, err
	default:
		return false, nil
	}
}

func (r *agentMarkdownStreamRenderer) renderFrame(frame []byte) (string, error) {
	if len(frame) == 0 {
		return "", nil
	}

	payload := frame[1:]
	switch frame[0] {
	case agentConversationTagAgentResponse:
		var block v1.AgentResponseBlock
		if err := proto.Unmarshal(payload, &block); err != nil {
			return "", fmt.Errorf("decode agent response block: %w", err)
		}
		return r.renderAgentResponseBlock(&block), nil
	case agentConversationTagUserInput:
		var block v1.UserInputBlock
		if err := proto.Unmarshal(payload, &block); err != nil {
			return "", fmt.Errorf("decode user input block: %w", err)
		}
		return r.closeOpenTextBlock() + renderUserInputBlockMarkdown(&block), nil
	case agentConversationTagAgentMessage:
		var message v1.AgentMessage
		if err := proto.Unmarshal(payload, &message); err != nil {
			return "", fmt.Errorf("decode agent message block: %w", err)
		}
		return r.closeOpenTextBlock() + renderBlockquoteMarkdown(message.GetPayload()), nil
	default:
		return "", nil
	}
}

func (r *agentMarkdownStreamRenderer) renderAgentResponseBlock(block *v1.AgentResponseBlock) string {
	text := block.GetText()
	if text == nil || text.GetContent() == "" || text.GetType() == v1.AgentResponseBlock_TextOutput_TYPE_THOUGHTS {
		return ""
	}

	blockID := block.GetId()
	isDelta := block.GetPhase() == v1.AgentResponseBlock_PHASE_DELTA || block.GetPhase() == v1.AgentResponseBlock_PHASE_UPDATE
	prefix := ""
	if r.openTextBlockID != "" && r.openTextBlockID != blockID {
		prefix = r.closeOpenTextBlock()
	}

	if blockID == "" {
		return prefix + markdownBlock(text.GetContent())
	}
	state := r.textBlocks[blockID]
	if isDelta {
		if text.GetSequenceId() > 0 && text.GetSequenceId() <= state.lastSequenceID {
			return ""
		}
		state.sawDelta = true
		state.lastSequenceID = max(state.lastSequenceID, text.GetSequenceId())
		r.textBlocks[blockID] = state
		r.openTextBlockID = blockID
		return prefix + text.GetContent()
	}

	delete(r.textBlocks, blockID)
	r.openTextBlockID = ""
	if state.sawDelta {
		return markdownBlock("")
	}
	return prefix + markdownBlock(text.GetContent())
}

func (r *agentMarkdownStreamRenderer) closeOpenTextBlock() string {
	if r.openTextBlockID == "" {
		return ""
	}
	r.openTextBlockID = ""
	return markdownBlock("")
}

func renderUserInputBlockMarkdown(block *v1.UserInputBlock) string {
	var out strings.Builder
	message := block.ProtoReflect()
	if field := message.Descriptor().Fields().ByName("text"); field != nil && message.Has(field) {
		text := message.Get(field).Message()
		content := text.Descriptor().Fields().ByName("content")
		out.WriteString(renderBlockquoteMarkdown(text.Get(content).String()))
	}
	for _, input := range block.GetInputs() {
		if text := input.GetText(); text != nil {
			out.WriteString(renderBlockquoteMarkdown(text.GetContent()))
			continue
		}
		if input.GetImage() != nil {
			out.WriteString(renderBlockquoteMarkdown("[image]"))
		}
	}
	return out.String()
}

func renderBlockquoteMarkdown(content string) string {
	if content == "" {
		return ""
	}
	var out strings.Builder
	for _, line := range strings.Split(content, "\n") {
		out.WriteString("> ")
		out.WriteString(line)
		out.WriteByte('\n')
	}
	out.WriteByte('\n')
	return out.String()
}

func markdownBlock(content string) string {
	if content == "" {
		return "\n\n"
	}
	if strings.HasSuffix(content, "\n\n") {
		return content
	}
	if strings.HasSuffix(content, "\n") {
		return content + "\n"
	}
	return content + "\n\n"
}

func (s *AgentSession) waitRunning(ctx context.Context) (*v1.AgentExecution, error) {
	operation := "agents.wait_running"
	raw, err := s.raw(operation)
	if err != nil {
		return nil, err
	}
	ctx = s.client.sdk.requestContext(ctx)
	if raw.AgentService() == nil {
		return nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("agent service is not configured")}}
	}

	initial, err := cache.NewAgentExecutionCache(ctx, raw.AgentService(), s.id, cache.WithNoFullSync())
	if err != nil {
		return nil, &SDKError{operationError{operation: operation, err: err}}
	}
	defer func() { _ = initial.Close() }()

	last, err := s.observeAgentExecution(ctx, operation, initial, nil)
	if err != nil {
		return last, err
	}

	running := func(exec *v1.AgentExecution) (bool, error) {
		phase := exec.GetStatus().GetPhase()
		switch phase {
		case v1.AgentExecution_PHASE_RUNNING:
			return true, nil
		case v1.AgentExecution_PHASE_STOPPED,
			v1.AgentExecution_PHASE_WAITING_FOR_INPUT:
			return false, &UnavailableError{operationError{
				operation: operation,
				err:       fmt.Errorf("agent execution %s reached %s before running; last status: %s", exec.GetId(), agentPhaseLabel(phase), AgentStatusLine(exec)),
			}}
		default:
			return false, nil
		}
	}
	if done, err := running(last); err != nil || done {
		return last, err
	}
	if raw.EventService() == nil {
		return last, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("event service is not configured")}}
	}

	s.client.sdk.logger().DebugContext(ctx, "waiting for agent execution to run",
		"operation", operation,
		"agent_execution_id", s.id,
	)
	return s.watchUntil(ctx, operation, raw, cache.AdaptWatchEvents(raw.EventService()), nil, running)
}

// WatchResult waits for the agent to stop or wait for input using an event-backed informer.
func (s *AgentSession) WatchResult(ctx context.Context, onUpdate AgentExecutionUpdateFunc) (*v1.AgentExecution, error) {
	operation := "agents.watch_result"
	raw, err := s.raw(operation)
	if err != nil {
		return nil, err
	}
	ctx = s.client.sdk.requestContext(ctx)
	if raw.AgentService() == nil {
		return nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("agent service is not configured")}}
	}
	if raw.EventService() == nil {
		return nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("event service is not configured")}}
	}

	s.client.sdk.logger().DebugContext(ctx, "watching agent execution result",
		"operation", operation,
		"agent_execution_id", s.id,
	)
	return s.watchUntil(ctx, operation, raw, cache.AdaptWatchEvents(raw.EventService()), onUpdate, func(exec *v1.AgentExecution) (bool, error) {
		switch exec.GetStatus().GetPhase() {
		case v1.AgentExecution_PHASE_STOPPED,
			v1.AgentExecution_PHASE_WAITING_FOR_INPUT:
			return true, nil
		default:
			return false, nil
		}
	})
}

func (s *AgentSession) watchUntil(ctx context.Context, operation string, raw *rawclient.ManagementPlane, watchEvents cache.WatchEventsFunc, onUpdate AgentExecutionUpdateFunc, condition func(*v1.AgentExecution) (bool, error)) (*v1.AgentExecution, error) {
	if watchEvents == nil {
		return nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("event watcher is not configured")}}
	}

	changes := make(chan struct{}, 64)
	watchCtx, cancel := context.WithCancel(ctx)

	s.client.sdk.logger().DebugContext(ctx, "creating agent execution informer",
		"operation", operation,
		"agent_execution_id", s.id,
	)
	informer, err := cache.NewAgentExecutionCache(watchCtx, raw.AgentService(), s.id,
		cache.WithNoFullSync(),
	)
	if err != nil {
		cancel()
		return nil, &SDKError{operationError{operation: operation, err: err}}
	}
	defer func() { _ = informer.Close() }()
	defer cancel()
	watchStarted := make(chan error, 1)
	watchDone := make(chan error, 1)

	go func() {
		s.client.sdk.logger().DebugContext(watchCtx, "watching agent execution events",
			"operation", operation,
			"agent_execution_id", s.id,
		)
		stream, err := watchEvents(watchCtx, connect.NewRequest(agentExecutionWatchRequest(s.id)))
		if err != nil {
			if !errors.Is(err, context.Canceled) {
				s.client.sdk.logger().ErrorContext(watchCtx, "agent execution informer event invalidation failed", "agent_execution_id", s.id, "err", err)
			}
			watchStarted <- err
			return
		}
		watchStarted <- nil
		defer func() { _ = stream.Close() }()

		for stream.Receive() {
			event := stream.Msg()
			if event.GetResourceType() != v1.ResourceType_RESOURCE_TYPE_AGENT_EXECUTION || event.GetResourceId() != s.id {
				continue
			}
			s.client.sdk.logger().DebugContext(watchCtx, "agent execution event received",
				"operation", operation,
				"agent_execution_id", s.id,
				"resource_operation", event.GetOperation().String(),
			)
			select {
			case changes <- struct{}{}:
			default:
			}
		}
		if err := stream.Err(); err != nil && !errors.Is(err, context.Canceled) {
			s.client.sdk.logger().ErrorContext(watchCtx, "agent execution informer event invalidation failed", "agent_execution_id", s.id, "err", err)
			watchDone <- err
			return
		}
		s.client.sdk.logger().DebugContext(watchCtx, "agent execution event watch ended",
			"operation", operation,
			"agent_execution_id", s.id,
		)
		watchDone <- nil
	}()

	select {
	case <-ctx.Done():
		return nil, MapError(operation, ctx.Err())
	case err := <-watchStarted:
		if err != nil {
			return nil, MapError(operation, err)
		}
	}

	last, err := s.observeAgentExecution(ctx, operation, informer, onUpdate)
	if err != nil {
		return last, err
	}
	if done, err := condition(last); err != nil || done {
		return last, err
	}
	informer.InvalidateItem(ctx, s.id)
	last, err = s.observeAgentExecution(ctx, operation, informer, onUpdate)
	if err != nil {
		return last, err
	}
	if done, err := condition(last); err != nil || done {
		return last, err
	}

	for {
		select {
		case <-ctx.Done():
			return last, MapError(operation, ctx.Err())
		case <-changes:
			informer.InvalidateItem(ctx, s.id)
			last, err = s.observeAgentExecution(ctx, operation, informer, onUpdate)
			if err != nil {
				return last, err
			}
			if done, err := condition(last); err != nil || done {
				return last, err
			}
		case watchErr := <-watchDone:
			for {
				select {
				case <-changes:
					informer.InvalidateItem(ctx, s.id)
					last, err = s.observeAgentExecution(ctx, operation, informer, onUpdate)
					if err != nil {
						return last, err
					}
					if done, err := condition(last); err != nil || done {
						return last, err
					}
				default:
					if watchErr != nil {
						return last, MapError(operation, watchErr)
					}
					return last, &UnavailableError{operationError{
						operation: operation,
						err:       fmt.Errorf("agent execution %s event stream ended before completion", s.id),
					}}
				}
			}
		}
	}
}

func (s *AgentSession) observeAgentExecution(ctx context.Context, operation string, informer *cache.ResourceCache[*v1.AgentExecution], onUpdate AgentExecutionUpdateFunc) (*v1.AgentExecution, error) {
	execution, err := informer.Get(ctx, s.id)
	if err != nil {
		return nil, MapError(operation, err)
	}

	previous := s.execution
	s.execution = execution
	if previous == nil || !proto.Equal(previous, execution) {
		s.client.sdk.logger().LogAttrs(ctx, slog.LevelDebug, "agent execution updated", agentExecutionLogAttrs(operation, execution)...)
	}
	if onUpdate != nil && (previous == nil || !proto.Equal(previous, execution)) {
		if err := onUpdate(ctx, execution); err != nil {
			return execution, &SDKError{operationError{operation: operation, err: fmt.Errorf("update callback: %w", err)}}
		}
	}
	if err := agentFailureError(operation, execution); err != nil {
		return execution, err
	}
	return execution, nil
}

func agentExecutionLogAttrs(operation string, execution *v1.AgentExecution) []slog.Attr {
	status := execution.GetStatus()
	attrs := []slog.Attr{
		slog.String("operation", operation),
		slog.String("agent_execution_id", execution.GetId()),
		slog.String("status", AgentStatusLine(execution)),
		slog.String("phase", status.GetPhase().String()),
		slog.String("activity", status.GetCurrentActivity()),
		slog.String("warning", status.GetWarningMessage()),
	}
	if failureReason := status.GetFailureReason(); failureReason != v1.AgentExecutionFailureReason_AGENT_EXECUTION_FAILURE_REASON_UNSPECIFIED {
		attrs = append(attrs, slog.String("failure_reason", failureReason.String()))
	}
	attrs = append(attrs,
		slog.String("failure_message", status.GetFailureMessage()),
		slog.String("support_bundle_url", safeURLForLog(status.GetSupportBundleUrl())),
	)
	return attrs
}

func (s *AgentSession) sendText(ctx context.Context, text string, inputID string) error {
	operation := "agents.send_message"
	if text == "" {
		return &ValidationError{operationError{operation: operation, err: errors.New("text is required")}}
	}
	if inputID == "" {
		inputID = uuid.NewString()
	}

	raw, err := s.raw(operation)
	if err != nil {
		return err
	}
	ctx = s.client.sdk.requestContext(ctx)
	if raw.AgentService() == nil {
		return &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("agent service is not configured")}}
	}
	s.client.sdk.logger().DebugContext(ctx, "sending agent text input",
		"operation", operation,
		"agent_execution_id", s.id,
		"input_id", inputID,
		"text_bytes", len(text),
	)
	_, err = raw.AgentService().SendToAgentExecution(ctx, connect.NewRequest(&v1.SendToAgentExecutionRequest{
		AgentExecutionId: s.id,
		Input: &v1.SendToAgentExecutionRequest_UserInput{
			UserInput: &v1.UserInputBlock{
				Id: inputID,
				Inputs: []*v1.UserInputBlock_Input{{
					Input: &v1.UserInputBlock_Input_Text{
						Text: &v1.UserInputBlock_TextInput{Content: text},
					},
				}},
			},
		},
	}))
	if err != nil {
		return MapError(operation, err)
	}
	s.client.sdk.logger().DebugContext(ctx, "agent text input sent",
		"operation", operation,
		"agent_execution_id", s.id,
		"input_id", inputID,
	)
	return nil
}

func (c *agentClient) raw(operation string) (*rawclient.ManagementPlane, error) {
	if c == nil || c.sdk == nil {
		return nil, &ValidationError{operationError{operation: operation, err: errSDKClientMissing}}
	}
	return c.sdk.requireRaw(operation)
}

func (s *AgentSession) raw(operation string) (*rawclient.ManagementPlane, error) {
	if s == nil || s.client == nil {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("agent session is not initialized")}}
	}
	if s.id == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("agent execution ID is required")}}
	}
	return s.client.raw(operation)
}

func buildCodexCodeContext(opts startCodexOptions) (*v1.AgentCodeContext, error) {
	if opts.EnvironmentID == "" {
		return nil, errors.New("environment ID is required")
	}
	return &v1.AgentCodeContext{
		Context: &v1.AgentCodeContext_EnvironmentId{EnvironmentId: opts.EnvironmentID},
	}, nil
}

func agentFailureError(operation string, exec *v1.AgentExecution) error {
	status := exec.GetStatus()
	if status.GetFailureMessage() == "" {
		return nil
	}
	return &UnavailableError{operationError{
		operation: operation,
		err:       fmt.Errorf("agent execution %s failed: %s; last status: %s", exec.GetId(), status.GetFailureMessage(), AgentStatusLine(exec)),
	}}
}

func agentExecutionWatchRequest(agentExecutionID string) *v1.WatchEventsRequest {
	return &v1.WatchEventsRequest{
		Scope: &v1.WatchEventsRequest_Organization{Organization: true},
		ResourceTypeFilters: []*v1.WatchEventsRequest_ResourceTypeFilter{{
			ResourceType: v1.ResourceType_RESOURCE_TYPE_AGENT_EXECUTION,
			ResourceIds:  []string{agentExecutionID},
		}},
	}
}

func agentPhaseLabel(phase v1.AgentExecution_Phase) string {
	switch phase {
	case v1.AgentExecution_PHASE_PENDING:
		return "pending"
	case v1.AgentExecution_PHASE_RUNNING:
		return "running"
	case v1.AgentExecution_PHASE_WAITING_FOR_INPUT:
		return "waiting for input"
	case v1.AgentExecution_PHASE_STOPPED:
		return "stopped"
	default:
		return phase.String()
	}
}

func agentCurrentOperation(status *v1.AgentExecution_Status) string {
	current := status.GetCurrentOperation()
	switch op := current.GetOperation().(type) {
	case *v1.AgentExecution_Status_CurrentOperation_ToolUse:
		if op.ToolUse.GetToolName() == "" {
			return ""
		}
		if op.ToolUse.GetComplete() {
			return fmt.Sprintf("completed tool %s", op.ToolUse.GetToolName())
		}
		return fmt.Sprintf("using tool %s", op.ToolUse.GetToolName())
	case *v1.AgentExecution_Status_CurrentOperation_Llm:
		if op.Llm.GetComplete() {
			return "completed model request"
		}
		return "calling model"
	default:
		return ""
	}
}

// AgentStatusLine returns a compact progress string suitable for logs and examples.
func AgentStatusLine(exec *v1.AgentExecution) string {
	status := exec.GetStatus()
	line := fmt.Sprintf("phase=%s", agentPhaseLabel(status.GetPhase()))
	if activity := status.GetCurrentActivity(); activity != "" {
		line += fmt.Sprintf(" activity=%q", activity)
	}
	if operation := agentCurrentOperation(status); operation != "" {
		line += fmt.Sprintf(" operation=%q", operation)
	}
	if warning := status.GetWarningMessage(); warning != "" {
		line += fmt.Sprintf(" warning=%q", warning)
	}
	if failureReason := status.GetFailureReason(); failureReason != v1.AgentExecutionFailureReason_AGENT_EXECUTION_FAILURE_REASON_UNSPECIFIED {
		line += fmt.Sprintf(" failure_reason=%s", failureReason.String())
	}
	if failure := status.GetFailureMessage(); failure != "" {
		line += fmt.Sprintf(" failure=%q", failure)
	}
	if supportBundleURL := status.GetSupportBundleUrl(); supportBundleURL != "" {
		line += fmt.Sprintf(" support_bundle=%s", safeURLForLog(supportBundleURL))
	}
	return line
}
