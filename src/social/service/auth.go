package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const(
	TokenLifespan = time.Hour*24*14

	KeyAuthUserId key = "auth_user_id"
)

var (
	ErrUnauthenticated = errors.New("unauthenticated")
)

type key string

type LoginOutput struct{
	Token string
	ExpiresAt time.Time
	AuthUser User
}

func (s *Service) AuthUserID(token string)(int64,error){
	s1,err := s.codec.DecodeToString(token)
	if err != nil{
		return 0,fmt.Errorf("could not decode token:%v",err)
	}
	i,err :=strconv.ParseInt(s1,10,64)

	if err != nil{
		return 0,fmt.Errorf("could not parse auth user id from token:%v",err)
	}
	return i,nil
}


func (s *Service) Login(ctx context.Context,email string)(LoginOutput,error){
	var out LoginOutput

	email = strings.TrimSpace(email)

	if !rxEmail.MatchString(email){
		return out,ErrInvalidEmail
	}

	query := "SELECT id,username from users WHERE email=?"
	err :=s.db.QueryRowContext(ctx,query,email).Scan(&out.AuthUser.ID,&out.AuthUser.Username)
	if err == sql.ErrNoRows{
		return out,ErrUserNotFound
	}
	if err != nil{
		return out,fmt.Errorf("could not query select user:%v",err)
	}
	out.Token,err = s.codec.EncodeToString(strconv.FormatInt(out.AuthUser.ID,10))
	if err != nil{
		return out,fmt.Errorf("could not create token:%v",err)
	}
	out.ExpiresAt = time.Now().Add(TokenLifespan)
	return out,nil
}

func (s *Service) AuthUser(ctx context.Context)(User,error){
	var u User
	uid,ok := ctx.Value(KeyAuthUserId).(int64)
	if !ok{
		return u,ErrUnauthenticated
	}
	query := "SELECT username FROM users WHERE  id = ?"
	err :=s.db.QueryRowContext(ctx,query,uid).Scan(&u.Username)

	if err != nil{
		return u,fmt.Errorf("could not query select auth user:%v",err)
	}
	u.ID = uid
	return u,nil

}