package app

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
