package lib

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gitpod-go"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"

	cryptorand "crypto/rand"
)

type Sandbox interface {
	io.Closer
	Exec(ctx context.Context, command string) (io.ReadCloser, error)
}

func NewSandbox(ctx context.Context, client *gitpod.Client, environmentID string) (Sandbox, error) {
	logTokenResp, err := client.Environments.NewLogsToken(ctx, gitpod.EnvironmentNewLogsTokenParams{
		EnvironmentID: gitpod.String(environmentID),
	})
	if err != nil {
		return nil, err
	}

	taskReference := fmt.Sprintf("sandbox-%d", rand.Uint64())
	resp, err := client.Environments.Automations.Tasks.New(ctx, gitpod.EnvironmentAutomationTaskNewParams{
		EnvironmentID: gitpod.String(environmentID),
		Metadata: gitpod.F(gitpod.EnvironmentAutomationTaskNewParamsMetadata{
			Reference:   gitpod.String(taskReference),
			Name:        gitpod.String("Sandbox Task"),
			Description: gitpod.String("Sandbox task"),
		}),
	})
	if err != nil {
		return nil, nil
	}

	/*signer, err := generateSSHKey()
	if err != nil {
		return nil, err
	}
	_, err = client.Environments.Update(ctx, gitpod.EnvironmentUpdateParams{
		Body: gitpod.EnvironmentUpdateParamsBodySpec{
			Spec: ,
		},
	})
	if err != nil {
		return nil, err
	}*/

	return &sandbox{
		taskReference:  taskReference,
		taskID:         resp.Task.ID,
		environmentID:  environmentID,
		logAccessToken: logTokenResp.AccessToken,
		client:         client,
	}, nil
}

type sandbox struct {
	taskReference  string
	taskID         string
	environmentID  string
	logAccessToken string
	client         *gitpod.Client
	sshSigner      ssh.Signer
}

func (s *sandbox) Close() error {
	_, err := s.client.Environments.Automations.Tasks.Delete(context.Background(), gitpod.EnvironmentAutomationTaskDeleteParams{
		ID: gitpod.String(s.taskID),
	})
	return err
}

func (s *sandbox) Exec(ctx context.Context, command string) (io.ReadCloser, error) {
	_, err := s.client.Environments.Automations.Tasks.Update(ctx, gitpod.EnvironmentAutomationTaskUpdateParams{
		ID: gitpod.String(s.taskID),
		Spec: gitpod.F[gitpod.EnvironmentAutomationTaskUpdateParamsSpecUnion](gitpod.EnvironmentAutomationTaskUpdateParamsSpecCommand{
			Command: gitpod.String(command),
		}),
	})
	if err != nil {
		return nil, err
	}
	startResp, err := s.client.Environments.Automations.Tasks.Start(ctx, gitpod.EnvironmentAutomationTaskStartParams{
		ID: gitpod.String(s.taskID),
	})
	if err != nil {
		return nil, err
	}
	taskExecutionID := startResp.TaskExecution.ID
	taskExecution, err := WaitFor(ctx, s.client, s.environmentID, gitpod.EventWatchResponseResourceTypeResourceTypeTaskExecution, taskExecutionID, func() (*gitpod.EnvironmentAutomationTaskExecutionGetResponseTaskExecution, bool, error) {
		resp, err := s.client.Environments.Automations.Tasks.Executions.Get(ctx, gitpod.EnvironmentAutomationTaskExecutionGetParams{
			ID: gitpod.String(taskExecutionID),
		})
		if err != nil {
			return nil, false, err
		}
		if resp.TaskExecution.Status.LogURL != "" {
			return &resp.TaskExecution, true, nil
		}
		return nil, false, nil
	})
	if err != nil {
		return nil, err
	}
	logURL := taskExecution.Status.LogURL

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, logURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.logAccessToken))
	logResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return logResp.Body, nil
}

type FS struct {
	*sftp.Client
	sshConn *ssh.Client
}

func (fs *FS) Close() error {
	return fs.Client.Close()
}

func (s *sandbox) FS(ctx context.Context) (*FS, error) {
	env, err := waitForRunning(ctx, s.client, s.environmentID, func(env *gitpod.EnvironmentGetResponseEnvironment) (bool, error) {
		keySpecified := false
		for _, key := range env.Spec.SSHPublicKeys {
			if key.ID == s.taskReference {
				keySpecified = true
				break
			}
		}
		if !keySpecified {
			return false, fmt.Errorf("public key not specified")
		}
		keyApplied := false
		for _, key := range env.Status.SSHPublicKeys {
			if key.ID == s.taskReference {
				if key.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseFailed {
					return false, fmt.Errorf("public key failed to apply")
				}
				if key.Phase == gitpod.EnvironmentGetResponseEnvironmentStatusSSHPublicKeysPhaseContentPhaseReady {
					keyApplied = true
					break
				}
			}
		}
		return keyApplied && env.Status.EnvironmentURLs.SSH.URL == "", nil
	})
	if err != nil {
		return nil, err
	}
	sshURL := env.Status.EnvironmentURLs.SSH.URL
	parsedURL, err := url.Parse(sshURL)
	if err != nil {
		return nil, err
	}

	cfg := ssh.ClientConfig{
		User: "gitpod_devcontainer",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(s.sshSigner),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	clnt, err := ssh.Dial("tcp", parsedURL.Host, &cfg)
	if err != nil {
		return nil, err
	}
	fs, err := sftp.NewClient(clnt)
	if err != nil {
		return nil, err
	}
	return &FS{
		Client:  fs,
		sshConn: clnt,
	}, nil
}

func generateSSHKey() (privateKeyPEM ssh.Signer, err error) {
	_, privateKey, err := ed25519.GenerateKey(cryptorand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key pair: %v", err)
	}

	signer, err := ssh.NewSignerFromKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create signer from private key: %v", err)
	}

	return signer, nil
}
