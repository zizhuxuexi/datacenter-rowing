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
func AddTrainingSummary(client pb.RowDataServiceClient, trainingSummary model.TrainingSummary) (res.CommonResponse, uint32) {
	req := &pb.TrainingSummaryAddRequest{
		TrainingName:    trainingSummary.TrainingName,
		TrainDate:       timestamppb.New(trainingSummary.TrainingDate),
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
	}, resp.TrainingId
}

func GetTrainingSummary(client pb.RowDataServiceClient, trainingSummaryGet model.TrainingSummaryGet) (res.CommonResponse, *[]model.TrainingSummary) {
	trainingSummary := BuildTrainingSummary(trainingSummaryGet.TrainingSummary)
	req := &pb.TrainingSummaryGetRequset{
		TrainingSummary:    trainingSummary,
		Set:                trainingSummaryGet.Set,
		SetTrainingName:    trainingSummaryGet.SetTrainingName,
		SetTrainingDate:    trainingSummaryGet.SetTrainingDate,
		SetEventGender:     trainingSummaryGet.SetEventGender,
		SetEventPeopleType: trainingSummaryGet.SetEventPeopleType,
		SetEventScale:      trainingSummaryGet.SetEventScale,
		SetEvent:           trainingSummaryGet.SetEvent,
		SetWeather:         trainingSummaryGet.SetWeather,
		SetTemp:            trainingSummaryGet.SetTemp,
		SetWindDir:         trainingSummaryGet.SetWindDir,
		SetLoc:             trainingSummaryGet.SetLoc,
		SetCoach:           trainingSummaryGet.SetCoach,
		SetSampleCount:     trainingSummaryGet.SetSampleCount,
		SetRemark:          trainingSummaryGet.SetRemark,
	}
	resp, err := client.TrainingSummaryGet(context.Background(), req)
	if err != nil {
		panic(err)
	}
	var trainingSummaryArr []model.TrainingSummary
	for _, ts := range resp.TrainingSummary {
		tsm := model.TrainingSummary{
			TrainingId:      uint(ts.TrainingId),
			TrainingName:    ts.TrainingName,
			TrainingDate:    ts.TrainingDate.AsTime(),
			EventGender:     ts.EventGender,
			EventPeopleType: ts.EventPeopleType,
			EventScale:      ts.EventScale,
			Event:           ts.Event,
			Weather:         ts.Weather,
			Temp:            int(ts.Temp),
			WindDir:         ts.WindDir,
			Loc:             ts.Loc,
			Coach:           ts.Coach,
			SampleCount:     int(ts.SampleCount),
			Remark:          ts.Remark,
		}
		trainingSummaryArr = append(trainingSummaryArr, tsm)
	}

	return res.CommonResponse{
		Code: resp.Code,
		Msg:  "查询后的训练数据",
	}, &trainingSummaryArr
}

func AddAthleteTrainingData(client pb.RowDataServiceClient, athleteTrainingData model.AthleteTrainingData) (res.CommonResponse, uint32) {
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
	}, resp.AthleteTrainingId
}

func GetAthleteTrainingData(client pb.RowDataServiceClient, atd model.AthleteTrainingDataGet) (res.CommonResponse, *[]model.AthleteTrainingData) {
	athleteTrainingData := BuildAthleteTrainingData(atd.AthleteTrainingData)
	req := &pb.AthleteTrainingDataGetRequest{
		AthleteTrainingData:  athleteTrainingData,
		Set:                  atd.Set,
		SetAthleteTrainingId: atd.SetAthleteTrainingId,
		SetTrainingId:        atd.SetTrainingId,
		SetName:              atd.SetName,
		SetGender:            atd.SetGender,
		SetSeat:              atd.SetSeat,
		SetSide:              atd.SetSide,
		SetHeight:            atd.SetHeight,
		SetWeight:            atd.SetWeight,
		SetOarInboard:        atd.SetOarInboard,
		SetOarLength:         atd.SetOarLength,
		SetOarBladeLength:    atd.SetOarBladeLength,
	}
	resp, err := client.AthleteTrainingDataGet(context.Background(), req)
	if err != nil {
		panic(err)
	}

	var athleteTrainingDataArr []model.AthleteTrainingData
	for _, atd := range resp.AthleteTrainingData {
		atdm := model.AthleteTrainingData{
			AthleteTrainingId: uint(atd.AthleteTrainingId),
			TrainingId:        uint(atd.TrainingId),
			Name:              atd.Name,
			Gender:            atd.Gender,
			Seat:              int(atd.Seat),
			Side:              int(atd.Side),
			Height:            atd.Height,
			Weight:            atd.Weight,
			OarInboard:        atd.OarInboard,
			OarLength:         atd.OarLength,
			OarBladeLength:    atd.OarBladeLength,
		}
		athleteTrainingDataArr = append(athleteTrainingDataArr, atdm)

	}
	return res.CommonResponse{
		Code: resp.Code,
		Msg:  "运动员训练数据",
	}, &athleteTrainingDataArr
}

