package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata-api/internal/model"
	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/rowdata"
	"github.com/zizhuxuexi/datacenter-rowing/pkg/res"
)

// 创建 gRPC 客户端连接
func NewGRPCClient(addr string) (pb.RowDataServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
		return nil, nil, err
	}
	client := pb.NewRowDataServiceClient(conn)
	return client, conn, nil
}

// 上传TrainingSummary的数据
func AddTrainingSummary(client pb.RowDataServiceClient, trainingSummary model.TrainingSummary) res.CommonResponse {
	req := &pb.TrainingSummaryAddRequest{
		TrainingName:    trainingSummary.TrainingName,
		TrainDate:       timestamppb.New(trainingSummary.TrainDate),
		EventGender:     trainingSummary.EventGender,
		EventPeopleType: trainingSummary.EventPeopleType,
		EventScale:      trainingSummary.EventScale,
		Event:           trainingSummary.Event,
		Weather:         trainingSummary.Weather,
		Temp:            int32(trainingSummary.Temp),
		WindDir:         trainingSummary.WindDir,
		Loc:             trainingSummary.Loc,
		Coach:           trainingSummary.Coach,
		SampleCount:     int32(trainingSummary.SampleCount),
		Remark:          trainingSummary.Remark,
	}

	//这里上下文线用空上下文
	resp, err := client.TrainingSummaryAdd(context.Background(), req)
	if err != nil {
		panic(err)
	}

	return res.CommonResponse{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: struct {
			TrainingId uint32
		}{TrainingId: resp.TrainingId},
	}
}

