package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     int
	}
	App struct {
		Name string
		Port string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
