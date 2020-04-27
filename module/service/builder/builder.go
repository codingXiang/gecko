package builder

import (
	"bytes"
	"fmt"
	"github.com/codingXiang/gecko/module/builder"
	"github.com/codingXiang/gecko/parser"
	"github.com/codingXiang/gecko/template/SERVICE"
	"github.com/codingXiang/gecko/template/base"
	"github.com/codingXiang/gecko/template/repository"
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
			//interface
			{
				itm := service.INTERFACE_ABSTRACT_METHOD
				if crud == "GetList" {
					itm = bytes.ReplaceAll(itm, []byte("{{verb}}"), []byte("Get"))
					itm = bytes.ReplaceAll(itm, []byte("{{model}}"), []byte(name+"List"))
				} else {
					itm = bytes.ReplaceAll(itm, []byte("{{verb}}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{model}}"), []byte(name))
				}
				if crud == "Modify" {
					itm = bytes.ReplaceAll(itm, []byte("{{param}}"), []byte("model "+packageName+"."+name+"Interface, data map[string]interface{}"))
				} else {
					itm = bytes.ReplaceAll(itm, []byte("{{param}}"), []byte("data "+packageName+"."+name+"Interface"))
				}
				if crud == "Delete" {
					itm = bytes.ReplaceAll(itm, []byte("{{type}}"), []byte("error"))
				} else {
					itm = bytes.ReplaceAll(itm, []byte("{{type}}"), []byte(packageName+"."+name+" ,error"))
				}
				itm = append(itm, repository.SUBSTITUTION...)
				_interface_abstract_method = append(_interface_abstract_method, itm...)
			}
			//implement
			{
				itm := service.INTERFACE_METHOD
				//取代方法的 module 名稱
				itm = bytes.ReplaceAll(itm, []byte("{{module}}"), []byte(name))
				itm = bytes.ReplaceAll(itm, []byte("{{variable}}"), repository.INTERFACE_METHOD_VARIABLE)
				itm = bytes.ReplaceAll(itm, []byte("{{action}}"), repository.INTERFACE_METHOD_ACTION)

				switch crud {
				case "GetList":
					itm = bytes.ReplaceAll(itm, []byte("{{variable}}"), []byte("make([]*"+packageName+"."+name+", 0)"))
					itm = bytes.ReplaceAll(itm, []byte("{{verb}}"), []byte("Get"))
					itm = bytes.ReplaceAll(itm, []byte("{{model}}"), []byte(name+"List"))
					itm = bytes.ReplaceAll(itm, []byte("{{param}}"), []byte("data map[string]interface{}"))
					itm = bytes.ReplaceAll(itm, []byte("{{type}}"), []byte("[]*"+packageName+"."+name+", error"))
					itm = bytes.ReplaceAll(itm, []byte("{{action}}"), []byte("Find(&result, data)"))
					itm = bytes.ReplaceAll(itm, []byte("{{return}}"), []byte("return in, err"))

					break;
				case "Get":
					itm = bytes.ReplaceAll(itm, []byte("{{variable}}"), []byte("data.(*"+packageName+"."+name+")"))
					itm = bytes.ReplaceAll(itm, []byte("{{verb}}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{model}}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{param}}"), []byte("data "+packageName+"."+name+"Interface"))
					itm = bytes.ReplaceAll(itm, []byte("{{type}}"), []byte("*"+packageName+"."+name+", error"))
					itm = bytes.ReplaceAll(itm, []byte("{{action}}"), []byte("First(&in)"))
					itm = bytes.ReplaceAll(itm, []byte("{{return}}"), []byte("return in, err"))
					break;
				case "Create":
					itm = bytes.ReplaceAll(itm, []byte("{{variable}}"), []byte("data.(*"+packageName+"."+name+")"))

					itm = bytes.ReplaceAll(itm, []byte("{{verb}}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{model}}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{param}}"), []byte("data "+packageName+"."+name+"Interface"))
					itm = bytes.ReplaceAll(itm, []byte("{{type}}"), []byte("*"+packageName+"."+name+", error"))
					itm = bytes.ReplaceAll(itm, []byte("{{action}}"), []byte("Create(&in)"))
					itm = bytes.ReplaceAll(itm, []byte("{{return}}"), []byte("return in, err"))

					break;
				case "Modify":
					itm = bytes.ReplaceAll(itm, []byte("{{variable}}"), []byte("data.(*"+packageName+"."+name+")"))

					itm = bytes.ReplaceAll(itm, []byte("{{verb}}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{model}}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{param}}"), []byte("data "+packageName+"."+name+"Interface, column map[string]interface{}"))
					itm = bytes.ReplaceAll(itm, []byte("{{type}}"), []byte("*"+packageName+"."+name+", error"))
					itm = bytes.ReplaceAll(itm, []byte("{{action}}"), []byte("Model(&in).Updates(column)"))
					itm = bytes.ReplaceAll(itm, []byte("{{return}}"), []byte("return in, err"))

					break;
				case "Update":
					itm = bytes.ReplaceAll(itm, []byte("{{variable}}"), []byte("data.(*"+packageName+"."+name+")"))

					itm = bytes.ReplaceAll(itm, []byte("{{verb}}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{model}}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{param}}"), []byte("data "+packageName+"."+name+"Interface"))
					itm = bytes.ReplaceAll(itm, []byte("{{type}}"), []byte("*"+packageName+"."+name+", error"))
					itm = bytes.ReplaceAll(itm, []byte("{{action}}"), []byte("Update(&in)"))
					itm = bytes.ReplaceAll(itm, []byte("{{return}}"), []byte("return in, err"))

					break;
				case "Delete":
					itm = bytes.ReplaceAll(itm, []byte("{{variable}}"), []byte("data.(*"+packageName+"."+name+")"))

					itm = bytes.ReplaceAll(itm, []byte("{{verb}}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{model}}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{param}}"), []byte("data "+packageName+"."+name+"Interface"))
					itm = bytes.ReplaceAll(itm, []byte("{{type}}"), []byte("error"))
					itm = bytes.ReplaceAll(itm, []byte("{{action}}"), []byte("Delete(&in)"))
					itm = bytes.ReplaceAll(itm, []byte("{{return}}"), []byte("return err"))

					break;
				}

				itm = append(itm, repository.SUBSTITUTION...)
				_interface_method = append(_interface_method, itm...)
			}
			//test
			{
				itm := service.INTERFACE_TEST
				//取代方法的 module 名稱
				if crud == "GetList" {
					itm = bytes.ReplaceAll(itm, []byte("{{verb}}"), []byte("Get"))
					itm = bytes.ReplaceAll(itm, []byte("{{model}}"), []byte(name+"List"))
				} else {
					itm = bytes.ReplaceAll(itm, []byte("{{verb}}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{model}}"), []byte(name))
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
