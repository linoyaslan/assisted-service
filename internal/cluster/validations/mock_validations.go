// Code generated by MockGen. DO NOT EDIT.
// Source: validations.go

// Package validations is a generated GoMock package.
package validations

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	auth "github.com/openshift/assisted-service/pkg/auth"
)

// MockPullSecretValidator is a mock of PullSecretValidator interface.
type MockPullSecretValidator struct {
	ctrl     *gomock.Controller
	recorder *MockPullSecretValidatorMockRecorder
}

// MockPullSecretValidatorMockRecorder is the mock recorder for MockPullSecretValidator.
type MockPullSecretValidatorMockRecorder struct {
	mock *MockPullSecretValidator
}

// NewMockPullSecretValidator creates a new mock instance.
func NewMockPullSecretValidator(ctrl *gomock.Controller) *MockPullSecretValidator {
	mock := &MockPullSecretValidator{ctrl: ctrl}
	mock.recorder = &MockPullSecretValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPullSecretValidator) EXPECT() *MockPullSecretValidatorMockRecorder {
	return m.recorder
}

// ValidatePullSecret mocks base method.
func (m *MockPullSecretValidator) ValidatePullSecret(secret, username string, authHandler auth.Authenticator) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePullSecret", secret, username, authHandler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidatePullSecret indicates an expected call of ValidatePullSecret.
func (mr *MockPullSecretValidatorMockRecorder) ValidatePullSecret(secret, username, authHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePullSecret", reflect.TypeOf((*MockPullSecretValidator)(nil).ValidatePullSecret), secret, username, authHandler)
}
