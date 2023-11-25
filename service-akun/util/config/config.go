package util

import (
	"github.com/spf13/viper"
)


type Config struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	RedisServiceAddress string `mapstructure:"REDIS_SERVICE_ADDRESS"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"` 
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}