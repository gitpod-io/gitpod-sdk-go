// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/internal/requestconfig"
	"github.com/gitpod-io/gitpod-sdk-go/option"
)

type AgentExecutionsPagePagination struct {
	NextToken string                            `json:"nextToken"`
	JSON      agentExecutionsPagePaginationJSON `json:"-"`
}

// agentExecutionsPagePaginationJSON contains the JSON metadata for the struct
// [AgentExecutionsPagePagination]
type agentExecutionsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentExecutionsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type AgentExecutionsPage[T any] struct {
	AgentExecutions []T                           `json:"agentExecutions"`
	Pagination      AgentExecutionsPagePagination `json:"pagination"`
	JSON            agentExecutionsPageJSON       `json:"-"`
	cfg             *requestconfig.RequestConfig
	res             *http.Response
}

// agentExecutionsPageJSON contains the JSON metadata for the struct
// [AgentExecutionsPage[T]]
type agentExecutionsPageJSON struct {
	AgentExecutions apijson.Field
	Pagination      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AgentExecutionsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentExecutionsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *AgentExecutionsPage[T]) GetNextPage() (res *AgentExecutionsPage[T], err error) {
	if len(r.AgentExecutions) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *AgentExecutionsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &AgentExecutionsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type AgentExecutionsPageAutoPager[T any] struct {
	page *AgentExecutionsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewAgentExecutionsPageAutoPager[T any](page *AgentExecutionsPage[T], err error) *AgentExecutionsPageAutoPager[T] {
	return &AgentExecutionsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *AgentExecutionsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.AgentExecutions) == 0 {
		return false
	}
	if r.idx >= len(r.page.AgentExecutions) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.AgentExecutions) == 0 {
			return false
		}
	}
	r.cur = r.page.AgentExecutions[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *AgentExecutionsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *AgentExecutionsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *AgentExecutionsPageAutoPager[T]) Index() int {
	return r.run
}

type AssignmentsPagePagination struct {
	NextToken string                        `json:"nextToken"`
	JSON      assignmentsPagePaginationJSON `json:"-"`
}

// assignmentsPagePaginationJSON contains the JSON metadata for the struct
// [AssignmentsPagePagination]
type assignmentsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssignmentsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assignmentsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type AssignmentsPage[T any] struct {
	Assignments []T                       `json:"assignments"`
	Pagination  AssignmentsPagePagination `json:"pagination"`
	JSON        assignmentsPageJSON       `json:"-"`
	cfg         *requestconfig.RequestConfig
	res         *http.Response
}

// assignmentsPageJSON contains the JSON metadata for the struct
// [AssignmentsPage[T]]
type assignmentsPageJSON struct {
	Assignments apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssignmentsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assignmentsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *AssignmentsPage[T]) GetNextPage() (res *AssignmentsPage[T], err error) {
	if len(r.Assignments) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *AssignmentsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &AssignmentsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type AssignmentsPageAutoPager[T any] struct {
	page *AssignmentsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewAssignmentsPageAutoPager[T any](page *AssignmentsPage[T], err error) *AssignmentsPageAutoPager[T] {
	return &AssignmentsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *AssignmentsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Assignments) == 0 {
		return false
	}
	if r.idx >= len(r.page.Assignments) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Assignments) == 0 {
			return false
		}
	}
	r.cur = r.page.Assignments[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *AssignmentsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *AssignmentsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *AssignmentsPageAutoPager[T]) Index() int {
	return r.run
}

type DomainVerificationsPagePagination struct {
	NextToken string                                `json:"nextToken"`
	JSON      domainVerificationsPagePaginationJSON `json:"-"`
}

// domainVerificationsPagePaginationJSON contains the JSON metadata for the struct
// [DomainVerificationsPagePagination]
type domainVerificationsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainVerificationsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainVerificationsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type DomainVerificationsPage[T any] struct {
	DomainVerifications []T                               `json:"domainVerifications"`
	Pagination          DomainVerificationsPagePagination `json:"pagination"`
	JSON                domainVerificationsPageJSON       `json:"-"`
	cfg                 *requestconfig.RequestConfig
	res                 *http.Response
}

// domainVerificationsPageJSON contains the JSON metadata for the struct
// [DomainVerificationsPage[T]]
type domainVerificationsPageJSON struct {
	DomainVerifications apijson.Field
	Pagination          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DomainVerificationsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainVerificationsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *DomainVerificationsPage[T]) GetNextPage() (res *DomainVerificationsPage[T], err error) {
	if len(r.DomainVerifications) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *DomainVerificationsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &DomainVerificationsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type DomainVerificationsPageAutoPager[T any] struct {
	page *DomainVerificationsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewDomainVerificationsPageAutoPager[T any](page *DomainVerificationsPage[T], err error) *DomainVerificationsPageAutoPager[T] {
	return &DomainVerificationsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *DomainVerificationsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.DomainVerifications) == 0 {
		return false
	}
	if r.idx >= len(r.page.DomainVerifications) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.DomainVerifications) == 0 {
			return false
		}
	}
	r.cur = r.page.DomainVerifications[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *DomainVerificationsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *DomainVerificationsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *DomainVerificationsPageAutoPager[T]) Index() int {
	return r.run
}

type EditorsPagePagination struct {
	NextToken string                    `json:"nextToken"`
	JSON      editorsPagePaginationJSON `json:"-"`
}

