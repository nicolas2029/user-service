// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IsUnauthorizeder is an autogenerated mock type for the IsUnauthorizeder type
type IsUnauthorizeder struct {
	mock.Mock
}

// IsUnauthorized provides a mock function with given fields:
func (_m *IsUnauthorizeder) IsUnauthorized() {
	_m.Called()
}

type mockConstructorTestingTNewIsUnauthorizeder interface {
	mock.TestingT
	Cleanup(func())
}

// NewIsUnauthorizeder creates a new instance of IsUnauthorizeder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIsUnauthorizeder(t mockConstructorTestingTNewIsUnauthorizeder) *IsUnauthorizeder {
	mock := &IsUnauthorizeder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
