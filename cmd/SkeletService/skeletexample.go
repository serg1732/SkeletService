package main

import (
	"fmt"
	"time"

	"github.com/serg1732/SkeletService/pkg/skeletservice"

	"github.com/serg1732/SkeletService/pkg/loggers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cLogger := loggers.NewConsoleLogger()

	router.GET("/test", func(c *gin.Context) {
		now := time.Now()
		fmt.Fprintf(c.Writer, "%s", now.String())
	})

	service := skeletservice.NewService(router, cLogger)
	service.Start()

}
