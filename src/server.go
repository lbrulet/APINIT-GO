package main

import (
	"fmt"

	"github.com/lbrulet/APINIT-GO/src/configs"
	"github.com/lbrulet/APINIT-GO/src/database"
	"github.com/lbrulet/APINIT-GO/src/routes"
)

func main() {
	configs.InitConfig()
	err := database.InitDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	srv := routes.InitRouter()
	if err := srv.Run(configs.Config.Port); err != nil {
		fmt.Println(err)
	}
}
