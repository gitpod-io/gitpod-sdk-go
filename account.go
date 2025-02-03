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

// AccountService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// GetAccount retrieves a single Account.
func (r *AccountService) Get(ctx context.Context, params AccountGetParams, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/GetAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// DeleteAccount deletes an Account. To Delete an Account, the Account must not be
// an active member of any Organization.
func (r *AccountService) Delete(ctx context.Context, params AccountDeleteParams, opts ...option.RequestOption) (res *AccountDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/DeleteAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetSSOLoginURL returns the URL to redirect the user to for SSO login.
func (r *AccountService) GetSSOLoginURL(ctx context.Context, params AccountGetSSOLoginURLParams, opts ...option.RequestOption) (res *AccountGetSSOLoginURLResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/GetSSOLoginURL"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// ListLoginProviders returns the list of login providers matching the provided
// filters.
func (r *AccountService) ListLoginProviders(ctx context.Context, params AccountListLoginProvidersParams, opts ...option.RequestOption) (res *AccountListLoginProvidersResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/ListLoginProviders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type AccountGetResponse struct {
	Account AccountGetResponseAccount `json:"account"`
	JSON    accountGetResponseJSON    `json:"-"`
}

// accountGetResponseJSON contains the JSON metadata for the struct
// [AccountGetResponse]
type accountGetResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetResponseAccount struct {
	ID        string `json:"id" format:"uuid"`
	AvatarURL string `json:"avatarUrl"`
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
	Email     string    `json:"email"`
	// This field can have the runtime type of
	// [[]AccountGetResponseAccountObjectJoinable].
	Joinables interface{} `json:"joinables"`
	// This field can have the runtime type of
	// [[]AccountGetResponseAccountObjectMembership].
	Memberships interface{} `json:"memberships"`
	Name        string      `json:"name"`
	// organization_id is the ID of the organization the account is owned by if it's
	// created through custom SSO
	OrganizationID string `json:"organizationId"`
	// public_email_provider is true if the email for the Account matches a known
	// public email provider
	PublicEmailProvider bool `json:"publicEmailProvider"`
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
	UpdatedAt time.Time                     `json:"updatedAt" format:"date-time"`
	JSON      accountGetResponseAccountJSON `json:"-"`
	union     AccountGetResponseAccountUnion
}

// accountGetResponseAccountJSON contains the JSON metadata for the struct
// [AccountGetResponseAccount]
type accountGetResponseAccountJSON struct {
	ID                  apijson.Field
	AvatarURL           apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	Joinables           apijson.Field
	Memberships         apijson.Field
	Name                apijson.Field
	OrganizationID      apijson.Field
	PublicEmailProvider apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r accountGetResponseAccountJSON) RawJSON() string {
	return r.raw
}

func (r *AccountGetResponseAccount) UnmarshalJSON(data []byte) (err error) {
	*r = AccountGetResponseAccount{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountGetResponseAccountUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [AccountGetResponseAccountObject],
// [AccountGetResponseAccountObject].
func (r AccountGetResponseAccount) AsUnion() AccountGetResponseAccountUnion {
	return r.union
}

// Union satisfied by [AccountGetResponseAccountObject] or
// [AccountGetResponseAccountObject].
type AccountGetResponseAccountUnion interface {
	implementsAccountGetResponseAccount()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountGetResponseAccountUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountGetResponseAccountObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountGetResponseAccountObject{}),
		},
	)
}

type AccountGetResponseAccountObject struct {
	// organization_id is the ID of the organization the account is owned by if it's
	// created through custom SSO
	OrganizationID string `json:"organizationId,required"`
	ID             string `json:"id" format:"uuid"`
	AvatarURL      string `json:"avatarUrl"`
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
	CreatedAt   time.Time                                   `json:"createdAt" format:"date-time"`
	Email       string                                      `json:"email"`
	Joinables   []AccountGetResponseAccountObjectJoinable   `json:"joinables"`
	Memberships []AccountGetResponseAccountObjectMembership `json:"memberships"`
	Name        string                                      `json:"name"`
	// public_email_provider is true if the email for the Account matches a known
	// public email provider
	PublicEmailProvider bool `json:"publicEmailProvider"`
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
	UpdatedAt time.Time                           `json:"updatedAt" format:"date-time"`
	JSON      accountGetResponseAccountObjectJSON `json:"-"`
}

// accountGetResponseAccountObjectJSON contains the JSON metadata for the struct
// [AccountGetResponseAccountObject]
type accountGetResponseAccountObjectJSON struct {
	OrganizationID      apijson.Field
	ID                  apijson.Field
	AvatarURL           apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	Joinables           apijson.Field
	Memberships         apijson.Field
	Name                apijson.Field
	PublicEmailProvider apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountGetResponseAccountObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseAccountObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountGetResponseAccountObject) implementsAccountGetResponseAccount() {}

type AccountGetResponseAccountObjectJoinable struct {
	// organization_id is the id of the organization the user can join
	OrganizationID string `json:"organizationId" format:"uuid"`
	// organization_member_count is the member count of the organization the user can
	// join
	OrganizationMemberCount int64 `json:"organizationMemberCount"`
	// organization_name is the name of the organization the user can join
	OrganizationName string                                      `json:"organizationName"`
	JSON             accountGetResponseAccountObjectJoinableJSON `json:"-"`
}

// accountGetResponseAccountObjectJoinableJSON contains the JSON metadata for the
// struct [AccountGetResponseAccountObjectJoinable]
type accountGetResponseAccountObjectJoinableJSON struct {
	OrganizationID          apijson.Field
	OrganizationMemberCount apijson.Field
	OrganizationName        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountGetResponseAccountObjectJoinable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseAccountObjectJoinableJSON) RawJSON() string {
	return r.raw
}

type AccountGetResponseAccountObjectMembership struct {
	// organization_id is the id of the organization the user is a member of
	OrganizationID string `json:"organizationId" format:"uuid"`
	// organization_name is the member count of the organization the user is a member
	// of
	OrganizationMemberCount int64 `json:"organizationMemberCount"`
	// organization_name is the name of the organization the user is a member of
	OrganizationName string `json:"organizationName"`
	// user_id is the ID the user has in the organization
	UserID string `json:"userId" format:"uuid"`
	// user_role is the role the user has in the organization
	UserRole AccountGetResponseAccountObjectMembershipsUserRole `json:"userRole"`
	JSON     accountGetResponseAccountObjectMembershipJSON      `json:"-"`
}

// accountGetResponseAccountObjectMembershipJSON contains the JSON metadata for the
// struct [AccountGetResponseAccountObjectMembership]
type accountGetResponseAccountObjectMembershipJSON struct {
	OrganizationID          apijson.Field
	OrganizationMemberCount apijson.Field
	OrganizationName        apijson.Field
	UserID                  apijson.Field
	UserRole                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountGetResponseAccountObjectMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseAccountObjectMembershipJSON) RawJSON() string {
	return r.raw
}

// user_role is the role the user has in the organization
type AccountGetResponseAccountObjectMembershipsUserRole string

const (
	AccountGetResponseAccountObjectMembershipsUserRoleOrganizationRoleUnspecified AccountGetResponseAccountObjectMembershipsUserRole = "ORGANIZATION_ROLE_UNSPECIFIED"
	AccountGetResponseAccountObjectMembershipsUserRoleOrganizationRoleAdmin       AccountGetResponseAccountObjectMembershipsUserRole = "ORGANIZATION_ROLE_ADMIN"
	AccountGetResponseAccountObjectMembershipsUserRoleOrganizationRoleMember      AccountGetResponseAccountObjectMembershipsUserRole = "ORGANIZATION_ROLE_MEMBER"
)

func (r AccountGetResponseAccountObjectMembershipsUserRole) IsKnown() bool {
	switch r {
	case AccountGetResponseAccountObjectMembershipsUserRoleOrganizationRoleUnspecified, AccountGetResponseAccountObjectMembershipsUserRoleOrganizationRoleAdmin, AccountGetResponseAccountObjectMembershipsUserRoleOrganizationRoleMember:
		return true
	}
	return false
}

type AccountDeleteResponse = interface{}

type AccountGetSSOLoginURLResponse struct {
	// login_url is the URL to redirect the user to for SSO login
	LoginURL string                            `json:"loginUrl"`
	JSON     accountGetSSOLoginURLResponseJSON `json:"-"`
}

// accountGetSSOLoginURLResponseJSON contains the JSON metadata for the struct
// [AccountGetSSOLoginURLResponse]
type accountGetSSOLoginURLResponseJSON struct {
	LoginURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetSSOLoginURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetSSOLoginURLResponseJSON) RawJSON() string {
	return r.raw
}

type AccountListLoginProvidersResponse struct {
	LoginProviders []AccountListLoginProvidersResponseLoginProvider `json:"loginProviders"`
	Pagination     AccountListLoginProvidersResponsePagination      `json:"pagination"`
	JSON           accountListLoginProvidersResponseJSON            `json:"-"`
}

// accountListLoginProvidersResponseJSON contains the JSON metadata for the struct
// [AccountListLoginProvidersResponse]
type accountListLoginProvidersResponseJSON struct {
	LoginProviders apijson.Field
	Pagination     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountListLoginProvidersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListLoginProvidersResponseJSON) RawJSON() string {
	return r.raw
}

type AccountListLoginProvidersResponseLoginProvider struct {
	// login_url is the URL to redirect the browser agent to for login
	LoginURL string `json:"loginUrl"`
	// provider is the provider used by this login method, e.g. "github", "google",
	// "custom"
	Provider string                                             `json:"provider"`
	JSON     accountListLoginProvidersResponseLoginProviderJSON `json:"-"`
}

// accountListLoginProvidersResponseLoginProviderJSON contains the JSON metadata
// for the struct [AccountListLoginProvidersResponseLoginProvider]
type accountListLoginProvidersResponseLoginProviderJSON struct {
	LoginURL    apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListLoginProvidersResponseLoginProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListLoginProvidersResponseLoginProviderJSON) RawJSON() string {
	return r.raw
}

type AccountListLoginProvidersResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                          `json:"nextToken"`
	JSON      accountListLoginProvidersResponsePaginationJSON `json:"-"`
}

// accountListLoginProvidersResponsePaginationJSON contains the JSON metadata for
// the struct [AccountListLoginProvidersResponsePagination]
type accountListLoginProvidersResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListLoginProvidersResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListLoginProvidersResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type AccountGetParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[AccountGetParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[AccountGetParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[AccountGetParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[AccountGetParamsConnect] `query:"connect"`
	Message param.Field[string]                  `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [AccountGetParams]'s query parameters as `url.Values`.
func (r AccountGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type AccountGetParamsEncoding string

const (
	AccountGetParamsEncodingProto AccountGetParamsEncoding = "proto"
	AccountGetParamsEncodingJson  AccountGetParamsEncoding = "json"
)

func (r AccountGetParamsEncoding) IsKnown() bool {
	switch r {
	case AccountGetParamsEncodingProto, AccountGetParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type AccountGetParamsConnectProtocolVersion float64

const (
	AccountGetParamsConnectProtocolVersion1 AccountGetParamsConnectProtocolVersion = 1
)

func (r AccountGetParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case AccountGetParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type AccountGetParamsCompression string

const (
	AccountGetParamsCompressionIdentity AccountGetParamsCompression = "identity"
	AccountGetParamsCompressionGzip     AccountGetParamsCompression = "gzip"
	AccountGetParamsCompressionBr       AccountGetParamsCompression = "br"
)

func (r AccountGetParamsCompression) IsKnown() bool {
	switch r {
	case AccountGetParamsCompressionIdentity, AccountGetParamsCompressionGzip, AccountGetParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type AccountGetParamsConnect string

const (
	AccountGetParamsConnectV1 AccountGetParamsConnect = "v1"
)

func (r AccountGetParamsConnect) IsKnown() bool {
	switch r {
	case AccountGetParamsConnectV1:
		return true
	}
	return false
}

type AccountDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[AccountDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	AccountID              param.Field[string]                                    `json:"accountId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r AccountDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type AccountDeleteParamsConnectProtocolVersion float64

const (
	AccountDeleteParamsConnectProtocolVersion1 AccountDeleteParamsConnectProtocolVersion = 1
)

func (r AccountDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case AccountDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type AccountGetSSOLoginURLParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[AccountGetSSOLoginURLParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[AccountGetSSOLoginURLParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[AccountGetSSOLoginURLParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[AccountGetSSOLoginURLParamsConnect] `query:"connect"`
	Message param.Field[string]                             `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [AccountGetSSOLoginURLParams]'s query parameters as
// `url.Values`.
func (r AccountGetSSOLoginURLParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type AccountGetSSOLoginURLParamsEncoding string

const (
	AccountGetSSOLoginURLParamsEncodingProto AccountGetSSOLoginURLParamsEncoding = "proto"
	AccountGetSSOLoginURLParamsEncodingJson  AccountGetSSOLoginURLParamsEncoding = "json"
)

func (r AccountGetSSOLoginURLParamsEncoding) IsKnown() bool {
	switch r {
	case AccountGetSSOLoginURLParamsEncodingProto, AccountGetSSOLoginURLParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type AccountGetSSOLoginURLParamsConnectProtocolVersion float64

const (
	AccountGetSSOLoginURLParamsConnectProtocolVersion1 AccountGetSSOLoginURLParamsConnectProtocolVersion = 1
)

func (r AccountGetSSOLoginURLParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case AccountGetSSOLoginURLParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type AccountGetSSOLoginURLParamsCompression string

const (
	AccountGetSSOLoginURLParamsCompressionIdentity AccountGetSSOLoginURLParamsCompression = "identity"
	AccountGetSSOLoginURLParamsCompressionGzip     AccountGetSSOLoginURLParamsCompression = "gzip"
	AccountGetSSOLoginURLParamsCompressionBr       AccountGetSSOLoginURLParamsCompression = "br"
)

func (r AccountGetSSOLoginURLParamsCompression) IsKnown() bool {
	switch r {
	case AccountGetSSOLoginURLParamsCompressionIdentity, AccountGetSSOLoginURLParamsCompressionGzip, AccountGetSSOLoginURLParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type AccountGetSSOLoginURLParamsConnect string

const (
	AccountGetSSOLoginURLParamsConnectV1 AccountGetSSOLoginURLParamsConnect = "v1"
)

func (r AccountGetSSOLoginURLParamsConnect) IsKnown() bool {
	switch r {
	case AccountGetSSOLoginURLParamsConnectV1:
		return true
	}
	return false
}

type AccountListLoginProvidersParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[AccountListLoginProvidersParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[AccountListLoginProvidersParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[AccountListLoginProvidersParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[AccountListLoginProvidersParamsConnect] `query:"connect"`
	Message param.Field[string]                                 `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [AccountListLoginProvidersParams]'s query parameters as
// `url.Values`.
func (r AccountListLoginProvidersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type AccountListLoginProvidersParamsEncoding string

const (
	AccountListLoginProvidersParamsEncodingProto AccountListLoginProvidersParamsEncoding = "proto"
	AccountListLoginProvidersParamsEncodingJson  AccountListLoginProvidersParamsEncoding = "json"
)

func (r AccountListLoginProvidersParamsEncoding) IsKnown() bool {
	switch r {
	case AccountListLoginProvidersParamsEncodingProto, AccountListLoginProvidersParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type AccountListLoginProvidersParamsConnectProtocolVersion float64

const (
	AccountListLoginProvidersParamsConnectProtocolVersion1 AccountListLoginProvidersParamsConnectProtocolVersion = 1
)

func (r AccountListLoginProvidersParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case AccountListLoginProvidersParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type AccountListLoginProvidersParamsCompression string

const (
	AccountListLoginProvidersParamsCompressionIdentity AccountListLoginProvidersParamsCompression = "identity"
	AccountListLoginProvidersParamsCompressionGzip     AccountListLoginProvidersParamsCompression = "gzip"
	AccountListLoginProvidersParamsCompressionBr       AccountListLoginProvidersParamsCompression = "br"
)

func (r AccountListLoginProvidersParamsCompression) IsKnown() bool {
	switch r {
	case AccountListLoginProvidersParamsCompressionIdentity, AccountListLoginProvidersParamsCompressionGzip, AccountListLoginProvidersParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type AccountListLoginProvidersParamsConnect string

const (
	AccountListLoginProvidersParamsConnectV1 AccountListLoginProvidersParamsConnect = "v1"
)

func (r AccountListLoginProvidersParamsConnect) IsKnown() bool {
	switch r {
	case AccountListLoginProvidersParamsConnectV1:
		return true
	}
	return false
}
