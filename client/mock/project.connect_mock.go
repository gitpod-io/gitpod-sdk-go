// Code generated by MockGen. DO NOT EDIT.
// Source: ../../v1/v1connect/project.connect.go
//
// Generated by this command:
//
//	mockgen -source=../../v1/v1connect/project.connect.go -destination=project.connect_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	connect "connectrpc.com/connect"
	v1 "github.com/gitpod-io/flex-go/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectServiceClient is a mock of ProjectServiceClient interface.
type MockProjectServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProjectServiceClientMockRecorder
	isgomock struct{}
}

// MockProjectServiceClientMockRecorder is the mock recorder for MockProjectServiceClient.
type MockProjectServiceClientMockRecorder struct {
	mock *MockProjectServiceClient
}

// NewMockProjectServiceClient creates a new mock instance.
func NewMockProjectServiceClient(ctrl *gomock.Controller) *MockProjectServiceClient {
	mock := &MockProjectServiceClient{ctrl: ctrl}
	mock.recorder = &MockProjectServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectServiceClient) EXPECT() *MockProjectServiceClientMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockProjectServiceClient) CreateProject(arg0 context.Context, arg1 *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.CreateProjectResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateProjectResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockProjectServiceClientMockRecorder) CreateProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectServiceClient)(nil).CreateProject), arg0, arg1)
}

// CreateProjectFromEnvironment mocks base method.
func (m *MockProjectServiceClient) CreateProjectFromEnvironment(arg0 context.Context, arg1 *connect.Request[v1.CreateProjectFromEnvironmentRequest]) (*connect.Response[v1.CreateProjectFromEnvironmentResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectFromEnvironment", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateProjectFromEnvironmentResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectFromEnvironment indicates an expected call of CreateProjectFromEnvironment.
func (mr *MockProjectServiceClientMockRecorder) CreateProjectFromEnvironment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectFromEnvironment", reflect.TypeOf((*MockProjectServiceClient)(nil).CreateProjectFromEnvironment), arg0, arg1)
}

// CreateProjectPolicy mocks base method.
func (m *MockProjectServiceClient) CreateProjectPolicy(arg0 context.Context, arg1 *connect.Request[v1.CreateProjectPolicyRequest]) (*connect.Response[v1.CreateProjectPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateProjectPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectPolicy indicates an expected call of CreateProjectPolicy.
func (mr *MockProjectServiceClientMockRecorder) CreateProjectPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectPolicy", reflect.TypeOf((*MockProjectServiceClient)(nil).CreateProjectPolicy), arg0, arg1)
}

// DeleteProject mocks base method.
func (m *MockProjectServiceClient) DeleteProject(arg0 context.Context, arg1 *connect.Request[v1.DeleteProjectRequest]) (*connect.Response[v1.DeleteProjectResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.DeleteProjectResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockProjectServiceClientMockRecorder) DeleteProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockProjectServiceClient)(nil).DeleteProject), arg0, arg1)
}

// DeleteProjectPolicy mocks base method.
func (m *MockProjectServiceClient) DeleteProjectPolicy(arg0 context.Context, arg1 *connect.Request[v1.DeleteProjectPolicyRequest]) (*connect.Response[v1.DeleteProjectPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProjectPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.DeleteProjectPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProjectPolicy indicates an expected call of DeleteProjectPolicy.
func (mr *MockProjectServiceClientMockRecorder) DeleteProjectPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectPolicy", reflect.TypeOf((*MockProjectServiceClient)(nil).DeleteProjectPolicy), arg0, arg1)
}

// GetProject mocks base method.
func (m *MockProjectServiceClient) GetProject(arg0 context.Context, arg1 *connect.Request[v1.GetProjectRequest]) (*connect.Response[v1.GetProjectResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.GetProjectResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockProjectServiceClientMockRecorder) GetProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockProjectServiceClient)(nil).GetProject), arg0, arg1)
}

// ListProjectPolicies mocks base method.
func (m *MockProjectServiceClient) ListProjectPolicies(arg0 context.Context, arg1 *connect.Request[v1.ListProjectPoliciesRequest]) (*connect.Response[v1.ListProjectPoliciesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectPolicies", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListProjectPoliciesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectPolicies indicates an expected call of ListProjectPolicies.
func (mr *MockProjectServiceClientMockRecorder) ListProjectPolicies(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectPolicies", reflect.TypeOf((*MockProjectServiceClient)(nil).ListProjectPolicies), arg0, arg1)
}

// ListProjects mocks base method.
func (m *MockProjectServiceClient) ListProjects(arg0 context.Context, arg1 *connect.Request[v1.ListProjectsRequest]) (*connect.Response[v1.ListProjectsResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListProjectsResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockProjectServiceClientMockRecorder) ListProjects(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockProjectServiceClient)(nil).ListProjects), arg0, arg1)
}

// UpdateProject mocks base method.
func (m *MockProjectServiceClient) UpdateProject(arg0 context.Context, arg1 *connect.Request[v1.UpdateProjectRequest]) (*connect.Response[v1.UpdateProjectResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.UpdateProjectResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockProjectServiceClientMockRecorder) UpdateProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockProjectServiceClient)(nil).UpdateProject), arg0, arg1)
}

// UpdateProjectPolicy mocks base method.
func (m *MockProjectServiceClient) UpdateProjectPolicy(arg0 context.Context, arg1 *connect.Request[v1.UpdateProjectPolicyRequest]) (*connect.Response[v1.UpdateProjectPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.UpdateProjectPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProjectPolicy indicates an expected call of UpdateProjectPolicy.
func (mr *MockProjectServiceClientMockRecorder) UpdateProjectPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectPolicy", reflect.TypeOf((*MockProjectServiceClient)(nil).UpdateProjectPolicy), arg0, arg1)
}

// MockProjectServiceHandler is a mock of ProjectServiceHandler interface.
type MockProjectServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockProjectServiceHandlerMockRecorder
	isgomock struct{}
}

