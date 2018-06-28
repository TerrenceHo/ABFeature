package config

import "github.com/spf13/viper"

func ReadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath("./config")
	err := v.ReadInConfig()
	return v, err
}
