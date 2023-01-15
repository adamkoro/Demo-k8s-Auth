package v1

import (
	"demo-k8s-auth/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin ping endpoint
func Ping(c *gin.Context) {
	var msg model.PingResponse
	msg.Message = "pong"
	c.JSON(http.StatusOK, msg)
}

// Gin health endpoint
func Health(c *gin.Context) {
	var msg model.HealthResponse
	msg.Status = "UP"
	msg.Message = "Service is running"
	c.JSON(http.StatusOK, msg)
}

// Gin user register endpoint
func Register(c *gin.Context) {

}

// Gin user login endpoint
func Login(c *gin.Context) {

}

// Gin user logout endpoint
func Logout(c *gin.Context) {

}
