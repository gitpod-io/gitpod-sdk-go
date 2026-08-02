package sdk

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"net/url"
	"reflect"
	"strings"
	"time"

	"connectrpc.com/connect"

	rawclient "github.com/gitpod-io/gitpod-sdk-go/internal/raw"
	"github.com/gitpod-io/gitpod-sdk-go/internal/raw/cache"
	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const defaultProjectUsageLookback = 90 * 24 * time.Hour

// EnvironmentClient exposes environment workflows that are useful from automation.
type EnvironmentClient struct {
	sdk *Client
}

// Environment is an Ona environment with convenience methods for automation tasks.
type Environment struct {
	sdk         *Client
	environment *environmentHandle
	ops         *environmentOps
}

// CreateEnvironmentOptions controls environment creation.
type CreateEnvironmentOptions struct {
	ContextURL string
	Name       string
}

// DeleteEnvironmentOptions controls environment deletion.
type DeleteEnvironmentOptions struct {
	Force bool
}

type environmentHandle struct {
	client      *EnvironmentClient
	id          string
	environment *v1.Environment
}

type createEnvironmentFromContextURLOptions struct {
	ContextURL       string
	EnvironmentClass string
	Name             string
}

type createScratchEnvironmentOptions struct {
	EnvironmentClass string
	Name             string
}

type createEnvironmentFromProjectOptions struct {
	ProjectID string
	Resolved  *v1.ParseContextURLResponse
	Name      string
}

// Create creates and starts a new environment from a context URL.
func (c *EnvironmentClient) Create(ctx context.Context, opts CreateEnvironmentOptions) (*Environment, error) {
	ctx = c.sdk.requestContext(ctx)
	operation := "environments.create"

	if opts.ContextURL == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("context URL is required")}}
	}

	logger := c.sdk.logger()
	logger.DebugContext(ctx, "creating environment",
		"operation", operation,
		"context_url", safeURLForLog(opts.ContextURL),
		"name", opts.Name,
	)

	resolved, err := c.sdk.scm().resolveContext(ctx, resolveContextOptions{
		contextURL: opts.ContextURL,
	})
	if err != nil {
		return nil, err
	}

	environmentClass := ""
	environmentClassRunnerID := ""
	environmentClassSource := ""
	selectEnvironmentClass := func() (string, error) {
		if environmentClass != "" {
			return environmentClass, nil
		}
		environmentClass = firstRecommendedEnvironmentClass(resolved)
		if environmentClass == "" {
			var err error
			environmentClass, environmentClassRunnerID, err = c.defaultEnvironmentClass(ctx)
			if err != nil {
				return "", err
			}
			environmentClassSource = "default"
		} else {
			environmentClassSource = "recommended"
		}
		return environmentClass, nil
	}

	environmentClassID, err := selectEnvironmentClass()
	if err != nil {
		return nil, err
	}
	authRunnerID := environmentClassRunnerID
	if authRunnerID == "" {
		authRunnerID, err = c.environmentClassRunnerID(ctx, environmentClassID)
		if err != nil {
			return nil, err
		}
	}
	if err := c.requireSCMAuthentication(ctx, authRunnerID, opts.ContextURL); err != nil {
		return nil, err
	}

	projectID := c.preferredProjectID(ctx, resolved.GetProjectIds())
	if projectID != "" {
		logger.DebugContext(ctx, "creating environment from project",
			"operation", operation,
			"project_id", projectID,
			"project_candidates", len(resolved.GetProjectIds()),
		)
		env, err := c.createFromProject(ctx, createEnvironmentFromProjectOptions{
			ProjectID: projectID,
			Resolved:  resolved,
			Name:      opts.Name,
		})
		if err != nil {
			return nil, err
		}
		return c.environment(ctx, env, true)
	}
	logger.DebugContext(ctx, "selected environment class",
		"operation", operation,
		"environment_class_id", environmentClass,
		"selection_source", environmentClassSource,
		"recommended_count", len(resolved.GetRecommendedEnvironmentClasses()),
	)

	env, err := c.createFromContextURL(ctx, createEnvironmentFromContextURLOptions{
		ContextURL:       opts.ContextURL,
		EnvironmentClass: environmentClass,
		Name:             opts.Name,
	})
	if err != nil {
		return nil, err
	}
	return c.environment(ctx, env, true)
}

