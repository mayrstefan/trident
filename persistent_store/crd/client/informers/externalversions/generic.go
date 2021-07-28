// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=trident.netapp.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("tridentbackends"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Trident().V1().TridentBackends().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("tridentbackendconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Trident().V1().TridentBackendConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("tridentmirrorrelationships"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Trident().V1().TridentMirrorRelationships().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("tridentnodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Trident().V1().TridentNodes().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("tridentsnapshots"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Trident().V1().TridentSnapshots().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("tridentsnapshotinfos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Trident().V1().TridentSnapshotInfos().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("tridentstorageclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Trident().V1().TridentStorageClasses().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("tridenttransactions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Trident().V1().TridentTransactions().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("tridentversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Trident().V1().TridentVersions().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("tridentvolumes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Trident().V1().TridentVolumes().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
