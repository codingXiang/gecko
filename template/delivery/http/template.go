package http

import "strings"

//Interface
var (
	INTERFACE_ABSTRACT_METHOD = []byte(`	{{ .method }}(c *gin.Context) error`)
)

var (
	IMPLEMENT = []byte(`package http
import (
	"github.com/gin-gonic/gin"
    cx "github.com/codingXiang/cxgateway/delivery"
)

type {{ .module }}HttpHandler struct {
	gateway cx.HttpHandler
	svc     {{ .package }}.Service
}

func New{{ .module }}HttpHandler(gateway cx.HttpHandler, svc {{ .package }}.Service) delivery.HttpHandler {
	var handler = &{{ .module }}HttpHandler{
		gateway: gateway,
		svc:     {{ .package }}.Service,
	}
	/*
		v1 版本的 {{ .module }} API
	  */
	v1 := gateway.GetApiRoute().Group("/v1/{{ .package }}")
{{ .api.method }}
	return handler
}

{{ .implement }}
`)
	API_METHOD = []byte(`	v1.{{ .http.method }}("{{ .api.path }}", e.Wrapper(handler.{{ .method }}))`)
	METHOD     = []byte(`func (g *{{ .module }}HttpHandler) {{ .method }}(c *gin.Context) error {
	panic("implement me")
}`)
)

func GetHttpMethodName(in string) string {
	if strings.Contains(in, "Create") {
		return "POST"
	}
	if strings.Contains(in, "Update") {
		return "PUT"
	}
	if strings.Contains(in, "Modify") {
		return "PATCH"
	}
	if strings.Contains(in, "Delete") {
		return "DELETE"
	}
	return "GET"
}

func GetApiPath(in string) string {
	result := "/{{ .package }}"
	if strings.Contains(in, "List") || strings.Contains(in, "Create") {
		return result
	} else {
		result += "/:id"
		return result
	}

}
