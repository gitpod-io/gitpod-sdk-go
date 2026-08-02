package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/sdk"
	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
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
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Minute)
	defer cancel()

	logger := sdk.NewDebugLogger(os.Stderr)
	ona, err := sdk.NewFromEnv(sdk.WithLogger(logger))
	if err != nil {
		return fmt.Errorf("create SDK: %w", err)
	}

	contextURL := "https://github.com/gitpod-io/template-golang-cli"
	initialPrompt := `Inspect the repository, create ona-codex-agent-notes.md in the workspace with a short summary and one suggested improvement, then finish without asking for confirmation.`
	run, err := ona.RunCodex(ctx, sdk.RunCodexOptions{
		RepositoryURL:   contextURL,
		Task:            initialPrompt,
		EnvironmentName: "ona sdk Codex agent example",
		AgentName:       "Create an SDK example note",
		Model:           v1.CodexOpenAIModel_CODEX_OPEN_AI_MODEL_GPT_5_6_SOL,
		ReasoningEffort: v1.CodexReasoningEffort_CODEX_REASONING_EFFORT_HIGH,
	})
	if run != nil {
		defer cleanupRun(ctx, run)
	}
	if err != nil {
		if run != nil {
			fmt.Fprintf(os.Stderr, "environment: %s\n", run.EnvironmentID())
			fmt.Fprintf(os.Stderr, "agent_execution_id: %s\n", run.ID())
		}
		return fmt.Errorf("run Codex for %s: %w", contextURL, err)
	}
	env := run.Environment()
	environmentID := run.EnvironmentID()

	fmt.Fprintf(os.Stderr, "environment: %s\n", environmentID)
	fmt.Fprintf(os.Stderr, "agent_execution_id: %s\n", run.ID())

	stream, err := run.MessageStream(ctx)
	if err != nil {
		return fmt.Errorf("open Codex message stream %s: %w", run.ID(), err)
	}
	streamDone := make(chan error, 1)
	go func() {
		_, err := io.Copy(os.Stdout, stream)
		streamDone <- err
	}()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
		for scanner.Scan() {
			message := scanner.Text()
			if strings.TrimSpace(message) == "" {
				continue
			}
			if err := run.SendMessage(ctx, message); err != nil {
				log.Printf("send stdin message to Codex %s: %v", run.ID(), err)
				return
			}
		}
		if err := scanner.Err(); err != nil {
			log.Printf("read stdin for Codex %s: %v", run.ID(), err)
		}
	}()

	lastLine := ""
	result, err := run.WatchResult(ctx, func(_ context.Context, exec *v1.AgentExecution) error {
		line := sdk.AgentStatusLine(exec)
		if line == lastLine {
			return nil
		}
		fmt.Fprintf(os.Stderr, "agent update: %s\n", line)
		lastLine = line
		return nil
	})
	if err != nil {
		_ = stream.Close()
		return fmt.Errorf("watch Codex result %s: %w", run.ID(), err)
	}
	_ = stream.Close()
	if err := <-streamDone; err != nil && !errors.Is(err, context.Canceled) && !errors.Is(err, io.ErrClosedPipe) {
		return fmt.Errorf("read Codex message stream %s: %w", run.ID(), err)
	}

	fmt.Fprintf(os.Stderr, "Codex finished: %s\n", sdk.AgentStatusLine(result))
	if result.GetStatus().GetPhase() == v1.AgentExecution_PHASE_WAITING_FOR_INPUT {
		fmt.Fprintln(os.Stderr, "Codex is waiting for user input.")
		return nil
	}

	changes, err := env.GitChanges(ctx, sdk.GitChangesOptions{Unified: 3})
	if err != nil {
		return fmt.Errorf("read git changes for environment %s: %w", environmentID, err)
	}

	status := changes.Status
	fmt.Fprintf(os.Stderr, "repository: %s\n", status.GetCloneUrl())
	fmt.Fprintf(os.Stderr, "branch: %s\n", status.GetBranch())
	fmt.Fprintf(os.Stderr, "changed_files: %d\n", status.GetTotalChangedFiles())
	for _, changed := range changes.Files {
		file := changed.File
		diff := changed.Diff
		fmt.Fprintf(os.Stderr, "%s %s hunks=%d binary=%t\n", file.GetChangeType().String(), file.GetPath(), len(diff.GetHunks()), diff.GetIsBinary())
	}
	return nil
}

func cleanupRun(ctx context.Context, run *sdk.CodexRun) {
	environmentID := run.EnvironmentID()
	deleteOpts := sdk.DeleteEnvironmentOptions{}
	stopCtx, stopCancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Minute)
	if err := run.StopEnvironment(stopCtx); err != nil {
		log.Printf("stop environment %s during cleanup: %v", environmentID, err)
		deleteOpts.Force = true
	}
	stopCancel()

	deleteCtx, deleteCancel := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Minute)
	defer deleteCancel()
	if err := run.DeleteEnvironment(deleteCtx, deleteOpts); err != nil {
		log.Printf("delete environment %s during cleanup: %v", environmentID, err)
	}
}