// Get retrieves an environment by ID.
func (c *EnvironmentClient) Get(ctx context.Context, environmentID string) (*Environment, error) {
	ctx = c.sdk.requestContext(ctx)
	env, err := c.get(ctx, environmentID)
	if err != nil {
		return nil, err
	}
	return c.environment(ctx, env, false)
}

// List returns a lazy sequence over caller-owned default environments.
func (c *EnvironmentClient) List(ctx context.Context) iter.Seq2[*Environment, error] {
	return func(yield func(*Environment, error) bool) {
		operation := "environments.list"
		if c == nil || c.sdk == nil {
			yield(nil, &ValidationError{operationError{operation: operation, err: errSDKClientMissing}})
			return
		}
		iterCtx := c.sdk.requestContext(ctx)
		raw, err := c.raw(operation)
		if err != nil {
			yield(nil, err)
			return
		}
		environmentService := raw.EnvironmentService()
		environmentServiceValue := reflect.ValueOf(environmentService)
		if environmentService == nil || environmentServiceValue.Kind() == reflect.Ptr && environmentServiceValue.IsNil() {
			yield(nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("environment service is not configured")}})
			return
		}
		identityService := raw.IdentityService()
		identityServiceValue := reflect.ValueOf(identityService)
		if identityService == nil || identityServiceValue.Kind() == reflect.Ptr && identityServiceValue.IsNil() {
			yield(nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("identity service is not configured")}})
			return
		}

		identity, err := identityService.GetAuthenticatedIdentity(iterCtx, connect.NewRequest(&v1.GetAuthenticatedIdentityRequest{}))
		if err != nil {
			yield(nil, MapError(operation, err))
			return
		}
		creatorID := identity.Msg.GetSubject().GetId()
		if creatorID == "" {
			yield(nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("authenticated identity did not include a subject ID")}})
			return
		}

		rawFilter := &v1.ListEnvironmentsRequest_Filter{
			CreatorIds:     []string{creatorID},
			Roles:          []v1.EnvironmentRole{v1.EnvironmentRole_ENVIRONMENT_ROLE_DEFAULT},
			ArchivalStatus: Pointer(v1.ListEnvironmentsRequest_ARCHIVAL_STATUS_ALL),
		}
		cfg := c.sdk.config()
		var pageToken string
		var count int
		for {
			c.sdk.logger().DebugContext(iterCtx, "listing environments",
				"operation", operation,
				"creator_id", creatorID,
				"page_size", cfg.pageSize,
				"has_page_token", pageToken != "",
			)
			resp, err := environmentService.ListEnvironments(iterCtx, connect.NewRequest(&v1.ListEnvironmentsRequest{
				Pagination: &v1.PaginationRequest{
					PageSize: cfg.pageSize,
					Token:    pageToken,
				},
				Filter: rawFilter,
			}))
			if err != nil {
				yield(nil, MapError(operation, err))
				return
			}
			for _, env := range resp.Msg.GetEnvironments() {
				count++
				if !yield(&Environment{
					sdk:         c.sdk,
					environment: c.handle(env),
				}, nil) {
					return
				}
			}

			nextToken := resp.Msg.GetPagination().GetNextToken()
			c.sdk.logger().DebugContext(iterCtx, "listed environment page",
				"operation", operation,
				"creator_id", creatorID,
				"page_count", len(resp.Msg.GetEnvironments()),
				"total_count", count,
				"has_next_page", nextToken != "",
			)
			if nextToken == "" {
				break
			}
			pageToken = nextToken
		}

		c.sdk.logger().DebugContext(iterCtx, "listed environments",
			"operation", operation,
			"creator_id", creatorID,
			"count", count,
		)
	}
}

// Delete deletes one environment by ID.
func (c *EnvironmentClient) Delete(ctx context.Context, environmentID string, opts DeleteEnvironmentOptions) error {
	ctx = c.sdk.requestContext(ctx)
	return c.delete(ctx, environmentID, opts)
}

