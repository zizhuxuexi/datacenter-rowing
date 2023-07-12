package service

import (
	"context"
	"sync"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata/internal/repository/db/dao"
	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/trainingSummary"
	"github.com/zizhuxuexi/datacenter-rowing/pkg/e"
)

var TrainingSummarySrvIns *TrainingSummarySrv
var TrainingSummarySrvOnce sync.Once

type TrainingSummarySrv struct {
	pb.UnimplementedTrainingSummaryServiceServer
}

func GetTrainingSummarySrv() *TrainingSummarySrv {
	TrainingSummarySrvOnce.Do(func() {
		TrainingSummarySrvIns = &TrainingSummarySrv{}
	})
	return TrainingSummarySrvIns
}

func (ts *TrainingSummarySrv) TrainingSummaryAdd(ctx context.Context, req *pb.TrainingSummaryAddRequest) (resp *pb.TrainingSummaryCommonResponse, err error) {
	resp = new(pb.TrainingSummaryCommonResponse)
	resp.Code = e.SUCCESS
	err = dao.NewTrainingSummaryDao(ctx).AddTrainingSummary(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.Data = e.GetMsg(int(resp.Code))
	return
}

func (ts *TrainingSummarySrv) TrainingSummaryGetByTrainingName(ctx context.Context, req *pb.TrainingSummaryGetByTrainingNameRequest) (resp *pb.TrainingSummaryResponse, err error) {
	resp = new(pb.TrainingSummaryResponse)
	resp.Code = e.SUCCESS
	r, err := dao.NewTrainingSummaryDao(ctx).GetTrainingSummaryByTrainingName(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.TrainingSummaryDetail = &pb.TrainingSummaryModel{
		TrainingId:      uint32(r.TrainingId),
		TrainingName:    r.TrainingName,
		TrainDate:       timestamppb.New(r.TrainDate),
		EventGender:     r.EventGender,
		EventPeopleType: r.EventPeopleType,
		EventScale:      r.EventScale,
		Event:           r.Event,
		Weather:         r.Weather,
		Temp:            int32(r.Temp),
		WindDir:         r.WindDir,
		Loc:             r.Loc,
		Coach:           r.Coach,
		SampleCount:     int32(r.SampleCount),
		Remark:          r.Remark,
	}
	return
}
