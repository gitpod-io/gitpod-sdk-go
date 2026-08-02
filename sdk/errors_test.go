package sdk

import (
	"errors"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/go-cmp/cmp"
)

func TestMapError(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Err         string
		Type        string
		ConnectCode connect.Code
	}

	tests := []struct {
		Name     string
		Err      error
		Expected Expectation
	}{
		{
			Name: "unauthenticated",
			Err:  connect.NewError(connect.CodeUnauthenticated, errors.New("missing token")),
			Expected: Expectation{
				Err:         "agents.start: unauthenticated: missing token",
				Type:        "authentication_required",
				ConnectCode: connect.CodeUnauthenticated,
			},
		},
		{
			Name: "permission_denied",
			Err:  connect.NewError(connect.CodePermissionDenied, errors.New("no access")),
			Expected: Expectation{
				Err:         "ops.exec: permission_denied: no access",
				Type:        "permission_denied",
				ConnectCode: connect.CodePermissionDenied,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			operation := "agents.start"
			if tc.Name == "permission_denied" {
				operation = "ops.exec"
			}
			err := MapError(operation, tc.Err)
			got := Expectation{
				Err:         err.Error(),
				Type:        sdkErrorType(err),
				ConnectCode: connect.CodeOf(err),
			}

			if diff := cmp.Diff(tc.Expected, got); diff != "" {
				t.Errorf("MapError() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func sdkErrorType(err error) string {
	var authenticationRequired *AuthenticationRequiredError
	if errors.As(err, &authenticationRequired) {
		return "authentication_required"
	}
	var permissionDenied *PermissionDeniedError
	if errors.As(err, &permissionDenied) {
		return "permission_denied"
	}
	return "unknown"
}
