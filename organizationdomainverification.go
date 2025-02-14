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
)

// OrganizationDomainVerificationService contains methods and other services that
// help with interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationDomainVerificationService] method instead.
type OrganizationDomainVerificationService struct {
	Options []option.RequestOption
}

// NewOrganizationDomainVerificationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrganizationDomainVerificationService(opts ...option.RequestOption) (r *OrganizationDomainVerificationService) {
	r = &OrganizationDomainVerificationService{}
	r.Options = opts
	return
}

// Initiates domain verification process to enable organization features.
//
// Use this method to:
//
// - Start domain ownership verification
// - Enable automatic team joining
// - Set up SSO restrictions
// - Configure email-based policies
//
// ### Examples
//
// - Verify primary domain:
//
//	Starts verification for main company domain.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	domain: "acme-corp.com"
//	```
//
// - Verify subsidiary domain:
//
//	Adds verification for additional company domain.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	domain: "acme-subsidiary.com"
//	```
func (r *OrganizationDomainVerificationService) New(ctx context.Context, body OrganizationDomainVerificationNewParams, opts ...option.RequestOption) (res *OrganizationDomainVerificationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/CreateDomainVerification"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the status of a domain verification request.
//
// Use this method to:
//
// - Check verification progress
// - View verification requirements
// - Monitor domain status
//
// ### Examples
//
// - Get verification status:
//
//	Checks the current state of a domain verification.
//
//	```yaml
//	domainVerificationId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *OrganizationDomainVerificationService) Get(ctx context.Context, body OrganizationDomainVerificationGetParams, opts ...option.RequestOption) (res *OrganizationDomainVerificationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/GetDomainVerification"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists and monitors domain verification status across an organization.
//
// Use this method to:
//
// - Track verification progress
// - View all verified domains
// - Monitor pending verifications
// - Audit domain settings
//
// ### Examples
//
// - List all verifications:
//
//	Shows all domain verifications regardless of status.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	```
//
// - List with pagination:
//
//	Retrieves next page of verifications.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	  token: "next-page-token-from-previous-response"
//	```
func (r *OrganizationDomainVerificationService) List(ctx context.Context, params OrganizationDomainVerificationListParams, opts ...option.RequestOption) (res *pagination.DomainVerificationsPage[DomainVerification], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "gitpod.v1.OrganizationService/ListDomainVerifications"
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

// Lists and monitors domain verification status across an organization.
//
// Use this method to:
//
// - Track verification progress
// - View all verified domains
// - Monitor pending verifications
// - Audit domain settings
//
// ### Examples
//
// - List all verifications:
//
//	Shows all domain verifications regardless of status.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	```
//
// - List with pagination:
//
//	Retrieves next page of verifications.
//
//	```yaml
//	organizationId: "b0e12f6c-4c67-429d-a4a6-d9838b5da047"
//	pagination:
//	  pageSize: 20
//	  token: "next-page-token-from-previous-response"
//	```
func (r *OrganizationDomainVerificationService) ListAutoPaging(ctx context.Context, params OrganizationDomainVerificationListParams, opts ...option.RequestOption) *pagination.DomainVerificationsPageAutoPager[DomainVerification] {
	return pagination.NewDomainVerificationsPageAutoPager(r.List(ctx, params, opts...))
}

// Removes a domain verification request.
//
// Use this method to:
//
// - Cancel pending verifications
// - Remove verified domains
// - Clean up unused domain records
//
// ### Examples
//
// - Delete verification:
//
//	Removes a domain verification request.
//
//	```yaml
//	domainVerificationId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *OrganizationDomainVerificationService) Delete(ctx context.Context, body OrganizationDomainVerificationDeleteParams, opts ...option.RequestOption) (res *OrganizationDomainVerificationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/DeleteDomainVerification"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Verifies domain ownership for an organization.
//
// Use this method to:
//
// - Complete domain verification process
// - Enable domain-based features
// - Validate DNS configuration
//
// ### Examples
//
// - Verify domain ownership:
//
//	Verifies ownership after DNS records are configured.
//
//	```yaml
//	domainVerificationId: "d2c94c27-3b76-4a42-b88c-95a85e392c68"
//	```
func (r *OrganizationDomainVerificationService) Verify(ctx context.Context, body OrganizationDomainVerificationVerifyParams, opts ...option.RequestOption) (res *OrganizationDomainVerificationVerifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.OrganizationService/VerifyDomain"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DomainVerification struct {
	ID             string                  `json:"id,required" format:"uuid"`
	Domain         string                  `json:"domain,required"`
	OrganizationID string                  `json:"organizationId,required" format:"uuid"`
	State          DomainVerificationState `json:"state,required"`
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
	VerifiedAt time.Time              `json:"verifiedAt" format:"date-time"`
	JSON       domainVerificationJSON `json:"-"`
}

// domainVerificationJSON contains the JSON metadata for the struct
// [DomainVerification]
type domainVerificationJSON struct {
	ID             apijson.Field
	Domain         apijson.Field
	OrganizationID apijson.Field
	State          apijson.Field
	VerifiedAt     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DomainVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainVerificationJSON) RawJSON() string {
	return r.raw
}

type DomainVerificationState string

const (
	DomainVerificationStateUnspecified DomainVerificationState = "DOMAIN_VERIFICATION_STATE_UNSPECIFIED"
	DomainVerificationStatePending     DomainVerificationState = "DOMAIN_VERIFICATION_STATE_PENDING"
	DomainVerificationStateVerified    DomainVerificationState = "DOMAIN_VERIFICATION_STATE_VERIFIED"
)

func (r DomainVerificationState) IsKnown() bool {
	switch r {
	case DomainVerificationStateUnspecified, DomainVerificationStatePending, DomainVerificationStateVerified:
		return true
	}
	return false
}

type OrganizationDomainVerificationNewResponse struct {
	DomainVerification DomainVerification                            `json:"domainVerification,required"`
	JSON               organizationDomainVerificationNewResponseJSON `json:"-"`
}

// organizationDomainVerificationNewResponseJSON contains the JSON metadata for the
// struct [OrganizationDomainVerificationNewResponse]
type organizationDomainVerificationNewResponseJSON struct {
	DomainVerification apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationDomainVerificationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDomainVerificationNewResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationDomainVerificationGetResponse struct {
	DomainVerification DomainVerification                            `json:"domainVerification,required"`
	JSON               organizationDomainVerificationGetResponseJSON `json:"-"`
}

// organizationDomainVerificationGetResponseJSON contains the JSON metadata for the
// struct [OrganizationDomainVerificationGetResponse]
type organizationDomainVerificationGetResponseJSON struct {
	DomainVerification apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationDomainVerificationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDomainVerificationGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationDomainVerificationDeleteResponse = interface{}

type OrganizationDomainVerificationVerifyResponse struct {
	DomainVerification DomainVerification                               `json:"domainVerification,required"`
	JSON               organizationDomainVerificationVerifyResponseJSON `json:"-"`
}

// organizationDomainVerificationVerifyResponseJSON contains the JSON metadata for
// the struct [OrganizationDomainVerificationVerifyResponse]
type organizationDomainVerificationVerifyResponseJSON struct {
	DomainVerification apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationDomainVerificationVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDomainVerificationVerifyResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationDomainVerificationNewParams struct {
	Domain         param.Field[string] `json:"domain,required"`
	OrganizationID param.Field[string] `json:"organizationId,required" format:"uuid"`
}

func (r OrganizationDomainVerificationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationDomainVerificationGetParams struct {
	DomainVerificationID param.Field[string] `json:"domainVerificationId,required" format:"uuid"`
}

func (r OrganizationDomainVerificationGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationDomainVerificationListParams struct {
	OrganizationID param.Field[string]                                             `json:"organizationId,required" format:"uuid"`
	Token          param.Field[string]                                             `query:"token"`
	PageSize       param.Field[int64]                                              `query:"pageSize"`
	Pagination     param.Field[OrganizationDomainVerificationListParamsPagination] `json:"pagination"`
}

func (r OrganizationDomainVerificationListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [OrganizationDomainVerificationListParams]'s query
// parameters as `url.Values`.
func (r OrganizationDomainVerificationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationDomainVerificationListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r OrganizationDomainVerificationListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationDomainVerificationDeleteParams struct {
	DomainVerificationID param.Field[string] `json:"domainVerificationId,required" format:"uuid"`
}

func (r OrganizationDomainVerificationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationDomainVerificationVerifyParams struct {
	DomainVerificationID param.Field[string] `json:"domainVerificationId,required" format:"uuid"`
}

func (r OrganizationDomainVerificationVerifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
