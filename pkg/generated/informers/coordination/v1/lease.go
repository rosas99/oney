// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	versioned "github.com/rosas/onex/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/rosas/onex/pkg/generated/informers/internalinterfaces"
	v1 "github.com/rosas/onex/pkg/generated/listers/coordination/v1"
	coordinationv1 "k8s.io/api/coordination/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LeaseInformer provides access to a shared informer and lister for
// Leases.
type LeaseInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.LeaseLister
}

type leaseInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLeaseInformer constructs a new informer for Lease type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLeaseInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLeaseInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLeaseInformer constructs a new informer for Lease type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLeaseInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoordinationV1().Leases(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoordinationV1().Leases(namespace).Watch(context.TODO(), options)
			},
		},
		&coordinationv1.Lease{},
		resyncPeriod,
		indexers,
	)
}

func (f *leaseInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLeaseInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *leaseInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&coordinationv1.Lease{}, f.defaultInformer)
}

func (f *leaseInformer) Lister() v1.LeaseLister {
	return v1.NewLeaseLister(f.Informer().GetIndexer())
}
