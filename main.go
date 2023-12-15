package main

import (
	"implement_middleware/config"
	"implement_middleware/routes"
)

func main() {
	e := routes.InitRoutes()

	config.InitDatabase()

	e.Logger.Fatal(e.Start(":8000"))

}
