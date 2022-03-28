package main

import (
	"fmt"
	"time"

	"github.com/serg1732/SkeletService/internal/loggers"
	"github.com/serg1732/SkeletService/internal/skeletservice"

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
