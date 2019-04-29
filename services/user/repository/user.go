package repository

import (
	"context"
	. "github.com/pqnguyen/tinyUrl/models"
	"github.com/pqnguyen/tinyUrl/services/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type userRepository struct {
}

func (u *userRepository) Exists(email string) (*User, bool) {
	result := DB().Collection(Col.User).FindOne(context.Background(), primitive.M{
		"email": email,
	})
	var usr User
	if err := result.Decode(&usr); err != nil {
		return &User{}, false
	}
	return &usr, true
}

func (u *userRepository) Create(name, email string, password string) (*User, error) {
	usr := User{
		Name:         name,
		Email:        email,
		Password:     password,
		CreationDate: time.Now(),
		LastLogin:    time.Now(),
	}
	buf, err := bson.Marshal(&usr)
	if err != nil {
		return &User{}, err
	}
	res, err := DB().Collection(Col.User).InsertOne(context.Background(), buf)
	if err != nil {
		return &User{}, err
	}
	usr.ID = res.InsertedID.(primitive.ObjectID)
	return &usr, nil
}

const FreeUser = "free_pretty_user"
const DefaultCreateDate = 1546300800

func (u *userRepository) GetFreeUser() *User {
	return &User{
		ID:           [12]byte{},
		Name:         FreeUser,
		Email:        "",
		Password:     "",
		CreationDate: time.Unix(DefaultCreateDate, 0),
		LastLogin:    time.Unix(DefaultCreateDate, 0),
	}
}

func NewUserRepository() user.Repository {
	return &userRepository{}
}
