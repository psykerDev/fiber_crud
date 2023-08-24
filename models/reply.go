package models

import "time"

type Reply struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Post_ID   uint   `json:"post_id"`
	Post      Post   `json:"foreignKey:Post_ID"`
	User_ID   uint   `json:"user_id"`
	User      User   `json:"foreignKey:User_ID"`
	Comment   string `json:"comment"`
	CreatedAt time.Time
}
