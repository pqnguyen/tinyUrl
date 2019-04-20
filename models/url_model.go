package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Url struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Hash           string             `bson:"hash,omitempty"`
	OriginalURL    string             `bson:"original_url,omitempty"`
	CreationDate   time.Time          `bson:"creation_date,omitempty"`
	ExpirationDate time.Time          `bson:"expiration_date,omitempty"`
	UserID         primitive.ObjectID `bson:"user_id,omitempty"`
}

type UrlRes struct {
	ID             primitive.ObjectID `json:"id"`
	Hash           string             `json:"hash"`
	OriginalURL    string             `json:"original_url"`
	CreationDate   time.Time          `json:"creation_date"`
	ExpirationDate time.Time          `json:"expiration_date"`
	UserID         primitive.ObjectID `json:"user_id"`
	TinyUrl        string             `json:"tiny_url"`
}

func (url *Url) ToRes(creator *User) *UrlRes {
	return &UrlRes{
		ID:             url.ID,
		Hash:           url.Hash,
		OriginalURL:    url.OriginalURL,
		CreationDate:   url.CreationDate,
		ExpirationDate: url.ExpirationDate,
		UserID:         url.UserID,
		TinyUrl:        url.Hash,
	}
}
