package service

import (
	"fmt"
	"sort"
	"time"

	"github.com/ronanzindev/echo-fly.io/src/database"
	"github.com/ronanzindev/echo-fly.io/src/models"
)

func CreateUser(user models.User) (err error) {
	err = database.DB.Create(&user).Error
	return
}

func GetUsers() (users []models.User, err error) {
	err = database.DB.Find(&users).Error
	return
}

func GetFormatedUsers(users []models.User) []models.User {
	if len(users) == 0 {
		return nil
	}
	u := make([]models.User, len(users))
	t := time.Now()
	for i, user := range users {
		if t.Before(user.Date) {
			user.Validate = false
		}
		ti, _ := time.Parse("2006-01-02", user.FormatedDate)
		y := ti.Year()
		m := ti.Month()
		d := ti.Day()
		user.FormatedDate = fmt.Sprintf("%d/%d/%d", d, m, y)
		u[i] = user
	}
	sort.SliceStable(u, func(i, j int) bool {
		return u[i].Name < u[j].Name
	})

	return u
}

func GetUserById(id string) (user models.User, err error) {
	err = database.DB.Where("id = ?", id).First(&user).Error
	return
}

func UpdateUser(user models.User) (err error) {
	t, err := time.Parse("2006-01-02", user.FormatedDate)
	if err != nil {
		return err
	}
	user.Validate = true
	user.Date = t
	err = database.DB.Save(&user).Error
	return
}
func DeleteUser(id string) (err error) {
	var user models.User
	err = database.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return
	}
	err = database.DB.Delete(&user).Error
	return
}
