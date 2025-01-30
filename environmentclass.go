// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// EnvironmentClassService contains methods and other services that help with
// interacting with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentClassService] method instead.
type EnvironmentClassService struct {
	Options []option.RequestOption
}

// NewEnvironmentClassService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnvironmentClassService(opts ...option.RequestOption) (r *EnvironmentClassService) {
	r = &EnvironmentClassService{}
	r.Options = opts
	return
}

// ListEnvironmentClasses returns the list of environment classes with runner
// details a user is able to use based on the
//
// query buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE
func (r *EnvironmentClassService) List(ctx context.Context, body EnvironmentClassListParams, opts ...option.RequestOption) (res *EnvironmentClassListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentService/ListEnvironmentClasses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EnvironmentClassListResponse struct {
	EnvironmentClasses []EnvironmentClassListResponseEnvironmentClass `json:"environmentClasses"`
	// pagination contains the pagination options for listing environment classes
	Pagination EnvironmentClassListResponsePagination `json:"pagination"`
	JSON       environmentClassListResponseJSON       `json:"-"`
}

// environmentClassListResponseJSON contains the JSON metadata for the struct
// [EnvironmentClassListResponse]
type environmentClassListResponseJSON struct {
	EnvironmentClasses apijson.Field
	Pagination         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentClassListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentClassListResponseEnvironmentClass struct {
	// id is the unique identifier of the environment class
	ID string `json:"id"`
	// configuration describes the configuration of the environment class
	Configuration []EnvironmentClassListResponseEnvironmentClassesConfiguration `json:"configuration"`
	// description is a human readable description of the environment class
	Description string `json:"description"`
	// display_name is the human readable name of the environment class
	DisplayName string `json:"displayName"`
	// enabled indicates whether the environment class can be used to create
	//
	// new environments.
	Enabled bool `json:"enabled"`
	// runner_id is the unique identifier of the runner the environment class belongs
	// to
	RunnerID string                                           `json:"runnerId"`
	JSON     environmentClassListResponseEnvironmentClassJSON `json:"-"`
}

// environmentClassListResponseEnvironmentClassJSON contains the JSON metadata for
// the struct [EnvironmentClassListResponseEnvironmentClass]
type environmentClassListResponseEnvironmentClassJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	RunnerID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EnvironmentClassListResponseEnvironmentClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassListResponseEnvironmentClassJSON) RawJSON() string {
	return r.raw
}

type EnvironmentClassListResponseEnvironmentClassesConfiguration struct {
	Key   string                                                          `json:"key"`
	Value string                                                          `json:"value"`
	JSON  environmentClassListResponseEnvironmentClassesConfigurationJSON `json:"-"`
}

// environmentClassListResponseEnvironmentClassesConfigurationJSON contains the
// JSON metadata for the struct
// [EnvironmentClassListResponseEnvironmentClassesConfiguration]
type environmentClassListResponseEnvironmentClassesConfigurationJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentClassListResponseEnvironmentClassesConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassListResponseEnvironmentClassesConfigurationJSON) RawJSON() string {
	return r.raw
}

// pagination contains the pagination options for listing environment classes
type EnvironmentClassListResponsePagination struct {
	// Token passed for retreiving the next set of results. Empty if there are no
	//
	// more results
	NextToken string                                     `json:"nextToken"`
	JSON      environmentClassListResponsePaginationJSON `json:"-"`
}

// environmentClassListResponsePaginationJSON contains the JSON metadata for the
// struct [EnvironmentClassListResponsePagination]
type environmentClassListResponsePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentClassListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type EnvironmentClassListParams struct {
	Filter param.Field[EnvironmentClassListParamsFilter] `json:"filter"`
	// pagination contains the pagination options for listing environment classes
	Pagination param.Field[EnvironmentClassListParamsPagination] `json:"pagination"`
}

func (r EnvironmentClassListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentClassListParamsFilter struct {
	// enabled filters the response to only enabled or disabled environment classes.
	//
	// If not set, all environment classes are returned.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r EnvironmentClassListParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pagination contains the pagination options for listing environment classes
type EnvironmentClassListParamsPagination struct {
	// Token for the next set of results that was returned as next_token of a
	//
	// PaginationResponse
	Token param.Field[string] `json:"token"`
	// Page size is the maximum number of results to retrieve per page. Defaults to 25.
	// Maximum 100.
	PageSize param.Field[int64] `json:"pageSize"`
}

func (r EnvironmentClassListParamsPagination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
