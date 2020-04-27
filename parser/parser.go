package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
)

type (
	ParserInterface interface {
		GetSource() string
		GetFile() *ast.File
		GetRawData() []byte
		GetInterface() <-chan ItInterface
		GetStruct() <-chan StInterface
	}
	ItInterface interface {
		GetName() *ast.Ident
		GetObj() *ast.InterfaceType
	}
	StInterface interface {
		GetName() *ast.Ident
		GetObj() *ast.StructType
		GetFields(raw []byte) <-chan FieldInterface
	}
)

type (
	Parser struct {
		_source    string
		_fileSet   *token.FileSet
		_file      *ast.File
		_rawData   []byte
		_interface chan ItInterface
		_struct    chan StInterface
	}
	namedInterface struct {
		_name *ast.Ident
		_obj  *ast.InterfaceType
	}
	namedStruct struct {
		_name   *ast.Ident
		_fields []*ast.Field
		obj     *ast.StructType
	}
)

func NewParser(source string) ParserInterface {
	p := &Parser{
		_source:  source,
		_fileSet: token.NewFileSet(),
	}
	var err error
	if raw, err := ioutil.ReadFile(source); err == nil {
		p._rawData = raw
	} else {
		log.Fatalln(err)
	}
	if p._file, err = parser.ParseFile(p._fileSet, p._source, nil, 0); err != nil {
		log.Fatal("read file " + p._source + " failed, err = " + err.Error())
	}
	return p
}

func (p *Parser) GetSource() string {
	return p._source
}

func (p *Parser) GetFile() *ast.File {
	return p._file
}
func (p *Parser) GetRawData() []byte {
	return p._rawData
}
func (p *Parser) GetInterface() <-chan ItInterface {
	ch := make(chan ItInterface)
	go func() {
		for _, decl := range p.GetFile().Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok || gd.Tok != token.TYPE {
				continue
			}
			for _, spec := range gd.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				it, ok := ts.Type.(*ast.InterfaceType)
				if !ok {
					continue
				}

				ch <- &namedInterface{ts.Name, it}
			}
		}
		close(ch)
	}()
	return ch
}

func (p *Parser) GetStruct() <-chan StInterface {
	ch := make(chan StInterface)
	go func() {
		for _, decl := range p.GetFile().Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok || gd.Tok != token.TYPE {
				continue
			}
			for _, spec := range gd.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				st, ok := ts.Type.(*ast.StructType)
				if !ok {
					continue
				}

				ch <- &namedStruct{ts.Name, st.Fields.List, st}
			}
		}
		close(ch)
	}()
	return ch
}

func (it *namedInterface) GetName() *ast.Ident {
	return it._name
}
func (it *namedInterface) GetObj() *ast.InterfaceType {
	return it._obj
}

func (st *namedStruct) GetName() *ast.Ident {
	return st._name
}
func (st *namedStruct) GetObj() *ast.StructType {
	return st.obj
}
func (st *namedStruct) GetFields(raw []byte) <-chan FieldInterface {
	ch := make(chan FieldInterface)
	go func() {
		for _, f := range st._fields {
			ch <- &Field{
				_name: f.Names[0].String(),
				_type: string(raw[f.Type.Pos()-1 : f.Type.End()-1]),
			}
		}
		close(ch)
	}()
	return ch
}

type (
	FieldInterface interface {
		GetName() string
		GetType() string
	}
	Field struct {
		_name string
		_type string
	}
)

func (f *Field) GetName() string {
	return f._name
}
func (f *Field) GetType() string {
	return f._type
}
