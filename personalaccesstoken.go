// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// PersonalAccessTokenService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPersonalAccessTokenService] method instead.
type PersonalAccessTokenService struct {
	Options []option.RequestOption
}

// NewPersonalAccessTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPersonalAccessTokenService(opts ...option.RequestOption) (r *PersonalAccessTokenService) {
	r = &PersonalAccessTokenService{}
	r.Options = opts
	return
}

// ListPersonalAccessTokens
func (r *PersonalAccessTokenService) List(ctx context.Context, params PersonalAccessTokenListParams, opts ...option.RequestOption) (res *PersonalAccessTokenListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.UserService/ListPersonalAccessTokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// DeletePersonalAccessToken
func (r *PersonalAccessTokenService) Delete(ctx context.Context, params PersonalAccessTokenDeleteParams, opts ...option.RequestOption) (res *PersonalAccessTokenDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.UserService/DeletePersonalAccessToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type PersonalAccessTokenListResponse struct {
	Pagination           PersonalAccessTokenListResponsePagination            `json:"pagination"`
	PersonalAccessTokens []PersonalAccessTokenListResponsePersonalAccessToken `json:"personalAccessTokens"`
	JSON                 personalAccessTokenListResponseJSON                  `json:"-"`
}

// personalAccessTokenListResponseJSON contains the JSON metadata for the struct
// [PersonalAccessTokenListResponse]
type personalAccessTokenListResponseJSON struct {
	Pagination           apijson.Field
	PersonalAccessTokens apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PersonalAccessTokenListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r personalAccessTokenListResponseJSON) RawJSON() string {
	return r.raw
}

type PersonalAccessTokenListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
	NextToken string                                        `json:"nextToken"`
	JSON      personalAccessTokenListResponsePaginationJSON `json:"-"`
}

// personalAccessTokenListResponsePaginationJSON contains the JSON metadata for the
// struct [PersonalAccessTokenListResponsePagination]
type personalAccessTokenListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PersonalAccessTokenListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r personalAccessTokenListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type PersonalAccessTokenListResponsePersonalAccessToken struct {
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
	CreatedAt   time.Time                                                  `json:"createdAt" format:"date-time"`
	Creator     PersonalAccessTokenListResponsePersonalAccessTokensCreator `json:"creator"`
	Description string                                                     `json:"description"`
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
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
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
	LastUsed time.Time                                              `json:"lastUsed" format:"date-time"`
	UserID   string                                                 `json:"userId" format:"uuid"`
	JSON     personalAccessTokenListResponsePersonalAccessTokenJSON `json:"-"`
}

// personalAccessTokenListResponsePersonalAccessTokenJSON contains the JSON
// metadata for the struct [PersonalAccessTokenListResponsePersonalAccessToken]
type personalAccessTokenListResponsePersonalAccessTokenJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Description apijson.Field
	ExpiresAt   apijson.Field
	LastUsed    apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PersonalAccessTokenListResponsePersonalAccessToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r personalAccessTokenListResponsePersonalAccessTokenJSON) RawJSON() string {
	return r.raw
}

type PersonalAccessTokenListResponsePersonalAccessTokensCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipal `json:"principal"`
	JSON      personalAccessTokenListResponsePersonalAccessTokensCreatorJSON      `json:"-"`
}

// personalAccessTokenListResponsePersonalAccessTokensCreatorJSON contains the JSON
// metadata for the struct
// [PersonalAccessTokenListResponsePersonalAccessTokensCreator]
type personalAccessTokenListResponsePersonalAccessTokensCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PersonalAccessTokenListResponsePersonalAccessTokensCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r personalAccessTokenListResponsePersonalAccessTokensCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipal string

const (
	PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalUnspecified    PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalAccount        PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalUser           PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipal = "PRINCIPAL_USER"
	PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalRunner         PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipal = "PRINCIPAL_RUNNER"
	PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalEnvironment    PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalServiceAccount PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipal) IsKnown() bool {
	switch r {
	case PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalUnspecified, PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalAccount, PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalUser, PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalRunner, PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalEnvironment, PersonalAccessTokenListResponsePersonalAccessTokensCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type PersonalAccessTokenDeleteResponse = interface{}

type PersonalAccessTokenListParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[PersonalAccessTokenListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	Filter                 param.Field[PersonalAccessTokenListParamsFilter]                 `json:"filter"`
	Pagination             param.Field[PersonalAccessTokenListParamsPagination]             `json:"pagination"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r PersonalAccessTokenListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type PersonalAccessTokenListParamsConnectProtocolVersion float64

const (
	PersonalAccessTokenListParamsConnectProtocolVersion1 PersonalAccessTokenListParamsConnectProtocolVersion = 1
)

func (r PersonalAccessTokenListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case PersonalAccessTokenListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type PersonalAccessTokenListParamsFilter struct {
	// creator_ids filters the response to only Environments created by specified
	// members
	UserIDs param.Field[[]string] `json:"userIds" format:"uuid"`
}

func (r PersonalAccessTokenListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PersonalAccessTokenListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r PersonalAccessTokenListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PersonalAccessTokenDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[PersonalAccessTokenDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	PersonalAccessTokenID  param.Field[string]                                                `json:"personalAccessTokenId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r PersonalAccessTokenDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type PersonalAccessTokenDeleteParamsConnectProtocolVersion float64

const (
	PersonalAccessTokenDeleteParamsConnectProtocolVersion1 PersonalAccessTokenDeleteParamsConnectProtocolVersion = 1
)

func (r PersonalAccessTokenDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case PersonalAccessTokenDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}
