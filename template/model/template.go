package model

var (
	INTERFACE_TYPE = []byte(`//此為自動產生的 Interface，建議不要進行更動
type {{name}}Interface interface {{{abstruct_method}}}
`)
	INTERFACE_ABSTRUCT_METHOD = []byte(`	{{verb}}{{name}}({{param}}) {{type}}`)
	INTERFACE_NEW = []byte(`//此為自動產生的方法，建議不要更動
func New{{struct}}() {{struct}}Interface {
	return &{{struct}}{}
}
`)
	INTERFACE_METHOD = []byte(`//此為自動產生的方法，建議不要更動
func (g *{{struct}}) {{verb}}{{name}}({{param}}) {{type}} {
    {{action}}
    {{return}}
}
`)
	SUBSTITUTION = []byte(`
`)
)
