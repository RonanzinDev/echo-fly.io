package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ronanzindev/echo-fly.io/src/models"
	"github.com/ronanzindev/echo-fly.io/src/props"
	"github.com/ronanzindev/echo-fly.io/src/service"
	"github.com/ronanzindev/echo-fly.io/src/utils"
)

func AddUserPage(c echo.Context) error {
	r := service.Render("Add.tsx", service.Engine, nil)
	return c.HTML(http.StatusOK, r)
}

func AddUser(c echo.Context) error {
	var body models.User
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.H{"err": err.Error()})
	}
	if err := service.CreateUser(body); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.H{"err": err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}
func GetUsers(c echo.Context) error {
	users, err := service.GetUsers()
	if len(users) == 0 {
		users := []models.User{
			{Name: "Enzo", Email: "enzo@gmail.com", Validate: true, Date: time.Now(), FormatedDate: "2100/12/12"},
		}
		r := service.Render("Users.tsx", service.Engine, props.UsersProps{Users: users})
		return c.HTML(http.StatusOK, r)
	}
	users = service.GetFormatedUsers(users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.H{"err": err.Error()})
	}
	r := service.Render("Users.tsx", service.Engine, props.UsersProps{Users: users})
	return c.HTML(http.StatusOK, r)
}

func EditUserPage(c echo.Context) error {
	id := c.Param("id")
	user, err := service.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.H{"err": err.Error()})

	}
	r := service.Render("Edit.tsx", service.Engine, props.User{User: user})
	return c.HTML(http.StatusOK, r)
}

func EditUser(c echo.Context) error {
	var body models.User
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.H{"err": err.Error()})
	}
	if err := service.UpdateUser(body); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.H{"err": err.Error()})
	}
	return c.JSON(http.StatusOK, nil)

}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if err := service.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.H{"err": err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}
