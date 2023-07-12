package controller

import (
	"github.com/gin-gonic/gin"
	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/rowdata"
)

// 自定义中间件函数
func GrpcMiddleware(c *gin.Context, client pb.RowDataServiceClient) {
	// 在这里执行中间件逻辑
	c.Set("grpc-client", client)
	c.Next()
}
