/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "foremast.ai/foremast/foremast-barrelman/pkg/apis/deployment/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeploymentMonitors implements DeploymentMonitorInterface
type FakeDeploymentMonitors struct {
	Fake *FakeDeploymentV1alpha1
	ns   string
}

var deploymentmonitorsResource = schema.GroupVersionResource{Group: "deployment.foremast.ai", Version: "v1alpha1", Resource: "deploymentmonitors"}

var deploymentmonitorsKind = schema.GroupVersionKind{Group: "deployment.foremast.ai", Version: "v1alpha1", Kind: "DeploymentMonitor"}

// Get takes name of the deploymentMonitor, and returns the corresponding deploymentMonitor object, and an error if there is any.
func (c *FakeDeploymentMonitors) Get(name string, options v1.GetOptions) (result *v1alpha1.DeploymentMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deploymentmonitorsResource, c.ns, name), &v1alpha1.DeploymentMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentMonitor), err
}

// List takes label and field selectors, and returns the list of DeploymentMonitors that match those selectors.
func (c *FakeDeploymentMonitors) List(opts v1.ListOptions) (result *v1alpha1.DeploymentMonitorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deploymentmonitorsResource, deploymentmonitorsKind, c.ns, opts), &v1alpha1.DeploymentMonitorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DeploymentMonitorList{ListMeta: obj.(*v1alpha1.DeploymentMonitorList).ListMeta}
	for _, item := range obj.(*v1alpha1.DeploymentMonitorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deploymentMonitors.
func (c *FakeDeploymentMonitors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deploymentmonitorsResource, c.ns, opts))

}

// Create takes the representation of a deploymentMonitor and creates it.  Returns the server's representation of the deploymentMonitor, and an error, if there is any.
func (c *FakeDeploymentMonitors) Create(deploymentMonitor *v1alpha1.DeploymentMonitor) (result *v1alpha1.DeploymentMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deploymentmonitorsResource, c.ns, deploymentMonitor), &v1alpha1.DeploymentMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentMonitor), err
}

// Update takes the representation of a deploymentMonitor and updates it. Returns the server's representation of the deploymentMonitor, and an error, if there is any.
func (c *FakeDeploymentMonitors) Update(deploymentMonitor *v1alpha1.DeploymentMonitor) (result *v1alpha1.DeploymentMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deploymentmonitorsResource, c.ns, deploymentMonitor), &v1alpha1.DeploymentMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentMonitor), err
}

// Delete takes name of the deploymentMonitor and deletes it. Returns an error if one occurs.
func (c *FakeDeploymentMonitors) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(deploymentmonitorsResource, c.ns, name), &v1alpha1.DeploymentMonitor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeploymentMonitors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deploymentmonitorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DeploymentMonitorList{})
	return err
}

// Patch applies the patch and returns the patched deploymentMonitor.
func (c *FakeDeploymentMonitors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeploymentMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentmonitorsResource, c.ns, name, data, subresources...), &v1alpha1.DeploymentMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentMonitor), err
}
