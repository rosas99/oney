// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ExternalDocumentationApplyConfiguration represents an declarative configuration of the ExternalDocumentation type for use
// with apply.
type ExternalDocumentationApplyConfiguration struct {
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// ExternalDocumentationApplyConfiguration constructs an declarative configuration of the ExternalDocumentation type for use with
// apply.
func ExternalDocumentation() *ExternalDocumentationApplyConfiguration {
	return &ExternalDocumentationApplyConfiguration{}
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *ExternalDocumentationApplyConfiguration) WithDescription(value string) *ExternalDocumentationApplyConfiguration {
	b.Description = &value
	return b
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *ExternalDocumentationApplyConfiguration) WithURL(value string) *ExternalDocumentationApplyConfiguration {
	b.URL = &value
	return b
}
