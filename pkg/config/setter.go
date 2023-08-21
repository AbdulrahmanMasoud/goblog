package config

import (
	"github.com/spf13/viper"
	"log"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error Reading the config")
	}

	err := viper.Unmarshal(&configurations)

	if err != nil {
		log.Fatal("Unable to decode into struct")
	}
}
