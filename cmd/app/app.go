package app

import "github.com/labstack/echo/v4"

type App struct {
	Echo *echo.Echo
	Port string
}

func NewApp(e *echo.Echo, port string) *App {
	return &App{e, port}
}
func (app *App) StartServer() {
	app.Echo.Logger.Fatal(app.Echo.Start(app.Port))
}
