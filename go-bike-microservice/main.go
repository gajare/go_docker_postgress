package main

import (
	"go-bike-microservice/config"
	"go-bike-microservice/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run()
}
