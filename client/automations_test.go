package client_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/gitpod-io/flex-sdk-go/client"
	v1 "github.com/gitpod-io/flex-sdk-go/v1"
)

func TestReadAutomationsFile(t *testing.T) {
	t.Parallel()
	type Expectation struct {
		Error           string
		AutomationsFile *v1.AutomationsFile
	}
	tests := []struct {
		Name        string
		Expectation Expectation
		Input       string
	}{
		{
			Name: "empty",
			Expectation: Expectation{
				Error: "EOF",
			},
		},
		{
			Name:  "JSON",
			Input: `{"services": {"ref1": {"name": "test"}}}`,
			Expectation: Expectation{
				AutomationsFile: &v1.AutomationsFile{
					Services: map[string]*v1.AutomationsFile_Service{
						"ref1": {
							Name: "test",
						},
					},
				},
			},
		},
		{
			Name: "all fields",
			Input: `services:
  ref1:
    name: test
    description: hello world
    commands:
      start: start
      stop: stop
      ready: ready
`,
			Expectation: Expectation{
				AutomationsFile: &v1.AutomationsFile{
					Services: map[string]*v1.AutomationsFile_Service{
						"ref1": {
							Name:        "test",
							Description: "hello world",
							Commands: &v1.ServiceSpec_Commands{
								Start: "start",
								Stop:  "stop",
								Ready: "ready",
							},
						},
					},
				},
			},
		},
		{
			Name: "unknown field",
			Input: `services:
  ref1:
    name: test
    description: hello world
    commands:
      start: start
    trigger:
      - manual
`,
			Expectation: Expectation{
				Error: "proto: (line 1:93): unknown field \"trigger\"",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			var act Expectation

			res, err := client.ReadAutomationsFile(strings.NewReader(test.Input))
			if err != nil {
				act.Error = strings.ReplaceAll(err.Error(), " ", " ")
			}
			act.AutomationsFile = res

			if diff := cmp.Diff(test.Expectation, act, protocmp.Transform()); diff != "" {
				t.Errorf("ReadAutomationsFile() mismatch (-want  got):\n%s", diff)
			}
		})
	}
}
