package support

import (
	"api.welooky/conf"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"net"
)

type Client interface {
	Service(string, string) ([]*consul.ServiceEntry, *consul.QueryMeta, error)
	Register(string, string, int, string, []string) error
	DeRegister(string) error
}

type client struct {
	consul *consul.Client
}

func NewConsulClient(register conf.Consul) (Client, error) {
	config := consul.DefaultConfig()
	config.Address = register.Host + ":" + register.Port
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &client{consul: c}, nil
}

func (c *client) Register(name string, addr string, port int, health string, tags []string) error {
	healthUrl := fmt.Sprintf("http://%s:%d%s", addr, port, health)
	reg := &consul.AgentServiceRegistration{
		ID:      name,
		Address: addr,
		Name:    name,
		Port:    port,
		Tags:    tags,
		Check: &consul.AgentServiceCheck{
			Status:   consul.HealthPassing,
			Interval: "5s",
			Timeout:  "20s",
			HTTP:     healthUrl,
			Method:   "GET",
		},
	}
	return c.consul.Agent().ServiceRegister(reg)
}

func (c *client) DeRegister(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}

func (c *client) Service(service, tag string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	addr, meta, err := c.consul.Health().Service(service, tag, true, nil)
	if len(addr) == 0 && err == nil {
		return nil, nil, fmt.Errorf("service ( %s ) was not found", service)
	}
	if err != nil {
		return nil, nil, err
	}
	return addr, meta, nil
}

func IPAddr() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil || ipnet.IP.To16() != nil {
				return ipnet.IP, nil
			}
		}
	}
	return nil, nil
}
