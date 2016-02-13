package etcd2

import (
	"github.com/bernardolins/clustereasy/setup/types"
	"strconv"
)

type Etcd2 struct {
	Name       string
	Parameters map[string]string
}

func New() *Etcd2 {
	etcd2 := new(Etcd2)

	etcd2.Name = "etcd2"
	etcd2.Parameters = make(map[string]string)

	return etcd2
}

func (etcd2 Etcd2) GetName() string {
	return etcd2.Name
}

func (etcd2 Etcd2) GetParameters() map[string]string {
	return etcd2.Parameters
}

func (etcd2 Etcd2) Configure(node types.Node, cluster types.Cluster) {
	etcd2.configureClusterParameters(cluster)
	etcd2.configureMachineParameters(node)
}

func (etcd2 Etcd2) configureClusterParameters(cluster types.Cluster) {
	etcd2.Parameters["initial-cluster-token"] = cluster.GetToken()
	etcd2.Parameters["initial-cluster-state"] = cluster.GetState()

	etcd2.Parameters["initial-cluster"] = setInitialClusterString(cluster)
}

func setInitialClusterString(cluster types.Cluster) string {
	initialClusterString := ""

	for _, node := range cluster.GetNodes() {
		if initialClusterString != "" {
			initialClusterString = initialClusterString + ","
		}

		nodeString := node.NodeName() + "=http://" + node.NodeIp() + ":2380"
		initialClusterString = initialClusterString + nodeString
	}

	return initialClusterString
}

func (etcd2 Etcd2) configureMachineParameters(node types.Node) {
	etcd2.Parameters["name"] = node.NodeName()
	etcd2.Parameters["initial-advertise-peer-urls"] = httpAddress(node.NodeIp(), 2380)
	etcd2.Parameters["listen-peer-urls"] = httpAddress(node.NodeIp(), 2380)
	etcd2.Parameters["listen-client-urls"] = httpAddress(node.NodeIp(), 2379) + "," + httpAddress(node.NodeIp(), 2379)
	etcd2.Parameters["advertise-client-urls"] = httpAddress(node.NodeIp(), 2379)
}

func httpAddress(ip string, port int) string {
	return "http://" + ip + ":" + strconv.Itoa(port)
}
