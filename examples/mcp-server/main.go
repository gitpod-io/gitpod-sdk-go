package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/gitpod-io/gitpod-sdk-go"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	var token string
	if envvar := os.Getenv("GITPOD_API_KEY"); envvar != "" {
		token = envvar
	} else if len(os.Args) > 1 {
		content, err := os.ReadFile(os.Args[1])
		if err != nil {
			slog.ErrorContext(context.Background(), "Failed to read token file", "error", err, "path", os.Args[1])
			return
		}
		slog.InfoContext(context.Background(), "Using token from file", "path", os.Args[1])
		token = strings.TrimSpace(string(content))
	}
	if token == "" {
		slog.ErrorContext(context.Background(), "No GITPOD_API_KEY environment variable or file found")
		return
	}

	client := gitpod.NewClient(option.WithBearerToken(token))
	identity, err := client.Identity.GetAuthenticatedIdentity(context.Background(), gitpod.IdentityGetAuthenticatedIdentityParams{})
	if err != nil {
		var apierr *gitpod.Error
		if errors.As(err, &apierr) && apierr.StatusCode == 401 {
			slog.ErrorContext(context.Background(), "Unauthorized. Did you set the GITPOD_API_KEY environment variable?", "error", apierr.Error())
		} else {
			slog.ErrorContext(context.Background(), "Failed to get authenticated identity", "error", err)
		}
		return
	}
	userID := identity.Subject.ID

	// Create MCP server
	s := server.NewMCPServer(
		"Gitpod MCP server",
		"0.1.0",
		server.WithResourceCapabilities(false, false),
	)

	s.AddResource(mcp.NewResource("gitpod://projects", "projects",
		mcp.WithResourceDescription("All available Gitpod projects"),
		mcp.WithMIMEType("application/json"),
		mcp.WithAnnotations([]mcp.Role{mcp.RoleAssistant}, 0.8),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]interface{}, error) {
		projects := client.Projects.ListAutoPaging(ctx, gitpod.ProjectListParams{})

		results := []interface{}{}
		for projects.Next() {
			results = append(results, projects.Current())
		}
		if err := projects.Err(); err != nil {
			return nil, err
		}

		return results, nil
	})
	s.AddResource(mcp.NewResource("gitpod://environments", "environments",
		mcp.WithResourceDescription("The user's current Gitpod environments"),
		mcp.WithMIMEType("application/json"),
		mcp.WithAnnotations([]mcp.Role{mcp.RoleAssistant}, 0.8),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]interface{}, error) {
		environments := client.Environments.ListAutoPaging(ctx, gitpod.EnvironmentListParams{
			Filter: gitpod.F(gitpod.EnvironmentListParamsFilter{
				CreatorIDs: gitpod.F([]string{userID}),
			}),
		})

		results := []interface{}{}
		for environments.Next() {
			results = append(results, environments.Current())
		}
		if err := environments.Err(); err != nil {
			return nil, err
		}

		return results, nil
	})

	s.AddTool(mcp.NewTool("create_environment",
		mcp.WithDescription("Create a new Gitpod environment"),
		mcp.WithString("project_id",
			mcp.Required(),
			mcp.Description("The ID of the project to create the environment in"),
		),
	), toolCreateEnvironment(client))
	s.AddTool(mcp.NewTool("stop_environment",
		mcp.WithDescription("Stop a Gitpod environment"),
		mcp.WithString("environment_id",
			mcp.Required(),
			mcp.Description("The ID of the environment to stop"),
		),
	), toolStopEnvironment(client))
	// execute_command
	s.AddTool(mcp.NewTool("execute_command",
		mcp.WithDescription("Execute a command in a Gitpod environment"),
		mcp.WithString("environment_id",
			mcp.Required(),
			mcp.Description("The ID of the environment to execute the command in"),
		),
		mcp.WithString("command",
			mcp.Required(),
			mcp.Description("The command to execute. Will be executed in the root of the project, as a bash script."),
		),
		mcp.WithString("description",
			mcp.Required(),
			mcp.Description("A short description of the command to execute (no more than 200 characters)"),
		),
	), toolExecuteCommand(client))

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func toolCreateEnvironment(client *gitpod.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		projectID, ok := request.Params.Arguments["project_id"].(string)
		if !ok {
			return mcp.NewToolResultError("project_id must be a string"), nil
		}

		environment, err := client.Environments.NewFromProject(ctx, gitpod.EnvironmentNewFromProjectParams{
			ProjectID: gitpod.F(projectID),
		})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create environment: %v", err)), nil
		}
		return watchEnvironment(client, environment.Environment.ID, func(event gitpod.EventWatchResponse) (*mcp.CallToolResult, error) {
			environment, err := client.Environments.Get(ctx, gitpod.EnvironmentGetParams{
				EnvironmentID: gitpod.F(environment.Environment.ID),
			})
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to get environment: %v", err)), nil
			}
			if environment.Environment.Status.Phase == gitpod.EnvironmentPhaseStopping || environment.Environment.Status.Phase == gitpod.EnvironmentPhaseStopped {
				return mcp.NewToolResultError(fmt.Sprintf("environment failed to start: %v", environment.Environment.Status.Phase)), nil
			}
			if environment.Environment.Status.Phase == gitpod.EnvironmentPhaseRunning {
				return mcp.NewToolResultText(fmt.Sprintf("Environment created for project %s", projectID)), nil
			}
			return nil, nil
		})
	}
}

