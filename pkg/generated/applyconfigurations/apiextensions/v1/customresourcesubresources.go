// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// CustomResourceSubresourcesApplyConfiguration represents an declarative configuration of the CustomResourceSubresources type for use
// with apply.
type CustomResourceSubresourcesApplyConfiguration struct {
	Status *v1.CustomResourceSubresourceStatus               `json:"status,omitempty"`
	Scale  *CustomResourceSubresourceScaleApplyConfiguration `json:"scale,omitempty"`
}

// CustomResourceSubresourcesApplyConfiguration constructs an declarative configuration of the CustomResourceSubresources type for use with
// apply.
func CustomResourceSubresources() *CustomResourceSubresourcesApplyConfiguration {
	return &CustomResourceSubresourcesApplyConfiguration{}
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *CustomResourceSubresourcesApplyConfiguration) WithStatus(value v1.CustomResourceSubresourceStatus) *CustomResourceSubresourcesApplyConfiguration {
	b.Status = &value
	return b
}

// WithScale sets the Scale field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scale field is set to the value of the last call.
func (b *CustomResourceSubresourcesApplyConfiguration) WithScale(value *CustomResourceSubresourceScaleApplyConfiguration) *CustomResourceSubresourcesApplyConfiguration {
	b.Scale = value
	return b
}
