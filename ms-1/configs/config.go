package configs

import "github.com/spf13/viper"

type conf struct {
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	UrlCollectior string `mapstructure:"URL_COLLECTOR"`
}

func LoadConfig(path string) *conf {
	var cfg *conf
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
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
