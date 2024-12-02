// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// FCVolumeSourceApplyConfiguration represents an declarative configuration of the FCVolumeSource type for use
// with apply.
type FCVolumeSourceApplyConfiguration struct {
	TargetWWNs []string `json:"targetWWNs,omitempty"`
	Lun        *int32   `json:"lun,omitempty"`
	FSType     *string  `json:"fsType,omitempty"`
	ReadOnly   *bool    `json:"readOnly,omitempty"`
	WWIDs      []string `json:"wwids,omitempty"`
}

// FCVolumeSourceApplyConfiguration constructs an declarative configuration of the FCVolumeSource type for use with
// apply.
func FCVolumeSource() *FCVolumeSourceApplyConfiguration {
	return &FCVolumeSourceApplyConfiguration{}
}

// WithTargetWWNs adds the given value to the TargetWWNs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TargetWWNs field.
func (b *FCVolumeSourceApplyConfiguration) WithTargetWWNs(values ...string) *FCVolumeSourceApplyConfiguration {
	for i := range values {
		b.TargetWWNs = append(b.TargetWWNs, values[i])
	}
	return b
}

// WithLun sets the Lun field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Lun field is set to the value of the last call.
func (b *FCVolumeSourceApplyConfiguration) WithLun(value int32) *FCVolumeSourceApplyConfiguration {
	b.Lun = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *FCVolumeSourceApplyConfiguration) WithFSType(value string) *FCVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *FCVolumeSourceApplyConfiguration) WithReadOnly(value bool) *FCVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithWWIDs adds the given value to the WWIDs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the WWIDs field.
func (b *FCVolumeSourceApplyConfiguration) WithWWIDs(values ...string) *FCVolumeSourceApplyConfiguration {
	for i := range values {
		b.WWIDs = append(b.WWIDs, values[i])
	}
	return b
}
