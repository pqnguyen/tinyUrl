package url_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "tinyUrl/context"
	"tinyUrl/models/views"
	"tinyUrl/services/url"
	"tinyUrl/services/usecase"
	"tinyUrl/types/code"
)

type urlHandler struct {
	UrlUCase url.UseCase
}

func NewUrlHandler(route *gin.RouterGroup) {
	handler := &urlHandler{
		UrlUCase: usecase.UrlUCase,
	}
	route.POST("/create_url", handler.createUrlAPI)
	route.GET("/:hash", handler.redirectUrlAPI)
}

func (handler *urlHandler) createUrlAPI(ctx *Context) {
	var view views.CreateURLView
	if err := ctx.ShouldBindJSON(&view); err != nil {
		SendError(ctx, code.ErrBadRequestedData)
		return
	}
	user := GetUser(ctx)
	urlObj, err := handler.UrlUCase.CreateUrl(user, view.Url, view.ExpiryDuration)
	if err != nil {
		SendError(ctx, err)
		return
	}
	SendSuccess(ctx, urlObj.ToRes(user))
}

func (handler *urlHandler) redirectUrlAPI(ctx *gin.Context) {
	var view views.RedirectUrlView
	if err := ctx.ShouldBindUri(&view); err != nil {
		SendError(ctx, code.ErrBadRequestedData)
		return
	}
	urlObj, err := handler.UrlUCase.GetRedirectUrl(view.Hash)
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Redirect(http.StatusPermanentRedirect, urlObj.OriginalURL)
}
