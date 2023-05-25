package discovery

import (
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

	}
}
