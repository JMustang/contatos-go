package model

import (
	"crypto/md5"
	"encoding/hex"

	errorhandler "github.com/JMustang/contatos-go/src/config/error_handler"
)

type UserDomain struct {
	// id       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type userDomainInterface interface {
	CreateUser() *errorhandler.HandlerErr
	UpdateUser(string) *errorhandler.HandlerErr
	FindUser(string) (*UserDomain, *errorhandler.HandlerErr)
	DeleteUser(string) *errorhandler.HandlerErr
}
