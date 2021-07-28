// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	scheme "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TridentMirrorRelationshipsGetter has a method to return a TridentMirrorRelationshipInterface.
// A group's client should implement this interface.
type TridentMirrorRelationshipsGetter interface {
	TridentMirrorRelationships(namespace string) TridentMirrorRelationshipInterface
}

// TridentMirrorRelationshipInterface has methods to work with TridentMirrorRelationship resources.
type TridentMirrorRelationshipInterface interface {
	Create(ctx context.Context, tridentMirrorRelationship *v1.TridentMirrorRelationship, opts metav1.CreateOptions) (*v1.TridentMirrorRelationship, error)
	Update(ctx context.Context, tridentMirrorRelationship *v1.TridentMirrorRelationship, opts metav1.UpdateOptions) (*v1.TridentMirrorRelationship, error)
	UpdateStatus(ctx context.Context, tridentMirrorRelationship *v1.TridentMirrorRelationship, opts metav1.UpdateOptions) (*v1.TridentMirrorRelationship, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TridentMirrorRelationship, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TridentMirrorRelationshipList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentMirrorRelationship, err error)
	TridentMirrorRelationshipExpansion
}

// tridentMirrorRelationships implements TridentMirrorRelationshipInterface
type tridentMirrorRelationships struct {
	client rest.Interface
	ns     string
}

// newTridentMirrorRelationships returns a TridentMirrorRelationships
func newTridentMirrorRelationships(c *TridentV1Client, namespace string) *tridentMirrorRelationships {
	return &tridentMirrorRelationships{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tridentMirrorRelationship, and returns the corresponding tridentMirrorRelationship object, and an error if there is any.
func (c *tridentMirrorRelationships) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TridentMirrorRelationship, err error) {
	result = &v1.TridentMirrorRelationship{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentmirrorrelationships").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TridentMirrorRelationships that match those selectors.
func (c *tridentMirrorRelationships) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TridentMirrorRelationshipList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TridentMirrorRelationshipList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentmirrorrelationships").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tridentMirrorRelationships.
func (c *tridentMirrorRelationships) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tridentmirrorrelationships").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tridentMirrorRelationship and creates it.  Returns the server's representation of the tridentMirrorRelationship, and an error, if there is any.
func (c *tridentMirrorRelationships) Create(ctx context.Context, tridentMirrorRelationship *v1.TridentMirrorRelationship, opts metav1.CreateOptions) (result *v1.TridentMirrorRelationship, err error) {
	result = &v1.TridentMirrorRelationship{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tridentmirrorrelationships").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentMirrorRelationship).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tridentMirrorRelationship and updates it. Returns the server's representation of the tridentMirrorRelationship, and an error, if there is any.
func (c *tridentMirrorRelationships) Update(ctx context.Context, tridentMirrorRelationship *v1.TridentMirrorRelationship, opts metav1.UpdateOptions) (result *v1.TridentMirrorRelationship, err error) {
	result = &v1.TridentMirrorRelationship{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentmirrorrelationships").
		Name(tridentMirrorRelationship.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentMirrorRelationship).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tridentMirrorRelationships) UpdateStatus(ctx context.Context, tridentMirrorRelationship *v1.TridentMirrorRelationship, opts metav1.UpdateOptions) (result *v1.TridentMirrorRelationship, err error) {
	result = &v1.TridentMirrorRelationship{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentmirrorrelationships").
		Name(tridentMirrorRelationship.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentMirrorRelationship).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tridentMirrorRelationship and deletes it. Returns an error if one occurs.
func (c *tridentMirrorRelationships) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentmirrorrelationships").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tridentMirrorRelationships) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentmirrorrelationships").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tridentMirrorRelationship.
func (c *tridentMirrorRelationships) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentMirrorRelationship, err error) {
	result = &v1.TridentMirrorRelationship{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tridentmirrorrelationships").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
