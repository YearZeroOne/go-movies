package models

import (
	"time"

	"gorm.io/gorm"
)


type Movie struct {
	gorm.Model
	ID       uint       `json:"id" gorm:"primaryKey"`
	Title    string     `json:"title"`
	ImageURL *string    `json:"url"`
	Genre    string     `json:"genre"`
	Release  *time.Time `json:"link"`
}

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
}
