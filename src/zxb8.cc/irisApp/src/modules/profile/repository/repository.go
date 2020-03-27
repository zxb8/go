package repository

import "zxb8/src/modules/profile/model"

type ProfileRepository interface{
	Save(*model.Profile) error
	Update(string,*model.Profile) error
	Delete(string) error
	FindById(string) (*model.Profile,error)
	FindByEmail(string)(*model.Profile,error)
	FindAll()(model.Profiles,error)
}