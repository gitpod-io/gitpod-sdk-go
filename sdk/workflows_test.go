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

func TestEnvironmentClientCreatePrefersProject(t *testing.T) {
	t.Parallel()

	type CapturedRequests struct {
		Parse             *v1.ParseContextURLRequest
		Class             *v1.GetEnvironmentClassRequest
		Auth              *v1.CheckAuthenticationForHostRequest
		Usage             *v1.GetTopProjectsRequest
		UsageHasDateRange bool
		Project           *v1.GetProjectRequest
		CreateFromProject *v1.CreateEnvironmentFromProjectRequest
		Get               *v1.GetEnvironmentRequest
	}
	type Expectation struct {
		EnvironmentID string
		Requests      CapturedRequests
		Err           string
	}

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	sdk := New(mp.Client())

	const (
		environmentClassID = "00000000-0000-0000-0000-000000000009"
		runnerID           = "11111111-1111-1111-1111-111111111111"
	)
	var got Expectation
	parse := mp.RunnerService.EXPECT().
		ParseContextURL(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.ParseContextURLRequest]) (*connect.Response[v1.ParseContextURLResponse], error) {
			got.Requests.Parse = req.Msg
			return connect.NewResponse(&v1.ParseContextURLResponse{
				Git: &v1.ParseContextURLResponse_GitContext{
					CloneUrl:          "https://github.com/acme/api.git",
					Branch:            "feature/sdk",
					Repo:              "api",
					UpstreamRemoteUrl: "https://github.com/acme/api",
				},
				ProjectIds:                    []string{"project-less", "project-more"},
				ScmId:                         "github",
				RecommendedEnvironmentClasses: []string{environmentClassID},
			}), nil
		})
	class := mp.RunnerConfigurationService.EXPECT().
		GetEnvironmentClass(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.GetEnvironmentClassRequest]) (*connect.Response[v1.GetEnvironmentClassResponse], error) {
			got.Requests.Class = req.Msg
			return connect.NewResponse(&v1.GetEnvironmentClassResponse{
				EnvironmentClass: &v1.EnvironmentClass{Id: environmentClassID, RunnerId: runnerID},
			}), nil
		}).
		After(parse)
	auth := mp.RunnerService.EXPECT().
		CheckAuthenticationForHost(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.CheckAuthenticationForHostRequest]) (*connect.Response[v1.CheckAuthenticationForHostResponse], error) {
			got.Requests.Auth = req.Msg
			return connect.NewResponse(&v1.CheckAuthenticationForHostResponse{Authenticated: true}), nil
		}).
		After(class)
	usage := mp.UsageService.EXPECT().
		GetTopProjects(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.GetTopProjectsRequest]) (*connect.Response[v1.GetTopProjectsResponse], error) {
			got.Requests.Usage = req.Msg
			got.Requests.UsageHasDateRange = req.Msg.GetDateRange().GetStartTime() != nil && req.Msg.GetDateRange().GetEndTime() != nil
			got.Requests.Usage.DateRange = nil
			return connect.NewResponse(&v1.GetTopProjectsResponse{
				Projects: []*v1.GetTopProjectsResponse_ProjectRuntimeInfo{
					{ProjectId: "project-more"},
					{ProjectId: "project-less"},
				},
				Pagination: &v1.PaginationResponse{},
			}), nil
		}).
		After(auth)
	project := mp.ProjectService.EXPECT().
		GetProject(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.GetProjectRequest]) (*connect.Response[v1.GetProjectResponse], error) {
			got.Requests.Project = req.Msg
			return connect.NewResponse(&v1.GetProjectResponse{
				Project: &v1.Project{
					Id: "project-more",
					Initializer: &v1.EnvironmentInitializer{
						Specs: []*v1.EnvironmentInitializer_Spec{{
							Spec: &v1.EnvironmentInitializer_Spec_Git{
								Git: &v1.GitInitializer{RemoteUri: "https://github.com/acme/api"},
							},
						}},
					},
				},
			}), nil
		}).
		After(usage)
	create := mp.EnvironmentService.EXPECT().
		CreateEnvironmentFromProject(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.CreateEnvironmentFromProjectRequest]) (*connect.Response[v1.CreateEnvironmentFromProjectResponse], error) {
			got.Requests.CreateFromProject = req.Msg
			return connect.NewResponse(&v1.CreateEnvironmentFromProjectResponse{
				Environment: &v1.Environment{
					Id:     "env-1",
					Status: &v1.EnvironmentStatus{Phase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_STARTING},
				},
			}), nil
		}).
		After(project)
	mp.EnvironmentService.EXPECT().
		GetEnvironment(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.GetEnvironmentRequest]) (*connect.Response[v1.GetEnvironmentResponse], error) {
			got.Requests.Get = req.Msg
			return connect.NewResponse(&v1.GetEnvironmentResponse{
				Environment: &v1.Environment{
					Id:     "env-1",
					Status: &v1.EnvironmentStatus{Phase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_RUNNING},
				},
			}), nil
		}).
		After(create)

	env, err := sdk.Environments().Create(t.Context(), CreateEnvironmentOptions{
		ContextURL: "https://github.com/acme/api",
		Name:       "api-debug",
	})
	if err != nil {
		got.Err = err.Error()
	} else {
		got.EnvironmentID = env.ID()
	}

	expected := Expectation{
		EnvironmentID: "env-1",
		Requests: CapturedRequests{
			Parse: &v1.ParseContextURLRequest{ContextUrl: "https://github.com/acme/api"},
			Class: &v1.GetEnvironmentClassRequest{EnvironmentClassId: environmentClassID},
			Auth:  &v1.CheckAuthenticationForHostRequest{RunnerId: runnerID, Host: "github.com"},
			Usage: &v1.GetTopProjectsRequest{
				Pagination: &v1.PaginationRequest{PageSize: defaultPageSize},
			},
			UsageHasDateRange: true,
			Project:           &v1.GetProjectRequest{ProjectId: "project-more"},
			CreateFromProject: &v1.CreateEnvironmentFromProjectRequest{
				ProjectId: "project-more",
				Spec: &v1.EnvironmentSpec{
					Content: &v1.EnvironmentSpec_Content{
						Initializer: &v1.EnvironmentInitializer{
							Specs: []*v1.EnvironmentInitializer_Spec{{
								Spec: &v1.EnvironmentInitializer_Spec_Git{
									Git: &v1.GitInitializer{
										RemoteUri:         "https://github.com/acme/api",
										TargetMode:        v1.GitInitializer_CLONE_TARGET_MODE_REMOTE_BRANCH,
										CloneTarget:       "feature/sdk",
										CheckoutLocation:  "api",
										UpstreamRemoteUri: "https://github.com/acme/api",
									},
								},
							}},
						},
					},
				},
				Name: Pointer("api-debug"),
			},
			Get: &v1.GetEnvironmentRequest{EnvironmentId: "env-1"},
		},
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("Create() mismatch (-want +got):\n%s", diff)
	}
}

