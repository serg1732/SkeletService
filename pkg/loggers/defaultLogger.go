package loggers

import "fmt"

type ILogger interface {
	Error(msg string, uuid string)
	Info(msg string, uuid string)
	Progress(msg string, uuid string)
	Debug(msg string, uuid string)
	System(msg string)
}

type ConsoleLog struct {
	formatLog string
}

func NewConsoleLogger() *ConsoleLog {
	return &ConsoleLog{
		formatLog: "%s: %s@%s",
	}
}

func printLog(msg string) {
	fmt.Println(msg)
}

func (c *ConsoleLog) Error(msg string, uuid string) {
	printLog(fmt.Sprintf(c.formatLog, "Error", msg, uuid))
}
func (c *ConsoleLog) Info(msg string, uuid string) {
	printLog(fmt.Sprintf(c.formatLog, "Info", msg, uuid))
}
func (c *ConsoleLog) Progress(msg string, uuid string) {
	printLog(fmt.Sprintf(c.formatLog, "Progress", msg, uuid))
}
func (c *ConsoleLog) Debug(msg string, uuid string) {
	printLog(fmt.Sprintf(c.formatLog, "Debug", msg, uuid))
}
func (c *ConsoleLog) System(msg string) {
	printLog(fmt.Sprintf(c.formatLog, "System", msg, ""))
}
