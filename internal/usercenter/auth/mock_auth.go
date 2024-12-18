// Copyright 2024 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rosas/onex/internal/usercenter/auth (interfaces: AuthProvider,AuthzInterface,AuthnInterface)

// Package auth is a generated GoMock package.
package auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	authn "github.com/rosas/onex/pkg/authn"
)

// MockAuthProvider is a mock of AuthProvider interface.
type MockAuthProvider struct {
	ctrl     *gomock.Controller
	recorder *MockAuthProviderMockRecorder
}

// MockAuthProviderMockRecorder is the mock recorder for MockAuthProvider.
type MockAuthProviderMockRecorder struct {
	mock *MockAuthProvider
}

// NewMockAuthProvider creates a new mock instance.
func NewMockAuthProvider(ctrl *gomock.Controller) *MockAuthProvider {
	mock := &MockAuthProvider{ctrl: ctrl}
	mock.recorder = &MockAuthProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthProvider) EXPECT() *MockAuthProviderMockRecorder {
	return m.recorder
}

// Authorize mocks base method.
func (m *MockAuthProvider) Authorize(arg0 ...any) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Authorize", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthProviderMockRecorder) Authorize(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthProvider)(nil).Authorize), arg0...)
}

// Sign mocks base method.
func (m *MockAuthProvider) Sign(arg0 context.Context, arg1 string) (authn.IToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0, arg1)
	ret0, _ := ret[0].(authn.IToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockAuthProviderMockRecorder) Sign(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockAuthProvider)(nil).Sign), arg0, arg1)
}

// Verify mocks base method.
func (m *MockAuthProvider) Verify(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockAuthProviderMockRecorder) Verify(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockAuthProvider)(nil).Verify), arg0)
}

// MockAuthzInterface is a mock of AuthzInterface interface.
type MockAuthzInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthzInterfaceMockRecorder
}

// MockAuthzInterfaceMockRecorder is the mock recorder for MockAuthzInterface.
type MockAuthzInterfaceMockRecorder struct {
	mock *MockAuthzInterface
}

// NewMockAuthzInterface creates a new mock instance.
func NewMockAuthzInterface(ctrl *gomock.Controller) *MockAuthzInterface {
	mock := &MockAuthzInterface{ctrl: ctrl}
	mock.recorder = &MockAuthzInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthzInterface) EXPECT() *MockAuthzInterfaceMockRecorder {
	return m.recorder
}

// Authorize mocks base method.
func (m *MockAuthzInterface) Authorize(arg0 ...any) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Authorize", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthzInterfaceMockRecorder) Authorize(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthzInterface)(nil).Authorize), arg0...)
}

// MockAuthnInterface is a mock of AuthnInterface interface.
type MockAuthnInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthnInterfaceMockRecorder
}

// MockAuthnInterfaceMockRecorder is the mock recorder for MockAuthnInterface.
type MockAuthnInterfaceMockRecorder struct {
	mock *MockAuthnInterface
}

// NewMockAuthnInterface creates a new mock instance.
func NewMockAuthnInterface(ctrl *gomock.Controller) *MockAuthnInterface {
	mock := &MockAuthnInterface{ctrl: ctrl}
	mock.recorder = &MockAuthnInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthnInterface) EXPECT() *MockAuthnInterfaceMockRecorder {
	return m.recorder
}

// Sign mocks base method.
func (m *MockAuthnInterface) Sign(arg0 context.Context, arg1 string) (authn.IToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0, arg1)
	ret0, _ := ret[0].(authn.IToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockAuthnInterfaceMockRecorder) Sign(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockAuthnInterface)(nil).Sign), arg0, arg1)
}

// Verify mocks base method.
func (m *MockAuthnInterface) Verify(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockAuthnInterfaceMockRecorder) Verify(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockAuthnInterface)(nil).Verify), arg0)
}
