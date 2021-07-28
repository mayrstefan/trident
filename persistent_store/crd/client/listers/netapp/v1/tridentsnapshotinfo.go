// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TridentSnapshotInfoLister helps list TridentSnapshotInfos.
type TridentSnapshotInfoLister interface {
	// List lists all TridentSnapshotInfos in the indexer.
	List(selector labels.Selector) (ret []*v1.TridentSnapshotInfo, err error)
	// TridentSnapshotInfos returns an object that can list and get TridentSnapshotInfos.
	TridentSnapshotInfos(namespace string) TridentSnapshotInfoNamespaceLister
	TridentSnapshotInfoListerExpansion
}

// tridentSnapshotInfoLister implements the TridentSnapshotInfoLister interface.
type tridentSnapshotInfoLister struct {
	indexer cache.Indexer
}

// NewTridentSnapshotInfoLister returns a new TridentSnapshotInfoLister.
func NewTridentSnapshotInfoLister(indexer cache.Indexer) TridentSnapshotInfoLister {
	return &tridentSnapshotInfoLister{indexer: indexer}
}

// List lists all TridentSnapshotInfos in the indexer.
func (s *tridentSnapshotInfoLister) List(selector labels.Selector) (ret []*v1.TridentSnapshotInfo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentSnapshotInfo))
	})
	return ret, err
}

// TridentSnapshotInfos returns an object that can list and get TridentSnapshotInfos.
func (s *tridentSnapshotInfoLister) TridentSnapshotInfos(namespace string) TridentSnapshotInfoNamespaceLister {
	return tridentSnapshotInfoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TridentSnapshotInfoNamespaceLister helps list and get TridentSnapshotInfos.
type TridentSnapshotInfoNamespaceLister interface {
	// List lists all TridentSnapshotInfos in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TridentSnapshotInfo, err error)
	// Get retrieves the TridentSnapshotInfo from the indexer for a given namespace and name.
	Get(name string) (*v1.TridentSnapshotInfo, error)
	TridentSnapshotInfoNamespaceListerExpansion
}

// tridentSnapshotInfoNamespaceLister implements the TridentSnapshotInfoNamespaceLister
// interface.
type tridentSnapshotInfoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TridentSnapshotInfos in the indexer for a given namespace.
func (s tridentSnapshotInfoNamespaceLister) List(selector labels.Selector) (ret []*v1.TridentSnapshotInfo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentSnapshotInfo))
	})
	return ret, err
}

// Get retrieves the TridentSnapshotInfo from the indexer for a given namespace and name.
func (s tridentSnapshotInfoNamespaceLister) Get(name string) (*v1.TridentSnapshotInfo, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tridentsnapshotinfo"), name)
	}
	return obj.(*v1.TridentSnapshotInfo), nil
}
