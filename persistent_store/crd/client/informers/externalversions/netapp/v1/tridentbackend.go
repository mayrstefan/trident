// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	netappv1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	versioned "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned"
	internalinterfaces "github.com/netapp/trident/persistent_store/crd/client/informers/externalversions/internalinterfaces"
	v1 "github.com/netapp/trident/persistent_store/crd/client/listers/netapp/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TridentBackendInformer provides access to a shared informer and lister for
// TridentBackends.
type TridentBackendInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TridentBackendLister
}

type tridentBackendInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTridentBackendInformer constructs a new informer for TridentBackend type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTridentBackendInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTridentBackendInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTridentBackendInformer constructs a new informer for TridentBackend type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTridentBackendInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TridentV1().TridentBackends(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TridentV1().TridentBackends(namespace).Watch(context.TODO(), options)
			},
		},
		&netappv1.TridentBackend{},
		resyncPeriod,
		indexers,
	)
}

func (f *tridentBackendInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTridentBackendInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tridentBackendInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&netappv1.TridentBackend{}, f.defaultInformer)
}

func (f *tridentBackendInformer) Lister() v1.TridentBackendLister {
	return v1.NewTridentBackendLister(f.Informer().GetIndexer())
}