// Stop stops an existing environment.
func (c *EnvironmentClient) Stop(ctx context.Context, environmentID string) error {
	ctx = c.sdk.requestContext(ctx)
	operation := "environments.stop"
	raw, err := c.raw(operation)
	if err != nil {
		return err
	}
	if environmentID == "" {
		return &ValidationError{operationError{operation: operation, err: errors.New("environment ID is required")}}
	}
	c.sdk.logger().DebugContext(ctx, "stopping environment", "operation", operation, "environment_id", environmentID)
	_, err = raw.EnvironmentService().StopEnvironment(ctx, connect.NewRequest(&v1.StopEnvironmentRequest{
		EnvironmentId: environmentID,
	}))
	if err != nil {
		return MapError(operation, err)
	}
	c.sdk.logger().DebugContext(ctx, "environment stop requested", "operation", operation, "environment_id", environmentID)
	env := &environmentHandle{client: c, id: environmentID}
	_, err = env.waitStopped(ctx)
	return err
}

// ID returns the environment ID.
func (e *Environment) ID() string {
	if e == nil || e.environment == nil {
		return ""
	}
	return e.environment.idString()
}

// Proto returns the latest fetched environment message.
func (e *Environment) Proto() *v1.Environment {
	if e == nil || e.environment == nil {
		return nil
	}
	return e.environment.proto()
}

func (c *EnvironmentClient) get(ctx context.Context, environmentID string) (*environmentHandle, error) {
	operation := "environments.get"
	raw, err := c.raw(operation)
	if err != nil {
		return nil, err
	}
	if environmentID == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("environment ID is required")}}
	}

	resp, err := raw.EnvironmentService().GetEnvironment(ctx, connect.NewRequest(&v1.GetEnvironmentRequest{
		EnvironmentId: environmentID,
	}))
	if err != nil {
		return nil, MapError(operation, err)
	}
	return c.handle(resp.Msg.GetEnvironment()), nil
}

func (c *EnvironmentClient) createFromContextURL(ctx context.Context, opts createEnvironmentFromContextURLOptions) (*environmentHandle, error) {
	operation := "environments.create_from_context_url"
	raw, err := c.raw(operation)
	if err != nil {
		return nil, err
	}
	if opts.ContextURL == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("context URL is required")}}
	}

	c.sdk.logger().DebugContext(ctx, "sending environment create request",
		"operation", operation,
		"context_url", safeURLForLog(opts.ContextURL),
		"environment_class_id", opts.EnvironmentClass,
		"name", opts.Name,
	)
	resp, err := raw.EnvironmentService().CreateEnvironment(ctx, connect.NewRequest(&v1.CreateEnvironmentRequest{
		Spec: environmentSpecFromContextURL(opts),
		Name: optionalString(opts.Name),
	}))
	if err != nil {
		return nil, MapError(operation, err)
	}
	c.sdk.logger().DebugContext(ctx, "environment created",
		"operation", operation,
		"environment_id", resp.Msg.GetEnvironment().GetId(),
		"phase", resp.Msg.GetEnvironment().GetStatus().GetPhase().String(),
		"environment_class_id", opts.EnvironmentClass,
	)
	return c.handle(resp.Msg.GetEnvironment()), nil
}

func (c *EnvironmentClient) createScratch(ctx context.Context, name string) (*Environment, error) {
	ctx = c.sdk.requestContext(ctx)
	environmentClass, _, err := c.defaultEnvironmentClass(ctx)
	if err != nil {
		return nil, err
	}
	handle, err := c.createScratchEnvironment(ctx, createScratchEnvironmentOptions{
		EnvironmentClass: environmentClass,
		Name:             name,
	})
	if err != nil {
		return nil, err
	}
	return c.environment(ctx, handle, true)
}

func (c *EnvironmentClient) createScratchEnvironment(ctx context.Context, opts createScratchEnvironmentOptions) (*environmentHandle, error) {
	operation := "environments.create_scratch"
	raw, err := c.raw(operation)
	if err != nil {
		return nil, err
	}

	c.sdk.logger().DebugContext(ctx, "sending scratch environment create request",
		"operation", operation,
		"environment_class_id", opts.EnvironmentClass,
		"name", opts.Name,
	)
	resp, err := raw.EnvironmentService().CreateEnvironment(ctx, connect.NewRequest(&v1.CreateEnvironmentRequest{
		Spec: &v1.EnvironmentSpec{
			DesiredPhase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_RUNNING,
			Machine:      &v1.EnvironmentSpec_Machine{Class: opts.EnvironmentClass},
		},
		Name: optionalString(opts.Name),
	}))
	if err != nil {
		return nil, MapError(operation, err)
	}
	return c.handle(resp.Msg.GetEnvironment()), nil
}

