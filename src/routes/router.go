package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ronanzindev/echo-fly.io/src/controller"
)

func NewRouter() (e *echo.Echo) {
	e = echo.New()
	e.Debug = true
	e.Use(middleware.Recover())
	e.Static("/assets", "./frontend")
	e.GET("/add", controller.AddUserPage)
	e.POST("/add", controller.AddUser)
	e.GET("/", controller.Home)
	e.POST("/login", controller.Login)
	e.GET("/users", controller.GetUsers)
	e.GET("/edit/:id", controller.EditUserPage)
	e.PUT("/edit", controller.EditUser)
	e.DELETE("/delete/:id", controller.DeleteUser)
	return
}
