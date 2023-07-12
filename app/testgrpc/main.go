package main

import (
	"context"
	"log"
	"time"

	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/rowdata"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	layout := "2006_01_02" // 定义日期字符串的格式
	// 解析日期字符串为时间对象
	trainingDate, err := time.Parse(layout, "2021_06_23")
	if err != nil {
		panic(err)
	}

	client := pb.NewRowDataServiceClient(conn)
	// req := &pb.AthleteTrainingDataAddRequest{
	// 	TrainingId:     1,
	// 	Name:           "zhhx",
	// 	Gender:         "男",
	// 	Seat:           1,
	// 	Side:           1,
	// 	Height:         1.74,
	// 	Weight:         67.5,
	// 	OarInboard:     1.1,
	// 	OarLength:      1.5,
	// 	OarBladeLength: 1.2,
	// }
	req2 := &pb.TrainingSummaryGetRequset{
		TrainingSummary: &pb.TrainingSummaryModel{
			TrainingId:      4,
			TrainingName:    "2021_06_23_1_WH8+",
			TrainDate:       timestamppb.New(trainingDate),
			EventGender:     "W",
			EventPeopleType: "8+",
			EventScale:      "H",
			Event:           "WH8+",
			Weather:         "晴天",
			Temp:            24,
			WindDir:         "顺风",
			Loc:             "千岛湖",
			Coach:           "张三",
			SampleCount:     7,
		},
		Set:                true,
		SetTrainingId:      true,
		SetTrainingName:    true,
		SetTrainingDate:    true,
		SetEventGender:     true,
		SetEventPeopleType: true,
		SetEventScale:      true,
		SetEvent:           true,
		SetWeather:         true,
		SetTemp:            true,
		SetWindDir:         true,
		SetLoc:             true,
		SetCoach:           true,
		SetSampleCount:     true,
	}

	// resp, err := client.AthleteTrainingDataAdd(context.Background(), req)
	// if err != nil {
	// 	panic(err)
	// }

	resp, err := client.TrainingSummaryGet(context.Background(), req2)
	if err != nil {
		panic(err)
	}

	for _, r := range resp.TrainingSummary {
		log.Println(*r)
	}
	// log.Println(resp.AthleteTrainingDataDetail)
}
