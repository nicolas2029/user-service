// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	model "users-service/model"

	mock "github.com/stretchr/testify/mock"
)

// Canner is an autogenerated mock type for the Canner type
type Canner struct {
	mock.Mock
}

// CanHire provides a mock function with given fields: _a0, _a1, _a2
func (_m *Canner) CanHire(_a0 uint, _a1 string, _a2 *model.UserRole) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, string, *model.UserRole) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanUpdate provides a mock function with given fields: from, target
func (_m *Canner) CanUpdate(from uint, target uint) error {
	ret := _m.Called(from, target)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(from, target)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Equal provides a mock function with given fields: uID, rID
func (_m *Canner) Equal(uID uint, rID model.RoleID) (uint, error) {
	ret := _m.Called(uID, rID)

	var r0 uint
	if rf, ok := ret.Get(0).(func(uint, model.RoleID) uint); ok {
		r0 = rf(uID, rID)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, model.RoleID) error); ok {
		r1 = rf(uID, rID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Greater provides a mock function with given fields: uID, rID
func (_m *Canner) Greater(uID uint, rID model.RoleID) error {
	ret := _m.Called(uID, rID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, model.RoleID) error); ok {
		r0 = rf(uID, rID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Job provides a mock function with given fields: _a0
func (_m *Canner) Job(_a0 uint) (model.UserRole, error) {
	ret := _m.Called(_a0)

	var r0 model.UserRole
	if rf, ok := ret.Get(0).(func(uint) model.UserRole); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.UserRole)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCanner interface {
	mock.TestingT
	Cleanup(func())
}

// NewCanner creates a new instance of Canner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCanner(t mockConstructorTestingTNewCanner) *Canner {
	mock := &Canner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
