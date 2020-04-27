package repository

import (
	"github.com/codingXiang/gecko/template/base"
)

type RepositoryTemplate interface {
	base.TemplateInterface
}

type Repository struct {
	base.Template
}

func NewRepositoryTemplate() RepositoryTemplate {
	var r = &Repository{}
	r.Interface = INTERFACE
	r.Implement = IMPLEMENT
	r.Test = TEST
	return r
}

func (r *Repository) GetInterface() []byte {
	return r.Interface
}

func (r *Repository) GetImplement() []byte {
	return r.Implement
}

func (r *Repository) GetTest() []byte {
	return r.Test
}
