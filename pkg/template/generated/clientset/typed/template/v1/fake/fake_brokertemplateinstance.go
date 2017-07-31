package fake

import (
	template_v1 "github.com/openshift/origin/pkg/template/apis/template/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBrokerTemplateInstances implements BrokerTemplateInstanceInterface
type FakeBrokerTemplateInstances struct {
	Fake *FakeTemplateV1
}

var brokertemplateinstancesResource = schema.GroupVersionResource{Group: "template.openshift.io", Version: "v1", Resource: "brokertemplateinstances"}

var brokertemplateinstancesKind = schema.GroupVersionKind{Group: "template.openshift.io", Version: "v1", Kind: "BrokerTemplateInstance"}

// Get takes name of the brokerTemplateInstance, and returns the corresponding brokerTemplateInstance object, and an error if there is any.
func (c *FakeBrokerTemplateInstances) Get(name string, options v1.GetOptions) (result *template_v1.BrokerTemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(brokertemplateinstancesResource, name), &template_v1.BrokerTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*template_v1.BrokerTemplateInstance), err
}

// List takes label and field selectors, and returns the list of BrokerTemplateInstances that match those selectors.
func (c *FakeBrokerTemplateInstances) List(opts v1.ListOptions) (result *template_v1.BrokerTemplateInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(brokertemplateinstancesResource, brokertemplateinstancesKind, opts), &template_v1.BrokerTemplateInstanceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &template_v1.BrokerTemplateInstanceList{}
	for _, item := range obj.(*template_v1.BrokerTemplateInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested brokerTemplateInstances.
func (c *FakeBrokerTemplateInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(brokertemplateinstancesResource, opts))
}

// Create takes the representation of a brokerTemplateInstance and creates it.  Returns the server's representation of the brokerTemplateInstance, and an error, if there is any.
func (c *FakeBrokerTemplateInstances) Create(brokerTemplateInstance *template_v1.BrokerTemplateInstance) (result *template_v1.BrokerTemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(brokertemplateinstancesResource, brokerTemplateInstance), &template_v1.BrokerTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*template_v1.BrokerTemplateInstance), err
}

// Update takes the representation of a brokerTemplateInstance and updates it. Returns the server's representation of the brokerTemplateInstance, and an error, if there is any.
func (c *FakeBrokerTemplateInstances) Update(brokerTemplateInstance *template_v1.BrokerTemplateInstance) (result *template_v1.BrokerTemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(brokertemplateinstancesResource, brokerTemplateInstance), &template_v1.BrokerTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*template_v1.BrokerTemplateInstance), err
}

// Delete takes name of the brokerTemplateInstance and deletes it. Returns an error if one occurs.
func (c *FakeBrokerTemplateInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(brokertemplateinstancesResource, name), &template_v1.BrokerTemplateInstance{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBrokerTemplateInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(brokertemplateinstancesResource, listOptions)

	_, err := c.Fake.Invokes(action, &template_v1.BrokerTemplateInstanceList{})
	return err
}

// Patch applies the patch and returns the patched brokerTemplateInstance.
func (c *FakeBrokerTemplateInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *template_v1.BrokerTemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(brokertemplateinstancesResource, name, data, subresources...), &template_v1.BrokerTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*template_v1.BrokerTemplateInstance), err
}
