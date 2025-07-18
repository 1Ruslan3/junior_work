package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
	}
	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
}

func LoadConfig() Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(getConfigPath())

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	return cfg
}

func getConfigPath() string {
	viper.AutomaticEnv()
	path := viper.GetString("CONFIG_PATH")
	if path == "" {
		path = "./config/config.yaml"
	}
	return path
}
