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

// OrganizationService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
	Invite  *OrganizationInviteService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.Invite = NewOrganizationInviteService(opts...)
	return
}

// LeaveOrganization lets the passed user leave an Organization.
func (r *OrganizationService) Leave(ctx context.Context, body OrganizationLeaveParams, opts ...option.RequestOption) (res *OrganizationLeaveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/LeaveOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListMembers lists all members of the specified organization.
func (r *OrganizationService) ListMembers(ctx context.Context, body OrganizationListMembersParams, opts ...option.RequestOption) (res *OrganizationListMembersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/ListMembers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// SetRole
func (r *OrganizationService) SetRole(ctx context.Context, body OrganizationSetRoleParams, opts ...option.RequestOption) (res *OrganizationSetRoleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/SetRole"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type OrganizationLeaveResponse = interface{}

type OrganizationListMembersResponse struct {
	// members are the members of the organization
	Members []OrganizationListMembersResponseMember `json:"members"`
	// pagination contains the pagination options for listing members
	Pagination OrganizationListMembersResponsePagination `json:"pagination"`
	JSON       organizationListMembersResponseJSON       `json:"-"`
}

// organizationListMembersResponseJSON contains the JSON metadata for the struct
// [OrganizationListMembersResponse]
type organizationListMembersResponseJSON struct {
	Members     apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListMembersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListMembersResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationListMembersResponseMember struct {
	AvatarURL string `json:"avatarUrl"`
	Email     string `json:"email"`
	FullName  string `json:"fullName"`
	// login_provider is the login provider the user uses to sign in
	LoginProvider string `json:"loginProvider"`
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
	MemberSince time.Time                                    `json:"memberSince" format:"date-time"`
	Role        OrganizationListMembersResponseMembersRole   `json:"role"`
	Status      OrganizationListMembersResponseMembersStatus `json:"status"`
	UserID      string                                       `json:"userId" format:"uuid"`
	JSON        organizationListMembersResponseMemberJSON    `json:"-"`
}

// organizationListMembersResponseMemberJSON contains the JSON metadata for the
// struct [OrganizationListMembersResponseMember]
type organizationListMembersResponseMemberJSON struct {
	AvatarURL     apijson.Field
	Email         apijson.Field
	FullName      apijson.Field
	LoginProvider apijson.Field
	MemberSince   apijson.Field
	Role          apijson.Field
	Status        apijson.Field
	UserID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationListMembersResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListMembersResponseMemberJSON) RawJSON() string {
	return r.raw
}

type OrganizationListMembersResponseMembersRole string

const (
	OrganizationListMembersResponseMembersRoleOrganizationRoleUnspecified OrganizationListMembersResponseMembersRole = "ORGANIZATION_ROLE_UNSPECIFIED"
	OrganizationListMembersResponseMembersRoleOrganizationRoleAdmin       OrganizationListMembersResponseMembersRole = "ORGANIZATION_ROLE_ADMIN"
	OrganizationListMembersResponseMembersRoleOrganizationRoleMember      OrganizationListMembersResponseMembersRole = "ORGANIZATION_ROLE_MEMBER"
)

func (r OrganizationListMembersResponseMembersRole) IsKnown() bool {
	switch r {
	case OrganizationListMembersResponseMembersRoleOrganizationRoleUnspecified, OrganizationListMembersResponseMembersRoleOrganizationRoleAdmin, OrganizationListMembersResponseMembersRoleOrganizationRoleMember:
		return true
	}
	return false
}

type OrganizationListMembersResponseMembersStatus string

const (
	OrganizationListMembersResponseMembersStatusUserStatusUnspecified OrganizationListMembersResponseMembersStatus = "USER_STATUS_UNSPECIFIED"
	OrganizationListMembersResponseMembersStatusUserStatusActive      OrganizationListMembersResponseMembersStatus = "USER_STATUS_ACTIVE"
	OrganizationListMembersResponseMembersStatusUserStatusSuspended   OrganizationListMembersResponseMembersStatus = "USER_STATUS_SUSPENDED"
	OrganizationListMembersResponseMembersStatusUserStatusLeft        OrganizationListMembersResponseMembersStatus = "USER_STATUS_LEFT"
)

func (r OrganizationListMembersResponseMembersStatus) IsKnown() bool {
	switch r {
	case OrganizationListMembersResponseMembersStatusUserStatusUnspecified, OrganizationListMembersResponseMembersStatusUserStatusActive, OrganizationListMembersResponseMembersStatusUserStatusSuspended, OrganizationListMembersResponseMembersStatusUserStatusLeft:
		return true
	}
	return false
}

// pagination contains the pagination options for listing members
type OrganizationListMembersResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                        `json:"nextToken"`
	JSON      organizationListMembersResponsePaginationJSON `json:"-"`
}

// organizationListMembersResponsePaginationJSON contains the JSON metadata for the
// struct [OrganizationListMembersResponsePagination]
type organizationListMembersResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListMembersResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListMembersResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type OrganizationSetRoleResponse = interface{}

type OrganizationLeaveParams struct {
	UserID param.Field[string] `json:"userId" format:"uuid"`
}

func (r OrganizationLeaveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationListMembersParams struct {
	// organization_id is the ID of the organization to list members for
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
	// pagination contains the pagination options for listing members
	Pagination param.Field[OrganizationListMembersParamsPagination] `json:"pagination"`
}

func (r OrganizationListMembersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing members
type OrganizationListMembersParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	//
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r OrganizationListMembersParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationSetRoleParams struct {
	OrganizationID param.Field[string]                        `json:"organizationId" format:"uuid"`
	Role           param.Field[OrganizationSetRoleParamsRole] `json:"role"`
	UserID         param.Field[string]                        `json:"userId" format:"uuid"`
}

func (r OrganizationSetRoleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationSetRoleParamsRole string

const (
	OrganizationSetRoleParamsRoleOrganizationRoleUnspecified OrganizationSetRoleParamsRole = "ORGANIZATION_ROLE_UNSPECIFIED"
	OrganizationSetRoleParamsRoleOrganizationRoleAdmin       OrganizationSetRoleParamsRole = "ORGANIZATION_ROLE_ADMIN"
	OrganizationSetRoleParamsRoleOrganizationRoleMember      OrganizationSetRoleParamsRole = "ORGANIZATION_ROLE_MEMBER"
)

func (r OrganizationSetRoleParamsRole) IsKnown() bool {
	switch r {
	case OrganizationSetRoleParamsRoleOrganizationRoleUnspecified, OrganizationSetRoleParamsRoleOrganizationRoleAdmin, OrganizationSetRoleParamsRoleOrganizationRoleMember:
		return true
	}
	return false
}
