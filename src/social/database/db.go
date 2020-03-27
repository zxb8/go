package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

const DRIVER = "mysql"
const DBNAME = "social"
const USER = "root"
const PASS = "123456"


func Connect() *sql.DB{
	URL := fmt.Sprintf("%s:%s@tcp(192.168.100.200:3306)/%s",USER,PASS,DBNAME)
	con,err := sql.Open(DRIVER,URL)

	if err != nil{
		log.Fatal(err)
		return nil
	}
	return con
}