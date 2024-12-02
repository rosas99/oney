// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

package controller

import (
	"context"

	"github.com/redis/go-redis/v9"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/rosas/onex/internal/controller/apis/config"
	chaincontroller "github.com/rosas/onex/internal/controller/chain"
	minercontroller "github.com/rosas/onex/internal/controller/miner"
	minerconfig "github.com/rosas/onex/internal/controller/miner/apis/config"
	minersetcontroller "github.com/rosas/onex/internal/controller/minerset"
	synccontroller "github.com/rosas/onex/internal/controller/sync"
	"github.com/rosas/onex/internal/gateway/store"
)

// Following types provides access to reconcilers implemented in internal/controller, thus
// allowing users to provide a single binary "batteries included" with OneX and providers of choice.

// ChainReconciler reconciles a Chain object.
type ChainReconciler struct {
	ComponentConfig *config.ChainControllerConfiguration

	// WatchFilterValue is the label value used to filter events prior to reconciliation.
	WatchFilterValue string
}

func (r *ChainReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	return (&chaincontroller.Reconciler{
		ComponentConfig:  r.ComponentConfig,
		WatchFilterValue: r.WatchFilterValue,
	}).SetupWithManager(ctx, mgr, options)
}

// MinerReconciler reconciles a Miner object.
type MinerReconciler struct {
	DryRun          bool
	ProviderClient  kubernetes.Interface
	ProviderCluster cluster.Cluster
	RedisClient     *redis.Client
	ComponentConfig *minerconfig.MinerControllerConfiguration

	// WatchFilterValue is the label value used to filter events prior to reconciliation.
	WatchFilterValue string
}

func (r *MinerReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	return (&minercontroller.Reconciler{
		DryRun:           r.DryRun,
		ProviderClient:   r.ProviderClient,
		RedisClient:      r.RedisClient,
		ComponentConfig:  r.ComponentConfig,
		WatchFilterValue: r.WatchFilterValue,
	}).SetupWithManager(ctx, mgr, options, r.ProviderCluster)
}

// MinerSetReconciler reconciles a MinerSet object.
type MinerSetReconciler struct {
	// WatchFilterValue is the label value used to filter events prior to reconciliation.
	WatchFilterValue string
}

func (r *MinerSetReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	return (&minersetcontroller.Reconciler{
		WatchFilterValue: r.WatchFilterValue,
	}).SetupWithManager(ctx, mgr, options)
}

// SyncReconciler sync onex resource to database.
type SyncReconciler struct {
	Store store.IStore
}

func (r *SyncReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	// setup chainSync controller
	if err := (&synccontroller.ChainSyncReconciler{
		Store: r.Store,
	}).SetupWithManager(ctx, mgr, options); err != nil {
		return err
	}

	// setup minerSetSync controller
	if err := (&synccontroller.MinerSetSyncReconciler{
		Store: r.Store,
	}).SetupWithManager(ctx, mgr, options); err != nil {
		return err
	}

	// setup minerSync controller
	if err := (&synccontroller.MinerSyncReconciler{
		Store: r.Store,
	}).SetupWithManager(ctx, mgr, options); err != nil {
		return err
	}

	return nil
}

/*
// NamespacedResourcesDeleterReconciler is a reconciler used to delete a namespace with all resources in it.
type NamespacedResourcesDeleterReconciler struct {
	Client         clientset.Interface
	MetadataClient metadata.Interface
}

func (r *NamespacedResourcesDeleterReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	return namespacecontroller.NewNamespacedResourcesDeleter(mgr, r.Client, r.MetadataClient).SetupWithManager(mgr, options)
}
*/
