// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// OrganizationMemberService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationMemberService] method instead.
type OrganizationMemberService struct {
	Options []option.RequestOption
}

// NewOrganizationMemberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationMemberService(opts ...option.RequestOption) (r *OrganizationMemberService) {
	r = &OrganizationMemberService{}
	r.Options = opts
	return
}

// ListMembers lists all members of the specified organization.
func (r *OrganizationMemberService) List(ctx context.Context, params OrganizationMemberListParams, opts ...option.RequestOption) (res *OrganizationMemberListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/ListMembers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type OrganizationMemberListResponse struct {
	// members are the members of the organization
	Members []OrganizationMemberListResponseMember `json:"members"`
	// pagination contains the pagination options for listing members
	Pagination OrganizationMemberListResponsePagination `json:"pagination"`
	JSON       organizationMemberListResponseJSON       `json:"-"`
}

// organizationMemberListResponseJSON contains the JSON metadata for the struct
// [OrganizationMemberListResponse]
type organizationMemberListResponseJSON struct {
	Members     apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMemberListResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationMemberListResponseMember struct {
	AvatarURL string `json:"avatarUrl"`
	Email     string `json:"email"`
	FullName  string `json:"fullName"`
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
	MemberSince time.Time                                   `json:"memberSince" format:"date-time"`
	Role        OrganizationMemberListResponseMembersRole   `json:"role"`
	Status      OrganizationMemberListResponseMembersStatus `json:"status"`
	UserID      string                                      `json:"userId" format:"uuid"`
	JSON        organizationMemberListResponseMemberJSON    `json:"-"`
}

// organizationMemberListResponseMemberJSON contains the JSON metadata for the
// struct [OrganizationMemberListResponseMember]
type organizationMemberListResponseMemberJSON struct {
	AvatarURL   apijson.Field
	Email       apijson.Field
	FullName    apijson.Field
	MemberSince apijson.Field
	Role        apijson.Field
	Status      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberListResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMemberListResponseMemberJSON) RawJSON() string {
	return r.raw
}

type OrganizationMemberListResponseMembersRole string

const (
	OrganizationMemberListResponseMembersRoleOrganizationRoleUnspecified OrganizationMemberListResponseMembersRole = "ORGANIZATION_ROLE_UNSPECIFIED"
	OrganizationMemberListResponseMembersRoleOrganizationRoleAdmin       OrganizationMemberListResponseMembersRole = "ORGANIZATION_ROLE_ADMIN"
	OrganizationMemberListResponseMembersRoleOrganizationRoleMember      OrganizationMemberListResponseMembersRole = "ORGANIZATION_ROLE_MEMBER"
)

func (r OrganizationMemberListResponseMembersRole) IsKnown() bool {
	switch r {
	case OrganizationMemberListResponseMembersRoleOrganizationRoleUnspecified, OrganizationMemberListResponseMembersRoleOrganizationRoleAdmin, OrganizationMemberListResponseMembersRoleOrganizationRoleMember:
		return true
	}
	return false
}

type OrganizationMemberListResponseMembersStatus string

const (
	OrganizationMemberListResponseMembersStatusUserStatusUnspecified OrganizationMemberListResponseMembersStatus = "USER_STATUS_UNSPECIFIED"
	OrganizationMemberListResponseMembersStatusUserStatusActive      OrganizationMemberListResponseMembersStatus = "USER_STATUS_ACTIVE"
	OrganizationMemberListResponseMembersStatusUserStatusSuspended   OrganizationMemberListResponseMembersStatus = "USER_STATUS_SUSPENDED"
	OrganizationMemberListResponseMembersStatusUserStatusLeft        OrganizationMemberListResponseMembersStatus = "USER_STATUS_LEFT"
)

func (r OrganizationMemberListResponseMembersStatus) IsKnown() bool {
	switch r {
	case OrganizationMemberListResponseMembersStatusUserStatusUnspecified, OrganizationMemberListResponseMembersStatusUserStatusActive, OrganizationMemberListResponseMembersStatusUserStatusSuspended, OrganizationMemberListResponseMembersStatusUserStatusLeft:
		return true
	}
	return false
}

// pagination contains the pagination options for listing members
type OrganizationMemberListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                       `json:"nextToken"`
	JSON      organizationMemberListResponsePaginationJSON `json:"-"`
}

// organizationMemberListResponsePaginationJSON contains the JSON metadata for the
// struct [OrganizationMemberListResponsePagination]
type organizationMemberListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMemberListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type OrganizationMemberListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationMemberListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// organization_id is the ID of the organization to list members for
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
	// pagination contains the pagination options for listing members
	Pagination param.Field[OrganizationMemberListParamsPagination] `json:"pagination"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationMemberListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationMemberListParamsConnectProtocolVersion float64

const (
	OrganizationMemberListParamsConnectProtocolVersion1 OrganizationMemberListParamsConnectProtocolVersion = 1
)

func (r OrganizationMemberListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationMemberListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// pagination contains the pagination options for listing members
type OrganizationMemberListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r OrganizationMemberListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
