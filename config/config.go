package config

import (
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	//viper.AddConfigPath("/home/gstrackme/gst-reports-api-go/config")
	//viper.AddConfigPath("$HOME/projects/gst-reports-api/config")
	//viper.AddConfigPath("D:/GO/gst-reports-api_users_feature/config")
	viper.AddConfigPath("C:/Clean_Survey")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		println(err)
	}
}
