// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by informer-gen. DO NOT EDIT.

package internalinterfaces

import (
	time "time"

	versioned "github.com/rosas/onex/pkg/generated/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	cache "k8s.io/client-go/tools/cache"
)

// NewInformerFunc takes versioned.Interface and time.Duration to return a SharedIndexInformer.
type NewInformerFunc func(versioned.Interface, time.Duration) cache.SharedIndexInformer

// SharedInformerFactory a small interface to allow for adding an informer without an import cycle
type SharedInformerFactory interface {
	Start(stopCh <-chan struct{})
	InformerFor(obj runtime.Object, newFunc NewInformerFunc) cache.SharedIndexInformer
}

// TweakListOptionsFunc is a function that transforms a v1.ListOptions.
type TweakListOptionsFunc func(*v1.ListOptions)
