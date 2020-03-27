package repository

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"zxb8/src/modules/profile/model"
)

type ProfileRepositoryMongo struct{
	db *mgo.Database
	collection string
}

func NewProfileRepositoryMongo(db *mgo.Database,collection string) *ProfileRepositoryMongo{
	return &ProfileRepositoryMongo{
		db:db,
		collection:collection,
	}
}

func(r *ProfileRepositoryMongo) Save(profile *model.Profile) error{
	err :=r.db.C(r.collection).Insert(profile)
	return err
}

func(r *ProfileRepositoryMongo) Update(id string,profile *model.Profile) error{
	profile.UpdateAt = time.Now()
	err := r.db.C(r.collection).Update(bson.M{"id":id},profile)
	return err
}

func(r *ProfileRepositoryMongo) Delete(id string) error{
	err := r.db.C(r.collection).Remove(bson.M{"id":id})
	return err
}

func(r *ProfileRepositoryMongo) FindById(id string) (*model.Profile,error){
	var profile model.Profile
	err := r.db.C(r.collection).Find(bson.M{"id":id}).One(&profile)
	if err != nil{
		return nil,err
	}
	return &profile,nil
}

func(r *ProfileRepositoryMongo) FindByEmail(email string)(*model.Profile,error){
	var profile model.Profile
	err := r.db.C(r.collection).Find(bson.M{"email":email}).One(&profile)
	if err != nil{
		return nil,err
	}
	return &profile,nil
}

func(r *ProfileRepositoryMongo) FindAll()(model.Profiles,error){
	var profiles model.Profiles
	err := r.db.C(r.collection).Find(bson.M{}).All(&profiles)
	if err != nil{
		return nil,err
	}
	return profiles,nil
}