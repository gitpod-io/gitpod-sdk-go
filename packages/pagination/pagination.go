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
