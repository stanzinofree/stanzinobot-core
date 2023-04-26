package initializer

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB       string `mapstructure:"DB"`
	Name     string `mapstructure:"NAME"`
	Main     int    `mapstructure:"MAIN"`
	Version  int    `mapstructure:"VERSION"`
	Release  int    `mapstructure:"REL"`
	Channel  string `mapstructure:"CHANNEL"`
	Codename string `mapstructure:"CODENAME"`
	Port     int    `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
