package repository

import (
	"fmt"
	"time"
	"tinyUrl/models"
	"tinyUrl/services/cache"
	"tinyUrl/types/enums"
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
