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
	"github.com/stainless-sdks/gitpod-go/packages/jsonl"
	"github.com/stainless-sdks/gitpod-go/packages/pagination"
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
func (r *EventService) List(ctx context.Context, params EventListParams, opts ...option.RequestOption) (res *pagination.PersonalAccessTokensPage[EventListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.EventService/ListAuditLogs"
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

// ListAuditLogs retrieves a paginated list of audit logs for the specified
// organization
func (r *EventService) ListAutoPaging(ctx context.Context, params EventListParams, opts ...option.RequestOption) *pagination.PersonalAccessTokensPageAutoPager[EventListResponse] {
	return pagination.NewPersonalAccessTokensPageAutoPager(r.List(ctx, params, opts...))
}

// WatchEvents streams all requests events to the client
func (r *EventService) WatchStreaming(ctx context.Context, body EventWatchParams, opts ...option.RequestOption) (stream *jsonl.Stream[EventWatchResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/jsonl")}, opts...)
	path := "gitpod.v1.EventService/WatchEvents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return jsonl.NewStream[EventWatchResponse](raw, err)
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
	// Token passed for retreiving the next set of results. Empty if there are no more
	// results
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
	Token    param.Field[string]                `query:"token"`
	PageSize param.Field[int64]                 `query:"pageSize"`
	Filter   param.Field[EventListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing environments
	Pagination param.Field[EventListParamsPagination] `json:"pagination"`
}

func (r EventListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventListParamsFilter struct {
	ActorIDs        param.Field[[]string]                              `json:"actorIds" format:"uuid"`
	ActorPrincipals param.Field[[]EventListParamsFilterActorPrincipal] `json:"actorPrincipals"`
	SubjectIDs      param.Field[[]string]                              `json:"subjectIds" format:"uuid"`
	SubjectTypes    param.Field[[]EventListParamsFilterSubjectType]    `json:"subjectTypes"`
}

func (r EventListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventListParamsFilterActorPrincipal string

const (
	EventListParamsFilterActorPrincipalPrincipalUnspecified    EventListParamsFilterActorPrincipal = "PRINCIPAL_UNSPECIFIED"
	EventListParamsFilterActorPrincipalPrincipalAccount        EventListParamsFilterActorPrincipal = "PRINCIPAL_ACCOUNT"
	EventListParamsFilterActorPrincipalPrincipalUser           EventListParamsFilterActorPrincipal = "PRINCIPAL_USER"
	EventListParamsFilterActorPrincipalPrincipalRunner         EventListParamsFilterActorPrincipal = "PRINCIPAL_RUNNER"
	EventListParamsFilterActorPrincipalPrincipalEnvironment    EventListParamsFilterActorPrincipal = "PRINCIPAL_ENVIRONMENT"
	EventListParamsFilterActorPrincipalPrincipalServiceAccount EventListParamsFilterActorPrincipal = "PRINCIPAL_SERVICE_ACCOUNT"
)

func (r EventListParamsFilterActorPrincipal) IsKnown() bool {
	switch r {
	case EventListParamsFilterActorPrincipalPrincipalUnspecified, EventListParamsFilterActorPrincipalPrincipalAccount, EventListParamsFilterActorPrincipalPrincipalUser, EventListParamsFilterActorPrincipalPrincipalRunner, EventListParamsFilterActorPrincipalPrincipalEnvironment, EventListParamsFilterActorPrincipalPrincipalServiceAccount:
		return true
	}
	return false
}

type EventListParamsFilterSubjectType string

const (
	EventListParamsFilterSubjectTypeResourceTypeUnspecified             EventListParamsFilterSubjectType = "RESOURCE_TYPE_UNSPECIFIED"
	EventListParamsFilterSubjectTypeResourceTypeEnvironment             EventListParamsFilterSubjectType = "RESOURCE_TYPE_ENVIRONMENT"
	EventListParamsFilterSubjectTypeResourceTypeRunner                  EventListParamsFilterSubjectType = "RESOURCE_TYPE_RUNNER"
	EventListParamsFilterSubjectTypeResourceTypeProject                 EventListParamsFilterSubjectType = "RESOURCE_TYPE_PROJECT"
	EventListParamsFilterSubjectTypeResourceTypeTask                    EventListParamsFilterSubjectType = "RESOURCE_TYPE_TASK"
	EventListParamsFilterSubjectTypeResourceTypeTaskExecution           EventListParamsFilterSubjectType = "RESOURCE_TYPE_TASK_EXECUTION"
	EventListParamsFilterSubjectTypeResourceTypeService                 EventListParamsFilterSubjectType = "RESOURCE_TYPE_SERVICE"
	EventListParamsFilterSubjectTypeResourceTypeOrganization            EventListParamsFilterSubjectType = "RESOURCE_TYPE_ORGANIZATION"
	EventListParamsFilterSubjectTypeResourceTypeUser                    EventListParamsFilterSubjectType = "RESOURCE_TYPE_USER"
	EventListParamsFilterSubjectTypeResourceTypeEnvironmentClass        EventListParamsFilterSubjectType = "RESOURCE_TYPE_ENVIRONMENT_CLASS"
	EventListParamsFilterSubjectTypeResourceTypeRunnerScmIntegration    EventListParamsFilterSubjectType = "RESOURCE_TYPE_RUNNER_SCM_INTEGRATION"
	EventListParamsFilterSubjectTypeResourceTypeHostAuthenticationToken EventListParamsFilterSubjectType = "RESOURCE_TYPE_HOST_AUTHENTICATION_TOKEN"
	EventListParamsFilterSubjectTypeResourceTypeGroup                   EventListParamsFilterSubjectType = "RESOURCE_TYPE_GROUP"
	EventListParamsFilterSubjectTypeResourceTypePersonalAccessToken     EventListParamsFilterSubjectType = "RESOURCE_TYPE_PERSONAL_ACCESS_TOKEN"
	EventListParamsFilterSubjectTypeResourceTypeUserPreference          EventListParamsFilterSubjectType = "RESOURCE_TYPE_USER_PREFERENCE"
	EventListParamsFilterSubjectTypeResourceTypeServiceAccount          EventListParamsFilterSubjectType = "RESOURCE_TYPE_SERVICE_ACCOUNT"
	EventListParamsFilterSubjectTypeResourceTypeSecret                  EventListParamsFilterSubjectType = "RESOURCE_TYPE_SECRET"
	EventListParamsFilterSubjectTypeResourceTypeSSOConfig               EventListParamsFilterSubjectType = "RESOURCE_TYPE_SSO_CONFIG"
)

func (r EventListParamsFilterSubjectType) IsKnown() bool {
	switch r {
	case EventListParamsFilterSubjectTypeResourceTypeUnspecified, EventListParamsFilterSubjectTypeResourceTypeEnvironment, EventListParamsFilterSubjectTypeResourceTypeRunner, EventListParamsFilterSubjectTypeResourceTypeProject, EventListParamsFilterSubjectTypeResourceTypeTask, EventListParamsFilterSubjectTypeResourceTypeTaskExecution, EventListParamsFilterSubjectTypeResourceTypeService, EventListParamsFilterSubjectTypeResourceTypeOrganization, EventListParamsFilterSubjectTypeResourceTypeUser, EventListParamsFilterSubjectTypeResourceTypeEnvironmentClass, EventListParamsFilterSubjectTypeResourceTypeRunnerScmIntegration, EventListParamsFilterSubjectTypeResourceTypeHostAuthenticationToken, EventListParamsFilterSubjectTypeResourceTypeGroup, EventListParamsFilterSubjectTypeResourceTypePersonalAccessToken, EventListParamsFilterSubjectTypeResourceTypeUserPreference, EventListParamsFilterSubjectTypeResourceTypeServiceAccount, EventListParamsFilterSubjectTypeResourceTypeSecret, EventListParamsFilterSubjectTypeResourceTypeSSOConfig:
		return true
	}
	return false
}

// pagination contains the pagination options for listing environments
type EventListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EventListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventWatchParams struct {
	Body EventWatchParamsBodyUnion `json:"body,required"`
}

func (r EventWatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type EventWatchParamsBody struct {
	// Environment scope produces events for the environment itself, all tasks, task
	// executions, and services associated with that environment.
	EnvironmentID param.Field[string] `json:"environmentId"`
	// Organization scope produces events for all projects, runners and environments
	// the caller can see within their organization. No task, task execution or service
	// events are produed.
	Organization param.Field[bool] `json:"organization"`
}

func (r EventWatchParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventWatchParamsBody) implementsEventWatchParamsBodyUnion() {}

// Satisfied by
// [EventWatchParamsBodyEnvironmentScopeProducesEventsForTheEnvironmentItselfAllTasksTaskExecutionsAndServicesAssociatedWithThatEnvironment],
// [EventWatchParamsBodyOrganizationScopeProducesEventsForAllProjectsRunnersAndEnvironmentsTheCallerCanSeeWithinTheirOrganizationNoTaskTaskExecutionOrServiceEventsAreProdued],
// [EventWatchParamsBody].
type EventWatchParamsBodyUnion interface {
	implementsEventWatchParamsBodyUnion()
}

type EventWatchParamsBodyEnvironmentScopeProducesEventsForTheEnvironmentItselfAllTasksTaskExecutionsAndServicesAssociatedWithThatEnvironment struct {
	// Environment scope produces events for the environment itself, all tasks, task
	// executions, and services associated with that environment.
	EnvironmentID param.Field[string] `json:"environmentId,required"`
}

func (r EventWatchParamsBodyEnvironmentScopeProducesEventsForTheEnvironmentItselfAllTasksTaskExecutionsAndServicesAssociatedWithThatEnvironment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventWatchParamsBodyEnvironmentScopeProducesEventsForTheEnvironmentItselfAllTasksTaskExecutionsAndServicesAssociatedWithThatEnvironment) implementsEventWatchParamsBodyUnion() {
}

type EventWatchParamsBodyOrganizationScopeProducesEventsForAllProjectsRunnersAndEnvironmentsTheCallerCanSeeWithinTheirOrganizationNoTaskTaskExecutionOrServiceEventsAreProdued struct {
	// Organization scope produces events for all projects, runners and environments
	// the caller can see within their organization. No task, task execution or service
	// events are produed.
	Organization param.Field[bool] `json:"organization,required"`
}

func (r EventWatchParamsBodyOrganizationScopeProducesEventsForAllProjectsRunnersAndEnvironmentsTheCallerCanSeeWithinTheirOrganizationNoTaskTaskExecutionOrServiceEventsAreProdued) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventWatchParamsBodyOrganizationScopeProducesEventsForAllProjectsRunnersAndEnvironmentsTheCallerCanSeeWithinTheirOrganizationNoTaskTaskExecutionOrServiceEventsAreProdued) implementsEventWatchParamsBodyUnion() {
}
