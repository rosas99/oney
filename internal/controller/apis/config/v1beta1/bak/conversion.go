// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

package v1beta1

import (
	"k8s.io/apimachinery/pkg/conversion"

	"github.com/rosas/onex/pkg/config"
	configv1beta1 "github.com/rosas/onex/pkg/config/v1beta1"
)

// Important! The public back-and-forth conversion functions for the types in this generic
// package with ComponentConfig types need to be manually exposed like this in order for
// other packages that reference this package to be able to call these conversion functions
// in an autogenerated manner.
// TODO: Fix the bug in conversion-gen so it automatically discovers these Convert_* functions
// in autogenerated code as well.

func Convert_v1beta1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(in *configv1beta1.GarbageCollectorControllerConfiguration, out *config.GarbageCollectorControllerConfiguration, s conversion.Scope) error {
	return configv1beta1.Convert_v1beta1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(in, out, s)
}

func Convert_config_GarbageCollectorControllerConfiguration_To_v1beta1_GarbageCollectorControllerConfiguration(in *config.GarbageCollectorControllerConfiguration, out *configv1beta1.GarbageCollectorControllerConfiguration, s conversion.Scope) error {
	return configv1beta1.Convert_config_GarbageCollectorControllerConfiguration_To_v1beta1_GarbageCollectorControllerConfiguration(in, out, s)
}

func Convert_v1beta1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(in *configv1beta1.GenericControllerManagerConfiguration, out *config.GenericControllerManagerConfiguration, s conversion.Scope) error {
	return configv1beta1.Convert_v1beta1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(in, out, s)
}

func Convert_config_GenericControllerManagerConfiguration_To_v1beta1_GenericControllerManagerConfiguration(in *config.GenericControllerManagerConfiguration, out *GenericControllerManagerConfiguration, s conversion.Scope) error {
	return configv1beta1.Convert_config_GenericControllerManagerConfiguration_To_v1beta1_GenericControllerManagerConfiguration(in, out, s)
}

func Convert_v1beta1_MySQLConfiguration_To_config_MySQLConfiguration(in *configv1beta1.MySQLConfiguration, out *config.MySQLConfiguration, s conversion.Scope) error {
	return configv1beta1.Convert_v1beta1_MySQLConfiguration_To_config_MySQLConfiguration(in, out, s)
}

func Convert_config_MySQLConfiguration_To_v1beta1_MySQLConfiguration(in *config.MySQLConfiguration, out *configv1beta1.MySQLConfiguration, s conversion.Scope) error {
	return configv1beta1.Convert_config_MySQLConfiguration_To_v1beta1_MySQLConfiguration(in, out, s)
}
func Convert_v1beta1_RedisConfiguration_To_config_RedisConfiguration(in *configv1beta1.RedisConfiguration, out *config.RedisConfiguration, s conversion.Scope) error {
	return configv1beta1.Convert_v1beta1_RedisConfiguration_To_config_RedisConfiguration(in, out, s)
}
func Convert_config_RedisConfiguration_To_v1beta1_RedisConfiguration(in *config.RedisConfiguration, out *configv1beta1.RedisConfiguration, s conversion.Scope) error {
	return configv1beta1.Convert_config_RedisConfiguration_To_v1beta1_RedisConfiguration(in, out, s)
}
