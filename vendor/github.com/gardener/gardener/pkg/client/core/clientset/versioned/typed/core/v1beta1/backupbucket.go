/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	scheme "github.com/gardener/gardener/pkg/client/core/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackupBucketsGetter has a method to return a BackupBucketInterface.
// A group's client should implement this interface.
type BackupBucketsGetter interface {
	BackupBuckets() BackupBucketInterface
}

// BackupBucketInterface has methods to work with BackupBucket resources.
type BackupBucketInterface interface {
	Create(ctx context.Context, backupBucket *v1beta1.BackupBucket, opts v1.CreateOptions) (*v1beta1.BackupBucket, error)
	Update(ctx context.Context, backupBucket *v1beta1.BackupBucket, opts v1.UpdateOptions) (*v1beta1.BackupBucket, error)
	UpdateStatus(ctx context.Context, backupBucket *v1beta1.BackupBucket, opts v1.UpdateOptions) (*v1beta1.BackupBucket, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.BackupBucket, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.BackupBucketList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackupBucket, err error)
	BackupBucketExpansion
}

// backupBuckets implements BackupBucketInterface
type backupBuckets struct {
	client rest.Interface
}

// newBackupBuckets returns a BackupBuckets
func newBackupBuckets(c *CoreV1beta1Client) *backupBuckets {
	return &backupBuckets{
		client: c.RESTClient(),
	}
}

// Get takes name of the backupBucket, and returns the corresponding backupBucket object, and an error if there is any.
func (c *backupBuckets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackupBucket, err error) {
	result = &v1beta1.BackupBucket{}
	err = c.client.Get().
		Resource("backupbuckets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupBuckets that match those selectors.
func (c *backupBuckets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackupBucketList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BackupBucketList{}
	err = c.client.Get().
		Resource("backupbuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupBuckets.
func (c *backupBuckets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("backupbuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backupBucket and creates it.  Returns the server's representation of the backupBucket, and an error, if there is any.
func (c *backupBuckets) Create(ctx context.Context, backupBucket *v1beta1.BackupBucket, opts v1.CreateOptions) (result *v1beta1.BackupBucket, err error) {
	result = &v1beta1.BackupBucket{}
	err = c.client.Post().
		Resource("backupbuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupBucket).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backupBucket and updates it. Returns the server's representation of the backupBucket, and an error, if there is any.
func (c *backupBuckets) Update(ctx context.Context, backupBucket *v1beta1.BackupBucket, opts v1.UpdateOptions) (result *v1beta1.BackupBucket, err error) {
	result = &v1beta1.BackupBucket{}
	err = c.client.Put().
		Resource("backupbuckets").
		Name(backupBucket.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupBucket).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *backupBuckets) UpdateStatus(ctx context.Context, backupBucket *v1beta1.BackupBucket, opts v1.UpdateOptions) (result *v1beta1.BackupBucket, err error) {
	result = &v1beta1.BackupBucket{}
	err = c.client.Put().
		Resource("backupbuckets").
		Name(backupBucket.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupBucket).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backupBucket and deletes it. Returns an error if one occurs.
func (c *backupBuckets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("backupbuckets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupBuckets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("backupbuckets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backupBucket.
func (c *backupBuckets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackupBucket, err error) {
	result = &v1beta1.BackupBucket{}
	err = c.client.Patch(pt).
		Resource("backupbuckets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
