package service

import "github.com/codingXiang/gecko/template/base"

type ServiceTemplate interface {
	base.TemplateInterface
}

type Service struct {
	base.Template
}

func NewServiceTemplate() ServiceTemplate {
	var r = &Service{}
	r.Interface = INTERFACE
	r.Implement = IMPLEMENT
	r.Test = TEST
	return r
}

func (s *Service) GetInterface() []byte {
	return s.Interface
}

func (s *Service) GetImplement() []byte {
	return s.Implement
}

func (s *Service) GetTest() []byte {
	return s.Test
}
