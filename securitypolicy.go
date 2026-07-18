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
)

// SecurityPolicyService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityPolicyService] method instead.
type SecurityPolicyService struct {
	Options []option.RequestOption
}

// NewSecurityPolicyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecurityPolicyService(opts ...option.RequestOption) (r *SecurityPolicyService) {
	r = &SecurityPolicyService{}
	r.Options = opts
	return
}

// Creates a new security policy.
//
// Use this method to:
//
// - Define environment access controls
// - Configure audited or blocked operations
// - Manage organization security posture
//
// ### Examples
//
// - Create security policy:
//
//	Creates an audit-first Veto Exec policy with one audited bare name and one
//	blocked absolute path. Creation stores an inactive definition; assigning it as
//	the organization default validates materializability.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	metadata:
//	  name: "Veto Exec audit-first"
//	spec:
//	  executables:
//	    defaultEffect: EFFECT_ALLOW
//	    rules:
//	      - path: "npx"
//	        effect: EFFECT_AUDIT
//	      - path: "/usr/bin/curl"
//	        effect: EFFECT_BLOCK
//	```
func (r *SecurityPolicyService) New(ctx context.Context, body SecurityPolicyNewParams, opts ...option.RequestOption) (res *SecurityPolicyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.SecurityService/CreateSecurityPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets details about a specific security policy.
//
// Use this method to:
//
// - View security policy configuration
// - Inspect enforcement rules
//
// ### Examples
//
// - Get security policy:
//
//	Retrieves a security policy by ID.
//
//	```yaml
//	securityPolicyId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *SecurityPolicyService) Get(ctx context.Context, body SecurityPolicyGetParams, opts ...option.RequestOption) (res *SecurityPolicyGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.SecurityService/GetSecurityPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates a security policy.
//
// Use this method to:
//
// - Rename a security policy
// - Change enforcement rules
// - Update auditing behavior
//
// ### Examples
//
// - Update security policy:
//
//	Promotes one executable rule from audit to block while leaving unmatched
//	executables allowed. Updating an assigned policy validates materializability;
//	updating an unassigned policy only stores its spec.
//
//	```yaml
//	securityPolicyId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	spec:
//	  executables:
//	    defaultEffect: EFFECT_ALLOW
//	    rules:
//	      - path: "npx"
//	        effect: EFFECT_BLOCK
//	      - path: "/usr/bin/curl"
//	        effect: EFFECT_BLOCK
//	```
func (r *SecurityPolicyService) Update(ctx context.Context, body SecurityPolicyUpdateParams, opts ...option.RequestOption) (res *SecurityPolicyUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.SecurityService/UpdateSecurityPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists security policies.
//
// Use this method to:
//
// - View all security policies in an organization
// - Audit configured security controls
//
// ### Examples
//
// - List organization policies:
//
//	Shows security policies with pagination.
//
//	```yaml
//	filter:
//	  organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	```
func (r *SecurityPolicyService) List(ctx context.Context, params SecurityPolicyListParams, opts ...option.RequestOption) (res *pagination.SecurityPoliciesPage[SecurityPolicy], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.SecurityService/ListSecurityPolicies"
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

// Lists security policies.
//
// Use this method to:
//
// - View all security policies in an organization
// - Audit configured security controls
//
// ### Examples
//
// - List organization policies:
//
//	Shows security policies with pagination.
//
//	```yaml
//	filter:
//	  organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	```
func (r *SecurityPolicyService) ListAutoPaging(ctx context.Context, params SecurityPolicyListParams, opts ...option.RequestOption) *pagination.SecurityPoliciesPageAutoPager[SecurityPolicy] {
	return pagination.NewSecurityPoliciesPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a security policy.
//
// Use this method to:
//
// - Remove obsolete security policies
// - Clean up unused policy definitions
//
// ### Examples
//
// - Delete security policy:
//
//	Permanently removes a security policy.
//
//	```yaml
//	securityPolicyId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *SecurityPolicyService) Delete(ctx context.Context, body SecurityPolicyDeleteParams, opts ...option.RequestOption) (res *SecurityPolicyDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.SecurityService/DeleteSecurityPolicy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SecurityPolicy struct {
	Metadata SecurityPolicyMetadata `json:"metadata" api:"required"`
	// Mandate/deploy security agents, e.g. CrowdStrike. Mandate credential
	// security/proxy use. These can be modeled later as explicit fields if needed.
	Spec SecurityPolicySpec `json:"spec" api:"required"`
	ID   string             `json:"id" format:"uuid"`
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
	CreatedAt      time.Time `json:"createdAt" format:"date-time"`
	OrganizationID string    `json:"organizationId" format:"uuid"`
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
	UpdatedAt time.Time          `json:"updatedAt" format:"date-time"`
	JSON      securityPolicyJSON `json:"-"`
}

// securityPolicyJSON contains the JSON metadata for the struct [SecurityPolicy]
type securityPolicyJSON struct {
	Metadata       apijson.Field
	Spec           apijson.Field
	ID             apijson.Field
	CreatedAt      apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SecurityPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityPolicyJSON) RawJSON() string {
	return r.raw
}

type SecurityPolicyMetadata struct {
	Name string                     `json:"name"`
	JSON securityPolicyMetadataJSON `json:"-"`
}

// securityPolicyMetadataJSON contains the JSON metadata for the struct
// [SecurityPolicyMetadata]
type securityPolicyMetadataJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecurityPolicyMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityPolicyMetadataJSON) RawJSON() string {
	return r.raw
}

// Mandate/deploy security agents, e.g. CrowdStrike. Mandate credential
// security/proxy use. These can be modeled later as explicit fields if needed.
type SecurityPolicySpec struct {
	// executables is the public Veto Exec GA policy surface.
	Executables SecurityPolicySpecExecutables `json:"executables"`
	JSON        securityPolicySpecJSON        `json:"-"`
}

// securityPolicySpecJSON contains the JSON metadata for the struct
// [SecurityPolicySpec]
type securityPolicySpecJSON struct {
	Executables apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecurityPolicySpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityPolicySpecJSON) RawJSON() string {
	return r.raw
}

// executables is the public Veto Exec GA policy surface.
type SecurityPolicySpecExecutables struct {
	// default_effect controls executables that do not match a rule. For Veto Exec,
	// omit this field or set it to EFFECT_ALLOW. EFFECT_UNSPECIFIED is normalized to
	// EFFECT_ALLOW.
	DefaultEffect SecurityPolicySpecExecutablesDefaultEffect `json:"defaultEffect"`
	// rules contains executable-specific audit or block decisions.
	Rules []SecurityPolicySpecExecutablesRule `json:"rules"`
	JSON  securityPolicySpecExecutablesJSON   `json:"-"`
}

// securityPolicySpecExecutablesJSON contains the JSON metadata for the struct
// [SecurityPolicySpecExecutables]
type securityPolicySpecExecutablesJSON struct {
	DefaultEffect apijson.Field
	Rules         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SecurityPolicySpecExecutables) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityPolicySpecExecutablesJSON) RawJSON() string {
	return r.raw
}

// default_effect controls executables that do not match a rule. For Veto Exec,
// omit this field or set it to EFFECT_ALLOW. EFFECT_UNSPECIFIED is normalized to
// EFFECT_ALLOW.
type SecurityPolicySpecExecutablesDefaultEffect string

const (
	SecurityPolicySpecExecutablesDefaultEffectEffectUnspecified SecurityPolicySpecExecutablesDefaultEffect = "EFFECT_UNSPECIFIED"
	SecurityPolicySpecExecutablesDefaultEffectEffectAllow       SecurityPolicySpecExecutablesDefaultEffect = "EFFECT_ALLOW"
	SecurityPolicySpecExecutablesDefaultEffectEffectBlock       SecurityPolicySpecExecutablesDefaultEffect = "EFFECT_BLOCK"
	SecurityPolicySpecExecutablesDefaultEffectEffectAudit       SecurityPolicySpecExecutablesDefaultEffect = "EFFECT_AUDIT"
)

func (r SecurityPolicySpecExecutablesDefaultEffect) IsKnown() bool {
	switch r {
	case SecurityPolicySpecExecutablesDefaultEffectEffectUnspecified, SecurityPolicySpecExecutablesDefaultEffectEffectAllow, SecurityPolicySpecExecutablesDefaultEffectEffectBlock, SecurityPolicySpecExecutablesDefaultEffectEffectAudit:
		return true
	}
	return false
}

type SecurityPolicySpecExecutablesRule struct {
	// effect must be EFFECT_AUDIT or EFFECT_BLOCK. EFFECT_ALLOW is not supported on an
	// executable rule.
	Effect SecurityPolicySpecExecutablesRulesEffect `json:"effect"`
	// path is either an absolute executable path, such as /usr/bin/curl, or a bare
	// executable name, such as npx. Bare names are expanded by runtime discovery.
	// Surrounding whitespace is ignored. Empty or whitespace-only selectors and
	// relative paths with directory separators are invalid. Enforcement uses
	// executable content hashes, so different paths with identical content share one
	// runtime decision and block wins conflicts.
	Path string                                `json:"path"`
	JSON securityPolicySpecExecutablesRuleJSON `json:"-"`
}

// securityPolicySpecExecutablesRuleJSON contains the JSON metadata for the struct
// [SecurityPolicySpecExecutablesRule]
type securityPolicySpecExecutablesRuleJSON struct {
	Effect      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecurityPolicySpecExecutablesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityPolicySpecExecutablesRuleJSON) RawJSON() string {
	return r.raw
}

// effect must be EFFECT_AUDIT or EFFECT_BLOCK. EFFECT_ALLOW is not supported on an
// executable rule.
type SecurityPolicySpecExecutablesRulesEffect string

const (
	SecurityPolicySpecExecutablesRulesEffectEffectUnspecified SecurityPolicySpecExecutablesRulesEffect = "EFFECT_UNSPECIFIED"
	SecurityPolicySpecExecutablesRulesEffectEffectAllow       SecurityPolicySpecExecutablesRulesEffect = "EFFECT_ALLOW"
	SecurityPolicySpecExecutablesRulesEffectEffectBlock       SecurityPolicySpecExecutablesRulesEffect = "EFFECT_BLOCK"
	SecurityPolicySpecExecutablesRulesEffectEffectAudit       SecurityPolicySpecExecutablesRulesEffect = "EFFECT_AUDIT"
)

func (r SecurityPolicySpecExecutablesRulesEffect) IsKnown() bool {
	switch r {
	case SecurityPolicySpecExecutablesRulesEffectEffectUnspecified, SecurityPolicySpecExecutablesRulesEffectEffectAllow, SecurityPolicySpecExecutablesRulesEffectEffectBlock, SecurityPolicySpecExecutablesRulesEffectEffectAudit:
		return true
	}
	return false
}

type SecurityPolicyNewResponse struct {
	SecurityPolicy SecurityPolicy                `json:"securityPolicy" api:"required"`
	JSON           securityPolicyNewResponseJSON `json:"-"`
}

// securityPolicyNewResponseJSON contains the JSON metadata for the struct
// [SecurityPolicyNewResponse]
type securityPolicyNewResponseJSON struct {
	SecurityPolicy apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SecurityPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

type SecurityPolicyGetResponse struct {
	SecurityPolicy SecurityPolicy                `json:"securityPolicy" api:"required"`
	JSON           securityPolicyGetResponseJSON `json:"-"`
}

// securityPolicyGetResponseJSON contains the JSON metadata for the struct
// [SecurityPolicyGetResponse]
type securityPolicyGetResponseJSON struct {
	SecurityPolicy apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SecurityPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

type SecurityPolicyUpdateResponse struct {
	SecurityPolicy SecurityPolicy                   `json:"securityPolicy" api:"required"`
	JSON           securityPolicyUpdateResponseJSON `json:"-"`
}

// securityPolicyUpdateResponseJSON contains the JSON metadata for the struct
// [SecurityPolicyUpdateResponse]
type securityPolicyUpdateResponseJSON struct {
	SecurityPolicy apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SecurityPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityPolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SecurityPolicyDeleteResponse = interface{}

type SecurityPolicyNewParams struct {
	Metadata param.Field[SecurityPolicyNewParamsMetadata] `json:"metadata" api:"required"`
	// Mandate/deploy security agents, e.g. CrowdStrike. Mandate credential
	// security/proxy use. These can be modeled later as explicit fields if needed.
	Spec           param.Field[SecurityPolicyNewParamsSpec] `json:"spec" api:"required"`
	OrganizationID param.Field[string]                      `json:"organizationId" format:"uuid"`
}

func (r SecurityPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecurityPolicyNewParamsMetadata struct {
	Name param.Field[string] `json:"name"`
}

func (r SecurityPolicyNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mandate/deploy security agents, e.g. CrowdStrike. Mandate credential
// security/proxy use. These can be modeled later as explicit fields if needed.
type SecurityPolicyNewParamsSpec struct {
	// executables is the public Veto Exec GA policy surface.
	Executables param.Field[SecurityPolicyNewParamsSpecExecutables] `json:"executables"`
}

func (r SecurityPolicyNewParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// executables is the public Veto Exec GA policy surface.
type SecurityPolicyNewParamsSpecExecutables struct {
	// default_effect controls executables that do not match a rule. For Veto Exec,
	// omit this field or set it to EFFECT_ALLOW. EFFECT_UNSPECIFIED is normalized to
	// EFFECT_ALLOW.
	DefaultEffect param.Field[SecurityPolicyNewParamsSpecExecutablesDefaultEffect] `json:"defaultEffect"`
	// rules contains executable-specific audit or block decisions.
	Rules param.Field[[]SecurityPolicyNewParamsSpecExecutablesRule] `json:"rules"`
}

func (r SecurityPolicyNewParamsSpecExecutables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// default_effect controls executables that do not match a rule. For Veto Exec,
// omit this field or set it to EFFECT_ALLOW. EFFECT_UNSPECIFIED is normalized to
// EFFECT_ALLOW.
type SecurityPolicyNewParamsSpecExecutablesDefaultEffect string

const (
	SecurityPolicyNewParamsSpecExecutablesDefaultEffectEffectUnspecified SecurityPolicyNewParamsSpecExecutablesDefaultEffect = "EFFECT_UNSPECIFIED"
	SecurityPolicyNewParamsSpecExecutablesDefaultEffectEffectAllow       SecurityPolicyNewParamsSpecExecutablesDefaultEffect = "EFFECT_ALLOW"
	SecurityPolicyNewParamsSpecExecutablesDefaultEffectEffectBlock       SecurityPolicyNewParamsSpecExecutablesDefaultEffect = "EFFECT_BLOCK"
	SecurityPolicyNewParamsSpecExecutablesDefaultEffectEffectAudit       SecurityPolicyNewParamsSpecExecutablesDefaultEffect = "EFFECT_AUDIT"
)

func (r SecurityPolicyNewParamsSpecExecutablesDefaultEffect) IsKnown() bool {
	switch r {
	case SecurityPolicyNewParamsSpecExecutablesDefaultEffectEffectUnspecified, SecurityPolicyNewParamsSpecExecutablesDefaultEffectEffectAllow, SecurityPolicyNewParamsSpecExecutablesDefaultEffectEffectBlock, SecurityPolicyNewParamsSpecExecutablesDefaultEffectEffectAudit:
		return true
	}
	return false
}

type SecurityPolicyNewParamsSpecExecutablesRule struct {
	// effect must be EFFECT_AUDIT or EFFECT_BLOCK. EFFECT_ALLOW is not supported on an
	// executable rule.
	Effect param.Field[SecurityPolicyNewParamsSpecExecutablesRulesEffect] `json:"effect"`
	// path is either an absolute executable path, such as /usr/bin/curl, or a bare
	// executable name, such as npx. Bare names are expanded by runtime discovery.
	// Surrounding whitespace is ignored. Empty or whitespace-only selectors and
	// relative paths with directory separators are invalid. Enforcement uses
	// executable content hashes, so different paths with identical content share one
	// runtime decision and block wins conflicts.
	Path param.Field[string] `json:"path"`
}

func (r SecurityPolicyNewParamsSpecExecutablesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// effect must be EFFECT_AUDIT or EFFECT_BLOCK. EFFECT_ALLOW is not supported on an
// executable rule.
type SecurityPolicyNewParamsSpecExecutablesRulesEffect string

const (
	SecurityPolicyNewParamsSpecExecutablesRulesEffectEffectUnspecified SecurityPolicyNewParamsSpecExecutablesRulesEffect = "EFFECT_UNSPECIFIED"
	SecurityPolicyNewParamsSpecExecutablesRulesEffectEffectAllow       SecurityPolicyNewParamsSpecExecutablesRulesEffect = "EFFECT_ALLOW"
	SecurityPolicyNewParamsSpecExecutablesRulesEffectEffectBlock       SecurityPolicyNewParamsSpecExecutablesRulesEffect = "EFFECT_BLOCK"
	SecurityPolicyNewParamsSpecExecutablesRulesEffectEffectAudit       SecurityPolicyNewParamsSpecExecutablesRulesEffect = "EFFECT_AUDIT"
)

func (r SecurityPolicyNewParamsSpecExecutablesRulesEffect) IsKnown() bool {
	switch r {
	case SecurityPolicyNewParamsSpecExecutablesRulesEffectEffectUnspecified, SecurityPolicyNewParamsSpecExecutablesRulesEffectEffectAllow, SecurityPolicyNewParamsSpecExecutablesRulesEffectEffectBlock, SecurityPolicyNewParamsSpecExecutablesRulesEffectEffectAudit:
		return true
	}
	return false
}

type SecurityPolicyGetParams struct {
	SecurityPolicyID param.Field[string] `json:"securityPolicyId" format:"uuid"`
}

func (r SecurityPolicyGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecurityPolicyUpdateParams struct {
	Metadata         param.Field[SecurityPolicyUpdateParamsMetadata] `json:"metadata"`
	SecurityPolicyID param.Field[string]                             `json:"securityPolicyId" format:"uuid"`
	// Mandate/deploy security agents, e.g. CrowdStrike. Mandate credential
	// security/proxy use. These can be modeled later as explicit fields if needed.
	Spec param.Field[SecurityPolicyUpdateParamsSpec] `json:"spec"`
}

func (r SecurityPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecurityPolicyUpdateParamsMetadata struct {
	Name param.Field[string] `json:"name"`
}

func (r SecurityPolicyUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mandate/deploy security agents, e.g. CrowdStrike. Mandate credential
// security/proxy use. These can be modeled later as explicit fields if needed.
type SecurityPolicyUpdateParamsSpec struct {
	// executables is the public Veto Exec GA policy surface.
	Executables param.Field[SecurityPolicyUpdateParamsSpecExecutables] `json:"executables"`
}

func (r SecurityPolicyUpdateParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// executables is the public Veto Exec GA policy surface.
type SecurityPolicyUpdateParamsSpecExecutables struct {
	// default_effect controls executables that do not match a rule. For Veto Exec,
	// omit this field or set it to EFFECT_ALLOW. EFFECT_UNSPECIFIED is normalized to
	// EFFECT_ALLOW.
	DefaultEffect param.Field[SecurityPolicyUpdateParamsSpecExecutablesDefaultEffect] `json:"defaultEffect"`
	// rules contains executable-specific audit or block decisions.
	Rules param.Field[[]SecurityPolicyUpdateParamsSpecExecutablesRule] `json:"rules"`
}

func (r SecurityPolicyUpdateParamsSpecExecutables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// default_effect controls executables that do not match a rule. For Veto Exec,
// omit this field or set it to EFFECT_ALLOW. EFFECT_UNSPECIFIED is normalized to
// EFFECT_ALLOW.
type SecurityPolicyUpdateParamsSpecExecutablesDefaultEffect string

const (
	SecurityPolicyUpdateParamsSpecExecutablesDefaultEffectEffectUnspecified SecurityPolicyUpdateParamsSpecExecutablesDefaultEffect = "EFFECT_UNSPECIFIED"
	SecurityPolicyUpdateParamsSpecExecutablesDefaultEffectEffectAllow       SecurityPolicyUpdateParamsSpecExecutablesDefaultEffect = "EFFECT_ALLOW"
	SecurityPolicyUpdateParamsSpecExecutablesDefaultEffectEffectBlock       SecurityPolicyUpdateParamsSpecExecutablesDefaultEffect = "EFFECT_BLOCK"
	SecurityPolicyUpdateParamsSpecExecutablesDefaultEffectEffectAudit       SecurityPolicyUpdateParamsSpecExecutablesDefaultEffect = "EFFECT_AUDIT"
)

func (r SecurityPolicyUpdateParamsSpecExecutablesDefaultEffect) IsKnown() bool {
	switch r {
	case SecurityPolicyUpdateParamsSpecExecutablesDefaultEffectEffectUnspecified, SecurityPolicyUpdateParamsSpecExecutablesDefaultEffectEffectAllow, SecurityPolicyUpdateParamsSpecExecutablesDefaultEffectEffectBlock, SecurityPolicyUpdateParamsSpecExecutablesDefaultEffectEffectAudit:
		return true
	}
	return false
}

type SecurityPolicyUpdateParamsSpecExecutablesRule struct {
	// effect must be EFFECT_AUDIT or EFFECT_BLOCK. EFFECT_ALLOW is not supported on an
	// executable rule.
	Effect param.Field[SecurityPolicyUpdateParamsSpecExecutablesRulesEffect] `json:"effect"`
	// path is either an absolute executable path, such as /usr/bin/curl, or a bare
	// executable name, such as npx. Bare names are expanded by runtime discovery.
	// Surrounding whitespace is ignored. Empty or whitespace-only selectors and
	// relative paths with directory separators are invalid. Enforcement uses
	// executable content hashes, so different paths with identical content share one
	// runtime decision and block wins conflicts.
	Path param.Field[string] `json:"path"`
}

func (r SecurityPolicyUpdateParamsSpecExecutablesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// effect must be EFFECT_AUDIT or EFFECT_BLOCK. EFFECT_ALLOW is not supported on an
// executable rule.
type SecurityPolicyUpdateParamsSpecExecutablesRulesEffect string

const (
	SecurityPolicyUpdateParamsSpecExecutablesRulesEffectEffectUnspecified SecurityPolicyUpdateParamsSpecExecutablesRulesEffect = "EFFECT_UNSPECIFIED"
	SecurityPolicyUpdateParamsSpecExecutablesRulesEffectEffectAllow       SecurityPolicyUpdateParamsSpecExecutablesRulesEffect = "EFFECT_ALLOW"
	SecurityPolicyUpdateParamsSpecExecutablesRulesEffectEffectBlock       SecurityPolicyUpdateParamsSpecExecutablesRulesEffect = "EFFECT_BLOCK"
	SecurityPolicyUpdateParamsSpecExecutablesRulesEffectEffectAudit       SecurityPolicyUpdateParamsSpecExecutablesRulesEffect = "EFFECT_AUDIT"
)

func (r SecurityPolicyUpdateParamsSpecExecutablesRulesEffect) IsKnown() bool {
	switch r {
	case SecurityPolicyUpdateParamsSpecExecutablesRulesEffectEffectUnspecified, SecurityPolicyUpdateParamsSpecExecutablesRulesEffectEffectAllow, SecurityPolicyUpdateParamsSpecExecutablesRulesEffectEffectBlock, SecurityPolicyUpdateParamsSpecExecutablesRulesEffectEffectAudit:
		return true
	}
	return false
}

type SecurityPolicyListParams struct {
	Token      param.Field[string]                             `query:"token"`
	PageSize   param.Field[int64]                              `query:"pageSize"`
	Filter     param.Field[SecurityPolicyListParamsFilter]     `json:"filter"`
	Pagination param.Field[SecurityPolicyListParamsPagination] `json:"pagination"`
}

func (r SecurityPolicyListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SecurityPolicyListParams]'s query parameters as
// `url.Values`.
func (r SecurityPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SecurityPolicyListParamsFilter struct {
	OrganizationID    param.Field[string]   `json:"organizationId" format:"uuid"`
	Search            param.Field[string]   `json:"search"`
	SecurityPolicyIDs param.Field[[]string] `json:"securityPolicyIds" format:"uuid"`
}

func (r SecurityPolicyListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecurityPolicyListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r SecurityPolicyListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecurityPolicyDeleteParams struct {
	SecurityPolicyID param.Field[string] `json:"securityPolicyId" format:"uuid"`
}

func (r SecurityPolicyDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
