package appconfig

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	TokenKey      string        `mapstructure:"TOKEN_KEY"`
	TokenDuration time.Duration `mapstructure:"TOKEN_DURATION"`
	Address       string        `mapstructure:"ADDRESS"`
}

func Load(path string) (config *Config, err error) {

	viper.SetConfigFile(path)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
