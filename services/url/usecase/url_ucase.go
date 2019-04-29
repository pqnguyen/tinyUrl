package usecase

import (
	"context"
	"fmt"
	"github.com/pqnguyen/tinyUrl/config/constant"
	"github.com/pqnguyen/tinyUrl/models"
	"github.com/pqnguyen/tinyUrl/services/cache"
	"github.com/pqnguyen/tinyUrl/services/url"
	"github.com/pqnguyen/tinyUrl/services/user"
	"github.com/pqnguyen/tinyUrl/types/code"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

type urlUseCase struct {
	urlRepo   url.Repository
	userRepo  user.Repository
	cacheRepo cache.Repository
}

func (u *urlUseCase) RecordStatistic(hash string) {
	now := time.Now()
	nWeek := now.Day() % 7
	month := fmt.Sprintf("%02d/%d", now.Month(), now.Year())
	upsert := true
	_, _ = models.DB().Collection(models.Col.Statistic).UpdateOne(
		context.Background(),
		bson.M{
			"hash":  hash,
			"month": month,
		}, bson.M{
			"$inc": bson.M{strconv.Itoa(nWeek): 1},
		}, &options.UpdateOptions{Upsert: &upsert})
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
		if leftExpiryDuration < 0 {
			return "", code.ErrTinyUrlExpired
		}
		if leftExpiryDuration > constant.DefaultLeftTimeCache {
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
