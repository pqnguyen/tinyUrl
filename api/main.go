package main

import (
	"github.com/pqnguyen/tinyUrl/api/server"
	"github.com/pqnguyen/tinyUrl/config/env"
	"github.com/pqnguyen/tinyUrl/models"
	"github.com/pqnguyen/tinyUrl/services/usecase"
	"github.com/pqnguyen/tinyUrl/types/enums"
)

func main() {
	port := "8080"
	env.InitEnvironment(enums.Docker)
	models.InitModels()
	usecase.InitUseCase()
	server.Serve(port)
}
