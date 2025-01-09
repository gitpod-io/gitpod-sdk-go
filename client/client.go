package client

import (
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"go.uber.org/mock/gomock"

	"github.com/gitpod-io/flex-go/client/mock"
	"github.com/gitpod-io/flex-go/v1/v1connect"
)

type options struct {
	baseURL      string
	bearerToken  string
	userAgent    string
	httpClient   *http.Client
	interceptors []connect.Interceptor
}

type Option func(*options) error

func WithUserAgent(userAgent string) Option {
	return func(o *options) error {
		o.userAgent = userAgent
		return nil
	}
}

func WithTracing(opts ...otelconnect.Option) Option {
	return func(o *options) error {
		interceptor, err := otelconnect.NewInterceptor(opts...)
		if err != nil {
			return fmt.Errorf("cannot create tracing interceptor: %w", err)
		}

		o.interceptors = append(o.interceptors, interceptor)

		return nil
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(o *options) error {
		o.httpClient = client
		return nil
	}
}

func WithInterceptor(interceptor connect.Interceptor) Option {
	return func(o *options) error {
		o.interceptors = append(o.interceptors, interceptor)
		return nil
	}
}

type ManagementPlaneMock struct {
	AccountService               *mock.MockAccountServiceClient
	EditorService                *mock.MockEditorServiceClient
	EnvironmentAutomationService *mock.MockEnvironmentAutomationServiceClient
	EnvironmentService           *mock.MockEnvironmentServiceClient
	EventService                 *mock.MockEventServiceClient
	IdentityService              *mock.MockIdentityServiceClient
	OrganizationService          *mock.MockOrganizationServiceClient
	ProjectService               *mock.MockProjectServiceClient
	RunnerConfigurationService   *mock.MockRunnerConfigurationServiceClient
	RunnerInteractionService     *mock.MockRunnerInteractionServiceClient
	RunnerService                *mock.MockRunnerServiceClient
	UserService                  *mock.MockUserServiceClient
	SecretService                *mock.MockSecretServiceClient
	LLMService                   *mock.MockLLMServiceClient
}

// Client returns a client for the control plane API
func (m *ManagementPlaneMock) Client() *ManagementPlane {
	return &ManagementPlane{
		accountService:               m.AccountService,
		editorService:                m.EditorService,
		environmentAutomationService: m.EnvironmentAutomationService,
		environmentService:           m.EnvironmentService,
		eventService:                 m.EventService,
		identityService:              m.IdentityService,
		organizationService:          m.OrganizationService,
		projectService:               m.ProjectService,
		runnerConfigurationService:   m.RunnerConfigurationService,
		runnerInteractionService:     m.RunnerInteractionService,
		runnerService:                m.RunnerService,
		userService:                  m.UserService,
		secretService:                m.SecretService,
		llmService:                   m.LLMService,
	}
}

// NewMock creates a new mock for the control plane API
func NewMock(ctrl *gomock.Controller) *ManagementPlaneMock {
	return &ManagementPlaneMock{
		AccountService:               mock.NewMockAccountServiceClient(ctrl),
		EditorService:                mock.NewMockEditorServiceClient(ctrl),
		EnvironmentAutomationService: mock.NewMockEnvironmentAutomationServiceClient(ctrl),
		EnvironmentService:           mock.NewMockEnvironmentServiceClient(ctrl),
		EventService:                 mock.NewMockEventServiceClient(ctrl),
		IdentityService:              mock.NewMockIdentityServiceClient(ctrl),
		OrganizationService:          mock.NewMockOrganizationServiceClient(ctrl),
		ProjectService:               mock.NewMockProjectServiceClient(ctrl),
		RunnerConfigurationService:   mock.NewMockRunnerConfigurationServiceClient(ctrl),
		RunnerInteractionService:     mock.NewMockRunnerInteractionServiceClient(ctrl),
		RunnerService:                mock.NewMockRunnerServiceClient(ctrl),
		UserService:                  mock.NewMockUserServiceClient(ctrl),
		SecretService:                mock.NewMockSecretServiceClient(ctrl),
		LLMService:                   mock.NewMockLLMServiceClient(ctrl),
	}
}

func New(baseURL, token string, opts ...Option) (*ManagementPlane, error) {
	o := options{
		httpClient:  http.DefaultClient,
		baseURL:     baseURL,
		bearerToken: token,
	}
	for _, opt := range opts {
		err := opt(&o)
		if err != nil {
			return nil, fmt.Errorf("cannot apply option: %w", err)
		}
	}

	interceptors := append([]connect.Interceptor{
		WithBearerToken(o.bearerToken),
		WithCustomUserAgent(o.userAgent),
	}, o.interceptors...)

	clientOpts := []connect.ClientOption{
		connect.WithInterceptors(interceptors...),
		// Use the binary gRPC protocol
		connect.WithGRPC(),
	}

	return &ManagementPlane{
		accountService:               v1connect.NewAccountServiceClient(o.httpClient, o.baseURL, clientOpts...),
		editorService:                v1connect.NewEditorServiceClient(o.httpClient, o.baseURL, clientOpts...),
		environmentAutomationService: v1connect.NewEnvironmentAutomationServiceClient(o.httpClient, o.baseURL, clientOpts...),
		environmentService:           v1connect.NewEnvironmentServiceClient(o.httpClient, o.baseURL, clientOpts...),
		eventService:                 v1connect.NewEventServiceClient(o.httpClient, o.baseURL, clientOpts...),
		identityService:              v1connect.NewIdentityServiceClient(o.httpClient, o.baseURL, clientOpts...),
		organizationService:          v1connect.NewOrganizationServiceClient(o.httpClient, o.baseURL, clientOpts...),
		projectService:               v1connect.NewProjectServiceClient(o.httpClient, o.baseURL, clientOpts...),
		runnerConfigurationService:   v1connect.NewRunnerConfigurationServiceClient(o.httpClient, o.baseURL, clientOpts...),
		runnerInteractionService:     v1connect.NewRunnerInteractionServiceClient(o.httpClient, o.baseURL, clientOpts...),
		runnerService:                v1connect.NewRunnerServiceClient(o.httpClient, o.baseURL, clientOpts...),
		userService:                  v1connect.NewUserServiceClient(o.httpClient, o.baseURL, clientOpts...),
		secretService:                v1connect.NewSecretServiceClient(o.httpClient, o.baseURL, clientOpts...),
		llmService:                   v1connect.NewLLMServiceClient(o.httpClient, o.baseURL, clientOpts...),
	}, nil
}

type ManagementPlane struct {
	accountService               v1connect.AccountServiceClient
	editorService                v1connect.EditorServiceClient
	environmentAutomationService v1connect.EnvironmentAutomationServiceClient
	environmentService           v1connect.EnvironmentServiceClient
	eventService                 v1connect.EventServiceClient
	identityService              v1connect.IdentityServiceClient
	organizationService          v1connect.OrganizationServiceClient
	projectService               v1connect.ProjectServiceClient
	runnerConfigurationService   v1connect.RunnerConfigurationServiceClient
	runnerInteractionService     v1connect.RunnerInteractionServiceClient
	runnerService                v1connect.RunnerServiceClient
	userService                  v1connect.UserServiceClient
	secretService                v1connect.SecretServiceClient
	llmService                   v1connect.LLMServiceClient
}

func (g *ManagementPlane) AccountService() v1connect.AccountServiceClient {
	return g.accountService
}

func (g *ManagementPlane) EditorService() v1connect.EditorServiceClient {
	return g.editorService
}

func (g *ManagementPlane) EnvironmentAutomationService() v1connect.EnvironmentAutomationServiceClient {
	return g.environmentAutomationService
}

func (g *ManagementPlane) EnvironmentService() v1connect.EnvironmentServiceClient {
	return g.environmentService
}

func (g *ManagementPlane) EventService() v1connect.EventServiceClient {
	return g.eventService
}

func (g *ManagementPlane) IdentityService() v1connect.IdentityServiceClient {
	return g.identityService
}

func (g *ManagementPlane) OrganizationService() v1connect.OrganizationServiceClient {
	return g.organizationService
}

func (g *ManagementPlane) ProjectService() v1connect.ProjectServiceClient {
	return g.projectService
}

func (g *ManagementPlane) RunnerConfigurationService() v1connect.RunnerConfigurationServiceClient {
	return g.runnerConfigurationService
}

func (g *ManagementPlane) RunnerInteractionService() v1connect.RunnerInteractionServiceClient {
	return g.runnerInteractionService
}

func (g *ManagementPlane) RunnerService() v1connect.RunnerServiceClient {
	return g.runnerService
}

func (g *ManagementPlane) UserService() v1connect.UserServiceClient {
	return g.userService
}

func (g *ManagementPlane) SecretService() v1connect.SecretServiceClient {
	return g.secretService
}

func (g *ManagementPlane) LLMService() v1connect.LLMServiceClient {
	return g.llmService
}
