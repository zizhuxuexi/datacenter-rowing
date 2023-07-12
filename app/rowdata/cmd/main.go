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

	// //在consul注册服务
	// consulHost := config.Conf.Consul.Host
	// consulPort := config.Conf.Consul.Port
	// //服务注册
	// registryDiscoveryClient, err := discovery.NewConsulServiceRegistry(consulHost, consulPort, "")
	// if err != nil {
	// 	panic(err)
	// }
	// //服务实例
	// serviceInstanceInfo, _ := discovery.NewDefaultServiceInstance("rowdata-service")

	grpcAddress := config.Conf.Services["row-data"].Addr
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
