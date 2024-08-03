package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title    string     `json:"title"`
	ImageURL *string    `json:"url"`
	Genre    string     `json:"genre"`
	Release  *time.Time `json:"link"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
}
