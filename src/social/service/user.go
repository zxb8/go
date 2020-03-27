package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var(
	rxEmail = regexp.MustCompile("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$")
	rxUsername = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_-]{0,17}$")
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidUsername = errors.New("invalid username")
)

type User struct{
	ID int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

func (s *Service) CreateUser(ctx context.Context,email,username string) error{
	email = strings.TrimSpace(email)
	if !rxEmail.MatchString(email){
		return ErrInvalidEmail
	}
	username = strings.TrimSpace(username)
	if !rxUsername.MatchString(username){
		return ErrInvalidUsername
	}

	query := "INSERT INTO users (email,username) VALUES (?,?)"
	_,err :=s.db.ExecContext(ctx,query,email,username)

	if err != nil{
		return fmt.Errorf("could not insert user:%v",err)
	}

	return nil
}