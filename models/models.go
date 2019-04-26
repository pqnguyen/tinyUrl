package models

import (
	"context"
	_redis "github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
	"tinyUrl/config/env"
	"tinyUrl/types/enums"
)

type Collection struct {
	Url       string
	User      string
	Statistic string
}

var Col Collection

var database = make(map[enums.DatabaseName]*mongo.Database)
var redis *_redis.Client

func InitModels() {
	initDatabase()
	initCache()
}

func initCache() {
	redis = _redis.NewClient(&_redis.Options{
		Addr:     env.Config.RedisConfig.GetAddr(),
		Password: env.Config.RedisConfig.Password,
		DB:       env.Config.RedisConfig.DB, // use default DB
	})

	pong, err := redis.Ping().Result()
	if err != nil {
		log.Printf("[ERROR] got error when connecting to redis server: %s", err)
	}
	log.Printf("[INFO] connect successful to redis server: %s", pong)
}

func initDatabase() {
	client, err := mongo.NewClient(options.Client().ApplyURI(env.Config.MongoConfig.GetURI()))
	if err != nil {
		log.Fatalf("[ERROR] got error when creating mongodb client: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("[ERROR] got error when ping to mongodb server: %v", err)
	}

	log.Printf("[INFO] connect successful to mongodb server")

	database[env.Config.MongoConfig.DBTinyUrl] = client.Database(env.Config.MongoConfig.DBTinyUrl)
	Col = Collection{
		Url:       "url",
		User:      "user",
		Statistic: "statistic",
	}
}

func DB() *mongo.Database {
	dbUrl, ok := database[env.Config.MongoConfig.DBTinyUrl]
	if !ok {
		log.Fatalf("[ERROR] init database before get one")
	}
	return dbUrl
}

func ClearDB() error {
	dbUrl, ok := database[env.Config.MongoConfig.DBTinyUrl]
	if !ok {
		log.Fatalf("[ERROR] init database before get one")
	}
	return dbUrl.Drop(context.Background())
}

func ClearCache() error {
	statusCmd := Redis().FlushAll()
	return statusCmd.Err()
}

func Redis() *_redis.Client {
	if redis == nil {
		initCache()
	}
	return redis
}
