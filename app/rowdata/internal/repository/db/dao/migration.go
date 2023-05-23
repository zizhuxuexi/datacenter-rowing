package dao

import (
	"fmt"
	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata/internal/repository/db/model"
)

func migration() {
	//自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.TrainingSummary{},
		)
	if err != nil {
		fmt.Println("migration err", err)
	}
}
