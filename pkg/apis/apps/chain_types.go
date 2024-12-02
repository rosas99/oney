// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

package apps

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ChainFinalizer is the finalizer used by the Chain controller to
	// clean up referenced template resources if necessary when a Chain is being deleted.
	ChainFinalizer = "chain.onex.io/finalizer"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Chain is the Schema for the chains API.
type Chain struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the chain.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec ChainSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status is the most recently observed status of the Chain.
	// This data may be out of date by some window of time.
	// Populated by the system.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Status ChainStatus `json:"status,omitempty"`
}

// ChainSpec defines the desired state of Chain.
type ChainSpec struct {
	// The display name of the chain.
	// +optional
	DisplayName string `json:"displayName,omitempty" protobuf:"bytes,1,opt,name=displayName"`

	// Genesis node machine configuration.
	// +optional
	MinerType string `json:"minerType,omitempty" protobuf:"bytes,2,opt,name=minerType"`

	// Image specify the blockchain node image.
	// +optional
	Image string `json:"image,omitempty" protobuf:"bytes,1,opt,name=image"`

	// Minimum number of seconds for the miners to mine a block.
	// +optional
	MinMineIntervalSeconds int32

	// Default bootstrap OneX's Genesis account with 1M TBB tokens.
	// This field is automatic generated by OneX, you should not set this field.
	// +optional
	BootstrapAccount *string `json:"bootstrapAccount,omitempty" protobuf:"bytes,3,opt,name=bootstrapAccount"`
}

// ChainStatus defines the observed state of Chain.
type ChainStatus struct {
	// +optional
	ConfigMapRef *LocalObjectReference `json:"configMapRef,omitempty"`

	// +optional
	MinerRef *LocalObjectReference `json:"minerRef,omitempty"`

	// ObservedGeneration is the latest generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions defines the current state of the Chain
	// +optional
	Conditions Conditions `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChainList is a list of Chain objects.
type ChainList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is a list of schema objects.
	Items []Chain `json:"items" protobuf:"bytes,2,rep,name=items"`
}
