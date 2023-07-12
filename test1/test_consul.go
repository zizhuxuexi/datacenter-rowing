package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zizhuxuexi/datacenter-rowing/pkg/discovery"
)

func main() {
	ConsulServiceRegistry()

}

func ConsulServiceRegistry() {
	host := "127.0.0.1"
	port := 8500
	registryDiscoveryClient, _ := discovery.NewConsulServiceRegistry(host, port, "")

	_, err := GetOutBoundIP()
	if err != nil {
		fmt.Println("ip error")
	}

	serviceInstanceInfo, _ := discovery.NewDefaultServiceInstance("test_service", "", 8090, false, map[string]string{"user": "zyn"}, "")

	registryDiscoveryClient.Register(serviceInstanceInfo)

	r := gin.Default()

	r.GET("/actuator/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err = r.Run(":8090")
	if err != nil {
		registryDiscoveryClient.Deregister()
	}
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return ip, err
}
