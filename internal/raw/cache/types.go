package cache

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
	"github.com/gitpod-io/gitpod-sdk-go/v1/v1connect"
	"google.golang.org/protobuf/proto"
)

// serviceRequester is a helper struct that provides methods to get and list Service items.
type serviceRequester struct {
	Client v1connect.EnvironmentAutomationServiceClient
	Filter *v1.ListServicesRequest_Filter
}

func (s *serviceRequester) Get(ctx context.Context, key string) (*v1.Service, error) {
	resp, err := s.Client.GetService(ctx, connect.NewRequest(&v1.GetServiceRequest{
		Id: key,
	}))
	if err != nil {
		return nil, err
	}
	return resp.Msg.Service, nil
}

// List retrieves a list of Service items with pagination.
func (s *serviceRequester) List(ctx context.Context, page *v1.PaginationRequest, ids []string) ([]*v1.Service, *v1.PaginationResponse, error) {
	filter := proto.Clone(s.Filter).(*v1.ListServicesRequest_Filter)
	if len(ids) > 0 {
		filter.ServiceIds = ids
	}

	resp, err := s.Client.ListServices(ctx, connect.NewRequest(&v1.ListServicesRequest{
		Pagination: page,
		Filter:     filter,
	}))
	if err != nil {
		return nil, nil, err
	}
	return resp.Msg.Services, resp.Msg.Pagination, nil
}

// NewServiceCache creates a new ResourceCache for v1.Service items. It uses the provided
// EnvironmentAutomationServiceClient and ListServicesRequest_Filter to retrieve the
// service items. The ResourceCache can be configured with the provided options. The context is only used for error logging.
func NewServiceCache(ctx context.Context, apiclnt v1connect.EnvironmentAutomationServiceClient, filter *v1.ListServicesRequest_Filter, opts ...ResourceCacheOption) (*ResourceCache[*v1.Service], error) {
	return NewResourceCache(
		ctx,
		UseRequester(&serviceRequester{
			Client: apiclnt,
			Filter: filter,
		}),
		opts...,
	)
}

// taskExecutionRequester is a helper struct that provides methods to get and list TaskExecution items.
type taskExecutionRequester struct {
	Client v1connect.EnvironmentAutomationServiceClient
	Filter *v1.ListTaskExecutionsRequest_Filter
}

func (t *taskExecutionRequester) Get(ctx context.Context, key string) (*v1.TaskExecution, error) {
	resp, err := t.Client.GetTaskExecution(ctx, connect.NewRequest(&v1.GetTaskExecutionRequest{
		Id: key,
	}))
	if err != nil {
		return nil, err
	}
	return resp.Msg.TaskExecution, nil
}

// List retrieves a list of TaskExecution items with pagination.
func (t *taskExecutionRequester) List(ctx context.Context, page *v1.PaginationRequest, ids []string) ([]*v1.TaskExecution, *v1.PaginationResponse, error) {
	if len(ids) > 0 {
		slog.DebugContext(ctx, "cannot list task executions by id - returning everything")
	}

	resp, err := t.Client.ListTaskExecutions(ctx, connect.NewRequest(&v1.ListTaskExecutionsRequest{
		Pagination: page,
		Filter:     t.Filter,
	}))
	if err != nil {
		return nil, nil, err
	}
	return resp.Msg.TaskExecutions, resp.Msg.Pagination, nil
}

// NewTaskExecutionCache creates a new ResourceCache for v1.TaskExecution items. It uses the provided
// EnvironmentAutomationServiceClient and ListTaskExecutionsRequest_Filter to retrieve the
// task execution items. The ResourceCache can be configured with the provided options. The context is only used for error logging.
func NewTaskExecutionCache(ctx context.Context, apiclnt v1connect.EnvironmentAutomationServiceClient, filter *v1.ListTaskExecutionsRequest_Filter, opts ...ResourceCacheOption) (*ResourceCache[*v1.TaskExecution], error) {
	return NewResourceCache(
		ctx,
		UseRequester(&taskExecutionRequester{
			Client: apiclnt,
			Filter: filter,
		}),
		opts...,
	)
}

// taskRequester is a helper struct that provides methods to get and list Task items.
type taskRequester struct {
	Client v1connect.EnvironmentAutomationServiceClient
	Filter *v1.ListTasksRequest_Filter
}

func (t *taskRequester) Get(ctx context.Context, key string) (*v1.Task, error) {
	resp, err := t.Client.GetTask(ctx, connect.NewRequest(&v1.GetTaskRequest{
		Id: key,
	}))
	if err != nil {
		return nil, err
	}
	return resp.Msg.Task, nil
}

// List retrieves a list of Task items with pagination.
func (t *taskRequester) List(ctx context.Context, page *v1.PaginationRequest, ids []string) ([]*v1.Task, *v1.PaginationResponse, error) {
	filter := proto.Clone(t.Filter).(*v1.ListTasksRequest_Filter)
	if len(ids) > 0 {
		filter.TaskIds = ids
	}

	resp, err := t.Client.ListTasks(ctx, connect.NewRequest(&v1.ListTasksRequest{
		Pagination: page,
		Filter:     filter,
	}))
	if err != nil {
		return nil, nil, err
	}
	return resp.Msg.Tasks, resp.Msg.Pagination, nil
}

