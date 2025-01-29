// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// RunnerConfigurationConfigurationSchemaService contains methods and other
// services that help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerConfigurationConfigurationSchemaService] method instead.
type RunnerConfigurationConfigurationSchemaService struct {
	Options []option.RequestOption
}

// NewRunnerConfigurationConfigurationSchemaService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRunnerConfigurationConfigurationSchemaService(opts ...option.RequestOption) (r *RunnerConfigurationConfigurationSchemaService) {
	r = &RunnerConfigurationConfigurationSchemaService{}
	r.Options = opts
	return
}

// GetRunnerConfigurationSchema retrieves the latest Runner configuration schema
func (r *RunnerConfigurationConfigurationSchemaService) New(ctx context.Context, params RunnerConfigurationConfigurationSchemaNewParams, opts ...option.RequestOption) (res *RunnerConfigurationConfigurationSchemaNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetRunnerConfigurationSchema"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type RunnerConfigurationConfigurationSchemaNewResponse struct {
	Schema RunnerConfigurationConfigurationSchemaNewResponseSchema `json:"schema"`
	JSON   runnerConfigurationConfigurationSchemaNewResponseJSON   `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseJSON contains the JSON metadata
// for the struct [RunnerConfigurationConfigurationSchemaNewResponse]
type runnerConfigurationConfigurationSchemaNewResponseJSON struct {
	Schema      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchema struct {
	EnvironmentClasses []RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass `json:"environmentClasses"`
	RunnerConfig       []RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig     `json:"runnerConfig"`
	Scm                []RunnerConfigurationConfigurationSchemaNewResponseSchemaScm              `json:"scm"`
	// The schema version
	Version string                                                      `json:"version"`
	JSON    runnerConfigurationConfigurationSchemaNewResponseSchemaJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchema]
type runnerConfigurationConfigurationSchemaNewResponseSchemaJSON struct {
	EnvironmentClasses apijson.Field
	RunnerConfig       apijson.Field
	Scm                apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass struct {
	JSON runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig struct {
	JSON runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigJSON contains
// the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaScm struct {
	DefaultHosts []string                                                        `json:"defaultHosts"`
	Name         string                                                          `json:"name"`
	OAuth        RunnerConfigurationConfigurationSchemaNewResponseSchemaScmOAuth `json:"oauth"`
	Pat          RunnerConfigurationConfigurationSchemaNewResponseSchemaScmPat   `json:"pat"`
	ScmID        string                                                          `json:"scmId"`
	JSON         runnerConfigurationConfigurationSchemaNewResponseSchemaScmJSON  `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaScmJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaScm]
type runnerConfigurationConfigurationSchemaNewResponseSchemaScmJSON struct {
	DefaultHosts apijson.Field
	Name         apijson.Field
	OAuth        apijson.Field
	Pat          apijson.Field
	ScmID        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaScm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaScmJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaScmOAuth struct {
	// callback_url is the URL the OAuth app will redirect to after the user has
	// authenticated.
	CallbackURL string                                                              `json:"callbackUrl"`
	JSON        runnerConfigurationConfigurationSchemaNewResponseSchemaScmOAuthJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaScmOAuthJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaScmOAuth]
type runnerConfigurationConfigurationSchemaNewResponseSchemaScmOAuthJSON struct {
	CallbackURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaScmOAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaScmOAuthJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaScmPat struct {
	// description is a human-readable description of the PAT.
	Description string `json:"description"`
	// docs_link is a link to the documentation on how to create a PAT for this SCM
	// system.
	DocsLink string                                                            `json:"docsLink"`
	JSON     runnerConfigurationConfigurationSchemaNewResponseSchemaScmPatJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaScmPatJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaScmPat]
type runnerConfigurationConfigurationSchemaNewResponseSchemaScmPatJSON struct {
	Description apijson.Field
	DocsLink    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaScmPat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaScmPatJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationConfigurationSchemaNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	RunnerID               param.Field[string]                                                                `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationConfigurationSchemaNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerConfigurationConfigurationSchemaNewParamsConnectProtocolVersion float64

const (
	RunnerConfigurationConfigurationSchemaNewParamsConnectProtocolVersion1 RunnerConfigurationConfigurationSchemaNewParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationConfigurationSchemaNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationConfigurationSchemaNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}