func (c *EnvironmentClient) createFromProject(ctx context.Context, opts createEnvironmentFromProjectOptions) (*environmentHandle, error) {
	operation := "environments.create_from_project"
	raw, err := c.raw(operation)
	if err != nil {
		return nil, err
	}
	if opts.ProjectID == "" {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("project ID is required")}}
	}

	gitRemoteURI := projectRemoteURIForParsedContext(ctx, raw, opts.ProjectID, opts.Resolved.GetGit())
	c.sdk.logger().DebugContext(ctx, "sending project environment create request",
		"operation", operation,
		"project_id", opts.ProjectID,
		"name", opts.Name,
	)
	resp, err := raw.EnvironmentService().CreateEnvironmentFromProject(ctx, connect.NewRequest(&v1.CreateEnvironmentFromProjectRequest{
		ProjectId: opts.ProjectID,
		Spec:      environmentSpecFromProject(opts, gitRemoteURI),
		Name:      optionalString(opts.Name),
	}))
	if err != nil {
		return nil, MapError(operation, err)
	}
	c.sdk.logger().DebugContext(ctx, "project environment created",
		"operation", operation,
		"environment_id", resp.Msg.GetEnvironment().GetId(),
		"project_id", opts.ProjectID,
		"phase", resp.Msg.GetEnvironment().GetStatus().GetPhase().String(),
	)
	return c.handle(resp.Msg.GetEnvironment()), nil
}

func (c *EnvironmentClient) delete(ctx context.Context, environmentID string, opts DeleteEnvironmentOptions) error {
	operation := "environments.delete"
	raw, err := c.raw(operation)
	if err != nil {
		return err
	}
	if environmentID == "" {
		return &ValidationError{operationError{operation: operation, err: errors.New("environment ID is required")}}
	}
	c.sdk.logger().DebugContext(ctx, "deleting environment",
		"operation", operation,
		"environment_id", environmentID,
		"force", opts.Force,
	)
	_, err = raw.EnvironmentService().DeleteEnvironment(ctx, connect.NewRequest(&v1.DeleteEnvironmentRequest{
		EnvironmentId: environmentID,
		Force:         opts.Force,
	}))
	if err != nil {
		return MapError(operation, err)
	}
	c.sdk.logger().DebugContext(ctx, "environment deleted",
		"operation", operation,
		"environment_id", environmentID,
		"force", opts.Force,
	)
	return nil
}

func (h *environmentHandle) idString() string {
	if h == nil {
		return ""
	}
	return h.id
}

func (h *environmentHandle) proto() *v1.Environment {
	if h == nil {
		return nil
	}
	return h.environment
}

func (h *environmentHandle) waitRunning(ctx context.Context) (*v1.Environment, error) {
	return h.waitForEnvironment(ctx, "environments.wait_running", func(env *v1.Environment) (bool, error) {
		if isTerminalBeforeRunning(env.GetStatus().GetPhase()) {
			return false, &UnavailableError{operationError{
				operation: "environments.wait_running",
				err:       fmt.Errorf("environment %s reached %s before running", env.GetId(), environmentPhaseLabel(env.GetStatus().GetPhase())),
			}}
		}
		return env.GetStatus().GetPhase() == v1.EnvironmentPhase_ENVIRONMENT_PHASE_RUNNING, nil
	})
}

func (h *environmentHandle) waitStopped(ctx context.Context) (*v1.Environment, error) {
	return h.waitForEnvironment(ctx, "environments.wait_stopped", func(env *v1.Environment) (bool, error) {
		phase := env.GetStatus().GetPhase()
		return phase == v1.EnvironmentPhase_ENVIRONMENT_PHASE_STOPPED || phase == v1.EnvironmentPhase_ENVIRONMENT_PHASE_DELETED, nil
	})
}

