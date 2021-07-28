// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	netappv1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTridentSnapshotInfos implements TridentSnapshotInfoInterface
type FakeTridentSnapshotInfos struct {
	Fake *FakeTridentV1
	ns   string
}

var tridentsnapshotinfosResource = schema.GroupVersionResource{Group: "trident.netapp.io", Version: "v1", Resource: "tridentsnapshotinfos"}

var tridentsnapshotinfosKind = schema.GroupVersionKind{Group: "trident.netapp.io", Version: "v1", Kind: "TridentSnapshotInfo"}

// Get takes name of the tridentSnapshotInfo, and returns the corresponding tridentSnapshotInfo object, and an error if there is any.
func (c *FakeTridentSnapshotInfos) Get(ctx context.Context, name string, options v1.GetOptions) (result *netappv1.TridentSnapshotInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tridentsnapshotinfosResource, c.ns, name), &netappv1.TridentSnapshotInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentSnapshotInfo), err
}

// List takes label and field selectors, and returns the list of TridentSnapshotInfos that match those selectors.
func (c *FakeTridentSnapshotInfos) List(ctx context.Context, opts v1.ListOptions) (result *netappv1.TridentSnapshotInfoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tridentsnapshotinfosResource, tridentsnapshotinfosKind, c.ns, opts), &netappv1.TridentSnapshotInfoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &netappv1.TridentSnapshotInfoList{ListMeta: obj.(*netappv1.TridentSnapshotInfoList).ListMeta}
	for _, item := range obj.(*netappv1.TridentSnapshotInfoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tridentSnapshotInfos.
func (c *FakeTridentSnapshotInfos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tridentsnapshotinfosResource, c.ns, opts))

}

// Create takes the representation of a tridentSnapshotInfo and creates it.  Returns the server's representation of the tridentSnapshotInfo, and an error, if there is any.
func (c *FakeTridentSnapshotInfos) Create(ctx context.Context, tridentSnapshotInfo *netappv1.TridentSnapshotInfo, opts v1.CreateOptions) (result *netappv1.TridentSnapshotInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tridentsnapshotinfosResource, c.ns, tridentSnapshotInfo), &netappv1.TridentSnapshotInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentSnapshotInfo), err
}

// Update takes the representation of a tridentSnapshotInfo and updates it. Returns the server's representation of the tridentSnapshotInfo, and an error, if there is any.
func (c *FakeTridentSnapshotInfos) Update(ctx context.Context, tridentSnapshotInfo *netappv1.TridentSnapshotInfo, opts v1.UpdateOptions) (result *netappv1.TridentSnapshotInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tridentsnapshotinfosResource, c.ns, tridentSnapshotInfo), &netappv1.TridentSnapshotInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentSnapshotInfo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTridentSnapshotInfos) UpdateStatus(ctx context.Context, tridentSnapshotInfo *netappv1.TridentSnapshotInfo, opts v1.UpdateOptions) (*netappv1.TridentSnapshotInfo, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tridentsnapshotinfosResource, "status", c.ns, tridentSnapshotInfo), &netappv1.TridentSnapshotInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentSnapshotInfo), err
}

// Delete takes name of the tridentSnapshotInfo and deletes it. Returns an error if one occurs.
func (c *FakeTridentSnapshotInfos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tridentsnapshotinfosResource, c.ns, name), &netappv1.TridentSnapshotInfo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTridentSnapshotInfos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tridentsnapshotinfosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &netappv1.TridentSnapshotInfoList{})
	return err
}

// Patch applies the patch and returns the patched tridentSnapshotInfo.
func (c *FakeTridentSnapshotInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *netappv1.TridentSnapshotInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tridentsnapshotinfosResource, c.ns, name, pt, data, subresources...), &netappv1.TridentSnapshotInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentSnapshotInfo), err
}