func TestEnvironmentClientCreateFallsBackToGenericContextURL(t *testing.T) {
	t.Parallel()

	type CapturedRequests struct {
		Parse  *v1.ParseContextURLRequest
		List   *v1.ListEnvironmentClassesRequest
		Auth   *v1.CheckAuthenticationForHostRequest
		Create *v1.CreateEnvironmentRequest
		Get    *v1.GetEnvironmentRequest
	}
	type Expectation struct {
		EnvironmentID string
		Requests      CapturedRequests
		Err           string
	}

	const (
		environmentClassID = "00000000-0000-0000-0000-000000000001"
		runnerID           = "11111111-1111-1111-1111-111111111111"
	)
	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	sdk := New(mp.Client())

	var got Expectation
	parse := mp.RunnerService.EXPECT().
		ParseContextURL(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.ParseContextURLRequest]) (*connect.Response[v1.ParseContextURLResponse], error) {
			got.Requests.Parse = req.Msg
			return connect.NewResponse(&v1.ParseContextURLResponse{
				Git:   &v1.ParseContextURLResponse_GitContext{CloneUrl: "https://github.com/acme/api.git", Branch: "main"},
				ScmId: "github",
			}), nil
		})
	list := mp.EnvironmentService.EXPECT().
		ListEnvironmentClasses(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.ListEnvironmentClassesRequest]) (*connect.Response[v1.ListEnvironmentClassesResponse], error) {
			got.Requests.List = req.Msg
			return connect.NewResponse(&v1.ListEnvironmentClassesResponse{
				Pagination:         &v1.PaginationResponse{},
				EnvironmentClasses: []*v1.EnvironmentClass{{Id: environmentClassID, RunnerId: runnerID}},
			}), nil
		}).
		After(parse)
	auth := mp.RunnerService.EXPECT().
		CheckAuthenticationForHost(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.CheckAuthenticationForHostRequest]) (*connect.Response[v1.CheckAuthenticationForHostResponse], error) {
			got.Requests.Auth = req.Msg
			return connect.NewResponse(&v1.CheckAuthenticationForHostResponse{Authenticated: true}), nil
		}).
		After(list)
	create := mp.EnvironmentService.EXPECT().
		CreateEnvironment(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.CreateEnvironmentRequest]) (*connect.Response[v1.CreateEnvironmentResponse], error) {
			got.Requests.Create = req.Msg
			return connect.NewResponse(&v1.CreateEnvironmentResponse{
				Environment: &v1.Environment{
					Id:     "env-1",
					Status: &v1.EnvironmentStatus{Phase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_STARTING},
				},
			}), nil
		}).
		After(auth)
	mp.EnvironmentService.EXPECT().
		GetEnvironment(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, req *connect.Request[v1.GetEnvironmentRequest]) (*connect.Response[v1.GetEnvironmentResponse], error) {
			got.Requests.Get = req.Msg
			return connect.NewResponse(&v1.GetEnvironmentResponse{
				Environment: &v1.Environment{
					Id:     "env-1",
					Status: &v1.EnvironmentStatus{Phase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_RUNNING},
				},
			}), nil
		}).
		After(create)

	env, err := sdk.Environments().Create(t.Context(), CreateEnvironmentOptions{
		ContextURL: "https://github.com/acme/api",
		Name:       "api-debug",
	})
	if err != nil {
		got.Err = err.Error()
	} else {
		got.EnvironmentID = env.ID()
	}

	expected := Expectation{
		EnvironmentID: "env-1",
		Requests: CapturedRequests{
			Parse: &v1.ParseContextURLRequest{ContextUrl: "https://github.com/acme/api"},
			List: &v1.ListEnvironmentClassesRequest{
				Pagination: &v1.PaginationRequest{PageSize: defaultPageSize},
				Filter: &v1.ListEnvironmentClassesRequest_Filter{
					Enabled:               Pointer(true),
					CanCreateEnvironments: Pointer(true),
				},
			},
			Auth: &v1.CheckAuthenticationForHostRequest{RunnerId: runnerID, Host: "github.com"},
			Create: &v1.CreateEnvironmentRequest{
				Spec: &v1.EnvironmentSpec{
					DesiredPhase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_RUNNING,
					Machine:      &v1.EnvironmentSpec_Machine{Class: environmentClassID},
					Content: &v1.EnvironmentSpec_Content{
						Initializer: &v1.EnvironmentInitializer{
							Specs: []*v1.EnvironmentInitializer_Spec{{
								Spec: &v1.EnvironmentInitializer_Spec_ContextUrl{
									ContextUrl: &v1.ContextURLInitializer{Url: "https://github.com/acme/api"},
								},
							}},
						},
					},
				},
				Name: Pointer("api-debug"),
			},
			Get: &v1.GetEnvironmentRequest{EnvironmentId: "env-1"},
		},
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("Create() mismatch (-want +got):\n%s", diff)
	}
}

