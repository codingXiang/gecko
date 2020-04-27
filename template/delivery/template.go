package delivery

const DELIVERY = "delivery"

// Base Template
var INTERFACE = []byte(`package delivery
import "github.com/gin-gonic/gin"

//HttpHandler http流量 handler
type HttpHandler interface {
//HttpImplement
}

//GRPCHandler gRPC流量 handler
type GRPCHandler interface {
//gRpcImplement
}

//CmdHandler cli handler
type CmdHandler interface {
//CmdImplement
}
`)