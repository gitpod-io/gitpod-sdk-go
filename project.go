// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
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
func (r *ProjectService) New(ctx context.Context, params ProjectNewParams, opts ...option.RequestOption) (res *ProjectNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/CreateProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetProject retrieves a single Project.
func (r *ProjectService) Get(ctx context.Context, params ProjectGetParams, opts ...option.RequestOption) (res *ProjectGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/GetProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// UpdateProject updates the properties of a Project.
func (r *ProjectService) Update(ctx context.Context, params ProjectUpdateParams, opts ...option.RequestOption) (res *ProjectUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/UpdateProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListProjects lists all projects the caller has access to.
func (r *ProjectService) List(ctx context.Context, params ProjectListParams, opts ...option.RequestOption) (res *ProjectListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/ListProjects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// DeleteProject deletes the specified project.
func (r *ProjectService) Delete(ctx context.Context, params ProjectDeleteParams, opts ...option.RequestOption) (res *ProjectDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/DeleteProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// CreateProject creates a new Project using an environment as template.
func (r *ProjectService) NewFromEnvironment(ctx context.Context, params ProjectNewFromEnvironmentParams, opts ...option.RequestOption) (res *ProjectNewFromEnvironmentResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.ProjectService/CreateProjectFromEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
// [ProjectNewResponseProjectEnvironmentClassObject],
// [ProjectNewResponseProjectEnvironmentClassObject],
// [ProjectNewResponseProjectEnvironmentClassObject].
func (r ProjectNewResponseProjectEnvironmentClass) AsUnion() ProjectNewResponseProjectEnvironmentClassUnion {
	return r.union
}

// Union satisfied by [ProjectNewResponseProjectEnvironmentClassObject],
// [ProjectNewResponseProjectEnvironmentClassObject] or
// [ProjectNewResponseProjectEnvironmentClassObject].
type ProjectNewResponseProjectEnvironmentClassUnion interface {
	implementsProjectNewResponseProjectEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewResponseProjectEnvironmentClassUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewResponseProjectEnvironmentClassObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewResponseProjectEnvironmentClassObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewResponseProjectEnvironmentClassObject{}),
		},
	)
}

type ProjectNewResponseProjectEnvironmentClassObject struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId,required" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                                                `json:"localRunner"`
	JSON        projectNewResponseProjectEnvironmentClassObjectJSON `json:"-"`
}

// projectNewResponseProjectEnvironmentClassObjectJSON contains the JSON metadata
// for the struct [ProjectNewResponseProjectEnvironmentClassObject]
type projectNewResponseProjectEnvironmentClassObjectJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectNewResponseProjectEnvironmentClassObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectEnvironmentClassObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseProjectEnvironmentClassObject) implementsProjectNewResponseProjectEnvironmentClass() {
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
	// [ProjectNewResponseProjectInitializerSpecsObjectContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [ProjectNewResponseProjectInitializerSpecsObjectGit].
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
// [ProjectNewResponseProjectInitializerSpecsObject],
// [ProjectNewResponseProjectInitializerSpecsObject],
// [ProjectNewResponseProjectInitializerSpecsObject].
func (r ProjectNewResponseProjectInitializerSpec) AsUnion() ProjectNewResponseProjectInitializerSpecsUnion {
	return r.union
}

// Union satisfied by [ProjectNewResponseProjectInitializerSpecsObject],
// [ProjectNewResponseProjectInitializerSpecsObject] or
// [ProjectNewResponseProjectInitializerSpecsObject].
type ProjectNewResponseProjectInitializerSpecsUnion interface {
	implementsProjectNewResponseProjectInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewResponseProjectInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewResponseProjectInitializerSpecsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewResponseProjectInitializerSpecsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewResponseProjectInitializerSpecsObject{}),
		},
	)
}

type ProjectNewResponseProjectInitializerSpecsObject struct {
	ContextURL ProjectNewResponseProjectInitializerSpecsObjectContextURL `json:"contextUrl,required"`
	Git        ProjectNewResponseProjectInitializerSpecsObjectGit        `json:"git"`
	JSON       projectNewResponseProjectInitializerSpecsObjectJSON       `json:"-"`
}

// projectNewResponseProjectInitializerSpecsObjectJSON contains the JSON metadata
// for the struct [ProjectNewResponseProjectInitializerSpecsObject]
type projectNewResponseProjectInitializerSpecsObjectJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseProjectInitializerSpecsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectInitializerSpecsObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseProjectInitializerSpecsObject) implementsProjectNewResponseProjectInitializerSpec() {
}

type ProjectNewResponseProjectInitializerSpecsObjectContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                        `json:"url" format:"uri"`
	JSON projectNewResponseProjectInitializerSpecsObjectContextURLJSON `json:"-"`
}

// projectNewResponseProjectInitializerSpecsObjectContextURLJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseProjectInitializerSpecsObjectContextURL]
type projectNewResponseProjectInitializerSpecsObjectContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseProjectInitializerSpecsObjectContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectInitializerSpecsObjectContextURLJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseProjectInitializerSpecsObjectGit struct {
	// a path relative to the environment root in which the code will be checked out
	//
	// to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode ProjectNewResponseProjectInitializerSpecsObjectGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                                 `json:"upstreamRemoteUri"`
	JSON              projectNewResponseProjectInitializerSpecsObjectGitJSON `json:"-"`
}

