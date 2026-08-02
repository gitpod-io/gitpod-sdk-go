package sdk

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"strings"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"golang.org/x/oauth2"

	rawclient "github.com/gitpod-io/gitpod-sdk-go/internal/raw"
	supervisorv1 "github.com/gitpod-io/gitpod-sdk-go/supervisor/v1"
	supervisorconnect "github.com/gitpod-io/gitpod-sdk-go/supervisor/v1/v1connect"
	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
)

type opsClient struct {
	sdk *Client
}

type environmentOps struct {
	raw           supervisorconnect.EnvironmentOpsServiceClient
	logger        *slog.Logger
	environmentID string
	opsURL        string
}

// ReadFileOptions controls an environment file read.
type ReadFileOptions struct {
	Offset int64
	Length int64
}

// WriteFileOptions controls an environment file write.
type WriteFileOptions struct {
	Mode WriteFileMode
}

// WriteFileMode controls how WriteFile handles an existing file.
type WriteFileMode int32

const (
	// WriteFileModeUnspecified uses the supervisor default.
	WriteFileModeUnspecified WriteFileMode = WriteFileMode(supervisorv1.WriteMode_WRITE_MODE_UNSPECIFIED)
	// WriteFileModeCreate creates a file and fails if it already exists.
	WriteFileModeCreate WriteFileMode = WriteFileMode(supervisorv1.WriteMode_WRITE_MODE_CREATE)
	// WriteFileModeTruncate overwrites an existing file and fails if it does not exist.
	WriteFileModeTruncate WriteFileMode = WriteFileMode(supervisorv1.WriteMode_WRITE_MODE_TRUNCATE)
	// WriteFileModeCreateOrTruncate creates or overwrites a file.
	WriteFileModeCreateOrTruncate WriteFileMode = WriteFileMode(supervisorv1.WriteMode_WRITE_MODE_CREATE_OR_TRUNCATE)
	// WriteFileModeAppend appends to a file and fails if it does not exist.
	WriteFileModeAppend WriteFileMode = WriteFileMode(supervisorv1.WriteMode_WRITE_MODE_APPEND)
)

// RunCommandOptions controls command execution inside an environment.
type RunCommandOptions struct {
	Command          string
	WorkingDirectory string
	TimeoutSeconds   int32
}

// GitChangesOptions controls environment git review.
type GitChangesOptions struct {
	Unified int32
	BaseRef string
}

// EnvironmentGitChanges contains git status plus per-file diffs.
type EnvironmentGitChanges struct {
	Status *v1.EnvironmentGitStatus
	Files  []*EnvironmentGitFileChange
}

// EnvironmentGitFileChange contains a changed file and its diff.
type EnvironmentGitFileChange struct {
	File *v1.FileChange
	Diff *supervisorv1.GetGitDiffResponse
}

// RunCommandResult includes the environment ID with the command result.
type RunCommandResult struct {
	EnvironmentID string
	ExitCode      int32
	Stdout        string
	Stderr        string
}

// RunCommand runs a command in the environment.
func (e *Environment) RunCommand(ctx context.Context, opts RunCommandOptions) (*RunCommandResult, error) {
	if e == nil || e.sdk == nil || e.environment == nil {
		return nil, &ValidationError{operationError{operation: "environments.run_command", err: errors.New("environment is not initialized")}}
	}
	ctx = e.sdk.requestContext(ctx)
	ops, err := e.opsClient(ctx)
	if err != nil {
		return nil, err
	}
	e.sdk.logger().DebugContext(ctx, "running command in environment",
		"operation", "environments.run_command",
		"environment_id", e.ID(),
		"working_directory", opts.WorkingDirectory,
		"timeout_seconds", opts.TimeoutSeconds,
		"command_bytes", len(opts.Command),
	)
	resp, err := ops.runCommand(ctx, opts)
	if err != nil {
		return nil, err
	}
	e.sdk.logger().DebugContext(ctx, "environment command finished",
		"operation", "environments.run_command",
		"environment_id", e.ID(),
		"exit_code", resp.GetExitCode(),
		"stdout_bytes", len(resp.GetStdout()),
		"stderr_bytes", len(resp.GetStderr()),
	)
	return &RunCommandResult{
		EnvironmentID: e.ID(),
		ExitCode:      resp.GetExitCode(),
		Stdout:        resp.GetStdout(),
		Stderr:        resp.GetStderr(),
	}, nil
}

