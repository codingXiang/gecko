package http

import (
	"strings"
)

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
		svc:     svc,
	}
	/*
		v1 版本的 {{ .module }} API
	  */
	v1 := gateway.GetApiRoute().Group("/v1")
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

var (
	TEST = []byte(`package http_test

import (
	"fmt"
	cx_delivery "github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/delivery/http"
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/codingXiang/cxgateway/pkg/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	http3 "net/http"
	"testing"
)

//Suite 為設定 mock repository
type Suite struct {
	suite.Suite
	gateway     cx_delivery.HttpHandler
	tester      util.HttpTesterInterface
	httpHandler delivery.HttpHandler
}

//測試變數


//初始化 Suite
func (s *Suite) SetupSuite() {
	//建立 Api Gateway，自定義 config
	s.gateway = http.NewApiGatewayWithData("config", []byte({{ .default.config }}))
	//設定 tester
	s.tester = util.NewHttpTester(s.gateway.GetEngine())
	// 建立 mock controller
	ctrl := gomock.NewController(s.T())
	// 透過 mock 建立 service
	service := mock.NewMockService(ctrl)
	//設定 mock 資料
	{
		//mock service
	}
	// 設定 http handler
	s.httpHandler = http.New{{ .module }}HttpHandler(s.gateway, service)
}

//TestStart 為測試程式進入點
func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

{{ .implement }}
`)
	TEST_METHOD = []byte(`//Test{{ .method }} 為測試 HttpHandler 中的 {{ .method }} 方法
func (s *Suite) Test{{ .method }}() {
    panic("implement me")
}`)
	TEST_CONFIG = []byte("`" + `
application:
  timeout:
    read: 1000
    write: 1000
  port: 8080
  mode: "release"
  log:
    level: "error"
    format: "text"
  appId: "app"
  appToken: ""
  apiBaseRoute: "/api"
i18n:
  defaultLanguage: "zh_Hant"
  file:
    path: "../../../../i18n"
    type: "yaml"
` + "`")
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
