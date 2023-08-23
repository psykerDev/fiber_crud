package models

import "time"

type Reply struct {
	CreatedAt time.Time
	ReplyID   uint   `json:"reply_id" gorm:"primaryKey"`
	Comment   string `json:"comment"`
}
