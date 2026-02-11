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

// ProjectService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options           []option.RequestOption
	EnvironmentClases *ProjectEnvironmentClaseService
	Policies          *ProjectPolicyService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r *ProjectService) {
	r = &ProjectService{}
	r.Options = opts
	r.EnvironmentClases = NewProjectEnvironmentClaseService(opts...)
	r.Policies = NewProjectPolicyService(opts...)
	return
}

// Creates a new project with specified configuration.
//
// Use this method to:
//
// - Set up development projects
// - Configure project environments
// - Define project settings
// - Initialize project content
//
// ### Examples
//
// - Create basic project:
//
//	Creates a project with minimal configuration.
//
//	```yaml
//	name: "Web Application"
//	initializer:
//	  specs:
//	    - git:
//	        remoteUri: "https://github.com/org/repo"
//	```
//
// - Create project with devcontainer:
//
//	Creates a project with custom development container.
//
//	```yaml
//	name: "Backend Service"
//	initializer:
//	  specs:
//	    - git:
//	        remoteUri: "https://github.com/org/backend"
//	devcontainerFilePath: ".devcontainer/devcontainer.json"
//	automationsFilePath: ".gitpod/automations.yaml"
//	```
func (r *ProjectService) New(ctx context.Context, body ProjectNewParams, opts ...option.RequestOption) (res *ProjectNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.ProjectService/CreateProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets details about a specific project.
//
// Use this method to:
//
// - View project configuration
// - Check project status
// - Get project metadata
//
// ### Examples
//
// - Get project details:
//
//	Retrieves information about a specific project.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *ProjectService) Get(ctx context.Context, body ProjectGetParams, opts ...option.RequestOption) (res *ProjectGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.ProjectService/GetProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates a project's configuration.
//
// Use this method to:
//
// - Modify project settings
// - Update environment class
// - Change project name
// - Configure initializers
// - Configure prebuild settings
//
// ### Examples
//
// - Update project name:
//
//	Changes the project's display name.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	name: "New Project Name"
//	```
//
// - Enable prebuilds with daily schedule:
//
//	Configures prebuilds to run daily at 2 AM UTC.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	prebuildConfiguration:
//	  enabled: true
//	  environmentClassIds:
//	    - "b0e12f6c-4c67-429d-a4a6-d9838b5da041"
//	  timeout: "3600s"
//	  trigger:
//	    dailySchedule:
//	      hourUtc: 2
//	```
func (r *ProjectService) Update(ctx context.Context, body ProjectUpdateParams, opts ...option.RequestOption) (res *ProjectUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.ProjectService/UpdateProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists projects with optional filtering.
//
// Use this method to:
//
// - View all accessible projects
// - Browse project configurations
// - Monitor project status
//
// ### Examples
//
// - List projects:
//
//	Shows all projects with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
func (r *ProjectService) List(ctx context.Context, params ProjectListParams, opts ...option.RequestOption) (res *pagination.ProjectsPage[Project], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.ProjectService/ListProjects"
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

// Lists projects with optional filtering.
//
// Use this method to:
//
// - View all accessible projects
// - Browse project configurations
// - Monitor project status
//
// ### Examples
//
// - List projects:
//
//	Shows all projects with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
func (r *ProjectService) ListAutoPaging(ctx context.Context, params ProjectListParams, opts ...option.RequestOption) *pagination.ProjectsPageAutoPager[Project] {
	return pagination.NewProjectsPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a project permanently.
//
// Use this method to:
//
// - Remove unused projects
// - Clean up test projects
// - Delete obsolete configurations
//
// ### Examples
//
// - Delete project:
//
//	Permanently removes a project.
//
//	```yaml
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *ProjectService) Delete(ctx context.Context, body ProjectDeleteParams, opts ...option.RequestOption) (res *ProjectDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.ProjectService/DeleteProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates multiple projects in a single request.
//
// Use this method to:
//
// - Onboard multiple repositories at once
// - Import a batch of projects during initial setup
//
// Returns successfully created projects and details about any failures. Each
// project in the request is processed independently — partial success is possible.
//
// ### Examples
//
// - Create multiple projects:
//
//	Creates several projects in one request.
//
//	```yaml
//	projects:
//	  - name: "Frontend"
//	    initializer:
//	      specs:
//	        - git:
//	            remoteUri: "https://github.com/org/frontend"
//	  - name: "Backend"
//	    initializer:
//	      specs:
//	        - git:
//	            remoteUri: "https://github.com/org/backend"
//	```
func (r *ProjectService) BulkNew(ctx context.Context, body ProjectBulkNewParams, opts ...option.RequestOption) (res *ProjectBulkNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.ProjectService/CreateProjects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes multiple projects in a single request.
//
// Use this method to:
//
// - Remove multiple unused projects at once
// - Clean up projects in batch
//
// Returns successfully deleted project IDs and details about any failures. Each
// project in the request is processed independently — partial success is possible.
//
// ### Examples
//
// - Delete multiple projects:
//
//	Permanently removes several projects in one request.
//
//	```yaml
//	projectIds:
//	  - "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	  - "c1f23g7d-5d78-430e-b5b7-e0949c6eb158"
//	```
func (r *ProjectService) BulkDelete(ctx context.Context, body ProjectBulkDeleteParams, opts ...option.RequestOption) (res *ProjectBulkDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.ProjectService/DeleteProjects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates multiple projects in a single request.
//
// Use this method to:
//
// - Modify settings across multiple projects at once
// - Apply configuration changes in batch
//
// Returns successfully updated projects and details about any failures. Each
// project in the request is processed independently — partial success is possible.
//
// ### Examples
//
// - Update multiple projects:
//
//	Updates several projects in one request.
//
//	```yaml
//	projects:
//	  - projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	    name: "Updated Frontend"
//	  - projectId: "c1f23g7d-5d78-430e-b5b7-e0949c6eb158"
//	    name: "Updated Backend"
//	```
func (r *ProjectService) BulkUpdate(ctx context.Context, body ProjectBulkUpdateParams, opts ...option.RequestOption) (res *ProjectBulkUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.ProjectService/UpdateProjects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a new project using an existing environment as a template.
//
// Use this method to:
//
// - Clone environment configurations
// - Create projects from templates
// - Share environment setups
//
// ### Examples
//
// - Create from environment:
//
//	Creates a project based on existing environment.
//
//	```yaml
//	name: "Frontend Project"
//	environmentId: "07e03a28-65a5-4d98-b532-8ea67b188048"
//	```
func (r *ProjectService) NewFromEnvironment(ctx context.Context, body ProjectNewFromEnvironmentParams, opts ...option.RequestOption) (res *ProjectNewFromEnvironmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.ProjectService/CreateProjectFromEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// EnvironmentInitializer specifies how an environment is to be initialized
type EnvironmentInitializer struct {
	Specs []EnvironmentInitializerSpec `json:"specs"`
	JSON  environmentInitializerJSON   `json:"-"`
}

// environmentInitializerJSON contains the JSON metadata for the struct
// [EnvironmentInitializer]
type environmentInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentInitializerJSON) RawJSON() string {
	return r.raw
}

type EnvironmentInitializerSpec struct {
	ContextURL EnvironmentInitializerSpecsContextURL `json:"contextUrl"`
	Git        EnvironmentInitializerSpecsGit        `json:"git"`
	JSON       environmentInitializerSpecJSON        `json:"-"`
}

// environmentInitializerSpecJSON contains the JSON metadata for the struct
// [EnvironmentInitializerSpec]
type environmentInitializerSpecJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentInitializerSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentInitializerSpecJSON) RawJSON() string {
	return r.raw
}

type EnvironmentInitializerSpecsContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                    `json:"url" format:"uri"`
	JSON environmentInitializerSpecsContextURLJSON `json:"-"`
}

// environmentInitializerSpecsContextURLJSON contains the JSON metadata for the
// struct [EnvironmentInitializerSpecsContextURL]
type environmentInitializerSpecsContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentInitializerSpecsContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentInitializerSpecsContextURLJSON) RawJSON() string {
	return r.raw
}

type EnvironmentInitializerSpecsGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// the target mode determines what gets checked out
	TargetMode EnvironmentInitializerSpecsGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                             `json:"upstreamRemoteUri"`
	JSON              environmentInitializerSpecsGitJSON `json:"-"`
}

// environmentInitializerSpecsGitJSON contains the JSON metadata for the struct
// [EnvironmentInitializerSpecsGit]
type environmentInitializerSpecsGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentInitializerSpecsGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentInitializerSpecsGitJSON) RawJSON() string {
	return r.raw
}

// the target mode determines what gets checked out
type EnvironmentInitializerSpecsGitTargetMode string

const (
	EnvironmentInitializerSpecsGitTargetModeCloneTargetModeUnspecified  EnvironmentInitializerSpecsGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteHead   EnvironmentInitializerSpecsGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteCommit EnvironmentInitializerSpecsGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteBranch EnvironmentInitializerSpecsGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	EnvironmentInitializerSpecsGitTargetModeCloneTargetModeLocalBranch  EnvironmentInitializerSpecsGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
	EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteTag    EnvironmentInitializerSpecsGitTargetMode = "CLONE_TARGET_MODE_REMOTE_TAG"
)

func (r EnvironmentInitializerSpecsGitTargetMode) IsKnown() bool {
	switch r {
	case EnvironmentInitializerSpecsGitTargetModeCloneTargetModeUnspecified, EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteHead, EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteCommit, EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteBranch, EnvironmentInitializerSpecsGitTargetModeCloneTargetModeLocalBranch, EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteTag:
		return true
	}
	return false
}

// EnvironmentInitializer specifies how an environment is to be initialized
type EnvironmentInitializerParam struct {
	Specs param.Field[[]EnvironmentInitializerSpecParam] `json:"specs"`
}

func (r EnvironmentInitializerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentInitializerSpecParam struct {
	ContextURL param.Field[EnvironmentInitializerSpecsContextURLParam] `json:"contextUrl"`
	Git        param.Field[EnvironmentInitializerSpecsGitParam]        `json:"git"`
}

func (r EnvironmentInitializerSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentInitializerSpecsContextURLParam struct {
	// url is the URL from which the environment is created
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r EnvironmentInitializerSpecsContextURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentInitializerSpecsGitParam struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation param.Field[string] `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget param.Field[string] `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri param.Field[string] `json:"remoteUri"`
	// the target mode determines what gets checked out
	TargetMode param.Field[EnvironmentInitializerSpecsGitTargetMode] `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri param.Field[string] `json:"upstreamRemoteUri"`
}

func (r EnvironmentInitializerSpecsGitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Project struct {
	// Use `environment_classes` instead.
	//
	// Deprecated: deprecated
	EnvironmentClass shared.ProjectEnvironmentClass `json:"environmentClass,required"`
	// id is the unique identifier for the project
	ID string `json:"id" format:"uuid"`
	// automations_file_path is the path to the automations file relative to the repo
	// root
	AutomationsFilePath string `json:"automationsFilePath"`
	// desired_phase is the desired phase of the project When set to DELETED, the
	// project is pending deletion
	DesiredPhase ProjectPhase `json:"desiredPhase"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// environment_classes is the list of environment classes for the project
	EnvironmentClasses []shared.ProjectEnvironmentClass `json:"environmentClasses"`
	// initializer is the content initializer
	Initializer EnvironmentInitializer `json:"initializer"`
	Metadata    ProjectMetadata        `json:"metadata"`
	// prebuild_configuration defines how prebuilds are created for this project.
	PrebuildConfiguration ProjectPrebuildConfiguration `json:"prebuildConfiguration"`
	// recommended_editors specifies the editors recommended for this project.
	RecommendedEditors RecommendedEditors `json:"recommendedEditors"`
	// technical_description is a detailed technical description of the project This
	// field is not returned by default in GetProject or ListProjects responses
	TechnicalDescription string        `json:"technicalDescription"`
	UsedBy               ProjectUsedBy `json:"usedBy"`
	JSON                 projectJSON   `json:"-"`
}

// projectJSON contains the JSON metadata for the struct [Project]
type projectJSON struct {
	EnvironmentClass      apijson.Field
	ID                    apijson.Field
	AutomationsFilePath   apijson.Field
	DesiredPhase          apijson.Field
	DevcontainerFilePath  apijson.Field
	EnvironmentClasses    apijson.Field
	Initializer           apijson.Field
	Metadata              apijson.Field
	PrebuildConfiguration apijson.Field
	RecommendedEditors    apijson.Field
	TechnicalDescription  apijson.Field
	UsedBy                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Project) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectJSON) RawJSON() string {
	return r.raw
}

type ProjectUsedBy struct {
	// Subjects are the 10 most recent subjects who have used the project to create an
	// environment
	Subjects []shared.Subject `json:"subjects"`
	// Total number of unique subjects who have used the project
	TotalSubjects int64             `json:"totalSubjects"`
	JSON          projectUsedByJSON `json:"-"`
}

// projectUsedByJSON contains the JSON metadata for the struct [ProjectUsedBy]
type projectUsedByJSON struct {
	Subjects      apijson.Field
	TotalSubjects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectUsedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUsedByJSON) RawJSON() string {
	return r.raw
}

type ProjectMetadata struct {
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// creator is the identity of the project creator
	Creator shared.Subject `json:"creator"`
	// name is the human readable name of the project
	Name string `json:"name"`
	// organization_id is the ID of the organization that contains the environment
	OrganizationID string `json:"organizationId" format:"uuid"`
	// A Timestamp represents a point in time independent of any time zone or local
	// calendar, encoded as a count of seconds and fractions of seconds at nanosecond
	// resolution. The count is relative to an epoch at UTC midnight on January 1,
	// 1970, in the proleptic Gregorian calendar which extends the Gregorian calendar
	// backwards to year one.
	//
	// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
	// second table is needed for interpretation, using a
	// [24-hour linear smear](https://developers.google.com/time/smear).
	//
	// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
	// restricting to that range, we ensure that we can convert to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
	//
	// # Examples
	//
	// Example 1: Compute Timestamp from POSIX `time()`.
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(time(NULL));
	//	timestamp.set_nanos(0);
	//
	// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
	//
	//	struct timeval tv;
	//	gettimeofday(&tv, NULL);
	//
	//	Timestamp timestamp;
	//	timestamp.set_seconds(tv.tv_sec);
	//	timestamp.set_nanos(tv.tv_usec * 1000);
	//
	// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
	//
	//	FILETIME ft;
	//	GetSystemTimeAsFileTime(&ft);
	//	UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
	//
	//	// A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
	//	// is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
	//	Timestamp timestamp;
	//	timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
	//	timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
	//
	// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
	//
	//	long millis = System.currentTimeMillis();
	//
	//	Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
	//	    .setNanos((int) ((millis % 1000) * 1000000)).build();
	//
	// Example 5: Compute Timestamp from Java `Instant.now()`.
	//
	//	Instant now = Instant.now();
	//
	//	Timestamp timestamp =
	//	    Timestamp.newBuilder().setSeconds(now.getEpochSecond())
	//	        .setNanos(now.getNano()).build();
	//
	// Example 6: Compute Timestamp from current time in Python.
	//
	//	timestamp = Timestamp()
	//	timestamp.GetCurrentTime()
	//
	// # JSON Mapping
	//
	// In JSON format, the Timestamp type is encoded as a string in the
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the format is
	// "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z" where {year} is always
	// expressed using four digits while {month}, {day}, {hour}, {min}, and {sec} are
	// zero-padded to two digits each. The fractional seconds, which can go up to 9
	// digits (i.e. up to 1 nanosecond resolution), are optional. The "Z" suffix
	// indicates the timezone ("UTC"); the timezone is required. A proto3 JSON
	// serializer should always use UTC (as indicated by "Z") when printing the
	// Timestamp type and a proto3 JSON parser should be able to accept both UTC and
	// other timezones (as indicated by an offset).
	//
	// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past 01:30 UTC on
	// January 15, 2017.
	//
	// In JavaScript, one can convert a Date object to this format using the standard
	// [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
	// method. In Python, a standard `datetime.datetime` object can be converted to
	// this format using
	// [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with the
	// time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use the
	// Joda Time's
	// [`ISODateTimeFormat.dateTime()`](<http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()>)
	// to obtain a formatter capable of generating timestamps in this format.
	UpdatedAt time.Time           `json:"updatedAt" format:"date-time"`
	JSON      projectMetadataJSON `json:"-"`
}

// projectMetadataJSON contains the JSON metadata for the struct [ProjectMetadata]
type projectMetadataJSON struct {
	CreatedAt      apijson.Field
	Creator        apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectMetadataJSON) RawJSON() string {
	return r.raw
}

type ProjectPhase string

const (
	ProjectPhaseUnspecified ProjectPhase = "PROJECT_PHASE_UNSPECIFIED"
	ProjectPhaseActive      ProjectPhase = "PROJECT_PHASE_ACTIVE"
	ProjectPhaseDeleted     ProjectPhase = "PROJECT_PHASE_DELETED"
)

func (r ProjectPhase) IsKnown() bool {
	switch r {
	case ProjectPhaseUnspecified, ProjectPhaseActive, ProjectPhaseDeleted:
		return true
	}
	return false
}

// ProjectPrebuildConfiguration defines how prebuilds are created for a project.
// Prebuilds create environment snapshots that enable faster environment startup
// times.
type ProjectPrebuildConfiguration struct {
	// enabled controls whether prebuilds are created for this project. When disabled,
	// no automatic prebuilds will be triggered.
	Enabled bool `json:"enabled"`
	// enable_jetbrains_warmup controls whether JetBrains IDE warmup runs during
	// prebuilds.
	EnableJetbrainsWarmup bool `json:"enableJetbrainsWarmup"`
	// environment_class_ids specifies which environment classes should have prebuilds
	// created. If empty, no prebuilds are created.
	EnvironmentClassIDs []string `json:"environmentClassIds" format:"uuid"`
	// executor specifies who runs prebuilds for this project. The executor's SCM
	// credentials are used to clone the repository. If not set, defaults to the
	// project creator.
	Executor shared.Subject `json:"executor"`
	// timeout is the maximum duration allowed for a prebuild to complete. If not
	// specified, defaults to 1 hour. Must be between 5 minutes and 2 hours.
	Timeout string `json:"timeout" format:"regex"`
	// trigger defines when prebuilds should be created.
	Trigger ProjectPrebuildConfigurationTrigger `json:"trigger"`
	JSON    projectPrebuildConfigurationJSON    `json:"-"`
}

// projectPrebuildConfigurationJSON contains the JSON metadata for the struct
// [ProjectPrebuildConfiguration]
type projectPrebuildConfigurationJSON struct {
	Enabled               apijson.Field
	EnableJetbrainsWarmup apijson.Field
	EnvironmentClassIDs   apijson.Field
	Executor              apijson.Field
	Timeout               apijson.Field
	Trigger               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ProjectPrebuildConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPrebuildConfigurationJSON) RawJSON() string {
	return r.raw
}

// trigger defines when prebuilds should be created.
type ProjectPrebuildConfigurationTrigger struct {
	// daily_schedule triggers a prebuild once per day at the specified hour (UTC). The
	// actual start time may vary slightly to distribute system load.
	DailySchedule ProjectPrebuildConfigurationTriggerDailySchedule `json:"dailySchedule,required"`
	JSON          projectPrebuildConfigurationTriggerJSON          `json:"-"`
}

// projectPrebuildConfigurationTriggerJSON contains the JSON metadata for the
// struct [ProjectPrebuildConfigurationTrigger]
type projectPrebuildConfigurationTriggerJSON struct {
	DailySchedule apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectPrebuildConfigurationTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPrebuildConfigurationTriggerJSON) RawJSON() string {
	return r.raw
}

// daily_schedule triggers a prebuild once per day at the specified hour (UTC). The
// actual start time may vary slightly to distribute system load.
type ProjectPrebuildConfigurationTriggerDailySchedule struct {
	// hour_utc is the hour of day (0-23) in UTC when the prebuild should start. The
	// actual start time may be adjusted by a few minutes to balance system load.
	HourUtc int64                                                `json:"hourUtc"`
	JSON    projectPrebuildConfigurationTriggerDailyScheduleJSON `json:"-"`
}

// projectPrebuildConfigurationTriggerDailyScheduleJSON contains the JSON metadata
// for the struct [ProjectPrebuildConfigurationTriggerDailySchedule]
type projectPrebuildConfigurationTriggerDailyScheduleJSON struct {
	HourUtc     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPrebuildConfigurationTriggerDailySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPrebuildConfigurationTriggerDailyScheduleJSON) RawJSON() string {
	return r.raw
}

// ProjectPrebuildConfiguration defines how prebuilds are created for a project.
// Prebuilds create environment snapshots that enable faster environment startup
// times.
type ProjectPrebuildConfigurationParam struct {
	// enabled controls whether prebuilds are created for this project. When disabled,
	// no automatic prebuilds will be triggered.
	Enabled param.Field[bool] `json:"enabled"`
	// enable_jetbrains_warmup controls whether JetBrains IDE warmup runs during
	// prebuilds.
	EnableJetbrainsWarmup param.Field[bool] `json:"enableJetbrainsWarmup"`
	// environment_class_ids specifies which environment classes should have prebuilds
	// created. If empty, no prebuilds are created.
	EnvironmentClassIDs param.Field[[]string] `json:"environmentClassIds" format:"uuid"`
	// executor specifies who runs prebuilds for this project. The executor's SCM
	// credentials are used to clone the repository. If not set, defaults to the
	// project creator.
	Executor param.Field[shared.SubjectParam] `json:"executor"`
	// timeout is the maximum duration allowed for a prebuild to complete. If not
	// specified, defaults to 1 hour. Must be between 5 minutes and 2 hours.
	Timeout param.Field[string] `json:"timeout" format:"regex"`
	// trigger defines when prebuilds should be created.
	Trigger param.Field[ProjectPrebuildConfigurationTriggerParam] `json:"trigger"`
}

func (r ProjectPrebuildConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// trigger defines when prebuilds should be created.
type ProjectPrebuildConfigurationTriggerParam struct {
	// daily_schedule triggers a prebuild once per day at the specified hour (UTC). The
	// actual start time may vary slightly to distribute system load.
	DailySchedule param.Field[ProjectPrebuildConfigurationTriggerDailyScheduleParam] `json:"dailySchedule,required"`
}

func (r ProjectPrebuildConfigurationTriggerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// daily_schedule triggers a prebuild once per day at the specified hour (UTC). The
// actual start time may vary slightly to distribute system load.
type ProjectPrebuildConfigurationTriggerDailyScheduleParam struct {
	// hour_utc is the hour of day (0-23) in UTC when the prebuild should start. The
	// actual start time may be adjusted by a few minutes to balance system load.
	HourUtc param.Field[int64] `json:"hourUtc"`
}

func (r ProjectPrebuildConfigurationTriggerDailyScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// RecommendedEditors contains the map of recommended editors and their versions.
type RecommendedEditors struct {
	// editors maps editor aliases to their recommended versions. Key is the editor
	// alias (e.g., "intellij", "goland", "vscode"). Value contains the list of
	// recommended versions for that editor. If versions list is empty, all available
	// versions are recommended. Example: {"intellij": {versions: ["2025.1",
	// "2024.3"]}, "goland": {}}
	Editors map[string]RecommendedEditorsEditor `json:"editors"`
	JSON    recommendedEditorsJSON              `json:"-"`
}

// recommendedEditorsJSON contains the JSON metadata for the struct
// [RecommendedEditors]
type recommendedEditorsJSON struct {
	Editors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecommendedEditors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recommendedEditorsJSON) RawJSON() string {
	return r.raw
}

// EditorVersions contains the recommended versions for an editor.
type RecommendedEditorsEditor struct {
	// versions is the list of recommended versions for this editor. If empty, all
	// available versions are recommended. Examples for JetBrains: ["2025.1", "2024.3"]
	Versions []string                     `json:"versions"`
	JSON     recommendedEditorsEditorJSON `json:"-"`
}

// recommendedEditorsEditorJSON contains the JSON metadata for the struct
// [RecommendedEditorsEditor]
type recommendedEditorsEditorJSON struct {
	Versions    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecommendedEditorsEditor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recommendedEditorsEditorJSON) RawJSON() string {
	return r.raw
}

// RecommendedEditors contains the map of recommended editors and their versions.
type RecommendedEditorsParam struct {
	// editors maps editor aliases to their recommended versions. Key is the editor
	// alias (e.g., "intellij", "goland", "vscode"). Value contains the list of
	// recommended versions for that editor. If versions list is empty, all available
	// versions are recommended. Example: {"intellij": {versions: ["2025.1",
	// "2024.3"]}, "goland": {}}
	Editors param.Field[map[string]RecommendedEditorsEditorParam] `json:"editors"`
}

func (r RecommendedEditorsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EditorVersions contains the recommended versions for an editor.
type RecommendedEditorsEditorParam struct {
	// versions is the list of recommended versions for this editor. If empty, all
	// available versions are recommended. Examples for JetBrains: ["2025.1", "2024.3"]
	Versions param.Field[[]string] `json:"versions"`
}

func (r RecommendedEditorsEditorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectNewResponse struct {
	Project Project                `json:"project"`
	JSON    projectNewResponseJSON `json:"-"`
}

// projectNewResponseJSON contains the JSON metadata for the struct
// [ProjectNewResponse]
type projectNewResponseJSON struct {
	Project     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponse struct {
	Project Project                `json:"project"`
	JSON    projectGetResponseJSON `json:"-"`
}

// projectGetResponseJSON contains the JSON metadata for the struct
// [ProjectGetResponse]
type projectGetResponseJSON struct {
	Project     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectUpdateResponse struct {
	Project Project                   `json:"project"`
	JSON    projectUpdateResponseJSON `json:"-"`
}

// projectUpdateResponseJSON contains the JSON metadata for the struct
// [ProjectUpdateResponse]
type projectUpdateResponseJSON struct {
	Project     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectDeleteResponse = interface{}

type ProjectBulkNewResponse struct {
	// created_projects contains the successfully created projects
	CreatedProjects []Project `json:"createdProjects"`
	// failed_projects contains details about projects that failed to create
	FailedProjects []ProjectBulkNewResponseFailedProject `json:"failedProjects"`
	JSON           projectBulkNewResponseJSON            `json:"-"`
}

// projectBulkNewResponseJSON contains the JSON metadata for the struct
// [ProjectBulkNewResponse]
type projectBulkNewResponseJSON struct {
	CreatedProjects apijson.Field
	FailedProjects  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProjectBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectBulkNewResponseFailedProject struct {
	// error describes why the project creation failed
	Error string `json:"error"`
	// index is the position in the request array (0-based)
	Index int64 `json:"index"`
	// name is the project name that failed
	Name string                                  `json:"name"`
	JSON projectBulkNewResponseFailedProjectJSON `json:"-"`
}

// projectBulkNewResponseFailedProjectJSON contains the JSON metadata for the
// struct [ProjectBulkNewResponseFailedProject]
type projectBulkNewResponseFailedProjectJSON struct {
	Error       apijson.Field
	Index       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectBulkNewResponseFailedProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectBulkNewResponseFailedProjectJSON) RawJSON() string {
	return r.raw
}

type ProjectBulkDeleteResponse struct {
	// deleted_project_ids contains the IDs of successfully deleted projects
	DeletedProjectIDs []string `json:"deletedProjectIds"`
	// failed_projects contains details about projects that failed to delete
	FailedProjects []ProjectBulkDeleteResponseFailedProject `json:"failedProjects"`
	JSON           projectBulkDeleteResponseJSON            `json:"-"`
}

// projectBulkDeleteResponseJSON contains the JSON metadata for the struct
// [ProjectBulkDeleteResponse]
type projectBulkDeleteResponseJSON struct {
	DeletedProjectIDs apijson.Field
	FailedProjects    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectBulkDeleteResponseFailedProject struct {
	// error describes why the project deletion failed
	Error string `json:"error"`
	// index is the position in the request array (0-based)
	Index int64 `json:"index"`
	// project_id is the project ID that failed
	ProjectID string                                     `json:"projectId"`
	JSON      projectBulkDeleteResponseFailedProjectJSON `json:"-"`
}

// projectBulkDeleteResponseFailedProjectJSON contains the JSON metadata for the
// struct [ProjectBulkDeleteResponseFailedProject]
type projectBulkDeleteResponseFailedProjectJSON struct {
	Error       apijson.Field
	Index       apijson.Field
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectBulkDeleteResponseFailedProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectBulkDeleteResponseFailedProjectJSON) RawJSON() string {
	return r.raw
}

type ProjectBulkUpdateResponse struct {
	// failed_projects contains details about projects that failed to update
	FailedProjects []ProjectBulkUpdateResponseFailedProject `json:"failedProjects"`
	// updated_projects contains the successfully updated projects
	UpdatedProjects []Project                     `json:"updatedProjects"`
	JSON            projectBulkUpdateResponseJSON `json:"-"`
}

// projectBulkUpdateResponseJSON contains the JSON metadata for the struct
// [ProjectBulkUpdateResponse]
type projectBulkUpdateResponseJSON struct {
	FailedProjects  apijson.Field
	UpdatedProjects apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProjectBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectBulkUpdateResponseFailedProject struct {
	// error describes why the project update failed
	Error string `json:"error"`
	// index is the position in the request array (0-based)
	Index int64 `json:"index"`
	// project_id is the project ID that failed
	ProjectID string                                     `json:"projectId"`
	JSON      projectBulkUpdateResponseFailedProjectJSON `json:"-"`
}

// projectBulkUpdateResponseFailedProjectJSON contains the JSON metadata for the
// struct [ProjectBulkUpdateResponseFailedProject]
type projectBulkUpdateResponseFailedProjectJSON struct {
	Error       apijson.Field
	Index       apijson.Field
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectBulkUpdateResponseFailedProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectBulkUpdateResponseFailedProjectJSON) RawJSON() string {
	return r.raw
}

type ProjectNewFromEnvironmentResponse struct {
	Project Project                               `json:"project"`
	JSON    projectNewFromEnvironmentResponseJSON `json:"-"`
}

// projectNewFromEnvironmentResponseJSON contains the JSON metadata for the struct
// [ProjectNewFromEnvironmentResponse]
type projectNewFromEnvironmentResponseJSON struct {
	Project     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectNewParams struct {
	// initializer is the content initializer
	Initializer param.Field[EnvironmentInitializerParam] `json:"initializer,required"`
	// automations_file_path is the path to the automations file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath"`
	Name                 param.Field[string] `json:"name"`
	// prebuild_configuration defines how prebuilds are created for this project. If
	// not set, prebuilds are disabled for the project.
	PrebuildConfiguration param.Field[ProjectPrebuildConfigurationParam] `json:"prebuildConfiguration"`
	// technical_description is a detailed technical description of the project This
	// field is not returned by default in GetProject or ListProjects responses 8KB max
	TechnicalDescription param.Field[string] `json:"technicalDescription"`
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectGetParams struct {
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
}

func (r ProjectGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectUpdateParams struct {
	// automations_file_path is the path to the automations file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath"`
	// initializer is the content initializer
	Initializer param.Field[EnvironmentInitializerParam] `json:"initializer"`
	Name        param.Field[string]                      `json:"name"`
	// prebuild_configuration defines how prebuilds are created for this project. If
	// not provided, the existing prebuild configuration is not modified. To disable
	// prebuilds, set enabled to false.
	PrebuildConfiguration param.Field[ProjectPrebuildConfigurationParam] `json:"prebuildConfiguration"`
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
	// recommended_editors specifies the editors recommended for this project. If not
	// provided, the existing recommended editors are not modified. To clear all
	// recommended editors, set to an empty RecommendedEditors message.
	RecommendedEditors param.Field[RecommendedEditorsParam] `json:"recommendedEditors"`
	// technical_description is a detailed technical description of the project This
	// field is not returned by default in GetProject or ListProjects responses 8KB max
	TechnicalDescription param.Field[string] `json:"technicalDescription"`
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectListParams struct {
	Token    param.Field[string]                  `query:"token"`
	PageSize param.Field[int64]                   `query:"pageSize"`
	Filter   param.Field[ProjectListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing organizations
	Pagination param.Field[ProjectListParamsPagination] `json:"pagination"`
}

func (r ProjectListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectListParamsFilter struct {
	// project_ids filters the response to only projects with these IDs
	ProjectIDs param.Field[[]string] `json:"projectIds" format:"uuid"`
	// runner_ids filters the response to only projects that use environment classes
	// from these runners
	RunnerIDs param.Field[[]string] `json:"runnerIds" format:"uuid"`
	// runner_kinds filters the response to only projects that use environment classes
	// from runners of these kinds
	RunnerKinds param.Field[[]RunnerKind] `json:"runnerKinds"`
	// search performs case-insensitive search across project name, project ID, and
	// repository name
	Search param.Field[string] `json:"search"`
}

func (r ProjectListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing organizations
type ProjectListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r ProjectListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectDeleteParams struct {
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
}

func (r ProjectDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectBulkNewParams struct {
	Projects param.Field[[]ProjectBulkNewParamsProject] `json:"projects"`
}

func (r ProjectBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectBulkNewParamsProject struct {
	// initializer is the content initializer
	Initializer param.Field[EnvironmentInitializerParam] `json:"initializer,required"`
	// automations_file_path is the path to the automations file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath"`
	Name                 param.Field[string] `json:"name"`
	// prebuild_configuration defines how prebuilds are created for this project. If
	// not set, prebuilds are disabled for the project.
	PrebuildConfiguration param.Field[ProjectPrebuildConfigurationParam] `json:"prebuildConfiguration"`
	// technical_description is a detailed technical description of the project This
	// field is not returned by default in GetProject or ListProjects responses 8KB max
	TechnicalDescription param.Field[string] `json:"technicalDescription"`
}

func (r ProjectBulkNewParamsProject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectBulkDeleteParams struct {
	ProjectIDs param.Field[[]string] `json:"projectIds" format:"uuid"`
}

func (r ProjectBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectBulkUpdateParams struct {
	Projects param.Field[[]ProjectBulkUpdateParamsProject] `json:"projects"`
}

func (r ProjectBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectBulkUpdateParamsProject struct {
	// automations_file_path is the path to the automations file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath"`
	// initializer is the content initializer
	Initializer param.Field[EnvironmentInitializerParam] `json:"initializer"`
	Name        param.Field[string]                      `json:"name"`
	// prebuild_configuration defines how prebuilds are created for this project. If
	// not provided, the existing prebuild configuration is not modified. To disable
	// prebuilds, set enabled to false.
	PrebuildConfiguration param.Field[ProjectPrebuildConfigurationParam] `json:"prebuildConfiguration"`
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
	// recommended_editors specifies the editors recommended for this project. If not
	// provided, the existing recommended editors are not modified. To clear all
	// recommended editors, set to an empty RecommendedEditors message.
	RecommendedEditors param.Field[RecommendedEditorsParam] `json:"recommendedEditors"`
	// technical_description is a detailed technical description of the project This
	// field is not returned by default in GetProject or ListProjects responses 8KB max
	TechnicalDescription param.Field[string] `json:"technicalDescription"`
}

func (r ProjectBulkUpdateParamsProject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectNewFromEnvironmentParams struct {
	// environment_id specifies the environment identifier
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
	Name          param.Field[string] `json:"name"`
}

func (r ProjectNewFromEnvironmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
