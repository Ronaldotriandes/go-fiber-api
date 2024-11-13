package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}
