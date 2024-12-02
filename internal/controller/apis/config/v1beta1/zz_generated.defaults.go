//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by defaulter-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&OneXControllerManagerConfiguration{}, func(obj interface{}) {
		SetObjectDefaults_OneXControllerManagerConfiguration(obj.(*OneXControllerManagerConfiguration))
	})
	return nil
}

func SetObjectDefaults_OneXControllerManagerConfiguration(in *OneXControllerManagerConfiguration) {
	SetDefaults_OneXControllerManagerConfiguration(in)
}
