// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PreferredSchedulingTermApplyConfiguration represents an declarative configuration of the PreferredSchedulingTerm type for use
// with apply.
type PreferredSchedulingTermApplyConfiguration struct {
	Weight     *int32                              `json:"weight,omitempty"`
	Preference *NodeSelectorTermApplyConfiguration `json:"preference,omitempty"`
}

// PreferredSchedulingTermApplyConfiguration constructs an declarative configuration of the PreferredSchedulingTerm type for use with
// apply.
func PreferredSchedulingTerm() *PreferredSchedulingTermApplyConfiguration {
	return &PreferredSchedulingTermApplyConfiguration{}
}

// WithWeight sets the Weight field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Weight field is set to the value of the last call.
func (b *PreferredSchedulingTermApplyConfiguration) WithWeight(value int32) *PreferredSchedulingTermApplyConfiguration {
	b.Weight = &value
	return b
}

// WithPreference sets the Preference field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Preference field is set to the value of the last call.
func (b *PreferredSchedulingTermApplyConfiguration) WithPreference(value *NodeSelectorTermApplyConfiguration) *PreferredSchedulingTermApplyConfiguration {
	b.Preference = value
	return b
}
