package url

import (
	"github.com/pqnguyen/tinyUrl/models"
	"time"
)

type Repository interface {
	CreateURL(user *models.User, originalURL string, expiryDuration time.Duration) (*models.Url, error)
	GetUrl(hash string) (*models.Url, error)
}
