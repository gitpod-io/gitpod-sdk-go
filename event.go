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
	"github.com/gitpod-io/gitpod-sdk-go/packages/jsonl"
	"github.com/gitpod-io/gitpod-sdk-go/packages/pagination"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
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

// Gets one audit-log entry, including any typed details stored for it.
//
// Use this method to:
//
// - Inspect the details of a specific audit-log entry
// - Retrieve the evidence associated with a Veto Exec audit event
//
// ### Examples
//
// - Get an audit-log entry:
//
//	```yaml
//	auditLogEntryId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *EventService) Get(ctx context.Context, body EventGetParams, opts ...option.RequestOption) (res *EventGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.EventService/GetAuditLog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists audit logs with filtering and pagination options.
//
// Use this method to:
//
// - View audit history
// - Track user actions
// - Monitor system changes
//
// ### Examples
//
// - List all logs:
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter by actor:
//
//	```yaml
//	filter:
//	  actorIds: ["d2c94c27-3b76-4a42-b88c-95a85e392c68"]
//	  actorPrincipals: ["PRINCIPAL_USER"]
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter by time range:
//
//	```yaml
//	filter:
//	  from: "2024-01-01T00:00:00Z"
//	  to: "2024-02-01T00:00:00Z"
//	pagination:
//	  pageSize: 20
//	```
func (r *EventService) List(ctx context.Context, params EventListParams, opts ...option.RequestOption) (res *pagination.EntriesPage[EventListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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

// Lists audit logs with filtering and pagination options.
//
// Use this method to:
//
// - View audit history
// - Track user actions
// - Monitor system changes
//
// ### Examples
//
// - List all logs:
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter by actor:
//
//	```yaml
//	filter:
//	  actorIds: ["d2c94c27-3b76-4a42-b88c-95a85e392c68"]
//	  actorPrincipals: ["PRINCIPAL_USER"]
//	pagination:
//	  pageSize: 20
//	```
//
// - Filter by time range:
//
//	```yaml
//	filter:
//	  from: "2024-01-01T00:00:00Z"
//	  to: "2024-02-01T00:00:00Z"
//	pagination:
//	  pageSize: 20
//	```
func (r *EventService) ListAutoPaging(ctx context.Context, params EventListParams, opts ...option.RequestOption) *pagination.EntriesPageAutoPager[EventListResponse] {
	return pagination.NewEntriesPageAutoPager(r.List(ctx, params, opts...))
}

// Streams events for all projects, runners, environments, tasks, and services
// based on the specified scope.
//
// Use this method to:
//
// - Monitor resource changes in real-time
// - Track system events
// - Receive notifications
//
// The scope parameter determines which events to watch:
//
//   - Organization scope (default): Watch all organization-wide events including
//     projects, runners and environments. Task and service events are not included.
//     Use by setting organization=true or omitting the scope.
//   - Environment scope: Watch events for a specific environment, including its
//     tasks, task executions, and services. Use by setting environment_id to the
//     UUID of the environment to watch.
func (r *EventService) WatchStreaming(ctx context.Context, body EventWatchParams, opts ...option.RequestOption) (stream *jsonl.Stream[EventWatchResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/jsonl")}, opts...)
	path := "gitpod.v1.EventService/WatchEvents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return jsonl.NewStream[EventWatchResponse](raw, err)
}

// AuditLogEntryDetails contains the typed evidence stored with an audit-log entry.
type AuditLogEntryDetails struct {
	// veto_exec contains Veto Exec event details without process.cmdline.
	VetoExec AuditLogEntryDetailsVetoExec `json:"vetoExec" api:"required"`
	JSON     auditLogEntryDetailsJSON     `json:"-"`
}

// auditLogEntryDetailsJSON contains the JSON metadata for the struct
// [AuditLogEntryDetails]
type auditLogEntryDetailsJSON struct {
	VetoExec    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogEntryDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogEntryDetailsJSON) RawJSON() string {
	return r.raw
}

// veto_exec contains Veto Exec event details without process.cmdline.
type AuditLogEntryDetailsVetoExec struct {
	// process contains metadata about the process that triggered the event.
	Process Process `json:"process" api:"required"`
	// timestamp is when the event occurred in the environment.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// action is the enforcement action taken (block or audit).
	Action shared.KernelControlsAction `json:"action"`
	// environment_id is the environment where the event occurred.
	EnvironmentID string `json:"environmentId" format:"uuid"`
	// executable is the digest of the binary content (e.g., "sha256:a1b2c3d4..."). 256
	// allows for longer hash algorithms or prefixed identifiers. May be empty when the
	// event source cannot compute the hash.
	Executable string `json:"executable"`
	// filename is the kernel-resolved path of the binary. Kernel PATH_MAX = 4096
	// (include/uapi/linux/limits.h). May be empty if the event source could not
	// resolve it.
	Filename string                           `json:"filename"`
	JSON     auditLogEntryDetailsVetoExecJSON `json:"-"`
}

// auditLogEntryDetailsVetoExecJSON contains the JSON metadata for the struct
// [AuditLogEntryDetailsVetoExec]
type auditLogEntryDetailsVetoExecJSON struct {
	Process       apijson.Field
	Timestamp     apijson.Field
	Action        apijson.Field
	EnvironmentID apijson.Field
	Executable    apijson.Field
	Filename      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AuditLogEntryDetailsVetoExec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogEntryDetailsVetoExecJSON) RawJSON() string {
	return r.raw
}

// Process describes process metadata for a security event.
//
// PID fields use int32 to match the kernel's pid_t (signed int). Linux PID max is
// 4,194,304 (2^22), well within int32 range. Postgres has no unsigned integer
// type: Ent maps uint32 to bigint (8 bytes) while int32 maps to integer (4 bytes).
// Using int32 aligns proto, Go, and Postgres types without wasting storage.
type Process struct {
	// name is the process name (comm). 2x kernel TASK_COMM_LEN=16
	Name string `json:"name"`
	// pgid is the process group ID.
	Pgid int64 `json:"pgid"`
	// pid is the userspace process ID (kernel thread group ID, tgid).
	Pid int64 `json:"pid"`
	// ppid is the parent process ID.
	Ppid int64 `json:"ppid"`
	// sid is the session ID.
	Sid int64 `json:"sid"`
	// started_at is when the process started.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// tid is the userspace thread ID (kernel pid).
	Tid  int64       `json:"tid"`
	JSON processJSON `json:"-"`
}

// processJSON contains the JSON metadata for the struct [Process]
type processJSON struct {
	Name        apijson.Field
	Pgid        apijson.Field
	Pid         apijson.Field
	Ppid        apijson.Field
	Sid         apijson.Field
	StartedAt   apijson.Field
	Tid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Process) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r processJSON) RawJSON() string {
	return r.raw
}

type ResourceOperation string

const (
	ResourceOperationUnspecified  ResourceOperation = "RESOURCE_OPERATION_UNSPECIFIED"
	ResourceOperationCreate       ResourceOperation = "RESOURCE_OPERATION_CREATE"
	ResourceOperationUpdate       ResourceOperation = "RESOURCE_OPERATION_UPDATE"
	ResourceOperationDelete       ResourceOperation = "RESOURCE_OPERATION_DELETE"
	ResourceOperationUpdateStatus ResourceOperation = "RESOURCE_OPERATION_UPDATE_STATUS"
)

func (r ResourceOperation) IsKnown() bool {
	switch r {
	case ResourceOperationUnspecified, ResourceOperationCreate, ResourceOperationUpdate, ResourceOperationDelete, ResourceOperationUpdateStatus:
		return true
	}
	return false
}

type EventGetResponse struct {
	// entry contains the common audit-log fields also returned by ListAuditLogs.
	Entry EventGetResponseEntry `json:"entry" api:"required"`
	// details contains typed evidence captured with the audit entry. It is absent when
	// the entry has no supported, valid details.
	Details AuditLogEntryDetails `json:"details"`
	JSON    eventGetResponseJSON `json:"-"`
}

// eventGetResponseJSON contains the JSON metadata for the struct
// [EventGetResponse]
type eventGetResponseJSON struct {
	Entry       apijson.Field
	Details     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventGetResponseJSON) RawJSON() string {
	return r.raw
}

// entry contains the common audit-log fields also returned by ListAuditLogs.
type EventGetResponseEntry struct {
	ID             string           `json:"id"`
	Action         string           `json:"action"`
	ActorID        string           `json:"actorId"`
	ActorPrincipal shared.Principal `json:"actorPrincipal"`
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
	// AuditLogEntryKind identifies the coarse query and rendering family of an
	// audit-log entry.
	Kind        EventGetResponseEntryKind `json:"kind"`
	SubjectID   string                    `json:"subjectId"`
	SubjectType shared.ResourceType       `json:"subjectType"`
	JSON        eventGetResponseEntryJSON `json:"-"`
}

// eventGetResponseEntryJSON contains the JSON metadata for the struct
// [EventGetResponseEntry]
type eventGetResponseEntryJSON struct {
	ID             apijson.Field
	Action         apijson.Field
	ActorID        apijson.Field
	ActorPrincipal apijson.Field
	CreatedAt      apijson.Field
	Kind           apijson.Field
	SubjectID      apijson.Field
	SubjectType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EventGetResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventGetResponseEntryJSON) RawJSON() string {
	return r.raw
}

// AuditLogEntryKind identifies the coarse query and rendering family of an
// audit-log entry.
type EventGetResponseEntryKind string

const (
	EventGetResponseEntryKindAuditLogEntryKindUnspecified              EventGetResponseEntryKind = "AUDIT_LOG_ENTRY_KIND_UNSPECIFIED"
	EventGetResponseEntryKindAuditLogEntryKindAgentSecurityExecBlocked EventGetResponseEntryKind = "AUDIT_LOG_ENTRY_KIND_AGENT_SECURITY_EXEC_BLOCKED"
	EventGetResponseEntryKindAuditLogEntryKindAgentSecurityExecAudited EventGetResponseEntryKind = "AUDIT_LOG_ENTRY_KIND_AGENT_SECURITY_EXEC_AUDITED"
	EventGetResponseEntryKindAuditLogEntryKindResourceChange           EventGetResponseEntryKind = "AUDIT_LOG_ENTRY_KIND_RESOURCE_CHANGE"
	EventGetResponseEntryKindAuditLogEntryKindCredentialAccess         EventGetResponseEntryKind = "AUDIT_LOG_ENTRY_KIND_CREDENTIAL_ACCESS"
	EventGetResponseEntryKindAuditLogEntryKindEnvironmentVeto          EventGetResponseEntryKind = "AUDIT_LOG_ENTRY_KIND_ENVIRONMENT_VETO"
)

func (r EventGetResponseEntryKind) IsKnown() bool {
	switch r {
	case EventGetResponseEntryKindAuditLogEntryKindUnspecified, EventGetResponseEntryKindAuditLogEntryKindAgentSecurityExecBlocked, EventGetResponseEntryKindAuditLogEntryKindAgentSecurityExecAudited, EventGetResponseEntryKindAuditLogEntryKindResourceChange, EventGetResponseEntryKindAuditLogEntryKindCredentialAccess, EventGetResponseEntryKindAuditLogEntryKindEnvironmentVeto:
		return true
	}
	return false
}

type EventListResponse struct {
	ID             string           `json:"id"`
	Action         string           `json:"action"`
	ActorID        string           `json:"actorId"`
	ActorPrincipal shared.Principal `json:"actorPrincipal"`
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
	// AuditLogEntryKind identifies the coarse query and rendering family of an
	// audit-log entry.
	Kind        EventListResponseKind `json:"kind"`
	SubjectID   string                `json:"subjectId"`
	SubjectType shared.ResourceType   `json:"subjectType"`
	JSON        eventListResponseJSON `json:"-"`
}

// eventListResponseJSON contains the JSON metadata for the struct
// [EventListResponse]
type eventListResponseJSON struct {
	ID             apijson.Field
	Action         apijson.Field
	ActorID        apijson.Field
	ActorPrincipal apijson.Field
	CreatedAt      apijson.Field
	Kind           apijson.Field
	SubjectID      apijson.Field
	SubjectType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseJSON) RawJSON() string {
	return r.raw
}

// AuditLogEntryKind identifies the coarse query and rendering family of an
// audit-log entry.
type EventListResponseKind string

const (
	EventListResponseKindAuditLogEntryKindUnspecified              EventListResponseKind = "AUDIT_LOG_ENTRY_KIND_UNSPECIFIED"
	EventListResponseKindAuditLogEntryKindAgentSecurityExecBlocked EventListResponseKind = "AUDIT_LOG_ENTRY_KIND_AGENT_SECURITY_EXEC_BLOCKED"
	EventListResponseKindAuditLogEntryKindAgentSecurityExecAudited EventListResponseKind = "AUDIT_LOG_ENTRY_KIND_AGENT_SECURITY_EXEC_AUDITED"
	EventListResponseKindAuditLogEntryKindResourceChange           EventListResponseKind = "AUDIT_LOG_ENTRY_KIND_RESOURCE_CHANGE"
	EventListResponseKindAuditLogEntryKindCredentialAccess         EventListResponseKind = "AUDIT_LOG_ENTRY_KIND_CREDENTIAL_ACCESS"
	EventListResponseKindAuditLogEntryKindEnvironmentVeto          EventListResponseKind = "AUDIT_LOG_ENTRY_KIND_ENVIRONMENT_VETO"
)

func (r EventListResponseKind) IsKnown() bool {
	switch r {
	case EventListResponseKindAuditLogEntryKindUnspecified, EventListResponseKindAuditLogEntryKindAgentSecurityExecBlocked, EventListResponseKindAuditLogEntryKindAgentSecurityExecAudited, EventListResponseKindAuditLogEntryKindResourceChange, EventListResponseKindAuditLogEntryKindCredentialAccess, EventListResponseKindAuditLogEntryKindEnvironmentVeto:
		return true
	}
	return false
}

type EventWatchResponse struct {
	Operation    ResourceOperation      `json:"operation"`
	ResourceID   string                 `json:"resourceId" format:"uuid"`
	ResourceType shared.ResourceType    `json:"resourceType"`
	JSON         eventWatchResponseJSON `json:"-"`
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

type EventGetParams struct {
	// audit_log_entry_id is the ID of the audit-log entry to retrieve.
	AuditLogEntryID param.Field[string] `json:"auditLogEntryId" api:"required" format:"uuid"`
}

func (r EventGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventListParams struct {
	Token    param.Field[string]                `query:"token"`
	PageSize param.Field[int64]                 `query:"pageSize"`
	Filter   param.Field[EventListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing audit logs
	Pagination param.Field[EventListParamsPagination] `json:"pagination"`
	// sort specifies the order of results. When unspecified, results are sorted by
	// creation time descending (newest first). Supported sort fields: createdAt.
	Sort param.Field[shared.SortParam] `json:"sort"`
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
	ActorIDs        param.Field[[]string]           `json:"actorIds" format:"uuid"`
	ActorPrincipals param.Field[[]shared.Principal] `json:"actorPrincipals"`
	// from filters audit logs created at or after this timestamp (inclusive).
	From         param.Field[time.Time]             `json:"from" format:"date-time"`
	SubjectIDs   param.Field[[]string]              `json:"subjectIds" format:"uuid"`
	SubjectTypes param.Field[[]shared.ResourceType] `json:"subjectTypes"`
	// to filters audit logs created before this timestamp (exclusive).
	To param.Field[time.Time] `json:"to" format:"date-time"`
}

func (r EventListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing audit logs
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
	// Environment scope produces events for the environment itself, all tasks, task
	// executions, and services associated with that environment.
	EnvironmentID param.Field[string] `json:"environmentId"`
	// Organization scope produces events for all projects, runners and environments
	// the caller can see within their organization. No task, task execution or service
	// events are produed.
	Organization param.Field[bool] `json:"organization"`
	// Filters to limit which events are delivered on organization-scoped streams. When
	// empty, all events for the scope are delivered. When populated, only events
	// matching at least one filter entry are forwarded. Not supported for
	// environment-scoped streams; setting this field returns an error.
	ResourceTypeFilters param.Field[[]EventWatchParamsResourceTypeFilter] `json:"resourceTypeFilters"`
}

func (r EventWatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ResourceTypeFilter restricts which events are delivered for a specific resource
// type.
type EventWatchParamsResourceTypeFilter struct {
	// If non-empty, only events where the resource was created by one of these user
	// IDs are delivered. Skipped for DELETE operations (creator info is unavailable
	// after deletion). Events with no creator information are skipped when this filter
	// is set (fail-closed).
	CreatorIDs param.Field[[]string] `json:"creatorIds" format:"uuid"`
	// If non-empty, only events for these specific resource IDs are delivered.
	ResourceIDs param.Field[[]string] `json:"resourceIds" format:"uuid"`
	// The resource type to filter for.
	ResourceType param.Field[shared.ResourceType] `json:"resourceType"`
}

func (r EventWatchParamsResourceTypeFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
