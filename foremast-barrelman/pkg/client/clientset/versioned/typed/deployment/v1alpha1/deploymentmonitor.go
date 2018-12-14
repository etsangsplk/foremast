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

package v1alpha1

import (
	v1alpha1 "foremast.ai/foremast/foremast-barrelman/pkg/apis/deployment/v1alpha1"
	scheme "foremast.ai/foremast/foremast-barrelman/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeploymentMonitorsGetter has a method to return a DeploymentMonitorInterface.
// A group's client should implement this interface.
type DeploymentMonitorsGetter interface {
	DeploymentMonitors(namespace string) DeploymentMonitorInterface
}

// DeploymentMonitorInterface has methods to work with DeploymentMonitor resources.
type DeploymentMonitorInterface interface {
	Create(*v1alpha1.DeploymentMonitor) (*v1alpha1.DeploymentMonitor, error)
	Update(*v1alpha1.DeploymentMonitor) (*v1alpha1.DeploymentMonitor, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DeploymentMonitor, error)
	List(opts v1.ListOptions) (*v1alpha1.DeploymentMonitorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeploymentMonitor, err error)
	DeploymentMonitorExpansion
}

// deploymentMonitors implements DeploymentMonitorInterface
type deploymentMonitors struct {
	client rest.Interface
	ns     string
}

// newDeploymentMonitors returns a DeploymentMonitors
func newDeploymentMonitors(c *DeploymentV1alpha1Client, namespace string) *deploymentMonitors {
	return &deploymentMonitors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deploymentMonitor, and returns the corresponding deploymentMonitor object, and an error if there is any.
func (c *deploymentMonitors) Get(name string, options v1.GetOptions) (result *v1alpha1.DeploymentMonitor, err error) {
	result = &v1alpha1.DeploymentMonitor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymentmonitors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeploymentMonitors that match those selectors.
func (c *deploymentMonitors) List(opts v1.ListOptions) (result *v1alpha1.DeploymentMonitorList, err error) {
	result = &v1alpha1.DeploymentMonitorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymentmonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deploymentMonitors.
func (c *deploymentMonitors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deploymentmonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a deploymentMonitor and creates it.  Returns the server's representation of the deploymentMonitor, and an error, if there is any.
func (c *deploymentMonitors) Create(deploymentMonitor *v1alpha1.DeploymentMonitor) (result *v1alpha1.DeploymentMonitor, err error) {
	result = &v1alpha1.DeploymentMonitor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deploymentmonitors").
		Body(deploymentMonitor).
		Do().
		Into(result)
	return
}

// Update takes the representation of a deploymentMonitor and updates it. Returns the server's representation of the deploymentMonitor, and an error, if there is any.
func (c *deploymentMonitors) Update(deploymentMonitor *v1alpha1.DeploymentMonitor) (result *v1alpha1.DeploymentMonitor, err error) {
	result = &v1alpha1.DeploymentMonitor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deploymentmonitors").
		Name(deploymentMonitor.Name).
		Body(deploymentMonitor).
		Do().
		Into(result)
	return
}

// Delete takes name of the deploymentMonitor and deletes it. Returns an error if one occurs.
func (c *deploymentMonitors) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymentmonitors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deploymentMonitors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymentmonitors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched deploymentMonitor.
func (c *deploymentMonitors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeploymentMonitor, err error) {
	result = &v1alpha1.DeploymentMonitor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deploymentmonitors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
