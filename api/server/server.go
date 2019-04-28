package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tinyUrl/api/server/auth_handler"
	"tinyUrl/api/server/url_handler"
	"tinyUrl/context"
)

func cors(ctx *context.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if ctx.Request.Method == "OPTIONS" {
		ctx.Status(http.StatusOK)
		return
	}
	ctx.Next()
}

func CreateServer() *gin.Engine {
	router := gin.Default()

	router.Use(cors)

	urlGroup := router.Group("")
	url_handler.NewUrlHandler(urlGroup)

	authGroup := router.Group("/auth")
	auth_handler.NewAuthHandler(authGroup)
	return router
}

func Serve(port string) {
	router := CreateServer()
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("[ERROR] api server was shutdown: %v", err)
	}
}
