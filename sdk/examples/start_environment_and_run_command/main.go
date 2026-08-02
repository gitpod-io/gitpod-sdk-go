package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/sdk"
)

func main() {
	if err := run(); err != nil {
		var authErr *sdk.AuthenticationRequiredError
		if errors.Is(err, sdk.ErrMissingAPIKey) {
			log.Fatalf("set %s to run this example", sdk.APIKeyEnvVar)
		}
		if errors.As(err, &authErr) {
			log.Fatalf("%v\nAuthenticate with github.com in Settings > Git authentications, then rerun this example.", authErr)
		}
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	logger := sdk.NewDebugLogger(os.Stderr)
	ona, err := sdk.NewFromEnv(sdk.WithLogger(logger))
	if err != nil {
		return fmt.Errorf("create SDK: %w", err)
	}

	contextURL := "https://github.com/gitpod-io/template-golang-cli"
	env, err := ona.Environments().Create(ctx, sdk.CreateEnvironmentOptions{
		ContextURL: contextURL,
		Name:       "ona sdk environment interactions example",
	})
	if err != nil {
		return fmt.Errorf("create environment from %s: %w", contextURL, err)
	}
	environmentID := env.ID()
	defer func() {
		deleteOpts := sdk.DeleteEnvironmentOptions{}
		stopCtx, stopCancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Minute)
		if err := ona.Environments().Stop(stopCtx, environmentID); err != nil {
			log.Printf("stop environment %s during cleanup: %v", environmentID, err)
			deleteOpts.Force = true
		}
		stopCancel()

		deleteCtx, deleteCancel := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Minute)
		defer deleteCancel()
		if err := ona.Environments().Delete(deleteCtx, environmentID, deleteOpts); err != nil {
			log.Printf("delete environment %s during cleanup: %v", environmentID, err)
		}
	}()

	workspaceDir := environmentWorkspaceDir(env)
	if workspaceDir == "" {
		return fmt.Errorf("environment %s did not report a workspace directory", environmentID)
	}

	quotedWorkspaceDir := shellQuote(workspaceDir)
	command := fmt.Sprintf(`set -u
cd %s
printf 'workspace_dir=%%s\n' %s
echo "pwd=$(pwd)"
echo "user=$(whoami)"
echo "kernel=$(uname -srm)"
echo "files:"
find . -maxdepth 2 -type f 2>/dev/null | sed -n '1,20p'
if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  echo "git_remote=$(git remote get-url origin 2>/dev/null || true)"
  git status --short
else
  echo "not inside a git work tree" >&2
fi
echo "stderr is captured too" >&2`, quotedWorkspaceDir, quotedWorkspaceDir)

	result, err := env.RunCommand(ctx, sdk.RunCommandOptions{
		WorkingDirectory: workspaceDir,
		Command:          command,
		TimeoutSeconds:   600,
	})
	if err != nil {
		return fmt.Errorf("run command in environment %s: %w", environmentID, err)
	}
	printExecResult(result)
	if result.ExitCode != 0 {
		return fmt.Errorf("command exited with code %d", result.ExitCode)
	}

	readPath := path.Join(workspaceDir, "README.md")
	readme, err := env.ReadFile(ctx, readPath, sdk.ReadFileOptions{Length: 32 * 1024})
	if err != nil {
		return fmt.Errorf("read %s from environment %s: %w", readPath, environmentID, err)
	}
	if content := readme.GetContent(); content != nil {
		fmt.Printf("read %d bytes from %s\n", len(content.GetData()), readPath)
	}

	writePath := path.Join(workspaceDir, "ona-sdk-example.txt")
	writeContent := []byte("created by the Ona SDK environment interactions example\n")
	written, err := env.WriteFile(ctx, writePath, writeContent, sdk.WriteFileOptions{Mode: sdk.WriteFileModeCreateOrTruncate})
	if err != nil {
		return fmt.Errorf("write %s in environment %s: %w", writePath, environmentID, err)
	}
	fmt.Printf("wrote %d bytes to %s\n", written.GetBytesWritten(), writePath)

	verified, err := env.ReadFile(ctx, writePath, sdk.ReadFileOptions{Length: int64(len(writeContent))})
	if err != nil {
		return fmt.Errorf("verify %s in environment %s: %w", writePath, environmentID, err)
	}
	fmt.Printf("verified file content: %s", string(verified.GetContent().GetData()))
	return nil
}

func environmentWorkspaceDir(env *sdk.Environment) string {
	if env == nil || env.Proto() == nil {
		return ""
	}
	status := env.Proto().GetStatus()
	if workspaceDir := status.GetDevcontainer().GetRemoteWorkspaceFolder(); workspaceDir != "" {
		return workspaceDir
	}
	return status.GetContent().GetContentLocationInMachine()
}

func printExecResult(result *sdk.RunCommandResult) {
	fmt.Printf("environment: %s\n", result.EnvironmentID)
	fmt.Printf("exit_code: %d\n", result.ExitCode)
	fmt.Println("stdout:")
	if result.Stdout == "" {
		fmt.Println("<empty>")
	} else {
		fmt.Print(result.Stdout)
	}
	fmt.Println("stderr:")
	if result.Stderr == "" {
		fmt.Println("<empty>")
	} else {
		fmt.Print(result.Stderr)
	}
}

func shellQuote(value string) string {
	if value == "" {
		return "''"
	}
	return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'"
}
