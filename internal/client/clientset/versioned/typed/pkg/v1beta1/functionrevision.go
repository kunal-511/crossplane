/*
Copyright 2021 The Crossplane Authors.

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/crossplane/crossplane/apis/pkg/v1beta1"
	scheme "github.com/crossplane/crossplane/internal/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FunctionRevisionsGetter has a method to return a FunctionRevisionInterface.
// A group's client should implement this interface.
type FunctionRevisionsGetter interface {
	FunctionRevisions() FunctionRevisionInterface
}

// FunctionRevisionInterface has methods to work with FunctionRevision resources.
type FunctionRevisionInterface interface {
	Create(ctx context.Context, functionRevision *v1beta1.FunctionRevision, opts v1.CreateOptions) (*v1beta1.FunctionRevision, error)
	Update(ctx context.Context, functionRevision *v1beta1.FunctionRevision, opts v1.UpdateOptions) (*v1beta1.FunctionRevision, error)
	UpdateStatus(ctx context.Context, functionRevision *v1beta1.FunctionRevision, opts v1.UpdateOptions) (*v1beta1.FunctionRevision, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.FunctionRevision, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.FunctionRevisionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FunctionRevision, err error)
	FunctionRevisionExpansion
}

// functionRevisions implements FunctionRevisionInterface
type functionRevisions struct {
	client rest.Interface
}

// newFunctionRevisions returns a FunctionRevisions
func newFunctionRevisions(c *PkgV1beta1Client) *functionRevisions {
	return &functionRevisions{
		client: c.RESTClient(),
	}
}

// Get takes name of the functionRevision, and returns the corresponding functionRevision object, and an error if there is any.
func (c *functionRevisions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FunctionRevision, err error) {
	result = &v1beta1.FunctionRevision{}
	err = c.client.Get().
		Resource("functionrevisions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FunctionRevisions that match those selectors.
func (c *functionRevisions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FunctionRevisionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.FunctionRevisionList{}
	err = c.client.Get().
		Resource("functionrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested functionRevisions.
func (c *functionRevisions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("functionrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a functionRevision and creates it.  Returns the server's representation of the functionRevision, and an error, if there is any.
func (c *functionRevisions) Create(ctx context.Context, functionRevision *v1beta1.FunctionRevision, opts v1.CreateOptions) (result *v1beta1.FunctionRevision, err error) {
	result = &v1beta1.FunctionRevision{}
	err = c.client.Post().
		Resource("functionrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(functionRevision).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a functionRevision and updates it. Returns the server's representation of the functionRevision, and an error, if there is any.
func (c *functionRevisions) Update(ctx context.Context, functionRevision *v1beta1.FunctionRevision, opts v1.UpdateOptions) (result *v1beta1.FunctionRevision, err error) {
	result = &v1beta1.FunctionRevision{}
	err = c.client.Put().
		Resource("functionrevisions").
		Name(functionRevision.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(functionRevision).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *functionRevisions) UpdateStatus(ctx context.Context, functionRevision *v1beta1.FunctionRevision, opts v1.UpdateOptions) (result *v1beta1.FunctionRevision, err error) {
	result = &v1beta1.FunctionRevision{}
	err = c.client.Put().
		Resource("functionrevisions").
		Name(functionRevision.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(functionRevision).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the functionRevision and deletes it. Returns an error if one occurs.
func (c *functionRevisions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("functionrevisions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *functionRevisions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("functionrevisions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched functionRevision.
func (c *functionRevisions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FunctionRevision, err error) {
	result = &v1beta1.FunctionRevision{}
	err = c.client.Patch(pt).
		Resource("functionrevisions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}