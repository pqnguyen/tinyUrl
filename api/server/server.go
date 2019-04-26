package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"tinyUrl/api/server/url_handler"
)

func CreateServer() *gin.Engine {
	router := gin.Default()

	urlGroup := router.Group("")
	url_handler.NewUrlHandler(urlGroup)
	return router
}

func Serve(port string) {
	router := CreateServer()
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("[ERROR] api server was shutdown: %v", err)
	}
}