func AddSampleMetrics(client pb.RowDataServiceClient, sampleMetrics model.SampleMetrics) (res.CommonResponse, uint32) {
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
	}, resp.Id

}

func GetSampleMetricsByAthleteTrainingId(client pb.RowDataServiceClient, athleteTrainingId uint32, set bool) (res.CommonResponse, *[]model.SampleMetrics) {
	req := &pb.SampleMetricsGetByAthleteTrainingIdRequest{AthleteTrainingId: athleteTrainingId, Set: set}
	resp, err := client.SampleMetricsGetByAthleteTrainingId(context.Background(), req)
	if err != nil {
		panic(err)
	}
	var sampleMetricsArr []model.SampleMetrics
	for _, sm := range resp.SampleMetrics {
		oarAngle := [51]float64{}
		copy(oarAngle[:], sm.OarAngle[:51])
		handleForce := [51]float64{}
		copy(handleForce[:], sm.HandleForce[:51])
		verticalAngleBoatRoll := [51]float64{}
		copy(verticalAngleBoatRoll[:], sm.VerticalAngleBoatRoll[:51])
		legsVelocity := [51]float64{}
		copy(legsVelocity[:], sm.LegsVelocity[:51])
		handleSpeed := [51]float64{}
		copy(handleSpeed[:], sm.HandleSpeed[:51])
		hdfFig := [51]float64{}
		copy(hdfFig[:], sm.HdfFig[:51])
		bladeDF := [51]float64{}
		copy(bladeDF[:], sm.BladeDf[:51])
		velocity := [51]float64{}
		copy(velocity[:], sm.Velocity[:51])
		boatAcceleration := [51]float64{}
		copy(boatAcceleration[:], sm.BoatAcceleration[:51])
		velocityRel := [51]float64{}
		copy(velocityRel[:], sm.VelocityRel[:51])

		smm := model.SampleMetrics{
			Id:                          uint(sm.Id),
			AthleteTrainingId:           uint(sm.AthleteTrainingId),
			DataSample:                  sm.DataSample,
			StrokeRate:                  sm.StrokeRate,
			DriveTime:                   sm.DriveTime,
			Rhythm:                      sm.Rhythm,
			CatchAngle:                  sm.CatchAngle,
			FinishAngle:                 sm.FinishAngle,
			TotalAngle:                  sm.TotalAngle,
			CatchSlip:                   sm.CatchSlip,
			ReleaseSlip:                 sm.ReleaseSlip,
			EffectiveAnglePercent:       sm.EffectiveAnglePercent,
			MaxBladeDepth:               sm.MaxBladeDepth,
			BladeEfficiency:             sm.BladeEfficiency,
			RowingPower:                 sm.RowingPower,
			WorkPerStroke:               sm.WorkPerStroke,
			RelativeWpS:                 sm.RelativeWpS,
			EffectiveAngleDegree:        sm.EffectiveAngleDegree,
			TargetAngle:                 sm.TargetAngle,
			TargetForce:                 sm.TargetForce,
			TargetWpS:                   sm.TargetWpS,
			AngleDivTarget:              sm.AngleDivTarget,
			ForceDivTarget:              sm.ForceDivTarget,
			WpSDivTarget:                sm.WpSDivTarget,
			AverageVelocity:             sm.AverageVelocity,
			BladeSpecificImpulse:        sm.BladeSpecificImpulse,
			TimeOver2000m:               sm.TimeOver_2000M,
			MaxForce:                    sm.MaxForce,
			AverageForce:                sm.AverageForce,
			RatioAverDivMaxForce:        sm.RatioAverDivMaxForce,
			PositionOfPeakForce:         sm.PositionOfPeakForce,
			CatchForceGradient:          sm.CatchForceGradient,
			FinishForceGradient:         sm.FinishForceGradient,
			MaxHandleVelocity:           sm.MaxHandleVelocity,
			HDF:                         sm.Hdf,
			LegsDrive:                   sm.LegsDrive,
			LegsMaxSpeed:                sm.LegsMaxSpeed,
			CatchFactor:                 sm.CatchFactor,
			RowingStyleFactor:           sm.RowingStyleFactor,
			ReleaseWash:                 sm.ReleaseWash,
			AverForceDivWeight:          sm.AverForceDivWeight,
			VseatAtCatch:                sm.VseatAtCatch,
			HandleTravelAtEntryForce:    sm.HandleTravelAtEntryForce,
			HandleTravelAt70PerForce:    sm.HandleTravelAt_70PerForce,
			HandleTravelAt0As:           sm.HandleTravelAt_0As,
			SeatTravelAtEntryForce:      sm.SeatTravelAtEntryForce,
			SeatTravelAt70PerForce:      sm.SeatTravelAt_70PerForce,
			SeatTravelAt0As:             sm.SeatTravelAt_0As,
			DTravelAtEntryForcePercent:  sm.DTravelAtEntryForcePercent,
			DTravelAt70PerForcePercent:  sm.DTravelAt_70PerForcePercent,
			DTravelAt0AsPercent:         sm.DTravelAt_0AsPercent,
			DTravelAtEntryForceDistance: sm.DTravelAtEntryForceDistance,
			DTravelAt70PerForceDistance: sm.DTravelAt_70PerForceDistance,
			DTravelAt0AsDistance:        sm.DTravelAt_0AsDistance,
			SeatOnRecovery:              sm.SeatOnRecovery,
			VertAtCatch:                 sm.VertAtCatch,
			EntryForce:                  sm.EntryForce,
			ForceUpto70Per:              sm.ForceUpto_70Per,
			MaxVseat:                    sm.MaxVseat,
			PeakForce:                   sm.PeakForce,
			ForceFrom70Per:              sm.ForceFrom_70Per,
			VertAtFinish:                sm.VertAtFinish,
			ForceAtFinish:               sm.ForceAtFinish,
			AverageBoatSpeed:            sm.AverageBoatSpeed,
			MinimalBoatSpeed:            sm.MinimalBoatSpeed,
			MaximalBoatSpeed:            sm.MaximalBoatSpeed,
			DistancePerStroke:           sm.DistancePerStroke,
			DragFactor:                  sm.DragFactor,
			WindForwardCompRelWater:     sm.WindForwardCompRelWater,
			WindDirectionRelWater:       sm.WindDirectionRelWater,
			Time250m:                    sm.Time_250M,
			BoatSpeedEfficiency:         sm.BoatSpeedEfficiency,
			TimeAtWaterTemp25Deg:        sm.TimeAtWaterTemp_25Deg,
			BoatSpeedVariation:          sm.BoatSpeedVariation,
			WindSpeedRelBoat:            sm.WindSpeedRelBoat,
			WindDirectionRelBoat:        sm.WindDirectionRelBoat,
			DragFactorF:                 sm.DragFactorF,
			DragFactorPprop:             sm.DragFactorPprop,
			DragFactorPold:              sm.DragFactorPold,
			DragFactorPtot:              sm.DragFactorPtot,
			AccelerationMinimun:         sm.AccelerationMinimun,
			AccelerationMaximum:         sm.AccelerationMaximum,
			ModelSpeed:                  sm.ModelSpeed,
			EffectiveWorkPerStroke:      sm.EffectiveWorkPerStroke,
			ModelDPS:                    sm.ModelDps,
			PropulsivePower:             sm.PropulsivePower,
			DriveMaximalAt:              sm.DriveMaximalAt,
			FirstPeak:                   sm.FirstPeak,
			ZeroBeforeCatch:             sm.ZeroBeforeCatch,
			MinimalFromCatch:            sm.MinimalFromCatch,
			ZeroAfterCatch:              sm.ZeroAfterCatch,
			StdDeviation:                sm.StdDeviation,
			OarAngle:                    oarAngle,
			HandleForce:                 handleForce,
			VerticalAngleBoatRoll:       verticalAngleBoatRoll,
			LegsVelocity:                legsVelocity,
			HandleSpeed:                 handleSpeed,
			HDFFig:                      hdfFig,
			BladeDF:                     bladeDF,
			Velocity:                    velocity,
			BoatAcceleration:            boatAcceleration,
			VelocityRel:                 velocityRel,
		}

		sampleMetricsArr = append(sampleMetricsArr, smm)
	}

	return res.CommonResponse{
		Code: resp.Code,
		Msg:  "查询🈯️定athleteTrainingId的详细指标",
	}, &sampleMetricsArr
}

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
