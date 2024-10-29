// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/gitpod-go/internal/apiform"
	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// RunnerInteractionService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerInteractionService] method instead.
type RunnerInteractionService struct {
	Options      []option.RequestOption
	Environments *RunnerInteractionEnvironmentService
}

// NewRunnerInteractionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRunnerInteractionService(opts ...option.RequestOption) (r *RunnerInteractionService) {
	r = &RunnerInteractionService{}
	r.Options = opts
	r.Environments = NewRunnerInteractionEnvironmentService(opts...)
	return
}

// GetRunnerHostAuthenticationToken returns an authentication token for the given
// host.
func (r *RunnerInteractionService) GetHostAuthenticationTokenValue(ctx context.Context, params RunnerInteractionGetHostAuthenticationTokenValueParams, opts ...option.RequestOption) (res *RunnerInteractionGetHostAuthenticationTokenValueResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/GetHostAuthenticationTokenValue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetLatestVersion returns the latest version of the runner.
func (r *RunnerInteractionService) GetLatestVersion(ctx context.Context, params RunnerInteractionGetLatestVersionParams, opts ...option.RequestOption) (res *RunnerInteractionGetLatestVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/GetLatestVersion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListRunnerEnvironmentClasses returns the environment classes configured for the
// runner.
func (r *RunnerInteractionService) ListRunnerEnvironmentClasses(ctx context.Context, params RunnerInteractionListRunnerEnvironmentClassesParams, opts ...option.RequestOption) (res *RunnerInteractionListRunnerEnvironmentClassesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/ListRunnerEnvironmentClasses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListRunnerSCMIntegrations
func (r *RunnerInteractionService) ListRunnerScmIntegrations(ctx context.Context, params RunnerInteractionListRunnerScmIntegrationsParams, opts ...option.RequestOption) (res *RunnerInteractionListRunnerScmIntegrationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/ListRunnerSCMIntegrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// MarkRunnerActive marks a runner as available. This must be called every 30
// seconds to keep the runner active.
func (r *RunnerInteractionService) MarkActive(ctx context.Context, params RunnerInteractionMarkActiveParams, opts ...option.RequestOption) (res *RunnerInteractionMarkActiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/MarkRunnerActive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// SendResponse sends a response to a request.
func (r *RunnerInteractionService) SendResponse(ctx context.Context, params RunnerInteractionSendResponseParams, opts ...option.RequestOption) (res *RunnerInteractionSendResponseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/SendResponse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Signup is called by a runner to sign up with the backend. This is the first call
// a runner makes.
func (r *RunnerInteractionService) Signup(ctx context.Context, params RunnerInteractionSignupParams, opts ...option.RequestOption) (res *RunnerInteractionSignupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/Signup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateRunnerConfigurationSchema updates the runner's configuration schema.
func (r *RunnerInteractionService) UpdateRunnerConfigurationSchema(ctx context.Context, params RunnerInteractionUpdateRunnerConfigurationSchemaParams, opts ...option.RequestOption) (res *RunnerInteractionUpdateRunnerConfigurationSchemaResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/UpdateRunnerConfigurationSchema"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateRunnerStatus updates the status of the runner.
func (r *RunnerInteractionService) UpdateStatus(ctx context.Context, params RunnerInteractionUpdateStatusParams, opts ...option.RequestOption) (res *RunnerInteractionUpdateStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerInteractionService/UpdateRunnerStatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type RunnerInteractionGetHostAuthenticationTokenValueResponse struct {
	// The host authentication token's ID
	TokenID string `json:"tokenId"`
	// The authentication token encrypted as NaCL anonymous sealed box using the
	// runner's public key.
	Value string                                                       `json:"value" format:"byte"`
	JSON  runnerInteractionGetHostAuthenticationTokenValueResponseJSON `json:"-"`
}

// runnerInteractionGetHostAuthenticationTokenValueResponseJSON contains the JSON
// metadata for the struct
// [RunnerInteractionGetHostAuthenticationTokenValueResponse]
type runnerInteractionGetHostAuthenticationTokenValueResponseJSON struct {
	TokenID     apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionGetHostAuthenticationTokenValueResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionGetHostAuthenticationTokenValueResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionGetLatestVersionResponse struct {
	// auto-update indicates if the runner should be updated automatically
	AutoUpdate bool `json:"autoUpdate"`
	// gitpod_cli_download_url is the URL to download the gitpod CLI
	GitpodCliDownloadURL string `json:"gitpodCliDownloadUrl"`
	// The container image of the runner
	RunnerImage string `json:"runnerImage"`
	// supervisor_download_url is the URL to download the supervisor
	SupervisorDownloadURL string `json:"supervisorDownloadUrl"`
	// The latest version of the runner
	Version string                                        `json:"version"`
	JSON    runnerInteractionGetLatestVersionResponseJSON `json:"-"`
}

// runnerInteractionGetLatestVersionResponseJSON contains the JSON metadata for the
// struct [RunnerInteractionGetLatestVersionResponse]
type runnerInteractionGetLatestVersionResponseJSON struct {
	AutoUpdate            apijson.Field
	GitpodCliDownloadURL  apijson.Field
	RunnerImage           apijson.Field
	SupervisorDownloadURL apijson.Field
	Version               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RunnerInteractionGetLatestVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionGetLatestVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionListRunnerEnvironmentClassesResponse struct {
	// The environment classes configured for the runner
	EnvironmentClasses []RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClass `json:"environmentClasses"`
	// pagination contains the pagination options for listing environment classes
	Pagination RunnerInteractionListRunnerEnvironmentClassesResponsePagination `json:"pagination"`
	JSON       runnerInteractionListRunnerEnvironmentClassesResponseJSON       `json:"-"`
}

// runnerInteractionListRunnerEnvironmentClassesResponseJSON contains the JSON
// metadata for the struct [RunnerInteractionListRunnerEnvironmentClassesResponse]
type runnerInteractionListRunnerEnvironmentClassesResponseJSON struct {
	EnvironmentClasses apijson.Field
	Pagination         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerEnvironmentClassesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerEnvironmentClassesResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID string `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration []RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfiguration `json:"configuration"`
	// description is a human readable description of the environment class
	Description string `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName string `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled bool `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID string                                                                    `json:"runnerId"`
	JSON     runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassJSON `json:"-"`
}

// runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassJSON
// contains the JSON metadata for the struct
// [RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClass]
type runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	RunnerID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfiguration struct {
	Key   string                                                                                   `json:"key"`
	Value string                                                                                   `json:"value"`
	JSON  runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfigurationJSON `json:"-"`
}

// runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfigurationJSON
// contains the JSON metadata for the struct
// [RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfiguration]
type runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfigurationJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerEnvironmentClassesResponseEnvironmentClassesConfigurationJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing environment classes
type RunnerInteractionListRunnerEnvironmentClassesResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                                              `json:"nextToken"`
	JSON      runnerInteractionListRunnerEnvironmentClassesResponsePaginationJSON `json:"-"`
}

// runnerInteractionListRunnerEnvironmentClassesResponsePaginationJSON contains the
// JSON metadata for the struct
// [RunnerInteractionListRunnerEnvironmentClassesResponsePagination]
type runnerInteractionListRunnerEnvironmentClassesResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerEnvironmentClassesResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerEnvironmentClassesResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionListRunnerScmIntegrationsResponse struct {
	// pagination contains the pagination options for listing SCM integrations
	Pagination RunnerInteractionListRunnerScmIntegrationsResponsePagination `json:"pagination"`
	// The SCM integrations configured for the runner
	ScmIntegrations []RunnerInteractionListRunnerScmIntegrationsResponseScmIntegrationsUnion `json:"scmIntegrations"`
	JSON            runnerInteractionListRunnerScmIntegrationsResponseJSON                   `json:"-"`
}

// runnerInteractionListRunnerScmIntegrationsResponseJSON contains the JSON
// metadata for the struct [RunnerInteractionListRunnerScmIntegrationsResponse]
type runnerInteractionListRunnerScmIntegrationsResponseJSON struct {
	Pagination      apijson.Field
	ScmIntegrations apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerScmIntegrationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerScmIntegrationsResponseJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing SCM integrations
type RunnerInteractionListRunnerScmIntegrationsResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                                           `json:"nextToken"`
	JSON      runnerInteractionListRunnerScmIntegrationsResponsePaginationJSON `json:"-"`
}

// runnerInteractionListRunnerScmIntegrationsResponsePaginationJSON contains the
// JSON metadata for the struct
// [RunnerInteractionListRunnerScmIntegrationsResponsePagination]
type runnerInteractionListRunnerScmIntegrationsResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionListRunnerScmIntegrationsResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionListRunnerScmIntegrationsResponsePaginationJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [RunnerInteractionListRunnerScmIntegrationsResponseScmIntegrationsUnknown] or
// [RunnerInteractionListRunnerScmIntegrationsResponseScmIntegrationsUnknown].
type RunnerInteractionListRunnerScmIntegrationsResponseScmIntegrationsUnion interface {
	implementsRunnerInteractionListRunnerScmIntegrationsResponseScmIntegrationsUnion()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*RunnerInteractionListRunnerScmIntegrationsResponseScmIntegrationsUnion)(nil)).Elem(), "")
}

type RunnerInteractionMarkActiveResponse = interface{}

type RunnerInteractionSendResponseResponse = interface{}

type RunnerInteractionSignupResponse struct {
	// The runner's identity
	RunnerID string                              `json:"runnerId" format:"uuid"`
	JSON     runnerInteractionSignupResponseJSON `json:"-"`
}

// runnerInteractionSignupResponseJSON contains the JSON metadata for the struct
// [RunnerInteractionSignupResponse]
type runnerInteractionSignupResponseJSON struct {
	RunnerID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerInteractionSignupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerInteractionSignupResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerInteractionUpdateRunnerConfigurationSchemaResponse = interface{}

type RunnerInteractionUpdateStatusResponse = interface{}

type RunnerInteractionGetHostAuthenticationTokenValueParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// The host to get the authentication token for
	Host param.Field[string] `json:"host"`
	// The principal's ID to get the authentication token for
	PrincipalID param.Field[string] `json:"principalId" format:"uuid"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionGetHostAuthenticationTokenValueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion float64

const (
	RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion1 RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionGetHostAuthenticationTokenValueParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionGetLatestVersionParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionGetLatestVersionParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// The current version of the runner
	CurrentVersion param.Field[string] `json:"currentVersion"`
	// The version of the infrastructure
	InfrastructureVersion param.Field[string] `json:"infrastructureVersion"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionGetLatestVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionGetLatestVersionParamsConnectProtocolVersion float64

const (
	RunnerInteractionGetLatestVersionParamsConnectProtocolVersion1 RunnerInteractionGetLatestVersionParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionGetLatestVersionParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionGetLatestVersionParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionListRunnerEnvironmentClassesParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Filter                 param.Field[RunnerInteractionListRunnerEnvironmentClassesParamsFilter]                 `json:"filter"`
	// pagination contains the pagination options for listing environment classes
	Pagination param.Field[RunnerInteractionListRunnerEnvironmentClassesParamsPagination] `json:"pagination"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionListRunnerEnvironmentClassesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion float64

const (
	RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion1 RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionListRunnerEnvironmentClassesParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionListRunnerEnvironmentClassesParamsFilter struct {
	// environment_class_ids filters the response to only environment classes with
	// these IDs
	EnvironmentClassIDs param.Field[[]string] `json:"environmentClassIds" format:"uuid"`
}

func (r RunnerInteractionListRunnerEnvironmentClassesParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing environment classes
type RunnerInteractionListRunnerEnvironmentClassesParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerInteractionListRunnerEnvironmentClassesParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionListRunnerScmIntegrationsParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Filter                 param.Field[RunnerInteractionListRunnerScmIntegrationsParamsFilter]                 `json:"filter"`
	// pagination contains the pagination options for listing SCM integrations
	Pagination param.Field[RunnerInteractionListRunnerScmIntegrationsParamsPagination] `json:"pagination"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionListRunnerScmIntegrationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion float64

const (
	RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion1 RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionListRunnerScmIntegrationsParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionListRunnerScmIntegrationsParamsFilter struct {
	// environment_class_ids filters the response to only SCM integrations with these
	// IDs
	ScmIntegrationIDs param.Field[[]string] `json:"scmIntegrationIds" format:"uuid"`
}

func (r RunnerInteractionListRunnerScmIntegrationsParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing SCM integrations
type RunnerInteractionListRunnerScmIntegrationsParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerInteractionListRunnerScmIntegrationsParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionMarkActiveParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionMarkActiveParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionMarkActiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionMarkActiveParamsConnectProtocolVersion float64

const (
	RunnerInteractionMarkActiveParamsConnectProtocolVersion1 RunnerInteractionMarkActiveParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionMarkActiveParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionMarkActiveParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionSendResponseParams struct {
	Body RunnerInteractionSendResponseParamsBody `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionSendResponseParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionSendResponseParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type RunnerInteractionSendResponseParamsBody struct {
}

func (r RunnerInteractionSendResponseParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionSendResponseParamsConnectProtocolVersion float64

const (
	RunnerInteractionSendResponseParamsConnectProtocolVersion1 RunnerInteractionSendResponseParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionSendResponseParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionSendResponseParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionSignupParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionSignupParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// The environment classes this runner has to offer
	EnvironmentClasses param.Field[[]RunnerInteractionSignupParamsEnvironmentClass] `json:"environmentClasses"`
	// The runner's public key. Must be an ECDH public key encoded in PKIX, ASN.1 DER
	// format.
	PublicKey param.Field[string] `json:"publicKey" format:"byte"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionSignupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionSignupParamsConnectProtocolVersion float64

const (
	RunnerInteractionSignupParamsConnectProtocolVersion1 RunnerInteractionSignupParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionSignupParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionSignupParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type RunnerInteractionSignupParamsEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID param.Field[string] `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration param.Field[[]RunnerInteractionSignupParamsEnvironmentClassesConfiguration] `json:"configuration"`
	// description is a human readable description of the environment class
	Description param.Field[string] `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName param.Field[string] `json:"displayName"`
	// enabled indicates whether the environment class can be used to create new
	// environments.
	Enabled param.Field[bool] `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID param.Field[string] `json:"runnerId"`
}

func (r RunnerInteractionSignupParamsEnvironmentClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionSignupParamsEnvironmentClassesConfiguration struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r RunnerInteractionSignupParamsEnvironmentClassesConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionUpdateRunnerConfigurationSchemaParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// config_schema is the schema for the runner's configuration
	ConfigSchema param.Field[RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchema] `json:"configSchema"`
	// The runner's identity
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion float64

const (
	RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion1 RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionUpdateRunnerConfigurationSchemaParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// config_schema is the schema for the runner's configuration
type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchema struct {
	EnvironmentClasses param.Field[[]RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClassUnion] `json:"environmentClasses"`
	RunnerConfig       param.Field[[]RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfigUnion]     `json:"runnerConfig"`
	Scm                param.Field[[]RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScm]                   `json:"scm"`
	// The schema version
	Version param.Field[string] `json:"version"`
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClassesUnknown],
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClassesUnknown],
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClassesUnknown],
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClassesUnknown],
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClassesUnknown],
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClassesUnknown].
type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClassUnion interface {
	implementsRunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaEnvironmentClassUnion()
}

// Satisfied by
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfigUnknown],
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfigUnknown],
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfigUnknown],
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfigUnknown],
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfigUnknown],
// [RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfigUnknown].
type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfigUnion interface {
	implementsRunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaRunnerConfigUnion()
}

type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScm struct {
	DefaultHosts param.Field[[]string]                                                                   `json:"defaultHosts"`
	Name         param.Field[string]                                                                     `json:"name"`
	OAuth        param.Field[RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmOAuth] `json:"oauth"`
	Pat          param.Field[RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmPat]   `json:"pat"`
	ScmID        param.Field[string]                                                                     `json:"scmId"`
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScm) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmOAuth struct {
	// callback_url is the URL the OAuth app will redirect to after the user has
	// authenticated.
	CallbackURL param.Field[string] `json:"callbackUrl"`
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmOAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmPat struct {
	// description is a human-readable description of the PAT.
	Description param.Field[string] `json:"description"`
	// docs_link is a link to the documentation on how to create a PAT for this SCM
	// system.
	DocsLink param.Field[string] `json:"docsLink"`
}

func (r RunnerInteractionUpdateRunnerConfigurationSchemaParamsConfigSchemaScmPat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerInteractionUpdateStatusParams struct {
	Body RunnerInteractionUpdateStatusParamsBody `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[RunnerInteractionUpdateStatusParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r RunnerInteractionUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerInteractionUpdateStatusParamsBody struct {
}

func (r RunnerInteractionUpdateStatusParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type RunnerInteractionUpdateStatusParamsConnectProtocolVersion float64

const (
	RunnerInteractionUpdateStatusParamsConnectProtocolVersion1 RunnerInteractionUpdateStatusParamsConnectProtocolVersion = 1
)

func (r RunnerInteractionUpdateStatusParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case RunnerInteractionUpdateStatusParamsConnectProtocolVersion1:
		return true
	}
	return false
}
