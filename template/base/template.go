package base

import (
	"bytes"
	"strings"
)

type (
	TemplateInterface interface {
		GetInterface() []byte
		GetImplement() []byte
		GetTest() []byte
	}
	ModuleInfoInterface interface {
		GetPackageName() string
		GetModuleName() string
		ReplaceTemplate(data []byte) []byte
		ReplaceImplement(in []byte, data []byte) []byte
	}
	Template struct {
		Interface []byte
		Implement []byte
		Test      []byte
	}
	ModuleInfo struct {
		Package string
		Module  string
	}
)

func NewModuleInfo(name string) ModuleInfoInterface {
	return &ModuleInfo{
		Package: strings.ToLower(name),
		Module:  strings.Title(strings.ToLower(name)),
	}
}

func (m *ModuleInfo) ReplaceTemplate(in []byte) []byte {
	in = bytes.ReplaceAll(in, []byte("{{package}}"), []byte(m.GetPackageName()))
	in = bytes.ReplaceAll(in, []byte("{{module}}"), []byte(m.GetModuleName()))
	return in
}
func (m *ModuleInfo) ReplaceImplement(in []byte, data []byte) []byte {
	in = bytes.ReplaceAll(in, []byte("{{implement}}"), data)
	return in
}
func (m *ModuleInfo) GetPackageName() string {
	return m.Package
}

func (m *ModuleInfo) GetModuleName() string {
	return m.Module
}
