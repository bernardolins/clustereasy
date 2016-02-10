package scope

import (
	"github.com/bernardolins/clustereasy/service"
	"github.com/bernardolins/clustereasy/unit"
)

type Scope struct {
	Name     string
	Services []service.Service
	Units    []*unit.Unit
}

func New(Name string) *Scope {
	s := new(Scope)

	s.Name = Name

	return s
}

func (s *Scope) AddService(service service.Service) {
	s.Services = append(s.Services, service)
}

func (s *Scope) GetServices() []service.Service {
	return s.Services
}

func (s *Scope) GetName() string {
	return s.Name
}

func (s *Scope) GetUnits() []*unit.Unit {
	return s.Units
}

func (s *Scope) AddUnit(unit *unit.Unit) {
	s.Units = append(s.Units, unit)
}
