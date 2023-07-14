package main

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata/internal/repository/db/dao"
	"github.com/zizhuxuexi/datacenter-rowing/app/rowdata/internal/service"
	"github.com/zizhuxuexi/datacenter-rowing/config"
	pb "github.com/zizhuxuexi/datacenter-rowing/idl/pb/rowdata"
	"google.golang.org/grpc"
)

func main() {
	config.InitConfig()
	dao.InitDB()

	// grpcAddress := config.Conf.Services["row-data"].Addr
	// serviceName := config.Conf.Services["rowdata"].Name
	// port := config.Conf.Services["rowdata"].Port
	// grpcAddress := fmt.Sprintf("%s:%d", serviceName, port)
	grpcAddress := "0.0.0.0:10001"

	server := grpc.NewServer()
	defer server.Stop()
	pb.RegisterRowDataServiceServer(server, service.GetRowDataSrv())
	//pb.RegisterTrainingSummaryServiceServer(server, service.GetTrainingSummarySrv())
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	logrus.Info("server started listen on ", grpcAddress)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
