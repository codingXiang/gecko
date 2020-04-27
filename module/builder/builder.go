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

//func (m *RepositoryBuilder) Create() {
//	m.loadTemplate("repository")
//	m.replaceTemplate("repository")
//	m.save("repository")
//}
//
//func (m *RepositoryBuilder) CreateService() {
//	m.loadTemplate("service")
//	m.replaceTemplate("service")
//	m.save("service")
//}
//func (m *RepositoryBuilder) init() {
//	if err := os.MkdirAll(fmt.Sprintf("%s/mock", m.Path), 0777); err == nil {
//		log.Println("Create mock folder success")
//	} else {
//		log.Fatalln("Create mock folder failed")
//	}
//	if err := os.MkdirAll(fmt.Sprintf("%s/repository", m.Path), 0777); err == nil {
//		log.Println("Create repository folder success")
//	} else {
//		log.Fatalln("Create repository folder failed")
//	}
//	if err := os.MkdirAll(fmt.Sprintf("%s/service", m.Path), 0777); err == nil {
//		log.Println("Create service folder success")
//	} else {
//		log.Fatalln("Create service folder failed")
//	}
//	log.Println("Create delivery folder")
//	if err := os.MkdirAll(fmt.Sprintf("%s/delivery/http", m.Path), 0777); err == nil {
//		log.Println("Create delivery http folder success")
//	} else {
//		log.Fatalln("Create delivery http folder failed")
//	}
//	if err := os.MkdirAll(fmt.Sprintf("%s/delivery/cmd", m.Path), 0777); err == nil {
//		log.Println("Create delivery cmd folder success")
//	} else {
//		log.Fatalln("Create delivery cmd folder failed")
//	}
//	if err := os.MkdirAll(fmt.Sprintf("%s/delivery/grpc", m.Path), 0777); err == nil {
//		log.Println("Create delivery grpc folder success")
//	} else {
//		log.Fatalln("Create delivery grpc folder failed")
//	}
//}
//
//func (m *RepositoryBuilder) loadTemplate(_type string) {
//	var (
//		err error
//	)
//	if _type == "repository" {
//		if m.Template.Repository.Interface, err = ioutil.ReadFile("./template/repository/interface.template"); err == nil {
//			log.Println("Load Repository interface template success")
//		} else {
//			log.Fatal("Template Repository interface load failed")
//		}
//		if m.Template.Repository.Implement, err = ioutil.ReadFile("./template/repository/implement.template"); err == nil {
//			log.Println("Load Repository implement template success")
//		} else {
//			log.Fatal("Template Repository implement load failed")
//		}
//		if m.Template.Repository.Test, err = ioutil.ReadFile("./template/repository/test.template"); err == nil {
//			log.Println("Load Repository test template success")
//		} else {
//			log.Fatal("Template Repository test load failed")
//		}
//	} else if _type == "service" {
//		if m.Template.Service.Interface, err = ioutil.ReadFile("./template/service/interface.template"); err == nil {
//			log.Println("Load Service interface template success")
//		} else {
//			log.Fatal("Template Service interface load failed")
//		}
//		if m.Template.Service.Implement, err = ioutil.ReadFile("./template/service/implement.template"); err == nil {
//			log.Println("Load Service implement template success")
//		} else {
//			log.Fatal("Template Service implement load failed")
//		}
//		if m.Template.Service.Test, err = ioutil.ReadFile("./template/service/test.template"); err == nil {
//			log.Println("Load Service test template success")
//		} else {
//			log.Fatal("Template Service test load failed")
//		}
//	}
//
//}
//
//func (m *RepositoryBuilder) replaceTemplate(_type string) {
//	switch _type {
//	case "repository":
//		m.Template.Repository.Interface = bytes.ReplaceAll(m.Template.Repository.Interface, []byte("{{package}}"), []byte(m.PackageName))
//		m.Template.Repository.Implement = bytes.ReplaceAll(m.Template.Repository.Implement, []byte("{{package}}"), []byte(m.PackageName))
//		m.Template.Repository.Test = bytes.ReplaceAll(m.Template.Repository.Test, []byte("{{package}}"), []byte(m.PackageName))
//		m.Template.Repository.Interface = bytes.ReplaceAll(m.Template.Repository.Interface, []byte("{{module}}"), []byte(m.ModuleName))
//		m.Template.Repository.Implement = bytes.ReplaceAll(m.Template.Repository.Implement, []byte("{{module}}"), []byte(m.ModuleName))
//		m.Template.Repository.Test = bytes.ReplaceAll(m.Template.Repository.Test, []byte("{{module}}"), []byte(m.ModuleName))
//		break
//	case "service":
//		m.Template.Service.Interface = bytes.ReplaceAll(m.Template.Service.Interface, []byte("{{package}}"), []byte(m.PackageName))
//		m.Template.Service.Implement = bytes.ReplaceAll(m.Template.Service.Implement, []byte("{{package}}"), []byte(m.PackageName))
//		m.Template.Service.Test = bytes.ReplaceAll(m.Template.Service.Test, []byte("{{package}}"), []byte(m.PackageName))
//		m.Template.Service.Interface = bytes.ReplaceAll(m.Template.Service.Interface, []byte("{{module}}"), []byte(m.ModuleName))
//		m.Template.Service.Implement = bytes.ReplaceAll(m.Template.Service.Implement, []byte("{{module}}"), []byte(m.ModuleName))
//		m.Template.Service.Test = bytes.ReplaceAll(m.Template.Service.Test, []byte("{{module}}"), []byte(m.ModuleName))
//		break
//	case "delivery":
//		break
//	default:
//		log.Fatalln("type " + _type + " is not support")
//		break;
//	}
//
//}
//func (m *ModuleBuilder) save(_type string) {
//	var (
//		interfacePath = m.Path + "/"
//	)
//	if _type == "repository" {
//		implementPath := interfacePath + "repository/"
//		if err := ioutil.WriteFile(interfacePath+"repository.go", m.Template.Repository.Interface, 0777); err == nil {
//			log.Println("Save Repository Interface success")
//		} else {
//			log.Fatalln(err)
//		}
//		if err := ioutil.WriteFile(implementPath+"repository.go", m.Template.Repository.Implement, 0777); err == nil {
//			log.Println("Save Repository Implement success")
//		} else {
//			log.Fatalln(err)
//		}
//		if err := ioutil.WriteFile(implementPath+"repository_test.go", m.Template.Repository.Test, 0777); err == nil {
//			log.Println("Save Repository Test success")
//		} else {
//			log.Fatalln(err)
//		}
//	} else if _type == "service" {
//		implementPath := interfacePath + "service/"
//		if err := ioutil.WriteFile(interfacePath+"service.go", m.Template.Service.Interface, 0777); err == nil {
//			log.Println("Save Service Interface success")
//		} else {
//			log.Fatalln(err)
//		}
//		if err := ioutil.WriteFile(implementPath+"service.go", m.Template.Service.Implement, 0777); err == nil {
//			log.Println("Save Service Implement success")
//		} else {
//			log.Fatalln(err)
//		}
//		if err := ioutil.WriteFile(implementPath+"service_test.go", m.Template.Service.Test, 0777); err == nil {
//			log.Println("Save Service Test success")
//		} else {
//			log.Fatalln(err)
//		}
//	}
//}
