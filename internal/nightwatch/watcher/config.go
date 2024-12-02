// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

// Package watcher provides functions used by all watchers.
package watcher

import (
	"github.com/rosas/onex/internal/nightwatch/store"
	"github.com/rosas/onex/internal/pkg/client/minio"
	aggregatestore "github.com/rosas/onex/internal/pkg/client/store"
	clientset "github.com/rosas/onex/pkg/generated/clientset/versioned"
)

// AggregateConfig aggregates the configurations of all watchers and serves as a configuration aggregator.
type AggregateConfig struct {
	Minio minio.IMinio
	Store store.IStore
	// The purpose of nightwatch is to handle asynchronous tasks on the onex platform
	// in a unified manner, so a store aggregation type is needed here.
	AggregateStore aggregatestore.Interface

	// Client is the client for onex-apiserver.
	Client clientset.Interface

	// Then maximum concurrency event of user watcher.
	UserWatcherMaxWorkers int64
}
