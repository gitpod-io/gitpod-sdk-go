// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gitpod-io/flex-api-go/client/cache (interfaces: Invalidator,EventStream)
//
// Generated by this command:
//
//	mockgen -package cache_test -destination cache_mock_test.go github.com/gitpod-io/flex-api-go/client/cache Invalidator,EventStream
//

// Package cache_test is a generated GoMock package.
package cache_test

import (
	context "context"
	reflect "reflect"

	v1 "github.com/gitpod-io/flex-api-go/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockInvalidator is a mock of Invalidator interface.
type MockInvalidator struct {
	ctrl     *gomock.Controller
	recorder *MockInvalidatorMockRecorder
}

// MockInvalidatorMockRecorder is the mock recorder for MockInvalidator.
type MockInvalidatorMockRecorder struct {
	mock *MockInvalidator
}

// NewMockInvalidator creates a new mock instance.
func NewMockInvalidator(ctrl *gomock.Controller) *MockInvalidator {
	mock := &MockInvalidator{ctrl: ctrl}
	mock.recorder = &MockInvalidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvalidator) EXPECT() *MockInvalidatorMockRecorder {
	return m.recorder
}

// InvalidateItem mocks base method.
func (m *MockInvalidator) InvalidateItem(arg0 context.Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InvalidateItem", arg0, arg1)
}

// InvalidateItem indicates an expected call of InvalidateItem.
func (mr *MockInvalidatorMockRecorder) InvalidateItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateItem", reflect.TypeOf((*MockInvalidator)(nil).InvalidateItem), arg0, arg1)
}

// RemoveItem mocks base method.
func (m *MockInvalidator) RemoveItem(arg0 context.Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveItem", arg0, arg1)
}

// RemoveItem indicates an expected call of RemoveItem.
func (mr *MockInvalidatorMockRecorder) RemoveItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveItem", reflect.TypeOf((*MockInvalidator)(nil).RemoveItem), arg0, arg1)
}

// MockEventStream is a mock of EventStream interface.
type MockEventStream struct {
	ctrl     *gomock.Controller
	recorder *MockEventStreamMockRecorder
}

// MockEventStreamMockRecorder is the mock recorder for MockEventStream.
type MockEventStreamMockRecorder struct {
	mock *MockEventStream
}

// NewMockEventStream creates a new mock instance.
func NewMockEventStream(ctrl *gomock.Controller) *MockEventStream {
	mock := &MockEventStream{ctrl: ctrl}
	mock.recorder = &MockEventStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventStream) EXPECT() *MockEventStreamMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockEventStream) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockEventStreamMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEventStream)(nil).Close))
}

// Err mocks base method.
func (m *MockEventStream) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockEventStreamMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockEventStream)(nil).Err))
}

// Msg mocks base method.
func (m *MockEventStream) Msg() *v1.WatchEventsResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Msg")
	ret0, _ := ret[0].(*v1.WatchEventsResponse)
	return ret0
}

// Msg indicates an expected call of Msg.
func (mr *MockEventStreamMockRecorder) Msg() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Msg", reflect.TypeOf((*MockEventStream)(nil).Msg))
}

// Receive mocks base method.
func (m *MockEventStream) Receive() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Receive indicates an expected call of Receive.
func (mr *MockEventStreamMockRecorder) Receive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockEventStream)(nil).Receive))
}
