package user

import "tinyUrl/models"

type Repository interface {
	GetFreeUser() *models.User
}
