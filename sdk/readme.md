# Ona Go SDK

The `sdk` package is the task-oriented layer on top of the raw Ona API client. Use it for production-ready workflows that should be easy to compose, such as creating an environment, running a command, reading files, and starting Codex inside an environment.

Use the raw `client` package when you need direct access to backend RPCs that are not part of the stable SDK surface.

For the common Codex workflow, use `RunCodex` to create a repository-backed or scratch environment and start the initial task in one call:

```go
run, err := ona.RunCodex(ctx, sdk.RunCodexOptions{
	RepositoryURL:   "https://github.com/gitpod-io/template-golang-cli",
	Task:            "Inspect the repository and improve its README.",
	Model:           v1.CodexOpenAIModel_CODEX_OPEN_AI_MODEL_GPT_5_6_SOL,
	ReasoningEffort: v1.CodexReasoningEffort_CODEX_REASONING_EFFORT_HIGH,
})
if err != nil {
	return err
}
defer run.DeleteEnvironment(context.WithoutCancel(ctx), sdk.DeleteEnvironmentOptions{Force: true})
```

Omit `RepositoryURL` to start from an empty workspace. The returned run exposes the environment, agent session, live-only Markdown stream, follow-up messages, and result watcher. The environment is caller-owned and is never deleted automatically.

Install the module with `go get github.com/gitpod-io/gitpod-sdk-go` and import `github.com/gitpod-io/gitpod-sdk-go/sdk`.

## Structure

The SDK is organized around resource clients and resource handles:

```text
sdk.Client
`-- Environments()
    |-- Create(ctx, CreateEnvironmentOptions) -> *Environment
    |-- Get(ctx, environmentID) -> *Environment
    |-- List(ctx) -> iter.Seq2[*Environment, error]
    |-- Stop(ctx, environmentID)
    `-- Delete(ctx, environmentID, DeleteEnvironmentOptions)

Environment
|-- ID()
|-- Proto()
|-- RunCommand(ctx, RunCommandOptions)
|-- ReadFile(ctx, path, ReadFileOptions)
|-- WriteFile(ctx, path, content, WriteFileOptions)
|-- GitChanges(ctx, GitChangesOptions)
`-- StartCodex(ctx, EnvironmentCodexOptions) -> *AgentSession

Client
`-- RunCodex(ctx, RunCodexOptions) -> *CodexRun

AgentSession
|-- ID()
|-- SendMessage(ctx, text)
|-- MessageStream(ctx) -> io.ReadCloser
`-- WatchResult(ctx, onUpdate)
```

The usual flow is:

1. Create an SDK client with `sdk.NewFromEnv`.
2. Use `ona.Environments()` to create, get, list, stop, or delete environments.
3. Use an `Environment` handle to run commands, access files, or start Codex.
4. Clean up environments with `Environments().Delete`.

The package files follow the same structure:

- `sdk.go`: client construction, options, and top-level resource clients.
- `environments.go`: environment create, get, list, stop, delete, and environment handles.
- `ops.go`: supervisor-backed environment operations such as command execution, file access, and git changes.
- `agents.go`: Codex startup, agent sessions, and agent result watching.
- `errors.go`: typed SDK errors.
- `logging.go`: human-readable `slog` helpers.

## Create a client

`NewFromEnv` reads `ONA_API_KEY` and falls back to `GITPOD_API_KEY`. If both are set, `ONA_API_KEY` takes precedence. It uses `https://app.ona.com/api` by default. Set `ONA_BASE_URL` or pass `WithBaseURL` for a custom management-plane domain, local development, or replay verification.

```go
ctx := context.Background()

ona, err := sdk.NewFromEnv()
if err != nil {
    return err
}

envs := ona.Environments()
```

Use `sdk.New(rawClient)` when you already have a configured raw management-plane client.

```go
ona, err := sdk.NewFromEnv(
    sdk.WithBaseURL("http://127.0.0.1:18080"),
)
if err != nil {
    return err
}
```

## Enable readable SDK logs

Use `WithLogger` to see resource creation, state changes, agent updates, and supervisor operations. `NewDebugLogger` returns a human-readable `slog` logger for examples and command-line tools.

```go
ona, err := sdk.NewFromEnv(
    sdk.WithLogger(sdk.NewDebugLogger(os.Stderr)),
)
if err != nil {
    return err
}
```

## Create or get an environment

`Create` creates an environment from a context URL and waits until it is running. The SDK resolves the context URL, prefers matching projects when available, selects an environment class, checks SCM authentication, and starts the environment.

