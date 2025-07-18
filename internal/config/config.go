package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	}
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
	}
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")
	_ = viper.ReadInConfig() // YAML не обязателен

	viper.AutomaticEnv() // .env-переменные

	// .env-файл (опционально)
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		if err := viper.MergeInConfig(); err != nil {
			log.Printf("Warning: failed to load .env: %v", err)
		}
	}

	cfg := &Config{}

	err := viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("Error decoding config: %v", err)
	}

	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
	return cfg
}
