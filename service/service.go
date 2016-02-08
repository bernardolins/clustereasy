package service

import (
	"github.com/bernardolins/clustereasy/setup/types"
)

//type Service struct {
//	name   string
//	fields []Field
//}
//
//func NewService(name string) *Service {
//	s := new(Service)
//
//	s.name = name
//
//	return s
//}
//
//func (s *Service) AddField(field Field) {
//	s.fields = append(s.fields, field)
//}
//
//func (s *Service) GetName() string {
//	return s.name
//}
//
//func (s *Service) GetFields() []Field {
//	return s.fields
//}

type Service interface {
	Configure(types.Node, types.Cluster)
}
