package context

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"tinyUrl/models"
	"tinyUrl/types/code"
	"tinyUrl/types/response"
)

type Context = gin.Context

func GetUser(ctx *Context) *models.User {
	return &models.User{ID: primitive.ObjectID{'1', '2', '3'}}
}

func SendError(ctx *Context, error error) {
	res := response.SendError(http.StatusBadRequest, error, "")
	ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func SendSuccess(ctx *Context, message interface{}) {
	res := response.SendSuccess(http.StatusOK, code.Success, message)
	ctx.JSON(http.StatusOK, res)
}
