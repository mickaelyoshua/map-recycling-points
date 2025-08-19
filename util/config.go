package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	GeocodingApiKey string `mapstructure:"GEOCODING_API_KEY"`
	AddressesDataPath string `mapstructure:"ADDRESSES_DATA_PATH"`
	FinalDataPath string `mapstructure:"FINAL_DATA_PATH"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv() // check if env variables match the existing keys

	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("Error reading env file: %v\n", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, fmt.Errorf("Error parsing to struct: %v\n", err)
	}

	return config, nil
}