// editorsPagePaginationJSON contains the JSON metadata for the struct
// [EditorsPagePagination]
type editorsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditorsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editorsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type EditorsPage[T any] struct {
	Editors    []T                   `json:"editors"`
	Pagination EditorsPagePagination `json:"pagination"`
	JSON       editorsPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// editorsPageJSON contains the JSON metadata for the struct [EditorsPage[T]]
type editorsPageJSON struct {
	Editors     apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditorsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editorsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *EditorsPage[T]) GetNextPage() (res *EditorsPage[T], err error) {
	if len(r.Editors) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *EditorsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &EditorsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type EditorsPageAutoPager[T any] struct {
	page *EditorsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewEditorsPageAutoPager[T any](page *EditorsPage[T], err error) *EditorsPageAutoPager[T] {
	return &EditorsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *EditorsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Editors) == 0 {
		return false
	}
	if r.idx >= len(r.page.Editors) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Editors) == 0 {
			return false
		}
	}
	r.cur = r.page.Editors[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *EditorsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *EditorsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *EditorsPageAutoPager[T]) Index() int {
	return r.run
}

type EntriesPagePagination struct {
	NextToken string                    `json:"nextToken"`
	JSON      entriesPagePaginationJSON `json:"-"`
}

// entriesPagePaginationJSON contains the JSON metadata for the struct
// [EntriesPagePagination]
type entriesPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntriesPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entriesPagePaginationJSON) RawJSON() string {
	return r.raw
}

type EntriesPage[T any] struct {
	Entries    []T                   `json:"entries"`
	Pagination EntriesPagePagination `json:"pagination"`
	JSON       entriesPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// entriesPageJSON contains the JSON metadata for the struct [EntriesPage[T]]
type entriesPageJSON struct {
	Entries     apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntriesPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entriesPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *EntriesPage[T]) GetNextPage() (res *EntriesPage[T], err error) {
	if len(r.Entries) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *EntriesPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &EntriesPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type EntriesPageAutoPager[T any] struct {
	page *EntriesPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewEntriesPageAutoPager[T any](page *EntriesPage[T], err error) *EntriesPageAutoPager[T] {
	return &EntriesPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *EntriesPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Entries) == 0 {
		return false
	}
	if r.idx >= len(r.page.Entries) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Entries) == 0 {
			return false
		}
	}
	r.cur = r.page.Entries[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *EntriesPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *EntriesPageAutoPager[T]) Err() error {
	return r.err
}

func (r *EntriesPageAutoPager[T]) Index() int {
	return r.run
}

type EnvironmentClassesPagePagination struct {
	NextToken string                               `json:"nextToken"`
	JSON      environmentClassesPagePaginationJSON `json:"-"`
}

// environmentClassesPagePaginationJSON contains the JSON metadata for the struct
// [EnvironmentClassesPagePagination]
type environmentClassesPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentClassesPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassesPagePaginationJSON) RawJSON() string {
	return r.raw
}

type EnvironmentClassesPage[T any] struct {
	EnvironmentClasses []T                              `json:"environmentClasses"`
	Pagination         EnvironmentClassesPagePagination `json:"pagination"`
	JSON               environmentClassesPageJSON       `json:"-"`
	cfg                *requestconfig.RequestConfig
	res                *http.Response
}

// environmentClassesPageJSON contains the JSON metadata for the struct
// [EnvironmentClassesPage[T]]
type environmentClassesPageJSON struct {
	EnvironmentClasses apijson.Field
	Pagination         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentClassesPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentClassesPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *EnvironmentClassesPage[T]) GetNextPage() (res *EnvironmentClassesPage[T], err error) {
	if len(r.EnvironmentClasses) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *EnvironmentClassesPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &EnvironmentClassesPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type EnvironmentClassesPageAutoPager[T any] struct {
	page *EnvironmentClassesPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewEnvironmentClassesPageAutoPager[T any](page *EnvironmentClassesPage[T], err error) *EnvironmentClassesPageAutoPager[T] {
	return &EnvironmentClassesPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *EnvironmentClassesPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.EnvironmentClasses) == 0 {
		return false
	}
	if r.idx >= len(r.page.EnvironmentClasses) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.EnvironmentClasses) == 0 {
			return false
		}
	}
	r.cur = r.page.EnvironmentClasses[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *EnvironmentClassesPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *EnvironmentClassesPageAutoPager[T]) Err() error {
	return r.err
}

func (r *EnvironmentClassesPageAutoPager[T]) Index() int {
	return r.run
}

type EnvironmentsPagePagination struct {
	NextToken string                         `json:"nextToken"`
	JSON      environmentsPagePaginationJSON `json:"-"`
}

// environmentsPagePaginationJSON contains the JSON metadata for the struct
// [EnvironmentsPagePagination]
type environmentsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type EnvironmentsPage[T any] struct {
	Environments []T                        `json:"environments"`
	Pagination   EnvironmentsPagePagination `json:"pagination"`
	JSON         environmentsPageJSON       `json:"-"`
	cfg          *requestconfig.RequestConfig
	res          *http.Response
}

// environmentsPageJSON contains the JSON metadata for the struct
// [EnvironmentsPage[T]]
type environmentsPageJSON struct {
	Environments apijson.Field
	Pagination   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *EnvironmentsPage[T]) GetNextPage() (res *EnvironmentsPage[T], err error) {
	if len(r.Environments) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *EnvironmentsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &EnvironmentsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type EnvironmentsPageAutoPager[T any] struct {
	page *EnvironmentsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewEnvironmentsPageAutoPager[T any](page *EnvironmentsPage[T], err error) *EnvironmentsPageAutoPager[T] {
	return &EnvironmentsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *EnvironmentsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Environments) == 0 {
		return false
	}
	if r.idx >= len(r.page.Environments) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Environments) == 0 {
			return false
		}
	}
	r.cur = r.page.Environments[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *EnvironmentsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *EnvironmentsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *EnvironmentsPageAutoPager[T]) Index() int {
	return r.run
}