func GetTrainingSummaryByTrainingName(client pb.RowDataServiceClient, trainingName string) res.CommonResponse {
	req := &pb.TrainingSummaryGetByTrainingNameRequest{TrainingName: trainingName}
	resp, err := client.TrainingSummaryGetByTrainingName(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return res.CommonResponse{
		Code: resp.Code,
		Msg:  "训练数据",
		Data: resp.TrainingSummaryDetail,
	}
}

func AddAthleteTrainingData(client pb.RowDataServiceClient, athleteTrainingData model.AthleteTrainingData) res.CommonResponse {
	req := &pb.AthleteTrainingDataAddRequest{
		TrainingId:     uint32(athleteTrainingData.TrainingId),
		Name:           athleteTrainingData.Name,
		Gender:         athleteTrainingData.Gender,
		Seat:           int32(athleteTrainingData.Seat),
		Side:           int32(athleteTrainingData.Side),
		Height:         athleteTrainingData.Height,
		Weight:         athleteTrainingData.Weight,
		OarInboard:     athleteTrainingData.OarInboard,
		OarLength:      athleteTrainingData.OarLength,
		OarBladeLength: athleteTrainingData.OarBladeLength,
	}

	resp, err := client.AthleteTrainingDataAdd(context.Background(), req)
	if err != nil {
		panic(err)
	}

	return res.CommonResponse{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: struct {
			AthleteTrainingId uint32
		}{AthleteTrainingId: resp.AthleteTrainingId},
	}
}

func GetAthleteTrainingDataByName(client pb.RowDataServiceClient, name string) res.CommonResponse {
	req := &pb.AthleteTrainingDataGetByName{Name: name}
	resp, err := client.AthleteTrainingDataGetByName(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return res.CommonResponse{
		Code: resp.Code,
		Msg:  "运动员训练数据",
		Data: resp.AthleteTrainingDataDetail,
	}
}

func AddSampleMetrics(client pb.RowDataServiceClient, sampleMetrics model.SampleMetrics) res.CommonResponse {
	req := &pb.SampleMetricsModel{
		Id:                           uint32(sampleMetrics.Id),
		AthleteTrainingId:            uint32(sampleMetrics.AthleteTrainingId),
		DataSample:                   sampleMetrics.DataSample,
		StrokeRate:                   sampleMetrics.StrokeRate,
		DriveTime:                    sampleMetrics.DriveTime,
		Rhythm:                       sampleMetrics.Rhythm,
		CatchAngle:                   sampleMetrics.CatchAngle,
		FinishAngle:                  sampleMetrics.FinishAngle,
		TotalAngle:                   sampleMetrics.TotalAngle,
		CatchSlip:                    sampleMetrics.CatchSlip,
		ReleaseSlip:                  sampleMetrics.ReleaseSlip,
		EffectiveAnglePercent:        sampleMetrics.EffectiveAnglePercent,
		MaxBladeDepth:                sampleMetrics.MaxBladeDepth,
		BladeEfficiency:              sampleMetrics.BladeEfficiency,
		RowingPower:                  sampleMetrics.RowingPower,
		WorkPerStroke:                sampleMetrics.WorkPerStroke,
		RelativeWpS:                  sampleMetrics.RelativeWpS,
		EffectiveAngleDegree:         sampleMetrics.EffectiveAngleDegree,
		TargetAngle:                  sampleMetrics.TargetAngle,
		TargetForce:                  sampleMetrics.TargetForce,
		TargetWpS:                    sampleMetrics.TargetWpS,
		AngleDivTarget:               sampleMetrics.AngleDivTarget,
		ForceDivTarget:               sampleMetrics.ForceDivTarget,
		WpSDivTarget:                 sampleMetrics.WpSDivTarget,
		AverageVelocity:              sampleMetrics.AverageVelocity,
		BladeSpecificImpulse:         sampleMetrics.BladeSpecificImpulse,
		TimeOver_2000M:               sampleMetrics.TimeOver2000m,
		MaxForce:                     sampleMetrics.MaxForce,
		AverageForce:                 sampleMetrics.AverageForce,
		RatioAverDivMaxForce:         sampleMetrics.RatioAverDivMaxForce,
		PositionOfPeakForce:          sampleMetrics.PositionOfPeakForce,
		CatchForceGradient:           sampleMetrics.CatchForceGradient,
		FinishForceGradient:          sampleMetrics.FinishForceGradient,
		MaxHandleVelocity:            sampleMetrics.MaxHandleVelocity,
		Hdf:                          sampleMetrics.HDF,
		LegsDrive:                    sampleMetrics.LegsDrive,
		LegsMaxSpeed:                 sampleMetrics.LegsMaxSpeed,
		CatchFactor:                  sampleMetrics.CatchFactor,
		RowingStyleFactor:            sampleMetrics.RowingStyleFactor,
		ReleaseWash:                  sampleMetrics.ReleaseWash,
		AverForceDivWeight:           sampleMetrics.AverForceDivWeight,
		VseatAtCatch:                 sampleMetrics.VseatAtCatch,
		HandleTravelAtEntryForce:     sampleMetrics.HandleTravelAtEntryForce,
		HandleTravelAt_70PerForce:    sampleMetrics.HandleTravelAt70PerForce,
		HandleTravelAt_0As:           sampleMetrics.HandleTravelAt0As,
		SeatTravelAtEntryForce:       sampleMetrics.SeatTravelAtEntryForce,
		SeatTravelAt_70PerForce:      sampleMetrics.SeatTravelAt70PerForce,
		SeatTravelAt_0As:             sampleMetrics.SeatTravelAt0As,
		DTravelAtEntryForcePercent:   sampleMetrics.DTravelAtEntryForcePercent,
		DTravelAt_70PerForcePercent:  sampleMetrics.DTravelAt70PerForcePercent,
		DTravelAt_0AsPercent:         sampleMetrics.DTravelAt0AsPercent,
		DTravelAtEntryForceDistance:  sampleMetrics.DTravelAtEntryForceDistance,
		DTravelAt_70PerForceDistance: sampleMetrics.DTravelAt70PerForceDistance,
		DTravelAt_0AsDistance:        sampleMetrics.DTravelAt0AsDistance,
		SeatOnRecovery:               sampleMetrics.SeatOnRecovery,
		VertAtCatch:                  sampleMetrics.VertAtCatch,
		EntryForce:                   sampleMetrics.EntryForce,
		ForceUpto_70Per:              sampleMetrics.ForceUpto70Per,
		MaxVseat:                     sampleMetrics.MaxVseat,
		PeakForce:                    sampleMetrics.PeakForce,
		ForceFrom_70Per:              sampleMetrics.ForceFrom70Per,
		VertAtFinish:                 sampleMetrics.VertAtFinish,
		ForceAtFinish:                sampleMetrics.ForceAtFinish,
		AverageBoatSpeed:             sampleMetrics.AverageBoatSpeed,
		MinimalBoatSpeed:             sampleMetrics.MinimalBoatSpeed,
		MaximalBoatSpeed:             sampleMetrics.MaximalBoatSpeed,
		DistancePerStroke:            sampleMetrics.DistancePerStroke,
		DragFactor:                   sampleMetrics.DragFactor,
		WindForwardCompRelWater:      sampleMetrics.WindForwardCompRelWater,
		WindDirectionRelWater:        sampleMetrics.WindDirectionRelWater,
		Time_250M:                    sampleMetrics.Time250m,
		BoatSpeedEfficiency:          sampleMetrics.BoatSpeedEfficiency,
		TimeAtWaterTemp_25Deg:        sampleMetrics.TimeAtWaterTemp25Deg,
		BoatSpeedVariation:           sampleMetrics.BoatSpeedVariation,
		WindSpeedRelBoat:             sampleMetrics.WindSpeedRelBoat,
		WindDirectionRelBoat:         sampleMetrics.WindDirectionRelBoat,
		DragFactorF:                  sampleMetrics.DragFactorF,
		DragFactorPprop:              sampleMetrics.DragFactorPprop,
		DragFactorPold:               sampleMetrics.DragFactorPold,
		DragFactorPtot:               sampleMetrics.DragFactorPtot,
		AccelerationMinimun:          sampleMetrics.AccelerationMinimun,
		AccelerationMaximum:          sampleMetrics.AccelerationMaximum,
		ModelSpeed:                   sampleMetrics.ModelSpeed,
		EffectiveWorkPerStroke:       sampleMetrics.EffectiveWorkPerStroke,
		ModelDps:                     sampleMetrics.ModelDPS,
		PropulsivePower:              sampleMetrics.PropulsivePower,
		DriveMaximalAt:               sampleMetrics.DriveMaximalAt,
		FirstPeak:                    sampleMetrics.FirstPeak,
		ZeroBeforeCatch:              sampleMetrics.ZeroBeforeCatch,
		MinimalFromCatch:             sampleMetrics.MinimalFromCatch,
		ZeroAfterCatch:               sampleMetrics.ZeroAfterCatch,
		StdDeviation:                 sampleMetrics.StdDeviation,
		OarAngle:                     sampleMetrics.OarAngle[:],
		HandleForce:                  sampleMetrics.HandleForce[:],
		VerticalAngleBoatRoll:        sampleMetrics.VerticalAngleBoatRoll[:],
		LegsVelocity:                 sampleMetrics.LegsVelocity[:],
		HandleSpeed:                  sampleMetrics.HandleSpeed[:],
		HdfFig:                       sampleMetrics.HDFFig[:],
		BladeDf:                      sampleMetrics.BladeDF[:],
		Velocity:                     sampleMetrics.Velocity[:],
		BoatAcceleration:             sampleMetrics.BoatAcceleration[:],
		VelocityRel:                  sampleMetrics.VelocityRel[:],
	}

	resp, err := client.SampleMetricsAdd(context.Background(), req)
	if err != nil {
		panic(err)
	}

	return res.CommonResponse{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: struct {
			Id uint32
		}{Id: resp.Id},
	}

}

func GetSampleMetricsByAthleteTrainingId(client pb.RowDataServiceClient, athleteTrainingId uint32) res.CommonResponse {
	req := &pb.SampleMetricsGetByAthleteTrainingIdRequest{AthleteTrainingId: athleteTrainingId}
	resp, err := client.SampleMetricsGetByAthleteTrainingId(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return res.CommonResponse{
		Code: resp.Code,
		Msg:  "查询🈯️定athleteTrainingId的详细指标",
		Data: resp.SampleMetrics,
	}
}
