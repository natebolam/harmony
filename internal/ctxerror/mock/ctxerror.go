// Code generated by MockGen. DO NOT EDIT.
// Source: ctxerror.go

// Package mock_ctxerror is a generated GoMock package.
package mock_ctxerror

import (
	gomock "github.com/golang/mock/gomock"
	ctxerror "github.com/harmony-one/harmony/internal/ctxerror"
	reflect "reflect"
)

// MockCtxError is a mock of CtxError interface
type MockCtxError struct {
	ctrl     *gomock.Controller
	recorder *MockCtxErrorMockRecorder
}

// MockCtxErrorMockRecorder is the mock recorder for MockCtxError
type MockCtxErrorMockRecorder struct {
	mock *MockCtxError
}

// NewMockCtxError creates a new mock instance
func NewMockCtxError(ctrl *gomock.Controller) *MockCtxError {
	mock := &MockCtxError{ctrl: ctrl}
	mock.recorder = &MockCtxErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCtxError) EXPECT() *MockCtxErrorMockRecorder {
	return m.recorder
}

// Error mocks base method
func (m *MockCtxError) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockCtxErrorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockCtxError)(nil).Error))
}

// Message mocks base method
func (m *MockCtxError) Message() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message")
	ret0, _ := ret[0].(string)
	return ret0
}

// Message indicates an expected call of Message
func (mr *MockCtxErrorMockRecorder) Message() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockCtxError)(nil).Message))
}

// Contexts mocks base method
func (m *MockCtxError) Contexts() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contexts")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Contexts indicates an expected call of Contexts
func (mr *MockCtxErrorMockRecorder) Contexts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contexts", reflect.TypeOf((*MockCtxError)(nil).Contexts))
}

// WithCause mocks base method
func (m *MockCtxError) WithCause(c error) ctxerror.CtxError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithCause", c)
	ret0, _ := ret[0].(ctxerror.CtxError)
	return ret0
}

// WithCause indicates an expected call of WithCause
func (mr *MockCtxErrorMockRecorder) WithCause(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithCause", reflect.TypeOf((*MockCtxError)(nil).WithCause), c)
}

// MockLog15Logger is a mock of Log15Logger interface
type MockLog15Logger struct {
	ctrl     *gomock.Controller
	recorder *MockLog15LoggerMockRecorder
}

// MockLog15LoggerMockRecorder is the mock recorder for MockLog15Logger
type MockLog15LoggerMockRecorder struct {
	mock *MockLog15Logger
}

// NewMockLog15Logger creates a new mock instance
func NewMockLog15Logger(ctrl *gomock.Controller) *MockLog15Logger {
	mock := &MockLog15Logger{ctrl: ctrl}
	mock.recorder = &MockLog15LoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLog15Logger) EXPECT() *MockLog15LoggerMockRecorder {
	return m.recorder
}

// Log15 mocks base method
func (m *MockLog15Logger) Log15(f ctxerror.Log15Func) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Log15", f)
}

// Log15 indicates an expected call of Log15
func (mr *MockLog15LoggerMockRecorder) Log15(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log15", reflect.TypeOf((*MockLog15Logger)(nil).Log15), f)
}
