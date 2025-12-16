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
)

// GroupService contains methods and other services that help with interacting with
// the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupService] method instead.
type GroupService struct {
	Options         []option.RequestOption
	Memberships     *GroupMembershipService
	RoleAssignments *GroupRoleAssignmentService
}

// NewGroupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGroupService(opts ...option.RequestOption) (r *GroupService) {
	r = &GroupService{}
	r.Options = opts
	r.Memberships = NewGroupMembershipService(opts...)
	r.RoleAssignments = NewGroupRoleAssignmentService(opts...)
	return
}

// Creates a new group within an organization.
//
// Use this method to:
//
// - Create teams for access control
// - Organize users by department or function
// - Set up role-based access groups
//
// ### Examples
//
// - Create a basic group:
//
//	Creates a group with name and description.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	name: "Backend Team"
//	description: "Backend engineering team"
//	```
//
// ### Authorization
//
// Requires `org:admin` role on the organization.
func (r *GroupService) New(ctx context.Context, body GroupNewParams, opts ...option.RequestOption) (res *GroupNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.GroupService/CreateGroup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets information about a specific group.
//
// Use this method to:
//
// - Retrieve group details and metadata
// - Check group configuration
// - View member count
//
// ### Examples
//
// - Get group details:
//
//	Retrieves information about a specific group.
//
//	```yaml
//	groupId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
//
// ### Authorization
//
// All organization members can view group information (transparency model).
func (r *GroupService) Get(ctx context.Context, body GroupGetParams, opts ...option.RequestOption) (res *GroupGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.GroupService/GetGroup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates group information.
//
// Use this method to:
//
// - Rename a group
// - Update group description
//
// ### Examples
//
// - Update group name:
//
//	Changes the name of an existing group.
//
//	```yaml
//	groupId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	name: "Platform Team"
//	description: "Platform engineering team"
//	```
//
// ### Authorization
//
// Requires `org:admin` permission on the organization or `group:admin` permission
// on the specific group.
func (r *GroupService) Update(ctx context.Context, body GroupUpdateParams, opts ...option.RequestOption) (res *GroupUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.GroupService/UpdateGroup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists groups with optional pagination.
//
// Use this method to:
//
// - View all groups in an organization
// - Check group memberships
// - Monitor group configurations
// - Audit group access
//
// ### Examples
//
// - List all groups:
//
//	Shows all groups with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - List with custom page size:
//
//	Shows groups with specified page size.
//
//	```yaml
//	pagination:
//	  pageSize: 50
//	  token: "next-page-token-from-previous-response"
//	```
//
// ### Authorization
//
// All organization members can list groups (transparency model).
func (r *GroupService) List(ctx context.Context, params GroupListParams, opts ...option.RequestOption) (res *pagination.GroupsPage[Group], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.GroupService/ListGroups"
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

// Lists groups with optional pagination.
//
// Use this method to:
//
// - View all groups in an organization
// - Check group memberships
// - Monitor group configurations
// - Audit group access
//
// ### Examples
//
// - List all groups:
//
//	Shows all groups with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - List with custom page size:
//
//	Shows groups with specified page size.
//
//	```yaml
//	pagination:
//	  pageSize: 50
//	  token: "next-page-token-from-previous-response"
//	```
//
// ### Authorization
//
// All organization members can list groups (transparency model).
func (r *GroupService) ListAutoPaging(ctx context.Context, params GroupListParams, opts ...option.RequestOption) *pagination.GroupsPageAutoPager[Group] {
	return pagination.NewGroupsPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a group and removes all its resource assignments.
//
// When a group is deleted, all resource assignments revert to org-level scope.
//
// Use this method to:
//
// - Remove unused groups
// - Clean up after team reorganization
//
// ### Examples
//
// - Delete a group:
//
//	Permanently removes a group.
//
//	```yaml
//	groupId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
//
// ### Authorization
//
// Requires `org:admin` role on the organization.
func (r *GroupService) Delete(ctx context.Context, body GroupDeleteParams, opts ...option.RequestOption) (res *GroupDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.GroupService/DeleteGroup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Group struct {
	ID string `json:"id" format:"uuid"`
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
	CreatedAt   time.Time `json:"createdAt" format:"date-time"`
	Description string    `json:"description"`
	// member_count is the total number of members in this group
	MemberCount    int64  `json:"memberCount"`
	Name           string `json:"name"`
	OrganizationID string `json:"organizationId" format:"uuid"`
	// system_managed indicates that this group is created by the system automatically
	SystemManaged bool `json:"systemManaged"`
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
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	JSON      groupJSON `json:"-"`
}

// groupJSON contains the JSON metadata for the struct [Group]
type groupJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Description    apijson.Field
	MemberCount    apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	SystemManaged  apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Group) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupJSON) RawJSON() string {
	return r.raw
}

type GroupNewResponse struct {
	Group Group                `json:"group"`
	JSON  groupNewResponseJSON `json:"-"`
}

// groupNewResponseJSON contains the JSON metadata for the struct
// [GroupNewResponse]
type groupNewResponseJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupNewResponseJSON) RawJSON() string {
	return r.raw
}

type GroupGetResponse struct {
	Group Group                `json:"group"`
	JSON  groupGetResponseJSON `json:"-"`
}

// groupGetResponseJSON contains the JSON metadata for the struct
// [GroupGetResponse]
type groupGetResponseJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupGetResponseJSON) RawJSON() string {
	return r.raw
}

type GroupUpdateResponse struct {
	Group Group                   `json:"group"`
	JSON  groupUpdateResponseJSON `json:"-"`
}

// groupUpdateResponseJSON contains the JSON metadata for the struct
// [GroupUpdateResponse]
type groupUpdateResponseJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type GroupDeleteResponse = interface{}

type GroupNewParams struct {
	Description    param.Field[string] `json:"description"`
	Name           param.Field[string] `json:"name"`
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
}

func (r GroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GroupGetParams struct {
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
}

func (r GroupGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GroupUpdateParams struct {
	Description param.Field[string] `json:"description"`
	GroupID     param.Field[string] `json:"groupId" format:"uuid"`
	Name        param.Field[string] `json:"name"`
}

func (r GroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GroupListParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// pagination contains the pagination options for listing groups
	Pagination param.Field[GroupListParamsPagination] `json:"pagination"`
}

func (r GroupListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [GroupListParams]'s query parameters as `url.Values`.
func (r GroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// pagination contains the pagination options for listing groups
type GroupListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r GroupListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GroupDeleteParams struct {
	GroupID param.Field[string] `json:"groupId" format:"uuid"`
}

func (r GroupDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
