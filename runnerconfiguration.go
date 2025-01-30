// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
	"github.com/tidwall/gjson"
)

// RunnerConfigurationService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerConfigurationService] method instead.
type RunnerConfigurationService struct {
	Options                  []option.RequestOption
	HostAuthenticationTokens *RunnerConfigurationHostAuthenticationTokenService
	ConfigurationSchema      *RunnerConfigurationConfigurationSchemaService
	ScmIntegration           *RunnerConfigurationScmIntegrationService
	EnvironmentClasses       *RunnerConfigurationEnvironmentClassService
}

// NewRunnerConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRunnerConfigurationService(opts ...option.RequestOption) (r *RunnerConfigurationService) {
	r = &RunnerConfigurationService{}
	r.Options = opts
	r.HostAuthenticationTokens = NewRunnerConfigurationHostAuthenticationTokenService(opts...)
	r.ConfigurationSchema = NewRunnerConfigurationConfigurationSchemaService(opts...)
	r.ScmIntegration = NewRunnerConfigurationScmIntegrationService(opts...)
	r.EnvironmentClasses = NewRunnerConfigurationEnvironmentClassService(opts...)
	return
}

// ValidateRunnerConfiguration validates a runner configuration (e.g. environment
// class, SCM integration)
//
// with the runner.
func (r *RunnerConfigurationService) Validate(ctx context.Context, params RunnerConfigurationValidateParams, opts ...option.RequestOption) (res *RunnerConfigurationValidateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/ValidateRunnerConfiguration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type RunnerConfigurationValidateResponse struct {
	// This field can have the runtime type of
	// [RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClass].
	EnvironmentClass interface{} `json:"environmentClass"`
	// This field can have the runtime type of
	// [RunnerConfigurationValidateResponseScmIntegrationScmIntegration].
	ScmIntegration interface{}                             `json:"scmIntegration"`
	JSON           runnerConfigurationValidateResponseJSON `json:"-"`
	union          RunnerConfigurationValidateResponseUnion
}

// runnerConfigurationValidateResponseJSON contains the JSON metadata for the
// struct [RunnerConfigurationValidateResponse]
type runnerConfigurationValidateResponseJSON struct {
	EnvironmentClass apijson.Field
	ScmIntegration   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r runnerConfigurationValidateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationValidateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationValidateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RunnerConfigurationValidateResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationValidateResponseEnvironmentClass],
// [RunnerConfigurationValidateResponseScmIntegration].
func (r RunnerConfigurationValidateResponse) AsUnion() RunnerConfigurationValidateResponseUnion {
	return r.union
}

// Union satisfied by [RunnerConfigurationValidateResponseEnvironmentClass] or
// [RunnerConfigurationValidateResponseScmIntegration].
type RunnerConfigurationValidateResponseUnion interface {
	implementsRunnerConfigurationValidateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationValidateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationValidateResponseEnvironmentClass{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationValidateResponseScmIntegration{}),
		},
	)
}

type RunnerConfigurationValidateResponseEnvironmentClass struct {
	EnvironmentClass RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClass `json:"environmentClass,required"`
	JSON             runnerConfigurationValidateResponseEnvironmentClassJSON             `json:"-"`
}

// runnerConfigurationValidateResponseEnvironmentClassJSON contains the JSON
// metadata for the struct [RunnerConfigurationValidateResponseEnvironmentClass]
type runnerConfigurationValidateResponseEnvironmentClassJSON struct {
	EnvironmentClass apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponseEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationValidateResponseEnvironmentClass) implementsRunnerConfigurationValidateResponse() {
}

type RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClass struct {
	DescriptionError string                                                                  `json:"descriptionError"`
	DisplayNameError string                                                                  `json:"displayNameError"`
	JSON             runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassJSON `json:"-"`
	union            RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassUnion
}

// runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassJSON contains
// the JSON metadata for the struct
// [RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClass]
type runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassJSON struct {
	DescriptionError apijson.Field
	DisplayNameError apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClass{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDescriptionError],
// [RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDisplayNameError].
func (r RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClass) AsUnion() RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassUnion {
	return r.union
}

// Union satisfied by
// [RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDescriptionError]
// or
// [RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDisplayNameError].
type RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassUnion interface {
	implementsRunnerConfigurationValidateResponseEnvironmentClassEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDescriptionError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDisplayNameError{}),
		},
	)
}

type RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDescriptionError struct {
	DescriptionError string                                                                                  `json:"descriptionError,required"`
	JSON             runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDescriptionErrorJSON `json:"-"`
}

// runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDescriptionErrorJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDescriptionError]
type runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDescriptionErrorJSON struct {
	DescriptionError apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDescriptionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDescriptionErrorJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDescriptionError) implementsRunnerConfigurationValidateResponseEnvironmentClassEnvironmentClass() {
}

type RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDisplayNameError struct {
	DisplayNameError string                                                                                  `json:"displayNameError,required"`
	JSON             runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDisplayNameErrorJSON `json:"-"`
}

// runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDisplayNameErrorJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDisplayNameError]
type runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDisplayNameErrorJSON struct {
	DisplayNameError apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDisplayNameError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDisplayNameErrorJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationValidateResponseEnvironmentClassEnvironmentClassDisplayNameError) implementsRunnerConfigurationValidateResponseEnvironmentClassEnvironmentClass() {
}

type RunnerConfigurationValidateResponseScmIntegration struct {
	ScmIntegration RunnerConfigurationValidateResponseScmIntegrationScmIntegration `json:"scmIntegration,required"`
	JSON           runnerConfigurationValidateResponseScmIntegrationJSON           `json:"-"`
}

// runnerConfigurationValidateResponseScmIntegrationJSON contains the JSON metadata
// for the struct [RunnerConfigurationValidateResponseScmIntegration]
type runnerConfigurationValidateResponseScmIntegrationJSON struct {
	ScmIntegration apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponseScmIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseScmIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationValidateResponseScmIntegration) implementsRunnerConfigurationValidateResponse() {
}

type RunnerConfigurationValidateResponseScmIntegrationScmIntegration struct {
	HostError  string                                                              `json:"hostError"`
	OAuthError string                                                              `json:"oauthError"`
	PatError   string                                                              `json:"patError"`
	ScmIDError string                                                              `json:"scmIdError"`
	JSON       runnerConfigurationValidateResponseScmIntegrationScmIntegrationJSON `json:"-"`
	union      RunnerConfigurationValidateResponseScmIntegrationScmIntegrationUnion
}

// runnerConfigurationValidateResponseScmIntegrationScmIntegrationJSON contains the
// JSON metadata for the struct
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegration]
type runnerConfigurationValidateResponseScmIntegrationScmIntegrationJSON struct {
	HostError   apijson.Field
	OAuthError  apijson.Field
	PatError    apijson.Field
	ScmIDError  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r runnerConfigurationValidateResponseScmIntegrationScmIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r *RunnerConfigurationValidateResponseScmIntegrationScmIntegration) UnmarshalJSON(data []byte) (err error) {
	*r = RunnerConfigurationValidateResponseScmIntegrationScmIntegration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationHostError],
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationOAuthError],
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationPatError],
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationScmIDError].
func (r RunnerConfigurationValidateResponseScmIntegrationScmIntegration) AsUnion() RunnerConfigurationValidateResponseScmIntegrationScmIntegrationUnion {
	return r.union
}

// Union satisfied by
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationHostError],
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationOAuthError],
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationPatError] or
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationScmIDError].
type RunnerConfigurationValidateResponseScmIntegrationScmIntegrationUnion interface {
	implementsRunnerConfigurationValidateResponseScmIntegrationScmIntegration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RunnerConfigurationValidateResponseScmIntegrationScmIntegrationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationValidateResponseScmIntegrationScmIntegrationHostError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationValidateResponseScmIntegrationScmIntegrationOAuthError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationValidateResponseScmIntegrationScmIntegrationPatError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RunnerConfigurationValidateResponseScmIntegrationScmIntegrationScmIDError{}),
		},
	)
}

type RunnerConfigurationValidateResponseScmIntegrationScmIntegrationHostError struct {
	HostError string                                                                       `json:"hostError,required"`
	JSON      runnerConfigurationValidateResponseScmIntegrationScmIntegrationHostErrorJSON `json:"-"`
}

