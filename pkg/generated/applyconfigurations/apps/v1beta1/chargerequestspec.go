// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// ChargeRequestSpecApplyConfiguration represents an declarative configuration of the ChargeRequestSpec type for use
// with apply.
type ChargeRequestSpecApplyConfiguration struct {
	From     *string `json:"from,omitempty"`
	Password *string `json:"password,omitempty"`
}

// ChargeRequestSpecApplyConfiguration constructs an declarative configuration of the ChargeRequestSpec type for use with
// apply.
func ChargeRequestSpec() *ChargeRequestSpecApplyConfiguration {
	return &ChargeRequestSpecApplyConfiguration{}
}

// WithFrom sets the From field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the From field is set to the value of the last call.
func (b *ChargeRequestSpecApplyConfiguration) WithFrom(value string) *ChargeRequestSpecApplyConfiguration {
	b.From = &value
	return b
}

// WithPassword sets the Password field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Password field is set to the value of the last call.
func (b *ChargeRequestSpecApplyConfiguration) WithPassword(value string) *ChargeRequestSpecApplyConfiguration {
	b.Password = &value
	return b
}
