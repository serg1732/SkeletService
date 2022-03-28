package config

import (
	"flag"
	"strconv"
)

type ServiceOption struct {
	ServiceName          string
	ServicePort          int
	ServiceIP            string
	ServiceAgentIP       string
	ServiceIntervalCheck string
	ServiceTimeoutCheck  string
	ServiceID            string
	ServiceTags          string
}

func NewDefaultOption() *ServiceOption {
	option := new(ServiceOption)
	flag.StringVar(
		&option.ServiceName, "sname",
		"service-name", "Usage -sname=SERVICE_NAME")
	flag.IntVar(
		&option.ServicePort, "sport",
		8002, "Usage -sport=SERVICE_PORT")
	flag.StringVar(
		&option.ServiceIP, "sip",
		"127.0.0.1", "Usage -sip=SERVICE_IP_ADDRESS")
	flag.StringVar(
		&option.ServiceAgentIP, "aip",
		"127.0.0.1", "Usage -aip=CONSUL_AGENT_IP_ADDRESS")
	flag.StringVar(
		&option.ServiceIntervalCheck, "sinterval",
		"5s", "Usage -sinterval=(TIME_IN_SECONDS)s")
	flag.StringVar(
		&option.ServiceTimeoutCheck, "stimeout",
		"1s", "Usage -stimeout=(TIME_IN_SECONDS)s")
	flag.StringVar(
		&option.ServiceTags, "stags",
		"Service", "Usage -stags=tag_one|tag_two|...|tag_last")
	flag.Parse()
	option.ServiceID = option.ServiceName + "-instance-" + strconv.Itoa(option.ServicePort)
	return option
}
