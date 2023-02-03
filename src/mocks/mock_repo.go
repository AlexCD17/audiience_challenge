// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/rates/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// GetRates mocks base method.
func (m *MockIRepository) GetRates(state, estimationType string) (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRates", state, estimationType)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRates indicates an expected call of GetRates.
func (mr *MockIRepositoryMockRecorder) GetRates(state, estimationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRates", reflect.TypeOf((*MockIRepository)(nil).GetRates), state, estimationType)
}