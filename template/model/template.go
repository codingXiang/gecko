package model

var (
	INTERFACE_TYPE = []byte(`//此為自動產生的 Interface，建議不要進行更動
type {{ .module }}Interface interface {{{ .interface.method }}}
`)
	INTERFACE_ABSTRUCT_METHOD = []byte(`	{{ .method.verb }}{{ .module }}({{ .method.param }}) {{ .method.return.type }}`)
	INTERFACE_NEW = []byte(`//此為自動產生的方法，建議不要更動
func New{{ .struct }}() {{ .struct }}Interface {
	return &{{ .struct }}{}
}
`)
	INTERFACE_METHOD = []byte(`//此為自動產生的方法，建議不要更動
func (g *{{ .struct }}) {{ .method.verb }}{{ .module }}({{ .method.param }}) {{ .method.return.type }} {
    {{ .method.action }}
    {{ .method.return }}
}
`)
	SUBSTITUTION = []byte(`
`)
)