// ReadFile reads a file or directory from the environment.
func (e *Environment) ReadFile(ctx context.Context, path string, opts ReadFileOptions) (*supervisorv1.ReadFileResponse, error) {
	if e == nil || e.sdk == nil || e.environment == nil {
		return nil, &ValidationError{operationError{operation: "environments.read_file", err: errors.New("environment is not initialized")}}
	}
	ctx = e.sdk.requestContext(ctx)
	ops, err := e.opsClient(ctx)
	if err != nil {
		return nil, err
	}
	return ops.readFile(ctx, path, opts)
}

// WriteFile writes bytes to an environment file.
func (e *Environment) WriteFile(ctx context.Context, path string, content []byte, opts WriteFileOptions) (*supervisorv1.WriteFileResponse, error) {
	if e == nil || e.sdk == nil || e.environment == nil {
		return nil, &ValidationError{operationError{operation: "environments.write_file", err: errors.New("environment is not initialized")}}
	}
	ctx = e.sdk.requestContext(ctx)
	ops, err := e.opsClient(ctx)
	if err != nil {
		return nil, err
	}
	return ops.writeFile(ctx, path, content, opts)
}

// GitChanges returns the current git status and diffs for each changed file.
func (e *Environment) GitChanges(ctx context.Context, opts GitChangesOptions) (*EnvironmentGitChanges, error) {
	if e == nil || e.sdk == nil || e.environment == nil {
		return nil, &ValidationError{operationError{operation: "environments.git_changes", err: errors.New("environment is not initialized")}}
	}
	ctx = e.sdk.requestContext(ctx)
	ops, err := e.opsClient(ctx)
	if err != nil {
		return nil, err
	}
	return ops.gitChanges(ctx, opts)
}

func (e *Environment) opsClient(ctx context.Context) (*environmentOps, error) {
	if e.ops != nil {
		e.sdk.logger().DebugContext(ctx, "using cached environment ops client", "operation", "environments.ops", "environment_id", e.environment.idString())
		return e.ops, nil
	}
	ops, err := e.sdk.ops().forEnvironment(ctx, e.environment.idString())
	if err != nil {
		return nil, err
	}
	e.ops = ops
	e.sdk.logger().DebugContext(ctx, "environment ops client ready", "operation", "environments.ops", "environment_id", e.environment.idString())
	return ops, nil
}

func (c *opsClient) forEnvironment(ctx context.Context, environmentID string) (*environmentOps, error) {
	operation := "ops.for_environment"
	raw, err := c.raw(operation)
	if err != nil {
		return nil, err
	}
	ctx = c.sdk.requestContext(ctx)
	if environmentID == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("environment ID is required")}}
	}

	c.sdk.logger().DebugContext(ctx, "connecting to environment ops",
		"operation", operation,
		"environment_id", environmentID,
	)
	envResp, err := raw.EnvironmentService().GetEnvironment(ctx, connect.NewRequest(&v1.GetEnvironmentRequest{
		EnvironmentId: environmentID,
	}))
	if err != nil {
		return nil, MapError(operation, err)
	}
	env := envResp.Msg.GetEnvironment()
	opsURL := env.GetStatus().GetEnvironmentUrls().GetOps()
	if opsURL == "" {
		return nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("environment ops URL is not available")}}
	}

	c.sdk.logger().DebugContext(ctx, "creating environment access token",
		"operation", operation,
		"environment_id", environmentID,
	)
	tokenResp, err := raw.EnvironmentService().CreateEnvironmentAccessToken(ctx, connect.NewRequest(&v1.CreateEnvironmentAccessTokenRequest{
		EnvironmentId: environmentID,
	}))
	if err != nil {
		return nil, MapError(operation, err)
	}
	httpClient := c.sdk.config().httpClient
	opsClient := supervisorconnect.NewEnvironmentOpsServiceClient(httpClient, opsURL, connect.WithInterceptors(
		rawclient.TokenSourceInterceptor(oauth2.StaticTokenSource(&oauth2.Token{AccessToken: tokenResp.Msg.GetAccessToken()})),
		rawclient.WithCustomUserAgent(rawclient.SDKUserAgent()),
	))

	return &environmentOps{
		raw:           opsClient,
		logger:        c.sdk.logger(),
		environmentID: environmentID,
		opsURL:        opsURL,
	}, nil
}

