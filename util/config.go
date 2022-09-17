package util

import "github.com/spf13/viper"

type Config struct {
	DbDriver string `mapstructure:"DB_DRIVER"`
	DbSource string `mapstructure:"DB_SOURCE"`
	Addr     string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path, filename string) (*Config, error) {
	var config Config

	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
