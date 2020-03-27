package usecase

import "zxb8/src/modules/profile/model"

//Profile usecase
type ProfileUsecase interface{
	SaveProfile(profile *model.Profile)(*model.Profile,error)
	UpdateProfile(string,profile *model.Profile)(*model.Profile,error)
	GetById(string)(*model.Profile,error)
	GetByEmail(string)(*model.Profile,error)
}
