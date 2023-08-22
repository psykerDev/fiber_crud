package models

import "time"

type Post struct {
	PostID    uint   `json:"post_id" gorm:"unique" gorm:"primaryKey"`
	Title     string `json:"title"`
	CreatedAt time.Time
	Body      string `json:"body"`
	User
	Reply
}
