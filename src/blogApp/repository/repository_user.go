package repository

import "blogApp/models"

type UserRepository interface{
	Save(user models.User)(models.User,error)
	FindAll()([]models.User,error)
	FindById(uint32)(models.User,error)
	Update(uint32,user models.User)(int64,error)
	Delete(uint32)(int64,error)
}