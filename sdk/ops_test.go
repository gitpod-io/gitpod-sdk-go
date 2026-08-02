package sdk

import (
	"context"
	"errors"
	"net"
	"testing"

	"connectrpc.com/connect"
	supervisorfake "github.com/gitpod-io/gitpod-sdk-go/supervisor/fake"
	supervisorv1 "github.com/gitpod-io/gitpod-sdk-go/supervisor/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestEnvironmentRunCommandPreservesStdoutAndStderr(t *testing.T) {
	t.Parallel()

	type CapturedRequests struct {
		Exec              *supervisorv1.ExecRequest
		OperationIDIsUUID bool
	}
	type Expectation struct {
		Result   *RunCommandResult
		Requests CapturedRequests
		Err      string
	}

	tests := []struct {
		Name     string
		Expected Expectation
	}{
		{
			Name: "returns_command_output_from_supervisor",
			Expected: Expectation{
				Result: &RunCommandResult{
					EnvironmentID: "env-1",
					ExitCode:      7,
					Stdout:        "stdout from environment\n",
					Stderr:        "stderr from environment\n",
				},
				Requests: CapturedRequests{
					Exec: &supervisorv1.ExecRequest{
						OperationId:      "<generated>",
						Command:          "printf stdout && printf stderr >&2",
						WorkingDirectory: "/workspace/repo",
						TimeoutSeconds:   30,
					},
					OperationIDIsUUID: true,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			opsClient := supervisorfake.NewMockEnvironmentOpsServiceClient(ctrl)
			env := &Environment{
				sdk: New(nil),
				environment: &environmentHandle{
					id: "env-1",
				},
				ops: &environmentOps{
					raw:           opsClient,
					environmentID: "env-1",
				},
			}

			var got Expectation
			opsClient.EXPECT().
				Exec(gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, req *connect.Request[supervisorv1.ExecRequest]) (*connect.Response[supervisorv1.ExecResponse], error) {
					got.Requests.Exec = req.Msg
					if _, err := uuid.Parse(req.Msg.GetOperationId()); err == nil {
						got.Requests.OperationIDIsUUID = true
						got.Requests.Exec.OperationId = "<generated>"
					}
					return connect.NewResponse(&supervisorv1.ExecResponse{
						ExitCode: tc.Expected.Result.ExitCode,
						Stdout:   tc.Expected.Result.Stdout,
						Stderr:   tc.Expected.Result.Stderr,
					}), nil
				})

			result, err := env.RunCommand(t.Context(), RunCommandOptions{
				Command:          "printf stdout && printf stderr >&2",
				WorkingDirectory: "/workspace/repo",
				TimeoutSeconds:   30,
			})
			if err != nil {
				got.Err = err.Error()
			}
			got.Result = result

			if diff := cmp.Diff(tc.Expected, got, protocmp.Transform()); diff != "" {
				t.Errorf("RunCommand() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEnvironmentConnectivityErrorIsExplicit(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Err           string
		IsUnreachable bool
	}

	tests := []struct {
		Name     string
		Expected Expectation
	}{
		{
			Name: "dns_failure_returns_environment_unreachable_error",
			Expected: Expectation{
				Err:           "ops.read_file: cannot reach environment env-1 ops service at https://unreachable.example: lookup unreachable.example: no such host",
				IsUnreachable: true,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			opsClient := supervisorfake.NewMockEnvironmentOpsServiceClient(ctrl)
			env := &Environment{
				sdk: New(nil),
				environment: &environmentHandle{
					id: "env-1",
				},
				ops: &environmentOps{
					raw:           opsClient,
					environmentID: "env-1",
					opsURL:        "https://unreachable.example",
				},
			}
			opsClient.EXPECT().
				ReadFile(gomock.Any(), gomock.Any()).
				Return(nil, &net.DNSError{Name: "unreachable.example", Err: "no such host"})

			var got Expectation
			_, err := env.ReadFile(t.Context(), "/workspace/README.md", ReadFileOptions{})
			if err != nil {
				got.Err = err.Error()
				var unreachable *EnvironmentUnreachableError
				got.IsUnreachable = errors.As(err, &unreachable)
			}

			if diff := cmp.Diff(tc.Expected, got); diff != "" {
				t.Errorf("ReadFile() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
