// Code generated by MockGen. DO NOT EDIT.
// Source: cars_service.go

// Package cars is a generated GoMock package.
package cars

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCarHandler is a mock of CarHandler interface.
type MockCarHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCarHandlerMockRecorder
}

// MockCarHandlerMockRecorder is the mock recorder for MockCarHandler.
type MockCarHandlerMockRecorder struct {
	mock *MockCarHandler
}

// NewMockCarHandler creates a new mock instance.
func NewMockCarHandler(ctrl *gomock.Controller) *MockCarHandler {
	mock := &MockCarHandler{ctrl: ctrl}
	mock.recorder = &MockCarHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarHandler) EXPECT() *MockCarHandlerMockRecorder {
	return m.recorder
}

// CreateCar mocks base method.
func (m *MockCarHandler) CreateCar(car *CarPreview) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCar", car)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCar indicates an expected call of CreateCar.
func (mr *MockCarHandlerMockRecorder) CreateCar(car interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCar", reflect.TypeOf((*MockCarHandler)(nil).CreateCar), car)
}

// FindCarByID mocks base method.
func (m *MockCarHandler) FindCarByID(id uint) (*CarPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCarByID", id)
	ret0, _ := ret[0].(*CarPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCarByID indicates an expected call of FindCarByID.
func (mr *MockCarHandlerMockRecorder) FindCarByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCarByID", reflect.TypeOf((*MockCarHandler)(nil).FindCarByID), id)
}

// Search mocks base method.
func (m *MockCarHandler) Search(c *CarSearch) ([]CarPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", c)
	ret0, _ := ret[0].([]CarPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockCarHandlerMockRecorder) Search(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockCarHandler)(nil).Search), c)
}

// ValidateCar mocks base method.
func (m *MockCarHandler) ValidateCar(car *CarPreview) (bool, []string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCar", car)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]string)
	return ret0, ret1
}

// ValidateCar indicates an expected call of ValidateCar.
func (mr *MockCarHandlerMockRecorder) ValidateCar(car interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCar", reflect.TypeOf((*MockCarHandler)(nil).ValidateCar), car)
}