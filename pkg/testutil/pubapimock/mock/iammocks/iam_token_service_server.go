// Code generated by mockery v2.38.0. DO NOT EDIT.

package iammocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	iam "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

// IamTokenServiceServer is an autogenerated mock type for the IamTokenServiceServer type
type IamTokenServiceServer struct {
	mock.Mock
}

type IamTokenServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *IamTokenServiceServer) EXPECT() *IamTokenServiceServer_Expecter {
	return &IamTokenServiceServer_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *IamTokenServiceServer) Create(_a0 context.Context, _a1 *iam.CreateIamTokenRequest) (*iam.CreateIamTokenResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *iam.CreateIamTokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.CreateIamTokenRequest) (*iam.CreateIamTokenResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.CreateIamTokenRequest) *iam.CreateIamTokenResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.CreateIamTokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.CreateIamTokenRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IamTokenServiceServer_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type IamTokenServiceServer_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *iam.CreateIamTokenRequest
func (_e *IamTokenServiceServer_Expecter) Create(_a0 interface{}, _a1 interface{}) *IamTokenServiceServer_Create_Call {
	return &IamTokenServiceServer_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *IamTokenServiceServer_Create_Call) Run(run func(_a0 context.Context, _a1 *iam.CreateIamTokenRequest)) *IamTokenServiceServer_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*iam.CreateIamTokenRequest))
	})
	return _c
}

func (_c *IamTokenServiceServer_Create_Call) Return(_a0 *iam.CreateIamTokenResponse, _a1 error) *IamTokenServiceServer_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IamTokenServiceServer_Create_Call) RunAndReturn(run func(context.Context, *iam.CreateIamTokenRequest) (*iam.CreateIamTokenResponse, error)) *IamTokenServiceServer_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateForServiceAccount provides a mock function with given fields: _a0, _a1
func (_m *IamTokenServiceServer) CreateForServiceAccount(_a0 context.Context, _a1 *iam.CreateIamTokenForServiceAccountRequest) (*iam.CreateIamTokenResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateForServiceAccount")
	}

	var r0 *iam.CreateIamTokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.CreateIamTokenForServiceAccountRequest) (*iam.CreateIamTokenResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.CreateIamTokenForServiceAccountRequest) *iam.CreateIamTokenResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.CreateIamTokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.CreateIamTokenForServiceAccountRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IamTokenServiceServer_CreateForServiceAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateForServiceAccount'
type IamTokenServiceServer_CreateForServiceAccount_Call struct {
	*mock.Call
}

// CreateForServiceAccount is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *iam.CreateIamTokenForServiceAccountRequest
func (_e *IamTokenServiceServer_Expecter) CreateForServiceAccount(_a0 interface{}, _a1 interface{}) *IamTokenServiceServer_CreateForServiceAccount_Call {
	return &IamTokenServiceServer_CreateForServiceAccount_Call{Call: _e.mock.On("CreateForServiceAccount", _a0, _a1)}
}

func (_c *IamTokenServiceServer_CreateForServiceAccount_Call) Run(run func(_a0 context.Context, _a1 *iam.CreateIamTokenForServiceAccountRequest)) *IamTokenServiceServer_CreateForServiceAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*iam.CreateIamTokenForServiceAccountRequest))
	})
	return _c
}

func (_c *IamTokenServiceServer_CreateForServiceAccount_Call) Return(_a0 *iam.CreateIamTokenResponse, _a1 error) *IamTokenServiceServer_CreateForServiceAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IamTokenServiceServer_CreateForServiceAccount_Call) RunAndReturn(run func(context.Context, *iam.CreateIamTokenForServiceAccountRequest) (*iam.CreateIamTokenResponse, error)) *IamTokenServiceServer_CreateForServiceAccount_Call {
	_c.Call.Return(run)
	return _c
}

// NewIamTokenServiceServer creates a new instance of IamTokenServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIamTokenServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *IamTokenServiceServer {
	mock := &IamTokenServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
