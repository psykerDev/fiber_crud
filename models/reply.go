package models

import "time"

type Reply struct {
	CreatedAt time.Time
	ReplyID   uint   `json:"reply_id" gorm:"primaryKey"`
	Post_ID   uint   `json:"post_id"`
	Comment   string `json:"comment"`
}
