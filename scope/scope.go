package scope

import (
	"github.com/bernardolins/clustereasy/service"
)

type Scope struct {
	name     string
	services []service.Service
}

func New(name string) *Scope {
	s := new(Scope)

	s.name = name

	return s
}

func (s *Scope) AddService(service service.Service) {
	s.services = append(s.services, service)
}

func (s *Scope) GetServices() []service.Service {
	return s.services
}
