package sdk

import (
	"context"
	"errors"
	"io"
	"strings"

	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
)

// RunCodexOptions controls a new environment and Codex task.
type RunCodexOptions struct {
	RepositoryURL   string
	Task            string
	EnvironmentName string
	AgentName       string
	Model           v1.CodexOpenAIModel
	ReasoningEffort v1.CodexReasoningEffort
}

// CodexRun groups the environment and agent session created for a Codex task.
type CodexRun struct {
	client      *Client
	environment *Environment
	session     *AgentSession
}

// RunCodex creates an environment and starts Codex with the supplied task.
func (s *Client) RunCodex(ctx context.Context, opts RunCodexOptions) (*CodexRun, error) {
	operation := "codex.run"
	if s == nil {
		return nil, &ValidationError{operationError{operation: operation, err: errSDKClientMissing}}
	}
	if strings.TrimSpace(opts.Task) == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("task is required")}}
	}
	if err := validateCodexSettings(opts.Model, opts.ReasoningEffort); err != nil {
		return nil, &ValidationError{operationError{operation: operation, err: err}}
	}

	var env *Environment
	var err error
	if strings.TrimSpace(opts.RepositoryURL) == "" {
		env, err = s.Environments().createScratch(ctx, opts.EnvironmentName)
	} else {
		env, err = s.Environments().Create(ctx, CreateEnvironmentOptions{
			ContextURL: opts.RepositoryURL,
			Name:       opts.EnvironmentName,
		})
	}
	if err != nil {
		return nil, err
	}

	run := &CodexRun{client: s, environment: env}
	session, err := env.StartCodex(ctx, EnvironmentCodexOptions{
		Prompt:          opts.Task,
		Name:            opts.AgentName,
		Model:           opts.Model,
		ReasoningEffort: opts.ReasoningEffort,
	})
	run.session = session
	if err != nil {
		return run, err
	}
	return run, nil
}

// Environment returns the environment created for the run.
func (r *CodexRun) Environment() *Environment {
	if r == nil {
		return nil
	}
	return r.environment
}

// Session returns the Codex agent session.
func (r *CodexRun) Session() *AgentSession {
	if r == nil {
		return nil
	}
	return r.session
}

// EnvironmentID returns the created environment ID.
func (r *CodexRun) EnvironmentID() string {
	if r == nil || r.environment == nil {
		return ""
	}
	return r.environment.ID()
}

// ID returns the agent execution ID.
func (r *CodexRun) ID() string {
	if r == nil || r.session == nil {
		return ""
	}
	return r.session.ID()
}

// SendMessage sends follow-up text to Codex.
func (r *CodexRun) SendMessage(ctx context.Context, text string) error {
	if r == nil || r.session == nil {
		return &ValidationError{operationError{operation: "codex.run.send_message", err: errors.New("codex session is not initialized")}}
	}
	return r.session.SendMessage(ctx, text)
}

// MessageStream opens the existing live-only Markdown stream.
func (r *CodexRun) MessageStream(ctx context.Context) (io.ReadCloser, error) {
	if r == nil || r.session == nil {
		return nil, &ValidationError{operationError{operation: "codex.run.message_stream", err: errors.New("codex session is not initialized")}}
	}
	return r.session.MessageStream(ctx)
}

// WatchResult waits for Codex to stop or request input.
func (r *CodexRun) WatchResult(ctx context.Context, onUpdate AgentExecutionUpdateFunc) (*v1.AgentExecution, error) {
	if r == nil || r.session == nil {
		return nil, &ValidationError{operationError{operation: "codex.run.watch_result", err: errors.New("codex session is not initialized")}}
	}
	return r.session.WatchResult(ctx, onUpdate)
}

// StopEnvironment stops the environment created for the run.
func (r *CodexRun) StopEnvironment(ctx context.Context) error {
	if r == nil || r.client == nil || r.environment == nil {
		return &ValidationError{operationError{operation: "codex.run.stop_environment", err: errors.New("environment is not initialized")}}
	}
	return r.client.Environments().Stop(ctx, r.environment.ID())
}

// DeleteEnvironment deletes the environment created for the run.
func (r *CodexRun) DeleteEnvironment(ctx context.Context, opts DeleteEnvironmentOptions) error {
	if r == nil || r.client == nil || r.environment == nil {
		return &ValidationError{operationError{operation: "codex.run.delete_environment", err: errors.New("environment is not initialized")}}
	}
	return r.client.Environments().Delete(ctx, r.environment.ID(), opts)
}
