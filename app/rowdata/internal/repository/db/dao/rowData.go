package dao

import (
	"context"
	"errors"

	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata/internal/repository/db/model"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/rowdata"
	"github.com/zizhuxuexi/datacenter-rowing/pkg/util/logger"
)

type RowDataDao struct {
	*gorm.DB
}

func NewRowDataDao(ctx context.Context) *RowDataDao {
	return &RowDataDao{NewDBClient(ctx)}
}

// 上传一级训练记录
func (dao *RowDataDao) AddTrainingSummary(req *pb.TrainingSummaryAddRequest) (uint, error) {
	//if exist := dao.CheckExist(req); exist {
	//	return errors.New("训练记录已存在，请勿重复上传。")
	//}
	var count int64
	var trainingSummary model.TrainingSummary
	dao.Model(&model.TrainingSummary{}).Where("training_name=?", req.TrainingName).Count(&count)
	if count != 0 {
		return 0, errors.New("训练记录已存在，请勿重复上传。")
	}
	trainingSummary = model.TrainingSummary{
		TrainingName:    req.TrainingName,
		TrainingDate:    req.TrainDate.AsTime(),
		EventGender:     req.EventGender,
		EventPeopleType: req.EventPeopleType,
		EventScale:      req.EventScale,
		Event:           req.Event,
		Weather:         req.Weather,
		Temp:            int(req.Temp),
		WindDir:         req.WindDir,
		Loc:             req.Loc,
		Coach:           req.Coach,
		SampleCount:     int(req.SampleCount),
		Remark:          req.Remark,
	}
	if err := dao.Model(&model.TrainingSummary{}).Create(&trainingSummary).Error; err != nil {
		logger.LogrusObj.Error("Insert User Error:" + err.Error())
		return 0, err
	}
	return trainingSummary.TrainingId, nil
}

// 功能可以背GetTrainingSummary替代
func (dao *RowDataDao) GetTrainingSummaryByTrainingName(req *pb.TrainingSummaryGetByTrainingNameRequest) (ts *model.TrainingSummary, err error) {
	err = dao.Model(&model.TrainingSummary{}).Where("training_name=?", req.TrainingName).First(&ts).Error
	return
}

// 查询制定条件的一级训练记录
func (dao *RowDataDao) GetTrainingSummary(req *pb.TrainingSummaryGetRequset) (ts []*model.TrainingSummary, err error) {
	query := dao.Model(&model.TrainingSummary{})

	//判断是否有需要查询的内容
	if req.Set {
		if req.SetTrainingId {
			query = query.Where("training_id=?", req.TrainingSummary.TrainingId)
		}
		if req.SetTrainingName {
			query = query.Where("training_name=?", req.TrainingSummary.TrainingName)
		}
		if req.SetTrainingDate {
			query = query.Where("training_date=?", req.TrainingSummary.TrainingDate.AsTime())
		}
		if req.SetEventGender {
			query = query.Where("event_gender=?", req.TrainingSummary.EventGender)
		}
		if req.SetEventPeopleType {
			query = query.Where("event_people_type=?", req.TrainingSummary.EventPeopleType)
		}
		if req.SetEventScale {
			query = query.Where("event_scale=?", req.TrainingSummary.EventScale)
		}
		if req.SetWeather {
			query = query.Where("weather=?", req.TrainingSummary.Weather)
		}
		if req.SetTemp {
			query = query.Where("temp=?", req.TrainingSummary.Temp)
		}
		if req.SetWindDir {
			query = query.Where("wind_dir=?", req.TrainingSummary.WindDir)
		}
		if req.SetLoc {
			query = query.Where("loc=?", req.TrainingSummary.Loc)
		}
		if req.SetCoach {
			query = query.Where("coach=?", req.TrainingSummary.Coach)
		}
		if req.SetSampleCount {
			query = query.Where("sample_count=?", req.TrainingSummary.SampleCount)
		}
		if req.SetRemark {
			query = query.Where("remark=?", req.TrainingSummary.Remark)
		}
	}
	err = query.Find(&ts).Error
	return

}

// 查询是否已经存在重复的一级训练记录（通过训练名判断）
func (dao *RowDataDao) CheckExist(req *pb.TrainingSummaryAddRequest) bool {
	var trainingSummary model.TrainingSummary
	if err := dao.Model(&model.TrainingSummary{}).Where("training_name=?", req.TrainingName).First(&trainingSummary).Error; err == gorm.ErrRecordNotFound {
		return false
	} else {
		return true
	}
}

// 将数据库读取的一级训练数据对象 转换为 grpc通信的对象
func BuildTrainingSummary(item model.TrainingSummary) *pb.TrainingSummaryModel {
	trainingSummaryModel := pb.TrainingSummaryModel{
		TrainingId:      uint32(item.TrainingId),
		TrainingName:    item.TrainingName,
		TrainingDate:    timestamppb.New(item.TrainingDate),
		EventGender:     item.EventGender,
		EventPeopleType: item.EventPeopleType,
		EventScale:      item.EventScale,
		Event:           item.Event,
		Weather:         item.Weather,
		Temp:            int32(item.Temp),
		WindDir:         item.WindDir,
		Loc:             item.Loc,
		Coach:           item.Coach,
		SampleCount:     int32(item.SampleCount),
		Remark:          item.Remark,
	}
	return &trainingSummaryModel
}

// 上传二级训练记录 运动选训练数据
func (dao *RowDataDao) AddAthleteTrainingData(req *pb.AthleteTrainingDataAddRequest) (uint, error) {
	//if exist := dao.CheckExist(req); exist {
	//	return errors.New("训练记录已存在，请勿重复上传。")
	//}
	//var count int64
	//athleteTrainingData还没想好怎么判断是否已经上传
	// dao.Model(&model.TrainingSummary{}).Where("training_name=?", req.TrainingName).Count(&count)
	// if count != 0 {
	// 	return errors.New("训练记录已存在，请勿重复上传。")
	// }
	athleteTrainingData := model.AthleteTrainingData{
		TrainingId:     uint(req.TrainingId),
		Name:           req.Name,
		Gender:         req.Gender,
		Seat:           int(req.Seat),
		Side:           int(req.Side),
		Height:         req.Height,
		Weight:         req.Weight,
		OarInboard:     req.OarInboard,
		OarLength:      req.OarLength,
		OarBladeLength: req.OarBladeLength,
	}
	if err := dao.Model(&model.AthleteTrainingData{}).Create(&athleteTrainingData).Error; err != nil {
		logger.LogrusObj.Error("Insert User Error:" + err.Error())
		return 0, err
	}
	return athleteTrainingData.AthleteTrainingId, nil
}

// 功能可以被GetAthleteTrainingData取代
func (dao *RowDataDao) GetAthleteTrainingDataByName(req *pb.AthleteTrainingDataGetByName) (ts *[]model.AthleteTrainingData, err error) {
	err = dao.Model(&model.AthleteTrainingData{}).Where("name=?", req.Name).Find(&ts).Error
	return
}

func (dao *RowDataDao) GetAthleteTrainingData(req *pb.AthleteTrainingDataGetRequest) (atd []*model.AthleteTrainingData, err error) {
	query := dao.Model(&model.AthleteTrainingData{})
	if req.Set {
		if req.SetAthleteTrainingId {
			query = query.Where("athlete_training_id=?", req.AthleteTrainingData.AthleteTrainingId)
		}
		if req.SetTrainingId {
			query = query.Where("training_id=?", req.AthleteTrainingData.TrainingId)
		}
		if req.SetName {
			query = query.Where("name=?", req.AthleteTrainingData.Name)
		}
		if req.SetGender {
			query = query.Where("gender=?", req.AthleteTrainingData.Gender)
		}
		if req.SetSeat {
			query = query.Where("seat=?", req.AthleteTrainingData.Seat)
		}
		if req.SetSide {
			query = query.Where("side=?", req.AthleteTrainingData.Side)
		}
		if req.SetHeight {
			query = query.Where("height=?", req.AthleteTrainingData.Height)
		}
		if req.SetWeight {
			query = query.Where("weight=?", req.AthleteTrainingData.Weight)
		}
		if req.SetOarInboard {
			query = query.Where("oar_inboard=?", req.AthleteTrainingData.OarInboard)
		}
		if req.SetOarLength {
			query = query.Where("oar_length=?", req.AthleteTrainingData.OarLength)
		}
		if req.SetOarBladeLength {
			query = query.Where("oar_blade_length=?", req.AthleteTrainingData.OarBladeLength)
		}
	}

	err = query.Find(&atd).Error
	return

}

// 将数据库读取的耳机训练数据对象 转换为 grpc通信对象
func BuildAthleteTrainingData(req model.AthleteTrainingData) *pb.AthleteTrainingDataModel {
	athleteTrainingDataModel := pb.AthleteTrainingDataModel{
		AthleteTrainingId: uint32(req.AthleteTrainingId),
		TrainingId:        uint32(req.TrainingId),
		Name:              req.Name,
		Gender:            req.Gender,
		Seat:              int32(req.Seat),
		Side:              int32(req.Side),
		Height:            req.Height,
		Weight:            req.Weight,
		OarInboard:        req.OarInboard,
		OarLength:         req.OarLength,
		OarBladeLength:    req.OarBladeLength,
	}
	return &athleteTrainingDataModel
}

