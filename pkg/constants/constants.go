package constants

import "strconv"

type RequestType int

const (
	GET    = 1
	POST   = 2
	PUT    = 3
	DELETE = 4
)

var ServiceName = "service-name"
var ServicePort = 8002
var ServiceIP = "127.0.0.1"
var ServiceAgentIP = "127.0.0.1"
var ServiceIntervalCheck = "5s"
var ServiceTimeoutCheck = "1s"
var ServiceID = "service-instance-" + strconv.Itoa(ServicePort)
var ServiceTags = "SERVICE"
