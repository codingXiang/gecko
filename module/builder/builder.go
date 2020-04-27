package builder

import (
	"github.com/codingXiang/gecko/parser"
	"github.com/codingXiang/gecko/template/base"
)

const (
	SERVICE = "service"
)

type ModuleBuilderInterface interface {
	SetSource(in string)
	GetSource() parser.ParserInterface
	SetDestination(in string, name string)
	GetDestination() string
	GetInfo() base.ModuleInfoInterface
	SetupFolder()
	Create() (_interface []byte, _implement []byte, _test []byte)
	General() (_interface []byte, _implement []byte, _test []byte)
	Save(_interface []byte, _implement []byte, _test []byte)
}

type (
	ModuleBuilder struct {
		Source      parser.ParserInterface
		Destination string
		Info        base.ModuleInfoInterface
	}

	DeliveryBuilder struct {
		ModuleBuilder
		//Template *De
	}
)
