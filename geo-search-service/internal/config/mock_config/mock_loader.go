// Code generated by MockGen. DO NOT EDIT.
// Source: loader.go

// Package mock_config is a generated GoMock package.
package mock_config

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConfigLoader is a mock of ConfigLoader interface.
type MockConfigLoader struct {
	ctrl     *gomock.Controller
	recorder *MockConfigLoaderMockRecorder
}

// MockConfigLoaderMockRecorder is the mock recorder for MockConfigLoader.
type MockConfigLoaderMockRecorder struct {
	mock *MockConfigLoader
}

// NewMockConfigLoader creates a new mock instance.
func NewMockConfigLoader(ctrl *gomock.Controller) *MockConfigLoader {
	mock := &MockConfigLoader{ctrl: ctrl}
	mock.recorder = &MockConfigLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigLoader) EXPECT() *MockConfigLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *MockConfigLoader) Load() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockConfigLoaderMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockConfigLoader)(nil).Load))
}