func (o *environmentOps) gitChanges(ctx context.Context, opts GitChangesOptions) (*EnvironmentGitChanges, error) {
	o.loggerOrDefault().DebugContext(ctx, "reading environment git changes",
		"operation", "ops.git_changes",
		"environment_id", o.environmentID,
		"base_ref", opts.BaseRef,
		"unified", opts.Unified,
	)
	status, err := o.gitStatus(ctx)
	if err != nil {
		return nil, err
	}

	changes := &EnvironmentGitChanges{Status: status}
	for _, changed := range status.GetChangedFiles() {
		diff, err := o.gitDiff(ctx, changed.GetPath(), opts)
		if err != nil {
			return changes, err
		}
		changes.Files = append(changes.Files, &EnvironmentGitFileChange{
			File: changed,
			Diff: diff,
		})
	}
	o.loggerOrDefault().DebugContext(ctx, "read environment git changes",
		"operation", "ops.git_changes",
		"environment_id", o.environmentID,
		"changed_files", len(changes.Files),
	)
	return changes, nil
}

func (o *environmentOps) readFile(ctx context.Context, path string, opts ReadFileOptions) (*supervisorv1.ReadFileResponse, error) {
	operation := "ops.read_file"
	raw, err := o.requireRaw(operation)
	if err != nil {
		return nil, err
	}
	if path == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("path is required")}}
	}
	o.loggerOrDefault().DebugContext(ctx, "reading environment file",
		"operation", operation,
		"environment_id", o.environmentID,
		"path", path,
		"offset", opts.Offset,
		"length", opts.Length,
	)
	resp, err := raw.ReadFile(ctx, connect.NewRequest(&supervisorv1.ReadFileRequest{
		Path:   path,
		Offset: opts.Offset,
		Length: opts.Length,
	}))
	if err != nil {
		return nil, o.mapError(operation, err)
	}
	resultType := "unknown"
	contentBytes := 0
	var totalSize int64
	directoryEntries := 0
	if content := resp.Msg.GetContent(); content != nil {
		resultType = "file"
		contentBytes = len(content.GetData())
		totalSize = content.GetTotalSize()
	}
	if directory := resp.Msg.GetDirectory(); directory != nil {
		resultType = "directory"
		directoryEntries = len(directory.GetEntries())
	}
	o.loggerOrDefault().DebugContext(ctx, "read environment file",
		"operation", operation,
		"environment_id", o.environmentID,
		"path", path,
		"result_type", resultType,
		"bytes", contentBytes,
		"total_size", totalSize,
		"directory_entries", directoryEntries,
	)
	return resp.Msg, nil
}

func (o *environmentOps) writeFile(ctx context.Context, path string, content []byte, opts WriteFileOptions) (*supervisorv1.WriteFileResponse, error) {
	operation := "ops.write_file"
	raw, err := o.requireRaw(operation)
	if err != nil {
		return nil, err
	}
	if path == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("path is required")}}
	}
	o.loggerOrDefault().DebugContext(ctx, "writing environment file",
		"operation", operation,
		"environment_id", o.environmentID,
		"path", path,
		"bytes", len(content),
		"mode", supervisorv1.WriteMode(opts.Mode).String(),
	)
	resp, err := raw.WriteFile(ctx, connect.NewRequest(&supervisorv1.WriteFileRequest{
		Path:    path,
		Content: content,
		Mode:    supervisorv1.WriteMode(opts.Mode),
	}))
	if err != nil {
		return nil, o.mapError(operation, err)
	}
	o.loggerOrDefault().DebugContext(ctx, "wrote environment file",
		"operation", operation,
		"environment_id", o.environmentID,
		"path", path,
	)
	return resp.Msg, nil
}

