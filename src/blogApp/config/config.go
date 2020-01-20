package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var (
	PORT = 0
	DBDRIVER = ""
	DBURL = ""
)

func Load(){
	var err error
	err  = godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	PORT,err = strconv.Atoi(os.Getenv("PORT"))
	if err !=nil{
		PORT = 9000
	}

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@(192.168.100.200)/%s?charset=utf8&parseTime=True&loc=Local",os.Getenv("DB_USER"),os.Getenv("DB_PASS"),os.Getenv("DB_NAME"))
}