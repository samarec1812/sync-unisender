package configs

import "github.com/spf13/viper"

func InitConfig() error {
	viper.AddConfigPath("app/auth/configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
