package usecase

import (
	"time"
	"tinyUrl/models"
	"tinyUrl/services/cache"
	"tinyUrl/services/url"
	"tinyUrl/services/user"
)

type urlUseCase struct {
	urlRepo   url.Repository
	userRepo  user.Repository
	cacheRepo cache.Repository
}

func (u *urlUseCase) CreateFreeUrl(originalURL string) (*models.Url, error) {
	freeUser := u.userRepo.GetFreeUser()
	expiryDuration := 24 * time.Hour
	urlObj, err := u.urlRepo.CreateURL(freeUser, originalURL, expiryDuration)
	if err != nil {
		return &models.Url{}, err
	}
	return urlObj, nil
}

func (u *urlUseCase) GetRedirectUrl(hash string) (string, error) {
	originalUrl, exists := u.cacheRepo.GetOriginalUrl(hash)
	if exists != true {
		urlObj, err := u.urlRepo.GetUrl(hash)
		if err != nil {
			return "", err
		}

		leftExpiryDuration := urlObj.ExpirationDate.Unix() - time.Now().UTC().Unix()
		if leftExpiryDuration > 0 {
			u.cacheRepo.SetUrl(hash, urlObj.OriginalURL, time.Duration(leftExpiryDuration)*time.Second)
		}
		originalUrl = urlObj.OriginalURL
	}
	return originalUrl, nil
}

func (u *urlUseCase) CreateUrl(user *models.User, originalURL string, expiryDuration uint) (*models.Url, error) {
	urlObj, err := u.urlRepo.CreateURL(user, originalURL, time.Duration(expiryDuration))
	if err != nil {
		return &models.Url{}, err
	}
	return urlObj, nil
}

func NewUrlUsecase(url url.Repository, user user.Repository, cache cache.Repository) url.UseCase {
	return &urlUseCase{
		urlRepo:   url,
		userRepo:  user,
		cacheRepo: cache,
	}
}
