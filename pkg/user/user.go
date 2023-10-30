package user

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Validate time.Time `json:"validate"`
}

var Users = []User{
	{
		Name:     "Jonh Doe",
		Email:    "jonhdoe@gmail.com",
		Validate: time.Date(2023, 10, 29, 0, 0, 0, 0, time.Local),
	},
	{
		Name:     "Jonh Doe",
		Email:    "jonhdoe@gmail.com",
		Validate: time.Date(2023, 10, 29, 0, 0, 0, 0, time.Local),
	},
	{
		Name:     "Jonh Doe",
		Email:    "jonhdoe@gmail.com",
		Validate: time.Date(2023, 10, 29, 0, 0, 0, 0, time.Local),
	},
	{
		Name:     "Jonh Doe",
		Email:    "jonhdoe@gmail.com",
		Validate: time.Date(2023, 10, 29, 0, 0, 0, 0, time.Local),
	},
}

func GetUsers(c echo.Context) error {
	type Data struct {
		Users []User
	}
	return c.Render(http.StatusOK, "users.html", Data{Users})
}
