package fleet

import (
	"github.com/bernardolins/clustereasy/setup/types"
)

type Fleet struct {
	Name       string
	Parameters map[string]string
}

func New() *Fleet {
	fleet := new(Fleet)

	fleet.Name = "fleet"
	fleet.Parameters = make(map[string]string)

	return fleet
}

func (fleet Fleet) GetName() string {
	return fleet.Name
}

func (fleet Fleet) GetParameters() map[string]string {
	return fleet.Parameters
}

func (fleet Fleet) Configure(node types.Node, cluster types.Cluster) {
	fleet.Parameters["public-ip"] = node.NodeIp()
}
