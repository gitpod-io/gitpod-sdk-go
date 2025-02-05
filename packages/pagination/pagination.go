// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	EnvironmentClasses []T                              `json:"environment_classes"`
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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

type EnvironmentClassesPage[T any] struct {
	EnvironmentClasses []T                              `json:"environment_classes"`
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	LoginProviders []T                          `json:"login_providers"`
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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

type OrganizationsPagePagination struct {
	NextToken string                          `json:"nextToken"`
	JSON      organizationsPagePaginationJSON `json:"-"`
}

// organizationsPagePaginationJSON contains the JSON metadata for the struct
// [OrganizationsPagePagination]
type organizationsPagePaginationJSON struct {
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationsPagePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationsPagePaginationJSON) RawJSON() string {
	return r.raw
}

type OrganizationsPage[T any] struct {
	Organizations []T                         `json:"organizations"`
	Pagination    OrganizationsPagePagination `json:"pagination"`
	JSON          organizationsPageJSON       `json:"-"`
	cfg           *requestconfig.RequestConfig
	res           *http.Response
}

// organizationsPageJSON contains the JSON metadata for the struct
// [OrganizationsPage[T]]
type organizationsPageJSON struct {
	Organizations apijson.Field
	Pagination    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OrganizationsPage[T]) GetNextPage() (res *OrganizationsPage[T], err error) {
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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

func (r *OrganizationsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OrganizationsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OrganizationsPageAutoPager[T any] struct {
	page *OrganizationsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewOrganizationsPageAutoPager[T any](page *OrganizationsPage[T], err error) *OrganizationsPageAutoPager[T] {
	return &OrganizationsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OrganizationsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Organizations) == 0 {
		return false
	}
	if r.idx >= len(r.page.Organizations) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Organizations) == 0 {
			return false
		}
	}
	r.cur = r.page.Organizations[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OrganizationsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *OrganizationsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *OrganizationsPageAutoPager[T]) Index() int {
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
	PersonalAccessTokens []T                                `json:"personal_access_tokens"`
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	SSOConfigurations []T                             `json:"sso_configurations"`
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	TaskExecutions []T                          `json:"task_executions"`
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
	next := r.Pagination.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("token", next))
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
