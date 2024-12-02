// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NodeConditionApplyConfiguration represents an declarative configuration of the NodeCondition type for use
// with apply.
type NodeConditionApplyConfiguration struct {
	Type               *v1.NodeConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus   `json:"status,omitempty"`
	LastHeartbeatTime  *metav1.Time          `json:"lastHeartbeatTime,omitempty"`
	LastTransitionTime *metav1.Time          `json:"lastTransitionTime,omitempty"`
	Reason             *string               `json:"reason,omitempty"`
	Message            *string               `json:"message,omitempty"`
}

// NodeConditionApplyConfiguration constructs an declarative configuration of the NodeCondition type for use with
// apply.
func NodeCondition() *NodeConditionApplyConfiguration {
	return &NodeConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *NodeConditionApplyConfiguration) WithType(value v1.NodeConditionType) *NodeConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *NodeConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *NodeConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastHeartbeatTime sets the LastHeartbeatTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastHeartbeatTime field is set to the value of the last call.
func (b *NodeConditionApplyConfiguration) WithLastHeartbeatTime(value metav1.Time) *NodeConditionApplyConfiguration {
	b.LastHeartbeatTime = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *NodeConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *NodeConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *NodeConditionApplyConfiguration) WithReason(value string) *NodeConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *NodeConditionApplyConfiguration) WithMessage(value string) *NodeConditionApplyConfiguration {
	b.Message = &value
	return b
}
