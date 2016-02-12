package unitdef

import (
	"github.com/bernardolins/clustereasy/setup/types"
	"github.com/bernardolins/clustereasy/unit"
	"github.com/bernardolins/clustereasy/unit/default/unitcont"
)

func DefaultUnits(node types.Node) []*unit.Unit {
	def := []*unit.Unit{
		etcd2(),
		fleet(),
		flannel(),
		staticNetwork(node),
	}

	return def
}

func defaultContentlessUnit(name, command string) *unit.Unit {
	u := unit.New(name)

	return u
}

func DefaultContentUnit(name, command, content string) *unit.Unit {
	u := unit.New(name)

	return u
}

func staticNetwork(node types.Node) *unit.Unit {
	if node.NodeIp() != "" {
		u := unit.New("10-" + node.NodeInterface() + ".network")
		u.OnRuntime()
		u.SetContent(unitcont.DefaultStaticIp(node))

		return u
	}

	return nil
}

func etcd2() *unit.Unit {
	u := unit.New("etcd2")
	u.AddCommand("start")

	return u
}

func fleet() *unit.Unit {
	u := unit.New("fleet")
	u.AddCommand("start")

	return u
}

func flannel() *unit.Unit {
	u := unit.New("flannel")
	u.AddCommand("start")

	return u
}
