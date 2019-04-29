package repository

import (
	"fmt"
	"github.com/pqnguyen/tinyUrl/models"
	"github.com/pqnguyen/tinyUrl/services/cache"
	"github.com/pqnguyen/tinyUrl/types/enums"
	"time"
)

type cacheRepository struct {
}

func generateKey(ns enums.RedisNameSpace, key string) string {
	return fmt.Sprintf("%s:%s", ns, key)
}

func (c *cacheRepository) GetOriginalUrl(hash string) (string, bool) {
	key := generateKey(enums.UrlNS, hash)
	originalUrl, err := models.Redis().Get(key).Result()
	if err != nil {
		return "", false
	}
	return originalUrl, true
}

func (c *cacheRepository) SetUrl(hash string, originalUrl string, d time.Duration) {
	key := generateKey(enums.UrlNS, hash)
	models.Redis().Set(key, originalUrl, d)
}

func NewCacheRepository() cache.Repository {
	return &cacheRepository{}
}
