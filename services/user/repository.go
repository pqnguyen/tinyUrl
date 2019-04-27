package user

import "tinyUrl/models"

type Repository interface {
	GetFreeUser() *models.User
	Exists(email string) (*models.User, bool)
	Create(name, email string, password string) (*models.User, error)
}
