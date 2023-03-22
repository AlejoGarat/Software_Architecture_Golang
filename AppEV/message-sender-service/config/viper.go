package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Setup() {
	viper.SetConfigName("config")

	viper.AddConfigPath("../")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var configuration Configurations

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}
