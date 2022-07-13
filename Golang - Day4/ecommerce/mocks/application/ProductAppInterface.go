// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "ecommerce/domain/entity"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ProductAppInterface is an autogenerated mock type for the ProductAppInterface type
type ProductAppInterface struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: _a0
func (_m *ProductAppInterface) AddProduct(_a0 *entity.Product) (*entity.Product, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Product
	if rf, ok := ret.Get(0).(func(*entity.Product) *entity.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Product) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllProducts provides a mock function with given fields:
func (_m *ProductAppInterface) GetAllProducts() ([]entity.Product, error) {
	ret := _m.Called()

	var r0 []entity.Product
	if rf, ok := ret.Get(0).(func() []entity.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductByID provides a mock function with given fields: _a0
func (_m *ProductAppInterface) GetProductByID(_a0 uuid.UUID) (*entity.Product, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Product
	if rf, ok := ret.Get(0).(func(uuid.UUID) *entity.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
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

// UpdateProduct provides a mock function with given fields: _a0
func (_m *ProductAppInterface) UpdateProduct(_a0 *entity.Product) (*entity.Product, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Product
	if rf, ok := ret.Get(0).(func(*entity.Product) *entity.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Product) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductAppInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductAppInterface creates a new instance of ProductAppInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductAppInterface(t mockConstructorTestingTNewProductAppInterface) *ProductAppInterface {
	mock := &ProductAppInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
