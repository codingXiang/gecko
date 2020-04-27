package builder

import (
	"bytes"
	"github.com/codingXiang/gecko/parser"
	"github.com/codingXiang/gecko/template/model"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ModelBuilderInterface interface {
	General()
	Save()
}

type ModelBuilder struct {
	_source      string
	_destination string
	_filename    string
	_parser      parser.ParserInterface
	_rawData     []byte
}

func NewModelBuilder(source string, destination string, filename string) ModelBuilderInterface {
	m := &ModelBuilder{
		_source:   source,
		_parser:   parser.NewParser(source + "/" + filename),
		_filename: filename,
	}

	if destination == "" {
		m._destination = m._source
	} else {
		m._destination = destination
	}

	return m
}

func (m *ModelBuilder) General() {
	var (
		inter = []byte(``)
	)
	mt := model.SUBSTITUTION
	for st := range m._parser.GetStruct() {
		typeT := model.INTERFACE_TYPE
		typeT = bytes.ReplaceAll(typeT, []byte("{{ .module }}"), []byte(st.GetName().Name))
		amt := model.SUBSTITUTION
		for f := range st.GetFields(m._parser.GetRawData()) {
			var (
				name  = f.GetName()
				_type = f.GetType()
			)
			// Getter
			{
				var verb = "Get"
				if f.GetType() == "bool" {
					verb = "Is"
				}
				//abstruct method
				{
					tmp := model.INTERFACE_ABSTRUCT_METHOD
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.verb }}"), []byte(verb))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .module }}"), []byte(strings.Title(name)))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.param }}"), []byte(``))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.return.type }}"), []byte(_type))
					amt = append(amt, tmp...)
					amt = append(amt, model.SUBSTITUTION...)
				}
				//implement method
				{
					tmp := model.INTERFACE_METHOD
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .struct }}"), []byte(st.GetName().Name))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.verb }}"), []byte(verb))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .module }}"), []byte(strings.Title(name)))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.param }}"), []byte(``))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.return.type }}"), []byte(_type))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.action }}"), []byte(``))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.return }}"), []byte("return g."+name))
					mt = append(mt, tmp...)
					mt = append(mt, model.SUBSTITUTION...)
				}
			}
			// Setter
			{
				//abstruct method
				{
					tmp := model.INTERFACE_ABSTRUCT_METHOD
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.verb }}"), []byte("Set"))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .module }}"), []byte(strings.Title(name)))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.param }}"), []byte("in "+f.GetType()))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.return.type }}"), []byte(`*`+st.GetName().Name))
					amt = append(amt, tmp...)
					amt = append(amt, model.SUBSTITUTION...)
				}
				//implement method
				{
					tmp := model.INTERFACE_METHOD
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .struct }}"), []byte(st.GetName().Name))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.verb }}"), []byte("Set"))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .module }}"), []byte(strings.Title(name)))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.param }}"), []byte("in "+f.GetType()))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.return.type }}"), []byte(`*`+st.GetName().Name))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.action }}"), []byte(`g.`+name+" = in"))
					tmp = bytes.ReplaceAll(tmp, []byte("{{ .method.return }}"), []byte("return g"))
					mt = append(mt, tmp...)
					mt = append(mt, model.SUBSTITUTION...)
				}
			}
		}
		n_str := model.INTERFACE_NEW
		//implement method
		{
			n_str = bytes.ReplaceAll(n_str, []byte("{{ .struct }}"), []byte(st.GetName().Name))
			mt = append(n_str, mt...)
		}

		typeT = bytes.ReplaceAll(typeT, []byte("{{ .interface.method }}"), amt)
		inter = append(inter, typeT...)
	}
	m._rawData = bytes.ReplaceAll(m._parser.GetRawData(), []byte("//exdev_interface_type"), inter)
	m._rawData = bytes.ReplaceAll(m._rawData, []byte("//exdev_interface_method"), mt)
}

func (m *ModelBuilder) Save() {
	if err := os.MkdirAll(m._destination, 0777); err == nil {
		log.Println("Create model folder" + m._destination + " success")
	} else {
		log.Println("Create model folder" + m._destination + " failed")
	}
	if err := ioutil.WriteFile(m._destination+"/"+m._filename, m._rawData, 0777); err == nil {
		log.Println("Save model " + m._destination + "/" + m._filename + " success")
	} else {
		log.Println("Save model " + m._destination + "/" + m._filename + " failed")
	}
}
