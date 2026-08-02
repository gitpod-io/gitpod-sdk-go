# Ona Go SDK

Install the module with:

```bash
go get github.com/gitpod-io/gitpod-sdk-go
```

Use `github.com/gitpod-io/gitpod-sdk-go/sdk` for task-oriented environment and Codex workflows. The same client exposes every public API service through `Client.Services`, with generated protobuf types in `github.com/gitpod-io/gitpod-sdk-go/v1`.

```go
package main

import (
    "context"
    "log"

    "github.com/gitpod-io/gitpod-sdk-go/sdk"
)

func main() {
    ctx := context.Background()
    ona, err := sdk.NewFromEnv()
    if err != nil {
        log.Fatal(err)
    }

    env, err := ona.Environments().Create(ctx, sdk.CreateEnvironmentOptions{
        ContextURL: "https://github.com/gitpod-io/template-golang-cli",
    })
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        _ = ona.Environments().Delete(ctx, env.ID(), sdk.DeleteEnvironmentOptions{Force: true})
    }()
}
```

Call a public RPC with its generated request type:

```go
response, err := ona.Services.Environment.ListEnvironments(
    ctx,
    connect.NewRequest(&v1.ListEnvironmentsRequest{}),
)
if err != nil {
    log.Fatal(err)
}
```

This example requires imports for `connectrpc.com/connect` and `github.com/gitpod-io/gitpod-sdk-go/v1`.

`sdk.NewFromEnv` reads `ONA_API_KEY` and falls back to `GITPOD_API_KEY`. If both are set, `ONA_API_KEY` takes precedence. It uses `https://app.ona.com/api` by default and accepts `ONA_BASE_URL` or `sdk.WithBaseURL(...)` for a custom management-plane domain.

See [`sdk/readme.md`](sdk/readme.md) for the full high-level API and [`sdk/examples`](sdk/examples) for runnable examples.
