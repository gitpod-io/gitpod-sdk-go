package client

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"go.uber.org/mock/gomock"
	"golang.org/x/oauth2"

	"github.com/gitpod-io/gitpod-sdk-go/internal/raw/mock"
	"github.com/gitpod-io/gitpod-sdk-go/v1/v1connect"
)

const (
	DefaultBaseURL = "https://app.ona.com/api"
	APIKeyEnvVar   = "ONA_API_KEY"
)

// ErrMissingAPIKey is returned by environment-based constructors when ONA_API_KEY is not set.
var ErrMissingAPIKey = errors.New("ONA_API_KEY is required")

type options struct {
	baseURL      string
	tokenSource  oauth2.TokenSource
	userAgent    string
	httpClient   *http.Client
	interceptors []connect.Interceptor
}

type Option func(*options) error

// WithUserAgent sets an application-specific User-Agent prefix.
//
// The client always appends its own telemetry token so API requests identify
// the SDK language, version, and whether the request came from the raw client
// or the higher-level SDK layer.
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

// WithMetrics adds a metrics interceptor that counts requests per procedure and status code.
// Create APICallMetrics once via NewAPICallMetrics and share across multiple clients.
func WithMetrics(m *APICallMetrics) Option {
	return func(o *options) error {
		o.interceptors = append(o.interceptors, m.Interceptor())
		return nil
	}
}

func WithTokenSource(source oauth2.TokenSource) Option {
	return func(o *options) error {
		o.tokenSource = source
		return nil
	}
}

func WithAccessToken(token string) Option {
	return func(o *options) error {
		o.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		return nil
	}
}

type ManagementPlaneMock struct {
	AccountService               *mock.MockAccountServiceClient
	AgentService                 *mock.MockAgentServiceClient
	BillingService               *mock.MockBillingServiceClient
	EditorService                *mock.MockEditorServiceClient
	EnvironmentAutomationService *mock.MockEnvironmentAutomationServiceClient
	EnvironmentService           *mock.MockEnvironmentServiceClient
	EventService                 *mock.MockEventServiceClient
	GroupService                 *mock.MockGroupServiceClient
	IdentityService              *mock.MockIdentityServiceClient
	InsightsService              *mock.MockInsightsServiceClient
	IntegrationService           *mock.MockIntegrationServiceClient
	NotificationService          *mock.MockNotificationServiceClient
	OrganizationService          *mock.MockOrganizationServiceClient
	PrebuildService              *mock.MockPrebuildServiceClient
	ProjectService               *mock.MockProjectServiceClient
	RunnerConfigurationService   *mock.MockRunnerConfigurationServiceClient
	RunnerService                *mock.MockRunnerServiceClient
	UserService                  *mock.MockUserServiceClient
	UsageService                 *mock.MockUsageServiceClient
	SecurityService              *mock.MockSecurityServiceClient
	SecretService                *mock.MockSecretServiceClient
	ServiceAccountService        *mock.MockServiceAccountServiceClient
	TeamService                  *mock.MockTeamServiceClient
	GatewayService               *mock.MockGatewayServiceClient
	WorkflowService              *mock.MockWorkflowServiceClient
	WebhookService               *mock.MockWebhookServiceClient
}

// Client returns a client for the control plane API
func (m *ManagementPlaneMock) Client() *ManagementPlane {
	return &ManagementPlane{
		accountService:               m.AccountService,
		agentService:                 m.AgentService,
		billingService:               m.BillingService,
		editorService:                m.EditorService,
		environmentAutomationService: m.EnvironmentAutomationService,
		environmentService:           m.EnvironmentService,
		eventService:                 m.EventService,
		groupService:                 m.GroupService,
		identityService:              m.IdentityService,
		insightsService:              m.InsightsService,
		integrationService:           m.IntegrationService,
		notificationService:          m.NotificationService,
		organizationService:          m.OrganizationService,
		prebuildService:              m.PrebuildService,
		projectService:               m.ProjectService,
		runnerConfigurationService:   m.RunnerConfigurationService,
		runnerService:                m.RunnerService,
		userService:                  m.UserService,
		usageService:                 m.UsageService,
		securityService:              m.SecurityService,
		secretService:                m.SecretService,
		serviceAccountService:        m.ServiceAccountService,
		teamService:                  m.TeamService,
		gatewayService:               m.GatewayService,
		workflowService:              m.WorkflowService,
		webhookService:               m.WebhookService,
	}
}

