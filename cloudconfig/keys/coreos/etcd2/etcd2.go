package etcd2

import (
	"github.com/bernardolins/clustereasy/setup/types"
	"strconv"
)

type Etcd2 struct {
	name       string
	parameters map[string]string
	clientPort int
	peerPort   int
	useHttps   bool
}

func Configure(node types.Node, cluster types.Cluster) Etcd2 {
	etcd2 := new(Etcd2)
	etcd2.name = "etcd2"
	etcd2.parameters = make(map[string]string)

	etcd2.clientPort = 2379
	etcd2.peerPort = 2380

	etcd2.configureMachine(node)
	etcd2.configureCluster(cluster)

	etcd2.useHttps = false

	return *etcd2
}

func (etcd2 Etcd2) Name() string {
	return etcd2.name
}

func (etcd2 Etcd2) Parameters() map[string]string {
	return etcd2.parameters
}

// --------------------

func (etcd2 Etcd2) configureCluster(cluster types.Cluster) {
	etcd2.parameters["initial-cluster-token"] = cluster.GetToken()
	etcd2.parameters["initial-cluster-state"] = cluster.GetState()

	etcd2.setInitialClusterString(cluster)
}

func (etcd2 Etcd2) configureMachine(node types.Node) {
	etcd2.parameters["name"] = node.NodeName()
	etcd2.parameters["initial-advertise-peer-urls"] = httpAddress(node.NodeIp(), etcd2.peerPort)
	etcd2.parameters["listen-peer-urls"] = httpAddress(node.NodeIp(), etcd2.peerPort)
	etcd2.parameters["listen-client-urls"] = httpAddress("0.0.0.0", etcd2.clientPort)
	etcd2.parameters["advertise-client-urls"] = httpAddress(node.NodeIp(), etcd2.clientPort)
}

func (etcd2 Etcd2) setInitialClusterString(cluster types.Cluster) {
	initialClusterString := ""

	for _, node := range cluster.GetNodes() {
		nodeString := ""

		if initialClusterString != "" {
			initialClusterString = initialClusterString + ","
		}

		if etcd2.useHttps {
			nodeString = node.NodeName() + "=" + httpsAddress(node.NodeIp(), etcd2.peerPort)
		} else {
			nodeString = node.NodeName() + "=" + httpAddress(node.NodeIp(), etcd2.peerPort)
		}

		initialClusterString = initialClusterString + nodeString
	}

	etcd2.parameters["initial-cluster"] = initialClusterString
}

// -----------------------------

func httpAddress(ip string, port int) string {
	return "http://" + ip + ":" + strconv.Itoa(port)
}

func httpsAddress(ip string, port int) string {
	return "https://" + ip + ":" + strconv.Itoa(port)
}
