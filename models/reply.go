package models

import "time"

type Reply struct {
	CreatedAt time.Time
	ReplyID   uint   `json:"reyly_id" gorm:"unique" gorm:"primaryKey"`
	Comment   string `json:"comment"`
}
