// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	table "github.com/modular-project/protobuffers/information/table"
)

// TableService is an autogenerated mock type for the TableService type
type TableService struct {
	mock.Mock
}

// CreateInBatch provides a mock function with given fields: ctx, eID, qua
func (_m *TableService) CreateInBatch(ctx context.Context, eID uint64, qua uint32) ([]uint64, error) {
	ret := _m.Called(ctx, eID, qua)

	var r0 []uint64
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint32) []uint64); ok {
		r0 = rf(ctx, eID, qua)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint32) error); ok {
		r1 = rf(ctx, eID, qua)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, eID, qua
func (_m *TableService) Delete(ctx context.Context, eID uint64, qua uint32) (uint32, error) {
	ret := _m.Called(ctx, eID, qua)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint32) uint32); ok {
		r0 = rf(ctx, eID, qua)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint32) error); ok {
		r1 = rf(ctx, eID, qua)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFromEstablishment provides a mock function with given fields: _a0, _a1
func (_m *TableService) GetFromEstablishment(_a0 context.Context, _a1 uint64) ([]*table.Table, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*table.Table
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []*table.Table); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*table.Table)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTableService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTableService creates a new instance of TableService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTableService(t mockConstructorTestingTNewTableService) *TableService {
	mock := &TableService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}