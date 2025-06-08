package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string `yaml:"port"`
	Host        string `yaml:"host"`
	DebugServer bool   `yaml:"debug_server"`
}

func GetConfig() Config {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	cfg := Config{
		Port:        viper.GetString("port"),
		Host:        viper.GetString("host"),
		DebugServer: viper.GetBool("debug_server"),
	}

	return cfg
}
