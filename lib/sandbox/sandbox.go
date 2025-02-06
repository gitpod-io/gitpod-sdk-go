package sandbox

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gitpod-go"
	"github.com/stainless-sdks/gitpod-go/internal/apierror"
)

type EnvironmentSandbox struct {
	Client *gitpod.Client
}

func NewEnvironmentSandbox(client *gitpod.Client) *EnvironmentSandbox {
	return &EnvironmentSandbox{
		Client: client,
	}
}

type CreateEnvironmentOptions struct {
	ProjectID        string
	ContextURL       string
	EnvironmentClass string
}

func (s *EnvironmentSandbox) Create(ctx context.Context, options *CreateEnvironmentOptions) (res *gitpod.EnvironmentGetResponseEnvironment, err error) {
	envID, err := s.create(ctx, options)
	if err != nil {
		return nil, err
	}
	return s.waitForRunning(ctx, envID)
}

type StartEnvironmentOptions struct {
	EnvironmentID string
}

func (s *EnvironmentSandbox) Start(ctx context.Context, options *StartEnvironmentOptions) (res *gitpod.EnvironmentGetResponseEnvironment, err error) {
	_, err = s.Client.Environments.Start(ctx, gitpod.EnvironmentStartParams{
		EnvironmentID: gitpod.String(options.EnvironmentID),
	})
	if err != nil {
		return nil, err
	}
	return s.waitForRunning(ctx, options.EnvironmentID)
}

type StopEnvironmentOptions struct {
	EnvironmentID string
}

func (s *EnvironmentSandbox) Stop(ctx context.Context, options *StopEnvironmentOptions) (res *gitpod.EnvironmentGetResponseEnvironment, err error) {
	_, err = s.Client.Environments.Stop(ctx, gitpod.EnvironmentStopParams{
		EnvironmentID: gitpod.String(options.EnvironmentID),
	})
	if err != nil {
		return nil, err
	}
	return s.waitForStopped(ctx, options.EnvironmentID)
}

type DeleteEnvironmentOptions struct {
	EnvironmentID string
}

func (s *EnvironmentSandbox) Delete(ctx context.Context, options *DeleteEnvironmentOptions) error {
	_, err := s.Client.Environments.Delete(ctx, gitpod.EnvironmentDeleteParams{
		EnvironmentID: gitpod.String(options.EnvironmentID),
	})
	if err != nil {
		apierr, ok := err.(*apierror.Error)
		if ok && apierr.StatusCode == http.StatusNotFound {
			return nil
		}
		return err
	}
	return s.waitForDeleted(ctx, options.EnvironmentID)
}

func (s *EnvironmentSandbox) waitForDeleted(ctx context.Context, envID string) error {
	_, err := s.waitFor(ctx, envID, func() (*gitpod.EnvironmentGetResponseEnvironment, bool, error) {
		resp, err := s.Client.Environments.Get(ctx, gitpod.EnvironmentGetParams{
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

func (s *EnvironmentSandbox) waitForStopped(ctx context.Context, envID string) (*gitpod.EnvironmentGetResponseEnvironment, error) {
	return s.waitFor(ctx, envID, func() (*gitpod.EnvironmentGetResponseEnvironment, bool, error) {
		resp, err := s.Client.Environments.Get(ctx, gitpod.EnvironmentGetParams{
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

func (s *EnvironmentSandbox) waitForRunning(ctx context.Context, envID string) (*gitpod.EnvironmentGetResponseEnvironment, error) {
	return s.waitFor(ctx, envID, func() (*gitpod.EnvironmentGetResponseEnvironment, bool, error) {
		resp, err := s.Client.Environments.Get(ctx, gitpod.EnvironmentGetParams{
			EnvironmentID: gitpod.String(envID),
		})
		if err != nil {
			return nil, false, err
		}

		if fm := resp.Environment.Status.FailureMessage; len(fm) > 0 {
			return nil, false, fmt.Errorf("environment creation failed: %s", fm)
		}
		if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseRunning {
			return &resp.Environment, true, nil
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

func (s *EnvironmentSandbox) waitFor(ctx context.Context, envID string, fetch func() (*gitpod.EnvironmentGetResponseEnvironment, bool, error)) (*gitpod.EnvironmentGetResponseEnvironment, error) {
	env, ok, err := fetch()
	if err != nil {
		return nil, err
	}
	if ok {
		return env, nil
	}

	stream := s.Client.Events.WatchStreaming(ctx, gitpod.EventWatchParams{
		Body: gitpod.EventWatchParamsBodyEnvironmentScopeProducesEventsForTheEnvironmentItselfAllTasksTaskExecutionsAndServicesAssociatedWithThatEnvironment{
			EnvironmentID: gitpod.String(envID),
		},
	})
	defer stream.Close()

	for stream.Next() {
		resp := stream.Current()
		if resp.ResourceType == gitpod.EventWatchResponseResourceTypeResourceTypeEnvironment && resp.ResourceID == envID {
			env, ok, err := fetch()
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

func (s *EnvironmentSandbox) create(ctx context.Context, options *CreateEnvironmentOptions) (string, error) {
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
