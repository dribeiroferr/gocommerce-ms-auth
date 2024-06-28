package config

import (
	"fmt"
	"os"
	"reflect"

	commons "github.com/dribeiroferr/gocommerce-ms-auth/common"
	"github.com/joho/godotenv"
)

// Config struct holds the configuration settings
type Config struct {
	// App Name
	AppName string
	// JWT Secret Key
	SecretKey string
	// Relational database
	RelationalDBHost     string
	RelationalDBPort     string
	RelationalDBUser     string
	RelationalDBPassword string
	RelationalDBName     string
	// Non-relational database
	NonRelationalDBHost     string
	NonRelationalDBPort     string
	NonRelationalDBUser     string
	NonRelationalDBPassword string
	NonRelationalDBName     string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	commons.Log.Info("loading configurations")

	// Define a map of default environment values
	defaultEnvs := map[string]string{
		"dev":  ".env.dev",
		"test": ".env.test",
		"prod": ".env.prod",
	}

	// Get the value of APP_ENV environment variable
	env := os.Getenv("APP_ENV")
	if _, exists := defaultEnvs[env]; !exists {
		env = "dev" // Default to "dev" if APP_ENV is not set or not in the map
	}

	// Load the environment variables from the corresponding .env file
	err := godotenv.Load(defaultEnvs[env])
	if err != nil {
		return nil, fmt.Errorf("error loading %s: %w", defaultEnvs[env], err)
	}

	// Create an instance of Config struct with environment variable values
	config := &Config{
		AppName:                 os.Getenv("APP_NAME"),
		SecretKey:               os.Getenv("SECRET_KEY"),
		RelationalDBHost:        os.Getenv("RELATIONAL_DB_HOST"),
		RelationalDBPort:        os.Getenv("RELATIONAL_DB_PORT"),
		RelationalDBUser:        os.Getenv("RELATIONAL_DB_USER"),
		RelationalDBPassword:    os.Getenv("RELATIONAL_DB_PASSWORD"),
		RelationalDBName:        os.Getenv("RELATIONAL_DB_NAME"),
		NonRelationalDBHost:     os.Getenv("NON_RELATIONAL_DB_HOST"),
		NonRelationalDBPort:     os.Getenv("NON_RELATIONAL_DB_PORT"),
		NonRelationalDBUser:     os.Getenv("NON_RELATIONAL_DB_USER"),
		NonRelationalDBPassword: os.Getenv("NON_RELATIONAL_DB_PASSWORD"),
		NonRelationalDBName:     os.Getenv("NON_RELATIONAL_DB_NAME"),
	}

	// Use reflection to check if all required environment variables are set
	missingVars := []string{}
	v := reflect.ValueOf(*config)
	typeOfConfig := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			missingVars = append(missingVars, typeOfConfig.Field(i).Name)
		}
	}

	if len(missingVars) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	commons.Log.Infof("[%s] initializing", config.AppName)
	commons.Log.Info("all required envs are loaded")

	return config, nil
}
