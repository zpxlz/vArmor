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
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVarmorPolicies implements VarmorPolicyInterface
type FakeVarmorPolicies struct {
	Fake *FakeCrdV1beta1
	ns   string
}

var varmorpoliciesResource = schema.GroupVersionResource{Group: "crd.varmor.org", Version: "v1beta1", Resource: "varmorpolicies"}

var varmorpoliciesKind = schema.GroupVersionKind{Group: "crd.varmor.org", Version: "v1beta1", Kind: "VarmorPolicy"}

// Get takes name of the varmorPolicy, and returns the corresponding varmorPolicy object, and an error if there is any.
func (c *FakeVarmorPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VarmorPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(varmorpoliciesResource, c.ns, name), &v1beta1.VarmorPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VarmorPolicy), err
}

// List takes label and field selectors, and returns the list of VarmorPolicies that match those selectors.
func (c *FakeVarmorPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VarmorPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(varmorpoliciesResource, varmorpoliciesKind, c.ns, opts), &v1beta1.VarmorPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VarmorPolicyList{ListMeta: obj.(*v1beta1.VarmorPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.VarmorPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested varmorPolicies.
func (c *FakeVarmorPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(varmorpoliciesResource, c.ns, opts))

}

// Create takes the representation of a varmorPolicy and creates it.  Returns the server's representation of the varmorPolicy, and an error, if there is any.
func (c *FakeVarmorPolicies) Create(ctx context.Context, varmorPolicy *v1beta1.VarmorPolicy, opts v1.CreateOptions) (result *v1beta1.VarmorPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(varmorpoliciesResource, c.ns, varmorPolicy), &v1beta1.VarmorPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VarmorPolicy), err
}

// Update takes the representation of a varmorPolicy and updates it. Returns the server's representation of the varmorPolicy, and an error, if there is any.
func (c *FakeVarmorPolicies) Update(ctx context.Context, varmorPolicy *v1beta1.VarmorPolicy, opts v1.UpdateOptions) (result *v1beta1.VarmorPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(varmorpoliciesResource, c.ns, varmorPolicy), &v1beta1.VarmorPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VarmorPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVarmorPolicies) UpdateStatus(ctx context.Context, varmorPolicy *v1beta1.VarmorPolicy, opts v1.UpdateOptions) (*v1beta1.VarmorPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(varmorpoliciesResource, "status", c.ns, varmorPolicy), &v1beta1.VarmorPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VarmorPolicy), err
}

// Delete takes name of the varmorPolicy and deletes it. Returns an error if one occurs.
func (c *FakeVarmorPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(varmorpoliciesResource, c.ns, name, opts), &v1beta1.VarmorPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVarmorPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(varmorpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VarmorPolicyList{})
	return err
}

// Patch applies the patch and returns the patched varmorPolicy.
func (c *FakeVarmorPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VarmorPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(varmorpoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.VarmorPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VarmorPolicy), err
}
