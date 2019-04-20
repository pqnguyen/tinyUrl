package env

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type MongoConfig struct {
	Host string
	Port string
}

func (mg *MongoConfig) GetURI() string {
	return fmt.Sprintf("mongodb://%s:%s", mg.Host, mg.Port)
}

type Configuration struct {
	MongoConfig
}

var Config *Configuration

func InitEnvironment() {
	viper.SetConfigFile("config/env/local_env.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] got error while read config file: %v", err)
	}
	Config = &Configuration{
		MongoConfig: MongoConfig{
			Host: viper.GetString("mongo.host"),
			Port: viper.GetString("mongo.port"),
		},
	}
}
