package configs

import "github.com/spf13/viper"

type conf struct {
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	Zipkin        string `mapstructure:"ZIPKIN_URL"`
}

func LoadConfig(path string) *conf {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigFile("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
