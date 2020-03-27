package config

import (
	"gopkg.in/mgo.v2"
)

func GetMongoDB()(*mgo.Database,error){
	host := "192.168.100.200"
	//host := os.Getenv("MONGO_HOST")
	dbName := "iris_demo"
	//dbName := os.Getenv("MONGO_DB_NAME")

	session,err := mgo.Dial(host)

	if err != nil{
		return nil,err
	}

	db := session.DB(dbName)
	return db,nil
}