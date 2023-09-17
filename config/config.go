package config

import (
	"os"
	"strconv"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type AppConfig struct {
	BuildEnv string `mapstructure:"BUILD_ENV"`

	// service info
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	ServiceHost    string `mapstructure:"SERVICE_HOST"`
	ServicePort    int    `mapstructure:"SERVICE_PORT"`
	ServiceTimeout int    `mapstructure:"SERVICE_TIMEOUT"`

	// database info
	DBHost                      string `mapstructure:"DB_HOST"`
	DBPort                      string `mapstructure:"DB_PORT"`
	DBName                      string `mapstructure:"DB_NAME"`
	DBUserName                  string `mapstructure:"DB_USERNAME"`
	DBPassword                  string `mapstructure:"DB_PASSWORD"`
	SSLCertAuthorityCertificate string `mapstructure:"SSL_CERT_AUTH"`

	// jwt info
	JWTSecretKey       string `mapstructure:"JWT_SECRET_KEY"`
	VerifyTokenExpiry  int    `mapstructure:"VERIFY_TOKEN_EXPIRY"`
	AccessTokenExpiry  int    `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry int    `mapstructure:"REFRESH_TOKEN_EXPIRY"`
}

func InitLoadAppConf() *AppConfig {
	configPath := os.Getenv("CONFIG_PATH")
	env := os.Getenv("CONFIG_ENV") // env = dev | prod

	appConfig, err := LoadConfigFromFile(configPath, env)
	if err != nil {
		panic(err)
	}

	var result AppConfig
	customConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
	}

	decoder, err := mapstructure.NewDecoder(customConfig)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(appConfig)
	if err != nil {
		panic(err)
	}

	return &result
}

func LoadConfigFromFile(path string, configName string) (config interface{}, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return
}

func (a AppConfig) ToInt(target string) int {
	result, err := strconv.Atoi(target)
	if err != nil {
		panic(err)
	}

	return result
}
