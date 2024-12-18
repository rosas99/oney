// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	flowcontrolv1 "github.com/rosas/onex/pkg/generated/applyconfigurations/flowcontrol/v1"
	v1 "k8s.io/api/flowcontrol/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePriorityLevelConfigurations implements PriorityLevelConfigurationInterface
type FakePriorityLevelConfigurations struct {
	Fake *FakeFlowcontrolV1
}

var prioritylevelconfigurationsResource = v1.SchemeGroupVersion.WithResource("prioritylevelconfigurations")

var prioritylevelconfigurationsKind = v1.SchemeGroupVersion.WithKind("PriorityLevelConfiguration")

// Get takes name of the priorityLevelConfiguration, and returns the corresponding priorityLevelConfiguration object, and an error if there is any.
func (c *FakePriorityLevelConfigurations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(prioritylevelconfigurationsResource, name), &v1.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PriorityLevelConfiguration), err
}

// List takes label and field selectors, and returns the list of PriorityLevelConfigurations that match those selectors.
func (c *FakePriorityLevelConfigurations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PriorityLevelConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(prioritylevelconfigurationsResource, prioritylevelconfigurationsKind, opts), &v1.PriorityLevelConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.PriorityLevelConfigurationList{ListMeta: obj.(*v1.PriorityLevelConfigurationList).ListMeta}
	for _, item := range obj.(*v1.PriorityLevelConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested priorityLevelConfigurations.
func (c *FakePriorityLevelConfigurations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(prioritylevelconfigurationsResource, opts))
}

// Create takes the representation of a priorityLevelConfiguration and creates it.  Returns the server's representation of the priorityLevelConfiguration, and an error, if there is any.
func (c *FakePriorityLevelConfigurations) Create(ctx context.Context, priorityLevelConfiguration *v1.PriorityLevelConfiguration, opts metav1.CreateOptions) (result *v1.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(prioritylevelconfigurationsResource, priorityLevelConfiguration), &v1.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PriorityLevelConfiguration), err
}

// Update takes the representation of a priorityLevelConfiguration and updates it. Returns the server's representation of the priorityLevelConfiguration, and an error, if there is any.
func (c *FakePriorityLevelConfigurations) Update(ctx context.Context, priorityLevelConfiguration *v1.PriorityLevelConfiguration, opts metav1.UpdateOptions) (result *v1.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(prioritylevelconfigurationsResource, priorityLevelConfiguration), &v1.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PriorityLevelConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePriorityLevelConfigurations) UpdateStatus(ctx context.Context, priorityLevelConfiguration *v1.PriorityLevelConfiguration, opts metav1.UpdateOptions) (*v1.PriorityLevelConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(prioritylevelconfigurationsResource, "status", priorityLevelConfiguration), &v1.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PriorityLevelConfiguration), err
}

// Delete takes name of the priorityLevelConfiguration and deletes it. Returns an error if one occurs.
func (c *FakePriorityLevelConfigurations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(prioritylevelconfigurationsResource, name, opts), &v1.PriorityLevelConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePriorityLevelConfigurations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(prioritylevelconfigurationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.PriorityLevelConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched priorityLevelConfiguration.
func (c *FakePriorityLevelConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(prioritylevelconfigurationsResource, name, pt, data, subresources...), &v1.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PriorityLevelConfiguration), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied priorityLevelConfiguration.
func (c *FakePriorityLevelConfigurations) Apply(ctx context.Context, priorityLevelConfiguration *flowcontrolv1.PriorityLevelConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PriorityLevelConfiguration, err error) {
	if priorityLevelConfiguration == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(priorityLevelConfiguration)
	if err != nil {
		return nil, err
	}
	name := priorityLevelConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(prioritylevelconfigurationsResource, *name, types.ApplyPatchType, data), &v1.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PriorityLevelConfiguration), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePriorityLevelConfigurations) ApplyStatus(ctx context.Context, priorityLevelConfiguration *flowcontrolv1.PriorityLevelConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PriorityLevelConfiguration, err error) {
	if priorityLevelConfiguration == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(priorityLevelConfiguration)
	if err != nil {
		return nil, err
	}
	name := priorityLevelConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(prioritylevelconfigurationsResource, *name, types.ApplyPatchType, data, "status"), &v1.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PriorityLevelConfiguration), err
}
