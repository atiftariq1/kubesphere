/*
Copyright 2020 The KubeSphere Authors.

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

package v1beta2

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	typesv1beta2 "kubesphere.io/api/types/v1beta2"
	versioned "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	internalinterfaces "kubesphere.io/kubesphere/pkg/client/informers/externalversions/internalinterfaces"
	v1beta2 "kubesphere.io/kubesphere/pkg/client/listers/types/v1beta2"
)

// FederatedNotificationConfigInformer provides access to a shared informer and lister for
// FederatedNotificationConfigs.
type FederatedNotificationConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta2.FederatedNotificationConfigLister
}

type federatedNotificationConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewFederatedNotificationConfigInformer constructs a new informer for FederatedNotificationConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedNotificationConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedNotificationConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedNotificationConfigInformer constructs a new informer for FederatedNotificationConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedNotificationConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TypesV1beta2().FederatedNotificationConfigs().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TypesV1beta2().FederatedNotificationConfigs().Watch(context.TODO(), options)
			},
		},
		&typesv1beta2.FederatedNotificationConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedNotificationConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedNotificationConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedNotificationConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&typesv1beta2.FederatedNotificationConfig{}, f.defaultInformer)
}

func (f *federatedNotificationConfigInformer) Lister() v1beta2.FederatedNotificationConfigLister {
	return v1beta2.NewFederatedNotificationConfigLister(f.Informer().GetIndexer())
}