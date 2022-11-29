// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	bank "gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/bank"

	tls "crypto/tls"
)

// Bank is an autogenerated mock type for the Bank type
type Bank struct {
	mock.Mock
}

type Bank_Expecter struct {
	mock *mock.Mock
}

func (_m *Bank) EXPECT() *Bank_Expecter {
	return &Bank_Expecter{mock: &_m.Mock}
}

// GetAuthUrl provides a mock function with given fields:
func (_m *Bank) GetAuthUrl() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Bank_GetAuthUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthUrl'
type Bank_GetAuthUrl_Call struct {
	*mock.Call
}

// GetAuthUrl is a helper method to define mock.On call
func (_e *Bank_Expecter) GetAuthUrl() *Bank_GetAuthUrl_Call {
	return &Bank_GetAuthUrl_Call{Call: _e.mock.On("GetAuthUrl")}
}

func (_c *Bank_GetAuthUrl_Call) Run(run func()) *Bank_GetAuthUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Bank_GetAuthUrl_Call) Return(_a0 string) *Bank_GetAuthUrl_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetAuthorization provides a mock function with given fields:
func (_m *Bank) GetAuthorization() bank.Authorization {
	ret := _m.Called()

	var r0 bank.Authorization
	if rf, ok := ret.Get(0).(func() bank.Authorization); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bank.Authorization)
	}

	return r0
}

// Bank_GetAuthorization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthorization'
type Bank_GetAuthorization_Call struct {
	*mock.Call
}

// GetAuthorization is a helper method to define mock.On call
func (_e *Bank_Expecter) GetAuthorization() *Bank_GetAuthorization_Call {
	return &Bank_GetAuthorization_Call{Call: _e.mock.On("GetAuthorization")}
}

func (_c *Bank_GetAuthorization_Call) Run(run func()) *Bank_GetAuthorization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Bank_GetAuthorization_Call) Return(_a0 bank.Authorization) *Bank_GetAuthorization_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetCertificate provides a mock function with given fields:
func (_m *Bank) GetCertificate() *tls.Certificate {
	ret := _m.Called()

	var r0 *tls.Certificate
	if rf, ok := ret.Get(0).(func() *tls.Certificate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tls.Certificate)
		}
	}

	return r0
}

// Bank_GetCertificate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCertificate'
type Bank_GetCertificate_Call struct {
	*mock.Call
}

// GetCertificate is a helper method to define mock.On call
func (_e *Bank_Expecter) GetCertificate() *Bank_GetCertificate_Call {
	return &Bank_GetCertificate_Call{Call: _e.mock.On("GetCertificate")}
}

func (_c *Bank_GetCertificate_Call) Run(run func()) *Bank_GetCertificate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Bank_GetCertificate_Call) Return(_a0 *tls.Certificate) *Bank_GetCertificate_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetCredentials provides a mock function with given fields:
func (_m *Bank) GetCredentials() bank.Credentials {
	ret := _m.Called()

	var r0 bank.Credentials
	if rf, ok := ret.Get(0).(func() bank.Credentials); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bank.Credentials)
	}

	return r0
}

// Bank_GetCredentials_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCredentials'
type Bank_GetCredentials_Call struct {
	*mock.Call
}

// GetCredentials is a helper method to define mock.On call
func (_e *Bank_Expecter) GetCredentials() *Bank_GetCredentials_Call {
	return &Bank_GetCredentials_Call{Call: _e.mock.On("GetCredentials")}
}

func (_c *Bank_GetCredentials_Call) Run(run func()) *Bank_GetCredentials_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Bank_GetCredentials_Call) Return(_a0 bank.Credentials) *Bank_GetCredentials_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetUrl provides a mock function with given fields:
func (_m *Bank) GetUrl() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Bank_GetUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUrl'
type Bank_GetUrl_Call struct {
	*mock.Call
}

// GetUrl is a helper method to define mock.On call
func (_e *Bank_Expecter) GetUrl() *Bank_GetUrl_Call {
	return &Bank_GetUrl_Call{Call: _e.mock.On("GetUrl")}
}

func (_c *Bank_GetUrl_Call) Run(run func()) *Bank_GetUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Bank_GetUrl_Call) Return(_a0 string) *Bank_GetUrl_Call {
	_c.Call.Return(_a0)
	return _c
}

// SetAuthorization provides a mock function with given fields: authorization
func (_m *Bank) SetAuthorization(authorization bank.Authorization) {
	_m.Called(authorization)
}

// Bank_SetAuthorization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAuthorization'
type Bank_SetAuthorization_Call struct {
	*mock.Call
}

// SetAuthorization is a helper method to define mock.On call
//  - authorization bank.Authorization
func (_e *Bank_Expecter) SetAuthorization(authorization interface{}) *Bank_SetAuthorization_Call {
	return &Bank_SetAuthorization_Call{Call: _e.mock.On("SetAuthorization", authorization)}
}

func (_c *Bank_SetAuthorization_Call) Run(run func(authorization bank.Authorization)) *Bank_SetAuthorization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bank.Authorization))
	})
	return _c
}

func (_c *Bank_SetAuthorization_Call) Return() *Bank_SetAuthorization_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewBank interface {
	mock.TestingT
	Cleanup(func())
}

// NewBank creates a new instance of Bank. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBank(t mockConstructorTestingTNewBank) *Bank {
	mock := &Bank{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
