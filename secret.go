// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/gitpod-io/flex-sdk-go/internal/apijson"
	"github.com/gitpod-io/flex-sdk-go/internal/apiquery"
	"github.com/gitpod-io/flex-sdk-go/internal/param"
	"github.com/gitpod-io/flex-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/flex-sdk-go/option"
	"github.com/gitpod-io/flex-sdk-go/packages/pagination"
	"github.com/tidwall/gjson"
)

// SecretService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecretService] method instead.
type SecretService struct {
	Options []option.RequestOption
}

// NewSecretService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSecretService(opts ...option.RequestOption) (r *SecretService) {
	r = &SecretService{}
	r.Options = opts
	return
}

// CreateSecret creates a new secret.
func (r *SecretService) New(ctx context.Context, body SecretNewParams, opts ...option.RequestOption) (res *SecretNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/CreateSecret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListSecrets lists secrets.
func (r *SecretService) List(ctx context.Context, params SecretListParams, opts ...option.RequestOption) (res *pagination.SecretsPage[SecretListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.SecretService/ListSecrets"
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

// ListSecrets lists secrets.
func (r *SecretService) ListAutoPaging(ctx context.Context, params SecretListParams, opts ...option.RequestOption) *pagination.SecretsPageAutoPager[SecretListResponse] {
	return pagination.NewSecretsPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteSecret deletes a secret.
func (r *SecretService) Delete(ctx context.Context, body SecretDeleteParams, opts ...option.RequestOption) (res *SecretDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/DeleteSecret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetSecretValue retrieves the value of a secret Only Environments can perform
// this operation, and only for secrets specified on the EnvironmentSpec.
func (r *SecretService) GetValue(ctx context.Context, body SecretGetValueParams, opts ...option.RequestOption) (res *SecretGetValueResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/GetSecretValue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateSecretValue updates the value of a secret.
func (r *SecretService) UpdateValue(ctx context.Context, body SecretUpdateValueParams, opts ...option.RequestOption) (res *SecretUpdateValueResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/UpdateSecretValue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SecretNewResponse struct {
	Secret SecretNewResponseSecret `json:"secret"`
	JSON   secretNewResponseJSON   `json:"-"`
}

// secretNewResponseJSON contains the JSON metadata for the struct
// [SecretNewResponse]
type secretNewResponseJSON struct {
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretNewResponseJSON) RawJSON() string {
	return r.raw
}

type SecretNewResponseSecret struct {
	// secret will be created as an Environment Variable with the same name as the
	// secret
	EnvironmentVariable bool `json:"environmentVariable"`
	// absolute path to the file where the secret is mounted
	FilePath string                      `json:"filePath"`
	JSON     secretNewResponseSecretJSON `json:"-"`
	union    SecretNewResponseSecretUnion
}

// secretNewResponseSecretJSON contains the JSON metadata for the struct
// [SecretNewResponseSecret]
type secretNewResponseSecretJSON struct {
	EnvironmentVariable apijson.Field
	FilePath            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r secretNewResponseSecretJSON) RawJSON() string {
	return r.raw
}

func (r *SecretNewResponseSecret) UnmarshalJSON(data []byte) (err error) {
	*r = SecretNewResponseSecret{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SecretNewResponseSecretUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [SecretNewResponseSecretSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret],
// [SecretNewResponseSecretAbsolutePathToTheFileWhereTheSecretIsMounted].
func (r SecretNewResponseSecret) AsUnion() SecretNewResponseSecretUnion {
	return r.union
}

// Union satisfied by
// [SecretNewResponseSecretSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret]
// or [SecretNewResponseSecretAbsolutePathToTheFileWhereTheSecretIsMounted].
type SecretNewResponseSecretUnion interface {
	implementsSecretNewResponseSecret()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SecretNewResponseSecretUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecretNewResponseSecretSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecretNewResponseSecretAbsolutePathToTheFileWhereTheSecretIsMounted{}),
		},
	)
}

type SecretNewResponseSecretSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret struct {
	// secret will be created as an Environment Variable with the same name as the
	// secret
	EnvironmentVariable bool                                                                                            `json:"environmentVariable,required"`
	JSON                secretNewResponseSecretSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretJSON `json:"-"`
}

// secretNewResponseSecretSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretJSON
// contains the JSON metadata for the struct
// [SecretNewResponseSecretSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret]
type secretNewResponseSecretSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretJSON struct {
	EnvironmentVariable apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SecretNewResponseSecretSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretNewResponseSecretSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretJSON) RawJSON() string {
	return r.raw
}

func (r SecretNewResponseSecretSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret) implementsSecretNewResponseSecret() {
}

type SecretNewResponseSecretAbsolutePathToTheFileWhereTheSecretIsMounted struct {
	// absolute path to the file where the secret is mounted
	FilePath string                                                                  `json:"filePath,required"`
	JSON     secretNewResponseSecretAbsolutePathToTheFileWhereTheSecretIsMountedJSON `json:"-"`
}

// secretNewResponseSecretAbsolutePathToTheFileWhereTheSecretIsMountedJSON contains
// the JSON metadata for the struct
// [SecretNewResponseSecretAbsolutePathToTheFileWhereTheSecretIsMounted]
type secretNewResponseSecretAbsolutePathToTheFileWhereTheSecretIsMountedJSON struct {
	FilePath    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretNewResponseSecretAbsolutePathToTheFileWhereTheSecretIsMounted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretNewResponseSecretAbsolutePathToTheFileWhereTheSecretIsMountedJSON) RawJSON() string {
	return r.raw
}

func (r SecretNewResponseSecretAbsolutePathToTheFileWhereTheSecretIsMounted) implementsSecretNewResponseSecret() {
}

type SecretListResponse struct {
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
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// This field can have the runtime type of
	// [SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreator],
	// [SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreator].
	Creator interface{} `json:"creator"`
	// secret will be created as an Environment Variable with the same name as the
	// secret
	EnvironmentVariable bool `json:"environmentVariable"`
	// absolute path to the file where the secret is mounted
	FilePath string `json:"filePath"`
	// Name of the secret for humans.
	Name string `json:"name"`
	// The Project ID this Secret belongs to
	ProjectID string `json:"projectId" format:"uuid"`
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
	UpdatedAt time.Time              `json:"updatedAt" format:"date-time"`
	JSON      secretListResponseJSON `json:"-"`
	union     SecretListResponseUnion
}

// secretListResponseJSON contains the JSON metadata for the struct
// [SecretListResponse]
type secretListResponseJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	Creator             apijson.Field
	EnvironmentVariable apijson.Field
	FilePath            apijson.Field
	Name                apijson.Field
	ProjectID           apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r secretListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SecretListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = SecretListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SecretListResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret],
// [SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMounted].
func (r SecretListResponse) AsUnion() SecretListResponseUnion {
	return r.union
}

// Union satisfied by
// [SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret]
// or [SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMounted].
type SecretListResponseUnion interface {
	implementsSecretListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SecretListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMounted{}),
		},
	)
}

type SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret struct {
	// secret will be created as an Environment Variable with the same name as the
	// secret
	EnvironmentVariable bool   `json:"environmentVariable,required"`
	ID                  string `json:"id" format:"uuid"`
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
	// creator is the identity of the creator of the secret
	Creator SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreator `json:"creator"`
	// Name of the secret for humans.
	Name string `json:"name"`
	// The Project ID this Secret belongs to
	ProjectID string `json:"projectId" format:"uuid"`
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
	UpdatedAt time.Time                                                                                  `json:"updatedAt" format:"date-time"`
	JSON      secretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretJSON `json:"-"`
}

// secretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretJSON
// contains the JSON metadata for the struct
// [SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret]
type secretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretJSON struct {
	EnvironmentVariable apijson.Field
	ID                  apijson.Field
	CreatedAt           apijson.Field
	Creator             apijson.Field
	Name                apijson.Field
	ProjectID           apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretJSON) RawJSON() string {
	return r.raw
}

func (r SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret) implementsSecretListResponse() {
}

// creator is the identity of the creator of the secret
type SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipal `json:"principal"`
	JSON      secretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorJSON      `json:"-"`
}

// secretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorJSON
// contains the JSON metadata for the struct
// [SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreator]
type secretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipal string

const (
	SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalUnspecified    SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalAccount        SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalUser           SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipal = "PRINCIPAL_USER"
	SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalRunner         SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipal = "PRINCIPAL_RUNNER"
	SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalEnvironment    SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalServiceAccount SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipal) IsKnown() bool {
	switch r {
	case SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalUnspecified, SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalAccount, SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalUser, SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalRunner, SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalEnvironment, SecretListResponseSecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecretCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMounted struct {
	// absolute path to the file where the secret is mounted
	FilePath string `json:"filePath,required"`
	ID       string `json:"id" format:"uuid"`
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
	// creator is the identity of the creator of the secret
	Creator SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreator `json:"creator"`
	// Name of the secret for humans.
	Name string `json:"name"`
	// The Project ID this Secret belongs to
	ProjectID string `json:"projectId" format:"uuid"`
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
	UpdatedAt time.Time                                                          `json:"updatedAt" format:"date-time"`
	JSON      secretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedJSON `json:"-"`
}

// secretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedJSON contains the
// JSON metadata for the struct
// [SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMounted]
type secretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedJSON struct {
	FilePath    apijson.Field
	ID          apijson.Field
	CreatedAt   apijson.Field
	Creator     apijson.Field
	Name        apijson.Field
	ProjectID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMounted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedJSON) RawJSON() string {
	return r.raw
}

func (r SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMounted) implementsSecretListResponse() {
}

// creator is the identity of the creator of the secret
type SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipal `json:"principal"`
	JSON      secretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorJSON      `json:"-"`
}

// secretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorJSON
// contains the JSON metadata for the struct
// [SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreator]
type secretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipal string

const (
	SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalUnspecified    SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalAccount        SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalUser           SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipal = "PRINCIPAL_USER"
	SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalRunner         SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipal = "PRINCIPAL_RUNNER"
	SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalEnvironment    SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalServiceAccount SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipal) IsKnown() bool {
	switch r {
	case SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalUnspecified, SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalAccount, SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalUser, SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalRunner, SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalEnvironment, SecretListResponseAbsolutePathToTheFileWhereTheSecretIsMountedCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type SecretDeleteResponse = interface{}

type SecretGetValueResponse struct {
	Value string                     `json:"value"`
	JSON  secretGetValueResponseJSON `json:"-"`
}

// secretGetValueResponseJSON contains the JSON metadata for the struct
// [SecretGetValueResponse]
type secretGetValueResponseJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretGetValueResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretGetValueResponseJSON) RawJSON() string {
	return r.raw
}

type SecretUpdateValueResponse = interface{}

type SecretNewParams struct {
	Body SecretNewParamsBodyUnion `json:"body,required"`
}

func (r SecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SecretNewParamsBody struct {
	// secret will be created as an Environment Variable with the same name as the
	// secret
	EnvironmentVariable param.Field[bool] `json:"environmentVariable"`
	// absolute path to the file where the secret is mounted value must be an absolute
	// path (start with a /):
	//
	// ```
	// this.matches('^/(?:[^/]*/)*.*$')
	// ```
	FilePath param.Field[string] `json:"filePath"`
	Name     param.Field[string] `json:"name"`
	// project_id is the ProjectID this Secret belongs to
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
	// value is the plaintext value of the secret
	Value param.Field[string] `json:"value"`
}

func (r SecretNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SecretNewParamsBody) implementsSecretNewParamsBodyUnion() {}

// Satisfied by
// [SecretNewParamsBodySecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret],
// [SecretNewParamsBodyAbsolutePathToTheFileWhereTheSecretIsMounted],
// [SecretNewParamsBody].
type SecretNewParamsBodyUnion interface {
	implementsSecretNewParamsBodyUnion()
}

type SecretNewParamsBodySecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret struct {
	// secret will be created as an Environment Variable with the same name as the
	// secret
	EnvironmentVariable param.Field[bool]   `json:"environmentVariable,required"`
	Name                param.Field[string] `json:"name"`
	// project_id is the ProjectID this Secret belongs to
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
	// value is the plaintext value of the secret
	Value param.Field[string] `json:"value"`
}

func (r SecretNewParamsBodySecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SecretNewParamsBodySecretWillBeCreatedAsAnEnvironmentVariableWithTheSameNameAsTheSecret) implementsSecretNewParamsBodyUnion() {
}

type SecretNewParamsBodyAbsolutePathToTheFileWhereTheSecretIsMounted struct {
	// absolute path to the file where the secret is mounted value must be an absolute
	// path (start with a /):
	//
	// ```
	// this.matches('^/(?:[^/]*/)*.*$')
	// ```
	FilePath param.Field[string] `json:"filePath,required"`
	Name     param.Field[string] `json:"name"`
	// project_id is the ProjectID this Secret belongs to
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
	// value is the plaintext value of the secret
	Value param.Field[string] `json:"value"`
}

func (r SecretNewParamsBodyAbsolutePathToTheFileWhereTheSecretIsMounted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SecretNewParamsBodyAbsolutePathToTheFileWhereTheSecretIsMounted) implementsSecretNewParamsBodyUnion() {
}

type SecretListParams struct {
	Token    param.Field[string]                 `query:"token"`
	PageSize param.Field[int64]                  `query:"pageSize"`
	Filter   param.Field[SecretListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing environments
	Pagination param.Field[SecretListParamsPagination] `json:"pagination"`
}

func (r SecretListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SecretListParams]'s query parameters as `url.Values`.
func (r SecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SecretListParamsFilter struct {
	// project_ids filters the response to only Secrets used by these Project IDs
	ProjectIDs param.Field[[]string] `json:"projectIds" format:"uuid"`
}

func (r SecretListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing environments
type SecretListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r SecretListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecretDeleteParams struct {
	SecretID param.Field[string] `json:"secretId" format:"uuid"`
}

func (r SecretDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecretGetValueParams struct {
	SecretID param.Field[string] `json:"secretId" format:"uuid"`
}

func (r SecretGetValueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecretUpdateValueParams struct {
	SecretID param.Field[string] `json:"secretId" format:"uuid"`
	// value is the plaintext value of the secret
	Value param.Field[string] `json:"value"`
}

func (r SecretUpdateValueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
