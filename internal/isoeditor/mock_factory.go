// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/internal/isoeditor (interfaces: Factory)

// Package isoeditor is a generated GoMock package.
package isoeditor

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// WithEditor mocks base method.
func (m *MockFactory) WithEditor(arg0 context.Context, arg1 string, arg2 logrus.FieldLogger, arg3 EditFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithEditor", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithEditor indicates an expected call of WithEditor.
func (mr *MockFactoryMockRecorder) WithEditor(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithEditor", reflect.TypeOf((*MockFactory)(nil).WithEditor), arg0, arg1, arg2, arg3)
}
