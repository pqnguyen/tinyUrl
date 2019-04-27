package auth_handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	. "tinyUrl/context"
	"tinyUrl/services/auth"
	"tinyUrl/services/auth/password"
	"tinyUrl/services/usecase"
	"tinyUrl/services/user"
)

type authHandler struct {
	auth.Provider
	UserUCase user.UseCase
}

func NewAuthHandler(route *gin.RouterGroup) {
	cfg := password.Config{
		UserUCase: usecase.UserUCase,
	}
	handler := &authHandler{
		Provider:  password.New(&cfg),
		UserUCase: usecase.UserUCase,
	}
	route.POST("/login", handler.Login)
	route.POST("/register", handler.Register)
}

func (handler *authHandler) Login(ctx *Context) {
	token, err := handler.Provider.Login(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}
	SendSuccess(ctx, token)
}

func (handler *authHandler) Register(ctx *Context) {
	token, err := handler.Provider.Register(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}
	SendSuccess(ctx, token)
}

func Authenticate(ctx *Context) {
	authorization := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authorization, "bear ") {
		SendForbidden(ctx)
		return
	}
	tokenString := strings.TrimPrefix(authorization, "bear ")
	token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.SignedString), nil
	})
	if err != nil {
		SendForbidden(ctx)
		return
	}
	_, ok := token.Claims.(*auth.Claims)
	if !ok || !token.Valid {
		SendForbidden(ctx)
		return
	}
	ctx.Next()
}
