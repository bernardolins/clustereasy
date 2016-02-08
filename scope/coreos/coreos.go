package coreos

import (
	"github.com/bernardolins/clustereasy/scope"
	"github.com/bernardolins/clustereasy/service/etcd"
	"github.com/bernardolins/clustereasy/service/flannel"
	"github.com/bernardolins/clustereasy/service/fleet"
	"github.com/bernardolins/clustereasy/setup/types"
)

func CreateScope(node types.Node, cluster types.Cluster) *scope.Scope {
	etcd2 := etcd2.New()
	etcd2.Configure(node, cluster)

	fleet := fleet.New()
	fleet.Configure(node, cluster)

	flannel := flannel.New()
	flannel.Configure(node, cluster)

	coreos := scope.New("coreos")

	coreos.AddService(*etcd2)
	coreos.AddService(*fleet)
	coreos.AddService(*flannel)

	return coreos
}
