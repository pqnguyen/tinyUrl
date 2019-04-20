package usecase

import (
	"tinyUrl/services/url"
	_urlRepo "tinyUrl/services/url/repository"
	_urlUCase "tinyUrl/services/url/usecase"
)

var UrlUCase url.UseCase

func InitUseCase() {
	urlRepo := _urlRepo.NewUrlRepository()

	UrlUCase = _urlUCase.NewUrlUsecase(urlRepo)
}
