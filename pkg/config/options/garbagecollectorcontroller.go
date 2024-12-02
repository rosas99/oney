// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

package options

import (
	"github.com/spf13/pflag"

	genericconfig "github.com/rosas/onex/pkg/config"
)

// GarbageCollectorControllerOptions holds the GarbageCollectorController options.
type GarbageCollectorControllerOptions struct {
	*genericconfig.GarbageCollectorControllerConfiguration
}

func NewGarbageCollectorControllerOptions(cfg *genericconfig.GarbageCollectorControllerConfiguration) *GarbageCollectorControllerOptions {
	return &GarbageCollectorControllerOptions{
		GarbageCollectorControllerConfiguration: cfg,
	}
}

// AddFlags adds flags related to GarbageCollectorController for controller manager to the specified FlagSet.
func (o *GarbageCollectorControllerOptions) AddFlags(fs *pflag.FlagSet) {
	if o == nil {
		return
	}

	fs.Int32Var(&o.ConcurrentGCSyncs, "concurrent-gc-syncs", o.ConcurrentGCSyncs, "The number of garbage collector workers that are allowed to sync concurrently."+
		"This parameter is ignored if a config file is specified by --config.")
	fs.BoolVar(&o.EnableGarbageCollector, "enable-garbage-collector", o.EnableGarbageCollector, "Enables the generic garbage collector. MUST be synced with "+
		"the corresponding flag of the kube-apiserver. This parameter is ignored if a config file is specified by --config.")
}

// ApplyTo fills up GarbageCollectorController config with options.
func (o *GarbageCollectorControllerOptions) ApplyTo(cfg *genericconfig.GarbageCollectorControllerConfiguration) error {
	if o == nil || cfg == nil {
		return nil
	}

	*cfg = *o.GarbageCollectorControllerConfiguration
	return nil
}

// Validate checks validation of GarbageCollectorController.
func (o *GarbageCollectorControllerOptions) Validate() []error {
	if o == nil {
		return nil
	}

	errs := []error{}
	return errs
}
