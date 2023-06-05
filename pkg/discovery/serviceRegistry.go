package discovery

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type consulServiceRegistry struct {
	serviceInstance      map[string]map[string]ServiceInstance
	client               api.Client
	localServiceInstance ServiceInstance
}

func (c consulServiceRegistry) Register(serviceInstance ServiceInstance) bool {
	registration := new(api.AgentServiceRegistration)
	registration.ID = serviceInstance.GetInstanceId()
	registration.Name = serviceInstance.GetServiceId()
	registration.Port = serviceInstance.GetPort()
	var tags []string
	if serviceInstance.IsSecure() {
		tags = append(tags, "secure=true")
	}else {
		tags = append(tags, "secure=false")
	}
	if serviceInstance.GetMetadata() != nil{
		for key, value := range serviceInstance.GetMetadata(){
			tags = append(tags, key+"="+value)
		}
	}
	registration.Tags = tags
	registration.Address = serviceInstance.GetHost()

	check := new(api.AgentServiceCheck)

	schema := "http"
	if serviceInstance.IsSecure(){
		schema = "https"
	}
	check.HTTP = fmt.Sprintf("%s://%s:%d/actuator/health", schema, registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "20s"
	registration.Check = check

	err := c.client.Agent().ServiceRegister(registration)
	if err != nil{
		fmt.Println(err)
		return false
	}

	if
}
