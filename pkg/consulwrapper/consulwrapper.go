package consulwrapper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/serg1732/SkeletService/pkg/config"

	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	client *api.Client
}

func NewConsulClient() *ConsulClient {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &ConsulClient{client: client}
}

func (cc *ConsulClient) RegisterService(option config.ServiceOption) error {
	var err error
	if cc.client == nil {
		cc.client, err = api.NewClient(api.DefaultConfig())
		if err != nil {
			return err
		}
	}

	err = cc.client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name:    option.ServiceName + "_tasks",
		Address: option.ServiceIP,
		Port:    option.ServicePort,
		ID:      option.ServiceID,
		Tags:    strings.Split(option.ServiceTags, "|"),
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + option.ServiceAgentIP + ":" + strconv.Itoa(option.ServicePort) + "/healthz",
			Interval: option.ServiceIntervalCheck,
			Timeout:  option.ServiceTimeoutCheck,
		},
		Connect: &api.AgentServiceConnect{
			Native: true,
		},
	})

	return err
}

func (cc *ConsulClient) DeRegisterService(serviceID string) error {
	return cc.client.Agent().ServiceDeregister(serviceID)
}

func (cc *ConsulClient) FindServices(serviceName string) []string {
	servicesData, _, err := cc.client.Health().Service(serviceName, "", true, &api.QueryOptions{})
	if err != nil {
		fmt.Println(err)
		//TODO: output into Logs
	}
	var servicesList []string
	servicesList = make([]string, len(servicesData))
	for _, entry := range servicesData {
		servicesList = append(servicesList, fmt.Sprintf("%s:%d", entry.Node.Address, entry.Service.Port))
	}
	return servicesList
}
