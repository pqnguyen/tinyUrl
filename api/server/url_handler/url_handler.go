package url_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tinyUrl/api/server/auth"
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

	route.POST("/create_free_url", handler.createFreeUrlAPI)
	route.GET("/:hash", handler.redirectUrlAPI)
	route.Use(auth.Authenticate)
	route.POST("/create_url", handler.createUrlAPI)
}

func (handler *urlHandler) createFreeUrlAPI(ctx *Context) {
	var view views.CreateFreeURLView
	if err := ctx.ShouldBindJSON(&view); err != nil {
		SendError(ctx, code.ErrBadRequestedData)
		return
	}
	urlObj, err := handler.UrlUCase.CreateFreeUrl(view.Url)
	if err != nil {
		SendError(ctx, err)
		return
	}
	SendSuccess(ctx, urlObj.ToRes())
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
	SendSuccess(ctx, urlObj.ToRes())
}

func (handler *urlHandler) redirectUrlAPI(ctx *gin.Context) {
	var view views.RedirectUrlView
	if err := ctx.ShouldBindUri(&view); err != nil {
		SendError(ctx, code.ErrBadRequestedData)
		return
	}
	originalURL, err := handler.UrlUCase.GetRedirectUrl(view.Hash)
	if err != nil {
		SendError(ctx, err)
		return
	}
	go func() {
		handler.UrlUCase.RecordStatistic(view.Hash)
	}()
	ctx.Redirect(http.StatusPermanentRedirect, originalURL)
}
