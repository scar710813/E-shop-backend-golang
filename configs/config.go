package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var configuration conf

type conf struct {
	DatabaseDriver   string `mapstructure:"DATABASE_DRIVER"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`

	ServerPort string `mapstructure:"SERVER_PORT"`

	JwtSecret    string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	JwtTokenAuth *jwtauth.JWTAuth
}

func LoadConfigurations(configPath string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(configPath)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	readConfigError := viper.ReadInConfig()
	if readConfigError != nil {
		return nil, readConfigError
	}

	unmarshalConfigError := viper.Unmarshal(&configuration)
	if unmarshalConfigError != nil {
		return nil, unmarshalConfigError
	}

	configuration.JwtTokenAuth = jwtauth.New("HS256", []byte(configuration.JwtSecret), nil)

	return &configuration, nil
}
