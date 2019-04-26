package usecase

import (
	"tinyUrl/services/cache"
)

type cacheUseCase struct {
	cacheRepo cache.Repository
}

func (c *cacheUseCase) GetOriginalUrl(hash string) (string, bool) {
	return c.cacheRepo.GetOriginalUrl(hash)
}

func NewCacheUsecase(cache cache.Repository) cache.UseCase {
	return &cacheUseCase{
		cacheRepo: cache,
	}
}
