package controller

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata-api/internal/model"
	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata-api/internal/services"
	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/rowdata"
	"github.com/zizhuxuexi/datacenter-rowing/pkg/util/convert"

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

// 上传单个文件（未使用）
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
	// ctx.String(http.StatusOK, fmt.Sprintf("upload ok %d excels", len(excels)))

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
		err = os.Remove(filename)
		if err != nil {
			panic(err)
		}

	}

	//返回上传成功的响应
	ctx.JSON(200, gin.H{"message": "上传文件成功", "status": 200, "trainingSummaryId": trainingId})

}

// 传入单个查询请求
func (c *RowDataController) FindTrainingSummary(ctx *gin.Context, client pb.RowDataServiceClient) {
	queries := ctx.Request.URL.Query()
	findReq := model.TrainingSummaryGet{}
	if len(queries) == 0 {
		findReq.Set = false
	} else {
		if queries.Get("training_id") != "" {
			trainingId, _ := strconv.ParseUint(queries.Get("training_id"), 10, 32)
			findReq.TrainingSummary.TrainingId = uint(trainingId)
			findReq.SetTrainingId = true
		} else {
			findReq.SetTrainingId = false
		}

		if queries.Get("training_name") != "" {
			findReq.TrainingSummary.TrainingName = queries.Get("training_name")
			findReq.SetTrainingName = true
		} else {
			findReq.SetTrainingName = false
		}

		if queries.Get("training_date") != "" {
			trainingDate := queries.Get("training_date")
			findReq.TrainingSummary.TrainingDate = convert.StringToTime(trainingDate)
			findReq.SetTrainingDate = true
		} else {
			findReq.SetTrainingDate = false
		}

		if queries.Get("event_gender") != "" {
			findReq.TrainingSummary.EventGender = queries.Get("event_gender")
			findReq.SetEventGender = true
		} else {
			findReq.SetEventGender = false
		}

		if queries.Get("event_people_type") != "" {
			findReq.TrainingSummary.EventPeopleType = queries.Get("event_people_type")
			findReq.SetEventPeopleType = true
		} else {
			findReq.SetEventPeopleType = false
		}

		if queries.Get("event_scale") != "" {
			findReq.TrainingSummary.EventScale = queries.Get("event_scale")
			findReq.SetEventScale = true
		} else {
			findReq.SetEventScale = false
		}

		if queries.Get("event") != "" {
			findReq.TrainingSummary.Event = queries.Get("event")
			findReq.SetEvent = true
		} else {
			findReq.SetEvent = false
		}

		if queries.Get("wheater") != "" {
			findReq.TrainingSummary.Weather = queries.Get("wheater")
			findReq.SetWeather = true
		} else {
			findReq.SetWeather = false
		}

		if queries.Get("temp") != "" {
			temp, _ := strconv.ParseInt(queries.Get("temp"), 10, 64)
			findReq.TrainingSummary.Temp = int(temp)
			findReq.SetTemp = true
		} else {
			findReq.SetTemp = false
		}

		if queries.Get("wind_dir") != "" {
			findReq.TrainingSummary.WindDir = queries.Get("wind_dir")
			findReq.SetWindDir = true
		} else {
			findReq.SetWindDir = false
		}
		if queries.Get("loc") != "" {
			findReq.TrainingSummary.Loc = queries.Get("loc")
			findReq.SetTemp = true
		} else {
			findReq.SetTemp = false
		}
		if queries.Get("coach") != "" {

			findReq.TrainingSummary.Coach = queries.Get("coach")
			findReq.SetCoach = true
		} else {
			findReq.SetCoach = false
		}
		if queries.Get("sample_count") != "" {
			sampleCount, _ := strconv.ParseInt(queries.Get("sample_count"), 10, 64)
			findReq.TrainingSummary.SampleCount = int(sampleCount)
			findReq.SetSampleCount = true
		} else {
			findReq.SetSampleCount = false
		}
		if queries.Get("remark") != "" {
			findReq.TrainingSummary.Remark = queries.Get("remark")
			findReq.SetRemark = true
		} else {
			findReq.SetRemark = false
		}
	}

	trainingSummaryArr := c.rowDataService.GetTrainingSummary(findReq, client)
	ctx.JSON(200, gin.H{"message": "查询成功", "status": 200, "len": len(*trainingSummaryArr), "trainingSummary": *trainingSummaryArr})

}

