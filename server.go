package main

import (
	"fmt"

	"github.com/lbrulet/APINIT-GO/src/configs"
	"github.com/lbrulet/APINIT-GO/src/database"
	"github.com/lbrulet/APINIT-GO/src/mongodb"
	"github.com/lbrulet/APINIT-GO/src/routes"
)

func main() {
	configs.InitConfig()
	err := database.InitDB()
	if err != nil {
		fmt.Println(err)
	} else {
		db := mongodb.Connect()
		defer db.GetDatabase().Close()

		srv := routes.InitRouter()

		srv.Run(configs.Config.Port)
	}
}