// projectNewResponseProjectInitializerSpecsObjectGitJSON contains the JSON
// metadata for the struct [ProjectNewResponseProjectInitializerSpecsObjectGit]
type projectNewResponseProjectInitializerSpecsObjectGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectNewResponseProjectInitializerSpecsObjectGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseProjectInitializerSpecsObjectGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectNewResponseProjectInitializerSpecsObjectGitTargetMode string

const (
	ProjectNewResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified  ProjectNewResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectNewResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead   ProjectNewResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectNewResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit ProjectNewResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectNewResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch ProjectNewResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectNewResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch  ProjectNewResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectNewResponseProjectInitializerSpecsObjectGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectNewResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified, ProjectNewResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead, ProjectNewResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit, ProjectNewResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch, ProjectNewResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectNewResponseProjectMetadata struct {
	// A Timestamp represents a point in time independent of any time zone or local
	//
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
	//
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
// [ProjectGetResponseProjectEnvironmentClassObject],
// [ProjectGetResponseProjectEnvironmentClassObject],
// [ProjectGetResponseProjectEnvironmentClassObject].
func (r ProjectGetResponseProjectEnvironmentClass) AsUnion() ProjectGetResponseProjectEnvironmentClassUnion {
	return r.union
}

// Union satisfied by [ProjectGetResponseProjectEnvironmentClassObject],
// [ProjectGetResponseProjectEnvironmentClassObject] or
// [ProjectGetResponseProjectEnvironmentClassObject].
type ProjectGetResponseProjectEnvironmentClassUnion interface {
	implementsProjectGetResponseProjectEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectGetResponseProjectEnvironmentClassUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectGetResponseProjectEnvironmentClassObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectGetResponseProjectEnvironmentClassObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectGetResponseProjectEnvironmentClassObject{}),
		},
	)
}

type ProjectGetResponseProjectEnvironmentClassObject struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId,required" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                                                `json:"localRunner"`
	JSON        projectGetResponseProjectEnvironmentClassObjectJSON `json:"-"`
}

// projectGetResponseProjectEnvironmentClassObjectJSON contains the JSON metadata
// for the struct [ProjectGetResponseProjectEnvironmentClassObject]
type projectGetResponseProjectEnvironmentClassObjectJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectGetResponseProjectEnvironmentClassObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectEnvironmentClassObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseProjectEnvironmentClassObject) implementsProjectGetResponseProjectEnvironmentClass() {
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
	// [ProjectGetResponseProjectInitializerSpecsObjectContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [ProjectGetResponseProjectInitializerSpecsObjectGit].
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
// [ProjectGetResponseProjectInitializerSpecsObject],
// [ProjectGetResponseProjectInitializerSpecsObject],
// [ProjectGetResponseProjectInitializerSpecsObject].
func (r ProjectGetResponseProjectInitializerSpec) AsUnion() ProjectGetResponseProjectInitializerSpecsUnion {
	return r.union
}

// Union satisfied by [ProjectGetResponseProjectInitializerSpecsObject],
// [ProjectGetResponseProjectInitializerSpecsObject] or
// [ProjectGetResponseProjectInitializerSpecsObject].
type ProjectGetResponseProjectInitializerSpecsUnion interface {
	implementsProjectGetResponseProjectInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectGetResponseProjectInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectGetResponseProjectInitializerSpecsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectGetResponseProjectInitializerSpecsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectGetResponseProjectInitializerSpecsObject{}),
		},
	)
}

type ProjectGetResponseProjectInitializerSpecsObject struct {
	ContextURL ProjectGetResponseProjectInitializerSpecsObjectContextURL `json:"contextUrl,required"`
	Git        ProjectGetResponseProjectInitializerSpecsObjectGit        `json:"git"`
	JSON       projectGetResponseProjectInitializerSpecsObjectJSON       `json:"-"`
}

// projectGetResponseProjectInitializerSpecsObjectJSON contains the JSON metadata
// for the struct [ProjectGetResponseProjectInitializerSpecsObject]
type projectGetResponseProjectInitializerSpecsObjectJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseProjectInitializerSpecsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectInitializerSpecsObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseProjectInitializerSpecsObject) implementsProjectGetResponseProjectInitializerSpec() {
}

type ProjectGetResponseProjectInitializerSpecsObjectContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                        `json:"url" format:"uri"`
	JSON projectGetResponseProjectInitializerSpecsObjectContextURLJSON `json:"-"`
}

// projectGetResponseProjectInitializerSpecsObjectContextURLJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseProjectInitializerSpecsObjectContextURL]
type projectGetResponseProjectInitializerSpecsObjectContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseProjectInitializerSpecsObjectContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectInitializerSpecsObjectContextURLJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseProjectInitializerSpecsObjectGit struct {
	// a path relative to the environment root in which the code will be checked out
	//
	// to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode ProjectGetResponseProjectInitializerSpecsObjectGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                                 `json:"upstreamRemoteUri"`
	JSON              projectGetResponseProjectInitializerSpecsObjectGitJSON `json:"-"`
}

