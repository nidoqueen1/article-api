package config

import (
	"github.com/spf13/viper"
)

// Initializes viper and loads the configuration
func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	_ = viper.BindEnv("db.url", "DATABASE_URL")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
