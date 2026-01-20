package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppPort                int    `mapstructure:"APP_PORT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBName                 string `mapstructure:"DB_NAME"`
	DBUsername             string `mapstructure:"DB_USERNAME"`
	DBPassword             string `mapstructure:"DB_PASSWORD"`
	DBPort                 int    `mapstructure:"DB_PORT"`
	JwtAccessSecret        string `mapstructure:"JWT_ACCESS_SECRET"`
	JwtRefreshSecret       string `mapstructure:"JWT_REFRESH_SECRET"`
	JwtAccessTokenExpired  int    `mapstructure:"JWT_ACCESS_TOKEN_EXPIRED"`
	JwtRefreshTokenExpired int    `mapstructure:"JWT_REFRESH_TOKEN_EXPIRED"`
}

var cfg Env

func LoadEnv() *Env {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("APP_PORT", 8080)
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "buytun")
	viper.SetDefault("DB_USERNAME", "root")
	viper.SetDefault("DB_PASSWORD", "")
	viper.SetDefault("DB_PORT", 5432)

	if err := viper.ReadInConfig(); err != nil {
		if errors.Is(err, viper.ConfigFileNotFoundError{}) {
			log.Println("no config file found, using default.")
		} else {
			log.Fatalf("fatal reading config %v", err)
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to unmarshal config into struct: %s \n", err)
	}

	return &cfg
}

func GetConfig() *Env {
	return &cfg
}
