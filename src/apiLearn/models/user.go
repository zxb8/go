package models

import (
	"apiLearn/utils"
	"fmt"
)

type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreateAt string `json:"createAt"`
}

func NewUser(user User)(bool,error){
	hash,err := utils.Hash(user.Password)

	if err != nil{
		return false,err
	}
	user.Password = fmt.Sprintf("%s",hash)

	con := Connect()
	sql :="insert into users (name,email,password) values(?,?,?)"
	stmt,err := con.Prepare(sql)
	if err != nil{
		return false,nil
	}
	_,err =stmt.Exec(user.Name,user.Email,user.Password)
	if err != nil{
		return false,nil
	}
	defer stmt.Close()
	defer con.Close()
	return true,nil
}

func GetUsers()([]User,error){
	con := Connect()
	sql := "select * from users"
	rs,err :=con.Query(sql)
	if err != nil{
		return nil,err
	}

	var users []User
	for rs.Next(){
		var user User
		err:=rs.Scan(&user.Id,&user.Name,&user.Email,&user.Password,&user.CreateAt)
		if err != nil{
			return nil,err
		}
		users = append(users,user)
	}
	defer rs.Close()
	defer con.Close()
	return users,nil
}

func GetUser(id int)(User ,error){
	con := Connect()
	sql := "select * from users where id=?"
	rs,err :=con.Query(sql,id)
	if err != nil{
		return User{},nil
	}

	var user User

	if rs.Next(){
		err:= rs.Scan(&user.Id,&user.Name,&user.Email,&user.Password,&user.CreateAt)
		if err != nil{
			return User{},err
		}
	}

	defer rs.Close()
	defer con.Close()
	return user,nil
}

func UpdateUser(user User)(int64,error){
	con := Connect()
	sql := "update users set name = ?,email=? where id= ?"
	stmt,err := con.Prepare(sql)
	if err != nil{
		return 0,err
	}

	rs,err :=stmt.Exec(user.Name,user.Email,user.Id)
	if err != nil{
		return 0,err
	}

	rows,err := rs.RowsAffected()
	if err != nil{
		return 0,err
	}

	defer stmt.Close()
	defer con.Close()
	return rows,nil
}

func DeleteUser(id int)(int64,error){
	con := Connect()
	sql := "delete from users where id=?"
	stmt,err := con.Prepare(sql)

	if err != nil{
		return 0,nil
	}
	rs,err :=stmt.Exec(id)
	if err !=nil{
		return 0,err
	}
	rows,err := rs.RowsAffected()
	if err != nil{
		return 0,nil
	}
	defer stmt.Close()
	defer con.Close()
 	return rows,nil
}