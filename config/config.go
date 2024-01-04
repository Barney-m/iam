package config

import (
	"github.com/spf13/viper"
)

func ReadConfig(env string) error {
	viper.AddConfigPath("./config")
	viper.SetConfigName(env)    // Register config file name (no extension)
	viper.SetConfigType("yaml") // Look for specific type
	return viper.ReadInConfig()
}
