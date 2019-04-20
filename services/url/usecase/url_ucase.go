package usecase

import (
	"tinyUrl/models"
	"tinyUrl/services/url"
)

type urlUseCase struct {
	urlRepo url.Repository
}

func (u *urlUseCase) GetRedirectUrl(hash string) (*models.Url, error) {
	urlObj, err := u.urlRepo.GetUrl(hash)
	if err != nil {
		return &models.Url{}, err
	}
	return urlObj, nil
}

func (u *urlUseCase) CreateUrl(user *models.User, originalURL string, expiryDuration uint) (*models.Url, error) {
	urlObj, err := u.urlRepo.CreateURL(user, originalURL, expiryDuration)
	if err != nil {
		return &models.Url{}, err
	}
	return urlObj, nil
}

func NewUrlUsecase(url url.Repository) url.UseCase {
	return &urlUseCase{
		urlRepo: url,
	}
}
