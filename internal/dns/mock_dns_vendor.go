// Code generated by MockGen. DO NOT EDIT.
// Source: mocks.go

// Package dns is a generated GoMock package.
package dns

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDNSProvider is a mock of DNSProvider interface.
type MockDNSProvider struct {
	ctrl     *gomock.Controller
	recorder *MockDNSProviderMockRecorder
}

// MockDNSProviderMockRecorder is the mock recorder for MockDNSProvider.
type MockDNSProviderMockRecorder struct {
	mock *MockDNSProvider
}

// NewMockDNSProvider creates a new mock instance.
func NewMockDNSProvider(ctrl *gomock.Controller) *MockDNSProvider {
	mock := &MockDNSProvider{ctrl: ctrl}
	mock.recorder = &MockDNSProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDNSProvider) EXPECT() *MockDNSProviderMockRecorder {
	return m.recorder
}

// CreateRecordSet mocks base method.
func (m *MockDNSProvider) CreateRecordSet(recordSetName, recordSetValue string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecordSet", recordSetName, recordSetValue)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRecordSet indicates an expected call of CreateRecordSet.
func (mr *MockDNSProviderMockRecorder) CreateRecordSet(recordSetName, recordSetValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecordSet", reflect.TypeOf((*MockDNSProvider)(nil).CreateRecordSet), recordSetName, recordSetValue)
}

// DeleteRecordSet mocks base method.
func (m *MockDNSProvider) DeleteRecordSet(recordSetName, recordSetValue string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecordSet", recordSetName, recordSetValue)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRecordSet indicates an expected call of DeleteRecordSet.
func (mr *MockDNSProviderMockRecorder) DeleteRecordSet(recordSetName, recordSetValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecordSet", reflect.TypeOf((*MockDNSProvider)(nil).DeleteRecordSet), recordSetName, recordSetValue)
}

// GetDomainName mocks base method.
func (m *MockDNSProvider) GetDomainName() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainName indicates an expected call of GetDomainName.
func (mr *MockDNSProviderMockRecorder) GetDomainName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainName", reflect.TypeOf((*MockDNSProvider)(nil).GetDomainName))
}

// GetRecordSet mocks base method.
func (m *MockDNSProvider) GetRecordSet(recordSetName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecordSet", recordSetName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordSet indicates an expected call of GetRecordSet.
func (mr *MockDNSProviderMockRecorder) GetRecordSet(recordSetName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordSet", reflect.TypeOf((*MockDNSProvider)(nil).GetRecordSet), recordSetName)
}

// UpdateRecordSet mocks base method.
func (m *MockDNSProvider) UpdateRecordSet(recordSetName, recordSetValue string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecordSet", recordSetName, recordSetValue)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRecordSet indicates an expected call of UpdateRecordSet.
func (mr *MockDNSProviderMockRecorder) UpdateRecordSet(recordSetName, recordSetValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecordSet", reflect.TypeOf((*MockDNSProvider)(nil).UpdateRecordSet), recordSetName, recordSetValue)
}
