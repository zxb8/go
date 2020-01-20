package crud

import (
	"blogApp/models"
	"blogApp/utils/channels"
	"github.com/jinzhu/gorm"
)

type repositoryUserCRUD struct{
	db *gorm.DB
}

func NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUserCRUD{
	return &repositoryUserCRUD{db}
}

func (r *repositoryUserCRUD) Save(user models.User)(models.User,error){
	var err error
	done := make(chan bool)

	go func(ch chan <- bool){
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil{
			ch <- false
			return
		}
		ch <-true
	}(done)

	if channels.OK(done){
		return user,nil
	}

	return models.User{},err
}