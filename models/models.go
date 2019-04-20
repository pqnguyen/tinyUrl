package models

import (
	"context"
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

func InitModels() {
	client, err := mongo.NewClient(options.Client().ApplyURI(env.Config.GetURI()))
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
