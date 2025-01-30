// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

type ServicesPage[T any] struct {
	Services  []T              `json:"services"`
	NextToken string           `json:"nextToken"`
	JSON      servicesPageJSON `json:"-"`
	cfg       *requestconfig.RequestConfig
	res       *http.Response
}

// servicesPageJSON contains the JSON metadata for the struct [ServicesPage[T]]
type servicesPageJSON struct {
	Services    apijson.Field
	NextToken   apijson.Field
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
	next := r.NextToken
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

type TasksPage[T any] struct {
	Tasks     []T           `json:"tasks"`
	NextToken string        `json:"nextToken"`
	JSON      tasksPageJSON `json:"-"`
	cfg       *requestconfig.RequestConfig
	res       *http.Response
}

// tasksPageJSON contains the JSON metadata for the struct [TasksPage[T]]
type tasksPageJSON struct {
	Tasks       apijson.Field
	NextToken   apijson.Field
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
	next := r.NextToken
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

type TaskExecutionsPage[T any] struct {
	TaskExecutions []T                    `json:"taskExecutions"`
	NextToken      string                 `json:"nextToken"`
	JSON           taskExecutionsPageJSON `json:"-"`
	cfg            *requestconfig.RequestConfig
	res            *http.Response
}

// taskExecutionsPageJSON contains the JSON metadata for the struct
// [TaskExecutionsPage[T]]
type taskExecutionsPageJSON struct {
	TaskExecutions apijson.Field
	NextToken      apijson.Field
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
	next := r.NextToken
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

type EnvironnmentClassesPage[T any] struct {
	EnvironmentClasses []T                         `json:"environmentClasses"`
	NextToken          string                      `json:"nextToken"`
	JSON               environnmentClassesPageJSON `json:"-"`
	cfg                *requestconfig.RequestConfig
	res                *http.Response
}

// environnmentClassesPageJSON contains the JSON metadata for the struct
// [EnvironnmentClassesPage[T]]
type environnmentClassesPageJSON struct {
	EnvironmentClasses apijson.Field
	NextToken          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironnmentClassesPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environnmentClassesPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *EnvironnmentClassesPage[T]) GetNextPage() (res *EnvironnmentClassesPage[T], err error) {
	next := r.NextToken
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

func (r *EnvironnmentClassesPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &EnvironnmentClassesPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type EnvironnmentClassesPageAutoPager[T any] struct {
	page *EnvironnmentClassesPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewEnvironnmentClassesPageAutoPager[T any](page *EnvironnmentClassesPage[T], err error) *EnvironnmentClassesPageAutoPager[T] {
	return &EnvironnmentClassesPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *EnvironnmentClassesPageAutoPager[T]) Next() bool {
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

func (r *EnvironnmentClassesPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *EnvironnmentClassesPageAutoPager[T]) Err() error {
	return r.err
}

func (r *EnvironnmentClassesPageAutoPager[T]) Index() int {
	return r.run
}

type EnvironmentsPage[T any] struct {
	Environments []T                  `json:"environments"`
	NextToken    string               `json:"nextToken"`
	JSON         environmentsPageJSON `json:"-"`
	cfg          *requestconfig.RequestConfig
	res          *http.Response
}

// environmentsPageJSON contains the JSON metadata for the struct
// [EnvironmentsPage[T]]
type environmentsPageJSON struct {
	Environments apijson.Field
	NextToken    apijson.Field
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
	next := r.NextToken
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

type EntriesPage[T any] struct {
	Entries   []T             `json:"entries"`
	NextToken string          `json:"nextToken"`
	JSON      entriesPageJSON `json:"-"`
	cfg       *requestconfig.RequestConfig
	res       *http.Response
}

// entriesPageJSON contains the JSON metadata for the struct [EntriesPage[T]]
type entriesPageJSON struct {
	Entries     apijson.Field
	NextToken   apijson.Field
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
	next := r.NextToken
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

type GroupsPage[T any] struct {
	Groups    []T            `json:"groups"`
	NextToken string         `json:"nextToken"`
	JSON      groupsPageJSON `json:"-"`
	cfg       *requestconfig.RequestConfig
	res       *http.Response
}

// groupsPageJSON contains the JSON metadata for the struct [GroupsPage[T]]
type groupsPageJSON struct {
	Groups      apijson.Field
	NextToken   apijson.Field
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
	next := r.NextToken
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

type MembersPage[T any] struct {
	Members   []T             `json:"members"`
	NextToken string          `json:"nextToken"`
	JSON      membersPageJSON `json:"-"`
	cfg       *requestconfig.RequestConfig
	res       *http.Response
}

// membersPageJSON contains the JSON metadata for the struct [MembersPage[T]]
type membersPageJSON struct {
	Members     apijson.Field
	NextToken   apijson.Field
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
	next := r.NextToken
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
