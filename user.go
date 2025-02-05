// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"time"

	"github.com/gitpod-io/flex-sdk-go/internal/apijson"
	"github.com/gitpod-io/flex-sdk-go/internal/param"
	"github.com/gitpod-io/flex-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/flex-sdk-go/option"
)

// UserService contains methods and other services that help with interacting with
// the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
	Pats    *UserPatService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.Pats = NewUserPatService(opts...)
	return
}

// GetAuthenticatedUser allows to retrieve the current user.
func (r *UserService) GetAuthenticatedUser(ctx context.Context, body UserGetAuthenticatedUserParams, opts ...option.RequestOption) (res *UserGetAuthenticatedUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.UserService/GetAuthenticatedUser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// SetSuspended sets the suspended state of the user.
func (r *UserService) SetSuspended(ctx context.Context, body UserSetSuspendedParams, opts ...option.RequestOption) (res *UserSetSuspendedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.UserService/SetSuspended"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UserGetAuthenticatedUserResponse struct {
	User UserGetAuthenticatedUserResponseUser `json:"user"`
	JSON userGetAuthenticatedUserResponseJSON `json:"-"`
}

// userGetAuthenticatedUserResponseJSON contains the JSON metadata for the struct
// [UserGetAuthenticatedUserResponse]
type userGetAuthenticatedUserResponseJSON struct {
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetAuthenticatedUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetAuthenticatedUserResponseJSON) RawJSON() string {
	return r.raw
}

type UserGetAuthenticatedUserResponseUser struct {
	// id is a UUID of the user
	ID string `json:"id" format:"uuid"`
	// avatar_url is a link to the user avatar
	AvatarURL string `json:"avatarUrl"`
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
	// name is the full name of the user
	Name string `json:"name"`
	// organization_id is the id of the organization this account is owned by.
	//
	// +optional if not set, this account is owned by the installation.
	OrganizationID string `json:"organizationId" format:"uuid"`
	// status is the status the user is in
	Status UserGetAuthenticatedUserResponseUserStatus `json:"status"`
	JSON   userGetAuthenticatedUserResponseUserJSON   `json:"-"`
}

// userGetAuthenticatedUserResponseUserJSON contains the JSON metadata for the
// struct [UserGetAuthenticatedUserResponseUser]
type userGetAuthenticatedUserResponseUserJSON struct {
	ID             apijson.Field
	AvatarURL      apijson.Field
	CreatedAt      apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserGetAuthenticatedUserResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetAuthenticatedUserResponseUserJSON) RawJSON() string {
	return r.raw
}

// status is the status the user is in
type UserGetAuthenticatedUserResponseUserStatus string

const (
	UserGetAuthenticatedUserResponseUserStatusUserStatusUnspecified UserGetAuthenticatedUserResponseUserStatus = "USER_STATUS_UNSPECIFIED"
	UserGetAuthenticatedUserResponseUserStatusUserStatusActive      UserGetAuthenticatedUserResponseUserStatus = "USER_STATUS_ACTIVE"
	UserGetAuthenticatedUserResponseUserStatusUserStatusSuspended   UserGetAuthenticatedUserResponseUserStatus = "USER_STATUS_SUSPENDED"
	UserGetAuthenticatedUserResponseUserStatusUserStatusLeft        UserGetAuthenticatedUserResponseUserStatus = "USER_STATUS_LEFT"
)

func (r UserGetAuthenticatedUserResponseUserStatus) IsKnown() bool {
	switch r {
	case UserGetAuthenticatedUserResponseUserStatusUserStatusUnspecified, UserGetAuthenticatedUserResponseUserStatusUserStatusActive, UserGetAuthenticatedUserResponseUserStatusUserStatusSuspended, UserGetAuthenticatedUserResponseUserStatusUserStatusLeft:
		return true
	}
	return false
}

type UserSetSuspendedResponse = interface{}

type UserGetAuthenticatedUserParams struct {
	Empty param.Field[bool] `json:"empty"`
}

func (r UserGetAuthenticatedUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetSuspendedParams struct {
	Suspended param.Field[bool]   `json:"suspended"`
	UserID    param.Field[string] `json:"userId" format:"uuid"`
}

func (r UserSetSuspendedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
