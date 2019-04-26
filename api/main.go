package main

import (
	"tinyUrl/api/server"
	"tinyUrl/config/env"
	"tinyUrl/models"
	"tinyUrl/services/usecase"
	"tinyUrl/types/enums"
)

func main() {
	port := "8080"
	env.InitEnvironment(enums.Local)
	models.InitModels()
	usecase.InitUseCase()
	server.Serve(port)
}
