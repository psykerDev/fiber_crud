package models

import "time"

type Post struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	CreatedAt time.Time
	Body      string `json:"body"`
	User_ID   uint   `json:"user_id"`
	User      User   `json:"foreignKey:User_ID"`
}
