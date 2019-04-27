package usecase

import (
	"tinyUrl/models"
	"tinyUrl/services/user"
)

type userUseCase struct {
	userRepo user.Repository
}

func (user *userUseCase) Exists(email string) (*models.User, bool) {
	return user.userRepo.Exists(email)
}

func (user *userUseCase) Create(name, email string, password string) (*models.User, error) {
	return user.userRepo.Create(name, email, password)
}

func NewUserUsecase(user user.Repository) user.UseCase {
	return &userUseCase{
		userRepo: user,
	}
}
