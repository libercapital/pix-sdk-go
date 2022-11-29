// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	pixsdk "gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git"
	pix "gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/services/pix"

	webhook "gitlab.com/bavatech/architecture/software/libs/go-modules/pix-sdk.git/services/webhook"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// CreateWebhook provides a mock function with given fields: ctx, key, _a2
func (_m *Service) CreateWebhook(ctx context.Context, key string, _a2 webhook.CreateWebhook) error {
	ret := _m.Called(ctx, key, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, webhook.CreateWebhook) error); ok {
		r0 = rf(ctx, key, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_CreateWebhook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateWebhook'
type Service_CreateWebhook_Call struct {
	*mock.Call
}

// CreateWebhook is a helper method to define mock.On call
//  - ctx context.Context
//  - key string
//  - _a2 webhook.CreateWebhook
func (_e *Service_Expecter) CreateWebhook(ctx interface{}, key interface{}, _a2 interface{}) *Service_CreateWebhook_Call {
	return &Service_CreateWebhook_Call{Call: _e.mock.On("CreateWebhook", ctx, key, _a2)}
}

func (_c *Service_CreateWebhook_Call) Run(run func(ctx context.Context, key string, _a2 webhook.CreateWebhook)) *Service_CreateWebhook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(webhook.CreateWebhook))
	})
	return _c
}

func (_c *Service_CreateWebhook_Call) Return(_a0 error) *Service_CreateWebhook_Call {
	_c.Call.Return(_a0)
	return _c
}

// DeleteWebhook provides a mock function with given fields: ctx, key
func (_m *Service) DeleteWebhook(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_DeleteWebhook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteWebhook'
type Service_DeleteWebhook_Call struct {
	*mock.Call
}

// DeleteWebhook is a helper method to define mock.On call
//  - ctx context.Context
//  - key string
func (_e *Service_Expecter) DeleteWebhook(ctx interface{}, key interface{}) *Service_DeleteWebhook_Call {
	return &Service_DeleteWebhook_Call{Call: _e.mock.On("DeleteWebhook", ctx, key)}
}

func (_c *Service_DeleteWebhook_Call) Run(run func(ctx context.Context, key string)) *Service_DeleteWebhook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Service_DeleteWebhook_Call) Return(_a0 error) *Service_DeleteWebhook_Call {
	_c.Call.Return(_a0)
	return _c
}

// FindPix provides a mock function with given fields: e2eId
func (_m *Service) FindPix(e2eId string) (pix.Pix, error) {
	ret := _m.Called(e2eId)

	var r0 pix.Pix
	if rf, ok := ret.Get(0).(func(string) pix.Pix); ok {
		r0 = rf(e2eId)
	} else {
		r0 = ret.Get(0).(pix.Pix)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(e2eId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_FindPix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindPix'
type Service_FindPix_Call struct {
	*mock.Call
}

// FindPix is a helper method to define mock.On call
//  - e2eId string
func (_e *Service_Expecter) FindPix(e2eId interface{}) *Service_FindPix_Call {
	return &Service_FindPix_Call{Call: _e.mock.On("FindPix", e2eId)}
}

func (_c *Service_FindPix_Call) Run(run func(e2eId string)) *Service_FindPix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Service_FindPix_Call) Return(_a0 pix.Pix, _a1 error) *Service_FindPix_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// FindWebhook provides a mock function with given fields: ctx, key
func (_m *Service) FindWebhook(ctx context.Context, key string) (webhook.Webhook, error) {
	ret := _m.Called(ctx, key)

	var r0 webhook.Webhook
	if rf, ok := ret.Get(0).(func(context.Context, string) webhook.Webhook); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(webhook.Webhook)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_FindWebhook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindWebhook'
type Service_FindWebhook_Call struct {
	*mock.Call
}

// FindWebhook is a helper method to define mock.On call
//  - ctx context.Context
//  - key string
func (_e *Service_Expecter) FindWebhook(ctx interface{}, key interface{}) *Service_FindWebhook_Call {
	return &Service_FindWebhook_Call{Call: _e.mock.On("FindWebhook", ctx, key)}
}

func (_c *Service_FindWebhook_Call) Run(run func(ctx context.Context, key string)) *Service_FindWebhook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Service_FindWebhook_Call) Return(_a0 webhook.Webhook, _a1 error) *Service_FindWebhook_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListPix provides a mock function with given fields: listPix
func (_m *Service) ListPix(listPix pix.ListPix) (pix.ListPixResponse, error) {
	ret := _m.Called(listPix)

	var r0 pix.ListPixResponse
	if rf, ok := ret.Get(0).(func(pix.ListPix) pix.ListPixResponse); ok {
		r0 = rf(listPix)
	} else {
		r0 = ret.Get(0).(pix.ListPixResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pix.ListPix) error); ok {
		r1 = rf(listPix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_ListPix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPix'
type Service_ListPix_Call struct {
	*mock.Call
}

// ListPix is a helper method to define mock.On call
//  - listPix pix.ListPix
func (_e *Service_Expecter) ListPix(listPix interface{}) *Service_ListPix_Call {
	return &Service_ListPix_Call{Call: _e.mock.On("ListPix", listPix)}
}

func (_c *Service_ListPix_Call) Run(run func(listPix pix.ListPix)) *Service_ListPix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(pix.ListPix))
	})
	return _c
}

func (_c *Service_ListPix_Call) Return(_a0 pix.ListPixResponse, _a1 error) *Service_ListPix_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SetConfig provides a mock function with given fields: config
func (_m *Service) SetConfig(config pixsdk.Config) {
	_m.Called(config)
}

// Service_SetConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConfig'
type Service_SetConfig_Call struct {
	*mock.Call
}

// SetConfig is a helper method to define mock.On call
//  - config pixsdk.Config
func (_e *Service_Expecter) SetConfig(config interface{}) *Service_SetConfig_Call {
	return &Service_SetConfig_Call{Call: _e.mock.On("SetConfig", config)}
}

func (_c *Service_SetConfig_Call) Run(run func(config pixsdk.Config)) *Service_SetConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(pixsdk.Config))
	})
	return _c
}

func (_c *Service_SetConfig_Call) Return() *Service_SetConfig_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
