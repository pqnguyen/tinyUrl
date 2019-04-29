package auth

import "github.com/pqnguyen/tinyUrl/context"

type Provider interface {
	Login(context *context.Context) (string, error)
	Register(context *context.Context) (string, error)
}
