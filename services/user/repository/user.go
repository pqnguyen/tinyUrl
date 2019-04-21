package repository

import (
	"time"
	"tinyUrl/models"
	"tinyUrl/services/user"
)

type userRepository struct {
}

const FreeUser = "free_pretty_user"
const DefaultCreateDate = 1546300800

func (u *userRepository) GetFreeUser() *models.User {
	return &models.User{
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
