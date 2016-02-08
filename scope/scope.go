package scope

import (
	"github.com/bernardolins/clustereasy/service"
)

type Scope struct {
	Name     string
	Services []service.Service
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
