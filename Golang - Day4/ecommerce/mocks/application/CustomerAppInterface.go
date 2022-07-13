// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "ecommerce/domain/entity"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// CustomerAppInterface is an autogenerated mock type for the CustomerAppInterface type
type CustomerAppInterface struct {
	mock.Mock
}

// AddCustomer provides a mock function with given fields: _a0
func (_m *CustomerAppInterface) AddCustomer(_a0 *entity.Customer) (*entity.Customer, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Customer
	if rf, ok := ret.Get(0).(func(*entity.Customer) *entity.Customer); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Customer) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerByID provides a mock function with given fields: _a0
func (_m *CustomerAppInterface) GetCustomerByID(_a0 uuid.UUID) (*entity.Customer, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Customer
	if rf, ok := ret.Get(0).(func(uuid.UUID) *entity.Customer); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCustomerAppInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewCustomerAppInterface creates a new instance of CustomerAppInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCustomerAppInterface(t mockConstructorTestingTNewCustomerAppInterface) *CustomerAppInterface {
	mock := &CustomerAppInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}