// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/apiquery"
	"github.com/gitpod-io/gitpod-sdk-go/internal/param"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
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

// Creates a new secret for a project.
//
// Use this method to:
//
// - Store sensitive configuration values
// - Set up environment variables
// - Configure registry authentication
// - Add file-based secrets
//
// ### Examples
//
// - Create environment variable:
//
//	Creates a secret that will be available as an environment variable.
//
//	```yaml
//	name: "DATABASE_URL"
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	value: "postgresql://user:pass@localhost:5432/db"
//	environmentVariable: true
//	```
//
// - Create file secret:
//
//	Creates a secret that will be mounted as a file.
//
//	```yaml
//	name: "SSH_KEY"
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	value: "-----BEGIN RSA PRIVATE KEY-----\n..."
//	filePath: "/home/gitpod/.ssh/id_rsa"
//	```
//
// - Create registry auth:
//
//	Creates credentials for private container registry.
//
//	```yaml
//	name: "DOCKER_AUTH"
//	projectId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	value: "username:password"
//	containerRegistryBasicAuthHost: "https://registry.example.com"
//	```
func (r *SecretService) New(ctx context.Context, body SecretNewParams, opts ...option.RequestOption) (res *SecretNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/CreateSecret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListSecrets
func (r *SecretService) List(ctx context.Context, params SecretListParams, opts ...option.RequestOption) (res *pagination.SecretsPage[Secret], err error) {
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

// ListSecrets
func (r *SecretService) ListAutoPaging(ctx context.Context, params SecretListParams, opts ...option.RequestOption) *pagination.SecretsPageAutoPager[Secret] {
	return pagination.NewSecretsPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a secret permanently.
//
// Use this method to:
//
// - Remove unused secrets
// - Clean up old credentials
//
// ### Examples
//
// - Delete secret:
//
//	Permanently removes a secret.
//
//	```yaml
//	secretId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *SecretService) Delete(ctx context.Context, body SecretDeleteParams, opts ...option.RequestOption) (res *SecretDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/DeleteSecret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets the value of a secret. Only available to environments that are authorized
// to access the secret.
//
// Use this method to:
//
// - Retrieve secret values
// - Access credentials
//
// ### Examples
//
// - Get secret value:
//
//	Retrieves the value of a specific secret.
//
//	```yaml
//	secretId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *SecretService) GetValue(ctx context.Context, body SecretGetValueParams, opts ...option.RequestOption) (res *SecretGetValueResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/GetSecretValue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates the value of an existing secret.
//
// Use this method to:
//
// - Rotate secret values
// - Update credentials
//
// ### Examples
//
// - Update secret value:
//
//	Changes the value of an existing secret.
//
//	```yaml
//	secretId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	value: "new-secret-value"
//	```
func (r *SecretService) UpdateValue(ctx context.Context, body SecretUpdateValueParams, opts ...option.RequestOption) (res *SecretUpdateValueResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.SecretService/UpdateSecretValue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Secret struct {
	ID string `json:"id" format:"uuid"`
	// secret will be mounted as a registry secret
	ContainerRegistryBasicAuthHost string `json:"containerRegistryBasicAuthHost" format:"uri"`
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
	Creator shared.Subject `json:"creator"`
	// secret will be created as an Environment Variable with the same name as the
	// secret
	EnvironmentVariable bool `json:"environmentVariable"`
	// absolute path to the file where the secret is mounted
	FilePath string `json:"filePath"`
	// Name of the secret for humans.
	Name string `json:"name"`
	// The Project ID this Secret belongs to Deprecated: use scope instead
	//
	// Deprecated: deprecated
	ProjectID string      `json:"projectId" format:"uuid"`
	Scope     SecretScope `json:"scope"`
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
	UpdatedAt time.Time  `json:"updatedAt" format:"date-time"`
	JSON      secretJSON `json:"-"`
}

// secretJSON contains the JSON metadata for the struct [Secret]
type secretJSON struct {
	ID                             apijson.Field
	ContainerRegistryBasicAuthHost apijson.Field
	CreatedAt                      apijson.Field
	Creator                        apijson.Field
	EnvironmentVariable            apijson.Field
	FilePath                       apijson.Field
	Name                           apijson.Field
	ProjectID                      apijson.Field
	Scope                          apijson.Field
	UpdatedAt                      apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *Secret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretJSON) RawJSON() string {
	return r.raw
}

type SecretScope struct {
	// project_id is the Project ID this Secret belongs to
	ProjectID string `json:"projectId" format:"uuid"`
	// user_id is the User ID this Secret belongs to
	UserID string          `json:"userId" format:"uuid"`
	JSON   secretScopeJSON `json:"-"`
}

// secretScopeJSON contains the JSON metadata for the struct [SecretScope]
type secretScopeJSON struct {
	ProjectID   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretScopeJSON) RawJSON() string {
	return r.raw
}

type SecretScopeParam struct {
	// project_id is the Project ID this Secret belongs to
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
	// user_id is the User ID this Secret belongs to
	UserID param.Field[string] `json:"userId" format:"uuid"`
}

func (r SecretScopeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecretNewResponse struct {
	Secret Secret                `json:"secret"`
	JSON   secretNewResponseJSON `json:"-"`
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
	// secret will be mounted as a docker config in the environment VM, mount will have
	// the docker registry host
	ContainerRegistryBasicAuthHost param.Field[string] `json:"containerRegistryBasicAuthHost"`
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
	// project_id is the ProjectID this Secret belongs to Deprecated: use scope instead
	ProjectID param.Field[string] `json:"projectId"`
	// scope is the scope of the secret
	Scope param.Field[SecretScopeParam] `json:"scope"`
	// value is the plaintext value of the secret
	Value param.Field[string] `json:"value"`
}

func (r SecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// Deprecated: use scope instead. Values in project_ids will be ignored.
	//
	// Deprecated: deprecated
	ProjectIDs param.Field[[]string] `json:"projectIds" format:"uuid"`
	// scope is the scope of the secrets to list
	Scope param.Field[SecretScopeParam] `json:"scope"`
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
