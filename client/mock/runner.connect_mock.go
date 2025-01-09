// Code generated by MockGen. DO NOT EDIT.
// Source: ../../v1/v1connect/runner.connect.go
//
// Generated by this command:
//
//	mockgen -source=../../v1/v1connect/runner.connect.go -destination=runner.connect_mock.go -package=mock
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

// MockRunnerServiceClient is a mock of RunnerServiceClient interface.
type MockRunnerServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerServiceClientMockRecorder
	isgomock struct{}
}

// MockRunnerServiceClientMockRecorder is the mock recorder for MockRunnerServiceClient.
type MockRunnerServiceClientMockRecorder struct {
	mock *MockRunnerServiceClient
}

// NewMockRunnerServiceClient creates a new mock instance.
func NewMockRunnerServiceClient(ctrl *gomock.Controller) *MockRunnerServiceClient {
	mock := &MockRunnerServiceClient{ctrl: ctrl}
	mock.recorder = &MockRunnerServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunnerServiceClient) EXPECT() *MockRunnerServiceClientMockRecorder {
	return m.recorder
}

// CheckAuthenticationForHost mocks base method.
func (m *MockRunnerServiceClient) CheckAuthenticationForHost(arg0 context.Context, arg1 *connect.Request[v1.CheckAuthenticationForHostRequest]) (*connect.Response[v1.CheckAuthenticationForHostResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAuthenticationForHost", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CheckAuthenticationForHostResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAuthenticationForHost indicates an expected call of CheckAuthenticationForHost.
func (mr *MockRunnerServiceClientMockRecorder) CheckAuthenticationForHost(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAuthenticationForHost", reflect.TypeOf((*MockRunnerServiceClient)(nil).CheckAuthenticationForHost), arg0, arg1)
}

// CreateRunner mocks base method.
func (m *MockRunnerServiceClient) CreateRunner(arg0 context.Context, arg1 *connect.Request[v1.CreateRunnerRequest]) (*connect.Response[v1.CreateRunnerResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRunner", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateRunnerResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRunner indicates an expected call of CreateRunner.
func (mr *MockRunnerServiceClientMockRecorder) CreateRunner(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRunner", reflect.TypeOf((*MockRunnerServiceClient)(nil).CreateRunner), arg0, arg1)
}

// CreateRunnerPolicy mocks base method.
func (m *MockRunnerServiceClient) CreateRunnerPolicy(arg0 context.Context, arg1 *connect.Request[v1.CreateRunnerPolicyRequest]) (*connect.Response[v1.CreateRunnerPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRunnerPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateRunnerPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRunnerPolicy indicates an expected call of CreateRunnerPolicy.
func (mr *MockRunnerServiceClientMockRecorder) CreateRunnerPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRunnerPolicy", reflect.TypeOf((*MockRunnerServiceClient)(nil).CreateRunnerPolicy), arg0, arg1)
}

// CreateRunnerToken mocks base method.
func (m *MockRunnerServiceClient) CreateRunnerToken(arg0 context.Context, arg1 *connect.Request[v1.CreateRunnerTokenRequest]) (*connect.Response[v1.CreateRunnerTokenResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRunnerToken", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateRunnerTokenResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRunnerToken indicates an expected call of CreateRunnerToken.
func (mr *MockRunnerServiceClientMockRecorder) CreateRunnerToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRunnerToken", reflect.TypeOf((*MockRunnerServiceClient)(nil).CreateRunnerToken), arg0, arg1)
}

// DeleteRunner mocks base method.
func (m *MockRunnerServiceClient) DeleteRunner(arg0 context.Context, arg1 *connect.Request[v1.DeleteRunnerRequest]) (*connect.Response[v1.DeleteRunnerResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRunner", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.DeleteRunnerResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRunner indicates an expected call of DeleteRunner.
func (mr *MockRunnerServiceClientMockRecorder) DeleteRunner(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRunner", reflect.TypeOf((*MockRunnerServiceClient)(nil).DeleteRunner), arg0, arg1)
}

// DeleteRunnerPolicy mocks base method.
func (m *MockRunnerServiceClient) DeleteRunnerPolicy(arg0 context.Context, arg1 *connect.Request[v1.DeleteRunnerPolicyRequest]) (*connect.Response[v1.DeleteRunnerPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRunnerPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.DeleteRunnerPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRunnerPolicy indicates an expected call of DeleteRunnerPolicy.
func (mr *MockRunnerServiceClientMockRecorder) DeleteRunnerPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRunnerPolicy", reflect.TypeOf((*MockRunnerServiceClient)(nil).DeleteRunnerPolicy), arg0, arg1)
}

// GetRunner mocks base method.
func (m *MockRunnerServiceClient) GetRunner(arg0 context.Context, arg1 *connect.Request[v1.GetRunnerRequest]) (*connect.Response[v1.GetRunnerResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunner", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.GetRunnerResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunner indicates an expected call of GetRunner.
func (mr *MockRunnerServiceClientMockRecorder) GetRunner(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunner", reflect.TypeOf((*MockRunnerServiceClient)(nil).GetRunner), arg0, arg1)
}

// ListRunnerPolicies mocks base method.
func (m *MockRunnerServiceClient) ListRunnerPolicies(arg0 context.Context, arg1 *connect.Request[v1.ListRunnerPoliciesRequest]) (*connect.Response[v1.ListRunnerPoliciesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRunnerPolicies", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListRunnerPoliciesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRunnerPolicies indicates an expected call of ListRunnerPolicies.
func (mr *MockRunnerServiceClientMockRecorder) ListRunnerPolicies(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRunnerPolicies", reflect.TypeOf((*MockRunnerServiceClient)(nil).ListRunnerPolicies), arg0, arg1)
}

// ListRunners mocks base method.
func (m *MockRunnerServiceClient) ListRunners(arg0 context.Context, arg1 *connect.Request[v1.ListRunnersRequest]) (*connect.Response[v1.ListRunnersResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRunners", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListRunnersResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRunners indicates an expected call of ListRunners.
func (mr *MockRunnerServiceClientMockRecorder) ListRunners(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRunners", reflect.TypeOf((*MockRunnerServiceClient)(nil).ListRunners), arg0, arg1)
}

// ParseContextURL mocks base method.
func (m *MockRunnerServiceClient) ParseContextURL(arg0 context.Context, arg1 *connect.Request[v1.ParseContextURLRequest]) (*connect.Response[v1.ParseContextURLResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseContextURL", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ParseContextURLResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseContextURL indicates an expected call of ParseContextURL.
func (mr *MockRunnerServiceClientMockRecorder) ParseContextURL(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseContextURL", reflect.TypeOf((*MockRunnerServiceClient)(nil).ParseContextURL), arg0, arg1)
}

// UpdateRunner mocks base method.
func (m *MockRunnerServiceClient) UpdateRunner(arg0 context.Context, arg1 *connect.Request[v1.UpdateRunnerRequest]) (*connect.Response[v1.UpdateRunnerResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRunner", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.UpdateRunnerResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRunner indicates an expected call of UpdateRunner.
func (mr *MockRunnerServiceClientMockRecorder) UpdateRunner(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRunner", reflect.TypeOf((*MockRunnerServiceClient)(nil).UpdateRunner), arg0, arg1)
}

// UpdateRunnerPolicy mocks base method.
func (m *MockRunnerServiceClient) UpdateRunnerPolicy(arg0 context.Context, arg1 *connect.Request[v1.UpdateRunnerPolicyRequest]) (*connect.Response[v1.UpdateRunnerPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRunnerPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.UpdateRunnerPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRunnerPolicy indicates an expected call of UpdateRunnerPolicy.
func (mr *MockRunnerServiceClientMockRecorder) UpdateRunnerPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRunnerPolicy", reflect.TypeOf((*MockRunnerServiceClient)(nil).UpdateRunnerPolicy), arg0, arg1)
}

// MockRunnerServiceHandler is a mock of RunnerServiceHandler interface.
type MockRunnerServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerServiceHandlerMockRecorder
	isgomock struct{}
}

// MockRunnerServiceHandlerMockRecorder is the mock recorder for MockRunnerServiceHandler.
type MockRunnerServiceHandlerMockRecorder struct {
	mock *MockRunnerServiceHandler
}

// NewMockRunnerServiceHandler creates a new mock instance.
func NewMockRunnerServiceHandler(ctrl *gomock.Controller) *MockRunnerServiceHandler {
	mock := &MockRunnerServiceHandler{ctrl: ctrl}
	mock.recorder = &MockRunnerServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunnerServiceHandler) EXPECT() *MockRunnerServiceHandlerMockRecorder {
	return m.recorder
}

// CheckAuthenticationForHost mocks base method.
func (m *MockRunnerServiceHandler) CheckAuthenticationForHost(arg0 context.Context, arg1 *connect.Request[v1.CheckAuthenticationForHostRequest]) (*connect.Response[v1.CheckAuthenticationForHostResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAuthenticationForHost", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CheckAuthenticationForHostResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAuthenticationForHost indicates an expected call of CheckAuthenticationForHost.
func (mr *MockRunnerServiceHandlerMockRecorder) CheckAuthenticationForHost(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAuthenticationForHost", reflect.TypeOf((*MockRunnerServiceHandler)(nil).CheckAuthenticationForHost), arg0, arg1)
}

// CreateRunner mocks base method.
func (m *MockRunnerServiceHandler) CreateRunner(arg0 context.Context, arg1 *connect.Request[v1.CreateRunnerRequest]) (*connect.Response[v1.CreateRunnerResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRunner", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateRunnerResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRunner indicates an expected call of CreateRunner.
func (mr *MockRunnerServiceHandlerMockRecorder) CreateRunner(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRunner", reflect.TypeOf((*MockRunnerServiceHandler)(nil).CreateRunner), arg0, arg1)
}

// CreateRunnerPolicy mocks base method.
func (m *MockRunnerServiceHandler) CreateRunnerPolicy(arg0 context.Context, arg1 *connect.Request[v1.CreateRunnerPolicyRequest]) (*connect.Response[v1.CreateRunnerPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRunnerPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateRunnerPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRunnerPolicy indicates an expected call of CreateRunnerPolicy.
func (mr *MockRunnerServiceHandlerMockRecorder) CreateRunnerPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRunnerPolicy", reflect.TypeOf((*MockRunnerServiceHandler)(nil).CreateRunnerPolicy), arg0, arg1)
}

// CreateRunnerToken mocks base method.
func (m *MockRunnerServiceHandler) CreateRunnerToken(arg0 context.Context, arg1 *connect.Request[v1.CreateRunnerTokenRequest]) (*connect.Response[v1.CreateRunnerTokenResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRunnerToken", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.CreateRunnerTokenResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRunnerToken indicates an expected call of CreateRunnerToken.
func (mr *MockRunnerServiceHandlerMockRecorder) CreateRunnerToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRunnerToken", reflect.TypeOf((*MockRunnerServiceHandler)(nil).CreateRunnerToken), arg0, arg1)
}

// DeleteRunner mocks base method.
func (m *MockRunnerServiceHandler) DeleteRunner(arg0 context.Context, arg1 *connect.Request[v1.DeleteRunnerRequest]) (*connect.Response[v1.DeleteRunnerResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRunner", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.DeleteRunnerResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRunner indicates an expected call of DeleteRunner.
func (mr *MockRunnerServiceHandlerMockRecorder) DeleteRunner(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRunner", reflect.TypeOf((*MockRunnerServiceHandler)(nil).DeleteRunner), arg0, arg1)
}

// DeleteRunnerPolicy mocks base method.
func (m *MockRunnerServiceHandler) DeleteRunnerPolicy(arg0 context.Context, arg1 *connect.Request[v1.DeleteRunnerPolicyRequest]) (*connect.Response[v1.DeleteRunnerPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRunnerPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.DeleteRunnerPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRunnerPolicy indicates an expected call of DeleteRunnerPolicy.
func (mr *MockRunnerServiceHandlerMockRecorder) DeleteRunnerPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRunnerPolicy", reflect.TypeOf((*MockRunnerServiceHandler)(nil).DeleteRunnerPolicy), arg0, arg1)
}

// GetRunner mocks base method.
func (m *MockRunnerServiceHandler) GetRunner(arg0 context.Context, arg1 *connect.Request[v1.GetRunnerRequest]) (*connect.Response[v1.GetRunnerResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunner", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.GetRunnerResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunner indicates an expected call of GetRunner.
func (mr *MockRunnerServiceHandlerMockRecorder) GetRunner(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunner", reflect.TypeOf((*MockRunnerServiceHandler)(nil).GetRunner), arg0, arg1)
}

// ListRunnerPolicies mocks base method.
func (m *MockRunnerServiceHandler) ListRunnerPolicies(arg0 context.Context, arg1 *connect.Request[v1.ListRunnerPoliciesRequest]) (*connect.Response[v1.ListRunnerPoliciesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRunnerPolicies", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListRunnerPoliciesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRunnerPolicies indicates an expected call of ListRunnerPolicies.
func (mr *MockRunnerServiceHandlerMockRecorder) ListRunnerPolicies(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRunnerPolicies", reflect.TypeOf((*MockRunnerServiceHandler)(nil).ListRunnerPolicies), arg0, arg1)
}

// ListRunners mocks base method.
func (m *MockRunnerServiceHandler) ListRunners(arg0 context.Context, arg1 *connect.Request[v1.ListRunnersRequest]) (*connect.Response[v1.ListRunnersResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRunners", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListRunnersResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRunners indicates an expected call of ListRunners.
func (mr *MockRunnerServiceHandlerMockRecorder) ListRunners(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRunners", reflect.TypeOf((*MockRunnerServiceHandler)(nil).ListRunners), arg0, arg1)
}

// ParseContextURL mocks base method.
func (m *MockRunnerServiceHandler) ParseContextURL(arg0 context.Context, arg1 *connect.Request[v1.ParseContextURLRequest]) (*connect.Response[v1.ParseContextURLResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseContextURL", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ParseContextURLResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseContextURL indicates an expected call of ParseContextURL.
func (mr *MockRunnerServiceHandlerMockRecorder) ParseContextURL(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseContextURL", reflect.TypeOf((*MockRunnerServiceHandler)(nil).ParseContextURL), arg0, arg1)
}

// UpdateRunner mocks base method.
func (m *MockRunnerServiceHandler) UpdateRunner(arg0 context.Context, arg1 *connect.Request[v1.UpdateRunnerRequest]) (*connect.Response[v1.UpdateRunnerResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRunner", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.UpdateRunnerResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRunner indicates an expected call of UpdateRunner.
func (mr *MockRunnerServiceHandlerMockRecorder) UpdateRunner(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRunner", reflect.TypeOf((*MockRunnerServiceHandler)(nil).UpdateRunner), arg0, arg1)
}

// UpdateRunnerPolicy mocks base method.
func (m *MockRunnerServiceHandler) UpdateRunnerPolicy(arg0 context.Context, arg1 *connect.Request[v1.UpdateRunnerPolicyRequest]) (*connect.Response[v1.UpdateRunnerPolicyResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRunnerPolicy", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.UpdateRunnerPolicyResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRunnerPolicy indicates an expected call of UpdateRunnerPolicy.
func (mr *MockRunnerServiceHandlerMockRecorder) UpdateRunnerPolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRunnerPolicy", reflect.TypeOf((*MockRunnerServiceHandler)(nil).UpdateRunnerPolicy), arg0, arg1)
}
