package sdk

import (
	"context"
	"errors"
	"testing"

	"connectrpc.com/connect"
	rawclient "github.com/gitpod-io/gitpod-sdk-go/internal/raw"
	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
)

// not parallel: modifies process environment.
func TestNewFromEnv(t *testing.T) {
	type Expectation struct {
		HasSDK  bool
		BaseURL string
		Err     string
	}

	tests := []struct {
		Name          string
		APIKey        string
		LegacyAPIKey  string
		EnvBaseURL    string
		OptionBaseURL string
		Expected      Expectation
	}{
		{
			Name: "requires_api_key",
			Expected: Expectation{
				Err: ErrMissingAPIKey.Error(),
			},
		},
		{
			Name:   "uses_default_base_url",
			APIKey: "ona_api_key_test",
			Expected: Expectation{
				HasSDK:  true,
				BaseURL: DefaultBaseURL,
			},
		},
		{
			Name:         "accepts_legacy_api_key",
			LegacyAPIKey: "gitpod_api_key_test",
			Expected: Expectation{
				HasSDK:  true,
				BaseURL: DefaultBaseURL,
			},
		},
		{
			Name:       "uses_environment_base_url",
			APIKey:     "ona_api_key_test",
			EnvBaseURL: "https://ona.example.com/api",
			Expected: Expectation{
				HasSDK:  true,
				BaseURL: "https://ona.example.com/api",
			},
		},
		{
			Name:          "explicit_base_url_overrides_environment",
			APIKey:        "ona_api_key_test",
			EnvBaseURL:    "https://ignored.example.com/api",
			OptionBaseURL: "https://ona.example.com/custom/api",
			Expected: Expectation{
				HasSDK:  true,
				BaseURL: "https://ona.example.com/custom/api",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			// not parallel: modifies process environment.
			t.Setenv(APIKeyEnvVar, tc.APIKey)
			t.Setenv(LegacyAPIKeyEnvVar, tc.LegacyAPIKey)
			t.Setenv(BaseURLEnvVar, tc.EnvBaseURL)

			var opts []Option
			if tc.OptionBaseURL != "" {
				opts = append(opts, WithBaseURL(tc.OptionBaseURL))
			}

			var got Expectation
			sdk, err := NewFromEnv(opts...)
			if err != nil {
				got.Err = err.Error()
				if !errors.Is(err, ErrMissingAPIKey) {
					t.Errorf("NewFromEnv() returned unexpected error type: %v", err)
				}
			} else {
				got.HasSDK = sdk != nil
				got.BaseURL = sdk.config().baseURL
			}

			if diff := cmp.Diff(tc.Expected, got); diff != "" {
				t.Errorf("NewFromEnv() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// not parallel: modifies process environment.
func TestAPIKeyFromEnv(t *testing.T) {
	type Expectation struct {
		APIKey string
	}

	tests := []struct {
		Name         string
		APIKey       string
		LegacyAPIKey string
		Expected     Expectation
	}{
		{
			Name:   "uses_ona_api_key",
			APIKey: "ona_api_key_test",
			Expected: Expectation{
				APIKey: "ona_api_key_test",
			},
		},
		{
			Name:         "falls_back_to_gitpod_api_key",
			LegacyAPIKey: "gitpod_api_key_test",
			Expected: Expectation{
				APIKey: "gitpod_api_key_test",
			},
		},
		{
			Name:         "ignores_blank_ona_api_key",
			APIKey:       "  ",
			LegacyAPIKey: "gitpod_api_key_test",
			Expected: Expectation{
				APIKey: "gitpod_api_key_test",
			},
		},
		{
			Name:         "prefers_ona_api_key",
			APIKey:       "ona_api_key_test",
			LegacyAPIKey: "gitpod_api_key_test",
			Expected: Expectation{
				APIKey: "ona_api_key_test",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			// not parallel: modifies process environment.
			t.Setenv(APIKeyEnvVar, tc.APIKey)
			t.Setenv(LegacyAPIKeyEnvVar, tc.LegacyAPIKey)

			got := Expectation{APIKey: apiKeyFromEnv()}
			if diff := cmp.Diff(tc.Expected, got); diff != "" {
				t.Errorf("apiKeyFromEnv() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEnvironmentClientDelete(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Request *v1.DeleteEnvironmentRequest
		Err     string
	}

	tests := []struct {
		Name          string
		EnvironmentID string
		Options       DeleteEnvironmentOptions
		Expected      Expectation
	}{
		{
			Name:          "deletes_environment_by_id",
			EnvironmentID: "env-1",
			Expected: Expectation{
				Request: &v1.DeleteEnvironmentRequest{EnvironmentId: "env-1"},
			},
		},
		{
			Name:          "force_deletes_environment_by_id",
			EnvironmentID: "env-1",
			Options:       DeleteEnvironmentOptions{Force: true},
			Expected: Expectation{
				Request: &v1.DeleteEnvironmentRequest{EnvironmentId: "env-1", Force: true},
			},
		},
		{
			Name: "requires_environment_id",
			Expected: Expectation{
				Err: "environments.delete: environment ID is required",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mp := rawclient.NewMock(ctrl)
			sdk := New(mp.Client())

			var got Expectation
			if tc.EnvironmentID != "" {
				mp.EnvironmentService.EXPECT().
					DeleteEnvironment(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, req *connect.Request[v1.DeleteEnvironmentRequest]) (*connect.Response[v1.DeleteEnvironmentResponse], error) {
						got.Request = req.Msg
						return connect.NewResponse(&v1.DeleteEnvironmentResponse{}), nil
					})
			}

			err := sdk.Environments().Delete(t.Context(), tc.EnvironmentID, tc.Options)
			if err != nil {
				got.Err = err.Error()
			}

			if diff := cmp.Diff(tc.Expected, got, protocmp.Transform()); diff != "" {
				t.Errorf("Delete() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEnvironmentClientGet(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Request       *v1.GetEnvironmentRequest
		EnvironmentID string
		Proto         *v1.Environment
		Err           string
	}

	tests := []struct {
		Name          string
		EnvironmentID string
		Expected      Expectation
	}{
		{
			Name:          "gets_environment_by_id",
			EnvironmentID: "env-1",
			Expected: Expectation{
				Request:       &v1.GetEnvironmentRequest{EnvironmentId: "env-1"},
				EnvironmentID: "env-1",
				Proto: &v1.Environment{
					Id:     "env-1",
					Status: &v1.EnvironmentStatus{Phase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_RUNNING},
				},
			},
		},
		{
			Name: "requires_environment_id",
			Expected: Expectation{
				Err: "environments.get: environment ID is required",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mp := rawclient.NewMock(ctrl)
			sdk := New(mp.Client())

			var got Expectation
			if tc.EnvironmentID != "" {
				mp.EnvironmentService.EXPECT().
					GetEnvironment(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, req *connect.Request[v1.GetEnvironmentRequest]) (*connect.Response[v1.GetEnvironmentResponse], error) {
						got.Request = req.Msg
						return connect.NewResponse(&v1.GetEnvironmentResponse{
							Environment: tc.Expected.Proto,
						}), nil
					})
			}

			env, err := sdk.Environments().Get(t.Context(), tc.EnvironmentID)
			if err != nil {
				got.Err = err.Error()
			} else {
				got.EnvironmentID = env.ID()
				got.Proto = env.Proto()
			}

			if diff := cmp.Diff(tc.Expected, got, protocmp.Transform()); diff != "" {
				t.Errorf("Get() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEnvironmentClientList(t *testing.T) {
	t.Parallel()

	t.Run("does_not_call_api_until_iterated", func(t *testing.T) {
		t.Parallel()

		type Expectation struct {
			HasSequence bool
		}

		ctrl := gomock.NewController(t)
		mp := rawclient.NewMock(ctrl)
		sdk := New(mp.Client())

		seq := sdk.Environments().List(t.Context())
		got := Expectation{HasSequence: seq != nil}
		expected := Expectation{HasSequence: true}
		if diff := cmp.Diff(expected, got); diff != "" {
			t.Errorf("List() mismatch (-want +got):\n%s", diff)
		}
	})

	allArchivalStatuses := v1.ListEnvironmentsRequest_ARCHIVAL_STATUS_ALL
	type Expectation struct {
		IdentityRequest bool
		ListRequests    []*v1.ListEnvironmentsRequest
		EnvironmentIDs  []string
		Protos          []*v1.Environment
		Err             string
	}

	tests := []struct {
		Name                      string
		WithoutEnvironmentService bool
		WithoutIdentityService    bool
		Identity                  *v1.GetAuthenticatedIdentityResponse
		IdentityErr               error
		Pages                     []*v1.ListEnvironmentsResponse
		ListErr                   error
		StopAfter                 int
		Expected                  Expectation
	}{
		{
			Name:                      "requires_environment_service",
			WithoutEnvironmentService: true,
			Expected: Expectation{
				Err: "environments.list: environment service is not configured",
			},
		},
		{
			Name:                   "requires_identity_service",
			WithoutIdentityService: true,
			Expected: Expectation{
				Err: "environments.list: identity service is not configured",
			},
		},
		{
			Name: "lists_caller_owned_default_environments_across_pages",
			Identity: &v1.GetAuthenticatedIdentityResponse{
				Subject: &v1.Subject{Id: "11111111-1111-1111-1111-111111111111", Principal: v1.Principal_PRINCIPAL_USER},
			},
			Pages: []*v1.ListEnvironmentsResponse{
				{
					Environments: []*v1.Environment{
						{
							Id:       "env-1",
							Metadata: &v1.EnvironmentMetadata{Role: v1.EnvironmentRole_ENVIRONMENT_ROLE_DEFAULT},
						},
					},
					Pagination: &v1.PaginationResponse{NextToken: "page-2"},
				},
				{
					Environments: []*v1.Environment{
						{
							Id:       "env-2",
							Metadata: &v1.EnvironmentMetadata{Role: v1.EnvironmentRole_ENVIRONMENT_ROLE_DEFAULT},
						},
					},
					Pagination: &v1.PaginationResponse{},
				},
			},
			Expected: Expectation{
				IdentityRequest: true,
				ListRequests: []*v1.ListEnvironmentsRequest{
					{
						Pagination: &v1.PaginationRequest{PageSize: defaultPageSize},
						Filter: &v1.ListEnvironmentsRequest_Filter{
							CreatorIds:     []string{"11111111-1111-1111-1111-111111111111"},
							Roles:          []v1.EnvironmentRole{v1.EnvironmentRole_ENVIRONMENT_ROLE_DEFAULT},
							ArchivalStatus: &allArchivalStatuses,
						},
					},
					{
						Pagination: &v1.PaginationRequest{PageSize: defaultPageSize, Token: "page-2"},
						Filter: &v1.ListEnvironmentsRequest_Filter{
							CreatorIds:     []string{"11111111-1111-1111-1111-111111111111"},
							Roles:          []v1.EnvironmentRole{v1.EnvironmentRole_ENVIRONMENT_ROLE_DEFAULT},
							ArchivalStatus: &allArchivalStatuses,
						},
					},
				},
				EnvironmentIDs: []string{"env-1", "env-2"},
				Protos: []*v1.Environment{
					{
						Id:       "env-1",
						Metadata: &v1.EnvironmentMetadata{Role: v1.EnvironmentRole_ENVIRONMENT_ROLE_DEFAULT},
					},
					{
						Id:       "env-2",
						Metadata: &v1.EnvironmentMetadata{Role: v1.EnvironmentRole_ENVIRONMENT_ROLE_DEFAULT},
					},
				},
			},
		},
		{
			Name: "stops_pagination_when_caller_stops",
			Identity: &v1.GetAuthenticatedIdentityResponse{
				Subject: &v1.Subject{Id: "11111111-1111-1111-1111-111111111111", Principal: v1.Principal_PRINCIPAL_USER},
			},
			Pages: []*v1.ListEnvironmentsResponse{
				{
					Environments: []*v1.Environment{
						{
							Id:       "env-1",
							Metadata: &v1.EnvironmentMetadata{Role: v1.EnvironmentRole_ENVIRONMENT_ROLE_DEFAULT},
						},
					},
					Pagination: &v1.PaginationResponse{NextToken: "page-2"},
				},
			},
			StopAfter: 1,
			Expected: Expectation{
				IdentityRequest: true,
				ListRequests: []*v1.ListEnvironmentsRequest{
					{
						Pagination: &v1.PaginationRequest{PageSize: defaultPageSize},
						Filter: &v1.ListEnvironmentsRequest_Filter{
							CreatorIds:     []string{"11111111-1111-1111-1111-111111111111"},
							Roles:          []v1.EnvironmentRole{v1.EnvironmentRole_ENVIRONMENT_ROLE_DEFAULT},
							ArchivalStatus: &allArchivalStatuses,
						},
					},
				},
				EnvironmentIDs: []string{"env-1"},
				Protos: []*v1.Environment{
					{
						Id:       "env-1",
						Metadata: &v1.EnvironmentMetadata{Role: v1.EnvironmentRole_ENVIRONMENT_ROLE_DEFAULT},
					},
				},
			},
		},
		{
			Name:        "maps_identity_errors",
			IdentityErr: connect.NewError(connect.CodeUnauthenticated, errors.New("missing token")),
			Expected: Expectation{
				IdentityRequest: true,
				Err:             "environments.list: unauthenticated: missing token",
			},
		},
		{
			Name:     "requires_authenticated_subject_id",
			Identity: &v1.GetAuthenticatedIdentityResponse{},
			Expected: Expectation{
				IdentityRequest: true,
				Err:             "environments.list: authenticated identity did not include a subject ID",
			},
		},
		{
			Name: "maps_list_errors",
			Identity: &v1.GetAuthenticatedIdentityResponse{
				Subject: &v1.Subject{Id: "11111111-1111-1111-1111-111111111111", Principal: v1.Principal_PRINCIPAL_USER},
			},
			ListErr: connect.NewError(connect.CodePermissionDenied, errors.New("no access")),
			Expected: Expectation{
				IdentityRequest: true,
				ListRequests: []*v1.ListEnvironmentsRequest{
					{
						Pagination: &v1.PaginationRequest{PageSize: defaultPageSize},
						Filter: &v1.ListEnvironmentsRequest_Filter{
							CreatorIds:     []string{"11111111-1111-1111-1111-111111111111"},
							Roles:          []v1.EnvironmentRole{v1.EnvironmentRole_ENVIRONMENT_ROLE_DEFAULT},
							ArchivalStatus: &allArchivalStatuses,
						},
					},
				},
				Err: "environments.list: permission_denied: no access",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mp := rawclient.NewMock(ctrl)
			if tc.WithoutEnvironmentService {
				mp.EnvironmentService = nil
			}
			if tc.WithoutIdentityService {
				mp.IdentityService = nil
			}
			sdk := New(mp.Client())

			var got Expectation
			if !tc.WithoutEnvironmentService && !tc.WithoutIdentityService {
				mp.IdentityService.EXPECT().
					GetAuthenticatedIdentity(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, req *connect.Request[v1.GetAuthenticatedIdentityRequest]) (*connect.Response[v1.GetAuthenticatedIdentityResponse], error) {
						got.IdentityRequest = req.Msg != nil
						if tc.IdentityErr != nil {
							return nil, tc.IdentityErr
						}
						return connect.NewResponse(tc.Identity), nil
					})
			}

			if tc.Identity != nil && tc.Identity.GetSubject().GetId() != "" {
				call := mp.EnvironmentService.EXPECT().
					ListEnvironments(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, req *connect.Request[v1.ListEnvironmentsRequest]) (*connect.Response[v1.ListEnvironmentsResponse], error) {
						got.ListRequests = append(got.ListRequests, req.Msg)
						if tc.ListErr != nil {
							return nil, tc.ListErr
						}
						return connect.NewResponse(tc.Pages[0]), nil
					})
				if len(tc.Pages) > 1 {
					for _, page := range tc.Pages[1:] {
						nextPage := page
						call = mp.EnvironmentService.EXPECT().
							ListEnvironments(gomock.Any(), gomock.Any()).
							DoAndReturn(func(ctx context.Context, req *connect.Request[v1.ListEnvironmentsRequest]) (*connect.Response[v1.ListEnvironmentsResponse], error) {
								got.ListRequests = append(got.ListRequests, req.Msg)
								return connect.NewResponse(nextPage), nil
							}).
							After(call)
					}
				}
			}

			for env, err := range sdk.Environments().List(t.Context()) {
				if err != nil {
					got.Err = err.Error()
					break
				}
				if env != nil {
					got.EnvironmentIDs = append(got.EnvironmentIDs, env.ID())
					got.Protos = append(got.Protos, env.Proto())
				}
				if tc.StopAfter > 0 && len(got.EnvironmentIDs) >= tc.StopAfter {
					break
				}
			}

			if diff := cmp.Diff(tc.Expected, got, protocmp.Transform()); diff != "" {
				t.Errorf("List() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEnvironmentClientStop(t *testing.T) {
	t.Parallel()

	type CapturedRequests struct {
		Stop *v1.StopEnvironmentRequest
		Get  *v1.GetEnvironmentRequest
	}
	type Expectation struct {
		Requests CapturedRequests
		Err      string
	}

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	sdk := New(mp.Client())

	var got Expectation
	stop := mp.EnvironmentService.EXPECT().
		StopEnvironment(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.StopEnvironmentRequest]) (*connect.Response[v1.StopEnvironmentResponse], error) {
			got.Requests.Stop = req.Msg
			return connect.NewResponse(&v1.StopEnvironmentResponse{}), nil
		})
	mp.EnvironmentService.EXPECT().
		GetEnvironment(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.GetEnvironmentRequest]) (*connect.Response[v1.GetEnvironmentResponse], error) {
			got.Requests.Get = req.Msg
			return connect.NewResponse(&v1.GetEnvironmentResponse{
				Environment: &v1.Environment{
					Id:     "env-1",
					Status: &v1.EnvironmentStatus{Phase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_STOPPED},
				},
			}), nil
		}).
		After(stop)

	err := sdk.Environments().Stop(t.Context(), "env-1")
	if err != nil {
		got.Err = err.Error()
	}

	expected := Expectation{
		Requests: CapturedRequests{
			Stop: &v1.StopEnvironmentRequest{EnvironmentId: "env-1"},
			Get:  &v1.GetEnvironmentRequest{EnvironmentId: "env-1"},
		},
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("Stop() mismatch (-want +got):\n%s", diff)
	}
}
