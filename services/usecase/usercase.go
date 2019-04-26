package usecase

import (
	"tinyUrl/services/cache"
	_cacheRepo "tinyUrl/services/cache/repository"
	_cacheUcase "tinyUrl/services/cache/usecase"
	"tinyUrl/services/url"
	_urlRepo "tinyUrl/services/url/repository"
	_urlUCase "tinyUrl/services/url/usecase"
	_userRepo "tinyUrl/services/user/repository"
)

var UrlUCase url.UseCase
var CacheUCase cache.UseCase

func InitUseCase() {
	urlRepo := _urlRepo.NewUrlRepository()
	userRepo := _userRepo.NewUserRepository()
	cacheRepo := _cacheRepo.NewCacheRepository()

	UrlUCase = _urlUCase.NewUrlUsecase(urlRepo, userRepo, cacheRepo)
	CacheUCase = _cacheUcase.NewCacheUsecase(cacheRepo)
}
