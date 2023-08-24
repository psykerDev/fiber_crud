package models

import "time"

type Reply struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Post_ID   uint   `json:"post_id"`
	Post      Post   `json:"foreignKey:Post_ID"`
	Comment   string `json:"comment"`
	CreatedAt time.Time
}