// NewMock creates a new mock for the control plane API
func NewMock(ctrl *gomock.Controller) *ManagementPlaneMock {
	return &ManagementPlaneMock{
		AccountService:               mock.NewMockAccountServiceClient(ctrl),
		AgentService:                 mock.NewMockAgentServiceClient(ctrl),
		BillingService:               mock.NewMockBillingServiceClient(ctrl),
		EditorService:                mock.NewMockEditorServiceClient(ctrl),
		EnvironmentAutomationService: mock.NewMockEnvironmentAutomationServiceClient(ctrl),
		EnvironmentService:           mock.NewMockEnvironmentServiceClient(ctrl),
		EventService:                 mock.NewMockEventServiceClient(ctrl),
		GroupService:                 mock.NewMockGroupServiceClient(ctrl),
		IdentityService:              mock.NewMockIdentityServiceClient(ctrl),
		InsightsService:              mock.NewMockInsightsServiceClient(ctrl),
		IntegrationService:           mock.NewMockIntegrationServiceClient(ctrl),
		NotificationService:          mock.NewMockNotificationServiceClient(ctrl),
		OrganizationService:          mock.NewMockOrganizationServiceClient(ctrl),
		PrebuildService:              mock.NewMockPrebuildServiceClient(ctrl),
		ProjectService:               mock.NewMockProjectServiceClient(ctrl),
		RunnerConfigurationService:   mock.NewMockRunnerConfigurationServiceClient(ctrl),
		RunnerService:                mock.NewMockRunnerServiceClient(ctrl),
		UserService:                  mock.NewMockUserServiceClient(ctrl),
		UsageService:                 mock.NewMockUsageServiceClient(ctrl),
		SecurityService:              mock.NewMockSecurityServiceClient(ctrl),
		SecretService:                mock.NewMockSecretServiceClient(ctrl),
		ServiceAccountService:        mock.NewMockServiceAccountServiceClient(ctrl),
		TeamService:                  mock.NewMockTeamServiceClient(ctrl),
		GatewayService:               mock.NewMockGatewayServiceClient(ctrl),
		WorkflowService:              mock.NewMockWorkflowServiceClient(ctrl),
		WebhookService:               mock.NewMockWebhookServiceClient(ctrl),
	}
}

// NewFromEnv creates a management-plane client for the production API using ONA_API_KEY.
func NewFromEnv(opts ...Option) (*ManagementPlane, error) {
	token := strings.TrimSpace(os.Getenv(APIKeyEnvVar))
	if token == "" {
		return nil, ErrMissingAPIKey
	}
	return New(DefaultBaseURL, append([]Option{WithAccessToken(token)}, opts...)...)
}

