package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/ronanzindev/echo-fly.io/src/models"
	"github.com/ronanzindev/echo-fly.io/src/utils"
)

func AdmPage(c echo.Context) error {

	return c.Render(http.StatusOK, "login.html", nil)
}

func Login(c echo.Context) error {
	var body models.LoginAdm
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.H{"error": err.Error()})
	}
	if body.Login != os.Getenv("LOGIN") {
		return c.JSON(http.StatusBadRequest, utils.H{"error": "Login incorreto"})
	}
	if body.Password != os.Getenv("PASSWORD") {
		return c.JSON(http.StatusBadRequest, utils.H{"error": "Senha incorreto"})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": body.Login,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(), // token expires in 30 days
	})
	secretKey := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.H{"error": err.Error()})

	}
	return c.JSON(http.StatusOK, utils.H{"message": "login feito com sucesso", "token": tokenString})
}
