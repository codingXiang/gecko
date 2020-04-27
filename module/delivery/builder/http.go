package builder

import (
	"bytes"
	"fmt"
	"github.com/codingXiang/gecko/module/builder"
	"github.com/codingXiang/gecko/parser"
	"github.com/codingXiang/gecko/template/base"
	"github.com/codingXiang/gecko/template/delivery"
	"github.com/codingXiang/gecko/template/delivery/http"
	"github.com/codingXiang/gecko/template/repository"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	HTTP = "http"
)

type HttpBuilder struct {
	builder.ModuleBuilder
	Template delivery.HttpTemplate
}

func NewHttpBuilder(source string, filename string, destination string, pkg string) builder.ModuleBuilderInterface {
	var builder = &HttpBuilder{
		Template: delivery.NewHttpTemplate(),
	}
	builder.Info = base.NewModuleInfo(pkg)
	builder.SetSource(source + "/" + filename)
	builder.SetDestination(destination, pkg)
	builder.SetupFolder()
	return builder
}

func (r *HttpBuilder) SetSource(in string) {
	r.Source = parser.NewParser(in)
}
func (r *HttpBuilder) GetSource() parser.ParserInterface {
	return r.Source
}
func (r *HttpBuilder) SetDestination(in string, name string) {
	r.Destination = in + "/" + strings.ToLower(name)
}
func (r *HttpBuilder) GetDestination() string {
	return r.Destination
}
func (r *HttpBuilder) GetInfo() base.ModuleInfoInterface {
	return r.Info
}

func (r *HttpBuilder) SetupFolder() {
	if err := os.MkdirAll(fmt.Sprintf("%s/%s/%s", r.GetDestination(), delivery.DELIVERY, HTTP), 0777); err == nil {
		log.Println("Create " + HTTP + " folder success")
	} else {
		log.Fatalln("Create " + HTTP + " folder failed")
	}
}
func (r *HttpBuilder) Create() (_interface []byte, _implement []byte, _test []byte) {
	_interface = r.Info.ReplaceTemplate(r.Template.GetInterface())
	_implement = r.Info.ReplaceTemplate(r.Template.GetImplement())
	_test = r.Info.ReplaceTemplate(r.Template.GetTest())
	return
}

func (r *HttpBuilder) General() (_interface []byte, _implement []byte, _test []byte) {
	_interface = r.Info.ReplaceTemplate(r.Template.GetInterface())
	_implement = r.Template.GetImplement()
	_test = r.Info.ReplaceTemplate(r.Template.GetTest())

	_test = bytes.ReplaceAll(_test, []byte("{{ .default.config }}"), http.TEST_CONFIG)

	_interface_abstract_method := []byte("")
	_implement_api := []byte("")
	_implement_method := []byte("")
	_test_method := []byte("")
	for it := range r.Source.GetInterface() {
		for _, m := range it.GetObj().Methods.List {
			var (
				methodName = m.Names[0].Name
			)
			itm := http.INTERFACE_ABSTRACT_METHOD
			itm = bytes.ReplaceAll(itm, []byte("{{ .method }}"), []byte(methodName))
			itm = append(itm, repository.SUBSTITUTION...)

			api := http.API_METHOD
			api = bytes.ReplaceAll(api, []byte("{{ .http.method }}"), []byte(http.GetHttpMethodName(methodName)))
			api = bytes.ReplaceAll(api, []byte("{{ .api.path }}"), []byte(http.GetApiPath(methodName)))
			api = bytes.ReplaceAll(api, []byte("{{ .method }}"), []byte(methodName))
			api = append(api, repository.SUBSTITUTION...)

			im := http.METHOD
			im = bytes.ReplaceAll(im, []byte("{{ .method }}"), []byte(methodName))
			im = append(im, repository.SUBSTITUTION...)

			tm := http.TEST_METHOD
			tm = bytes.ReplaceAll(tm, []byte("{{ .method }}"), []byte(methodName))
			tm = append(tm, repository.SUBSTITUTION...)

			_interface_abstract_method = append(_interface_abstract_method, itm...)
			_implement_api = append(_implement_api, api...)
			_implement_method = append(_implement_method, im...)
			_test_method = append(_test_method, tm...)
		}
	}
	_interface = bytes.ReplaceAll(_interface, []byte("//HttpImplement"), _interface_abstract_method)
	_implement = bytes.ReplaceAll(_implement, []byte("{{ .api.method }}"), _implement_api)
	_implement = bytes.ReplaceAll(_implement, []byte("{{ .implement }}"), _implement_method)
	_implement = r.Info.ReplaceTemplate(_implement)
	_test = bytes.ReplaceAll(_test, []byte("{{ .implement }}"), _test_method)
	//_interface = r.Info.ReplaceImplement(_interface, _interface_abstract_method)
	//_implement = r.Info.ReplaceImplement(_implement, _interface_method)

	return
}

func (r *HttpBuilder) Save(_interface []byte, _implement []byte, _test []byte) {
	var (
		interfacePath = r.GetDestination() + "/" + delivery.DELIVERY + "/"
		implementPath = interfacePath + HTTP + "/"
		fileName      = HTTP + ".go"
		testFileName  = HTTP + "_test.go"
	)

	if err := ioutil.WriteFile(interfacePath+"handler.go", _interface, 0777); err == nil {
		log.Println("Save " + HTTP + " Interface success")
	} else {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile(implementPath+fileName, _implement, 0777); err == nil {
		log.Println("Save " + HTTP + " success")
	} else {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile(implementPath+testFileName, _test, 0777); err == nil {
		log.Println("Save " + HTTP + " Test success")
	} else {
		log.Fatalln(err)
	}
}
