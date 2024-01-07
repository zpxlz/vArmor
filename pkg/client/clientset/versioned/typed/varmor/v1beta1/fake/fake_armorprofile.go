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

package fake

import (
	"context"

	v1beta1 "github.com/bytedance/vArmor/apis/varmor/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArmorProfiles implements ArmorProfileInterface
type FakeArmorProfiles struct {
	Fake *FakeCrdV1beta1
	ns   string
}

var armorprofilesResource = v1beta1.SchemeGroupVersion.WithResource("armorprofiles")

var armorprofilesKind = v1beta1.SchemeGroupVersion.WithKind("ArmorProfile")

// Get takes name of the armorProfile, and returns the corresponding armorProfile object, and an error if there is any.
func (c *FakeArmorProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ArmorProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(armorprofilesResource, c.ns, name), &v1beta1.ArmorProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArmorProfile), err
}

// List takes label and field selectors, and returns the list of ArmorProfiles that match those selectors.
func (c *FakeArmorProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ArmorProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(armorprofilesResource, armorprofilesKind, c.ns, opts), &v1beta1.ArmorProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ArmorProfileList{ListMeta: obj.(*v1beta1.ArmorProfileList).ListMeta}
	for _, item := range obj.(*v1beta1.ArmorProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested armorProfiles.
func (c *FakeArmorProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(armorprofilesResource, c.ns, opts))

}

// Create takes the representation of a armorProfile and creates it.  Returns the server's representation of the armorProfile, and an error, if there is any.
func (c *FakeArmorProfiles) Create(ctx context.Context, armorProfile *v1beta1.ArmorProfile, opts v1.CreateOptions) (result *v1beta1.ArmorProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(armorprofilesResource, c.ns, armorProfile), &v1beta1.ArmorProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArmorProfile), err
}

// Update takes the representation of a armorProfile and updates it. Returns the server's representation of the armorProfile, and an error, if there is any.
func (c *FakeArmorProfiles) Update(ctx context.Context, armorProfile *v1beta1.ArmorProfile, opts v1.UpdateOptions) (result *v1beta1.ArmorProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(armorprofilesResource, c.ns, armorProfile), &v1beta1.ArmorProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArmorProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArmorProfiles) UpdateStatus(ctx context.Context, armorProfile *v1beta1.ArmorProfile, opts v1.UpdateOptions) (*v1beta1.ArmorProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(armorprofilesResource, "status", c.ns, armorProfile), &v1beta1.ArmorProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArmorProfile), err
}

// Delete takes name of the armorProfile and deletes it. Returns an error if one occurs.
func (c *FakeArmorProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(armorprofilesResource, c.ns, name, opts), &v1beta1.ArmorProfile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArmorProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(armorprofilesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ArmorProfileList{})
	return err
}

// Patch applies the patch and returns the patched armorProfile.
func (c *FakeArmorProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ArmorProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(armorprofilesResource, c.ns, name, pt, data, subresources...), &v1beta1.ArmorProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ArmorProfile), err
}