func TestEnvironmentStartCodex(t *testing.T) {
	t.Parallel()

	type CapturedRequests struct {
		Start *v1.StartAgentRequest
		Send  *v1.SendToAgentExecutionRequest
		Get   *v1.GetAgentExecutionRequest
	}
	type Expectation struct {
		AgentExecutionID string
		Requests         CapturedRequests
		Err              string
	}

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	sdk := New(mp.Client())
	env := &Environment{
		sdk: sdk,
		environment: &environmentHandle{
			client: sdk.Environments(),
			id:     "env-1",
		},
	}

	var got Expectation
	gomock.InOrder(
		mp.AgentService.EXPECT().
			StartAgent(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, req *connect.Request[v1.StartAgentRequest]) (*connect.Response[v1.StartAgentResponse], error) {
				got.Requests.Start = req.Msg
				return connect.NewResponse(&v1.StartAgentResponse{AgentExecutionId: "exec-1"}), nil
			}),
		mp.AgentService.EXPECT().
			SendToAgentExecution(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, req *connect.Request[v1.SendToAgentExecutionRequest]) (*connect.Response[v1.SendToAgentExecutionResponse], error) {
				got.Requests.Send = req.Msg
				got.Requests.Send.GetUserInput().Id = ""
				return connect.NewResponse(&v1.SendToAgentExecutionResponse{}), nil
			}),
		mp.AgentService.EXPECT().
			GetAgentExecution(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, req *connect.Request[v1.GetAgentExecutionRequest]) (*connect.Response[v1.GetAgentExecutionResponse], error) {
				got.Requests.Get = req.Msg
				return connect.NewResponse(&v1.GetAgentExecutionResponse{
					AgentExecution: &v1.AgentExecution{
						Id:     "exec-1",
						Status: &v1.AgentExecution_Status{Phase: v1.AgentExecution_PHASE_RUNNING},
					},
				}), nil
			}),
	)

	agent, err := env.StartCodex(t.Context(), EnvironmentCodexOptions{
		Name:   "SDK example task",
		Prompt: "Create a note.",
	})
	if err != nil {
		got.Err = err.Error()
	} else {
		got.AgentExecutionID = agent.ID()
	}

	expected := Expectation{
		AgentExecutionID: "exec-1",
		Requests: CapturedRequests{
			Start: &v1.StartAgentRequest{
				AgentId: codexAppInEnvironmentAgentID,
				CodeContext: &v1.AgentCodeContext{
					Context: &v1.AgentCodeContext_EnvironmentId{EnvironmentId: "env-1"},
				},
				Name: "SDK example task",
			},
			Send: &v1.SendToAgentExecutionRequest{
				AgentExecutionId: "exec-1",
				Input: &v1.SendToAgentExecutionRequest_UserInput{
					UserInput: &v1.UserInputBlock{
						Inputs: []*v1.UserInputBlock_Input{{
							Input: &v1.UserInputBlock_Input_Text{
								Text: &v1.UserInputBlock_TextInput{Content: "Create a note."},
							},
						}},
					},
				},
			},
			Get: &v1.GetAgentExecutionRequest{AgentExecutionId: "exec-1"},
		},
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("StartCodex() mismatch (-want +got):\n%s", diff)
	}
}

