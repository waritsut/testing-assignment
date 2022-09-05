// Code generated by MockGen. DO NOT EDIT.
// Source: sequence-number-service/internal/services/sequence_number_service (interfaces: Service)

// Package mock_sequence_number_service is a generated GoMock package.
package mock_sequence_number_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CalTheDataSet mocks base method.
func (m *MockService) CalTheDataSet(arg0 uint64) []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalTheDataSet", arg0)
	ret0, _ := ret[0].([]int)
	return ret0
}

// CalTheDataSet indicates an expected call of CalTheDataSet.
func (mr *MockServiceMockRecorder) CalTheDataSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalTheDataSet", reflect.TypeOf((*MockService)(nil).CalTheDataSet), arg0)
}