func (h *environmentHandle) waitForEnvironment(ctx context.Context, operation string, done func(*v1.Environment) (bool, error)) (*v1.Environment, error) {
	if h == nil || h.client == nil {
		return nil, &ValidationError{operationError{operation: operation, err: errors.New("environment handle is not initialized")}}
	}
	raw, err := h.client.raw(operation)
	if err != nil {
		return nil, err
	}
	if raw.EnvironmentService() == nil {
		return nil, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("environment service is not configured")}}
	}
	var lastLogged v1.EnvironmentPhase
	logged := false
	changes := make(chan struct{}, 64)

	informer, err := cache.NewEnvironmentCache(ctx, raw.EnvironmentService(), h.id,
		cache.WithNoFullSync(),
	)
	if err != nil {
		return nil, &SDKError{operationError{operation: operation, err: err}}
	}
	defer func() { _ = informer.Close() }()

	observe := func(ctx context.Context) (*v1.Environment, bool, error) {
		env, err := informer.Get(ctx, h.id)
		if err != nil {
			return nil, false, MapError(operation, err)
		}
		h.environment = env
		phase := env.GetStatus().GetPhase()
		if !logged || phase != lastLogged {
			h.client.sdk.logger().DebugContext(ctx, "environment phase changed",
				"operation", operation,
				"environment_id", env.GetId(),
				"phase", phase.String(),
			)
			lastLogged = phase
			logged = true
		}
		if err := environmentFailureError(operation, env); err != nil {
			h.client.sdk.logger().DebugContext(ctx, "environment failed while waiting",
				"operation", operation,
				"environment_id", env.GetId(),
				"err", err,
			)
			return env, false, err
		}
		done, err := done(env)
		return env, done, err
	}

	last, reached, err := observe(ctx)
	if err != nil || reached {
		return last, err
	}

	if raw.EventService() == nil {
		return last, &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("event service is not configured")}}
	}

	watchCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	watchDone := make(chan error, 1)

	go func() {
		h.client.sdk.logger().DebugContext(watchCtx, "watching environment events",
			"operation", operation,
			"environment_id", h.id,
		)
		stream, err := raw.EventService().WatchEvents(watchCtx, connect.NewRequest(environmentWatchRequest(h.id)))
		if err != nil && !errors.Is(err, context.Canceled) {
			h.client.sdk.logger().ErrorContext(watchCtx, "environment informer event watch failed",
				"operation", operation,
				"environment_id", h.id,
				"err", err,
			)
			watchDone <- err
			return
		}
		if err != nil {
			return
		}
		defer func() { _ = stream.Close() }()

		for stream.Receive() {
			event := stream.Msg()
			if event.GetResourceType() != v1.ResourceType_RESOURCE_TYPE_ENVIRONMENT || event.GetResourceId() != h.id {
				continue
			}
			h.client.sdk.logger().DebugContext(watchCtx, "environment event received",
				"operation", operation,
				"environment_id", h.id,
				"resource_operation", event.GetOperation().String(),
			)
			select {
			case changes <- struct{}{}:
			default:
			}
		}
		if err := stream.Err(); err != nil && !errors.Is(err, context.Canceled) {
			h.client.sdk.logger().ErrorContext(watchCtx, "environment informer event watch failed",
				"operation", operation,
				"environment_id", h.id,
				"err", err,
			)
			watchDone <- err
			return
		}
		h.client.sdk.logger().DebugContext(watchCtx, "environment event watch ended",
			"operation", operation,
			"environment_id", h.id,
		)
		watchDone <- nil
	}()

	informer.InvalidateItem(ctx, h.id)
	last, reached, err = observe(ctx)
	if err != nil || reached {
		return last, err
	}

	for {
		select {
		case <-ctx.Done():
			return last, MapError(operation, ctx.Err())
		case <-changes:
			informer.InvalidateItem(ctx, h.id)
			last, reached, err = observe(ctx)
			if err != nil || reached {
				return last, err
			}
		case watchErr := <-watchDone:
			for {
				select {
				case <-changes:
					informer.InvalidateItem(ctx, h.id)
					last, reached, err = observe(ctx)
					if err != nil || reached {
						return last, err
					}
				default:
					if watchErr != nil {
						return last, MapError(operation, watchErr)
					}
					return last, &UnavailableError{operationError{
						operation: operation,
						err:       fmt.Errorf("environment %s event stream ended before the target phase", h.id),
					}}
				}
			}
		}
	}
}

func (c *EnvironmentClient) environment(ctx context.Context, handle *environmentHandle, waitForRunning bool) (*Environment, error) {
	if waitForRunning {
		c.sdk.logger().DebugContext(ctx, "waiting for environment to run", "environment_id", handle.idString())
		if _, err := handle.waitRunning(ctx); err != nil {
			return nil, err
		}
	}

	env := &Environment{
		sdk:         c.sdk,
		environment: handle,
	}
	return env, nil
}

