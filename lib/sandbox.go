package lib

import (
	"context"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"

	"github.com/stainless-sdks/gitpod-go"
)

type Sandbox interface {
	io.Closer
	Exec(ctx context.Context, command string) (io.ReadCloser, error)
}

func NewSandbox(ctx context.Context, client *gitpod.Client, environmentID string) (Sandbox, error) {
	logTokenResp, err := client.Environments.NewLogsToken(ctx, gitpod.EnvironmentNewLogsTokenParams{
		EnvironmentID: gitpod.String(environmentID),
	})
	if err != nil {
		return nil, err
	}

	taskReference := fmt.Sprintf("sandbox-%d", rand.Uint64())
	resp, err := client.Environments.Automations.Tasks.New(ctx, gitpod.EnvironmentAutomationTaskNewParams{
		EnvironmentID: gitpod.String(environmentID),
		Metadata: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsMetadata{
			Reference:   gitpod.String(taskReference),
			Name:        gitpod.String("Sandbox Task"),
			Description: gitpod.String("Sandbox task"),
		}),
	})
	if err != nil {
		return nil, nil
	}

	return &sandbox{
		taskID:         resp.Task.ID,
		environmentID:  environmentID,
		logAccessToken: logTokenResp.AccessToken,
		client:         client,
	}, nil
}

type sandbox struct {
	taskID         string
	environmentID  string
	logAccessToken string
	client         *gitpod.Client
}

func (s *sandbox) Close() error {
	_, err := s.client.Environments.Automations.Tasks.Delete(context.Background(), gitpod.EnvironmentAutomationTaskDeleteParams{
		ID: gitpod.String(s.taskID),
	})
	return err
}

func (s *sandbox) Exec(ctx context.Context, command string) (io.ReadCloser, error) {
	_, err := s.client.Environments.Automations.Tasks.Update(ctx, gitpod.EnvironmentAutomationTaskUpdateParams{
		ID: gitpod.String(s.taskID),
		Spec: gitpod.F[gitpod.EnvironmentAutomationTaskUpdateParamsSpecUnion](gitpod.EnvironmentAutomationTaskUpdateParamsSpecCommand{
			Command: gitpod.String(command),
		}),
	})
	if err != nil {
		return nil, err
	}
	startResp, err := s.client.Environments.Automations.Tasks.Start(ctx, gitpod.EnvironmentAutomationTaskStartParams{
		ID: gitpod.String(s.taskID),
	})
	if err != nil {
		return nil, err
	}
	taskExecutionID := startResp.TaskExecution.ID
	taskExecution, err := waitFor(ctx, s.client, s.environmentID, gitpod.EventWatchResponseResourceTypeResourceTypeTaskExecution, taskExecutionID, func() (*gitpod.EnvironmentAutomationTaskExecutionGetResponseTaskExecution, bool, error) {
		resp, err := s.client.Environments.Automations.Tasks.Executions.Get(ctx, gitpod.EnvironmentAutomationTaskExecutionGetParams{
			ID: gitpod.String(taskExecutionID),
		})
		if err != nil {
			return nil, false, err
		}
		if resp.TaskExecution.Status.LogURL != "" {
			return &resp.TaskExecution, true, nil
		}
		return nil, false, nil
	})
	if err != nil {
		return nil, err
	}
	logURL := taskExecution.Status.LogURL

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, logURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.logAccessToken))
	logResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return logResp.Body, nil
}