// projectGetResponseProjectInitializerSpecsObjectGitJSON contains the JSON
// metadata for the struct [ProjectGetResponseProjectInitializerSpecsObjectGit]
type projectGetResponseProjectInitializerSpecsObjectGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectGetResponseProjectInitializerSpecsObjectGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseProjectInitializerSpecsObjectGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectGetResponseProjectInitializerSpecsObjectGitTargetMode string

const (
	ProjectGetResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified  ProjectGetResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectGetResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead   ProjectGetResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectGetResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit ProjectGetResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectGetResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch ProjectGetResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectGetResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch  ProjectGetResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectGetResponseProjectInitializerSpecsObjectGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectGetResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified, ProjectGetResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead, ProjectGetResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit, ProjectGetResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch, ProjectGetResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectGetResponseProjectMetadata struct {
	// A Timestamp represents a point in time independent of any time zone or local
	//
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
	//
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
// [ProjectUpdateResponseProjectEnvironmentClassObject],
// [ProjectUpdateResponseProjectEnvironmentClassObject],
// [ProjectUpdateResponseProjectEnvironmentClassObject].
func (r ProjectUpdateResponseProjectEnvironmentClass) AsUnion() ProjectUpdateResponseProjectEnvironmentClassUnion {
	return r.union
}

// Union satisfied by [ProjectUpdateResponseProjectEnvironmentClassObject],
// [ProjectUpdateResponseProjectEnvironmentClassObject] or
// [ProjectUpdateResponseProjectEnvironmentClassObject].
type ProjectUpdateResponseProjectEnvironmentClassUnion interface {
	implementsProjectUpdateResponseProjectEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectUpdateResponseProjectEnvironmentClassUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectUpdateResponseProjectEnvironmentClassObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectUpdateResponseProjectEnvironmentClassObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectUpdateResponseProjectEnvironmentClassObject{}),
		},
	)
}

type ProjectUpdateResponseProjectEnvironmentClassObject struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId,required" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                                                   `json:"localRunner"`
	JSON        projectUpdateResponseProjectEnvironmentClassObjectJSON `json:"-"`
}

// projectUpdateResponseProjectEnvironmentClassObjectJSON contains the JSON
// metadata for the struct [ProjectUpdateResponseProjectEnvironmentClassObject]
type projectUpdateResponseProjectEnvironmentClassObjectJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectEnvironmentClassObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectEnvironmentClassObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectUpdateResponseProjectEnvironmentClassObject) implementsProjectUpdateResponseProjectEnvironmentClass() {
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
	// [ProjectUpdateResponseProjectInitializerSpecsObjectContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [ProjectUpdateResponseProjectInitializerSpecsObjectGit].
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
// [ProjectUpdateResponseProjectInitializerSpecsObject],
// [ProjectUpdateResponseProjectInitializerSpecsObject],
// [ProjectUpdateResponseProjectInitializerSpecsObject].
func (r ProjectUpdateResponseProjectInitializerSpec) AsUnion() ProjectUpdateResponseProjectInitializerSpecsUnion {
	return r.union
}

// Union satisfied by [ProjectUpdateResponseProjectInitializerSpecsObject],
// [ProjectUpdateResponseProjectInitializerSpecsObject] or
// [ProjectUpdateResponseProjectInitializerSpecsObject].
type ProjectUpdateResponseProjectInitializerSpecsUnion interface {
	implementsProjectUpdateResponseProjectInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectUpdateResponseProjectInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectUpdateResponseProjectInitializerSpecsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectUpdateResponseProjectInitializerSpecsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectUpdateResponseProjectInitializerSpecsObject{}),
		},
	)
}

type ProjectUpdateResponseProjectInitializerSpecsObject struct {
	ContextURL ProjectUpdateResponseProjectInitializerSpecsObjectContextURL `json:"contextUrl,required"`
	Git        ProjectUpdateResponseProjectInitializerSpecsObjectGit        `json:"git"`
	JSON       projectUpdateResponseProjectInitializerSpecsObjectJSON       `json:"-"`
}

// projectUpdateResponseProjectInitializerSpecsObjectJSON contains the JSON
// metadata for the struct [ProjectUpdateResponseProjectInitializerSpecsObject]
type projectUpdateResponseProjectInitializerSpecsObjectJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectInitializerSpecsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectInitializerSpecsObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectUpdateResponseProjectInitializerSpecsObject) implementsProjectUpdateResponseProjectInitializerSpec() {
}

type ProjectUpdateResponseProjectInitializerSpecsObjectContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                           `json:"url" format:"uri"`
	JSON projectUpdateResponseProjectInitializerSpecsObjectContextURLJSON `json:"-"`
}

// projectUpdateResponseProjectInitializerSpecsObjectContextURLJSON contains the
// JSON metadata for the struct
// [ProjectUpdateResponseProjectInitializerSpecsObjectContextURL]
type projectUpdateResponseProjectInitializerSpecsObjectContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectInitializerSpecsObjectContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectInitializerSpecsObjectContextURLJSON) RawJSON() string {
	return r.raw
}

type ProjectUpdateResponseProjectInitializerSpecsObjectGit struct {
	// a path relative to the environment root in which the code will be checked out
	//
	// to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                                    `json:"upstreamRemoteUri"`
	JSON              projectUpdateResponseProjectInitializerSpecsObjectGitJSON `json:"-"`
}

