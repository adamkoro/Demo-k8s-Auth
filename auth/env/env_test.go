package env

import (
	"os"
	"testing"
)

// Test if DbUsername is not empty and return value
func TestGetDbUsernameEnvSetted(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_USERNAME", "test")
	// Get environment variables
	DbUsername = getDbUsername()
	// Check if DbUsername is empty and set default value
	if DbUsername != CheckDbUsername(DbUsername) {
		t.Errorf("Expected %s, got %s", "test", DbUsername)
	}
}

// Test if DbUsername is empty and set default value
func TestGetDbUsernameEnvNotSetted(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_USERNAME", "")
	// Set default value
	testDbUsername := "postgres"
	// Check if DbUsername is empty and set default value
	if testDbUsername != CheckDbUsername(getDbUsername()) {
		t.Errorf("Expected %s, got %s", "guest", DbUsername)
	}
}

// Test if DbPassword is not empty and return value
func TestGetDbPasswordEnvSetted(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_PASSWORD", "test")
	// Get environment variables
	DbPassword = getDbPassword()
	// Check if DbPassword is empty and set default value
	if DbPassword != CheckDbPassword(DbPassword) {
		t.Errorf("Expected %s, got %s", "test", DbPassword)
	}
}

// Test if DbPassword is empty and set default value
func TestGetDbPasswordEnvNotSetted(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_PASSWORD", "")
	// Set default value
	testDbPassword := "postgres"
	// Check if DbPassword is empty and set default value
	if testDbPassword != CheckDbPassword(getDbPassword()) {
		t.Errorf("Expected %s, got %s", "guest", DbPassword)
	}
}

// Test if DbHost is not empty and return value
func TestGetDbHostEnvSetted(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_HOST", "test.domain.local")
	// Get environment variables
	DbHost = getDbHost()
	// Check if DbHost is empty and set default value
	if DbHost != CheckDbHost(DbHost) {
		t.Errorf("Expected %s, got %s", "test.domain.local", DbHost)
	}
}

// Test if DbHost is empty and set default value
func TestGetDbHostEnvNotSetted(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_HOST", "")
	// Set default value
	testDbHost := "localhost"
	// Check if DbHost is empty and set default value
	if testDbHost != CheckDbHost(getDbHost()) {
		t.Errorf("Expected %s, got %s", "localhost", DbHost)
	}
}

// Test if DbPort is not empty and return value
func TestGetDbPortEnvSetted(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_PORT", "5432")
	// Get environment variables
	testDbPort := getDbPort()
	// Check if DbPort is empty and set default value
	if testDbPort != CheckDbPort(getDbPort()) {
		t.Errorf("Expected %d, got %d", 5432, DbPort)
	}
}

// Test if DbPort is empty and return value
func TestGetDbPortEnvNotSetted(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_PORT", "")
	// Set default value
	testDbPort := "5432"
	// Check if DbPort is empty and set default value
	if testDbPort != CheckDbPort((getDbPort())) {
		t.Errorf("Expected %d, got %d", 5432, DbPort)
	}
}

// Test if DbName is not empty and return value
func TestGetDbNameEnvSetted(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_NAME", "test")
	// Get environment variables
	DbName = getDbName()
	// Check if DbName is empty and set default value
	if DbName != CheckDbName(DbName) {
		t.Errorf("Expected %s, got %s", "test", DbName)
	}
}

// Test if DbName is empty and set default value
func TestGetDbNameEnvNotSetted(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_NAME", "")
	// Set default value
	testDbName := "postgres"
	// Check if DbName is empty and set default value
	if testDbName != CheckDbName(getDbName()) {
		t.Errorf("Expected %s, got %s", "guest", DbName)
	}
}

// Test integer conversion
func TestConvertToInteger(t *testing.T) {
	tmpInt := 12345
	tmpStr := "12345"
	convInt, err := ConvertToInteger(tmpStr)
	if err != nil {
		t.Errorf("Expected %s, got %s", "nil", err)
	}

	if tmpInt != convInt {
		t.Errorf("Expected %d, got %d", tmpInt, convInt)
	}
}

// Test CheckEnvs function
func TestCheckEnvs(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_USERNAME", "test")
	os.Setenv("DB_PASSWORD", "test")
	os.Setenv("DB_HOST", "test.domain.local")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "test")
	// Get environment variables
	CheckEnvs()
	// Check if DbUsername is empty and set default value
	if DbUsername != CheckDbUsername(DbUsername) {
		t.Errorf("Expected %s, got %s", "test", DbUsername)
	}
	// Check if DbPassword is empty and set default value
	if DbPassword != CheckDbPassword(DbPassword) {
		t.Errorf("Expected %s, got %s", "test", DbPassword)
	}
	// Check if DbHost is empty and set default value
	if DbHost != CheckDbHost(DbHost) {
		t.Errorf("Expected %s, got %s", "test.domain.local", DbHost)
	}
	expectedDbPort, err := ConvertToInteger(CheckDbPort((getDbPort())))
	if err != nil {
		t.Errorf("Expected %d, got %d", 5432, DbPort)
	}
	// Check if DbPort is empty and set default value
	if DbPort != expectedDbPort {
		t.Errorf("Expected %d, got %d", 5432, DbPort)
	}
	// Check if DbName is empty and set default value
	if DbName != CheckDbName(DbName) {
		t.Errorf("Expected %s, got %s", "test", DbName)
	}
}