// runnerConfigurationValidateResponseScmIntegrationScmIntegrationHostErrorJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationHostError]
type runnerConfigurationValidateResponseScmIntegrationScmIntegrationHostErrorJSON struct {
	HostError   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponseScmIntegrationScmIntegrationHostError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseScmIntegrationScmIntegrationHostErrorJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationValidateResponseScmIntegrationScmIntegrationHostError) implementsRunnerConfigurationValidateResponseScmIntegrationScmIntegration() {
}

type RunnerConfigurationValidateResponseScmIntegrationScmIntegrationOAuthError struct {
	OAuthError string                                                                        `json:"oauthError,required"`
	JSON       runnerConfigurationValidateResponseScmIntegrationScmIntegrationOAuthErrorJSON `json:"-"`
}

// runnerConfigurationValidateResponseScmIntegrationScmIntegrationOAuthErrorJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationOAuthError]
type runnerConfigurationValidateResponseScmIntegrationScmIntegrationOAuthErrorJSON struct {
	OAuthError  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponseScmIntegrationScmIntegrationOAuthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseScmIntegrationScmIntegrationOAuthErrorJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationValidateResponseScmIntegrationScmIntegrationOAuthError) implementsRunnerConfigurationValidateResponseScmIntegrationScmIntegration() {
}

type RunnerConfigurationValidateResponseScmIntegrationScmIntegrationPatError struct {
	PatError string                                                                      `json:"patError,required"`
	JSON     runnerConfigurationValidateResponseScmIntegrationScmIntegrationPatErrorJSON `json:"-"`
}

// runnerConfigurationValidateResponseScmIntegrationScmIntegrationPatErrorJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationPatError]
type runnerConfigurationValidateResponseScmIntegrationScmIntegrationPatErrorJSON struct {
	PatError    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponseScmIntegrationScmIntegrationPatError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseScmIntegrationScmIntegrationPatErrorJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationValidateResponseScmIntegrationScmIntegrationPatError) implementsRunnerConfigurationValidateResponseScmIntegrationScmIntegration() {
}

type RunnerConfigurationValidateResponseScmIntegrationScmIntegrationScmIDError struct {
	ScmIDError string                                                                        `json:"scmIdError,required"`
	JSON       runnerConfigurationValidateResponseScmIntegrationScmIntegrationScmIDErrorJSON `json:"-"`
}

// runnerConfigurationValidateResponseScmIntegrationScmIntegrationScmIDErrorJSON
// contains the JSON metadata for the struct
// [RunnerConfigurationValidateResponseScmIntegrationScmIntegrationScmIDError]
type runnerConfigurationValidateResponseScmIntegrationScmIntegrationScmIDErrorJSON struct {
	ScmIDError  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationValidateResponseScmIntegrationScmIntegrationScmIDError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationValidateResponseScmIntegrationScmIntegrationScmIDErrorJSON) RawJSON() string {
	return r.raw
}

func (r RunnerConfigurationValidateResponseScmIntegrationScmIntegrationScmIDError) implementsRunnerConfigurationValidateResponseScmIntegrationScmIntegration() {
}

type RunnerConfigurationValidateParams struct {
	Body RunnerConfigurationValidateParamsBodyUnion `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerConfigurationValidateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerConfigurationValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationValidateParamsBody struct {
	EnvironmentClass param.Field[interface{}] `json:"environmentClass"`
	RunnerID         param.Field[string]      `json:"runnerId" format:"uuid"`
	ScmIntegration   param.Field[interface{}] `json:"scmIntegration"`
}

func (r RunnerConfigurationValidateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationValidateParamsBody) implementsRunnerConfigurationValidateParamsBodyUnion() {
}

// Satisfied by [RunnerConfigurationValidateParamsBodyEnvironmentClass],
// [RunnerConfigurationValidateParamsBodyScmIntegration],
// [RunnerConfigurationValidateParamsBody].
type RunnerConfigurationValidateParamsBodyUnion interface {
	implementsRunnerConfigurationValidateParamsBodyUnion()
}

type RunnerConfigurationValidateParamsBodyEnvironmentClass struct {
	EnvironmentClass param.Field[RunnerConfigurationValidateParamsBodyEnvironmentClassEnvironmentClass] `json:"environmentClass,required"`
	RunnerID         param.Field[string]                                                                `json:"runnerId" format:"uuid"`
}

func (r RunnerConfigurationValidateParamsBodyEnvironmentClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationValidateParamsBodyEnvironmentClass) implementsRunnerConfigurationValidateParamsBodyUnion() {
}

type RunnerConfigurationValidateParamsBodyEnvironmentClassEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID param.Field[string] `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration param.Field[[]RunnerConfigurationValidateParamsBodyEnvironmentClassEnvironmentClassConfiguration] `json:"configuration"`
	// description is a human readable description of the environment class
	Description param.Field[string] `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName param.Field[string] `json:"displayName"`
	// enabled indicates whether the environment class can be used to create
	//
	// new environments.
	Enabled param.Field[bool] `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID param.Field[string] `json:"runnerId"`
}

