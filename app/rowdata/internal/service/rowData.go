package service

import (
	"context"
	"sync"

	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata/internal/repository/db/dao"
	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/rowdata"
	"github.com/zizhuxuexi/datacenter-rowing/pkg/e"
)

var RowDataSrvIns *RowDataSrv
var RowDataSrvOnce sync.Once

type RowDataSrv struct {
	pb.UnimplementedRowDataServiceServer
}

func GetRowDataSrv() *RowDataSrv {
	RowDataSrvOnce.Do(func() {
		RowDataSrvIns = &RowDataSrv{}
	})
	return RowDataSrvIns
}

func (rd *RowDataSrv) TrainingSummaryAdd(ctx context.Context, req *pb.TrainingSummaryAddRequest) (resp *pb.TrainingSummaryAddResponse, err error) {
	resp = new(pb.TrainingSummaryAddResponse)
	resp.Code = e.SUCCESS
	trainingId, err := dao.NewRowDataDao(ctx).AddTrainingSummary(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.Data = e.GetMsg(int(resp.Code))
	resp.TrainingId = uint32(trainingId)
	return
}

func (rd *RowDataSrv) TrainingSummaryGetByTrainingName(ctx context.Context, req *pb.TrainingSummaryGetByTrainingNameRequest) (resp *pb.TrainingSummaryResponse, err error) {
	resp = new(pb.TrainingSummaryResponse)
	resp.Code = e.SUCCESS
	r, err := dao.NewRowDataDao(ctx).GetTrainingSummaryByTrainingName(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	// resp.TrainingSummaryDetail = &pb.TrainingSummaryModel{
	// 	TrainingId:      uint32(r.TrainingId),
	// 	TrainingName:    r.TrainingName,
	// 	TrainDate:       timestamppb.New(r.TrainDate),
	// 	EventGender:     r.EventGender,
	// 	EventPeopleType: r.EventPeopleType,
	// 	EventScale:      r.EventScale,
	// 	Event:           r.Event,
	// 	Weather:         r.Weather,
	// 	Temp:            int32(r.Temp),
	// 	WindDir:         r.WindDir,
	// 	Loc:             r.Loc,
	// 	Coach:           r.Coach,
	// 	SampleCount:     int32(r.SampleCount),
	// 	Remark:          r.Remark,
	// }
	resp.TrainingSummaryDetail = dao.BuildTrainingSummary(*r)
	return
}

func (rd *RowDataSrv) TrainingSummaryGet(ctx context.Context, req *pb.TrainingSummaryGetRequset) (resp *pb.TrainingSummaryGetResponse, err error) {
	resp = new(pb.TrainingSummaryGetResponse)
	resp.Code = e.SUCCESS
	r, err := dao.NewRowDataDao(ctx).GetTrainingSummary(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	var tss []*pb.TrainingSummaryModel
	for _, ts := range r {
		tss = append(tss, dao.BuildTrainingSummary(*ts))
	}
	resp.TrainingSummary = tss
	return
}

func (rd *RowDataSrv) AthleteTrainingDataAdd(ctx context.Context, req *pb.AthleteTrainingDataAddRequest) (resp *pb.AthleteTrainingDataAddResponse, err error) {
	resp = new(pb.AthleteTrainingDataAddResponse)
	resp.Code = e.SUCCESS
	athleteTrainingId, err := dao.NewRowDataDao(ctx).AddAthleteTrainingData(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.Data = e.GetMsg(int(resp.Code))
	resp.AthleteTrainingId = uint32(athleteTrainingId)
	return
}

func (rd *RowDataSrv) AthleteTrainingDataGetByName(ctx context.Context, req *pb.AthleteTrainingDataGetByName) (resp *pb.AthleteTrainingDataResponse, err error) {
	resp = new(pb.AthleteTrainingDataResponse)
	resp.Code = e.SUCCESS
	r, err := dao.NewRowDataDao(ctx).GetAthleteTrainingDataByName(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	for _, rowdata := range *r {
		resp.AthleteTrainingDataDetail = append(resp.AthleteTrainingDataDetail, dao.BuildAthleteTrainingData(*&rowdata))
	}
	return
}

func (rd *RowDataSrv) AthleteTrainingDataGet(ctx context.Context, req *pb.AthleteTrainingDataGetRequest) (resp *pb.AthleteTrainingDataGetResponse, err error) {
	resp = new(pb.AthleteTrainingDataGetResponse)
	resp.Code = e.SUCCESS
	r, err := dao.NewRowDataDao(ctx).GetAthleteTrainingData(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	var atds []*pb.AthleteTrainingDataModel
	for _, atd := range r {
		atds = append(atds, dao.BuildAthleteTrainingData(*atd))
	}
	resp.AthleteTrainingData = atds
	return
}

func (rd *RowDataSrv) SampleMetricsAdd(ctx context.Context, req *pb.SampleMetricsModel) (resp *pb.SampleMetricsAddResponse, err error) {
	resp = new(pb.SampleMetricsAddResponse)
	resp.Code = e.SUCCESS
	id, err := dao.NewRowDataDao(ctx).AddSampleMetrics(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.Data = e.GetMsg(int(resp.Code))
	resp.Id = uint32(id)
	return
}

func (rd *RowDataSrv) SampleMetricsGetByAthleteTrainingId(ctx context.Context, req *pb.SampleMetricsGetByAthleteTrainingIdRequest) (resp *pb.SampleMetricsGetByAthleteTrainingIdResponse, err error) {
	resp = new(pb.SampleMetricsGetByAthleteTrainingIdResponse)
	resp.Code = e.SUCCESS
	r, err := dao.NewRowDataDao(ctx).GetSampleMetricsByAthleteTrainingId(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	var sms []*pb.SampleMetricsModel
	for _, sm := range r {
		sms = append(sms, dao.BuildSampleMetrics(*sm))
	}
	resp.SampleMetrics = sms
	return
}
