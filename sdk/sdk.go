package sdk

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	rawclient "github.com/gitpod-io/gitpod-sdk-go/internal/raw"
	"github.com/gitpod-io/gitpod-sdk-go/v1/v1connect"
)

const (
	APIKeyEnvVar        = "ONA_API_KEY"
	LegacyAPIKeyEnvVar  = "GITPOD_API_KEY"
	BaseURLEnvVar       = "ONA_BASE_URL"
	DefaultBaseURL      = "https://app.ona.com/api"
	defaultPollInterval = 2 * time.Second
	defaultPageSize     = int32(100)
)

// Client exposes task-oriented Ona workflows on top of the raw API client.
type Client struct {
	Services ServiceClients
	raw      *rawclient.ManagementPlane
	cfg      config
}

// ServiceClients exposes the generated clients for every public API service.
type ServiceClients struct {
	Account               v1connect.AccountServiceClient
	Agent                 v1connect.AgentServiceClient
	Billing               v1connect.BillingServiceClient
	Editor                v1connect.EditorServiceClient
	Environment           v1connect.EnvironmentServiceClient
	EnvironmentAutomation v1connect.EnvironmentAutomationServiceClient
	Event                 v1connect.EventServiceClient
	Gateway               v1connect.GatewayServiceClient
	Group                 v1connect.GroupServiceClient
	Identity              v1connect.IdentityServiceClient
	Insights              v1connect.InsightsServiceClient
	Integration           v1connect.IntegrationServiceClient
	Notification          v1connect.NotificationServiceClient
	Organization          v1connect.OrganizationServiceClient
	Prebuild              v1connect.PrebuildServiceClient
	Project               v1connect.ProjectServiceClient
	Runner                v1connect.RunnerServiceClient
	RunnerConfiguration   v1connect.RunnerConfigurationServiceClient
	Secret                v1connect.SecretServiceClient
	Security              v1connect.SecurityServiceClient
	ServiceAccount        v1connect.ServiceAccountServiceClient
	Team                  v1connect.TeamServiceClient
	Usage                 v1connect.UsageServiceClient
	User                  v1connect.UserServiceClient
	Webhook               v1connect.WebhookServiceClient
	Workflow              v1connect.WorkflowServiceClient
}

type config struct {
	baseURL      string
	pollInterval time.Duration
	pageSize     int32
	httpClient   *http.Client
	logger       *slog.Logger
}

// Option configures New.
type Option func(*config)

// New wraps a raw management-plane client with high-level workflows.
func New(raw *rawclient.ManagementPlane, opts ...Option) *Client {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	normalizeConfig(&cfg)

	return &Client{
		Services: newServiceClients(raw),
		raw:      raw,
		cfg:      cfg,
	}
}

// NewFromEnv creates a production SDK using ONA_API_KEY or GITPOD_API_KEY.
func NewFromEnv(opts ...Option) (*Client, error) {
	cfg := defaultConfig()
	if baseURL := strings.TrimSpace(os.Getenv(BaseURLEnvVar)); baseURL != "" {
		cfg.baseURL = baseURL
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	normalizeConfig(&cfg)

	token := apiKeyFromEnv()
	if token == "" {
		return nil, ErrMissingAPIKey
	}
	raw, err := rawclient.New(cfg.baseURL,
		rawclient.WithAccessToken(token),
		rawclient.WithHTTPClient(cfg.httpClient),
	)
	if err != nil {
		return nil, err
	}
	return &Client{Services: newServiceClients(raw), raw: raw, cfg: cfg}, nil
}

func newServiceClients(raw *rawclient.ManagementPlane) ServiceClients {
	if raw == nil {
		return ServiceClients{}
	}
	return ServiceClients{
		Account:               raw.AccountService(),
		Agent:                 raw.AgentService(),
		Billing:               raw.BillingService(),
		Editor:                raw.EditorService(),
		Environment:           raw.EnvironmentService(),
		EnvironmentAutomation: raw.EnvironmentAutomationService(),
		Event:                 raw.EventService(),
		Gateway:               raw.GatewayService(),
		Group:                 raw.GroupService(),
		Identity:              raw.IdentityService(),
		Insights:              raw.InsightsService(),
		Integration:           raw.IntegrationService(),
		Notification:          raw.NotificationService(),
		Organization:          raw.OrganizationService(),
		Prebuild:              raw.PrebuildService(),
		Project:               raw.ProjectService(),
		Runner:                raw.RunnerService(),
		RunnerConfiguration:   raw.RunnerConfigurationService(),
		Secret:                raw.SecretService(),
		Security:              raw.SecurityService(),
		ServiceAccount:        raw.ServiceAccountService(),
		Team:                  raw.TeamService(),
		Usage:                 raw.UsageService(),
		User:                  raw.UserService(),
		Webhook:               raw.WebhookService(),
		Workflow:              raw.WorkflowService(),
	}
}

func apiKeyFromEnv() string {
	if token := strings.TrimSpace(os.Getenv(APIKeyEnvVar)); token != "" {
		return token
	}
	return strings.TrimSpace(os.Getenv(LegacyAPIKeyEnvVar))
}

// WithBaseURL sets the management-plane base URL used by NewFromEnv.
func WithBaseURL(baseURL string) Option {
	return func(cfg *config) {
		cfg.baseURL = strings.TrimSpace(baseURL)
	}
}

// WithHTTPClient sets the HTTP client used by SDK workflows and NewFromEnv raw clients.
func WithHTTPClient(client *http.Client) Option {
	return func(cfg *config) {
		cfg.httpClient = client
	}
}

// WithLogger sets the logger used by high-level SDK workflows.
func WithLogger(logger *slog.Logger) Option {
	return func(cfg *config) {
		cfg.logger = logger
	}
}

// Environments returns environment cleanup workflows.
func (s *Client) Environments() *EnvironmentClient {
	return &EnvironmentClient{sdk: s}
}

func (s *Client) scm() *scmClient {
	return &scmClient{sdk: s}
}

func (s *Client) ops() *opsClient {
	return &opsClient{sdk: s}
}

func (s *Client) agents() *agentClient {
	return &agentClient{sdk: s}
}

func defaultConfig() config {
	return config{
		baseURL:      DefaultBaseURL,
		pollInterval: defaultPollInterval,
		pageSize:     defaultPageSize,
		httpClient:   http.DefaultClient,
		logger:       slog.Default(),
	}
}

func normalizeConfig(cfg *config) {
	if cfg.baseURL == "" {
		cfg.baseURL = DefaultBaseURL
	}
	if cfg.pollInterval <= 0 {
		cfg.pollInterval = defaultPollInterval
	}
	if cfg.pageSize <= 0 {
		cfg.pageSize = defaultPageSize
	}
	if cfg.httpClient == nil {
		cfg.httpClient = http.DefaultClient
	}
	if cfg.logger == nil {
		cfg.logger = slog.Default()
	}
}

func (s *Client) config() config {
	if s == nil {
		return defaultConfig()
	}
	cfg := s.cfg
	normalizeConfig(&cfg)
	return cfg
}

func (s *Client) logger() *slog.Logger {
	return s.config().logger
}

func (s *Client) requestContext(ctx context.Context) context.Context {
	return rawclient.WithSDKUserAgentLayer(ctx)
}

func (s *Client) requireRaw(operation string) (*rawclient.ManagementPlane, error) {
	if s == nil || s.raw == nil {
		return nil, &ValidationError{operationError{operation: operation, err: errSDKClientMissing}}
	}
	return s.raw, nil
}
