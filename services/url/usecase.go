package url

import "tinyUrl/models"

type UseCase interface {
	CreateUrl(user *models.User, s string, u uint) (*models.Url, error)
	GetRedirectUrl(hash string) (*models.Url, error)
}