// 上传三级训练记录 采样指标
func (dao *RowDataDao) AddSampleMetrics(req *pb.SampleMetricsModel) (uint, error) {
	sampleMetrics := model.SampleMetrics{
		AthleteTrainingId:           uint(req.AthleteTrainingId),
		DataSample:                  req.DataSample,
		StrokeRate:                  req.StrokeRate,
		DriveTime:                   req.DriveTime,
		Rhythm:                      req.Rhythm,
		CatchAngle:                  req.CatchAngle,
		FinishAngle:                 req.FinishAngle,
		TotalAngle:                  req.TotalAngle,
		CatchSlip:                   req.CatchSlip,
		ReleaseSlip:                 req.ReleaseSlip,
		EffectiveAnglePercent:       req.EffectiveAnglePercent,
		MaxBladeDepth:               req.MaxBladeDepth,
		BladeEfficiency:             req.BladeEfficiency,
		RowingPower:                 req.RowingPower,
		WorkPerStroke:               req.WorkPerStroke,
		RelativeWpS:                 req.RelativeWpS,
		EffectiveAngleDegree:        req.EffectiveAngleDegree,
		TargetAngle:                 req.TargetAngle,
		TargetForce:                 req.TargetForce,
		TargetWpS:                   req.TargetWpS,
		AngleDivTarget:              req.AngleDivTarget,
		ForceDivTarget:              req.ForceDivTarget,
		WpSDivTarget:                req.WpSDivTarget,
		AverageVelocity:             req.AverageVelocity,
		BladeSpecificImpulse:        req.BladeSpecificImpulse,
		TimeOver2000m:               req.TimeOver_2000M,
		MaxForce:                    req.MaxForce,
		AverageForce:                req.AverageForce,
		RatioAverDivMaxForce:        req.RatioAverDivMaxForce,
		PositionOfPeakForce:         req.PositionOfPeakForce,
		CatchForceGradient:          req.CatchForceGradient,
		FinishForceGradient:         req.FinishForceGradient,
		MaxHandleVelocity:           req.MaxHandleVelocity,
		HDF:                         req.Hdf,
		LegsDrive:                   req.LegsDrive,
		LegsMaxSpeed:                req.LegsMaxSpeed,
		CatchFactor:                 req.CatchFactor,
		RowingStyleFactor:           req.RowingStyleFactor,
		ReleaseWash:                 req.ReleaseWash,
		AverForceDivWeight:          req.AverForceDivWeight,
		VseatAtCatch:                req.VseatAtCatch,
		HandleTravelAtEntryForce:    req.HandleTravelAtEntryForce,
		HandleTravelAt70PerForce:    req.HandleTravelAt_70PerForce,
		HandleTravelAt0As:           req.HandleTravelAt_0As,
		SeatTravelAtEntryForce:      req.SeatTravelAtEntryForce,
		SeatTravelAt70PerForce:      req.SeatTravelAt_70PerForce,
		SeatTravelAt0As:             req.SeatTravelAt_0As,
		DTravelAtEntryForcePercent:  req.DTravelAtEntryForcePercent,
		DTravelAt70PerForcePercent:  req.DTravelAt_70PerForcePercent,
		DTravelAt0AsPercent:         req.DTravelAt_0AsPercent,
		DTravelAtEntryForceDistance: req.DTravelAtEntryForceDistance,
		DTravelAt70PerForceDistance: req.DTravelAt_70PerForceDistance,
		DTravelAt0AsDistance:        req.DTravelAt_0AsDistance,
		SeatOnRecovery:              req.SeatOnRecovery,
		VertAtCatch:                 req.VertAtCatch,
		EntryForce:                  req.EntryForce,
		ForceUpto70Per:              req.ForceUpto_70Per,
		MaxVseat:                    req.MaxVseat,
		PeakForce:                   req.PeakForce,
		ForceFrom70Per:              req.ForceFrom_70Per,
		VertAtFinish:                req.VertAtFinish,
		ForceAtFinish:               req.ForceAtFinish,
		AverageBoatSpeed:            req.AverageBoatSpeed,
		MinimalBoatSpeed:            req.MinimalBoatSpeed,
		MaximalBoatSpeed:            req.MaximalBoatSpeed,
		DistancePerStroke:           req.DistancePerStroke,
		DragFactor:                  req.DragFactor,
		WindForwardCompRelWater:     req.WindForwardCompRelWater,
		WindDirectionRelWater:       req.WindDirectionRelWater,
		Time250m:                    req.Time_250M,
		BoatSpeedEfficiency:         req.BoatSpeedEfficiency,
		TimeAtWaterTemp25Deg:        req.TimeAtWaterTemp_25Deg,
		BoatSpeedVariation:          req.BoatSpeedVariation,
		WindSpeedRelBoat:            req.WindSpeedRelBoat,
		WindDirectionRelBoat:        req.WindDirectionRelBoat,
		DragFactorF:                 req.DragFactorF,
		DragFactorPprop:             req.DragFactorPprop,
		DragFactorPold:              req.DragFactorPold,
		DragFactorPtot:              req.DragFactorPtot,
		AccelerationMinimun:         req.AccelerationMinimun,
		AccelerationMaximum:         req.AccelerationMaximum,
		ModelSpeed:                  req.ModelSpeed,
		EffectiveWorkPerStroke:      req.EffectiveWorkPerStroke,
		ModelDPS:                    req.ModelDps,
		PropulsivePower:             req.PropulsivePower,
		DriveMaximalAt:              req.DriveMaximalAt,
		FirstPeak:                   req.FirstPeak,
		ZeroBeforeCatch:             req.ZeroBeforeCatch,
		MinimalFromCatch:            req.MinimalFromCatch,
		ZeroAfterCatch:              req.ZeroAfterCatch,
		StdDeviation:                req.StdDeviation,

		OarAngle_0:  req.OarAngle[0],
		OarAngle_1:  req.OarAngle[1],
		OarAngle_2:  req.OarAngle[2],
		OarAngle_3:  req.OarAngle[3],
		OarAngle_4:  req.OarAngle[4],
		OarAngle_5:  req.OarAngle[5],
		OarAngle_6:  req.OarAngle[6],
		OarAngle_7:  req.OarAngle[7],
		OarAngle_8:  req.OarAngle[8],
		OarAngle_9:  req.OarAngle[9],
		OarAngle_10: req.OarAngle[10],
		OarAngle_11: req.OarAngle[11],
		OarAngle_12: req.OarAngle[12],
		OarAngle_13: req.OarAngle[13],
		OarAngle_14: req.OarAngle[14],
		OarAngle_15: req.OarAngle[15],
		OarAngle_16: req.OarAngle[16],
		OarAngle_17: req.OarAngle[17],
		OarAngle_18: req.OarAngle[18],
		OarAngle_19: req.OarAngle[19],
		OarAngle_20: req.OarAngle[20],
		OarAngle_21: req.OarAngle[21],
		OarAngle_22: req.OarAngle[22],
		OarAngle_23: req.OarAngle[23],
		OarAngle_24: req.OarAngle[24],
		OarAngle_25: req.OarAngle[25],
		OarAngle_26: req.OarAngle[26],
		OarAngle_27: req.OarAngle[27],
		OarAngle_28: req.OarAngle[28],
		OarAngle_29: req.OarAngle[29],
		OarAngle_30: req.OarAngle[30],
		OarAngle_31: req.OarAngle[31],
		OarAngle_32: req.OarAngle[32],
		OarAngle_33: req.OarAngle[33],
		OarAngle_34: req.OarAngle[34],
		OarAngle_35: req.OarAngle[35],
		OarAngle_36: req.OarAngle[36],
		OarAngle_37: req.OarAngle[37],
		OarAngle_38: req.OarAngle[38],
		OarAngle_39: req.OarAngle[39],
		OarAngle_40: req.OarAngle[40],
		OarAngle_41: req.OarAngle[41],
		OarAngle_42: req.OarAngle[42],
		OarAngle_43: req.OarAngle[43],
		OarAngle_44: req.OarAngle[44],
		OarAngle_45: req.OarAngle[45],
		OarAngle_46: req.OarAngle[46],
		OarAngle_47: req.OarAngle[47],
		OarAngle_48: req.OarAngle[48],
		OarAngle_49: req.OarAngle[49],
		OarAngle_50: req.OarAngle[50],

		HandleForce_0:  req.HandleForce[0],
		HandleForce_1:  req.HandleForce[1],
		HandleForce_2:  req.HandleForce[2],
		HandleForce_3:  req.HandleForce[3],
		HandleForce_4:  req.HandleForce[4],
		HandleForce_5:  req.HandleForce[5],
		HandleForce_6:  req.HandleForce[6],
		HandleForce_7:  req.HandleForce[7],
		HandleForce_8:  req.HandleForce[8],
		HandleForce_9:  req.HandleForce[9],
		HandleForce_10: req.HandleForce[10],
		HandleForce_11: req.HandleForce[11],
		HandleForce_12: req.HandleForce[12],
		HandleForce_13: req.HandleForce[13],
		HandleForce_14: req.HandleForce[14],
		HandleForce_15: req.HandleForce[15],
		HandleForce_16: req.HandleForce[16],
		HandleForce_17: req.HandleForce[17],
		HandleForce_18: req.HandleForce[18],
		HandleForce_19: req.HandleForce[19],
		HandleForce_20: req.HandleForce[20],
		HandleForce_21: req.HandleForce[21],
		HandleForce_22: req.HandleForce[22],
		HandleForce_23: req.HandleForce[23],
		HandleForce_24: req.HandleForce[24],
		HandleForce_25: req.HandleForce[25],
		HandleForce_26: req.HandleForce[26],
		HandleForce_27: req.HandleForce[27],
		HandleForce_28: req.HandleForce[28],
		HandleForce_29: req.HandleForce[29],
		HandleForce_30: req.HandleForce[30],
		HandleForce_31: req.HandleForce[31],
		HandleForce_32: req.HandleForce[32],
		HandleForce_33: req.HandleForce[33],
		HandleForce_34: req.HandleForce[34],
		HandleForce_35: req.HandleForce[35],
		HandleForce_36: req.HandleForce[36],
		HandleForce_37: req.HandleForce[37],
		HandleForce_38: req.HandleForce[38],
		HandleForce_39: req.HandleForce[39],
		HandleForce_40: req.HandleForce[40],
		HandleForce_41: req.HandleForce[41],
		HandleForce_42: req.HandleForce[42],
		HandleForce_43: req.HandleForce[43],
		HandleForce_44: req.HandleForce[44],
		HandleForce_45: req.HandleForce[45],
		HandleForce_46: req.HandleForce[46],
		HandleForce_47: req.HandleForce[47],
		HandleForce_48: req.HandleForce[48],
		HandleForce_49: req.HandleForce[49],
		HandleForce_50: req.HandleForce[50],

		VerticalAngleBoatRoll_0:  req.VerticalAngleBoatRoll[0],
		VerticalAngleBoatRoll_1:  req.VerticalAngleBoatRoll[1],
		VerticalAngleBoatRoll_2:  req.VerticalAngleBoatRoll[2],
		VerticalAngleBoatRoll_3:  req.VerticalAngleBoatRoll[3],
		VerticalAngleBoatRoll_4:  req.VerticalAngleBoatRoll[4],
		VerticalAngleBoatRoll_5:  req.VerticalAngleBoatRoll[5],
		VerticalAngleBoatRoll_6:  req.VerticalAngleBoatRoll[6],
		VerticalAngleBoatRoll_7:  req.VerticalAngleBoatRoll[7],
		VerticalAngleBoatRoll_8:  req.VerticalAngleBoatRoll[8],
		VerticalAngleBoatRoll_9:  req.VerticalAngleBoatRoll[9],
		VerticalAngleBoatRoll_10: req.VerticalAngleBoatRoll[10],
		VerticalAngleBoatRoll_11: req.VerticalAngleBoatRoll[11],
		VerticalAngleBoatRoll_12: req.VerticalAngleBoatRoll[12],
		VerticalAngleBoatRoll_13: req.VerticalAngleBoatRoll[13],
		VerticalAngleBoatRoll_14: req.VerticalAngleBoatRoll[14],
		VerticalAngleBoatRoll_15: req.VerticalAngleBoatRoll[15],
		VerticalAngleBoatRoll_16: req.VerticalAngleBoatRoll[16],
		VerticalAngleBoatRoll_17: req.VerticalAngleBoatRoll[17],
		VerticalAngleBoatRoll_18: req.VerticalAngleBoatRoll[18],
		VerticalAngleBoatRoll_19: req.VerticalAngleBoatRoll[19],
		VerticalAngleBoatRoll_20: req.VerticalAngleBoatRoll[20],
		VerticalAngleBoatRoll_21: req.VerticalAngleBoatRoll[21],
		VerticalAngleBoatRoll_22: req.VerticalAngleBoatRoll[22],
		VerticalAngleBoatRoll_23: req.VerticalAngleBoatRoll[23],
		VerticalAngleBoatRoll_24: req.VerticalAngleBoatRoll[24],
		VerticalAngleBoatRoll_25: req.VerticalAngleBoatRoll[25],
		VerticalAngleBoatRoll_26: req.VerticalAngleBoatRoll[26],
		VerticalAngleBoatRoll_27: req.VerticalAngleBoatRoll[27],
		VerticalAngleBoatRoll_28: req.VerticalAngleBoatRoll[28],
		VerticalAngleBoatRoll_29: req.VerticalAngleBoatRoll[29],
		VerticalAngleBoatRoll_30: req.VerticalAngleBoatRoll[30],
		VerticalAngleBoatRoll_31: req.VerticalAngleBoatRoll[31],
		VerticalAngleBoatRoll_32: req.VerticalAngleBoatRoll[32],
		VerticalAngleBoatRoll_33: req.VerticalAngleBoatRoll[33],
		VerticalAngleBoatRoll_34: req.VerticalAngleBoatRoll[34],
		VerticalAngleBoatRoll_35: req.VerticalAngleBoatRoll[35],
		VerticalAngleBoatRoll_36: req.VerticalAngleBoatRoll[36],
		VerticalAngleBoatRoll_37: req.VerticalAngleBoatRoll[37],
		VerticalAngleBoatRoll_38: req.VerticalAngleBoatRoll[38],
		VerticalAngleBoatRoll_39: req.VerticalAngleBoatRoll[39],
		VerticalAngleBoatRoll_40: req.VerticalAngleBoatRoll[40],
		VerticalAngleBoatRoll_41: req.VerticalAngleBoatRoll[41],
		VerticalAngleBoatRoll_42: req.VerticalAngleBoatRoll[42],
		VerticalAngleBoatRoll_43: req.VerticalAngleBoatRoll[43],
		VerticalAngleBoatRoll_44: req.VerticalAngleBoatRoll[44],
		VerticalAngleBoatRoll_45: req.VerticalAngleBoatRoll[45],
		VerticalAngleBoatRoll_46: req.VerticalAngleBoatRoll[46],
		VerticalAngleBoatRoll_47: req.VerticalAngleBoatRoll[47],
		VerticalAngleBoatRoll_48: req.VerticalAngleBoatRoll[48],
		VerticalAngleBoatRoll_49: req.VerticalAngleBoatRoll[49],
		VerticalAngleBoatRoll_50: req.VerticalAngleBoatRoll[50],

		LegsVelocity_0:  req.LegsVelocity[0],
		LegsVelocity_1:  req.LegsVelocity[1],
		LegsVelocity_2:  req.LegsVelocity[2],
		LegsVelocity_3:  req.LegsVelocity[3],
		LegsVelocity_4:  req.LegsVelocity[4],
		LegsVelocity_5:  req.LegsVelocity[5],
		LegsVelocity_6:  req.LegsVelocity[6],
		LegsVelocity_7:  req.LegsVelocity[7],
		LegsVelocity_8:  req.LegsVelocity[8],
		LegsVelocity_9:  req.LegsVelocity[9],
		LegsVelocity_10: req.LegsVelocity[10],
		LegsVelocity_11: req.LegsVelocity[11],
		LegsVelocity_12: req.LegsVelocity[12],
		LegsVelocity_13: req.LegsVelocity[13],
		LegsVelocity_14: req.LegsVelocity[14],
		LegsVelocity_15: req.LegsVelocity[15],
		LegsVelocity_16: req.LegsVelocity[16],
		LegsVelocity_17: req.LegsVelocity[17],
		LegsVelocity_18: req.LegsVelocity[18],
		LegsVelocity_19: req.LegsVelocity[19],
		LegsVelocity_20: req.LegsVelocity[20],
		LegsVelocity_21: req.LegsVelocity[21],
		LegsVelocity_22: req.LegsVelocity[22],
		LegsVelocity_23: req.LegsVelocity[23],
		LegsVelocity_24: req.LegsVelocity[24],
		LegsVelocity_25: req.LegsVelocity[25],
		LegsVelocity_26: req.LegsVelocity[26],
		LegsVelocity_27: req.LegsVelocity[27],
		LegsVelocity_28: req.LegsVelocity[28],
		LegsVelocity_29: req.LegsVelocity[29],
		LegsVelocity_30: req.LegsVelocity[30],
		LegsVelocity_31: req.LegsVelocity[31],
		LegsVelocity_32: req.LegsVelocity[32],
		LegsVelocity_33: req.LegsVelocity[33],
		LegsVelocity_34: req.LegsVelocity[34],
		LegsVelocity_35: req.LegsVelocity[35],
		LegsVelocity_36: req.LegsVelocity[36],
		LegsVelocity_37: req.LegsVelocity[37],
		LegsVelocity_38: req.LegsVelocity[38],
		LegsVelocity_39: req.LegsVelocity[39],
		LegsVelocity_40: req.LegsVelocity[40],
		LegsVelocity_41: req.LegsVelocity[41],
		LegsVelocity_42: req.LegsVelocity[42],
		LegsVelocity_43: req.LegsVelocity[43],
		LegsVelocity_44: req.LegsVelocity[44],
		LegsVelocity_45: req.LegsVelocity[45],
		LegsVelocity_46: req.LegsVelocity[46],
		LegsVelocity_47: req.LegsVelocity[47],
		LegsVelocity_48: req.LegsVelocity[48],
		LegsVelocity_49: req.LegsVelocity[49],
		LegsVelocity_50: req.LegsVelocity[50],

		HandleSpeed_0:  req.HandleSpeed[0],
		HandleSpeed_1:  req.HandleSpeed[1],
		HandleSpeed_2:  req.HandleSpeed[2],
		HandleSpeed_3:  req.HandleSpeed[3],
		HandleSpeed_4:  req.HandleSpeed[4],
		HandleSpeed_5:  req.HandleSpeed[5],
		HandleSpeed_6:  req.HandleSpeed[6],
		HandleSpeed_7:  req.HandleSpeed[7],
		HandleSpeed_8:  req.HandleSpeed[8],
		HandleSpeed_9:  req.HandleSpeed[9],
		HandleSpeed_10: req.HandleSpeed[10],
		HandleSpeed_11: req.HandleSpeed[11],
		HandleSpeed_12: req.HandleSpeed[12],
		HandleSpeed_13: req.HandleSpeed[13],
		HandleSpeed_14: req.HandleSpeed[14],
		HandleSpeed_15: req.HandleSpeed[15],
		HandleSpeed_16: req.HandleSpeed[16],
		HandleSpeed_17: req.HandleSpeed[17],
		HandleSpeed_18: req.HandleSpeed[18],
		HandleSpeed_19: req.HandleSpeed[19],
		HandleSpeed_20: req.HandleSpeed[20],
		HandleSpeed_21: req.HandleSpeed[21],
		HandleSpeed_22: req.HandleSpeed[22],
		HandleSpeed_23: req.HandleSpeed[23],
		HandleSpeed_24: req.HandleSpeed[24],
		HandleSpeed_25: req.HandleSpeed[25],
		HandleSpeed_26: req.HandleSpeed[26],
		HandleSpeed_27: req.HandleSpeed[27],
		HandleSpeed_28: req.HandleSpeed[28],
		HandleSpeed_29: req.HandleSpeed[29],
		HandleSpeed_30: req.HandleSpeed[30],
		HandleSpeed_31: req.HandleSpeed[31],
		HandleSpeed_32: req.HandleSpeed[32],
		HandleSpeed_33: req.HandleSpeed[33],
		HandleSpeed_34: req.HandleSpeed[34],
		HandleSpeed_35: req.HandleSpeed[35],
		HandleSpeed_36: req.HandleSpeed[36],
		HandleSpeed_37: req.HandleSpeed[37],
		HandleSpeed_38: req.HandleSpeed[38],
		HandleSpeed_39: req.HandleSpeed[39],
		HandleSpeed_40: req.HandleSpeed[40],
		HandleSpeed_41: req.HandleSpeed[41],
		HandleSpeed_42: req.HandleSpeed[42],
		HandleSpeed_43: req.HandleSpeed[43],
		HandleSpeed_44: req.HandleSpeed[44],
		HandleSpeed_45: req.HandleSpeed[45],
		HandleSpeed_46: req.HandleSpeed[46],
		HandleSpeed_47: req.HandleSpeed[47],
		HandleSpeed_48: req.HandleSpeed[48],
		HandleSpeed_49: req.HandleSpeed[49],
		HandleSpeed_50: req.HandleSpeed[50],

		HDF_0:  req.HdfFig[0],
		HDF_1:  req.HdfFig[1],
		HDF_2:  req.HdfFig[2],
		HDF_3:  req.HdfFig[3],
		HDF_4:  req.HdfFig[4],
		HDF_5:  req.HdfFig[5],
		HDF_6:  req.HdfFig[6],
		HDF_7:  req.HdfFig[7],
		HDF_8:  req.HdfFig[8],
		HDF_9:  req.HdfFig[9],
		HDF_10: req.HdfFig[10],
		HDF_11: req.HdfFig[11],
		HDF_12: req.HdfFig[12],
		HDF_13: req.HdfFig[13],
		HDF_14: req.HdfFig[14],
		HDF_15: req.HdfFig[15],
		HDF_16: req.HdfFig[16],
		HDF_17: req.HdfFig[17],
		HDF_18: req.HdfFig[18],
		HDF_19: req.HdfFig[19],
		HDF_20: req.HdfFig[20],
		HDF_21: req.HdfFig[21],
		HDF_22: req.HdfFig[22],
		HDF_23: req.HdfFig[23],
		HDF_24: req.HdfFig[24],
		HDF_25: req.HdfFig[25],
		HDF_26: req.HdfFig[26],
		HDF_27: req.HdfFig[27],
		HDF_28: req.HdfFig[28],
		HDF_29: req.HdfFig[29],
		HDF_30: req.HdfFig[30],
		HDF_31: req.HdfFig[31],
		HDF_32: req.HdfFig[32],
		HDF_33: req.HdfFig[33],
		HDF_34: req.HdfFig[34],
		HDF_35: req.HdfFig[35],
		HDF_36: req.HdfFig[36],
		HDF_37: req.HdfFig[37],
		HDF_38: req.HdfFig[38],
		HDF_39: req.HdfFig[39],
		HDF_40: req.HdfFig[40],
		HDF_41: req.HdfFig[41],
		HDF_42: req.HdfFig[42],
		HDF_43: req.HdfFig[43],
		HDF_44: req.HdfFig[44],
		HDF_45: req.HdfFig[45],
		HDF_46: req.HdfFig[46],
		HDF_47: req.HdfFig[47],
		HDF_48: req.HdfFig[48],
		HDF_49: req.HdfFig[49],
		HDF_50: req.HdfFig[50],

		BladeDF_0:  req.BladeDf[0],
		BladeDF_1:  req.BladeDf[1],
		BladeDF_2:  req.BladeDf[2],
		BladeDF_3:  req.BladeDf[3],
		BladeDF_4:  req.BladeDf[4],
		BladeDF_5:  req.BladeDf[5],
		BladeDF_6:  req.BladeDf[6],
		BladeDF_7:  req.BladeDf[7],
		BladeDF_8:  req.BladeDf[8],
		BladeDF_9:  req.BladeDf[9],
		BladeDF_10: req.BladeDf[10],
		BladeDF_11: req.BladeDf[11],
		BladeDF_12: req.BladeDf[12],
		BladeDF_13: req.BladeDf[13],
		BladeDF_14: req.BladeDf[14],
		BladeDF_15: req.BladeDf[15],
		BladeDF_16: req.BladeDf[16],
		BladeDF_17: req.BladeDf[17],
		BladeDF_18: req.BladeDf[18],
		BladeDF_19: req.BladeDf[19],
		BladeDF_20: req.BladeDf[20],
		BladeDF_21: req.BladeDf[21],
		BladeDF_22: req.BladeDf[22],
		BladeDF_23: req.BladeDf[23],
		BladeDF_24: req.BladeDf[24],
		BladeDF_25: req.BladeDf[25],
		BladeDF_26: req.BladeDf[26],
		BladeDF_27: req.BladeDf[27],
		BladeDF_28: req.BladeDf[28],
		BladeDF_29: req.BladeDf[29],
		BladeDF_30: req.BladeDf[30],
		BladeDF_31: req.BladeDf[31],
		BladeDF_32: req.BladeDf[32],
		BladeDF_33: req.BladeDf[33],
		BladeDF_34: req.BladeDf[34],
		BladeDF_35: req.BladeDf[35],
		BladeDF_36: req.BladeDf[36],
		BladeDF_37: req.BladeDf[37],
		BladeDF_38: req.BladeDf[38],
		BladeDF_39: req.BladeDf[39],
		BladeDF_40: req.BladeDf[40],
		BladeDF_41: req.BladeDf[41],
		BladeDF_42: req.BladeDf[42],
		BladeDF_43: req.BladeDf[43],
		BladeDF_44: req.BladeDf[44],
		BladeDF_45: req.BladeDf[45],
		BladeDF_46: req.BladeDf[46],
		BladeDF_47: req.BladeDf[47],
		BladeDF_48: req.BladeDf[48],
		BladeDF_49: req.BladeDf[49],
		BladeDF_50: req.BladeDf[50],

		Velocity_0:          req.Velocity[0],
		Velocity_1:          req.Velocity[1],
		Velocity_2:          req.Velocity[2],
		Velocity_3:          req.Velocity[3],
		Velocity_4:          req.Velocity[4],
		Velocity_5:          req.Velocity[5],
		Velocity_6:          req.Velocity[6],
		Velocity_7:          req.Velocity[7],
		Velocity_8:          req.Velocity[8],
		Velocity_9:          req.Velocity[9],
		Velocity_10:         req.Velocity[10],
		Velocity_11:         req.Velocity[11],
		Velocity_12:         req.Velocity[12],
		Velocity_13:         req.Velocity[13],
		Velocity_14:         req.Velocity[14],
		Velocity_15:         req.Velocity[15],
		Velocity_16:         req.Velocity[16],
		Velocity_17:         req.Velocity[17],
		Velocity_18:         req.Velocity[18],
		Velocity_19:         req.Velocity[19],
		Velocity_20:         req.Velocity[20],
		Velocity_21:         req.Velocity[21],
		Velocity_22:         req.Velocity[22],
		Velocity_23:         req.Velocity[23],
		Velocity_24:         req.Velocity[24],
		Velocity_25:         req.Velocity[25],
		Velocity_26:         req.Velocity[26],
		Velocity_27:         req.Velocity[27],
		Velocity_28:         req.Velocity[28],
		Velocity_29:         req.Velocity[29],
		Velocity_30:         req.Velocity[30],
		Velocity_31:         req.Velocity[31],
		Velocity_32:         req.Velocity[32],
		Velocity_33:         req.Velocity[33],
		Velocity_34:         req.Velocity[34],
		Velocity_35:         req.Velocity[35],
		Velocity_36:         req.Velocity[36],
		Velocity_37:         req.Velocity[37],
		Velocity_38:         req.Velocity[38],
		Velocity_39:         req.Velocity[39],
		Velocity_40:         req.Velocity[40],
		Velocity_41:         req.Velocity[41],
		Velocity_42:         req.Velocity[42],
		Velocity_43:         req.Velocity[43],
		Velocity_44:         req.Velocity[44],
		Velocity_45:         req.Velocity[45],
		Velocity_46:         req.Velocity[46],
		Velocity_47:         req.Velocity[47],
		Velocity_48:         req.Velocity[48],
		Velocity_49:         req.Velocity[49],
		Velocity_50:         req.Velocity[50],
		BoatAcceleration_0:  req.BoatAcceleration[0],
		BoatAcceleration_1:  req.BoatAcceleration[1],
		BoatAcceleration_2:  req.BoatAcceleration[2],
		BoatAcceleration_3:  req.BoatAcceleration[3],
		BoatAcceleration_4:  req.BoatAcceleration[4],
		BoatAcceleration_5:  req.BoatAcceleration[5],
		BoatAcceleration_6:  req.BoatAcceleration[6],
		BoatAcceleration_7:  req.BoatAcceleration[7],
		BoatAcceleration_8:  req.BoatAcceleration[8],
		BoatAcceleration_9:  req.BoatAcceleration[9],
		BoatAcceleration_10: req.BoatAcceleration[10],
		BoatAcceleration_11: req.BoatAcceleration[11],
		BoatAcceleration_12: req.BoatAcceleration[12],
		BoatAcceleration_13: req.BoatAcceleration[13],
		BoatAcceleration_14: req.BoatAcceleration[14],
		BoatAcceleration_15: req.BoatAcceleration[15],
		BoatAcceleration_16: req.BoatAcceleration[16],
		BoatAcceleration_17: req.BoatAcceleration[17],
		BoatAcceleration_18: req.BoatAcceleration[18],
		BoatAcceleration_19: req.BoatAcceleration[19],
		BoatAcceleration_20: req.BoatAcceleration[20],
		BoatAcceleration_21: req.BoatAcceleration[21],
		BoatAcceleration_22: req.BoatAcceleration[22],
		BoatAcceleration_23: req.BoatAcceleration[23],
		BoatAcceleration_24: req.BoatAcceleration[24],
		BoatAcceleration_25: req.BoatAcceleration[25],
		BoatAcceleration_26: req.BoatAcceleration[26],
		BoatAcceleration_27: req.BoatAcceleration[27],
		BoatAcceleration_28: req.BoatAcceleration[28],
		BoatAcceleration_29: req.BoatAcceleration[29],
		BoatAcceleration_30: req.BoatAcceleration[30],
		BoatAcceleration_31: req.BoatAcceleration[31],
		BoatAcceleration_32: req.BoatAcceleration[32],
		BoatAcceleration_33: req.BoatAcceleration[33],
		BoatAcceleration_34: req.BoatAcceleration[34],
		BoatAcceleration_35: req.BoatAcceleration[35],
		BoatAcceleration_36: req.BoatAcceleration[36],
		BoatAcceleration_37: req.BoatAcceleration[37],
		BoatAcceleration_38: req.BoatAcceleration[38],
		BoatAcceleration_39: req.BoatAcceleration[39],
		BoatAcceleration_40: req.BoatAcceleration[40],
		BoatAcceleration_41: req.BoatAcceleration[41],
		BoatAcceleration_42: req.BoatAcceleration[42],
		BoatAcceleration_43: req.BoatAcceleration[43],
		BoatAcceleration_44: req.BoatAcceleration[44],
		BoatAcceleration_45: req.BoatAcceleration[45],
		BoatAcceleration_46: req.BoatAcceleration[46],
		BoatAcceleration_47: req.BoatAcceleration[47],
		BoatAcceleration_48: req.BoatAcceleration[48],
		BoatAcceleration_49: req.BoatAcceleration[49],
		BoatAcceleration_50: req.BoatAcceleration[50],

		VelocityRel_0:  req.VelocityRel[0],
		VelocityRel_1:  req.VelocityRel[1],
		VelocityRel_2:  req.VelocityRel[2],
		VelocityRel_3:  req.VelocityRel[3],
		VelocityRel_4:  req.VelocityRel[4],
		VelocityRel_5:  req.VelocityRel[5],
		VelocityRel_6:  req.VelocityRel[6],
		VelocityRel_7:  req.VelocityRel[7],
		VelocityRel_8:  req.VelocityRel[8],
		VelocityRel_9:  req.VelocityRel[9],
		VelocityRel_10: req.VelocityRel[10],
		VelocityRel_11: req.VelocityRel[11],
		VelocityRel_12: req.VelocityRel[12],
		VelocityRel_13: req.VelocityRel[13],
		VelocityRel_14: req.VelocityRel[14],
		VelocityRel_15: req.VelocityRel[15],
		VelocityRel_16: req.VelocityRel[16],
		VelocityRel_17: req.VelocityRel[17],
		VelocityRel_18: req.VelocityRel[18],
		VelocityRel_19: req.VelocityRel[19],
		VelocityRel_20: req.VelocityRel[20],
		VelocityRel_21: req.VelocityRel[21],
		VelocityRel_22: req.VelocityRel[22],
		VelocityRel_23: req.VelocityRel[23],
		VelocityRel_24: req.VelocityRel[24],
		VelocityRel_25: req.VelocityRel[25],
		VelocityRel_26: req.VelocityRel[26],
		VelocityRel_27: req.VelocityRel[27],
		VelocityRel_28: req.VelocityRel[28],
		VelocityRel_29: req.VelocityRel[29],
		VelocityRel_30: req.VelocityRel[30],
		VelocityRel_31: req.VelocityRel[31],
		VelocityRel_32: req.VelocityRel[32],
		VelocityRel_33: req.VelocityRel[33],
		VelocityRel_34: req.VelocityRel[34],
		VelocityRel_35: req.VelocityRel[35],
		VelocityRel_36: req.VelocityRel[36],
		VelocityRel_37: req.VelocityRel[37],
		VelocityRel_38: req.VelocityRel[38],
		VelocityRel_39: req.VelocityRel[39],
		VelocityRel_40: req.VelocityRel[40],
		VelocityRel_41: req.VelocityRel[41],
		VelocityRel_42: req.VelocityRel[42],
		VelocityRel_43: req.VelocityRel[43],
		VelocityRel_44: req.VelocityRel[44],
		VelocityRel_45: req.VelocityRel[45],
		VelocityRel_46: req.VelocityRel[46],
		VelocityRel_47: req.VelocityRel[47],
		VelocityRel_48: req.VelocityRel[48],
		VelocityRel_49: req.VelocityRel[49],
		VelocityRel_50: req.VelocityRel[50],
	}

	if err := dao.Model(&model.SampleMetrics{}).Create(&sampleMetrics).Error; err != nil {
		logger.LogrusObj.Error("Insert SampleMetrics Error:" + err.Error())
		return 0, err
	}
	return sampleMetrics.Id, nil
}

