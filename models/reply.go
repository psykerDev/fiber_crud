package models

import "time"

type Reply struct {
	ReplyID   uint   `json:"reply_id" gorm:"primaryKey"`
	Post_ID   uint   `json:"post_id"`
	Comment   string `json:"comment"`
	CreatedAt time.Time
}
