package url

import (
	"time"
	"tinyUrl/models"
)

type Repository interface {
	CreateURL(user *models.User, originalURL string, expiryDuration time.Duration) (*models.Url, error)
	GetUrl(hash string) (*models.Url, error)
}
