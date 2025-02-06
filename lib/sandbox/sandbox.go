package sandbox

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/stainless-sdks/gitpod-go"
)

type EnvironmentSandbox struct {
	Client *gitpod.Client
}

func NewEnvironmentSandbox(client *gitpod.Client) *EnvironmentSandbox {
	return &EnvironmentSandbox{
		Client: client,
	}
}

type CreateEnvironmentParams struct {
	ProjectID        string
	ContextURL       string
	EnvironmentClass string
}

func (s *EnvironmentSandbox) Create(ctx context.Context, params *CreateEnvironmentParams) (res *gitpod.EnvironmentGetResponseEnvironment, err error) {
	envID, err := s.create(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.waitForRunning(ctx, envID)
}

func (s *EnvironmentSandbox) waitForRunning(ctx context.Context, envID string) (*gitpod.EnvironmentGetResponseEnvironment, error) {
	tick := time.NewTicker(500 * time.Microsecond)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-tick.C:
			resp, err := s.Client.Environments.Get(ctx, gitpod.EnvironmentGetParams{
				EnvironmentID: gitpod.String(envID),
			})
			if err != nil {
				// TODO: if transient we should retry?
				return nil, err
			}

			if fm := resp.Environment.Status.FailureMessage; len(fm) > 0 {
				return nil, fmt.Errorf("environment creation failed: %s", fm)
			}
			if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseRunning {
				return &resp.Environment, nil
			}
			if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseStopping {
				return nil, errors.New("environment creation failed: environment is stopping")
			}
			if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseStopped {
				return nil, errors.New("environment creation failed: environment is stopped")
			}
			if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseDeleting {
				return nil, errors.New("environment creation failed: environment is deleting")
			}
			if resp.Environment.Status.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusPhaseEnvironmentPhaseDeleted {
				return nil, errors.New("environment creation failed: environment is deleted")
			}
		}
	}
}

func (s *EnvironmentSandbox) create(ctx context.Context, params *CreateEnvironmentParams) (string, error) {
	if params.ProjectID != "" {
		resp, err := s.Client.Environments.NewFromProject(ctx, gitpod.EnvironmentNewFromProjectParams{
			ProjectID: gitpod.String(params.ProjectID),
		})
		if err != nil {
			return "", err
		}
		return resp.Environment.ID, nil
	}
	if params.ContextURL != "" {
		if params.EnvironmentClass == "" {
			return "", errors.New("environmentClass must be provided when contextURL is provided")
		}
		resp, err := s.Client.Environments.New(ctx, gitpod.EnvironmentNewParams{
			Spec: gitpod.F(gitpod.EnvironmentNewParamsSpec{
				DesiredPhase: gitpod.F(gitpod.EnvironmentNewParamsSpecDesiredPhaseEnvironmentPhaseRunning),
				Machine: gitpod.F(gitpod.EnvironmentNewParamsSpecMachine{
					Class: gitpod.String(params.EnvironmentClass),
				}),
				Content: gitpod.F(gitpod.EnvironmentNewParamsSpecContent{
					Initializer: gitpod.F(gitpod.EnvironmentNewParamsSpecContentInitializer{
						Specs: gitpod.F([]gitpod.EnvironmentNewParamsSpecContentInitializerSpecUnion{
							gitpod.EnvironmentNewParamsSpecContentInitializerSpecsContextURL{
								ContextURL: gitpod.F(gitpod.EnvironmentNewParamsSpecContentInitializerSpecsContextURLContextURL{
									URL: gitpod.String(params.ContextURL),
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
