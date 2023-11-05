package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/ronanzindev/echo-fly.io/src/database"
	"github.com/ronanzindev/echo-fly.io/src/models"
)

func SeedUsers() {
	l := byte(97)
	b := []bool{true, false}
	for i := 0; i < 20; i++ {
		letter := int(l) + i
		user := models.User{
			Name:         string((letter)),
			Email:        string(byte(letter)) + "@gmail.com",
			Validate:     b[rand.Intn(2)],
			Date:         time.Date(2023, 10, 29, 0, 0, 0, 0, time.Local),
			FormatedDate: time.Date(2023, 10, 29, 0, 0, 0, 0, time.Local).Format("2006-01-02"),
		}
		user.ID = uuid.New()
		if err := database.DB.Create(&user); err != nil {
			fmt.Println(err)
		}

	}

}
