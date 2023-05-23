package main

import (
	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata/internal/repository/db/dao"
	"github.com/zizhuxuexi/datacenter-rowing/config"
)

func main() {
	config.InitConfig()
	dao.InitDB()

}
