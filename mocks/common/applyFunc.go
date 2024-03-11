// Code generated by mockery v2.35.3. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// applyFunc is an autogenerated mock type for the applyFunc type
type applyFunc struct {
	mock.Mock
}

type applyFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *applyFunc) EXPECT() *applyFunc_Expecter {
	return &applyFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: r
func (_m *applyFunc) Execute(r *http.Request) {
	_m.Called(r)
}

// applyFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type applyFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - r *http.Request
func (_e *applyFunc_Expecter) Execute(r interface{}) *applyFunc_Execute_Call {
	return &applyFunc_Execute_Call{Call: _e.mock.On("Execute", r)}
}

func (_c *applyFunc_Execute_Call) Run(run func(r *http.Request)) *applyFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*http.Request))
	})
	return _c
}

func (_c *applyFunc_Execute_Call) Return() *applyFunc_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *applyFunc_Execute_Call) RunAndReturn(run func(*http.Request)) *applyFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// newApplyFunc creates a new instance of applyFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newApplyFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *applyFunc {
	mock := &applyFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
