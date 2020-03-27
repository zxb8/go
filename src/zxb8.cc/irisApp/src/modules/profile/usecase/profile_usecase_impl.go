package usecase

import (
	"zxb8/src/modules/profile/model"
	"zxb8/src/modules/profile/repository"
)

type ProfileUsecaseImpl struct{
	profileRepository repository.ProfileRepository
}

func NewProfileUsecase(profileRepository repository.ProfileRepository) *ProfileUsecaseImpl{
	return &ProfileUsecaseImpl{profileRepository}
}

func(profileUsecase *ProfileUsecaseImpl) SaveProfile(profile *model.Profile)(*model.Profile,error){
	err := profileUsecase.profileRepository.Save(profile)
	if err != nil{
		return nil,err
	}
	return profile,nil
}

func(profileUsecase *ProfileUsecaseImpl) UpdateProfile(id string,profile *model.Profile)(*model.Profile,error){
	err := profileUsecase.profileRepository.Update(id,profile)
	if err != nil{
		return nil,err
	}
	return profile,nil
}

func(profileUsecase *ProfileUsecaseImpl) GetById(id string)(*model.Profile,error){
	profile,err :=profileUsecase.profileRepository.FindById(id)
	if err != nil{
		return nil,err
	}
	return profile,nil
}

func(profileUsecase *ProfileUsecaseImpl) GetByEmail(email string)(*model.Profile,error){
	profile,err :=profileUsecase.profileRepository.FindByEmail(email)
	if err != nil{
		return nil,err
	}
	return profile,nil
}

