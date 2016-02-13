package flannel

import (
	"github.com/bernardolins/clustereasy/setup/types"
)

type Flannel struct {
	Name       string
	Parameters map[string]string
}

func New() *Flannel {
	flannel := new(Flannel)

	flannel.Name = "flannel"
	flannel.Parameters = make(map[string]string)

	return flannel
}

func (flannel Flannel) GetName() string {
	return flannel.Name
}

func (flannel Flannel) GetParameters() map[string]string {
	return flannel.Parameters
}

func (flannel Flannel) Configure(node types.Node, cluster types.Cluster) {
	flannel.Parameters["interface"] = node.NodeIp()
}
