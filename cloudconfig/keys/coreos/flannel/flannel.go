package flannel

import (
	"github.com/bernardolins/clustereasy/setup/types"
)

type Flannel struct {
	name       string
	parameters map[string]string
}

func Configure(node types.Node, cluster types.Cluster) Flannel {
	flannel := new(Flannel)
	flannel.name = "flannel"
	flannel.parameters = make(map[string]string)

	flannel.configure(node, cluster)

	return *flannel
}

func (flannel Flannel) Name() string {
	return flannel.name
}

func (flannel Flannel) Parameters() map[string]string {
	return flannel.parameters
}

func (flannel Flannel) configure(node types.Node, cluster types.Cluster) {
	flannel.parameters["interface"] = node.NodeIp()
}
