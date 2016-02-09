package coreos

import (
	"github.com/bernardolins/clustereasy/setup/types"
	"testing"
)

func TestCreateScope(t *testing.T) {
	cluster := types.Cluster{}
	node := types.Node{}

	scope := CreateScope(node, cluster)

	expect := 3
	got := len(scope.GetServices())

	if expect != got {
		t.Fatalf("expect scope services array len to be %d\ngot %d", expect, got)
	}
}
