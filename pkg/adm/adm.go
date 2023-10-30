package adm

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func AdmPage(c echo.Context) error {

	return c.Render(http.StatusOK, "login.html", nil)
}

func Login(c echo.Context) error {
	// login := os.Getenv("LOGIN")
	// password := os.Getenv("PASSWORD")
	login := c.FormValue("login")
	password := c.FormValue("password")
	if login != os.Getenv("LOGIN") {
		return c.JSON(http.StatusBadRequest, "login errado")
	}
	if password != os.Getenv("PASSWORD") {
		return c.JSON(http.StatusBadRequest, "senha errada")
	}
	// url := url.URL{Path: "http://localhost:8080" + "/html"}
	return c.JSON(http.StatusOK, "sucesso")
}
