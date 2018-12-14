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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	deploymentv1alpha1 "foremast.ai/foremast/foremast-barrelman/pkg/apis/deployment/v1alpha1"
	versioned "foremast.ai/foremast/foremast-barrelman/pkg/client/clientset/versioned"
	internalinterfaces "foremast.ai/foremast/foremast-barrelman/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "foremast.ai/foremast/foremast-barrelman/pkg/client/listers/deployment/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DeploymentMetadataInformer provides access to a shared informer and lister for
// DeploymentMetadatas.
type DeploymentMetadataInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DeploymentMetadataLister
}

type deploymentMetadataInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDeploymentMetadataInformer constructs a new informer for DeploymentMetadata type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeploymentMetadataInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDeploymentMetadataInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDeploymentMetadataInformer constructs a new informer for DeploymentMetadata type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDeploymentMetadataInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DeploymentV1alpha1().DeploymentMetadatas(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DeploymentV1alpha1().DeploymentMetadatas(namespace).Watch(options)
			},
		},
		&deploymentv1alpha1.DeploymentMetadata{},
		resyncPeriod,
		indexers,
	)
}

func (f *deploymentMetadataInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDeploymentMetadataInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *deploymentMetadataInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&deploymentv1alpha1.DeploymentMetadata{}, f.defaultInformer)
}

func (f *deploymentMetadataInformer) Lister() v1alpha1.DeploymentMetadataLister {
	return v1alpha1.NewDeploymentMetadataLister(f.Informer().GetIndexer())
}
