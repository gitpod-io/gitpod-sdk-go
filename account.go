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

// GetAccount retrieves a single Account.
func (r *AccountService) Get(ctx context.Context, body AccountGetParams, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/GetAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// DeleteAccount deletes an Account. To Delete an Account, the Account must not be
// an active member of any Organization.
func (r *AccountService) Delete(ctx context.Context, body AccountDeleteParams, opts ...option.RequestOption) (res *AccountDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/DeleteAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GetSSOLoginURL returns the URL to redirect the user to for SSO login.
func (r *AccountService) GetSSOLoginURL(ctx context.Context, body AccountGetSSOLoginURLParams, opts ...option.RequestOption) (res *AccountGetSSOLoginURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.AccountService/GetSSOLoginURL"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ListLoginProviders returns the list of login providers matching the provided
// filters.
func (r *AccountService) ListLoginProviders(ctx context.Context, params AccountListLoginProvidersParams, opts ...option.RequestOption) (res *pagination.LoginProvidersPage[LoginProvider], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// ListLoginProviders returns the list of login providers matching the provided
// filters.
func (r *AccountService) ListLoginProvidersAutoPaging(ctx context.Context, params AccountListLoginProvidersParams, opts ...option.RequestOption) *pagination.LoginProvidersPageAutoPager[LoginProvider] {
	return pagination.NewLoginProvidersPageAutoPager(r.ListLoginProviders(ctx, params, opts...))
}

type Account struct {
	ID        string `json:"id" format:"uuid"`
	AvatarURL string `json:"avatarUrl"`
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
	CreatedAt   time.Time              `json:"createdAt" format:"date-time"`
	Email       string                 `json:"email"`
	Joinables   []JoinableOrganization `json:"joinables"`
	Memberships []AccountMembership    `json:"memberships"`
	Name        string                 `json:"name"`
	// organization_id is the ID of the organization the account is owned by if it's
	// created through custom SSO
	OrganizationID string `json:"organizationId,nullable"`
	// public_email_provider is true if the email for the Account matches a known
	// public email provider
	PublicEmailProvider bool `json:"publicEmailProvider"`
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
	UpdatedAt time.Time   `json:"updatedAt" format:"date-time"`
	JSON      accountJSON `json:"-"`
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	ID                  apijson.Field
	AvatarURL           apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	Joinables           apijson.Field
	Memberships         apijson.Field
	Name                apijson.Field
	OrganizationID      apijson.Field
	PublicEmailProvider apijson.Field
	UpdatedAt           apijson.Field
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
	OrganizationID string `json:"organizationId" format:"uuid"`
	// organization_name is the member count of the organization the user is a member
	// of
	OrganizationMemberCount int64 `json:"organizationMemberCount"`
	// organization_name is the name of the organization the user is a member of
	OrganizationName string `json:"organizationName"`
	// user_id is the ID the user has in the organization
	UserID string `json:"userId" format:"uuid"`
	// user_role is the role the user has in the organization
	UserRole shared.OrganizationRole `json:"userRole"`
	JSON     accountMembershipJSON   `json:"-"`
}

// accountMembershipJSON contains the JSON metadata for the struct
// [AccountMembership]
type accountMembershipJSON struct {
	OrganizationID          apijson.Field
	OrganizationMemberCount apijson.Field
	OrganizationName        apijson.Field
	UserID                  apijson.Field
	UserRole                apijson.Field
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
	OrganizationID string `json:"organizationId" format:"uuid"`
	// organization_member_count is the member count of the organization the user can
	// join
	OrganizationMemberCount int64 `json:"organizationMemberCount"`
	// organization_name is the name of the organization the user can join
	OrganizationName string                   `json:"organizationName"`
	JSON             joinableOrganizationJSON `json:"-"`
}

// joinableOrganizationJSON contains the JSON metadata for the struct
// [JoinableOrganization]
type joinableOrganizationJSON struct {
	OrganizationID          apijson.Field
	OrganizationMemberCount apijson.Field
	OrganizationName        apijson.Field
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
	// login_url is the URL to redirect the browser agent to for login
	LoginURL string `json:"loginUrl"`
	// provider is the provider used by this login method, e.g. "github", "google",
	// "custom"
	Provider string            `json:"provider"`
	JSON     loginProviderJSON `json:"-"`
}

// loginProviderJSON contains the JSON metadata for the struct [LoginProvider]
type loginProviderJSON struct {
	LoginURL    apijson.Field
	Provider    apijson.Field
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
	Account Account                `json:"account"`
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
	LoginURL string                            `json:"loginUrl"`
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

type AccountGetParams struct {
	Empty param.Field[bool] `json:"empty"`
}

func (r AccountGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDeleteParams struct {
	AccountID param.Field[string] `json:"accountId" format:"uuid"`
}

func (r AccountDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGetSSOLoginURLParams struct {
	// email is the email the user wants to login with
	Email param.Field[string] `json:"email" format:"email"`
	// return_to is the URL the user will be redirected to after login
	ReturnTo param.Field[string] `json:"returnTo" format:"uri"`
}

func (r AccountGetSSOLoginURLParams) MarshalJSON() (data []byte, err error) {
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
