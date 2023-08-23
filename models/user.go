package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Email     string `json:"email" gorm:"unique"`
	Name      string `json:"name"`
}
