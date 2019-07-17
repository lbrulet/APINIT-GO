package main

import (
	"github.com/lbrulet/APINIT-GO/mongodb"
	"github.com/lbrulet/APINIT-GO/routes"
)

func main() {
	db := mongodb.Connect()
	defer db.GetDatabase().Close()

	srv := routes.StartServer()

	srv.Run()
}
