package cache

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	v1 "github.com/gitpod-io/flex-sdk-go/v1"
	"github.com/gitpod-io/flex-sdk-go/v1/v1connect"
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
