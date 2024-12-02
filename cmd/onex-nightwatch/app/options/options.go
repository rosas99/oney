// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

// Package options contains flags and options for initializing an nightwatch.
package options

import (
	"math"

	"github.com/spf13/viper"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/tools/clientcmd"
	cliflag "k8s.io/component-base/cli/flag"

	"github.com/rosas/onex/internal/nightwatch"
	"github.com/rosas/onex/internal/pkg/feature"
	kubeutil "github.com/rosas/onex/internal/pkg/util/kube"
	"github.com/rosas/onex/pkg/app"
	clientset "github.com/rosas/onex/pkg/generated/clientset/versioned"
	"github.com/rosas/onex/pkg/log"
	genericoptions "github.com/rosas/onex/pkg/options"
	"github.com/rosas/onex/pkg/watch"
)

const (
	// UserAgent is the userAgent name when starting onex-nightwatch server.
	UserAgent = "onex-nightwatch"
)

var _ app.CliOptions = (*Options)(nil)

// Options contains everything necessary to create and run a nightwatch server.
type Options struct {
	HealthOptions         *genericoptions.HealthOptions  `json:"health" mapstructure:"health"`
	MySQLOptions          *genericoptions.MySQLOptions   `json:"mysql" mapstructure:"mysql"`
	RedisOptions          *genericoptions.RedisOptions   `json:"redis" mapstructure:"redis"`
	WatchOptions          *watch.Options                 `json:"nightwatch" mapstructure:"nightwatch"`
	HTTPOptions           *genericoptions.HTTPOptions    `json:"http" mapstructure:"http"`
	TLSOptions            *genericoptions.TLSOptions     `json:"tls" mapstructure:"tls"`
	UserWatcherMaxWorkers int64                          `json:"user-watcher-max-workers" mapstructure:"user-watcher-max-workers"`
	DisableRESTServer     bool                           `json:"disable-rest-server" mapstructure:"disable-rest-server"`
	Metrics               *genericoptions.MetricsOptions `json:"metrics" mapstructure:"metrics"`
	// Path to kubeconfig file with authorization and master location information.
	Kubeconfig   string          `json:"kubeconfig" mapstructure:"kubeconfig"`
	FeatureGates map[string]bool `json:"feature-gates"`
	Log          *log.Options    `json:"log" mapstructure:"log"`
}

// NewOptions returns initialized Options.
func NewOptions() *Options {
	o := &Options{
		HealthOptions:         genericoptions.NewHealthOptions(),
		MySQLOptions:          genericoptions.NewMySQLOptions(),
		RedisOptions:          genericoptions.NewRedisOptions(),
		HTTPOptions:           genericoptions.NewHTTPOptions(),
		TLSOptions:            genericoptions.NewTLSOptions(),
		DisableRESTServer:     false,
		UserWatcherMaxWorkers: math.MaxInt64,
		WatchOptions:          watch.NewOptions(),
		Metrics:               genericoptions.NewMetricsOptions(),
		Log:                   log.NewOptions(),
	}

	return o
}

// Flags returns flags for a specific server by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.HealthOptions.AddFlags(fss.FlagSet("health"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.HTTPOptions.AddFlags(fss.FlagSet("http"))
	o.TLSOptions.AddFlags(fss.FlagSet("tls"))
	o.WatchOptions.AddFlags(fss.FlagSet("watch"))
	o.Metrics.AddFlags(fss.FlagSet("metrics"))
	o.Log.AddFlags(fss.FlagSet("log"))

	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs := fss.FlagSet("misc")
	fs.StringVar(&o.Kubeconfig, "kubeconfig", o.Kubeconfig, "Path to kubeconfig file with authorization and master location information.")
	fs.BoolVar(&o.DisableRESTServer, "disable-rest-server", o.DisableRESTServer, "Disable the REST server functionality.")
	fs.Int64Var(&o.UserWatcherMaxWorkers, "user-watcher-max-workers", o.UserWatcherMaxWorkers, "Specify the maximum concurrency event of user watcher.")
	feature.DefaultMutableFeatureGate.AddFlag(fs)

	return fss
}

// Complete completes all the required options.
func (o *Options) Complete() error {
	if err := viper.Unmarshal(&o); err != nil {
		return err
	}

	if o.UserWatcherMaxWorkers < 1 {
		o.UserWatcherMaxWorkers = math.MaxInt64
	}

	_ = feature.DefaultMutableFeatureGate.SetFromMap(o.FeatureGates)
	return nil
}

// Validate validates all the required options.
func (o *Options) Validate() error {
	errs := []error{}

	errs = append(errs, o.HealthOptions.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.RedisOptions.Validate()...)
	errs = append(errs, o.HTTPOptions.Validate()...)
	errs = append(errs, o.TLSOptions.Validate()...)
	errs = append(errs, o.WatchOptions.Validate()...)
	errs = append(errs, o.Metrics.Validate()...)
	errs = append(errs, o.Log.Validate()...)

	return utilerrors.NewAggregate(errs)
}

// ApplyTo fills up onex-nightwatch config with options.
func (o *Options) ApplyTo(c *nightwatch.Config) error {
	c.HealthOptions = o.HealthOptions
	c.MySQLOptions = o.MySQLOptions
	c.RedisOptions = o.RedisOptions
	c.HTTPOptions = o.HTTPOptions
	c.TLSOptions = o.TLSOptions
	c.WatchOptions = o.WatchOptions
	c.DisableRESTServer = o.DisableRESTServer
	c.UserWatcherMaxWorkers = o.UserWatcherMaxWorkers
	return nil
}

// Config return an onex-nightwatch config object.
func (o *Options) Config() (*nightwatch.Config, error) {
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", o.Kubeconfig)
	if err != nil {
		return nil, err
	}
	kubeutil.SetDefaultClientOptions(kubeutil.AddUserAgent(kubeconfig, UserAgent))

	client, err := clientset.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}

	c := &nightwatch.Config{
		Client: client,
	}

	if err := o.ApplyTo(c); err != nil {
		return nil, err
	}

	return c, nil
}
