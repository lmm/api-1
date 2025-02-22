// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	scheme "github.com/tigera/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AuthorizationReviewsGetter has a method to return a AuthorizationReviewInterface.
// A group's client should implement this interface.
type AuthorizationReviewsGetter interface {
	AuthorizationReviews() AuthorizationReviewInterface
}

// AuthorizationReviewInterface has methods to work with AuthorizationReview resources.
type AuthorizationReviewInterface interface {
	Create(ctx context.Context, authorizationReview *v3.AuthorizationReview, opts v1.CreateOptions) (*v3.AuthorizationReview, error)
	Update(ctx context.Context, authorizationReview *v3.AuthorizationReview, opts v1.UpdateOptions) (*v3.AuthorizationReview, error)
	UpdateStatus(ctx context.Context, authorizationReview *v3.AuthorizationReview, opts v1.UpdateOptions) (*v3.AuthorizationReview, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.AuthorizationReview, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.AuthorizationReviewList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.AuthorizationReview, err error)
	AuthorizationReviewExpansion
}

// authorizationReviews implements AuthorizationReviewInterface
type authorizationReviews struct {
	client rest.Interface
}

// newAuthorizationReviews returns a AuthorizationReviews
func newAuthorizationReviews(c *ProjectcalicoV3Client) *authorizationReviews {
	return &authorizationReviews{
		client: c.RESTClient(),
	}
}

// Get takes name of the authorizationReview, and returns the corresponding authorizationReview object, and an error if there is any.
func (c *authorizationReviews) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.AuthorizationReview, err error) {
	result = &v3.AuthorizationReview{}
	err = c.client.Get().
		Resource("authorizationreviews").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AuthorizationReviews that match those selectors.
func (c *authorizationReviews) List(ctx context.Context, opts v1.ListOptions) (result *v3.AuthorizationReviewList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.AuthorizationReviewList{}
	err = c.client.Get().
		Resource("authorizationreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested authorizationReviews.
func (c *authorizationReviews) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("authorizationreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a authorizationReview and creates it.  Returns the server's representation of the authorizationReview, and an error, if there is any.
func (c *authorizationReviews) Create(ctx context.Context, authorizationReview *v3.AuthorizationReview, opts v1.CreateOptions) (result *v3.AuthorizationReview, err error) {
	result = &v3.AuthorizationReview{}
	err = c.client.Post().
		Resource("authorizationreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authorizationReview).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a authorizationReview and updates it. Returns the server's representation of the authorizationReview, and an error, if there is any.
func (c *authorizationReviews) Update(ctx context.Context, authorizationReview *v3.AuthorizationReview, opts v1.UpdateOptions) (result *v3.AuthorizationReview, err error) {
	result = &v3.AuthorizationReview{}
	err = c.client.Put().
		Resource("authorizationreviews").
		Name(authorizationReview.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authorizationReview).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *authorizationReviews) UpdateStatus(ctx context.Context, authorizationReview *v3.AuthorizationReview, opts v1.UpdateOptions) (result *v3.AuthorizationReview, err error) {
	result = &v3.AuthorizationReview{}
	err = c.client.Put().
		Resource("authorizationreviews").
		Name(authorizationReview.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authorizationReview).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the authorizationReview and deletes it. Returns an error if one occurs.
func (c *authorizationReviews) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("authorizationreviews").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *authorizationReviews) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("authorizationreviews").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched authorizationReview.
func (c *authorizationReviews) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.AuthorizationReview, err error) {
	result = &v3.AuthorizationReview{}
	err = c.client.Patch(pt).
		Resource("authorizationreviews").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
