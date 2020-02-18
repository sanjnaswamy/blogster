package main

import (
	"github.com/blogster/routers"
)

func main() {
	app := routers.RegisterRoutes()
	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
