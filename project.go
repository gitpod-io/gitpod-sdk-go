// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
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
	Options  []option.RequestOption
	Policies *ProjectPolicyService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r *ProjectService) {
	r = &ProjectService{}
	r.Options = opts
	r.Policies = NewProjectPolicyService(opts...)
	return
}

// CreateProject creates a new Project.
func (r *ProjectService) New(ctx context.Context, body ProjectNewParams, opts ...option.RequestOption) (res *ProjectNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/CreateProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetProject retrieves a single Project.
func (r *ProjectService) Get(ctx context.Context, body ProjectGetParams, opts ...option.RequestOption) (res *ProjectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/GetProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateProject updates the properties of a Project.
func (r *ProjectService) Update(ctx context.Context, body ProjectUpdateParams, opts ...option.RequestOption) (res *ProjectUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/UpdateProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListProjects lists all projects the caller has access to.
func (r *ProjectService) List(ctx context.Context, params ProjectListParams, opts ...option.RequestOption) (res *pagination.ProjectsPage[Project], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// ListProjects lists all projects the caller has access to.
func (r *ProjectService) ListAutoPaging(ctx context.Context, params ProjectListParams, opts ...option.RequestOption) *pagination.ProjectsPageAutoPager[Project] {
	return pagination.NewProjectsPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteProject deletes the specified project.
func (r *ProjectService) Delete(ctx context.Context, body ProjectDeleteParams, opts ...option.RequestOption) (res *ProjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/DeleteProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// CreateProject creates a new Project using an environment as template.
func (r *ProjectService) NewFromEnvironment(ctx context.Context, body ProjectNewFromEnvironmentParams, opts ...option.RequestOption) (res *ProjectNewFromEnvironmentResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
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

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type EnvironmentInitializerSpecsGitTargetMode string

const (
	EnvironmentInitializerSpecsGitTargetModeCloneTargetModeUnspecified  EnvironmentInitializerSpecsGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteHead   EnvironmentInitializerSpecsGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteCommit EnvironmentInitializerSpecsGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteBranch EnvironmentInitializerSpecsGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	EnvironmentInitializerSpecsGitTargetModeCloneTargetModeLocalBranch  EnvironmentInitializerSpecsGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r EnvironmentInitializerSpecsGitTargetMode) IsKnown() bool {
	switch r {
	case EnvironmentInitializerSpecsGitTargetModeCloneTargetModeUnspecified, EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteHead, EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteCommit, EnvironmentInitializerSpecsGitTargetModeCloneTargetModeRemoteBranch, EnvironmentInitializerSpecsGitTargetModeCloneTargetModeLocalBranch:
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
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode param.Field[EnvironmentInitializerSpecsGitTargetMode] `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri param.Field[string] `json:"upstreamRemoteUri"`
}

func (r EnvironmentInitializerSpecsGitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Project struct {
	EnvironmentClass ProjectEnvironmentClass `json:"environmentClass,required"`
	// id is the unique identifier for the project
	ID string `json:"id" format:"uuid"`
	// automations_file_path is the path to the automations file relative to the repo
	// root
	AutomationsFilePath string `json:"automationsFilePath"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer EnvironmentInitializer `json:"initializer"`
	Metadata    ProjectMetadata        `json:"metadata"`
	UsedBy      ProjectUsedBy          `json:"usedBy"`
	JSON        projectJSON            `json:"-"`
}

// projectJSON contains the JSON metadata for the struct [Project]
type projectJSON struct {
	EnvironmentClass     apijson.Field
	ID                   apijson.Field
	AutomationsFilePath  apijson.Field
	DevcontainerFilePath apijson.Field
	Initializer          apijson.Field
	Metadata             apijson.Field
	UsedBy               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
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

type ProjectEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                        `json:"localRunner"`
	JSON        projectEnvironmentClassJSON `json:"-"`
}

// projectEnvironmentClassJSON contains the JSON metadata for the struct
// [ProjectEnvironmentClass]
type projectEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type ProjectEnvironmentClassParam struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID param.Field[string] `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner param.Field[bool] `json:"localRunner"`
}

func (r ProjectEnvironmentClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	EnvironmentClass param.Field[ProjectEnvironmentClassParam] `json:"environmentClass,required"`
	// EnvironmentInitializer specifies how an environment is to be initialized
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
	DevcontainerFilePath param.Field[string]                       `json:"devcontainerFilePath"`
	EnvironmentClass     param.Field[ProjectEnvironmentClassParam] `json:"environmentClass"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer param.Field[EnvironmentInitializerParam] `json:"initializer"`
	Name        param.Field[string]                      `json:"name"`
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
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

type ProjectNewFromEnvironmentParams struct {
	// environment_id specifies the environment identifier
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
	Name          param.Field[string] `json:"name"`
}

func (r ProjectNewFromEnvironmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
