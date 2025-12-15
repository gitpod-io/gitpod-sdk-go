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

// AccountService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Gets information about the currently authenticated account.
//
// Use this method to:
//
// - Retrieve account profile information
// - Check organization memberships
// - View account settings
// - Get joinable organizations
//
// ### Examples
//
// - Get account details:
//
//	Retrieves information about the authenticated account.
//
//	```yaml
//	{}
//	```
func (r *AccountService) Get(ctx context.Context, body AccountGetParams, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AccountService/GetAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes an account permanently.
//
// Use this method to:
//
// - Remove unused accounts
// - Clean up test accounts
// - Complete account deletion requests
//
// The account must not be an active member of any organization.
//
// ### Examples
//
// - Delete account:
//
//	Permanently removes an account.
//
//	```yaml
//	accountId: "f53d2330-3795-4c5d-a1f3-453121af9c60"
//	```
func (r *AccountService) Delete(ctx context.Context, body AccountDeleteParams, opts ...option.RequestOption) (res *AccountDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AccountService/DeleteAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets the SSO login URL for a specific email domain.
//
// Use this method to:
//
// - Initiate SSO authentication
// - Get organization-specific login URLs
// - Handle SSO redirects
//
// ### Examples
//
// - Get login URL:
//
//	Retrieves SSO URL for email domain.
//
//	```yaml
//	email: "user@company.com"
//	```
//
// - Get URL with return path:
//
//	Gets SSO URL with specific return location.
//
//	```yaml
//	email: "user@company.com"
//	returnTo: "https://gitpod.io/workspaces"
//	```
func (r *AccountService) GetSSOLoginURL(ctx context.Context, body AccountGetSSOLoginURLParams, opts ...option.RequestOption) (res *AccountGetSSOLoginURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "gitpod.v1.AccountService/GetSSOLoginURL"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists organizations that the currently authenticated account can join.
//
// Use this method to:
//
// - Discover organizations associated with the account's email domain.
// - Allow users to join existing organizations.
// - Display potential organizations during onboarding.
//
// ### Examples
//
// - List joinable organizations:
//
//	Retrieves a list of organizations the account can join.
//
//	```yaml
//	{}
//	```
func (r *AccountService) ListJoinableOrganizations(ctx context.Context, params AccountListJoinableOrganizationsParams, opts ...option.RequestOption) (res *pagination.JoinableOrganizationsPage[JoinableOrganization], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.AccountService/ListJoinableOrganizations"
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

// Lists organizations that the currently authenticated account can join.
//
// Use this method to:
//
// - Discover organizations associated with the account's email domain.
// - Allow users to join existing organizations.
// - Display potential organizations during onboarding.
//
// ### Examples
//
// - List joinable organizations:
//
//	Retrieves a list of organizations the account can join.
//
//	```yaml
//	{}
//	```
func (r *AccountService) ListJoinableOrganizationsAutoPaging(ctx context.Context, params AccountListJoinableOrganizationsParams, opts ...option.RequestOption) *pagination.JoinableOrganizationsPageAutoPager[JoinableOrganization] {
	return pagination.NewJoinableOrganizationsPageAutoPager(r.ListJoinableOrganizations(ctx, params, opts...))
}

// Lists available login providers with optional filtering.
//
// Use this method to:
//
// - View supported authentication methods
// - Get provider-specific login URLs
// - Filter providers by invite
//
// ### Examples
//
// - List all providers:
//
//	Shows all available login providers.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - List for specific invite:
//
//	Shows providers available for an invite.
//
//	```yaml
//	filter:
//	  inviteId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	pagination:
//	  pageSize: 20
//	```
func (r *AccountService) ListLoginProviders(ctx context.Context, params AccountListLoginProvidersParams, opts ...option.RequestOption) (res *pagination.LoginProvidersPage[LoginProvider], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.AccountService/ListLoginProviders"
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

// Lists available login providers with optional filtering.
//
// Use this method to:
//
// - View supported authentication methods
// - Get provider-specific login URLs
// - Filter providers by invite
//
// ### Examples
//
// - List all providers:
//
//	Shows all available login providers.
//
//	```yaml
//	pagination:
//	  pageSize: 20
//	```
//
// - List for specific invite:
//
//	Shows providers available for an invite.
//
//	```yaml
//	filter:
//	  inviteId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	pagination:
//	  pageSize: 20
//	```
func (r *AccountService) ListLoginProvidersAutoPaging(ctx context.Context, params AccountListLoginProvidersParams, opts ...option.RequestOption) *pagination.LoginProvidersPageAutoPager[LoginProvider] {
	return pagination.NewLoginProvidersPageAutoPager(r.ListLoginProviders(ctx, params, opts...))
}

// ListSSOLogins
func (r *AccountService) ListSSOLogins(ctx context.Context, params AccountListSSOLoginsParams, opts ...option.RequestOption) (res *pagination.LoginsPage[AccountListSSOLoginsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.AccountService/ListSSOLogins"
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

// ListSSOLogins
func (r *AccountService) ListSSOLoginsAutoPaging(ctx context.Context, params AccountListSSOLoginsParams, opts ...option.RequestOption) *pagination.LoginsPageAutoPager[AccountListSSOLoginsResponse] {
	return pagination.NewLoginsPageAutoPager(r.ListSSOLogins(ctx, params, opts...))
}

type Account struct {
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
	Email     string    `json:"email,required"`
	Name      string    `json:"name,required"`
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
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	AvatarURL string    `json:"avatarUrl"`
	// joinables is deprecated. Use ListJoinableOrganizations instead.
	//
	// Deprecated: deprecated
	Joinables   []JoinableOrganization `json:"joinables"`
	Memberships []AccountMembership    `json:"memberships"`
	// organization_id is the ID of the organization the account is owned by if it's
	// created through custom SSO
	OrganizationID string `json:"organizationId,nullable"`
	// public_email_provider is true if the email for the Account matches a known
	// public email provider
	PublicEmailProvider bool        `json:"publicEmailProvider"`
	JSON                accountJSON `json:"-"`
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	AvatarURL           apijson.Field
	Joinables           apijson.Field
	Memberships         apijson.Field
	OrganizationID      apijson.Field
	PublicEmailProvider apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountJSON) RawJSON() string {
	return r.raw
}

type AccountMembership struct {
	// organization_id is the id of the organization the user is a member of
	OrganizationID string `json:"organizationId,required" format:"uuid"`
	// organization_name is the name of the organization the user is a member of
	OrganizationName string `json:"organizationName,required"`
	// user_id is the ID the user has in the organization
	UserID string `json:"userId,required" format:"uuid"`
	// user_role is the role the user has in the organization
	UserRole shared.OrganizationRole `json:"userRole,required"`
	// organization_name is the member count of the organization the user is a member
	// of
	OrganizationMemberCount int64                 `json:"organizationMemberCount"`
	JSON                    accountMembershipJSON `json:"-"`
}

// accountMembershipJSON contains the JSON metadata for the struct
// [AccountMembership]
type accountMembershipJSON struct {
	OrganizationID          apijson.Field
	OrganizationName        apijson.Field
	UserID                  apijson.Field
	UserRole                apijson.Field
	OrganizationMemberCount apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMembershipJSON) RawJSON() string {
	return r.raw
}

type JoinableOrganization struct {
	// organization_id is the id of the organization the user can join
	OrganizationID string `json:"organizationId,required" format:"uuid"`
	// organization_name is the name of the organization the user can join
	OrganizationName string `json:"organizationName,required"`
	// organization_member_count is the member count of the organization the user can
	// join
	OrganizationMemberCount int64                    `json:"organizationMemberCount"`
	JSON                    joinableOrganizationJSON `json:"-"`
}

// joinableOrganizationJSON contains the JSON metadata for the struct
// [JoinableOrganization]
type joinableOrganizationJSON struct {
	OrganizationID          apijson.Field
	OrganizationName        apijson.Field
	OrganizationMemberCount apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *JoinableOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r joinableOrganizationJSON) RawJSON() string {
	return r.raw
}

type LoginProvider struct {
	// provider is the provider used by this login method, e.g. "github", "google",
	// "custom"
	Provider string `json:"provider,required"`
	// login_url is the URL to redirect the browser agent to for login, when provider
	// is "custom"
	LoginURL string            `json:"loginUrl"`
	JSON     loginProviderJSON `json:"-"`
}

// loginProviderJSON contains the JSON metadata for the struct [LoginProvider]
type loginProviderJSON struct {
	Provider    apijson.Field
	LoginURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoginProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loginProviderJSON) RawJSON() string {
	return r.raw
}

type AccountGetResponse struct {
	Account Account                `json:"account,required"`
	JSON    accountGetResponseJSON `json:"-"`
}

// accountGetResponseJSON contains the JSON metadata for the struct
// [AccountGetResponse]
type accountGetResponseJSON struct {
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDeleteResponse = interface{}

type AccountGetSSOLoginURLResponse struct {
	// login_url is the URL to redirect the user to for SSO login
	LoginURL string                            `json:"loginUrl,required"`
	JSON     accountGetSSOLoginURLResponseJSON `json:"-"`
}

// accountGetSSOLoginURLResponseJSON contains the JSON metadata for the struct
// [AccountGetSSOLoginURLResponse]
type accountGetSSOLoginURLResponseJSON struct {
	LoginURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetSSOLoginURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetSSOLoginURLResponseJSON) RawJSON() string {
	return r.raw
}

type AccountListSSOLoginsResponse struct {
	// provider is the provider used by this login method, e.g. "github", "google",
	// "custom"
	DisplayName string `json:"displayName,required"`
	// login_url is the URL to redirect the user to for SSO login
	LoginURL string                           `json:"loginUrl,required"`
	JSON     accountListSSOLoginsResponseJSON `json:"-"`
}

// accountListSSOLoginsResponseJSON contains the JSON metadata for the struct
// [AccountListSSOLoginsResponse]
type accountListSSOLoginsResponseJSON struct {
	DisplayName apijson.Field
	LoginURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListSSOLoginsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListSSOLoginsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetParams struct {
	Empty param.Field[bool] `json:"empty"`
}

func (r AccountGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDeleteParams struct {
	AccountID param.Field[string] `json:"accountId,required" format:"uuid"`
	// reason is an optional field for the reason for account deletion
	Reason param.Field[string] `json:"reason"`
}

func (r AccountDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGetSSOLoginURLParams struct {
	// email is the email the user wants to login with
	Email param.Field[string] `json:"email,required" format:"email"`
	// return_to is the URL the user will be redirected to after login
	ReturnTo param.Field[string] `json:"returnTo" format:"uri"`
}

func (r AccountGetSSOLoginURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountListJoinableOrganizationsParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// pagination contains the pagination options for listing joinable organizations
	Pagination param.Field[AccountListJoinableOrganizationsParamsPagination] `json:"pagination"`
}

func (r AccountListJoinableOrganizationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountListJoinableOrganizationsParams]'s query parameters
// as `url.Values`.
func (r AccountListJoinableOrganizationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// pagination contains the pagination options for listing joinable organizations
type AccountListJoinableOrganizationsParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r AccountListJoinableOrganizationsParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountListLoginProvidersParams struct {
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// filter contains the filter options for listing login methods
	Filter param.Field[AccountListLoginProvidersParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing login methods
	Pagination param.Field[AccountListLoginProvidersParamsPagination] `json:"pagination"`
}

func (r AccountListLoginProvidersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountListLoginProvidersParams]'s query parameters as
// `url.Values`.
func (r AccountListLoginProvidersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// filter contains the filter options for listing login methods
type AccountListLoginProvidersParamsFilter struct {
	// email is the email address to filter SSO providers by
	Email param.Field[string] `json:"email"`
	// invite_id is the ID of the invite URL the user wants to login with
	InviteID param.Field[string] `json:"inviteId" format:"uuid"`
}

func (r AccountListLoginProvidersParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing login methods
type AccountListLoginProvidersParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r AccountListLoginProvidersParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountListSSOLoginsParams struct {
	// email is the email the user wants to login with
	Email    param.Field[string] `json:"email,required" format:"email"`
	Token    param.Field[string] `query:"token"`
	PageSize param.Field[int64]  `query:"pageSize"`
	// pagination contains the pagination options for listing SSO logins
	Pagination param.Field[AccountListSSOLoginsParamsPagination] `json:"pagination"`
	// return_to is the URL the user will be redirected to after login
	ReturnTo param.Field[string] `json:"returnTo" format:"uri"`
}

func (r AccountListSSOLoginsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountListSSOLoginsParams]'s query parameters as
// `url.Values`.
func (r AccountListSSOLoginsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// pagination contains the pagination options for listing SSO logins
type AccountListSSOLoginsParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r AccountListSSOLoginsParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
