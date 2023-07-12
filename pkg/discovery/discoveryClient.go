package discovery

func (c ConsulServiceRegistry) GetInstances(serviceId string) ([]ServiceInstance, error) {
	catalogService, _, _ := c.client.Catalog().Service(serviceId, "", nil)
	if len(catalogService) > 0 {
		result := make([]ServiceInstance, len(catalogService))
		for index, server := range catalogService {
			s := DefaultServiceInstance{
				InstanceId: server.ServiceID,
				ServiceId:  server.ServiceName,
				Host:       server.Address,
				Port:       server.ServicePort,
				Metadata:   server.ServiceMeta,
			}
			result[index] = s
		}
		return result, nil
	} else {
		return nil, nil
	}
}

func (c ConsulServiceRegistry) GetServices() ([]string, error) {
	services, _, _ := c.client.Catalog().Services(nil)
	resutl := make([]string, len(services))
	index := 0
	for serviceName, _ := range services {
		resutl[index] = serviceName
		index++
	}
	return resutl, nil
}
