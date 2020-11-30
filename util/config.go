package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	WebhookURL string `mapstructure:"WEBHOOK_URL"`
	Interval   int    `mapstructure:"INTERVAL"`
}

func NewConfig() *Config {
	return &Config{
		WebhookURL: "",
		Interval:   30,
	}
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := NewConfig()
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil

}