func (c *EnvironmentClient) requireSCMAuthentication(ctx context.Context, runnerID string, contextURL string) error {
	operation := "environments.require_scm_authentication"
	host, err := hostFromURL(contextURL)
	if err != nil {
		return &ValidationError{operationError{operation: operation, err: err}}
	}
	_, err = c.sdk.scm().checkAuthenticationForHost(ctx, checkAuthenticationOptions{
		runnerID: runnerID,
		host:     host,
	})
	return err
}

func (c *EnvironmentClient) environmentClassRunnerID(ctx context.Context, environmentClassID string) (string, error) {
	operation := "environments.environment_class_runner"
	raw, err := c.raw(operation)
	if err != nil {
		return "", err
	}
	if environmentClassID == "" {
		return "", &ValidationError{operationError{operation: operation, err: errors.New("environment class ID is required")}}
	}
	if raw.RunnerConfigurationService() == nil {
		return "", &CapabilityUnavailableError{operationError{operation: operation, err: errors.New("runner configuration service is not configured")}}
	}

	c.sdk.logger().DebugContext(ctx, "resolving runner for environment class",
		"operation", operation,
		"environment_class_id", environmentClassID,
	)
	resp, err := raw.RunnerConfigurationService().GetEnvironmentClass(ctx, connect.NewRequest(&v1.GetEnvironmentClassRequest{
		EnvironmentClassId: environmentClassID,
	}))
	if err != nil {
		return "", MapError(operation, err)
	}
	runnerID := resp.Msg.GetEnvironmentClass().GetRunnerId()
	if runnerID == "" {
		return "", &CapabilityUnavailableError{operationError{
			operation: operation,
			err:       fmt.Errorf("environment class %s did not include a runner ID", environmentClassID),
		}}
	}
	c.sdk.logger().DebugContext(ctx, "resolved runner for environment class",
		"operation", operation,
		"environment_class_id", environmentClassID,
		"runner_id", runnerID,
	)
	return runnerID, nil
}

func (c *EnvironmentClient) defaultEnvironmentClass(ctx context.Context) (string, string, error) {
	operation := "environments.default_class"
	raw, err := c.raw(operation)
	if err != nil {
		return "", "", err
	}

	filter := &v1.ListEnvironmentClassesRequest_Filter{
		Enabled:               Pointer(true),
		CanCreateEnvironments: Pointer(true),
	}

	c.sdk.logger().DebugContext(ctx, "listing environment classes for default selection", "operation", operation)
	resp, err := raw.EnvironmentService().ListEnvironmentClasses(ctx, connect.NewRequest(&v1.ListEnvironmentClassesRequest{
		Pagination: &v1.PaginationRequest{PageSize: c.sdk.config().pageSize},
		Filter:     filter,
	}))
	if err != nil {
		return "", "", MapError(operation, err)
	}
	for _, class := range resp.Msg.GetEnvironmentClasses() {
		if class.GetId() != "" {
			c.sdk.logger().DebugContext(ctx, "selected default environment class",
				"operation", operation,
				"environment_class_id", class.GetId(),
				"runner_id", class.GetRunnerId(),
			)
			return class.GetId(), class.GetRunnerId(), nil
		}
	}

	return "", "", &CapabilityUnavailableError{operationError{
		operation: operation,
		err:       errors.New("no environment class available for creating environments"),
	}}
}

func (c *EnvironmentClient) raw(operation string) (*rawclient.ManagementPlane, error) {
	if c == nil || c.sdk == nil {
		return nil, &ValidationError{operationError{operation: operation, err: errSDKClientMissing}}
	}
	return c.sdk.requireRaw(operation)
}

func (c *EnvironmentClient) handle(env *v1.Environment) *environmentHandle {
	return &environmentHandle{
		client:      c,
		id:          env.GetId(),
		environment: env,
	}
}

func environmentSpecFromContextURL(opts createEnvironmentFromContextURLOptions) *v1.EnvironmentSpec {
	spec := &v1.EnvironmentSpec{
		DesiredPhase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_RUNNING,
	}
	spec.Content = &v1.EnvironmentSpec_Content{
		Initializer: &v1.EnvironmentInitializer{
			Specs: []*v1.EnvironmentInitializer_Spec{{
				Spec: &v1.EnvironmentInitializer_Spec_ContextUrl{
					ContextUrl: &v1.ContextURLInitializer{Url: opts.ContextURL},
				},
			}},
		},
	}
	if opts.EnvironmentClass != "" {
		spec.Machine = &v1.EnvironmentSpec_Machine{Class: opts.EnvironmentClass}
	}
	return spec
}

