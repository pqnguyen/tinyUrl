package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"tinyUrl/api/server/url_handler"
	"tinyUrl/services/usecase"
)

func Serve(port string) {
	usecase.InitUseCase()
	router := gin.Default()

	urlGroup := router.Group("")
	url_handler.NewUrlHandler(urlGroup)


	if err := router.Run(":" + port); err != nil {
		log.Fatalf("[ERROR] api server was shutdown: %v", err)
	}
}
