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
func (r *SecretService) New(ctx context.Context, params SecretNewParams, opts ...option.RequestOption) (res *SecretNewResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/CreateSecret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ListSecrets lists secrets.
func (r *SecretService) List(ctx context.Context, params SecretListParams, opts ...option.RequestOption) (res *SecretListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/ListSecrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// DeleteSecret deletes a secret.
func (r *SecretService) Delete(ctx context.Context, params SecretDeleteParams, opts ...option.RequestOption) (res *SecretDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/DeleteSecret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GetSecretValue retrieves the value of a secret Only Environments can perform
// this operation, and only for secrets specified on the EnvironmentSpec.
func (r *SecretService) GetValue(ctx context.Context, params SecretGetValueParams, opts ...option.RequestOption) (res *SecretGetValueResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/GetSecretValue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// UpdateSecretValue updates the value of a secret.
func (r *SecretService) UpdateValue(ctx context.Context, params SecretUpdateValueParams, opts ...option.RequestOption) (res *SecretUpdateValueResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/UpdateSecretValue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// This field can have the runtime type of [SecretNewResponseSecretObjectCreator].
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
	UpdatedAt time.Time                   `json:"updatedAt" format:"date-time"`
	JSON      secretNewResponseSecretJSON `json:"-"`
	union     SecretNewResponseSecretUnion
}

// secretNewResponseSecretJSON contains the JSON metadata for the struct
// [SecretNewResponseSecret]
type secretNewResponseSecretJSON struct {
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
// Possible runtime types of the union are [SecretNewResponseSecretObject],
// [SecretNewResponseSecretObject], [SecretNewResponseSecretObject].
func (r SecretNewResponseSecret) AsUnion() SecretNewResponseSecretUnion {
	return r.union
}

// Union satisfied by [SecretNewResponseSecretObject],
// [SecretNewResponseSecretObject] or [SecretNewResponseSecretObject].
type SecretNewResponseSecretUnion interface {
	implementsSecretNewResponseSecret()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SecretNewResponseSecretUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecretNewResponseSecretObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecretNewResponseSecretObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecretNewResponseSecretObject{}),
		},
	)
}

type SecretNewResponseSecretObject struct {
	// secret will be created as an Environment Variable with the same name as the
	// secret
	EnvironmentVariable bool   `json:"environmentVariable,required"`
	ID                  string `json:"id" format:"uuid"`
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
	// creator is the identity of the creator of the secret
	Creator SecretNewResponseSecretObjectCreator `json:"creator"`
	// absolute path to the file where the secret is mounted
	FilePath string `json:"filePath"`
	// Name of the secret for humans.
	Name string `json:"name"`
	// The Project ID this Secret belongs to
	ProjectID string `json:"projectId" format:"uuid"`
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
	UpdatedAt time.Time                         `json:"updatedAt" format:"date-time"`
	JSON      secretNewResponseSecretObjectJSON `json:"-"`
}

// secretNewResponseSecretObjectJSON contains the JSON metadata for the struct
// [SecretNewResponseSecretObject]
type secretNewResponseSecretObjectJSON struct {
	EnvironmentVariable apijson.Field
	ID                  apijson.Field
	CreatedAt           apijson.Field
	Creator             apijson.Field
	FilePath            apijson.Field
	Name                apijson.Field
	ProjectID           apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SecretNewResponseSecretObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretNewResponseSecretObjectJSON) RawJSON() string {
	return r.raw
}

func (r SecretNewResponseSecretObject) implementsSecretNewResponseSecret() {}

// creator is the identity of the creator of the secret
type SecretNewResponseSecretObjectCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal SecretNewResponseSecretObjectCreatorPrincipal `json:"principal"`
	JSON      secretNewResponseSecretObjectCreatorJSON      `json:"-"`
}

// secretNewResponseSecretObjectCreatorJSON contains the JSON metadata for the
// struct [SecretNewResponseSecretObjectCreator]
type secretNewResponseSecretObjectCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretNewResponseSecretObjectCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretNewResponseSecretObjectCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type SecretNewResponseSecretObjectCreatorPrincipal string

const (
	SecretNewResponseSecretObjectCreatorPrincipalPrincipalUnspecified    SecretNewResponseSecretObjectCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	SecretNewResponseSecretObjectCreatorPrincipalPrincipalAccount        SecretNewResponseSecretObjectCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	SecretNewResponseSecretObjectCreatorPrincipalPrincipalUser           SecretNewResponseSecretObjectCreatorPrincipal = "PRINCIPAL_USER"
	SecretNewResponseSecretObjectCreatorPrincipalPrincipalRunner         SecretNewResponseSecretObjectCreatorPrincipal = "PRINCIPAL_RUNNER"
	SecretNewResponseSecretObjectCreatorPrincipalPrincipalEnvironment    SecretNewResponseSecretObjectCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	SecretNewResponseSecretObjectCreatorPrincipalPrincipalServiceAccount SecretNewResponseSecretObjectCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r SecretNewResponseSecretObjectCreatorPrincipal) IsKnown() bool {
	switch r {
	case SecretNewResponseSecretObjectCreatorPrincipalPrincipalUnspecified, SecretNewResponseSecretObjectCreatorPrincipalPrincipalAccount, SecretNewResponseSecretObjectCreatorPrincipalPrincipalUser, SecretNewResponseSecretObjectCreatorPrincipalPrincipalRunner, SecretNewResponseSecretObjectCreatorPrincipalPrincipalEnvironment, SecretNewResponseSecretObjectCreatorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type SecretListResponse struct {
	// pagination contains the pagination options for listing secrets
	Pagination SecretListResponsePagination `json:"pagination"`
	Secrets    []SecretListResponseSecret   `json:"secrets"`
	JSON       secretListResponseJSON       `json:"-"`
}

// secretListResponseJSON contains the JSON metadata for the struct
// [SecretListResponse]
type secretListResponseJSON struct {
	Pagination  apijson.Field
	Secrets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretListResponseJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing secrets
type SecretListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                           `json:"nextToken"`
	JSON      secretListResponsePaginationJSON `json:"-"`
}

// secretListResponsePaginationJSON contains the JSON metadata for the struct
// [SecretListResponsePagination]
type secretListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type SecretListResponseSecret struct {
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
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// This field can have the runtime type of
	// [SecretListResponseSecretsObjectCreator].
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
	UpdatedAt time.Time                    `json:"updatedAt" format:"date-time"`
	JSON      secretListResponseSecretJSON `json:"-"`
	union     SecretListResponseSecretsUnion
}

// secretListResponseSecretJSON contains the JSON metadata for the struct
// [SecretListResponseSecret]
type secretListResponseSecretJSON struct {
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

func (r secretListResponseSecretJSON) RawJSON() string {
	return r.raw
}

func (r *SecretListResponseSecret) UnmarshalJSON(data []byte) (err error) {
	*r = SecretListResponseSecret{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SecretListResponseSecretsUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [SecretListResponseSecretsObject],
// [SecretListResponseSecretsObject], [SecretListResponseSecretsObject].
func (r SecretListResponseSecret) AsUnion() SecretListResponseSecretsUnion {
	return r.union
}

// Union satisfied by [SecretListResponseSecretsObject],
// [SecretListResponseSecretsObject] or [SecretListResponseSecretsObject].
type SecretListResponseSecretsUnion interface {
	implementsSecretListResponseSecret()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SecretListResponseSecretsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecretListResponseSecretsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecretListResponseSecretsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecretListResponseSecretsObject{}),
		},
	)
}

type SecretListResponseSecretsObject struct {
	// secret will be created as an Environment Variable with the same name as the
	// secret
	EnvironmentVariable bool   `json:"environmentVariable,required"`
	ID                  string `json:"id" format:"uuid"`
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
	// creator is the identity of the creator of the secret
	Creator SecretListResponseSecretsObjectCreator `json:"creator"`
	// absolute path to the file where the secret is mounted
	FilePath string `json:"filePath"`
	// Name of the secret for humans.
	Name string `json:"name"`
	// The Project ID this Secret belongs to
	ProjectID string `json:"projectId" format:"uuid"`
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
	JSON      secretListResponseSecretsObjectJSON `json:"-"`
}

// secretListResponseSecretsObjectJSON contains the JSON metadata for the struct
// [SecretListResponseSecretsObject]
type secretListResponseSecretsObjectJSON struct {
	EnvironmentVariable apijson.Field
	ID                  apijson.Field
	CreatedAt           apijson.Field
	Creator             apijson.Field
	FilePath            apijson.Field
	Name                apijson.Field
	ProjectID           apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SecretListResponseSecretsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretListResponseSecretsObjectJSON) RawJSON() string {
	return r.raw
}

func (r SecretListResponseSecretsObject) implementsSecretListResponseSecret() {}

// creator is the identity of the creator of the secret
type SecretListResponseSecretsObjectCreator struct {
	// id is the UUID of the subject
	ID string `json:"id"`
	// Principal is the principal of the subject
	Principal SecretListResponseSecretsObjectCreatorPrincipal `json:"principal"`
	JSON      secretListResponseSecretsObjectCreatorJSON      `json:"-"`
}

// secretListResponseSecretsObjectCreatorJSON contains the JSON metadata for the
// struct [SecretListResponseSecretsObjectCreator]
type secretListResponseSecretsObjectCreatorJSON struct {
	ID          apijson.Field
	Principal   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretListResponseSecretsObjectCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretListResponseSecretsObjectCreatorJSON) RawJSON() string {
	return r.raw
}

// Principal is the principal of the subject
type SecretListResponseSecretsObjectCreatorPrincipal string

const (
	SecretListResponseSecretsObjectCreatorPrincipalPrincipalUnspecified    SecretListResponseSecretsObjectCreatorPrincipal = "PRINCIPAL_UNSPECIFIED"
	SecretListResponseSecretsObjectCreatorPrincipalPrincipalAccount        SecretListResponseSecretsObjectCreatorPrincipal = "PRINCIPAL_ACCOUNT"
	SecretListResponseSecretsObjectCreatorPrincipalPrincipalUser           SecretListResponseSecretsObjectCreatorPrincipal = "PRINCIPAL_USER"
	SecretListResponseSecretsObjectCreatorPrincipalPrincipalRunner         SecretListResponseSecretsObjectCreatorPrincipal = "PRINCIPAL_RUNNER"
	SecretListResponseSecretsObjectCreatorPrincipalPrincipalEnvironment    SecretListResponseSecretsObjectCreatorPrincipal = "PRINCIPAL_ENVIRONMENT"
	SecretListResponseSecretsObjectCreatorPrincipalPrincipalServiceAccount SecretListResponseSecretsObjectCreatorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r SecretListResponseSecretsObjectCreatorPrincipal) IsKnown() bool {
	switch r {
	case SecretListResponseSecretsObjectCreatorPrincipalPrincipalUnspecified, SecretListResponseSecretsObjectCreatorPrincipalPrincipalAccount, SecretListResponseSecretsObjectCreatorPrincipalPrincipalUser, SecretListResponseSecretsObjectCreatorPrincipalPrincipalRunner, SecretListResponseSecretsObjectCreatorPrincipalPrincipalEnvironment, SecretListResponseSecretsObjectCreatorPrincipalPrincipalServiceAccount:
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
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[SecretNewParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
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

// Satisfied by [SecretNewParamsBodyObject], [SecretNewParamsBodyObject],
// [SecretNewParamsBodyObject], [SecretNewParamsBody].
type SecretNewParamsBodyUnion interface {
	implementsSecretNewParamsBodyUnion()
}

type SecretNewParamsBodyObject struct {
	// secret will be created as an Environment Variable with the same name as the
	// secret
	EnvironmentVariable param.Field[bool] `json:"environmentVariable,required"`
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

func (r SecretNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SecretNewParamsBodyObject) implementsSecretNewParamsBodyUnion() {}

// Define the version of the Connect protocol
type SecretNewParamsConnectProtocolVersion float64

const (
	SecretNewParamsConnectProtocolVersion1 SecretNewParamsConnectProtocolVersion = 1
)

func (r SecretNewParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case SecretNewParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type SecretListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[SecretListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[SecretListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[SecretListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[SecretListParamsConnect] `query:"connect"`
	Message param.Field[string]                  `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [SecretListParams]'s query parameters as `url.Values`.
func (r SecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type SecretListParamsEncoding string

const (
	SecretListParamsEncodingProto SecretListParamsEncoding = "proto"
	SecretListParamsEncodingJson  SecretListParamsEncoding = "json"
)

func (r SecretListParamsEncoding) IsKnown() bool {
	switch r {
	case SecretListParamsEncodingProto, SecretListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type SecretListParamsConnectProtocolVersion float64

const (
	SecretListParamsConnectProtocolVersion1 SecretListParamsConnectProtocolVersion = 1
)

func (r SecretListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case SecretListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type SecretListParamsCompression string

const (
	SecretListParamsCompressionIdentity SecretListParamsCompression = "identity"
	SecretListParamsCompressionGzip     SecretListParamsCompression = "gzip"
	SecretListParamsCompressionBr       SecretListParamsCompression = "br"
)

func (r SecretListParamsCompression) IsKnown() bool {
	switch r {
	case SecretListParamsCompressionIdentity, SecretListParamsCompressionGzip, SecretListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type SecretListParamsConnect string

const (
	SecretListParamsConnectV1 SecretListParamsConnect = "v1"
)

func (r SecretListParamsConnect) IsKnown() bool {
	switch r {
	case SecretListParamsConnectV1:
		return true
	}
	return false
}

type SecretDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[SecretDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	SecretID               param.Field[string]                                   `json:"secretId" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r SecretDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type SecretDeleteParamsConnectProtocolVersion float64

const (
	SecretDeleteParamsConnectProtocolVersion1 SecretDeleteParamsConnectProtocolVersion = 1
)

func (r SecretDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case SecretDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type SecretGetValueParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[SecretGetValueParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[SecretGetValueParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[SecretGetValueParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[SecretGetValueParamsConnect] `query:"connect"`
	Message param.Field[string]                      `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [SecretGetValueParams]'s query parameters as `url.Values`.
func (r SecretGetValueParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type SecretGetValueParamsEncoding string

const (
	SecretGetValueParamsEncodingProto SecretGetValueParamsEncoding = "proto"
	SecretGetValueParamsEncodingJson  SecretGetValueParamsEncoding = "json"
)

func (r SecretGetValueParamsEncoding) IsKnown() bool {
	switch r {
	case SecretGetValueParamsEncodingProto, SecretGetValueParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type SecretGetValueParamsConnectProtocolVersion float64

const (
	SecretGetValueParamsConnectProtocolVersion1 SecretGetValueParamsConnectProtocolVersion = 1
)

func (r SecretGetValueParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case SecretGetValueParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type SecretGetValueParamsCompression string

const (
	SecretGetValueParamsCompressionIdentity SecretGetValueParamsCompression = "identity"
	SecretGetValueParamsCompressionGzip     SecretGetValueParamsCompression = "gzip"
	SecretGetValueParamsCompressionBr       SecretGetValueParamsCompression = "br"
)

func (r SecretGetValueParamsCompression) IsKnown() bool {
	switch r {
	case SecretGetValueParamsCompressionIdentity, SecretGetValueParamsCompressionGzip, SecretGetValueParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type SecretGetValueParamsConnect string

const (
	SecretGetValueParamsConnectV1 SecretGetValueParamsConnect = "v1"
)

func (r SecretGetValueParamsConnect) IsKnown() bool {
	switch r {
	case SecretGetValueParamsConnectV1:
		return true
	}
	return false
}

type SecretUpdateValueParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[SecretUpdateValueParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	SecretID               param.Field[string]                                        `json:"secretId" format:"uuid"`
	// value is the plaintext value of the secret
	Value param.Field[string] `json:"value"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r SecretUpdateValueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type SecretUpdateValueParamsConnectProtocolVersion float64

const (
	SecretUpdateValueParamsConnectProtocolVersion1 SecretUpdateValueParamsConnectProtocolVersion = 1
)

func (r SecretUpdateValueParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case SecretUpdateValueParamsConnectProtocolVersion1:
		return true
	}
	return false
}
