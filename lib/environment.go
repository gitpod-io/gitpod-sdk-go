package lib

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gitpod-go"
	"github.com/stainless-sdks/gitpod-go/internal/apierror"
)

type CreateOptions struct {
	ProjectID        string
	ContextURL       string
	EnvironmentClass string
}

type EnvironmentService interface {
	Create(ctx context.Context, options *CreateOptions) (*gitpod.EnvironmentGetResponseEnvironment, error)
	Start(ctx context.Context, environmentID string) (*gitpod.EnvironmentGetResponseEnvironment, error)
	Stop(ctx context.Context, environmentID string) (*gitpod.EnvironmentGetResponseEnvironment, error)
	Delete(ctx context.Context, environmentID string) error
}

func NewEnvironmentService(client *gitpod.Client) EnvironmentService {
	return &environmentService{
		Client: client,
	}
}

type environmentService struct {
	Client *gitpod.Client
}

func (s *environmentService) Create(ctx context.Context, options *CreateOptions) (*gitpod.EnvironmentGetResponseEnvironment, error) {
	envID, err := s.create(ctx, options)
	if err != nil {
		return nil, err
	}
	return waitForRunning(ctx, s.Client, envID, func(env *gitpod.EnvironmentGetResponseEnvironment) (bool, error) {
		return true, nil
	})
}

func (s *environmentService) Start(ctx context.Context, environmentID string) (*gitpod.EnvironmentGetResponseEnvironment, error) {
	_, err := s.Client.Environments.Start(ctx, gitpod.EnvironmentStartParams{
		EnvironmentID: gitpod.String(environmentID),
	})
	if err != nil {
		return nil, err
	}
	return waitForRunning(ctx, s.Client, environmentID, func(env *gitpod.EnvironmentGetResponseEnvironment) (bool, error) {
		return true, nil
	})
}

func (s *environmentService) Stop(ctx context.Context, environmentID string) (*gitpod.EnvironmentGetResponseEnvironment, error) {
	_, err := s.Client.Environments.Stop(ctx, gitpod.EnvironmentStopParams{
		EnvironmentID: gitpod.String(environmentID),
	})
	if err != nil {
		return nil, err
	}
	return waitForStopped(ctx, s.Client, environmentID)
}

