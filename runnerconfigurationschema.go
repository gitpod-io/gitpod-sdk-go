// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
)

// RunnerConfigurationSchemaService contains methods and other services that help
// with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerConfigurationSchemaService] method instead.
type RunnerConfigurationSchemaService struct {
	Options []option.RequestOption
}

// NewRunnerConfigurationSchemaService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRunnerConfigurationSchemaService(opts ...option.RequestOption) (r *RunnerConfigurationSchemaService) {
	r = &RunnerConfigurationSchemaService{}
	r.Options = opts
	return
}

// GetRunnerConfigurationSchema retrieves the latest Runner configuration schema
func (r *RunnerConfigurationSchemaService) Get(ctx context.Context, body RunnerConfigurationSchemaGetParams, opts ...option.RequestOption) (res *RunnerConfigurationSchemaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetRunnerConfigurationSchema"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RunnerConfigurationSchema struct {
	EnvironmentClasses []RunnerConfigurationSchemaEnvironmentClass `json:"environmentClasses"`
	RunnerConfig       []RunnerConfigurationSchemaRunnerConfig     `json:"runnerConfig"`
	Scm                []RunnerConfigurationSchemaScm              `json:"scm"`
	// The schema version
	Version string                        `json:"version"`
	JSON    runnerConfigurationSchemaJSON `json:"-"`
}

