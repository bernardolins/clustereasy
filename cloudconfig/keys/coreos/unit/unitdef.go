package unit

import (
	"fmt"
	"github.com/bernardolins/clustereasy/setup/types"
)

func DefaultUnits() []*Unit {
	def := []*Unit{
		etcd2(),
		fleet(),
		flannel(),
	}

	return def
}

func etcd2() *Unit {
	u := New("etcd2")
	u.SetParameter("command", "start")

	return u
}

func fleet() *Unit {
	u := New("fleet")
	u.SetParameter("command", "start")

	return u
}

func flannel() *Unit {
	u := New("flannel")
	u.SetParameter("command", "start")

	return u
}

func Network(node types.Node) *Unit {
	content := fmt.Sprintf("[Match]\nName=%s\n[Network]\nAddress=%s", node.NodeInterface(), node.NodeIp())

	u := New("static.network")
	u.SetParameter("runtime", true)
	u.SetContent(content)

	return u
}
