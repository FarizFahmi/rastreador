package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var now = time.Now().Format("2006/01/02 15:04:05")

func GinLogger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[%s] [GIN] [INFO] [%d] [%s] [%s] [%s] [%dms] [%s] %s \n",
		now,
		param.StatusCode,
		param.Method,
		param.Path,
		param.ClientIP,
		param.Latency.Milliseconds(),
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func RecoveryLogger() gin.HandlerFunc {
	return func (c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log := New("RECOVER")
				log.ErrorWithoutTrace(err)

				c.AbortWithStatusJSON(500, gin.H{
					"error": "An unexpected error occurred. Please try again later / contact admin",
				})
			}
		}()
		c.Next()
	}
}

func GinDebugRoute(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	fmt.Printf("[%s] [GIN] [INFO] %v %v %v %v \n", now, httpMethod, absolutePath, handlerName, nuHandlers)
}

func GinDebugPrint(format string, values ...interface{}) {
	fmt.Printf("[%s] [GIN] [INFO] %v \n", now, values)
}