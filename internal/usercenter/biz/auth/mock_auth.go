// Copyright 2024 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rosas/onex/internal/usercenter/biz/auth (interfaces: AuthBiz)

// Package auth is a generated GoMock package.
package auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/rosas/onex/pkg/api/usercenter/v1"
)

// MockAuthBiz is a mock of AuthBiz interface.
type MockAuthBiz struct {
	ctrl     *gomock.Controller
	recorder *MockAuthBizMockRecorder
}

// MockAuthBizMockRecorder is the mock recorder for MockAuthBiz.
type MockAuthBizMockRecorder struct {
	mock *MockAuthBiz
}

// NewMockAuthBiz creates a new mock instance.
func NewMockAuthBiz(ctrl *gomock.Controller) *MockAuthBiz {
	mock := &MockAuthBiz{ctrl: ctrl}
	mock.recorder = &MockAuthBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthBiz) EXPECT() *MockAuthBizMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuthBiz) Authenticate(arg0 context.Context, arg1 string) (*v1.AuthenticateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0, arg1)
	ret0, _ := ret[0].(*v1.AuthenticateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthBizMockRecorder) Authenticate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthBiz)(nil).Authenticate), arg0, arg1)
}

// Authorize mocks base method.
func (m *MockAuthBiz) Authorize(arg0 context.Context, arg1, arg2, arg3 string) (*v1.AuthorizeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.AuthorizeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthBizMockRecorder) Authorize(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthBiz)(nil).Authorize), arg0, arg1, arg2, arg3)
}

// Login mocks base method.
func (m *MockAuthBiz) Login(arg0 context.Context, arg1 *v1.LoginRequest) (*v1.LoginReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*v1.LoginReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthBizMockRecorder) Login(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthBiz)(nil).Login), arg0, arg1)
}

// Logout mocks base method.
func (m *MockAuthBiz) Logout(arg0 context.Context, arg1 *v1.LogoutRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockAuthBizMockRecorder) Logout(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthBiz)(nil).Logout), arg0, arg1)
}

// RefreshToken mocks base method.
func (m *MockAuthBiz) RefreshToken(arg0 context.Context, arg1 *v1.RefreshTokenRequest) (*v1.LoginReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", arg0, arg1)
	ret0, _ := ret[0].(*v1.LoginReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockAuthBizMockRecorder) RefreshToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockAuthBiz)(nil).RefreshToken), arg0, arg1)
}
