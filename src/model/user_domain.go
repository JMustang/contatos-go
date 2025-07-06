package model

import errorhandler "github.com/JMustang/contatos-go/src/config/error_handler"

type UserDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

type userDomainInterface interface {
	CreateUser(UserDomain) *errorhandler.HandlerErr
	UpdateUser(string, UserDomain) *errorhandler.HandlerErr
	FindUser(string) (*UserDomain, *errorhandler.HandlerErr)
	DeleteUser(string) *errorhandler.HandlerErr
}
