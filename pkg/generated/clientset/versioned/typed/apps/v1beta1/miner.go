// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1beta1 "github.com/rosas/onex/pkg/apis/apps/v1beta1"
	appsv1beta1 "github.com/rosas/onex/pkg/generated/applyconfigurations/apps/v1beta1"
	scheme "github.com/rosas/onex/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MinersGetter has a method to return a MinerInterface.
// A group's client should implement this interface.
type MinersGetter interface {
	Miners(namespace string) MinerInterface
}

// MinerInterface has methods to work with Miner resources.
type MinerInterface interface {
	Create(ctx context.Context, miner *v1beta1.Miner, opts v1.CreateOptions) (*v1beta1.Miner, error)
	Update(ctx context.Context, miner *v1beta1.Miner, opts v1.UpdateOptions) (*v1beta1.Miner, error)
	UpdateStatus(ctx context.Context, miner *v1beta1.Miner, opts v1.UpdateOptions) (*v1beta1.Miner, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Miner, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.MinerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Miner, err error)
	Apply(ctx context.Context, miner *appsv1beta1.MinerApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Miner, err error)
	ApplyStatus(ctx context.Context, miner *appsv1beta1.MinerApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Miner, err error)
	MinerExpansion
}

// miners implements MinerInterface
type miners struct {
	client rest.Interface
	ns     string
}

// newMiners returns a Miners
func newMiners(c *AppsV1beta1Client, namespace string) *miners {
	return &miners{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the miner, and returns the corresponding miner object, and an error if there is any.
func (c *miners) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Miner, err error) {
	result = &v1beta1.Miner{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("miners").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Miners that match those selectors.
func (c *miners) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MinerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.MinerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("miners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested miners.
func (c *miners) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("miners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a miner and creates it.  Returns the server's representation of the miner, and an error, if there is any.
func (c *miners) Create(ctx context.Context, miner *v1beta1.Miner, opts v1.CreateOptions) (result *v1beta1.Miner, err error) {
	result = &v1beta1.Miner{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("miners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(miner).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a miner and updates it. Returns the server's representation of the miner, and an error, if there is any.
func (c *miners) Update(ctx context.Context, miner *v1beta1.Miner, opts v1.UpdateOptions) (result *v1beta1.Miner, err error) {
	result = &v1beta1.Miner{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("miners").
		Name(miner.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(miner).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *miners) UpdateStatus(ctx context.Context, miner *v1beta1.Miner, opts v1.UpdateOptions) (result *v1beta1.Miner, err error) {
	result = &v1beta1.Miner{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("miners").
		Name(miner.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(miner).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the miner and deletes it. Returns an error if one occurs.
func (c *miners) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("miners").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *miners) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("miners").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched miner.
func (c *miners) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Miner, err error) {
	result = &v1beta1.Miner{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("miners").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied miner.
func (c *miners) Apply(ctx context.Context, miner *appsv1beta1.MinerApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Miner, err error) {
	if miner == nil {
		return nil, fmt.Errorf("miner provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(miner)
	if err != nil {
		return nil, err
	}
	name := miner.Name
	if name == nil {
		return nil, fmt.Errorf("miner.Name must be provided to Apply")
	}
	result = &v1beta1.Miner{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("miners").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *miners) ApplyStatus(ctx context.Context, miner *appsv1beta1.MinerApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Miner, err error) {
	if miner == nil {
		return nil, fmt.Errorf("miner provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(miner)
	if err != nil {
		return nil, err
	}

	name := miner.Name
	if name == nil {
		return nil, fmt.Errorf("miner.Name must be provided to Apply")
	}

	result = &v1beta1.Miner{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("miners").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
