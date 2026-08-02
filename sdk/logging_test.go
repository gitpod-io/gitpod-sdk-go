package sdk

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"strings"
	"sync"
	"testing"

	"connectrpc.com/connect"
	rawclient "github.com/gitpod-io/gitpod-sdk-go/internal/raw"
	v1 "github.com/gitpod-io/gitpod-sdk-go/v1"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

type capturedSlogRecord struct {
	Message string
	Attrs   map[string]string
}

type captureSlogSink struct {
	mu      sync.Mutex
	records []capturedSlogRecord
}

type captureSlogHandler struct {
	sink  *captureSlogSink
	attrs []slog.Attr
}

func newCaptureSlogHandler() *captureSlogHandler {
	return &captureSlogHandler{sink: &captureSlogSink{}}
}

func (h *captureSlogHandler) Enabled(context.Context, slog.Level) bool {
	return true
}

func (h *captureSlogHandler) Handle(_ context.Context, record slog.Record) error {
	attrs := make(map[string]string)
	for _, attr := range h.attrs {
		attrs[attr.Key] = attr.Value.Resolve().String()
	}
	record.Attrs(func(attr slog.Attr) bool {
		attrs[attr.Key] = attr.Value.Resolve().String()
		return true
	})

	h.sink.mu.Lock()
	defer h.sink.mu.Unlock()
	h.sink.records = append(h.sink.records, capturedSlogRecord{
		Message: record.Message,
		Attrs:   attrs,
	})
	return nil
}

func (h *captureSlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	next := &captureSlogHandler{
		sink:  h.sink,
		attrs: append([]slog.Attr{}, h.attrs...),
	}
	next.attrs = append(next.attrs, attrs...)
	return next
}

func (h *captureSlogHandler) WithGroup(string) slog.Handler {
	return &captureSlogHandler{
		sink:  h.sink,
		attrs: append([]slog.Attr{}, h.attrs...),
	}
}

