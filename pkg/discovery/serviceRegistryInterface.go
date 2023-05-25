package discovery

type ServiceRegistry interface {
	Register(serviceInstance ServiceInstance) bool
	Deregister()
}
