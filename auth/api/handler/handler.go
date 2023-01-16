package handler

import (
	v1 "demo-k8s-auth/api/v1"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Create Gin router group to api/v1 endpoint
func Router() *gin.Engine {
	router := gin.New()

	// Custom logger: [HTTP] 2023/01/08 18:47:25 | Code: 404 | Method: GET | IP: 127.0.0.1 | Path: /producer/v1/test
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[HTTP] %s | Code: %d | Method: %s | IP: %s | Path: %s\n",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			param.StatusCode,
			param.Method,
			param.ClientIP,
			param.Path,
		)
	}))

	api := router.Group("/api/v1")
	{
		api.GET("/ping", v1.Ping)
		api.GET("/health", v1.Health)
	}
	return router
}