// runnerConfigurationSchemaJSON contains the JSON metadata for the struct
// [RunnerConfigurationSchema]
type runnerConfigurationSchemaJSON struct {
	EnvironmentClasses apijson.Field
	RunnerConfig       apijson.Field
	Scm                apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RunnerConfigurationSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaEnvironmentClass struct {
	ID          string                                             `json:"id"`
	Bool        RunnerConfigurationSchemaEnvironmentClassesBool    `json:"bool"`
	Description string                                             `json:"description"`
	Display     RunnerConfigurationSchemaEnvironmentClassesDisplay `json:"display"`
	Enum        RunnerConfigurationSchemaEnvironmentClassesEnum    `json:"enum"`
	Int         RunnerConfigurationSchemaEnvironmentClassesInt     `json:"int"`
	Name        string                                             `json:"name"`
	Required    bool                                               `json:"required"`
	Secret      bool                                               `json:"secret"`
	String      RunnerConfigurationSchemaEnvironmentClassesString  `json:"string"`
	JSON        runnerConfigurationSchemaEnvironmentClassJSON      `json:"-"`
}

// runnerConfigurationSchemaEnvironmentClassJSON contains the JSON metadata for the
// struct [RunnerConfigurationSchemaEnvironmentClass]
type runnerConfigurationSchemaEnvironmentClassJSON struct {
	ID          apijson.Field
	Bool        apijson.Field
	Description apijson.Field
	Display     apijson.Field
	Enum        apijson.Field
	Int         apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	String      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaEnvironmentClassesBool struct {
	Default bool                                                `json:"default"`
	JSON    runnerConfigurationSchemaEnvironmentClassesBoolJSON `json:"-"`
}

// runnerConfigurationSchemaEnvironmentClassesBoolJSON contains the JSON metadata
// for the struct [RunnerConfigurationSchemaEnvironmentClassesBool]
type runnerConfigurationSchemaEnvironmentClassesBoolJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaEnvironmentClassesBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaEnvironmentClassesBoolJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaEnvironmentClassesDisplay struct {
	Default string                                                 `json:"default"`
	JSON    runnerConfigurationSchemaEnvironmentClassesDisplayJSON `json:"-"`
}

// runnerConfigurationSchemaEnvironmentClassesDisplayJSON contains the JSON
// metadata for the struct [RunnerConfigurationSchemaEnvironmentClassesDisplay]
type runnerConfigurationSchemaEnvironmentClassesDisplayJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaEnvironmentClassesDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaEnvironmentClassesDisplayJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaEnvironmentClassesEnum struct {
	Default string                                              `json:"default"`
	Values  []string                                            `json:"values"`
	JSON    runnerConfigurationSchemaEnvironmentClassesEnumJSON `json:"-"`
}

// runnerConfigurationSchemaEnvironmentClassesEnumJSON contains the JSON metadata
// for the struct [RunnerConfigurationSchemaEnvironmentClassesEnum]
type runnerConfigurationSchemaEnvironmentClassesEnumJSON struct {
	Default     apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaEnvironmentClassesEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaEnvironmentClassesEnumJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaEnvironmentClassesInt struct {
	Default int64                                              `json:"default"`
	Max     int64                                              `json:"max"`
	Min     int64                                              `json:"min"`
	JSON    runnerConfigurationSchemaEnvironmentClassesIntJSON `json:"-"`
}

// runnerConfigurationSchemaEnvironmentClassesIntJSON contains the JSON metadata
// for the struct [RunnerConfigurationSchemaEnvironmentClassesInt]
type runnerConfigurationSchemaEnvironmentClassesIntJSON struct {
	Default     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaEnvironmentClassesInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaEnvironmentClassesIntJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaEnvironmentClassesString struct {
	Default string                                                `json:"default"`
	Pattern string                                                `json:"pattern"`
	JSON    runnerConfigurationSchemaEnvironmentClassesStringJSON `json:"-"`
}

// runnerConfigurationSchemaEnvironmentClassesStringJSON contains the JSON metadata
// for the struct [RunnerConfigurationSchemaEnvironmentClassesString]
type runnerConfigurationSchemaEnvironmentClassesStringJSON struct {
	Default     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaEnvironmentClassesString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaEnvironmentClassesStringJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaRunnerConfig struct {
	ID          string                                       `json:"id"`
	Bool        RunnerConfigurationSchemaRunnerConfigBool    `json:"bool"`
	Description string                                       `json:"description"`
	Display     RunnerConfigurationSchemaRunnerConfigDisplay `json:"display"`
	Enum        RunnerConfigurationSchemaRunnerConfigEnum    `json:"enum"`
	Int         RunnerConfigurationSchemaRunnerConfigInt     `json:"int"`
	Name        string                                       `json:"name"`
	Required    bool                                         `json:"required"`
	Secret      bool                                         `json:"secret"`
	String      RunnerConfigurationSchemaRunnerConfigString  `json:"string"`
	JSON        runnerConfigurationSchemaRunnerConfigJSON    `json:"-"`
}

// runnerConfigurationSchemaRunnerConfigJSON contains the JSON metadata for the
// struct [RunnerConfigurationSchemaRunnerConfig]
type runnerConfigurationSchemaRunnerConfigJSON struct {
	ID          apijson.Field
	Bool        apijson.Field
	Description apijson.Field
	Display     apijson.Field
	Enum        apijson.Field
	Int         apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	String      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaRunnerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaRunnerConfigJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaRunnerConfigBool struct {
	Default bool                                          `json:"default"`
	JSON    runnerConfigurationSchemaRunnerConfigBoolJSON `json:"-"`
}

// runnerConfigurationSchemaRunnerConfigBoolJSON contains the JSON metadata for the
// struct [RunnerConfigurationSchemaRunnerConfigBool]
type runnerConfigurationSchemaRunnerConfigBoolJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaRunnerConfigBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaRunnerConfigBoolJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaRunnerConfigDisplay struct {
	Default string                                           `json:"default"`
	JSON    runnerConfigurationSchemaRunnerConfigDisplayJSON `json:"-"`
}

// runnerConfigurationSchemaRunnerConfigDisplayJSON contains the JSON metadata for
// the struct [RunnerConfigurationSchemaRunnerConfigDisplay]
type runnerConfigurationSchemaRunnerConfigDisplayJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaRunnerConfigDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaRunnerConfigDisplayJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaRunnerConfigEnum struct {
	Default string                                        `json:"default"`
	Values  []string                                      `json:"values"`
	JSON    runnerConfigurationSchemaRunnerConfigEnumJSON `json:"-"`
}

// runnerConfigurationSchemaRunnerConfigEnumJSON contains the JSON metadata for the
// struct [RunnerConfigurationSchemaRunnerConfigEnum]
type runnerConfigurationSchemaRunnerConfigEnumJSON struct {
	Default     apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaRunnerConfigEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaRunnerConfigEnumJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaRunnerConfigInt struct {
	Default int64                                        `json:"default"`
	Max     int64                                        `json:"max"`
	Min     int64                                        `json:"min"`
	JSON    runnerConfigurationSchemaRunnerConfigIntJSON `json:"-"`
}

// runnerConfigurationSchemaRunnerConfigIntJSON contains the JSON metadata for the
// struct [RunnerConfigurationSchemaRunnerConfigInt]
type runnerConfigurationSchemaRunnerConfigIntJSON struct {
	Default     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaRunnerConfigInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaRunnerConfigIntJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaRunnerConfigString struct {
	Default string                                          `json:"default"`
	Pattern string                                          `json:"pattern"`
	JSON    runnerConfigurationSchemaRunnerConfigStringJSON `json:"-"`
}

// runnerConfigurationSchemaRunnerConfigStringJSON contains the JSON metadata for
// the struct [RunnerConfigurationSchemaRunnerConfigString]
type runnerConfigurationSchemaRunnerConfigStringJSON struct {
	Default     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaRunnerConfigString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaRunnerConfigStringJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaScm struct {
	DefaultHosts []string                          `json:"defaultHosts"`
	Name         string                            `json:"name"`
	OAuth        RunnerConfigurationSchemaScmOAuth `json:"oauth"`
	Pat          RunnerConfigurationSchemaScmPat   `json:"pat"`
	ScmID        string                            `json:"scmId"`
	JSON         runnerConfigurationSchemaScmJSON  `json:"-"`
}

// runnerConfigurationSchemaScmJSON contains the JSON metadata for the struct
// [RunnerConfigurationSchemaScm]
type runnerConfigurationSchemaScmJSON struct {
	DefaultHosts apijson.Field
	Name         apijson.Field
	OAuth        apijson.Field
	Pat          apijson.Field
	ScmID        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaScm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaScmJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaScmOAuth struct {
	// callback_url is the URL the OAuth app will redirect to after the user has
	// authenticated.
	CallbackURL string                                `json:"callbackUrl"`
	JSON        runnerConfigurationSchemaScmOAuthJSON `json:"-"`
}

// runnerConfigurationSchemaScmOAuthJSON contains the JSON metadata for the struct
// [RunnerConfigurationSchemaScmOAuth]
type runnerConfigurationSchemaScmOAuthJSON struct {
	CallbackURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaScmOAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaScmOAuthJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaScmPat struct {
	// description is a human-readable description of the PAT.
	Description string `json:"description"`
	// docs_link is a link to the documentation on how to create a PAT for this SCM
	// system.
	DocsLink string                              `json:"docsLink"`
	JSON     runnerConfigurationSchemaScmPatJSON `json:"-"`
}

// runnerConfigurationSchemaScmPatJSON contains the JSON metadata for the struct
// [RunnerConfigurationSchemaScmPat]
type runnerConfigurationSchemaScmPatJSON struct {
	Description apijson.Field
	DocsLink    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaScmPat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaScmPatJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponse struct {
	Schema RunnerConfigurationSchema                `json:"schema"`
	JSON   runnerConfigurationSchemaGetResponseJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseJSON contains the JSON metadata for the
// struct [RunnerConfigurationSchemaGetResponse]
type runnerConfigurationSchemaGetResponseJSON struct {
	Schema      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetParams struct {
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerConfigurationSchemaGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
