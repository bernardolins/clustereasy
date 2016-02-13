package coreos

import (
	"github.com/bernardolins/clustereasy/cloudconfig/keys/coreos/etcd2"
	"github.com/bernardolins/clustereasy/cloudconfig/keys/coreos/flannel"
	"github.com/bernardolins/clustereasy/cloudconfig/keys/coreos/fleet"
	"github.com/bernardolins/clustereasy/cloudconfig/keys/coreos/unit"
	"github.com/bernardolins/clustereasy/setup/types"
)

type CoreOS struct {
	etcd2   etcd2.Etcd2
	fleet   fleet.Fleet
	flannel flannel.Flannel
	units   map[string]*unit.Unit
}

func New() *CoreOS {
	c := new(CoreOS)
	c.units = make(map[string]*unit.Unit)

	return c
}

func ConfigureCoreOS(node types.Node, cluster types.Cluster) CoreOS {
	coreos := New()

	coreos.etcd2 = etcd2.Configure(node, cluster)
	coreos.fleet = fleet.Configure(node, cluster)
	coreos.flannel = flannel.Configure(node, cluster)

	configureUnits(coreos, cluster)

	return *coreos
}

func configureUnits(coreos *CoreOS, cluster types.Cluster) {
	for _, unt := range cluster.GetUnits() {
		coreUnit := unit.New(unt.UnitName())
		coreUnit.Configure(unt)
		coreos.units[unt.UnitName()] = coreUnit
	}
}

func (c CoreOS) Etcd2() etcd2.Etcd2 {
	return c.etcd2
}

func (c CoreOS) Fleet() fleet.Fleet {
	return c.fleet
}

func (c CoreOS) Flannel() flannel.Flannel {
	return c.flannel

}

func (c CoreOS) Units() map[string]*unit.Unit {
	return c.units
}
