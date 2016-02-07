package types

type Cluster struct {
	Token string `yaml: "token,omitempty"`
	State string `yaml: "state,omitempty"`
	Nodes []Node `yaml: "nodes"`
}

func (c *Cluster) GetToken() string {
	if c.Token != "" {
		return c.Token
	}

	return "default.tkn"
}

func (c *Cluster) GetState() string {
	if c.State != "" {
		return c.State
	}

	return "new"
}

func (c *Cluster) GetNodes() []Node {
	return c.Nodes
}
