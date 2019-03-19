/*
Copyright 2018 Openstorage.org

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

package fake

import (
	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterDomainsStatuses implements ClusterDomainsStatusInterface
type FakeClusterDomainsStatuses struct {
	Fake *FakeStorkV1alpha1
	ns   string
}

var clusterdomainsstatusesResource = schema.GroupVersionResource{Group: "stork.libopenstorage.org", Version: "v1alpha1", Resource: "clusterdomainsstatuses"}

var clusterdomainsstatusesKind = schema.GroupVersionKind{Group: "stork.libopenstorage.org", Version: "v1alpha1", Kind: "ClusterDomainsStatus"}

// Get takes name of the clusterDomainsStatus, and returns the corresponding clusterDomainsStatus object, and an error if there is any.
func (c *FakeClusterDomainsStatuses) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterDomainsStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterdomainsstatusesResource, c.ns, name), &v1alpha1.ClusterDomainsStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainsStatus), err
}

// List takes label and field selectors, and returns the list of ClusterDomainsStatuses that match those selectors.
func (c *FakeClusterDomainsStatuses) List(opts v1.ListOptions) (result *v1alpha1.ClusterDomainsStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterdomainsstatusesResource, clusterdomainsstatusesKind, c.ns, opts), &v1alpha1.ClusterDomainsStatusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterDomainsStatusList{ListMeta: obj.(*v1alpha1.ClusterDomainsStatusList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterDomainsStatusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterDomainsStatuses.
func (c *FakeClusterDomainsStatuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterdomainsstatusesResource, c.ns, opts))

}

// Create takes the representation of a clusterDomainsStatus and creates it.  Returns the server's representation of the clusterDomainsStatus, and an error, if there is any.
func (c *FakeClusterDomainsStatuses) Create(clusterDomainsStatus *v1alpha1.ClusterDomainsStatus) (result *v1alpha1.ClusterDomainsStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterdomainsstatusesResource, c.ns, clusterDomainsStatus), &v1alpha1.ClusterDomainsStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainsStatus), err
}

// Update takes the representation of a clusterDomainsStatus and updates it. Returns the server's representation of the clusterDomainsStatus, and an error, if there is any.
func (c *FakeClusterDomainsStatuses) Update(clusterDomainsStatus *v1alpha1.ClusterDomainsStatus) (result *v1alpha1.ClusterDomainsStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterdomainsstatusesResource, c.ns, clusterDomainsStatus), &v1alpha1.ClusterDomainsStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainsStatus), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterDomainsStatuses) UpdateStatus(clusterDomainsStatus *v1alpha1.ClusterDomainsStatus) (*v1alpha1.ClusterDomainsStatus, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusterdomainsstatusesResource, "status", c.ns, clusterDomainsStatus), &v1alpha1.ClusterDomainsStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainsStatus), err
}

// Delete takes name of the clusterDomainsStatus and deletes it. Returns an error if one occurs.
func (c *FakeClusterDomainsStatuses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusterdomainsstatusesResource, c.ns, name), &v1alpha1.ClusterDomainsStatus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterDomainsStatuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterdomainsstatusesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterDomainsStatusList{})
	return err
}

// Patch applies the patch and returns the patched clusterDomainsStatus.
func (c *FakeClusterDomainsStatuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterDomainsStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterdomainsstatusesResource, c.ns, name, data, subresources...), &v1alpha1.ClusterDomainsStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainsStatus), err
}