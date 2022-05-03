/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/nokia/danm/crd/apis/danm/v1"
	scheme "github.com/nokia/danm/crd/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TenantConfigsGetter has a method to return a TenantConfigInterface.
// A group's client should implement this interface.
type TenantConfigsGetter interface {
	TenantConfigs() TenantConfigInterface
}

// TenantConfigInterface has methods to work with TenantConfig resources.
type TenantConfigInterface interface {
	Create(ctx context.Context, tenantConfig *v1.TenantConfig, opts metav1.CreateOptions) (*v1.TenantConfig, error)
	Update(ctx context.Context, tenantConfig *v1.TenantConfig, opts metav1.UpdateOptions) (*v1.TenantConfig, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TenantConfig, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TenantConfigList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TenantConfig, err error)
	TenantConfigExpansion
}

// tenantConfigs implements TenantConfigInterface
type tenantConfigs struct {
	client rest.Interface
}

// newTenantConfigs returns a TenantConfigs
func newTenantConfigs(c *DanmV1Client) *tenantConfigs {
	return &tenantConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the tenantConfig, and returns the corresponding tenantConfig object, and an error if there is any.
func (c *tenantConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TenantConfig, err error) {
	result = &v1.TenantConfig{}
	err = c.client.Get().
		Resource("tenantconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TenantConfigs that match those selectors.
func (c *tenantConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TenantConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TenantConfigList{}
	err = c.client.Get().
		Resource("tenantconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tenantConfigs.
func (c *tenantConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("tenantconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tenantConfig and creates it.  Returns the server's representation of the tenantConfig, and an error, if there is any.
func (c *tenantConfigs) Create(ctx context.Context, tenantConfig *v1.TenantConfig, opts metav1.CreateOptions) (result *v1.TenantConfig, err error) {
	result = &v1.TenantConfig{}
	err = c.client.Post().
		Resource("tenantconfigs").Namespace(tenantConfig.Namespace).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tenantConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tenantConfig and updates it. Returns the server's representation of the tenantConfig, and an error, if there is any.
func (c *tenantConfigs) Update(ctx context.Context, tenantConfig *v1.TenantConfig, opts metav1.UpdateOptions) (result *v1.TenantConfig, err error) {
	result = &v1.TenantConfig{}
	err = c.client.Put().
		Resource("tenantconfigs").
		Name(tenantConfig.Name).Namespace(tenantConfig.Namespace).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tenantConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tenantConfig and deletes it. Returns an error if one occurs.
func (c *tenantConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("tenantconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tenantConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("tenantconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tenantConfig.
func (c *tenantConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TenantConfig, err error) {
	result = &v1.TenantConfig{}
	err = c.client.Patch(pt).
		Resource("tenantconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
