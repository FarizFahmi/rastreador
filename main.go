package main

import (
	logger "packages/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(logger.RecoveryLogger())
	r.Use(gin.LoggerWithFormatter(logger.GinLogger))

	gin.LoggerWithFormatter(logger.GinLogger)

	gin.DebugPrintRouteFunc = logger.GinDebugRoute
	gin.DebugPrintFunc = logger.GinDebugPrint

	r.GET("/ping", func(c *gin.Context) {
		log := logger.New("PING")
		
		log.Log("Test log ping")
		var a any
		a = 1
		a = a.(string)


		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8082")
}
