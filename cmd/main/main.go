package main

import (
	"github.com/ronanzindev/echo-fly.io/cmd/app"
	"github.com/ronanzindev/echo-fly.io/pkg/routes"
)

func main() {
	app := app.NewApp(routes.NewRouter(), ":8080")
	app.StartServer()
}
