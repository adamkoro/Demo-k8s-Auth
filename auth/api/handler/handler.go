package handler

import (
	v1 "demo-k8s-auth/api/v1"

	"github.com/gin-gonic/gin"
)

// Create Gin router group to api/v1 endpoint
func ApiV1() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.GET("/ping", v1.Ping)
		api.GET("/health", v1.Health)
	}
	return router
}
