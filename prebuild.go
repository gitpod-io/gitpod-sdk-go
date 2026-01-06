// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// PrebuildService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrebuildService] method instead.
type PrebuildService struct {
	Options []option.RequestOption
}

// NewPrebuildService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrebuildService(opts ...option.RequestOption) (r *PrebuildService) {
	r = &PrebuildService{}
	r.Options = opts
	return
}

// Creates a prebuild for a project.
//
// Use this method to:
//
// - Create on-demand prebuilds for faster environment startup
// - Trigger prebuilds after repository changes
// - Generate prebuilds for specific environment classes
//
// The prebuild process creates an environment, runs the devcontainer prebuild
// lifecycle, and creates a snapshot for future environment provisioning.
//
// ### Examples
//
// - Create basic prebuild:
//
//	Creates a prebuild for a project using default settings.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	spec:
//	  timeout: "3600s" # 60 minutes default
//	```
//
// - Create prebuild with custom environment class:
//
//	Creates a prebuild with a specific environment class and timeout.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	environmentClassId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	spec:
//	  timeout: "3600s" # 1 hour
//	```
func (r *PrebuildService) New(ctx context.Context, body PrebuildNewParams, opts ...option.RequestOption) (res *PrebuildNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.PrebuildService/CreatePrebuild"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets details about a specific prebuild.
//
// Use this method to:
//
// - Check prebuild status and progress
// - Access prebuild logs for debugging
//
// ### Examples
//
// - Get prebuild details:
//
//	Retrieves comprehensive information about a prebuild.
//
//	```yaml
//	prebuildId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *PrebuildService) Get(ctx context.Context, body PrebuildGetParams, opts ...option.RequestOption) (res *PrebuildGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.PrebuildService/GetPrebuild"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListPrebuilds
func (r *PrebuildService) List(ctx context.Context, params PrebuildListParams, opts ...option.RequestOption) (res *pagination.PrebuildsPage[Prebuild], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.PrebuildService/ListPrebuilds"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// ListPrebuilds
func (r *PrebuildService) ListAutoPaging(ctx context.Context, params PrebuildListParams, opts ...option.RequestOption) *pagination.PrebuildsPageAutoPager[Prebuild] {
	return pagination.NewPrebuildsPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a prebuild.
//
// Prebuilds are automatically deleted after some time. Use this method to manually
// delete a prebuild before automatic cleanup, for example to remove a prebuild
// that should no longer be used.
//
// Deletion is processed asynchronously. The prebuild will be marked for deletion
// and removed from the system in the background.
//
// ### Examples
//
// - Delete prebuild:
//
//	Marks a prebuild for deletion and removes it from the system.
//
//	```yaml
//	prebuildId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *PrebuildService) Delete(ctx context.Context, body PrebuildDeleteParams, opts ...option.RequestOption) (res *PrebuildDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.PrebuildService/DeletePrebuild"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Cancels a running prebuild.
//
// Use this method to:
//
// - Stop prebuilds that are no longer needed
// - Free up resources for other operations
//
// ### Examples
//
// - Cancel prebuild:
//
//	Stops a running prebuild and cleans up resources.
//
//	```yaml
//	prebuildId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *PrebuildService) Cancel(ctx context.Context, body PrebuildCancelParams, opts ...option.RequestOption) (res *PrebuildCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.PrebuildService/CancelPrebuild"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a logs access token for a prebuild.
//
// Use this method to:
//
// - Stream logs from a running prebuild
// - Access archived logs from completed prebuilds
//
// Generated tokens are valid for one hour.
//
// ### Examples
//
// - Create prebuild logs token:
//
//	Generates a token for accessing prebuild logs.
//
//	```yaml
//	prebuildId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *PrebuildService) NewLogsToken(ctx context.Context, body PrebuildNewLogsTokenParams, opts ...option.RequestOption) (res *PrebuildNewLogsTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.PrebuildService/CreatePrebuildLogsToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Prebuild represents a prebuild for a project that creates a snapshot for faster
// environment startup times.
type Prebuild struct {
	// metadata contains organizational and ownership information
	Metadata PrebuildMetadata `json:"metadata,required"`
	// spec contains the configuration used to create this prebuild
	Spec PrebuildSpec `json:"spec,required"`
	// status contains the current status and progress of the prebuild
	Status PrebuildStatus `json:"status,required"`
	// id is the unique identifier for the prebuild
	ID   string       `json:"id" format:"uuid"`
	JSON prebuildJSON `json:"-"`
}

// prebuildJSON contains the JSON metadata for the struct [Prebuild]
type prebuildJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Prebuild) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prebuildJSON) RawJSON() string {
	return r.raw
}

// PrebuildMetadata contains metadata about the prebuild
type PrebuildMetadata struct {
	// created_at is when the prebuild was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// creator is the identity of who created the prebuild. For manual prebuilds, this
	// is the user who triggered it. For scheduled prebuilds, this is the configured
	// executor.
	Creator shared.Subject `json:"creator,required"`
	// updated_at is when the prebuild was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// environment_class_id is the environment class used to create this prebuild.
	// While the prebuild is created with a specific environment class, environments
	// with different classes (e.g., smaller or larger instance sizes) can be created
	// from the same prebuild, as long as they run on the same runner. If not specified
	// in create requests, uses the project's default environment class.
	EnvironmentClassID string `json:"environmentClassId" format:"uuid"`
	// executor is the identity used to run the prebuild. The executor's SCM
	// credentials are used to clone the repository. If not set, the creator's identity
	// is used.
	Executor shared.Subject `json:"executor"`
	// organization_id is the ID of the organization that owns the prebuild
	OrganizationID string `json:"organizationId" format:"uuid"`
	// project_id is the ID of the project this prebuild was created for
	ProjectID string `json:"projectId" format:"uuid"`
	// trigger describes the trigger that created this prebuild.
	TriggeredBy PrebuildTrigger      `json:"triggeredBy"`
	JSON        prebuildMetadataJSON `json:"-"`
}

// prebuildMetadataJSON contains the JSON metadata for the struct
// [PrebuildMetadata]
type prebuildMetadataJSON struct {
	CreatedAt          apijson.Field
	Creator            apijson.Field
	UpdatedAt          apijson.Field
	EnvironmentClassID apijson.Field
	Executor           apijson.Field
	OrganizationID     apijson.Field
	ProjectID          apijson.Field
	TriggeredBy        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PrebuildMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prebuildMetadataJSON) RawJSON() string {
	return r.raw
}

// PrebuildPhase represents the lifecycle phase of a prebuild
type PrebuildPhase string

const (
	PrebuildPhaseUnspecified  PrebuildPhase = "PREBUILD_PHASE_UNSPECIFIED"
	PrebuildPhasePending      PrebuildPhase = "PREBUILD_PHASE_PENDING"
	PrebuildPhaseStarting     PrebuildPhase = "PREBUILD_PHASE_STARTING"
	PrebuildPhaseRunning      PrebuildPhase = "PREBUILD_PHASE_RUNNING"
	PrebuildPhaseStopping     PrebuildPhase = "PREBUILD_PHASE_STOPPING"
	PrebuildPhaseSnapshotting PrebuildPhase = "PREBUILD_PHASE_SNAPSHOTTING"
	PrebuildPhaseCompleted    PrebuildPhase = "PREBUILD_PHASE_COMPLETED"
	PrebuildPhaseFailed       PrebuildPhase = "PREBUILD_PHASE_FAILED"
	PrebuildPhaseCancelling   PrebuildPhase = "PREBUILD_PHASE_CANCELLING"
	PrebuildPhaseCancelled    PrebuildPhase = "PREBUILD_PHASE_CANCELLED"
	PrebuildPhaseDeleting     PrebuildPhase = "PREBUILD_PHASE_DELETING"
	PrebuildPhaseDeleted      PrebuildPhase = "PREBUILD_PHASE_DELETED"
)

func (r PrebuildPhase) IsKnown() bool {
	switch r {
	case PrebuildPhaseUnspecified, PrebuildPhasePending, PrebuildPhaseStarting, PrebuildPhaseRunning, PrebuildPhaseStopping, PrebuildPhaseSnapshotting, PrebuildPhaseCompleted, PrebuildPhaseFailed, PrebuildPhaseCancelling, PrebuildPhaseCancelled, PrebuildPhaseDeleting, PrebuildPhaseDeleted:
		return true
	}
	return false
}

// PrebuildSpec contains the configuration used to create a prebuild
type PrebuildSpec struct {
	// desired_phase is the desired phase of the prebuild. Used to signal cancellation
	// or other state changes. This field is managed by the API and reconciler.
	DesiredPhase PrebuildPhase `json:"desiredPhase"`
	// spec_version is incremented each time the spec is updated. Used for optimistic
	// concurrency control.
	SpecVersion string `json:"specVersion"`
	// timeout is the maximum time allowed for the prebuild to complete. Defaults to 60
	// minutes if not specified. Maximum allowed timeout is 2 hours.
	Timeout string           `json:"timeout" format:"regex"`
	JSON    prebuildSpecJSON `json:"-"`
}

// prebuildSpecJSON contains the JSON metadata for the struct [PrebuildSpec]
type prebuildSpecJSON struct {
	DesiredPhase apijson.Field
	SpecVersion  apijson.Field
	Timeout      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrebuildSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prebuildSpecJSON) RawJSON() string {
	return r.raw
}

// PrebuildSpec contains the configuration used to create a prebuild
type PrebuildSpecParam struct {
	// desired_phase is the desired phase of the prebuild. Used to signal cancellation
	// or other state changes. This field is managed by the API and reconciler.
	DesiredPhase param.Field[PrebuildPhase] `json:"desiredPhase"`
	// spec_version is incremented each time the spec is updated. Used for optimistic
	// concurrency control.
	SpecVersion param.Field[string] `json:"specVersion"`
	// timeout is the maximum time allowed for the prebuild to complete. Defaults to 60
	// minutes if not specified. Maximum allowed timeout is 2 hours.
	Timeout param.Field[string] `json:"timeout" format:"regex"`
}

func (r PrebuildSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PrebuildStatus contains the current status and progress of a prebuild
type PrebuildStatus struct {
	// phase is the current phase of the prebuild lifecycle
	Phase PrebuildPhase `json:"phase,required"`
	// completion_time is when the prebuild completed (successfully or with failure)
	CompletionTime time.Time `json:"completionTime" format:"date-time"`
	// environment_id is the ID of the environment used to create this prebuild. This
	// field is set when the prebuild environment is created.
	EnvironmentID string `json:"environmentId" format:"uuid"`
	// failure_message contains details about why the prebuild failed
	FailureMessage string `json:"failureMessage"`
	// log_url provides access to prebuild logs. During prebuild execution, this
	// references the environment logs. After completion, this may reference archived
	// logs.
	LogURL string `json:"logUrl" format:"uri"`
	// snapshot_completion_percentage is the progress of snapshot creation (0-100).
	// Only populated when phase is SNAPSHOTTING and progress is available from the
	// cloud provider. This value may update infrequently or remain at 0 depending on
	// the provider.
	SnapshotCompletionPercentage int64 `json:"snapshotCompletionPercentage"`
	// status_version is incremented each time the status is updated. Used for
	// optimistic concurrency control.
	StatusVersion string `json:"statusVersion"`
	// warning_message contains warnings from the prebuild environment that indicate
	// something went wrong but the prebuild could still complete. For example, the
	// devcontainer failed to build but the environment is still usable. These warnings
	// will likely affect any environment started from this prebuild.
	WarningMessage string             `json:"warningMessage"`
	JSON           prebuildStatusJSON `json:"-"`
}

// prebuildStatusJSON contains the JSON metadata for the struct [PrebuildStatus]
type prebuildStatusJSON struct {
	Phase                        apijson.Field
	CompletionTime               apijson.Field
	EnvironmentID                apijson.Field
	FailureMessage               apijson.Field
	LogURL                       apijson.Field
	SnapshotCompletionPercentage apijson.Field
	StatusVersion                apijson.Field
	WarningMessage               apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *PrebuildStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prebuildStatusJSON) RawJSON() string {
	return r.raw
}

// PrebuildTrigger indicates how the prebuild was triggered
type PrebuildTrigger string

const (
	PrebuildTriggerUnspecified PrebuildTrigger = "PREBUILD_TRIGGER_UNSPECIFIED"
	PrebuildTriggerManual      PrebuildTrigger = "PREBUILD_TRIGGER_MANUAL"
	PrebuildTriggerScheduled   PrebuildTrigger = "PREBUILD_TRIGGER_SCHEDULED"
)

func (r PrebuildTrigger) IsKnown() bool {
	switch r {
	case PrebuildTriggerUnspecified, PrebuildTriggerManual, PrebuildTriggerScheduled:
		return true
	}
	return false
}

type PrebuildNewResponse struct {
	// Prebuild represents a prebuild for a project that creates a snapshot for faster
	// environment startup times.
	Prebuild Prebuild                `json:"prebuild,required"`
	JSON     prebuildNewResponseJSON `json:"-"`
}

// prebuildNewResponseJSON contains the JSON metadata for the struct
// [PrebuildNewResponse]
type prebuildNewResponseJSON struct {
	Prebuild    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrebuildNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prebuildNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrebuildGetResponse struct {
	// Prebuild represents a prebuild for a project that creates a snapshot for faster
	// environment startup times.
	Prebuild Prebuild                `json:"prebuild,required"`
	JSON     prebuildGetResponseJSON `json:"-"`
}

// prebuildGetResponseJSON contains the JSON metadata for the struct
// [PrebuildGetResponse]
type prebuildGetResponseJSON struct {
	Prebuild    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrebuildGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prebuildGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrebuildDeleteResponse = interface{}

type PrebuildCancelResponse struct {
	// Prebuild represents a prebuild for a project that creates a snapshot for faster
	// environment startup times.
	Prebuild Prebuild                   `json:"prebuild,required"`
	JSON     prebuildCancelResponseJSON `json:"-"`
}

// prebuildCancelResponseJSON contains the JSON metadata for the struct
// [PrebuildCancelResponse]
type prebuildCancelResponseJSON struct {
	Prebuild    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrebuildCancelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prebuildCancelResponseJSON) RawJSON() string {
	return r.raw
}

type PrebuildNewLogsTokenResponse struct {
	// access_token is the token that can be used to access the logs of the prebuild
	AccessToken string                           `json:"accessToken,required"`
	JSON        prebuildNewLogsTokenResponseJSON `json:"-"`
}

// prebuildNewLogsTokenResponseJSON contains the JSON metadata for the struct
// [PrebuildNewLogsTokenResponse]
type prebuildNewLogsTokenResponseJSON struct {
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrebuildNewLogsTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prebuildNewLogsTokenResponseJSON) RawJSON() string {
	return r.raw
}

type PrebuildNewParams struct {
	// project_id specifies the project to create a prebuild for
	ProjectID param.Field[string] `json:"projectId,required" format:"uuid"`
	// spec contains the configuration for creating the prebuild
	Spec param.Field[PrebuildSpecParam] `json:"spec,required"`
	// environment_class_id specifies which environment class to use for the prebuild.
	// If not specified, uses the project's default environment class.
	EnvironmentClassID param.Field[string] `json:"environmentClassId" format:"uuid"`
}

func (r PrebuildNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrebuildGetParams struct {
	// prebuild_id specifies the prebuild to retrieve
	PrebuildID param.Field[string] `json:"prebuildId,required" format:"uuid"`
}

func (r PrebuildGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrebuildListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// filter contains the filter options for listing prebuilds
	Filter param.Field[PrebuildListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing prebuilds
	Pagination param.Field[PrebuildListParamsPagination] `json:"pagination"`
}

func (r PrebuildListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [PrebuildListParams]'s query parameters as `url.Values`.
func (r PrebuildListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// filter contains the filter options for listing prebuilds
type PrebuildListParamsFilter struct {
	// phases filters prebuilds by their current phase
	Phases param.Field[[]PrebuildPhase] `json:"phases"`
	// project_ids filters prebuilds to specific projects
	ProjectIDs param.Field[[]string] `json:"projectIds" format:"uuid"`
}

func (r PrebuildListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing prebuilds
type PrebuildListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r PrebuildListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrebuildDeleteParams struct {
	// prebuild_id specifies the prebuild to delete
	PrebuildID param.Field[string] `json:"prebuildId,required" format:"uuid"`
}

func (r PrebuildDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrebuildCancelParams struct {
	// prebuild_id specifies the prebuild to cancel
	PrebuildID param.Field[string] `json:"prebuildId,required" format:"uuid"`
}

func (r PrebuildCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrebuildNewLogsTokenParams struct {
	// prebuild_id specifies the prebuild for which the logs token should be created.
	//
	// +required
	PrebuildID param.Field[string] `json:"prebuildId,required" format:"uuid"`
}

func (r PrebuildNewLogsTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
