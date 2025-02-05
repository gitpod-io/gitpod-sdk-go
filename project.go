// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
	"github.com/stainless-sdks/gitpod-go/packages/pagination"
	"github.com/tidwall/gjson"
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
func (r *ProjectService) List(ctx context.Context, params ProjectListParams, opts ...option.RequestOption) (res *pagination.ProjectsPage[ProjectListResponse], err error) {
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
func (r *ProjectService) ListAutoPaging(ctx context.Context, params ProjectListParams, opts ...option.RequestOption) *pagination.ProjectsPageAutoPager[ProjectListResponse] {
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

type ProjectNewResponse struct {
	Project ProjectNewResponseProject `json:"project"`
	JSON    projectNewResponseJSON    `json:"-"`
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

type ProjectNewResponseProject struct {
	EnvironmentClass ProjectNewResponseProjectEnvironmentClass `json:"environmentClass,required"`
	// id is the unique identifier for the project
	ID string `json:"id" format:"uuid"`
	// automations_file_path is the path to the automations file relative to the repo
	// root
	AutomationsFilePath string `json:"automationsFilePath"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer ProjectNewResponseProjectInitializer `json:"initializer"`
	Metadata    ProjectNewResponseProjectMetadata    `json:"metadata"`
	UsedBy      ProjectNewResponseProjectUsedBy      `json:"usedBy"`
	JSON        projectNewResponseProjectJSON        `json:"-"`
}

// projectNewResponseProjectJSON contains the JSON metadata for the struct
// [ProjectNewResponseProject]
type projectNewResponseProjectJSON struct {
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

func (r *ProjectNewResponseProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseProjectEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                                          `json:"localRunner"`
	JSON        projectNewResponseProjectEnvironmentClassJSON `json:"-"`
	union       ProjectNewResponseProjectEnvironmentClassUnion
}

// projectNewResponseProjectEnvironmentClassJSON contains the JSON metadata for the
// struct [ProjectNewResponseProjectEnvironmentClass]
type projectNewResponseProjectEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r projectNewResponseProjectEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectNewResponseProjectEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectNewResponseProjectEnvironmentClass{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectNewResponseProjectEnvironmentClassUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectNewResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass],
// [ProjectNewResponseProjectEnvironmentClassUseALocalRunnerForTheUser].
func (r ProjectNewResponseProjectEnvironmentClass) AsUnion() ProjectNewResponseProjectEnvironmentClassUnion {
	return r.union
}

// Union satisfied by
// [ProjectNewResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass]
// or [ProjectNewResponseProjectEnvironmentClassUseALocalRunnerForTheUser].
type ProjectNewResponseProjectEnvironmentClassUnion interface {
	implementsProjectNewResponseProjectEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewResponseProjectEnvironmentClassUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewResponseProjectEnvironmentClassUseALocalRunnerForTheUser{}),
		},
	)
}

type ProjectNewResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string                                                                                                                        `json:"environmentClassId,required" format:"uuid"`
	JSON               projectNewResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON `json:"-"`
}

// projectNewResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass]
type projectNewResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectNewResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) implementsProjectNewResponseProjectEnvironmentClass() {
}

type ProjectNewResponseProjectEnvironmentClassUseALocalRunnerForTheUser struct {
	// Use a local runner for the user
	LocalRunner bool                                                                   `json:"localRunner,required"`
	JSON        projectNewResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON `json:"-"`
}

// projectNewResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseProjectEnvironmentClassUseALocalRunnerForTheUser]
type projectNewResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON struct {
	LocalRunner apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseProjectEnvironmentClassUseALocalRunnerForTheUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseProjectEnvironmentClassUseALocalRunnerForTheUser) implementsProjectNewResponseProjectEnvironmentClass() {
}

// EnvironmentInitializer specifies how an environment is to be initialized
type ProjectNewResponseProjectInitializer struct {
	Specs []ProjectNewResponseProjectInitializerSpec `json:"specs"`
	JSON  projectNewResponseProjectInitializerJSON   `json:"-"`
}

// projectNewResponseProjectInitializerJSON contains the JSON metadata for the
// struct [ProjectNewResponseProjectInitializer]
type projectNewResponseProjectInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseProjectInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectInitializerJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseProjectInitializerSpec struct {
	// This field can have the runtime type of
	// [ProjectNewResponseProjectInitializerSpecsContextURLContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [ProjectNewResponseProjectInitializerSpecsGitGit].
	Git   interface{}                                  `json:"git"`
	JSON  projectNewResponseProjectInitializerSpecJSON `json:"-"`
	union ProjectNewResponseProjectInitializerSpecsUnion
}

// projectNewResponseProjectInitializerSpecJSON contains the JSON metadata for the
// struct [ProjectNewResponseProjectInitializerSpec]
type projectNewResponseProjectInitializerSpecJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectNewResponseProjectInitializerSpecJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectNewResponseProjectInitializerSpec) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectNewResponseProjectInitializerSpec{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectNewResponseProjectInitializerSpecsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectNewResponseProjectInitializerSpecsContextURL],
// [ProjectNewResponseProjectInitializerSpecsGit].
func (r ProjectNewResponseProjectInitializerSpec) AsUnion() ProjectNewResponseProjectInitializerSpecsUnion {
	return r.union
}

// Union satisfied by [ProjectNewResponseProjectInitializerSpecsContextURL] or
// [ProjectNewResponseProjectInitializerSpecsGit].
type ProjectNewResponseProjectInitializerSpecsUnion interface {
	implementsProjectNewResponseProjectInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewResponseProjectInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewResponseProjectInitializerSpecsContextURL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewResponseProjectInitializerSpecsGit{}),
		},
	)
}

type ProjectNewResponseProjectInitializerSpecsContextURL struct {
	ContextURL ProjectNewResponseProjectInitializerSpecsContextURLContextURL `json:"contextUrl,required"`
	JSON       projectNewResponseProjectInitializerSpecsContextURLJSON       `json:"-"`
}

// projectNewResponseProjectInitializerSpecsContextURLJSON contains the JSON
// metadata for the struct [ProjectNewResponseProjectInitializerSpecsContextURL]
type projectNewResponseProjectInitializerSpecsContextURLJSON struct {
	ContextURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseProjectInitializerSpecsContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectInitializerSpecsContextURLJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseProjectInitializerSpecsContextURL) implementsProjectNewResponseProjectInitializerSpec() {
}

type ProjectNewResponseProjectInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                            `json:"url" format:"uri"`
	JSON projectNewResponseProjectInitializerSpecsContextURLContextURLJSON `json:"-"`
}

// projectNewResponseProjectInitializerSpecsContextURLContextURLJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseProjectInitializerSpecsContextURLContextURL]
type projectNewResponseProjectInitializerSpecsContextURLContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseProjectInitializerSpecsContextURLContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectInitializerSpecsContextURLContextURLJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseProjectInitializerSpecsGit struct {
	Git  ProjectNewResponseProjectInitializerSpecsGitGit  `json:"git,required"`
	JSON projectNewResponseProjectInitializerSpecsGitJSON `json:"-"`
}

// projectNewResponseProjectInitializerSpecsGitJSON contains the JSON metadata for
// the struct [ProjectNewResponseProjectInitializerSpecsGit]
type projectNewResponseProjectInitializerSpecsGitJSON struct {
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseProjectInitializerSpecsGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectInitializerSpecsGitJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseProjectInitializerSpecsGit) implementsProjectNewResponseProjectInitializerSpec() {
}

type ProjectNewResponseProjectInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode ProjectNewResponseProjectInitializerSpecsGitGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                              `json:"upstreamRemoteUri"`
	JSON              projectNewResponseProjectInitializerSpecsGitGitJSON `json:"-"`
}

// projectNewResponseProjectInitializerSpecsGitGitJSON contains the JSON metadata
// for the struct [ProjectNewResponseProjectInitializerSpecsGitGit]
type projectNewResponseProjectInitializerSpecsGitGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectNewResponseProjectInitializerSpecsGitGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectInitializerSpecsGitGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectNewResponseProjectInitializerSpecsGitGitTargetMode string

const (
	ProjectNewResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  ProjectNewResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectNewResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   ProjectNewResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectNewResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit ProjectNewResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectNewResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch ProjectNewResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectNewResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  ProjectNewResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectNewResponseProjectInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectNewResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, ProjectNewResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, ProjectNewResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, ProjectNewResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, ProjectNewResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectNewResponseProjectMetadata struct {
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
	Creator ProjectNewResponseProjectMetadataCreator `json:"creator"`
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
	UpdatedAt time.Time                             `json:"updatedAt" format:"date-time"`
	JSON      projectNewResponseProjectMetadataJSON `json:"-"`
}

// projectNewResponseProjectMetadataJSON contains the JSON metadata for the struct
// [ProjectNewResponseProjectMetadata]
type projectNewResponseProjectMetadataJSON struct {
	CreatedAt      apijson.Field
	Creator        apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectNewResponseProjectMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the project creator
type ProjectNewResponseProjectMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectNewResponseProjectMetadataCreatorPrincipal `json:"principal"`
	JSON      projectNewResponseProjectMetadataCreatorJSON      `json:"-"`
}

// projectNewResponseProjectMetadataCreatorJSON contains the JSON metadata for the
// struct [ProjectNewResponseProjectMetadataCreator]
type projectNewResponseProjectMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseProjectMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectNewResponseProjectMetadataCreatorPrincipal string

const (
	ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalUnspecified    ProjectNewResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalAccount        ProjectNewResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalUser           ProjectNewResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_USER"
	ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalRunner         ProjectNewResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalEnvironment    ProjectNewResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalServiceAccount ProjectNewResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectNewResponseProjectMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalUnspecified, ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalAccount, ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalUser, ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalRunner, ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalEnvironment, ProjectNewResponseProjectMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type ProjectNewResponseProjectUsedBy struct {
	// Subjects are the 10 most recent subjects who have used the project to create an
	// environment
	Subjects []ProjectNewResponseProjectUsedBySubject `json:"subjects"`
	// Total number of unique subjects who have used the project
	TotalSubjects int64                               `json:"totalSubjects"`
	JSON          projectNewResponseProjectUsedByJSON `json:"-"`
}

// projectNewResponseProjectUsedByJSON contains the JSON metadata for the struct
// [ProjectNewResponseProjectUsedBy]
type projectNewResponseProjectUsedByJSON struct {
	Subjects      apijson.Field
	TotalSubjects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectNewResponseProjectUsedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectUsedByJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseProjectUsedBySubject struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectNewResponseProjectUsedBySubjectsPrincipal `json:"principal"`
	JSON      projectNewResponseProjectUsedBySubjectJSON       `json:"-"`
}

// projectNewResponseProjectUsedBySubjectJSON contains the JSON metadata for the
// struct [ProjectNewResponseProjectUsedBySubject]
type projectNewResponseProjectUsedBySubjectJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseProjectUsedBySubject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectUsedBySubjectJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectNewResponseProjectUsedBySubjectsPrincipal string

const (
	ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalUnspecified    ProjectNewResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalAccount        ProjectNewResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalUser           ProjectNewResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_USER"
	ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalRunner         ProjectNewResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_RUNNER"
	ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalEnvironment    ProjectNewResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalServiceAccount ProjectNewResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectNewResponseProjectUsedBySubjectsPrincipal) IsKnown() bool {
	switch r {
	case ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalUnspecified, ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalAccount, ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalUser, ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalRunner, ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalEnvironment, ProjectNewResponseProjectUsedBySubjectsPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type ProjectGetResponse struct {
	Project ProjectGetResponseProject `json:"project"`
	JSON    projectGetResponseJSON    `json:"-"`
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

type ProjectGetResponseProject struct {
	EnvironmentClass ProjectGetResponseProjectEnvironmentClass `json:"environmentClass,required"`
	// id is the unique identifier for the project
	ID string `json:"id" format:"uuid"`
	// automations_file_path is the path to the automations file relative to the repo
	// root
	AutomationsFilePath string `json:"automationsFilePath"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer ProjectGetResponseProjectInitializer `json:"initializer"`
	Metadata    ProjectGetResponseProjectMetadata    `json:"metadata"`
	UsedBy      ProjectGetResponseProjectUsedBy      `json:"usedBy"`
	JSON        projectGetResponseProjectJSON        `json:"-"`
}

// projectGetResponseProjectJSON contains the JSON metadata for the struct
// [ProjectGetResponseProject]
type projectGetResponseProjectJSON struct {
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

func (r *ProjectGetResponseProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseProjectEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                                          `json:"localRunner"`
	JSON        projectGetResponseProjectEnvironmentClassJSON `json:"-"`
	union       ProjectGetResponseProjectEnvironmentClassUnion
}

// projectGetResponseProjectEnvironmentClassJSON contains the JSON metadata for the
// struct [ProjectGetResponseProjectEnvironmentClass]
type projectGetResponseProjectEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r projectGetResponseProjectEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectGetResponseProjectEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectGetResponseProjectEnvironmentClass{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectGetResponseProjectEnvironmentClassUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectGetResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass],
// [ProjectGetResponseProjectEnvironmentClassUseALocalRunnerForTheUser].
func (r ProjectGetResponseProjectEnvironmentClass) AsUnion() ProjectGetResponseProjectEnvironmentClassUnion {
	return r.union
}

// Union satisfied by
// [ProjectGetResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass]
// or [ProjectGetResponseProjectEnvironmentClassUseALocalRunnerForTheUser].
type ProjectGetResponseProjectEnvironmentClassUnion interface {
	implementsProjectGetResponseProjectEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectGetResponseProjectEnvironmentClassUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectGetResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectGetResponseProjectEnvironmentClassUseALocalRunnerForTheUser{}),
		},
	)
}

type ProjectGetResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string                                                                                                                        `json:"environmentClassId,required" format:"uuid"`
	JSON               projectGetResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON `json:"-"`
}

// projectGetResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass]
type projectGetResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectGetResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) implementsProjectGetResponseProjectEnvironmentClass() {
}

type ProjectGetResponseProjectEnvironmentClassUseALocalRunnerForTheUser struct {
	// Use a local runner for the user
	LocalRunner bool                                                                   `json:"localRunner,required"`
	JSON        projectGetResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON `json:"-"`
}

// projectGetResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseProjectEnvironmentClassUseALocalRunnerForTheUser]
type projectGetResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON struct {
	LocalRunner apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseProjectEnvironmentClassUseALocalRunnerForTheUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseProjectEnvironmentClassUseALocalRunnerForTheUser) implementsProjectGetResponseProjectEnvironmentClass() {
}

// EnvironmentInitializer specifies how an environment is to be initialized
type ProjectGetResponseProjectInitializer struct {
	Specs []ProjectGetResponseProjectInitializerSpec `json:"specs"`
	JSON  projectGetResponseProjectInitializerJSON   `json:"-"`
}

// projectGetResponseProjectInitializerJSON contains the JSON metadata for the
// struct [ProjectGetResponseProjectInitializer]
type projectGetResponseProjectInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseProjectInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectInitializerJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseProjectInitializerSpec struct {
	// This field can have the runtime type of
	// [ProjectGetResponseProjectInitializerSpecsContextURLContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [ProjectGetResponseProjectInitializerSpecsGitGit].
	Git   interface{}                                  `json:"git"`
	JSON  projectGetResponseProjectInitializerSpecJSON `json:"-"`
	union ProjectGetResponseProjectInitializerSpecsUnion
}

// projectGetResponseProjectInitializerSpecJSON contains the JSON metadata for the
// struct [ProjectGetResponseProjectInitializerSpec]
type projectGetResponseProjectInitializerSpecJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectGetResponseProjectInitializerSpecJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectGetResponseProjectInitializerSpec) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectGetResponseProjectInitializerSpec{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectGetResponseProjectInitializerSpecsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectGetResponseProjectInitializerSpecsContextURL],
// [ProjectGetResponseProjectInitializerSpecsGit].
func (r ProjectGetResponseProjectInitializerSpec) AsUnion() ProjectGetResponseProjectInitializerSpecsUnion {
	return r.union
}

// Union satisfied by [ProjectGetResponseProjectInitializerSpecsContextURL] or
// [ProjectGetResponseProjectInitializerSpecsGit].
type ProjectGetResponseProjectInitializerSpecsUnion interface {
	implementsProjectGetResponseProjectInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectGetResponseProjectInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectGetResponseProjectInitializerSpecsContextURL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectGetResponseProjectInitializerSpecsGit{}),
		},
	)
}

