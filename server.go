package main

import (
	"github.com/lbrulet/APINIT-GO/configs"
	"github.com/lbrulet/APINIT-GO/mongodb"
	"github.com/lbrulet/APINIT-GO/routes"
)

func main() {
	configs.InitConfig()
	db := mongodb.Connect()
	defer db.GetDatabase().Close()

	srv := routes.InitRouter()

	srv.Run(configs.Config.Port)
}
