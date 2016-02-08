package types

type Node struct {
	Name      string `yaml: "name"`
	Interface string `yaml: "interface"`
	Ip        string `yaml: "ip"`
}

func (n *Node) NodeName() string {
	return n.Name
}

func (n *Node) NodeInterface() string {
	return n.Interface
}

func (n *Node) NodeIp() string {
	return n.Ip
}
