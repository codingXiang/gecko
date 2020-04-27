package builder

import (
	"bytes"
	"fmt"
	"github.com/codingXiang/gecko/module/builder"
	"github.com/codingXiang/gecko/parser"
	"github.com/codingXiang/gecko/template/base"
	"github.com/codingXiang/gecko/template/repository"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const REPOSITORY = "repository"

type RepositoryBuilder struct {
	builder.ModuleBuilder
	Template repository.RepositoryTemplate
}

func NewRepositoryBuilder(source string, filename string, destination string, pkg string) builder.ModuleBuilderInterface {
	var builder = &RepositoryBuilder{
		Template: repository.NewRepositoryTemplate(),
	}
	builder.Info = base.NewModuleInfo(pkg)
	builder.SetSource(source + "/" + filename)
	builder.SetDestination(destination, pkg)
	builder.SetupFolder()
	return builder
}

func (r *RepositoryBuilder) SetSource(in string) {
	r.Source = parser.NewParser(in)
}
func (r *RepositoryBuilder) GetSource() parser.ParserInterface {
	return r.Source
}
func (r *RepositoryBuilder) SetDestination(in string, name string) {
	r.Destination = in + "/" + strings.ToLower(name)
}
func (r *RepositoryBuilder) GetDestination() string {
	return r.Destination
}
func (r *RepositoryBuilder) GetInfo() base.ModuleInfoInterface {
	return r.Info
}
func (r *RepositoryBuilder) SetupFolder() {
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", r.GetDestination(), REPOSITORY), 0777); err == nil {
		log.Println("Create " + REPOSITORY + " folder success")
	} else {
		log.Fatalln("Create " + REPOSITORY + " folder failed")
	}
}

func (r *RepositoryBuilder) Create() (_interface []byte, _implement []byte, _test []byte) {
	_interface = r.Info.ReplaceTemplate(r.Template.GetInterface())
	_interface = r.Info.ReplaceImplement(_interface, []byte(""))
	_implement = r.Info.ReplaceTemplate(r.Template.GetImplement())
	_implement = r.Info.ReplaceImplement(_implement, []byte(""))
	_test = r.Info.ReplaceTemplate(r.Template.GetTest())
	_test = r.Info.ReplaceImplement(_test, []byte(""))
	return
}
func (r *RepositoryBuilder) General() (_interface []byte, _implement []byte, _test []byte) {
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
				itm := repository.INTERFACE_ABSTRACT_METHOD
				itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))

				if crud == "GetList" {
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte("Get"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.extension }}"), []byte("List"))
				} else {
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.extension }}"), []byte(""))
				}
				if crud == "Modify" {
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.param }}"), []byte("model "+packageName+"."+name+"Interface, data map[string]interface{}"))
				} else {
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.param }}"), []byte("data "+packageName+"."+name+"Interface"))
				}
				if crud == "Delete" {
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return.type }}"), []byte("error"))
				} else {
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return.type }}"), []byte(packageName+"."+name+" ,error"))
				}
				itm = append(itm, repository.SUBSTITUTION...)
				_interface_abstract_method = append(_interface_abstract_method, itm...)
			}
			//implement
			{
				itm := repository.INTERFACE_METHOD
				//取代方法的 module 名稱
				itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))
				itm = bytes.ReplaceAll(itm, []byte("{{ .module.variable }}"), repository.INTERFACE_METHOD_VARIABLE)
				itm = bytes.ReplaceAll(itm, []byte("{{ .method.action }}"), repository.INTERFACE_METHOD_ACTION)

				switch crud {
				case "GetList":
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.variable }}"), []byte("make([]*"+packageName+"."+name+", 0)"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte("Get"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.extension }}"), []byte("List"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.param }}"), []byte("data map[string]interface{}"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return.type }}"), []byte("[]*"+packageName+"."+name+", error"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.action }}"), []byte("Find(&result, data)"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return }}"), []byte("return in, err"))

					break;
				case "Get":
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.variable }}"), []byte("data.(*"+packageName+"."+name+")"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.extension }}"), []byte(""))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.param }}"), []byte("data "+packageName+"."+name+"Interface"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return.type }}"), []byte("*"+packageName+"."+name+", error"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.action }}"), []byte("First(&in)"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return }}"), []byte("return in, err"))
					break;
				case "Create":
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.variable }}"), []byte("data.(*"+packageName+"."+name+")"))

					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.extension }}"), []byte(""))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.param }}"), []byte("data "+packageName+"."+name+"Interface"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return.type }}"), []byte("*"+packageName+"."+name+", error"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.action }}"), []byte("Create(&in)"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return }}"), []byte("return in, err"))

					break;
				case "Modify":
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.variable }}"), []byte("data.(*"+packageName+"."+name+")"))

					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.extension }}"), []byte(""))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.param }}"), []byte("data "+packageName+"."+name+"Interface, column map[string]interface{}"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return.type }}"), []byte("*"+packageName+"."+name+", error"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.action }}"), []byte("Model(&in).Updates(column)"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return }}"), []byte("return in, err"))

					break;
				case "Update":
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.variable }}"), []byte("data.(*"+packageName+"."+name+")"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.extension }}"), []byte(""))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.param }}"), []byte("data "+packageName+"."+name+"Interface"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return.type }}"), []byte("*"+packageName+"."+name+", error"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.action }}"), []byte("Update(&in)"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return }}"), []byte("return in, err"))

					break;
				case "Delete":
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.variable }}"), []byte("data.(*"+packageName+"."+name+")"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.extension }}"), []byte(""))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.param }}"), []byte("data "+packageName+"."+name+"Interface"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return.type }}"), []byte("error"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.action }}"), []byte("Delete(&in)"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.return }}"), []byte("return err"))

					break;
				}

				itm = append(itm, repository.SUBSTITUTION...)
				_interface_method = append(_interface_method, itm...)
			}
			//test
			{
				itm := repository.INTERFACE_TEST
				//取代方法的 module 名稱
				if crud == "GetList" {
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte("Get"))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module.extension }}"), []byte("List"))
				} else {
					itm = bytes.ReplaceAll(itm, []byte("{{ .method.verb }}"), []byte(crud))
					itm = bytes.ReplaceAll(itm, []byte("{{ .module }}"), []byte(name))
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
func (r *RepositoryBuilder) Save(_interface []byte, _implement []byte, _test []byte) {
	var (
		interfacePath = r.GetDestination() + "/"
		implementPath = interfacePath + REPOSITORY + "/"
		fileName      = REPOSITORY + ".go"
		testFileName  = REPOSITORY + "_test.go"
	)
	if err := ioutil.WriteFile(interfacePath+fileName, _interface, 0777); err == nil {
		log.Println("Save " + REPOSITORY + " Interface success")
	} else {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile(implementPath+fileName, _implement, 0777); err == nil {
		log.Println("Save " + REPOSITORY + " success")
	} else {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile(implementPath+testFileName, _test, 0777); err == nil {
		log.Println("Save " + REPOSITORY + " Test success")
	} else {
		log.Fatalln(err)
	}
}