type ProjectGetResponseProjectInitializerSpecsContextURL struct {
	ContextURL ProjectGetResponseProjectInitializerSpecsContextURLContextURL `json:"contextUrl,required"`
	JSON       projectGetResponseProjectInitializerSpecsContextURLJSON       `json:"-"`
}

// projectGetResponseProjectInitializerSpecsContextURLJSON contains the JSON
// metadata for the struct [ProjectGetResponseProjectInitializerSpecsContextURL]
type projectGetResponseProjectInitializerSpecsContextURLJSON struct {
	ContextURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseProjectInitializerSpecsContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectInitializerSpecsContextURLJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseProjectInitializerSpecsContextURL) implementsProjectGetResponseProjectInitializerSpec() {
}

type ProjectGetResponseProjectInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                            `json:"url" format:"uri"`
	JSON projectGetResponseProjectInitializerSpecsContextURLContextURLJSON `json:"-"`
}

// projectGetResponseProjectInitializerSpecsContextURLContextURLJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseProjectInitializerSpecsContextURLContextURL]
type projectGetResponseProjectInitializerSpecsContextURLContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseProjectInitializerSpecsContextURLContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectInitializerSpecsContextURLContextURLJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseProjectInitializerSpecsGit struct {
	Git  ProjectGetResponseProjectInitializerSpecsGitGit  `json:"git,required"`
	JSON projectGetResponseProjectInitializerSpecsGitJSON `json:"-"`
}

// projectGetResponseProjectInitializerSpecsGitJSON contains the JSON metadata for
// the struct [ProjectGetResponseProjectInitializerSpecsGit]
type projectGetResponseProjectInitializerSpecsGitJSON struct {
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseProjectInitializerSpecsGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectInitializerSpecsGitJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseProjectInitializerSpecsGit) implementsProjectGetResponseProjectInitializerSpec() {
}

type ProjectGetResponseProjectInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode ProjectGetResponseProjectInitializerSpecsGitGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                              `json:"upstreamRemoteUri"`
	JSON              projectGetResponseProjectInitializerSpecsGitGitJSON `json:"-"`
}

// projectGetResponseProjectInitializerSpecsGitGitJSON contains the JSON metadata
// for the struct [ProjectGetResponseProjectInitializerSpecsGitGit]
type projectGetResponseProjectInitializerSpecsGitGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectGetResponseProjectInitializerSpecsGitGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectInitializerSpecsGitGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectGetResponseProjectInitializerSpecsGitGitTargetMode string

