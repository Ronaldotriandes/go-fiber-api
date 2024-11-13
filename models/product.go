package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Product_name string `json:"product_name"`
	Description  string `json:"description"`
	User_id      string `json:"user_id"`
}
