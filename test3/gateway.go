package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zizhuxuexi/datacenter-rowing/pkg/discovery"
)

type Gateway struct {
	consulClient *discovery.ConsulServiceRegistry
	proxy        *httputil.ReverseProxy
	router       *gin.Engine
}

func NewGateway(host string, port int) (*Gateway, error) {
	g := &Gateway{
		proxy: &httputil.ReverseProxy{},
	}

	g.consulClient, _ = discovery.NewConsulServiceRegistry(host, port, "")

	// 创建 Consul 客户端
	// config := api.DefaultConfig()
	// config.Address = consulAddr
	// consulClient, err := api.NewClient(config)
	// if err != nil {
	// 	return nil, err
	// }
	// g.consulClient = consulClient

	// 注册路由
	router := gin.Default()
	router.Any("/*path", g.handleRequest)
	g.router = router

	return g, nil
}

func (g *Gateway) handleRequest(c *gin.Context) {
	// 获取请求路径
	requestPath := c.Request.URL.Path
	prefix := ""
	firstSlashIndex := strings.Index(requestPath, "/")
	if firstSlashIndex == -1 {
		// 如果没有找到斜杠，则返回整个字符串
		prefix = requestPath
	}

	// 截取第一个斜杠后面的子串
	substring := requestPath[firstSlashIndex+1:]

	// 查找第二个斜杠的索引
	secondSlashIndex := strings.Index(substring, "/")
	if secondSlashIndex == -1 {
		// 如果没有找到第二个斜杠，则返回截取的子串
		prefix = substring
	}

	// 截取第一个斜杠和第二个斜杠之间的子串
	prefix = substring[:secondSlashIndex]

	// 通过 Consul 获取服务实例

	serviceInstances, err := g.consulClient.GetInstances(prefix)
	if err != nil {
		log.Println("Failed to get service instances from Consul:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// serviceInstances, _, err := g.consulClient.Catalog().Service(g.serviceName, "", nil)
	// if err != nil {
	// 	log.Println("Failed to get service instances from Consul:", err)
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// 	return
	// }

	fmt.Println(serviceInstances)

	//TODO 如果没有这个注册服务 就直接从网关上转发 或者 提示错误
	if len(serviceInstances) == 0 {
		log.Println("No instances found for service:", prefix)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// 选择负载均衡策略，这里简单地选择第一个实例
	serviceInstance := serviceInstances[0]

	// 创建转发目标 URL
	schema := "http"
	if serviceInstance.IsSecure() {
		schema = "https"
	}
	targetURL := fmt.Sprintf("%s://%s:%d%s", schema, serviceInstance.GetHost(), serviceInstance.GetPort(), strings.TrimPrefix(requestPath, "/"+prefix))

	// 设置代理目标
	target, err := url.Parse(targetURL)
	if err != nil {
		log.Println("Failed to parse target URL:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	g.proxy.Director = func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
		req.URL.RawQuery = target.RawQuery
		req.Host = target.Host
	}

	// 执行代理转发
	g.proxy.ServeHTTP(c.Writer, c.Request)

	// // 检查是否是要转发的服务请求
	// if strings.HasPrefix(requestPath, "/"+g.serviceName+"/") || requestPath == "/"+g.serviceName || requestPath == "/"+g.serviceName+"/" {
	// 	// 通过 Consul 获取服务实例

	// 	serviceInstances, err := g.consulClient.GetInstances(g.serviceName)
	// 	if err != nil {
	// 		log.Println("Failed to get service instances from Consul:", err)
	// 		c.AbortWithStatus(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	// serviceInstances, _, err := g.consulClient.Catalog().Service(g.serviceName, "", nil)
	// 	// if err != nil {
	// 	// 	log.Println("Failed to get service instances from Consul:", err)
	// 	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// 	// 	return
	// 	// }

	// 	fmt.Println(serviceInstances)

	// 	if len(serviceInstances) == 0 {
	// 		log.Println("No instances found for service:", g.serviceName)
	// 		c.AbortWithStatus(http.StatusNotFound)
	// 		return
	// 	}

	// 	// 选择负载均衡策略，这里简单地选择第一个实例
	// 	serviceInstance := serviceInstances[0]

	// 	// 创建转发目标 URL
	// 	schema := "http"
	// 	if serviceInstance.IsSecure() {
	// 		schema = "https"
	// 	}
	// 	targetURL := fmt.Sprintf("%s://%s:%d%s", schema, serviceInstance.GetHost(), serviceInstance.GetPort(), strings.TrimPrefix(requestPath, "/"+g.serviceName))

	// 	// 设置代理目标
	// 	target, err := url.Parse(targetURL)
	// 	if err != nil {
	// 		log.Println("Failed to parse target URL:", err)
	// 		c.AbortWithStatus(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	g.proxy.Director = func(req *http.Request) {
	// 		req.URL.Scheme = target.Scheme
	// 		req.URL.Host = target.Host
	// 		req.URL.Path = target.Path
	// 		req.URL.RawQuery = target.RawQuery
	// 		req.Host = target.Host
	// 	}

	// 	// 执行代理转发
	// 	g.proxy.ServeHTTP(c.Writer, c.Request)
	// } else {
	// 	// 如果不是服务请求，可以在这里处理其他自定义逻辑
	// 	// ...
	// }
}

func main() {
	// 创建网关
	gateway, err := NewGateway("127.0.0.1", 8500)
	if err != nil {
		log.Fatal("Failed to create gateway:", err)
	}

	// 启动网关服务
	err = gateway.router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start gateway server:", err)
	}
}
