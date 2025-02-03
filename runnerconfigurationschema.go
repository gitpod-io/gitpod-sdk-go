// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
	"github.com/tidwall/gjson"
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
func (r *RunnerConfigurationSchemaService) Get(ctx context.Context, params RunnerConfigurationSchemaGetParams, opts ...option.RequestOption) (res *RunnerConfigurationSchemaGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetRunnerConfigurationSchema"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type RunnerConfigurationSchemaGetResponse struct {
	Schema RunnerConfigurationSchemaGetResponseSchema `json:"schema"`
	JSON   runnerConfigurationSchemaGetResponseJSON   `json:"-"`
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

type RunnerConfigurationSchemaGetResponseSchema struct {
	EnvironmentClasses []RunnerConfigurationSchemaGetResponseSchemaEnvironmentClass `json:"environmentClasses"`
	RunnerConfig       []RunnerConfigurationSchemaGetResponseSchemaRunnerConfig     `json:"runnerConfig"`
	Scm                []RunnerConfigurationSchemaGetResponseSchemaScm              `json:"scm"`
	// The schema version
	Version string                                         `json:"version"`
	JSON    runnerConfigurationSchemaGetResponseSchemaJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaJSON contains the JSON metadata for
// the struct [RunnerConfigurationSchemaGetResponseSchema]
type runnerConfigurationSchemaGetResponseSchemaJSON struct {
	EnvironmentClasses apijson.Field
	RunnerConfig       apijson.Field
	Scm                apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaEnvironmentClass struct {
	ID string `json:"id"`
	// This field can have the runtime type of
	// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectBool].
	Bool        interface{} `json:"bool"`
	Description string      `json:"description"`
	// This field can have the runtime type of
	// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectDisplay].
	Display interface{} `json:"display"`
	// This field can have the runtime type of
	// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectEnum].
	Enum interface{} `json:"enum"`
	// This field can have the runtime type of
	// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectInt].
	Int      interface{} `json:"int"`
	Name     string      `json:"name"`
	Required bool        `json:"required"`
	Secret   bool        `json:"secret"`
	// This field can have the runtime type of
	// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectString].
	String interface{}                                                    `json:"string"`
	JSON   runnerConfigurationSchemaGetResponseSchemaEnvironmentClassJSON `json:"-"`
	union  RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesUnion
}

// runnerConfigurationSchemaGetResponseSchemaEnvironmentClassJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClass]
type runnerConfigurationSchemaGetResponseSchemaEnvironmentClassJSON struct {
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

func (r runnerConfigurationSchemaGetResponseSchemaEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationSchemaGetResponseSchemaEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationSchemaGetResponseSchemaEnvironmentClass{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject],
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject],
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject],
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject],
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject],
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject].
func (r RunnerConfigurationSchemaGetResponseSchemaEnvironmentClass) AsUnion() RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesUnion {
	return r.union
}

// Union satisfied by
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject],
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject],
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject],
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject],
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject] or
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject].
type RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesUnion interface {
	implementsRunnerConfigurationSchemaGetResponseSchemaEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject{}),
		},
	)
}

type RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject struct {
	Bool        RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectBool    `json:"bool,required"`
	ID          string                                                                    `json:"id"`
	Description string                                                                    `json:"description"`
	Display     RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectDisplay `json:"display"`
	Enum        RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectEnum    `json:"enum"`
	Int         RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectInt     `json:"int"`
	Name        string                                                                    `json:"name"`
	Required    bool                                                                      `json:"required"`
	Secret      bool                                                                      `json:"secret"`
	String      RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectString  `json:"string"`
	JSON        runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectJSON    `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectJSON contains
// the JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject]
type runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectJSON struct {
	Bool        apijson.Field
	ID          apijson.Field
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

func (r *RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObject) implementsRunnerConfigurationSchemaGetResponseSchemaEnvironmentClass() {
}

type RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectBool struct {
	Default bool                                                                       `json:"default"`
	JSON    runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectBoolJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectBoolJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectBool]
type runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectBoolJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectBoolJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectDisplay struct {
	Default string                                                                        `json:"default"`
	JSON    runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectDisplayJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectDisplayJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectDisplay]
type runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectDisplayJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectDisplayJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectEnum struct {
	Default string                                                                     `json:"default"`
	Values  []string                                                                   `json:"values"`
	JSON    runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectEnumJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectEnumJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectEnum]
type runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectEnumJSON struct {
	Default     apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectEnumJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectInt struct {
	Default int64                                                                     `json:"default"`
	Max     int64                                                                     `json:"max"`
	Min     int64                                                                     `json:"min"`
	JSON    runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectIntJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectIntJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectInt]
type runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectIntJSON struct {
	Default     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectIntJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectString struct {
	Default string                                                                       `json:"default"`
	Pattern string                                                                       `json:"pattern"`
	JSON    runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectStringJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectStringJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectString]
type runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectStringJSON struct {
	Default     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaEnvironmentClassesObjectStringJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaRunnerConfig struct {
	ID string `json:"id"`
	// This field can have the runtime type of
	// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectBool].
	Bool        interface{} `json:"bool"`
	Description string      `json:"description"`
	// This field can have the runtime type of
	// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectDisplay].
	Display interface{} `json:"display"`
	// This field can have the runtime type of
	// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectEnum].
	Enum interface{} `json:"enum"`
	// This field can have the runtime type of
	// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectInt].
	Int      interface{} `json:"int"`
	Name     string      `json:"name"`
	Required bool        `json:"required"`
	Secret   bool        `json:"secret"`
	// This field can have the runtime type of
	// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectString].
	String interface{}                                                `json:"string"`
	JSON   runnerConfigurationSchemaGetResponseSchemaRunnerConfigJSON `json:"-"`
	union  RunnerConfigurationSchemaGetResponseSchemaRunnerConfigUnion
}

// runnerConfigurationSchemaGetResponseSchemaRunnerConfigJSON contains the JSON
// metadata for the struct [RunnerConfigurationSchemaGetResponseSchemaRunnerConfig]
type runnerConfigurationSchemaGetResponseSchemaRunnerConfigJSON struct {
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

func (r runnerConfigurationSchemaGetResponseSchemaRunnerConfigJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationSchemaGetResponseSchemaRunnerConfig) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationSchemaGetResponseSchemaRunnerConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject],
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject],
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject],
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject],
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject],
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject].
func (r RunnerConfigurationSchemaGetResponseSchemaRunnerConfig) AsUnion() RunnerConfigurationSchemaGetResponseSchemaRunnerConfigUnion {
	return r.union
}

// Union satisfied by
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject],
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject],
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject],
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject],
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject] or
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject].
type RunnerConfigurationSchemaGetResponseSchemaRunnerConfigUnion interface {
	implementsRunnerConfigurationSchemaGetResponseSchemaRunnerConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationSchemaGetResponseSchemaRunnerConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject{}),
		},
	)
}

type RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject struct {
	Bool        RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectBool    `json:"bool,required"`
	ID          string                                                              `json:"id"`
	Description string                                                              `json:"description"`
	Display     RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectDisplay `json:"display"`
	Enum        RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectEnum    `json:"enum"`
	Int         RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectInt     `json:"int"`
	Name        string                                                              `json:"name"`
	Required    bool                                                                `json:"required"`
	Secret      bool                                                                `json:"secret"`
	String      RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectString  `json:"string"`
	JSON        runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectJSON    `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject]
type runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectJSON struct {
	Bool        apijson.Field
	ID          apijson.Field
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

func (r *RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObject) implementsRunnerConfigurationSchemaGetResponseSchemaRunnerConfig() {
}

type RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectBool struct {
	Default bool                                                                 `json:"default"`
	JSON    runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectBoolJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectBoolJSON contains
// the JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectBool]
type runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectBoolJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectBoolJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectDisplay struct {
	Default string                                                                  `json:"default"`
	JSON    runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectDisplayJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectDisplayJSON contains
// the JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectDisplay]
type runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectDisplayJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectDisplayJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectEnum struct {
	Default string                                                               `json:"default"`
	Values  []string                                                             `json:"values"`
	JSON    runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectEnumJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectEnumJSON contains
// the JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectEnum]
type runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectEnumJSON struct {
	Default     apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectEnumJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectInt struct {
	Default int64                                                               `json:"default"`
	Max     int64                                                               `json:"max"`
	Min     int64                                                               `json:"min"`
	JSON    runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectIntJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectIntJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectInt]
type runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectIntJSON struct {
	Default     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectIntJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectString struct {
	Default string                                                                 `json:"default"`
	Pattern string                                                                 `json:"pattern"`
	JSON    runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectStringJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectStringJSON contains
// the JSON metadata for the struct
// [RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectString]
type runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectStringJSON struct {
	Default     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaRunnerConfigObjectStringJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaScm struct {
	DefaultHosts []string                                           `json:"defaultHosts"`
	Name         string                                             `json:"name"`
	OAuth        RunnerConfigurationSchemaGetResponseSchemaScmOAuth `json:"oauth"`
	Pat          RunnerConfigurationSchemaGetResponseSchemaScmPat   `json:"pat"`
	ScmID        string                                             `json:"scmId"`
	JSON         runnerConfigurationSchemaGetResponseSchemaScmJSON  `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaScmJSON contains the JSON metadata for
// the struct [RunnerConfigurationSchemaGetResponseSchemaScm]
type runnerConfigurationSchemaGetResponseSchemaScmJSON struct {
	DefaultHosts apijson.Field
	Name         apijson.Field
	OAuth        apijson.Field
	Pat          apijson.Field
	ScmID        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaScm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaScmJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaScmOAuth struct {
	// callback_url is the URL the OAuth app will redirect to after the user has
	// authenticated.
	CallbackURL string                                                 `json:"callbackUrl"`
	JSON        runnerConfigurationSchemaGetResponseSchemaScmOAuthJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaScmOAuthJSON contains the JSON
// metadata for the struct [RunnerConfigurationSchemaGetResponseSchemaScmOAuth]
type runnerConfigurationSchemaGetResponseSchemaScmOAuthJSON struct {
	CallbackURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaScmOAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaScmOAuthJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetResponseSchemaScmPat struct {
	// description is a human-readable description of the PAT.
	Description string `json:"description"`
	// docs_link is a link to the documentation on how to create a PAT for this SCM
	// system.
	DocsLink string                                               `json:"docsLink"`
	JSON     runnerConfigurationSchemaGetResponseSchemaScmPatJSON `json:"-"`
}

// runnerConfigurationSchemaGetResponseSchemaScmPatJSON contains the JSON metadata
// for the struct [RunnerConfigurationSchemaGetResponseSchemaScmPat]
type runnerConfigurationSchemaGetResponseSchemaScmPatJSON struct {
	Description apijson.Field
	DocsLink    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationSchemaGetResponseSchemaScmPat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationSchemaGetResponseSchemaScmPatJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationSchemaGetParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[RunnerConfigurationSchemaGetParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationSchemaGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[RunnerConfigurationSchemaGetParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[RunnerConfigurationSchemaGetParamsConnect] `query:"connect"`
	Message param.Field[string]                                    `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [RunnerConfigurationSchemaGetParams]'s query parameters as
// `url.Values`.
func (r RunnerConfigurationSchemaGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type RunnerConfigurationSchemaGetParamsEncoding string

const (
	RunnerConfigurationSchemaGetParamsEncodingProto RunnerConfigurationSchemaGetParamsEncoding = "proto"
	RunnerConfigurationSchemaGetParamsEncodingJson  RunnerConfigurationSchemaGetParamsEncoding = "json"
)

func (r RunnerConfigurationSchemaGetParamsEncoding) IsKnown() bool {
	switch r {
	case RunnerConfigurationSchemaGetParamsEncodingProto, RunnerConfigurationSchemaGetParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerConfigurationSchemaGetParamsConnectProtocolVersion float64

const (
	RunnerConfigurationSchemaGetParamsConnectProtocolVersion1 RunnerConfigurationSchemaGetParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationSchemaGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationSchemaGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type RunnerConfigurationSchemaGetParamsCompression string

const (
	RunnerConfigurationSchemaGetParamsCompressionIdentity RunnerConfigurationSchemaGetParamsCompression = "identity"
	RunnerConfigurationSchemaGetParamsCompressionGzip     RunnerConfigurationSchemaGetParamsCompression = "gzip"
	RunnerConfigurationSchemaGetParamsCompressionBr       RunnerConfigurationSchemaGetParamsCompression = "br"
)

func (r RunnerConfigurationSchemaGetParamsCompression) IsKnown() bool {
	switch r {
	case RunnerConfigurationSchemaGetParamsCompressionIdentity, RunnerConfigurationSchemaGetParamsCompressionGzip, RunnerConfigurationSchemaGetParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type RunnerConfigurationSchemaGetParamsConnect string

const (
	RunnerConfigurationSchemaGetParamsConnectV1 RunnerConfigurationSchemaGetParamsConnect = "v1"
)

func (r RunnerConfigurationSchemaGetParamsConnect) IsKnown() bool {
	switch r {
	case RunnerConfigurationSchemaGetParamsConnectV1:
		return true
	}
	return false
}
