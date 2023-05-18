// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned/typed/netapp/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTridentV1 struct {
	*testing.Fake
}

func (c *FakeTridentV1) TridentActionMirrorUpdates(namespace string) v1.TridentActionMirrorUpdateInterface {
	return &FakeTridentActionMirrorUpdates{c, namespace}
}

func (c *FakeTridentV1) TridentBackends(namespace string) v1.TridentBackendInterface {
	return &FakeTridentBackends{c, namespace}
}

func (c *FakeTridentV1) TridentBackendConfigs(namespace string) v1.TridentBackendConfigInterface {
	return &FakeTridentBackendConfigs{c, namespace}
}

func (c *FakeTridentV1) TridentMirrorRelationships(namespace string) v1.TridentMirrorRelationshipInterface {
	return &FakeTridentMirrorRelationships{c, namespace}
}

func (c *FakeTridentV1) TridentNodes(namespace string) v1.TridentNodeInterface {
	return &FakeTridentNodes{c, namespace}
}

func (c *FakeTridentV1) TridentSnapshots(namespace string) v1.TridentSnapshotInterface {
	return &FakeTridentSnapshots{c, namespace}
}

func (c *FakeTridentV1) TridentSnapshotInfos(namespace string) v1.TridentSnapshotInfoInterface {
	return &FakeTridentSnapshotInfos{c, namespace}
}

func (c *FakeTridentV1) TridentStorageClasses(namespace string) v1.TridentStorageClassInterface {
	return &FakeTridentStorageClasses{c, namespace}
}

func (c *FakeTridentV1) TridentTransactions(namespace string) v1.TridentTransactionInterface {
	return &FakeTridentTransactions{c, namespace}
}

func (c *FakeTridentV1) TridentVersions(namespace string) v1.TridentVersionInterface {
	return &FakeTridentVersions{c, namespace}
}

func (c *FakeTridentV1) TridentVolumes(namespace string) v1.TridentVolumeInterface {
	return &FakeTridentVolumes{c, namespace}
}

func (c *FakeTridentV1) TridentVolumePublications(namespace string) v1.TridentVolumePublicationInterface {
	return &FakeTridentVolumePublications{c, namespace}
}

func (c *FakeTridentV1) TridentVolumeReferences(namespace string) v1.TridentVolumeReferenceInterface {
	return &FakeTridentVolumeReferences{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTridentV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
