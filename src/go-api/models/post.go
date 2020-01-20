package models

import "time"

type Post struct{
	Id uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	ImageUrl string `gorm:"type:varchar(255)" json:"image_url"`
	Subtitle string `gorm:"type:varchar(100)" json:"subtitle"`
	UserId uint32 `gorm:"not null" json:"user_id"`
	User User `json:"user"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	Feedbacks []Feedback `gorm:"ForeignKey:PostId" json:"feedbacks"`
}


func NewPost(post Post) error{
	db := Connect()
	defer db.Close()
	return  db.Create(&post).Error
}