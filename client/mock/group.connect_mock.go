// Code generated by MockGen. DO NOT EDIT.
// Source: ../../v1/v1connect/group.connect.go
//
// Generated by this command:
//
//	mockgen -source=../../v1/v1connect/group.connect.go -destination=group.connect_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	connect "connectrpc.com/connect"
	v1 "github.com/gitpod-io/flex-api-go/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockGroupServiceClient is a mock of GroupServiceClient interface.
type MockGroupServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockGroupServiceClientMockRecorder
	isgomock struct{}
}

// MockGroupServiceClientMockRecorder is the mock recorder for MockGroupServiceClient.
type MockGroupServiceClientMockRecorder struct {
	mock *MockGroupServiceClient
}

// NewMockGroupServiceClient creates a new mock instance.
func NewMockGroupServiceClient(ctrl *gomock.Controller) *MockGroupServiceClient {
	mock := &MockGroupServiceClient{ctrl: ctrl}
	mock.recorder = &MockGroupServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupServiceClient) EXPECT() *MockGroupServiceClientMockRecorder {
	return m.recorder
}

// ListGroups mocks base method.
func (m *MockGroupServiceClient) ListGroups(arg0 context.Context, arg1 *connect.Request[v1.ListGroupsRequest]) (*connect.Response[v1.ListGroupsResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroups", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListGroupsResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroups indicates an expected call of ListGroups.
func (mr *MockGroupServiceClientMockRecorder) ListGroups(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockGroupServiceClient)(nil).ListGroups), arg0, arg1)
}

// MockGroupServiceHandler is a mock of GroupServiceHandler interface.
type MockGroupServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockGroupServiceHandlerMockRecorder
	isgomock struct{}
}

// MockGroupServiceHandlerMockRecorder is the mock recorder for MockGroupServiceHandler.
type MockGroupServiceHandlerMockRecorder struct {
	mock *MockGroupServiceHandler
}

// NewMockGroupServiceHandler creates a new mock instance.
func NewMockGroupServiceHandler(ctrl *gomock.Controller) *MockGroupServiceHandler {
	mock := &MockGroupServiceHandler{ctrl: ctrl}
	mock.recorder = &MockGroupServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupServiceHandler) EXPECT() *MockGroupServiceHandlerMockRecorder {
	return m.recorder
}

// ListGroups mocks base method.
func (m *MockGroupServiceHandler) ListGroups(arg0 context.Context, arg1 *connect.Request[v1.ListGroupsRequest]) (*connect.Response[v1.ListGroupsResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroups", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[v1.ListGroupsResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroups indicates an expected call of ListGroups.
func (mr *MockGroupServiceHandlerMockRecorder) ListGroups(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockGroupServiceHandler)(nil).ListGroups), arg0, arg1)
}
