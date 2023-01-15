package env

import (
	"os"
	"strconv"
)

var (
	DbUsername string
	DbPassword string
	DbHost     string
	DbPort     int
	DbName     string
)

// Get DbUsername from environment variables
func getDbUsername() string {
	return os.Getenv("DB_USERNAME")
}

// Get DbPassword from environment variables
func getDbPassword() string {
	return os.Getenv("DB_PASSWORD")
}

// Get DbHost from environment variables
func getDbHost() string {
	return os.Getenv("DB_HOST")
}

// Get DbPort from environment variables
func getDbPort() string {
	return os.Getenv("DB_PORT")
}

// Get DbName from environment variables
func getDbName() string {
	return os.Getenv("DB_NAME")
}

// Check if DbUsername is empty and set default value
func CheckDbUsername(username string) string {
	if len(username) != 0 {
		return username
	}
	username = "postgres"
	return username
}

// Check if DbPassword is empty and set default value
func CheckDbPassword(password string) string {
	if len(password) != 0 {
		return password
	}
	password = "postgres"
	return password
}

// Check if DbHost is empty and set default value
func CheckDbHost(host string) string {
	if len(host) != 0 {
		return host
	}
	host = "localhost"
	return host
}

// Check if DbPort is empty and set default value
func CheckDbPort(port string) string {
	if len(port) != 0 {
		return port
	}
	port = "5432"
	return port
}

// Check if DbName is empty and set default value
func CheckDbName(name string) string {
	if len(name) != 0 {
		return name
	}
	name = "postgres"
	return name
}

// Convert DbPort value to integer
func ConvertToInteger(port string) (int, error) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return 0, err
	}
	return portInt, nil
}

// Override variables
func CheckEnvs() {
	var err error
	DbUsername = CheckDbUsername(getDbUsername())
	DbPassword = CheckDbPassword(getDbPassword())
	DbHost = CheckDbHost(getDbHost())
	DbPort, _ = ConvertToInteger(CheckDbPort(getDbPort()))
	if err != nil {
		panic(err)
	}
	DbName = CheckDbName(getDbName())
}
