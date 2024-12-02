// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// ResourceRequirementsApplyConfiguration represents an declarative configuration of the ResourceRequirements type for use
// with apply.
type ResourceRequirementsApplyConfiguration struct {
	Limits   *v1.ResourceList                  `json:"limits,omitempty"`
	Requests *v1.ResourceList                  `json:"requests,omitempty"`
	Claims   []ResourceClaimApplyConfiguration `json:"claims,omitempty"`
}

// ResourceRequirementsApplyConfiguration constructs an declarative configuration of the ResourceRequirements type for use with
// apply.
func ResourceRequirements() *ResourceRequirementsApplyConfiguration {
	return &ResourceRequirementsApplyConfiguration{}
}

// WithLimits sets the Limits field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Limits field is set to the value of the last call.
func (b *ResourceRequirementsApplyConfiguration) WithLimits(value v1.ResourceList) *ResourceRequirementsApplyConfiguration {
	b.Limits = &value
	return b
}

// WithRequests sets the Requests field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Requests field is set to the value of the last call.
func (b *ResourceRequirementsApplyConfiguration) WithRequests(value v1.ResourceList) *ResourceRequirementsApplyConfiguration {
	b.Requests = &value
	return b
}

// WithClaims adds the given value to the Claims field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Claims field.
func (b *ResourceRequirementsApplyConfiguration) WithClaims(values ...*ResourceClaimApplyConfiguration) *ResourceRequirementsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithClaims")
		}
		b.Claims = append(b.Claims, *values[i])
	}
	return b
}