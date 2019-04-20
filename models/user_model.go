package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	CreationDate time.Time          `bson:"creation_date"`
	LastLogin    time.Time          `bson:"last_login"`
}
