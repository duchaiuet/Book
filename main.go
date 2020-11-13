package main

import (
	"Book/Api"
	"Book/Database"
)

// @title Swagger Book project API
// @version 2.0
// @description This is list api for vvt product project
// @localhost:1234
// @BasePath /api/v1/book
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	Api.HandleHttpServer(Database.HostPort)
}
