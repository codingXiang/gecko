package delivery

var (
	INTERFACE_ABSTRACT_METHOD = []byte(`{{verb}}{{model}}(c *gin.Context) error`)
	INTERFACE = []byte(`/*
Auto Create By Moduler
{{package}} 模組的 Handler Interface
*/
package {{package}}
import "github.com/gin-gonic/gin"

//HttpHandler http流量 handler
type HttpHandler interface {
	    /*
        寫入封裝方法
    */
{{implement}}
}`)
)