func TestEnvironmentStartCodexReturnsSessionWhenWaitFails(t *testing.T) {
	t.Parallel()

	type CapturedRequests struct {
		Start *v1.StartAgentRequest
		Send  *v1.SendToAgentExecutionRequest
		Get   *v1.GetAgentExecutionRequest
	}
	type Expectation struct {
		AgentExecutionID string
		Requests         CapturedRequests
		Err              string
	}

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	sdk := New(mp.Client())
	env := &Environment{
		sdk: sdk,
		environment: &environmentHandle{
			client: sdk.Environments(),
			id:     "env-1",
		},
	}

	var got Expectation
	gomock.InOrder(
		mp.AgentService.EXPECT().
			StartAgent(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, req *connect.Request[v1.StartAgentRequest]) (*connect.Response[v1.StartAgentResponse], error) {
				got.Requests.Start = req.Msg
				return connect.NewResponse(&v1.StartAgentResponse{AgentExecutionId: "exec-1"}), nil
			}),
		mp.AgentService.EXPECT().
			SendToAgentExecution(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, req *connect.Request[v1.SendToAgentExecutionRequest]) (*connect.Response[v1.SendToAgentExecutionResponse], error) {
				got.Requests.Send = req.Msg
				got.Requests.Send.GetUserInput().Id = ""
				return connect.NewResponse(&v1.SendToAgentExecutionResponse{}), nil
			}),
		mp.AgentService.EXPECT().
			GetAgentExecution(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, req *connect.Request[v1.GetAgentExecutionRequest]) (*connect.Response[v1.GetAgentExecutionResponse], error) {
				got.Requests.Get = req.Msg
				return connect.NewResponse(&v1.GetAgentExecutionResponse{
					AgentExecution: &v1.AgentExecution{
						Id: "exec-1",
						Status: &v1.AgentExecution_Status{
							Phase:            v1.AgentExecution_PHASE_STOPPED,
							FailureReason:    v1.AgentExecutionFailureReason_AGENT_EXECUTION_FAILURE_REASON_SERVICE,
							FailureMessage:   "agent process exited",
							SupportBundleUrl: "https://example.com/support-bundle?token=secret",
						},
					},
				}), nil
			}),
	)

	agent, err := env.StartCodex(t.Context(), EnvironmentCodexOptions{
		Name:   "SDK example task",
		Prompt: "Create a note.",
	})
	if err != nil {
		got.Err = err.Error()
	}
	if agent != nil {
		got.AgentExecutionID = agent.ID()
	}

	expected := Expectation{
		AgentExecutionID: "exec-1",
		Requests: CapturedRequests{
			Start: &v1.StartAgentRequest{
				AgentId: codexAppInEnvironmentAgentID,
				CodeContext: &v1.AgentCodeContext{
					Context: &v1.AgentCodeContext_EnvironmentId{EnvironmentId: "env-1"},
				},
				Name: "SDK example task",
			},
			Send: &v1.SendToAgentExecutionRequest{
				AgentExecutionId: "exec-1",
				Input: &v1.SendToAgentExecutionRequest_UserInput{
					UserInput: &v1.UserInputBlock{
						Inputs: []*v1.UserInputBlock_Input{{
							Input: &v1.UserInputBlock_Input_Text{
								Text: &v1.UserInputBlock_TextInput{Content: "Create a note."},
							},
						}},
					},
				},
			},
			Get: &v1.GetAgentExecutionRequest{AgentExecutionId: "exec-1"},
		},
		Err: `agents.wait_running: agent execution exec-1 failed: agent process exited; last status: phase=stopped failure_reason=AGENT_EXECUTION_FAILURE_REASON_SERVICE failure="agent process exited" support_bundle=https://example.com/support-bundle`,
	}
	if diff := cmp.Diff(expected, got, protocmp.Transform()); diff != "" {
		t.Errorf("StartCodex() mismatch (-want +got):\n%s", diff)
	}
}

