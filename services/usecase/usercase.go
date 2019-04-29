package usecase

import (
	"github.com/pqnguyen/tinyUrl/services/cache"
	_cacheRepo "github.com/pqnguyen/tinyUrl/services/cache/repository"
	_cacheUcase "github.com/pqnguyen/tinyUrl/services/cache/usecase"
	"github.com/pqnguyen/tinyUrl/services/url"
	_urlRepo "github.com/pqnguyen/tinyUrl/services/url/repository"
	_urlUCase "github.com/pqnguyen/tinyUrl/services/url/usecase"
	"github.com/pqnguyen/tinyUrl/services/user"
	_userRepo "github.com/pqnguyen/tinyUrl/services/user/repository"
	_userUcase "github.com/pqnguyen/tinyUrl/services/user/usecase"
)

var UrlUCase url.UseCase
var CacheUCase cache.UseCase
var UserUCase user.UseCase

func InitUseCase() {
	urlRepo := _urlRepo.NewUrlRepository()
	userRepo := _userRepo.NewUserRepository()
	cacheRepo := _cacheRepo.NewCacheRepository()

	UrlUCase = _urlUCase.NewUrlUsecase(urlRepo, userRepo, cacheRepo)
	CacheUCase = _cacheUcase.NewCacheUsecase(cacheRepo)
	UserUCase = _userUcase.NewUserUsecase(userRepo)
}
