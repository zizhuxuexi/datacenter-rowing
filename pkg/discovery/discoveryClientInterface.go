package discovery

type DiscoveryClient interface {
	GetInstances(serviceId string) ([]ServiceInstance, error)

	GerServices() ([]string, error)
}
