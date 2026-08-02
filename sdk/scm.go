package sdk

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"

	rawclient "github.com/gitpod-io/gitpod-sdk-go/internal/raw"
	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
)

type scmClient struct {
	sdk *Client
}

type resolveContextOptions struct {
	runnerID   string
	contextURL string
}

type checkAuthenticationOptions struct {
	runnerID string
	host     string
}

func (c *scmClient) resolveContext(ctx context.Context, opts resolveContextOptions) (*v1.ParseContextURLResponse, error) {
	operation := "scm.resolve_context"
	raw, err := c.raw(operation)
	if err != nil {
		return nil, err
	}
	if opts.contextURL == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("context URL is required")}}
	}

	c.sdk.logger().DebugContext(ctx, "resolving context URL",
		"operation", operation,
		"context_url", safeURLForLog(opts.contextURL),
		"runner_id", opts.runnerID,
	)
	resp, err := raw.RunnerService().ParseContextURL(ctx, connect.NewRequest(&v1.ParseContextURLRequest{
		RunnerId:   opts.runnerID,
		ContextUrl: opts.contextURL,
	}))
	if err != nil {
		return nil, MapError(operation, err)
	}
	c.sdk.logger().DebugContext(ctx, "resolved context URL",
		"operation", operation,
		"context_url", safeURLForLog(resp.Msg.GetOriginalContextUrl()),
		"runner_id", opts.runnerID,
		"scm_id", resp.Msg.GetScmId(),
		"git_host", resp.Msg.GetGit().GetHost(),
		"git_owner", resp.Msg.GetGit().GetOwner(),
		"git_repo", resp.Msg.GetGit().GetRepo(),
		"project_count", len(resp.Msg.GetProjectIds()),
		"recommended_environment_class_count", len(resp.Msg.GetRecommendedEnvironmentClasses()),
	)
	return resp.Msg, nil
}

func (c *scmClient) checkAuthenticationForHost(ctx context.Context, opts checkAuthenticationOptions) (*v1.CheckAuthenticationForHostResponse, error) {
	operation := "scm.check_authentication_for_host"
	raw, err := c.raw(operation)
	if err != nil {
		return nil, err
	}
	if opts.host == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("host is required")}}
	}

	c.sdk.logger().DebugContext(ctx, "checking SCM authentication",
		"operation", operation,
		"runner_id", opts.runnerID,
		"host", opts.host,
	)
	resp, err := raw.RunnerService().CheckAuthenticationForHost(ctx, connect.NewRequest(&v1.CheckAuthenticationForHostRequest{
		RunnerId: opts.runnerID,
		Host:     opts.host,
	}))
	if err != nil {
		return nil, MapError(operation, err)
	}
	if !resp.Msg.GetAuthenticated() {
		patSupported := resp.Msg.GetSupportsPat() != nil
		if !patSupported {
			field := resp.Msg.ProtoReflect().Descriptor().Fields().ByName("pat_supported")
			patSupported = field != nil && resp.Msg.ProtoReflect().Get(field).Bool()
		}
		c.sdk.logger().DebugContext(ctx, "SCM authentication required",
			"operation", operation,
			"runner_id", opts.runnerID,
			"host", opts.host,
			"scm_id", resp.Msg.GetScmId(),
			"scm_name", resp.Msg.GetScmName(),
			"supports_oauth2", resp.Msg.GetSupportsOauth2() != nil,
			"supports_pat", patSupported,
		)
		return resp.Msg, &AuthenticationRequiredError{operationError{
			operation: operation,
			err:       fmt.Errorf("authentication required for %s", opts.host),
		}}
	}
	c.sdk.logger().DebugContext(ctx, "SCM authentication available",
		"operation", operation,
		"runner_id", opts.runnerID,
		"host", opts.host,
		"scm_id", resp.Msg.GetScmId(),
		"scm_name", resp.Msg.GetScmName(),
	)
	return resp.Msg, nil
}

func (c *scmClient) raw(operation string) (*rawclient.ManagementPlane, error) {
	if c == nil || c.sdk == nil {
		return nil, &ValidationError{operationError{operation: operation, err: errSDKClientMissing}}
	}
	return c.sdk.requireRaw(operation)
}
