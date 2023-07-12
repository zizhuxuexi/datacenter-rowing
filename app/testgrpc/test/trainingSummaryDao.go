package dao

import (
	"context"
	"errors"

	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata/internal/repository/db/model"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/trainingSummary"
	"github.com/zizhuxuexi/datacenter-rowing/pkg/util/logger"
)

type TrainingSummaryDao struct {
	*gorm.DB
}

func NewTrainingSummaryDao(ctx context.Context) *TrainingSummaryDao {
	return &TrainingSummaryDao{NewDBClient(ctx)}
}

// 上传训练记录
func (dao *TrainingSummaryDao) AddTrainingSummary(req *pb.TrainingSummaryAddRequest) error {
	//if exist := dao.CheckExist(req); exist {
	//	return errors.New("训练记录已存在，请勿重复上传。")
	//}
	var count int64
	var trainingSummary model.TrainingSummary
	dao.Model(&model.TrainingSummary{}).Where("training_name=?", req.TrainingName).Count(&count)
	if count != 0 {
		return errors.New("训练记录已存在，请勿重复上传。")
	}
	trainingSummary = model.TrainingSummary{
		TrainingName:    req.TrainingName,
		TrainDate:       req.TrainDate.AsTime(),
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
		return err
	}
	return nil
}

func (dao *TrainingSummaryDao) GetTrainingSummaryByTrainingName(req *pb.TrainingSummaryGetByTrainingNameRequest) (ts *model.TrainingSummary, err error) {
	err = dao.Model(&model.TrainingSummary{}).Where("training_name=?", req.TrainingName).First(&ts).Error
	return
}

func (dao *TrainingSummaryDao) CheckExist(req *pb.TrainingSummaryAddRequest) bool {
	var trainingSummary model.TrainingSummary
	if err := dao.Model(&model.TrainingSummary{}).Where("training_name=?", req.TrainingName).First(&trainingSummary).Error; err == gorm.ErrRecordNotFound {
		return false
	} else {
		return true
	}
}

func BuildTrainingSummarytest(item model.TrainingSummary) *pb.TrainingSummaryModel {
	trainingSummaryModel := pb.TrainingSummaryModel{
		TrainingId:      uint32(item.TrainingId),
		TrainingName:    item.TrainingName,
		TrainDate:       timestamppb.New(item.TrainDate),
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
