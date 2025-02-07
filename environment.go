// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gitpod-io/flex-sdk-go/internal/apijson"
	"github.com/gitpod-io/flex-sdk-go/internal/apiquery"
	"github.com/gitpod-io/flex-sdk-go/internal/param"
	"github.com/gitpod-io/flex-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/flex-sdk-go/option"
	"github.com/gitpod-io/flex-sdk-go/packages/pagination"
	"github.com/gitpod-io/flex-sdk-go/shared"
)

// EnvironmentService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentService] method instead.
type EnvironmentService struct {
	Options     []option.RequestOption
	Automations *EnvironmentAutomationService
	Classes     *EnvironmentClassService
}

// NewEnvironmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnvironmentService(opts ...option.RequestOption) (r *EnvironmentService) {
	r = &EnvironmentService{}
	r.Options = opts
	r.Automations = NewEnvironmentAutomationService(opts...)
	r.Classes = NewEnvironmentClassService(opts...)
	return
}

// CreateEnvironment creates a new environment and starts it.
//
// The `class` field must be a valid environment class ID. You can find a list of
// available environment classes with the `ListEnvironmentClasses` method.
//
// ### Examples
//
// - from context URL:
//
//	Creates an environment from a context URL, e.g. a GitHub repository.
//
//	```yaml
//	spec:
//	  machine:
//	    class: "61000000-0000-0000-0000-000000000000"
//	  content:
//	    initializer:
//	      specs:
//	        - contextUrl:
//	            url: "https://github.com/gitpod-io/gitpod"
//	```
//
// - from Git repository:
//
//	Creates an environment from a Git repository directly. While less convenient,
//	this is useful if you want to specify a specific branch, commit, etc.
//
//	```yaml
//	spec:
//	  machine:
//	    class: "61000000-0000-0000-0000-000000000000"
//	  content:
//	    initializer:
//	      specs:
//	        - git:
//	            remoteUri: "https://github.com/gitpod-io/gitpod"
//	            cloneTarget: "main"
//	            targetMode: "CLONE_TARGET_MODE_REMOTE_BRANCH"
//	```
func (r *EnvironmentService) New(ctx context.Context, body EnvironmentNewParams, opts ...option.RequestOption) (res *EnvironmentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/CreateEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetEnvironment returns a single environment.
func (r *EnvironmentService) Get(ctx context.Context, body EnvironmentGetParams, opts ...option.RequestOption) (res *EnvironmentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/GetEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// UpdateEnvironment updates the environment partially.
func (r *EnvironmentService) Update(ctx context.Context, body EnvironmentUpdateParams, opts ...option.RequestOption) (res *EnvironmentUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/UpdateEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListEnvironments returns a list of environments that match the query.
func (r *EnvironmentService) List(ctx context.Context, params EnvironmentListParams, opts ...option.RequestOption) (res *pagination.EnvironmentsPage[Environment], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.EnvironmentService/ListEnvironments"
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

// ListEnvironments returns a list of environments that match the query.
func (r *EnvironmentService) ListAutoPaging(ctx context.Context, params EnvironmentListParams, opts ...option.RequestOption) *pagination.EnvironmentsPageAutoPager[Environment] {
	return pagination.NewEnvironmentsPageAutoPager(r.List(ctx, params, opts...))
}

// DeleteEnvironment deletes an environment. When the environment is running, it
// will be stopped as well. Deleted environments cannot be started again.
func (r *EnvironmentService) Delete(ctx context.Context, body EnvironmentDeleteParams, opts ...option.RequestOption) (res *EnvironmentDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/DeleteEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// CreateAbdStartEnvironmentFromProject creates a new environment from a project
// and starts it.
func (r *EnvironmentService) NewFromProject(ctx context.Context, body EnvironmentNewFromProjectParams, opts ...option.RequestOption) (res *EnvironmentNewFromProjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/CreateEnvironmentFromProject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// CreateEnvironmentLogsToken creates a token that can be used to access the logs
// of an environment.
func (r *EnvironmentService) NewLogsToken(ctx context.Context, body EnvironmentNewLogsTokenParams, opts ...option.RequestOption) (res *EnvironmentNewLogsTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/CreateEnvironmentLogsToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// MarkEnvironmentActive allows tools to signal activity for an environment.
func (r *EnvironmentService) MarkActive(ctx context.Context, body EnvironmentMarkActiveParams, opts ...option.RequestOption) (res *EnvironmentMarkActiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/MarkEnvironmentActive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// StartEnvironment starts an environment. This function is idempotent, i.e. if the
// environment is already running no error is returned.
func (r *EnvironmentService) Start(ctx context.Context, body EnvironmentStartParams, opts ...option.RequestOption) (res *EnvironmentStartResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/StartEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// StopEnvironment stops a running environment.
func (r *EnvironmentService) Stop(ctx context.Context, body EnvironmentStopParams, opts ...option.RequestOption) (res *EnvironmentStopResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/StopEnvironment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Admission level describes who can access an environment instance and its ports.
type AdmissionLevel string

const (
	AdmissionLevelUnspecified AdmissionLevel = "ADMISSION_LEVEL_UNSPECIFIED"
	AdmissionLevelOwnerOnly   AdmissionLevel = "ADMISSION_LEVEL_OWNER_ONLY"
	AdmissionLevelEveryone    AdmissionLevel = "ADMISSION_LEVEL_EVERYONE"
)

func (r AdmissionLevel) IsKnown() bool {
	switch r {
	case AdmissionLevelUnspecified, AdmissionLevelOwnerOnly, AdmissionLevelEveryone:
		return true
	}
	return false
}

// +resource get environment
type Environment struct {
	// ID is a unique identifier of this environment. No other environment with the
	// same name must be managed by this environment manager
	ID string `json:"id"`
	// EnvironmentMetadata is data associated with an environment that's required for
	// other parts of the system to function
	Metadata EnvironmentMetadata `json:"metadata"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec EnvironmentSpec `json:"spec"`
	// EnvironmentStatus describes an environment status
	Status EnvironmentStatus `json:"status"`
	JSON   environmentJSON   `json:"-"`
}

// environmentJSON contains the JSON metadata for the struct [Environment]
type environmentJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Environment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentJSON) RawJSON() string {
	return r.raw
}

// EnvironmentActivitySignal used to signal activity for an environment.
type EnvironmentActivitySignal struct {
	// source of the activity signal, such as "VS Code", "SSH", or "Automations". It
	// should be a human-readable string that describes the source of the activity
	// signal.
	Source string `json:"source"`
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
	Timestamp time.Time                     `json:"timestamp" format:"date-time"`
	JSON      environmentActivitySignalJSON `json:"-"`
}

// environmentActivitySignalJSON contains the JSON metadata for the struct
// [EnvironmentActivitySignal]
type environmentActivitySignalJSON struct {
	Source      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentActivitySignal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentActivitySignalJSON) RawJSON() string {
	return r.raw
}

// EnvironmentActivitySignal used to signal activity for an environment.
type EnvironmentActivitySignalParam struct {
	// source of the activity signal, such as "VS Code", "SSH", or "Automations". It
	// should be a human-readable string that describes the source of the activity
	// signal.
	Source param.Field[string] `json:"source"`
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
	Timestamp param.Field[time.Time] `json:"timestamp" format:"date-time"`
}

func (r EnvironmentActivitySignalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnvironmentMetadata is data associated with an environment that's required for
// other parts of the system to function
type EnvironmentMetadata struct {
	// annotations are key/value pairs that gets attached to the environment.
	// +internal - not yet implemented
	Annotations map[string]string `json:"annotations"`
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
	// creator is the identity of the creator of the environment
	Creator shared.Subject `json:"creator"`
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
	LastStartedAt time.Time `json:"lastStartedAt" format:"date-time"`
	// name is the name of the environment as specified by the user
	Name string `json:"name"`
	// organization_id is the ID of the organization that contains the environment
	OrganizationID string `json:"organizationId" format:"uuid"`
	// original_context_url is the normalized URL from which the environment was
	// created
	OriginalContextURL string `json:"originalContextUrl"`
	// If the Environment was started from a project, the project_id will reference the
	// project.
	ProjectID string `json:"projectId"`
	// Runner is the ID of the runner that runs this environment.
	RunnerID string                  `json:"runnerId"`
	JSON     environmentMetadataJSON `json:"-"`
}

// environmentMetadataJSON contains the JSON metadata for the struct
// [EnvironmentMetadata]
type environmentMetadataJSON struct {
	Annotations        apijson.Field
	CreatedAt          apijson.Field
	Creator            apijson.Field
	LastStartedAt      apijson.Field
	Name               apijson.Field
	OrganizationID     apijson.Field
	OriginalContextURL apijson.Field
	ProjectID          apijson.Field
	RunnerID           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentMetadataJSON) RawJSON() string {
	return r.raw
}

type EnvironmentPhase string

const (
	EnvironmentPhaseUnspecified EnvironmentPhase = "ENVIRONMENT_PHASE_UNSPECIFIED"
	EnvironmentPhaseCreating    EnvironmentPhase = "ENVIRONMENT_PHASE_CREATING"
	EnvironmentPhaseStarting    EnvironmentPhase = "ENVIRONMENT_PHASE_STARTING"
	EnvironmentPhaseRunning     EnvironmentPhase = "ENVIRONMENT_PHASE_RUNNING"
	EnvironmentPhaseUpdating    EnvironmentPhase = "ENVIRONMENT_PHASE_UPDATING"
	EnvironmentPhaseStopping    EnvironmentPhase = "ENVIRONMENT_PHASE_STOPPING"
	EnvironmentPhaseStopped     EnvironmentPhase = "ENVIRONMENT_PHASE_STOPPED"
	EnvironmentPhaseDeleting    EnvironmentPhase = "ENVIRONMENT_PHASE_DELETING"
	EnvironmentPhaseDeleted     EnvironmentPhase = "ENVIRONMENT_PHASE_DELETED"
)

func (r EnvironmentPhase) IsKnown() bool {
	switch r {
	case EnvironmentPhaseUnspecified, EnvironmentPhaseCreating, EnvironmentPhaseStarting, EnvironmentPhaseRunning, EnvironmentPhaseUpdating, EnvironmentPhaseStopping, EnvironmentPhaseStopped, EnvironmentPhaseDeleting, EnvironmentPhaseDeleted:
		return true
	}
	return false
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type EnvironmentSpec struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission AdmissionLevel `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile EnvironmentSpecAutomationsFile `json:"automationsFile"`
	// content is the content spec of the environment
	Content EnvironmentSpecContent `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase EnvironmentPhase `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer EnvironmentSpecDevcontainer `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine EnvironmentSpecMachine `json:"machine"`
	// ports is the set of ports which ought to be exposed to the internet
	Ports []EnvironmentSpecPort `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets []EnvironmentSpecSecret `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion string `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys []EnvironmentSpecSSHPublicKey `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout EnvironmentSpecTimeout `json:"timeout"`
	JSON    environmentSpecJSON    `json:"-"`
}

// environmentSpecJSON contains the JSON metadata for the struct [EnvironmentSpec]
type environmentSpecJSON struct {
	Admission       apijson.Field
	AutomationsFile apijson.Field
	Content         apijson.Field
	DesiredPhase    apijson.Field
	Devcontainer    apijson.Field
	Machine         apijson.Field
	Ports           apijson.Field
	Secrets         apijson.Field
	SpecVersion     apijson.Field
	SSHPublicKeys   apijson.Field
	Timeout         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EnvironmentSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecJSON) RawJSON() string {
	return r.raw
}

// automations_file is the automations file spec of the environment
type EnvironmentSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath string                             `json:"automationsFilePath"`
	Session             string                             `json:"session"`
	JSON                environmentSpecAutomationsFileJSON `json:"-"`
}

// environmentSpecAutomationsFileJSON contains the JSON metadata for the struct
// [EnvironmentSpecAutomationsFile]
type environmentSpecAutomationsFileJSON struct {
	AutomationsFilePath apijson.Field
	Session             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentSpecAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// content is the content spec of the environment
type EnvironmentSpecContent struct {
	// The Git email address
	GitEmail string `json:"gitEmail"`
	// The Git username
	GitUsername string `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer EnvironmentInitializer     `json:"initializer"`
	Session     string                     `json:"session"`
	JSON        environmentSpecContentJSON `json:"-"`
}

// environmentSpecContentJSON contains the JSON metadata for the struct
// [EnvironmentSpecContent]
type environmentSpecContentJSON struct {
	GitEmail    apijson.Field
	GitUsername apijson.Field
	Initializer apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentSpecContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecContentJSON) RawJSON() string {
	return r.raw
}

// devcontainer is the devcontainer spec of the environment
type EnvironmentSpecDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath string                          `json:"devcontainerFilePath"`
	Session              string                          `json:"session"`
	JSON                 environmentSpecDevcontainerJSON `json:"-"`
}

// environmentSpecDevcontainerJSON contains the JSON metadata for the struct
// [EnvironmentSpecDevcontainer]
type environmentSpecDevcontainerJSON struct {
	DevcontainerFilePath apijson.Field
	Session              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentSpecDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecDevcontainerJSON) RawJSON() string {
	return r.raw
}

// machine is the machine spec of the environment
type EnvironmentSpecMachine struct {
	// Class denotes the class of the environment we ought to start
	Class   string                     `json:"class" format:"uuid"`
	Session string                     `json:"session"`
	JSON    environmentSpecMachineJSON `json:"-"`
}

// environmentSpecMachineJSON contains the JSON metadata for the struct
// [EnvironmentSpecMachine]
type environmentSpecMachineJSON struct {
	Class       apijson.Field
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentSpecMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecMachineJSON) RawJSON() string {
	return r.raw
}

type EnvironmentSpecPort struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission AdmissionLevel `json:"admission"`
	// name of this port
	Name string `json:"name"`
	// port number
	Port int64                   `json:"port"`
	JSON environmentSpecPortJSON `json:"-"`
}

// environmentSpecPortJSON contains the JSON metadata for the struct
// [EnvironmentSpecPort]
type environmentSpecPortJSON struct {
	Admission   apijson.Field
	Name        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentSpecPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecPortJSON) RawJSON() string {
	return r.raw
}

type EnvironmentSpecSecret struct {
	EnvironmentVariable string `json:"environmentVariable"`
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath          string `json:"filePath"`
	GitCredentialHost string `json:"gitCredentialHost"`
	// name is the human readable description of the secret
	Name string `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session string `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source string `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef string                    `json:"sourceRef"`
	JSON      environmentSpecSecretJSON `json:"-"`
}

// environmentSpecSecretJSON contains the JSON metadata for the struct
// [EnvironmentSpecSecret]
type environmentSpecSecretJSON struct {
	EnvironmentVariable apijson.Field
	FilePath            apijson.Field
	GitCredentialHost   apijson.Field
	Name                apijson.Field
	Session             apijson.Field
	Source              apijson.Field
	SourceRef           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EnvironmentSpecSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecSecretJSON) RawJSON() string {
	return r.raw
}

type EnvironmentSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// value is the actual public key in the public key file format
	Value string                          `json:"value"`
	JSON  environmentSpecSSHPublicKeyJSON `json:"-"`
}

// environmentSpecSSHPublicKeyJSON contains the JSON metadata for the struct
// [EnvironmentSpecSSHPublicKey]
type environmentSpecSSHPublicKeyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentSpecSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// Timeout configures the environment timeout
type EnvironmentSpecTimeout struct {
	// A Duration represents a signed, fixed-length span of time represented as a count
	// of seconds and fractions of seconds at nanosecond resolution. It is independent
	// of any calendar and concepts like "day" or "month". It is related to Timestamp
	// in that the difference between two Timestamp values is a Duration and it can be
	// added or subtracted from a Timestamp. Range is approximately +-10,000 years.
	//
	// # Examples
	//
	// Example 1: Compute Duration from two Timestamps in pseudo code.
	//
	//	Timestamp start = ...;
	//	Timestamp end = ...;
	//	Duration duration = ...;
	//
	//	duration.seconds = end.seconds - start.seconds;
	//	duration.nanos = end.nanos - start.nanos;
	//
	//	if (duration.seconds < 0 && duration.nanos > 0) {
	//	  duration.seconds += 1;
	//	  duration.nanos -= 1000000000;
	//	} else if (duration.seconds > 0 && duration.nanos < 0) {
	//	  duration.seconds -= 1;
	//	  duration.nanos += 1000000000;
	//	}
	//
	// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
	//
	//	Timestamp start = ...;
	//	Duration duration = ...;
	//	Timestamp end = ...;
	//
	//	end.seconds = start.seconds + duration.seconds;
	//	end.nanos = start.nanos + duration.nanos;
	//
	//	if (end.nanos < 0) {
	//	  end.seconds -= 1;
	//	  end.nanos += 1000000000;
	//	} else if (end.nanos >= 1000000000) {
	//	  end.seconds += 1;
	//	  end.nanos -= 1000000000;
	//	}
	//
	// Example 3: Compute Duration from datetime.timedelta in Python.
	//
	//	td = datetime.timedelta(days=3, minutes=10)
	//	duration = Duration()
	//	duration.FromTimedelta(td)
	//
	// # JSON Mapping
	//
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by
	// the number of seconds, with nanoseconds expressed as fractional seconds. For
	// example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as
	// "3.000000001s", and 3 seconds and 1 microsecond should be expressed in JSON
	// format as "3.000001s".
	Disconnected string                     `json:"disconnected" format:"regex"`
	JSON         environmentSpecTimeoutJSON `json:"-"`
}

// environmentSpecTimeoutJSON contains the JSON metadata for the struct
// [EnvironmentSpecTimeout]
type environmentSpecTimeoutJSON struct {
	Disconnected apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentSpecTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentSpecTimeoutJSON) RawJSON() string {
	return r.raw
}

// EnvironmentSpec specifies the configuration of an environment for an environment
// start
type EnvironmentSpecParam struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission param.Field[AdmissionLevel] `json:"admission"`
	// automations_file is the automations file spec of the environment
	AutomationsFile param.Field[EnvironmentSpecAutomationsFileParam] `json:"automationsFile"`
	// content is the content spec of the environment
	Content param.Field[EnvironmentSpecContentParam] `json:"content"`
	// Phase is the desired phase of the environment
	DesiredPhase param.Field[EnvironmentPhase] `json:"desiredPhase"`
	// devcontainer is the devcontainer spec of the environment
	Devcontainer param.Field[EnvironmentSpecDevcontainerParam] `json:"devcontainer"`
	// machine is the machine spec of the environment
	Machine param.Field[EnvironmentSpecMachineParam] `json:"machine"`
	// ports is the set of ports which ought to be exposed to the internet
	Ports param.Field[[]EnvironmentSpecPortParam] `json:"ports"`
	// secrets are confidential data that is mounted into the environment
	Secrets param.Field[[]EnvironmentSpecSecretParam] `json:"secrets"`
	// version of the spec. The value of this field has no semantic meaning (e.g. don't
	// interpret it as as a timestamp), but it can be used to impose a partial order.
	// If a.spec_version < b.spec_version then a was the spec before b.
	SpecVersion param.Field[string] `json:"specVersion"`
	// ssh_public_keys are the public keys used to ssh into the environment
	SSHPublicKeys param.Field[[]EnvironmentSpecSSHPublicKeyParam] `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout param.Field[EnvironmentSpecTimeoutParam] `json:"timeout"`
}

func (r EnvironmentSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// automations_file is the automations file spec of the environment
type EnvironmentSpecAutomationsFileParam struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath"`
	Session             param.Field[string] `json:"session"`
}

func (r EnvironmentSpecAutomationsFileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// content is the content spec of the environment
type EnvironmentSpecContentParam struct {
	// The Git email address
	GitEmail param.Field[string] `json:"gitEmail"`
	// The Git username
	GitUsername param.Field[string] `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer param.Field[EnvironmentInitializerParam] `json:"initializer"`
	Session     param.Field[string]                      `json:"session"`
}

func (r EnvironmentSpecContentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// devcontainer is the devcontainer spec of the environment
type EnvironmentSpecDevcontainerParam struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath"`
	Session              param.Field[string] `json:"session"`
}

func (r EnvironmentSpecDevcontainerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// machine is the machine spec of the environment
type EnvironmentSpecMachineParam struct {
	// Class denotes the class of the environment we ought to start
	Class   param.Field[string] `json:"class" format:"uuid"`
	Session param.Field[string] `json:"session"`
}

func (r EnvironmentSpecMachineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentSpecPortParam struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission param.Field[AdmissionLevel] `json:"admission"`
	// name of this port
	Name param.Field[string] `json:"name"`
	// port number
	Port param.Field[int64] `json:"port"`
}

func (r EnvironmentSpecPortParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentSpecSecretParam struct {
	EnvironmentVariable param.Field[string] `json:"environmentVariable"`
	// file_path is the path inside the devcontainer where the secret is mounted
	FilePath          param.Field[string] `json:"filePath"`
	GitCredentialHost param.Field[string] `json:"gitCredentialHost"`
	// name is the human readable description of the secret
	Name param.Field[string] `json:"name"`
	// session indicated the current session of the secret. When the session does not
	// change, secrets are not reloaded in the environment.
	Session param.Field[string] `json:"session"`
	// source is the source of the secret, for now control-plane or runner
	Source param.Field[string] `json:"source"`
	// source_ref into the source, in case of control-plane this is uuid of the secret
	SourceRef param.Field[string] `json:"sourceRef"`
}

func (r EnvironmentSpecSecretParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentSpecSSHPublicKeyParam struct {
	// id is the unique identifier of the public key
	ID param.Field[string] `json:"id"`
	// value is the actual public key in the public key file format
	Value param.Field[string] `json:"value"`
}

func (r EnvironmentSpecSSHPublicKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Timeout configures the environment timeout
type EnvironmentSpecTimeoutParam struct {
	// A Duration represents a signed, fixed-length span of time represented as a count
	// of seconds and fractions of seconds at nanosecond resolution. It is independent
	// of any calendar and concepts like "day" or "month". It is related to Timestamp
	// in that the difference between two Timestamp values is a Duration and it can be
	// added or subtracted from a Timestamp. Range is approximately +-10,000 years.
	//
	// # Examples
	//
	// Example 1: Compute Duration from two Timestamps in pseudo code.
	//
	//	Timestamp start = ...;
	//	Timestamp end = ...;
	//	Duration duration = ...;
	//
	//	duration.seconds = end.seconds - start.seconds;
	//	duration.nanos = end.nanos - start.nanos;
	//
	//	if (duration.seconds < 0 && duration.nanos > 0) {
	//	  duration.seconds += 1;
	//	  duration.nanos -= 1000000000;
	//	} else if (duration.seconds > 0 && duration.nanos < 0) {
	//	  duration.seconds -= 1;
	//	  duration.nanos += 1000000000;
	//	}
	//
	// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
	//
	//	Timestamp start = ...;
	//	Duration duration = ...;
	//	Timestamp end = ...;
	//
	//	end.seconds = start.seconds + duration.seconds;
	//	end.nanos = start.nanos + duration.nanos;
	//
	//	if (end.nanos < 0) {
	//	  end.seconds -= 1;
	//	  end.nanos += 1000000000;
	//	} else if (end.nanos >= 1000000000) {
	//	  end.seconds += 1;
	//	  end.nanos -= 1000000000;
	//	}
	//
	// Example 3: Compute Duration from datetime.timedelta in Python.
	//
	//	td = datetime.timedelta(days=3, minutes=10)
	//	duration = Duration()
	//	duration.FromTimedelta(td)
	//
	// # JSON Mapping
	//
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by
	// the number of seconds, with nanoseconds expressed as fractional seconds. For
	// example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as
	// "3.000000001s", and 3 seconds and 1 microsecond should be expressed in JSON
	// format as "3.000001s".
	Disconnected param.Field[string] `json:"disconnected" format:"regex"`
}

func (r EnvironmentSpecTimeoutParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnvironmentStatus describes an environment status
type EnvironmentStatus struct {
	// EnvironmentActivitySignal used to signal activity for an environment.
	ActivitySignal EnvironmentActivitySignal `json:"activitySignal"`
	// automations_file contains the status of the automations file.
	AutomationsFile EnvironmentStatusAutomationsFile `json:"automationsFile"`
	// content contains the status of the environment content.
	Content EnvironmentStatusContent `json:"content"`
	// devcontainer contains the status of the devcontainer.
	Devcontainer EnvironmentStatusDevcontainer `json:"devcontainer"`
	// environment_url contains the URL at which the environment can be accessed. This
	// field is only set if the environment is running.
	EnvironmentURLs EnvironmentStatusEnvironmentURLs `json:"environmentUrls"`
	// failure_message summarises why the environment failed to operate. If this is
	// non-empty the environment has failed to operate and will likely transition to a
	// stopped state.
	FailureMessage []string `json:"failureMessage"`
	// machine contains the status of the environment machine
	Machine EnvironmentStatusMachine `json:"machine"`
	// the phase of an environment is a simple, high-level summary of where the
	// environment is in its lifecycle
	Phase EnvironmentPhase `json:"phase"`
	// RunnerACK is the acknowledgement from the runner that is has received the
	// environment spec.
	RunnerAck EnvironmentStatusRunnerAck `json:"runnerAck"`
	// secrets contains the status of the environment secrets
	Secrets []EnvironmentStatusSecret `json:"secrets"`
	// ssh_public_keys contains the status of the environment ssh public keys
	SSHPublicKeys []EnvironmentStatusSSHPublicKey `json:"sshPublicKeys"`
	// version of the status update. Environment instances themselves are unversioned,
	// but their status has different versions. The value of this field has no semantic
	// meaning (e.g. don't interpret it as as a timestamp), but it can be used to
	// impose a partial order. If a.status_version < b.status_version then a was the
	// status before b.
	StatusVersion string `json:"statusVersion"`
	// warning_message contains warnings, e.g. when the environment is present but not
	// in the expected state.
	WarningMessage []string              `json:"warningMessage"`
	JSON           environmentStatusJSON `json:"-"`
}

// environmentStatusJSON contains the JSON metadata for the struct
// [EnvironmentStatus]
type environmentStatusJSON struct {
	ActivitySignal  apijson.Field
	AutomationsFile apijson.Field
	Content         apijson.Field
	Devcontainer    apijson.Field
	EnvironmentURLs apijson.Field
	FailureMessage  apijson.Field
	Machine         apijson.Field
	Phase           apijson.Field
	RunnerAck       apijson.Field
	Secrets         apijson.Field
	SSHPublicKeys   apijson.Field
	StatusVersion   apijson.Field
	WarningMessage  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EnvironmentStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusJSON) RawJSON() string {
	return r.raw
}

// automations_file contains the status of the automations file.
type EnvironmentStatusAutomationsFile struct {
	// automations_file_path is the path to the automations file relative to the repo
	// root.
	AutomationsFilePath string `json:"automationsFilePath"`
	// automations_file_presence indicates how an automations file is present in the
	// environment.
	AutomationsFilePresence EnvironmentStatusAutomationsFileAutomationsFilePresence `json:"automationsFilePresence"`
	// failure_message contains the reason the automations file failed to be applied.
	// This is only set if the phase is FAILED.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the automations file.
	Phase EnvironmentStatusAutomationsFilePhase `json:"phase"`
	// session is the automations file session that is currently applied in the
	// environment.
	Session string                               `json:"session"`
	JSON    environmentStatusAutomationsFileJSON `json:"-"`
}

// environmentStatusAutomationsFileJSON contains the JSON metadata for the struct
// [EnvironmentStatusAutomationsFile]
type environmentStatusAutomationsFileJSON struct {
	AutomationsFilePath     apijson.Field
	AutomationsFilePresence apijson.Field
	FailureMessage          apijson.Field
	Phase                   apijson.Field
	Session                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *EnvironmentStatusAutomationsFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusAutomationsFileJSON) RawJSON() string {
	return r.raw
}

// automations_file_presence indicates how an automations file is present in the
// environment.
type EnvironmentStatusAutomationsFileAutomationsFilePresence string

const (
	EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceUnspecified EnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceAbsent      EnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_ABSENT"
	EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceDiscovered  EnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceSpecified   EnvironmentStatusAutomationsFileAutomationsFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentStatusAutomationsFileAutomationsFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceUnspecified, EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceAbsent, EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceDiscovered, EnvironmentStatusAutomationsFileAutomationsFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the automations file.
type EnvironmentStatusAutomationsFilePhase string

const (
	EnvironmentStatusAutomationsFilePhaseContentPhaseUnspecified  EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentStatusAutomationsFilePhaseContentPhaseCreating     EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_CREATING"
	EnvironmentStatusAutomationsFilePhaseContentPhaseInitializing EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentStatusAutomationsFilePhaseContentPhaseReady        EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_READY"
	EnvironmentStatusAutomationsFilePhaseContentPhaseUpdating     EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_UPDATING"
	EnvironmentStatusAutomationsFilePhaseContentPhaseFailed       EnvironmentStatusAutomationsFilePhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentStatusAutomationsFilePhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusAutomationsFilePhaseContentPhaseUnspecified, EnvironmentStatusAutomationsFilePhaseContentPhaseCreating, EnvironmentStatusAutomationsFilePhaseContentPhaseInitializing, EnvironmentStatusAutomationsFilePhaseContentPhaseReady, EnvironmentStatusAutomationsFilePhaseContentPhaseUpdating, EnvironmentStatusAutomationsFilePhaseContentPhaseFailed:
		return true
	}
	return false
}

// content contains the status of the environment content.
type EnvironmentStatusContent struct {
	// content_location_in_machine is the location of the content in the machine
	ContentLocationInMachine string `json:"contentLocationInMachine"`
	// failure_message contains the reason the content initialization failed.
	FailureMessage string `json:"failureMessage"`
	// git is the Git working copy status of the environment. Note: this is a
	// best-effort field and more often than not will not be present. Its absence does
	// not indicate the absence of a working copy.
	Git EnvironmentStatusContentGit `json:"git"`
	// phase is the current phase of the environment content
	Phase EnvironmentStatusContentPhase `json:"phase"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the content is present but not in
	// the expected state.
	WarningMessage string                       `json:"warningMessage"`
	JSON           environmentStatusContentJSON `json:"-"`
}

// environmentStatusContentJSON contains the JSON metadata for the struct
// [EnvironmentStatusContent]
type environmentStatusContentJSON struct {
	ContentLocationInMachine apijson.Field
	FailureMessage           apijson.Field
	Git                      apijson.Field
	Phase                    apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentStatusContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusContentJSON) RawJSON() string {
	return r.raw
}

// git is the Git working copy status of the environment. Note: this is a
// best-effort field and more often than not will not be present. Its absence does
// not indicate the absence of a working copy.
type EnvironmentStatusContentGit struct {
	// branch is branch we're currently on
	Branch string `json:"branch"`
	// changed_files is an array of changed files in the environment, possibly
	// truncated
	ChangedFiles []EnvironmentStatusContentGitChangedFile `json:"changedFiles"`
	// clone_url is the repository url as you would pass it to "git clone". Only HTTPS
	// clone URLs are supported.
	CloneURL string `json:"cloneUrl"`
	// latest_commit is the most recent commit on the current branch
	LatestCommit      string `json:"latestCommit"`
	TotalChangedFiles int64  `json:"totalChangedFiles"`
	// the total number of unpushed changes
	TotalUnpushedCommits int64 `json:"totalUnpushedCommits"`
	// unpushed_commits is an array of unpushed changes in the environment, possibly
	// truncated
	UnpushedCommits []string                        `json:"unpushedCommits"`
	JSON            environmentStatusContentGitJSON `json:"-"`
}

// environmentStatusContentGitJSON contains the JSON metadata for the struct
// [EnvironmentStatusContentGit]
type environmentStatusContentGitJSON struct {
	Branch               apijson.Field
	ChangedFiles         apijson.Field
	CloneURL             apijson.Field
	LatestCommit         apijson.Field
	TotalChangedFiles    apijson.Field
	TotalUnpushedCommits apijson.Field
	UnpushedCommits      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *EnvironmentStatusContentGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusContentGitJSON) RawJSON() string {
	return r.raw
}

type EnvironmentStatusContentGitChangedFile struct {
	// ChangeType is the type of change that happened to the file
	ChangeType EnvironmentStatusContentGitChangedFilesChangeType `json:"changeType"`
	// path is the path of the file
	Path string                                     `json:"path"`
	JSON environmentStatusContentGitChangedFileJSON `json:"-"`
}

// environmentStatusContentGitChangedFileJSON contains the JSON metadata for the
// struct [EnvironmentStatusContentGitChangedFile]
type environmentStatusContentGitChangedFileJSON struct {
	ChangeType  apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentStatusContentGitChangedFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusContentGitChangedFileJSON) RawJSON() string {
	return r.raw
}

// ChangeType is the type of change that happened to the file
type EnvironmentStatusContentGitChangedFilesChangeType string

const (
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified        EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNSPECIFIED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeAdded              EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_ADDED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeModified           EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_MODIFIED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeDeleted            EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_DELETED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeRenamed            EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_RENAMED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeCopied             EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_COPIED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UPDATED_BUT_UNMERGED"
	EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUntracked          EnvironmentStatusContentGitChangedFilesChangeType = "CHANGE_TYPE_UNTRACKED"
)

func (r EnvironmentStatusContentGitChangedFilesChangeType) IsKnown() bool {
	switch r {
	case EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUnspecified, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeAdded, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeModified, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeDeleted, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeRenamed, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeCopied, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUpdatedButUnmerged, EnvironmentStatusContentGitChangedFilesChangeTypeChangeTypeUntracked:
		return true
	}
	return false
}

// phase is the current phase of the environment content
type EnvironmentStatusContentPhase string

const (
	EnvironmentStatusContentPhaseContentPhaseUnspecified  EnvironmentStatusContentPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentStatusContentPhaseContentPhaseCreating     EnvironmentStatusContentPhase = "CONTENT_PHASE_CREATING"
	EnvironmentStatusContentPhaseContentPhaseInitializing EnvironmentStatusContentPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentStatusContentPhaseContentPhaseReady        EnvironmentStatusContentPhase = "CONTENT_PHASE_READY"
	EnvironmentStatusContentPhaseContentPhaseUpdating     EnvironmentStatusContentPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentStatusContentPhaseContentPhaseFailed       EnvironmentStatusContentPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentStatusContentPhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusContentPhaseContentPhaseUnspecified, EnvironmentStatusContentPhaseContentPhaseCreating, EnvironmentStatusContentPhaseContentPhaseInitializing, EnvironmentStatusContentPhaseContentPhaseReady, EnvironmentStatusContentPhaseContentPhaseUpdating, EnvironmentStatusContentPhaseContentPhaseFailed:
		return true
	}
	return false
}

// devcontainer contains the status of the devcontainer.
type EnvironmentStatusDevcontainer struct {
	// container_id is the ID of the container.
	ContainerID string `json:"containerId"`
	// container_name is the name of the container that is used to connect to the
	// devcontainer
	ContainerName string `json:"containerName"`
	// devcontainerconfig_in_sync indicates if the devcontainer is up to date w.r.t.
	// the devcontainer config file.
	DevcontainerconfigInSync bool `json:"devcontainerconfigInSync"`
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root
	DevcontainerFilePath string `json:"devcontainerFilePath"`
	// devcontainer_file_presence indicates how the devcontainer file is present in the
	// repo.
	DevcontainerFilePresence EnvironmentStatusDevcontainerDevcontainerFilePresence `json:"devcontainerFilePresence"`
	// failure_message contains the reason the devcontainer failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the devcontainer
	Phase EnvironmentStatusDevcontainerPhase `json:"phase"`
	// remote_user is the user that is used to connect to the devcontainer
	RemoteUser string `json:"remoteUser"`
	// remote_workspace_folder is the folder that is used to connect to the
	// devcontainer
	RemoteWorkspaceFolder string `json:"remoteWorkspaceFolder"`
	// secrets_in_sync indicates if the secrets are up to date w.r.t. the running
	// devcontainer.
	SecretsInSync bool `json:"secretsInSync"`
	// session is the session that is currently active in the devcontainer.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the devcontainer is present but not
	// in the expected state.
	WarningMessage string                            `json:"warningMessage"`
	JSON           environmentStatusDevcontainerJSON `json:"-"`
}

// environmentStatusDevcontainerJSON contains the JSON metadata for the struct
// [EnvironmentStatusDevcontainer]
type environmentStatusDevcontainerJSON struct {
	ContainerID              apijson.Field
	ContainerName            apijson.Field
	DevcontainerconfigInSync apijson.Field
	DevcontainerFilePath     apijson.Field
	DevcontainerFilePresence apijson.Field
	FailureMessage           apijson.Field
	Phase                    apijson.Field
	RemoteUser               apijson.Field
	RemoteWorkspaceFolder    apijson.Field
	SecretsInSync            apijson.Field
	Session                  apijson.Field
	WarningMessage           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EnvironmentStatusDevcontainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusDevcontainerJSON) RawJSON() string {
	return r.raw
}

// devcontainer_file_presence indicates how the devcontainer file is present in the
// repo.
type EnvironmentStatusDevcontainerDevcontainerFilePresence string

const (
	EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceUnspecified EnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_UNSPECIFIED"
	EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceGenerated   EnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_GENERATED"
	EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceDiscovered  EnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_DISCOVERED"
	EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceSpecified   EnvironmentStatusDevcontainerDevcontainerFilePresence = "PRESENCE_SPECIFIED"
)

func (r EnvironmentStatusDevcontainerDevcontainerFilePresence) IsKnown() bool {
	switch r {
	case EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceUnspecified, EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceGenerated, EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceDiscovered, EnvironmentStatusDevcontainerDevcontainerFilePresencePresenceSpecified:
		return true
	}
	return false
}

// phase is the current phase of the devcontainer
type EnvironmentStatusDevcontainerPhase string

const (
	EnvironmentStatusDevcontainerPhasePhaseUnspecified EnvironmentStatusDevcontainerPhase = "PHASE_UNSPECIFIED"
	EnvironmentStatusDevcontainerPhasePhaseCreating    EnvironmentStatusDevcontainerPhase = "PHASE_CREATING"
	EnvironmentStatusDevcontainerPhasePhaseRunning     EnvironmentStatusDevcontainerPhase = "PHASE_RUNNING"
	EnvironmentStatusDevcontainerPhasePhaseStopped     EnvironmentStatusDevcontainerPhase = "PHASE_STOPPED"
	EnvironmentStatusDevcontainerPhasePhaseFailed      EnvironmentStatusDevcontainerPhase = "PHASE_FAILED"
)

func (r EnvironmentStatusDevcontainerPhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusDevcontainerPhasePhaseUnspecified, EnvironmentStatusDevcontainerPhasePhaseCreating, EnvironmentStatusDevcontainerPhasePhaseRunning, EnvironmentStatusDevcontainerPhasePhaseStopped, EnvironmentStatusDevcontainerPhasePhaseFailed:
		return true
	}
	return false
}

// environment_url contains the URL at which the environment can be accessed. This
// field is only set if the environment is running.
type EnvironmentStatusEnvironmentURLs struct {
	// logs is the URL at which the environment logs can be accessed.
	Logs  string                                 `json:"logs"`
	Ports []EnvironmentStatusEnvironmentURLsPort `json:"ports"`
	// SSH is the URL at which the environment can be accessed via SSH.
	SSH  EnvironmentStatusEnvironmentURLsSSH  `json:"ssh"`
	JSON environmentStatusEnvironmentURLsJSON `json:"-"`
}

// environmentStatusEnvironmentURLsJSON contains the JSON metadata for the struct
// [EnvironmentStatusEnvironmentURLs]
type environmentStatusEnvironmentURLsJSON struct {
	Logs        apijson.Field
	Ports       apijson.Field
	SSH         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentStatusEnvironmentURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusEnvironmentURLsJSON) RawJSON() string {
	return r.raw
}

type EnvironmentStatusEnvironmentURLsPort struct {
	// port is the port number of the environment port
	Port int64 `json:"port"`
	// url is the URL at which the environment port can be accessed
	URL  string                                   `json:"url"`
	JSON environmentStatusEnvironmentURLsPortJSON `json:"-"`
}

// environmentStatusEnvironmentURLsPortJSON contains the JSON metadata for the
// struct [EnvironmentStatusEnvironmentURLsPort]
type environmentStatusEnvironmentURLsPortJSON struct {
	Port        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentStatusEnvironmentURLsPort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusEnvironmentURLsPortJSON) RawJSON() string {
	return r.raw
}

// SSH is the URL at which the environment can be accessed via SSH.
type EnvironmentStatusEnvironmentURLsSSH struct {
	URL  string                                  `json:"url"`
	JSON environmentStatusEnvironmentURLsSSHJSON `json:"-"`
}

// environmentStatusEnvironmentURLsSSHJSON contains the JSON metadata for the
// struct [EnvironmentStatusEnvironmentURLsSSH]
type environmentStatusEnvironmentURLsSSHJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentStatusEnvironmentURLsSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusEnvironmentURLsSSHJSON) RawJSON() string {
	return r.raw
}

// machine contains the status of the environment machine
type EnvironmentStatusMachine struct {
	// failure_message contains the reason the machine failed to operate.
	FailureMessage string `json:"failureMessage"`
	// phase is the current phase of the environment machine
	Phase EnvironmentStatusMachinePhase `json:"phase"`
	// session is the session that is currently active in the machine.
	Session string `json:"session"`
	// timeout contains the reason the environment has timed out. If this field is
	// empty, the environment has not timed out.
	Timeout string `json:"timeout"`
	// versions contains the versions of components in the machine.
	Versions EnvironmentStatusMachineVersions `json:"versions"`
	// warning_message contains warnings, e.g. when the machine is present but not in
	// the expected state.
	WarningMessage string                       `json:"warningMessage"`
	JSON           environmentStatusMachineJSON `json:"-"`
}

// environmentStatusMachineJSON contains the JSON metadata for the struct
// [EnvironmentStatusMachine]
type environmentStatusMachineJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	Session        apijson.Field
	Timeout        apijson.Field
	Versions       apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentStatusMachine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusMachineJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the environment machine
type EnvironmentStatusMachinePhase string

const (
	EnvironmentStatusMachinePhasePhaseUnspecified EnvironmentStatusMachinePhase = "PHASE_UNSPECIFIED"
	EnvironmentStatusMachinePhasePhaseCreating    EnvironmentStatusMachinePhase = "PHASE_CREATING"
	EnvironmentStatusMachinePhasePhaseStarting    EnvironmentStatusMachinePhase = "PHASE_STARTING"
	EnvironmentStatusMachinePhasePhaseRunning     EnvironmentStatusMachinePhase = "PHASE_RUNNING"
	EnvironmentStatusMachinePhasePhaseStopping    EnvironmentStatusMachinePhase = "PHASE_STOPPING"
	EnvironmentStatusMachinePhasePhaseStopped     EnvironmentStatusMachinePhase = "PHASE_STOPPED"
	EnvironmentStatusMachinePhasePhaseDeleting    EnvironmentStatusMachinePhase = "PHASE_DELETING"
	EnvironmentStatusMachinePhasePhaseDeleted     EnvironmentStatusMachinePhase = "PHASE_DELETED"
)

func (r EnvironmentStatusMachinePhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusMachinePhasePhaseUnspecified, EnvironmentStatusMachinePhasePhaseCreating, EnvironmentStatusMachinePhasePhaseStarting, EnvironmentStatusMachinePhasePhaseRunning, EnvironmentStatusMachinePhasePhaseStopping, EnvironmentStatusMachinePhasePhaseStopped, EnvironmentStatusMachinePhasePhaseDeleting, EnvironmentStatusMachinePhasePhaseDeleted:
		return true
	}
	return false
}

// versions contains the versions of components in the machine.
type EnvironmentStatusMachineVersions struct {
	SupervisorCommit  string                               `json:"supervisorCommit"`
	SupervisorVersion string                               `json:"supervisorVersion"`
	JSON              environmentStatusMachineVersionsJSON `json:"-"`
}

// environmentStatusMachineVersionsJSON contains the JSON metadata for the struct
// [EnvironmentStatusMachineVersions]
type environmentStatusMachineVersionsJSON struct {
	SupervisorCommit  apijson.Field
	SupervisorVersion apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EnvironmentStatusMachineVersions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusMachineVersionsJSON) RawJSON() string {
	return r.raw
}

// RunnerACK is the acknowledgement from the runner that is has received the
// environment spec.
type EnvironmentStatusRunnerAck struct {
	Message     string                               `json:"message"`
	SpecVersion string                               `json:"specVersion"`
	StatusCode  EnvironmentStatusRunnerAckStatusCode `json:"statusCode"`
	JSON        environmentStatusRunnerAckJSON       `json:"-"`
}

// environmentStatusRunnerAckJSON contains the JSON metadata for the struct
// [EnvironmentStatusRunnerAck]
type environmentStatusRunnerAckJSON struct {
	Message     apijson.Field
	SpecVersion apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentStatusRunnerAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusRunnerAckJSON) RawJSON() string {
	return r.raw
}

type EnvironmentStatusRunnerAckStatusCode string

const (
	EnvironmentStatusRunnerAckStatusCodeStatusCodeUnspecified        EnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_UNSPECIFIED"
	EnvironmentStatusRunnerAckStatusCodeStatusCodeOk                 EnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_OK"
	EnvironmentStatusRunnerAckStatusCodeStatusCodeInvalidResource    EnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_INVALID_RESOURCE"
	EnvironmentStatusRunnerAckStatusCodeStatusCodeFailedPrecondition EnvironmentStatusRunnerAckStatusCode = "STATUS_CODE_FAILED_PRECONDITION"
)

func (r EnvironmentStatusRunnerAckStatusCode) IsKnown() bool {
	switch r {
	case EnvironmentStatusRunnerAckStatusCodeStatusCodeUnspecified, EnvironmentStatusRunnerAckStatusCodeStatusCodeOk, EnvironmentStatusRunnerAckStatusCodeStatusCodeInvalidResource, EnvironmentStatusRunnerAckStatusCodeStatusCodeFailedPrecondition:
		return true
	}
	return false
}

type EnvironmentStatusSecret struct {
	// failure_message contains the reason the secret failed to be materialize.
	FailureMessage string                        `json:"failureMessage"`
	Phase          EnvironmentStatusSecretsPhase `json:"phase"`
	SecretName     string                        `json:"secretName"`
	// session is the session that is currently active in the environment.
	Session string `json:"session"`
	// warning_message contains warnings, e.g. when the secret is present but not in
	// the expected state.
	WarningMessage string                      `json:"warningMessage"`
	JSON           environmentStatusSecretJSON `json:"-"`
}

// environmentStatusSecretJSON contains the JSON metadata for the struct
// [EnvironmentStatusSecret]
type environmentStatusSecretJSON struct {
	FailureMessage apijson.Field
	Phase          apijson.Field
	SecretName     apijson.Field
	Session        apijson.Field
	WarningMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EnvironmentStatusSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusSecretJSON) RawJSON() string {
	return r.raw
}

type EnvironmentStatusSecretsPhase string

const (
	EnvironmentStatusSecretsPhaseContentPhaseUnspecified  EnvironmentStatusSecretsPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentStatusSecretsPhaseContentPhaseCreating     EnvironmentStatusSecretsPhase = "CONTENT_PHASE_CREATING"
	EnvironmentStatusSecretsPhaseContentPhaseInitializing EnvironmentStatusSecretsPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentStatusSecretsPhaseContentPhaseReady        EnvironmentStatusSecretsPhase = "CONTENT_PHASE_READY"
	EnvironmentStatusSecretsPhaseContentPhaseUpdating     EnvironmentStatusSecretsPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentStatusSecretsPhaseContentPhaseFailed       EnvironmentStatusSecretsPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentStatusSecretsPhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusSecretsPhaseContentPhaseUnspecified, EnvironmentStatusSecretsPhaseContentPhaseCreating, EnvironmentStatusSecretsPhaseContentPhaseInitializing, EnvironmentStatusSecretsPhaseContentPhaseReady, EnvironmentStatusSecretsPhaseContentPhaseUpdating, EnvironmentStatusSecretsPhaseContentPhaseFailed:
		return true
	}
	return false
}

type EnvironmentStatusSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID string `json:"id"`
	// phase is the current phase of the public key
	Phase EnvironmentStatusSSHPublicKeysPhase `json:"phase"`
	JSON  environmentStatusSSHPublicKeyJSON   `json:"-"`
}

// environmentStatusSSHPublicKeyJSON contains the JSON metadata for the struct
// [EnvironmentStatusSSHPublicKey]
type environmentStatusSSHPublicKeyJSON struct {
	ID          apijson.Field
	Phase       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentStatusSSHPublicKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentStatusSSHPublicKeyJSON) RawJSON() string {
	return r.raw
}

// phase is the current phase of the public key
type EnvironmentStatusSSHPublicKeysPhase string

const (
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseUnspecified  EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_UNSPECIFIED"
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseCreating     EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_CREATING"
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseInitializing EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_INITIALIZING"
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseReady        EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_READY"
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseUpdating     EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_UPDATING"
	EnvironmentStatusSSHPublicKeysPhaseContentPhaseFailed       EnvironmentStatusSSHPublicKeysPhase = "CONTENT_PHASE_FAILED"
)

func (r EnvironmentStatusSSHPublicKeysPhase) IsKnown() bool {
	switch r {
	case EnvironmentStatusSSHPublicKeysPhaseContentPhaseUnspecified, EnvironmentStatusSSHPublicKeysPhaseContentPhaseCreating, EnvironmentStatusSSHPublicKeysPhaseContentPhaseInitializing, EnvironmentStatusSSHPublicKeysPhaseContentPhaseReady, EnvironmentStatusSSHPublicKeysPhaseContentPhaseUpdating, EnvironmentStatusSSHPublicKeysPhaseContentPhaseFailed:
		return true
	}
	return false
}

type EnvironmentNewResponse struct {
	// +resource get environment
	Environment Environment                `json:"environment"`
	JSON        environmentNewResponseJSON `json:"-"`
}

// environmentNewResponseJSON contains the JSON metadata for the struct
// [EnvironmentNewResponse]
type environmentNewResponseJSON struct {
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentGetResponse struct {
	// +resource get environment
	Environment Environment                `json:"environment"`
	JSON        environmentGetResponseJSON `json:"-"`
}

// environmentGetResponseJSON contains the JSON metadata for the struct
// [EnvironmentGetResponse]
type environmentGetResponseJSON struct {
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentGetResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentUpdateResponse = interface{}

type EnvironmentDeleteResponse = interface{}

type EnvironmentNewFromProjectResponse struct {
	// +resource get environment
	Environment Environment                           `json:"environment"`
	JSON        environmentNewFromProjectResponseJSON `json:"-"`
}

// environmentNewFromProjectResponseJSON contains the JSON metadata for the struct
// [EnvironmentNewFromProjectResponse]
type environmentNewFromProjectResponseJSON struct {
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewFromProjectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewFromProjectResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewLogsTokenResponse struct {
	// access_token is the token that can be used to access the logs of the environment
	AccessToken string                              `json:"accessToken"`
	JSON        environmentNewLogsTokenResponseJSON `json:"-"`
}

// environmentNewLogsTokenResponseJSON contains the JSON metadata for the struct
// [EnvironmentNewLogsTokenResponse]
type environmentNewLogsTokenResponseJSON struct {
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewLogsTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewLogsTokenResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentMarkActiveResponse = interface{}

type EnvironmentStartResponse = interface{}

type EnvironmentStopResponse = interface{}

type EnvironmentNewParams struct {
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec param.Field[EnvironmentSpecParam] `json:"spec"`
}

func (r EnvironmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentGetParams struct {
	// environment_id specifies the environment to get
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParams struct {
	// environment_id specifies which environment should be updated.
	//
	// +required
	EnvironmentID param.Field[string]                      `json:"environmentId" format:"uuid"`
	Metadata      param.Field[interface{}]                 `json:"metadata"`
	Spec          param.Field[EnvironmentUpdateParamsSpec] `json:"spec"`
}

func (r EnvironmentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsSpec struct {
	// automations_file is the automations file spec of the environment
	AutomationsFile param.Field[EnvironmentUpdateParamsSpecAutomationsFile] `json:"automationsFile"`
	Content         param.Field[EnvironmentUpdateParamsSpecContent]         `json:"content"`
	Devcontainer    param.Field[EnvironmentUpdateParamsSpecDevcontainer]    `json:"devcontainer"`
	// ports controls port sharing
	Ports param.Field[[]EnvironmentUpdateParamsSpecPort] `json:"ports"`
	// ssh_public_keys are the public keys to update empty array means nothing to
	// update
	SSHPublicKeys param.Field[[]EnvironmentUpdateParamsSpecSSHPublicKey] `json:"sshPublicKeys"`
	// Timeout configures the environment timeout
	Timeout param.Field[EnvironmentUpdateParamsSpecTimeout] `json:"timeout"`
}

func (r EnvironmentUpdateParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// automations_file is the automations file spec of the environment
type EnvironmentUpdateParamsSpecAutomationsFile struct {
	// automations_file_path is the path to the automations file that is applied in the
	// environment, relative to the repo root. path must not be absolute (start with a
	// /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	AutomationsFilePath param.Field[string] `json:"automationsFilePath"`
	Session             param.Field[string] `json:"session"`
}

func (r EnvironmentUpdateParamsSpecAutomationsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsSpecContent struct {
	// The Git email address
	GitEmail param.Field[string] `json:"gitEmail"`
	// The Git username
	GitUsername param.Field[string] `json:"gitUsername"`
	// EnvironmentInitializer specifies how an environment is to be initialized
	Initializer param.Field[EnvironmentInitializerParam] `json:"initializer"`
	// session should be changed to trigger a content reinitialization
	Session param.Field[string] `json:"session"`
}

func (r EnvironmentUpdateParamsSpecContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsSpecDevcontainer struct {
	// devcontainer_file_path is the path to the devcontainer file relative to the repo
	// root path must not be absolute (start with a /):
	//
	// ```
	// this.matches('^$|^[^/].*')
	// ```
	DevcontainerFilePath param.Field[string] `json:"devcontainerFilePath"`
	// session should be changed to trigger a devcontainer rebuild
	Session param.Field[string] `json:"session"`
}

func (r EnvironmentUpdateParamsSpecDevcontainer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsSpecPort struct {
	// Admission level describes who can access an environment instance and its ports.
	Admission param.Field[AdmissionLevel] `json:"admission"`
	// name of this port
	Name param.Field[string] `json:"name"`
	// port number
	Port param.Field[int64] `json:"port"`
}

func (r EnvironmentUpdateParamsSpecPort) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsSpecSSHPublicKey struct {
	// id is the unique identifier of the public key
	ID param.Field[string] `json:"id"`
	// value is the actual public key in the public key file format if not provided,
	// the public key will be removed
	Value param.Field[string] `json:"value"`
}

func (r EnvironmentUpdateParamsSpecSSHPublicKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Timeout configures the environment timeout
type EnvironmentUpdateParamsSpecTimeout struct {
	// A Duration represents a signed, fixed-length span of time represented as a count
	// of seconds and fractions of seconds at nanosecond resolution. It is independent
	// of any calendar and concepts like "day" or "month". It is related to Timestamp
	// in that the difference between two Timestamp values is a Duration and it can be
	// added or subtracted from a Timestamp. Range is approximately +-10,000 years.
	//
	// # Examples
	//
	// Example 1: Compute Duration from two Timestamps in pseudo code.
	//
	//	Timestamp start = ...;
	//	Timestamp end = ...;
	//	Duration duration = ...;
	//
	//	duration.seconds = end.seconds - start.seconds;
	//	duration.nanos = end.nanos - start.nanos;
	//
	//	if (duration.seconds < 0 && duration.nanos > 0) {
	//	  duration.seconds += 1;
	//	  duration.nanos -= 1000000000;
	//	} else if (duration.seconds > 0 && duration.nanos < 0) {
	//	  duration.seconds -= 1;
	//	  duration.nanos += 1000000000;
	//	}
	//
	// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
	//
	//	Timestamp start = ...;
	//	Duration duration = ...;
	//	Timestamp end = ...;
	//
	//	end.seconds = start.seconds + duration.seconds;
	//	end.nanos = start.nanos + duration.nanos;
	//
	//	if (end.nanos < 0) {
	//	  end.seconds -= 1;
	//	  end.nanos += 1000000000;
	//	} else if (end.nanos >= 1000000000) {
	//	  end.seconds += 1;
	//	  end.nanos -= 1000000000;
	//	}
	//
	// Example 3: Compute Duration from datetime.timedelta in Python.
	//
	//	td = datetime.timedelta(days=3, minutes=10)
	//	duration = Duration()
	//	duration.FromTimedelta(td)
	//
	// # JSON Mapping
	//
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by
	// the number of seconds, with nanoseconds expressed as fractional seconds. For
	// example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as
	// "3.000000001s", and 3 seconds and 1 microsecond should be expressed in JSON
	// format as "3.000001s".
	Disconnected param.Field[string] `json:"disconnected" format:"regex"`
}

func (r EnvironmentUpdateParamsSpecTimeout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentListParams struct {
	Token    param.Field[string]                      `query:"token"`
	PageSize param.Field[int64]                       `query:"pageSize"`
	Filter   param.Field[EnvironmentListParamsFilter] `json:"filter"`
	// organization_id is the ID of the organization that contains the environments
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
	// pagination contains the pagination options for listing environments
	Pagination param.Field[EnvironmentListParamsPagination] `json:"pagination"`
}

func (r EnvironmentListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [EnvironmentListParams]'s query parameters as `url.Values`.
func (r EnvironmentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EnvironmentListParamsFilter struct {
	// creator_ids filters the response to only Environments created by specified
	// members
	CreatorIDs param.Field[[]string] `json:"creatorIds" format:"uuid"`
	// project_ids filters the response to only Environments associated with the
	// specified projects
	ProjectIDs param.Field[[]string] `json:"projectIds" format:"uuid"`
	// runner_ids filters the response to only Environments running on these Runner IDs
	RunnerIDs param.Field[[]string] `json:"runnerIds" format:"uuid"`
	// runner_kinds filters the response to only Environments running on these Runner
	// Kinds
	RunnerKinds param.Field[[]RunnerKind] `json:"runnerKinds"`
	// actual_phases is a list of phases the environment must be in for it to be
	// returned in the API call
	StatusPhases param.Field[[]EnvironmentPhase] `json:"statusPhases"`
}

func (r EnvironmentListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing environments
type EnvironmentListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EnvironmentListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentDeleteParams struct {
	// environment_id specifies the environment that is going to delete.
	//
	// +required
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
	// force indicates whether the environment should be deleted forcefully When force
	// deleting an Environment, the Environment is removed immediately and environment
	// lifecycle is not respected. Force deleting can result in data loss on the
	// environment.
	Force param.Field[bool] `json:"force"`
}

func (r EnvironmentDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewFromProjectParams struct {
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
	// EnvironmentSpec specifies the configuration of an environment for an environment
	// start
	Spec param.Field[EnvironmentSpecParam] `json:"spec"`
}

func (r EnvironmentNewFromProjectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewLogsTokenParams struct {
	// environment_id specifies the environment for which the logs token should be
	// created.
	//
	// +required
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentNewLogsTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentMarkActiveParams struct {
	// EnvironmentActivitySignal used to signal activity for an environment.
	ActivitySignal param.Field[EnvironmentActivitySignalParam] `json:"activitySignal"`
	// The ID of the environment to update activity for.
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentMarkActiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentStartParams struct {
	// environment_id specifies which environment should be started.
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentStopParams struct {
	// environment_id specifies which environment should be stopped.
	//
	// +required
	EnvironmentID param.Field[string] `json:"environmentId" format:"uuid"`
}

func (r EnvironmentStopParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