const (
	ProjectGetResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  ProjectGetResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectGetResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   ProjectGetResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectGetResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit ProjectGetResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectGetResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch ProjectGetResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectGetResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  ProjectGetResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectGetResponseProjectInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectGetResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, ProjectGetResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, ProjectGetResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, ProjectGetResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, ProjectGetResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectGetResponseProjectMetadata struct {
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
	Creator ProjectGetResponseProjectMetadataCreator `json:"creator"`
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
	UpdatedAt time.Time                             `json:"updatedAt" format:"date-time"`
	JSON      projectGetResponseProjectMetadataJSON `json:"-"`
}

// projectGetResponseProjectMetadataJSON contains the JSON metadata for the struct
// [ProjectGetResponseProjectMetadata]
type projectGetResponseProjectMetadataJSON struct {
	CreatedAt      apijson.Field
	Creator        apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectGetResponseProjectMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the project creator
type ProjectGetResponseProjectMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectGetResponseProjectMetadataCreatorPrincipal `json:"principal"`
	JSON      projectGetResponseProjectMetadataCreatorJSON      `json:"-"`
}

// projectGetResponseProjectMetadataCreatorJSON contains the JSON metadata for the
// struct [ProjectGetResponseProjectMetadataCreator]
type projectGetResponseProjectMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseProjectMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectGetResponseProjectMetadataCreatorPrincipal string

const (
	ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalUnspecified    ProjectGetResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalAccount        ProjectGetResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalUser           ProjectGetResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_USER"
	ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalRunner         ProjectGetResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalEnvironment    ProjectGetResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalServiceAccount ProjectGetResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectGetResponseProjectMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalUnspecified, ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalAccount, ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalUser, ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalRunner, ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalEnvironment, ProjectGetResponseProjectMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type ProjectGetResponseProjectUsedBy struct {
	// Subjects are the 10 most recent subjects who have used the project to create an
	// environment
	Subjects []ProjectGetResponseProjectUsedBySubject `json:"subjects"`
	// Total number of unique subjects who have used the project
	TotalSubjects int64                               `json:"totalSubjects"`
	JSON          projectGetResponseProjectUsedByJSON `json:"-"`
}

// projectGetResponseProjectUsedByJSON contains the JSON metadata for the struct
// [ProjectGetResponseProjectUsedBy]
type projectGetResponseProjectUsedByJSON struct {
	Subjects      apijson.Field
	TotalSubjects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectGetResponseProjectUsedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectUsedByJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseProjectUsedBySubject struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectGetResponseProjectUsedBySubjectsPrincipal `json:"principal"`
	JSON      projectGetResponseProjectUsedBySubjectJSON       `json:"-"`
}

// projectGetResponseProjectUsedBySubjectJSON contains the JSON metadata for the
// struct [ProjectGetResponseProjectUsedBySubject]
type projectGetResponseProjectUsedBySubjectJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseProjectUsedBySubject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectUsedBySubjectJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectGetResponseProjectUsedBySubjectsPrincipal string

const (
	ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalUnspecified    ProjectGetResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalAccount        ProjectGetResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalUser           ProjectGetResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_USER"
	ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalRunner         ProjectGetResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_RUNNER"
	ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalEnvironment    ProjectGetResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalServiceAccount ProjectGetResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectGetResponseProjectUsedBySubjectsPrincipal) IsKnown() bool {
	switch r {
	case ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalUnspecified, ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalAccount, ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalUser, ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalRunner, ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalEnvironment, ProjectGetResponseProjectUsedBySubjectsPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type ProjectUpdateResponse struct {
	Project ProjectUpdateResponseProject `json:"project"`
	JSON    projectUpdateResponseJSON    `json:"-"`
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

type ProjectUpdateResponseProject struct {
	EnvironmentClass ProjectUpdateResponseProjectEnvironmentClass `json:"environmentClass,required"`
	// id is the unique identifier for the project
	ID string `json:"id" format:"uuid"`
	// automations_file_path is the path to the automations file relative to the repo
	// root
	AutomationsFilePath string `json:"automationsFilePath"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer ProjectUpdateResponseProjectInitializer `json:"initializer"`
	Metadata    ProjectUpdateResponseProjectMetadata    `json:"metadata"`
	UsedBy      ProjectUpdateResponseProjectUsedBy      `json:"usedBy"`
	JSON        projectUpdateResponseProjectJSON        `json:"-"`
}

// projectUpdateResponseProjectJSON contains the JSON metadata for the struct
// [ProjectUpdateResponseProject]
type projectUpdateResponseProjectJSON struct {
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

func (r *ProjectUpdateResponseProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectJSON) RawJSON() string {
	return r.raw
}

type ProjectUpdateResponseProjectEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                                             `json:"localRunner"`
	JSON        projectUpdateResponseProjectEnvironmentClassJSON `json:"-"`
	union       ProjectUpdateResponseProjectEnvironmentClassUnion
}

// projectUpdateResponseProjectEnvironmentClassJSON contains the JSON metadata for
// the struct [ProjectUpdateResponseProjectEnvironmentClass]
type projectUpdateResponseProjectEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r projectUpdateResponseProjectEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectUpdateResponseProjectEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectUpdateResponseProjectEnvironmentClass{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectUpdateResponseProjectEnvironmentClassUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectUpdateResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass],
// [ProjectUpdateResponseProjectEnvironmentClassUseALocalRunnerForTheUser].
func (r ProjectUpdateResponseProjectEnvironmentClass) AsUnion() ProjectUpdateResponseProjectEnvironmentClassUnion {
	return r.union
}

// Union satisfied by
// [ProjectUpdateResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass]
// or [ProjectUpdateResponseProjectEnvironmentClassUseALocalRunnerForTheUser].
type ProjectUpdateResponseProjectEnvironmentClassUnion interface {
	implementsProjectUpdateResponseProjectEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectUpdateResponseProjectEnvironmentClassUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectUpdateResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectUpdateResponseProjectEnvironmentClassUseALocalRunnerForTheUser{}),
		},
	)
}

type ProjectUpdateResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string                                                                                                                           `json:"environmentClassId,required" format:"uuid"`
	JSON               projectUpdateResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON `json:"-"`
}

// projectUpdateResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON
// contains the JSON metadata for the struct
// [ProjectUpdateResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass]
type projectUpdateResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r ProjectUpdateResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) implementsProjectUpdateResponseProjectEnvironmentClass() {
}

type ProjectUpdateResponseProjectEnvironmentClassUseALocalRunnerForTheUser struct {
	// Use a local runner for the user
	LocalRunner bool                                                                      `json:"localRunner,required"`
	JSON        projectUpdateResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON `json:"-"`
}

// projectUpdateResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON
// contains the JSON metadata for the struct
// [ProjectUpdateResponseProjectEnvironmentClassUseALocalRunnerForTheUser]
type projectUpdateResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON struct {
	LocalRunner apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectEnvironmentClassUseALocalRunnerForTheUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON) RawJSON() string {
	return r.raw
}

func (r ProjectUpdateResponseProjectEnvironmentClassUseALocalRunnerForTheUser) implementsProjectUpdateResponseProjectEnvironmentClass() {
}

// EnvironmentInitializer specifies how an environment is to be initialized
type ProjectUpdateResponseProjectInitializer struct {
	Specs []ProjectUpdateResponseProjectInitializerSpec `json:"specs"`
	JSON  projectUpdateResponseProjectInitializerJSON   `json:"-"`
}

// projectUpdateResponseProjectInitializerJSON contains the JSON metadata for the
// struct [ProjectUpdateResponseProjectInitializer]
type projectUpdateResponseProjectInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectInitializerJSON) RawJSON() string {
	return r.raw
}

type ProjectUpdateResponseProjectInitializerSpec struct {
	// This field can have the runtime type of
	// [ProjectUpdateResponseProjectInitializerSpecsContextURLContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [ProjectUpdateResponseProjectInitializerSpecsGitGit].
	Git   interface{}                                     `json:"git"`
	JSON  projectUpdateResponseProjectInitializerSpecJSON `json:"-"`
	union ProjectUpdateResponseProjectInitializerSpecsUnion
}

// projectUpdateResponseProjectInitializerSpecJSON contains the JSON metadata for
// the struct [ProjectUpdateResponseProjectInitializerSpec]
type projectUpdateResponseProjectInitializerSpecJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectUpdateResponseProjectInitializerSpecJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectUpdateResponseProjectInitializerSpec) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectUpdateResponseProjectInitializerSpec{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectUpdateResponseProjectInitializerSpecsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectUpdateResponseProjectInitializerSpecsContextURL],
// [ProjectUpdateResponseProjectInitializerSpecsGit].
func (r ProjectUpdateResponseProjectInitializerSpec) AsUnion() ProjectUpdateResponseProjectInitializerSpecsUnion {
	return r.union
}

// Union satisfied by [ProjectUpdateResponseProjectInitializerSpecsContextURL] or
// [ProjectUpdateResponseProjectInitializerSpecsGit].
type ProjectUpdateResponseProjectInitializerSpecsUnion interface {
	implementsProjectUpdateResponseProjectInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectUpdateResponseProjectInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectUpdateResponseProjectInitializerSpecsContextURL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectUpdateResponseProjectInitializerSpecsGit{}),
		},
	)
}

type ProjectUpdateResponseProjectInitializerSpecsContextURL struct {
	ContextURL ProjectUpdateResponseProjectInitializerSpecsContextURLContextURL `json:"contextUrl,required"`
	JSON       projectUpdateResponseProjectInitializerSpecsContextURLJSON       `json:"-"`
}

// projectUpdateResponseProjectInitializerSpecsContextURLJSON contains the JSON
// metadata for the struct [ProjectUpdateResponseProjectInitializerSpecsContextURL]
type projectUpdateResponseProjectInitializerSpecsContextURLJSON struct {
	ContextURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectInitializerSpecsContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectInitializerSpecsContextURLJSON) RawJSON() string {
	return r.raw
}

func (r ProjectUpdateResponseProjectInitializerSpecsContextURL) implementsProjectUpdateResponseProjectInitializerSpec() {
}

type ProjectUpdateResponseProjectInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                               `json:"url" format:"uri"`
	JSON projectUpdateResponseProjectInitializerSpecsContextURLContextURLJSON `json:"-"`
}

// projectUpdateResponseProjectInitializerSpecsContextURLContextURLJSON contains
// the JSON metadata for the struct
// [ProjectUpdateResponseProjectInitializerSpecsContextURLContextURL]
type projectUpdateResponseProjectInitializerSpecsContextURLContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectInitializerSpecsContextURLContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectInitializerSpecsContextURLContextURLJSON) RawJSON() string {
	return r.raw
}

type ProjectUpdateResponseProjectInitializerSpecsGit struct {
	Git  ProjectUpdateResponseProjectInitializerSpecsGitGit  `json:"git,required"`
	JSON projectUpdateResponseProjectInitializerSpecsGitJSON `json:"-"`
}

// projectUpdateResponseProjectInitializerSpecsGitJSON contains the JSON metadata
// for the struct [ProjectUpdateResponseProjectInitializerSpecsGit]
type projectUpdateResponseProjectInitializerSpecsGitJSON struct {
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectInitializerSpecsGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectInitializerSpecsGitJSON) RawJSON() string {
	return r.raw
}

func (r ProjectUpdateResponseProjectInitializerSpecsGit) implementsProjectUpdateResponseProjectInitializerSpec() {
}

type ProjectUpdateResponseProjectInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode ProjectUpdateResponseProjectInitializerSpecsGitGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                                 `json:"upstreamRemoteUri"`
	JSON              projectUpdateResponseProjectInitializerSpecsGitGitJSON `json:"-"`
}

// projectUpdateResponseProjectInitializerSpecsGitGitJSON contains the JSON
// metadata for the struct [ProjectUpdateResponseProjectInitializerSpecsGitGit]
type projectUpdateResponseProjectInitializerSpecsGitGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectInitializerSpecsGitGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectInitializerSpecsGitGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectUpdateResponseProjectInitializerSpecsGitGitTargetMode string

const (
	ProjectUpdateResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  ProjectUpdateResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectUpdateResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   ProjectUpdateResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectUpdateResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit ProjectUpdateResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectUpdateResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch ProjectUpdateResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectUpdateResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  ProjectUpdateResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectUpdateResponseProjectInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectUpdateResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, ProjectUpdateResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, ProjectUpdateResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, ProjectUpdateResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, ProjectUpdateResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectUpdateResponseProjectMetadata struct {
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
	Creator ProjectUpdateResponseProjectMetadataCreator `json:"creator"`
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
	UpdatedAt time.Time                                `json:"updatedAt" format:"date-time"`
	JSON      projectUpdateResponseProjectMetadataJSON `json:"-"`
}

// projectUpdateResponseProjectMetadataJSON contains the JSON metadata for the
// struct [ProjectUpdateResponseProjectMetadata]
type projectUpdateResponseProjectMetadataJSON struct {
	CreatedAt      apijson.Field
	Creator        apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the project creator
type ProjectUpdateResponseProjectMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectUpdateResponseProjectMetadataCreatorPrincipal `json:"principal"`
	JSON      projectUpdateResponseProjectMetadataCreatorJSON      `json:"-"`
}

// projectUpdateResponseProjectMetadataCreatorJSON contains the JSON metadata for
// the struct [ProjectUpdateResponseProjectMetadataCreator]
type projectUpdateResponseProjectMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectUpdateResponseProjectMetadataCreatorPrincipal string

const (
	ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalUnspecified    ProjectUpdateResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalAccount        ProjectUpdateResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalUser           ProjectUpdateResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_USER"
	ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalRunner         ProjectUpdateResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalEnvironment    ProjectUpdateResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalServiceAccount ProjectUpdateResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectUpdateResponseProjectMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalUnspecified, ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalAccount, ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalUser, ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalRunner, ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalEnvironment, ProjectUpdateResponseProjectMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type ProjectUpdateResponseProjectUsedBy struct {
	// Subjects are the 10 most recent subjects who have used the project to create an
	// environment
	Subjects []ProjectUpdateResponseProjectUsedBySubject `json:"subjects"`
	// Total number of unique subjects who have used the project
	TotalSubjects int64                                  `json:"totalSubjects"`
	JSON          projectUpdateResponseProjectUsedByJSON `json:"-"`
}

// projectUpdateResponseProjectUsedByJSON contains the JSON metadata for the struct
// [ProjectUpdateResponseProjectUsedBy]
type projectUpdateResponseProjectUsedByJSON struct {
	Subjects      apijson.Field
	TotalSubjects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectUsedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectUsedByJSON) RawJSON() string {
	return r.raw
}

type ProjectUpdateResponseProjectUsedBySubject struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectUpdateResponseProjectUsedBySubjectsPrincipal `json:"principal"`
	JSON      projectUpdateResponseProjectUsedBySubjectJSON       `json:"-"`
}

// projectUpdateResponseProjectUsedBySubjectJSON contains the JSON metadata for the
// struct [ProjectUpdateResponseProjectUsedBySubject]
type projectUpdateResponseProjectUsedBySubjectJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectUsedBySubject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectUsedBySubjectJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectUpdateResponseProjectUsedBySubjectsPrincipal string

const (
	ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalUnspecified    ProjectUpdateResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalAccount        ProjectUpdateResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalUser           ProjectUpdateResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_USER"
	ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalRunner         ProjectUpdateResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_RUNNER"
	ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalEnvironment    ProjectUpdateResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalServiceAccount ProjectUpdateResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectUpdateResponseProjectUsedBySubjectsPrincipal) IsKnown() bool {
	switch r {
	case ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalUnspecified, ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalAccount, ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalUser, ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalRunner, ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalEnvironment, ProjectUpdateResponseProjectUsedBySubjectsPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type ProjectListResponse struct {
	EnvironmentClass ProjectListResponseEnvironmentClass `json:"environmentClass,required"`
	// id is the unique identifier for the project
	ID string `json:"id" format:"uuid"`
	// automations_file_path is the path to the automations file relative to the repo
	// root
	AutomationsFilePath string `json:"automationsFilePath"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer ProjectListResponseInitializer `json:"initializer"`
	Metadata    ProjectListResponseMetadata    `json:"metadata"`
	UsedBy      ProjectListResponseUsedBy      `json:"usedBy"`
	JSON        projectListResponseJSON        `json:"-"`
}

// projectListResponseJSON contains the JSON metadata for the struct
// [ProjectListResponse]
type projectListResponseJSON struct {
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

func (r *ProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                                    `json:"localRunner"`
	JSON        projectListResponseEnvironmentClassJSON `json:"-"`
	union       ProjectListResponseEnvironmentClassUnion
}

// projectListResponseEnvironmentClassJSON contains the JSON metadata for the
// struct [ProjectListResponseEnvironmentClass]
type projectListResponseEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r projectListResponseEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectListResponseEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectListResponseEnvironmentClass{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectListResponseEnvironmentClassUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectListResponseEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass],
// [ProjectListResponseEnvironmentClassUseALocalRunnerForTheUser].
func (r ProjectListResponseEnvironmentClass) AsUnion() ProjectListResponseEnvironmentClassUnion {
	return r.union
}

// Union satisfied by
// [ProjectListResponseEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass]
// or [ProjectListResponseEnvironmentClassUseALocalRunnerForTheUser].
type ProjectListResponseEnvironmentClassUnion interface {
	implementsProjectListResponseEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectListResponseEnvironmentClassUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectListResponseEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectListResponseEnvironmentClassUseALocalRunnerForTheUser{}),
		},
	)
}

type ProjectListResponseEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string                                                                                                                  `json:"environmentClassId,required" format:"uuid"`
	JSON               projectListResponseEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON `json:"-"`
}

// projectListResponseEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON
// contains the JSON metadata for the struct
// [ProjectListResponseEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass]
type projectListResponseEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectListResponseEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) implementsProjectListResponseEnvironmentClass() {
}

type ProjectListResponseEnvironmentClassUseALocalRunnerForTheUser struct {
	// Use a local runner for the user
	LocalRunner bool                                                             `json:"localRunner,required"`
	JSON        projectListResponseEnvironmentClassUseALocalRunnerForTheUserJSON `json:"-"`
}

// projectListResponseEnvironmentClassUseALocalRunnerForTheUserJSON contains the
// JSON metadata for the struct
// [ProjectListResponseEnvironmentClassUseALocalRunnerForTheUser]
type projectListResponseEnvironmentClassUseALocalRunnerForTheUserJSON struct {
	LocalRunner apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseEnvironmentClassUseALocalRunnerForTheUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseEnvironmentClassUseALocalRunnerForTheUserJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseEnvironmentClassUseALocalRunnerForTheUser) implementsProjectListResponseEnvironmentClass() {
}

// EnvironmentInitializer specifies how an environment is to be initialized
type ProjectListResponseInitializer struct {
	Specs []ProjectListResponseInitializerSpec `json:"specs"`
	JSON  projectListResponseInitializerJSON   `json:"-"`
}

// projectListResponseInitializerJSON contains the JSON metadata for the struct
// [ProjectListResponseInitializer]
type projectListResponseInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseInitializerJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseInitializerSpec struct {
	// This field can have the runtime type of
	// [ProjectListResponseInitializerSpecsContextURLContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [ProjectListResponseInitializerSpecsGitGit].
	Git   interface{}                            `json:"git"`
	JSON  projectListResponseInitializerSpecJSON `json:"-"`
	union ProjectListResponseInitializerSpecsUnion
}

// projectListResponseInitializerSpecJSON contains the JSON metadata for the struct
// [ProjectListResponseInitializerSpec]
type projectListResponseInitializerSpecJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectListResponseInitializerSpecJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectListResponseInitializerSpec) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectListResponseInitializerSpec{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectListResponseInitializerSpecsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectListResponseInitializerSpecsContextURL],
// [ProjectListResponseInitializerSpecsGit].
func (r ProjectListResponseInitializerSpec) AsUnion() ProjectListResponseInitializerSpecsUnion {
	return r.union
}

