package props

import (
	"github.com/ronanzindev/echo-fly.io/src/models"
)

type UsersProps struct {
	Users []models.User `json:"users"`
}
type User struct {
	User models.User `json:"user"`
}
