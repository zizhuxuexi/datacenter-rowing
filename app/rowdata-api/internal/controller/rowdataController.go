package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata-api/internal/model"
	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata-api/internal/services"
	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/rowdata"

	"google.golang.org/grpc"
)

//这个函数主要实现实现几个网络请求的处理
//增删改查

type RowDataController struct {
	rowDataService *services.RowDataService
	conn           *grpc.ClientConn
	client         pb.RowDataServiceClient
}

func NewRowDataController(conn *grpc.ClientConn, client pb.RowDataServiceClient) *RowDataController {
	return &RowDataController{
		rowDataService: services.NewRowdataService(),
		conn:           conn,
		client:         client,
	}
}

func (c *RowDataController) UploadRowData(ctx *gin.Context, client pb.RowDataServiceClient) {
	_, excel, err := ctx.Request.FormFile("excel")
	if err != nil {
		panic(err)
	}
	temp, _ := strconv.ParseInt(ctx.Request.FormValue("temp"), 10, 64)
	rowDataUploadInfo := model.RowDataUploadInfo{
		Weather: ctx.Request.FormValue("weather"),
		Temp:    int(temp),
		WindDir: ctx.Request.FormValue("windDir"),
		Loc:     ctx.Request.FormValue("loc"),
		Coach:   ctx.Request.FormValue("coach"),
		Remark:  ctx.Request.FormValue("remark"),
	}

	// fmt.Println(excel.Header.Get("Content-Type"))
	ctx.SaveUploadedFile(excel, "./excel/"+excel.Filename)
	ctx.String(http.StatusOK, excel.Filename+"上传成功")

	c.rowDataService.ProcessExcel("./excel/"+excel.Filename, &rowDataUploadInfo, client)
}

func (c *RowDataController) UploadMultiRowData(ctx *gin.Context, client pb.RowDataServiceClient) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		panic(err)
	}
	excels := form.File["excels"]
	if len(excels) < 1 {
		panic(errors.New("未上传文件"))

	}
	var excelnames []string
	for _, excel := range excels {
		if err := ctx.SaveUploadedFile(excel, "./excel/"+excel.Filename); err != nil {
			ctx.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			panic(err)
		}
		excelnames = append(excelnames, "./excel/"+excel.Filename)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("upload ok %d excels", len(excels)))

	temp, _ := strconv.ParseInt(ctx.Request.FormValue("temp"), 10, 64)

	rowDataUploadInfo := model.RowDataUploadInfo{
		Weather: ctx.Request.FormValue("weather"),
		Temp:    int(temp),
		WindDir: ctx.Request.FormValue("windDir"),
		Loc:     ctx.Request.FormValue("loc"),
		Coach:   ctx.Request.FormValue("coach"),
		Remark:  ctx.Request.FormValue("remark"),
	}

	trainingId := c.rowDataService.ProcessExcels(excelnames, &rowDataUploadInfo, client)
	for _, filename := range excelnames {
		c.rowDataService.ProcessAthleteTrainingData(filename, trainingId, client)
	}

}

func (c *RowDataController) FindTrainingSummaryByTrainingName(ctx *gin.Context, client pb.RowDataServiceClient) {

}

func (c *RowDataController) FindAthleteTrainingDataByName(ctx *gin.Context, client pb.RowDataServiceClient) {

}
