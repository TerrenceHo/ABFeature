package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadConfig(filepath string) (*viper.Viper, error) {
	fmt.Println("Configuration file:", filepath)
	v := viper.New()
	v.SetConfigFile(filepath)
	err := v.ReadInConfig()
	return v, err
}
