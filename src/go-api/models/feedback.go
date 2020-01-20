package models

import "time"

type Feedback struct{
	Id uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Comment string `gorm:"size:255;not null" json:"comment"`
	UserId uint32 `gorm:"not null" json:"user_id"`
	User User `json:"user"`
	PostId uint32 `gorm:"not null" json:"post_id"`
	Post Post `json;"post"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func NewFeedback(feedback Feedback) error{
	db := Connect()
	defer db.Close()
	return db.Create(feedback).Error
}