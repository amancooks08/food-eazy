package config

import "github.com/spf13/viper"

func Load() {
	viper.SetConfigFile("./application.yaml") // name of config file
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}