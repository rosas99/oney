// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// LimitedPriorityLevelConfigurationApplyConfiguration represents an declarative configuration of the LimitedPriorityLevelConfiguration type for use
// with apply.
type LimitedPriorityLevelConfigurationApplyConfiguration struct {
	NominalConcurrencyShares *int32                           `json:"nominalConcurrencyShares,omitempty"`
	LimitResponse            *LimitResponseApplyConfiguration `json:"limitResponse,omitempty"`
	LendablePercent          *int32                           `json:"lendablePercent,omitempty"`
	BorrowingLimitPercent    *int32                           `json:"borrowingLimitPercent,omitempty"`
}

// LimitedPriorityLevelConfigurationApplyConfiguration constructs an declarative configuration of the LimitedPriorityLevelConfiguration type for use with
// apply.
func LimitedPriorityLevelConfiguration() *LimitedPriorityLevelConfigurationApplyConfiguration {
	return &LimitedPriorityLevelConfigurationApplyConfiguration{}
}

// WithNominalConcurrencyShares sets the NominalConcurrencyShares field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NominalConcurrencyShares field is set to the value of the last call.
func (b *LimitedPriorityLevelConfigurationApplyConfiguration) WithNominalConcurrencyShares(value int32) *LimitedPriorityLevelConfigurationApplyConfiguration {
	b.NominalConcurrencyShares = &value
	return b
}

// WithLimitResponse sets the LimitResponse field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LimitResponse field is set to the value of the last call.
func (b *LimitedPriorityLevelConfigurationApplyConfiguration) WithLimitResponse(value *LimitResponseApplyConfiguration) *LimitedPriorityLevelConfigurationApplyConfiguration {
	b.LimitResponse = value
	return b
}

// WithLendablePercent sets the LendablePercent field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LendablePercent field is set to the value of the last call.
func (b *LimitedPriorityLevelConfigurationApplyConfiguration) WithLendablePercent(value int32) *LimitedPriorityLevelConfigurationApplyConfiguration {
	b.LendablePercent = &value
	return b
}

// WithBorrowingLimitPercent sets the BorrowingLimitPercent field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BorrowingLimitPercent field is set to the value of the last call.
func (b *LimitedPriorityLevelConfigurationApplyConfiguration) WithBorrowingLimitPercent(value int32) *LimitedPriorityLevelConfigurationApplyConfiguration {
	b.BorrowingLimitPercent = &value
	return b
}
