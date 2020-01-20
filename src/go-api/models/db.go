package models

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

const(
	USER = "root"
	PASS = "123456"
	HOST = "192.168.100.200"
	PORT = 3306
	DBNAME = "gorm"
)

func Connect() *gorm.DB{
	URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",USER,PASS,HOST,PORT,DBNAME)
	db,err :=gorm.Open("mysql",URL)
	if err != nil{
		log.Fatal(err)
	}
	return db
}



