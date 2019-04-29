package context

import (
	"github.com/gin-gonic/gin"
	"github.com/pqnguyen/tinyUrl/models"
	"github.com/pqnguyen/tinyUrl/types/code"
	"github.com/pqnguyen/tinyUrl/types/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
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

func SendForbidden(ctx *Context) {
	res := response.SendForbidden(http.StatusForbidden)
	ctx.AbortWithStatusJSON(http.StatusForbidden, res)
}
