package main

import (
	"tinyUrl/api/server"
	"tinyUrl/config/env"
	"tinyUrl/models"
)

func main() {
	port := "8080"
	env.InitEnvironment()
	models.InitModels()
	server.Serve(port)
}
