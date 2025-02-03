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

// EventService contains methods and other services that help with interacting with
// the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

// ListAuditLogs retrieves a paginated list of audit logs for the specified
// organization
func (r *EventService) List(ctx context.Context, params EventListParams, opts ...option.RequestOption) (res *EventListResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EventService/ListAuditLogs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// WatchEvents streams all requests events to the client
func (r *EventService) Watch(ctx context.Context, params EventWatchParams, opts ...option.RequestOption) (res *EventWatchResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/connect+json")}, opts...)
	path := "gitpod.v1.EventService/WatchEvents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type EventListResponse struct {
	Entries []EventListResponseEntry `json:"entries"`
	// pagination contains the pagination options for listing environments
	Pagination EventListResponsePagination `json:"pagination"`
	JSON       eventListResponseJSON       `json:"-"`
}

// eventListResponseJSON contains the JSON metadata for the struct
// [EventListResponse]
type eventListResponseJSON struct {
	Entries     apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEntry struct {
	ID             string                                 `json:"id"`
	Action         string                                 `json:"action"`
	ActorID        string                                 `json:"actorId"`
	ActorPrincipal EventListResponseEntriesActorPrincipal `json:"actorPrincipal"`
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
	CreatedAt   time.Time                           `json:"createdAt" format:"date-time"`
	SubjectID   string                              `json:"subjectId"`
	SubjectType EventListResponseEntriesSubjectType `json:"subjectType"`
	JSON        eventListResponseEntryJSON          `json:"-"`
}

// eventListResponseEntryJSON contains the JSON metadata for the struct
// [EventListResponseEntry]
type eventListResponseEntryJSON struct {
	ID             apijson.Field
	Action         apijson.Field
	ActorID        apijson.Field
	ActorPrincipal apijson.Field
	CreatedAt      apijson.Field
	SubjectID      apijson.Field
	SubjectType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EventListResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEntryJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEntriesActorPrincipal string

const (
	EventListResponseEntriesActorPrincipalPrincipalUnspecified    EventListResponseEntriesActorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EventListResponseEntriesActorPrincipalPrincipalAccount        EventListResponseEntriesActorPrincipal = "PRINCIPAL_ACCOUNT"
	EventListResponseEntriesActorPrincipalPrincipalUser           EventListResponseEntriesActorPrincipal = "PRINCIPAL_USER"
	EventListResponseEntriesActorPrincipalPrincipalRunner         EventListResponseEntriesActorPrincipal = "PRINCIPAL_RUNNER"
	EventListResponseEntriesActorPrincipalPrincipalEnvironment    EventListResponseEntriesActorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EventListResponseEntriesActorPrincipalPrincipalServiceAccount EventListResponseEntriesActorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EventListResponseEntriesActorPrincipal) IsKnown() bool {
	switch r {
	case EventListResponseEntriesActorPrincipalPrincipalUnspecified, EventListResponseEntriesActorPrincipalPrincipalAccount, EventListResponseEntriesActorPrincipalPrincipalUser, EventListResponseEntriesActorPrincipalPrincipalRunner, EventListResponseEntriesActorPrincipalPrincipalEnvironment, EventListResponseEntriesActorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type EventListResponseEntriesSubjectType string

const (
	EventListResponseEntriesSubjectTypeResourceTypeUnspecified             EventListResponseEntriesSubjectType = "RESOURCE_TYPE_UNSPECIFIED"
	EventListResponseEntriesSubjectTypeResourceTypeEnvironment             EventListResponseEntriesSubjectType = "RESOURCE_TYPE_ENVIRONMENT"
	EventListResponseEntriesSubjectTypeResourceTypeRunner                  EventListResponseEntriesSubjectType = "RESOURCE_TYPE_RUNNER"
	EventListResponseEntriesSubjectTypeResourceTypeProject                 EventListResponseEntriesSubjectType = "RESOURCE_TYPE_PROJECT"
	EventListResponseEntriesSubjectTypeResourceTypeTask                    EventListResponseEntriesSubjectType = "RESOURCE_TYPE_TASK"
	EventListResponseEntriesSubjectTypeResourceTypeTaskExecution           EventListResponseEntriesSubjectType = "RESOURCE_TYPE_TASK_EXECUTION"
	EventListResponseEntriesSubjectTypeResourceTypeService                 EventListResponseEntriesSubjectType = "RESOURCE_TYPE_SERVICE"
	EventListResponseEntriesSubjectTypeResourceTypeOrganization            EventListResponseEntriesSubjectType = "RESOURCE_TYPE_ORGANIZATION"
	EventListResponseEntriesSubjectTypeResourceTypeUser                    EventListResponseEntriesSubjectType = "RESOURCE_TYPE_USER"
	EventListResponseEntriesSubjectTypeResourceTypeEnvironmentClass        EventListResponseEntriesSubjectType = "RESOURCE_TYPE_ENVIRONMENT_CLASS"
	EventListResponseEntriesSubjectTypeResourceTypeRunnerScmIntegration    EventListResponseEntriesSubjectType = "RESOURCE_TYPE_RUNNER_SCM_INTEGRATION"
	EventListResponseEntriesSubjectTypeResourceTypeHostAuthenticationToken EventListResponseEntriesSubjectType = "RESOURCE_TYPE_HOST_AUTHENTICATION_TOKEN"
	EventListResponseEntriesSubjectTypeResourceTypeGroup                   EventListResponseEntriesSubjectType = "RESOURCE_TYPE_GROUP"
	EventListResponseEntriesSubjectTypeResourceTypePersonalAccessToken     EventListResponseEntriesSubjectType = "RESOURCE_TYPE_PERSONAL_ACCESS_TOKEN"
	EventListResponseEntriesSubjectTypeResourceTypeUserPreference          EventListResponseEntriesSubjectType = "RESOURCE_TYPE_USER_PREFERENCE"
	EventListResponseEntriesSubjectTypeResourceTypeServiceAccount          EventListResponseEntriesSubjectType = "RESOURCE_TYPE_SERVICE_ACCOUNT"
	EventListResponseEntriesSubjectTypeResourceTypeSecret                  EventListResponseEntriesSubjectType = "RESOURCE_TYPE_SECRET"
	EventListResponseEntriesSubjectTypeResourceTypeSSOConfig               EventListResponseEntriesSubjectType = "RESOURCE_TYPE_SSO_CONFIG"
)

func (r EventListResponseEntriesSubjectType) IsKnown() bool {
	switch r {
	case EventListResponseEntriesSubjectTypeResourceTypeUnspecified, EventListResponseEntriesSubjectTypeResourceTypeEnvironment, EventListResponseEntriesSubjectTypeResourceTypeRunner, EventListResponseEntriesSubjectTypeResourceTypeProject, EventListResponseEntriesSubjectTypeResourceTypeTask, EventListResponseEntriesSubjectTypeResourceTypeTaskExecution, EventListResponseEntriesSubjectTypeResourceTypeService, EventListResponseEntriesSubjectTypeResourceTypeOrganization, EventListResponseEntriesSubjectTypeResourceTypeUser, EventListResponseEntriesSubjectTypeResourceTypeEnvironmentClass, EventListResponseEntriesSubjectTypeResourceTypeRunnerScmIntegration, EventListResponseEntriesSubjectTypeResourceTypeHostAuthenticationToken, EventListResponseEntriesSubjectTypeResourceTypeGroup, EventListResponseEntriesSubjectTypeResourceTypePersonalAccessToken, EventListResponseEntriesSubjectTypeResourceTypeUserPreference, EventListResponseEntriesSubjectTypeResourceTypeServiceAccount, EventListResponseEntriesSubjectTypeResourceTypeSecret, EventListResponseEntriesSubjectTypeResourceTypeSSOConfig:
		return true
	}
	return false
}

// pagination contains the pagination options for listing environments
type EventListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                          `json:"nextToken"`
	JSON      eventListResponsePaginationJSON `json:"-"`
}

// eventListResponsePaginationJSON contains the JSON metadata for the struct
// [EventListResponsePagination]
type eventListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type EventWatchResponse struct {
	Operation    EventWatchResponseOperation    `json:"operation"`
	ResourceID   string                         `json:"resourceId" format:"uuid"`
	ResourceType EventWatchResponseResourceType `json:"resourceType"`
	JSON         eventWatchResponseJSON         `json:"-"`
}

// eventWatchResponseJSON contains the JSON metadata for the struct
// [EventWatchResponse]
type eventWatchResponseJSON struct {
	Operation    apijson.Field
	ResourceID   apijson.Field
	ResourceType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EventWatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventWatchResponseJSON) RawJSON() string {
	return r.raw
}

type EventWatchResponseOperation string

const (
	EventWatchResponseOperationResourceOperationUnspecified  EventWatchResponseOperation = "RESOURCE_OPERATION_UNSPECIFIED"
	EventWatchResponseOperationResourceOperationCreate       EventWatchResponseOperation = "RESOURCE_OPERATION_CREATE"
	EventWatchResponseOperationResourceOperationUpdate       EventWatchResponseOperation = "RESOURCE_OPERATION_UPDATE"
	EventWatchResponseOperationResourceOperationDelete       EventWatchResponseOperation = "RESOURCE_OPERATION_DELETE"
	EventWatchResponseOperationResourceOperationUpdateStatus EventWatchResponseOperation = "RESOURCE_OPERATION_UPDATE_STATUS"
)

func (r EventWatchResponseOperation) IsKnown() bool {
	switch r {
	case EventWatchResponseOperationResourceOperationUnspecified, EventWatchResponseOperationResourceOperationCreate, EventWatchResponseOperationResourceOperationUpdate, EventWatchResponseOperationResourceOperationDelete, EventWatchResponseOperationResourceOperationUpdateStatus:
		return true
	}
	return false
}

type EventWatchResponseResourceType string

const (
	EventWatchResponseResourceTypeResourceTypeUnspecified             EventWatchResponseResourceType = "RESOURCE_TYPE_UNSPECIFIED"
	EventWatchResponseResourceTypeResourceTypeEnvironment             EventWatchResponseResourceType = "RESOURCE_TYPE_ENVIRONMENT"
	EventWatchResponseResourceTypeResourceTypeRunner                  EventWatchResponseResourceType = "RESOURCE_TYPE_RUNNER"
	EventWatchResponseResourceTypeResourceTypeProject                 EventWatchResponseResourceType = "RESOURCE_TYPE_PROJECT"
	EventWatchResponseResourceTypeResourceTypeTask                    EventWatchResponseResourceType = "RESOURCE_TYPE_TASK"
	EventWatchResponseResourceTypeResourceTypeTaskExecution           EventWatchResponseResourceType = "RESOURCE_TYPE_TASK_EXECUTION"
	EventWatchResponseResourceTypeResourceTypeService                 EventWatchResponseResourceType = "RESOURCE_TYPE_SERVICE"
	EventWatchResponseResourceTypeResourceTypeOrganization            EventWatchResponseResourceType = "RESOURCE_TYPE_ORGANIZATION"
	EventWatchResponseResourceTypeResourceTypeUser                    EventWatchResponseResourceType = "RESOURCE_TYPE_USER"
	EventWatchResponseResourceTypeResourceTypeEnvironmentClass        EventWatchResponseResourceType = "RESOURCE_TYPE_ENVIRONMENT_CLASS"
	EventWatchResponseResourceTypeResourceTypeRunnerScmIntegration    EventWatchResponseResourceType = "RESOURCE_TYPE_RUNNER_SCM_INTEGRATION"
	EventWatchResponseResourceTypeResourceTypeHostAuthenticationToken EventWatchResponseResourceType = "RESOURCE_TYPE_HOST_AUTHENTICATION_TOKEN"
	EventWatchResponseResourceTypeResourceTypeGroup                   EventWatchResponseResourceType = "RESOURCE_TYPE_GROUP"
	EventWatchResponseResourceTypeResourceTypePersonalAccessToken     EventWatchResponseResourceType = "RESOURCE_TYPE_PERSONAL_ACCESS_TOKEN"
	EventWatchResponseResourceTypeResourceTypeUserPreference          EventWatchResponseResourceType = "RESOURCE_TYPE_USER_PREFERENCE"
	EventWatchResponseResourceTypeResourceTypeServiceAccount          EventWatchResponseResourceType = "RESOURCE_TYPE_SERVICE_ACCOUNT"
	EventWatchResponseResourceTypeResourceTypeSecret                  EventWatchResponseResourceType = "RESOURCE_TYPE_SECRET"
	EventWatchResponseResourceTypeResourceTypeSSOConfig               EventWatchResponseResourceType = "RESOURCE_TYPE_SSO_CONFIG"
)

func (r EventWatchResponseResourceType) IsKnown() bool {
	switch r {
	case EventWatchResponseResourceTypeResourceTypeUnspecified, EventWatchResponseResourceTypeResourceTypeEnvironment, EventWatchResponseResourceTypeResourceTypeRunner, EventWatchResponseResourceTypeResourceTypeProject, EventWatchResponseResourceTypeResourceTypeTask, EventWatchResponseResourceTypeResourceTypeTaskExecution, EventWatchResponseResourceTypeResourceTypeService, EventWatchResponseResourceTypeResourceTypeOrganization, EventWatchResponseResourceTypeResourceTypeUser, EventWatchResponseResourceTypeResourceTypeEnvironmentClass, EventWatchResponseResourceTypeResourceTypeRunnerScmIntegration, EventWatchResponseResourceTypeResourceTypeHostAuthenticationToken, EventWatchResponseResourceTypeResourceTypeGroup, EventWatchResponseResourceTypeResourceTypePersonalAccessToken, EventWatchResponseResourceTypeResourceTypeUserPreference, EventWatchResponseResourceTypeResourceTypeServiceAccount, EventWatchResponseResourceTypeResourceTypeSecret, EventWatchResponseResourceTypeResourceTypeSSOConfig:
		return true
	}
	return false
}

type EventListParams struct {
	// Define which encoding or 'Message-Codec' to use
	Encoding param.Field[EventListParamsEncoding] `query:"encoding,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EventListParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Specifies if the message query param is base64 encoded, which may be required
	// for binary data
	Base64 param.Field[bool] `query:"base64"`
	// Which compression algorithm to use for this request
	Compression param.Field[EventListParamsCompression] `query:"compression"`
	// Define the version of the Connect protocol
	Connect param.Field[EventListParamsConnect] `query:"connect"`
	Message param.Field[string]                 `query:"message"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Define which encoding or 'Message-Codec' to use
type EventListParamsEncoding string

const (
	EventListParamsEncodingProto EventListParamsEncoding = "proto"
	EventListParamsEncodingJson  EventListParamsEncoding = "json"
)

func (r EventListParamsEncoding) IsKnown() bool {
	switch r {
	case EventListParamsEncodingProto, EventListParamsEncodingJson:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EventListParamsConnectProtocolVersion float64

const (
	EventListParamsConnectProtocolVersion1 EventListParamsConnectProtocolVersion = 1
)

func (r EventListParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EventListParamsConnectProtocolVersion1:
		return true
	}
	return false
}

// Which compression algorithm to use for this request
type EventListParamsCompression string

const (
	EventListParamsCompressionIdentity EventListParamsCompression = "identity"
	EventListParamsCompressionGzip     EventListParamsCompression = "gzip"
	EventListParamsCompressionBr       EventListParamsCompression = "br"
)

func (r EventListParamsCompression) IsKnown() bool {
	switch r {
	case EventListParamsCompressionIdentity, EventListParamsCompressionGzip, EventListParamsCompressionBr:
		return true
	}
	return false
}

// Define the version of the Connect protocol
type EventListParamsConnect string

const (
	EventListParamsConnectV1 EventListParamsConnect = "v1"
)

func (r EventListParamsConnect) IsKnown() bool {
	switch r {
	case EventListParamsConnectV1:
		return true
	}
	return false
}

type EventWatchParams struct {
	Body EventWatchParamsBodyUnion `json:"body,required"`
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[EventWatchParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r EventWatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type EventWatchParamsBody struct {
	// Environment scope produces events for the environment itself, all tasks, task
	// executions,
	//
	// and services associated with that environment.
	EnvironmentID param.Field[string] `json:"environmentId"`
	// Organization scope produces events for all projects, runners and environments
	//
	// the caller can see within their organization. No task, task execution or service
	// events are produed.
	Organization param.Field[bool] `json:"organization"`
}

func (r EventWatchParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventWatchParamsBody) implementsEventWatchParamsBodyUnion() {}

// Satisfied by [EventWatchParamsBodyObject], [EventWatchParamsBodyObject],
// [EventWatchParamsBodyObject], [EventWatchParamsBody].
type EventWatchParamsBodyUnion interface {
	implementsEventWatchParamsBodyUnion()
}

type EventWatchParamsBodyObject struct {
	// Environment scope produces events for the environment itself, all tasks, task
	// executions,
	//
	// and services associated with that environment.
	EnvironmentID param.Field[string] `json:"environmentId,required"`
	// Organization scope produces events for all projects, runners and environments
	//
	// the caller can see within their organization. No task, task execution or service
	// events are produed.
	Organization param.Field[bool] `json:"organization"`
}

func (r EventWatchParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventWatchParamsBodyObject) implementsEventWatchParamsBodyUnion() {}

// Define the version of the Connect protocol
type EventWatchParamsConnectProtocolVersion float64

const (
	EventWatchParamsConnectProtocolVersion1 EventWatchParamsConnectProtocolVersion = 1
)

func (r EventWatchParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case EventWatchParamsConnectProtocolVersion1:
		return true
	}
	return false
}