type GatewaysPagePagination struct {
	NextToken string                     `json:"nextToken"`
	JSON      gatewaysPagePaginationJSON `json:"-"`
}

// gatewaysPagePaginationJSON contains the JSON metadata for the struct
// [GatewaysPagePagination]
type gatewaysPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewaysPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewaysPagePaginationJSON) RawJSON() string {
	return r.raw
}

type GatewaysPage[T any] struct {
	Gateways   []T                    `json:"gateways"`
	Pagination GatewaysPagePagination `json:"pagination"`
	JSON       gatewaysPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// gatewaysPageJSON contains the JSON metadata for the struct [GatewaysPage[T]]
type gatewaysPageJSON struct {
	Gateways    apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewaysPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewaysPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *GatewaysPage[T]) GetNextPage() (res *GatewaysPage[T], err error) {
	if len(r.Gateways) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *GatewaysPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &GatewaysPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type GatewaysPageAutoPager[T any] struct {
	page *GatewaysPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewGatewaysPageAutoPager[T any](page *GatewaysPage[T], err error) *GatewaysPageAutoPager[T] {
	return &GatewaysPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *GatewaysPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Gateways) == 0 {
		return false
	}
	if r.idx >= len(r.page.Gateways) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Gateways) == 0 {
			return false
		}
	}
	r.cur = r.page.Gateways[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *GatewaysPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *GatewaysPageAutoPager[T]) Err() error {
	return r.err
}

func (r *GatewaysPageAutoPager[T]) Index() int {
	return r.run
}

type GroupsPagePagination struct {
	NextToken string                   `json:"nextToken"`
	JSON      groupsPagePaginationJSON `json:"-"`
}

// groupsPagePaginationJSON contains the JSON metadata for the struct
// [GroupsPagePagination]
type groupsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GroupsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type GroupsPage[T any] struct {
	Groups     []T                  `json:"groups"`
	Pagination GroupsPagePagination `json:"pagination"`
	JSON       groupsPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// groupsPageJSON contains the JSON metadata for the struct [GroupsPage[T]]
type groupsPageJSON struct {
	Groups      apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GroupsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *GroupsPage[T]) GetNextPage() (res *GroupsPage[T], err error) {
	if len(r.Groups) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *GroupsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &GroupsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type GroupsPageAutoPager[T any] struct {
	page *GroupsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewGroupsPageAutoPager[T any](page *GroupsPage[T], err error) *GroupsPageAutoPager[T] {
	return &GroupsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *GroupsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Groups) == 0 {
		return false
	}
	if r.idx >= len(r.page.Groups) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Groups) == 0 {
			return false
		}
	}
	r.cur = r.page.Groups[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *GroupsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *GroupsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *GroupsPageAutoPager[T]) Index() int {
	return r.run
}

type IntegrationsPagePagination struct {
	NextToken string                         `json:"nextToken"`
	JSON      integrationsPagePaginationJSON `json:"-"`
}

// integrationsPagePaginationJSON contains the JSON metadata for the struct
// [IntegrationsPagePagination]
type integrationsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type IntegrationsPage[T any] struct {
	Integrations []T                        `json:"integrations"`
	Pagination   IntegrationsPagePagination `json:"pagination"`
	JSON         integrationsPageJSON       `json:"-"`
	cfg          *requestconfig.RequestConfig
	res          *http.Response
}

// integrationsPageJSON contains the JSON metadata for the struct
// [IntegrationsPage[T]]
type integrationsPageJSON struct {
	Integrations apijson.Field
	Pagination   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IntegrationsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *IntegrationsPage[T]) GetNextPage() (res *IntegrationsPage[T], err error) {
	if len(r.Integrations) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *IntegrationsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &IntegrationsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type IntegrationsPageAutoPager[T any] struct {
	page *IntegrationsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewIntegrationsPageAutoPager[T any](page *IntegrationsPage[T], err error) *IntegrationsPageAutoPager[T] {
	return &IntegrationsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *IntegrationsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Integrations) == 0 {
		return false
	}
	if r.idx >= len(r.page.Integrations) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Integrations) == 0 {
			return false
		}
	}
	r.cur = r.page.Integrations[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *IntegrationsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *IntegrationsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *IntegrationsPageAutoPager[T]) Index() int {
	return r.run
}

type JoinableOrganizationsPagePagination struct {
	NextToken string                                  `json:"nextToken"`
	JSON      joinableOrganizationsPagePaginationJSON `json:"-"`
}

// joinableOrganizationsPagePaginationJSON contains the JSON metadata for the
// struct [JoinableOrganizationsPagePagination]
type joinableOrganizationsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JoinableOrganizationsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r joinableOrganizationsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type JoinableOrganizationsPage[T any] struct {
	JoinableOrganizations []T                                 `json:"joinableOrganizations"`
	Pagination            JoinableOrganizationsPagePagination `json:"pagination"`
	JSON                  joinableOrganizationsPageJSON       `json:"-"`
	cfg                   *requestconfig.RequestConfig
	res                   *http.Response
}

// joinableOrganizationsPageJSON contains the JSON metadata for the struct
// [JoinableOrganizationsPage[T]]
type joinableOrganizationsPageJSON struct {
	JoinableOrganizations apijson.Field
	Pagination            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *JoinableOrganizationsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r joinableOrganizationsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *JoinableOrganizationsPage[T]) GetNextPage() (res *JoinableOrganizationsPage[T], err error) {
	if len(r.JoinableOrganizations) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *JoinableOrganizationsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &JoinableOrganizationsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type JoinableOrganizationsPageAutoPager[T any] struct {
	page *JoinableOrganizationsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewJoinableOrganizationsPageAutoPager[T any](page *JoinableOrganizationsPage[T], err error) *JoinableOrganizationsPageAutoPager[T] {
	return &JoinableOrganizationsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *JoinableOrganizationsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.JoinableOrganizations) == 0 {
		return false
	}
	if r.idx >= len(r.page.JoinableOrganizations) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.JoinableOrganizations) == 0 {
			return false
		}
	}
	r.cur = r.page.JoinableOrganizations[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *JoinableOrganizationsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *JoinableOrganizationsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *JoinableOrganizationsPageAutoPager[T]) Index() int {
	return r.run
}

type LoginProvidersPagePagination struct {
	NextToken string                           `json:"nextToken"`
	JSON      loginProvidersPagePaginationJSON `json:"-"`
}

// loginProvidersPagePaginationJSON contains the JSON metadata for the struct
// [LoginProvidersPagePagination]
type loginProvidersPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoginProvidersPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loginProvidersPagePaginationJSON) RawJSON() string {
	return r.raw
}

type LoginProvidersPage[T any] struct {
	LoginProviders []T                          `json:"loginProviders"`
	Pagination     LoginProvidersPagePagination `json:"pagination"`
	JSON           loginProvidersPageJSON       `json:"-"`
	cfg            *requestconfig.RequestConfig
	res            *http.Response
}

// loginProvidersPageJSON contains the JSON metadata for the struct
// [LoginProvidersPage[T]]
type loginProvidersPageJSON struct {
	LoginProviders apijson.Field
	Pagination     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoginProvidersPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loginProvidersPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *LoginProvidersPage[T]) GetNextPage() (res *LoginProvidersPage[T], err error) {
	if len(r.LoginProviders) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *LoginProvidersPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &LoginProvidersPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type LoginProvidersPageAutoPager[T any] struct {
	page *LoginProvidersPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewLoginProvidersPageAutoPager[T any](page *LoginProvidersPage[T], err error) *LoginProvidersPageAutoPager[T] {
	return &LoginProvidersPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *LoginProvidersPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.LoginProviders) == 0 {
		return false
	}
	if r.idx >= len(r.page.LoginProviders) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.LoginProviders) == 0 {
			return false
		}
	}
	r.cur = r.page.LoginProviders[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *LoginProvidersPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *LoginProvidersPageAutoPager[T]) Err() error {
	return r.err
}

func (r *LoginProvidersPageAutoPager[T]) Index() int {
	return r.run
}

type LoginsPagePagination struct {
	NextToken string                   `json:"nextToken"`
	JSON      loginsPagePaginationJSON `json:"-"`
}

// loginsPagePaginationJSON contains the JSON metadata for the struct
// [LoginsPagePagination]
type loginsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoginsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loginsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type LoginsPage[T any] struct {
	Logins     []T                  `json:"logins"`
	Pagination LoginsPagePagination `json:"pagination"`
	JSON       loginsPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// loginsPageJSON contains the JSON metadata for the struct [LoginsPage[T]]
type loginsPageJSON struct {
	Logins      apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoginsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loginsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *LoginsPage[T]) GetNextPage() (res *LoginsPage[T], err error) {
	if len(r.Logins) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *LoginsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &LoginsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type LoginsPageAutoPager[T any] struct {
	page *LoginsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewLoginsPageAutoPager[T any](page *LoginsPage[T], err error) *LoginsPageAutoPager[T] {
	return &LoginsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *LoginsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Logins) == 0 {
		return false
	}
	if r.idx >= len(r.page.Logins) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Logins) == 0 {
			return false
		}
	}
	r.cur = r.page.Logins[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *LoginsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *LoginsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *LoginsPageAutoPager[T]) Index() int {
	return r.run
}

type MembersPagePagination struct {
	NextToken string                    `json:"nextToken"`
	JSON      membersPagePaginationJSON `json:"-"`
}

// membersPagePaginationJSON contains the JSON metadata for the struct
// [MembersPagePagination]
type membersPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembersPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membersPagePaginationJSON) RawJSON() string {
	return r.raw
}

type MembersPage[T any] struct {
	Members    []T                   `json:"members"`
	Pagination MembersPagePagination `json:"pagination"`
	JSON       membersPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// membersPageJSON contains the JSON metadata for the struct [MembersPage[T]]
type membersPageJSON struct {
	Members     apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembersPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membersPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *MembersPage[T]) GetNextPage() (res *MembersPage[T], err error) {
	if len(r.Members) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *MembersPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &MembersPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type MembersPageAutoPager[T any] struct {
	page *MembersPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewMembersPageAutoPager[T any](page *MembersPage[T], err error) *MembersPageAutoPager[T] {
	return &MembersPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *MembersPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Members) == 0 {
		return false
	}
	if r.idx >= len(r.page.Members) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Members) == 0 {
			return false
		}
	}
	r.cur = r.page.Members[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *MembersPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *MembersPageAutoPager[T]) Err() error {
	return r.err
}

func (r *MembersPageAutoPager[T]) Index() int {
	return r.run
}

type PersonalAccessTokensPagePagination struct {
	NextToken string                                 `json:"nextToken"`
	JSON      personalAccessTokensPagePaginationJSON `json:"-"`
}

// personalAccessTokensPagePaginationJSON contains the JSON metadata for the struct
// [PersonalAccessTokensPagePagination]
type personalAccessTokensPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PersonalAccessTokensPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r personalAccessTokensPagePaginationJSON) RawJSON() string {
	return r.raw
}

type PersonalAccessTokensPage[T any] struct {
	Pagination           PersonalAccessTokensPagePagination `json:"pagination"`
	PersonalAccessTokens []T                                `json:"personalAccessTokens"`
	JSON                 personalAccessTokensPageJSON       `json:"-"`
	cfg                  *requestconfig.RequestConfig
	res                  *http.Response
}

// personalAccessTokensPageJSON contains the JSON metadata for the struct
// [PersonalAccessTokensPage[T]]
type personalAccessTokensPageJSON struct {
	Pagination           apijson.Field
	PersonalAccessTokens apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PersonalAccessTokensPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r personalAccessTokensPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PersonalAccessTokensPage[T]) GetNextPage() (res *PersonalAccessTokensPage[T], err error) {
	if len(r.PersonalAccessTokens) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PersonalAccessTokensPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PersonalAccessTokensPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PersonalAccessTokensPageAutoPager[T any] struct {
	page *PersonalAccessTokensPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPersonalAccessTokensPageAutoPager[T any](page *PersonalAccessTokensPage[T], err error) *PersonalAccessTokensPageAutoPager[T] {
	return &PersonalAccessTokensPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PersonalAccessTokensPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.PersonalAccessTokens) == 0 {
		return false
	}
	if r.idx >= len(r.page.PersonalAccessTokens) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.PersonalAccessTokens) == 0 {
			return false
		}
	}
	r.cur = r.page.PersonalAccessTokens[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PersonalAccessTokensPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PersonalAccessTokensPageAutoPager[T]) Err() error {
	return r.err
}

func (r *PersonalAccessTokensPageAutoPager[T]) Index() int {
	return r.run
}

type PoliciesPagePagination struct {
	NextToken string                     `json:"nextToken"`
	JSON      policiesPagePaginationJSON `json:"-"`
}

// policiesPagePaginationJSON contains the JSON metadata for the struct
// [PoliciesPagePagination]
type policiesPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policiesPagePaginationJSON) RawJSON() string {
	return r.raw
}

