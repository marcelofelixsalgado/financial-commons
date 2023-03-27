package settings

import (
	"log"

	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// ConfigType struct to resolve env vars
type ConfigType struct {
	AppName     string `env:"APP_NAME" default:"financial-api-response-module"`
	Environment string `env:"ENVIRONMENT" default:"development"`

	// Key used to sign the token
	SecretKey []byte `env:"SECRET_KEY"`

	// Log files
	LogAccessFile string `env:"LOG_ACCESS_FILE" default:"./access.log"`
	LogAppFile    string `env:"LOG_APP_FILE" default:"./app.log"`
	LogLevel      string `env:"LOG_LEVEL" default:"INFO"`
}

var Config ConfigType

// InitConfigs initializes the environment settings
func Load() {
	// load .env (if exists)
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	}

	// bind env vars
	if err := env.Set(&Config); err != nil {
		log.Fatal(err)
	}

	if _, err := logrus.ParseLevel(Config.LogLevel); err != nil {
		log.Fatal(err)
	}
}