func (c *RowDataController) FindAthleteTrainingData(ctx *gin.Context, client pb.RowDataServiceClient) {
	queries := ctx.Request.URL.Query()
	findReq := model.AthleteTrainingDataGet{}
	if len(queries) == 0 {
		findReq.Set = false
	} else {
		if queries.Get("athlete_training_id") != "" {
			athleteTrainingId, _ := strconv.ParseUint(queries.Get("athlete_training_id"), 10, 32)
			findReq.AthleteTrainingData.AthleteTrainingId = uint(athleteTrainingId)
			findReq.SetAthleteTrainingId = true
		} else {
			findReq.SetAthleteTrainingId = false
		}

		if queries.Get("training_id") != "" {
			trainingId, _ := strconv.ParseUint(queries.Get("training_id"), 10, 32)
			findReq.AthleteTrainingData.TrainingId = uint(trainingId)
			findReq.SetTrainingId = true
		} else {
			findReq.SetTrainingId = false
		}

		if queries.Get("name") != "" {
			findReq.AthleteTrainingData.Name = queries.Get("name")
			findReq.SetName = true
		} else {
			findReq.SetName = false
		}

		if queries.Get("gender") != "" {
			findReq.AthleteTrainingData.Gender = queries.Get("gender")
			findReq.SetGender = true
		} else {
			findReq.SetGender = false
		}

		if queries.Get("seat") != "" {
			seat, _ := strconv.ParseInt(queries.Get("seat"), 10, 64)
			findReq.AthleteTrainingData.Seat = int(seat)
			findReq.SetSeat = true
		} else {
			findReq.SetSeat = false
		}

		if queries.Get("side") != "" {
			side, _ := strconv.ParseInt(queries.Get("side"), 10, 64)
			findReq.AthleteTrainingData.Side = int(side)
			findReq.SetSide = true
		} else {
			findReq.SetSide = false
		}

		if queries.Get("height") != "" {
			height, _ := strconv.ParseFloat(queries.Get("height"), 32)
			findReq.AthleteTrainingData.Height = float32(height)
			findReq.SetHeight = true
		} else {
			findReq.SetHeight = false
		}
		if queries.Get("weight") != "" {
			weight, _ := strconv.ParseFloat(queries.Get("weight"), 32)
			findReq.AthleteTrainingData.Weight = float32(weight)
			findReq.SetWeight = true
		} else {
			findReq.SetWeight = false
		}

		if queries.Get("oar_inboard") != "" {
			oarInboard, _ := strconv.ParseFloat(queries.Get("oar_inboard"), 32)
			findReq.AthleteTrainingData.OarInboard = float32(oarInboard)
			findReq.SetOarInboard = true
		} else {
			findReq.SetOarInboard = false
		}

		if queries.Get("oar_length") != "" {
			oarLength, _ := strconv.ParseFloat(queries.Get("oar_length"), 32)
			findReq.AthleteTrainingData.OarLength = float32(oarLength)
			findReq.SetOarLength = true
		} else {
			findReq.SetOarLength = false
		}

		if queries.Get("oar_blade_length") != "" {
			oarBladeLength, _ := strconv.ParseFloat(queries.Get("oar_blade_length"), 32)
			findReq.AthleteTrainingData.OarBladeLength = float32(oarBladeLength)
			findReq.SetOarBladeLength = true
		} else {
			findReq.SetOarBladeLength = false
		}

	}

	athleteTrainingDataArr := c.rowDataService.GetAthleteTrainingData(findReq, client)
	ctx.JSON(200, gin.H{"message": "查询成功", "status": 200, "len": len(*athleteTrainingDataArr), "athleteTrainingData": *athleteTrainingDataArr})
}

func (c *RowDataController) FindSampleMetricsByAthleteTrainingId(ctx *gin.Context, client pb.RowDataServiceClient) {
	athleteTrainingID := ctx.Query("athlete_training_id")
	var sampleMetricsArr *[]model.SampleMetrics
	if athleteTrainingID != "" {
		id, _ := strconv.ParseUint(athleteTrainingID, 10, 32)
		sampleMetricsArr = c.rowDataService.GetSampleMetricsByAthleteTrainingId(uint32(id), true, client)
	} else {
		sampleMetricsArr = c.rowDataService.GetSampleMetricsByAthleteTrainingId(0, false, client)
	}
	ctx.JSON(200, gin.H{"message": "查询成功", "status": 200, "len": len(*sampleMetricsArr), "sampleMetrics": *sampleMetricsArr})

}
