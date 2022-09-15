// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	bank "gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk/bank"
)

// ServiceAuthorizer is an autogenerated mock type for the ServiceAuthorizer type
type ServiceAuthorizer struct {
	mock.Mock
}

// Authorize provides a mock function with given fields:
func (_m *ServiceAuthorizer) Authorize() (bank.Authorization, error) {
	ret := _m.Called()

	var r0 bank.Authorization
	if rf, ok := ret.Get(0).(func() bank.Authorization); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bank.Authorization)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewServiceAuthorizer interface {
	mock.TestingT
	Cleanup(func())
}

// NewServiceAuthorizer creates a new instance of ServiceAuthorizer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServiceAuthorizer(t mockConstructorTestingTNewServiceAuthorizer) *ServiceAuthorizer {
	mock := &ServiceAuthorizer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
