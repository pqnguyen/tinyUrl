package url

import (
	"tinyUrl/models"
)

type Repository interface {
	CreateURL(user *models.User, originalURL string, expiryDuration uint) (*models.Url, error)
	GetUrl(hash string) (*models.Url, error)
}
