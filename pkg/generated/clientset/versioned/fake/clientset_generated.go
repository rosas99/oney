// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/rosas/onex/pkg/generated/clientset/versioned"
	apiextensionsv1 "github.com/rosas/onex/pkg/generated/clientset/versioned/typed/apiextensions/v1"
	fakeapiextensionsv1 "github.com/rosas/onex/pkg/generated/clientset/versioned/typed/apiextensions/v1/fake"
	appsv1beta1 "github.com/rosas/onex/pkg/generated/clientset/versioned/typed/apps/v1beta1"
	fakeappsv1beta1 "github.com/rosas/onex/pkg/generated/clientset/versioned/typed/apps/v1beta1/fake"
	coordinationv1 "github.com/rosas/onex/pkg/generated/clientset/versioned/typed/coordination/v1"
	fakecoordinationv1 "github.com/rosas/onex/pkg/generated/clientset/versioned/typed/coordination/v1/fake"
	corev1 "github.com/rosas/onex/pkg/generated/clientset/versioned/typed/core/v1"
	fakecorev1 "github.com/rosas/onex/pkg/generated/clientset/versioned/typed/core/v1/fake"
	flowcontrolv1 "github.com/rosas/onex/pkg/generated/clientset/versioned/typed/flowcontrol/v1"
	fakeflowcontrolv1 "github.com/rosas/onex/pkg/generated/clientset/versioned/typed/flowcontrol/v1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// ApiextensionsV1 retrieves the ApiextensionsV1Client
func (c *Clientset) ApiextensionsV1() apiextensionsv1.ApiextensionsV1Interface {
	return &fakeapiextensionsv1.FakeApiextensionsV1{Fake: &c.Fake}
}

// AppsV1beta1 retrieves the AppsV1beta1Client
func (c *Clientset) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	return &fakeappsv1beta1.FakeAppsV1beta1{Fake: &c.Fake}
}

// CoordinationV1 retrieves the CoordinationV1Client
func (c *Clientset) CoordinationV1() coordinationv1.CoordinationV1Interface {
	return &fakecoordinationv1.FakeCoordinationV1{Fake: &c.Fake}
}

// CoreV1 retrieves the CoreV1Client
func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return &fakecorev1.FakeCoreV1{Fake: &c.Fake}
}

// FlowcontrolV1 retrieves the FlowcontrolV1Client
func (c *Clientset) FlowcontrolV1() flowcontrolv1.FlowcontrolV1Interface {
	return &fakeflowcontrolv1.FakeFlowcontrolV1{Fake: &c.Fake}
}
