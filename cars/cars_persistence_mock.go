// Code generated by MockGen. DO NOT EDIT.
// Source: cars_persistence.go

// Package cars is a generated GoMock package.
package cars

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCarPersistence is a mock of CarPersistence interface.
type MockCarPersistence struct {
	ctrl     *gomock.Controller
	recorder *MockCarPersistenceMockRecorder
}

// MockCarPersistenceMockRecorder is the mock recorder for MockCarPersistence.
type MockCarPersistenceMockRecorder struct {
	mock *MockCarPersistence
}

// NewMockCarPersistence creates a new mock instance.
func NewMockCarPersistence(ctrl *gomock.Controller) *MockCarPersistence {
	mock := &MockCarPersistence{ctrl: ctrl}
	mock.recorder = &MockCarPersistenceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarPersistence) EXPECT() *MockCarPersistenceMockRecorder {
	return m.recorder
}

// AddFeatures mocks base method.
func (m *MockCarPersistence) AddFeatures(car []CarsFeature) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFeatures", car)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFeatures indicates an expected call of AddFeatures.
func (mr *MockCarPersistenceMockRecorder) AddFeatures(car interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFeatures", reflect.TypeOf((*MockCarPersistence)(nil).AddFeatures), car)
}

// CreateCar mocks base method.
func (m *MockCarPersistence) CreateCar(car *Car) (*Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCar", car)
	ret0, _ := ret[0].(*Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCar indicates an expected call of CreateCar.
func (mr *MockCarPersistenceMockRecorder) CreateCar(car interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCar", reflect.TypeOf((*MockCarPersistence)(nil).CreateCar), car)
}

// Find mocks base method.
func (m *MockCarPersistence) Find(c *CarSearch) ([]CarPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", c)
	ret0, _ := ret[0].([]CarPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCarPersistenceMockRecorder) Find(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCarPersistence)(nil).Find), c)
}

// FindByID mocks base method.
func (m *MockCarPersistence) FindByID(id uint) (*CarPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*CarPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockCarPersistenceMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCarPersistence)(nil).FindByID), id)
}