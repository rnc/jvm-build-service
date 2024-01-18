/*
Copyright 2021-2022 Red Hat, Inc.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/redhat-appstudio/jvm-build-service/pkg/apis/jvmbuildservice/v1alpha1"
	scheme "github.com/redhat-appstudio/jvm-build-service/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// JBSConfigsGetter has a method to return a JBSConfigInterface.
// A group's client should implement this interface.
type JBSConfigsGetter interface {
	JBSConfigs(namespace string) JBSConfigInterface
}

// JBSConfigInterface has methods to work with JBSConfig resources.
type JBSConfigInterface interface {
	Create(ctx context.Context, jBSConfig *v1alpha1.JBSConfig, opts v1.CreateOptions) (*v1alpha1.JBSConfig, error)
	Update(ctx context.Context, jBSConfig *v1alpha1.JBSConfig, opts v1.UpdateOptions) (*v1alpha1.JBSConfig, error)
	UpdateStatus(ctx context.Context, jBSConfig *v1alpha1.JBSConfig, opts v1.UpdateOptions) (*v1alpha1.JBSConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.JBSConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.JBSConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JBSConfig, err error)
	JBSConfigExpansion
}

// jBSConfigs implements JBSConfigInterface
type jBSConfigs struct {
	client rest.Interface
	ns     string
}

// newJBSConfigs returns a JBSConfigs
func newJBSConfigs(c *JvmbuildserviceV1alpha1Client, namespace string) *jBSConfigs {
	return &jBSConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the jBSConfig, and returns the corresponding jBSConfig object, and an error if there is any.
func (c *jBSConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.JBSConfig, err error) {
	result = &v1alpha1.JBSConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jbsconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of JBSConfigs that match those selectors.
func (c *jBSConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.JBSConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.JBSConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jbsconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jBSConfigs.
func (c *jBSConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("jbsconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a jBSConfig and creates it.  Returns the server's representation of the jBSConfig, and an error, if there is any.
func (c *jBSConfigs) Create(ctx context.Context, jBSConfig *v1alpha1.JBSConfig, opts v1.CreateOptions) (result *v1alpha1.JBSConfig, err error) {
	result = &v1alpha1.JBSConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("jbsconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jBSConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a jBSConfig and updates it. Returns the server's representation of the jBSConfig, and an error, if there is any.
func (c *jBSConfigs) Update(ctx context.Context, jBSConfig *v1alpha1.JBSConfig, opts v1.UpdateOptions) (result *v1alpha1.JBSConfig, err error) {
	result = &v1alpha1.JBSConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jbsconfigs").
		Name(jBSConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jBSConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *jBSConfigs) UpdateStatus(ctx context.Context, jBSConfig *v1alpha1.JBSConfig, opts v1.UpdateOptions) (result *v1alpha1.JBSConfig, err error) {
	result = &v1alpha1.JBSConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jbsconfigs").
		Name(jBSConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jBSConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the jBSConfig and deletes it. Returns an error if one occurs.
func (c *jBSConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jbsconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *jBSConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jbsconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched jBSConfig.
func (c *jBSConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JBSConfig, err error) {
	result = &v1alpha1.JBSConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("jbsconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