func (o *environmentOps) runCommand(ctx context.Context, opts RunCommandOptions) (*supervisorv1.ExecResponse, error) {
	operation := "ops.run_command"
	raw, err := o.requireRaw(operation)
	if err != nil {
		return nil, err
	}
	if opts.Command == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("command is required")}}
	}
	operationID := uuid.NewString()
	o.loggerOrDefault().DebugContext(ctx, "executing environment command",
		"operation", operation,
		"environment_id", o.environmentID,
		"operation_id", operationID,
		"working_directory", opts.WorkingDirectory,
		"timeout_seconds", opts.TimeoutSeconds,
		"command_bytes", len(opts.Command),
	)
	resp, err := raw.Exec(ctx, connect.NewRequest(&supervisorv1.ExecRequest{
		OperationId:      operationID,
		Command:          opts.Command,
		WorkingDirectory: opts.WorkingDirectory,
		TimeoutSeconds:   opts.TimeoutSeconds,
	}))
	if err != nil {
		return nil, o.mapError(operation, err)
	}
	o.loggerOrDefault().DebugContext(ctx, "executed environment command",
		"operation", operation,
		"environment_id", o.environmentID,
		"operation_id", operationID,
		"exit_code", resp.Msg.GetExitCode(),
		"stdout_bytes", len(resp.Msg.GetStdout()),
		"stderr_bytes", len(resp.Msg.GetStderr()),
	)
	return resp.Msg, nil
}

func (o *environmentOps) gitStatus(ctx context.Context) (*v1.EnvironmentGitStatus, error) {
	operation := "ops.git_status"
	raw, err := o.requireRaw(operation)
	if err != nil {
		return nil, err
	}
	resp, err := raw.GetGitStatus(ctx, connect.NewRequest(&supervisorv1.GetGitStatusRequest{}))
	if err != nil {
		return nil, o.mapError(operation, err)
	}
	o.loggerOrDefault().DebugContext(ctx, "read environment git status",
		"operation", operation,
		"environment_id", o.environmentID,
		"changed_files", len(resp.Msg.GetStatus().GetChangedFiles()),
	)
	return resp.Msg.GetStatus(), nil
}

func (o *environmentOps) gitDiff(ctx context.Context, path string, opts GitChangesOptions) (*supervisorv1.GetGitDiffResponse, error) {
	operation := "ops.git_diff"
	raw, err := o.requireRaw(operation)
	if err != nil {
		return nil, err
	}
	resp, err := raw.GetGitDiff(ctx, connect.NewRequest(&supervisorv1.GetGitDiffRequest{
		Path:    path,
		Unified: opts.Unified,
		BaseRef: opts.BaseRef,
	}))
	if err != nil {
		return nil, o.mapError(operation, err)
	}
	o.loggerOrDefault().DebugContext(ctx, "read environment git diff",
		"operation", operation,
		"environment_id", o.environmentID,
		"path", path,
		"base_ref", opts.BaseRef,
		"unified", opts.Unified,
	)
	return resp.Msg, nil
}

func (o *environmentOps) requireRaw(operation string) (supervisorconnect.EnvironmentOpsServiceClient, error) {
	if o == nil || o.raw == nil {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("environment ops client is not initialized")}}
	}
	return o.raw, nil
}

func (o *environmentOps) mapError(operation string, err error) error {
	if err == nil {
		return nil
	}
	if isEnvironmentOpsReachabilityError(err) {
		return &EnvironmentUnreachableError{operationError{operation: operation, err: fmt.Errorf("cannot reach environment %s ops service at %s: %w", o.environmentID, o.opsURL, err)}}
	}
	return MapError(operation, err)
}

func isEnvironmentOpsReachabilityError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return true
	}
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		if urlErr.Timeout() {
			return true
		}
	}
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return true
	}

	switch connect.CodeOf(err) {
	case connect.CodeUnavailable, connect.CodeDeadlineExceeded:
		return true
	}

	message := strings.ToLower(err.Error())
	return strings.Contains(message, "no such host") ||
		strings.Contains(message, "connection refused") ||
		strings.Contains(message, "network is unreachable") ||
		strings.Contains(message, "i/o timeout")
}

func (o *environmentOps) loggerOrDefault() *slog.Logger {
	if o == nil || o.logger == nil {
		return slog.Default()
	}
	return o.logger
}

func (c *opsClient) raw(operation string) (*rawclient.ManagementPlane, error) {
	if c == nil || c.sdk == nil {
		return nil, &ValidationError{operationError{operation: operation, err: errSDKClientMissing}}
	}
	return c.sdk.requireRaw(operation)
}
