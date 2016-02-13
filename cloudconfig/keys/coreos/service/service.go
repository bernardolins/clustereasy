package service

import (
	"github.com/bernardolins/clustereasy/setup/types"
)

type Service interface {
	GetName() string
	GetParameters() map[string]string
	Configure(types.Node, types.Cluster)
}
