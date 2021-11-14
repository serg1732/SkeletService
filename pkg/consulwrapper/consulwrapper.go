package consulwrapper

import (
	"fmt"
	"strconv"

	"github.com/serg1732/SkeletService/pkg/constants"

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

func (cc *ConsulClient) RegisterService() error {
	var err error
	if cc.client == nil {
		cc.client, err = api.NewClient(api.DefaultConfig())
		if err != nil {
			return err
		}
	}

	err = cc.client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name:    constants.ServiceName + "_tasks",
		Address: constants.ServiceIP,
		Port:    constants.ServicePort,
		ID:      constants.ServiceID,
		Tags:    []string{constants.ServiceTags},
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + constants.ServiceAgentIP + ":" + strconv.Itoa(constants.ServicePort) + "/healthz",
			Interval: constants.ServiceIntervalCheck,
			Timeout:  constants.ServiceTimeoutCheck,
		},
		Connect: &api.AgentServiceConnect{
			Native: true,
		},
	})

	return err
}

func (cc *ConsulClient) DeRegisterService() error {
	return cc.client.Agent().ServiceDeregister(constants.ServiceID)
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
