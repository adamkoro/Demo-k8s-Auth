package v1

import (
	logger "demo-k8s-auth/log"
	"demo-k8s-auth/model"
	"demo-k8s-auth/pkg/db"
	"demo-k8s-auth/pkg/jwt"
	"demo-k8s-auth/service"
	"demo-k8s-auth/vars"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// WaitGroup for sync
var wg sync.WaitGroup

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
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Username == "" || user.Password == "" || user.Email == "" {
		var msg model.StandardError
		msg.Error = "username, password or email is empty"
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	/*wg.Add(5)
	go func() {
		if !validate.Username(user.Username) {
			fmt.Println("username is invalid")
			c.JSON(http.StatusBadRequest, gin.H{"error": "username is invalid"})
			return
		}
		defer wg.Done()
	}()
	go func() {
		if !validate.Email(user.Email) {
			fmt.Println("email is invalid")
			c.JSON(http.StatusBadRequest, gin.H{"error": "email is invalid"})
			return
		}
		defer wg.Done()
	}()
	go func() {
		if !validate.Phone(user.Phone) {
			fmt.Println("phone is invalid")
			c.JSON(http.StatusBadRequest, gin.H{"error": "phone is invalid"})
			return
		}
		defer wg.Done()
	}()
	go func() {
		if !validate.Name(user.LastName) {
			fmt.Println("Lastname is invalid")
			c.JSON(http.StatusBadRequest, gin.H{"error": "lastname is invalid"})
			return
		}
		defer wg.Done()
	}()
	go func() {
		if !validate.Name(user.FirstName) {
			fmt.Println("firstname is invalid")
			c.JSON(http.StatusBadRequest, gin.H{"error": "firstname is invalid"})
			return
		}
		defer wg.Done()
	}()
	wg.Wait()*/
	// Generate ID
	user.ID = jwt.GenerateID()
	// Get current date
	user.CreatedAt = time.Now().Local().Format("2006-01-02 15:04:05")
	fmt.Println(user.ID, user.CreatedAt, user.Username, user.Password, user.Email, user.Phone, user.FirstName, user.LastName)
	err := db.Ping(vars.DBConn)
	if err != nil {
		logger.ErrorLogger.Println(err)
	} else {
		log.Println("Connected to database successfully!")
	}
	usr, err := service.CreateUser(vars.DBConn, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usr)
	/*err := db.Ping(vars.DBConn)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username or password is incorrect quit"})
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
	c.JSON(http.StatusOK, gin.H{"token": token})*/
}

// Gin user login endpoint
func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Username == "" || user.Password == "" {
		var msg model.StandardError
		msg.Error = "username or password is empty"
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	/*wg.Add(2)
	go func() {
		if !validate.Username(user.Username) {
			fmt.Println("username is invalid")
			c.JSON(http.StatusBadRequest, gin.H{"error": "username is invalid"})
			return
		}
		defer wg.Done()
	}()
	go func() {
		if !validate.Email(user.Email) {
			fmt.Println("email is invalid")
			c.JSON(http.StatusBadRequest, gin.H{"error": "email is invalid"})
			return
		}
		defer wg.Done()
	}()
	wg.Wait()*/
	err := db.Ping(vars.DBConn)
	if err != nil {
		logger.ErrorLogger.Println(err)
	} else {
		log.Println("Connected to database successfully!")
	}
	dbUser, err := service.GetUserByUsername(vars.DBConn, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		logger.ErrorLogger.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username or password is incorrect"})
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
	/*err := db.Ping(vars.DBConn)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username or password is incorrect quit"})
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
	c.JSON(http.StatusOK, gin.H{"token": token})*/
}

// Gin user logout endpoint
func Logout(c *gin.Context) {

}
