package usermodel

import "time"

type User struct {
	ID uint `gorm:"primaryKey"`
	UserId int 
	Username string
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

