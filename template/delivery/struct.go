package delivery

import (
	"github.com/codingXiang/gecko/template/base"
	"github.com/codingXiang/gecko/template/delivery/http"
)

type (
	HttpTemplate interface {
		base.TemplateInterface
	}
	GRPCTemplate interface {
		base.TemplateInterface
	}
	CliTemplate interface {
		base.TemplateInterface
	}
)

type (
	Http struct {
		base.Template
	}
	GRPC struct {
		base.Template
	}
	Cli struct {
		base.Template
	}
)

func NewHttpTemplate() HttpTemplate {
	var r = &Http{}
	r.Interface = INTERFACE
	r.Implement = http.IMPLEMENT
	r.Test = http.TEST
	return r
}

func (r *Http) GetInterface() []byte {
	return r.Interface
}

func (r *Http) GetImplement() []byte {
	return r.Implement
}

func (r *Http) GetTest() []byte {
	return r.Test
}
