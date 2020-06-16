//lint:file-ignore ST1005 Generated code

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/szuecs/routegroup-client/client/clientset/versioned/typed/zalando.org/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeZalandoV1 struct {
	*testing.Fake
}

func (c *FakeZalandoV1) RouteGroups(namespace string) v1.RouteGroupInterface {
	return &FakeRouteGroups{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeZalandoV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}