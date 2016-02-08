package init

import (
	"github.com/bernardolins/clustereasy/init/types"
)

type InitData struct {
	Cluster types.Cluster
}

func (i InitData) GetClusterInfo() types.Cluster {
	return i.Cluster
}
