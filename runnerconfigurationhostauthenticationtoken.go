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
	"github.com/gitpod-io/gitpod-sdk-go/shared"
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

// Creates a new authentication token for accessing remote hosts.
//
// Use this method to:
//
// - Set up SCM authentication
// - Configure OAuth credentials
// - Manage PAT tokens
//
// ### Examples
//
// - Create OAuth token:
//
//	Creates a new OAuth-based authentication token.
//
//	```yaml
//	runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	userId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	host: "github.com"
//	token: "gho_xxxxxxxxxxxx"
//	source: HOST_AUTHENTICATION_TOKEN_SOURCE_OAUTH
//	expiresAt: "2024-12-31T23:59:59Z"
//	refreshToken: "ghr_xxxxxxxxxxxx"
//	```
func (r *RunnerConfigurationHostAuthenticationTokenService) New(ctx context.Context, body RunnerConfigurationHostAuthenticationTokenNewParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerConfigurationService/CreateHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets details about a specific host authentication token.
//
// Use this method to:
//
// - View token information
// - Check token expiration
// - Verify token validity
//
// ### Examples
//
// - Get token details:
//
//	Retrieves information about a specific token.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *RunnerConfigurationHostAuthenticationTokenService) Get(ctx context.Context, body RunnerConfigurationHostAuthenticationTokenGetParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerConfigurationService/GetHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing host authentication token.
//
// Use this method to:
//
// - Refresh token values
// - Update expiration
// - Modify token settings
//
// ### Examples
//
// - Update token:
//
//	Updates token value and expiration.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	token: "gho_xxxxxxxxxxxx"
//	expiresAt: "2024-12-31T23:59:59Z"
//	refreshToken: "ghr_xxxxxxxxxxxx"
//	```
func (r *RunnerConfigurationHostAuthenticationTokenService) Update(ctx context.Context, body RunnerConfigurationHostAuthenticationTokenUpdateParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerConfigurationService/UpdateHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists host authentication tokens with optional filtering.
//
// Use this method to:
//
// - View all tokens
// - Filter by runner or user
// - Monitor token status
//
// ### Examples
//
// - List all tokens:
//
//	Shows all tokens with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter by runner:
//
//	Lists tokens for a specific runner.
//
//	```yaml
//	filter:
//	  runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	pagination:
//	  pageSize: 20
//	```
func (r *RunnerConfigurationHostAuthenticationTokenService) List(ctx context.Context, params RunnerConfigurationHostAuthenticationTokenListParams, opts ...option.RequestOption) (res *pagination.TokensPage[HostAuthenticationToken], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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

// Lists host authentication tokens with optional filtering.
//
// Use this method to:
//
// - View all tokens
// - Filter by runner or user
// - Monitor token status
//
// ### Examples
//
// - List all tokens:
//
//	Shows all tokens with pagination.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter by runner:
//
//	Lists tokens for a specific runner.
//
//	```yaml
//	filter:
//	  runnerId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	pagination:
//	  pageSize: 20
//	```
func (r *RunnerConfigurationHostAuthenticationTokenService) ListAutoPaging(ctx context.Context, params RunnerConfigurationHostAuthenticationTokenListParams, opts ...option.RequestOption) *pagination.TokensPageAutoPager[HostAuthenticationToken] {
	return pagination.NewTokensPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a host authentication token.
//
// Use this method to:
//
// - Remove unused tokens
// - Revoke access
// - Clean up expired tokens
//
// ### Examples
//
// - Delete token:
//
//	Permanently removes a token.
//
//	```yaml
//	id: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *RunnerConfigurationHostAuthenticationTokenService) Delete(ctx context.Context, body RunnerConfigurationHostAuthenticationTokenDeleteParams, opts ...option.RequestOption) (res *RunnerConfigurationHostAuthenticationTokenDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.RunnerConfigurationService/DeleteHostAuthenticationToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type HostAuthenticationToken struct {
	ID string `json:"id,required"`
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
	Host      string    `json:"host"`
	// links to integration instance
	IntegrationID string `json:"integrationId"`
	RunnerID      string `json:"runnerId"`
	// token permissions
	Scopes []string `json:"scopes"`
	// auth_type
	Source HostAuthenticationTokenSource `json:"source"`
	// Subject identifies the principal (user or service account) for the token Note:
	// actual token and refresh_token values are retrieved via
	// GetHostAuthenticationTokenValue API
	Subject shared.Subject `json:"subject"`
	// Deprecated: Use principal_id and principal_type instead principal (user)
	//
	// Deprecated: deprecated
	UserID string                      `json:"userId"`
	JSON   hostAuthenticationTokenJSON `json:"-"`
}

// hostAuthenticationTokenJSON contains the JSON metadata for the struct
// [HostAuthenticationToken]
type hostAuthenticationTokenJSON struct {
	ID            apijson.Field
	ExpiresAt     apijson.Field
	Host          apijson.Field
	IntegrationID apijson.Field
	RunnerID      apijson.Field
	Scopes        apijson.Field
	Source        apijson.Field
	Subject       apijson.Field
	UserID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *HostAuthenticationToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostAuthenticationTokenJSON) RawJSON() string {
	return r.raw
}

type HostAuthenticationTokenSource string

const (
	HostAuthenticationTokenSourceUnspecified HostAuthenticationTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_UNSPECIFIED"
	HostAuthenticationTokenSourceOAuth       HostAuthenticationTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_OAUTH"
	HostAuthenticationTokenSourcePat         HostAuthenticationTokenSource = "HOST_AUTHENTICATION_TOKEN_SOURCE_PAT"
)

func (r HostAuthenticationTokenSource) IsKnown() bool {
	switch r {
	case HostAuthenticationTokenSourceUnspecified, HostAuthenticationTokenSourceOAuth, HostAuthenticationTokenSourcePat:
		return true
	}
	return false
}

type RunnerConfigurationHostAuthenticationTokenNewResponse struct {
	Token HostAuthenticationToken                                   `json:"token,required"`
	JSON  runnerConfigurationHostAuthenticationTokenNewResponseJSON `json:"-"`
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

type RunnerConfigurationHostAuthenticationTokenGetResponse struct {
	Token HostAuthenticationToken                                   `json:"token,required"`
	JSON  runnerConfigurationHostAuthenticationTokenGetResponseJSON `json:"-"`
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

type RunnerConfigurationHostAuthenticationTokenUpdateResponse = interface{}

type RunnerConfigurationHostAuthenticationTokenDeleteResponse = interface{}

type RunnerConfigurationHostAuthenticationTokenNewParams struct {
	// stored encrypted, retrieved via GetHostAuthenticationTokenValue
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
	ExpiresAt     param.Field[time.Time] `json:"expiresAt" format:"date-time"`
	Host          param.Field[string]    `json:"host"`
	IntegrationID param.Field[string]    `json:"integrationId" format:"uuid"`
	// stored encrypted, retrieved via GetHostAuthenticationTokenValue
	RefreshToken param.Field[string] `json:"refreshToken"`
	RunnerID     param.Field[string] `json:"runnerId" format:"uuid"`
	// Maximum 100 scopes allowed (101 for validation purposes)
	Scopes param.Field[[]string]                      `json:"scopes"`
	Source param.Field[HostAuthenticationTokenSource] `json:"source"`
	// Subject identifies the principal (user or service account) for the token
	Subject param.Field[shared.SubjectParam] `json:"subject"`
	// Deprecated: Use principal_id and principal_type instead
	UserID param.Field[string] `json:"userId" format:"uuid"`
}

func (r RunnerConfigurationHostAuthenticationTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationHostAuthenticationTokenGetParams struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r RunnerConfigurationHostAuthenticationTokenGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationHostAuthenticationTokenUpdateParams struct {
	ID    param.Field[string] `json:"id" format:"uuid"`
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
	Scopes       param.Field[[]string]  `json:"scopes"`
}

func (r RunnerConfigurationHostAuthenticationTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RunnerConfigurationHostAuthenticationTokenListParams struct {
	Token      param.Field[string]                                                         `query:"token"`
	PageSize   param.Field[int64]                                                          `query:"pageSize"`
	Filter     param.Field[RunnerConfigurationHostAuthenticationTokenListParamsFilter]     `json:"filter"`
	Pagination param.Field[RunnerConfigurationHostAuthenticationTokenListParamsPagination] `json:"pagination"`
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
	// Filter by subject (user or service account)
	SubjectID param.Field[string] `json:"subjectId" format:"uuid"`
	// Deprecated: Use principal_id instead
	//
	// Deprecated: deprecated
	UserID param.Field[string] `json:"userId" format:"uuid"`
}

func (r RunnerConfigurationHostAuthenticationTokenListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
