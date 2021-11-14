package config

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
