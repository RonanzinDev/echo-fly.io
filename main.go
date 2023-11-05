package main

import (
	"github.com/ronanzindev/echo-fly.io/src/app"
	"github.com/ronanzindev/echo-fly.io/src/database"
	"github.com/ronanzindev/echo-fly.io/src/routes"
	"github.com/ronanzindev/echo-fly.io/src/utils"
)

func init() {
	utils.LoadEnv()
}
func main() {
	database.ConnectPostgresDB()
	// service.SeedUsers()
	app := app.NewApp(routes.NewRouter(), ":8080")
	app.StartServer()
}
