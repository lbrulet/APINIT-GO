package main

import (
	"github.com/lbrulet/APINIT-GO/configs"
	"github.com/lbrulet/APINIT-GO/mongodb"
	"github.com/lbrulet/APINIT-GO/routes"
)

// @title APINIT-GO
// @version 1.0
// @description This is a sample golang server.

// @contact.name Luc Brulet
// @contact.url https://www.linkedin.com/in/luc-brulet/
// @contact.email luc.brulet@epitech.eu

// @host localhost:8080
// @BasePath /api/

// @securityDefinitions.basic BasicAuth
func main() {
	configs.InitConfig()
	db := mongodb.Connect()
	defer db.GetDatabase().Close()

	srv := routes.InitRouter()

	srv.Run(configs.Config.Port)
}
