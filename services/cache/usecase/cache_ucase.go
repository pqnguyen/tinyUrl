package usecase

import (
	"tinyUrl/services/cache"
)

type cacheUseCase struct {
	cacheRepo cache.Repository
}

func NewCacheUsecase(cache cache.Repository) cache.UseCase {
	return &cacheUseCase{
		cacheRepo: cache,
	}
}
