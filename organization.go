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

// OrganizationService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options             []option.RequestOption
	CustomDomains       *OrganizationCustomDomainService
	DomainVerifications *OrganizationDomainVerificationService
	Invites             *OrganizationInviteService
	Policies            *OrganizationPolicyService
	SSOConfigurations   *OrganizationSSOConfigurationService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.CustomDomains = NewOrganizationCustomDomainService(opts...)
	r.DomainVerifications = NewOrganizationDomainVerificationService(opts...)
	r.Invites = NewOrganizationInviteService(opts...)
	r.Policies = NewOrganizationPolicyService(opts...)
	r.SSOConfigurations = NewOrganizationSSOConfigurationService(opts...)
	return
}

// Creates a new organization with the specified name and settings.
//
// Use this method to:
//
// - Create a new organization for team collaboration
// - Set up automatic domain-based invites for team members
// - Join the organization immediately upon creation
//
// ### Examples
//
// - Create a basic organization:
//
//	Creates an organization with just a name.
//
//	```yaml
//	name: "Acme Corp Engineering"
//	joinOrganization: true
//	```
//
// - Create with domain-based invites:
//
//	Creates an organization that automatically invites users with matching email
//	domains.
//
//	```yaml
//	name: "Acme Corp"
//	joinOrganization: true
//	inviteAccountsWithMatchingDomain: true
//	```
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (res *OrganizationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/CreateOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets details about a specific organization.
//
// Use this method to:
//
// - Retrieve organization settings and configuration
// - Check organization membership status
// - View domain verification settings
//
// ### Examples
//
// - Get organization details:
//
//	Retrieves information about a specific organization.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *OrganizationService) Get(ctx context.Context, body OrganizationGetParams, opts ...option.RequestOption) (res *OrganizationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/GetOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an organization's settings including name, invite domains, and member
// policies.
//
// Use this method to:
//
// - Modify organization display name
// - Configure email domain restrictions
// - Update organization-wide settings
// - Manage member access policies
//
// ### Examples
//
// - Update basic settings:
//
//	Changes organization name and invite domains.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	name: "New Company Name"
//	inviteDomains:
//	  domains:
//	    - "company.com"
//	    - "subsidiary.com"
//	```
//
// - Remove domain restrictions:
//
//	Clears all domain-based invite restrictions.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	inviteDomains:
//	  domains: []
//	```
func (r *OrganizationService) Update(ctx context.Context, body OrganizationUpdateParams, opts ...option.RequestOption) (res *OrganizationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/UpdateOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Permanently deletes an organization.
//
// Use this method to:
//
// - Remove unused organizations
// - Clean up test organizations
// - Complete organization migration
//
// ### Examples
//
// - Delete organization:
//
//	Permanently removes an organization and all its data.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
func (r *OrganizationService) Delete(ctx context.Context, body OrganizationDeleteParams, opts ...option.RequestOption) (res *OrganizationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/DeleteOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Allows users to join an organization through direct ID, invite link, or
// domain-based auto-join.
//
// Use this method to:
//
// - Join an organization via direct ID or invite
// - Join automatically based on email domain
// - Accept organization invitations
//
// ### Examples
//
// - Join via organization ID:
//
//	Joins an organization directly when you have the ID.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	```
//
// - Join via invite:
//
//	Accepts an organization invitation link.
//
//	```yaml
//	inviteId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *OrganizationService) Join(ctx context.Context, body OrganizationJoinParams, opts ...option.RequestOption) (res *OrganizationJoinResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/JoinOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes a user from an organization while preserving organization data.
//
// Use this method to:
//
// - Remove yourself from an organization
// - Clean up inactive memberships
// - Transfer project ownership before leaving
// - Manage team transitions
//
// ### Examples
//
// - Leave organization:
//
//	Removes user from organization membership.
//
//	```yaml
//	userId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	```
//
// Note: Ensure all projects and resources are transferred before leaving.
func (r *OrganizationService) Leave(ctx context.Context, body OrganizationLeaveParams, opts ...option.RequestOption) (res *OrganizationLeaveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/LeaveOrganization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists and filters organization members with optional pagination.
//
// Use this method to:
//
// - View all organization members
// - Monitor member activity
// - Manage team membership
//
// ### Examples
//
// - List active members:
//
//	Retrieves active members with pagination.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	```
//
// - List with pagination:
//
//	Retrieves next page of members.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 50
//	  token: "next-page-token-from-previous-response"
//	```
func (r *OrganizationService) ListMembers(ctx context.Context, params OrganizationListMembersParams, opts ...option.RequestOption) (res *pagination.MembersPage[OrganizationMember], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.OrganizationService/ListMembers"
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

// Lists and filters organization members with optional pagination.
//
// Use this method to:
//
// - View all organization members
// - Monitor member activity
// - Manage team membership
//
// ### Examples
//
// - List active members:
//
//	Retrieves active members with pagination.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	```
//
// - List with pagination:
//
//	Retrieves next page of members.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 50
//	  token: "next-page-token-from-previous-response"
//	```
func (r *OrganizationService) ListMembersAutoPaging(ctx context.Context, params OrganizationListMembersParams, opts ...option.RequestOption) *pagination.MembersPageAutoPager[OrganizationMember] {
	return pagination.NewMembersPageAutoPager(r.ListMembers(ctx, params, opts...))
}

// Manages organization membership and roles by setting a user's role within the
// organization.
//
// Use this method to:
//
// - Promote members to admin role
// - Change member permissions
// - Demote admins to regular members
//
// ### Examples
//
// - Promote to admin:
//
//	Makes a user an organization administrator.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	userId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	role: ORGANIZATION_ROLE_ADMIN
//	```
//
// - Change to member:
//
//	Changes a user's role to regular member.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	userId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	role: ORGANIZATION_ROLE_MEMBER
//	```
func (r *OrganizationService) SetRole(ctx context.Context, body OrganizationSetRoleParams, opts ...option.RequestOption) (res *OrganizationSetRoleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.OrganizationService/SetRole"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type InviteDomains struct {
	// domains is the list of domains that are allowed to join the organization
	Domains []string          `json:"domains"`
	JSON    inviteDomainsJSON `json:"-"`
}

// inviteDomainsJSON contains the JSON metadata for the struct [InviteDomains]
type inviteDomainsJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteDomains) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteDomainsJSON) RawJSON() string {
	return r.raw
}

type InviteDomainsParam struct {
	// domains is the list of domains that are allowed to join the organization
	Domains param.Field[[]string] `json:"domains"`
}

func (r InviteDomainsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Organization struct {
	ID string `json:"id,required" format:"uuid"`
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
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	Name      string    `json:"name,required"`
	// The tier of the organization - free, enterprise or core
	Tier OrganizationTier `json:"tier,required"`
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
	UpdatedAt     time.Time        `json:"updatedAt,required" format:"date-time"`
	InviteDomains InviteDomains    `json:"inviteDomains"`
	JSON          organizationJSON `json:"-"`
}

// organizationJSON contains the JSON metadata for the struct [Organization]
type organizationJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Name          apijson.Field
	Tier          apijson.Field
	UpdatedAt     apijson.Field
	InviteDomains apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Organization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationJSON) RawJSON() string {
	return r.raw
}

type OrganizationMember struct {
	Email    string `json:"email,required"`
	FullName string `json:"fullName,required"`
	// login_provider is the login provider the user uses to sign in
	LoginProvider string `json:"loginProvider,required"`
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
	MemberSince time.Time               `json:"memberSince,required" format:"date-time"`
	Role        shared.OrganizationRole `json:"role,required"`
	Status      shared.UserStatus       `json:"status,required"`
	UserID      string                  `json:"userId,required" format:"uuid"`
	AvatarURL   string                  `json:"avatarUrl"`
	JSON        organizationMemberJSON  `json:"-"`
}

// organizationMemberJSON contains the JSON metadata for the struct
// [OrganizationMember]
type organizationMemberJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LoginProvider apijson.Field
	MemberSince   apijson.Field
	Role          apijson.Field
	Status        apijson.Field
	UserID        apijson.Field
	AvatarURL     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMemberJSON) RawJSON() string {
	return r.raw
}

type OrganizationTier string

const (
	OrganizationTierUnspecified OrganizationTier = "ORGANIZATION_TIER_UNSPECIFIED"
	OrganizationTierFree        OrganizationTier = "ORGANIZATION_TIER_FREE"
	OrganizationTierEnterprise  OrganizationTier = "ORGANIZATION_TIER_ENTERPRISE"
	OrganizationTierCore        OrganizationTier = "ORGANIZATION_TIER_CORE"
	OrganizationTierFreeOna     OrganizationTier = "ORGANIZATION_TIER_FREE_ONA"
)

func (r OrganizationTier) IsKnown() bool {
	switch r {
	case OrganizationTierUnspecified, OrganizationTierFree, OrganizationTierEnterprise, OrganizationTierCore, OrganizationTierFreeOna:
		return true
	}
	return false
}

type OrganizationNewResponse struct {
	// organization is the created organization
	Organization Organization `json:"organization,required"`
	// member is the member that joined the org on creation. Only set if specified
	// "join_organization" is "true" in the request.
	Member OrganizationMember          `json:"member"`
	JSON   organizationNewResponseJSON `json:"-"`
}

// organizationNewResponseJSON contains the JSON metadata for the struct
// [OrganizationNewResponse]
type organizationNewResponseJSON struct {
	Organization apijson.Field
	Member       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponse struct {
	// organization is the requested organization
	Organization Organization                `json:"organization,required"`
	JSON         organizationGetResponseJSON `json:"-"`
}

// organizationGetResponseJSON contains the JSON metadata for the struct
// [OrganizationGetResponse]
type organizationGetResponseJSON struct {
	Organization apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponse struct {
	// organization is the updated organization
	Organization Organization                   `json:"organization,required"`
	JSON         organizationUpdateResponseJSON `json:"-"`
}

// organizationUpdateResponseJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponse]
type organizationUpdateResponseJSON struct {
	Organization apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationDeleteResponse = interface{}

type OrganizationJoinResponse struct {
	// member is the member that was created by joining the organization.
	Member OrganizationMember           `json:"member,required"`
	JSON   organizationJoinResponseJSON `json:"-"`
}

// organizationJoinResponseJSON contains the JSON metadata for the struct
// [OrganizationJoinResponse]
type organizationJoinResponseJSON struct {
	Member      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationJoinResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationJoinResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationLeaveResponse = interface{}

type OrganizationSetRoleResponse = interface{}

type OrganizationNewParams struct {
	// name is the organization name
	Name param.Field[string] `json:"name,required"`
	// Should other Accounts with the same domain be automatically invited to the
	// organization?
	InviteAccountsWithMatchingDomain param.Field[bool] `json:"inviteAccountsWithMatchingDomain"`
	// join_organization decides whether the Identity issuing this request joins the
	// org on creation
	JoinOrganization param.Field[bool] `json:"joinOrganization"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationGetParams struct {
	// organization_id is the unique identifier of the Organization to retreive.
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r OrganizationGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateParams struct {
	// organization_id is the ID of the organization to update the settings for.
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
	// invite_domains is the domain allowlist of the organization
	InviteDomains param.Field[InviteDomainsParam] `json:"inviteDomains"`
	// name is the new name of the organization
	Name param.Field[string] `json:"name"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationDeleteParams struct {
	// organization_id is the ID of the organization to delete
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r OrganizationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationJoinParams struct {
	// invite_id is the unique identifier of the invite to join the organization.
	InviteID param.Field[string] `json:"inviteId" format:"uuid"`
	// organization_id is the unique identifier of the Organization to join.
	OrganizationID param.Field[string] `json:"organizationId" format:"uuid"`
}

func (r OrganizationJoinParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationLeaveParams struct {
	UserID param.Field[string] `json:"userId,required" format:"uuid"`
}

func (r OrganizationLeaveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationListMembersParams struct {
	// organization_id is the ID of the organization to list members for
	OrganizationID param.Field[string]                              `json:"organizationId,required" format:"uuid"`
	Token          param.Field[string]                              `query:"token"`
	PageSize       param.Field[int64]                               `query:"pageSize"`
	Filter         param.Field[OrganizationListMembersParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing members
	Pagination param.Field[OrganizationListMembersParamsPagination] `json:"pagination"`
	// sort specifies the order of results. When unspecified, the authenticated user is
	// returned first, followed by other members sorted by name ascending. When an
	// explicit sort is specified, results are sorted purely by the requested field
	// without any special handling for the authenticated user.
	Sort param.Field[OrganizationListMembersParamsSort] `json:"sort"`
}

func (r OrganizationListMembersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [OrganizationListMembersParams]'s query parameters as
// `url.Values`.
func (r OrganizationListMembersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationListMembersParamsFilter struct {
	// search performs case-insensitive search across member name and email
	Search param.Field[string] `json:"search"`
}

func (r OrganizationListMembersParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing members
type OrganizationListMembersParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r OrganizationListMembersParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// sort specifies the order of results. When unspecified, the authenticated user is
// returned first, followed by other members sorted by name ascending. When an
// explicit sort is specified, results are sorted purely by the requested field
// without any special handling for the authenticated user.
type OrganizationListMembersParamsSort struct {
	Field param.Field[OrganizationListMembersParamsSortField] `json:"field"`
	Order param.Field[OrganizationListMembersParamsSortOrder] `json:"order"`
}

func (r OrganizationListMembersParamsSort) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationListMembersParamsSortField string

const (
	OrganizationListMembersParamsSortFieldSortFieldUnspecified OrganizationListMembersParamsSortField = "SORT_FIELD_UNSPECIFIED"
	OrganizationListMembersParamsSortFieldSortFieldName        OrganizationListMembersParamsSortField = "SORT_FIELD_NAME"
	OrganizationListMembersParamsSortFieldSortFieldDateJoined  OrganizationListMembersParamsSortField = "SORT_FIELD_DATE_JOINED"
)

func (r OrganizationListMembersParamsSortField) IsKnown() bool {
	switch r {
	case OrganizationListMembersParamsSortFieldSortFieldUnspecified, OrganizationListMembersParamsSortFieldSortFieldName, OrganizationListMembersParamsSortFieldSortFieldDateJoined:
		return true
	}
	return false
}

type OrganizationListMembersParamsSortOrder string

const (
	OrganizationListMembersParamsSortOrderSortOrderUnspecified OrganizationListMembersParamsSortOrder = "SORT_ORDER_UNSPECIFIED"
	OrganizationListMembersParamsSortOrderSortOrderAsc         OrganizationListMembersParamsSortOrder = "SORT_ORDER_ASC"
	OrganizationListMembersParamsSortOrderSortOrderDesc        OrganizationListMembersParamsSortOrder = "SORT_ORDER_DESC"
)

func (r OrganizationListMembersParamsSortOrder) IsKnown() bool {
	switch r {
	case OrganizationListMembersParamsSortOrderSortOrderUnspecified, OrganizationListMembersParamsSortOrderSortOrderAsc, OrganizationListMembersParamsSortOrderSortOrderDesc:
		return true
	}
	return false
}

type OrganizationSetRoleParams struct {
	OrganizationID param.Field[string]                  `json:"organizationId,required" format:"uuid"`
	UserID         param.Field[string]                  `json:"userId,required" format:"uuid"`
	Role           param.Field[shared.OrganizationRole] `json:"role"`
}

func (r OrganizationSetRoleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