func New(baseURL string, opts ...Option) (*ManagementPlane, error) {
	o := options{
		httpClient: http.DefaultClient,
		baseURL:    baseURL,
	}
	for _, opt := range opts {
		err := opt(&o)
		if err != nil {
			return nil, fmt.Errorf("cannot apply option: %w", err)
		}
	}
	o.httpClient = httpClientPreservingAuthorization(o.httpClient, o.tokenSource)

	interceptors := append([]connect.Interceptor{
		TokenSourceInterceptor(o.tokenSource),
	}, o.interceptors...)
	interceptors = append(interceptors, withOnaUserAgent(o.userAgent))

	clientOpts := []connect.ClientOption{
		connect.WithInterceptors(interceptors...),
	}

	return &ManagementPlane{
		accountService:               v1connect.NewAccountServiceClient(o.httpClient, o.baseURL, clientOpts...),
		agentService:                 v1connect.NewAgentServiceClient(o.httpClient, o.baseURL, clientOpts...),
		billingService:               v1connect.NewBillingServiceClient(o.httpClient, o.baseURL, clientOpts...),
		editorService:                v1connect.NewEditorServiceClient(o.httpClient, o.baseURL, clientOpts...),
		environmentAutomationService: v1connect.NewEnvironmentAutomationServiceClient(o.httpClient, o.baseURL, clientOpts...),
		environmentService:           v1connect.NewEnvironmentServiceClient(o.httpClient, o.baseURL, clientOpts...),
		eventService:                 v1connect.NewEventServiceClient(o.httpClient, o.baseURL, clientOpts...),
		groupService:                 v1connect.NewGroupServiceClient(o.httpClient, o.baseURL, clientOpts...),
		identityService:              v1connect.NewIdentityServiceClient(o.httpClient, o.baseURL, clientOpts...),
		insightsService:              v1connect.NewInsightsServiceClient(o.httpClient, o.baseURL, clientOpts...),
		integrationService:           v1connect.NewIntegrationServiceClient(o.httpClient, o.baseURL, clientOpts...),
		notificationService:          v1connect.NewNotificationServiceClient(o.httpClient, o.baseURL, clientOpts...),
		organizationService:          v1connect.NewOrganizationServiceClient(o.httpClient, o.baseURL, clientOpts...),
		prebuildService:              v1connect.NewPrebuildServiceClient(o.httpClient, o.baseURL, clientOpts...),
		projectService:               v1connect.NewProjectServiceClient(o.httpClient, o.baseURL, clientOpts...),
		runnerConfigurationService:   v1connect.NewRunnerConfigurationServiceClient(o.httpClient, o.baseURL, clientOpts...),
		runnerService:                v1connect.NewRunnerServiceClient(o.httpClient, o.baseURL, clientOpts...),
		userService:                  v1connect.NewUserServiceClient(o.httpClient, o.baseURL, clientOpts...),
		usageService:                 v1connect.NewUsageServiceClient(o.httpClient, o.baseURL, clientOpts...),
		securityService:              v1connect.NewSecurityServiceClient(o.httpClient, o.baseURL, clientOpts...),
		secretService:                v1connect.NewSecretServiceClient(o.httpClient, o.baseURL, clientOpts...),
		serviceAccountService:        v1connect.NewServiceAccountServiceClient(o.httpClient, o.baseURL, clientOpts...),
		teamService:                  v1connect.NewTeamServiceClient(o.httpClient, o.baseURL, clientOpts...),
		gatewayService:               v1connect.NewGatewayServiceClient(o.httpClient, o.baseURL, clientOpts...),
		workflowService:              v1connect.NewWorkflowServiceClient(o.httpClient, o.baseURL, clientOpts...),
		webhookService:               v1connect.NewWebhookServiceClient(o.httpClient, o.baseURL, clientOpts...),
	}, nil
}

func httpClientPreservingAuthorization(client *http.Client, source oauth2.TokenSource) *http.Client {
	redirectClient := *client
	transport := client.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	redirectClient.Transport = authorizationRoundTripper{base: transport, source: source}
	return &redirectClient
}

type authorizationRoundTripper struct {
	base   http.RoundTripper
	source oauth2.TokenSource
}

func (transport authorizationRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	// net/http removes authorization on cross-host redirects before the redirected
	// request reaches the transport, so obtain it again from the configured source.
	if request.Header.Get("Authorization") != "" || transport.source == nil {
		return transport.base.RoundTrip(request)
	}
	token, err := transport.source.Token()
	if err != nil || token.AccessToken == "" {
		return transport.base.RoundTrip(request)
	}
	request = request.Clone(request.Context())
	request.Header = request.Header.Clone()
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", BearerPrefix, token.AccessToken))
	return transport.base.RoundTrip(request)
}