func toolStopEnvironment(client *gitpod.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		environmentID, ok := request.Params.Arguments["environment_id"].(string)
		if !ok {
			return mcp.NewToolResultError("environment_id must be a string"), nil
		}

		_, err := client.Environments.Stop(ctx, gitpod.EnvironmentStopParams{
			EnvironmentID: gitpod.F(environmentID),
		})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to stop environment: %v", err)), nil
		}

		return watchEnvironment(client, environmentID, func(event gitpod.EventWatchResponse) (*mcp.CallToolResult, error) {
			environment, err := client.Environments.Get(ctx, gitpod.EnvironmentGetParams{
				EnvironmentID: gitpod.F(environmentID),
			})
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to get environment: %v", err)), nil
			}
			if environment.Environment.Status.Phase == gitpod.EnvironmentPhaseStopped {
				return mcp.NewToolResultText("Environment stopped"), nil
			}
			return nil, nil
		})
	}
}

func watchEnvironment[T any](client *gitpod.Client, environmentID string, onEvent func(event gitpod.EventWatchResponse) (*T, error)) (*T, error) {
	events := client.Events.WatchStreaming(context.Background(), gitpod.EventWatchParams{
		EnvironmentID: gitpod.F(environmentID),
	})
	for events.Next() {
		event := events.Current()
		res, err := onEvent(event)
		if err != nil {
			return nil, err
		}
		if res != nil {
			return res, nil
		}
	}
	if err := events.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}

func toolExecuteCommand(client *gitpod.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		environmentID, ok := request.Params.Arguments["environment_id"].(string)
		if !ok {
			return mcp.NewToolResultError("environment_id must be a string"), nil
		}
		command, ok := request.Params.Arguments["command"].(string)
		if !ok {
			return mcp.NewToolResultError("command must be a string"), nil
		}
		description, ok := request.Params.Arguments["description"].(string)
		if !ok {
			return mcp.NewToolResultError("description must be a string"), nil
		}

		task, err := client.Environments.Automations.Tasks.New(ctx, gitpod.EnvironmentAutomationTaskNewParams{
			EnvironmentID: gitpod.F(environmentID),
			Metadata: gitpod.F(shared.TaskMetadataParam{
				Name:        gitpod.F(command),
				Description: gitpod.F(description),
			}),
			Spec: gitpod.F(shared.TaskSpecParam{
				Command: gitpod.F(command),
			}),
		})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create task: %v", err)), nil
		}
		defer client.Environments.Automations.Tasks.Delete(ctx, gitpod.EnvironmentAutomationTaskDeleteParams{
			ID: gitpod.F(task.Task.ID),
		})

		run, err := client.Environments.Automations.Tasks.Start(ctx, gitpod.EnvironmentAutomationTaskStartParams{
			ID: gitpod.F(task.Task.ID),
		})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to start task: %v", err)), nil
		}
		logURL, err := watchEnvironment(client, environmentID, func(event gitpod.EventWatchResponse) (*string, error) {
			resp, err := client.Environments.Automations.Tasks.Executions.Get(ctx, gitpod.EnvironmentAutomationTaskExecutionGetParams{
				ID: gitpod.F(run.TaskExecution.ID),
			})
			if err != nil {
				return nil, fmt.Errorf("Failed to get task: %v", err)
			}
			if resp.TaskExecution.Status.LogURL != "" {
				return &resp.TaskExecution.Status.LogURL, nil
			}
			return nil, nil
		})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get task: %v", err)), nil
		}

		logToken, err := client.Environments.NewLogsToken(ctx, gitpod.EnvironmentNewLogsTokenParams{
			EnvironmentID: gitpod.F(environmentID),
		})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get log token: %v", err)), nil
		}

		req, err := http.NewRequestWithContext(ctx, "GET", *logURL, nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create request: %v", err)), nil
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", logToken.AccessToken))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get log: %v", err)), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to read log: %v", err)), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("%s", string(body))), nil
	}
}
