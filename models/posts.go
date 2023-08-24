package models

import "time"

type Post struct {
	Post_ID   uint   `json:"post_id" gorm:"primaryKey"`
	Title     string `json:"title"`
	CreatedAt time.Time
	Body      string `json:"body"`
	User_ID   uint   `json:"user_id"`
	//Reply_ID  uint   `json:"reply_id"`
	// User      User   `gorm:"foreignKey:ID;references:User_ID"`
	// Reply     Reply  `gorm:"foreignKey:ReplyID;references:Reply_ID"`
}
