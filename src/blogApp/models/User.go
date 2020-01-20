package models

import (
	"blogApp/security"
	"time"
)

type User struct{
	ID uint32 `gorm:"primary_key:auto_increment" json:"id"`
	Nickname string `gorm:"size:20;not null;unique" json:"nickname"`
	Email string `gorm:"size:50;not null;unique" json:"email"`
	Password string `gorm:"size:60;not null" json:"password"`
	CreateAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdateAt time.Time `gorm:"default:current_timestamp()" json:"update_at"`
}

func(u *User) BeforeSave() error{
	hashedPassword,err :=security.Hash(u.Password)
	if err != nil{
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}