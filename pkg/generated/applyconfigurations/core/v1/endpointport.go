// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// EndpointPortApplyConfiguration represents an declarative configuration of the EndpointPort type for use
// with apply.
type EndpointPortApplyConfiguration struct {
	Name        *string      `json:"name,omitempty"`
	Port        *int32       `json:"port,omitempty"`
	Protocol    *v1.Protocol `json:"protocol,omitempty"`
	AppProtocol *string      `json:"appProtocol,omitempty"`
}

// EndpointPortApplyConfiguration constructs an declarative configuration of the EndpointPort type for use with
// apply.
func EndpointPort() *EndpointPortApplyConfiguration {
	return &EndpointPortApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *EndpointPortApplyConfiguration) WithName(value string) *EndpointPortApplyConfiguration {
	b.Name = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *EndpointPortApplyConfiguration) WithPort(value int32) *EndpointPortApplyConfiguration {
	b.Port = &value
	return b
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *EndpointPortApplyConfiguration) WithProtocol(value v1.Protocol) *EndpointPortApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithAppProtocol sets the AppProtocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AppProtocol field is set to the value of the last call.
func (b *EndpointPortApplyConfiguration) WithAppProtocol(value string) *EndpointPortApplyConfiguration {
	b.AppProtocol = &value
	return b
}