// MockProjectServiceHandlerMockRecorder is the mock recorder for MockProjectServiceHandler.
type MockProjectServiceHandlerMockRecorder struct {
	mock *MockProjectServiceHandler
}

// NewMockProjectServiceHandler creates a new mock instance.
func NewMockProjectServiceHandler(ctrl *gomock.Controller) *MockProjectServiceHandler {
	mock := &MockProjectServiceHandler{ctrl: ctrl}
	mock.recorder = &MockProjectServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectServiceHandler) EXPECT() *MockProjectServiceHandlerMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockProjectServiceHandler) CreateProject(arg0 context.Context, arg1 *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.CreateProjectResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateProjectResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockProjectServiceHandlerMockRecorder) CreateProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectServiceHandler)(nil).CreateProject), arg0, arg1)
}

// CreateProjectFromEnvironment mocks base method.
func (m *MockProjectServiceHandler) CreateProjectFromEnvironment(arg0 context.Context, arg1 *connect.Request[v1.CreateProjectFromEnvironmentRequest]) (*connect.Response[v1.CreateProjectFromEnvironmentResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectFromEnvironment", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateProjectFromEnvironmentResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectFromEnvironment indicates an expected call of CreateProjectFromEnvironment.
func (mr *MockProjectServiceHandlerMockRecorder) CreateProjectFromEnvironment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectFromEnvironment", reflect.TypeOf((*MockProjectServiceHandler)(nil).CreateProjectFromEnvironment), arg0, arg1)
}

// CreateProjectPolicy mocks base method.
func (m *MockProjectServiceHandler) CreateProjectPolicy(arg0 context.Context, arg1 *connect.Request[v1.CreateProjectPolicyRequest]) (*connect.Response[v1.CreateProjectPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateProjectPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectPolicy indicates an expected call of CreateProjectPolicy.
func (mr *MockProjectServiceHandlerMockRecorder) CreateProjectPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectPolicy", reflect.TypeOf((*MockProjectServiceHandler)(nil).CreateProjectPolicy), arg0, arg1)
}

// DeleteProject mocks base method.
func (m *MockProjectServiceHandler) DeleteProject(arg0 context.Context, arg1 *connect.Request[v1.DeleteProjectRequest]) (*connect.Response[v1.DeleteProjectResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.DeleteProjectResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockProjectServiceHandlerMockRecorder) DeleteProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockProjectServiceHandler)(nil).DeleteProject), arg0, arg1)
}

// DeleteProjectPolicy mocks base method.
func (m *MockProjectServiceHandler) DeleteProjectPolicy(arg0 context.Context, arg1 *connect.Request[v1.DeleteProjectPolicyRequest]) (*connect.Response[v1.DeleteProjectPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProjectPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.DeleteProjectPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProjectPolicy indicates an expected call of DeleteProjectPolicy.
func (mr *MockProjectServiceHandlerMockRecorder) DeleteProjectPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectPolicy", reflect.TypeOf((*MockProjectServiceHandler)(nil).DeleteProjectPolicy), arg0, arg1)
}

// GetProject mocks base method.
func (m *MockProjectServiceHandler) GetProject(arg0 context.Context, arg1 *connect.Request[v1.GetProjectRequest]) (*connect.Response[v1.GetProjectResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.GetProjectResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockProjectServiceHandlerMockRecorder) GetProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockProjectServiceHandler)(nil).GetProject), arg0, arg1)
}

// ListProjectPolicies mocks base method.
func (m *MockProjectServiceHandler) ListProjectPolicies(arg0 context.Context, arg1 *connect.Request[v1.ListProjectPoliciesRequest]) (*connect.Response[v1.ListProjectPoliciesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectPolicies", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListProjectPoliciesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectPolicies indicates an expected call of ListProjectPolicies.
func (mr *MockProjectServiceHandlerMockRecorder) ListProjectPolicies(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectPolicies", reflect.TypeOf((*MockProjectServiceHandler)(nil).ListProjectPolicies), arg0, arg1)
}

// ListProjects mocks base method.
func (m *MockProjectServiceHandler) ListProjects(arg0 context.Context, arg1 *connect.Request[v1.ListProjectsRequest]) (*connect.Response[v1.ListProjectsResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListProjectsResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockProjectServiceHandlerMockRecorder) ListProjects(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockProjectServiceHandler)(nil).ListProjects), arg0, arg1)
}

// UpdateProject mocks base method.
func (m *MockProjectServiceHandler) UpdateProject(arg0 context.Context, arg1 *connect.Request[v1.UpdateProjectRequest]) (*connect.Response[v1.UpdateProjectResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.UpdateProjectResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockProjectServiceHandlerMockRecorder) UpdateProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockProjectServiceHandler)(nil).UpdateProject), arg0, arg1)
}

// UpdateProjectPolicy mocks base method.
func (m *MockProjectServiceHandler) UpdateProjectPolicy(arg0 context.Context, arg1 *connect.Request[v1.UpdateProjectPolicyRequest]) (*connect.Response[v1.UpdateProjectPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.UpdateProjectPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProjectPolicy indicates an expected call of UpdateProjectPolicy.
func (mr *MockProjectServiceHandlerMockRecorder) UpdateProjectPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectPolicy", reflect.TypeOf((*MockProjectServiceHandler)(nil).UpdateProjectPolicy), arg0, arg1)
}
