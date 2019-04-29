package url

import "github.com/pqnguyen/tinyUrl/models"

type UseCase interface {
	CreateUrl(user *models.User, url string, duration uint) (*models.Url, error)
	CreateFreeUrl(url string) (*models.Url, error)
	GetRedirectUrl(hash string) (string, error)
	RecordStatistic(hash string)
}
