package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata-api/internal/controller"
	"github.com/zizhuxuexi/datacenter-rowing/config"
	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/rowdata"

	"google.golang.org/grpc"
)

func main() {

	config.InitConfig()
	rowDataGrpcAddr := config.Conf.Services["row-data"].Addr
	// 创建 gRPC 连接
	grpcConn, err := grpc.Dial(rowDataGrpcAddr, grpc.WithInsecure())
	if err != nil {
		// 错误处理
		panic(err)
	}
	client := pb.NewRowDataServiceClient(grpcConn)

	// 创建控制器并传递 gRPC 连接
	rowDataController := controller.NewRowDataController(grpcConn, client)
	defer grpcConn.Close()

	router := gin.Default()
	//router.MaxMultipartMemory = 8 << 20

	router.POST("/rowdata", func(ctx *gin.Context) {
		rowDataController.UploadMultiRowData(ctx, client) // 调用方法
	})
	router.GET("/rowdata", func(ctx *gin.Context) {
		rowDataController.FindTrainingSummary(ctx, client)
	})
	router.GET("/athleteTrainingData", func(ctx *gin.Context) {
		rowDataController.FindAthleteTrainingData(ctx, client)
	})
	router.GET("/sampleMetrics", func(ctx *gin.Context) {
		rowDataController.FindSampleMetricsByAthleteTrainingId(ctx, client)
	})

	router.Run(":10002")

}
