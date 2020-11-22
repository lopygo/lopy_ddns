package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App     App
	Drivers []Driver
}

func LoadFromFile() (Config, error) {
	c := Config{}
	viperOb := viper.New()
	viperOb.SetConfigName("config")
	viperOb.AddConfigPath(".")
	viperOb.SetConfigType("yaml")
	err := viperOb.ReadInConfig()
	if err != nil {
		return c, err
	}

	err = viperOb.Unmarshal(&c)
	return c, err
}
