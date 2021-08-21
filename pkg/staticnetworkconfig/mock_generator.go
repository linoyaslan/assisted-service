// Code generated by MockGen. DO NOT EDIT.
// Source: generator.go

// Package staticnetworkconfig is a generated GoMock package.
package staticnetworkconfig

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/openshift/assisted-service/models"
)

// MockStaticNetworkConfig is a mock of StaticNetworkConfig interface.
type MockStaticNetworkConfig struct {
	ctrl     *gomock.Controller
	recorder *MockStaticNetworkConfigMockRecorder
}

// MockStaticNetworkConfigMockRecorder is the mock recorder for MockStaticNetworkConfig.
type MockStaticNetworkConfigMockRecorder struct {
	mock *MockStaticNetworkConfig
}

// NewMockStaticNetworkConfig creates a new mock instance.
func NewMockStaticNetworkConfig(ctrl *gomock.Controller) *MockStaticNetworkConfig {
	mock := &MockStaticNetworkConfig{ctrl: ctrl}
	mock.recorder = &MockStaticNetworkConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStaticNetworkConfig) EXPECT() *MockStaticNetworkConfigMockRecorder {
	return m.recorder
}

// FormatStaticNetworkConfigForDB mocks base method.
func (m *MockStaticNetworkConfig) FormatStaticNetworkConfigForDB(staticNetworkConfig []*models.HostStaticNetworkConfig) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatStaticNetworkConfigForDB", staticNetworkConfig)
	ret0, _ := ret[0].(string)
	return ret0
}

// FormatStaticNetworkConfigForDB indicates an expected call of FormatStaticNetworkConfigForDB.
func (mr *MockStaticNetworkConfigMockRecorder) FormatStaticNetworkConfigForDB(staticNetworkConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatStaticNetworkConfigForDB", reflect.TypeOf((*MockStaticNetworkConfig)(nil).FormatStaticNetworkConfigForDB), staticNetworkConfig)
}

// GenerateStaticNetworkConfigData mocks base method.
func (m *MockStaticNetworkConfig) GenerateStaticNetworkConfigData(hostsYAMLS string) ([]StaticNetworkConfigData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateStaticNetworkConfigData", hostsYAMLS)
	ret0, _ := ret[0].([]StaticNetworkConfigData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateStaticNetworkConfigData indicates an expected call of GenerateStaticNetworkConfigData.
func (mr *MockStaticNetworkConfigMockRecorder) GenerateStaticNetworkConfigData(hostsYAMLS interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateStaticNetworkConfigData", reflect.TypeOf((*MockStaticNetworkConfig)(nil).GenerateStaticNetworkConfigData), hostsYAMLS)
}

// ValidateStaticConfigParams mocks base method.
func (m *MockStaticNetworkConfig) ValidateStaticConfigParams(staticNetworkConfig []*models.HostStaticNetworkConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateStaticConfigParams", staticNetworkConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateStaticConfigParams indicates an expected call of ValidateStaticConfigParams.
func (mr *MockStaticNetworkConfigMockRecorder) ValidateStaticConfigParams(staticNetworkConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateStaticConfigParams", reflect.TypeOf((*MockStaticNetworkConfig)(nil).ValidateStaticConfigParams), staticNetworkConfig)
}