// NewTaskCache creates a new ResourceCache for v1.Task items. It uses the provided
// EnvironmentAutomationServiceClient and ListTasksRequest_Filter to retrieve the
// task items. The ResourceCache can be configured with the provided options. The context is only used for error logging.
func NewTaskCache(ctx context.Context, apiclnt v1connect.EnvironmentAutomationServiceClient, filter *v1.ListTasksRequest_Filter, opts ...ResourceCacheOption) (*ResourceCache[*v1.Task], error) {
	return NewResourceCache(
		ctx,
		UseRequester(&taskRequester{
			Client: apiclnt,
			Filter: filter,
		}),
		opts...,
	)
}

// agentExecutionRequester is a helper struct that provides methods to get and list AgentExecution items.
type agentExecutionRequester struct {
	Client           v1connect.AgentServiceClient
	AgentExecutionID string
}

func (a *agentExecutionRequester) Get(ctx context.Context, key string) (*v1.AgentExecution, error) {
	resp, err := a.Client.GetAgentExecution(ctx, connect.NewRequest(&v1.GetAgentExecutionRequest{
		AgentExecutionId: key,
	}))
	if err != nil {
		return nil, err
	}
	return resp.Msg.AgentExecution, nil
}

func (a *agentExecutionRequester) List(ctx context.Context, page *v1.PaginationRequest, ids []string) ([]*v1.AgentExecution, *v1.PaginationResponse, error) {
	if len(ids) == 0 && a.AgentExecutionID != "" {
		ids = []string{a.AgentExecutionID}
	}

	resp, err := a.Client.ListAgentExecutions(ctx, connect.NewRequest(&v1.ListAgentExecutionsRequest{
		Pagination: page,
		Filter: &v1.ListAgentExecutionsRequest_Filter{
			AgentExecutionIds: ids,
		},
	}))
	if err != nil {
		return nil, nil, err
	}
	return resp.Msg.AgentExecutions, resp.Msg.Pagination, nil
}

// NewAgentExecutionCache creates a new ResourceCache for a single v1.AgentExecution item.
// It uses GetAgentExecution for cache misses and ListAgentExecutions filtered by ID for refetches and full syncs.
func NewAgentExecutionCache(ctx context.Context, apiclnt v1connect.AgentServiceClient, agentExecutionID string, opts ...ResourceCacheOption) (*ResourceCache[*v1.AgentExecution], error) {
	return NewResourceCache(
		ctx,
		UseRequester(&agentExecutionRequester{
			Client:           apiclnt,
			AgentExecutionID: agentExecutionID,
		}),
		opts...,
	)
}

// environmentRequester is a helper struct that provides methods to get and list Environment items.
type environmentRequester struct {
	Client        v1connect.EnvironmentServiceClient
	EnvironmentID string
}

func (e *environmentRequester) Get(ctx context.Context, key string) (*v1.Environment, error) {
	resp, err := e.Client.GetEnvironment(ctx, connect.NewRequest(&v1.GetEnvironmentRequest{
		EnvironmentId: key,
	}))
	if err != nil {
		return nil, err
	}
	return resp.Msg.Environment, nil
}

// List retrieves a list of Environment items with pagination.
func (e *environmentRequester) List(ctx context.Context, page *v1.PaginationRequest, ids []string) ([]*v1.Environment, *v1.PaginationResponse, error) {
	// For environment cache, we only cache a single environment, so we just get it directly
	if len(ids) > 0 {
		// If specific IDs are requested, only return the one we care about if it's in the list
		for _, id := range ids {
			if id == e.EnvironmentID {
				env, err := e.Get(ctx, id)
				if err != nil {
					return nil, nil, err
				}
				return []*v1.Environment{env}, &v1.PaginationResponse{}, nil
			}
		}
		return []*v1.Environment{}, &v1.PaginationResponse{}, nil
	}

	// Get our single environment
	env, err := e.Get(ctx, e.EnvironmentID)
	if err != nil {
		return nil, nil, err
	}
	return []*v1.Environment{env}, &v1.PaginationResponse{}, nil
}

// NewEnvironmentCache creates a new ResourceCache for a single v1.Environment item. It uses the provided
// EnvironmentServiceClient to retrieve the environment item. The ResourceCache can be configured with the provided options.
func NewEnvironmentCache(ctx context.Context, apiclnt v1connect.EnvironmentServiceClient, environmentID string, opts ...ResourceCacheOption) (*ResourceCache[*v1.Environment], error) {
	return NewResourceCache(
		ctx,
		UseRequester(&environmentRequester{
			Client:        apiclnt,
			EnvironmentID: environmentID,
		}),
		opts...,
	)
}