func TestCheckEnvsEmpty(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_USERNAME", "test")
	os.Setenv("DB_PASSWORD", "test")
	os.Setenv("DB_HOST", "test.domain.local")
	//os.Setenv("DB_PORT", "")
	//os.Setenv("DB_NAME", "")
	// Get environment variables
	CheckEnvs()
	// Check if DbUsername is empty and set default value
	if DbUsername != CheckDbUsername(DbUsername) {
		t.Errorf("Expected %s, got %s", "test", DbUsername)
	}
	// Check if DbPassword is empty and set default value
	if DbPassword != CheckDbPassword(DbPassword) {
		t.Errorf("Expected %s, got %s", "test", DbPassword)
	}
	// Check if DbHost is empty and set default value
	if DbHost != CheckDbHost(DbHost) {
		t.Errorf("Expected %s, got %s", "test.domain.local", DbHost)
	}
	// Test value for DbPort
	expectedDbPort := 5432
	// Function to get DbPort
	getTestDbPort, err := ConvertToInteger(CheckDbPort((getDbPort())))
	if err != nil {
		t.Errorf("Expected %d, got %d", 5432, getTestDbPort)
	}

	// Check if DbPort is empty and set default value
	if expectedDbPort != getTestDbPort {
		t.Errorf("Expected %d, got %d", expectedDbPort, getTestDbPort)
	}
	// Check if DbName is empty and set default value
	if DbName != CheckDbName(DbName) {
		t.Errorf("Expected %s, got %s", "postgres", DbName)
	}
}

func TestCheckEnvsHalfEmpty(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_USERNAME", "test")
	os.Setenv("DB_PASSWORD", "test")
	os.Setenv("DB_HOST", "test.domain.local")
	os.Setenv("DB_PORT", "")
	os.Setenv("DB_NAME", "")
	// Get environment variables
	CheckEnvs()
	// Check if DbUsername is empty and set default value
	if DbUsername != CheckDbUsername(DbUsername) {
		t.Errorf("Expected %s, got %s", "test", DbUsername)
	}
	// Check if DbPassword is empty and set default value
	if DbPassword != CheckDbPassword(DbPassword) {
		t.Errorf("Expected %s, got %s", "test", DbPassword)
	}
	// Check if DbHost is empty and set default value
	if DbHost != CheckDbHost(DbHost) {
		t.Errorf("Expected %s, got %s", "test.domain.local", DbHost)
	}
	// Test value for DbPort
	expectedDbPort := 5432
	// Function to get DbPort
	getTestDbPort, err := ConvertToInteger(CheckDbPort((getDbPort())))
	if err != nil {
		t.Errorf("Expected %d, got %d", 5432, getTestDbPort)
	}

	// Check if DbPort is empty and set default value
	if expectedDbPort != getTestDbPort {
		t.Errorf("Expected %d, got %d", expectedDbPort, getTestDbPort)
	}
	// Check if DbName is empty and set default value
	if DbName != CheckDbName(DbName) {
		t.Errorf("Expected %s, got %s", "postgres", DbName)
	}
}

func TestCheckEnvsCustomValues(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_USERNAME", "guest1")
	os.Setenv("DB_PASSWORD", "quest1")
	os.Setenv("DB_HOST", "dev.postgres.local")
	os.Setenv("DB_PORT", "5444")
	os.Setenv("DB_NAME", "dev-data")
	// Get environment variables
	CheckEnvs()
	// Check if DbUsername is empty and set default value
	if DbUsername != CheckDbUsername(DbUsername) {
		t.Errorf("Expected %s, got %s", "guest1", DbUsername)
	}
	// Check if DbPassword is empty and set default value
	if DbPassword != CheckDbPassword(DbPassword) {
		t.Errorf("Expected %s, got %s", "guest1", DbPassword)
	}
	// Check if DbHost is empty and set default value
	if DbHost != CheckDbHost(DbHost) {
		t.Errorf("Expected %s, got %s", "dev.postgres.local", DbHost)
	}
	// Test value for DbPort
	testDbPort := 5444
	// Function to get DbPort
	testDbPort, err := ConvertToInteger(CheckDbPort((getDbPort())))
	if err != nil {
		t.Errorf("Expected %d, got %d", 5432, testDbPort)
	}
	// Check if DbName is empty and set default value
	if DbName != CheckDbName(DbName) {
		t.Errorf("Expected %s, got %s", "dev-data", DbName)
	}
}
