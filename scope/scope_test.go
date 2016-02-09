package scope

import (
	"github.com/bernardolins/clustereasy/service/etcd"
	"testing"
)

func TestAddService(t *testing.T) {
	service := etcd2.New()
	scope := New("test")

	scope.AddService(service)

	expect := 1
	got := len(scope.GetServices())

	if expect != got {
		t.Fatalf("expect services array len to be %d\ngot %d", expect, got)
	}
}
