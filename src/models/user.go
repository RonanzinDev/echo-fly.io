package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Validate     bool      `json:"validate"`
	Date         time.Time `json:"date"`
	FormatedDate string    `json:"formated_date"`
	Schema
}

type Schema struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid()`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