func (dao *RowDataDao) GetSampleMetricsByAthleteTrainingId(req *pb.SampleMetricsGetByAthleteTrainingIdRequest) (sm []*model.SampleMetrics, err error) {
	query := dao.Model(&model.SampleMetrics{})
	if req.Set {
		query = query.Where("athlete_training_id=?", req.AthleteTrainingId)
	}
	err = query.Find(&sm).Error
	return
}

func BuildSampleMetrics(req model.SampleMetrics) *pb.SampleMetricsModel {
	var oarAngle [51]float64
	var handleForce [51]float64
	var verticalAngleBoatRoll [51]float64
	var legsVelocity [51]float64
	var handleSpeed [51]float64
	var hdfFig [51]float64
	var bladeDf [51]float64
	var velocity [51]float64
	var boatAcceleration [51]float64
	var velocityRel [51]float64

	oarAngle[0] = req.OarAngle_0
	oarAngle[1] = req.OarAngle_1
	oarAngle[2] = req.OarAngle_2
	oarAngle[3] = req.OarAngle_3
	oarAngle[4] = req.OarAngle_4
	oarAngle[5] = req.OarAngle_5
	oarAngle[6] = req.OarAngle_6
	oarAngle[7] = req.OarAngle_7
	oarAngle[8] = req.OarAngle_8
	oarAngle[9] = req.OarAngle_9
	oarAngle[10] = req.OarAngle_10
	oarAngle[11] = req.OarAngle_11
	oarAngle[12] = req.OarAngle_12
	oarAngle[13] = req.OarAngle_13
	oarAngle[14] = req.OarAngle_14
	oarAngle[15] = req.OarAngle_15
	oarAngle[16] = req.OarAngle_16
	oarAngle[17] = req.OarAngle_17
	oarAngle[18] = req.OarAngle_18
	oarAngle[19] = req.OarAngle_19
	oarAngle[20] = req.OarAngle_20
	oarAngle[21] = req.OarAngle_21
	oarAngle[22] = req.OarAngle_22
	oarAngle[23] = req.OarAngle_23
	oarAngle[24] = req.OarAngle_24
	oarAngle[25] = req.OarAngle_25
	oarAngle[26] = req.OarAngle_26
	oarAngle[27] = req.OarAngle_27
	oarAngle[28] = req.OarAngle_28
	oarAngle[29] = req.OarAngle_29
	oarAngle[30] = req.OarAngle_30
	oarAngle[31] = req.OarAngle_31
	oarAngle[32] = req.OarAngle_32
	oarAngle[33] = req.OarAngle_33
	oarAngle[34] = req.OarAngle_34
	oarAngle[35] = req.OarAngle_35
	oarAngle[36] = req.OarAngle_36
	oarAngle[37] = req.OarAngle_37
	oarAngle[38] = req.OarAngle_38
	oarAngle[39] = req.OarAngle_39
	oarAngle[40] = req.OarAngle_40
	oarAngle[41] = req.OarAngle_41
	oarAngle[42] = req.OarAngle_42
	oarAngle[43] = req.OarAngle_43
	oarAngle[44] = req.OarAngle_44
	oarAngle[45] = req.OarAngle_45
	oarAngle[46] = req.OarAngle_46
	oarAngle[47] = req.OarAngle_47
	oarAngle[48] = req.OarAngle_48
	oarAngle[49] = req.OarAngle_49
	oarAngle[50] = req.OarAngle_50

	handleForce[0] = req.HandleForce_0
	handleForce[1] = req.HandleForce_1
	handleForce[2] = req.HandleForce_2
	handleForce[3] = req.HandleForce_3
	handleForce[4] = req.HandleForce_4
	handleForce[5] = req.HandleForce_5
	handleForce[6] = req.HandleForce_6
	handleForce[7] = req.HandleForce_7
	handleForce[8] = req.HandleForce_8
	handleForce[9] = req.HandleForce_9
	handleForce[10] = req.HandleForce_10
	handleForce[11] = req.HandleForce_11
	handleForce[12] = req.HandleForce_12
	handleForce[13] = req.HandleForce_13
	handleForce[14] = req.HandleForce_14
	handleForce[15] = req.HandleForce_15
	handleForce[16] = req.HandleForce_16
	handleForce[17] = req.HandleForce_17
	handleForce[18] = req.HandleForce_18
	handleForce[19] = req.HandleForce_19
	handleForce[20] = req.HandleForce_20
	handleForce[21] = req.HandleForce_21
	handleForce[22] = req.HandleForce_22
	handleForce[23] = req.HandleForce_23
	handleForce[24] = req.HandleForce_24
	handleForce[25] = req.HandleForce_25
	handleForce[26] = req.HandleForce_26
	handleForce[27] = req.HandleForce_27
	handleForce[28] = req.HandleForce_28
	handleForce[29] = req.HandleForce_29
	handleForce[30] = req.HandleForce_30
	handleForce[31] = req.HandleForce_31
	handleForce[32] = req.HandleForce_32
	handleForce[33] = req.HandleForce_33
	handleForce[34] = req.HandleForce_34
	handleForce[35] = req.HandleForce_35
	handleForce[36] = req.HandleForce_36
	handleForce[37] = req.HandleForce_37
	handleForce[38] = req.HandleForce_38
	handleForce[39] = req.HandleForce_39
	handleForce[40] = req.HandleForce_40
	handleForce[41] = req.HandleForce_41
	handleForce[42] = req.HandleForce_42
	handleForce[43] = req.HandleForce_43
	handleForce[44] = req.HandleForce_44
	handleForce[45] = req.HandleForce_45
	handleForce[46] = req.HandleForce_46
	handleForce[47] = req.HandleForce_47
	handleForce[48] = req.HandleForce_48
	handleForce[49] = req.HandleForce_49
	handleForce[50] = req.HandleForce_50

	verticalAngleBoatRoll[0] = req.VerticalAngleBoatRoll_0
	verticalAngleBoatRoll[1] = req.VerticalAngleBoatRoll_1
	verticalAngleBoatRoll[2] = req.VerticalAngleBoatRoll_2
	verticalAngleBoatRoll[3] = req.VerticalAngleBoatRoll_3
	verticalAngleBoatRoll[4] = req.VerticalAngleBoatRoll_4
	verticalAngleBoatRoll[5] = req.VerticalAngleBoatRoll_5
	verticalAngleBoatRoll[6] = req.VerticalAngleBoatRoll_6
	verticalAngleBoatRoll[7] = req.VerticalAngleBoatRoll_7
	verticalAngleBoatRoll[8] = req.VerticalAngleBoatRoll_8
	verticalAngleBoatRoll[9] = req.VerticalAngleBoatRoll_9
	verticalAngleBoatRoll[10] = req.VerticalAngleBoatRoll_10
	verticalAngleBoatRoll[11] = req.VerticalAngleBoatRoll_11
	verticalAngleBoatRoll[12] = req.VerticalAngleBoatRoll_12
	verticalAngleBoatRoll[13] = req.VerticalAngleBoatRoll_13
	verticalAngleBoatRoll[14] = req.VerticalAngleBoatRoll_14
	verticalAngleBoatRoll[15] = req.VerticalAngleBoatRoll_15
	verticalAngleBoatRoll[16] = req.VerticalAngleBoatRoll_16
	verticalAngleBoatRoll[17] = req.VerticalAngleBoatRoll_17
	verticalAngleBoatRoll[18] = req.VerticalAngleBoatRoll_18
	verticalAngleBoatRoll[19] = req.VerticalAngleBoatRoll_19
	verticalAngleBoatRoll[20] = req.VerticalAngleBoatRoll_20
	verticalAngleBoatRoll[21] = req.VerticalAngleBoatRoll_21
	verticalAngleBoatRoll[22] = req.VerticalAngleBoatRoll_22
	verticalAngleBoatRoll[23] = req.VerticalAngleBoatRoll_23
	verticalAngleBoatRoll[24] = req.VerticalAngleBoatRoll_24
	verticalAngleBoatRoll[25] = req.VerticalAngleBoatRoll_25
	verticalAngleBoatRoll[26] = req.VerticalAngleBoatRoll_26
	verticalAngleBoatRoll[27] = req.VerticalAngleBoatRoll_27
	verticalAngleBoatRoll[28] = req.VerticalAngleBoatRoll_28
	verticalAngleBoatRoll[29] = req.VerticalAngleBoatRoll_29
	verticalAngleBoatRoll[30] = req.VerticalAngleBoatRoll_30
	verticalAngleBoatRoll[31] = req.VerticalAngleBoatRoll_31
	verticalAngleBoatRoll[32] = req.VerticalAngleBoatRoll_32
	verticalAngleBoatRoll[33] = req.VerticalAngleBoatRoll_33
	verticalAngleBoatRoll[34] = req.VerticalAngleBoatRoll_34
	verticalAngleBoatRoll[35] = req.VerticalAngleBoatRoll_35
	verticalAngleBoatRoll[36] = req.VerticalAngleBoatRoll_36
	verticalAngleBoatRoll[37] = req.VerticalAngleBoatRoll_37
	verticalAngleBoatRoll[38] = req.VerticalAngleBoatRoll_38
	verticalAngleBoatRoll[39] = req.VerticalAngleBoatRoll_39
	verticalAngleBoatRoll[40] = req.VerticalAngleBoatRoll_40
	verticalAngleBoatRoll[41] = req.VerticalAngleBoatRoll_41
	verticalAngleBoatRoll[42] = req.VerticalAngleBoatRoll_42
	verticalAngleBoatRoll[43] = req.VerticalAngleBoatRoll_43
	verticalAngleBoatRoll[44] = req.VerticalAngleBoatRoll_44
	verticalAngleBoatRoll[45] = req.VerticalAngleBoatRoll_45
	verticalAngleBoatRoll[46] = req.VerticalAngleBoatRoll_46
	verticalAngleBoatRoll[47] = req.VerticalAngleBoatRoll_47
	verticalAngleBoatRoll[48] = req.VerticalAngleBoatRoll_48
	verticalAngleBoatRoll[49] = req.VerticalAngleBoatRoll_49
	verticalAngleBoatRoll[50] = req.VerticalAngleBoatRoll_50

	legsVelocity[0] = req.LegsVelocity_0
	legsVelocity[1] = req.LegsVelocity_1
	legsVelocity[2] = req.LegsVelocity_2
	legsVelocity[3] = req.LegsVelocity_3
	legsVelocity[4] = req.LegsVelocity_4
	legsVelocity[5] = req.LegsVelocity_5
	legsVelocity[6] = req.LegsVelocity_6
	legsVelocity[7] = req.LegsVelocity_7
	legsVelocity[8] = req.LegsVelocity_8
	legsVelocity[9] = req.LegsVelocity_9
	legsVelocity[10] = req.LegsVelocity_10
	legsVelocity[11] = req.LegsVelocity_11
	legsVelocity[12] = req.LegsVelocity_12
	legsVelocity[13] = req.LegsVelocity_13
	legsVelocity[14] = req.LegsVelocity_14
	legsVelocity[15] = req.LegsVelocity_15
	legsVelocity[16] = req.LegsVelocity_16
	legsVelocity[17] = req.LegsVelocity_17
	legsVelocity[18] = req.LegsVelocity_18
	legsVelocity[19] = req.LegsVelocity_19
	legsVelocity[20] = req.LegsVelocity_20
	legsVelocity[21] = req.LegsVelocity_21
	legsVelocity[22] = req.LegsVelocity_22
	legsVelocity[23] = req.LegsVelocity_23
	legsVelocity[24] = req.LegsVelocity_24
	legsVelocity[25] = req.LegsVelocity_25
	legsVelocity[26] = req.LegsVelocity_26
	legsVelocity[27] = req.LegsVelocity_27
	legsVelocity[28] = req.LegsVelocity_28
	legsVelocity[29] = req.LegsVelocity_29
	legsVelocity[30] = req.LegsVelocity_30
	legsVelocity[31] = req.LegsVelocity_31
	legsVelocity[32] = req.LegsVelocity_32
	legsVelocity[33] = req.LegsVelocity_33
	legsVelocity[34] = req.LegsVelocity_34
	legsVelocity[35] = req.LegsVelocity_35
	legsVelocity[36] = req.LegsVelocity_36
	legsVelocity[37] = req.LegsVelocity_37
	legsVelocity[38] = req.LegsVelocity_38
	legsVelocity[39] = req.LegsVelocity_39
	legsVelocity[40] = req.LegsVelocity_40
	legsVelocity[41] = req.LegsVelocity_41
	legsVelocity[42] = req.LegsVelocity_42
	legsVelocity[43] = req.LegsVelocity_43
	legsVelocity[44] = req.LegsVelocity_44
	legsVelocity[45] = req.LegsVelocity_45
	legsVelocity[46] = req.LegsVelocity_46
	legsVelocity[47] = req.LegsVelocity_47
	legsVelocity[48] = req.LegsVelocity_48
	legsVelocity[49] = req.LegsVelocity_49
	legsVelocity[50] = req.LegsVelocity_50

	handleSpeed[0] = req.HandleSpeed_0
	handleSpeed[1] = req.HandleSpeed_1
	handleSpeed[2] = req.HandleSpeed_2
	handleSpeed[3] = req.HandleSpeed_3
	handleSpeed[4] = req.HandleSpeed_4
	handleSpeed[5] = req.HandleSpeed_5
	handleSpeed[6] = req.HandleSpeed_6
	handleSpeed[7] = req.HandleSpeed_7
	handleSpeed[8] = req.HandleSpeed_8
	handleSpeed[9] = req.HandleSpeed_9
	handleSpeed[10] = req.HandleSpeed_10
	handleSpeed[11] = req.HandleSpeed_11
	handleSpeed[12] = req.HandleSpeed_12
	handleSpeed[13] = req.HandleSpeed_13
	handleSpeed[14] = req.HandleSpeed_14
	handleSpeed[15] = req.HandleSpeed_15
	handleSpeed[16] = req.HandleSpeed_16
	handleSpeed[17] = req.HandleSpeed_17
	handleSpeed[18] = req.HandleSpeed_18
	handleSpeed[19] = req.HandleSpeed_19
	handleSpeed[20] = req.HandleSpeed_20
	handleSpeed[21] = req.HandleSpeed_21
	handleSpeed[22] = req.HandleSpeed_22
	handleSpeed[23] = req.HandleSpeed_23
	handleSpeed[24] = req.HandleSpeed_24
	handleSpeed[25] = req.HandleSpeed_25
	handleSpeed[26] = req.HandleSpeed_26
	handleSpeed[27] = req.HandleSpeed_27
	handleSpeed[28] = req.HandleSpeed_28
	handleSpeed[29] = req.HandleSpeed_29
	handleSpeed[30] = req.HandleSpeed_30
	handleSpeed[31] = req.HandleSpeed_31
	handleSpeed[32] = req.HandleSpeed_32
	handleSpeed[33] = req.HandleSpeed_33
	handleSpeed[34] = req.HandleSpeed_34
	handleSpeed[35] = req.HandleSpeed_35
	handleSpeed[36] = req.HandleSpeed_36
	handleSpeed[37] = req.HandleSpeed_37
	handleSpeed[38] = req.HandleSpeed_38
	handleSpeed[39] = req.HandleSpeed_39
	handleSpeed[40] = req.HandleSpeed_40
	handleSpeed[41] = req.HandleSpeed_41
	handleSpeed[42] = req.HandleSpeed_42
	handleSpeed[43] = req.HandleSpeed_43
	handleSpeed[44] = req.HandleSpeed_44
	handleSpeed[45] = req.HandleSpeed_45
	handleSpeed[46] = req.HandleSpeed_46
	handleSpeed[47] = req.HandleSpeed_47
	handleSpeed[48] = req.HandleSpeed_48
	handleSpeed[49] = req.HandleSpeed_49
	handleSpeed[50] = req.HandleSpeed_50

	hdfFig[0] = req.HDF_0
	hdfFig[1] = req.HDF_1
	hdfFig[2] = req.HDF_2
	hdfFig[3] = req.HDF_3
	hdfFig[4] = req.HDF_4
	hdfFig[5] = req.HDF_5
	hdfFig[6] = req.HDF_6
	hdfFig[7] = req.HDF_7
	hdfFig[8] = req.HDF_8
	hdfFig[9] = req.HDF_9
	hdfFig[10] = req.HDF_10
	hdfFig[11] = req.HDF_11
	hdfFig[12] = req.HDF_12
	hdfFig[13] = req.HDF_13
	hdfFig[14] = req.HDF_14
	hdfFig[15] = req.HDF_15
	hdfFig[16] = req.HDF_16
	hdfFig[17] = req.HDF_17
	hdfFig[18] = req.HDF_18
	hdfFig[19] = req.HDF_19
	hdfFig[20] = req.HDF_20
	hdfFig[21] = req.HDF_21
	hdfFig[22] = req.HDF_22
	hdfFig[23] = req.HDF_23
	hdfFig[24] = req.HDF_24
	hdfFig[25] = req.HDF_25
	hdfFig[26] = req.HDF_26
	hdfFig[27] = req.HDF_27
	hdfFig[28] = req.HDF_28
	hdfFig[29] = req.HDF_29
	hdfFig[30] = req.HDF_30
	hdfFig[31] = req.HDF_31
	hdfFig[32] = req.HDF_32
	hdfFig[33] = req.HDF_33
	hdfFig[34] = req.HDF_34
	hdfFig[35] = req.HDF_35
	hdfFig[36] = req.HDF_36
	hdfFig[37] = req.HDF_37
	hdfFig[38] = req.HDF_38
	hdfFig[39] = req.HDF_39
	hdfFig[40] = req.HDF_40
	hdfFig[41] = req.HDF_41
	hdfFig[42] = req.HDF_42
	hdfFig[43] = req.HDF_43
	hdfFig[44] = req.HDF_44
	hdfFig[45] = req.HDF_45
	hdfFig[46] = req.HDF_46
	hdfFig[47] = req.HDF_47
	hdfFig[48] = req.HDF_48
	hdfFig[49] = req.HDF_49
	hdfFig[50] = req.HDF_50

	bladeDf[0] = req.BladeDF_0
	bladeDf[1] = req.BladeDF_1
	bladeDf[2] = req.BladeDF_2
	bladeDf[3] = req.BladeDF_3
	bladeDf[4] = req.BladeDF_4
	bladeDf[5] = req.BladeDF_5
	bladeDf[6] = req.BladeDF_6
	bladeDf[7] = req.BladeDF_7
	bladeDf[8] = req.BladeDF_8
	bladeDf[9] = req.BladeDF_9
	bladeDf[10] = req.BladeDF_10
	bladeDf[11] = req.BladeDF_11
	bladeDf[12] = req.BladeDF_12
	bladeDf[13] = req.BladeDF_13
	bladeDf[14] = req.BladeDF_14
	bladeDf[15] = req.BladeDF_15
	bladeDf[16] = req.BladeDF_16
	bladeDf[17] = req.BladeDF_17
	bladeDf[18] = req.BladeDF_18
	bladeDf[19] = req.BladeDF_19
	bladeDf[20] = req.BladeDF_20
	bladeDf[21] = req.BladeDF_21
	bladeDf[22] = req.BladeDF_22
	bladeDf[23] = req.BladeDF_23
	bladeDf[24] = req.BladeDF_24
	bladeDf[25] = req.BladeDF_25
	bladeDf[26] = req.BladeDF_26
	bladeDf[27] = req.BladeDF_27
	bladeDf[28] = req.BladeDF_28
	bladeDf[29] = req.BladeDF_29
	bladeDf[30] = req.BladeDF_30
	bladeDf[31] = req.BladeDF_31
	bladeDf[32] = req.BladeDF_32
	bladeDf[33] = req.BladeDF_33
	bladeDf[34] = req.BladeDF_34
	bladeDf[35] = req.BladeDF_35
	bladeDf[36] = req.BladeDF_36
	bladeDf[37] = req.BladeDF_37
	bladeDf[38] = req.BladeDF_38
	bladeDf[39] = req.BladeDF_39
	bladeDf[40] = req.BladeDF_40
	bladeDf[41] = req.BladeDF_41
	bladeDf[42] = req.BladeDF_42
	bladeDf[43] = req.BladeDF_43
	bladeDf[44] = req.BladeDF_44
	bladeDf[45] = req.BladeDF_45
	bladeDf[46] = req.BladeDF_46
	bladeDf[47] = req.BladeDF_47
	bladeDf[48] = req.BladeDF_48
	bladeDf[49] = req.BladeDF_49
	bladeDf[50] = req.BladeDF_50

	velocity[0] = req.Velocity_0
	velocity[1] = req.Velocity_1
	velocity[2] = req.Velocity_2
	velocity[3] = req.Velocity_3
	velocity[4] = req.Velocity_4
	velocity[5] = req.Velocity_5
	velocity[6] = req.Velocity_6
	velocity[7] = req.Velocity_7
	velocity[8] = req.Velocity_8
	velocity[9] = req.Velocity_9
	velocity[10] = req.Velocity_10
	velocity[11] = req.Velocity_11
	velocity[12] = req.Velocity_12
	velocity[13] = req.Velocity_13
	velocity[14] = req.Velocity_14
	velocity[15] = req.Velocity_15
	velocity[16] = req.Velocity_16
	velocity[17] = req.Velocity_17
	velocity[18] = req.Velocity_18
	velocity[19] = req.Velocity_19
	velocity[20] = req.Velocity_20
	velocity[21] = req.Velocity_21
	velocity[22] = req.Velocity_22
	velocity[23] = req.Velocity_23
	velocity[24] = req.Velocity_24
	velocity[25] = req.Velocity_25
	velocity[26] = req.Velocity_26
	velocity[27] = req.Velocity_27
	velocity[28] = req.Velocity_28
	velocity[29] = req.Velocity_29
	velocity[30] = req.Velocity_30
	velocity[31] = req.Velocity_31
	velocity[32] = req.Velocity_32
	velocity[33] = req.Velocity_33
	velocity[34] = req.Velocity_34
	velocity[35] = req.Velocity_35
	velocity[36] = req.Velocity_36
	velocity[37] = req.Velocity_37
	velocity[38] = req.Velocity_38
	velocity[39] = req.Velocity_39
	velocity[40] = req.Velocity_40
	velocity[41] = req.Velocity_41
	velocity[42] = req.Velocity_42
	velocity[43] = req.Velocity_43
	velocity[44] = req.Velocity_44
	velocity[45] = req.Velocity_45
	velocity[46] = req.Velocity_46
	velocity[47] = req.Velocity_47
	velocity[48] = req.Velocity_48
	velocity[49] = req.Velocity_49
	velocity[50] = req.Velocity_50

	boatAcceleration[0] = req.BoatAcceleration_0
	boatAcceleration[1] = req.BoatAcceleration_1
	boatAcceleration[2] = req.BoatAcceleration_2
	boatAcceleration[3] = req.BoatAcceleration_3
	boatAcceleration[4] = req.BoatAcceleration_4
	boatAcceleration[5] = req.BoatAcceleration_5
	boatAcceleration[6] = req.BoatAcceleration_6
	boatAcceleration[7] = req.BoatAcceleration_7
	boatAcceleration[8] = req.BoatAcceleration_8
	boatAcceleration[9] = req.BoatAcceleration_9
	boatAcceleration[10] = req.BoatAcceleration_10
	boatAcceleration[11] = req.BoatAcceleration_11
	boatAcceleration[12] = req.BoatAcceleration_12
	boatAcceleration[13] = req.BoatAcceleration_13
	boatAcceleration[14] = req.BoatAcceleration_14
	boatAcceleration[15] = req.BoatAcceleration_15
	boatAcceleration[16] = req.BoatAcceleration_16
	boatAcceleration[17] = req.BoatAcceleration_17
	boatAcceleration[18] = req.BoatAcceleration_18
	boatAcceleration[19] = req.BoatAcceleration_19
	boatAcceleration[20] = req.BoatAcceleration_20
	boatAcceleration[21] = req.BoatAcceleration_21
	boatAcceleration[22] = req.BoatAcceleration_22
	boatAcceleration[23] = req.BoatAcceleration_23
	boatAcceleration[24] = req.BoatAcceleration_24
	boatAcceleration[25] = req.BoatAcceleration_25
	boatAcceleration[26] = req.BoatAcceleration_26
	boatAcceleration[27] = req.BoatAcceleration_27
	boatAcceleration[28] = req.BoatAcceleration_28
	boatAcceleration[29] = req.BoatAcceleration_29
	boatAcceleration[30] = req.BoatAcceleration_30
	boatAcceleration[31] = req.BoatAcceleration_31
	boatAcceleration[32] = req.BoatAcceleration_32
	boatAcceleration[33] = req.BoatAcceleration_33
	boatAcceleration[34] = req.BoatAcceleration_34
	boatAcceleration[35] = req.BoatAcceleration_35
	boatAcceleration[36] = req.BoatAcceleration_36
	boatAcceleration[37] = req.BoatAcceleration_37
	boatAcceleration[38] = req.BoatAcceleration_38
	boatAcceleration[39] = req.BoatAcceleration_39
	boatAcceleration[40] = req.BoatAcceleration_40
	boatAcceleration[41] = req.BoatAcceleration_41
	boatAcceleration[42] = req.BoatAcceleration_42
	boatAcceleration[43] = req.BoatAcceleration_43
	boatAcceleration[44] = req.BoatAcceleration_44
	boatAcceleration[45] = req.BoatAcceleration_45
	boatAcceleration[46] = req.BoatAcceleration_46
	boatAcceleration[47] = req.BoatAcceleration_47
	boatAcceleration[48] = req.BoatAcceleration_48
	boatAcceleration[49] = req.BoatAcceleration_49
	boatAcceleration[50] = req.BoatAcceleration_50

	velocityRel[0] = req.VelocityRel_0
	velocityRel[1] = req.VelocityRel_1
	velocityRel[2] = req.VelocityRel_2
	velocityRel[3] = req.VelocityRel_3
	velocityRel[4] = req.VelocityRel_4
	velocityRel[5] = req.VelocityRel_5
	velocityRel[6] = req.VelocityRel_6
	velocityRel[7] = req.VelocityRel_7
	velocityRel[8] = req.VelocityRel_8
	velocityRel[9] = req.VelocityRel_9
	velocityRel[10] = req.VelocityRel_10
	velocityRel[11] = req.VelocityRel_11
	velocityRel[12] = req.VelocityRel_12
	velocityRel[13] = req.VelocityRel_13
	velocityRel[14] = req.VelocityRel_14
	velocityRel[15] = req.VelocityRel_15
	velocityRel[16] = req.VelocityRel_16
	velocityRel[17] = req.VelocityRel_17
	velocityRel[18] = req.VelocityRel_18
	velocityRel[19] = req.VelocityRel_19
	velocityRel[20] = req.VelocityRel_20
	velocityRel[21] = req.VelocityRel_21
	velocityRel[22] = req.VelocityRel_22
	velocityRel[23] = req.VelocityRel_23
	velocityRel[24] = req.VelocityRel_24
	velocityRel[25] = req.VelocityRel_25
	velocityRel[26] = req.VelocityRel_26
	velocityRel[27] = req.VelocityRel_27
	velocityRel[28] = req.VelocityRel_28
	velocityRel[29] = req.VelocityRel_29
	velocityRel[30] = req.VelocityRel_30
	velocityRel[31] = req.VelocityRel_31
	velocityRel[32] = req.VelocityRel_32
	velocityRel[33] = req.VelocityRel_33
	velocityRel[34] = req.VelocityRel_34
	velocityRel[35] = req.VelocityRel_35
	velocityRel[36] = req.VelocityRel_36
	velocityRel[37] = req.VelocityRel_37
	velocityRel[38] = req.VelocityRel_38
	velocityRel[39] = req.VelocityRel_39
	velocityRel[40] = req.VelocityRel_40
	velocityRel[41] = req.VelocityRel_41
	velocityRel[42] = req.VelocityRel_42
	velocityRel[43] = req.VelocityRel_43
	velocityRel[44] = req.VelocityRel_44
	velocityRel[45] = req.VelocityRel_45
	velocityRel[46] = req.VelocityRel_46
	velocityRel[47] = req.VelocityRel_47
	velocityRel[48] = req.VelocityRel_48
	velocityRel[49] = req.VelocityRel_49
	velocityRel[50] = req.VelocityRel_50

	sampleMetricsModel := pb.SampleMetricsModel{
		Id:                           uint32(req.Id),
		AthleteTrainingId:            uint32(req.AthleteTrainingId),
		DataSample:                   req.DataSample,
		StrokeRate:                   req.StrokeRate,
		DriveTime:                    req.DriveTime,
		Rhythm:                       req.Rhythm,
		CatchAngle:                   req.CatchAngle,
		FinishAngle:                  req.FinishAngle,
		TotalAngle:                   req.TotalAngle,
		CatchSlip:                    req.CatchSlip,
		ReleaseSlip:                  req.ReleaseSlip,
		EffectiveAnglePercent:        req.EffectiveAnglePercent,
		MaxBladeDepth:                req.MaxBladeDepth,
		BladeEfficiency:              req.BladeEfficiency,
		RowingPower:                  req.RowingPower,
		WorkPerStroke:                req.WorkPerStroke,
		RelativeWpS:                  req.RelativeWpS,
		EffectiveAngleDegree:         req.EffectiveAngleDegree,
		TargetAngle:                  req.TargetAngle,
		TargetForce:                  req.TargetForce,
		TargetWpS:                    req.TargetWpS,
		AngleDivTarget:               req.AngleDivTarget,
		ForceDivTarget:               req.ForceDivTarget,
		WpSDivTarget:                 req.WpSDivTarget,
		AverageVelocity:              req.AverageVelocity,
		BladeSpecificImpulse:         req.BladeSpecificImpulse,
		TimeOver_2000M:               req.TimeOver2000m,
		MaxForce:                     req.MaxForce,
		AverageForce:                 req.AverageForce,
		RatioAverDivMaxForce:         req.RatioAverDivMaxForce,
		PositionOfPeakForce:          req.PositionOfPeakForce,
		CatchForceGradient:           req.CatchForceGradient,
		FinishForceGradient:          req.FinishForceGradient,
		MaxHandleVelocity:            req.MaxHandleVelocity,
		Hdf:                          req.HDF,
		LegsDrive:                    req.LegsDrive,
		LegsMaxSpeed:                 req.LegsMaxSpeed,
		CatchFactor:                  req.CatchFactor,
		RowingStyleFactor:            req.RowingStyleFactor,
		ReleaseWash:                  req.ReleaseWash,
		AverForceDivWeight:           req.AverForceDivWeight,
		VseatAtCatch:                 req.VseatAtCatch,
		HandleTravelAtEntryForce:     req.HandleTravelAtEntryForce,
		HandleTravelAt_70PerForce:    req.HandleTravelAt70PerForce,
		HandleTravelAt_0As:           req.HandleTravelAt0As,
		SeatTravelAtEntryForce:       req.SeatTravelAtEntryForce,
		SeatTravelAt_70PerForce:      req.SeatTravelAt70PerForce,
		SeatTravelAt_0As:             req.SeatTravelAt0As,
		DTravelAtEntryForcePercent:   req.DTravelAtEntryForcePercent,
		DTravelAt_70PerForcePercent:  req.DTravelAt70PerForcePercent,
		DTravelAt_0AsPercent:         req.DTravelAt0AsPercent,
		DTravelAtEntryForceDistance:  req.DTravelAtEntryForceDistance,
		DTravelAt_70PerForceDistance: req.DTravelAt70PerForceDistance,
		DTravelAt_0AsDistance:        req.DTravelAt0AsDistance,
		SeatOnRecovery:               req.SeatOnRecovery,
		VertAtCatch:                  req.VertAtCatch,
		EntryForce:                   req.EntryForce,
		ForceUpto_70Per:              req.ForceUpto70Per,
		MaxVseat:                     req.MaxVseat,
		PeakForce:                    req.PeakForce,
		ForceFrom_70Per:              req.ForceFrom70Per,
		VertAtFinish:                 req.VertAtFinish,
		ForceAtFinish:                req.ForceAtFinish,
		AverageBoatSpeed:             req.AverageBoatSpeed,
		MinimalBoatSpeed:             req.MinimalBoatSpeed,
		MaximalBoatSpeed:             req.MaximalBoatSpeed,
		DistancePerStroke:            req.DistancePerStroke,
		DragFactor:                   req.DragFactor,
		WindForwardCompRelWater:      req.WindForwardCompRelWater,
		WindDirectionRelWater:        req.WindDirectionRelWater,
		Time_250M:                    req.Time250m,
		BoatSpeedEfficiency:          req.BoatSpeedEfficiency,
		TimeAtWaterTemp_25Deg:        req.TimeAtWaterTemp25Deg,
		BoatSpeedVariation:           req.BoatSpeedVariation,
		WindSpeedRelBoat:             req.WindSpeedRelBoat,
		WindDirectionRelBoat:         req.WindDirectionRelBoat,
		DragFactorF:                  req.DragFactorF,
		DragFactorPprop:              req.DragFactorPprop,
		DragFactorPold:               req.DragFactorPold,
		DragFactorPtot:               req.DragFactorPtot,
		AccelerationMinimun:          req.AccelerationMinimun,
		AccelerationMaximum:          req.AccelerationMaximum,
		ModelSpeed:                   req.ModelSpeed,
		EffectiveWorkPerStroke:       req.EffectiveWorkPerStroke,
		ModelDps:                     req.ModelDPS,
		PropulsivePower:              req.PropulsivePower,
		DriveMaximalAt:               req.DriveMaximalAt,
		FirstPeak:                    req.FirstPeak,
		ZeroBeforeCatch:              req.ZeroBeforeCatch,
		MinimalFromCatch:             req.MinimalFromCatch,
		ZeroAfterCatch:               req.ZeroAfterCatch,
		StdDeviation:                 req.StdDeviation,
		OarAngle:                     oarAngle[:],
		HandleForce:                  handleForce[:],
		VerticalAngleBoatRoll:        verticalAngleBoatRoll[:],
		LegsVelocity:                 legsVelocity[:],
		HandleSpeed:                  handleSpeed[:],
		HdfFig:                       hdfFig[:],
		BladeDf:                      bladeDf[:],
		Velocity:                     velocity[:],
		BoatAcceleration:             boatAcceleration[:],
		VelocityRel:                  velocityRel[:],
	}

	return &sampleMetricsModel
}