type PoliciesPage[T any] struct {
	Pagination PoliciesPagePagination `json:"pagination"`
	Policies   []T                    `json:"policies"`
	JSON       policiesPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// policiesPageJSON contains the JSON metadata for the struct [PoliciesPage[T]]
type policiesPageJSON struct {
	Pagination  apijson.Field
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policiesPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PoliciesPage[T]) GetNextPage() (res *PoliciesPage[T], err error) {
	if len(r.Policies) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PoliciesPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PoliciesPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PoliciesPageAutoPager[T any] struct {
	page *PoliciesPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPoliciesPageAutoPager[T any](page *PoliciesPage[T], err error) *PoliciesPageAutoPager[T] {
	return &PoliciesPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PoliciesPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Policies) == 0 {
		return false
	}
	if r.idx >= len(r.page.Policies) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Policies) == 0 {
			return false
		}
	}
	r.cur = r.page.Policies[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PoliciesPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PoliciesPageAutoPager[T]) Err() error {
	return r.err
}

func (r *PoliciesPageAutoPager[T]) Index() int {
	return r.run
}

type PrebuildsPagePagination struct {
	NextToken string                      `json:"nextToken"`
	JSON      prebuildsPagePaginationJSON `json:"-"`
}

// prebuildsPagePaginationJSON contains the JSON metadata for the struct
// [PrebuildsPagePagination]
type prebuildsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrebuildsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prebuildsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type PrebuildsPage[T any] struct {
	Pagination PrebuildsPagePagination `json:"pagination"`
	Prebuilds  []T                     `json:"prebuilds"`
	JSON       prebuildsPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// prebuildsPageJSON contains the JSON metadata for the struct [PrebuildsPage[T]]
type prebuildsPageJSON struct {
	Pagination  apijson.Field
	Prebuilds   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrebuildsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prebuildsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PrebuildsPage[T]) GetNextPage() (res *PrebuildsPage[T], err error) {
	if len(r.Prebuilds) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PrebuildsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PrebuildsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PrebuildsPageAutoPager[T any] struct {
	page *PrebuildsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPrebuildsPageAutoPager[T any](page *PrebuildsPage[T], err error) *PrebuildsPageAutoPager[T] {
	return &PrebuildsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PrebuildsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Prebuilds) == 0 {
		return false
	}
	if r.idx >= len(r.page.Prebuilds) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Prebuilds) == 0 {
			return false
		}
	}
	r.cur = r.page.Prebuilds[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PrebuildsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PrebuildsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *PrebuildsPageAutoPager[T]) Index() int {
	return r.run
}

type ProjectEnvironmentClassesPagePagination struct {
	NextToken string                                      `json:"nextToken"`
	JSON      projectEnvironmentClassesPagePaginationJSON `json:"-"`
}

// projectEnvironmentClassesPagePaginationJSON contains the JSON metadata for the
// struct [ProjectEnvironmentClassesPagePagination]
type projectEnvironmentClassesPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEnvironmentClassesPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEnvironmentClassesPagePaginationJSON) RawJSON() string {
	return r.raw
}

type ProjectEnvironmentClassesPage[T any] struct {
	Pagination                ProjectEnvironmentClassesPagePagination `json:"pagination"`
	ProjectEnvironmentClasses []T                                     `json:"projectEnvironmentClasses"`
	JSON                      projectEnvironmentClassesPageJSON       `json:"-"`
	cfg                       *requestconfig.RequestConfig
	res                       *http.Response
}

// projectEnvironmentClassesPageJSON contains the JSON metadata for the struct
// [ProjectEnvironmentClassesPage[T]]
type projectEnvironmentClassesPageJSON struct {
	Pagination                apijson.Field
	ProjectEnvironmentClasses apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ProjectEnvironmentClassesPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEnvironmentClassesPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ProjectEnvironmentClassesPage[T]) GetNextPage() (res *ProjectEnvironmentClassesPage[T], err error) {
	if len(r.ProjectEnvironmentClasses) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ProjectEnvironmentClassesPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ProjectEnvironmentClassesPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ProjectEnvironmentClassesPageAutoPager[T any] struct {
	page *ProjectEnvironmentClassesPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewProjectEnvironmentClassesPageAutoPager[T any](page *ProjectEnvironmentClassesPage[T], err error) *ProjectEnvironmentClassesPageAutoPager[T] {
	return &ProjectEnvironmentClassesPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ProjectEnvironmentClassesPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.ProjectEnvironmentClasses) == 0 {
		return false
	}
	if r.idx >= len(r.page.ProjectEnvironmentClasses) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.ProjectEnvironmentClasses) == 0 {
			return false
		}
	}
	r.cur = r.page.ProjectEnvironmentClasses[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ProjectEnvironmentClassesPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ProjectEnvironmentClassesPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ProjectEnvironmentClassesPageAutoPager[T]) Index() int {
	return r.run
}

type ProjectsPagePagination struct {
	NextToken string                     `json:"nextToken"`
	JSON      projectsPagePaginationJSON `json:"-"`
}

// projectsPagePaginationJSON contains the JSON metadata for the struct
// [ProjectsPagePagination]
type projectsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type ProjectsPage[T any] struct {
	Pagination ProjectsPagePagination `json:"pagination"`
	Projects   []T                    `json:"projects"`
	JSON       projectsPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// projectsPageJSON contains the JSON metadata for the struct [ProjectsPage[T]]
type projectsPageJSON struct {
	Pagination  apijson.Field
	Projects    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ProjectsPage[T]) GetNextPage() (res *ProjectsPage[T], err error) {
	if len(r.Projects) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ProjectsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ProjectsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ProjectsPageAutoPager[T any] struct {
	page *ProjectsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewProjectsPageAutoPager[T any](page *ProjectsPage[T], err error) *ProjectsPageAutoPager[T] {
	return &ProjectsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ProjectsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Projects) == 0 {
		return false
	}
	if r.idx >= len(r.page.Projects) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Projects) == 0 {
			return false
		}
	}
	r.cur = r.page.Projects[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ProjectsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ProjectsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ProjectsPageAutoPager[T]) Index() int {
	return r.run
}

type PromptsPagePagination struct {
	NextToken string                    `json:"nextToken"`
	JSON      promptsPagePaginationJSON `json:"-"`
}

// promptsPagePaginationJSON contains the JSON metadata for the struct
// [PromptsPagePagination]
type promptsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type PromptsPage[T any] struct {
	Pagination PromptsPagePagination `json:"pagination"`
	Prompts    []T                   `json:"prompts"`
	JSON       promptsPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// promptsPageJSON contains the JSON metadata for the struct [PromptsPage[T]]
type promptsPageJSON struct {
	Pagination  apijson.Field
	Prompts     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PromptsPage[T]) GetNextPage() (res *PromptsPage[T], err error) {
	if len(r.Prompts) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PromptsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PromptsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PromptsPageAutoPager[T any] struct {
	page *PromptsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPromptsPageAutoPager[T any](page *PromptsPage[T], err error) *PromptsPageAutoPager[T] {
	return &PromptsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PromptsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Prompts) == 0 {
		return false
	}
	if r.idx >= len(r.page.Prompts) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Prompts) == 0 {
			return false
		}
	}
	r.cur = r.page.Prompts[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PromptsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PromptsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *PromptsPageAutoPager[T]) Index() int {
	return r.run
}

type RecordsPagePagination struct {
	NextToken string                    `json:"nextToken"`
	JSON      recordsPagePaginationJSON `json:"-"`
}

// recordsPagePaginationJSON contains the JSON metadata for the struct
// [RecordsPagePagination]
type recordsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type RecordsPage[T any] struct {
	Pagination RecordsPagePagination `json:"pagination"`
	Records    []T                   `json:"records"`
	JSON       recordsPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// recordsPageJSON contains the JSON metadata for the struct [RecordsPage[T]]
type recordsPageJSON struct {
	Pagination  apijson.Field
	Records     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *RecordsPage[T]) GetNextPage() (res *RecordsPage[T], err error) {
	if len(r.Records) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *RecordsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &RecordsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type RecordsPageAutoPager[T any] struct {
	page *RecordsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewRecordsPageAutoPager[T any](page *RecordsPage[T], err error) *RecordsPageAutoPager[T] {
	return &RecordsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *RecordsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Records) == 0 {
		return false
	}
	if r.idx >= len(r.page.Records) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Records) == 0 {
			return false
		}
	}
	r.cur = r.page.Records[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *RecordsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *RecordsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *RecordsPageAutoPager[T]) Index() int {
	return r.run
}

type RepositoriesPagePagination struct {
	NextToken string                         `json:"nextToken"`
	JSON      repositoriesPagePaginationJSON `json:"-"`
}

// repositoriesPagePaginationJSON contains the JSON metadata for the struct
// [RepositoriesPagePagination]
type repositoriesPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RepositoriesPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoriesPagePaginationJSON) RawJSON() string {
	return r.raw
}

type RepositoriesPage[T any] struct {
	Pagination   RepositoriesPagePagination `json:"pagination"`
	Repositories []T                        `json:"repositories"`
	JSON         repositoriesPageJSON       `json:"-"`
	cfg          *requestconfig.RequestConfig
	res          *http.Response
}

// repositoriesPageJSON contains the JSON metadata for the struct
// [RepositoriesPage[T]]
type repositoriesPageJSON struct {
	Pagination   apijson.Field
	Repositories apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RepositoriesPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repositoriesPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *RepositoriesPage[T]) GetNextPage() (res *RepositoriesPage[T], err error) {
	if len(r.Repositories) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *RepositoriesPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &RepositoriesPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type RepositoriesPageAutoPager[T any] struct {
	page *RepositoriesPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewRepositoriesPageAutoPager[T any](page *RepositoriesPage[T], err error) *RepositoriesPageAutoPager[T] {
	return &RepositoriesPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *RepositoriesPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Repositories) == 0 {
		return false
	}
	if r.idx >= len(r.page.Repositories) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Repositories) == 0 {
			return false
		}
	}
	r.cur = r.page.Repositories[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *RepositoriesPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *RepositoriesPageAutoPager[T]) Err() error {
	return r.err
}

func (r *RepositoriesPageAutoPager[T]) Index() int {
	return r.run
}

type RunnersPagePagination struct {
	NextToken string                    `json:"nextToken"`
	JSON      runnersPagePaginationJSON `json:"-"`
}

// runnersPagePaginationJSON contains the JSON metadata for the struct
// [RunnersPagePagination]
type runnersPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnersPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnersPagePaginationJSON) RawJSON() string {
	return r.raw
}

type RunnersPage[T any] struct {
	Pagination RunnersPagePagination `json:"pagination"`
	Runners    []T                   `json:"runners"`
	JSON       runnersPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// runnersPageJSON contains the JSON metadata for the struct [RunnersPage[T]]
type runnersPageJSON struct {
	Pagination  apijson.Field
	Runners     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunnersPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runnersPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *RunnersPage[T]) GetNextPage() (res *RunnersPage[T], err error) {
	if len(r.Runners) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *RunnersPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &RunnersPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type RunnersPageAutoPager[T any] struct {
	page *RunnersPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewRunnersPageAutoPager[T any](page *RunnersPage[T], err error) *RunnersPageAutoPager[T] {
	return &RunnersPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *RunnersPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Runners) == 0 {
		return false
	}
	if r.idx >= len(r.page.Runners) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Runners) == 0 {
			return false
		}
	}
	r.cur = r.page.Runners[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *RunnersPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *RunnersPageAutoPager[T]) Err() error {
	return r.err
}

func (r *RunnersPageAutoPager[T]) Index() int {
	return r.run
}

type SecretsPagePagination struct {
	NextToken string                    `json:"nextToken"`
	JSON      secretsPagePaginationJSON `json:"-"`
}

// secretsPagePaginationJSON contains the JSON metadata for the struct
// [SecretsPagePagination]
type secretsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type SecretsPage[T any] struct {
	Pagination SecretsPagePagination `json:"pagination"`
	Secrets    []T                   `json:"secrets"`
	JSON       secretsPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// secretsPageJSON contains the JSON metadata for the struct [SecretsPage[T]]
type secretsPageJSON struct {
	Pagination  apijson.Field
	Secrets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SecretsPage[T]) GetNextPage() (res *SecretsPage[T], err error) {
	if len(r.Secrets) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SecretsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SecretsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SecretsPageAutoPager[T any] struct {
	page *SecretsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewSecretsPageAutoPager[T any](page *SecretsPage[T], err error) *SecretsPageAutoPager[T] {
	return &SecretsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SecretsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Secrets) == 0 {
		return false
	}
	if r.idx >= len(r.page.Secrets) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Secrets) == 0 {
			return false
		}
	}
	r.cur = r.page.Secrets[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SecretsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SecretsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *SecretsPageAutoPager[T]) Index() int {
	return r.run
}

type ServicesPagePagination struct {
	NextToken string                     `json:"nextToken"`
	JSON      servicesPagePaginationJSON `json:"-"`
}

// servicesPagePaginationJSON contains the JSON metadata for the struct
// [ServicesPagePagination]
type servicesPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServicesPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r servicesPagePaginationJSON) RawJSON() string {
	return r.raw
}

type ServicesPage[T any] struct {
	Pagination ServicesPagePagination `json:"pagination"`
	Services   []T                    `json:"services"`
	JSON       servicesPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// servicesPageJSON contains the JSON metadata for the struct [ServicesPage[T]]
type servicesPageJSON struct {
	Pagination  apijson.Field
	Services    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServicesPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r servicesPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ServicesPage[T]) GetNextPage() (res *ServicesPage[T], err error) {
	if len(r.Services) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ServicesPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ServicesPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ServicesPageAutoPager[T any] struct {
	page *ServicesPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewServicesPageAutoPager[T any](page *ServicesPage[T], err error) *ServicesPageAutoPager[T] {
	return &ServicesPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ServicesPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Services) == 0 {
		return false
	}
	if r.idx >= len(r.page.Services) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Services) == 0 {
			return false
		}
	}
	r.cur = r.page.Services[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ServicesPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ServicesPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ServicesPageAutoPager[T]) Index() int {
	return r.run
}

type SSOConfigurationsPagePagination struct {
	NextToken string                              `json:"nextToken"`
	JSON      ssoConfigurationsPagePaginationJSON `json:"-"`
}

// ssoConfigurationsPagePaginationJSON contains the JSON metadata for the struct
// [SSOConfigurationsPagePagination]
type ssoConfigurationsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOConfigurationsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoConfigurationsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type SSOConfigurationsPage[T any] struct {
	Pagination        SSOConfigurationsPagePagination `json:"pagination"`
	SSOConfigurations []T                             `json:"ssoConfigurations"`
	JSON              ssoConfigurationsPageJSON       `json:"-"`
	cfg               *requestconfig.RequestConfig
	res               *http.Response
}

// ssoConfigurationsPageJSON contains the JSON metadata for the struct
// [SSOConfigurationsPage[T]]
type ssoConfigurationsPageJSON struct {
	Pagination        apijson.Field
	SSOConfigurations apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SSOConfigurationsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoConfigurationsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SSOConfigurationsPage[T]) GetNextPage() (res *SSOConfigurationsPage[T], err error) {
	if len(r.SSOConfigurations) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SSOConfigurationsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SSOConfigurationsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SSOConfigurationsPageAutoPager[T any] struct {
	page *SSOConfigurationsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewSSOConfigurationsPageAutoPager[T any](page *SSOConfigurationsPage[T], err error) *SSOConfigurationsPageAutoPager[T] {
	return &SSOConfigurationsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SSOConfigurationsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.SSOConfigurations) == 0 {
		return false
	}
	if r.idx >= len(r.page.SSOConfigurations) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.SSOConfigurations) == 0 {
			return false
		}
	}
	r.cur = r.page.SSOConfigurations[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SSOConfigurationsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SSOConfigurationsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *SSOConfigurationsPageAutoPager[T]) Index() int {
	return r.run
}

type TaskExecutionsPagePagination struct {
	NextToken string                           `json:"nextToken"`
	JSON      taskExecutionsPagePaginationJSON `json:"-"`
}

// taskExecutionsPagePaginationJSON contains the JSON metadata for the struct
// [TaskExecutionsPagePagination]
type taskExecutionsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskExecutionsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type TaskExecutionsPage[T any] struct {
	Pagination     TaskExecutionsPagePagination `json:"pagination"`
	TaskExecutions []T                          `json:"taskExecutions"`
	JSON           taskExecutionsPageJSON       `json:"-"`
	cfg            *requestconfig.RequestConfig
	res            *http.Response
}

// taskExecutionsPageJSON contains the JSON metadata for the struct
// [TaskExecutionsPage[T]]
type taskExecutionsPageJSON struct {
	Pagination     apijson.Field
	TaskExecutions apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TaskExecutionsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskExecutionsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TaskExecutionsPage[T]) GetNextPage() (res *TaskExecutionsPage[T], err error) {
	if len(r.TaskExecutions) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *TaskExecutionsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TaskExecutionsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TaskExecutionsPageAutoPager[T any] struct {
	page *TaskExecutionsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewTaskExecutionsPageAutoPager[T any](page *TaskExecutionsPage[T], err error) *TaskExecutionsPageAutoPager[T] {
	return &TaskExecutionsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TaskExecutionsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.TaskExecutions) == 0 {
		return false
	}
	if r.idx >= len(r.page.TaskExecutions) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.TaskExecutions) == 0 {
			return false
		}
	}
	r.cur = r.page.TaskExecutions[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TaskExecutionsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *TaskExecutionsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *TaskExecutionsPageAutoPager[T]) Index() int {
	return r.run
}

type TasksPagePagination struct {
	NextToken string                  `json:"nextToken"`
	JSON      tasksPagePaginationJSON `json:"-"`
}

// tasksPagePaginationJSON contains the JSON metadata for the struct
// [TasksPagePagination]
type tasksPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TasksPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tasksPagePaginationJSON) RawJSON() string {
	return r.raw
}

type TasksPage[T any] struct {
	Pagination TasksPagePagination `json:"pagination"`
	Tasks      []T                 `json:"tasks"`
	JSON       tasksPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// tasksPageJSON contains the JSON metadata for the struct [TasksPage[T]]
type tasksPageJSON struct {
	Pagination  apijson.Field
	Tasks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TasksPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tasksPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TasksPage[T]) GetNextPage() (res *TasksPage[T], err error) {
	if len(r.Tasks) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *TasksPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TasksPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TasksPageAutoPager[T any] struct {
	page *TasksPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewTasksPageAutoPager[T any](page *TasksPage[T], err error) *TasksPageAutoPager[T] {
	return &TasksPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TasksPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Tasks) == 0 {
		return false
	}
	if r.idx >= len(r.page.Tasks) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Tasks) == 0 {
			return false
		}
	}
	r.cur = r.page.Tasks[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TasksPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *TasksPageAutoPager[T]) Err() error {
	return r.err
}

func (r *TasksPageAutoPager[T]) Index() int {
	return r.run
}

type TokensPagePagination struct {
	NextToken string                   `json:"nextToken"`
	JSON      tokensPagePaginationJSON `json:"-"`
}

// tokensPagePaginationJSON contains the JSON metadata for the struct
// [TokensPagePagination]
type tokensPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokensPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokensPagePaginationJSON) RawJSON() string {
	return r.raw
}

type TokensPage[T any] struct {
	Pagination TokensPagePagination `json:"pagination"`
	Tokens     []T                  `json:"tokens"`
	JSON       tokensPageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// tokensPageJSON contains the JSON metadata for the struct [TokensPage[T]]
type tokensPageJSON struct {
	Pagination  apijson.Field
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokensPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokensPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TokensPage[T]) GetNextPage() (res *TokensPage[T], err error) {
	if len(r.Tokens) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *TokensPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TokensPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TokensPageAutoPager[T any] struct {
	page *TokensPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewTokensPageAutoPager[T any](page *TokensPage[T], err error) *TokensPageAutoPager[T] {
	return &TokensPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TokensPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Tokens) == 0 {
		return false
	}
	if r.idx >= len(r.page.Tokens) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Tokens) == 0 {
			return false
		}
	}
	r.cur = r.page.Tokens[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TokensPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *TokensPageAutoPager[T]) Err() error {
	return r.err
}

func (r *TokensPageAutoPager[T]) Index() int {
	return r.run
}
