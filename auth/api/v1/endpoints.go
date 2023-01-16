package v1

import (
	"demo-k8s-auth/model"
	"demo-k8s-auth/pkg/jwt"
	"demo-k8s-auth/vars"
	"fmt"
	"log"
	"net/http"

	logger "demo-k8s-auth/log"
	"demo-k8s-auth/pkg/db"

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
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is empty"})
		return
	}
	//defer db.Close(vars.DBConn)
	err := db.Ping(vars.DBConn)
	if err != nil {
		logger.ErrorLogger.Println(err)
	} else {
		log.Println("Connected to database successfully!")
	}
	var dbUser model.User
	fmt.Println(user)
	err = db.QueryRow(vars.DBConn, "SELECT * FROM users WHERE username = $1", user.Username).Scan(&dbUser.Username, &dbUser.Password)
	if err != nil {
		logger.ErrorLogger.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect quit"})
		return
	}
	if user.Password != dbUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect"})
		return
	}
	token, err := jwt.Create(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Gin user logout endpoint
func Logout(c *gin.Context) {

}