func environmentSpecFromProject(opts createEnvironmentFromProjectOptions, gitRemoteURI string) *v1.EnvironmentSpec {
	spec := &v1.EnvironmentSpec{}
	hasSpec := false

	if initializer := gitInitializerFromParsedContext(opts.Resolved.GetGit(), gitRemoteURI); initializer != nil {
		spec.Content = &v1.EnvironmentSpec_Content{
			Initializer: &v1.EnvironmentInitializer{
				Specs: []*v1.EnvironmentInitializer_Spec{{
					Spec: &v1.EnvironmentInitializer_Spec_Git{Git: initializer},
				}},
			},
		}
		hasSpec = true
	}

	if !hasSpec {
		return nil
	}
	return spec
}

func gitInitializerFromParsedContext(gitCtx *v1.ParseContextURLResponse_GitContext, remoteURI string) *v1.GitInitializer {
	if gitCtx == nil {
		return nil
	}
	if remoteURI == "" {
		remoteURI = gitCtx.GetCloneUrl()
	}
	if remoteURI == "" {
		return nil
	}

	target := gitCtx.GetCommit()
	targetMode := v1.GitInitializer_CLONE_TARGET_MODE_REMOTE_COMMIT
	if target == "" {
		target = gitCtx.GetTag()
		targetMode = v1.GitInitializer_CLONE_TARGET_MODE_REMOTE_TAG
	}
	if target == "" {
		target = gitCtx.GetBranch()
		targetMode = v1.GitInitializer_CLONE_TARGET_MODE_REMOTE_BRANCH
	}
	if target == "" {
		return nil
	}

	return &v1.GitInitializer{
		RemoteUri:         remoteURI,
		TargetMode:        targetMode,
		CloneTarget:       target,
		CheckoutLocation:  gitCtx.GetRepo(),
		UpstreamRemoteUri: gitCtx.GetUpstreamRemoteUrl(),
	}
}

func projectRemoteURIForParsedContext(ctx context.Context, raw *rawclient.ManagementPlane, projectID string, gitCtx *v1.ParseContextURLResponse_GitContext) string {
	fallback := gitCtx.GetCloneUrl()
	if raw == nil || raw.ProjectService() == nil || projectID == "" || fallback == "" {
		return fallback
	}

	resp, err := raw.ProjectService().GetProject(ctx, connect.NewRequest(&v1.GetProjectRequest{ProjectId: projectID}))
	if err != nil {
		return fallback
	}

	target := normalizedRepositoryIdentifier(fallback)
	for _, spec := range resp.Msg.GetProject().GetInitializer().GetSpecs() {
		remoteURI := spec.GetGit().GetRemoteUri()
		if remoteURI == "" {
			continue
		}
		if normalizedRepositoryIdentifier(remoteURI) == target {
			return remoteURI
		}
	}
	return fallback
}

func normalizedRepositoryIdentifier(raw string) string {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimSuffix(strings.TrimRight(raw, "/"), ".git")

	if strings.HasPrefix(raw, "git@") {
		withoutPrefix := strings.TrimPrefix(raw, "git@")
		parts := strings.SplitN(withoutPrefix, ":", 2)
		if len(parts) == 2 {
			return strings.ToLower(parts[0]) + "/" + strings.TrimSuffix(strings.Trim(parts[1], "/"), ".git")
		}
	}

	parsed, err := url.Parse(raw)
	if err != nil || parsed.Host == "" {
		return raw
	}
	return strings.ToLower(parsed.Host) + "/" + strings.TrimSuffix(strings.Trim(parsed.Path, "/"), ".git")
}

