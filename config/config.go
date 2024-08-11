package config

import (
	"github.com/pooya-hajjar/terminal-car-driving/models"
	"github.com/spf13/viper"
)

// LoadConfig loads the configuration from the specified path using Viper,
// which includes reading from an .env file and environment variables.
// It returns a models.Config instance populated with the loaded configuration data.
func LoadConfig(path string) (config models.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("setting")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
