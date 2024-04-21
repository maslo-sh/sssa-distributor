package main

import (
	_ "github.com/SSSaaS/sssa-golang"
	"privileges-management/database"
	"privileges-management/server"
)

func main() {
	database.InitializeSchema()
	r := server.NewRouter()

	r.Run(":8080")
}
