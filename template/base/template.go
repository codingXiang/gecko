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
		ReplaceTemplateMethod(in []byte, data map[string]string) []byte
		ReplaceImplement(in []byte, data []byte) []byte
		ReplaceImplementMethod(in []byte, data map[string]string) []byte
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
	in = bytes.ReplaceAll(in, []byte("{{ .package }}"), []byte(m.GetPackageName()))
	in = bytes.ReplaceAll(in, []byte("{{ .module }}"), []byte(m.GetModuleName()))
	return in
}
func (m *ModuleInfo) ReplaceTemplateMethod(in []byte, data map[string]string) []byte {
	in = bytes.ReplaceAll(in, []byte("{{ .method.variable }}"), []byte(data["method_variable"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.verb }}"), []byte(data["method_verb"]))
	in = bytes.ReplaceAll(in, []byte("{{ .module }}"), []byte(data["module"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.extension }}"), []byte(data["method_extension"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.param }}"), []byte(data["method_param"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.param.value }}"), []byte(data["method_param_value"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.return.type }}"), []byte(data["method_return_type"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.return }}"), []byte(data["method_return"]))
	return in
}
func (m *ModuleInfo) ReplaceImplement(in []byte, data []byte) []byte {
	in = bytes.ReplaceAll(in, []byte("{{ .implement }}"), data)
	return in
}
func (m *ModuleInfo) ReplaceImplementMethod(in []byte, data map[string]string) []byte {
	in = bytes.ReplaceAll(in, []byte("{{ .method.variable }}"), []byte(data["method_variable"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.verb }}"), []byte(data["method_verb"]))
	in = bytes.ReplaceAll(in, []byte("{{ .module }}"), []byte(data["module"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.extension }}"), []byte(data["method_extension"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.param }}"), []byte(data["method_param"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.param.value }}"), []byte(data["method_param_value"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.return.type }}"), []byte(data["method_return_type"]))
	in = bytes.ReplaceAll(in, []byte("{{ .method.return }}"), []byte(data["method_return"]))
	return in
}
func (m *ModuleInfo) GetPackageName() string {
	return m.Package
}

func (m *ModuleInfo) GetModuleName() string {
	return m.Module
}
