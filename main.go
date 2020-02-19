package main

import (
	"github.com/blogster/database"
	"github.com/blogster/routers"
)

func main() {
	database.Migrate()
	database.Apply()
	app := routers.RegisterRoutes()
	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
