package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerAddress      string `mapstructure:"server_address"`
	GrpcAddress        string `mapstructure:"grpc_address"`
	BinanceAPIKey      string `mapstructure:"binance_api_key"`
	BinanceSecret      string `mapstructure:"binance_secret"`
	BinanceAPIEndpoint string `mapstructure:"binance_api_endpoint"`
	BybitAPIKey        string `mapstructure:"bybit_api_key"`
	BybitSecret        string `mapstructure:"bybit_secret"`
	BybitAPIEndpoint   string `mapstructure:"bybit_api_endpoint"`
	BybitWSEndpoint    string `mapstructure:"bybit_ws_endpoint"`
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
