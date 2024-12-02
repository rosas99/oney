// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// WindowsSecurityContextOptionsApplyConfiguration represents an declarative configuration of the WindowsSecurityContextOptions type for use
// with apply.
type WindowsSecurityContextOptionsApplyConfiguration struct {
	GMSACredentialSpecName *string `json:"gmsaCredentialSpecName,omitempty"`
	GMSACredentialSpec     *string `json:"gmsaCredentialSpec,omitempty"`
	RunAsUserName          *string `json:"runAsUserName,omitempty"`
	HostProcess            *bool   `json:"hostProcess,omitempty"`
}

// WindowsSecurityContextOptionsApplyConfiguration constructs an declarative configuration of the WindowsSecurityContextOptions type for use with
// apply.
func WindowsSecurityContextOptions() *WindowsSecurityContextOptionsApplyConfiguration {
	return &WindowsSecurityContextOptionsApplyConfiguration{}
}

// WithGMSACredentialSpecName sets the GMSACredentialSpecName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GMSACredentialSpecName field is set to the value of the last call.
func (b *WindowsSecurityContextOptionsApplyConfiguration) WithGMSACredentialSpecName(value string) *WindowsSecurityContextOptionsApplyConfiguration {
	b.GMSACredentialSpecName = &value
	return b
}

// WithGMSACredentialSpec sets the GMSACredentialSpec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GMSACredentialSpec field is set to the value of the last call.
func (b *WindowsSecurityContextOptionsApplyConfiguration) WithGMSACredentialSpec(value string) *WindowsSecurityContextOptionsApplyConfiguration {
	b.GMSACredentialSpec = &value
	return b
}

// WithRunAsUserName sets the RunAsUserName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunAsUserName field is set to the value of the last call.
func (b *WindowsSecurityContextOptionsApplyConfiguration) WithRunAsUserName(value string) *WindowsSecurityContextOptionsApplyConfiguration {
	b.RunAsUserName = &value
	return b
}

// WithHostProcess sets the HostProcess field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostProcess field is set to the value of the last call.
func (b *WindowsSecurityContextOptionsApplyConfiguration) WithHostProcess(value bool) *WindowsSecurityContextOptionsApplyConfiguration {
	b.HostProcess = &value
	return b
}