```go
env, err := ona.Environments().Create(ctx, sdk.CreateEnvironmentOptions{
    ContextURL: "https://github.com/gitpod-io/template-golang-cli",
    Name:       "sdk-example",
})
if err != nil {
    return err
}
defer ona.Environments().Delete(ctx, env.ID(), sdk.DeleteEnvironmentOptions{Force: true})
```

Use `Get` when you already have an environment ID:

```go
env, err := ona.Environments().Get(ctx, "019f0000-0000-7000-8000-000000000000")
if err != nil {
    return err
}
```

## List environments lazily

`List` returns an `iter.Seq2[*Environment, error]`. The SDK does not call the API until the sequence is iterated, and it fetches the next page only when needed. The list is scoped to default environments owned by the authenticated caller.

```go
for env, err := range ona.Environments().List(ctx) {
    if err != nil {
        return err
    }

    fmt.Println(env.ID())
}
```

Break early to stop pagination:

```go
for env, err := range ona.Environments().List(ctx) {
    if err != nil {
        return err
    }

    fmt.Println(env.ID())
    break
}
```

## Run commands and access files

Use `Environment.RunCommand` to run a command in an environment. The SDK creates and caches the internal supervisor connection when the first environment operation runs.

```go
result, err := env.RunCommand(ctx, sdk.RunCommandOptions{
    Command:          "pwd && ls -la",
    WorkingDirectory: "/workspace",
    TimeoutSeconds:   60,
})
if err != nil {
    return err
}

fmt.Println(result.Stdout)
fmt.Fprintln(os.Stderr, result.Stderr)
```

Use the file and git methods on the same `Environment` handle:

```go
file, err := env.ReadFile(ctx, "/workspace/README.md", sdk.ReadFileOptions{})
if err != nil {
    return err
}

changes, err := env.GitChanges(ctx, sdk.GitChangesOptions{Unified: 3})
if err != nil {
    return err
}

_, err = env.WriteFile(ctx, "/workspace/ona-sdk-example.txt", []byte("hello\n"), sdk.WriteFileOptions{
    Mode: sdk.WriteFileModeCreateOrTruncate,
})
if err != nil {
    return err
}

_, _ = file, changes
```

## Start Codex inside an environment

`StartCodex` starts the Codex agent in an existing environment, sends the required initial prompt, and waits until the agent is running. `MessageStream` returns a live markdown stream for new conversation messages only; it does not fetch history. Use `SendMessage` for follow-up messages after the session starts.

```go
session, err := env.StartCodex(ctx, sdk.EnvironmentCodexOptions{
    Name:   "review-readme",
    Prompt: "Inspect the repository and summarize what the README says.",
})
if err != nil {
    return err
}

stream, err := session.MessageStream(ctx)
if err != nil {
    return err
}
defer stream.Close()

go func() {
    _, _ = io.Copy(os.Stdout, stream)
}()

if err := session.SendMessage(ctx, "Now check whether the project has tests."); err != nil {
    return err
}

execution, err := session.WatchResult(ctx, func(ctx context.Context, update *v1.AgentExecution) error {
    fmt.Println(update.GetStatus().GetPhase())
    return nil
})
if err != nil {
    return err
}

fmt.Println(execution.GetStatus().GetPhase())
```

## Error handling

SDK methods wrap common API failures in typed errors:

- `AuthenticationRequiredError`
- `PermissionDeniedError`
- `NotFoundError`
- `RateLimitedError`
- `UnavailableError`
- `DeadlineExceededError`
- `ValidationError`
- `CapabilityUnavailableError`
- `EnvironmentPolicyError`
- `EnvironmentUnreachableError`

Use `errors.As` when callers need to branch on an expected failure:

```go
env, err := ona.Environments().Create(ctx, sdk.CreateEnvironmentOptions{
    ContextURL: "https://github.com/private/repo",
})
if err != nil {
    var authErr *sdk.AuthenticationRequiredError
    if errors.As(err, &authErr) {
        return fmt.Errorf("authenticate Git access in Ona settings: %w", err)
    }
    return err
}

_ = env
```

`EnvironmentUnreachableError` means the SDK could not reach the environment supervisor ops service. This usually means the SDK is running somewhere that cannot resolve or connect to the environment ops URL, or the connection timed out.

## Examples

Runnable examples live under `examples/`:

- `examples/start_environment_and_run_command`
- `examples/run_codex_agent_in_environment`

Run them with `go run` from the example directory after setting `ONA_API_KEY`.