func TestEnvironmentStartCodexValidation(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Err string
	}

	tests := []struct {
		Name     string
		Options  EnvironmentCodexOptions
		Expected Expectation
	}{
		{
			Name: "requires_prompt",
			Expected: Expectation{
				Err: "environments.start_codex: prompt is required",
			},
		},
		{
			Name:    "requires_non_blank_prompt",
			Options: EnvironmentCodexOptions{Prompt: " \t\n"},
			Expected: Expectation{
				Err: "environments.start_codex: prompt is required",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			sdk := New(nil)
			env := &Environment{
				sdk: sdk,
				environment: &environmentHandle{
					client: sdk.Environments(),
					id:     "env-1",
				},
			}

			var got Expectation
			_, err := env.StartCodex(t.Context(), tc.Options)
			if err != nil {
				got.Err = err.Error()
			}

			if diff := cmp.Diff(tc.Expected, got); diff != "" {
				t.Errorf("StartCodex() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEnvironmentClientCreateValidation(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Err string
	}

	tests := []struct {
		Name     string
		Options  CreateEnvironmentOptions
		Expected Expectation
	}{
		{
			Name: "requires_context_url",
			Expected: Expectation{
				Err: "environments.create: context URL is required",
			},
		},
		{
			Name: "returns_required_scm_authentication",
			Options: CreateEnvironmentOptions{
				ContextURL: "https://github.com/acme/api",
			},
			Expected: Expectation{
				Err: "scm.check_authentication_for_host: authentication required for github.com",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mp := rawclient.NewMock(ctrl)
			sdk := New(mp.Client())

			if tc.Options.ContextURL != "" {
				parse := mp.RunnerService.EXPECT().
					ParseContextURL(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&v1.ParseContextURLResponse{
						Git:                           &v1.ParseContextURLResponse_GitContext{CloneUrl: "https://github.com/acme/api.git"},
						RecommendedEnvironmentClasses: []string{"class-1"},
					}), nil)
				class := mp.RunnerConfigurationService.EXPECT().
					GetEnvironmentClass(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&v1.GetEnvironmentClassResponse{
						EnvironmentClass: &v1.EnvironmentClass{Id: "class-1", RunnerId: "runner-1"},
					}), nil).
					After(parse)
				mp.RunnerService.EXPECT().
					CheckAuthenticationForHost(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&v1.CheckAuthenticationForHostResponse{Authenticated: false}), nil).
					After(class)
			}

			var got Expectation
			_, err := sdk.Environments().Create(t.Context(), tc.Options)
			if err != nil {
				got.Err = err.Error()
			}

			if diff := cmp.Diff(tc.Expected, got); diff != "" {
				t.Errorf("Create() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEnvironmentClientCreatePropagatesParseError(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mp := rawclient.NewMock(ctrl)
	sdk := New(mp.Client())
	mp.RunnerService.EXPECT().
		ParseContextURL(gomock.Any(), gomock.Any()).
		Return(nil, connect.NewError(connect.CodeUnavailable, errors.New("parse unavailable")))

	type Expectation struct {
		Err string
	}
	var got Expectation
	_, err := sdk.Environments().Create(t.Context(), CreateEnvironmentOptions{
		ContextURL: "https://github.com/acme/api",
	})
	if err != nil {
		got.Err = err.Error()
	}

	expected := Expectation{Err: "scm.resolve_context: unavailable: parse unavailable"}
	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("Create() mismatch (-want +got):\n%s", diff)
	}
}
