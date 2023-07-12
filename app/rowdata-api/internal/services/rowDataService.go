package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata-api/internal/model"
	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata-api/internal/services/grpc"
	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/rowdata"
)

//主要实现不同的方法 包括

type RowDataService struct {
}

func NewRowdataService() *RowDataService {
	return &RowDataService{}
}

//一个excel包含了一次训练所有的信息 1-n n-n*samplenum
//处理excel之后 返回

func (s *RowDataService) ProcessExcel(filename string, uploadInfo *model.RowDataUploadInfo, client pb.RowDataServiceClient) {
	excel, err := excelize.OpenFile(filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := excel.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//TrainingSummary 结构体中
	//获取TrainingName string
	trainingName, err := excel.GetCellValue("Summary", "B1")
	if err != nil {
		fmt.Println("训练名缺少")
		panic(err)
	}
	//获取TrainingData time.Time
	trainingDateString := trainingName[:10]
	layout := "2006_01_02" // 定义日期字符串的格式
	// 解析日期字符串为时间对象
	trainingDate, err := time.Parse(layout, trainingDateString)
	if err != nil {
		fmt.Println("日期解析错误:", err)
		panic(err)
	}
	//获取EventGender string
	eventGender, err := excel.GetCellValue("Summary", "A3")
	if err != nil {
		fmt.Println("性别缺少")
		panic(err)
	}
	//获取EventPeopleType string
	eventPeopleType, err := excel.GetCellValue("Summary", "B3")
	if err != nil {
		fmt.Println("人数类型缺少")
		panic(err)
	}
	//获取EventScale string
	eventScale, err := excel.GetCellValue("Summary", "C3")
	if err != nil {
		fmt.Println("量级缺少")
		panic(err)
	}
	//获取Event string
	event := eventGender + eventScale + eventPeopleType
	//获取SampleCount int
	sampleCountString, err := excel.GetCellValue("Summary", "A6")
	if err != nil {
		fmt.Println("采样数目缺少")
		panic(err)
	}
	sampleCount, _ := strconv.ParseInt(sampleCountString, 10, 64)

	trainingSummary := model.TrainingSummary{
		TrainingName:    trainingName,
		TrainingDate:    trainingDate,
		EventGender:     eventGender,
		EventPeopleType: eventPeopleType,
		EventScale:      eventScale,
		Event:           event,
		Weather:         uploadInfo.Weather,
		Temp:            uploadInfo.Temp,
		WindDir:         uploadInfo.WindDir,
		Loc:             uploadInfo.Loc,
		Coach:           uploadInfo.Coach,
		SampleCount:     int(sampleCount),
		Remark:          uploadInfo.Remark,
	}

	grpc.AddTrainingSummary(client, trainingSummary)

	fmt.Println(trainingSummary)

}

func (s *RowDataService) ProcessExcels(filenames []string, uploadInfo *model.RowDataUploadInfo, client pb.RowDataServiceClient) uint32 {
	excel, err := excelize.OpenFile(filenames[0])
	if err != nil {
		panic(err)
	}

	//TrainingSummary 结构体中
	//获取TrainingName string
	trainingName, err := excel.GetCellValue("Summary", "B1")
	if err != nil {
		fmt.Println("训练名缺少")
		panic(err)
	}
	//获取TrainingData time.Time
	trainingDateString := trainingName[:10]
	layout := "2006_01_02" // 定义日期字符串的格式
	// 解析日期字符串为时间对象
	trainingDate, err := time.Parse(layout, trainingDateString)
	if err != nil {
		fmt.Println("日期解析错误:", err)
		panic(err)
	}
	//获取EventGender string
	eventGender, err := excel.GetCellValue("Summary", "A3")
	if err != nil {
		fmt.Println("性别缺少")
		panic(err)
	}
	//获取EventPeopleType string
	eventPeopleType, err := excel.GetCellValue("Summary", "B3")
	if err != nil {
		fmt.Println("人数类型缺少")
		panic(err)
	}
	//获取EventScale string
	eventScale, err := excel.GetCellValue("Summary", "C3")
	if err != nil {
		fmt.Println("量级缺少")
		panic(err)
	}
	//获取Event string
	event := eventGender + eventScale + eventPeopleType
	//获取SampleCount int
	sampleCountString, err := excel.GetCellValue("Summary", "A6")
	if err != nil {
		fmt.Println("采样数目缺少")
		panic(err)
	}
	sampleCount, _ := strconv.ParseInt(sampleCountString, 10, 64)

	trainingSummary := model.TrainingSummary{
		TrainingName:    trainingName,
		TrainingDate:    trainingDate,
		EventGender:     eventGender,
		EventPeopleType: eventPeopleType,
		EventScale:      eventScale,
		Event:           event,
		Weather:         uploadInfo.Weather,
		Temp:            uploadInfo.Temp,
		WindDir:         uploadInfo.WindDir,
		Loc:             uploadInfo.Loc,
		Coach:           uploadInfo.Coach,
		SampleCount:     int(sampleCount),
		Remark:          uploadInfo.Remark,
	}

	//这里的状态码 之后考虑怎么优化
	_, trainingId := grpc.AddTrainingSummary(client, trainingSummary)

	if err := excel.Close(); err != nil {
		fmt.Println(err)
	}
	return trainingId

}

// 这个函数负责将excel里面运动员数据层的数据提取 并转换为对象串
func (s *RowDataService) ProcessAthleteTrainingData(filename string, trainingId uint32, client pb.RowDataServiceClient) {
	excel, err := excelize.OpenFile(filename)
	if err != nil {
		panic(err)
	}

	Name, err := excel.GetCellValue("Summary", "A4")
	if err != nil {
		fmt.Println("运动员姓名缺少")
		panic(err)
	}

	Gender, err := excel.GetCellValue("Summary", "B4")
	if err != nil {
		fmt.Println("运动员性别缺少")
		panic(err)
	}

	Seat, err := excel.GetCellValue("Summary", "C4")
	if err != nil {
		fmt.Println("运动员座位缺少")
		panic(err)
	}
	seat, _ := strconv.ParseInt(Seat, 10, 64)

	Side, err := excel.GetCellValue("Summary", "D4")
	if err != nil {
		fmt.Println("运动员侧方位缺少")
		panic(err)
	}
	side, _ := strconv.ParseInt(Side, 10, 64)

	Height, err := excel.GetCellValue("Summary", "E4")
	if err != nil {
		fmt.Println("运动员身高缺少")
		panic(err)
	}
	height, _ := strconv.ParseFloat(Height, 32)

	Weight, err := excel.GetCellValue("Summary", "F4")
	if err != nil {
		fmt.Println("运动员体重缺少")
		panic(err)
	}
	weight, _ := strconv.ParseFloat(Weight, 32)

	OarInboard, err := excel.GetCellValue("Summary", "A5")
	if err != nil {
		fmt.Println("缺少OarInboard")
		panic(err)
	}
	oarInboard, _ := strconv.ParseFloat(OarInboard, 32)

	OarLength, err := excel.GetCellValue("Summary", "B5")
	if err != nil {
		fmt.Println("缺少OarLength")
		panic(err)
	}
	oarLength, _ := strconv.ParseFloat(OarLength, 32)

	OarBladeLength, err := excel.GetCellValue("Summary", "C5")
	if err != nil {
		fmt.Println("缺少OarBladeLength")
		panic(err)
	}
	oarBladeLength, _ := strconv.ParseFloat(OarBladeLength, 32)

	athletetrainingdata := model.AthleteTrainingData{
		TrainingId:     uint(trainingId),
		Name:           Name,
		Gender:         Gender,
		Seat:           int(seat),
		Side:           int(side),
		Height:         float32(height),
		Weight:         float32(weight),
		OarInboard:     float32(oarInboard),
		OarLength:      float32(oarLength),
		OarBladeLength: float32(oarBladeLength),
	}

	_, athleteTrainingId := grpc.AddAthleteTrainingData(client, athletetrainingdata)

	sampleTimes, err := excel.GetCellValue("Summary", "A8")
	if err != nil {
		fmt.Println("采样组有问题")
		panic(err)
	}
	st, _ := strconv.ParseInt(sampleTimes, 10, 64)
	sampleArr := make([]int, st)
	for num := range sampleArr {
		col := string('A' + num)
		sampleNum, err := excel.GetCellValue("Summary", col+"9")
		if err != nil {
			fmt.Println("采样组有问题")
			panic(err)
		}
		sn, _ := strconv.ParseInt(sampleNum, 10, 64)
		sampleArr[num] = int(sn)
	}

	for key, value := range sampleArr {
		sampleNum := value
		fmt.Println(key)
		sheetnameT8d := "T8d_" + strconv.Itoa(key+1)
		sheetnameB8d := "B8d_" + strconv.Itoa(key+1)
		sheetnameT8dFig := "T8d_Fig_" + strconv.Itoa(key+1)
		sheetnameB8dFig := "B8d_Fig_" + strconv.Itoa(key+1)

		rowsT8d, err := excel.GetRows(sheetnameT8d)
		if err != nil {
			fmt.Println("T8d有问题")
			panic(err)
		}
		rowsB8d, err := excel.GetRows(sheetnameB8d)
		if err != nil {
			fmt.Println("B8d有问题")
			panic(err)
		}
		rowsT8dFig, err := excel.GetRows(sheetnameT8dFig)
		if err != nil {
			fmt.Println("T8d_Fig有问题")
			panic(err)
		}
		rowsB8dFig, err := excel.GetRows(sheetnameB8dFig)
		if err != nil {
			fmt.Println("B8d_Fig有问题")
			panic(err)
		}

		for i := 0; i < sampleNum; i++ {
			sampleMetrics := model.SampleMetrics{}
			sampleMetrics.AthleteTrainingId = uint(athleteTrainingId)
			sampleMetrics.DataSample = rowsT8d[3+i][1]
			sampleMetrics.StrokeRate, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][2]), 64)
			sampleMetrics.DriveTime, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][3]), 64)
			sampleMetrics.Rhythm, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][4]), 64)
			sampleMetrics.CatchAngle, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][5]), 64)
			sampleMetrics.FinishAngle, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][6]), 64)
			sampleMetrics.TotalAngle, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][7]), 64)
			sampleMetrics.CatchSlip, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][8]), 64)
			sampleMetrics.ReleaseSlip, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][9]), 64)
			sampleMetrics.EffectiveAnglePercent, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][10]), 64)
			sampleMetrics.MaxBladeDepth, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][11]), 64)
			sampleMetrics.BladeEfficiency, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][12]), 64)
			sampleMetrics.RowingPower, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][13]), 64)
			sampleMetrics.WorkPerStroke, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][14]), 64)
			sampleMetrics.RelativeWpS, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][16]), 64)
			sampleMetrics.EffectiveAngleDegree, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][18]), 64)
			sampleMetrics.TargetAngle, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][21]), 64)
			sampleMetrics.TargetForce, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][22]), 64)
			sampleMetrics.TargetWpS, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][23]), 64)
			sampleMetrics.AngleDivTarget, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][24]), 64)
			sampleMetrics.ForceDivTarget, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][25]), 64)
			sampleMetrics.WpSDivTarget, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][26]), 64)
			sampleMetrics.AverageVelocity, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][28]), 64)
			sampleMetrics.BladeSpecificImpulse, _ = strconv.ParseFloat(IsZero(rowsT8d[3+i][29]), 64)

			sampleMetrics.TimeOver2000m, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][1]), 64)
			sampleMetrics.MaxForce, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][2]), 64)
			sampleMetrics.AverageForce, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][3]), 64)
			sampleMetrics.RatioAverDivMaxForce, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][4]), 64)
			sampleMetrics.PositionOfPeakForce, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][5]), 64)
			sampleMetrics.CatchForceGradient, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][6]), 64)
			sampleMetrics.FinishForceGradient, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][7]), 64)
			sampleMetrics.MaxHandleVelocity, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][8]), 64)
			sampleMetrics.HDF, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][9]), 64)
			sampleMetrics.LegsDrive, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][10]), 64)
			sampleMetrics.LegsMaxSpeed, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][11]), 64)
			sampleMetrics.CatchFactor, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][12]), 64)
			sampleMetrics.RowingStyleFactor, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][13]), 64)
			sampleMetrics.ReleaseWash, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][18]), 64)
			sampleMetrics.AverForceDivWeight, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][19]), 64)
			sampleMetrics.VseatAtCatch, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][20]), 64)
			sampleMetrics.HandleTravelAtEntryForce, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][21]), 64)
			sampleMetrics.HandleTravelAt70PerForce, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][22]), 64)
			sampleMetrics.HandleTravelAt0As, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][23]), 64)
			sampleMetrics.SeatTravelAtEntryForce, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][24]), 64)
			sampleMetrics.SeatTravelAt70PerForce, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][25]), 64)
			sampleMetrics.SeatTravelAt0As, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][26]), 64)
			sampleMetrics.DTravelAtEntryForcePercent, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][27]), 64)
			sampleMetrics.DTravelAt70PerForcePercent, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][28]), 64)
			sampleMetrics.DTravelAt0AsPercent, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][29]), 64)
			sampleMetrics.DTravelAtEntryForceDistance, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][30]), 64)
			sampleMetrics.DTravelAt70PerForceDistance, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][31]), 64)
			sampleMetrics.DTravelAt0AsDistance, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][32]), 64)
			sampleMetrics.SeatOnRecovery, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][33]), 64)
			sampleMetrics.VertAtCatch, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][34]), 64)
			sampleMetrics.EntryForce, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][35]), 64)
			sampleMetrics.ForceUpto70Per, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][36]), 64)
			sampleMetrics.MaxVseat, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][37]), 64)
			sampleMetrics.PeakForce, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][38]), 64)
			sampleMetrics.ForceFrom70Per, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][39]), 64)
			sampleMetrics.VertAtFinish, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][40]), 64)
			sampleMetrics.ForceAtFinish, _ = strconv.ParseFloat(IsZero(rowsT8d[14+i][41]), 64)

			sampleMetrics.AverageBoatSpeed, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][4]), 64)
			sampleMetrics.MinimalBoatSpeed, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][5]), 64)
			sampleMetrics.MaximalBoatSpeed, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][6]), 64)
			sampleMetrics.DistancePerStroke, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][7]), 64)
			sampleMetrics.DragFactor, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][8]), 64)
			sampleMetrics.WindForwardCompRelWater, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][10]), 64)
			sampleMetrics.WindDirectionRelWater, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][11]), 64)
			sampleMetrics.Time250m, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][12]), 64)
			sampleMetrics.BoatSpeedEfficiency, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][13]), 64)
			sampleMetrics.TimeAtWaterTemp25Deg, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][14]), 64)
			sampleMetrics.BoatSpeedVariation, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][15]), 64)
			sampleMetrics.WindSpeedRelBoat, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][16]), 64)
			sampleMetrics.WindDirectionRelBoat, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][17]), 64)
			sampleMetrics.DragFactorF, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][24]), 64)
			sampleMetrics.DragFactorPprop, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][25]), 64)
			sampleMetrics.DragFactorPold, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][26]), 64)
			sampleMetrics.DragFactorPtot, _ = strconv.ParseFloat(IsZero(rowsB8d[3+i][27]), 64)
			sampleMetrics.AccelerationMinimun, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][2]), 64)
			sampleMetrics.AccelerationMaximum, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][3]), 64)
			sampleMetrics.ModelSpeed, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][7]), 64)
			sampleMetrics.EffectiveWorkPerStroke, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][8]), 64)
			sampleMetrics.ModelDPS, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][10]), 64)
			sampleMetrics.PropulsivePower, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][13]), 64)
			sampleMetrics.DriveMaximalAt, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][14]), 64)
			sampleMetrics.FirstPeak, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][15]), 64)
			sampleMetrics.ZeroBeforeCatch, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][16]), 64)
			sampleMetrics.MinimalFromCatch, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][17]), 64)
			sampleMetrics.ZeroAfterCatch, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][18]), 64)
			sampleMetrics.StdDeviation, _ = strconv.ParseFloat(IsZero(rowsB8d[13+i][22]), 64)

			for j := 0; j < 51; j++ {
				sampleMetrics.OarAngle[j], _ = strconv.ParseFloat(IsZero(rowsT8dFig[1+i][j]), 64)
				sampleMetrics.HandleForce[j], _ = strconv.ParseFloat(IsZero(rowsT8dFig[10+i][j]), 64)
				sampleMetrics.VerticalAngleBoatRoll[j], _ = strconv.ParseFloat(IsZero(rowsT8dFig[19+i][j]), 64)
				sampleMetrics.LegsVelocity[j], _ = strconv.ParseFloat(IsZero(rowsT8dFig[28+i][j]), 64)
				sampleMetrics.HandleSpeed[j], _ = strconv.ParseFloat(IsZero(rowsT8dFig[37+i][j]), 64)
				sampleMetrics.HDFFig[j], _ = strconv.ParseFloat(IsZero(rowsT8dFig[46+i][j]), 64)
				sampleMetrics.BladeDF[j], _ = strconv.ParseFloat(IsZero(rowsT8dFig[55+i][j]), 64)
				sampleMetrics.Velocity[j], _ = strconv.ParseFloat(IsZero(rowsB8dFig[1+i][j]), 64)
				sampleMetrics.BoatAcceleration[j], _ = strconv.ParseFloat(IsZero(rowsB8dFig[10+i][j]), 64)
				sampleMetrics.VelocityRel[j], _ = strconv.ParseFloat(IsZero(rowsB8dFig[28+i][j]), 64)
			}

			// addSampleMetricsResp := grpc.AddSampleMetrics(client, sampleMetrics)
			grpc.AddSampleMetrics(client, sampleMetrics)
		}

	}
}

func (s *RowDataService) GetTrainingSummary(findReq model.TrainingSummaryGet, client pb.RowDataServiceClient) *[]model.TrainingSummary {
	_, tsArr := grpc.GetTrainingSummary(client, findReq)
	return tsArr
}

func (s *RowDataService) GetAthleteTrainingData(findReq model.AthleteTrainingDataGet, client pb.RowDataServiceClient) *[]model.AthleteTrainingData {
	_, atdArr := grpc.GetAthleteTrainingData(client, findReq)
	return atdArr
}

func (s *RowDataService) GetSampleMetricsByAthleteTrainingId(id uint32, set bool, client pb.RowDataServiceClient) *[]model.SampleMetrics {
	_, smArr := grpc.GetSampleMetricsByAthleteTrainingId(client, id, set)
	return smArr
}

func IsZero(data string) string {
	if data == "#NAME?" || data == "-" || data == "#VALUE!" || data == "#N/A" || data == "#REF!" || data == "DIV/0!" {
		return "0"
	} else {
		return data
	}
}
