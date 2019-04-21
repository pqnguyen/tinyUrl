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

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func (mg *MongoConfig) GetURI() string {
	return fmt.Sprintf("mongodb://%s:%s", mg.Host, mg.Port)
}

type Configuration struct {
	MongoConfig
	RedisConfig
}

func (redis RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", redis.Host, redis.Port)
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
		RedisConfig: RedisConfig{
			Host:     viper.GetString("redis.host"),
			Port:     viper.GetString("redis.port"),
			Password: viper.GetString("redis.password"),
		},
	}
}