// Union satisfied by [ProjectListResponseInitializerSpecsContextURL] or
// [ProjectListResponseInitializerSpecsGit].
type ProjectListResponseInitializerSpecsUnion interface {
	implementsProjectListResponseInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectListResponseInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectListResponseInitializerSpecsContextURL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectListResponseInitializerSpecsGit{}),
		},
	)
}

type ProjectListResponseInitializerSpecsContextURL struct {
	ContextURL ProjectListResponseInitializerSpecsContextURLContextURL `json:"contextUrl,required"`
	JSON       projectListResponseInitializerSpecsContextURLJSON       `json:"-"`
}

// projectListResponseInitializerSpecsContextURLJSON contains the JSON metadata for
// the struct [ProjectListResponseInitializerSpecsContextURL]
type projectListResponseInitializerSpecsContextURLJSON struct {
	ContextURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseInitializerSpecsContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseInitializerSpecsContextURLJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseInitializerSpecsContextURL) implementsProjectListResponseInitializerSpec() {
}

type ProjectListResponseInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                      `json:"url" format:"uri"`
	JSON projectListResponseInitializerSpecsContextURLContextURLJSON `json:"-"`
}

// projectListResponseInitializerSpecsContextURLContextURLJSON contains the JSON
// metadata for the struct
// [ProjectListResponseInitializerSpecsContextURLContextURL]
type projectListResponseInitializerSpecsContextURLContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseInitializerSpecsContextURLContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseInitializerSpecsContextURLContextURLJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseInitializerSpecsGit struct {
	Git  ProjectListResponseInitializerSpecsGitGit  `json:"git,required"`
	JSON projectListResponseInitializerSpecsGitJSON `json:"-"`
}

// projectListResponseInitializerSpecsGitJSON contains the JSON metadata for the
// struct [ProjectListResponseInitializerSpecsGit]
type projectListResponseInitializerSpecsGitJSON struct {
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseInitializerSpecsGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseInitializerSpecsGitJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseInitializerSpecsGit) implementsProjectListResponseInitializerSpec() {}

type ProjectListResponseInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode ProjectListResponseInitializerSpecsGitGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                        `json:"upstreamRemoteUri"`
	JSON              projectListResponseInitializerSpecsGitGitJSON `json:"-"`
}

// projectListResponseInitializerSpecsGitGitJSON contains the JSON metadata for the
// struct [ProjectListResponseInitializerSpecsGitGit]
type projectListResponseInitializerSpecsGitGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectListResponseInitializerSpecsGitGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseInitializerSpecsGitGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectListResponseInitializerSpecsGitGitTargetMode string

