package model

import "time"

type TrainingSummary struct {
	TrainingId      uint      `gorm:"<-:false;primary_key;AUTO_INCREMENT;comment:'训练ID主键'"`
	TrainingName    string    `gorm:"type:varchar(50);not null;comment:'训练名称'"`
	TrainDate       time.Time `gorm:"type:Date;not null;comment:'训练日期'"`
	EventGender     string    `gorm:"type:char(1);not null;comment:'项目选手性别'"`
	EventPeopleType string    `gorm:"type:char(2);not null;comment:'项目总人数与类型'"`
	EventScale      string    `gorm:"type:char(1);not null;comment:'项目量级'"`
	Event           string    `gorm:"type:varchar(50);not null;comment:'项目类型'"`
	Weather         string    `gorm:"type:varchar(50);not null;comment:'天气'"`
	Temp            int       `gorm:"not null;comment:'气温'"`
	WindDir         string    `gorm:"type:varchar(50);not null;comment:'风向'"`
	Loc             string    `gorm:"type:varchar(50);not null;comment:'训练地点'"`
	Coach           string    `gorm:"type:varchar(50);not null;comment:'教练员名称'"`
	SampleCount     int       `gorm:"not null;comment:'采样桨频的数目'"`
	Remark          string    `gorm:"type:varchar(50);comment:'备注'"`
}

func (*TrainingSummary) TableName() string {
	return "training_summary"
}
