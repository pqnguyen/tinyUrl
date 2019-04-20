package repository

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"tinyUrl/config/constant"
	. "tinyUrl/models"
	"tinyUrl/services/url"
)

type urlRepository struct {
}

func (u *urlRepository) GetUrl(hash string) (*Url, error) {
	filter := Url{Hash: hash}
	b, err := bson.Marshal(&filter)
	if err != nil {
		return &Url{}, err
	}
	urlObj := Url{}
	err = DB().Collection(Col.Url).FindOne(context.Background(), b).Decode(&urlObj)
	if err != nil {
		return &Url{}, err
	}
	return &urlObj, nil
}

func (u *urlRepository) CreateURL(user *User, originalURL string, expiryDuration uint) (*Url, error) {
	str := fmt.Sprintf("%s%d%d", originalURL, user.ID, time.Now().Unix())
	byteHash := md5.Sum([]byte(str))
	hash := base64.StdEncoding.EncodeToString(byteHash[:])
	urlObj := Url{
		Hash:           string(hash[:constant.DefaultHashLength]),
		OriginalURL:    originalURL,
		CreationDate:   time.Now(),
		ExpirationDate: time.Now().Add(time.Duration(expiryDuration)),
		UserID:         user.ID,
	}
	b, err := bson.Marshal(&urlObj)
	if err != nil {
		return &Url{}, err
	}
	result, err := DB().Collection(Col.Url).InsertOne(context.Background(), b)
	if err != nil {
		return &Url{}, err
	}
	urlObj.ID = result.InsertedID.(primitive.ObjectID)
	return &urlObj, nil
}

func NewUrlRepository() url.Repository {
	return &urlRepository{}
}
