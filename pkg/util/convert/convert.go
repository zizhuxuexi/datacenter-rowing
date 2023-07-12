package convert

import (
	"time"
)

func StringToTime(date string) time.Time {
	// 获取TrainingData time.Time
	layout := "2006_01_02" // 定义日期字符串的格式
	// 解析日期字符串为时间对象
	trainingDate, _ := time.Parse(layout, date)
	return trainingDate
}
