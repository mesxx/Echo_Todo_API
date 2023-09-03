package model

import "time"

type Todo struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	UserID    uint   `json:"userId"`
	CreatedAt time.Time
}
