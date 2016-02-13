package fleet

import (
	"github.com/bernardolins/clustereasy/setup/types"
)

type Fleet struct {
	name       string
	parameters map[string]string
}

func Configure(node types.Node, cluster types.Cluster) Fleet {
	fleet := new(Fleet)
	fleet.name = "fleet"
	fleet.parameters = make(map[string]string)

	fleet.configure(node, cluster)

	return *fleet
}

func (fleet Fleet) Name() string {
	return fleet.name
}

func (fleet Fleet) Parameters() map[string]string {
	return fleet.parameters
}

func (fleet Fleet) configure(node types.Node, cluster types.Cluster) {
	fleet.parameters["public-ip"] = node.NodeIp()
}
