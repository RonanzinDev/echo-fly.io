package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ronanzindev/echo-fly.io/src/service"
)

func Home(c echo.Context) error {
	r := service.Render("App.tsx", service.Engine, nil)
	return c.HTML(http.StatusOK, r)
}
