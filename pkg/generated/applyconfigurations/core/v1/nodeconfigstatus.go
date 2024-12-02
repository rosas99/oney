// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NodeConfigStatusApplyConfiguration represents an declarative configuration of the NodeConfigStatus type for use
// with apply.
type NodeConfigStatusApplyConfiguration struct {
	Assigned      *NodeConfigSourceApplyConfiguration `json:"assigned,omitempty"`
	Active        *NodeConfigSourceApplyConfiguration `json:"active,omitempty"`
	LastKnownGood *NodeConfigSourceApplyConfiguration `json:"lastKnownGood,omitempty"`
	Error         *string                             `json:"error,omitempty"`
}

// NodeConfigStatusApplyConfiguration constructs an declarative configuration of the NodeConfigStatus type for use with
// apply.
func NodeConfigStatus() *NodeConfigStatusApplyConfiguration {
	return &NodeConfigStatusApplyConfiguration{}
}

// WithAssigned sets the Assigned field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Assigned field is set to the value of the last call.
func (b *NodeConfigStatusApplyConfiguration) WithAssigned(value *NodeConfigSourceApplyConfiguration) *NodeConfigStatusApplyConfiguration {
	b.Assigned = value
	return b
}

// WithActive sets the Active field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Active field is set to the value of the last call.
func (b *NodeConfigStatusApplyConfiguration) WithActive(value *NodeConfigSourceApplyConfiguration) *NodeConfigStatusApplyConfiguration {
	b.Active = value
	return b
}

// WithLastKnownGood sets the LastKnownGood field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastKnownGood field is set to the value of the last call.
func (b *NodeConfigStatusApplyConfiguration) WithLastKnownGood(value *NodeConfigSourceApplyConfiguration) *NodeConfigStatusApplyConfiguration {
	b.LastKnownGood = value
	return b
}

// WithError sets the Error field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Error field is set to the value of the last call.
func (b *NodeConfigStatusApplyConfiguration) WithError(value string) *NodeConfigStatusApplyConfiguration {
	b.Error = &value
	return b
}
