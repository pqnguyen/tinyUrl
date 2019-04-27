package user

import "tinyUrl/models"

type UseCase interface {
	Exists(email string) (*models.User, bool)
	Create(name, email string, password string) (*models.User, error)
}
