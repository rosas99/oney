// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// ChainSpecApplyConfiguration represents an declarative configuration of the ChainSpec type for use
// with apply.
type ChainSpecApplyConfiguration struct {
	DisplayName            *string `json:"displayName,omitempty"`
	MinerType              *string `json:"minerType,omitempty"`
	Image                  *string `json:"image,omitempty"`
	MinMineIntervalSeconds *int32  `json:"minMineIntervalSeconds,omitempty"`
	BootstrapAccount       *string `json:"bootstrapAccount,omitempty"`
}

// ChainSpecApplyConfiguration constructs an declarative configuration of the ChainSpec type for use with
// apply.
func ChainSpec() *ChainSpecApplyConfiguration {
	return &ChainSpecApplyConfiguration{}
}

// WithDisplayName sets the DisplayName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisplayName field is set to the value of the last call.
func (b *ChainSpecApplyConfiguration) WithDisplayName(value string) *ChainSpecApplyConfiguration {
	b.DisplayName = &value
	return b
}

// WithMinerType sets the MinerType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinerType field is set to the value of the last call.
func (b *ChainSpecApplyConfiguration) WithMinerType(value string) *ChainSpecApplyConfiguration {
	b.MinerType = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *ChainSpecApplyConfiguration) WithImage(value string) *ChainSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithMinMineIntervalSeconds sets the MinMineIntervalSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinMineIntervalSeconds field is set to the value of the last call.
func (b *ChainSpecApplyConfiguration) WithMinMineIntervalSeconds(value int32) *ChainSpecApplyConfiguration {
	b.MinMineIntervalSeconds = &value
	return b
}

// WithBootstrapAccount sets the BootstrapAccount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BootstrapAccount field is set to the value of the last call.
func (b *ChainSpecApplyConfiguration) WithBootstrapAccount(value string) *ChainSpecApplyConfiguration {
	b.BootstrapAccount = &value
	return b
}
