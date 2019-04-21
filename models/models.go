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
)

type DatabaseName = string

const (
	DBTinyUrl = "tiny_url"
)

type Collection struct {
	Url  string
	User string
}

var Col Collection

var database = make(map[DatabaseName]*mongo.Database)
var redis *_redis.Client

func InitModels() {
	initDatabase()
	initCache()
}

func initCache() {
	redis = _redis.NewClient(&_redis.Options{
		Addr:     env.Config.RedisConfig.GetAddr(),
		Password: env.Config.RedisConfig.Password,
		DB:       0, // use default DB
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

	database[DBTinyUrl] = client.Database(DBTinyUrl)
	Col = Collection{
		Url:  "url",
		User: "user",
	}
}

func DB() *mongo.Database {
	dbUrl, ok := database[DBTinyUrl]
	if !ok {
		log.Fatalf("[ERROR] init database before get one")
	}
	return dbUrl
}

func Redis() *_redis.Client {
	if redis == nil {
		initCache()
	}
	return redis
}
