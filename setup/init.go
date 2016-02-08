package setup

import (
	"github.com/bernardolins/clustereasy/setup/types"
)

type InitData struct {
	Cluster types.Cluster
}

func (i InitData) GetClusterInfo() types.Cluster {
	return i.Cluster
}
