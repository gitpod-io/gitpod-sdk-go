package sdk

import (
	"testing"

	rawclient "github.com/gitpod-io/gitpod-sdk-go/internal/raw"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

func TestNewExposesPublicServices(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Services []bool
	}

	raw := rawclient.NewMock(gomock.NewController(t))
	services := New(raw.Client()).Services
	got := Expectation{Services: []bool{
		services.Account == raw.AccountService,
		services.Agent == raw.AgentService,
		services.Billing == raw.BillingService,
		services.Editor == raw.EditorService,
		services.Environment == raw.EnvironmentService,
		services.EnvironmentAutomation == raw.EnvironmentAutomationService,
		services.Event == raw.EventService,
		services.Gateway == raw.GatewayService,
		services.Group == raw.GroupService,
		services.Identity == raw.IdentityService,
		services.Insights == raw.InsightsService,
		services.Integration == raw.IntegrationService,
		services.Notification == raw.NotificationService,
		services.Organization == raw.OrganizationService,
		services.Prebuild == raw.PrebuildService,
		services.Project == raw.ProjectService,
		services.Runner == raw.RunnerService,
		services.RunnerConfiguration == raw.RunnerConfigurationService,
		services.Secret == raw.SecretService,
		services.Security == raw.SecurityService,
		services.ServiceAccount == raw.ServiceAccountService,
		services.Team == raw.TeamService,
		services.Usage == raw.UsageService,
		services.User == raw.UserService,
		services.Webhook == raw.WebhookService,
		services.Workflow == raw.WorkflowService,
	}}
	want := Expectation{Services: []bool{
		true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true,
	}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("New() public services mismatch (-want +got):\n%s", diff)
	}
}