type ManagementPlane struct {
	accountService               v1connect.AccountServiceClient
	agentService                 v1connect.AgentServiceClient
	billingService               v1connect.BillingServiceClient
	editorService                v1connect.EditorServiceClient
	environmentAutomationService v1connect.EnvironmentAutomationServiceClient
	environmentService           v1connect.EnvironmentServiceClient
	eventService                 v1connect.EventServiceClient
	groupService                 v1connect.GroupServiceClient
	identityService              v1connect.IdentityServiceClient
	insightsService              v1connect.InsightsServiceClient
	integrationService           v1connect.IntegrationServiceClient
	notificationService          v1connect.NotificationServiceClient
	organizationService          v1connect.OrganizationServiceClient
	prebuildService              v1connect.PrebuildServiceClient
	projectService               v1connect.ProjectServiceClient
	runnerConfigurationService   v1connect.RunnerConfigurationServiceClient
	runnerService                v1connect.RunnerServiceClient
	userService                  v1connect.UserServiceClient
	usageService                 v1connect.UsageServiceClient
	securityService              v1connect.SecurityServiceClient
	secretService                v1connect.SecretServiceClient
	serviceAccountService        v1connect.ServiceAccountServiceClient
	teamService                  v1connect.TeamServiceClient
	gatewayService               v1connect.GatewayServiceClient
	workflowService              v1connect.WorkflowServiceClient
	webhookService               v1connect.WebhookServiceClient
}

func (g *ManagementPlane) AccountService() v1connect.AccountServiceClient {
	return g.accountService
}

func (g *ManagementPlane) AgentService() v1connect.AgentServiceClient {
	return g.agentService
}

func (g *ManagementPlane) BillingService() v1connect.BillingServiceClient {
	return g.billingService
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

func (g *ManagementPlane) GroupService() v1connect.GroupServiceClient {
	return g.groupService
}

func (g *ManagementPlane) IdentityService() v1connect.IdentityServiceClient {
	return g.identityService
}

func (g *ManagementPlane) InsightsService() v1connect.InsightsServiceClient {
	return g.insightsService
}

func (g *ManagementPlane) IntegrationService() v1connect.IntegrationServiceClient {
	return g.integrationService
}

func (g *ManagementPlane) NotificationService() v1connect.NotificationServiceClient {
	return g.notificationService
}

func (g *ManagementPlane) OrganizationService() v1connect.OrganizationServiceClient {
	return g.organizationService
}

func (g *ManagementPlane) PrebuildService() v1connect.PrebuildServiceClient {
	return g.prebuildService
}

func (g *ManagementPlane) ProjectService() v1connect.ProjectServiceClient {
	return g.projectService
}

func (g *ManagementPlane) RunnerConfigurationService() v1connect.RunnerConfigurationServiceClient {
	return g.runnerConfigurationService
}

func (g *ManagementPlane) RunnerService() v1connect.RunnerServiceClient {
	return g.runnerService
}

func (g *ManagementPlane) UserService() v1connect.UserServiceClient {
	return g.userService
}

func (g *ManagementPlane) UsageService() v1connect.UsageServiceClient {
	return g.usageService
}

func (g *ManagementPlane) SecurityService() v1connect.SecurityServiceClient {
	return g.securityService
}

func (g *ManagementPlane) SecretService() v1connect.SecretServiceClient {
	return g.secretService
}

func (g *ManagementPlane) ServiceAccountService() v1connect.ServiceAccountServiceClient {
	return g.serviceAccountService
}

func (g *ManagementPlane) TeamService() v1connect.TeamServiceClient {
	return g.teamService
}

func (g *ManagementPlane) GatewayService() v1connect.GatewayServiceClient {
	return g.gatewayService
}

func (g *ManagementPlane) WorkflowService() v1connect.WorkflowServiceClient {
	return g.workflowService
}

func (g *ManagementPlane) WebhookService() v1connect.WebhookServiceClient {
	return g.webhookService
}