func (r RunnerConfigurationValidateParamsBodyEnvironmentClassEnvironmentClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationValidateParamsBodyEnvironmentClassEnvironmentClassConfiguration struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r RunnerConfigurationValidateParamsBodyEnvironmentClassEnvironmentClassConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationValidateParamsBodyScmIntegration struct {
	ScmIntegration param.Field[RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationUnion] `json:"scmIntegration,required"`
	RunnerID       param.Field[string]                                                                 `json:"runnerId" format:"uuid"`
}

func (r RunnerConfigurationValidateParamsBodyScmIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationValidateParamsBodyScmIntegration) implementsRunnerConfigurationValidateParamsBodyUnion() {
}

type RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegration struct {
	// oauth_client_id is the OAuth app's client ID, if OAuth is configured.
	//
	// If configured, oauth_client_secret must also be set.
	OAuthClientID param.Field[string] `json:"oauthClientId"`
	// oauth_encrypted_client_secret is the OAuth app's client secret encrypted with
	// the runner's public key,
	//
	// if OAuth is configured. This can be used to e.g. validate an already encrypted
	// client secret of an existing SCM integration.
	OAuthEncryptedClientSecret param.Field[string] `json:"oauthEncryptedClientSecret" format:"byte"`
	// oauth_plaintext_client_secret is the OAuth app's client secret in clear text, if
	// OAuth is configured.
	//
	// This can be set to validate any new client secret before it is encrypted and
	// stored. This value will not be stored and get encrypted with the runner's public
	// key before passing it to the runner.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret"`
}

func (r RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegration) implementsRunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationUnion() {
}

// Satisfied by
// [RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthClientID],
// [RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthEncryptedClientSecret],
// [RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthPlaintextClientSecret],
// [RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegration].
type RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationUnion interface {
	implementsRunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationUnion()
}

type RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthClientID struct {
	// oauth_client_id is the OAuth app's client ID, if OAuth is configured.
	//
	// If configured, oauth_client_secret must also be set.
	OAuthClientID param.Field[string] `json:"oauthClientId,required"`
}

func (r RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthClientID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthClientID) implementsRunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationUnion() {
}

type RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthEncryptedClientSecret struct {
	// oauth_encrypted_client_secret is the OAuth app's client secret encrypted with
	// the runner's public key,
	//
	// if OAuth is configured. This can be used to e.g. validate an already encrypted
	// client secret of an existing SCM integration.
	OAuthEncryptedClientSecret param.Field[string] `json:"oauthEncryptedClientSecret,required" format:"byte"`
}

func (r RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthEncryptedClientSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthEncryptedClientSecret) implementsRunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationUnion() {
}

type RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthPlaintextClientSecret struct {
	// oauth_plaintext_client_secret is the OAuth app's client secret in clear text, if
	// OAuth is configured.
	//
	// This can be set to validate any new client secret before it is encrypted and
	// stored. This value will not be stored and get encrypted with the runner's public
	// key before passing it to the runner.
	OAuthPlaintextClientSecret param.Field[string] `json:"oauthPlaintextClientSecret,required"`
}

func (r RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthPlaintextClientSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationOAuthPlaintextClientSecret) implementsRunnerConfigurationValidateParamsBodyScmIntegrationScmIntegrationUnion() {
}

// Define the version of the Connect protocol
type RunnerConfigurationValidateParamsConnectProtocolVersion float64

const (
	RunnerConfigurationValidateParamsConnectProtocolVersion1 RunnerConfigurationValidateParamsConnectProtocolVersion = 1
)

func (r RunnerConfigurationValidateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerConfigurationValidateParamsConnectProtocolVersion1:
		return true
	}
	return false
}
