package password

import (
	"github.com/pqnguyen/tinyUrl/context"
	"github.com/pqnguyen/tinyUrl/models/views"
	"github.com/pqnguyen/tinyUrl/services/user"
	"github.com/pqnguyen/tinyUrl/types/code"
)

type Config struct {
	UserUCase user.UseCase
	Authority PasswordAuthority
}

type PasswordProvider struct {
	*Config
}

func New(cfg *Config) *PasswordProvider {
	if cfg == nil {
		cfg = &Config{}
	}
	if cfg.Authority == nil {
		cfg.Authority = &DefaultPasswordAuthority{}
	}
	return &PasswordProvider{cfg}
}

func (password *PasswordProvider) Login(context *context.Context) (string, error) {
	var data views.PasswordLoginView
	if err := context.ShouldBindJSON(&data); err != nil {
		return "", err
	}
	var config = password.Config
	u, exists := config.UserUCase.Exists(data.Email)
	if !exists {
		return "", code.ErrUserNotExists
	}
	if !config.Authority.VerifyUser(u, data.Password) {
		return "", code.ErrUserNotExists
	}
	token := config.Authority.GenerateJWT(u)
	return token, nil
}

func (password *PasswordProvider) Register(context *context.Context) (string, error) {
	var data views.PasswordRegisterView
	if err := context.ShouldBindJSON(&data); err != nil {
		return "", err
	}
	u, exists := password.Config.UserUCase.Exists(data.Email)
	if exists {
		return "", code.ErrUserAlreadyExists
	}

	var config = password.Config
	pwd := config.Authority.HashPassword(data.Password)
	u, err := config.UserUCase.Create(data.Name, data.Email, pwd)
	if err != nil {
		return "", err
	}
	token := config.Authority.GenerateJWT(u)
	return token, nil
}
