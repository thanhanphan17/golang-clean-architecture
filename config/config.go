package config

import (
	"os"

	"github.com/mitchellh/mapstructure"

	"github.com/spf13/viper"
)

type AppConfig struct {
	BuildEnv string `mapstructure:"BUILD_ENV"`

	// service info
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	ServiceHost    string `mapstructure:"SERVICE_HOST"`
	ServicePort    uint   `mapstructure:"SERVICE_PORT"`
	ServiceTimeout uint   `mapstructure:"SERVICE_TIMEOUT"`

	// database info
	DBHost                      string `mapstructure:"DB_HOST"`
	DBPort                      string `mapstructure:"DB_PORT"`
	DBName                      string `mapstructure:"DB_NAME"`
	DBUserName                  string `mapstructure:"DB_USERNAME"`
	DBPassword                  string `mapstructure:"DB_PASSWORD"`
	SSLCertAuthorityCertificate string `mapstructure:"SSL_CERT_AUTH"`
	// pool connection
	MaxOpenConnections    int `mapstructure:"MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections    int `mapstructure:"MAX_IDLE_CONNECTIONS"`
	ConnectionMaxIdleTime int `mapstructure:"CONNECTION_MAX_IDLE_TIME"`
	ConnectionMaxLifeTime int `mapstructure:"CONNECTION_MAX_LIFE_TIME"`
	ConnectionTimeout     int `mapstructure:"CONNECTION_TIMEOUT"`

	// jwt info
	JWTSecretKey       string `mapstructure:"JWT_SECRET_KEY"`
	VerifyTokenExpiry  uint   `mapstructure:"VERIFY_TOKEN_EXPIRY"`
	AccessTokenExpiry  uint   `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry uint   `mapstructure:"REFRESH_TOKEN_EXPIRY"`

	// mail info
	MailFrom   string `mapstructure:"MAIL_FROM"`
	MailServer string `mapstructure:"MAIL_SERVER"`
	MailPort   int    `mapstructure:"MAIL_PORT"`
	MailPass   string `mapstructure:"MAIL_PASS"`
}

// InitLoadAppConf initializes and returns the application configuration.
func InitLoadAppConf() *AppConfig {
	// Get the path of the configuration file from environment variable.
	configPath := os.Getenv("CONFIG_PATH")

	// Get the environment from environment variable. It can be either "dev" or "prod".
	env := os.Getenv("CONFIG_ENV")

	// Load the configuration from the file.
	appConfig, err := LoadConfigFromFile(configPath, env)
	if err != nil {
		panic(err)
	}

	// Create a new AppConfig struct to store the decoded configuration.
	var result AppConfig

	// Configure the decoder to allow weakly typed input and set the result to the AppConfig struct.
	customConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
	}

	// Create a new decoder with the custom configuration.
	decoder, err := mapstructure.NewDecoder(customConfig)
	if err != nil {
		panic(err)
	}

	// Decode the appConfig into the result struct.
	err = decoder.Decode(appConfig)
	if err != nil {
		panic(err)
	}

	// Return a pointer to the result struct.
	return &result
}

// LoadConfigFromFile loads a configuration file from the given path and returns the loaded configuration.
// It expects the path to the configuration file and the name of the configuration file as parameters.
func LoadConfigFromFile(path string, configName string) (config interface{}, err error) {
	// Add the given path to the list of search paths for the configuration file.
	viper.AddConfigPath(path)

	// Set the name of the configuration file to be loaded.
	viper.SetConfigName(configName)

	// Set the type of the configuration file to be loaded as environment variable.
	viper.SetConfigType("env")

	// Automatically read in environment variables that match the configuration keys.
	viper.AutomaticEnv()

	// Read in the configuration file.
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Unmarshal the configuration file into the given config interface.
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return
}
