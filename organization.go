// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
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
	Options           []option.RequestOption
	Invites           *OrganizationInviteService
	SSOConfigurations *OrganizationSSOConfigurationService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.Invites = NewOrganizationInviteService(opts...)
	r.SSOConfigurations = NewOrganizationSSOConfigurationService(opts...)
	return
}

// CreateOrganization creates a new Organization.
func (r *OrganizationService) New(ctx context.Context, params OrganizationNewParams, opts ...option.RequestOption) (res *OrganizationNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/CreateOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetOrganization retrieves a single Organization.
func (r *OrganizationService) Get(ctx context.Context, params OrganizationGetParams, opts ...option.RequestOption) (res *OrganizationGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/GetOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// UpdateOrganization updates the properties of an Organization.
func (r *OrganizationService) Update(ctx context.Context, params OrganizationUpdateParams, opts ...option.RequestOption) (res *OrganizationUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/UpdateOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListOrganizations lists all organization the caller has access to.
func (r *OrganizationService) List(ctx context.Context, params OrganizationListParams, opts ...option.RequestOption) (res *OrganizationListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/ListOrganizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// DeleteOrganization deletes the specified organization.
func (r *OrganizationService) Delete(ctx context.Context, params OrganizationDeleteParams, opts ...option.RequestOption) (res *OrganizationDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/DeleteOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// JoinOrganization lets accounts join an Organization.
func (r *OrganizationService) Join(ctx context.Context, params OrganizationJoinParams, opts ...option.RequestOption) (res *OrganizationJoinResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/JoinOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// LeaveOrganization lets the passed user leave an Organization.
func (r *OrganizationService) Leave(ctx context.Context, params OrganizationLeaveParams, opts ...option.RequestOption) (res *OrganizationLeaveResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/LeaveOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListMembers lists all members of the specified organization.
func (r *OrganizationService) ListMembers(ctx context.Context, params OrganizationListMembersParams, opts ...option.RequestOption) (res *OrganizationListMembersResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/ListMembers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// SetRole
func (r *OrganizationService) SetRole(ctx context.Context, params OrganizationSetRoleParams, opts ...option.RequestOption) (res *OrganizationSetRoleResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/SetRole"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type OrganizationNewResponse struct {
	// member is the member that joined the org on creation. Only set if specified
	// "join_organization" is "true" in the request.
	Member OrganizationNewResponseMember `json:"member"`
	// organization is the created organization
	Organization OrganizationNewResponseOrganization `json:"organization"`
	JSON         organizationNewResponseJSON         `json:"-"`
}

// organizationNewResponseJSON contains the JSON metadata for the struct
// [OrganizationNewResponse]
type organizationNewResponseJSON struct {
	Member       apijson.Field
	Organization apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseJSON) RawJSON() string {
	return r.raw
}

// member is the member that joined the org on creation. Only set if specified
// "join_organization" is "true" in the request.
type OrganizationNewResponseMember struct {
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
	MemberSince time.Time                           `json:"memberSince" format:"date-time"`
	Role        OrganizationNewResponseMemberRole   `json:"role"`
	Status      OrganizationNewResponseMemberStatus `json:"status"`
	UserID      string                              `json:"userId" format:"uuid"`
	JSON        organizationNewResponseMemberJSON   `json:"-"`
}

// organizationNewResponseMemberJSON contains the JSON metadata for the struct
// [OrganizationNewResponseMember]
type organizationNewResponseMemberJSON struct {
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

func (r *OrganizationNewResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseMemberJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseMemberRole string

const (
	OrganizationNewResponseMemberRoleOrganizationRoleUnspecified OrganizationNewResponseMemberRole = "ORGANIZATION_ROLE_UNSPECIFIED"
	OrganizationNewResponseMemberRoleOrganizationRoleAdmin       OrganizationNewResponseMemberRole = "ORGANIZATION_ROLE_ADMIN"
	OrganizationNewResponseMemberRoleOrganizationRoleMember      OrganizationNewResponseMemberRole = "ORGANIZATION_ROLE_MEMBER"
)

func (r OrganizationNewResponseMemberRole) IsKnown() bool {
	switch r {
	case OrganizationNewResponseMemberRoleOrganizationRoleUnspecified, OrganizationNewResponseMemberRoleOrganizationRoleAdmin, OrganizationNewResponseMemberRoleOrganizationRoleMember:
		return true
	}
	return false
}

type OrganizationNewResponseMemberStatus string

const (
	OrganizationNewResponseMemberStatusUserStatusUnspecified OrganizationNewResponseMemberStatus = "USER_STATUS_UNSPECIFIED"
	OrganizationNewResponseMemberStatusUserStatusActive      OrganizationNewResponseMemberStatus = "USER_STATUS_ACTIVE"
	OrganizationNewResponseMemberStatusUserStatusSuspended   OrganizationNewResponseMemberStatus = "USER_STATUS_SUSPENDED"
	OrganizationNewResponseMemberStatusUserStatusLeft        OrganizationNewResponseMemberStatus = "USER_STATUS_LEFT"
)

func (r OrganizationNewResponseMemberStatus) IsKnown() bool {
	switch r {
	case OrganizationNewResponseMemberStatusUserStatusUnspecified, OrganizationNewResponseMemberStatusUserStatusActive, OrganizationNewResponseMemberStatusUserStatusSuspended, OrganizationNewResponseMemberStatusUserStatusLeft:
		return true
	}
	return false
}

// organization is the created organization
type OrganizationNewResponseOrganization struct {
	ID string `json:"id" format:"uuid"`
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
	CreatedAt     time.Time                                        `json:"createdAt" format:"date-time"`
	InviteDomains OrganizationNewResponseOrganizationInviteDomains `json:"inviteDomains"`
	Name          string                                           `json:"name"`
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
	JSON      organizationNewResponseOrganizationJSON `json:"-"`
}

// organizationNewResponseOrganizationJSON contains the JSON metadata for the
// struct [OrganizationNewResponseOrganization]
type organizationNewResponseOrganizationJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	InviteDomains apijson.Field
	Name          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationNewResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseOrganizationInviteDomains struct {
	// domains is the list of domains that are allowed to join the organization
	Domains []string                                             `json:"domains"`
	JSON    organizationNewResponseOrganizationInviteDomainsJSON `json:"-"`
}

// organizationNewResponseOrganizationInviteDomainsJSON contains the JSON metadata
// for the struct [OrganizationNewResponseOrganizationInviteDomains]
type organizationNewResponseOrganizationInviteDomainsJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationNewResponseOrganizationInviteDomains) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseOrganizationInviteDomainsJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponse struct {
	// organization is the requested organization
	Organization OrganizationGetResponseOrganization `json:"organization"`
	JSON         organizationGetResponseJSON         `json:"-"`
}

// organizationGetResponseJSON contains the JSON metadata for the struct
// [OrganizationGetResponse]
type organizationGetResponseJSON struct {
	Organization apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseJSON) RawJSON() string {
	return r.raw
}

// organization is the requested organization
type OrganizationGetResponseOrganization struct {
	ID string `json:"id" format:"uuid"`
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
	CreatedAt     time.Time                                        `json:"createdAt" format:"date-time"`
	InviteDomains OrganizationGetResponseOrganizationInviteDomains `json:"inviteDomains"`
	Name          string                                           `json:"name"`
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
	JSON      organizationGetResponseOrganizationJSON `json:"-"`
}

// organizationGetResponseOrganizationJSON contains the JSON metadata for the
// struct [OrganizationGetResponseOrganization]
type organizationGetResponseOrganizationJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	InviteDomains apijson.Field
	Name          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationGetResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponseOrganizationInviteDomains struct {
	// domains is the list of domains that are allowed to join the organization
	Domains []string                                             `json:"domains"`
	JSON    organizationGetResponseOrganizationInviteDomainsJSON `json:"-"`
}

// organizationGetResponseOrganizationInviteDomainsJSON contains the JSON metadata
// for the struct [OrganizationGetResponseOrganizationInviteDomains]
type organizationGetResponseOrganizationInviteDomainsJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponseOrganizationInviteDomains) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseOrganizationInviteDomainsJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponse struct {
	// organization is the updated organization
	Organization OrganizationUpdateResponseOrganization `json:"organization"`
	JSON         organizationUpdateResponseJSON         `json:"-"`
}

// organizationUpdateResponseJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponse]
type organizationUpdateResponseJSON struct {
	Organization apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// organization is the updated organization
type OrganizationUpdateResponseOrganization struct {
	ID string `json:"id" format:"uuid"`
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
	CreatedAt     time.Time                                           `json:"createdAt" format:"date-time"`
	InviteDomains OrganizationUpdateResponseOrganizationInviteDomains `json:"inviteDomains"`
	Name          string                                              `json:"name"`
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
	UpdatedAt time.Time                                  `json:"updatedAt" format:"date-time"`
	JSON      organizationUpdateResponseOrganizationJSON `json:"-"`
}

// organizationUpdateResponseOrganizationJSON contains the JSON metadata for the
// struct [OrganizationUpdateResponseOrganization]
type organizationUpdateResponseOrganizationJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	InviteDomains apijson.Field
	Name          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationUpdateResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseOrganizationInviteDomains struct {
	// domains is the list of domains that are allowed to join the organization
	Domains []string                                                `json:"domains"`
	JSON    organizationUpdateResponseOrganizationInviteDomainsJSON `json:"-"`
}

// organizationUpdateResponseOrganizationInviteDomainsJSON contains the JSON
// metadata for the struct [OrganizationUpdateResponseOrganizationInviteDomains]
type organizationUpdateResponseOrganizationInviteDomainsJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponseOrganizationInviteDomains) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseOrganizationInviteDomainsJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponse struct {
	// organizations are the organizations that matched the query
	Organizations []OrganizationListResponseOrganization `json:"organizations"`
	// pagination contains the pagination options for listing organizations
	Pagination OrganizationListResponsePagination `json:"pagination"`
	JSON       organizationListResponseJSON       `json:"-"`
}

// organizationListResponseJSON contains the JSON metadata for the struct
// [OrganizationListResponse]
type organizationListResponseJSON struct {
	Organizations apijson.Field
	Pagination    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseOrganization struct {
	ID string `json:"id" format:"uuid"`
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
	CreatedAt     time.Time                                          `json:"createdAt" format:"date-time"`
	InviteDomains OrganizationListResponseOrganizationsInviteDomains `json:"inviteDomains"`
	Name          string                                             `json:"name"`
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
	JSON      organizationListResponseOrganizationJSON `json:"-"`
}

// organizationListResponseOrganizationJSON contains the JSON metadata for the
// struct [OrganizationListResponseOrganization]
type organizationListResponseOrganizationJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	InviteDomains apijson.Field
	Name          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationListResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseOrganizationsInviteDomains struct {
	// domains is the list of domains that are allowed to join the organization
	Domains []string                                               `json:"domains"`
	JSON    organizationListResponseOrganizationsInviteDomainsJSON `json:"-"`
}

// organizationListResponseOrganizationsInviteDomainsJSON contains the JSON
// metadata for the struct [OrganizationListResponseOrganizationsInviteDomains]
type organizationListResponseOrganizationsInviteDomainsJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponseOrganizationsInviteDomains) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseOrganizationsInviteDomainsJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing organizations
type OrganizationListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                 `json:"nextToken"`
	JSON      organizationListResponsePaginationJSON `json:"-"`
}

// organizationListResponsePaginationJSON contains the JSON metadata for the struct
// [OrganizationListResponsePagination]
type organizationListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type OrganizationDeleteResponse = interface{}

type OrganizationJoinResponse struct {
	// member is the member that was created by joining the organization.
	Member OrganizationJoinResponseMember `json:"member"`
	JSON   organizationJoinResponseJSON   `json:"-"`
}

// organizationJoinResponseJSON contains the JSON metadata for the struct
// [OrganizationJoinResponse]
type organizationJoinResponseJSON struct {
	Member      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationJoinResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationJoinResponseJSON) RawJSON() string {
	return r.raw
}

// member is the member that was created by joining the organization.
type OrganizationJoinResponseMember struct {
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
	MemberSince time.Time                            `json:"memberSince" format:"date-time"`
	Role        OrganizationJoinResponseMemberRole   `json:"role"`
	Status      OrganizationJoinResponseMemberStatus `json:"status"`
	UserID      string                               `json:"userId" format:"uuid"`
	JSON        organizationJoinResponseMemberJSON   `json:"-"`
}

// organizationJoinResponseMemberJSON contains the JSON metadata for the struct
// [OrganizationJoinResponseMember]
type organizationJoinResponseMemberJSON struct {
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

func (r *OrganizationJoinResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationJoinResponseMemberJSON) RawJSON() string {
	return r.raw
}

type OrganizationJoinResponseMemberRole string

const (
	OrganizationJoinResponseMemberRoleOrganizationRoleUnspecified OrganizationJoinResponseMemberRole = "ORGANIZATION_ROLE_UNSPECIFIED"
	OrganizationJoinResponseMemberRoleOrganizationRoleAdmin       OrganizationJoinResponseMemberRole = "ORGANIZATION_ROLE_ADMIN"
	OrganizationJoinResponseMemberRoleOrganizationRoleMember      OrganizationJoinResponseMemberRole = "ORGANIZATION_ROLE_MEMBER"
)

func (r OrganizationJoinResponseMemberRole) IsKnown() bool {
	switch r {
	case OrganizationJoinResponseMemberRoleOrganizationRoleUnspecified, OrganizationJoinResponseMemberRoleOrganizationRoleAdmin, OrganizationJoinResponseMemberRoleOrganizationRoleMember:
		return true
	}
	return false
}

type OrganizationJoinResponseMemberStatus string

const (
	OrganizationJoinResponseMemberStatusUserStatusUnspecified OrganizationJoinResponseMemberStatus = "USER_STATUS_UNSPECIFIED"
	OrganizationJoinResponseMemberStatusUserStatusActive      OrganizationJoinResponseMemberStatus = "USER_STATUS_ACTIVE"
	OrganizationJoinResponseMemberStatusUserStatusSuspended   OrganizationJoinResponseMemberStatus = "USER_STATUS_SUSPENDED"
	OrganizationJoinResponseMemberStatusUserStatusLeft        OrganizationJoinResponseMemberStatus = "USER_STATUS_LEFT"
)

func (r OrganizationJoinResponseMemberStatus) IsKnown() bool {
	switch r {
	case OrganizationJoinResponseMemberStatusUserStatusUnspecified, OrganizationJoinResponseMemberStatusUserStatusActive, OrganizationJoinResponseMemberStatusUserStatusSuspended, OrganizationJoinResponseMemberStatusUserStatusLeft:
		return true
	}
	return false
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

type OrganizationNewParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Should other Accounts with the same domain be automatically invited to the
	// organization?
	InviteAccountsWithMatchingDomain param.Field[bool] `json:"inviteAccountsWithMatchingDomain"`
	// join_organization decides whether the Identity issuing this request joins the
	// org on creation
	JoinOrganization param.Field[bool] `json:"joinOrganization"`
	// name is the organization name
	Name param.Field[string] `json:"name"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationNewParamsConnectProtocolVersion float64

const (
	OrganizationNewParamsConnectProtocolVersion1 OrganizationNewParamsConnectProtocolVersion = 1
)

func (r OrganizationNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationGetParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// organization_id is the unique identifier of the Organization to retreive.
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationGetParamsConnectProtocolVersion float64

const (
	OrganizationGetParamsConnectProtocolVersion1 OrganizationGetParamsConnectProtocolVersion = 1
)

func (r OrganizationGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationUpdateParams struct {
	Body OrganizationUpdateParamsBodyUnion `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type OrganizationUpdateParamsBody struct {
	InviteDomains param.Field[interface{}] `json:"inviteDomains"`
	// name is the new name of the organization
	Name param.Field[string] `json:"name"`
}

func (r OrganizationUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationUpdateParamsBody) implementsOrganizationUpdateParamsBodyUnion() {}

// Satisfied by
// [OrganizationUpdateParamsBodyInviteDomainsIsTheDomainAllowlistOfTheOrganization],
// [OrganizationUpdateParamsBodyNameIsTheNewNameOfTheOrganization],
// [OrganizationUpdateParamsBody].
type OrganizationUpdateParamsBodyUnion interface {
	implementsOrganizationUpdateParamsBodyUnion()
}

type OrganizationUpdateParamsBodyInviteDomainsIsTheDomainAllowlistOfTheOrganization struct {
	// invite_domains is the domain allowlist of the organization
	InviteDomains param.Field[OrganizationUpdateParamsBodyInviteDomainsIsTheDomainAllowlistOfTheOrganizationInviteDomains] `json:"inviteDomains,required"`
}

func (r OrganizationUpdateParamsBodyInviteDomainsIsTheDomainAllowlistOfTheOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationUpdateParamsBodyInviteDomainsIsTheDomainAllowlistOfTheOrganization) implementsOrganizationUpdateParamsBodyUnion() {
}

// invite_domains is the domain allowlist of the organization
type OrganizationUpdateParamsBodyInviteDomainsIsTheDomainAllowlistOfTheOrganizationInviteDomains struct {
	// domains is the list of domains that are allowed to join the organization
	Domains param.Field[[]string] `json:"domains"`
}

func (r OrganizationUpdateParamsBodyInviteDomainsIsTheDomainAllowlistOfTheOrganizationInviteDomains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateParamsBodyNameIsTheNewNameOfTheOrganization struct {
	// name is the new name of the organization
	Name param.Field[string] `json:"name,required"`
}

func (r OrganizationUpdateParamsBodyNameIsTheNewNameOfTheOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationUpdateParamsBodyNameIsTheNewNameOfTheOrganization) implementsOrganizationUpdateParamsBodyUnion() {
}

// Define the version of the Connect protocol
type OrganizationUpdateParamsConnectProtocolVersion float64

const (
	OrganizationUpdateParamsConnectProtocolVersion1 OrganizationUpdateParamsConnectProtocolVersion = 1
)

func (r OrganizationUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[OrganizationListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[OrganizationListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[OrganizationListParamsConnect] `query:"connect"`
	Message param.Field[string]                        `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type OrganizationListParamsEncoding string

const (
	OrganizationListParamsEncodingProto OrganizationListParamsEncoding = "proto"
	OrganizationListParamsEncodingJson  OrganizationListParamsEncoding = "json"
)

func (r OrganizationListParamsEncoding) IsKnown() bool {
	switch r {
	case OrganizationListParamsEncodingProto, OrganizationListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type OrganizationListParamsConnectProtocolVersion float64

const (
	OrganizationListParamsConnectProtocolVersion1 OrganizationListParamsConnectProtocolVersion = 1
)

func (r OrganizationListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type OrganizationListParamsCompression string

const (
	OrganizationListParamsCompressionIdentity OrganizationListParamsCompression = "identity"
	OrganizationListParamsCompressionGzip     OrganizationListParamsCompression = "gzip"
	OrganizationListParamsCompressionBr       OrganizationListParamsCompression = "br"
)

func (r OrganizationListParamsCompression) IsKnown() bool {
	switch r {
	case OrganizationListParamsCompressionIdentity, OrganizationListParamsCompressionGzip, OrganizationListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type OrganizationListParamsConnect string

const (
	OrganizationListParamsConnectV1 OrganizationListParamsConnect = "v1"
)

func (r OrganizationListParamsConnect) IsKnown() bool {
	switch r {
	case OrganizationListParamsConnectV1:
		return true
	}
	return false
}

type OrganizationDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// organization_id is the ID of the organization to delete
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationDeleteParamsConnectProtocolVersion float64

const (
	OrganizationDeleteParamsConnectProtocolVersion1 OrganizationDeleteParamsConnectProtocolVersion = 1
)

func (r OrganizationDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationJoinParams struct {
	Body OrganizationJoinParamsBodyUnion `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationJoinParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationJoinParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type OrganizationJoinParamsBody struct {
	// invite_id is the unique identifier of the invite to join the organization.
	InviteID param.Field[string] `json:"inviteId" format:"uuid"`
	// organization_id is the unique identifier of the Organization to join.
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
}

func (r OrganizationJoinParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationJoinParamsBody) implementsOrganizationJoinParamsBodyUnion() {}

// Satisfied by
// [OrganizationJoinParamsBodyInviteIDIsTheUniqueIdentifierOfTheInviteToJoinTheOrganization],
// [OrganizationJoinParamsBodyOrganizationIDIsTheUniqueIdentifierOfTheOrganizationToJoin],
// [OrganizationJoinParamsBody].
type OrganizationJoinParamsBodyUnion interface {
	implementsOrganizationJoinParamsBodyUnion()
}

type OrganizationJoinParamsBodyInviteIDIsTheUniqueIdentifierOfTheInviteToJoinTheOrganization struct {
	// invite_id is the unique identifier of the invite to join the organization.
	InviteID param.Field[string] `json:"inviteId,required" format:"uuid"`
}

func (r OrganizationJoinParamsBodyInviteIDIsTheUniqueIdentifierOfTheInviteToJoinTheOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationJoinParamsBodyInviteIDIsTheUniqueIdentifierOfTheInviteToJoinTheOrganization) implementsOrganizationJoinParamsBodyUnion() {
}

type OrganizationJoinParamsBodyOrganizationIDIsTheUniqueIdentifierOfTheOrganizationToJoin struct {
	// organization_id is the unique identifier of the Organization to join.
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r OrganizationJoinParamsBodyOrganizationIDIsTheUniqueIdentifierOfTheOrganizationToJoin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrganizationJoinParamsBodyOrganizationIDIsTheUniqueIdentifierOfTheOrganizationToJoin) implementsOrganizationJoinParamsBodyUnion() {
}

// Define the version of the Connect protocol
type OrganizationJoinParamsConnectProtocolVersion float64

const (
	OrganizationJoinParamsConnectProtocolVersion1 OrganizationJoinParamsConnectProtocolVersion = 1
)

func (r OrganizationJoinParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationJoinParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationLeaveParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationLeaveParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	UserID                 param.Field[string]                                        `json:"userId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationLeaveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationLeaveParamsConnectProtocolVersion float64

const (
	OrganizationLeaveParamsConnectProtocolVersion1 OrganizationLeaveParamsConnectProtocolVersion = 1
)

func (r OrganizationLeaveParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationLeaveParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type OrganizationListMembersParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[OrganizationListMembersParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationListMembersParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[OrganizationListMembersParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[OrganizationListMembersParamsConnect] `query:"connect"`
	Message param.Field[string]                               `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [OrganizationListMembersParams]'s query parameters as
// `url.Values`.
func (r OrganizationListMembersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type OrganizationListMembersParamsEncoding string

const (
	OrganizationListMembersParamsEncodingProto OrganizationListMembersParamsEncoding = "proto"
	OrganizationListMembersParamsEncodingJson  OrganizationListMembersParamsEncoding = "json"
)

func (r OrganizationListMembersParamsEncoding) IsKnown() bool {
	switch r {
	case OrganizationListMembersParamsEncodingProto, OrganizationListMembersParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type OrganizationListMembersParamsConnectProtocolVersion float64

const (
	OrganizationListMembersParamsConnectProtocolVersion1 OrganizationListMembersParamsConnectProtocolVersion = 1
)

func (r OrganizationListMembersParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationListMembersParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type OrganizationListMembersParamsCompression string

const (
	OrganizationListMembersParamsCompressionIdentity OrganizationListMembersParamsCompression = "identity"
	OrganizationListMembersParamsCompressionGzip     OrganizationListMembersParamsCompression = "gzip"
	OrganizationListMembersParamsCompressionBr       OrganizationListMembersParamsCompression = "br"
)

func (r OrganizationListMembersParamsCompression) IsKnown() bool {
	switch r {
	case OrganizationListMembersParamsCompressionIdentity, OrganizationListMembersParamsCompressionGzip, OrganizationListMembersParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type OrganizationListMembersParamsConnect string

const (
	OrganizationListMembersParamsConnectV1 OrganizationListMembersParamsConnect = "v1"
)

func (r OrganizationListMembersParamsConnect) IsKnown() bool {
	switch r {
	case OrganizationListMembersParamsConnectV1:
		return true
	}
	return false
}

type OrganizationSetRoleParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[OrganizationSetRoleParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	OrganizationID         param.Field[string]                                          `json:"organizationId" format:"uuid"`
	Role                   param.Field[OrganizationSetRoleParamsRole]                   `json:"role"`
	UserID                 param.Field[string]                                          `json:"userId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r OrganizationSetRoleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type OrganizationSetRoleParamsConnectProtocolVersion float64

const (
	OrganizationSetRoleParamsConnectProtocolVersion1 OrganizationSetRoleParamsConnectProtocolVersion = 1
)

func (r OrganizationSetRoleParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case OrganizationSetRoleParamsConnectProtocolVersion1:
		return true
	}
	return false
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
