package builder

import (
	"bytes"
	"fmt"
	"github.com/codingXiang/gecko/module/builder"
	"github.com/codingXiang/gecko/parser"
	"github.com/codingXiang/gecko/template/base"
	"github.com/codingXiang/gecko/template/repository"
	"github.com/codingXiang/gecko/template/service"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const SERVICE = "service"

type ServiceBuilder struct {
	builder.ModuleBuilder
	Template service.ServiceTemplate
}

func NewServiceBuilder(source string, filename string, destination string, pkg string) builder.ModuleBuilderInterface {
	var builder = &ServiceBuilder{
		Template: service.NewServiceTemplate(),
	}
	builder.Info = base.NewModuleInfo(pkg)
	builder.SetSource(source + "/" + filename)
	builder.SetDestination(destination, pkg)
	builder.SetupFolder()
	return builder
}

func (r *ServiceBuilder) SetSource(in string) {
	r.Source = parser.NewParser(in)
}
func (r *ServiceBuilder) GetSource() parser.ParserInterface {
	return r.Source
}
func (r *ServiceBuilder) SetDestination(in string, name string) {
	r.Destination = in + "/" + strings.ToLower(name)
}
func (r *ServiceBuilder) GetDestination() string {
	return r.Destination
}
func (r *ServiceBuilder) GetInfo() base.ModuleInfoInterface {
	return r.Info
}

func (r *ServiceBuilder) SetupFolder() {
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", r.GetDestination(), SERVICE), 0777); err == nil {
		log.Println("Create " + SERVICE + " folder success")
	} else {
		log.Fatalln("Create " + SERVICE + " folder failed")
	}
}

func (r *ServiceBuilder) Create() (_interface []byte, _implement []byte, _test []byte) {
	_interface = r.Info.ReplaceTemplate(r.Template.GetInterface())
	_implement = r.Info.ReplaceTemplate(r.Template.GetImplement())
	_test = r.Info.ReplaceTemplate(r.Template.GetTest())
	return
}

func (r *ServiceBuilder) General() (_interface []byte, _implement []byte, _test []byte) {
	_interface = r.Info.ReplaceTemplate(r.Template.GetInterface())
	_implement = r.Info.ReplaceTemplate(r.Template.GetImplement())
	_test = r.Info.ReplaceTemplate(r.Template.GetTest())

	_interface_abstract_method := []byte("")
	_interface_method := []byte("")
	_interface_test := []byte("")
	packageName := r.Source.GetFile().Name.Name
	for it := range r.Source.GetStruct() {
		name := it.GetName().Name
		for _, crud := range repository.CRUD {
			//共用參數
			data := map[string]string{
				"method_variable":    "result, err",
				"module":             name,
				"method_extension":   "",
				"method_verb":        crud,
				"method_param":       "data " + packageName + "." + name + "Interface",
				"method_param_value": "data",
				"method_return_type": "*" + packageName + "." + name + ", error",
				"method_return":      "result, err",
			}
			switch crud {
			case "GetList":
				data["method_verb"] = "Get"
				data["method_extension"] = "List"
				data["method_param"] = "data map[string]interface{}"
				data["method_return_type"] = "[]*" + packageName + "." + name + ", error"
				break;
			case "Modify":
				data["method_param"] = "data " + packageName + "." + name + "Interface, column map[string]interface{}"
				data["method_param_value"] = "data, column"
				break;
			case "Delete":
				data["method_variable"] = "err"
				data["method_return_type"] = "error"
				data["method_return"] = "err"
				break;
			}
			//interface
			{
				itm := service.INTERFACE_ABSTRACT_METHOD
				itm = r.Info.ReplaceTemplateMethod(itm, data)

				itm = append(itm, repository.SUBSTITUTION...)
				_interface_abstract_method = append(_interface_abstract_method, itm...)
			}
			//implement
			{
				itm := service.INTERFACE_METHOD
				//取代方法的 module 名稱
				itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))
				//itm = bytes.ReplaceAll(itm, []byte("{{ .method.variable }}"), service.INTERFACE_METHOD_VARIABLE)
				itm = bytes.ReplaceAll(itm, []byte("{{ .method.action }}"), service.INTERFACE_METHOD_ACTION)
				// method 取代參數

				itm = r.Info.ReplaceImplementMethod(itm, data)
				itm = bytes.ReplaceAll(itm, []byte("{{ .package }}"), []byte(r.Info.GetPackageName()))
				itm = append(itm, repository.SUBSTITUTION...)
				_interface_method = append(_interface_method, itm...)
			}
			//test
			{
				itm := service.INTERFACE_TEST
				itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))
				//取代方法的 module 名稱
				if crud == "GetList" {
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte("Get"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.extension }}"), []byte("List"))
				} else {
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.extension }}"), []byte(""))
				}

				itm = append(itm, repository.SUBSTITUTION...)
				_interface_test = append(_interface_test, itm...)
			}
		}
	}
	_interface = r.Info.ReplaceImplement(_interface, _interface_abstract_method)
	_implement = r.Info.ReplaceImplement(_implement, _interface_method)
	_test = r.Info.ReplaceImplement(_test, _interface_test)

	return
}
func (r *ServiceBuilder) Save(_interface []byte, _implement []byte, _test []byte) {
	var (
		interfacePath = r.GetDestination() + "/"
		implementPath = interfacePath + SERVICE + "/"
		fileName      = SERVICE + ".go"
		testFileName  = SERVICE + "_test.go"
	)
	if err := ioutil.WriteFile(interfacePath+fileName, _interface, 0777); err == nil {
		log.Println("Save " + SERVICE + " Interface success")
	} else {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile(implementPath+fileName, _implement, 0777); err == nil {
		log.Println("Save " + SERVICE + " success")
	} else {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile(implementPath+testFileName, _test, 0777); err == nil {
		log.Println("Save " + SERVICE + " Test success")
	} else {
		log.Fatalln(err)
	}
}
