package usecase

import (
	_cacheRepo "tinyUrl/services/cache/repository"
	"tinyUrl/services/url"
	_urlRepo "tinyUrl/services/url/repository"
	_urlUCase "tinyUrl/services/url/usecase"
	_userRepo "tinyUrl/services/user/repository"
)

var UrlUCase url.UseCase

func InitUseCase() {
	urlRepo := _urlRepo.NewUrlRepository()
	userRepo := _userRepo.NewUserRepository()
	cacheRepo := _cacheRepo.NewCacheRepository()

	UrlUCase = _urlUCase.NewUrlUsecase(urlRepo, userRepo, cacheRepo)
}
