//lint:file-ignore ST1005 Generated code

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	zalandoorgv1 "github.com/szuecs/routegroup-client/apis/zalando.org/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRouteGroups implements RouteGroupInterface
type FakeRouteGroups struct {
	Fake *FakeZalandoV1
	ns   string
}

var routegroupsResource = schema.GroupVersionResource{Group: "zalando.org", Version: "v1", Resource: "routegroups"}

var routegroupsKind = schema.GroupVersionKind{Group: "zalando.org", Version: "v1", Kind: "RouteGroup"}

// Get takes name of the routeGroup, and returns the corresponding routeGroup object, and an error if there is any.
func (c *FakeRouteGroups) Get(name string, options v1.GetOptions) (result *zalandoorgv1.RouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(routegroupsResource, c.ns, name), &zalandoorgv1.RouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*zalandoorgv1.RouteGroup), err
}

// List takes label and field selectors, and returns the list of RouteGroups that match those selectors.
func (c *FakeRouteGroups) List(opts v1.ListOptions) (result *zalandoorgv1.RouteGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(routegroupsResource, routegroupsKind, c.ns, opts), &zalandoorgv1.RouteGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &zalandoorgv1.RouteGroupList{ListMeta: obj.(*zalandoorgv1.RouteGroupList).ListMeta}
	for _, item := range obj.(*zalandoorgv1.RouteGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routeGroups.
func (c *FakeRouteGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(routegroupsResource, c.ns, opts))

}

// Create takes the representation of a routeGroup and creates it.  Returns the server's representation of the routeGroup, and an error, if there is any.
func (c *FakeRouteGroups) Create(routeGroup *zalandoorgv1.RouteGroup) (result *zalandoorgv1.RouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(routegroupsResource, c.ns, routeGroup), &zalandoorgv1.RouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*zalandoorgv1.RouteGroup), err
}

// Update takes the representation of a routeGroup and updates it. Returns the server's representation of the routeGroup, and an error, if there is any.
func (c *FakeRouteGroups) Update(routeGroup *zalandoorgv1.RouteGroup) (result *zalandoorgv1.RouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(routegroupsResource, c.ns, routeGroup), &zalandoorgv1.RouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*zalandoorgv1.RouteGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRouteGroups) UpdateStatus(routeGroup *zalandoorgv1.RouteGroup) (*zalandoorgv1.RouteGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(routegroupsResource, "status", c.ns, routeGroup), &zalandoorgv1.RouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*zalandoorgv1.RouteGroup), err
}

// Delete takes name of the routeGroup and deletes it. Returns an error if one occurs.
func (c *FakeRouteGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(routegroupsResource, c.ns, name), &zalandoorgv1.RouteGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRouteGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(routegroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &zalandoorgv1.RouteGroupList{})
	return err
}

// Patch applies the patch and returns the patched routeGroup.
func (c *FakeRouteGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *zalandoorgv1.RouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(routegroupsResource, c.ns, name, pt, data, subresources...), &zalandoorgv1.RouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*zalandoorgv1.RouteGroup), err
}
