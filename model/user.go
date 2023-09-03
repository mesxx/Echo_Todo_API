package model

import "time"

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Username  string `gorm:"unique" json:"username"`
	Password  string `json:"password"`
	Todo      []Todo `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}
