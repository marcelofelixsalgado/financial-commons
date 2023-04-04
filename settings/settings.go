package settings

// Common ConfigType struct to resolve env vars
type ConfigType struct {
	AppName     string `env:"APP_NAME" default:"financial-api-response-module"`
	Environment string `env:"ENVIRONMENT" default:"development"`

	// Database Connection String (MySQL)
	DatabaseConnectionUser          string `env:"DATABASE_USER"`
	DatabaseConnectionPassword      string `env:"DATABASE_PASSWORD"`
	DatabaseConnectionServerAddress string `env:"DATABASE_SERVER_ADDRESS"`
	DatabaseConnectionServerPort    int    `env:"DATABASE_SERVER_PORT" default:"3306"`
	DatabaseName                    string `env:"DATABASE_NAME"`

	// HTTP Port to expose the API
	ApiHttpPort int `env:"API_PORT"`

	// Key used to sign the token
	SecretKey []byte `env:"SECRET_KEY"`

	ServerCloseWait int `env:"SERVER_CLOSEWAIT" default:"10"`

	// Log files
	LogAccessFile string `env:"LOG_ACCESS_FILE" default:"./access.log"`
	LogAppFile    string `env:"LOG_APP_FILE" default:"./app.log"`
	LogLevel      string `env:"LOG_LEVEL" default:"INFO"`
}

var Config ConfigType
