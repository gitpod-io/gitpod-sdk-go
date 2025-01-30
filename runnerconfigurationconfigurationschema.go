// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
	"github.com/tidwall/gjson"
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
func (r *RunnerConfigurationConfigurationSchemaService) New(ctx context.Context, body RunnerConfigurationConfigurationSchemaNewParams, opts ...option.RequestOption) (res *RunnerConfigurationConfigurationSchemaNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetRunnerConfigurationSchema"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetRunnerConfigurationSchema retrieves the latest Runner configuration schema
func (r *RunnerConfigurationConfigurationSchemaService) Get(ctx context.Context, body RunnerConfigurationConfigurationSchemaGetParams, opts ...option.RequestOption) (res *RunnerConfigurationConfigurationSchemaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetRunnerConfigurationSchema"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
	ID string `json:"id"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolBool].
	Bool        interface{} `json:"bool"`
	Description string      `json:"description"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayDisplay].
	Display interface{} `json:"display"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumEnum].
	Enum interface{} `json:"enum"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntInt].
	Int      interface{} `json:"int"`
	Name     string      `json:"name"`
	Required bool        `json:"required"`
	Secret   bool        `json:"secret"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringString].
	String interface{}                                                                 `json:"string"`
	JSON   runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassJSON `json:"-"`
	union  RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesUnion
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassJSON struct {
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

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBool],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplay],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnum],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesInt],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesString].
func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass) AsUnion() RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesUnion {
	return r.union
}

// Union satisfied by
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBool],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplay],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnum],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesInt]
// or
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesString].
type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesUnion interface {
	implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBool{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplay{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnum{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesInt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesString{}),
		},
	)
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBool struct {
	Bool        RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolBool `json:"bool,required"`
	ID          string                                                                            `json:"id"`
	Description string                                                                            `json:"description"`
	Name        string                                                                            `json:"name"`
	Required    bool                                                                              `json:"required"`
	Secret      bool                                                                              `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBool]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolJSON struct {
	Bool        apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBool) implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass() {
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolBool struct {
	Default bool                                                                                  `json:"default"`
	JSON    runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolBoolJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolBoolJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolBool]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolBoolJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesBoolBoolJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplay struct {
	Display     RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayDisplay `json:"display,required"`
	ID          string                                                                                  `json:"id"`
	Description string                                                                                  `json:"description"`
	Name        string                                                                                  `json:"name"`
	Required    bool                                                                                    `json:"required"`
	Secret      bool                                                                                    `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayJSON    `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplay]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayJSON struct {
	Display     apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplay) implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass() {
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayDisplay struct {
	Default string                                                                                      `json:"default"`
	JSON    runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayDisplayJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayDisplayJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayDisplay]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayDisplayJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesDisplayDisplayJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnum struct {
	Enum        RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumEnum `json:"enum,required"`
	ID          string                                                                            `json:"id"`
	Description string                                                                            `json:"description"`
	Name        string                                                                            `json:"name"`
	Required    bool                                                                              `json:"required"`
	Secret      bool                                                                              `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnum]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumJSON struct {
	Enum        apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnum) implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass() {
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumEnum struct {
	Default string                                                                                `json:"default"`
	Values  []string                                                                              `json:"values"`
	JSON    runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumEnumJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumEnumJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumEnum]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumEnumJSON struct {
	Default     apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesEnumEnumJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesInt struct {
	Int         RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntInt  `json:"int,required"`
	ID          string                                                                           `json:"id"`
	Description string                                                                           `json:"description"`
	Name        string                                                                           `json:"name"`
	Required    bool                                                                             `json:"required"`
	Secret      bool                                                                             `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesInt]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntJSON struct {
	Int         apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesInt) implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass() {
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntInt struct {
	Default int64                                                                               `json:"default"`
	Max     int64                                                                               `json:"max"`
	Min     int64                                                                               `json:"min"`
	JSON    runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntIntJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntIntJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntInt]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntIntJSON struct {
	Default     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesIntIntJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesString struct {
	String      RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringString `json:"string,required"`
	ID          string                                                                                `json:"id"`
	Description string                                                                                `json:"description"`
	Name        string                                                                                `json:"name"`
	Required    bool                                                                                  `json:"required"`
	Secret      bool                                                                                  `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringJSON   `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesString]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringJSON struct {
	String      apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesString) implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClass() {
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringString struct {
	Default string                                                                                    `json:"default"`
	Pattern string                                                                                    `json:"pattern"`
	JSON    runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringStringJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringStringJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringString]
type runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringStringJSON struct {
	Default     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaEnvironmentClassesStringStringJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig struct {
	ID string `json:"id"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolBool].
	Bool        interface{} `json:"bool"`
	Description string      `json:"description"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayDisplay].
	Display interface{} `json:"display"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumEnum].
	Enum interface{} `json:"enum"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntInt].
	Int      interface{} `json:"int"`
	Name     string      `json:"name"`
	Required bool        `json:"required"`
	Secret   bool        `json:"secret"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringString].
	String interface{}                                                             `json:"string"`
	JSON   runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigJSON `json:"-"`
	union  RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigUnion
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigJSON contains
// the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigJSON struct {
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

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBool],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplay],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnum],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigInt],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigString].
func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig) AsUnion() RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigUnion {
	return r.union
}

// Union satisfied by
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBool],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplay],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnum],
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigInt] or
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigString].
type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigUnion interface {
	implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBool{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplay{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnum{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigInt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigString{}),
		},
	)
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBool struct {
	Bool        RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolBool `json:"bool,required"`
	ID          string                                                                      `json:"id"`
	Description string                                                                      `json:"description"`
	Name        string                                                                      `json:"name"`
	Required    bool                                                                        `json:"required"`
	Secret      bool                                                                        `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBool]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolJSON struct {
	Bool        apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBool) implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig() {
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolBool struct {
	Default bool                                                                            `json:"default"`
	JSON    runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolBoolJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolBoolJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolBool]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolBoolJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigBoolBoolJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplay struct {
	Display     RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayDisplay `json:"display,required"`
	ID          string                                                                            `json:"id"`
	Description string                                                                            `json:"description"`
	Name        string                                                                            `json:"name"`
	Required    bool                                                                              `json:"required"`
	Secret      bool                                                                              `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayJSON    `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplay]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayJSON struct {
	Display     apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplay) implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig() {
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayDisplay struct {
	Default string                                                                                `json:"default"`
	JSON    runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayDisplayJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayDisplayJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayDisplay]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayDisplayJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigDisplayDisplayJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnum struct {
	Enum        RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumEnum `json:"enum,required"`
	ID          string                                                                      `json:"id"`
	Description string                                                                      `json:"description"`
	Name        string                                                                      `json:"name"`
	Required    bool                                                                        `json:"required"`
	Secret      bool                                                                        `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnum]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumJSON struct {
	Enum        apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnum) implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig() {
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumEnum struct {
	Default string                                                                          `json:"default"`
	Values  []string                                                                        `json:"values"`
	JSON    runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumEnumJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumEnumJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumEnum]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumEnumJSON struct {
	Default     apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigEnumEnumJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigInt struct {
	Int         RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntInt  `json:"int,required"`
	ID          string                                                                     `json:"id"`
	Description string                                                                     `json:"description"`
	Name        string                                                                     `json:"name"`
	Required    bool                                                                       `json:"required"`
	Secret      bool                                                                       `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigInt]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntJSON struct {
	Int         apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigInt) implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig() {
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntInt struct {
	Default int64                                                                         `json:"default"`
	Max     int64                                                                         `json:"max"`
	Min     int64                                                                         `json:"min"`
	JSON    runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntIntJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntIntJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntInt]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntIntJSON struct {
	Default     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigIntIntJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigString struct {
	String      RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringString `json:"string,required"`
	ID          string                                                                          `json:"id"`
	Description string                                                                          `json:"description"`
	Name        string                                                                          `json:"name"`
	Required    bool                                                                            `json:"required"`
	Secret      bool                                                                            `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringJSON   `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigString]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringJSON struct {
	String      apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigString) implementsRunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfig() {
}

type RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringString struct {
	Default string                                                                              `json:"default"`
	Pattern string                                                                              `json:"pattern"`
	JSON    runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringStringJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringStringJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringString]
type runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringStringJSON struct {
	Default     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaNewResponseSchemaRunnerConfigStringStringJSON) RawJSON() string {
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

type RunnerConfigurationConfigurationSchemaGetResponse struct {
	Schema RunnerConfigurationConfigurationSchemaGetResponseSchema `json:"schema"`
	JSON   runnerConfigurationConfigurationSchemaGetResponseJSON   `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseJSON contains the JSON metadata
// for the struct [RunnerConfigurationConfigurationSchemaGetResponse]
type runnerConfigurationConfigurationSchemaGetResponseJSON struct {
	Schema      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchema struct {
	EnvironmentClasses []RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass `json:"environmentClasses"`
	RunnerConfig       []RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig     `json:"runnerConfig"`
	Scm                []RunnerConfigurationConfigurationSchemaGetResponseSchemaScm              `json:"scm"`
	// The schema version
	Version string                                                      `json:"version"`
	JSON    runnerConfigurationConfigurationSchemaGetResponseSchemaJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchema]
type runnerConfigurationConfigurationSchemaGetResponseSchemaJSON struct {
	EnvironmentClasses apijson.Field
	RunnerConfig       apijson.Field
	Scm                apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass struct {
	ID string `json:"id"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolBool].
	Bool        interface{} `json:"bool"`
	Description string      `json:"description"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayDisplay].
	Display interface{} `json:"display"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumEnum].
	Enum interface{} `json:"enum"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntInt].
	Int      interface{} `json:"int"`
	Name     string      `json:"name"`
	Required bool        `json:"required"`
	Secret   bool        `json:"secret"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringString].
	String interface{}                                                                 `json:"string"`
	JSON   runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassJSON `json:"-"`
	union  RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesUnion
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass]
type runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassJSON struct {
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

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBool],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplay],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnum],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesInt],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesString].
func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass) AsUnion() RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesUnion {
	return r.union
}

// Union satisfied by
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBool],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplay],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnum],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesInt]
// or
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesString].
type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesUnion interface {
	implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBool{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplay{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnum{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesInt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesString{}),
		},
	)
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBool struct {
	Bool        RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolBool `json:"bool,required"`
	ID          string                                                                            `json:"id"`
	Description string                                                                            `json:"description"`
	Name        string                                                                            `json:"name"`
	Required    bool                                                                              `json:"required"`
	Secret      bool                                                                              `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBool]
type runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolJSON struct {
	Bool        apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBool) implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass() {
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolBool struct {
	Default bool                                                                                  `json:"default"`
	JSON    runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolBoolJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolBoolJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolBool]
type runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolBoolJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesBoolBoolJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplay struct {
	Display     RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayDisplay `json:"display,required"`
	ID          string                                                                                  `json:"id"`
	Description string                                                                                  `json:"description"`
	Name        string                                                                                  `json:"name"`
	Required    bool                                                                                    `json:"required"`
	Secret      bool                                                                                    `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayJSON    `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplay]
type runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayJSON struct {
	Display     apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplay) implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass() {
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayDisplay struct {
	Default string                                                                                      `json:"default"`
	JSON    runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayDisplayJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayDisplayJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayDisplay]
type runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayDisplayJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesDisplayDisplayJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnum struct {
	Enum        RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumEnum `json:"enum,required"`
	ID          string                                                                            `json:"id"`
	Description string                                                                            `json:"description"`
	Name        string                                                                            `json:"name"`
	Required    bool                                                                              `json:"required"`
	Secret      bool                                                                              `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnum]
type runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumJSON struct {
	Enum        apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnum) implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass() {
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumEnum struct {
	Default string                                                                                `json:"default"`
	Values  []string                                                                              `json:"values"`
	JSON    runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumEnumJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumEnumJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumEnum]
type runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumEnumJSON struct {
	Default     apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesEnumEnumJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesInt struct {
	Int         RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntInt  `json:"int,required"`
	ID          string                                                                           `json:"id"`
	Description string                                                                           `json:"description"`
	Name        string                                                                           `json:"name"`
	Required    bool                                                                             `json:"required"`
	Secret      bool                                                                             `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesInt]
type runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntJSON struct {
	Int         apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesInt) implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass() {
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntInt struct {
	Default int64                                                                               `json:"default"`
	Max     int64                                                                               `json:"max"`
	Min     int64                                                                               `json:"min"`
	JSON    runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntIntJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntIntJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntInt]
type runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntIntJSON struct {
	Default     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesIntIntJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesString struct {
	String      RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringString `json:"string,required"`
	ID          string                                                                                `json:"id"`
	Description string                                                                                `json:"description"`
	Name        string                                                                                `json:"name"`
	Required    bool                                                                                  `json:"required"`
	Secret      bool                                                                                  `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringJSON   `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesString]
type runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringJSON struct {
	String      apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesString) implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClass() {
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringString struct {
	Default string                                                                                    `json:"default"`
	Pattern string                                                                                    `json:"pattern"`
	JSON    runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringStringJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringStringJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringString]
type runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringStringJSON struct {
	Default     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaEnvironmentClassesStringStringJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig struct {
	ID string `json:"id"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolBool].
	Bool        interface{} `json:"bool"`
	Description string      `json:"description"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayDisplay].
	Display interface{} `json:"display"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumEnum].
	Enum interface{} `json:"enum"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntInt].
	Int      interface{} `json:"int"`
	Name     string      `json:"name"`
	Required bool        `json:"required"`
	Secret   bool        `json:"secret"`
	// This field can have the runtime type of
	// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringString].
	String interface{}                                                             `json:"string"`
	JSON   runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigJSON `json:"-"`
	union  RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigUnion
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigJSON contains
// the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig]
type runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigJSON struct {
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

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBool],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplay],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnum],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigInt],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigString].
func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig) AsUnion() RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigUnion {
	return r.union
}

// Union satisfied by
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBool],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplay],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnum],
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigInt] or
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigString].
type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigUnion interface {
	implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBool{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplay{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnum{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigInt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigString{}),
		},
	)
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBool struct {
	Bool        RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolBool `json:"bool,required"`
	ID          string                                                                      `json:"id"`
	Description string                                                                      `json:"description"`
	Name        string                                                                      `json:"name"`
	Required    bool                                                                        `json:"required"`
	Secret      bool                                                                        `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBool]
type runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolJSON struct {
	Bool        apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBool) implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig() {
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolBool struct {
	Default bool                                                                            `json:"default"`
	JSON    runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolBoolJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolBoolJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolBool]
type runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolBoolJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigBoolBoolJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplay struct {
	Display     RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayDisplay `json:"display,required"`
	ID          string                                                                            `json:"id"`
	Description string                                                                            `json:"description"`
	Name        string                                                                            `json:"name"`
	Required    bool                                                                              `json:"required"`
	Secret      bool                                                                              `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayJSON    `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplay]
type runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayJSON struct {
	Display     apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplay) implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig() {
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayDisplay struct {
	Default string                                                                                `json:"default"`
	JSON    runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayDisplayJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayDisplayJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayDisplay]
type runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayDisplayJSON struct {
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayDisplay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigDisplayDisplayJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnum struct {
	Enum        RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumEnum `json:"enum,required"`
	ID          string                                                                      `json:"id"`
	Description string                                                                      `json:"description"`
	Name        string                                                                      `json:"name"`
	Required    bool                                                                        `json:"required"`
	Secret      bool                                                                        `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnum]
type runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumJSON struct {
	Enum        apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnum) implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig() {
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumEnum struct {
	Default string                                                                          `json:"default"`
	Values  []string                                                                        `json:"values"`
	JSON    runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumEnumJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumEnumJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumEnum]
type runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumEnumJSON struct {
	Default     apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumEnum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigEnumEnumJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigInt struct {
	Int         RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntInt  `json:"int,required"`
	ID          string                                                                     `json:"id"`
	Description string                                                                     `json:"description"`
	Name        string                                                                     `json:"name"`
	Required    bool                                                                       `json:"required"`
	Secret      bool                                                                       `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigInt]
type runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntJSON struct {
	Int         apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigInt) implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig() {
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntInt struct {
	Default int64                                                                         `json:"default"`
	Max     int64                                                                         `json:"max"`
	Min     int64                                                                         `json:"min"`
	JSON    runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntIntJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntIntJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntInt]
type runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntIntJSON struct {
	Default     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntInt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigIntIntJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigString struct {
	String      RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringString `json:"string,required"`
	ID          string                                                                          `json:"id"`
	Description string                                                                          `json:"description"`
	Name        string                                                                          `json:"name"`
	Required    bool                                                                            `json:"required"`
	Secret      bool                                                                            `json:"secret"`
	JSON        runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringJSON   `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigString]
type runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringJSON struct {
	String      apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigString) implementsRunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfig() {
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringString struct {
	Default string                                                                              `json:"default"`
	Pattern string                                                                              `json:"pattern"`
	JSON    runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringStringJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringStringJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringString]
type runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringStringJSON struct {
	Default     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaRunnerConfigStringStringJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaScm struct {
	DefaultHosts []string                                                        `json:"defaultHosts"`
	Name         string                                                          `json:"name"`
	OAuth        RunnerConfigurationConfigurationSchemaGetResponseSchemaScmOAuth `json:"oauth"`
	Pat          RunnerConfigurationConfigurationSchemaGetResponseSchemaScmPat   `json:"pat"`
	ScmID        string                                                          `json:"scmId"`
	JSON         runnerConfigurationConfigurationSchemaGetResponseSchemaScmJSON  `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaScmJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaScm]
type runnerConfigurationConfigurationSchemaGetResponseSchemaScmJSON struct {
	DefaultHosts apijson.Field
	Name         apijson.Field
	OAuth        apijson.Field
	Pat          apijson.Field
	ScmID        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaScm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaScmJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaScmOAuth struct {
	// callback_url is the URL the OAuth app will redirect to after the user has
	// authenticated.
	CallbackURL string                                                              `json:"callbackUrl"`
	JSON        runnerConfigurationConfigurationSchemaGetResponseSchemaScmOAuthJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaScmOAuthJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaScmOAuth]
type runnerConfigurationConfigurationSchemaGetResponseSchemaScmOAuthJSON struct {
	CallbackURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaScmOAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaScmOAuthJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaGetResponseSchemaScmPat struct {
	// description is a human-readable description of the PAT.
	Description string `json:"description"`
	// docs_link is a link to the documentation on how to create a PAT for this SCM
	// system.
	DocsLink string                                                            `json:"docsLink"`
	JSON     runnerConfigurationConfigurationSchemaGetResponseSchemaScmPatJSON `json:"-"`
}

// runnerConfigurationConfigurationSchemaGetResponseSchemaScmPatJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationConfigurationSchemaGetResponseSchemaScmPat]
type runnerConfigurationConfigurationSchemaGetResponseSchemaScmPatJSON struct {
	Description apijson.Field
	DocsLink    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationConfigurationSchemaGetResponseSchemaScmPat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationConfigurationSchemaGetResponseSchemaScmPatJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationConfigurationSchemaNewParams struct {
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerConfigurationConfigurationSchemaNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationConfigurationSchemaGetParams struct {
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
}

func (r RunnerConfigurationConfigurationSchemaGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
