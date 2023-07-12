package dao

import (
	"context"
	"fmt"
	"time"

	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/trainingSummary"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateTest() {
	fmt.Println("---------------------------1---------------")
	c := new(context.Context)
	dao := NewTrainingSummaryDao(*c)
	fmt.Println("---------------------------2----------------")
	req := new(pb.TrainingSummaryAddRequest)
	req.TrainingName = "2021_06_24_1_WH8+"
	req.TrainDate = timestamppb.New(time.Now())
	req.EventGender = "W"
	req.EventPeopleType = "8+"
	req.EventScale = "H"
	req.Event = "WH8+"
	req.Weather = "晴天"
	req.Temp = 25
	req.WindDir = "顺风"
	req.Loc = "千岛湖"
	req.Coach = "教练员2"
	req.SampleCount = 7
	err := dao.AddTrainingSummary(req)
	fmt.Println(err)
	resp, err1 := dao.GetTrainingSummaryByTrainingName(&pb.TrainingSummaryGetByTrainingNameRequest{
		TrainingName: "2021_06_23_1_WH8+",
	})
	fmt.Println(resp)
	fmt.Println(err1)

}

func CreateTest2() {
	fmt.Println("---------------------------1---------------")
	c := new(context.Context)
	dao := NewTrainingSummaryDao(*c)
	fmt.Println("---------------------------2----------------")
	req := new(pb.TrainingSummaryAddRequest)
	req.TrainingName = "2021_06_24_1_WH8+"
	req.TrainDate = timestamppb.New(time.Now())
	req.EventGender = "W"
	req.EventPeopleType = "8+"
	req.EventScale = "H"
	req.Event = "WH8+"
	req.Weather = "晴天"
	req.Temp = 25
	req.WindDir = "顺风"
	req.Loc = "千岛湖"
	req.Coach = "教练员2"
	req.SampleCount = 7
	err := dao.AddTrainingSummary(req)
	fmt.Println(err)
	resp, err1 := dao.GetTrainingSummaryByTrainingName(&pb.TrainingSummaryGetByTrainingNameRequest{
		TrainingName: "2021_06_23_1_WH8+",
	})
	fmt.Println(resp)
	fmt.Println(err1)

}

func main() {
	fmt.Println(time.Now())
	fmt.Println("aaaaaa")
}
