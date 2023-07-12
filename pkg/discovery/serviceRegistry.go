package discovery

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
)

type ConsulServiceRegistry struct {
	serviceInstances     map[string]map[string]ServiceInstance
	client               *api.Client
	localServiceInstance ServiceInstance
}

func (c ConsulServiceRegistry) Register(serviceInstance ServiceInstance) bool {
	registration := new(api.AgentServiceRegistration)
	registration.ID = serviceInstance.GetInstanceId()
	registration.Name = serviceInstance.GetServiceId()
	registration.Port = serviceInstance.GetPort()
	var tags []string
	if serviceInstance.IsSecure() {
		tags = append(tags, "secure=true")
	} else {
		tags = append(tags, "secure=false")
	}
	if serviceInstance.GetMetadata() != nil {
		for key, value := range serviceInstance.GetMetadata() {
			tags = append(tags, key+"="+value)
		}
	}
	registration.Tags = tags
	registration.Address = serviceInstance.GetHost()

	check := new(api.AgentServiceCheck)

	schema := "http"
	if serviceInstance.IsSecure() {
		schema = "https"
	}
	check.HTTP = fmt.Sprintf("%s://%s:%d/actuator/health", schema, registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "20s"
	registration.Check = check

	err := c.client.Agent().ServiceRegister(registration)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if c.serviceInstances == nil {
		c.serviceInstances = map[string]map[string]ServiceInstance{}
	}

	services := c.serviceInstances[serviceInstance.GetServiceId()]

	if services == nil {
		services = map[string]ServiceInstance{}
	}

	services[serviceInstance.GetInstanceId()] = serviceInstance

	c.serviceInstances[serviceInstance.GetServiceId()] = services

	c.localServiceInstance = serviceInstance

	return true
}

func (c ConsulServiceRegistry) Deregister() {
	if c.serviceInstances == nil {
		return
	}

	services := c.serviceInstances[c.localServiceInstance.GetServiceId()]

	if services == nil {
		return
	}

	delete(services, c.localServiceInstance.GetInstanceId())

	if len(services) == 0 {
		delete(c.serviceInstances, c.localServiceInstance.GetServiceId())
	}

	_ = c.client.Agent().ServiceDeregister(c.localServiceInstance.GetInstanceId())

	c.localServiceInstance = nil

}

func NewConsulServiceRegistry(host string, port int, token string) (*ConsulServiceRegistry, error) {
	if len(host) < 3 {
		return nil, errors.New("check host")
	}

	if port <= 0 || port > 65535 {
		return nil, errors.New("check port")
	}

	config := api.DefaultConfig()
	config.Address = host + ":" + strconv.Itoa(port)
	config.Token = token
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ConsulServiceRegistry{client: client}, nil
}
