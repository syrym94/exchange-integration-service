package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerAddress      string
	BinanceAPIKey      string
	BinanceSecret      string
	BinanceAPIEndpoint string
	BybitAPIKey        string
	BybitSecret        string
	BybitAPIEndpoint   string
	OkexAPIKey         string
	OkexSecret         string
	OkexAPIEndpoint    string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Looks for config.yaml in the root directory

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into config struct: %s", err)
	}

	return &cfg
}
