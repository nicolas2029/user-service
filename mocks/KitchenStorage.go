// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	model "users-service/model"

	mock "github.com/stretchr/testify/mock"
)

// KitchenStorage is an autogenerated mock type for the KitchenStorage type
type KitchenStorage struct {
	mock.Mock
}

// Delete provides a mock function with given fields: eID, kID
func (_m *KitchenStorage) Delete(eID uint, kID uint) error {
	ret := _m.Called(eID, kID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(eID, kID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInESTB provides a mock function with given fields: _a0
func (_m *KitchenStorage) GetInESTB(_a0 uint) ([]model.Kitchen, error) {
	ret := _m.Called(_a0)

	var r0 []model.Kitchen
	if rf, ok := ret.Get(0).(func(uint) []model.Kitchen); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Kitchen)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: eID, kID, k
func (_m *KitchenStorage) Update(eID uint, kID uint, k *model.Kitchen) error {
	ret := _m.Called(eID, kID, k)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, *model.Kitchen) error); ok {
		r0 = rf(eID, kID, k)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewKitchenStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewKitchenStorage creates a new instance of KitchenStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKitchenStorage(t mockConstructorTestingTNewKitchenStorage) *KitchenStorage {
	mock := &KitchenStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