func (s *environmentService) Delete(ctx context.Context, environmentID string) error {
	_, err := s.Client.Environments.Delete(ctx, gitpod.EnvironmentDeleteParams{
		EnvironmentID: gitpod.String(environmentID),
	})
	if err != nil {
		apierr, ok := err.(*apierror.Error)
		if ok && apierr.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return waitForDeleted(ctx, s.Client, environmentID)
}

func waitForDeleted(ctx context.Context, client *gitpod.Client, envID string) error {
	_, err := waitForEnvironment(ctx, client, envID, func() (*gitpod.EnvironmentGetResponseEnvironment, bool, error) {
		resp, err := client.Environments.Get(ctx, gitpod.EnvironmentGetParams{
			EnvironmentID: gitpod.String(envID),
		})
		if err != nil {
			apierr, ok := err.(*apierror.Error)
			if ok && apierr.StatusCode == http.StatusNotFound {
				return nil, true, nil
			}
			return nil, false, err
		}
		return nil, resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseDeleted, nil
	})
	return err
}

func waitForStopped(ctx context.Context, client *gitpod.Client, envID string) (*gitpod.EnvironmentGetResponseEnvironment, error) {
	return waitForEnvironment(ctx, client, envID, func() (*gitpod.EnvironmentGetResponseEnvironment, bool, error) {
		resp, err := client.Environments.Get(ctx, gitpod.EnvironmentGetParams{
			EnvironmentID: gitpod.String(envID),
		})
		if err != nil {
			return nil, false, err
		}
		if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseStopped ||
			resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseDeleting ||
			resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseDeleted {
			return &resp.Environment, true, nil
		}
		return nil, false, nil
	})
}

func waitForRunning(ctx context.Context, client *gitpod.Client, envID string, check func(*gitpod.EnvironmentGetResponseEnvironment) (bool, error)) (*gitpod.EnvironmentGetResponseEnvironment, error) {
	return waitForEnvironment(ctx, client, envID, func() (*gitpod.EnvironmentGetResponseEnvironment, bool, error) {
		resp, err := client.Environments.Get(ctx, gitpod.EnvironmentGetParams{
			EnvironmentID: gitpod.String(envID),
		})
		if err != nil {
			return nil, false, err
		}

		if fm := resp.Environment.Status.FailureMessage; len(fm) > 0 {
			return nil, false, fmt.Errorf("environment creation failed: %s", fm)
		}
		if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseRunning {
			ok, err := check(&resp.Environment)
			if err != nil {
				return nil, false, err
			}
			return &resp.Environment, ok, nil
		}
		if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseStopping {
			return nil, false, errors.New("environment creation failed: environment is stopping")
		}
		if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseStopped {
			return nil, false, errors.New("environment creation failed: environment is stopped")
		}
		if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseDeleting {
			return nil, false, errors.New("environment creation failed: environment is deleting")
		}
		if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseDeleted {
			return nil, false, errors.New("environment creation failed: environment is deleted")
		}
		return nil, false, nil
	})
}

func waitForEnvironment(ctx context.Context, client *gitpod.Client, envID string, check func() (*gitpod.EnvironmentGetResponseEnvironment, bool, error)) (*gitpod.EnvironmentGetResponseEnvironment, error) {
	return WaitFor(ctx, client, envID, gitpod.EventWatchResponseResourceTypeResourceTypeEnvironment, envID, check)
}

func WaitFor[T any](ctx context.Context, client *gitpod.Client, envID string, resourceType gitpod.EventWatchResponseResourceType, resourceID string, check func() (*T, bool, error)) (*T, error) {
	env, ok, err := check()
	if err != nil {
		return nil, err
	}
	if ok {
		return env, nil
	}

	stream := client.Events.WatchStreaming(ctx, gitpod.EventWatchParams{
		Body: gitpod.EventWatchParamsBodyEnvironmentScopeProducesEventsForTheEnvironmentItselfAllTasksTaskExecutionsAndServicesAssociatedWithThatEnvironment{
			EnvironmentID: gitpod.String(envID),
		},
	})
	defer stream.Close()

	for stream.Next() {
		resp := stream.Current()
		if resp.ResourceType == resourceType && resp.ResourceID == resourceID {
			env, ok, err := check()
			if err != nil {
				// TODO: if transient we should retry?
				return nil, err
			}
			if ok {
				return env, nil
			}
		}
	}

	return nil, stream.Err()
}

func (s *environmentService) create(ctx context.Context, options *CreateOptions) (string, error) {
	if options.ProjectID != "" {
		resp, err := s.Client.Environments.NewFromProject(ctx, gitpod.EnvironmentNewFromProjectParams{
			ProjectID: gitpod.String(options.ProjectID),
		})
		if err != nil {
			return "", err
		}
		return resp.Environment.ID, nil
	}
	if options.ContextURL != "" {
		if options.EnvironmentClass == "" {
			return "", errors.New("environmentClass must be provided when contextURL is provided")
		}
		resp, err := s.Client.Environments.New(ctx, gitpod.EnvironmentNewParams{
			Spec: gitpod.F(gitpod.EnvironmentNewParamsSpec{
				DesiredPhase: gitpod.F(gitpod.EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseRunning),
				Machine: gitpod.F(gitpod.EnvironmentNewParamsSpecMachine{
					Class: gitpod.String(options.EnvironmentClass),
				}),
				Content: gitpod.F(gitpod.EnvironmentNewParamsSpecContent{
					Initializer: gitpod.F(gitpod.EnvironmentNewParamsSpecContentInitializer{
						Specs: gitpod.F([]gitpod.EnvironmentNewParamsSpecContentInitializerSpecUnion{
							gitpod.EnvironmentNewParamsSpecContentInitializerSpecsContextURL{
								ContextURL: gitpod.F(gitpod.EnvironmentNewParamsSpecContentInitializerSpecsContextURLContextURL{
									URL: gitpod.String(options.ContextURL),
								}),
							},
						}),
					}),
				}),
			}),
		})
		if err != nil {
			return "", err
		}
		return resp.Environment.ID, nil
	}
	return "", errors.New("either projectID or contextURL must be provided")
}
