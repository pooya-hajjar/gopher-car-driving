package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	CarModel     string `mapstructure:"CAR_MODEL"`
	CarBrand     string `mapstructure:"CAR_BRAND"`
	DriverName   string `mapstructure:"DRIVER_NAME"`
	MaxSpeed     int    `mapstructure:"MAX_SPEED"`
	MaxRpm       int    `mapstructure:"MAX_RPM"`
	DangerousRpm int    `mapstructure:"DANGEROUS_RPM"`

	GasKey       string `mapstructure:"GAS_KEY"`
	BrakeKey     string `mapstructure:"BRAKE_KEY"`
	ClutchKey    string `mapstructure:"CLUTCH_KEY"`
	IGearKey     string `mapstructure:"INCREASE_GEAR_KEY"`
	DGearKey     string `mapstructure:"DECREASE_GEAR_KEY"`
	PowerKey     string `mapstructure:"POWER_KEY"`
	HandBrakeKey string `mapstructure:"HAND_BRAKE_KEY"`
	GearNKey     string `mapstructure:"GEAR_N_KEY"`
	Gear1Key     string `mapstructure:"GEAR_1_KEY"`
	Gear2Key     string `mapstructure:"GEAR_2_KEY"`
	Gear3Key     string `mapstructure:"GEAR_3_KEY"`
	Gear4Key     string `mapstructure:"GEAR_4_KEY"`
	Gear5Key     string `mapstructure:"GEAR_5_KEY"`
	GearRKey     string `mapstructure:"GEAR_R_KEY"`
}

// LoadConfig loads the configuration from the specified path using Viper,
// which includes reading from an .env file and environment variables.
// It returns a Config instance populated with the loaded configuration data.
func LoadConfig(path string) (config Config, err error) {
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
