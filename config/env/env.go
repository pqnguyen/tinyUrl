package env

import (
	"fmt"
	"github.com/pqnguyen/tinyUrl/types/enums"
	"github.com/spf13/viper"
	"log"
)

type MongoConfig struct {
	Host      string
	Port      string
	DBTinyUrl string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
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

func InitEnvironment(env enums.Environment) {
	fileConfig := fmt.Sprintf("config/env/%s_env.json", env)
	viper.SetConfigFile(fileConfig)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] got error while read config file: %v", err)
	}
	Config = &Configuration{
		MongoConfig: MongoConfig{
			Host:      viper.GetString("mongo.host"),
			Port:      viper.GetString("mongo.port"),
			DBTinyUrl: viper.GetString("mongo.db_tiny_url"),
		},
		RedisConfig: RedisConfig{
			Host:     viper.GetString("redis.host"),
			Port:     viper.GetString("redis.port"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		},
	}
}
