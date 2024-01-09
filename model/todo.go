package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	UserID uint   `json:"userId"`
}
