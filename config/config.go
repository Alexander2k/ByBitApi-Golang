package config

import "log"
import "github.com/spf13/viper"

type (
	Config struct {
		Api Api `mapstructure:"api,omitempty"`
	}

	Api struct {
		ApiSite   string `mapstructure:"api_site,omitempty"`
		ApiKey    string `mapstructure:"api_key,omitempty"`
		ApiSecret string `mapstructure:"api_secret,omitempty"`
	}
)

func NewConfig() (c *Config, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Cant read config %s", err)
		return
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Printf("Cant unmarshal config %s", err)
		return
	}
	return
}