func TestSDKHumanLogHandlerFormatsReadableOutput(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		Output string
	}

	tests := []struct {
		NoColor  bool
		Name     string
		Expected Expectation
	}{
		{
			NoColor: true,
			Name:    "formats_debug_record_without_ansi_when_color_disabled",
			Expected: Expectation{
				Output: "DBG creating environment operation=environments.create environment_id=env-1 count=2 group.phase=ENVIRONMENT_PHASE_RUNNING err=boom\n",
			},
		},
		{
			Name: "formats_debug_record_with_color_and_bold_text",
			Expected: Expectation{
				Output: "\x1b[1;96mDBG\x1b[0m \x1b[1mcreating environment\x1b[22m " +
					"\x1b[2;95moperation=\x1b[22menvironments.create\x1b[0m " +
					"\x1b[2;94menvironment_id=\x1b[22menv-1\x1b[0m " +
					"\x1b[2;92mcount=\x1b[22m2\x1b[0m " +
					"\x1b[2;93mgroup.phase=\x1b[22mENVIRONMENT_PHASE_RUNNING\x1b[0m " +
					"\x1b[2;91merr=\x1b[22m\x1b[1mboom\x1b[22m\x1b[0m\n",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			var output bytes.Buffer
			logger := slog.New(NewHumanLogHandler(&output, &HumanLogHandlerOptions{Level: slog.LevelDebug, NoColor: tc.NoColor}))
			logger.DebugContext(t.Context(), "creating environment",
				slog.String("operation", "environments.create"),
				slog.String("empty", ""),
				slog.String("environment_id", "env-1"),
				slog.Int("count", 2),
				slog.Group("group", slog.String("phase", "ENVIRONMENT_PHASE_RUNNING")),
				slog.Any("err", errors.New("boom")),
			)

			got := Expectation{Output: output.String()}
			if diff := cmp.Diff(tc.Expected, got); diff != "" {
				t.Errorf("NewHumanLogHandler() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func (h *captureSlogHandler) records() []capturedSlogRecord {
	h.sink.mu.Lock()
	defer h.sink.mu.Unlock()
	return append([]capturedSlogRecord{}, h.sink.records...)
}

func TestSDKLoggerRecordsEnvironmentLifecycleWithoutURLSecrets(t *testing.T) {
	t.Parallel()

	type Expectation struct {
		EnvironmentID          string
		RequestContextURL      string
		LogMessages            []string
		LoggedContextURLs      []string
		ContainsSensitiveValue bool
		Err                    string
	}

	tests := []struct {
		Name     string
		Expected Expectation
	}{
		{
			Name: "logs_lifecycle_and_sanitizes_context_urls",
			Expected: Expectation{
				EnvironmentID:     "env-1",
				RequestContextURL: "https://sdk-user:sdk-secret@github.com/acme/api?token=topsecret#readme",
				LogMessages: []string{
					"creating environment",
					"resolving context URL",
					"resolved context URL",
					"resolving runner for environment class",
					"resolved runner for environment class",
					"checking SCM authentication",
					"SCM authentication available",
					"selected environment class",
					"sending environment create request",
					"environment created",
					"waiting for environment to run",
					"environment phase changed",
				},
				LoggedContextURLs: []string{
					"https://github.com/acme/api",
					"https://github.com/acme/api",
					"https://github.com/acme/api",
					"https://github.com/acme/api",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mp := rawclient.NewMock(ctrl)
			logs := newCaptureSlogHandler()
			sdk := New(mp.Client(), WithLogger(slog.New(logs)))
			contextURL := tc.Expected.RequestContextURL

			var got Expectation
			parse := mp.RunnerService.EXPECT().
				ParseContextURL(gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, req *connect.Request[v1.ParseContextURLRequest]) (*connect.Response[v1.ParseContextURLResponse], error) {
					got.RequestContextURL = req.Msg.GetContextUrl()
					return connect.NewResponse(&v1.ParseContextURLResponse{
						OriginalContextUrl:            contextURL,
						RecommendedEnvironmentClasses: []string{"class-recommended"},
						Git: &v1.ParseContextURLResponse_GitContext{
							Host:     "github.com",
							Owner:    "acme",
							Repo:     "api",
							CloneUrl: "https://github.com/acme/api.git",
						},
					}), nil
				})
			class := mp.RunnerConfigurationService.EXPECT().
				GetEnvironmentClass(gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, _ *connect.Request[v1.GetEnvironmentClassRequest]) (*connect.Response[v1.GetEnvironmentClassResponse], error) {
					return connect.NewResponse(&v1.GetEnvironmentClassResponse{
						EnvironmentClass: &v1.EnvironmentClass{
							Id:       "class-recommended",
							RunnerId: "runner-1",
						},
					}), nil
				}).
				After(parse)
			auth := mp.RunnerService.EXPECT().
				CheckAuthenticationForHost(gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, _ *connect.Request[v1.CheckAuthenticationForHostRequest]) (*connect.Response[v1.CheckAuthenticationForHostResponse], error) {
					return connect.NewResponse(&v1.CheckAuthenticationForHostResponse{
						Authenticated: true,
					}), nil
				}).
				After(class)
			create := mp.EnvironmentService.EXPECT().
				CreateEnvironment(gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, _ *connect.Request[v1.CreateEnvironmentRequest]) (*connect.Response[v1.CreateEnvironmentResponse], error) {
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
				Return(connect.NewResponse(&v1.GetEnvironmentResponse{
					Environment: &v1.Environment{
						Id:     "env-1",
						Status: &v1.EnvironmentStatus{Phase: v1.EnvironmentPhase_ENVIRONMENT_PHASE_RUNNING},
					},
				}), nil).
				After(create)

			env, err := sdk.Environments().Create(t.Context(), CreateEnvironmentOptions{
				ContextURL: contextURL,
			})
			if err != nil {
				got.Err = err.Error()
			} else {
				got.EnvironmentID = env.ID()
			}

			for _, record := range logs.records() {
				got.LogMessages = append(got.LogMessages, record.Message)
				if contextURL := record.Attrs["context_url"]; contextURL != "" {
					got.LoggedContextURLs = append(got.LoggedContextURLs, contextURL)
				}
				for _, value := range record.Attrs {
					if strings.Contains(value, "sdk-secret") || strings.Contains(value, "topsecret") {
						got.ContainsSensitiveValue = true
					}
				}
			}

			if diff := cmp.Diff(tc.Expected, got); diff != "" {
				t.Errorf("Create() logging mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
