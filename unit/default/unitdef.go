package unitdef

import (
	"github.com/bernardolins/clustereasy/unit"
)

func DefaultUnits() []unit.Unit {
	def := []unit.Unit{
		defaultContentlessUnit("etcd2", "start"),
		defaultContentlessUnit("fleet", "start"),
		defaultContentlessUnit("flannel", "start"),
	}

	return def
}

func defaultContentlessUnit(name, command string) unit.Unit {
	u := unit.New(name, command)

	return *u
}
