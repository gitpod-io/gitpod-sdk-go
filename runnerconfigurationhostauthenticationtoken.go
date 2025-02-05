// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/apiquery"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
	"github.com/stainless-sdks/gitpod-go/packages/pagination"
)

// RunnerConfigurationHostAuthenticationTokenService contains methods and other
// services that help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunnerConfigurationHostAuthenticationTokenService] method instead.
type RunnerConfigurationHostAuthenticationTokenService struct {
	Options []option.RequestOption
}

// NewRunnerConfigurationHostAuthenticationTokenService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewRunnerConfigurationHostAuthenticationTokenService(opts ...option.RequestOption) (r *RunnerConfigurationHostAuthenticationTokenService) {
	r = &RunnerConfigurationHostAuthenticationTokenService{}
	r.Options = opts
	return
}

// CreateHostAuthenticationToken
func (r *RunnerConfigurationHostAuthenticationTokenService) New(ctx context.Context, body RunnerConfigurationHostAuthenticationTokenNewParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/CreateHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetHostAuthenticationToken
func (r *RunnerConfigurationHostAuthenticationTokenService) Get(ctx context.Context, body RunnerConfigurationHostAuthenticationTokenGetParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/GetHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateHostAuthenticationToken
func (r *RunnerConfigurationHostAuthenticationTokenService) Update(ctx context.Context, body RunnerConfigurationHostAuthenticationTokenUpdateParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/UpdateHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListHostAuthenticationTokens
func (r *RunnerConfigurationHostAuthenticationTokenService) List(ctx context.Context, params RunnerConfigurationHostAuthenticationTokenListParams, opts ...option.RequestOption) (res *pagination.TokensPage[RunnerConfigurationHostAuthenticationTokenListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.RunnerConfigurationService/ListHostAuthenticationTokens"
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

// ListHostAuthenticationTokens
func (r *RunnerConfigurationHostAuthenticationTokenService) ListAutoPaging(ctx context.Context, params RunnerConfigurationHostAuthenticationTokenListParams, opts ...option.RequestOption) *pagination.TokensPageAutoPager[RunnerConfigurationHostAuthenticationTokenListResponse] {
	return pagination.NewTokensPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteHostAuthenticationToken
func (r *RunnerConfigurationHostAuthenticationTokenService) Delete(ctx context.Context, body RunnerConfigurationHostAuthenticationTokenDeleteParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.RunnerConfigurationService/DeleteHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RunnerConfigurationHostAuthenticationTokenNewResponse struct {
	Token RunnerConfigurationHostAuthenticationTokenNewResponseToken `json:"token"`
	JSON  runnerConfigurationHostAuthenticationTokenNewResponseJSON  `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenNewResponseJSON contains the JSON
// metadata for the struct [RunnerConfigurationHostAuthenticationTokenNewResponse]
type runnerConfigurationHostAuthenticationTokenNewResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenNewResponseToken struct {
	ID string `json:"id"`
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
	ExpiresAt time.Time                                                        `json:"expiresAt" format:"date-time"`
	Host      string                                                           `json:"host"`
	RunnerID  string                                                           `json:"runnerId"`
	Source    RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource `json:"source"`
	UserID    string                                                           `json:"userId"`
	JSON      runnerConfigurationHostAuthenticationTokenNewResponseTokenJSON   `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenNewResponseTokenJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationHostAuthenticationTokenNewResponseToken]
type runnerConfigurationHostAuthenticationTokenNewResponseTokenJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	Host        apijson.Field
	RunnerID    apijson.Field
	Source      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenNewResponseToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenNewResponseTokenJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource string

const (
	RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourceUnspecified RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_UNSPECIFIED"
	RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourceOAuth       RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_OAUTH"
	RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourcePat         RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_PAT"
)

func (r RunnerConfigurationHostAuthenticationTokenNewResponseTokenSource) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourceUnspecified, RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourceOAuth, RunnerConfigurationHostAuthenticationTokenNewResponseTokenSourceHostAuthenticationTokenSourcePat:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenGetResponse struct {
	Token RunnerConfigurationHostAuthenticationTokenGetResponseToken `json:"token"`
	JSON  runnerConfigurationHostAuthenticationTokenGetResponseJSON  `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenGetResponseJSON contains the JSON
// metadata for the struct [RunnerConfigurationHostAuthenticationTokenGetResponse]
type runnerConfigurationHostAuthenticationTokenGetResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenGetResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenGetResponseToken struct {
	ID string `json:"id"`
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
	ExpiresAt time.Time                                                        `json:"expiresAt" format:"date-time"`
	Host      string                                                           `json:"host"`
	RunnerID  string                                                           `json:"runnerId"`
	Source    RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource `json:"source"`
	UserID    string                                                           `json:"userId"`
	JSON      runnerConfigurationHostAuthenticationTokenGetResponseTokenJSON   `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenGetResponseTokenJSON contains the JSON
// metadata for the struct
// [RunnerConfigurationHostAuthenticationTokenGetResponseToken]
type runnerConfigurationHostAuthenticationTokenGetResponseTokenJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	Host        apijson.Field
	RunnerID    apijson.Field
	Source      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenGetResponseToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenGetResponseTokenJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource string

const (
	RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourceUnspecified RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_UNSPECIFIED"
	RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourceOAuth       RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_OAUTH"
	RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourcePat         RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_PAT"
)

func (r RunnerConfigurationHostAuthenticationTokenGetResponseTokenSource) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourceUnspecified, RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourceOAuth, RunnerConfigurationHostAuthenticationTokenGetResponseTokenSourceHostAuthenticationTokenSourcePat:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenUpdateResponse = interface{}

type RunnerConfigurationHostAuthenticationTokenListResponse struct {
	ID string `json:"id"`
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
	ExpiresAt time.Time                                                    `json:"expiresAt" format:"date-time"`
	Host      string                                                       `json:"host"`
	RunnerID  string                                                       `json:"runnerId"`
	Source    RunnerConfigurationHostAuthenticationTokenListResponseSource `json:"source"`
	UserID    string                                                       `json:"userId"`
	JSON      runnerConfigurationHostAuthenticationTokenListResponseJSON   `json:"-"`
}

// runnerConfigurationHostAuthenticationTokenListResponseJSON contains the JSON
// metadata for the struct [RunnerConfigurationHostAuthenticationTokenListResponse]
type runnerConfigurationHostAuthenticationTokenListResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	Host        apijson.Field
	RunnerID    apijson.Field
	Source      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnerConfigurationHostAuthenticationTokenListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnerConfigurationHostAuthenticationTokenListResponseJSON) RawJSON() string {
	return r.raw
}

type RunnerConfigurationHostAuthenticationTokenListResponseSource string

const (
	RunnerConfigurationHostAuthenticationTokenListResponseSourceHostAuthenticationTokenSourceUnspecified RunnerConfigurationHostAuthenticationTokenListResponseSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_UNSPECIFIED"
	RunnerConfigurationHostAuthenticationTokenListResponseSourceHostAuthenticationTokenSourceOAuth       RunnerConfigurationHostAuthenticationTokenListResponseSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_OAUTH"
	RunnerConfigurationHostAuthenticationTokenListResponseSourceHostAuthenticationTokenSourcePat         RunnerConfigurationHostAuthenticationTokenListResponseSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_PAT"
)

func (r RunnerConfigurationHostAuthenticationTokenListResponseSource) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenListResponseSourceHostAuthenticationTokenSourceUnspecified, RunnerConfigurationHostAuthenticationTokenListResponseSourceHostAuthenticationTokenSourceOAuth, RunnerConfigurationHostAuthenticationTokenListResponseSourceHostAuthenticationTokenSourcePat:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenDeleteResponse = interface{}

type RunnerConfigurationHostAuthenticationTokenNewParams struct {
	Token param.Field[string] `json:"token"`
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
	ExpiresAt    param.Field[time.Time]                                                 `json:"expiresAt" format:"date-time"`
	Host         param.Field[string]                                                    `json:"host"`
	RefreshToken param.Field[string]                                                    `json:"refreshToken"`
	RunnerID     param.Field[string]                                                    `json:"runnerId" format:"uuid"`
	Source       param.Field[RunnerConfigurationHostAuthenticationTokenNewParamsSource] `json:"source"`
	UserID       param.Field[string]                                                    `json:"userId" format:"uuid"`
}

func (r RunnerConfigurationHostAuthenticationTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationHostAuthenticationTokenNewParamsSource string

const (
	RunnerConfigurationHostAuthenticationTokenNewParamsSourceHostAuthenticationTokenSourceUnspecified RunnerConfigurationHostAuthenticationTokenNewParamsSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_UNSPECIFIED"
	RunnerConfigurationHostAuthenticationTokenNewParamsSourceHostAuthenticationTokenSourceOAuth       RunnerConfigurationHostAuthenticationTokenNewParamsSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_OAUTH"
	RunnerConfigurationHostAuthenticationTokenNewParamsSourceHostAuthenticationTokenSourcePat         RunnerConfigurationHostAuthenticationTokenNewParamsSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_PAT"
)

func (r RunnerConfigurationHostAuthenticationTokenNewParamsSource) IsKnown() bool {
	switch r {
	case RunnerConfigurationHostAuthenticationTokenNewParamsSourceHostAuthenticationTokenSourceUnspecified, RunnerConfigurationHostAuthenticationTokenNewParamsSourceHostAuthenticationTokenSourceOAuth, RunnerConfigurationHostAuthenticationTokenNewParamsSourceHostAuthenticationTokenSourcePat:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenGetParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r RunnerConfigurationHostAuthenticationTokenGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationHostAuthenticationTokenUpdateParams struct {
	Body RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyUnion `json:"body,required"`
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RunnerConfigurationHostAuthenticationTokenUpdateParamsBody struct {
	Token param.Field[string] `json:"token"`
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
	ExpiresAt    param.Field[time.Time] `json:"expiresAt" format:"date-time"`
	RefreshToken param.Field[string]    `json:"refreshToken"`
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParamsBody) implementsRunnerConfigurationHostAuthenticationTokenUpdateParamsBodyUnion() {
}

// Satisfied by
// [RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyExpiresAt],
// [RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyRefreshToken],
// [RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyToken],
// [RunnerConfigurationHostAuthenticationTokenUpdateParamsBody].
type RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyUnion interface {
	implementsRunnerConfigurationHostAuthenticationTokenUpdateParamsBodyUnion()
}

type RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyExpiresAt struct {
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
	ExpiresAt param.Field[time.Time] `json:"expiresAt,required" format:"date-time"`
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyExpiresAt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyExpiresAt) implementsRunnerConfigurationHostAuthenticationTokenUpdateParamsBodyUnion() {
}

type RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyRefreshToken struct {
	RefreshToken param.Field[string] `json:"refreshToken,required"`
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyRefreshToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyRefreshToken) implementsRunnerConfigurationHostAuthenticationTokenUpdateParamsBodyUnion() {
}

type RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyToken struct {
	Token param.Field[string] `json:"token,required"`
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParamsBodyToken) implementsRunnerConfigurationHostAuthenticationTokenUpdateParamsBodyUnion() {
}

type RunnerConfigurationHostAuthenticationTokenListParams struct {
	Token      param.Field[string]                                                          `query:"token"`
	PageSize   param.Field[int64]                                                           `query:"pageSize"`
	Filter     param.Field[RunnerConfigurationHostAuthenticationTokenListParamsFilterUnion] `json:"filter"`
	Pagination param.Field[RunnerConfigurationHostAuthenticationTokenListParamsPagination]  `json:"pagination"`
}

func (r RunnerConfigurationHostAuthenticationTokenListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [RunnerConfigurationHostAuthenticationTokenListParams]'s
// query parameters as `url.Values`.
func (r RunnerConfigurationHostAuthenticationTokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RunnerConfigurationHostAuthenticationTokenListParamsFilter struct {
	RunnerID param.Field[string] `json:"runnerId" format:"uuid"`
	UserID   param.Field[string] `json:"userId" format:"uuid"`
}

func (r RunnerConfigurationHostAuthenticationTokenListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationHostAuthenticationTokenListParamsFilter) implementsRunnerConfigurationHostAuthenticationTokenListParamsFilterUnion() {
}

// Satisfied by
// [RunnerConfigurationHostAuthenticationTokenListParamsFilterRunnerID],
// [RunnerConfigurationHostAuthenticationTokenListParamsFilterUserID],
// [RunnerConfigurationHostAuthenticationTokenListParamsFilter].
type RunnerConfigurationHostAuthenticationTokenListParamsFilterUnion interface {
	implementsRunnerConfigurationHostAuthenticationTokenListParamsFilterUnion()
}

type RunnerConfigurationHostAuthenticationTokenListParamsFilterRunnerID struct {
	RunnerID param.Field[string] `json:"runnerId,required" format:"uuid"`
}

func (r RunnerConfigurationHostAuthenticationTokenListParamsFilterRunnerID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationHostAuthenticationTokenListParamsFilterRunnerID) implementsRunnerConfigurationHostAuthenticationTokenListParamsFilterUnion() {
}

type RunnerConfigurationHostAuthenticationTokenListParamsFilterUserID struct {
	UserID param.Field[string] `json:"userId,required" format:"uuid"`
}

func (r RunnerConfigurationHostAuthenticationTokenListParamsFilterUserID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RunnerConfigurationHostAuthenticationTokenListParamsFilterUserID) implementsRunnerConfigurationHostAuthenticationTokenListParamsFilterUnion() {
}

type RunnerConfigurationHostAuthenticationTokenListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r RunnerConfigurationHostAuthenticationTokenListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationHostAuthenticationTokenDeleteParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r RunnerConfigurationHostAuthenticationTokenDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
