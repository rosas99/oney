// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// LifecycleHandlerApplyConfiguration represents an declarative configuration of the LifecycleHandler type for use
// with apply.
type LifecycleHandlerApplyConfiguration struct {
	Exec      *ExecActionApplyConfiguration      `json:"exec,omitempty"`
	HTTPGet   *HTTPGetActionApplyConfiguration   `json:"httpGet,omitempty"`
	TCPSocket *TCPSocketActionApplyConfiguration `json:"tcpSocket,omitempty"`
	Sleep     *SleepActionApplyConfiguration     `json:"sleep,omitempty"`
}

// LifecycleHandlerApplyConfiguration constructs an declarative configuration of the LifecycleHandler type for use with
// apply.
func LifecycleHandler() *LifecycleHandlerApplyConfiguration {
	return &LifecycleHandlerApplyConfiguration{}
}

// WithExec sets the Exec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Exec field is set to the value of the last call.
func (b *LifecycleHandlerApplyConfiguration) WithExec(value *ExecActionApplyConfiguration) *LifecycleHandlerApplyConfiguration {
	b.Exec = value
	return b
}

// WithHTTPGet sets the HTTPGet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPGet field is set to the value of the last call.
func (b *LifecycleHandlerApplyConfiguration) WithHTTPGet(value *HTTPGetActionApplyConfiguration) *LifecycleHandlerApplyConfiguration {
	b.HTTPGet = value
	return b
}

// WithTCPSocket sets the TCPSocket field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TCPSocket field is set to the value of the last call.
func (b *LifecycleHandlerApplyConfiguration) WithTCPSocket(value *TCPSocketActionApplyConfiguration) *LifecycleHandlerApplyConfiguration {
	b.TCPSocket = value
	return b
}

// WithSleep sets the Sleep field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Sleep field is set to the value of the last call.
func (b *LifecycleHandlerApplyConfiguration) WithSleep(value *SleepActionApplyConfiguration) *LifecycleHandlerApplyConfiguration {
	b.Sleep = value
	return b
}
