package usecase

import (
	"tinyUrl/services/user"
)

type userUseCase struct {
	userRepo user.Repository
}

func NewUserUsecase(user user.Repository) user.UseCase {
	return &userUseCase{
		userRepo: user,
	}
}