func (c *EnvironmentClient) preferredProjectID(ctx context.Context, projectIDs []string) string {
	candidates := make(map[string]struct{}, len(projectIDs))
	var fallback string
	for _, projectID := range projectIDs {
		if projectID == "" {
			continue
		}
		if fallback == "" {
			fallback = projectID
		}
		candidates[projectID] = struct{}{}
	}
	if len(candidates) <= 1 {
		if fallback != "" {
			c.sdk.logger().DebugContext(ctx, "selected project for context URL",
				"operation", "environments.preferred_project",
				"project_id", fallback,
				"selection_source", "only_match",
				"candidate_count", len(candidates),
			)
		}
		return fallback
	}

	raw, err := c.raw("environments.preferred_project")
	if err != nil || raw.UsageService() == nil {
		c.sdk.logger().DebugContext(ctx, "selected project for context URL",
			"operation", "environments.preferred_project",
			"project_id", fallback,
			"selection_source", "first_match",
			"candidate_count", len(candidates),
			"err", err,
		)
		return fallback
	}

	cfg := c.sdk.config()
	now := time.Now()
	resp, err := raw.UsageService().GetTopProjects(ctx, connect.NewRequest(&v1.GetTopProjectsRequest{
		Pagination: &v1.PaginationRequest{PageSize: cfg.pageSize},
		DateRange: &v1.DateRange{
			StartTime: timestamppb.New(now.Add(-defaultProjectUsageLookback)),
			EndTime:   timestamppb.New(now),
		},
	}))
	if err != nil {
		c.sdk.logger().DebugContext(ctx, "selected project for context URL",
			"operation", "environments.preferred_project",
			"project_id", fallback,
			"selection_source", "first_match",
			"candidate_count", len(candidates),
			"err", err,
		)
		return fallback
	}

	for _, project := range resp.Msg.GetProjects() {
		if _, ok := candidates[project.GetProjectId()]; ok {
			c.sdk.logger().DebugContext(ctx, "selected project for context URL",
				"operation", "environments.preferred_project",
				"project_id", project.GetProjectId(),
				"selection_source", "usage",
				"candidate_count", len(candidates),
			)
			return project.GetProjectId()
		}
	}
	c.sdk.logger().DebugContext(ctx, "selected project for context URL",
		"operation", "environments.preferred_project",
		"project_id", fallback,
		"selection_source", "first_match",
		"candidate_count", len(candidates),
	)
	return fallback
}

func firstRecommendedEnvironmentClass(resolved *v1.ParseContextURLResponse) string {
	for _, classID := range resolved.GetRecommendedEnvironmentClasses() {
		if classID != "" {
			return classID
		}
	}
	return ""
}

func hostFromURL(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("parse context URL: %w", err)
	}
	if parsed.Host == "" {
		return "", errors.New("context URL must include a host")
	}
	return parsed.Host, nil
}

func environmentFailureError(operation string, env *v1.Environment) error {
	failures := env.GetStatus().GetFailureMessage()
	if len(failures) == 0 {
		return nil
	}
	return &EnvironmentPolicyError{operationError{
		operation: operation,
		err:       fmt.Errorf("environment %s failed: %v", env.GetId(), failures),
	}}
}

func environmentWatchRequest(environmentID string) *v1.WatchEventsRequest {
	return &v1.WatchEventsRequest{
		Scope: &v1.WatchEventsRequest_Organization{Organization: true},
		ResourceTypeFilters: []*v1.WatchEventsRequest_ResourceTypeFilter{{
			ResourceType: v1.ResourceType_RESOURCE_TYPE_ENVIRONMENT,
			ResourceIds:  []string{environmentID},
		}},
	}
}

func isTerminalBeforeRunning(phase v1.EnvironmentPhase) bool {
	switch phase {
	case v1.EnvironmentPhase_ENVIRONMENT_PHASE_STOPPED,
		v1.EnvironmentPhase_ENVIRONMENT_PHASE_DELETED:
		return true
	default:
		return false
	}
}

func environmentPhaseLabel(phase v1.EnvironmentPhase) string {
	switch phase {
	case v1.EnvironmentPhase_ENVIRONMENT_PHASE_STARTING:
		return "starting"
	case v1.EnvironmentPhase_ENVIRONMENT_PHASE_RUNNING:
		return "running"
	case v1.EnvironmentPhase_ENVIRONMENT_PHASE_STOPPING:
		return "stopping"
	case v1.EnvironmentPhase_ENVIRONMENT_PHASE_STOPPED:
		return "stopped"
	case v1.EnvironmentPhase_ENVIRONMENT_PHASE_DELETING:
		return "deleting"
	case v1.EnvironmentPhase_ENVIRONMENT_PHASE_DELETED:
		return "deleted"
	default:
		return phase.String()
	}
}

func optionalString(value string) *string {
	if value == "" {
		return nil
	}
	return Pointer(value)
}