// projectUpdateResponseProjectInitializerSpecsObjectGitJSON contains the JSON
// metadata for the struct [ProjectUpdateResponseProjectInitializerSpecsObjectGit]
type projectUpdateResponseProjectInitializerSpecsObjectGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectUpdateResponseProjectInitializerSpecsObjectGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectUpdateResponseProjectInitializerSpecsObjectGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetMode string

const (
	ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified  ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead   ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch  ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified, ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead, ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit, ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch, ProjectUpdateResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectUpdateResponseProjectMetadata struct {
	// A Timestamp represents a point in time independent of any time zone or local
	//
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
	//
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
	// pagination contains the pagination options for listing organizations
	Pagination ProjectListResponsePagination `json:"pagination"`
	Projects   []ProjectListResponseProject  `json:"projects"`
	JSON       projectListResponseJSON       `json:"-"`
}

// projectListResponseJSON contains the JSON metadata for the struct
// [ProjectListResponse]
type projectListResponseJSON struct {
	Pagination  apijson.Field
	Projects    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing organizations
type ProjectListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                            `json:"nextToken"`
	JSON      projectListResponsePaginationJSON `json:"-"`
}

// projectListResponsePaginationJSON contains the JSON metadata for the struct
// [ProjectListResponsePagination]
type projectListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseProject struct {
	EnvironmentClass ProjectListResponseProjectsEnvironmentClass `json:"environmentClass,required"`
	// id is the unique identifier for the project
	ID string `json:"id" format:"uuid"`
	// automations_file_path is the path to the automations file relative to the repo
	// root
	AutomationsFilePath string `json:"automationsFilePath"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer ProjectListResponseProjectsInitializer `json:"initializer"`
	Metadata    ProjectListResponseProjectsMetadata    `json:"metadata"`
	UsedBy      ProjectListResponseProjectsUsedBy      `json:"usedBy"`
	JSON        projectListResponseProjectJSON         `json:"-"`
}

// projectListResponseProjectJSON contains the JSON metadata for the struct
// [ProjectListResponseProject]
type projectListResponseProjectJSON struct {
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

func (r *ProjectListResponseProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseProjectJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseProjectsEnvironmentClass struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                                            `json:"localRunner"`
	JSON        projectListResponseProjectsEnvironmentClassJSON `json:"-"`
	union       ProjectListResponseProjectsEnvironmentClassUnion
}

// projectListResponseProjectsEnvironmentClassJSON contains the JSON metadata for
// the struct [ProjectListResponseProjectsEnvironmentClass]
type projectListResponseProjectsEnvironmentClassJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r projectListResponseProjectsEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectListResponseProjectsEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectListResponseProjectsEnvironmentClass{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectListResponseProjectsEnvironmentClassUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectListResponseProjectsEnvironmentClassObject],
// [ProjectListResponseProjectsEnvironmentClassObject],
// [ProjectListResponseProjectsEnvironmentClassObject].
func (r ProjectListResponseProjectsEnvironmentClass) AsUnion() ProjectListResponseProjectsEnvironmentClassUnion {
	return r.union
}

// Union satisfied by [ProjectListResponseProjectsEnvironmentClassObject],
// [ProjectListResponseProjectsEnvironmentClassObject] or
// [ProjectListResponseProjectsEnvironmentClassObject].
type ProjectListResponseProjectsEnvironmentClassUnion interface {
	implementsProjectListResponseProjectsEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectListResponseProjectsEnvironmentClassUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectListResponseProjectsEnvironmentClassObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectListResponseProjectsEnvironmentClassObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectListResponseProjectsEnvironmentClassObject{}),
		},
	)
}

type ProjectListResponseProjectsEnvironmentClassObject struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId,required" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                                                  `json:"localRunner"`
	JSON        projectListResponseProjectsEnvironmentClassObjectJSON `json:"-"`
}

// projectListResponseProjectsEnvironmentClassObjectJSON contains the JSON metadata
// for the struct [ProjectListResponseProjectsEnvironmentClassObject]
type projectListResponseProjectsEnvironmentClassObjectJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectListResponseProjectsEnvironmentClassObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseProjectsEnvironmentClassObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseProjectsEnvironmentClassObject) implementsProjectListResponseProjectsEnvironmentClass() {
}

// EnvironmentInitializer specifies how an environment is to be initialized
type ProjectListResponseProjectsInitializer struct {
	Specs []ProjectListResponseProjectsInitializerSpec `json:"specs"`
	JSON  projectListResponseProjectsInitializerJSON   `json:"-"`
}

// projectListResponseProjectsInitializerJSON contains the JSON metadata for the
// struct [ProjectListResponseProjectsInitializer]
type projectListResponseProjectsInitializerJSON struct {
	Specs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseProjectsInitializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseProjectsInitializerJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseProjectsInitializerSpec struct {
	// This field can have the runtime type of
	// [ProjectListResponseProjectsInitializerSpecsObjectContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [ProjectListResponseProjectsInitializerSpecsObjectGit].
	Git   interface{}                                    `json:"git"`
	JSON  projectListResponseProjectsInitializerSpecJSON `json:"-"`
	union ProjectListResponseProjectsInitializerSpecsUnion
}

// projectListResponseProjectsInitializerSpecJSON contains the JSON metadata for
// the struct [ProjectListResponseProjectsInitializerSpec]
type projectListResponseProjectsInitializerSpecJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectListResponseProjectsInitializerSpecJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectListResponseProjectsInitializerSpec) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectListResponseProjectsInitializerSpec{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectListResponseProjectsInitializerSpecsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectListResponseProjectsInitializerSpecsObject],
// [ProjectListResponseProjectsInitializerSpecsObject],
// [ProjectListResponseProjectsInitializerSpecsObject].
func (r ProjectListResponseProjectsInitializerSpec) AsUnion() ProjectListResponseProjectsInitializerSpecsUnion {
	return r.union
}

// Union satisfied by [ProjectListResponseProjectsInitializerSpecsObject],
// [ProjectListResponseProjectsInitializerSpecsObject] or
// [ProjectListResponseProjectsInitializerSpecsObject].
type ProjectListResponseProjectsInitializerSpecsUnion interface {
	implementsProjectListResponseProjectsInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectListResponseProjectsInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectListResponseProjectsInitializerSpecsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectListResponseProjectsInitializerSpecsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectListResponseProjectsInitializerSpecsObject{}),
		},
	)
}

type ProjectListResponseProjectsInitializerSpecsObject struct {
	ContextURL ProjectListResponseProjectsInitializerSpecsObjectContextURL `json:"contextUrl,required"`
	Git        ProjectListResponseProjectsInitializerSpecsObjectGit        `json:"git"`
	JSON       projectListResponseProjectsInitializerSpecsObjectJSON       `json:"-"`
}

// projectListResponseProjectsInitializerSpecsObjectJSON contains the JSON metadata
// for the struct [ProjectListResponseProjectsInitializerSpecsObject]
type projectListResponseProjectsInitializerSpecsObjectJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseProjectsInitializerSpecsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseProjectsInitializerSpecsObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseProjectsInitializerSpecsObject) implementsProjectListResponseProjectsInitializerSpec() {
}

type ProjectListResponseProjectsInitializerSpecsObjectContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                          `json:"url" format:"uri"`
	JSON projectListResponseProjectsInitializerSpecsObjectContextURLJSON `json:"-"`
}

// projectListResponseProjectsInitializerSpecsObjectContextURLJSON contains the
// JSON metadata for the struct
// [ProjectListResponseProjectsInitializerSpecsObjectContextURL]
type projectListResponseProjectsInitializerSpecsObjectContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseProjectsInitializerSpecsObjectContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseProjectsInitializerSpecsObjectContextURLJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseProjectsInitializerSpecsObjectGit struct {
	// a path relative to the environment root in which the code will be checked out
	//
	// to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode ProjectListResponseProjectsInitializerSpecsObjectGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                                   `json:"upstreamRemoteUri"`
	JSON              projectListResponseProjectsInitializerSpecsObjectGitJSON `json:"-"`
}

// projectListResponseProjectsInitializerSpecsObjectGitJSON contains the JSON
// metadata for the struct [ProjectListResponseProjectsInitializerSpecsObjectGit]
type projectListResponseProjectsInitializerSpecsObjectGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectListResponseProjectsInitializerSpecsObjectGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseProjectsInitializerSpecsObjectGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectListResponseProjectsInitializerSpecsObjectGitTargetMode string

const (
	ProjectListResponseProjectsInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified  ProjectListResponseProjectsInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectListResponseProjectsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead   ProjectListResponseProjectsInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectListResponseProjectsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit ProjectListResponseProjectsInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectListResponseProjectsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch ProjectListResponseProjectsInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectListResponseProjectsInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch  ProjectListResponseProjectsInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectListResponseProjectsInitializerSpecsObjectGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectListResponseProjectsInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified, ProjectListResponseProjectsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead, ProjectListResponseProjectsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit, ProjectListResponseProjectsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch, ProjectListResponseProjectsInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectListResponseProjectsMetadata struct {
	// A Timestamp represents a point in time independent of any time zone or local
	//
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
	Creator ProjectListResponseProjectsMetadataCreator `json:"creator"`
	// name is the human readable name of the project
	Name string `json:"name"`
	// organization_id is the ID of the organization that contains the environment
	OrganizationID string `json:"organizationId" format:"uuid"`
	// A Timestamp represents a point in time independent of any time zone or local
	//
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
	UpdatedAt time.Time                               `json:"updatedAt" format:"date-time"`
	JSON      projectListResponseProjectsMetadataJSON `json:"-"`
}

// projectListResponseProjectsMetadataJSON contains the JSON metadata for the
// struct [ProjectListResponseProjectsMetadata]
type projectListResponseProjectsMetadataJSON struct {
	CreatedAt      apijson.Field
	Creator        apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectListResponseProjectsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseProjectsMetadataJSON) RawJSON() string {
	return r.raw
}

// creator is the identity of the project creator
type ProjectListResponseProjectsMetadataCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectListResponseProjectsMetadataCreatorPrincipal `json:"principal"`
	JSON      projectListResponseProjectsMetadataCreatorJSON      `json:"-"`
}

// projectListResponseProjectsMetadataCreatorJSON contains the JSON metadata for
// the struct [ProjectListResponseProjectsMetadataCreator]
type projectListResponseProjectsMetadataCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseProjectsMetadataCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseProjectsMetadataCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectListResponseProjectsMetadataCreatorPrincipal string

const (
	ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalUnspecified    ProjectListResponseProjectsMetadataCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalAccount        ProjectListResponseProjectsMetadataCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalUser           ProjectListResponseProjectsMetadataCreatorPrincipal = "PRINCIPAL_USER"
	ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalRunner         ProjectListResponseProjectsMetadataCreatorPrincipal = "PRINCIPAL_RUNNER"
	ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalEnvironment    ProjectListResponseProjectsMetadataCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalServiceAccount ProjectListResponseProjectsMetadataCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectListResponseProjectsMetadataCreatorPrincipal) IsKnown() bool {
	switch r {
	case ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalUnspecified, ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalAccount, ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalUser, ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalRunner, ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalEnvironment, ProjectListResponseProjectsMetadataCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type ProjectListResponseProjectsUsedBy struct {
	// Subjects are the 10 most recent subjects who have used the project to create an
	// environment
	Subjects []ProjectListResponseProjectsUsedBySubject `json:"subjects"`
	// Total number of unique subjects who have used the project
	TotalSubjects int64                                 `json:"totalSubjects"`
	JSON          projectListResponseProjectsUsedByJSON `json:"-"`
}

// projectListResponseProjectsUsedByJSON contains the JSON metadata for the struct
// [ProjectListResponseProjectsUsedBy]
type projectListResponseProjectsUsedByJSON struct {
	Subjects      apijson.Field
	TotalSubjects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectListResponseProjectsUsedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseProjectsUsedByJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseProjectsUsedBySubject struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal ProjectListResponseProjectsUsedBySubjectsPrincipal `json:"principal"`
	JSON      projectListResponseProjectsUsedBySubjectJSON       `json:"-"`
}

// projectListResponseProjectsUsedBySubjectJSON contains the JSON metadata for the
// struct [ProjectListResponseProjectsUsedBySubject]
type projectListResponseProjectsUsedBySubjectJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseProjectsUsedBySubject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseProjectsUsedBySubjectJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type ProjectListResponseProjectsUsedBySubjectsPrincipal string

const (
	ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalUnspecified    ProjectListResponseProjectsUsedBySubjectsPrincipal = "PRINCIPAL_UNSPECIFIED"
	ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalAccount        ProjectListResponseProjectsUsedBySubjectsPrincipal = "PRINCIPAL_ACCOUNT"
	ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalUser           ProjectListResponseProjectsUsedBySubjectsPrincipal = "PRINCIPAL_USER"
	ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalRunner         ProjectListResponseProjectsUsedBySubjectsPrincipal = "PRINCIPAL_RUNNER"
	ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalEnvironment    ProjectListResponseProjectsUsedBySubjectsPrincipal = "PRINCIPAL_ENVIRONMENT"
	ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalServiceAccount ProjectListResponseProjectsUsedBySubjectsPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r ProjectListResponseProjectsUsedBySubjectsPrincipal) IsKnown() bool {
	switch r {
	case ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalUnspecified, ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalAccount, ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalUser, ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalRunner, ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalEnvironment, ProjectListResponseProjectsUsedBySubjectsPrincipalPrincipalServiceAccount:
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
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject],
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject],
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject].
func (r ProjectNewFromEnvironmentResponseProjectEnvironmentClass) AsUnion() ProjectNewFromEnvironmentResponseProjectEnvironmentClassUnion {
	return r.union
}

// Union satisfied by
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject],
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject] or
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject].
type ProjectNewFromEnvironmentResponseProjectEnvironmentClassUnion interface {
	implementsProjectNewFromEnvironmentResponseProjectEnvironmentClass()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewFromEnvironmentResponseProjectEnvironmentClassUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject{}),
		},
	)
}

type ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID string `json:"environmentClassId,required" format:"uuid"`
	// Use a local runner for the user
	LocalRunner bool                                                               `json:"localRunner"`
	JSON        projectNewFromEnvironmentResponseProjectEnvironmentClassObjectJSON `json:"-"`
}

// projectNewFromEnvironmentResponseProjectEnvironmentClassObjectJSON contains the
// JSON metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject]
type projectNewFromEnvironmentResponseProjectEnvironmentClassObjectJSON struct {
	EnvironmentClassID apijson.Field
	LocalRunner        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectEnvironmentClassObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewFromEnvironmentResponseProjectEnvironmentClassObject) implementsProjectNewFromEnvironmentResponseProjectEnvironmentClass() {
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
	// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectContextURL].
	ContextURL interface{} `json:"contextUrl"`
	// This field can have the runtime type of
	// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGit].
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
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject],
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject],
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject].
func (r ProjectNewFromEnvironmentResponseProjectInitializerSpec) AsUnion() ProjectNewFromEnvironmentResponseProjectInitializerSpecsUnion {
	return r.union
}

// Union satisfied by
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject],
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject] or
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject].
type ProjectNewFromEnvironmentResponseProjectInitializerSpecsUnion interface {
	implementsProjectNewFromEnvironmentResponseProjectInitializerSpec()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewFromEnvironmentResponseProjectInitializerSpecsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject{}),
		},
	)
}

type ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject struct {
	ContextURL ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectContextURL `json:"contextUrl,required"`
	Git        ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGit        `json:"git"`
	JSON       projectNewFromEnvironmentResponseProjectInitializerSpecsObjectJSON       `json:"-"`
}

// projectNewFromEnvironmentResponseProjectInitializerSpecsObjectJSON contains the
// JSON metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject]
type projectNewFromEnvironmentResponseProjectInitializerSpecsObjectJSON struct {
	ContextURL  apijson.Field
	Git         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectInitializerSpecsObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewFromEnvironmentResponseProjectInitializerSpecsObject) implementsProjectNewFromEnvironmentResponseProjectInitializerSpec() {
}

type ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectContextURL struct {
	// url is the URL from which the environment is created
	URL  string                                                                       `json:"url" format:"uri"`
	JSON projectNewFromEnvironmentResponseProjectInitializerSpecsObjectContextURLJSON `json:"-"`
}

// projectNewFromEnvironmentResponseProjectInitializerSpecsObjectContextURLJSON
// contains the JSON metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectContextURL]
type projectNewFromEnvironmentResponseProjectInitializerSpecsObjectContextURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectContextURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectInitializerSpecsObjectContextURLJSON) RawJSON() string {
	return r.raw
}

type ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGit struct {
	// a path relative to the environment root in which the code will be checked out
	//
	// to
	CheckoutLocation string `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget string `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri string `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetMode `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri string                                                                `json:"upstreamRemoteUri"`
	JSON              projectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitJSON `json:"-"`
}

// projectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitJSON contains
// the JSON metadata for the struct
// [ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGit]
type projectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitJSON struct {
	CheckoutLocation  apijson.Field
	CloneTarget       apijson.Field
	RemoteUri         apijson.Field
	TargetMode        apijson.Field
	UpstreamRemoteUri apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitJSON) RawJSON() string {
	return r.raw
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetMode string

const (
	ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified  ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead   ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch  ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified, ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead, ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit, ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch, ProjectNewFromEnvironmentResponseProjectInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

type ProjectNewFromEnvironmentResponseProjectMetadata struct {
	// A Timestamp represents a point in time independent of any time zone or local
	//
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
	//
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
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ProjectNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
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
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
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

// Satisfied by [ProjectNewParamsEnvironmentClassObject],
// [ProjectNewParamsEnvironmentClassObject],
// [ProjectNewParamsEnvironmentClassObject], [ProjectNewParamsEnvironmentClass].
type ProjectNewParamsEnvironmentClassUnion interface {
	implementsProjectNewParamsEnvironmentClassUnion()
}

type ProjectNewParamsEnvironmentClassObject struct {
	// Use a fixed environment class on a given Runner. This cannot be a local runner's
	// environment class.
	EnvironmentClassID param.Field[string] `json:"environmentClassId,required" format:"uuid"`
	// Use a local runner for the user
	LocalRunner param.Field[bool] `json:"localRunner"`
}

func (r ProjectNewParamsEnvironmentClassObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsEnvironmentClassObject) implementsProjectNewParamsEnvironmentClassUnion() {}

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

// Satisfied by [ProjectNewParamsInitializerSpecsObject],
// [ProjectNewParamsInitializerSpecsObject],
// [ProjectNewParamsInitializerSpecsObject], [ProjectNewParamsInitializerSpec].
type ProjectNewParamsInitializerSpecUnion interface {
	implementsProjectNewParamsInitializerSpecUnion()
}

type ProjectNewParamsInitializerSpecsObject struct {
	ContextURL param.Field[ProjectNewParamsInitializerSpecsObjectContextURL] `json:"contextUrl,required"`
	Git        param.Field[ProjectNewParamsInitializerSpecsObjectGit]        `json:"git"`
}

func (r ProjectNewParamsInitializerSpecsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsInitializerSpecsObject) implementsProjectNewParamsInitializerSpecUnion() {}

type ProjectNewParamsInitializerSpecsObjectContextURL struct {
	// url is the URL from which the environment is created
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r ProjectNewParamsInitializerSpecsObjectContextURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectNewParamsInitializerSpecsObjectGit struct {
	// a path relative to the environment root in which the code will be checked out
	//
	// to
	CheckoutLocation param.Field[string] `json:"checkoutLocation"`
	// the value for the clone target mode - use depends on the target mode
	CloneTarget param.Field[string] `json:"cloneTarget"`
	// remote_uri is the Git remote origin
	RemoteUri param.Field[string] `json:"remoteUri"`
	// CloneTargetMode is the target state in which we want to leave a GitEnvironment
	TargetMode param.Field[ProjectNewParamsInitializerSpecsObjectGitTargetMode] `json:"targetMode"`
	// upstream_Remote_uri is the fork upstream of a repository
	UpstreamRemoteUri param.Field[string] `json:"upstreamRemoteUri"`
}

func (r ProjectNewParamsInitializerSpecsObjectGit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CloneTargetMode is the target state in which we want to leave a GitEnvironment
type ProjectNewParamsInitializerSpecsObjectGitTargetMode string

const (
	ProjectNewParamsInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified  ProjectNewParamsInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_UNSPECIFIED"
	ProjectNewParamsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead   ProjectNewParamsInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_HEAD"
	ProjectNewParamsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit ProjectNewParamsInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_COMMIT"
	ProjectNewParamsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch ProjectNewParamsInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_REMOTE_BRANCH"
	ProjectNewParamsInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch  ProjectNewParamsInitializerSpecsObjectGitTargetMode = "CLONE_TARGET_MODE_LOCAL_BRANCH"
)

func (r ProjectNewParamsInitializerSpecsObjectGitTargetMode) IsKnown() bool {
	switch r {
	case ProjectNewParamsInitializerSpecsObjectGitTargetModeCloneTargetModeUnspecified, ProjectNewParamsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteHead, ProjectNewParamsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteCommit, ProjectNewParamsInitializerSpecsObjectGitTargetModeCloneTargetModeRemoteBranch, ProjectNewParamsInitializerSpecsObjectGitTargetModeCloneTargetModeLocalBranch:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type ProjectNewParamsConnectProtocolVersion float64

const (
	ProjectNewParamsConnectProtocolVersion1 ProjectNewParamsConnectProtocolVersion = 1
)

func (r ProjectNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ProjectNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ProjectGetParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[ProjectGetParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ProjectGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[ProjectGetParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[ProjectGetParamsConnect] `query:"connect"`
	Message param.Field[string]                  `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [ProjectGetParams]'s query parameters as `url.Values`.
func (r ProjectGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type ProjectGetParamsEncoding string

const (
	ProjectGetParamsEncodingProto ProjectGetParamsEncoding = "proto"
	ProjectGetParamsEncodingJson  ProjectGetParamsEncoding = "json"
)

func (r ProjectGetParamsEncoding) IsKnown() bool {
	switch r {
	case ProjectGetParamsEncodingProto, ProjectGetParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type ProjectGetParamsConnectProtocolVersion float64

const (
	ProjectGetParamsConnectProtocolVersion1 ProjectGetParamsConnectProtocolVersion = 1
)

func (r ProjectGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ProjectGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type ProjectGetParamsCompression string

const (
	ProjectGetParamsCompressionIdentity ProjectGetParamsCompression = "identity"
	ProjectGetParamsCompressionGzip     ProjectGetParamsCompression = "gzip"
	ProjectGetParamsCompressionBr       ProjectGetParamsCompression = "br"
)

func (r ProjectGetParamsCompression) IsKnown() bool {
	switch r {
	case ProjectGetParamsCompressionIdentity, ProjectGetParamsCompressionGzip, ProjectGetParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type ProjectGetParamsConnect string

const (
	ProjectGetParamsConnectV1 ProjectGetParamsConnect = "v1"
)

func (r ProjectGetParamsConnect) IsKnown() bool {
	switch r {
	case ProjectGetParamsConnectV1:
		return true
	}
	return false
}

type ProjectUpdateParams struct {
	Body ProjectUpdateParamsBody `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ProjectUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectUpdateParamsBody struct {
}

func (r ProjectUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ProjectUpdateParamsConnectProtocolVersion float64

const (
	ProjectUpdateParamsConnectProtocolVersion1 ProjectUpdateParamsConnectProtocolVersion = 1
)

func (r ProjectUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ProjectUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ProjectListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[ProjectListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ProjectListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[ProjectListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[ProjectListParamsConnect] `query:"connect"`
	Message param.Field[string]                   `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type ProjectListParamsEncoding string

const (
	ProjectListParamsEncodingProto ProjectListParamsEncoding = "proto"
	ProjectListParamsEncodingJson  ProjectListParamsEncoding = "json"
)

func (r ProjectListParamsEncoding) IsKnown() bool {
	switch r {
	case ProjectListParamsEncodingProto, ProjectListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type ProjectListParamsConnectProtocolVersion float64

const (
	ProjectListParamsConnectProtocolVersion1 ProjectListParamsConnectProtocolVersion = 1
)

func (r ProjectListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ProjectListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type ProjectListParamsCompression string

const (
	ProjectListParamsCompressionIdentity ProjectListParamsCompression = "identity"
	ProjectListParamsCompressionGzip     ProjectListParamsCompression = "gzip"
	ProjectListParamsCompressionBr       ProjectListParamsCompression = "br"
)

func (r ProjectListParamsCompression) IsKnown() bool {
	switch r {
	case ProjectListParamsCompressionIdentity, ProjectListParamsCompressionGzip, ProjectListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type ProjectListParamsConnect string

const (
	ProjectListParamsConnectV1 ProjectListParamsConnect = "v1"
)

func (r ProjectListParamsConnect) IsKnown() bool {
	switch r {
	case ProjectListParamsConnectV1:
		return true
	}
	return false
}

type ProjectDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ProjectDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// project_id specifies the project identifier
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ProjectDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ProjectDeleteParamsConnectProtocolVersion float64

const (
	ProjectDeleteParamsConnectProtocolVersion1 ProjectDeleteParamsConnectProtocolVersion = 1
)

func (r ProjectDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ProjectDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ProjectNewFromEnvironmentParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ProjectNewFromEnvironmentParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// environment_id specifies the environment identifier
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
	Name          param.Field[string] `json:"name"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ProjectNewFromEnvironmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ProjectNewFromEnvironmentParamsConnectProtocolVersion float64

const (
	ProjectNewFromEnvironmentParamsConnectProtocolVersion1 ProjectNewFromEnvironmentParamsConnectProtocolVersion = 1
)

func (r ProjectNewFromEnvironmentParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ProjectNewFromEnvironmentParamsConnectProtocolVersion1:
		return true
	}
	return false
}