const (
	ProjectListResponseInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  ProjectListResponseInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectListResponseInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   ProjectListResponseInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectListResponseInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit ProjectListResponseInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectListResponseInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch ProjectListResponseInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectListResponseInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  ProjectListResponseInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectListResponseInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectListResponseInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, ProjectListResponseInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, ProjectListResponseInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, ProjectListResponseInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, ProjectListResponseInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectListResponseMetadata struct {
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
	Creator ProjectListResponseMetadataCreator `json:"creator"`
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
	UpdatedAt time.Time                       `json:"updatedAt" format:"date-time"`
	JSON      projectListResponseMetadataJSON `json:"-"`
}

// projectListResponseMetadataJSON contains the JSON metadata for the struct
// [ProjectListResponseMetadata]
type projectListResponseMetadataJSON struct {
	CreatedAt      apijson.Field
	Creator        apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectListResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the project creator
type ProjectListResponseMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectListResponseMetadataCreatorPrincipal `json:"principal"`
	JSON      projectListResponseMetadataCreatorJSON      `json:"-"`
}

// projectListResponseMetadataCreatorJSON contains the JSON metadata for the struct
// [ProjectListResponseMetadataCreator]
type projectListResponseMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectListResponseMetadataCreatorPrincipal string

const (
	ProjectListResponseMetadataCreatorPrincipalPrincipalUnspecified    ProjectListResponseMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectListResponseMetadataCreatorPrincipalPrincipalAccount        ProjectListResponseMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectListResponseMetadataCreatorPrincipalPrincipalUser           ProjectListResponseMetadataCreatorPrincipal = "PRINCIPAL_USER"
	ProjectListResponseMetadataCreatorPrincipalPrincipalRunner         ProjectListResponseMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	ProjectListResponseMetadataCreatorPrincipalPrincipalEnvironment    ProjectListResponseMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectListResponseMetadataCreatorPrincipalPrincipalServiceAccount ProjectListResponseMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectListResponseMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case ProjectListResponseMetadataCreatorPrincipalPrincipalUnspecified, ProjectListResponseMetadataCreatorPrincipalPrincipalAccount, ProjectListResponseMetadataCreatorPrincipalPrincipalUser, ProjectListResponseMetadataCreatorPrincipalPrincipalRunner, ProjectListResponseMetadataCreatorPrincipalPrincipalEnvironment, ProjectListResponseMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type ProjectListResponseUsedBy struct {
	// Subjects are the 10 most recent subjects who have used the project to create an
	// environment
	Subjects []ProjectListResponseUsedBySubject `json:"subjects"`
	// Total number of unique subjects who have used the project
	TotalSubjects int64                         `json:"totalSubjects"`
	JSON          projectListResponseUsedByJSON `json:"-"`
}

// projectListResponseUsedByJSON contains the JSON metadata for the struct
// [ProjectListResponseUsedBy]
type projectListResponseUsedByJSON struct {
	Subjects      apijson.Field
	TotalSubjects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectListResponseUsedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseUsedByJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseUsedBySubject struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectListResponseUsedBySubjectsPrincipal `json:"principal"`
	JSON      projectListResponseUsedBySubjectJSON       `json:"-"`
}

// projectListResponseUsedBySubjectJSON contains the JSON metadata for the struct
// [ProjectListResponseUsedBySubject]
type projectListResponseUsedBySubjectJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseUsedBySubject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseUsedBySubjectJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectListResponseUsedBySubjectsPrincipal string

const (
	ProjectListResponseUsedBySubjectsPrincipalPrincipalUnspecified    ProjectListResponseUsedBySubjectsPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectListResponseUsedBySubjectsPrincipalPrincipalAccount        ProjectListResponseUsedBySubjectsPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectListResponseUsedBySubjectsPrincipalPrincipalUser           ProjectListResponseUsedBySubjectsPrincipal = "PRINCIPAL_USER"
	ProjectListResponseUsedBySubjectsPrincipalPrincipalRunner         ProjectListResponseUsedBySubjectsPrincipal = "PRINCIPAL_RUNNER"
	ProjectListResponseUsedBySubjectsPrincipalPrincipalEnvironment    ProjectListResponseUsedBySubjectsPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectListResponseUsedBySubjectsPrincipalPrincipalServiceAccount ProjectListResponseUsedBySubjectsPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectListResponseUsedBySubjectsPrincipal) IsKnown() bool {
	switch r {
	case ProjectListResponseUsedBySubjectsPrincipalPrincipalUnspecified, ProjectListResponseUsedBySubjectsPrincipalPrincipalAccount, ProjectListResponseUsedBySubjectsPrincipalPrincipalUser, ProjectListResponseUsedBySubjectsPrincipalPrincipalRunner, ProjectListResponseUsedBySubjectsPrincipalPrincipalEnvironment, ProjectListResponseUsedBySubjectsPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type ProjectDeleteResponse = interface{}

type ProjectNewFromEnvironmentResponse struct {
	Project ProjectNewFromEnvironmentResponseProject `json:"project"`
	JSON    projectNewFromEnvironmentResponseJSON    `json:"-"`
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

type ProjectNewFromEnvironmentResponseProject struct {
	EnvironmentClass ProjectNewFromEnvironmentResponseProjectEnvironmentClass `json:"environmentClass,required"`
	// id is the unique identifier for the project
	ID string `json:"id" format:"uuid"`
	// automations_file_path is the path to the automations file relative to the repo
	// root
	AutomationsFilePath string `json:"automationsFilePath"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer ProjectNewFromEnvironmentResponseProjectInitializer `json:"initializer"`
	Metadata    ProjectNewFromEnvironmentResponseProjectMetadata    `json:"metadata"`
	UsedBy      ProjectNewFromEnvironmentResponseProjectUsedBy      `json:"usedBy"`
	JSON        projectNewFromEnvironmentResponseProjectJSON        `json:"-"`
}

// projectNewFromEnvironmentResponseProjectJSON contains the JSON metadata for the
// struct [ProjectNewFromEnvironmentResponseProject]
type projectNewFromEnvironmentResponseProjectJSON struct {
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

func (r *ProjectNewFromEnvironmentResponseProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectJSON) RawJSON() string {
	return r.raw
}

type ProjectNewFromEnvironmentResponseProjectEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                                                         `json:"localRunner"`
	JSON        projectNewFromEnvironmentResponseProjectEnvironmentClassJSON `json:"-"`
	union       ProjectNewFromEnvironmentResponseProjectEnvironmentClassUnion
}

// projectNewFromEnvironmentResponseProjectEnvironmentClassJSON contains the JSON
// metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClass]
type projectNewFromEnvironmentResponseProjectEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r projectNewFromEnvironmentResponseProjectEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectNewFromEnvironmentResponseProjectEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectNewFromEnvironmentResponseProjectEnvironmentClass{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass],
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseALocalRunnerForTheUser].
func (r ProjectNewFromEnvironmentResponseProjectEnvironmentClass) AsUnion() ProjectNewFromEnvironmentResponseProjectEnvironmentClassUnion {
	return r.union
}

// Union satisfied by
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass]
// or
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseALocalRunnerForTheUser].
type ProjectNewFromEnvironmentResponseProjectEnvironmentClassUnion interface {
	implementsProjectNewFromEnvironmentResponseProjectEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewFromEnvironmentResponseProjectEnvironmentClassUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseALocalRunnerForTheUser{}),
		},
	)
}

type ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string                                                                                                                                       `json:"environmentClassId,required" format:"uuid"`
	JSON               projectNewFromEnvironmentResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON `json:"-"`
}

// projectNewFromEnvironmentResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON
// contains the JSON metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass]
type projectNewFromEnvironmentResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) implementsProjectNewFromEnvironmentResponseProjectEnvironmentClass() {
}

type ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseALocalRunnerForTheUser struct {
	// Use a local runner for the user
	LocalRunner bool                                                                                  `json:"localRunner,required"`
	JSON        projectNewFromEnvironmentResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON `json:"-"`
}

// projectNewFromEnvironmentResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON
// contains the JSON metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseALocalRunnerForTheUser]
type projectNewFromEnvironmentResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON struct {
	LocalRunner apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseALocalRunnerForTheUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectEnvironmentClassUseALocalRunnerForTheUserJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewFromEnvironmentResponseProjectEnvironmentClassUseALocalRunnerForTheUser) implementsProjectNewFromEnvironmentResponseProjectEnvironmentClass() {
}

// EnvironmentInitializer specifies how an environment is to be initialized
type ProjectNewFromEnvironmentResponseProjectInitializer struct {
	Specs []ProjectNewFromEnvironmentResponseProjectInitializerSpec `json:"specs"`
	JSON  projectNewFromEnvironmentResponseProjectInitializerJSON   `json:"-"`
}

// projectNewFromEnvironmentResponseProjectInitializerJSON contains the JSON
// metadata for the struct [ProjectNewFromEnvironmentResponseProjectInitializer]
type projectNewFromEnvironmentResponseProjectInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectInitializerJSON) RawJSON() string {
	return r.raw
}

type ProjectNewFromEnvironmentResponseProjectInitializerSpec struct {
	// This field can have the runtime type of
	// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURLContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGit].
	Git   interface{}                                                 `json:"git"`
	JSON  projectNewFromEnvironmentResponseProjectInitializerSpecJSON `json:"-"`
	union ProjectNewFromEnvironmentResponseProjectInitializerSpecsUnion
}

// projectNewFromEnvironmentResponseProjectInitializerSpecJSON contains the JSON
// metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectInitializerSpec]
type projectNewFromEnvironmentResponseProjectInitializerSpecJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectNewFromEnvironmentResponseProjectInitializerSpecJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectNewFromEnvironmentResponseProjectInitializerSpec) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectNewFromEnvironmentResponseProjectInitializerSpec{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURL],
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsGit].
func (r ProjectNewFromEnvironmentResponseProjectInitializerSpec) AsUnion() ProjectNewFromEnvironmentResponseProjectInitializerSpecsUnion {
	return r.union
}

// Union satisfied by
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURL] or
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsGit].
type ProjectNewFromEnvironmentResponseProjectInitializerSpecsUnion interface {
	implementsProjectNewFromEnvironmentResponseProjectInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewFromEnvironmentResponseProjectInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewFromEnvironmentResponseProjectInitializerSpecsGit{}),
		},
	)
}

type ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURL struct {
	ContextURL ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURLContextURL `json:"contextUrl,required"`
	JSON       projectNewFromEnvironmentResponseProjectInitializerSpecsContextURLJSON       `json:"-"`
}

// projectNewFromEnvironmentResponseProjectInitializerSpecsContextURLJSON contains
// the JSON metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURL]
type projectNewFromEnvironmentResponseProjectInitializerSpecsContextURLJSON struct {
	ContextURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectInitializerSpecsContextURLJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURL) implementsProjectNewFromEnvironmentResponseProjectInitializerSpec() {
}

type ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                                           `json:"url" format:"uri"`
	JSON projectNewFromEnvironmentResponseProjectInitializerSpecsContextURLContextURLJSON `json:"-"`
}

// projectNewFromEnvironmentResponseProjectInitializerSpecsContextURLContextURLJSON
// contains the JSON metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURLContextURL]
type projectNewFromEnvironmentResponseProjectInitializerSpecsContextURLContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectInitializerSpecsContextURLContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectInitializerSpecsContextURLContextURLJSON) RawJSON() string {
	return r.raw
}

type ProjectNewFromEnvironmentResponseProjectInitializerSpecsGit struct {
	Git  ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGit  `json:"git,required"`
	JSON projectNewFromEnvironmentResponseProjectInitializerSpecsGitJSON `json:"-"`
}

// projectNewFromEnvironmentResponseProjectInitializerSpecsGitJSON contains the
// JSON metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsGit]
type projectNewFromEnvironmentResponseProjectInitializerSpecsGitJSON struct {
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectInitializerSpecsGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectInitializerSpecsGitJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewFromEnvironmentResponseProjectInitializerSpecsGit) implementsProjectNewFromEnvironmentResponseProjectInitializerSpec() {
}

type ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                                             `json:"upstreamRemoteUri"`
	JSON              projectNewFromEnvironmentResponseProjectInitializerSpecsGitGitJSON `json:"-"`
}

// projectNewFromEnvironmentResponseProjectInitializerSpecsGitGitJSON contains the
// JSON metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGit]
type projectNewFromEnvironmentResponseProjectInitializerSpecsGitGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectInitializerSpecsGitGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetMode string

const (
	ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, ProjectNewFromEnvironmentResponseProjectInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectNewFromEnvironmentResponseProjectMetadata struct {
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
	Creator ProjectNewFromEnvironmentResponseProjectMetadataCreator `json:"creator"`
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
	UpdatedAt time.Time                                            `json:"updatedAt" format:"date-time"`
	JSON      projectNewFromEnvironmentResponseProjectMetadataJSON `json:"-"`
}

// projectNewFromEnvironmentResponseProjectMetadataJSON contains the JSON metadata
// for the struct [ProjectNewFromEnvironmentResponseProjectMetadata]
type projectNewFromEnvironmentResponseProjectMetadataJSON struct {
	CreatedAt      apijson.Field
	Creator        apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the project creator
type ProjectNewFromEnvironmentResponseProjectMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipal `json:"principal"`
	JSON      projectNewFromEnvironmentResponseProjectMetadataCreatorJSON      `json:"-"`
}

// projectNewFromEnvironmentResponseProjectMetadataCreatorJSON contains the JSON
// metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectMetadataCreator]
type projectNewFromEnvironmentResponseProjectMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipal string

const (
	ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalUnspecified    ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalAccount        ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalUser           ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_USER"
	ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalRunner         ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalEnvironment    ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalServiceAccount ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalUnspecified, ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalAccount, ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalUser, ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalRunner, ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalEnvironment, ProjectNewFromEnvironmentResponseProjectMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type ProjectNewFromEnvironmentResponseProjectUsedBy struct {
	// Subjects are the 10 most recent subjects who have used the project to create an
	// environment
	Subjects []ProjectNewFromEnvironmentResponseProjectUsedBySubject `json:"subjects"`
	// Total number of unique subjects who have used the project
	TotalSubjects int64                                              `json:"totalSubjects"`
	JSON          projectNewFromEnvironmentResponseProjectUsedByJSON `json:"-"`
}

// projectNewFromEnvironmentResponseProjectUsedByJSON contains the JSON metadata
// for the struct [ProjectNewFromEnvironmentResponseProjectUsedBy]
type projectNewFromEnvironmentResponseProjectUsedByJSON struct {
	Subjects      apijson.Field
	TotalSubjects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectUsedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectUsedByJSON) RawJSON() string {
	return r.raw
}

type ProjectNewFromEnvironmentResponseProjectUsedBySubject struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipal `json:"principal"`
	JSON      projectNewFromEnvironmentResponseProjectUsedBySubjectJSON       `json:"-"`
}

// projectNewFromEnvironmentResponseProjectUsedBySubjectJSON contains the JSON
// metadata for the struct [ProjectNewFromEnvironmentResponseProjectUsedBySubject]
type projectNewFromEnvironmentResponseProjectUsedBySubjectJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectUsedBySubject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectUsedBySubjectJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipal string

const (
	ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalUnspecified    ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalAccount        ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalUser           ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_USER"
	ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalRunner         ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_RUNNER"
	ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalEnvironment    ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalServiceAccount ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipal) IsKnown() bool {
	switch r {
	case ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalUnspecified, ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalAccount, ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalUser, ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalRunner, ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalEnvironment, ProjectNewFromEnvironmentResponseProjectUsedBySubjectsPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type ProjectNewParams struct {
	EnvironmentClass param.Field[ProjectNewParamsEnvironmentClassUnion] `json:"environmentClass,required"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer param.Field[ProjectNewParamsInitializer] `json:"initializer,required"`
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

type ProjectNewParamsEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID param.Field[string] `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner param.Field[bool] `json:"localRunner"`
}

func (r ProjectNewParamsEnvironmentClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsEnvironmentClass) implementsProjectNewParamsEnvironmentClassUnion() {}

// Satisfied by
// [ProjectNewParamsEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass],
// [ProjectNewParamsEnvironmentClassUseALocalRunnerForTheUser],
// [ProjectNewParamsEnvironmentClass].
type ProjectNewParamsEnvironmentClassUnion interface {
	implementsProjectNewParamsEnvironmentClassUnion()
}

type ProjectNewParamsEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID param.Field[string] `json:"environmentClassId,required" format:"uuid"`
}

func (r ProjectNewParamsEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) implementsProjectNewParamsEnvironmentClassUnion() {
}

type ProjectNewParamsEnvironmentClassUseALocalRunnerForTheUser struct {
	// Use a local runner for the user
	LocalRunner param.Field[bool] `json:"localRunner,required"`
}

func (r ProjectNewParamsEnvironmentClassUseALocalRunnerForTheUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsEnvironmentClassUseALocalRunnerForTheUser) implementsProjectNewParamsEnvironmentClassUnion() {
}

// EnvironmentInitializer specifies how an environment is to be initialized
type ProjectNewParamsInitializer struct {
	Specs param.Field[[]ProjectNewParamsInitializerSpecUnion] `json:"specs"`
}

func (r ProjectNewParamsInitializer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectNewParamsInitializerSpec struct {
	ContextURL param.Field[interface{}] `json:"contextUrl"`
	Git        param.Field[interface{}] `json:"git"`
}

func (r ProjectNewParamsInitializerSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsInitializerSpec) implementsProjectNewParamsInitializerSpecUnion() {}

// Satisfied by [ProjectNewParamsInitializerSpecsContextURL],
// [ProjectNewParamsInitializerSpecsGit], [ProjectNewParamsInitializerSpec].
type ProjectNewParamsInitializerSpecUnion interface {
	implementsProjectNewParamsInitializerSpecUnion()
}

type ProjectNewParamsInitializerSpecsContextURL struct {
	ContextURL param.Field[ProjectNewParamsInitializerSpecsContextURLContextURL] `json:"contextUrl,required"`
}

func (r ProjectNewParamsInitializerSpecsContextURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsInitializerSpecsContextURL) implementsProjectNewParamsInitializerSpecUnion() {
}

type ProjectNewParamsInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r ProjectNewParamsInitializerSpecsContextURLContextURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectNewParamsInitializerSpecsGit struct {
	Git param.Field[ProjectNewParamsInitializerSpecsGitGit] `json:"git,required"`
}

func (r ProjectNewParamsInitializerSpecsGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsInitializerSpecsGit) implementsProjectNewParamsInitializerSpecUnion() {}

type ProjectNewParamsInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation param.Field[string] `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget param.Field[string] `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri param.Field[string] `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode param.Field[ProjectNewParamsInitializerSpecsGitGitTargetMode] `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri param.Field[string] `json:"upstreamRemoteUri"`
}

func (r ProjectNewParamsInitializerSpecsGitGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectNewParamsInitializerSpecsGitGitTargetMode string

const (
	ProjectNewParamsInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  ProjectNewParamsInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectNewParamsInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   ProjectNewParamsInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectNewParamsInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit ProjectNewParamsInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectNewParamsInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch ProjectNewParamsInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectNewParamsInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  ProjectNewParamsInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectNewParamsInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectNewParamsInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, ProjectNewParamsInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, ProjectNewParamsInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, ProjectNewParamsInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, ProjectNewParamsInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectGetParams struct {
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
}

func (r ProjectGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectUpdateParams struct {
	Body ProjectUpdateParamsBodyUnion `json:"body,required"`
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectUpdateParamsBody struct {
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
	DevcontainerFilePath param.Field[string]      `json:"devcontainerFilePath"`
	EnvironmentClass     param.Field[interface{}] `json:"environmentClass"`
	Initializer          param.Field[interface{}] `json:"initializer"`
	Name                 param.Field[string]      `json:"name"`
}

func (r ProjectUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBody) implementsProjectUpdateParamsBodyUnion() {}

// Satisfied by
// [ProjectUpdateParamsBodyAutomationsFilePathIsThePathToTheAutomationsFileRelativeToTheRepoRoot],
// [ProjectUpdateParamsBodyDevcontainerFilePathIsThePathToTheDevcontainerFileRelativeToTheRepoRoot],
// [ProjectUpdateParamsBodyEnvironmentClass],
// [ProjectUpdateParamsBodyInitializerIsTheContentInitializer],
// [ProjectUpdateParamsBodyName], [ProjectUpdateParamsBody].
type ProjectUpdateParamsBodyUnion interface {
	implementsProjectUpdateParamsBodyUnion()
}

type ProjectUpdateParamsBodyAutomationsFilePathIsThePathToTheAutomationsFileRelativeToTheRepoRoot struct {
	// automations_file_path is the path to the automations file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath,required"`
}

func (r ProjectUpdateParamsBodyAutomationsFilePathIsThePathToTheAutomationsFileRelativeToTheRepoRoot) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBodyAutomationsFilePathIsThePathToTheAutomationsFileRelativeToTheRepoRoot) implementsProjectUpdateParamsBodyUnion() {
}

type ProjectUpdateParamsBodyDevcontainerFilePathIsThePathToTheDevcontainerFileRelativeToTheRepoRoot struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath,required"`
}

func (r ProjectUpdateParamsBodyDevcontainerFilePathIsThePathToTheDevcontainerFileRelativeToTheRepoRoot) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBodyDevcontainerFilePathIsThePathToTheDevcontainerFileRelativeToTheRepoRoot) implementsProjectUpdateParamsBodyUnion() {
}

type ProjectUpdateParamsBodyEnvironmentClass struct {
	EnvironmentClass param.Field[ProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUnion] `json:"environmentClass,required"`
}

func (r ProjectUpdateParamsBodyEnvironmentClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBodyEnvironmentClass) implementsProjectUpdateParamsBodyUnion() {}

type ProjectUpdateParamsBodyEnvironmentClassEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID param.Field[string] `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner param.Field[bool] `json:"localRunner"`
}

func (r ProjectUpdateParamsBodyEnvironmentClassEnvironmentClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBodyEnvironmentClassEnvironmentClass) implementsProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUnion() {
}

// Satisfied by
// [ProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass],
// [ProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUseALocalRunnerForTheUser],
// [ProjectUpdateParamsBodyEnvironmentClassEnvironmentClass].
type ProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUnion interface {
	implementsProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUnion()
}

type ProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID param.Field[string] `json:"environmentClassId,required" format:"uuid"`
}

func (r ProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUseAFixedEnvironmentClassOnAGivenRunnerThisCannotBeALocalRunnerSEnvironmentClass) implementsProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUnion() {
}

type ProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUseALocalRunnerForTheUser struct {
	// Use a local runner for the user
	LocalRunner param.Field[bool] `json:"localRunner,required"`
}

func (r ProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUseALocalRunnerForTheUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUseALocalRunnerForTheUser) implementsProjectUpdateParamsBodyEnvironmentClassEnvironmentClassUnion() {
}

type ProjectUpdateParamsBodyInitializerIsTheContentInitializer struct {
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer param.Field[ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializer] `json:"initializer,required"`
}

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializer) implementsProjectUpdateParamsBodyUnion() {
}

// EnvironmentInitializer specifies how an environment is to be initialized
type ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializer struct {
	Specs param.Field[[]ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecUnion] `json:"specs"`
}

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpec struct {
	ContextURL param.Field[interface{}] `json:"contextUrl"`
	Git        param.Field[interface{}] `json:"git"`
}

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpec) implementsProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecUnion() {
}

// Satisfied by
// [ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsContextURL],
// [ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGit],
// [ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpec].
type ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecUnion interface {
	implementsProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecUnion()
}

type ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsContextURL struct {
	ContextURL param.Field[ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsContextURLContextURL] `json:"contextUrl,required"`
}

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsContextURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsContextURL) implementsProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecUnion() {
}

type ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsContextURLContextURL struct {
	// url is the URL from which the environment is created
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsContextURLContextURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGit struct {
	Git param.Field[ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGit] `json:"git,required"`
}

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGit) implementsProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecUnion() {
}

type ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGit struct {
	// a path relative to the environment root in which the code will be checked out to
	CheckoutLocation param.Field[string] `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget param.Field[string] `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri param.Field[string] `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode param.Field[ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetMode] `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri param.Field[string] `json:"upstreamRemoteUri"`
}

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetMode string

const (
	ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified  ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead   ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch  ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetModeCloneTargetModeUnspecified, ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetModeCloneTargetModeRemoteHead, ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetModeCloneTargetModeRemoteCommit, ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetModeCloneTargetModeRemoteBranch, ProjectUpdateParamsBodyInitializerIsTheContentInitializerInitializerSpecsGitGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectUpdateParamsBodyName struct {
	Name param.Field[string] `json:"name,required"`
}

func (r ProjectUpdateParamsBodyName) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParamsBodyName) implementsProjectUpdateParamsBodyUnion() {}

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
